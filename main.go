package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	ACCUMULATOR        = "accumulator"
	IMMEDIATE          = "immediate"
	ZEROPAGE           = "zeropage"
	ZEROPAGEX          = "zeropagex"
	ZEROPAGEY          = "zeropagey"
	ABSOLUTE           = "absolute"
	ABSOLUTEX          = "absolutex"
	ABSOLUTEY          = "absolutey"
	INDIRECT           = "indirect"
	INDIRECTX          = "indirectx"
	INDIRECTY          = "indirecty"
	SPBaseAddress uint = 0x0100
)

var (
	allsuitea    = flag.Bool("allsuitea", false, "AllSuiteA ROM")
	klausd       = flag.Bool("klausd", false, "Klaus Dormann's 6502 functional test ROM")
	plus4        = flag.Bool("plus4", false, "Plus/4 ROMs")
	disassemble  = flag.Bool("dis", false, "Disassembler mode")
	stateMonitor = flag.Bool("state", false, "State monitor mode")

	disassembledInstruction string

	instructionCounter = 0

	// Fixed Addresses
	BASICROM      = make([]byte, 16384)
	KERNALROM     = make([]byte, 16384)
	THREEPLUS1ROM = make([]byte, 16384)
	AllSuiteAROM  = make([]byte, 16384)
	KlausDTestROM = make([]byte, 65536)

	basicROMAddress      = 0x8000
	kernalROMAddress     = 0xC000
	threePlus1ROMAddress = 0xD000
	AllSuiteAROMAddress  = 0x4000

	resetVectorAddress     = 0xFFFC
	interruptVectorAddress = 0xFFFE

	// CPURegisters and RAM
	A      byte        = 0x0    // Accumulator
	X      byte        = 0x0    // X register
	Y      byte        = 0x0    // Y register		(76543210) SR Bit 5 is always set
	SR     byte                 // Status Register	(NVEBDIZC)
	SP     uint        = 0x01FF // Stack Pointer
	PC     int                  // Program Counter
	memory [65536]byte          // Memory
)

func reset() {
	SP = SPBaseAddress
	// Set SR to 0b00110100
	SR = 0b00110110
	if *klausd {
		// Set PC to $0400
		PC = 0x0400
	} else {
		// Set PC to value stored at reset vector address
		PC = int(memory[resetVectorAddress]) + int(memory[resetVectorAddress+1])*256
	}
}

func loadROMs() {
	if *plus4 {
		// Copy the BASIC ROM into memory
		file, _ := os.Open("roms/plus4/basic-318006-01.bin")
		_, _ = io.ReadFull(file, BASICROM)
		fmt.Printf("Copying BASIC ROM into memory at $%04X to $%04X\n\n", basicROMAddress, basicROMAddress+len(BASICROM))
		copy(memory[basicROMAddress:], BASICROM)

		// Copy the Kernal ROM into memory
		file, _ = os.Open("roms/plus4/kernal-318005-05.bin")
		_, _ = io.ReadFull(file, KERNALROM)
		fmt.Printf("Copying KERNALROM into memory at $%04X to $%04X\n\n", kernalROMAddress, kernalROMAddress+len(KERNALROM))
		copy(memory[kernalROMAddress:], KERNALROM)
	}

	if *allsuitea {
		// Copy AllSuiteA ROM into memory
		file, _ := os.Open("roms/AllSuiteA.bin")
		_, _ = io.ReadFull(file, AllSuiteAROM)
		fmt.Printf("Copying AllSuiteA ROM into memory at $%04X to $%04X\n\n", AllSuiteAROMAddress, AllSuiteAROMAddress+len(AllSuiteAROM))
		copy(memory[AllSuiteAROMAddress:], AllSuiteAROM)
		// Set the interrupt vector addresses manually
		//memory[0xFFFE] = AllSuiteAROM[len(AllSuiteAROM)-2]
		//memory[0xFFFF] = AllSuiteAROM[len(AllSuiteAROM)-1]
		memory[0xFFFE] = 0x00 // Low byte of 0x4000
		memory[0xFFFF] = 0x40 // High byte of 0x4000
	}

	if *klausd {
		//Copy roms/6502_functional_test.bin into memory at $0000
		file, _ := os.Open("roms/6502_functional_test.bin")
		_, _ = io.ReadFull(file, KlausDTestROM)
		copy(memory[0x0000:], KlausDTestROM)
		fmt.Printf("Copying Klaus Dormann's 6502 functional test ROM into memory at $%04X to $%04X\n\n", 0x0000, 0x0000+len(memory))
		// Set the interrupt vector addresses manually
		memory[0xFFFE] = 0x00 // Low byte of 0x4000
		memory[0xFFFF] = 0x40 // High byte of 0x4000
	}
}

func main() {
	fmt.Printf("Six5go2 v2.0 - 6502 Emulator and Disassembler in Golang (c) 2022 Zayn Otley\n\n")
	flag.Parse()

	/*if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}*/

	fmt.Printf("Size of addressable memory is %v ($%04X) bytes\n\n", len(memory), len(memory))

	// Start emulation
	loadROMs()
	reset()
	execute()
}

func printMemoryAroundSP() {
	fmt.Printf("Memory around SP:\n")
	startAddress := uint16(SPBaseAddress) + uint16(SP) - 5
	for i := 0; i < 11; i++ {
		fmt.Printf("0x%04X: 0x%02X\n", startAddress+uint16(i), memory[startAddress+uint16(i)])
	}
}

func disassembleOpcode() {
	if *disassemble {
		fmt.Printf("%s\n", disassembledInstruction)
	}
}
func petsciiToAscii(petscii uint8) uint8 {
	// Convert PETSCII to ASCII
	if petscii >= 65 && petscii <= 90 { // PETSCII uppercase
		return petscii + 32 // To ASCII lowercase
	} else if petscii >= 193 && petscii <= 218 { // PETSCII lowercase
		return petscii - 96 // To ASCII uppercase
	} else if petscii >= 219 && petscii <= 250 { // PETSCII graphics
		return petscii - 128 // To ASCII graphics
	}
	return petscii
}

func kernalRoutines() {
	// CHROUT routine is at $FFD2
	switch PC {
	case 0xFFD2:
		// This is a CHROUT call
		ch := petsciiToAscii(A) // Convert PETSCII to ASCII

		// Handle control characters
		switch ch {
		case 13: // Carriage return
			fmt.Print("\r")
		case 10: // Line feed
			fmt.Print("\n")
		case 8: // Backspace
			fmt.Print("\b")
		case 9: // Tab
			fmt.Print("\t")
		default: // Not a control character
			if ch >= 32 && ch <= 126 { // Check if the ASCII value is a printable character
				fmt.Printf("%c", ch)
			} else {
				fmt.Printf("Invalid ASCII value: %d\n", ch)
			}
		}
	case 0xFFF6:
		// This is a RESET call
		reset()
	}
	// print "kernal rom call address"
	//fmt.Printf("\n\u001B[32;5mKernal ROM call address: $%04X\u001B[0m\n", PC)
}

func opcode() byte {
	return memory[PC]
}
func operand1() byte {
	return memory[PC+1]
}
func operand2() byte {
	return memory[PC+2]
}
func incCount(amount int) {
	if *stateMonitor {
		if disassembledInstruction != "BRK" {
			printMachineState()
		}
	}
	PC += amount
	if PC > 0xFFFF {
		PC = 0x0000 + (PC & 0xFFFF)
	}

	// If amount is 0, then we are in a branch instruction and we don't want to increment the instruction counter
	if amount != 0 {
		instructionCounter++
	}
}

