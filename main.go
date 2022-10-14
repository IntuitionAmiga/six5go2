package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

var (
	printHex           bool
	file               []byte
	bytecounter        int // Byte position counter
	machineMonitor         = false
	disassemble            = false
	once                   = true
	instructionCounter     = 0

	// CPURegisters and RAM
	A      byte        = 0x0 // Accumulator
	X      byte        = 0x0 // X register
	Y      byte        = 0x0 // Y register		(76543210) SR Bit 5 is always set
	SR     byte              // Status Register	(NVEBDIZC)
	SP     uint              // Stack Pointer
	PC     int               // Program Counter
	memory [65536]byte       // Memory
)

const (
	ACCUMULATOR = "accumulator"
	IMMEDIATE   = "immediate"
	ZEROPAGE    = "zeropage"
	ZEROPAGEX   = "zeropagex"
	ZEROPAGEY   = "zeropagey"
	ABSOLUTE    = "absolute"
	ABSOLUTEX   = "absolutex"
	ABSOLUTEY   = "absolutey"
	INDIRECT    = "indirect"
	INDIRECTX   = "indirectx"
	INDIRECTY   = "indirecty"
)

func main() {
	fmt.Printf("Six5go2 - 6502 Emulator and Disassembler in Golang (c) 2022 Zayn Otley\n\n")

	if len(os.Args) <= 2 {
		instructions()
		os.Exit(0)
	}
	if len(os.Args) > 2 {
		parseUint, _ := strconv.ParseUint(os.Args[2], 16, 16)
		PC = int(parseUint)
	}
	if len(os.Args) > 3 && os.Args[3] == "dis" {
		disassemble = true
	}
	if len(os.Args) > 3 && os.Args[3] == "mon" {
		machineMonitor = true
	}
	if len(os.Args) > 4 && os.Args[4] == "hex" {
		printHex = true
	}

	//  Read file
	file, _ = os.ReadFile(os.Args[1])
	fmt.Printf("Length of file %s is %v ($%04X) bytes\n\n", os.Args[1], len(file), len(file))

	fmt.Printf("Size of addressable memory is %v ($%04X) bytes\n\n", len(memory), len(memory))

	// Copy file into memory at PC
	fmt.Printf("Copying file into memory at $%04X\n\n", PC)
	copy(memory[PC:], file)

	// Start emulation
	fmt.Printf("Starting emulation at $%04X\n\n", PC)
	reset()
	printMachineState()
	execute()
}
func instructions() {
	fmt.Printf("USAGE   - %s <target_filename> <hex_entry_point> <dis>/<mon> (Disassembler/Machine Monitor) <hex> (Hex opcodes as comments with disassembly)\n\n", os.Args[0])
	fmt.Printf("EXAMPLE - %s AllSuiteA.bin 4000 mon\n\n", os.Args[0])
	fmt.Printf("EXAMPLE - %s AllSuiteA.bin 4000 dis\n\n", os.Args[0])
	fmt.Printf("EXAMPLE - %s AllSuiteA.bin 4000 dis hex\n\n", os.Args[0])
}
func opcode() byte {
	return memory[bytecounter]
}
func operand1() byte {
	return memory[bytecounter+1]
}
func operand2() byte {
	return memory[bytecounter+2]
}
func incCount(amount int) {
	if bytecounter+amount < len(file)-1 && amount != 0 {
		bytecounter += amount
	}
	PC += amount
	printMachineState()
}
func getTermDim() (width, height int, err error) {
	var termDim [4]uint16
	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(0), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&termDim)), 0, 0, 0); err != 0 {
		return -1, -1, err
	}
	return int(termDim[1]), int(termDim[0]), nil
}