func printMachineState() {
	// Imitate Virtual 6502 and print PC, opcode, operand1 if two byte instruction, operand2 if three byte instruction, disassembled instruction and any operands with correct addressing mode, "|",accumulator hex value, X register hex value, Y register hex value, SP as hex value,"|", SP as binary digits
	fmt.Printf("%04X ", PC)

	// If opcode() is a 1 byte instruction, print opcode
	if opcode() == 0x00 || opcode() == 0x08 || opcode() == 0x10 || opcode() == 0x18 || opcode() == 0x20 || opcode() == 0x28 || opcode() == 0x30 || opcode() == 0x38 || opcode() == 0x40 || opcode() == 0x48 || opcode() == 0x50 || opcode() == 0x58 || opcode() == 0x68 || opcode() == 0x70 || opcode() == 0x78 || opcode() == 0x88 || opcode() == 0x8A || opcode() == 0x98 || opcode() == 0x9A || opcode() == 0xA8 || opcode() == 0xAA || opcode() == 0xB8 || opcode() == 0xBA || opcode() == 0xC8 || opcode() == 0xCA || opcode() == 0xD8 || opcode() == 0xDA || opcode() == 0xE8 || opcode() == 0xEA || opcode() == 0xF8 || opcode() == 0xFA || opcode() == 0x2A || opcode() == 0x6A || opcode() == 0x60 {
		fmt.Printf("%02X ", opcode())
	}

	// If opcode() is a 2 byte instruction, print opcode and operand1
	// 		fmt.Printf("%02X %02X ", opcode(), operand1())
	// The 0x hex opcodes for the 2 byte instructions on the 6502 are
	if opcode() == 0x69 || opcode() == 0x29 || opcode() == 0xC9 || opcode() == 0xE0 || opcode() == 0xC0 || opcode() == 0x49 || opcode() == 0xA9 || opcode() == 0xA2 || opcode() == 0xA0 || opcode() == 0x09 || opcode() == 0xE9 || opcode() == 0x65 || opcode() == 0x25 || opcode() == 0x06 || opcode() == 0x24 || opcode() == 0xC5 || opcode() == 0xE4 || opcode() == 0xC4 || opcode() == 0xC6 || opcode() == 0x45 || opcode() == 0xE6 || opcode() == 0xA5 || opcode() == 0xA6 || opcode() == 0xA4 || opcode() == 0x46 || opcode() == 0x05 || opcode() == 0x26 || opcode() == 0x66 || opcode() == 0xE5 || opcode() == 0x85 || opcode() == 0x86 || opcode() == 0x84 || opcode() == 0x90 || opcode() == 0xB0 || opcode() == 0xF0 || opcode() == 0x30 || opcode() == 0xD0 || opcode() == 0x10 || opcode() == 0x50 || opcode() == 0x70 {
		fmt.Printf("%02X %02X ", opcode(), operand1())
	}

	// If opcode() is a 3 byte instruction, print opcode, operand1 and operand2
	// 			fmt.Printf("%02X %02X %02X ", opcode(), operand1(), operand2())
	if opcode() == 0x6D || opcode() == 0x2D || opcode() == 0x0E || opcode() == 0x2C || opcode() == 0xCD || opcode() == 0xEC || opcode() == 0xCC || opcode() == 0xCE || opcode() == 0x4D || opcode() == 0xEE || opcode() == 0x4C || opcode() == 0x20 || opcode() == 0xAD || opcode() == 0xAC || opcode() == 0xAE || opcode() == 0x4E || opcode() == 0x0D || opcode() == 0x2E || opcode() == 0x6E || opcode() == 0xED || opcode() == 0x8D || opcode() == 0x8E || opcode() == 0x8C || opcode() == 0x7D || opcode() == 0x79 || opcode() == 0x3D || opcode() == 0x39 || opcode() == 0x1E || opcode() == 0xDD || opcode() == 0xD9 || opcode() == 0xDE || opcode() == 0x5D || opcode() == 0x59 || opcode() == 0xFE || opcode() == 0xBD || opcode() == 0xB9 || opcode() == 0xBC || opcode() == 0xBE || opcode() == 0x5E || opcode() == 0x1D || opcode() == 0x19 || opcode() == 0x3E || opcode() == 0x7E || opcode() == 0xFD || opcode() == 0xF9 || opcode() == 0x9D || opcode() == 0x95 || opcode() == 0x99 || opcode() == 0xB5 || opcode() == 0x91 || opcode() == 0xB1 || opcode() == 0x81 || opcode() == 0xA1 || opcode() == 0x94 || opcode() == 0x96 || opcode() == 0xB4 || opcode() == 0xB6 || opcode() == 0x35 || opcode() == 0x15 || opcode() == 0x55 || opcode() == 0x21 || opcode() == 0x01 || opcode() == 0x41 || opcode() == 0x31 || opcode() == 0x11 || opcode() == 0x51 || opcode() == 0xF6 || opcode() == 0xD6 || opcode() == 0x4A || opcode() == 0x0A || opcode() == 0x16 || opcode() == 0x56 || opcode() == 0x36 || opcode() == 0x76 || opcode() == 0x75 || opcode() == 0xF5 || opcode() == 0xD5 || opcode() == 0xC1 || opcode() == 0xD1 || opcode() == 0x61 || opcode() == 0xE1 || opcode() == 0x71 || opcode() == 0xF1 {
		fmt.Printf("%02X %02X %02X ", opcode(), operand1(), operand2())
	}

	// Print disassembled instruction
	fmt.Printf("\t%s", disassembledInstruction)
	// Print A,X,Y,SP as hex values
	fmt.Printf("\t|%02X %02X %02X %02X| ", A, X, Y, SP)

	// SR flags are     N,V,-,B,D,I,Z,C
	// SR flag bits are 7,6,5,4,3,2,1,0

	//getSRBit flags for 7,6,3,2,1,0 and print as binary digits
	//getSRBit flags for N,V,D,I,Z and C and print as binary digits
	fmt.Printf("%b", getSRBit(7))
	fmt.Printf("%b", getSRBit(6))
	fmt.Printf("%b", getSRBit(3))
	fmt.Printf("%b", getSRBit(2))
	fmt.Printf("%b", getSRBit(1))
	fmt.Printf("%b |\n", getSRBit(0))

	// Print full SR as binary digits with zero padding
	//fmt.Printf("%08b |\n", SR)

	// Wait for keypress
	//fmt.Scanln()
}
func getSRBit(x byte) byte {
	return (SR >> x) & 1
}
func setSRBitOn(x byte) {
	SR |= 1 << x
}
func setSRBitOff(x byte) {
	SR &= ^(1 << x)
}
func getABit(x byte) byte {
	return (A >> x) & 1
}
func getXBit(x byte) byte {
	return (X >> x) & 1
}
func getYBit(x byte) byte {
	return (Y >> x) & 1
}
func setNegativeFlag() {
	setSRBitOn(7)
}
func unsetNegativeFlag() {
	setSRBitOff(7)
}
func setOverflowFlag() {
	setSRBitOn(6)
}
func unsetOverflowFlag() {
	setSRBitOff(6)
}
func setBreakFlag() {
	setSRBitOn(4)
}
func unsetBreakFlag() {
	setSRBitOff(4)
}
func setDecimalFlag() {
	setSRBitOn(3)
}
func unsetDecimalFlag() {
	setSRBitOff(3)
}
func setInterruptFlag() {
	setSRBitOn(2)
}
func unsetInterruptFlag() {
	setSRBitOff(2)
}
func setZeroFlag() {
	setSRBitOn(1)
}
func unsetZeroFlag() {
	setSRBitOff(1)
}
func setCarryFlag() {
	setSRBitOn(0)
}
func unsetCarryFlag() {
	setSRBitOff(0)
}
func readBit(bit byte, value byte) int {
	// Read bit from value and return it
	return int((value >> bit) & 1)
}
func decSP() {
	if SP == 0x00 {
		// Wrap around to top of stack at 0xFF
		SP = 0xFF
	} else {
		// Decrement SP normally
		SP--
	}
}

func incSP() {
	if SP == 0xFF {
		// Wrap around to bottom of stack at 0x00
		SP = 0x00
	} else {
		// Increment SP normally
		SP++
	}
}

// 6502 mnemonics with multiple addressing modes
func LDA(addressingMode string) {
	setFlags := func() {
		// If A is zero, set the SR Zero flag to 1 else set SR Zero flag to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If bit 7 of accumulator is 1, set the SR negative flag to 1 else set the SR negative flag to 0
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE: // Immediate
		A = operand1()
		setFlags()
		incCount(2)
	case ZEROPAGE: // Zero Page
		// Get address
		address := operand1()
		// Get value from memory at address
		value := memory[address]
		// Set accumulator to value
		A = value
		setFlags()
		incCount(2)
	case ZEROPAGEX: // Zero Page, X
		// Get address
		address := operand1() + X
		value := memory[address]
		// Set accumulator to value
		A = value
		setFlags()
		incCount(2)
	case ABSOLUTE: // Absolute
		// Get 16 bit address from operand 1 and operand 2
		address := int(operand2())<<8 | int(operand1())
		value := memory[address]
		// Set accumulator to value
		A = value
		setFlags()
		incCount(3)
	case ABSOLUTEX: // Absolute, X
		// Get the 16bit X indexed absolute memory address
		address := (int(operand2())<<8 | int(operand1())) + int(X)
		value := memory[address]
		// Set accumulator to value
		A = value
		setFlags()
		incCount(3)
	case ABSOLUTEY: // Absolute, Y
		// Get 16 bit address from operand 1 and operand 2
		address := (int(operand2())<<8 | int(operand1())) + int(Y)
		value := memory[address]
		// Set accumulator to value
		A = value
		setFlags()
		incCount(3)
	case INDIRECTX: // Indirect, X
		// Get the 16bit X indexed zero page indirect address
		indirectAddress := uint16(int(operand1()) + int(X)&0xFF)
		// Get the value at the indirect address
		indirectValue := memory[indirectAddress]
		// Get the value at the indirect address + 1
		indirectValue2 := memory[indirectAddress+1]
		// Corrected line: Combine the two values to get the address
		indirectAddress = uint16(int(indirectValue2)<<8 + int(indirectValue))
		// Get the value at the address
		value := memory[indirectAddress]
		// Set the accumulator to the value
		A = value
		setFlags()
		incCount(2)
	case INDIRECTY: // Indirect, Y
		// Get address
		zeroPageAddress := operand1()
		address := (uint16(memory[zeroPageAddress+1])<<8 | uint16(memory[zeroPageAddress])) + uint16(Y)
		// Get the value at the address
		value := memory[address]
		// Set the accumulator to the value
		A = value
		setFlags()
		incCount(2)
	}
}
func LDX(addressingMode string) {
	setFlags := func() {
		// If bit 7 of X is set, set the SR negative flag else reset it to 0
		if getXBit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If X is zero, set the SR zero flag else reset it to 0
		if X == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE: // Immediate
		// Load the value of the operand1() into the X register.
		X = operand1()
		setFlags()
		incCount(2)
	case ZEROPAGE: // Zero Page
		// Get address
		address := operand1()
		value := memory[address]
		// Load the value at the address into X
		X = value
		setFlags()
		incCount(2)
	case ZEROPAGEY: // Zero Page, Y
		// Get Y indexed Zero Page address
		address := operand1() + Y
		value := memory[address]
		// Load the X register with the Y indexed value in the operand
		X = value
		setFlags()
		incCount(2)
	case ABSOLUTE: // Absolute
		// Get 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		value := memory[address]
		// Update X with the value stored at the address in the operands
		X = value
		setFlags()
		incCount(3)
	case ABSOLUTEY: // Absolute, Y
		// Get 16 bit Y indexed address from operands
		address := int(operand2())<<8 | int(operand1()) + int(Y)
		value := memory[address]
		X = value
		setFlags()
		incCount(3)
	}
}
func LDY(addressingMode string) {
	setFlags := func() {
		// If bit 7 of Y is set, set the SR negative flag else reset it to 0
		if getYBit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If Y is zero, set the SR zero flag else reset it to 0
		if Y == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE: // Immediate
		// Load the value of the operand1() into the Y register.
		Y = operand1()
		setFlags()
		incCount(2)
	case ZEROPAGE: // Zero Page
		// Get address
		address := operand1()
		value := memory[address]
		// Load the value at the address into Y
		Y = value
		setFlags()
		incCount(2)
	case ZEROPAGEX: // Zero Page, X
		// Get the X indexed address
		address := operand1() + X
		value := memory[address]
		// Load the Y register with the X indexed value in the operand
		Y = value
		setFlags()
		incCount(2)
	case ABSOLUTE: // Absolute
		// Get 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		value := memory[address]
		// Update Y with the value stored at the address in the operands
		Y = value
		setFlags()
		incCount(3)
	case ABSOLUTEX: // Absolute, X
		// Get the 16bit X indexed absolute memory address
		address := (int(operand2())<<8 | int(operand1())) + int(X)
		value := memory[address]
		// Update Y with the value stored at the address
		Y = value
		setFlags()
		incCount(3)
	}
}
func STA(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE:
		address := operand1()
		memory[address] = A
		incCount(2)
	case ZEROPAGEX:
		address := (operand1() + X) & 0xFF // Ensure wraparound in Zero Page
		memory[address] = A
		incCount(2)
	case ABSOLUTE:
		address := uint16(operand2())<<8 | uint16(operand1())
		memory[address] = A
		incCount(3)
	case ABSOLUTEX:
		address := (uint16(operand2())<<8 | uint16(operand1())) + uint16(X)
		memory[address] = A
		incCount(3)
	case ABSOLUTEY:
		address := (uint16(operand2())<<8 | uint16(operand1())) + uint16(Y)
		memory[address] = A
		incCount(3)
	case INDIRECTX:
		zeroPageAddress := (operand1() + X) & 0xFF
		address := uint16(memory[zeroPageAddress]) | uint16(memory[(zeroPageAddress+1)&0xFF])<<8
		//print out the address and the value of A
		memory[address] = A
		incCount(2)
	case INDIRECTY:
		zeroPageAddress := operand1()
		address := uint16(memory[zeroPageAddress]) | uint16(memory[(zeroPageAddress+1)&0xFF])<<8
		finalAddress := (address + uint16(Y)) & 0xFFFF
		memory[finalAddress] = A
		incCount(2)
	}
}

func STX(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE:
		address := operand1()
		memory[address] = X
		incCount(2)
	case ZEROPAGEY: // Note the change from ZEROPAGEX to ZEROPAGEY
		address := (operand1() + Y) & 0xFF // We add Y here, not X
		memory[address] = X
		incCount(2)
	case ABSOLUTE:
		address := uint16(operand2())<<8 | uint16(operand1())
		memory[address] = X
		incCount(3)
	}
}

func STY(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE:
		address := operand1()
		memory[address] = Y
		incCount(2)
	case ZEROPAGEX:
		address := (operand1() + X) & 0xFF
		memory[address] = Y
		incCount(2)
	case ABSOLUTE:
		address := uint16(operand2())<<8 | uint16(operand1())
		memory[address] = Y
		incCount(3)
	}
}

func CMP(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// Subtract the value from the accumulator
		result = A - value
		//fmt.Printf("A: %X, value: %X, result: %X\n", A, value, result)
		// If the result is 0, set the zero flag
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		if readBit(7, result) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If the value is less than or equal to the accumulator, set the carry flag, else reset it
		if value <= A {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		if addressingMode == IMMEDIATE || addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX || addressingMode == INDIRECTX || addressingMode == INDIRECTY {
			incCount(2)
		} else {
			incCount(3)
		}
	}
	switch addressingMode {
	case IMMEDIATE: // Immediate
		// Get value from operand1()
		value = operand1()
		setFlags()
	case ZEROPAGE: // Zero Page
		// Get address
		address := operand1()
		// Subtract the operand from the accumulator
		value = memory[address]
		setFlags()
	case ZEROPAGEX: // Zero Page, X
		// Get address
		address := operand1() + X
		// Get value at address
		value = memory[address]
		setFlags()
	case ABSOLUTE: // Absolute
		// Get 16bit absolute address
		address := int(operand2())<<8 | int(operand1())
		// Get the value at the address
		value = memory[address]
		setFlags()
	case ABSOLUTEX: // Absolute, X
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Get value at address
		value = memory[address]
		setFlags()
	case ABSOLUTEY: // Absolute, Y
		// Get address
		address := int(operand2())<<8 | int(operand1()) + int(Y)
		// Get the value at the address
		value = memory[address]
		setFlags()
	case INDIRECTX: // Indirect, X
		// Get the address of the operand
		address := int(operand1()) + int(X)
		// Get the value of the operand
		value = memory[address]
		setFlags()
	case INDIRECTY: // Indirect, Y
		// Get address from operand1() and add Y to it
		address := memory[operand1()] + Y
		// Get value at address
		value = memory[address]
		setFlags()
	}
}

func JMP(addressingMode string) {
	incCount(0)
	switch addressingMode {
	case ABSOLUTE:
		// Get the 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		// Set the program counter to the absolute address
		PC = int(address)
	case INDIRECT:
		// Get the 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		// Handle 6502 page boundary bug
		loByteAddress := address
		hiByteAddress := (address & 0xFF00) | ((address + 1) & 0xFF) // Ensure it wraps within the page
		indirectAddress := uint16(memory[hiByteAddress])<<8 | uint16(memory[loByteAddress])
		// Set the program counter to the indirect address
		PC = int(indirectAddress)
	}
}