func printMachineState() {
	if machineMonitor {
		// fmt.Print("\033[H\033[2J") // ANSI escape code hack to clear the screen
		// Clear the screen once
		if once {
			fmt.Printf("\033[2J")
		}
		once = false
		// Move cursor to top left
		fmt.Printf("\033[0;0H")
	}

	fmt.Printf(";; PC=$%04X A=$%02X X=$%02X Y=$%02X SR=%08b (NVEBDIZC) SP=$%04X\n\n", PC, A, X, Y, SR, SP)

	if machineMonitor {
		// Get terminal width and height
		width, height, _ := getTermDim()
		fmt.Printf("RAM dump $0000 - $%04X:\n\n", (height-5)*(width/4+8))

		for i := 0; i < height-5; i++ {
			for j := 0; j < width/4+8; j++ {
				if memory[i*32+j] == 0 {
					fmt.Printf("\u001B[37m %02X", memory[i*32+j])
				} else {
					fmt.Printf("\u001B[3%dm %02X", (memory[i*32+j])%7+1, memory[i*32+j])
				}
			}
			fmt.Printf("\n")
		}
		time.Sleep(50 * time.Millisecond)
	}
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

/*
	func unsetBreakFlag() {
		setSRBitOff(4)
	}
*/
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

// 6502 mnemonics with multiple addressing modes
func LDA(addressingMode string) {
	switch addressingMode {
	case IMMEDIATE: // Immediate
		A = operand1()
		incCount(2)
	case ZEROPAGE: // Zero Page
		// Get address
		address := operand1()
		// Get value from memory at address
		value := memory[address]
		// Set accumulator to value
		A = value
		incCount(2)
	case ZEROPAGEX: // Zero Page, X
		// Get address
		address := operand1() + X
		value := memory[address]
		// Set accumulator to value
		A = value
		incCount(2)
	case ABSOLUTE: // Absolute
		// Get 16 bit address from operand 1 and operand 2
		address := int(operand2())<<8 | int(operand1())
		value := memory[address]
		// Set accumulator to value
		A = value
		incCount(3)
	case ABSOLUTEX: // Absolute, X
		// Get the 16bit X indexed absolute memory address
		address := int(operand2())<<8 | int(operand1()) + int(X)
		value := memory[address]
		// Set accumulator to value
		A = value
		incCount(3)
	case ABSOLUTEY: // Absolute, Y
		// Get 16 bit address from operand 1 and operand 2
		address := int(operand2())<<8 | int(operand1()) + int(Y)
		value := memory[address]
		// Set accumulator to value
		A = value
		incCount(3)
	case INDIRECTX: // Indirect, X
		// Get the 16bit X indexed zero page indirect address
		indirectAddress := uint16(int(operand1()) + int(X)&0xFF)
		// Get the value at the indirect address
		indirectValue := memory[indirectAddress]
		// Get the value at the indirect address + 1
		indirectValue2 := memory[(indirectAddress + 1)]
		// Combine the two values to get the address
		indirectAddress = uint16(int(indirectValue) + int(indirectValue2)<<8)
		// Get the value at the address
		value := memory[indirectAddress]
		// Set the accumulator to the value
		A = value
		incCount(2)
	case INDIRECTY: // Indirect, Y
		// Get address
		address := memory[operand1()]
		// Get the value at the address
		value := memory[address+Y]
		// Set the accumulator to the value
		A = value
		incCount(2)
	}
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
func LDX(addressingMode string) {
	switch addressingMode {
	case IMMEDIATE: // Immediate
		// Load the value of the operand1() into the X register.
		X = operand1()
		incCount(2)
	case ZEROPAGE: // Zero Page
		// Get address
		address := operand1()
		value := memory[address]
		// Load the value at the address into X
		X = value
		incCount(2)
	case ZEROPAGEY: // Zero Page, Y
		// Get Y indexed Zero Page address
		address := operand1() + Y
		value := memory[address]
		// Load the X register with the Y indexed value in the operand
		X = value
		incCount(2)
	case ABSOLUTE: // Absolute
		// Get 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		value := memory[address]
		// Update X with the value stored at the address in the operands
		X = value
		incCount(3)
	case ABSOLUTEY: // Absolute, Y
		// Get 16 bit Y indexed address from operands
		address := int(operand2())<<8 | int(operand1()) + int(Y)
		value := memory[address]
		X = value
		incCount(3)
	}
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
func LDY(addressingMode string) {
	switch addressingMode {
	case IMMEDIATE: // Immediate
		// Load the value of the operand1() into the Y register.
		Y = operand1()
		incCount(2)
	case ZEROPAGE: // Zero Page
		// Get address
		address := operand1()
		value := memory[address]
		// Load the value at the address into Y
		Y = value
		incCount(2)
	case ZEROPAGEX: // Zero Page, X
		// Get the X indexed address
		address := operand1() + X
		value := memory[address]
		// Load the Y register with the X indexed value in the operand
		Y = value
		incCount(2)
	case ABSOLUTE: // Absolute
		// Get 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		value := memory[address]
		// Update Y with the value stored at the address in the operands
		Y = value
		incCount(3)
	case ABSOLUTEX: // Absolute, X
		// Get the 16bit X indexed absolute memory address
		address := int(operand2())<<8 | int(operand1()) + int(X)
		value := memory[address]
		// Update Y with the value stored at the address
		Y = value
		incCount(3)
	}
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
func STA(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE: // Zero Page
		// Get address from operand1()
		address := operand1()
		// Store contents of Accumulator in memory
		memory[address] = A
		incCount(2)
	case ZEROPAGEX: // Zero Page, X
		// Get the X Indexed Zero Page address
		address := operand1() + X
		// Store contents of Accumulator in X indexed memory
		memory[address] = A
		incCount(2)
	case ABSOLUTE: // Absolute
		// Get 16 bit absolute address from operand 1 and operand 2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Update the memory at the address stored in operand 1 and operand 2 with the value of the accumulator
		memory[address] = A
		incCount(3)
	case ABSOLUTEX: // Absolute, X
		// Get 16 bit X indexed absolute memory address
		address := int(operand2())<<8 | int(operand1()) + int(X)
		memory[address] = A
		incCount(3)
	case ABSOLUTEY: // Absolute, Y
		// Get 16bit absolute address
		address := uint16(operand2())<<8 | uint16(operand1())
		// Update the memory at the Y indexed address stored in operand 1 and operand 2 with the value of the accumulator
		memory[int(address)+int(Y)] = A
		incCount(3)
	case INDIRECTX: // Indirect, X
		// Get the 16bit X indexed zero page indirect address
		indirectAddress := uint16(int(operand1()) + int(X)&0xFF)
		// Get the value at the indirect address
		indirectValue := memory[indirectAddress]
		// Get the value at the indirect address + 1
		indirectValue2 := memory[(indirectAddress + 1)]
		// Combine the two values to get the address
		indirectAddress = uint16(int(indirectValue) + int(indirectValue2)<<8)
		// Set the value at the address to the value of A
		memory[indirectAddress] = A
		incCount(2)
	case INDIRECTY: // Indirect, Y
		// Get address
		address := memory[operand1()]
		// Load accumulator with address+Y index value
		memory[address+Y] = A
		incCount(2)
	}
}
func STX(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE: // Zero Page
		// Get address from operand1()
		address := operand1()
		// Store contents of X register in memory address at operand1()
		memory[address] = X
		incCount(2)
	case ZEROPAGEY: // Zero Page, Y
		// Get Y indexed Zero Page address
		address := operand1() + Y
		// Store contents of X register in Y indexed memory address
		memory[address] = X
		incCount(2)
	case ABSOLUTE: // Absolute
		// Get the 16 bit address from operand 1 and operand 2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Update the memory at the address stored in operand 1 and operand 2 with the value of the X register
		memory[address] = X
		incCount(3)
	}
}
func STY(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE: // Zero Page
		// Get address
		address := operand1()
		// Store Y register in memory at address in operand1()
		memory[address] = Y
		incCount(2)
	case ZEROPAGEX: // Zero Page, X
		// Get X indexed Zero Page address
		address := operand1() + X
		// Store contents of Y register in X indexed memory address
		memory[address] = Y
		incCount(2)
	case ABSOLUTE: // Absolute
		// Get the 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		// Update the memory at the address stored in operand 1 and operand 2 with the value of the Y register
		memory[address] = Y
		incCount(3)
	}
}
func CMP(addressingMode string) {
	var value, result byte
	switch addressingMode {
	case IMMEDIATE: // Immediate
		// Get value from operand1()
		value = operand1()
		// Subtract the value from the accumulator
		result = A - value
		incCount(2)
	case ZEROPAGE: // Zero Page
		// Get address
		address := operand1()
		// Subtract the operand from the accumulator
		value = A - memory[address]
		// Subtract the value from the accumulator
		result = A - value
		incCount(2)
	case ZEROPAGEX: // Zero Page, X
		// Get address
		address := operand1() + X
		// Get value at address
		value = memory[address]
		// Subtract the value from the accumulator
		result = A - value
		incCount(2)
	case ABSOLUTE: // Absolute
		// Get 16bit absolute address
		address := int(operand2())<<8 | int(operand1())
		// Get the value at the address
		value = memory[address]
		// Subtract the value from the accumulator
		result = value - A
		incCount(3)
	case ABSOLUTEX: // Absolute, X
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Get value at address
		value = memory[address]
		// Subtract the value from the accumulator
		result = A - value
		incCount(3)
	case ABSOLUTEY: // Absolute, Y
		// Get address
		address := int(operand2())<<8 | int(operand1()) + int(Y)
		// Get the value at the address
		value = memory[address]
		// Subtract the value from the accumulator
		result = A - value
		incCount(3)
	case INDIRECTX: // Indirect, X
		// Get the address of the operand
		address := int(operand1()) + int(X)
		// Get the value of the operand
		value = memory[address]
		// Subtract the operand from the accumulator
		result = A - value
		incCount(2)
	case INDIRECTY: // Indirect, Y
		// Get address from operand1() and add Y to it
		address := memory[operand1()] + Y
		// Get value at address
		value = memory[address]
		// Subtract the operand from the accumulator
		result = A - value
		incCount(2)
	}
	// If A is equal to the value, set the zero flag
	if A == value {
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
}
func JMP(addressingMode string) {
	switch addressingMode {
	case ABSOLUTE:
		// Get the address from the operand
		address := int(operand2())<<8 | int(operand1())
		// Set the program counter to the address
		PC = address
		incCount(3)
	case INDIRECT:
		// Get the address from the operand
		address := int(operand2())<<8 | int(operand1())
		// Get the value at the address
		value := memory[address]
		// Set the program counter to the value
		PC = int(value)
		incCount(3)
	}
}
func AND(addressingMode string) {
	var value, result byte
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = operand1()
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
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
		incCount(2)
	}
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
func EOR(addressingMode string) {
	switch addressingMode {
	case IMMEDIATE:
		// XOR the accumulator with the operand
		A ^= operand1()
		// If accumulator is 0the  set Zero flag to 1 else set Zero flag to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If bit 7 of accumulator is set, set SR Negative flag to 1
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(2)
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Perform EOR operation
		A ^= memory[address]
		// If A==0, set zero flag
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If accumulator bit 7 is 1, set negative SR flag
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(2)
	case ZEROPAGEX:
		// Get address
		address := uint16(operand1()) + uint16(X)
		// Get value at address
		value := memory[address]
		// Perform EOR
		A ^= value
		// If accumulator is 0 then set Zero flag to 1 else set Zero flag to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If bit 7 of accumulator is set, set SR Negative flag to 1
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(2)
	case ABSOLUTE:
		// Get 16bit address from operand 1 and operand 2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Perform EOR operation
		A ^= memory[address]
		// If A==0 then set SR Zero flag bit 1 to 1
		if A == 0 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(3)
	case ABSOLUTEX:
		// Get address of X indexed memory location
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Perform EOR operation on A and the value at the X indexed memory location
		A ^= memory[address]
		if A == 0 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		// If bit 7 of A is 1 then set SR Negative flag bit 7 to 1
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(3)
	case ABSOLUTEY:
		// Get the address
		address := int(operand2())<<8 | int(operand1()) + int(Y)
		// EOR the accumulator with the memory at the address
		A ^= memory[address]
		// If A==0 then set SR Zero flag bit 1 to 1
		if A == 0 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(3)
	case INDIRECTX:
		// Get the X indexed zero page indirect address
		indirectAddress := uint16(int(operand1())) + uint16(X)
		// Get the value at the indirect address
		indirectValue := memory[indirectAddress]
		// Get the value at the indirect address + 1
		indirectValue2 := memory[(indirectAddress + 1)]
		// Combine the two values to get the address
		indirectAddress = uint16(int(indirectValue) + int(indirectValue2)<<8)
		// Get the value at the address
		value := memory[indirectAddress]
		// XOR the value with the accumulator
		A ^= value
		// If A==0, set zero flag
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If A bit 7 is 1, set negative flag
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(2)
	case INDIRECTY:
		// Get the 16bit address
		address := uint16(int(operand1()))
		// Get the indirect address
		indirectAddress1 := memory[address]
		indirectAddress2 := memory[address+1]
		indirectAddress := uint16(int(indirectAddress1)+int(indirectAddress2)<<8) + uint16(Y)
		// Get the value at the address
		value := memory[indirectAddress]
		// AND the value with the accumulator
		A ^= value
		// If accumulator is 0 then set Zero flag to 1 else set Zero flag to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If bit 7 of accumulator is set, set SR Negative flag to 1
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(2)
	}
}
func ORA(addressingMode string) {
	switch addressingMode {
	case IMMEDIATE:
		// OR the accumulator with the operand
		A |= operand1()
		// If A==0 then set SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If A bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(2)
	case ZEROPAGE:
		// OR the accumulator with the memory value at the address in the operand
		A |= memory[operand1()]
		// If the accumulator is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If accumulator bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(2)
	case ZEROPAGEX:
		// Get address
		address := operand1() + X
		// Get the value at the address
		value := memory[address]
		// OR the value with the accumulator
		A |= value
		// If A==0 then set SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If A bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(2)
	case ABSOLUTE:
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value from memory
		value := memory[address]
		// OR the accumulator with the value stored at the address
		A |= value
		// If accumulator is 0 then set SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If accumulator bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(3)
	case ABSOLUTEX:
		// Get address
		address := int(operand2())<<8 | int(operand1()) + int(X)
		// Update A with the value stored at the address
		A |= memory[address]
		// If accumulator is 0 then set SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If accumulator bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(3)
	case ABSOLUTEY:
		// Get 16 bit Y indexed address from operands
		address := int(operand2())<<8 | int(operand1()) + int(Y)
		// OR the value stored at the address with the accumulator
		A |= memory[address]
		// If accumulator is 0 then set SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If accumulator bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(3)
	case INDIRECTX:
		// Get the X indexed zero page indirect address
		indirectAddress := uint16(int(operand1()) + int(X)&0xFF)
		// Get the value at the indirect address
		indirectValue := memory[indirectAddress]
		// Get the value at the indirect address + 1
		indirectValue2 := memory[(indirectAddress + 1)]
		// Combine the two values to get the address
		indirectAddress = uint16(int(indirectValue) + int(indirectValue2)<<8)
		// Get the value at the address
		value := memory[indirectAddress]
		// OR the value with the accumulator
		A |= value
		// If A==0 set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If A bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(2)
	case INDIRECTY:
		// Get the 16bit address
		address := uint16(int(operand1()))
		// Get the indirect address
		indirectAddress1 := memory[address]
		indirectAddress2 := memory[address+1]
		indirectAddress := uint16(int(indirectAddress1)+int(indirectAddress2)<<8) + uint16(Y)
		// Get the value at the address
		value := memory[indirectAddress]
		// AND the value with the accumulator
		A |= value

		// If A==0 then set SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If A bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(2)
	}
}
func BIT(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get value at address
		value := memory[address]
		// Store result of AND between A and memory stored at location in operand in a result variable
		result := A & value
		// If bit 7 of result is set then set SR negative value to 1 else set it to 0
		if readBit(7, result) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If bit 6 of result is set then set SR overflow flag to 1 else set it to 0
		if readBit(6, result) == 1 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// If result is 0 then set SR zero flag to 1 else set it to 0
		if result == 0 {
			setZeroFlag()
		}
		incCount(2)
	case ABSOLUTE:
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value := memory[address]
		// Store the result of the AND between the accumulator and the operands in a result var
		result := A & value
		// Set the SR Negative flag bit 7 to the value of bit 7 of result
		if readBit(7, result) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// Set the SR Overflow flag bit 6 to the value of bit 6 of result
		if readBit(6, result) == 1 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// If result==0 then set the SR Zero flag bit 1 to the result of result else set SR negative flag to 0
		if result == 0 {
			// If bit 7 of result is 1 then set SR negative flag to 1
			if readBit(7, result) == 1 {
				setNegativeFlag()
			}
		} else {
			unsetNegativeFlag()
		}
		incCount(3)
	}
}
func INC(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get value at address
		value := memory[address]
		// Increment value
		value++
		// Store value at address
		memory[address] = value
		// If value==0 then set SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
		if value == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If value bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
		if readBit(7, value) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(2)
	case ZEROPAGEX:
		// Get address
		address := uint16(int(operand1()) + int(X)&0xFF)
		// Get value at address
		value := memory[address]
		// Increment value
		value++
		// Store value at address
		memory[address] = value
		// If value==0 then set SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
		if value == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If value bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
		if readBit(7, value) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(2)
	case ABSOLUTE:
		// Get 16bit address from operand 1 and operand 2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Increment the value stored in memory at the address stored in operand 1 and operand 2
		memory[address]++
		// If bit 7 of the value stored in memory is 1 then set SR negative bit 7 to 1 else set SR negative bit 7 to 0
		if readBit(7, memory[address]) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If the value stored in memory is 0 then set SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
		if memory[address] == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		incCount(3)
	case ABSOLUTEX:
		// Get 16 bit address of X indexed memory location
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Increment the value at the X indexed memory location
		memory[address]++
		// If bit 7 of the value in memory is set then set SR negative bit 7 to 1 else set SR negative bit 7 to 0
		if readBit(7, memory[address]) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If the value in memory is 0 then set SR zero bit 1 to 1 else set SR zero bit 1 to 0
		if memory[address] == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		incCount(3)
	}
}
func DEC(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get value at address
		value := memory[address]
		// Decrement value
		value--
		// Store value at address
		memory[address] = value
		// If value==0 then set SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
		if value == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If value bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
		if readBit(7, value) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(2)
	case ZEROPAGEX:
		// Get address
		address := uint16(int(operand1()) + int(X)&0xFF)
		// Get value at address
		value := memory[address]
		// Decrement value
		value--
		// Store value at address
		memory[address] = value
		// If value==0 then set SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
		if value == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If value bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
		if readBit(7, value) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(2)
	case ABSOLUTE:
		// Get 16 bit address
		address := uint16(operand2())<<8 | uint16(operand1())
		// Decrement value stored at address by 1
		memory[address]--
		// If bit 7 of the value stored in memory is 1 then set SR negative bit 7 to 1 else set SR negative bit 7 to 0
		if readBit(7, memory[address]) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If the value stored in memory is 0 then set SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
		if memory[address] == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		incCount(3)
	case ABSOLUTEX:
		// Get 16bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Decrement the value at the address by 1
		memory[address]--
		// If bit 7 of memory is set then set SR negative bit 7 to 1 else set SR negative bit 7 to 0
		if readBit(7, memory[address]) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If result is 0 set SR Zero flag bit 1 to 1 else set SR Zero flag bit 1 to 0
		if memory[address] == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		incCount(3)
	}
}
func ADC(addressingMode string) {
	switch addressingMode {
	case IMMEDIATE:
		// Get value from operand
		value := operand1()
		// If A+memory > 255, set SR carry flag
		if int(A)+int(value) > 255 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		// If bit 7 of A+memory is different from bit 7 of A, set SR overflow flag bit 6
		if readBit(7, A) != readBit(7, value) {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// Update A with A+memory
		A += value
		// If bit 7 of A is set then set SR negative flag bit 7 else clear it
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If A=0 then set SR zero flag bit 1 else clear it
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		incCount(2)
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get value from memory
		value := memory[address]
		// If A+memory > 255, set SR carry flag
		if int(A)+int(value) > 255 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		// If bit 7 of A+memory is different from bit 7 of A, set SR overflow flag bit 6
		if readBit(7, A) != readBit(7, value) {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// Update A with a+memory
		A += value
		// If bit 7 of A is set then set SR negative flag bit 7 else clear it
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If A=0 then set SR zero flag bit 1 else clear it
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		incCount(2)
	case ZEROPAGEX:
		// Get address
		address := uint16(int(operand1()) + int(X)&0xFF)
		// Get value from memory
		value := memory[address]
		// Add result to accumulator
		A += value
		// If result>255 then set SR carry flag bit 0 to 1
		if A > 0x00FF {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		// If result>127 or result<128 then set SR overflow flag bit 6 to 1 else set SR overflow flag bit 6 to 0
		if A > 127 || A < 128 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// If accumulator bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If accumulator is 0 then set SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		incCount(2)
	case ABSOLUTE:
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value from memory
		value := memory[address]
		// Read bit 7 of the accumulator into a result var
		temp := getABit(7)
		// If SR carry bit is set, add 1 to A
		if getSRBit(0) == 1 {
			A++
		}
		// Add the value of operand
		A += value
		// If bit 7 of accumulator is not equal to bit 7 of result then set SR overflow flag bit 6 to 1
		if getABit(7) != (temp & 0b10000000) {
			setOverflowFlag()
		}
		// If bit 7 of accumulator is 1 then set negative flag
		if getABit(7) == 1 {
			setNegativeFlag()
		}
		// If accumulator is 0 then set zero flag else set SR zero flag to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		incCount(3)
	case ABSOLUTEX:
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Get value from memory
		value := memory[address]

		// Add the X indexed address from the operands plus the carry bit to the accumulator
		A += value
		// If the carry flag is 1 then add 1 to the accumulator else reset it
		if getSRBit(0) == 1 {
			A++
		} else {
			unsetCarryFlag()
		}
		// If the accumulator is greater than 255 then set the carry flag else reset it
		if value > 0x00FF {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		// If the accumulator is greater than 99 then set the decimal flag else reset it
		if A > 99 {
			setDecimalFlag()
		} else {
			unsetDecimalFlag()
		}
		// If the accumulator is greater than 127 then set the overflow flag else reset it
		if A > 127 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// If the accumulator is less than 0 then set the overflow flag else reset it
		// If A bit 7 is unset then it's negative
		if getABit(7) == 0 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// If the accumulator is greater than 127 then set the negative flag else reset it
		if A > 127 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If the accumulator is less than 0 then set the negative flag else reset it
		// If A bit 7 is unset then it's negative
		if getABit(7) == 0 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If the accumulator is 0 then set the zero flag else reset it
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		incCount(3)
	case ABSOLUTEY:
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(Y)
		// Get value from memory
		value := memory[address]

		// Add the Y indexed address from the operands plus the carry bit to the accumulator
		A += value
		// If the carry flag is 1 then add 1 to the accumulator else reset it
		if getSRBit(0) == 1 {
			A++
		} else {
			unsetCarryFlag()
		}
		// If the accumulator is greater than 255 then set the carry flag else reset it
		// 256 rolls over to 0
		// If the accumulator is 0 then set the zero flag else reset it
		if A == 255 {
			setCarryFlag()
			setZeroFlag()
		} else {
			unsetCarryFlag()
			unsetZeroFlag()
		}
		// If the accumulator is greater than 99 then set the decimal flag else reset it
		if A > 99 {
			setDecimalFlag()
		} else {
			unsetDecimalFlag()
		}
		// If the accumulator is greater than 127 then set the overflow flag else reset it
		if A > 127 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// If the accumulator is less than 0 then set the overflow flag else reset it
		// If accumulator bit 7 is set then set SR bit 0 to 0 as number is negative
		// If the accumulator is less than 0 then set the negative flag else reset it
		if getABit(7) == 1 {
			setOverflowFlag()
			setNegativeFlag()
		} else {
			unsetOverflowFlag()
			unsetNegativeFlag()
		}
		// If the accumulator is greater than 127 then set the negative flag else reset it
		if A > 127 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		incCount(3)
	case INDIRECTX:
		// Get the X Indexed Zero Page Indirect address from the operand
		indirectAddress := int(operand1()) + int(X)
		// Get the address from the indirect address
		address := int(memory[indirectAddress]) + int(memory[indirectAddress+1])<<8
		// Get the value from the address
		value := memory[address]
		// Add the value to the accumulator
		A += value
		// If the carry flag is set, add 1 to the accumulator
		if getSRBit(0) == 1 {
			A++
		}
		// If the accumulator is 0, set the Zero flag to 1
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If bit 7 of accumulator is set, set SR Negative flag to 1
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If the accumulator is greater than 255, set the carry flag to 1
		if A+value > 255 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		// If the accumulator is greater than 127 or less than -128, set the overflow flag to 1
		if int(A) > 127 || int(A) < (-128) {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		incCount(2)
	case INDIRECTY:
		// Get address
		address := memory[operand1()]
		// Get value from memory
		value := memory[address+Y]
		// Load accumulator with address+Y index value
		A = value
		// If accumulator>255 then set SR carry flag bit 0 to 1
		if value > 0x00FF {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		// If accumulator>127 or accumulator<128 then set SR overflow flag bit 6 to 1 else set SR overflow flag bit 6 to 0
		if A > 127 || int(A) < 128 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// If accumulator bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If accumulator is 0 then set SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		incCount(2)
	}
}
func SBC(addressingMode string) {
	var value, result byte
	switch addressingMode {
	case IMMEDIATE:
		// Get value from operand
		value = operand1()
		// Store result of A-memory stored at operand1() in result variable
		result = A - value

		/*
			This instruction subtracts the value of memory and borrow from the value of the accumulator,
			using two's complement arithmetic, and stores the result in the accumulator.
			Borrow is defined as the carry flag complemented;
			therefore, a resultant carry flag indicates that a borrow has not occurred.
				Subtraction is performed by adding the two's complement of the value of memory to the accumulator.
				The carry flag is set if the result is greater than or equal to 0.
				The carry flag is reset when the result is less than 0, indicating a borrow.
				The overflow flag is set when the result exceeds +127 or -127, otherwise it is reset.
				The negative flag is set if the result in the accumulator has bit 7 on, otherwise it is reset.
				The Z flag is set if the result in the accumulator is 0, otherwise it is reset.
		*/

		// If

		// If result is greater than or equal to 0, set carry flag to 1
		// If result bit 7 is not set then set SR bit 0 to 1 as number is not negative
		if readBit(7, result) == 0 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		// If result is less than 0, set carry flag to 0
		// If result bit 7 is set then set SR bit 0 to 0 as number is negative
		// If bit 7 of result is set, set N flag to 1 else reset it
		if readBit(7, result) == 1 {
			unsetCarryFlag()
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If result is greater than 127 or less than -127, set overflow flag to 1
		if result > 127 || result == 0x80 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// If result is equal to 0, set Z flag to 1 else reset it
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// Set A to result
		A = result
		incCount(2)
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get value from memory
		value := memory[address]
		// Store result of A-memory stored at operand1() in result variable
		result := A - value
		// If result is greater than or equal to 0, set carry flag to 1
		// If result bit 7 is not set then set SR bit 0 to 1 as number is not negative
		if readBit(7, result) == 0 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		// If result is less than 0, set carry flag to 0
		// If result bit 7 is set then set SR bit 0 to 0 as number is negative
		// If bit 7 of result is set, set N flag to 1 else reset it
		if readBit(7, result) == 1 {
			unsetCarryFlag()
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If result is greater than 127 or less than -127, set overflow flag to 1
		if result > 127 || result == 0x80 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// If result is equal to 0, set Z flag to 1 else reset it
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// Set A to result
		A = result
		incCount(2)
	case ZEROPAGEX:
		// Get address
		address := operand1()
		// Get value from memory
		value := memory[address+X]
		// Store result of A-memory stored at operand1() in result variable
		result := A - value
		//result := A - memory[value] - (1 - SR&1)
		// If result is greater than or equal to 1 then set carry flag bit 0 to 1 else set carry flag bit 0 to 0
		// If result bit 7 is not set then set SR bit 0 to 1 as number is not negative
		if readBit(7, result) == 0 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		// If result is > 127 or < -127 then set overflow flag bit 6 to 1 else set overflow flag bit 6 to 0
		if result > 127 || result == 0x80 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// If result is < 0 then set Negative flag bit 7 to 1 else set Negative flag bit 7 to 0
		// If result bit 7 is set then set SR bit 0 to 0 as number is negative
		if readBit(7, result) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If result is 0 then set Z flag bit 1 to 1 else set Z flag bit 1 to 0
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// Update the accumulator
		A = result
		incCount(2)
	case ABSOLUTE:
		// Get address
		address := operand2()
		// Get value from memory
		value := memory[address]
		// Store result of A-memory stored at operand1() in result variable
		//result := A - memory[value] - (1 - SR&1)
		result := A - value - getSRBit(0)

		// If result is greater than or equal to 1 then set carry flag bit 0 to 1 else set carry flag bit 0 to 0
		if result >= 1 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		// If accumulator<0 then set SR carry bit 0 to 0 else set SR carry bit 0 to 1
		// If bit 7 of A is set then it is negative
		// If bit 7 of accumulator is set then set SR negative bit 7 to 1 else set SR negative bit 7 to 0
		if readBit(7, result) == 1 {
			unsetCarryFlag()
			setNegativeFlag()
		} else {
			setCarryFlag()
			unsetNegativeFlag()
		}
		// If accumulator > 127 or accumulator < -127 then set SR overflow bit 6 to 1 else set SR overflow bit 6 to 0
		if result > 127 || result == 0x80 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// If accumulator is 0 then set SR zero bit 1 to 1 else set SR zero bit 1 to 0
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// Update the accumulator
		A = result
		incCount(3)
	case ABSOLUTEX:
		// Get address
		address := operand2()
		// Get value from memory
		value := memory[address+X]
		// Subtract the value in the X indexed memory address from the accumulator with borrow
		// If result is greater than A then set SR carry bit 0 to 0 else set SR carry bit 0 to 1
		if value > A {
			unsetCarryFlag()
		} else {
			setCarryFlag()
		}
		// If result <0 then set carry bit to 0 indicating a borrow
		// If result bit 7 is set then set SR bit 0 to 0 as number is negative
		if readBit(7, value) == 1 {
			unsetCarryFlag()
		}
		// Subtract the value in memory from the accumulator with borrow
		A -= value - (1 - getSRBit(0))
		// If accumulator > 127 or accumulator < -127 then set SR overflow bit 6 to 1 else set SR overflow bit 6 to 0
		if A > 127 || A == 0x80 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// If bit 7 of accumulator is set then set SR negative bit 7 to 1 else set SR negative bit 7 to 0
		if getABit(7) == 1 {
		} else {
			unsetNegativeFlag()
		}
		// If accumulator is 0 then set SR zero bit 1 to 1 else set SR zero bit 1 to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		incCount(3)
	case ABSOLUTEY:
		// Get address
		address := operand2()
		// Get value from memory
		value := memory[address+Y]
		// Subtract the value in the Y indexed memory address from the accumulator with borrow
		// If result is greater than A then set SR carry bit 0 to 0 else set SR carry bit 0 to 1
		if value > A {
			unsetCarryFlag()
		} else {
			setCarryFlag()
		}
		// If result <0 then set carry bit to 0 indicating a borrow
		// If A bit 7 is unset then it's negative
		if readBit(7, A) == 1 {
			unsetCarryFlag()
		}
		// Subtract the value in memory from the accumulator with borrow
		A = A - value - (1 - getSRBit(0))
		// If accumulator > 127 or accumulator < -127 then set SR overflow bit 6 to 1 else set SR overflow bit 6 to 0
		// If accumulator bit 7 is set then set SR bit 0 to 0 as number is negative
		if A > 127 || A == 0x80 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// If bit 7 of accumulator is set then set SR negative bit 7 to 1 else set SR negative bit 7 to 0
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If accumulator is 0 then set SR zero bit 1 to 1 else set SR zero bit 1 to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		incCount(3)
	case INDIRECTX:
		// Get the value of the X Indexed Zero Page Address from operand
		indirectAddress := operand1() + X
		// Get the value of the memory location pointed to by the indirect address
		indirectValue := memory[indirectAddress]
		// Get the value of the memory location pointed to by the indirect value
		indirectValue2 := memory[indirectValue]
		// Combine the two values to get the final address
		finalAddress := uint16(indirectValue2) + uint16(indirectValue)<<8
		//  Get the value of the memory location pointed to by the final address
		finalValue := memory[finalAddress]
		// Subtract operand from A
		result := A - finalValue
		// If operand is greater than A, set carry flag to 0
		if result > A {
			unsetCarryFlag()
		}
		// If tempresult <0 Set the carry flag
		// If result bit 7 is set then set SR bit 0 to 0 as number is negative
		// If result bit 7 is not set then set SR bit 0 to 1 as number is not negative
		if readBit(7, result) == 1 {
			setCarryFlag()
			setNegativeFlag()
		} else {
			unsetCarryFlag()
			unsetNegativeFlag()
		}
		// If operand is equal to A, set Z flag to 1 else set Zero flag to 0
		if result == A {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If tempresult is greater than 127 or less than -127, set overflow flag to 1
		if result > 127 || result == 0x80 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// Set A to the result of the subtraction
		A = result
		incCount(2)
	case INDIRECTY:
		// Get address
		address := memory[operand1()]
		// Get value from memory
		value := memory[address+Y]
		// Load accumulator with address+Y index value
		A = value - (1 - SR&1)
		// Set carry flag bit 0 if result is greater than or equal to 1
		// If accumulator bit 7 is not set then set SR bit 0 to 1 as number is not negative
		if getABit(7) == 0 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		// Set overflow flag bit 6 if accumulator is greater than 127 or less than -127
		if A > 127 || A == 0x80 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// If accumulator bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
		if getABit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// Set Z flag bit 1 if accumulator is 0 else set Z flag bit 1 to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		incCount(2)
	}
}
func ROR(addressingMode string) {
	var value, result byte
	switch addressingMode {
	case ACCUMULATOR:
		// Get value from accumulator
		value = A
		// Rotate right one bit
		result = value >> 1
		// Store the result in the accumulator
		A = result
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
		incCount(2)
	case ZEROPAGEX:
		// Get X indexed zero page address
		address := operand1() + X
		// Get the value at the address
		value = memory[address]
		// Shift the value right 1 bit
		result = value >> 1
		// Store the value back into memory
		memory[address] = result
		incCount(2)
	case ABSOLUTE:
		// Get 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get the value stored at the address in the operands
		value = memory[address]
		// Shift the value right 1 bit
		result = value >> 1
		// Store the value back into memory
		memory[address] = result
		incCount(3)
	case ABSOLUTEX:
		// Get 16 bit address
		address := int(operand2())<<8 | int(operand1()) + int(X)
		// Get value stored at address
		value = memory[address]
		// Shift right the value by 1 bit
		result = value >> 1
		// Store value in memory at address in operand
		memory[address] = result
		incCount(3)
	}
	// Set bit 7 of result and negative flag to the carry flag
	if getSRBit(0) == 1 {
		result = result | 0x80
		setNegativeFlag()
	} else {
		result = result & 0x7F
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
}
func ROL(addressingMode string) {
	var value, result byte
	switch addressingMode {
	case ACCUMULATOR:
		// Get the value of the accumulator
		value = A
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | getSRBit(0)
		// Store the result in the accumulator
		A = result
		incCount(1)
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get the value at the address
		value = memory[address]
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | getSRBit(0)
		// Store the value back into memory
		memory[address] = result
		incCount(2)
	case ZEROPAGEX:
		// Get X indexed zero page address
		address := operand1() + X
		// Get the value at the address
		value = memory[address]
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | getSRBit(0)
		// Store the value back into memory
		memory[address] = result
		incCount(2)
	case ABSOLUTE:
		// Get 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get the value stored at the address in the operands
		value = memory[address]
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | getSRBit(0)
		// Store the value back into memory
		memory[address] = result
		incCount(3)
	case ABSOLUTEX:
		// Get 16bit X indexed absolute memory address
		address := int(operand2())<<8 | int(operand1()) + int(X)
		// Get the value stored at the address
		value = memory[address]
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | getSRBit(0)
		// Store the shifted value back in memory
		memory[address] = result
		incCount(3)
	}
	// Set SR carry flag to bit 7 of value
	if readBit(7, value) == 1 {
		setCarryFlag()
	} else {
		unsetCarryFlag()
	}
	// Set SR negative flag to bit 6 of value (bit 7 of result)
	if readBit(6, value) == 1 {
		setNegativeFlag()
	} else {
		unsetNegativeFlag()
	}
	// If result is 0 then set zero flag else reset it
	if result == 0 {
		setZeroFlag()
	} else {
		unsetZeroFlag()
	}
	//fmt.Printf("ROL: value = %04X, result = %04X\n", value, result)
}
func LSR(addressingMode string) {
	var value, result byte
	switch addressingMode {
	case ACCUMULATOR:
		// Get the value of the accumulator
		value = A
		// Shift the value right 1 bit
		result = value >> 1
		// Store the result back into the accumulator
		A = result
		incCount(1)
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get the value at the address
		value = memory[address]
		// Shift the value right 1 bit
		value >>= 1
		// Store the value back into memory
		memory[address] = value
		incCount(2)
	case ZEROPAGEX:
		// Get the X indexed address
		address := operand1() + X
		// Get the value at the X indexed address
		value = memory[address]
		// Shift the value right 1 bit
		value >>= 1
		// Store the shifted value in memory
		memory[address] = value
		incCount(2)
	case ABSOLUTE:
		// Get 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get the value stored at the address in the operands
		value = memory[address]
		// Shift the value right 1 bit
		value >>= 1
		// Store the shifted value back in memory
		memory[address] = value
		incCount(3)
	case ABSOLUTEX:
		// Get the 16bit X indexed absolute memory address
		address := int(operand2())<<8 | int(operand1()) + int(X)
		// Get the value stored at the address
		value = memory[address]
		// Shift the value right 1 bit
		value >>= 1
		// Store the shifted value back in memory
		memory[address] = value
		incCount(3)
	}
	// Reset the SR negative flag
	unsetNegativeFlag()
	// If result is 0 then set SR zero flag else reset it
	if result == 0 {
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
func ASL(addressingMode string) {
	var value, result byte
	switch addressingMode {
	case ACCUMULATOR:
		// Set value to accumulator
		value = A
		// Shift the value left 1 bit
		result = value << 1
		// Update the accumulator with the result
		A = result
		incCount(1)
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get value at address
		value = memory[address]
		// Shift value left 1 bit
		value <<= 1
		// Store value at address
		memory[address] = value
		incCount(2)
	case ZEROPAGEX:
		// Get address
		address := operand1() + X
		// Get value at address
		value := memory[address]
		// Shift value left 1 bit
		value <<= 1
		// Store value at address
		memory[address] = value
		incCount(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get the value at the address
		value = memory[address]
		// Shift the value left 1 bit
		value <<= 1
		// Store the value at the address
		memory[address] = value
		incCount(3)
	case ABSOLUTEX:
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Get value at address
		value = memory[address]
		// Shift left 1 bit
		value <<= 1
		// Store value back in memory
		memory[address] = value
		incCount(3)
	}
	// Set the SR Negative flag to the bit 7 of the result
	if readBit(7, result) == 1 {
		setNegativeFlag()
	} else {
		unsetNegativeFlag()
	}
	// If the result is 0, set the Zero flag to 1 else unset zero flag and set carry flag to bit 7 of value
	if result == 0 {
		setZeroFlag()
	} else {
		unsetZeroFlag()
		// Set the Carry flag to the bit 7 of input value
		if readBit(7, value) == 1 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
	}
}
func CPX(addressingMode string) {
	var value, result byte
	switch addressingMode {
	case IMMEDIATE:
		// Get value from operand1
		value = operand1()
		// Compare X with value
		result = X - value
		incCount(2)
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get value at address
		value = memory[address]
		// Store result of X-memory stored at operand1() in result variable
		result = X - value
		incCount(2)
	case ABSOLUTE:
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = memory[address]
		incCount(3)
	}
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
func CPY(addressingMode string) {
	var value, result byte
	switch addressingMode {
	case IMMEDIATE:
		// Get value from operand1
		value = operand1()
		// Subtract operand from Y
		result = Y - operand1()
		incCount(2)
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get value at address
		value = memory[address]
		// Store result of Y-memory stored at operand1() in result variable
		result = Y - value
		incCount(2)
	case ABSOLUTE:
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = memory[address]
		incCount(3)
	}
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

func reset() {
	// Set PC to 0xFFFC
	PC = int(uint16(memory[0xFFFD])<<8 | uint16(memory[0xFFFC]))
	// Set SP to 0x0100
	SP = 0x0100
	// Set SR to 0x24
	SR = 0x24
}

func execute() {
	if disassemble {
		fmt.Printf(" *= $%04X\n\n", PC)
	}
	for bytecounter = PC; bytecounter < len(file); instructionCounter++ {
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
				Operation: PC + 2↓, [FFFE] → PCL, [FFFF] → PCH

				The break command causes the microprocessor to go through an interrupt sequence under program control.

				This means that the program counter of the second byte after the BRK is automatically stored on the
				stack along with the processor status at the beginning of the break instruction.

				The microprocessor then transfers control to the interrupt vector.

				Other than changing the program counter, the break instruction changes no values in either the
				registers or the flags.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("BRK\n")
			}

			//  Push PC onto stack
			memory[SP] = byte(PC >> 8)
			// Store SR on stack
			memory[SP-1] = SR
			// Decrement SP
			SP--
			SP--
			// Set PC to 0xFFFE
			PC = 0xFFFE
			// Set SR interrupt disable bit to 1
			setInterruptFlag()
			// Set SR break command bit to 1
			setBreakFlag()
			// Set SR decimal mode bit to 0
			unsetDecimalFlag()
			// Set SR overflow bit to 0
			unsetOverflowFlag()
			// Set SR carry bit to 0
			unsetCarryFlag()
			// Set SR negative bit to 0
			unsetNegativeFlag()
			// Set SR zero bit to 0
			unsetZeroFlag()
			// Set PC to the value stored at 0xFFFE
			PC = int(uint16(memory[0xFFFE])<<8 | uint16(memory[0xFFFF]))
			bytecounter = PC
			incCount(0)
		case 0x18:
			/*
				CLC - Clear Carry Flag
				Operation: 0 → C

				This instruction initializes the carry flag to a 0. This operation should normally precede an ADC loop.
				It is also useful when used with a R0L instruction to clear a bit in memory.

				This instruction affects no registers in the microprocessor and no flags other than the carry flag which is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("CLC\n")
			}

			// Set SR carry flag bit 0 to 0
			unsetCarryFlag()
			incCount(1)
		case 0xD8:
			/*
				CLD - Clear Decimal Mode
				Operation: 0 → D

				This instruction sets the decimal mode flag to a 0. This all subsequent ADC and SBC instructions
				to operate as simple operations.

				CLD affects no registers in the microprocessor and no flags other than the decimal mode flag which
				is set to a 0.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("CLD\n")
			}

			unsetDecimalFlag()
			incCount(1)
		case 0x58:
			/*
				CLI - Clear Interrupt Disable
				Operation: 0 → I

				This instruction initializes the interrupt disable to a 0.
				his allows the microprocessor to receive interrupts.

				It affects no registers in the microprocessor and no flags other than the interrupt disable
				which is cleared.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("CLI\n")
			}

			// Set SR interrupt disable bit 2 to 0
			unsetInterruptFlag()
			incCount(1)
		case 0xB8:
			/*
				CLV - Clear Overflow Flag
				Operation: 0 → V

				This instruction clears the overflow flag to a 0. This command is used in conjunction with the
				set overflow pin which can change the state of the overflow flag with an external signal.

				CLV affects no registers in the microprocessor and no flags other than the overflow flag which
				is set to a 0.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("CLV\n")
			}

			// Set SR overflow flag bit 6 to 0
			unsetOverflowFlag()
			incCount(1)
		case 0xCA:
			/*
				DEX - Decrement Index Register X By One
				Operation: X - 1 → X

				This instruction subtracts one from the current value of the index register X and stores the result
				in the index register X.

				DEX does not affect the carry or overflow flag, it
				sets the N flag if it has bit 7 on as a result of the decrement, otherwise it resets the N flag;
				sets the Z flag if X is a 0 as a result of the decrement, otherwise it resets the Z flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("DEX\n")
			}

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
				Operation: Y - 1 → Y

				This instruction subtracts one from the current value in the index register Y and stores the result
				into the index register Y. The result does not affect or consider carry so that the value in the index
				register Y is decremented to 0 and then through 0 to FF.

				Decrement Y does not affect the carry or overflow flags;
				if the Y register contains bit 7 on as a result of the decrement the N flag is set,
				otherwise the N flag is reset.
				If the Y register is 0 as a result of the decrement, the Z flag is set otherwise the Z flag is reset.
				This instruction only affects the index register Y.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("DEY\n")
			}

			// Decrement the  Y register by 1
			Y--
			// If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getYBit(7) == 1 {
				setNegativeFlag()
			} else {
				unsetNegativeFlag()
			}
			incCount(1)
		case 0xE8:
			/*
				INX - Increment Index Register X By One
				Operation: X + 1 → X

				Increment X adds 1 to the current value of the X register.
				This is an 8-bit increment which does not affect the carry operation, therefore,
				if the value of X before the increment was FF, the resulting value is 00.

				INX does not affect the carry or overflow flags;
				it sets the N flag if the result of the increment has a one in bit 7, otherwise resets N;
				sets the Z flag if the result of the increment is 0, otherwise it resets the Z flag.

				INX does not affect any other register other than the X register.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("INX\n")
			}

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
				Operation: Y + 1 → Y

				Increment Y increments or adds one to the current value in the Y register,
				storing the result in the Y register.

				As in the case of INX the primary application is to step thru a set of values using the Y register.

				The INY does not affect the carry or overflow flags, sets the N flag if the result of the increment
				has a one in bit 7, otherwise resets N,
				sets Z if as a result of the increment the Y register is zero otherwise resets the Z flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("INY\n")
			}

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
				Operation: No operation
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("NOP\n")
			}
			incCount(1)
		case 0x48:
			/*
				PHA - Push Accumulator On Stack
				Operation: A↓

				This instruction transfers the current value of the accumulator to the next location on the stack,
				automatically decrementing the stack to point to the next empty location.

				The Push A instruction only affects the stack pointer register which is decremented by 1 as a result of
				the operation. It affects no flags.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("PHA\n")
			}

			// Update memory address pointed to by SP with value stored in accumulator
			memory[SP] = A
			// Decrement the stack pointer by 1 byte
			SP--
			incCount(1)
		case 0x08:
			/*
				PHP - Push Processor Status On Stack
				Operation: P↓

				This instruction transfers the contents of the processor status register unchanged to the stack,
				as governed by the stack pointer.

				The PHP instruction affects no registers or flags in the microprocessor.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("PHP\n")
			}

			// Push SR to stack
			memory[SP] = SR
			// Decrement the stack pointer by 1 byte
			SP--
			incCount(1)
		case 0x68:
			/*
				PLA - Pull Accumulator From Stack
				Operation: A↑

				This instruction adds 1 to the current value of the stack pointer and uses it to address the stack
				and loads the contents of the stack into the A register.

				The PLA instruction does not affect the carry or overflow flags.
				It sets N if the bit 7 is on in accumulator A as a result of instructions, otherwise it is reset.
				If accumulator A is zero as a result of the PLA, then the Z flag is set, otherwise it is reset.

				The PLA instruction changes content of the accumulator A to the contents of the memory location at
				stack register plus 1 and also increments the stack register.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("PLA\n")
			}

			// Increment the stack pointer by 1 byte
			SP++
			// Update accumulator with value stored in memory address pointed to by SP
			A = memory[SP]
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
				Operation: P↑

				This instruction transfers the next value on the stack to the Processor Status register,
				thereby changing all of the flags and setting the mode switches to the values from the stack.

				The PLP instruction affects no registers in the processor other than the status register.

				This instruction could affect all flags in the status register.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("PLP\n")
			}
			// Increment the stack pointer by 1 byte
			SP++
			// Update SR with the value stored at the address pointed to by SP
			SR = memory[SP]
			incCount(1)
		case 0x40:
			/*
				RTI - Return From Interrupt
				Operation: P↑ PC↑

				This instruction transfers from the stack into the microprocessor the processor status and the
				program counter location for the instruction which was interrupted.

				By virtue of the interrupt having stored this data before executing the instruction and the fact
				that the RTI re-initialises the microprocessor to the same state as when it was interrupted, the
				combination of interrupt plus RTI allows truly reentrant coding.

				The RTI instruction re-initialises all flags to the position to the point they were at the time
				the interrupt was taken and sets the program counter back to its pre-interrupt state.

				It affects no other registers in the microprocessor.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("RTI\n")
			}

			// Update SR with the value stored at the address pointed to by SP
			SR = memory[SP]
			// Update PC with the value stored at the address pointed to by SP+1
			PC = int(memory[SP] + 1)
			incCount(1)
		case 0x60:
			/*
				RTS - Return From Subroutine
				Operation: PC↑, PC + 1 → PC

				This instruction loads the program count low and program count high from the stack into the program
				counter and increments the program counter so that it points to the instruction following the JSR.

				The stack pointer is adjusted by incrementing it twice.

				The RTS instruction does not affect any flags and affects only PCL and PCH.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("RTS\n")
			}
			// Get the value stored at the address pointed to by SP
			PC = int(memory[SP])
			// Increment the stack pointer by 1 byte
			SP++
			// Increment the stack pointer by 1 byte
			SP++
			bytecounter = PC
			incCount(0)
		case 0x38:
			/*
				SEC - Set Carry Flag
				Operation: 1 → C

				This instruction initializes the carry flag to a 1.
				This operation should normally precede an SBC loop.
				It is also useful when used with a ROL instruction to initialize a bit in memory to a 1.

				This instruction affects no registers in the microprocessor and no flags other than the carry
				flag which is set.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, opcode())
				}
				fmt.Printf("SEC\n")
			}

			// Set SR carry flag bit 0 to 1
			setCarryFlag()
			incCount(1)
		case 0xF8:
			/*
				SED - Set Decimal Mode
				Operation: 1 → D

				This instruction sets the decimal mode flag D to a 1.
				This makes all subsequent ADC and SBC instructions operate as a decimal arithmetic operation.

				SED affects no registers in the microprocessor and no flags other than the decimal mode which
				is set to a 1.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("SED\n")
			}

			// Set SR decimal mode flag to 1
			setDecimalFlag()
			incCount(1)
		case 0x78:
			/*
				SEI - Set Interrupt Disable
				Operation: 1 → I

				This instruction initializes the interrupt disable to a 1.
				It is used to mask interrupt requests during system reset operations and during interrupt commands.

				It affects no registers in the microprocessor and no flags other than the interrupt disable which is set.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("SEI\n")
			}

			// Set SR interrupt disable bit 2 to 1
			setInterruptFlag()
			incCount(1)
		case 0xAA:
			/*
				TAX - Transfer Accumulator To Index X
				Operation: A → X

				This instruction takes the value from accumulator A and transfers or loads it into the index register X
				without disturbing the content of the accumulator A.

				TAX only affects the index register X, does not affect the carry or overflow flags.
				The N flag is set if the resultant value in the index register X has bit 7 on, otherwise N is reset.
				The Z bit is set if the content of the register X is 0 as a result of the operation, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("TAX\n")
			}

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
				Operation: A → Y

				This instruction moves the value of the accumulator into index register Y without affecting
				the accumulator.

				TAY instruction only affects the Y register and does not affect either the carry or overflow flags.
				If the index register Y has bit 7 on, then N is set, otherwise it is reset.
				If the content of the index register Y equals 0 as a result of the operation, Z is set on, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("TAY\n")
			}

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
				Operation: S → X

				This instruction transfers the value in the stack pointer to the index register X.

				TSX does not affect the carry or overflow flags.
				It sets N if bit 7 is on in index X as a result of the instruction, otherwise it is reset.
				If index X is zero as a result of the TSX, the Z flag is set, otherwise it is reset.
				TSX changes the value of index X, making it equal to the content of the stack pointer.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("TSX\n")
			}

			// Update X with the value stored at the address pointed to by SP
			X = memory[SP]
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
				Operation: X → A

				This instruction moves the value that is in the index register X to the accumulator A without disturbing
				the content of the index register X.

				TXA does not affect any register other than the accumulator and does not affect the carry or overflow flag.
				If the result in A has bit 7 on, then the N flag is set, otherwise it is reset.
				If the resultant value in the accumulator is 0, then the Z flag is set, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("TXA\n")
			}

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
				Operation: X → S

				This instruction transfers the value in the index register X to the stack pointer.

				TXS changes only the stack pointer, making it equal to the content of the index register X.
				It does not affect any of the flags.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("TXS\n")
			}

			// Set stack pointer to value of X register
			memory[SP] = X
			incCount(1)
		case 0x98:
			/*
				TYA - Transfer Index Y To Accumulator
				Operation: Y → A

				This instruction moves the value that is in the index register Y to accumulator A without disturbing
				the content of the register Y.

				TYA does not affect any other register other than the accumulator and does not affect the carry
				or overflow flag.
				If the result in the accumulator A has bit 7 on, the N flag is set, otherwise it is reset.
				If the resultant value in the accumulator A is 0, then the Z flag is set, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
				}
				fmt.Printf("TYA\n")
			}

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
				Operation: C ← /M7...M0/ ← 0

				The shift left instruction shifts either the accumulator or the address memory location 1 bit to
				the left, with the bit 0 always being set to 0 and the the input bit 7 being stored in the carry flag.

				ASL either shifts the accumulator left 1 bit or is a read/modify/write instruction that affects only memory.

				The instruction does not affect the overflow bit, sets N equal to the result bit 7 (bit 6 in the input),
				sets Z flag if the result is equal to 0, otherwise resets Z and stores the input bit 7 in the carry flag
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, opcode())
				}
				fmt.Printf("ASL\n")
			}
			ASL("accumulator")
		case 0x4A:
			/*
				LSR - Logical Shift Right
				Operation: 0 → /M7...M0/ → C

				This instruction shifts either the accumulator or a specified memory location 1 bit to the right,
				with the higher bit of the result always being set to 0, and the low bit which is shifted out of
				the field being stored in the carry flag.

				The shift right instruction either affects the accumulator by shifting it right 1 or is a
				read/modify/write instruction which changes a specified memory location but does not affect
				any internal registers. The shift right does not affect the overflow flag.
				The N flag is always reset.
				The Z flag is set if the result of the shift is 0 and reset otherwise.
				The carry is set equal to bit 0 of the input.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, opcode())
				}
				fmt.Printf("LSR\n")
			}

			LSR("accumulator")
		case 0x2A:
			/*
				ROL - Rotate Left
				Operation: C ← /M7...M0/ ← C

				The rotate left instruction shifts either the accumulator or addressed memory left 1 bit, with
				the input carry being stored in bit 0 and with the input bit 7 being stored in the carry flags.

				The ROL instruction either shifts the accumulator left 1 bit and stores the carry in accumulator bit 0
				or does not affect the internal registers at all.
				The ROL instruction sets carry equal to the input bit 7,
				sets N equal to the input bit 6,
				sets the Z flag if the result of the rotate is 0,
				otherwise it resets Z and does not affect the overflow flag at all.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, opcode())
				}
				fmt.Printf("ROL\n")
			}

			ROL("accumulator")
		case 0x6A:
			/*
				ROR - Rotate Right
				Operation: C → /M7...M0/ → C

				The rotate right instruction shifts either the accumulator or addressed memory right 1 bit with
				bit 0 shifted into the carry and carry shifted into bit 7.

				The ROR instruction either shifts the accumulator right 1 bit and stores the carry in accumulator
				bit 7 or does not affect the internal registers at all.
				The ROR instruction sets carry equal to input bit 0,
				sets N equal to the input carry and sets the Z flag if the result of the rotate is 0;
				It otherwise it resets Z and does not affect the overflow flag at all.

				(Available on Microprocessors after June, 1976)
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, opcode())
				}
				fmt.Printf("ROR\n")
			}
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
				Operation: A + M + C → A, C

				This instruction adds the value of memory and carry from the previous operation to the value of the
				accumulator and stores the result in the accumulator.

				This instruction affects the accumulator; sets the carry flag when the sum of a binary add exceeds
				255 or when the sum of a decimal add exceeds 99, otherwise carry is reset.

				The overflow flag is set when the sign or bit 7 is changed due to the result exceeding +127 or -128,
				otherwise overflow is reset.

				The negative flag is set if the accumulator result contains bit 7 on, otherwise the negative flag is reset.

				The zero flag is set if the accumulator result is 0, otherwise the zero flag is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("ADC #$%02X\n", operand1())
			}

			ADC("immediate")
		case 0x29:
			/*
				AND - "AND" Memory with Accumulator
				Operation: A ∧ M → A

				The AND instruction transfer the accumulator and memory to the adder which performs a bit-by-bit
				AND operation and stores the result back in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("AND #$%02X\n", operand1())
			}

			AND("immediate")
		case 0xC9:
			/*
				CMP - Compare Memory and Accumulator
				Operation: A - M

				This instruction subtracts the contents of memory from the contents of the accumulator.

				The use of the CMP affects the following flags: Z flag is set on an equal comparison, reset otherwise;
				the N flag is set or reset by the result bit 7,
				the carry flag is set when the value in memory is less than or equal to the accumulator,
				reset when it is greater than the accumulator.

				The accumulator is not affected.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("CMP #$%02X\n", operand1())
			}

			CMP("immediate")
		case 0xE0:
			/*
				CPX - Compare Index Register X To Memory
				Operation: X - M

				This instruction subtracts the value of the addressed memory location from the content of index
				register X using the adder but does not store the result;
				therefore, its only use is to set the N, Z and C flags to allow for comparison between the index
				register X and the value in memory.

				The CPX instruction does not affect any register in the machine; it also does not affect the overflow flag.
				It causes the carry to be set on if the absolute value of the index register X is equal to or greater
				than the data from memory.
				If the value of the memory is greater than the content of the index register X, carry is reset.
				If the results of the subtraction contain a bit 7, then the N flag is set, if not, it is reset.
				If the value in memory is equal to the value in index register X, the Z flag is set, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("CPX #$%02X\n", operand1())
			}

			CPX("immediate")
		case 0xC0:
			/*
				CPY - Compare Index Register Y To Memory
				Operation: Y - M

				This instruction performs a two's complement subtraction between the index register Y and the
				specified memory location. The results of the subtraction are not stored anywhere. The instruction is
				strictly used to set the flags.

				CPY affects no registers in the microprocessor and also does not affect the overflow flag.

				If the value in the index register Y is equal to or greater than the value in the memory,
				the carry flag will be set, otherwise it will be cleared.

				If the results of the subtraction contain bit 7 on the N bit will be set, otherwise it will be cleared.

				If the value in the index register Y and the value in the memory are equal, the zero flag will be set,
				otherwise it will be cleared.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("CPY #$%02X\n", operand1())
			}

			CPY("immediate")
		case 0x49:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
				Operation: A ⊻ M → A

				The EOR instruction transfers the memory and the accumulator to the adder which performs a binary
				"EXCLUSIVE OR" on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("EOR #$%02X\n", operand1())
			}

			EOR("immediate")
		case 0xA9:
			/*
				LDA - Load Accumulator with Memory
				Operation: M → A

				When instruction LDA is executed by the microprocessor, data is transferred from memory to the
				accumulator and stored in the accumulator.

				LDA affects the contents of the accumulator, does not affect the carry or overflow flags;
				sets the zero flag if the accumulator is zero as a result of the LDA, otherwise resets the zero flag;
				sets the negative flag if bit 7 of the accumulator is a 1, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("LDA #$%02X\n", operand1())
			}
			LDA("immediate")
		case 0xA2:
			/*
				LDX - Load Index Register X From Memory
				Operation: M → X

				Load the index register X from memory.

				LDX does not affect the C or V flags; sets Z if the value loaded was zero, otherwise resets it;
				sets N if the value loaded in bit 7 is a 1; otherwise N is reset, and affects only the X register.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("LDX #$%02X\n", operand1())
			}
			LDX("immediate")
		case 0xA0:
			/*
				LDY - Load Index Register Y From Memory
				Operation: M → Y

				Load the index register Y from memory.

				LDY does not affect the C or V flags,
				sets the N flag if the value loaded in bit 7 is a 1, otherwise resets N,
				sets Z flag if the loaded value is zero otherwise resets Z and only affects the Y register.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("LDY #$%02X\n", operand1())
			}

			LDY("immediate")
		case 0x09:
			/*
				ORA - "OR" Memory with Accumulator
				Operation: A ∨ M → A

				The ORA instruction transfers the memory and the accumulator to the adder which performs a binary "OR"
				on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("ORA #$%02x\n", operand1())
			}

			ORA("immediate")
		case 0xE9:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
				Operation: A - M - ~C → A

				This instruction subtracts the value of memory and borrow from the value of the accumulator,
				using two's complement arithmetic, and stores the result in the accumulator.

				Borrow is defined as the carry flag complemented; therefore, a resultant carry flag indicates
				that a borrow has not occurred.

				This instruction affects the accumulator.
				The carry flag is set if the result is greater than or equal to 0.
				The carry flag is reset when the result is less than 0, indicating a borrow.
				The overflow flag is set when the result exceeds +127 or -127, otherwise it is reset.
				The negative flag is set if the result in the accumulator has bit 7 on, otherwise it is reset.
				The Z flag is set if the result in the accumulator is 0, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("SBC #$%02X\n", operand1())
			}
			SBC("immediate")

		// Zero Page addressing mode instructions
		/*
			$nn

			The zero page instructions allow for shorter code and execution times by only fetching the second byte of the instruction and assuming a zero high address byte. Careful use of the zero page can result in significant increase in code efficiency.

			Bytes: 2
		*/
		case 0x65:
			/*
				ADC - Add Memory to Accumulator with Carry
				Operation: A + M + C → A, C

				This instruction adds the value of memory and carry from the previous operation to the value of the
				accumulator and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the carry flag when the sum of a binary add exceeds 255 or when the sum of a decimal add
				exceeds 99, otherwise carry is reset.
				The overflow flag is set when the sign or bit 7 is changed due to the result exceeding +127 or -128,
				otherwise overflow is reset.
				The negative flag is set if the accumulator result contains bit 7 on,
				otherwise the negative flag is reset.
				The zero flag is set if the accumulator result is 0, otherwise the zero flag is reset.

				Note on the MOS 6502:

				In decimal mode, the N, V and Z flags are not consistent with the decimal result.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("ADC $%02X\n", operand1())
			}

			ADC("zeroPage")
		case 0x25:
			/*
				AND - "AND" Memory with Accumulator
				Operation: A ∧ M → A

				The AND instruction transfer the accumulator and memory to the adder which performs a bit-by-bit
				AND operation and stores the result back in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("AND $%02X\n", operand1())
			}

			AND("zeropage")
		case 0x06:
			/*
				ASL - Arithmetic Shift Left
				Operation: C ← /M7...M0/ ← 0

				The shift left instruction shifts either the accumulator or the address memory location 1 bit to
				the left, with the bit 0 always being set to 0 and the the input bit 7 being stored in the carry flag.

				ASL either shifts the accumulator left 1 bit or is a read/modify/write instruction that affects only memory.

				The instruction does not affect the overflow bit,
				sets N equal to the result bit 7 (bit 6 in the input),
				sets Z flag if the result is equal to 0, otherwise resets Z and stores the input bit 7 in the carry flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("ASL $%02x\n", operand1())
			}

			ASL("zeropage")
		case 0x24:
			/*
				BIT - Test Bits in Memory with Accumulator
				Operation: A ∧ M, M7 → N, M6 → V

				This instruction performs an AND between a memory location and the accumulator but does not store
				the result of the AND into the accumulator.

				The bit instruction affects the N flag with N being set to the value of bit 7 of the memory being tested
				the V flag with V being set equal to bit 6 of the memory being tested and
				Z being set by the result of the AND operation between the accumulator and the memory if
				the result is Zero, Z is reset otherwise.
				It does not affect the accumulator.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("BIT $%02X\n", operand1())
			}

			BIT("zeropage")
		case 0xC5:
			/*
				CMP - Compare Memory and Accumulator
				Operation: A - M

				This instruction subtracts the contents of memory from the contents of the accumulator.

				The use of the CMP affects the following flags: Z flag is set on an equal comparison, reset otherwise;
				the N flag is set or reset by the result bit 7,
				the carry flag is set when the value in memory is less than or equal to the accumulator,
				reset when it is greater than the accumulator.
				The accumulator is not affected.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("CMP $%02X\n", operand1())
			}
			CMP("zeropage")
		case 0xE4:
			/*
				CPX - Compare Index Register X To Memory
				Operation: X - M

				This instruction subtracts the value of the addressed memory location from the content of
				index register X using the adder but does not store the result;
				therefore, its only use is to set the N, Z and C flags to allow for comparison between the
				index register X and the value in memory.

				The CPX instruction does not affect any register in the machine;
				it also does not affect the overflow flag.
				It causes the carry to be set on if the absolute value of the index register X is equal to
				or greater than the data from memory.
				If the value of the memory is greater than the content of the index register X, carry is reset.
				If the results of the subtraction contain a bit 7, then the N flag is set, if not, it is reset.
				If the value in memory is equal to the value in index register X, the Z flag is set,
				otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("CPX $%02X\n", operand1())
			}

			CPX("zeropage")
		case 0xC4:
			/*
				CPY - Compare Index Register Y To Memory
				Operation: Y - M

				This instruction performs a two's complement subtraction between the index register Y and the
				specified memory location.
				The results of the subtraction are not stored anywhere.
				The instruction is strictly used to set the flags.

				CPY affects no registers in the microprocessor and also does not affect the overflow flag.
				If the value in the index register Y is equal to or greater than the value in the memory,
				the carry flag will be set, otherwise it will be cleared.
				If the results of the subtraction contain bit 7 on the N bit will be set, otherwise it will be cleared.
				If the value in the index register Y and the value in the memory are equal, the zero flag will be set,
				otherwise it will be cleared.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("CPY $%02X\n", operand1())
			}
			CPY("zeropage")
		case 0xC6:
			/*
				DEC - Decrement Memory By One
				Operation: M - 1 → M

				This instruction subtracts 1, in two's complement, from the contents of the addressed memory location.

				The decrement instruction does not affect any internal register in the microprocessor.

				It does not affect the carry or overflow flags.
				If bit 7 is on as a result of the decrement, then the N flag is set, otherwise it is reset.
				If the result of the decrement is 0, the Z flag is set, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("DEC $%02X\n", operand1())
			}

			DEC("zeropage")
		case 0x45:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
				Operation: A ⊻ M → A

				The EOR instruction transfers the memory and the accumulator to the adder which performs a binary
				"EXCLUSIVE OR" on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("EOR $%02X\n", operand1())
			}

			EOR("zeropage")
		case 0xE6:
			/*
				INC - Increment Memory By One
				Operation: M + 1 → M

				This instruction adds 1 to the contents of the addressed memory location.

				The increment memory instruction does not affect any internal registers and does not affect the
				carry or overflow flags.
				If bit 7 is on as the result of the increment,N is set, otherwise it is reset;
				if the increment causes the result to become 0, the Z flag is set on, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("INC $%02X\n", operand1())
			}

			INC("zeropage")
		case 0xA5:
			/*
				LDA - Load Accumulator with Memory
				Operation: M → A

				When instruction LDA is executed by the microprocessor, data is transferred from memory to the
				accumulator and stored in the accumulator.

				LDA affects the contents of the accumulator, does not affect the carry or overflow flags;
				sets the zero flag if the accumulator is zero as a result of the LDA, otherwise resets the zero flag;
				sets the negative flag if bit 7 of the accumulator is a 1, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("LDA $%02X\n", operand1())
			}

			LDA("zeropage")
		case 0xA6:
			/*
				LDX - Load Index Register X From Memory
				Operation: M → X

				Load the index register X from memory.

				LDX does not affect the C or V flags;
				sets Z if the value loaded was zero, otherwise resets it;
				sets N if the value loaded in bit 7 is a 1; otherwise N is reset, and affects only the X register.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("LDX $%02X\n", operand1())
			}
			LDX("zeropage")
		case 0xA4:
			/*
				LDY - Load Index Register Y From Memory
				Operation: M → Y

				Load the index register Y from memory.

				LDY does not affect the C or V flags,
				sets the N flag if the value loaded in bit 7 is a 1, otherwise resets N,
				sets Z flag if the loaded value is zero otherwise resets Z and only affects the Y register.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("LDY $%02X\n", operand1())
			}

			LDY("zeropage")
		case 0x46:
			/*
				LSR - Logical Shift Right
				Operation: 0 → /M7...M0/ → C

				This instruction shifts either the accumulator or a specified memory location 1 bit to the right,
				with the higher bit of the result always being set to 0, and the low bit which is shifted out of the
				field being stored in the carry flag.

				The shift right instruction either affects the accumulator by shifting it right 1 or is a
				read/modify/write instruction which changes a specified memory location but does not affect any
				internal registers.
				The shift right does not affect the overflow flag.
				The N flag is always reset.
				The Z flag is set if the result of the shift is 0 and reset otherwise.
				The carry is set equal to bit 0 of the input.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("LSR $%02X\n", operand1())
			}

			LSR("zeropage")
		case 0x05:
			/*
				ORA - "OR" Memory with Accumulator
				Operation: A ∨ M → A

				The ORA instruction transfers the memory and the accumulator to the adder which performs a binary "OR"
				on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on,
				otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("ORA $%02x\n", operand1())
			}

			ORA("zeropage")
		case 0x26:
			/*
				ROL - Rotate Left
				Operation: C ← /M7...M0/ ← C

				The rotate left instruction shifts either the accumulator or addressed memory left 1 bit,
				with the input carry being stored in bit 0 and with the input bit 7 being stored in the carry flags.

				The ROL instruction either shifts the accumulator left 1 bit and stores the carry in accumulator bit 0
				or does not affect the internal registers at all.
				The ROL instruction sets carry equal to the input bit 7,
				sets N equal to the input bit 6 ,
				sets the Z flag if the result of the rotate is 0, otherwise it resets Z and does not affect
				the overflow flag at all.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("ROL $%02X\n", operand1())
			}

			ROL("zeropage")
		case 0x66:
			/*
				ROR - Rotate Right
				Operation: C → /M7...M0/ → C

				The rotate right instruction shifts either the accumulator or addressed memory right 1 bit with
				bit 0 shifted into the carry and carry shifted into bit 7.

				The ROR instruction either shifts the accumulator right 1 bit and stores the carry in accumulator bit 7
				or does not affect the internal registers at all.
				The ROR instruction sets carry equal to input bit 0,
				sets N equal to the input carry and
				sets the Z flag if the result of the rotate is 0; otherwise it resets Z and
				does not affect the overflow flag at all.

				(Available on Microprocessors after June, 1976)
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("ROR $%02X\n", operand1())
			}

			ROR("zeropage")
		case 0xE5:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
				Operation: A - M - ~C → A

				This instruction subtracts the value of memory and borrow from the value of the accumulator,
				using two's complement arithmetic, and stores the result in the accumulator.
				Borrow is defined as the carry flag complemented; therefore, a resultant carry flag indicates that
				a borrow has not occurred.

				This instruction affects the accumulator.
				The carry flag is set if the result is greater than or equal to 0.
				The carry flag is reset when the result is less than 0, indicating a borrow.
				The overflow flag is set when the result exceeds +127 or -127, otherwise it is reset.
				The negative flag is set if the result in the accumulator has bit 7 on, otherwise it is reset.
				The Z flag is set if the result in the accumulator is 0, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("SBC $%02X\n", operand1())
			}

			SBC("zeropage")
		case 0x85:
			/*
				STA - Store Accumulator in Memory

				Operation: A → M

				This instruction transfers the contents of the accumulator to memory.

				This instruction affects none of the flags in the processor status register and does not affect the accumulator.
			*/

			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("STA $%02X\n", operand1())
			}

			STA("zeropage")
		case 0x86:
			/*
				STX - Store Index Register X In Memory
				Operation: X → M

				Transfers value of X register to addressed memory location.

				No flags or registers in the microprocessor are affected by the store operation.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("STX $%02X\n", operand1())
			}
			STX("zeropage")
		case 0x84:
			/*
				STY - Store Index Register Y In Memory
				Operation: Y → M

				Transfer the value of the Y register to the addressed memory location.

				STY does not affect any flags or registers in the microprocessor.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("STY $%02X\n", operand1())
			}

			STY("zeropage")

		// X Indexed Zero Page addressing mode instructions
		/*
			$nn,X

			This form of addressing is used in conjunction with the X index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the high order 8 bits of memory and crossing of page boundaries does not occur.

			Bytes: 2
		*/
		case 0x75:
			/*
				ADC - Add Memory to Accumulator with Carry
				Operation: A + M + C → A, C

				This instruction adds the value of memory and carry from the previous operation to the value of the
				accumulator and stores the result in the accumulator.

				This instruction affects the accumulator; sets the carry flag when the sum of a binary add exceeds
				255 or when the sum of a decimal add exceeds 99, otherwise carry is reset.
				The overflow flag is set when the sign or bit 7 is changed due to the result exceeding +127 or -128,
				otherwise overflow is reset.
				The negative flag is set if the accumulator result contains bit 7 on, otherwise the negative flag is reset.
				The zero flag is set if the accumulator result is 0, otherwise the zero flag is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("ADC $%02X,X\n", operand1())
			}

			ADC("zeropagex")
		case 0x35:
			/*
				AND - "AND" Memory with Accumulator
				Operation: A ∧ M → A

				The AND instruction transfer the accumulator and memory to the adder which performs a bit-by-bit
				AND operation and stores the result back in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t(Zero Page,X)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("AND $%02X,X\n", operand1())
			}

			AND("zeropagex")
		case 0x16:
			/*
				ASL - Arithmetic Shift Left
				Operation: C ← /M7...M0/ ← 0

				The shift left instruction shifts either the accumulator or the address memory location 1 bit to the
				left, with the bit 0 always being set to 0 and the the input bit 7 being stored in the carry flag.
				ASL either shifts the accumulator left 1 bit or is a read/modify/write instruction that affects only memory.

				The instruction does not affect the overflow bit,
				sets N equal to the result bit 7 (bit 6 in the input),
				sets Z flag if the result is equal to 0, otherwise resets Z and stores the input bit 7 in the carry flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(ASL - Zero Page,X)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("ASL $%02X,X\n", operand1())
			}

			ASL("zeropagex")
		case 0xD5:
			/*
				CMP - Compare Memory and Accumulator
				Operation: A - M

				This instruction subtracts the contents of memory from the contents of the accumulator.

				The use of the CMP affects the following flags:
				Z flag is set on an equal comparison, reset otherwise;
				the N flag is set or reset by the result bit 7,
				the carry flag is set when the value in memory is less than or equal to the accumulator, reset when it
				is greater than the accumulator.
				The accumulator is not affected.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("CMP $%02X,X\n", operand1())
			}

			CMP("zeropagex")
		case 0xD6:
			/*
				DEC - Decrement Memory By One
				Operation: M - 1 → M

				This instruction subtracts 1, in two's complement, from the contents of the addressed memory location.

				The decrement instruction does not affect any internal register in the microprocessor.
				It does not affect the carry or overflow flags.
				If bit 7 is on as a result of the decrement, then the N flag is set, otherwise it is reset.
				If the result of the decrement is 0, the Z flag is set, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("DEC $%02X,X\n", operand1())
			}

			DEC("zeropagex")
		case 0xB5:
			/*
				LDA - Load Accumulator with Memory
				Operation: M → A

				When instruction LDA is executed by the microprocessor, data is transferred from memory to the accumulator
				and stored in the accumulator.

				LDA affects the contents of the accumulator, does not affect the carry or overflow flags;
				sets the zero flag if the accumulator is zero as a result of the LDA, otherwise resets the zero flag;
				sets the negative flag if bit 7 of the accumulator is a 1, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("LDA $%02X,X\n", operand1())
			}

			LDA("zeropagex")
		case 0xB4:
			/*
				LDY - Load Index Register Y From Memory
				Operation: M → Y

				Load the index register Y from memory.

				LDY does not affect the C or V flags, sets the N flag if the value loaded in bit 7 is a 1, otherwise resets N,
				sets Z flag if the loaded value is zero otherwise resets Z and only affects the Y register.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("LDY $%02X,X\n", operand1())
			}

			LDY("zeropagex")
		case 0x56:
			/*
				LSR - Logical Shift Right
				Operation: 0 → /M7...M0/ → C

				This instruction shifts either the accumulator or a specified memory location 1 bit to the right,
				with the higher bit of the result always being set to 0, and the low bit which is shifted out of the
				field being stored in the carry flag.

				The shift right instruction either affects the accumulator by shifting it right 1 or is a read/modify/write
				instruction which changes a specified memory location but does not affect any internal registers.

				The shift right does not affect the overflow flag.
				The N flag is always reset.
				The Z flag is set if the result of the shift is 0 and reset otherwise.
				The carry is set equal to bit 0 of the input.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("LSR $%02X,X\n", operand1())
			}

			LSR("zeropagex")
		case 0x15:
			/*
				ORA - "OR" Memory with Accumulator
				Operation: A ∨ M → A

				The ORA instruction transfers the memory and the accumulator to the adder which performs a binary "OR"
				on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("ORA $%02x,X\n", operand1())
			}

			ORA("zeropagex")
		case 0x36:
			/*
				ROL - Rotate Left
				Operation: C ← /M7...M0/ ← C

				The rotate left instruction shifts either the accumulator or addressed memory left 1 bit, with the
				input carry being stored in bit 0 and with the input bit 7 being stored in the carry flags.

				The ROL instruction either shifts the accumulator left 1 bit and stores the carry in accumulator bit 0
				or does not affect the internal registers at all.
				The ROL instruction sets carry equal to the input bit 7,
				sets N equal to the input bit 6 ,
				sets the Z flag if the result of the rotate is 0, otherwise it resets Z and
				does not affect the overflow flag at all.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("ROL $%02X,X\n", operand1())
			}
			ROL("zeropagex")
		case 0x76:
			/*
				ROR - Rotate Right
				Operation: C → /M7...M0/ → C

				The rotate right instruction shifts either the accumulator or addressed memory right 1 bit with bit 0
				shifted into the carry and carry shifted into bit 7.

				The ROR instruction either shifts the accumulator right 1 bit and stores the carry in accumulator
				bit 7 or does not affect the internal registers at all.
				The ROR instruction sets carry equal to input bit 0,
				sets N equal to the input carry and
				sets the Z flag if the result of the rotate is 0; otherwise it resets Z and
				does not affect the overflow flag at all.

				(Available on Microprocessors after June, 1976)
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("ROR $%02X,X\n", operand1())
			}
			ROR("zeropagex")
		case 0xF5:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
				Operation: A - M - ~C → A

				This instruction subtracts the value of memory and borrow from the value of the accumulator,
				using two's complement arithmetic, and stores the result in the accumulator.
				Borrow is defined as the carry flag complemented; therefore, a resultant carry flag indicates
				that a borrow has not occurred.

				This instruction affects the accumulator.
				The carry flag is set if the result is greater than or equal to 0.
				The carry flag is reset when the result is less than 0, indicating a borrow.
				The overflow flag is set when the result exceeds +127 or -127, otherwise it is reset.
				The negative flag is set if the result in the accumulator has bit 7 on, otherwise it is reset.
				The Z flag is set if the result in the accumulator is 0, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("SBC $%02X,X\n", operand1())
			}
			SBC("zeropagex")
		case 0x95:
			/*
				STA - Store Accumulator in Memory
				Operation: A → M

				This instruction transfers the contents of the accumulator to memory.

				This instruction affects none of the flags in the processor status register and does not affect
				the accumulator.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("STA $%02X,X\n", operand1())
			}
			STA("zeropagex")
		case 0x94:
			/*
				STY - Store Index Register Y In Memory
				Operation: Y → M

				Transfer the value of the Y register to the addressed memory location.

				STY does not affect any flags or registers in the microprocessor.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("STY $%02X,X\n", operand1())
			}

			STY("zeropagex")

		// Y Indexed Zero Page addressing mode instructions
		/*
			$nn,Y

			This form of addressing is used in conjunction with the Y index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the high order 8 bits of memory and crossing of page boundaries does not occur.

			Bytes: 2
		*/
		case 0xB6:
			/*
				LDX - Load Index Register X From Memory
				Operation: M → X

				Load the index register X from memory.

				LDX does not affect the C or V flags;
				sets Z if the value loaded was zero, otherwise resets it;
				sets N if the value loaded in bit 7 is a 1; otherwise N is reset, and affects only the X register.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Y)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("LDX $%02X,Y\n", operand1())
			}
			LDX("zeropagey")
		case 0x96:
			/*
				STX - Store Index Register X In Memory
				Operation: X → M

				Transfers value of X register to addressed memory location.

				No flags or registers in the microprocessor are affected by the store operation.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Y)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("STX $%02X,Y\n", operand1())
			}

			STX("zeropagey")

		// X Indexed Zero Page Indirect addressing mode instructions
		/*
			($nn,X)

			In indexed indirect addressing, the second byte of the instruction is added to the contents of the X index register, discarding the carry. The result of this addition points to a memory location on page zero whose contents is the low order eight bits of the effective address. The next memory location in page zero contains the high order eight bits of the effective address. Both memory locations specifying the high and low order bytes of the effective address must be in page zero.

			Bytes: 2
		*/
		case 0x61:
			/*
				ADC - Add Memory to Accumulator with Carry
				Operation: A + M + C → A, C

				This instruction adds the value of memory and carry from the previous operation to the value of the
				accumulator and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the carry flag when the sum of a binary add exceeds 255 or when the sum of a decimal add exceeds 99
				otherwise carry is reset.
				The overflow flag is set when the sign or bit 7 is changed due to the result exceeding +127 or -128,
				otherwise overflow is reset.
				The negative flag is set if the accumulator result contains bit 7 on, otherwise the negative flag is reset.
				The zero flag is set if the accumulator result is 0, otherwise the zero flag is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, opcode(), operand1())
				}
				fmt.Printf("ADC ($%02X,X)\n", operand1())
			}
			ADC("indirectx")
		case 0x21:
			/*
				AND - "AND" Memory with Accumulator
				Operation: A ∧ M → A

				The AND instruction transfer the accumulator and memory to the adder which performs a bit-by-bit
				AND operation and stores the result back in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page Indirect))\n", PC, opcode(), operand1())
				}
				fmt.Printf("AND ($%02X,X)\n", operand1())
			}

			AND("indirectx")
		case 0xC1:
			/*
				CMP - Compare Memory and Accumulator
				Operation: A - M

				This instruction subtracts the contents of memory from the contents of the accumulator.

				The use of the CMP affects the following flags:
				Z flag is set on an equal comparison, reset otherwise;
				the N flag is set or reset by the result bit 7,
				the carry flag is set when the value in memory is less than or equal to the accumulator,
				reset when it is greater than the accumulator.
				The accumulator is not affected.

			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, opcode(), operand1())
				}
				fmt.Printf("CMP ($%02X,X)\n", operand1())
			}

			CMP("indirectx")
		case 0x41:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
				Operation: A ⊻ M → A

				The EOR instruction transfers the memory and the accumulator to the adder which performs a binary
				"EXCLUSIVE OR" on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page, Indirect))\n", PC, opcode(), operand1())
				}
				fmt.Printf("EOR ($%02X,X)\n", operand1())
			}

			EOR("indirectx")
		case 0xA1:
			/*
				LDA - Load Accumulator with Memory
				Operation: M → A

				When instruction LDA is executed by the microprocessor, data is transferred from memory to the
				accumulator and stored in the accumulator.

				LDA affects the contents of the accumulator, does not affect the carry or overflow flags;
				sets the zero flag if the accumulator is zero as a result of the LDA, otherwise resets the zero flag;
				sets the negative flag if bit 7 of the accumulator is a 1, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, opcode(), operand1())
				}
				fmt.Printf("LDA ($%02X,X)\n", operand1())
			}

			LDA("indirectx")
		case 0x01:
			/*
				ORA - "OR" Memory with Accumulator
				Operation: A ∨ M → A

				The ORA instruction transfers the memory and the accumulator to the adder which performs a binary
				"OR" on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page Indirect))\n", PC, opcode(), operand1())
				}
				fmt.Printf("ORA ($%02x,X)\n", operand1())
			}

			ORA("indirectx")
		case 0xE1:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
				Operation: A - M - ~C → A

				This instruction subtracts the value of memory and borrow from the value of the accumulator,
				using two's complement arithmetic, and stores the result in the accumulator.
				Borrow is defined as the carry flag complemented; therefore, a resultant carry flag indicates that
				a borrow has not occurred.

				This instruction affects the accumulator.
				The carry flag is set if the result is greater than or equal to 0.
				The carry flag is reset when the result is less than 0, indicating a borrow.
				The overflow flag is set when the result exceeds +127 or -127, otherwise it is reset.
				The negative flag is set if the result in the accumulator has bit 7 on, otherwise it is reset.
				The Z flag is set if the result in the accumulator is 0, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, opcode(), operand1())
				}
				fmt.Printf("SBC ($%02X,X)\n", operand1())
			}

			SBC("indirectx")
		case 0x81:
			/*
				STA - Store Accumulator in Memory
				Operation: A → M

				This instruction transfers the contents of the accumulator to memory.

				This instruction affects none of the flags in the processor status register and does not
				affect the accumulator.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, opcode(), operand1())
				}
				fmt.Printf("STA ($%02X,X)\n", operand1())
			}
			STA("indirectx")

		// Zero Page Indirect Y Indexed addressing mode instructions
		/*
			($nn),Y

			In indirect indexed addressing, the second byte of the instruction points to a memory location in page zero. The contents of this memory location is added to the contents of the Y index register, the result being the low order eight bits of the effective address. The carry from this addition is added to the contents of the next page zero memory location, the result being the high order eight bits of the effective address.

			Bytes: 2
		*/
		case 0x71:
			/*
				ADC - Add Memory to Accumulator with Carry
				Operation: A + M + C → A, C

				This instruction adds the value of memory and carry from the previous operation to the value of the
				accumulator and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the carry flag when the sum of a binary add exceeds 255 or when the sum of a decimal add exceeds 99
				otherwise carry is reset.
				The overflow flag is set when the sign or bit 7 is changed due to the result exceeding +127 or -128,
				otherwise overflow is reset.
				The negative flag is set if the accumulator result contains bit 7 on, otherwise the negative flag is reset.
				The zero flag is set if the accumulator result is 0, otherwise the zero flag is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, opcode(), operand1())
				}
				fmt.Printf("ADC ($%02X),Y\n", operand1())
			}

			ADC("indirecty")
		case 0x31:
			/*
				AND - "AND" Memory with Accumulator
				Operation: A ∧ M → A

				The AND instruction transfer the accumulator and memory to the adder which performs a bit-by-bit
				AND operation and stores the result back in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, opcode(), operand1())
				}
				fmt.Printf("AND ($%02X),Y\n", operand1())
			}

			AND("indirecty")
		case 0xD1:
			/*
				CMP - Compare Memory and Accumulator
				Operation: A - M

				This instruction subtracts the contents of memory from the contents of the accumulator.

				The use of the CMP affects the following flags: Z flag is set on an equal comparison, reset otherwise;
				the N flag is set or reset by the result bit 7,
				the carry flag is set when the value in memory is less than or equal to the accumulator,
				reset when it is greater than the accumulator.
				The accumulator is not affected.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, opcode(), operand1())
				}
				fmt.Printf("CMP ($%02X),Y\n", operand1())
			}

			CMP("indirecty")
		case 0x51:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
				Operation: A ⊻ M → A

				The EOR instruction transfers the memory and the accumulator to the adder which performs a binary
				"EXCLUSIVE OR" on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, opcode(), operand1())
				}
				fmt.Printf("EOR ($%02X),Y\n", operand1())
			}

			EOR("indirecty")
		case 0xB1:
			/*
				LDA - Load Accumulator with Memory
				Operation: M → A

				When instruction LDA is executed by the microprocessor, data is transferred from memory to the
				accumulator and stored in the accumulator.

				LDA affects the contents of the accumulator, does not affect the carry or overflow flags;
				sets the zero flag if the accumulator is zero as a result of the LDA, otherwise resets the zero flag;
				sets the negative flag if bit 7 of the accumulator is a 1, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, opcode(), operand1())
				}
				fmt.Printf("LDA ($%02X),Y\n", operand1())
			}

			LDA("indirecty")
		case 0x11:
			/*
				ORA - "OR" Memory with Accumulator
				Operation: A ∨ M → A

				The ORA instruction transfers the memory and the accumulator to the adder which performs a binary
				"OR" on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page),Indirect Y)\n", PC, opcode(), operand1())
				}
				fmt.Printf("ORA ($%02x),Y\n", operand1())
			}
			ORA("indirecty")

		case 0xF1:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
				Operation: A - M - ~C → A

				This instruction subtracts the value of memory and borrow from the value of the accumulator,
				using two's complement arithmetic, and stores the result in the accumulator.
				Borrow is defined as the carry flag complemented; therefore, a resultant carry flag
				indicates that a borrow has not occurred.

				This instruction affects the accumulator.
				The carry flag is set if the result is greater than or equal to 0.
				The carry flag is reset when the result is less than 0, indicating a borrow.
				The overflow flag is set when the result exceeds +127 or -127, otherwise it is reset.
				The negative flag is set if the result in the accumulator has bit 7 on, otherwise it is reset.
				The Z flag is set if the result in the accumulator is 0, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, opcode(), operand1())
				}
				fmt.Printf("SBC ($%02X),Y\n", operand1())
			}

			SBC("indirecty")
		case 0x91:
			/*
				STA - Store Accumulator in Memory
				Operation: A → M

				This instruction transfers the contents of the accumulator to memory.

				This instruction affects none of the flags in the processor status register and does not affect the accumulator.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, opcode(), operand1())
				}
				fmt.Printf("STA ($%02X),Y\n", operand1())
			}

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
				Operation: Branch on N = 0

				This instruction is the complementary branch to branch on result minus.

				It is a conditional branch which takes the branch when the N bit is reset (0).

				BPL is used to test if the previous result bit 7 was off (0) and branch on result minus is used to
				determine if the previous result was minus or bit 7 was on (1).

				The instruction affects no flags or other registers other than the P counter and only affects the
				P counter when the N bit is reset.

				Relative addressing is used only with branch instructions and establishes a destination for
				the conditional branch.

				The second byte of-the instruction becomes the operand which is an “Offset" added to the
				contents of the lower eight bits of the program counter when the counter is set at the next
				instruction.
				The range of the offset is —128 to +127 bytes from the next instruction.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("BPL $%02X\n", (bytecounter+2+int(operand1()))&0xFF)
			}

			// Get offset from relative address in operand
			offset := operand1()
			// If SR negative flag bit 7 is 0 then branch
			if getSRBit(7) == 0 {
				// Branch
				// Add offset to lower 8bits of PC
				PC += int(offset)
				// If the offset is negative, decrement the PC by 1
				// If bit 7 is unset then it's negative
				if readBit(7, offset) == 0 {
					PC--
				}
				bytecounter += 2
				incCount(0)
			} else {
				// Don't branch
				incCount(2)
			}
		case 0x30:
			/*
				BMI - Branch on Result Minus
				Operation: Branch on N = 1

				This instruction takes the conditional branch if the N bit is set.

				BMI does not affect any of the flags or any other part of the machine other than the program counter
				and then only if the N bit is on.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("BMI $%02X\n", (bytecounter+2+int(operand1()))&0xFF)
			}

			// Get offset from relative address in operand
			offset := operand1()
			// If SR negative flag bit 7 is 1 then branch
			if getSRBit(7) == 1 {
				// Branch
				// Add offset to lower 8bits of PC
				PC += int(offset)
				// If the offset is negative, decrement the PC by 1
				// If bit 7 is unset then it's negative
				if readBit(7, offset) == 0 {
					PC--
				}
				bytecounter += 2
				incCount(0)
			} else {
				// Don't branch
				incCount(2)
			}
		case 0x50:
			/*
				BVC - Branch on Overflow Clear
				Operation: Branch on V = 0

				This instruction tests the status of the V flag and takes the conditional branch if the flag is not set.

				BVC does not affect any of the flags and registers other than the program counter and only
				when the overflow flag is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("BVC $%02X\n", bytecounter+2+int(operand1()))
			}

			// Get offset from relative address in operand
			offset := operand1()
			// If the overflow flag is not set, branch to the address specified by the operand
			if getSRBit(6) == 0 {
				// Branch
				// Add offset to lower 8bits of PC
				PC += int(offset)
				// If the offset is negative, decrement the PC by 1
				// If bit 7 is unset then it's negative
				if readBit(7, offset) == 0 {
					PC--
				}
				bytecounter += 2
				incCount(0)
			} else {
				// Don't branch
				incCount(2)
			}
		case 0x55:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
				Operation: A ⊻ M → A

				The EOR instruction transfers the memory and the accumulator to the adder which performs a binary
				"EXCLUSIVE OR" on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("EOR $%02X,X\n", operand1())
			}

			EOR("zeropagex")
		case 0x70:
			/*
				BVS - Branch on Overflow Set
				Operation: Branch on V = 1

				This instruction tests the V flag and takes the conditional branch if V is on.

				BVS does not affect any flags or registers other than the program, counter and only
				when the overflow flag is set.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("BVS $%04X\n", bytecounter+2+int(operand1()))
			}

			// Get offset from relative address in operand
			offset := int(operand1())
			// If SR overflow flag bit 6 is set then branch to operand 1
			if getSRBit(6) == 1 {
				PC = bytecounter + 2 + offset
			}
			incCount(2)

		case 0x90:
			/*
				BCC - Branch on Carry Clear
				Operation: Branch on C = 0

				This instruction tests the state of the carry bit and takes a conditional branch if the carry bit is reset.

				It affects no flags or registers other than the program bytecounter and then only if the C flag is not on.
			*/

			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("BCC $%02X\n", (bytecounter+2+int(operand1()))&0xFF)
			}

			// Get offset from relative address in operand
			offset := operand1()
			// If carry flag is not set then branch to offset
			if getSRBit(0) == 0 {
				// Add offset to program counter
				PC += int(offset)
			}
			incCount(2)
		case 0xB0:
			/*
				BCS - Branch on Carry Set
				Operation: Branch on C = 1

				This instruction takes the conditional branch if the carry flag is on.

				BCS does not affect any of the flags or registers except for the program counter and only
				then if the carry flag is on.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("BCS $%02X\n", (bytecounter+2+int(operand1()))&0xFF)
			}
			// Get offset from relative address in operand
			offset := operand1()

			// If the carry flag is set, branch to the address in the operand
			if getSRBit(0) == 1 {
				PC = bytecounter + 2 + int(offset)&0xFF
			}
			incCount(2)

		case 0xD0:
			/*
				BNE - Branch on Result Not Zero
				Operation: Branch on Z = 0

				This instruction could also be called "Branch on Not Equal."
				It tests the Z flag and takes the conditional branch if the Z flag is not on,
				indicating that the previous result was not zero.

				BNE does not affect any of the flags or registers other than the program counter
				and only then if the Z flag is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("BNE $%02X\n", (bytecounter+2+int(operand1()))&0xFF)
			}

			// Get offset from relative address in operand
			offset := operand1()
			// If Z flag is not set, branch to address
			if getSRBit(1) != 1 {
				// Branch
				// Add offset to lower 8bits of PC
				PC += int(offset)
				// If the offset is negative, decrement the PC by 1
				// If bit 7 is unset then it's negative
				if readBit(7, offset) == 0 {
					PC--
				}
				bytecounter += 2
				incCount(0)
			} else {
				// Don't branch
				incCount(2)
			}
		case 0xF0:
			/*
				BEQ - Branch on Result Zero
				Operation: Branch on Z = 1

				This instruction could also be called "Branch on Equal."

				It takes a conditional branch whenever the Z flag is on or the previous result is equal to 0.

				BEQ does not affect any of the flags or registers other than the program bytecounter and only then
				when the Z flag is set.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("BEQ $%04X\n", bytecounter+2+int(operand1()))
			}

			// Get offset from address in operand
			offset := operand1()
			// Get relative address from offset
			relativeAddress := PC + 2 + int(offset)

			// If Z flag is set, branch to relative address
			if getSRBit(1) == 1 {
				// If relative address is negative, subtract from PC
				if relativeAddress<<7 == 0b10000000 {
					PC -= int(offset)
					bytecounter = PC
					incCount(0)
				} else {
					// If relative address is positive, add to PC
					PC = relativeAddress
					bytecounter = PC
					incCount(0)
				}
			} else {
				// If Z flag is not set, don't branch
				incCount(2)
			}
		case 0xF6:
			/*
				INC - Increment Memory By One
				Operation: M + 1 → M

				This instruction adds 1 to the contents of the addressed memory location.

				The increment memory instruction does not affect any internal registers and does not affect the
				carry or overflow flags.
				If bit 7 is on as the result of the increment,N is set, otherwise it is reset;
				if the increment causes the result to become 0, the Z flag is set on, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
				}
				fmt.Printf("INC $%02X,X\n", operand1())
			}

			INC("zeropagex")
		}

		// 3 byte instructions with 2 operands
		switch opcode() {
		// Absolute addressing mode instructions
		/*
			$nnnn

			In absolute addressing, the second byte of the instruction specifies the eight low order bits of the effective address while the third byte specifies the eight high order bits. Thus, the absolute addressing mode allows access to the entire 65 K bytes of addressable memory.

			Bytes: 3
		*/
		case 0x6D:
			/*
				ADC - Add Memory to Accumulator with Carry
				Operation: A + M + C → A, C

				This instruction adds the value of memory and carry from the previous operation to the value of the
				accumulator and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the carry flag when the sum of a binary add exceeds 255 or when the sum of a decimal add exceeds 99
				otherwise carry is reset.

				The overflow flag is set when the sign or bit 7 is changed due to the result exceeding +127 or -128,
				otherwise overflow is reset.

				The negative flag is set if the accumulator result contains bit 7 on, otherwise the negative flag is reset.
				The zero flag is set if the accumulator result is 0, otherwise the zero flag is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("ADC $%02X%02X\n", operand2(), operand1())
			}
			ADC("absolute")
		case 0x2D:
			/*
				AND - "AND" Memory with Accumulator
				Operation: A ∧ M → A

				The AND instruction transfer the accumulator and memory to the adder which performs a bit-by-bit
				AND operation and stores the result back in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("AND $%02X%02X\n", operand2(), operand1())
			}

			AND("absolute")
		case 0x0E:
			/*
				ASL - Arithmetic Shift Left
				Operation: C ← /M7...M0/ ← 0

				The shift left instruction shifts either the accumulator or the address memory location
				1 bit to the left, with the bit 0 always being set to 0 and the the input bit 7 being stored in
				the carry flag.

				ASL either shifts the accumulator left 1 bit or is a read/modify/write instruction that affects only memory.

				The instruction does not affect the overflow bit,
				sets N equal to the result bit 7 (bit 6 in the input),
				sets Z flag if the result is equal to 0, otherwise resets Z
				and stores the input bit 7 in the carry flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("ASL $%02X%02X\n", operand2(), operand1())
			}

			ASL("absolute")
		case 0x2C:
			/*
				BIT - Test Bits in Memory with Accumulator
				Operation: A ∧ M, M7 → N, M6 → V

				This instruction performs an AND between a memory location and the accumulator but does not store the
				result of the AND into the accumulator.

				The bit instruction affects the N flag with
				N being set to the value of bit 7 of the memory being tested, the V flag with
				V being set equal to bit 6 of the memory being tested and
				Z being set by the result of the AND operation between the accumulator and the memory if the
				result is Zero, Z is reset otherwise.

				It does not affect the accumulator.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("BIT $%02X%02X\n", operand2(), operand1())
			}

			BIT("absolute")
		case 0xCD:
			/*
				CMP - Compare Memory and Accumulator
				Operation: A - M

				This instruction subtracts the contents of memory from the contents of the accumulator.

				The use of the CMP affects the following flags:
				Z flag is set on an equal comparison, reset otherwise;
				the N flag is set or reset by the result bit 7,
				the carry flag is set when the value in memory is less than or equal to the accumulator,
				reset when it is greater than the accumulator.
				The accumulator is not affected.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("CMP $%02X%02X\n", operand2(), operand1())
			}
			CMP("absolute")

		case 0xEC:
			/*
				CPX - Compare Index Register X To Memory
				Operation: X - M

				This instruction subtracts the value of the addressed memory location from the content of
				index register X using the adder but does not store the result;
				therefore, its only use is to set the N, Z and C flags to allow for comparison between the
				index register X and the value in memory.

				The CPX instruction does not affect any register in the machine;
				it also does not affect the overflow flag.
				It causes the carry to be set on if the absolute value of the index register X is equal to
				or greater than the data from memory.
				If the value of the memory is greater than the content of the index register X, carry is reset.
				If the results of the subtraction contain a bit 7, then the N flag is set, if not, it is reset.
				If the value in memory is equal to the value in index register X, the Z flag is set,
				otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("CPX $%02X%02X\n", operand2(), operand1())
			}

			CPX("absolute")
		case 0xCC:
			/*
				CPY - Compare Index Register Y To Memory
				Operation: Y - M

				This instruction performs a two's complement subtraction between the index register Y and the specified
				memory location. The results of the subtraction are not stored anywhere. The instruction is strictly
				used to set the flags.

				CPY affects no registers in the microprocessor and also does not affect the overflow flag.
				If the value in the index register Y is equal to or greater than the value in the memory,
				the carry flag will be set, otherwise it will be cleared.
				If the results of the subtracttion contain bit 7 on the N bit will be set, otherwise it will be cleared.
				If the value in the index register Y and the value in the memory are equal, the zero flag will be set,
				otherwise it will be cleared.


			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("CPY $%02X%02X\n", operand2(), operand1())
			}

			CPY("absolute")
		case 0xCE:
			/*
				DEC - Decrement Memory By One
				Operation: M - 1 → M

				This instruction subtracts 1, in two's complement, from the contents of the addressed memory location.

				The decrement instruction does not affect any internal register in the microprocessor.
				It does not affect the carry or overflow flags.
				If bit 7 is on as a result of the decrement, then the N flag is set, otherwise it is reset.
				If the result of the decrement is 0, the Z flag is set, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("DEC $%02X%02X\n", operand2(), operand1())
			}

			DEC("absolute")
		case 0x4D:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
				Operation: A ⊻ M → A

				The EOR instruction transfers the memory and the accumulator to the adder which performs a binary
				"EXCLUSIVE OR" on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("EOR $%02X%02X\n", operand2(), operand1())
			}

			EOR("absolute")
		case 0xEE:
			/*
				INC - Increment Memory By One
				Operation: M + 1 → M

				This instruction adds 1 to the contents of the addressed memory location.

				The increment memory instruction does not affect any internal registers and does not affect the carry
				or overflow flags.
				If bit 7 is on as the result of the increment,N is set, otherwise it is reset;
				if the increment causes the result to become 0, the Z flag is set on, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("INC $%02X%02X\n", operand2(), operand1())
			}

			INC("absolute")
		case 0x4C:
			/*
				JMP - JMP Indirect
				Operation: [PC + 1] → PCL, [PC + 2] → PCH

				This instruction establishes a new valne for the program counter.

				It affects only the program counter in the microprocessor and affects no flags in the status register.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("JMP $%04X\n", int(operand2())<<8|int(operand1()))
			}
			JMP("absolute")
		case 0x20:
			/*
				JSR - Jump To Subroutine
				Operation: PC + 2↓, [PC + 1] → PCL, [PC + 2] → PCH

				This instruction transfers control of the program counter to a subroutine location but leaves a
				return pointer on the stack to allow the user to return to perform the next instruction in the
				main program after the subroutine is complete.

				To accomplish this, JSR instruction stores the program counter address which points to the last byte
				of the jump instruction onto the stack using the stack pointer. The stack byte contains the
				program count high first, followed by program count low. The JSR then transfers the addresses following
				the jump instruction to the	program counter low and the program counter high, thereby directing the
				program to begin at that new address.

				The JSR instruction affects no flags, causes the stack pointer to be decremented by 2 and substitutes
				new values into the program bytecounter low and the program bytecounter high.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("JSR $%04X\n", int(operand2())<<8|int(operand1()))
			}
			// Push the program counter onto the stack
			memory[SP] = byte(PC >> 8)
			SP--
			memory[SP] = byte(PC & 0xFF)
			SP--
			// Set the program counter to the absolute address from the operands
			PC = int(operand2())<<8 | int(operand1())
			bytecounter = PC
			incCount(0)
		case 0xAD:
			/*
				LDA - Load Accumulator with Memory
				Operation: M → A

				When instruction LDA is executed by the microprocessor, data is transferred from memory to
				the accumulator and stored in the accumulator.

				 LDA affects the contents of the accumulator, does not affect the carry or overflow flags;
				sets the zero flag if the accumulator is zero as a result of the LDA, otherwise resets the zero flag;
				sets the negative flag if bit 7 of the accumulator is a 1, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("LDA $%04X\n", uint16(operand2())<<8|uint16(operand1()))
			}

			LDA("absolute")
		case 0xAE:
			/*
				LDX - Load Index Register X From Memory
				Operation: M → X

				Load the index register X from memory.

				LDX does not affect the C or V flags;
				sets Z if the value loaded was zero, otherwise resets it;
				sets N if the value loaded in bit 7 is a 1; otherwise N is reset, and affects only the X register.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("LDX $%02X%02X\n", operand2(), operand1())
			}

			LDX("absolute")
		case 0xAC:
			/*
				LDY - Load Index Register Y From Memory
				Operation: M → Y

				Load the index register Y from memory.

				LDY does not affect the C or V flags,
				sets the N flag if the value loaded in bit 7 is a 1, otherwise resets N,
				sets Z flag if the loaded value is zero otherwise resets Z and only affects the Y register.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("LDY $%02X%02X\n", operand2(), operand1())
			}

			LDY("absolute")
		case 0x4E:
			/*
				LSR - Logical Shift Right
				Operation: 0 → /M7...M0/ → C

				This instruction shifts either the accumulator or a specified memory location 1 bit to the right,
				with the higher bit of the result always being set to 0, and the low bit which is shifted out of the
				field being stored in the carry flag.

				The shift right instruction either affects the accumulator by shifting it right 1 or is a
				read/modify/write instruction which changes a specified memory location but does not affect any
				internal registers.
				The shift right does not affect the overflow flag.
				The N flag is always reset.
				The Z flag is set if the result of the shift is 0 and reset otherwise.
				The carry is set equal to bit 0 of the input.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("LSR $%02X%02X\n", operand2(), operand1())
			}

			LSR("absolute")
		case 0x0D:
			/*
				ORA - "OR" Memory with Accumulator
				Operation: A ∨ M → A

				The ORA instruction transfers the memory and the accumulator to the adder which performs a binary
				"OR" on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("ORA $%02X%02X\n", operand2(), operand1())
			}

			ORA("absolute")
		case 0x2E:
			/*
				ROL - Rotate Left
				Operation: C ← /M7...M0/ ← C

				The rotate left instruction shifts either the accumulator or addressed memory left 1 bit,
				with the input carry being stored in bit 0 and with the input bit 7 being stored in the carry flags.

				The ROL instruction either shifts the accumulator left 1 bit and stores the carry in accumulator bit 0
				or does not affect the internal registers at all.
				The ROL instruction sets carry equal to the input bit 7,
				sets N equal to the input bit 6,
				sets the Z flag if the result of the rotate is 0, otherwise it resets Z and
				does not affect the overflow flag at all.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("ROL $%02X%02X\n", operand2(), operand1())
			}

			ROL("absolute")
		case 0x6E:
			/*
				ROR - Rotate Right
				Operation: C → /M7...M0/ → C

				The rotate right instruction shifts either the accumulator or addressed memory right 1 bit with bit 0
				shifted into the carry and carry shifted into bit 7.

				The ROR instruction either shifts the accumulator right 1 bit and stores the carry in accumulator bit 7
				or does not affect the internal registers at all.
				The ROR instruction sets carry equal to input bit 0,
				sets N equal to the input carry and
				sets the Z flag if the result of the rotate is 0; otherwise it resets Z
				and does not affect the overflow flag at all.

				(Available on Microprocessors after June, 1976)
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("ROR $%02X%02X\n", operand2(), operand1())
			}
			ROR("absolute")
		case 0xED:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
				Operation: A - M - ~C → A

				This instruction subtracts the value of memory and borrow from the value of the accumulator,
				using two's complement arithmetic, and stores the result in the accumulator.

				Borrow is defined as the carry flag complemented;
				therefore, a resultant carry flag indicates that a borrow has not occurred.

				This instruction affects the accumulator.
				The carry flag is set if the result is greater than or equal to 0.
				The carry flag is reset when the result is less than 0, indicating a borrow.
				The overflow flag is set when the result exceeds +127 or -127, otherwise it is reset.
				The negative flag is set if the result in the accumulator has bit 7 on, otherwise it is reset.
				The Z flag is set if the result in the accumulator is 0, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("SBC $%02X%02X\n", operand2(), operand1())
			}

			SBC("absolute")
		case 0x8D:
			/*
				STA - Store Accumulator in Memory
				Operation: A → M

				This instruction transfers the contents of the accumulator to memory.

				This instruction affects none of the flags in the processor status register and
				does not affect the accumulator.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("STA $%04X\n", uint16(operand2())<<8|uint16(operand1()))
			}
			STA("absolute")
		case 0x8E:
			/*
				STX - Store Index Register X In Memory
				Operation: X → M

				Transfers value of X register to addressed memory location.

				No flags or registers in the microprocessor are affected by the store operation.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("STX $%02X%02X\n", operand2(), operand1())
			}

			STX("absolute")
		case 0x8C:
			/*
				STY - Store Index Register Y In Memory
				Operation: Y → M

				Transfer the value of the Y register to the addressed memory location.

				STY does not affect any flags or registers in the microprocessor.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("STY $%02X%02X\n", operand2(), operand1())
			}

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
				Operation: A + M + C → A, C

				This instruction adds the value of memory and carry from the previous operation to the value of the
				accumulator and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the carry flag when the sum of a binary add exceeds 255 or when the sum of a decimal add exceeds 99
				otherwise carry is reset.
				The overflow flag is set when the sign or bit 7 is changed due to the result exceeding +127 or -128,
				otherwise overflow is reset.
				The negative flag is set if the accumulator result contains bit 7 on, otherwise the negative flag is reset
				The zero flag is set if the accumulator result is 0, otherwise the zero flag is reset.

			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("ADC $%02X%02X,X\n", operand2(), operand1())
			}

			ADC("absolutex")
		case 0x3D:
			/*
				AND - "AND" Memory with Accumulator
				Operation: A ∧ M → A

				The AND instruction transfer the accumulator and memory to the adder which performs a bit-by-bit
				AND operation and stores the result back in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("AND $%02X%02X,X\n", operand2(), operand1())
			}

			AND("absolutex")
		case 0x1E:
			/*
				ASL - Arithmetic Shift Left
				Operation: C ← /M7...M0/ ← 0

				The shift left instruction shifts either the accumulator or the address memory location
				1 bit to the left, with the bit 0 always being set to 0 and the the input bit 7 being stored
				in the carry flag.
				ASL either shifts the accumulator left 1 bit or is a read/modify/write instruction that
				affects only memory.

				The instruction does not affect the overflow bit,
				sets N equal to the result bit 7 (bit 6 in the input),
				sets Z flag if the result is equal to 0, otherwise resets Z and
				stores the input bit 7 in the carry flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("ASL $%02X%02X,X\n", operand2(), operand1())
			}

			ASL("absolutex")
		case 0xDD:
			/*
				CMP - Compare Memory and Accumulator
				Operation: A - M

				This instruction subtracts the contents of memory from the contents of the accumulator.

				The use of the CMP affects the following flags:
				Z flag is set on an equal comparison, reset otherwise;
				the N flag is set or reset by the result bit 7,
				the carry flag is set when the value in memory is less than or equal to the accumulator,
				reset when it is greater than the accumulator.
				The accumulator is not affected.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("CMP $%02X%02X,X\n", operand2(), operand1())
			}

			CMP("absolutex")
		case 0xDE:
			/*
				DEC - Decrement Memory By One
				Operation: M - 1 → M

				This instruction subtracts 1, in two's complement, from the contents of the addressed memory location.

				The decrement instruction does not affect any internal register in the microprocessor.
				It does not affect the carry or overflow flags.
				If bit 7 is on as a result of the decrement, then the N flag is set, otherwise it is reset.
				If the result of the decrement is 0, the Z flag is set, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("DEC $%02X%02X,X\n", operand2(), operand1())
			}

			DEC("absolutex")
		case 0x5D:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
				Operation: A ⊻ M → A

				The EOR instruction transfers the memory and the accumulator to the adder which performs a binary
				"EXCLUSIVE OR" on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("EOR $%02X%02X,X\n", operand2(), operand1())
			}

			EOR("absolutex")
		case 0xFE:
			/*
				INC - Increment Memory By One
				Operation: M + 1 → M

				This instruction adds 1 to the contents of the addressed memory location.

				The increment memory instruction does not affect any internal registers and does not affect the
				carry or overflow flags.
				If bit 7 is on as the result of the increment,N is set, otherwise it is reset;
				if the increment causes the result to become 0, the Z flag is set on, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("INC $%04X,X\n", int(operand2())<<8|int(operand1()))
			}
			INC("absolutex")
		case 0xBD:
			/*
				LDA - Load Accumulator with Memory
				Operation: M → A

				When instruction LDA is executed by the microprocessor, data is transferred from memory to the
				accumulator and stored in the accumulator.

				LDA affects the contents of the accumulator,
				does not affect the carry or overflow flags;
				sets the zero flag if the accumulator is zero as a result of the LDA, otherwise resets the zero flag;
				sets the negative flag if bit 7 of the accumulator is a 1, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("LDA $%02X%02X,X\n", operand2(), operand1())
			}

			LDA("absolutex")
		case 0xBC:
			/*
				LDY - Load Index Register Y From Memory
				Operation: M → Y

				Load the index register Y from memory.

				LDY does not affect the C or V flags, sets the N flag if the value loaded in bit 7 is a 1, otherwise resets N,
				sets Z flag if the loaded value is zero otherwise resets Z and only affects the Y register.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("LDY $%02X%02X,X\n", operand2(), operand1())
			}

			LDY("absolutex")
		case 0x5E:
			/*
				LSR - Logical Shift Right
				Operation: 0 → /M7...M0/ → C

				This instruction shifts either the accumulator or a specified memory location 1 bit to the right,
				with the higher bit of the result always being set to 0, and the low bit which is shifted out of
				the field being stored in the carry flag.

				The shift right instruction either affects the accumulator by shifting it right 1 or is a
				read/modify/write instruction which changes a specified memory location but does not affect any
				internal registers.

				The shift right does not affect the overflow flag.
				The N flag is always reset.
				The Z flag is set if the result of the shift is 0 and reset otherwise.
				The carry is set equal to bit 0 of the input.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("LSR $%02X%02X,X\n", operand2(), operand1())
			}

			LSR("absolutex")
		case 0x1D:
			/*
				ORA - "OR" Memory with Accumulator
				Operation: A ∨ M → A

				The ORA instruction transfers the memory and the accumulator to the adder which performs a binary
				"OR" on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("ORA $%02X%02X,X\n", operand2(), operand1())
			}

			ORA("absolutex")
		case 0x3E:
			/*
				ROL - Rotate Left
				Operation: C ← /M7...M0/ ← C

				The rotate left instruction shifts either the accumulator or addressed memory left 1 bit,
				with the input carry being stored in bit 0 and with the input bit 7 being stored in the carry flags.

				The ROL instruction either shifts the accumulator left 1 bit and stores the carry in accumulator bit 0
				or does not affect the internal registers at all.
				The ROL instruction sets carry equal to the input bit 7,
				sets N equal to the input bit 6,
				sets the Z flag if the result of the rotate is 0, otherwise it resets Z and
				does not affect the overflow flag at all.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("ROL $%02X%02X,X\n", operand2(), operand1())
			}
			ROL("absolutex")
		case 0x7E:
			/*
				ROR - Rotate Right
				Operation: C → /M7...M0/ → C

				The rotate right instruction shifts either the accumulator or addressed memory right 1 bit with bit 0
				shifted into the carry and carry shifted into bit 7.

				The ROR instruction either shifts the accumulator right 1 bit and stores the carry in accumulator bit 7
				or does not affect the internal registers at all.
				The ROR instruction sets carry equal to input bit 0,
				sets N equal to the input carry and sets the Z flag if the result of the rotate is 0;
				otherwise it resets Z and does not affect the overflow flag at all.

				(Available on Microprocessors after June, 1976)
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("ROR $%02X%02X,X\n", operand2(), operand1())
			}

			ROR("absolutex")
		case 0xFD:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
				Operation: A - M - ~C → A

				This instruction subtracts the value of memory and borrow from the value of the accumulator,
				using two's complement arithmetic, and stores the result in the accumulator.
				Borrow is defined as the carry flag complemented; therefore, a resultant carry flag indicates
				that a borrow has not occurred.

				This instruction affects the accumulator.
				The carry flag is set if the result is greater than or equal to 0.
				The carry flag is reset when the result is less than 0, indicating a borrow.
				The overflow flag is set when the result exceeds +127 or -127, otherwise it is reset.
				The negative flag is set if the result in the accumulator has bit 7 on, otherwise it is reset.
				The Z flag is set if the result in the accumulator is 0, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("SBC $%02X%02X,X\n", operand2(), operand1())
			}

			SBC("absolutex")
		case 0x9D:
			/*
				STA - Store Accumulator in Memory
				Operation: A → M

				This instruction transfers the contents of the accumulator to memory.

				This instruction affects none of the flags in the processor status register and does not affect the accumulator.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("STA $%02X%02X,X\n", operand2(), operand1())
			}

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
				Operation: A + M + C → A, C

				This instruction adds the value of memory and carry from the previous operation to the value of the
				accumulator and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the carry flag when the sum of a binary add exceeds 255 or when the sum of a decimal add exceeds 99
				otherwise carry is reset.
				The overflow flag is set when the sign or bit 7 is changed due to the result exceeding +127 or -128
				otherwise overflow is reset.
				The negative flag is set if the accumulator result contains bit 7 on, otherwise the negative flag is reset.
				The zero flag is set if the accumulator result is 0, otherwise the zero flag is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("ADC $%04X,Y\n", int(operand2())<<8|int(operand1()))
			}

			ADC("absolutey")
		case 0x39:
			/*
				AND - "AND" Memory with Accumulator
				Operation: A ∧ M → A

				The AND instruction transfer the accumulator and memory to the adder which performs a bit-by-bit
				AND operation and stores the result back in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("AND $%02X%02X,Y\n", operand2(), operand1())
			}

			AND("absolutey")
		case 0xD9:
			/*
				CMP - Compare Memory and Accumulator
				Operation: A - M

				This instruction subtracts the contents of memory from the contents of the accumulator.

				The use of the CMP affects the following flags:
				Z flag is set on an equal comparison, reset otherwise;
				the N flag is set or reset by the result bit 7,
				the carry flag is set when the value in memory is less than or equal to the accumulator, reset when
				it is greater than the accumulator.
				The accumulator is not affected.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("CMP $%02X%02X,Y\n", operand2(), operand1())
			}

			CMP("absolutey")
		case 0x59:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
				Operation: A ⊻ M → A

				The EOR instruction transfers the memory and the accumulator to the adder which performs a
				binary "EXCLUSIVE OR" on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("EOE $%02X%02X,Y\n", operand2(), operand1())
			}

			EOR("absolutey")
		case 0xB9:
			/*
				LDA - Load Accumulator with Memory
				Operation: M → A

				When instruction LDA is executed by the microprocessor, data is transferred from memory to the
				accumulator and stored in the accumulator.

				LDA affects the contents of the accumulator, does not affect the carry or overflow flags;
				sets the zero flag if the accumulator is zero as a result of the LDA, otherwise resets the zero flag;
				sets the negative flag if bit 7 of the accumulator is a 1, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("LDA $%02X%02X,Y\n", operand2(), operand1())
			}

			LDA("absolutey")
		case 0xBE:
			/*
				LDX - Load Index Register X From Memory
				Operation: M → X

				Load the index register X from memory.

				LDX does not affect the C or V flags;
				sets Z if the value loaded was zero, otherwise resets it;
				sets N if the value loaded in bit 7 is a 1; otherwise N is reset, and affects only the X register.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("LDX $%02X%02X,Y\n", operand2(), operand1())
			}
			LDX("absolutey")
		case 0x19:
			/*
				ORA - "OR" Memory with Accumulator
				Operation: A ∨ M → A

				The ORA instruction transfers the memory and the accumulator to the adder which performs a binary
				"OR" on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("ORA $%02X%02X,Y\n", operand2(), operand1())
			}

			ORA("absolutey")
		case 0xF9:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
				Operation: A - M - ~C → A

				This instruction subtracts the value of memory and borrow from the value of the accumulator,
				using two's complement arithmetic, and stores the result in the accumulator.
				Borrow is defined as the carry flag complemented; therefore, a resultant carry flag indicates
				that a borrow has not occurred.

				This instruction affects the accumulator.
				The carry flag is set if the result is greater than or equal to 0.
				The carry flag is reset when the result is less than 0, indicating a borrow.
				The overflow flag is set when the result exceeds +127 or -127, otherwise it is reset.
				The negative flag is set if the result in the accumulator has bit 7 on, otherwise it is reset.
				The Z flag is set if the result in the accumulator is 0, otherwise it is reset.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("SBC $%02X%02X,Y\n", operand2(), operand1())
			}

			SBC("absolutey")
		case 0x99:
			/*
				STA - Store Accumulator in Memory
				Operation: A → M

				This instruction transfers the contents of the accumulator to memory.

				This instruction affects none of the flags in the processor status register
				and does not affect the accumulator.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("STA $%02X%02X,Y\n", operand2(), operand1())
			}

			STA("absolutey")

		// Absolute Indirect addressing mode instructions
		case 0x6C:
			/*
				JMP - JMP Indirect
				Operation: [PC + 1] → PCL, [PC + 2] → PCH

				This instruction establishes a new value for the program counter.

				It affects only the program counter in the microprocessor and affects no flags in the status register.
			*/
			if disassemble {
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute Indirect)\n", PC, opcode(), operand1(), operand2())
				}
				fmt.Printf("JMP ($%04X)\n", int(operand2())<<8|int(operand1()))
			}

			JMP("indirect")
			/*
				// Get 16bit absolute address
				address := uint16(operand2())<<8 | uint16(operand1())
				// Get the indirect address
				indirectAddress := uint16(memory[address])<<8 | uint16(memory[address])
				// Set the program counter to the indirect address
				PC = int(indirectAddress)
				incCount(3)

			*/
		}
	}
	fmt.Printf("Number of instructions executed: %d\n", instructionCounter)
	if memory[0x0210] == 0xFF {
		fmt.Println("Program completed successfully!")
	} else {
		fmt.Printf("Program failed on opcode $%02X.\n", memory[0x0210])
	}
}