func AND(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// If the result is 0, set the zero flag
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		if readBit(7, result) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = operand1()
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(2)
	case ZEROPAGE:
		// Get the address from the operand
		address := operand1()
		// Get the value at the address
		value = memory[address]
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(2)
	case ZEROPAGEX:
		// Get address
		address := operand1() + X
		// Get value at address
		value = memory[address]
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = memory[address]
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(3)
	case ABSOLUTEX:
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Get value at address
		value = memory[address]
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(3)
	case ABSOLUTEY:
		// Get the address
		address := int(operand2())<<8 | int(operand1()) + int(Y)
		// Get the value at the address
		value = memory[address]
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(3)
	case INDIRECTX:
		// Get the address
		indirectAddress := int(operand1()) + int(X)
		address := int(memory[indirectAddress]) + int(memory[indirectAddress+1])<<8
		// Get the value from the address
		value = memory[address]
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(2)
	case INDIRECTY:
		// Get the 16bit address
		address := uint16(int(operand1()))
		// Get the indirect address
		indirectAddress1 := memory[address]
		indirectAddress2 := memory[address+1]
		indirectAddress := uint16(int(indirectAddress1)+int(indirectAddress2)<<8) + uint16(Y)
		// Get the value at the address
		value = memory[indirectAddress]
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(2)
	}
}
func EOR(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// If the result is 0, set the zero flag
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		if readBit(7, result) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = operand1()
		// XOR the value with the accumulator
		result = A ^ value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(2)
	case ZEROPAGE:
		// Get the address from the operand
		address := operand1()
		// Get the value at the address
		value = memory[address]
		// XOR the value with the accumulator
		result = A ^ value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(2)
	case ZEROPAGEX:
		// Get address
		address := operand1() + X
		// Get value at address
		value = memory[address]
		// XOR the value with the accumulator
		result = A ^ value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = memory[address]
		// XOR the value with the accumulator
		result = A ^ value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(3)
	case ABSOLUTEX:
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Get value at address
		value = memory[address]
		// XOR the value with the accumulator
		result = A ^ value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(3)
	case ABSOLUTEY:
		// Get the address
		address := int(operand2())<<8 | int(operand1()) + int(Y)
		// Get the value at the address
		value = memory[address]
		// XOR the value with the accumulator
		result = A ^ value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(3)
	case INDIRECTX:
		// Get the address
		indirectAddress := int(operand1()) + int(X)
		address := int(memory[indirectAddress]) + int(memory[indirectAddress+1])<<8
		// Get the value from the address
		value = memory[address]
		// XOR the value with the accumulator
		result = A ^ value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(2)
	case INDIRECTY:
		// Get the 16bit address
		address := uint16(int(operand1()))
		// Get the indirect address
		indirectAddress1 := memory[address]
		indirectAddress2 := memory[address+1]
		indirectAddress := uint16(int(indirectAddress1)+int(indirectAddress2)<<8) + uint16(Y)
		// Get the value at the address
		value = memory[indirectAddress]
		// XOR the value with the accumulator
		result = A ^ value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(2)
	}
}
func ORA(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// If the result is 0, set the zero flag
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		if readBit(7, result) == 1 {
			setNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = operand1()
		// OR the value with the accumulator
		result = A | value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(2)
	case ZEROPAGE:
		// Get the address from the operand
		address := operand1()
		// Get the value at the address
		value = memory[address]
		// OR the value with the accumulator
		result = A | value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(2)
	case ZEROPAGEX:
		// Get address
		address := operand1() + X
		// Get value at address
		value = memory[address]
		// OR the value with the accumulator
		result = A | value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = memory[address]
		// OR the value with the accumulator
		result = A | value
		// Set the accumulator to the result
		A = result
		setFlags()
		incCount(3)
	case ABSOLUTEX:
		address := (uint16(operand1()) + uint16(X)) | uint16(operand2())<<8
		value = memory[address]
		A |= value
		setFlags()
		incCount(3)
	case ABSOLUTEY:
		address := (uint16(operand1()) + uint16(Y)) | uint16(operand2())<<8
		value = memory[address]
		A |= value
		setFlags()
		incCount(3)
	case INDIRECTX:
		zeroPageAddress := (operand1() + X) & 0xFF
		effectiveAddrLo := memory[zeroPageAddress]
		effectiveAddrHi := memory[(zeroPageAddress+1)&0xFF]
		address := uint16(effectiveAddrHi)<<8 | uint16(effectiveAddrLo)
		value = memory[address]
		A |= value
		setFlags()
		incCount(2)
	case INDIRECTY:
		zeroPageAddress := operand1()
		effectiveAddrLo := memory[zeroPageAddress]
		effectiveAddrHi := memory[(zeroPageAddress+1)&0xFF]
		address := (uint16(effectiveAddrHi)<<8 | uint16(effectiveAddrLo)) + uint16(Y)
		value = memory[address]
		A |= value
		setFlags()
		incCount(2)
	}
}
func BIT(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// Set Negative flag to bit 7 of the value
		if readBit(7, value) == 1 {
			setNegativeFlag()
		}
		// Set Overflow flag to bit 6 of the value
		if readBit(6, value) == 1 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// If the result is 0, set the zero flag
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
	}
	switch addressingMode {
	case ZEROPAGE:
		// Get the address from the operand
		address := operand1()
		// Get the value at the address
		value = memory[address]
		// AND the value with the accumulator
		result = A & value
		setFlags()
		incCount(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = memory[address]
		// AND the value with the accumulator
		result = A & value
		setFlags()
		incCount(3)
	}
}
func INC(addressingMode string) {
	var address uint16
	var result byte

	setFlags := func() {
		// Fetch the value from the address
		value := memory[address]
		// Increment the value (wrapping around for 8-bit values)
		result = value + 1
		// Write the result back to memory
		memory[address] = result

		// Update status flags
		// If bit 7 of the result is set, set the negative flag
		if readBit(7, result) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If the result is 0, set the zero flag
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
	}
	switch addressingMode {
	case ZEROPAGE:
		// Get the address from the operand
		address = uint16(operand1())
		setFlags()
		incCount(2)
	case ZEROPAGEX:
		// Get the address from the operand with X offset
		address = uint16(operand1() + X)
		setFlags()
		incCount(2)
	case ABSOLUTE:
		// Get 16-bit address from operand1 and operand2
		address = uint16(operand2())<<8 | uint16(operand1())
		setFlags()
		incCount(3)
	case ABSOLUTEX:
		// Get 16-bit address from operand1 and operand2 with X offset
		address = (uint16(operand2())<<8 | uint16(operand1())) + uint16(X)
		setFlags()
		incCount(3)
	}
}
func DEC(addressingMode string) {
	var address uint16
	var result byte

	setFlags := func() {
		// Fetch the value from the address
		value := memory[address]
		// Decrement the value (wrapping around for 8-bit values)
		result = value - 1
		// Write the result back to memory
		memory[address] = result

		// Update status flags
		// If bit 7 of the result is set, set the negative flag
		if readBit(7, result) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If the result is 0, set the zero flag
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
	}
	switch addressingMode {
	case ZEROPAGE:
		// Get the address from the operand
		address = uint16(operand1())
		setFlags()
		incCount(2)
	case ZEROPAGEX:
		// Get the address from the operand with X offset
		address = uint16(operand1() + X)
		setFlags()
		incCount(2)
	case ABSOLUTE:
		// Get 16-bit address from operand1 and operand2
		address = uint16(operand2())<<8 | uint16(operand1())
		setFlags()
		incCount(3)
	case ABSOLUTEX:
		// Get 16-bit address from operand1 and operand2 with X offset
		address = (uint16(operand2())<<8 | uint16(operand1())) + uint16(X)
		setFlags()
		incCount(3)
	}
}
func ADC(addressingMode string) {
	var value byte
	var result int

	setFlags := func() {
		// Add the value to the accumulator
		result = int(A) + int(value)
		// If the carry flag is set, add 1 to the result
		if getSRBit(0) == 1 {
			result++
		}

		// Check for BCD mode (Decimal flag set)
		if getSRBit(3) == 1 {
			// Handle BCD addition
			lowerNibble := (A & 0x0F) + (value & 0x0F)
			if getSRBit(0) == 1 {
				lowerNibble++
			}
			if lowerNibble > 9 {
				lowerNibble += 6
			}
			upperNibble := (A >> 4) + (value >> 4)
			if lowerNibble > 0x0F {
				upperNibble++
			}
			if upperNibble > 9 {
				upperNibble += 6
				setCarryFlag()
			} else {
				unsetCarryFlag()
			}
			result = int((upperNibble << 4) | (lowerNibble & 0x0F))
		} else {
			// Binary mode
			// If the result is greater than 255, set the carry flag
			if result > 255 {
				setCarryFlag()
			} else {
				unsetCarryFlag()
			}
		}

		// If result is positive and value is negative, or result is negative and value is positive set the overflow flag
		if (readBit(7, byte(result)) != readBit(7, value)) && (readBit(7, byte(result)) != readBit(7, A)) {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		if readBit(7, byte(result)) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		A = byte(result)

		if addressingMode == IMMEDIATE {
			incCount(2)
		}
		if addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX || addressingMode == INDIRECTX || addressingMode == INDIRECTY {
			incCount(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX || addressingMode == ABSOLUTEY {
			incCount(3)
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = operand1()
		setFlags()
	case ZEROPAGE:
		// Get the address from the operand
		address := operand1()
		// Get the value at the address
		value = memory[address]
		setFlags()
	case ZEROPAGEX:
		// Get the address from the operand
		address := operand1() + X
		// Get the value at the address
		value = memory[address]
		setFlags()
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = memory[address]
		setFlags()
	case ABSOLUTEX:
		// Get 16 bit address from operand1 and operand2
		address := (uint16(operand2())<<8 | uint16(operand1())) + uint16(X)
		// Get value at address
		value = memory[address]
		setFlags()
	case ABSOLUTEY:
		// Get 16 bit address from operand1 and operand2
		address := (uint16(operand2())<<8 | uint16(operand1())) + uint16(Y)
		// Get value at address
		value = memory[address]
		setFlags()
	case INDIRECTX:
		// Get the indirect address from the operand
		indirectAddress := operand1() + X
		// Get the address from the indirect address
		address := uint16(memory[indirectAddress+1])<<8 | uint16(memory[indirectAddress])
		// Get the value at the address
		value = memory[address]
		setFlags()
	case INDIRECTY:
		// Get the indirect address from the operand
		indirectAddress := operand1()
		// Get the address from the indirect address
		address := uint16(memory[indirectAddress+1])<<8 | uint16(memory[indirectAddress]) + uint16(Y)
		// Get the value at the address
		value = memory[address]
		setFlags()
	}
}
func SBC(addressingMode string) {
	var value byte
	var result int

	setFlags := func() {
		// Check for BCD mode (Decimal flag set)
		if getSRBit(3) == 1 {
			// Handle BCD subtraction
			lowerNibble := (A & 0x0F) - (value & 0x0F)
			if getSRBit(0) == 0 {
				lowerNibble--
			}
			upperNibble := (A >> 4) - (value >> 4)
			if lowerNibble&0x10 != 0 {
				upperNibble--
				lowerNibble -= 6
			}
			if upperNibble&0x10 != 0 {
				unsetCarryFlag()
				upperNibble -= 6
			} else {
				setCarryFlag()
			}
			result = int((upperNibble << 4) | (lowerNibble & 0x0F))
		} else {
			// Binary mode
			result = int(A) - int(value)
			if getSRBit(0) == 0 {
				result--
			}
			if int(A) >= int(value) {
				setCarryFlag()
			} else {
				unsetCarryFlag()
			}
		}

		// Overflow, Negative, and Zero flag checks remain the same
		if readBit(7, A) != readBit(7, value) && readBit(7, A) != readBit(7, byte(result)) {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		if readBit(7, byte(result)) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		if result == 0 {
			setZeroFlag()
		}
		A = byte(result)

		// Your addressing mode cycle counts remain the same
		if addressingMode == IMMEDIATE || addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX || addressingMode == INDIRECTX || addressingMode == INDIRECTY {
			incCount(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX || addressingMode == ABSOLUTEY {
			incCount(3)
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = operand1()
		setFlags()
	case ZEROPAGE:
		// Get the address from the operand
		address := operand1()
		// Get the value at the address
		value = memory[address]
		setFlags()
	case ZEROPAGEX:
		// Get the address from the operand
		address := operand1() + X
		// Get the value at the address
		value = memory[address]
		setFlags()
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = memory[address]
		setFlags()
	case ABSOLUTEX:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Get value at address
		value = memory[address]
		setFlags()
	case ABSOLUTEY:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(Y)
		// Get value at address
		value = memory[address]
		setFlags()
	case INDIRECTX:
		// Get the indirect address from the operand
		indirectAddress := operand1() + X
		// Get the address from the indirect address
		address := uint16(memory[indirectAddress+1])<<8 | uint16(memory[indirectAddress])
		// Get the value at the address
		value = memory[address]
		setFlags()
	case INDIRECTY:
		// Get the indirect address from the operand
		indirectAddress := operand1()
		// Get the address from the indirect address
		address := uint16(memory[indirectAddress+1])<<8 | uint16(memory[indirectAddress]) + uint16(Y)
		// Get the value at the address
		value = memory[address]
		setFlags()
	}
}
func ROR(addressingMode string) {
	var address, value, result byte
	var address16 uint16

	setFlags := func() {
		// Set bit 7 of result and negative flag to the carry flag
		if getSRBit(0) == 1 {
			result |= 0x80
			setNegativeFlag()
		} else {
			result &= 0x7F
			unsetNegativeFlag()
		}
		// Set carry flag to bit 0 of value
		if readBit(0, value) == 1 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		if addressingMode == ACCUMULATOR {
			// Store the result in the accumulator
			A = result
			incCount(1)
		}
		if addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX {
			// Store the value back into memory
			memory[address] = result
			incCount(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX {
			// Store the value back into memory
			memory[address16] = result
			incCount(3)
		}
	}
	switch addressingMode {
	case ACCUMULATOR:
		// Get value from accumulator
		value = A
		// Rotate right one bit
		result = value >> 1
		setFlags()
	case ZEROPAGE:
		// Get address
		address = operand1()
		// Get the value at the address
		value = memory[address]
		// Shift the value right 1 bit
		result = value >> 1
		setFlags()
	case ZEROPAGEX:
		// Get X indexed zero page address
		address = operand1() + X
		// Get the value at the address
		value = memory[address]
		// Shift the value right 1 bit
		result = value >> 1
		setFlags()
	case ABSOLUTE:
		// Get 16 bit address from operands
		address16 = uint16(operand2())<<8 | uint16(operand1())
		// Get the value stored at the address in the operands
		value = memory[address16]
		// Shift the value right 1 bit
		result = value >> 1
		setFlags()
	case ABSOLUTEX:
		// Get 16 bit address
		address16 = uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Get value stored at address
		value = memory[address16]
		// Shift right the value by 1 bit
		result = value >> 1
		setFlags()
	}
}
func ROL(addressingMode string) {
	var address, value, result byte
	var address16 uint16

	setFlags := func() {
		// Set SR carry flag to bit 7 of value
		if readBit(7, value) == 1 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}

		// Update the zero flag
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}

		// Update the negative flag
		if readBit(7, A) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		if addressingMode == ACCUMULATOR {
			// Store the result in the accumulator
			A = result
			incCount(1)
		}
		if addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX {
			// Store the value back into memory
			memory[address] = result
			incCount(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX {
			// Store the value back into memory
			memory[address16] = result
			incCount(3)
		}
	}
	switch addressingMode {
	case ACCUMULATOR:
		// Get the value of the accumulator
		value = A
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | getSRBit(0)
		setFlags()
	case ZEROPAGE:
		// Get address
		address = operand1()
		// Get the value at the address
		value = memory[address]
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | getSRBit(0)
		setFlags()
	case ZEROPAGEX:
		// Get X indexed zero page address
		address = operand1() + X
		// Get the value at the address
		value = memory[address]
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | getSRBit(0)
		setFlags()
	case ABSOLUTE:
		// Get 16 bit address from operands
		address16 = uint16(operand2())<<8 | uint16(operand1())
		// Get the value stored at the address in the operands
		value = memory[address16]
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | getSRBit(0)
		setFlags()
	case ABSOLUTEX:
		// Get 16bit X indexed absolute memory address
		address16 = uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Get the value stored at the address
		value = memory[address16]
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | getSRBit(0)
		setFlags()
	}
}
func LSR(addressingMode string) {
	var value, result byte

	setFlags := func() {
		// Reset the SR negative flag
		unsetNegativeFlag()
		// If A is 0 then set SR zero flag else reset it
		// Update the zero flag
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If bit 0 of value is 1 then set SR carry flag else reset it
		if readBit(0, value) == 1 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
	}

	switch addressingMode {
	case ACCUMULATOR:
		// Get the value of the accumulator
		value = A
		// Shift the value right 1 bit
		result = value >> 1
		// Store the result back into the accumulator
		A = result
		setFlags()
		incCount(1)
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get the value at the address
		value = memory[address]
		// Shift the value right 1 bit
		result = value >> 1
		// Store the value back into memory
		memory[address] = result
		setFlags()
		incCount(2)
	case ZEROPAGEX:
		// Get the X indexed address
		address := operand1() + X
		// Get the value at the X indexed address
		value = memory[address]
		// Shift the value right 1 bit
		result = value >> 1
		// Store the shifted value in memory
		memory[address] = result
		setFlags()
		incCount(2)
	case ABSOLUTE:
		// Get 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get the value stored at the address in the operands
		value = memory[address]
		// Shift the value right 1 bit
		result = value >> 1
		// Store the shifted value back in memory
		memory[address] = result
		setFlags()
		incCount(3)
	case ABSOLUTEX:
		// Get the 16bit X indexed absolute memory address
		address := uint16(operand2())<<8 | uint16(operand1())
		address += uint16(X)
		// Get the value stored at the address
		value = memory[address]
		// Shift the value right 1 bit
		result = value >> 1
		// Store the shifted value back in memory
		memory[address] = result
		setFlags()
		incCount(3)
	}
}

func ASL(addressingMode string) {
	var value, result byte

	setFlags := func() {
		// Update the zero flag
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}

		// Update the negative flag
		if readBit(7, A) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}

		// Set the Carry flag based on the original value's bit 7 before the shift operation
		if readBit(7, value) == 1 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
	}
	switch addressingMode {
	case ACCUMULATOR:
		// Set value to accumulator
		value = A
		// Shift the value left 1 bit
		result = value << 1
		// Update the accumulator with the result
		A = result
		setFlags()
		incCount(1)
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get the value at the address
		value = memory[address]
		// Shift the value left 1 bit
		result = value << 1
		// Store the value back into memory
		memory[address] = result
		setFlags()
		incCount(2)
	case ZEROPAGEX:
		// Get the X indexed address
		address := operand1() + X
		// Get the value at the X indexed address
		value = memory[address]
		// Shift the value left 1 bit
		result = value << 1
		// Store the shifted value in memory
		memory[address] = result
		setFlags()
		incCount(2)
	case ABSOLUTE:
		// Get 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get the value stored at the address in the operands
		value = memory[address]
		// Shift the value left 1 bit
		result = value << 1
		// Store the shifted value back in memory
		memory[address] = result
		setFlags()
		incCount(3)
	case ABSOLUTEX:
		// Get the 16bit X indexed absolute memory address
		address := int(operand2())<<8 | int(operand1()) + int(X)
		// Get the value stored at the address
		value = memory[address]
		// Shift the value left 1 bit
		result = value << 1
		// Store the shifted value back in memory
		memory[address] = result
		setFlags()
		incCount(3)
	}
}

func CPX(addressingMode string) {
	var value, result byte

	setFlags := func() {
		// If X >= value then set carry flag bit 0 to 1 set carry flag bit 0 to 0
		if X >= value {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		// If value> X then reset carry flag
		if value > X {
			unsetCarryFlag()
		}
		// If bit 7 of result is 1 then set negative flag else unset negative flag
		if readBit(7, result) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If value == X then set zero flag else unset zero flag
		if value == X {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get value from operand1
		value = operand1()
		// Compare X with value
		result = X - value
		setFlags()
		incCount(2)
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get value at address
		value = memory[address]
		// Store result of X-memory stored at operand1() in result variable
		result = X - value
		setFlags()
		incCount(2)
	case ABSOLUTE:
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = memory[address]
		setFlags()
		incCount(3)
	}
}
func CPY(addressingMode string) {
	var value, result byte

	setFlags := func() {
		// If Y>value then set carry flag to 1 else set carry flag to 0
		if Y >= value {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		// If bit 7 of result is set, set N flag to 1 else reset it
		if readBit(7, result) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If Y==value then set Z flag to 1 else reset it
		if Y == value {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get value from operand1
		value = operand1()
		// Subtract operand from Y
		result = Y - operand1()
		setFlags()
		incCount(2)
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get value at address
		value = memory[address]
		// Store result of Y-memory stored at operand1() in result variable
		result = Y - value
		setFlags()
		incCount(2)
	case ABSOLUTE:
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = memory[address]
		setFlags()
		incCount(3)
	}
}

func execute() {
	for PC < len(memory) {
		//debug option to break out from borked opcode infinite loop
		/*if PC == 0x45C0 {
			// exit to OS
			os.Exit(0)
		}*/

		//  1 byte instructions with no operands
		switch opcode() {
		// Implied addressing mode instructions
		/*
			In the implied addressing mode, the address containing the operand is implicitly stated in the operation code of the instruction.

			Bytes: 1
		*/
		case 0x00:
			/*
				BRK - Break Command
			*/
			disassembledInstruction = fmt.Sprintf("BRK")
			disassembleOpcode()
			// Increment PC
			PC++

			// Decrement SP and Push high byte of (PC+1) onto stack
			decSP()
			memory[SP] = byte((PC + 1) >> 8) // High byte

			// Decrement SP and Push low byte of (PC+1) onto stack
			decSP()
			memory[SP] = byte((PC + 1) & 0xFF) // Low byte

			// Set SR break flag
			setBreakFlag()

			// Decrement SP and Store SR on stack
			decSP()
			memory[SP] = SR

			// Set SR interrupt disable bit to 1
			setInterruptFlag()

			// Set PC to interrupt vector address
			//PC = interruptVectorAddress | interruptVectorAddress + 1
			PC = int(uint16(memory[0xFFFF])<<8 | uint16(memory[0xFFFE])) //PC = int(memory[0xFFFE]) + int(memory[0xFFFF])*256

			incCount(0)
		case 0x18:
			/*
				CLC - Clear Carry Flag
			*/
			// print the SR as binary digits
			disassembledInstruction = fmt.Sprintf("CLC")
			disassembleOpcode()
			// Set SR carry flag bit 0 to 0
			unsetCarryFlag()
			incCount(1)
		case 0xD8:
			/*
				CLD - Clear Decimal Mode
			*/

			disassembledInstruction = fmt.Sprintf("CLD")
			disassembleOpcode()
			unsetDecimalFlag()
			incCount(1)
		case 0x58:
			/*
				CLI - Clear Interrupt Disable
			*/
			disassembledInstruction = fmt.Sprintf("CLI")
			disassembleOpcode()
			// Set SR interrupt disable bit 2 to 0
			unsetInterruptFlag()
			incCount(1)
		case 0xB8:
			/*
				CLV - Clear Overflow Flag
			*/
			disassembledInstruction = fmt.Sprintf("CLV")
			disassembleOpcode()
			// Set SR overflow flag bit 6 to 0
			unsetOverflowFlag()
			incCount(1)
		case 0xCA:
			/*
				DEX - Decrement Index Register X By One
			*/
			disassembledInstruction = fmt.Sprintf("DEX")
			disassembleOpcode()

			// Decrement the X register by 1
			X--
			// If X register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getXBit(7) == 1 {
				setNegativeFlag()
			} else {
				unsetNegativeFlag()
			}
			// If X register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
			if X == 0 {
				setZeroFlag()
			} else {
				unsetZeroFlag()
			}
			incCount(1)
		case 0x88:
			/*
				DEY - Decrement Index Register Y By One
			*/
			disassembledInstruction = fmt.Sprintf("DEY")
			disassembleOpcode()

			// Decrement the  Y register by 1
			Y--
			// If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getYBit(7) == 1 {
				setNegativeFlag()
			} else {
				unsetNegativeFlag()
			}
			// If Y==0 then set the Zero flag
			if Y == 0 {
				setZeroFlag()
			} else {
				unsetZeroFlag()
			}
			incCount(1)
		case 0xE8:
			/*
				INX - Increment Index Register X By One
			*/
			disassembledInstruction = fmt.Sprintf("INX")
			disassembleOpcode()

			// Increment the X register by 1
			X++
			// If X register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getXBit(7) == 1 {
				setNegativeFlag()
			} else {
				unsetNegativeFlag()
			}
			// If X register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
			if X == 0 {
				setZeroFlag()
			} else {
				unsetZeroFlag()
			}
			incCount(1)
		case 0xC8:
			/*
				INY - Increment Index Register Y By One
			*/
			disassembledInstruction = fmt.Sprintf("INY")
			disassembleOpcode()

			// Increment the  Y register by 1
			Y++
			// If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getYBit(7) == 1 {
				setNegativeFlag()
			} else {
				unsetNegativeFlag()
			}
			// If Y register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
			if Y == 0 {
				setZeroFlag()
			} else {
				unsetZeroFlag()
			}
			incCount(1)
		case 0xEA:
			/*
				NOP - No Operation
			*/
			disassembledInstruction = fmt.Sprintf("NOP")
			disassembleOpcode()
			incCount(1)
		case 0x48:
			/*
				PHA - Push Accumulator On Stack
			*/
			disassembledInstruction = fmt.Sprintf("PHA")
			disassembleOpcode()

			// Update memory address pointed to by SP with value stored in accumulator
			memory[SPBaseAddress+SP] = A
			// Decrement the stack pointer by 1 byte
			if SP > 0 {
				SP--
			} else {
				SP = 0xFF
			}
			incCount(1)
		case 0x08:
			/*
			   PHP - Push Processor Status On Stack
			*/
			disassembledInstruction = fmt.Sprintf("PHP")
			disassembleOpcode()

			// Set the break flag and the unused bit before pushing
			SR |= (1 << 4) // Set break flag
			SR |= (1 << 5) // Set unused bit

			// Push the SR onto the stack
			memory[SPBaseAddress+SP] = SR

			// Decrement the stack pointer
			decSP()

			incCount(1)
		case 0x68:
			/*
			   PLA - Pull Accumulator From Stack
			*/
			disassembledInstruction = fmt.Sprintf("PLA")
			disassembleOpcode()

			// Increment the stack pointer first
			incSP()

			// Ensure all arithmetic is done in uint16
			expectedAddress := uint16(SPBaseAddress) + uint16(SP)

			// Now, update accumulator with value stored in memory address pointed to by SP
			A = memory[expectedAddress]

			// If bit 7 of accumulator is set, set negative SR flag else set negative SR flag to 0
			if getABit(7) == 1 {
				setNegativeFlag()
			} else {
				unsetNegativeFlag()
			}

			// If accumulator is 0, set zero SR flag else set zero SR flag to 0
			if A == 0 {
				setZeroFlag()
			} else {
				unsetZeroFlag()
			}

			incCount(1)
		case 0x28:
			/*
				PLP - Pull Processor Status From Stack
			*/
			disassembledInstruction = fmt.Sprintf("PLP")
			disassembleOpcode()

			// Update SR with the value stored at the address pointed to by SP
			SR = memory[SPBaseAddress+SP]
			incSP()
			incCount(1)
		case 0x40:
			/*
			   RTI - Return From Interrupt
			*/

			disassembledInstruction = fmt.Sprintf("RTI")
			disassembleOpcode()

			SR = memory[SPBaseAddress+SP] & 0xCF
			incSP()

			// Increment the stack pointer to get low byte of PC
			incSP()

			// Get low byte of PC
			low := uint16(memory[SPBaseAddress+SP])

			// Increment the stack pointer to get high byte of PC
			incSP()

			// Get high byte of PC
			high := uint16(memory[SPBaseAddress+SP])

			incCount(0)
			// Update PC with the value stored in memory at the address pointed to by SP
			PC = int((high << 8) | low)
		case 0x60:
			/*
				RTS - Return From Subroutine
			*/
			disassembledInstruction = fmt.Sprintf("RTS")
			disassembleOpcode()
			// Increment the stack pointer by 1 byte
			//SP++
			//Get low byte of PC
			low := uint16(memory[SP])
			// Increment the stack pointer by 1 byte
			//SP++
			//Get high byte of PC
			high := uint16(memory[SP+1])
			//Update PC with the value stored in memory at the address pointed to by SP
			incCount(0)
			PC = int((high << 8) | low)
		case 0x38:
			/*
				SEC - Set Carry Flag
			*/
			disassembledInstruction = fmt.Sprintf("SEC")
			disassembleOpcode()

			// Set SR carry flag bit 0 to 1
			setCarryFlag()
			incCount(1)
		case 0xF8:
			/*
				SED - Set Decimal Mode
			*/
			disassembledInstruction = fmt.Sprintf("SED")
			disassembleOpcode()

			// Set SR decimal mode flag to 1
			setDecimalFlag()
			incCount(1)
		case 0x78:
			/*
				SEI - Set Interrupt Disable
			*/
			disassembledInstruction = fmt.Sprintf("SEI")
			disassembleOpcode()

			// Set SR interrupt disable bit 2 to 1
			setInterruptFlag()
			incCount(1)
		case 0xAA:
			/*
				TAX - Transfer Accumulator To Index X
			*/
			disassembledInstruction = fmt.Sprintf("TAX")
			disassembleOpcode()

			// Update X with the value of A
			X = A
			// If X register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getXBit(7) == 1 {
				setNegativeFlag()
			} else {
				unsetNegativeFlag()
			}
			// If X register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
			if X == 0 {
				setZeroFlag()
			} else {
				unsetZeroFlag()
			}
			incCount(1)
		case 0xA8:
			/*
				TAY - Transfer Accumulator To Index Y
			*/
			disassembledInstruction = fmt.Sprintf("TAY")
			disassembleOpcode()

			// Set Y register to the value of the accumulator
			Y = A
			// If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getYBit(7) == 1 {
				setNegativeFlag()
			} else {
				unsetNegativeFlag()
			}
			// If Y register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
			if A == 0 {
				setZeroFlag()
			} else {
				unsetZeroFlag()
			}
			incCount(1)
		case 0xBA:
			/*
				TSX - Transfer Stack Pointer To Index X
			*/
			disassembledInstruction = fmt.Sprintf("TSX")
			disassembleOpcode()

			// Update X with the SP
			X = byte(SP)
			// If X register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getXBit(7) == 1 {
				setNegativeFlag()
			} else {
				unsetNegativeFlag()
			}
			// If X register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
			if X == 0 {
				setZeroFlag()
			} else {
				unsetZeroFlag()
			}
			incCount(1)
		case 0x8A:
			/*
				TXA - Transfer Index X To Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("TXA")
			disassembleOpcode()

			// Set accumulator to value of X register
			A = X
			// If bit 7 of accumulator is set, set negative SR flag else set negative SR flag to 0
			if getABit(7) == 1 {
				setNegativeFlag()
			} else {
				unsetNegativeFlag()
			}
			// If accumulator is 0, set zero SR flag else set zero SR flag to 0
			if A == 0 {
				setZeroFlag()
			} else {
				unsetZeroFlag()
			}
			incCount(1)
		case 0x9A:
			/*
				TXS - Transfer Index X To Stack Pointer
			*/
			disassembledInstruction = fmt.Sprintf("TXS")
			disassembleOpcode()

			// Set stack pointer to value of X register
			SP = uint(X)
			incCount(1)
		case 0x98:
			/*
				TYA - Transfer Index Y To Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("TYA")
			disassembleOpcode()

			// Set accumulator to value of Y register
			A = Y
			// If bit 7 of accumulator is set, set negative SR flag else set negative SR flag to 0
			if getABit(7) == 1 {
				setNegativeFlag()
			} else {
				unsetNegativeFlag()
			}
			// If accumulator is 0, set zero SR flag else set zero SR flag to 0
			if A == 0 {
				setZeroFlag()
			} else {
				unsetZeroFlag()
			}
			incCount(1)

		// Accumulator instructions
		/*
			A

			This form of addressing is represented with a one byte instruction, implying an operation on the accumulator.

			Bytes: 1
		*/
		case 0x0A:
			/*
				ASL - Arithmetic Shift Left
			*/
			disassembledInstruction = fmt.Sprintf("ASL")
			disassembleOpcode()
			ASL("accumulator")
		case 0x4A:
			/*
				LSR - Logical Shift Right
			*/
			disassembledInstruction = fmt.Sprintf("LSR")
			disassembleOpcode()

			LSR("accumulator")
		case 0x2A:
			/*
				ROL - Rotate Left
			*/
			disassembledInstruction = fmt.Sprintf("ROL")
			disassembleOpcode()

			ROL("accumulator")
		case 0x6A:
			/*
				ROR - Rotate Right
			*/
			disassembledInstruction = fmt.Sprintf("ROR")
			disassembleOpcode()
			ROR("accumulator")
		}

		// 2 byte instructions with 1 operand
		switch opcode() {
		// Immediate addressing mode instructions
		/*
			#$nn

			In immediate addressing, the operand is contained in the second byte of the instruction, with no further memory addressing required.

			Bytes: 2
		*/
		case 0x69:
			/*
				ADC - Add Memory to Accumulator with Carry
			*/
			disassembledInstruction = fmt.Sprintf("ADC #$%02X", operand1())
			disassembleOpcode()

			ADC("immediate")
		case 0x29:
			/*
				AND - "AND" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("AND #$%02X", operand1())
			disassembleOpcode()

			AND("immediate")
		case 0xC9:
			/*
				CMP - Compare Memory and Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("CMP #$%02X", operand1())
			disassembleOpcode()

			CMP("immediate")
		case 0xE0:
			/*
				CPX - Compare Index Register X To Memory
			*/
			disassembledInstruction = fmt.Sprintf("CPX #$%02X", operand1())
			disassembleOpcode()

			CPX("immediate")
		case 0xC0:
			/*
				CPY - Compare Index Register Y To Memory
			*/
			disassembledInstruction = fmt.Sprintf("CPY #$%02X", operand1())
			disassembleOpcode()

			CPY("immediate")
		case 0x49:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("EOR #$%02X", operand1())
			disassembleOpcode()

			EOR("immediate")
		case 0xA9:
			/*
				LDA - Load Accumulator with Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDA #$%02X", operand1())
			disassembleOpcode()
			LDA("immediate")
		case 0xA2:
			/*
				LDX - Load Index Register X From Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDX #$%02X", operand1())
			disassembleOpcode()

			LDX("immediate")
		case 0xA0:
			/*
				LDY - Load Index Register Y From Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDY #$%02X", operand1())
			disassembleOpcode()

			LDY("immediate")
		case 0x09:
			/*
				ORA - "OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("ORA #$%02X", operand1())
			disassembleOpcode()

			ORA("immediate")
		case 0xE9:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
			*/
			disassembledInstruction = fmt.Sprintf("SBC #$%02X", operand1())
			disassembleOpcode()
			SBC("immediate")

		// Zero Page addressing mode instructions
		/*
			$nn

			The zero page instructions allow for shorter code and execution times by only fetching the second byte of the instruction and assuming a zero low address byte. Careful use of the zero page can result in significant increase in code efficiency.

			Bytes: 2
		*/
		case 0x65:
			/*
				ADC - Add Memory to Accumulator with Carry
			*/
			disassembledInstruction = fmt.Sprintf("ADC $%02X", operand1())
			disassembleOpcode()

			ADC("zeropage")
		case 0x25:
			/*
				AND - "AND" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("AND $%02X", operand1())
			disassembleOpcode()

			AND("zeropage")
		case 0x06:
			/*
				ASL - Arithmetic Shift Left
			*/
			disassembledInstruction = fmt.Sprintf("ASL $%02X", operand1())
			disassembleOpcode()

			ASL("zeropage")
		case 0x24:
			/*
				BIT - Test Bits in Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("BIT $%02X", operand1())
			disassembleOpcode()

			BIT("zeropage")
		case 0xC5:
			/*
				CMP - Compare Memory and Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("CMP $%02X", operand1())
			disassembleOpcode()
			CMP("zeropage")
		case 0xE4:
			/*
				CPX - Compare Index Register X To Memory
			*/
			disassembledInstruction = fmt.Sprintf("CPX $%02X", operand1())
			disassembleOpcode()

			CPX("zeropage")
		case 0xC4:
			/*
				CPY - Compare Index Register Y To Memory
			*/
			disassembledInstruction = fmt.Sprintf("CPY $%02X", operand1())
			disassembleOpcode()
			CPY("zeropage")
		case 0xC6:
			/*
				DEC - Decrement Memory By One
			*/
			disassembledInstruction = fmt.Sprintf("DEC $%02X", operand1())
			disassembleOpcode()

			DEC("zeropage")
		case 0x45:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("EOR $%02X", operand1())
			disassembleOpcode()

			EOR("zeropage")
		case 0xE6:
			/*
				INC - Increment Memory By One
			*/
			disassembledInstruction = fmt.Sprintf("INC $%02X", operand1())
			disassembleOpcode()

			INC("zeropage")
		case 0xA5:
			/*
				LDA - Load Accumulator with Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDA $%02X", operand1())
			disassembleOpcode()

			LDA("zeropage")
		case 0xA6:
			/*
				LDX - Load Index Register X From Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDX $%02X", operand1())
			disassembleOpcode()
			LDX("zeropage")
		case 0xA4:
			/*
				LDY - Load Index Register Y From Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDY $%02X", operand1())
			disassembleOpcode()

			LDY("zeropage")
		case 0x46:
			/*
				LSR - Logical Shift Right
			*/
			disassembledInstruction = fmt.Sprintf("LSR $%02X", operand1())
			disassembleOpcode()

			LSR("zeropage")
		case 0x05:
			/*
				ORA - "OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("ORA $%02X", operand1())
			disassembleOpcode()

			ORA("zeropage")
		case 0x26:
			/*
				ROL - Rotate Left
			*/
			disassembledInstruction = fmt.Sprintf("ROL $%02X", operand1())
			disassembleOpcode()

			ROL("zeropage")
		case 0x66:
			/*
				ROR - Rotate Right
			*/
			disassembledInstruction = fmt.Sprintf("ROR $%02X", operand1())
			disassembleOpcode()

			ROR("zeropage")
		case 0xE5:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
			*/
			disassembledInstruction = fmt.Sprintf("SBC $%02X", operand1())
			disassembleOpcode()

			SBC("zeropage")
		case 0x85:
			/*
				STA - Store Accumulator in Memory
			*/

			disassembledInstruction = fmt.Sprintf("STA $%02X", operand1())
			disassembleOpcode()

			STA("zeropage")
		case 0x86:
			/*
				STX - Store Index Register X In Memory
			*/
			disassembledInstruction = fmt.Sprintf("STX $%02X", operand1())
			disassembleOpcode()
			STX("zeropage")
		case 0x84:
			/*
				STY - Store Index Register Y In Memory
			*/
			disassembledInstruction = fmt.Sprintf("STY $%02X", operand1())
			disassembleOpcode()

			STY("zeropage")

		// X Indexed Zero Page addressing mode instructions
		/*
			$nn,X

			This form of addressing is used in conjunction with the X index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the low order 8 bits of memory and crossing of page boundaries does not occur.

			Bytes: 2
		*/
		case 0x75:
			/*
				ADC - Add Memory to Accumulator with Carry
			*/
			disassembledInstruction = fmt.Sprintf("ADC $%02X,X", operand1())
			disassembleOpcode()

			ADC("zeropagex")
		case 0x35:
			/*
				AND - "AND" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("AND $%02X,X", operand1())
			disassembleOpcode()

			AND("zeropagex")
		case 0x16:
			/*
				ASL - Arithmetic Shift Left
			*/
			disassembledInstruction = fmt.Sprintf("ASL $%02X,X", operand1())
			disassembleOpcode()

			ASL("zeropagex")
		case 0xD5:
			/*
				CMP - Compare Memory and Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("CMP $%02X,X", operand1())
			disassembleOpcode()

			CMP("zeropagex")
		case 0xD6:
			/*
				DEC - Decrement Memory By One
			*/
			disassembledInstruction = fmt.Sprintf("DEC $%02X,X", operand1())
			disassembleOpcode()

			DEC("zeropagex")
		case 0xB5:
			/*
				LDA - Load Accumulator with Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDA $%02X,X", operand1())
			disassembleOpcode()

			LDA("zeropagex")
		case 0xB4:
			/*
				LDY - Load Index Register Y From Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDY $%02X,X", operand1())
			disassembleOpcode()

			LDY("zeropagex")
		case 0x56:
			/*
				LSR - Logical Shift Right
			*/
			disassembledInstruction = fmt.Sprintf("LSR $%02X,X", operand1())
			disassembleOpcode()

			LSR("zeropagex")
		case 0x15:
			/*
				ORA - "OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("ORA $%02X,X", operand1())
			disassembleOpcode()

			ORA("zeropagex")
		case 0x36:
			/*
				ROL - Rotate Left
			*/
			disassembledInstruction = fmt.Sprintf("ROL $%02X,X", operand1())
			disassembleOpcode()
			ROL("zeropagex")
		case 0x76:
			/*
				ROR - Rotate Right
			*/
			disassembledInstruction = fmt.Sprintf("ROR $%02X,X", operand1())
			disassembleOpcode()
			ROR("zeropagex")
		case 0xF5:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
			*/
			disassembledInstruction = fmt.Sprintf("SBC $%02X,X", operand1())
			disassembleOpcode()
			SBC("zeropagex")
		case 0x95:
			/*
				STA - Store Accumulator in Memory
			*/
			disassembledInstruction = fmt.Sprintf("STA $%02X,X", operand1())
			disassembleOpcode()
			STA("zeropagex")
		case 0x94:
			/*
				STY - Store Index Register Y In Memory
			*/
			disassembledInstruction = fmt.Sprintf("STY $%02X,X", operand1())
			disassembleOpcode()

			STY("zeropagex")

		// Y Indexed Zero Page addressing mode instructions
		/*
			$nn,Y

			This form of addressing is used in conjunction with the Y index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the low order 8 bits of memory and crossing of page boundaries does not occur.

			Bytes: 2
		*/
		case 0xB6:
			/*
				LDX - Load Index Register X From Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDX $%02X,Y", operand1())
			disassembleOpcode()
			LDX("zeropagey")
		case 0x96:
			/*
				STX - Store Index Register X In Memory
			*/
			disassembledInstruction = fmt.Sprintf("STX $%02X,Y", operand1())
			disassembleOpcode()

			STX("zeropagey")

		// X Indexed Zero Page Indirect addressing mode instructions
		/*
			($nn,X)

			In indexed indirect addressing, the second byte of the instruction is added to the contents of the X index register, discarding the carry. The result of this addition points to a memory location on page zero whose contents is the high order eight bits of the effective address. The next memory location in page zero contains the low order eight bits of the effective address. Both memory locations specifying the low and high order bytes of the effective address must be in page zero.

			Bytes: 2
		*/
		case 0x61:
			/*
				ADC - Add Memory to Accumulator with Carry
			*/
			disassembledInstruction = fmt.Sprintf("ADC ($%02X,X)", operand1())
			disassembleOpcode()
			ADC("indirectx")
		case 0x21:
			/*
				AND - "AND" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("AND ($%02X,X)", operand1())
			disassembleOpcode()

			AND("indirectx")
		case 0xC1:
			/*
				CMP - Compare Memory and Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("CMP ($%02X,X)", operand1())
			disassembleOpcode()

			CMP("indirectx")
		case 0x41:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("EOR ($%02X,X)", operand1())
			disassembleOpcode()

			EOR("indirectx")
		case 0xA1:
			/*
				LDA - Load Accumulator with Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDA ($%02X,X)", operand1())
			disassembleOpcode()

			LDA("indirectx")
		case 0x01:
			/*
				ORA - "OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("ORA ($%02X,X)", operand1())
			disassembleOpcode()

			ORA("indirectx")
		case 0xE1:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
			*/
			disassembledInstruction = fmt.Sprintf("SBC ($%02X,X)", operand1())
			disassembleOpcode()

			SBC("indirectx")
		case 0x81:
			/*
				STA - Store Accumulator in Memory
			*/
			disassembledInstruction = fmt.Sprintf("STA ($%02X,X)", operand1())
			disassembleOpcode()
			STA("indirectx")

		// Zero Page Indirect Y Indexed addressing mode instructions
		/*
			($nn),Y

			In indirect indexed addressing, the second byte of the instruction points to a memory location in page zero. The contents of this memory location is added to the contents of the Y index register, the result being the high order eight bits of the effective address. The carry from this addition is added to the contents of the next page zero memory location, the result being the low order eight bits of the effective address.

			Bytes: 2
		*/
		case 0x71:
			/*
				ADC - Add Memory to Accumulator with Carry
			*/
			disassembledInstruction = fmt.Sprintf("ADC ($%02X),Y", operand1())
			disassembleOpcode()

			ADC("indirecty")
		case 0x31:
			/*
				AND - "AND" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("AND ($%02X),Y", operand1())
			disassembleOpcode()

			AND("indirecty")
		case 0xD1:
			/*
				CMP - Compare Memory and Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("CMP ($%02X),Y", operand1())
			disassembleOpcode()

			CMP("indirecty")
		case 0x51:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("EOR ($%02X),Y", operand1())
			disassembleOpcode()

			EOR("indirecty")
		case 0xB1:
			/*
				LDA - Load Accumulator with Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDA ($%02X),Y", operand1())
			disassembleOpcode()

			LDA("indirecty")
		case 0x11:
			/*
				ORA - "OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("ORA ($%02X),Y", operand1())
			disassembleOpcode()
			ORA("indirecty")

		case 0xF1:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
			*/
			disassembledInstruction = fmt.Sprintf("SBC ($%02X),Y", operand1())
			disassembleOpcode()

			SBC("indirecty")
		case 0x91:
			/*
				STA - Store Accumulator in Memory
			*/
			disassembledInstruction = fmt.Sprintf("STA ($%02X),Y", operand1())
			disassembleOpcode()

			STA("indirecty")

		// Relative addressing mode instructions
		/*
			$nnnn

			Relative addressing is used only with branch instructions and establishes a destination for the conditional branch.

			The second byte of-the instruction becomes the operand which is an “Offset" added to the contents of the lower eight bits of the program counter when the counter is set at the next instruction. The range of the offset is —128 to +127 bytes from the next instruction.

			Bytes: 2
		*/
		case 0x10:
			/*
				BPL - Branch on Result Plus
			*/
			disassembledInstruction = fmt.Sprintf("BPL $%02X", (PC+2+int(operand1()))&0xFF)
			disassembleOpcode()

			offset := operand1()
			signedOffset := int8(offset)
			// Calculate the branch target address
			targetAddress := PC + 2 + int(signedOffset)
			// If N flag is not set, branch to address
			if getSRBit(7) == 0 {
				// Branch
				PC = targetAddress
				//incCount(0)
				instructionCounter++
			} else {
				// Don't branch
				// Increment the instruction counter by 2
				incCount(2)
			}

		case 0x30:
			/*
				BMI - Branch on Result Minus
			*/
			disassembledInstruction = fmt.Sprintf("BMI $%02X", (PC+2+int(operand1()))&0xFF)
			disassembleOpcode()

			// Get offset from operand
			offset := operand1()
			// If N flag is set, branch to address
			if getSRBit(7) == 1 {
				incCount(0)
				// Branch
				// Add offset to lower 8bits of PC
				PC = PC + 3 + int(offset)&0xFF
				// If the offset is negative, decrement the PC by 1
				// If bit 7 is unset then it's negative
				if readBit(7, offset) == 0 {
					PC--
				}
			} else {
				// Don't branch
				incCount(2)
			}
		case 0x50:
			/*
				BVC - Branch on Overflow Clear
			*/
			disassembledInstruction = fmt.Sprintf("BVC $%02X", PC+2+int(operand1()))
			disassembleOpcode()

			// Get offset from operand
			offset := operand1()
			// If overflow flag is not set, branch to address
			if getSRBit(6) == 0 {
				incCount(0)
				// Branch
				// Add offset to lower 8bits of PC
				PC = PC + 3 + int(offset)&0xFF
				// If the offset is negative, decrement the PC by 1
				// If bit 7 is unset then it's negative
				if readBit(7, offset) == 0 {
					PC--
				}
			} else {
				// Don't branch
				incCount(2)
			}
		case 0x55:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("EOR $%02X,X", operand1())
			disassembleOpcode()

			EOR("zeropagex")
		case 0x70:
			/*
				BVS - Branch on Overflow Set
			*/
			disassembledInstruction = fmt.Sprintf("BVS $%02X", PC+2+int(operand1()))
			disassembleOpcode()

			// Get offset from operand
			offset := operand1()
			// If overflow flag is set, branch to address
			if getSRBit(6) == 1 {
				incCount(0)
				// Branch
				// Add offset to lower 8bits of PC
				PC = PC + 3 + int(offset)&0xFF
				// If the offset is negative, decrement the PC by 1
				// If bit 7 is unset then it's negative
				if readBit(7, offset) == 0 {
					PC--
				}
			} else {
				// Don't branch
				incCount(2)
			}
		case 0x90:
			/*
				BCC - Branch on Carry Clear
			*/
			disassembledInstruction = fmt.Sprintf("BCC $%02X", (PC + 2 + int(operand1())))
			disassembleOpcode()

			// Get offset from operand
			offset := operand1()
			// If carry flag is unset, branch to address
			if getSRBit(0) == 0 {
				incCount(0)
				// Branch
				// Add offset to lower 8bits of PC
				PC = PC + 3 + int(offset)&0xFF
				// If the offset is negative, decrement the PC by 1
				// If bit 7 is unset then it's negative
				if readBit(7, offset) == 0 {
					PC--
				}
			} else {
				// Don't branch
				incCount(2)
			}
		case 0xB0:
			/*
				BCS - Branch on Carry Set
			*/
			disassembledInstruction = fmt.Sprintf("BCS $%02X", (PC+2+int(operand1()))&0xFF)
			disassembleOpcode()
			// Get offset from operand
			offset := operand1()
			// If carry flag is set, branch to address
			if getSRBit(0) == 1 {
				incCount(0)
				// Branch
				// Add offset to lower 8bits of PC
				PC = PC + 3 + int(offset)&0xFF
				// If the offset is negative, decrement the PC by 1
				// If bit 7 is unset then it's negative
				if readBit(7, offset) == 0 {
					PC--
				}
			} else {
				// Don't branch
				incCount(2)
			}

		case 0xD0:
			/*
				BNE - Branch on Result Not Zero
			*/

			disassembledInstruction = fmt.Sprintf("BNE $%02X", (PC+2+int(operand1()))&0xFF)
			disassembleOpcode()

			// Fetch offset from operand
			offset := operand1()

			// Check Z flag to determine if branching is needed
			if getSRBit(1) == 0 {
				// Calculate the branch target address
				targetAddr := PC + 2 + int(int8(offset))

				// Check if the branch crosses a page boundary
				if (PC & 0xFF00) != (targetAddr & 0xFF00) {
					incCount(2) // Add extra cycle if branch crosses page boundary
				} else {
					incCount(1) // Account for the cycle used by the branch instruction
				}

				// Update the program counter
				PC = targetAddr & 0xFFFF // Ensure the address wraps around correctly
			} else {
				// If Z flag is set, don't branch
				incCount(2)
			}

		case 0xF0:
			/*
				BEQ - Branch on Result Zero
			*/
			disassembledInstruction = fmt.Sprintf("BEQ $%02X", (PC + 2 + int(operand1())))
			disassembleOpcode()

			// Get offset from address in operand
			offset := int8(operand1()) // Cast to signed 8-bit integer to handle negative offsets
			// Get relative address from offset
			relativeAddress := PC + 2 + int(offset)

			// If Z flag is set, branch to address
			if getSRBit(1) == 1 {
				incCount(0)
				PC = relativeAddress
			} else {
				incCount(2)
			}

		case 0xF6:
			/*
				INC - Increment Memory By One
			*/
			disassembledInstruction = fmt.Sprintf("INC $%02X,X", operand1())
			disassembleOpcode()

			INC("zeropagex")
		}

		// 3 byte instructions with 2 operands
		switch opcode() {
		// Absolute addressing mode instructions
		/*
			$nnnn

			In absolute addressing, the second byte of the instruction specifies the eight high order bits of the effective address while the third byte specifies the eight low order bits. Thus, the absolute addressing mode allows access to the entire 65 K bytes of addressable memory.

			Bytes: 3
		*/
		case 0x6D:
			/*
				ADC - Add Memory to Accumulator with Carry
			*/
			disassembledInstruction = fmt.Sprintf("ADC $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			ADC("absolute")
		case 0x2D:
			/*
				AND - "AND" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("AND $%02X%02X", operand2(), operand1())
			disassembleOpcode()

			AND("absolute")
		case 0x0E:
			/*
				ASL - Arithmetic Shift Left
			*/
			disassembledInstruction = fmt.Sprintf("ASL $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			ASL("absolute")
		case 0x2C:
			/*
				BIT - Test Bits in Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("BIT $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			BIT("absolute")
		case 0xCD:
			/*
				CMP - Compare Memory and Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("CMP $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			CMP("absolute")
		case 0xEC:
			/*
				CPX - Compare Index Register X To Memory
			*/
			disassembledInstruction = fmt.Sprintf("CPX $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			CPX("absolute")
		case 0xCC:
			/*
				CPY - Compare Index Register Y To Memory
			*/
			disassembledInstruction = fmt.Sprintf("CPY $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			CPY("absolute")
		case 0xCE:
			/*
				DEC - Decrement Memory By One
			*/
			disassembledInstruction = fmt.Sprintf("DEC $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			DEC("absolute")
		case 0x4D:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("EOR $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			EOR("absolute")
		case 0xEE:
			/*
				INC - Increment Memory By One
			*/
			disassembledInstruction = fmt.Sprintf("INC $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			INC("absolute")
		case 0x4C:
			/*
				JMP - JMP Absolute
			*/
			disassembledInstruction = fmt.Sprintf("JMP $%04X", int(operand2())<<8|int(operand1()))
			disassembleOpcode()
			// For AllSuiteA.bin 6502 opcode test suite
			if *allsuitea && memory[0x210] == 0xFF {
				fmt.Printf("\n\u001B[32;5mMemory address $210 == $%02X. All opcodes succesfully tested and passed!\u001B[0m\n", memory[0x210])
				os.Exit(0)
			}
			JMP("absolute")
		case 0x20:
			/*
				JSR - Jump To Subroutine
			*/
			disassembledInstruction = fmt.Sprintf("JSR $%04X", int(operand2())<<8|int(operand1()))
			disassembleOpcode()

			// First, push the high byte
			decSP()
			memory[SP] = byte((PC + 1) >> 8)

			// Then, push the low byte
			decSP()
			memory[SP] = byte((PC + 1) & 0xFF)

			// Now, jump to the subroutine address specified by the operands
			PC = int(operand2())<<8 | int(operand1())
			incCount(0)
		case 0xAD:
			/*
				LDA - Load Accumulator with Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDA $%04X", uint16(operand2())<<8|uint16(operand1()))
			disassembleOpcode()
			LDA("absolute")
		case 0xAE:
			/*
				LDX - Load Index Register X From Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDX $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			LDX("absolute")
		case 0xAC:
			/*
				LDY - Load Index Register Y From Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDY $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			LDY("absolute")
		case 0x4E:
			/*
				LSR - Logical Shift Right
			*/
			disassembledInstruction = fmt.Sprintf("LSR $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			LSR("absolute")
		case 0x0D:
			/*
				ORA - "OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("ORA $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			ORA("absolute")
		case 0x2E:
			/*
				ROL - Rotate Left
			*/
			disassembledInstruction = fmt.Sprintf("ROL $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			ROL("absolute")
		case 0x6E:
			/*
				ROR - Rotate Right
			*/
			disassembledInstruction = fmt.Sprintf("ROR $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			ROR("absolute")
		case 0xED:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
			*/
			disassembledInstruction = fmt.Sprintf("SBC $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			SBC("absolute")
		case 0x8D:
			/*
				STA - Store Accumulator in Memory
			*/
			disassembledInstruction = fmt.Sprintf("STA $%04X", uint16(operand2())<<8|uint16(operand1()))
			disassembleOpcode()
			STA("absolute")
		case 0x8E:
			/*
				STX - Store Index Register X In Memory
			*/
			disassembledInstruction = fmt.Sprintf("STX $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			STX("absolute")
		case 0x8C:
			/*
				STY - Store Index Register Y In Memory
			*/
			disassembledInstruction = fmt.Sprintf("STY $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			STY("absolute")

		// X Indexed Absolute addressing mode instructions
		/*
			$nnnn,X

			This form of addressing is used in conjunction with the X index register. The effective address is formed by adding the contents of X to the address contained in the second and third bytes of the instruction. This mode allows the index register to contain the index or count value and the instruction to contain the base address. This type of indexing allows any location referencing and the index to modify multiple fields resulting in reduced coding and execution time.

			Note on the MOS 6502:

			The value at the specified address, ignoring the the addressing mode's X offset, is read (and discarded) before the final address is read. This may cause side effects in I/O registers.


			Bytes: 3
		*/
		case 0x7D:
			/*
				ADC - Add Memory to Accumulator with Carry
			*/
			disassembledInstruction = fmt.Sprintf("ADC $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			ADC("absolutex")
		case 0x3D:
			/*
				AND - "AND" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("AND $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			AND("absolutex")
		case 0x1E:
			/*
				ASL - Arithmetic Shift Left
			*/
			disassembledInstruction = fmt.Sprintf("ASL $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			ASL("absolutex")
		case 0xDD:
			/*
				CMP - Compare Memory and Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("CMP $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			CMP("absolutex")
		case 0xDE:
			/*
				DEC - Decrement Memory By One
			*/
			disassembledInstruction = fmt.Sprintf("DEC $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()

			DEC("absolutex")
		case 0x5D:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("EOR $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			EOR("absolutex")
		case 0xFE:
			/*
				INC - Increment Memory By One
			*/
			disassembledInstruction = fmt.Sprintf("INC $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			INC("absolutex")
		case 0xBD:
			/*
				LDA - Load Accumulator with Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDA $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			LDA("absolutex")
		case 0xBC:
			/*
				LDY - Load Index Register Y From Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDY $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			LDY("absolutex")
		case 0x5E:
			/*
				LSR - Logical Shift Right
			*/
			disassembledInstruction = fmt.Sprintf("LSR $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			LSR("absolutex")
		case 0x1D:
			/*
				ORA - "OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("ORA $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			ORA("absolutex")
		case 0x3E:
			/*
			 */
			disassembledInstruction = fmt.Sprintf("ROL $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			ROL("absolutex")
		case 0x7E:
			/*
				ROR - Rotate Right
			*/
			disassembledInstruction = fmt.Sprintf("ROR $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			ROR("absolutex")
		case 0xFD:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
			*/
			disassembledInstruction = fmt.Sprintf("SBC $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			SBC("absolutex")
		case 0x9D:
			/*
				STA - Store Accumulator in Memory
			*/
			disassembledInstruction = fmt.Sprintf("STA $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			STA("absolutex")

		// Y Indexed Absolute addressing mode instructions
		/*
			$nnnn,Y

			This form of addressing is used in conjunction with the Y index register. The effective address is formed by adding the contents of Y to the address contained in the second and third bytes of the instruction. This mode allows the index register to contain the index or count value and the instruction to contain the base address. This type of indexing allows any location referencing and the index to modify multiple fields resulting in reduced coding and execution time.

			Note on the MOS 6502:

			The value at the specified address, ignoring the the addressing mode's Y offset, is read (and discarded) before the final address is read. This may cause side effects in I/O registers.

			Bytes: 3
		*/
		case 0x79:
			/*
				ADC - Add Memory to Accumulator with Carry
			*/
			disassembledInstruction = fmt.Sprintf("ADC $%04X,Y", int(operand2())<<8|int(operand1()))
			disassembleOpcode()
			ADC("absolutey")
		case 0x39:
			/*
				AND - "AND" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("AND $%02X%02X,Y", operand2(), operand1())
			disassembleOpcode()
			AND("absolutey")
		case 0xD9:
			/*
				CMP - Compare Memory and Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("CMP $%02X%02X,Y", operand2(), operand1())
			disassembleOpcode()
			CMP("absolutey")
		case 0x59:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("EOR $%02X%02X,Y", operand2(), operand1())
			disassembleOpcode()
			EOR("absolutey")
		case 0xB9:
			/*
				LDA - Load Accumulator with Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDA $%02X%02X,Y", operand2(), operand1())
			disassembleOpcode()
			LDA("absolutey")
		case 0xBE:
			/*
				LDX - Load Index Register X From Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDX $%02X%02X,Y", operand2(), operand1())
			disassembleOpcode()
			LDX("absolutey")
		case 0x19:
			/*
				ORA - "OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("ORA $%02X%02X,Y", operand2(), operand1())
			disassembleOpcode()
			ORA("absolutey")
		case 0xF9:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
			*/
			disassembledInstruction = fmt.Sprintf("SBC $%02X%02X,Y", operand2(), operand1())
			disassembleOpcode()
			SBC("absolutey")
		case 0x99:
			/*
				STA - Store Accumulator in Memory
			*/
			disassembledInstruction = fmt.Sprintf("STA $%02X%02X,Y", operand2(), operand1())
			disassembleOpcode()
			STA("absolutey")
		// Absolute Indirect addressing mode instructions
		case 0x6C:
			/*
				JMP - JMP Indirect
			*/
			disassembledInstruction = fmt.Sprintf("JMP ($%02X%02X)", operand2(), operand1())
			disassembleOpcode()
			JMP("indirect")
		}
		if *plus4 {
			kernalRoutines()
		}
	}
}
