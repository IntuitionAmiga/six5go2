package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	allsuitea    = flag.Bool("allsuitea", false, "AllSuiteA ROM")
	klausd       = flag.Bool("klausd", false, "Klaus Dormann's 6502 functional test ROM")
	plus4        = flag.Bool("plus4", false, "Plus/4 ROMs")
	ruudb        = flag.Bool("ruudb", false, "RuudB's 8K Test ROM")
	c64          = flag.Bool("c64", false, "C64 ROMs")
	disassemble  = flag.Bool("dis", false, "Disassembler mode (Optional)")
	stateMonitor = flag.Bool("state", false, "State monitor mode (Optional)")

	disassembledInstruction string
	instructionCounter      uint32 = 0

	C64BASICROM  = make([]byte, 8192)
	C64KERNALROM = make([]byte, 8192)
	C64CHARROM   = make([]byte, 4096)

	PLUS4BASICROM  = make([]byte, 16384)
	PLUS4KERNALROM = make([]byte, 16384)
	PLUS4CHARROM   = make([]byte, 4096)
	THREEPLUS1ROM  = make([]byte, 16384)

	AllSuiteAROM  = make([]byte, 16384)
	KlausDTestROM = make([]byte, 65536)
	RuudBTestROM  = make([]byte, 8192)
)

func loadROMs() {
	if *plus4 {

		// Open the KERNAL ROM file
		file, _ := os.Open("roms/plus4/kernal")
		// Read the KERNAL ROM data into PLUS4KERNALROM
		_, _ = io.ReadFull(file, PLUS4KERNALROM)
		// Copy the KERNAL ROM data into memory at plus4kernalROMAddress
		copy(memory[plus4kernalROMAddress:], PLUS4KERNALROM)
		fmt.Printf("Copying %vKB Commodore Plus/4 KERNAL ROM into memory at $%04X to $%04X\n\n", len(PLUS4KERNALROM)/1024, plus4kernalROMAddress, plus4kernalROMAddress+len(PLUS4KERNALROM)-1)
		err := file.Close()
		if err != nil {
			return
		}

		// Open the BASIC ROM file
		file, _ = os.Open("roms/plus4/basic")
		_, _ = io.ReadFull(file, PLUS4BASICROM)
		fmt.Printf("Copying %vKB Commodore Plus/4 BASIC ROM into memory at $%04X to $%04X\n\n", len(PLUS4BASICROM)/1024, plus4basicROMAddress, plus4basicROMAddress+len(PLUS4BASICROM)-1)
		copy(memory[plus4basicROMAddress:plus4basicROMAddress+len(PLUS4BASICROM)], PLUS4BASICROM)
		err = file.Close()
		if err != nil {
			return
		}
	}
	if *c64 {
		// Load the BASIC ROM
		file, _ := os.Open("roms/c64/basic")
		_, _ = io.ReadFull(file, C64BASICROM)
		fmt.Printf("Copying %vKB Commodore 64 BASIC ROM into memory from $%04X to $%04X\n\n", len(C64BASICROM)/1024, c64basicROMAddress, c64basicROMAddress+len(C64BASICROM)-1)
		copy(memory[c64basicROMAddress:c64basicROMAddress+len(C64BASICROM)], C64BASICROM)
		err := file.Close()
		if err != nil {
			return
		}

		// Load the KERNAL ROM
		file, _ = os.Open("roms/c64/kernal")
		_, _ = io.ReadFull(file, C64KERNALROM)
		fmt.Printf("Copying %vKB Commodore 64 KERNAL ROM into memory from $%04X to $%04X\n\n", len(C64KERNALROM)/1024, c64kernalROMAddress, c64kernalROMAddress+len(C64KERNALROM)-1)
		copy(memory[c64kernalROMAddress:c64kernalROMAddress+len(C64KERNALROM)], C64KERNALROM)
		err = file.Close()
		if err != nil {
			return
		}

		// Load the CHARACTER ROM
		file, _ = os.Open("roms/c64/chargen")
		_, _ = io.ReadFull(file, C64CHARROM)
		fmt.Printf("Copying %vKB Commodore 64 CHARACTER ROM into memory from $%04X to $%04X\n\n", len(C64CHARROM)/1024, c64charROMAddress, c64charROMAddress+len(C64CHARROM)-1)
		copy(memory[c64charROMAddress:c64charROMAddress+len(C64CHARROM)], C64CHARROM)
		err = file.Close()
		if err != nil {
			return
		}
	}

	if *allsuitea {
		// Copy AllSuiteA ROM into memory
		file, _ := os.Open("roms/tests/AllSuiteA.bin")
		_, _ = io.ReadFull(file, AllSuiteAROM)
		fmt.Printf("Copying %vKB AllSuiteA ROM into memory from $%04X to $%04X\n\n", len(AllSuiteAROM)/1024, AllSuiteAROMAddress, AllSuiteAROMAddress+len(AllSuiteAROM)-1)
		copy(memory[AllSuiteAROMAddress:], AllSuiteAROM)
		// Set the interrupt vector addresses manually
		writeMemory(IRQVectorAddress, 0x00)   // Low byte of 0x4000
		writeMemory(IRQVectorAddress+1, 0x40) // High byte of 0x4000
		err := file.Close()
		if err != nil {
			return
		}
	}

	if *klausd {
		//Copy roms/6502_functional_test.bin into memory
		file, _ := os.Open("roms/tests/6502_functional_test.bin")
		_, _ = io.ReadFull(file, KlausDTestROM)
		copy(memory[KlausDTestROMAddress:], KlausDTestROM)
		fmt.Printf("Copying Klaus Dormann's %vKB 6502 Functional Test ROM into memory from $%04X to $%04X\n\n", len(KlausDTestROM)/1024, KlausDTestROMAddress, KlausDTestROMAddress+len(KlausDTestROM)-1)
		// Set the interrupt vector addresses manually
		writeMemory(IRQVectorAddress, 0x00)   // Low byte of 0x4000
		writeMemory(IRQVectorAddress+1, 0x40) // High byte of 0x4000
		err := file.Close()
		if err != nil {
			return
		}
	}

	if *ruudb {
		//Copy roms/TTL6502.BIN into memory
		file, _ := os.Open("roms/tests/TTL6502.BIN")
		_, _ = io.ReadFull(file, RuudBTestROM)
		copy(memory[RuudBTestROMAddress:], RuudBTestROM)
		fmt.Printf("Copying Ruud Baltissen's %vKB Test ROM into memory from $%04X to $%04X\n\n", len(RuudBTestROM)/1024, RuudBTestROMAddress, RuudBTestROMAddress+len(RuudBTestROM)-1)
		err := file.Close()
		if err != nil {
			return
		}
	}
	if *stateMonitor {
		fmt.Printf("|  PC  | OP |OPERANDS|    DISASSEMBLY   | \t REGISTERS\t  |  STACK  | SR FLAGS | INST COUNT | CYCLE COUNT | TIME SPENT  |\n")
		fmt.Printf("|------|----|--------|------------------|-------------------------|---------|----------|------------|-------------|-------------|\n")
	}
}
func dumpMemoryToFile(memory [65536]byte) {
	file, err := os.Create("memorydump.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	for _, byteValue := range memory {
		_, err := fmt.Fprintf(file, "%02X ", byteValue)
		if err != nil {
			return
		}
	}
	fmt.Println("Memory dump completed!")
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
func plus4KernalRoutines() {
	// CHROUT routine is at $FFD2
	switch PC {
	case 0xFFB1:
		// This is a CHROUT call
		fmt.Printf("Call to CHROUT!!!!\n")
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
		resetCPU()
	}
	// print "kernal rom call address"
	//fmt.Printf("\n\u001B[32;5mKernal ROM call address: $%04X\u001B[0m\n", PC)
}
func printMachineState() {
	if BRKtrue || (previousOpcode != 0x20 && previousOpcode != 0x4C && previousOpcode != 0x6C) {
		if previousOpcode == 0x60 || (previousOpcode == 0x00 && BRKtrue || previousOpcode == 0x40) {
			fmt.Printf("| %04X | ", previousPC)
			fmt.Printf("%02X |        |", previousOpcode)
			previousOpcode = 0x00
			previousPC = 0x0000
			previousOperand1 = 0x00
			previousOperand2 = 0x00
			BRKtrue = false
		} else {
			fmt.Printf("| %04X | ", PC)
			// If opcode() is a 1 byte instruction, print opcode
			if opcode() == 0x08 || opcode() == 0x18 || opcode() == 0x28 || opcode() == 0x30 || opcode() == 0x38 || opcode() == 0x48 || opcode() == 0x58 || opcode() == 0x68 || opcode() == 0x78 || opcode() == 0x88 || opcode() == 0x8A || opcode() == 0x98 || opcode() == 0x9A || opcode() == 0xA8 || opcode() == 0xAA || opcode() == 0xB8 || opcode() == 0xBA || opcode() == 0xC8 || opcode() == 0xCA || opcode() == 0xD8 || opcode() == 0xDA || opcode() == 0xE8 || opcode() == 0xEA || opcode() == 0xF8 || opcode() == 0xFA || opcode() == 0x2A || opcode() == 0x6A {
				fmt.Printf("%02X |        |", opcode())
			}

			// If opcode() is a 2 byte instruction, print opcode and operand1
			// 		fmt.Printf("%02X %02X ", opcode(), operand1())
			// The 0x hex opcodes for the 2 byte instructions on the 6502 are
			if opcode() == 0x69 || opcode() == 0x29 || opcode() == 0xC9 || opcode() == 0xE0 || opcode() == 0xC0 || opcode() == 0x49 || opcode() == 0xA9 || opcode() == 0xA2 || opcode() == 0xA0 || opcode() == 0x09 || opcode() == 0xE9 || opcode() == 0x65 || opcode() == 0x25 || opcode() == 0x06 || opcode() == 0x24 || opcode() == 0xC5 || opcode() == 0xE4 || opcode() == 0xC4 || opcode() == 0xC6 || opcode() == 0x45 || opcode() == 0xE6 || opcode() == 0xA5 || opcode() == 0xA6 || opcode() == 0xA4 || opcode() == 0x46 || opcode() == 0x05 || opcode() == 0x26 || opcode() == 0x66 || opcode() == 0xE5 || opcode() == 0x85 || opcode() == 0x86 || opcode() == 0x84 || opcode() == 0x90 || opcode() == 0xB0 || opcode() == 0xF0 || opcode() == 0x30 || opcode() == 0xD0 || opcode() == 0x10 || opcode() == 0x50 || opcode() == 0x70 {
				fmt.Printf("%02X | %02X     |", opcode(), operand1())
			}

			// If opcode() is a 3 byte instruction, print opcode, operand1 and operand2
			// 			fmt.Printf("%02X %02X %02X ", opcode(), operand1(), operand2())
			if opcode() == 0x6D || opcode() == 0x2D || opcode() == 0x0E || opcode() == 0x2C || opcode() == 0xCD || opcode() == 0xEC || opcode() == 0xCC || opcode() == 0xCE || opcode() == 0x4D || opcode() == 0xEE || opcode() == 0xAD || opcode() == 0xAC || opcode() == 0xAE || opcode() == 0x4E || opcode() == 0x0D || opcode() == 0x2E || opcode() == 0x6E || opcode() == 0xED || opcode() == 0x8D || opcode() == 0x8E || opcode() == 0x8C || opcode() == 0x7D || opcode() == 0x79 || opcode() == 0x3D || opcode() == 0x39 || opcode() == 0x1E || opcode() == 0xDD || opcode() == 0xD9 || opcode() == 0xDE || opcode() == 0x5D || opcode() == 0x59 || opcode() == 0xFE || opcode() == 0xBD || opcode() == 0xB9 || opcode() == 0xBC || opcode() == 0xBE || opcode() == 0x5E || opcode() == 0x1D || opcode() == 0x19 || opcode() == 0x3E || opcode() == 0x7E || opcode() == 0xFD || opcode() == 0xF9 || opcode() == 0x9D || opcode() == 0x95 || opcode() == 0x99 || opcode() == 0xB5 || opcode() == 0x91 || opcode() == 0xB1 || opcode() == 0x81 || opcode() == 0xA1 || opcode() == 0x94 || opcode() == 0x96 || opcode() == 0xB4 || opcode() == 0xB6 || opcode() == 0x35 || opcode() == 0x15 || opcode() == 0x55 || opcode() == 0x21 || opcode() == 0x01 || opcode() == 0x41 || opcode() == 0x31 || opcode() == 0x11 || opcode() == 0x51 || opcode() == 0xF6 || opcode() == 0xD6 || opcode() == 0x4A || opcode() == 0x0A || opcode() == 0x16 || opcode() == 0x56 || opcode() == 0x36 || opcode() == 0x76 || opcode() == 0x75 || opcode() == 0xF5 || opcode() == 0xD5 || opcode() == 0xC1 || opcode() == 0xD1 || opcode() == 0x61 || opcode() == 0xE1 || opcode() == 0x71 || opcode() == 0xF1 {
				fmt.Printf("%02X | %02X %02X  |", opcode(), operand1(), operand2())
			}
		}
	} else if previousOpcode == 0x20 || previousOpcode == 0x4C || previousOpcode == 0x6C {
		fmt.Printf("| %04X | %02X | %02X %02X  |", previousPC, previousOpcode, previousOperand1, previousOperand2)
		previousOpcode = 0x00
		previousPC = 0x0000
		previousOperand1 = 0x00
		previousOperand2 = 0x00
	} else {
		fmt.Printf("| %04X | ", previousPC)
		fmt.Printf("%02X | %02X %02X  |", previousOpcode, previousOperand1, previousOperand2)

	}

	// Print disassembled instruction
	fmt.Printf("\t %s\t|", disassembledInstruction)
	// Print A,X,Y,SP as hex values
	fmt.Printf(" A:%02X X:%02X Y:%02X SP:$%04X |  $%04X  | ", A, X, Y, SPBaseAddress+SP, readStack())

	// Print full SR as binary digits with zero padding
	//fmt.Printf("%08b | ", SR)

	// Print N if SR bit 7 is 1 else print -
	if getSRBit(7) == 1 {
		fmt.Printf("N")
	} else {
		fmt.Printf("-")
	}
	// Print V if SR bit 6 is 1 else print -
	if getSRBit(6) == 1 {
		fmt.Printf("V")
	} else {
		fmt.Printf("-")
	}
	// Print - for SR bit 5
	fmt.Printf("-")
	// Print B if SR bit 4 is 1 else print -
	if getSRBit(4) == 1 {
		fmt.Printf("B")
	} else {
		fmt.Printf("-")
	}
	// Print D if SR bit 3 is 1 else print -
	if getSRBit(3) == 1 {
		fmt.Printf("D")
	} else {
		fmt.Printf("-")
	}
	// Print I if SR bit 2 is 1 else print -
	if getSRBit(2) == 1 {
		fmt.Printf("I")
	} else {
		fmt.Printf("-")
	}
	// Print Z if SR bit 1 is 1 else print -
	if getSRBit(1) == 1 {
		fmt.Printf("Z")
	} else {
		fmt.Printf("-")
	}
	// Print C if SR bit 0 is 1 else print -
	if getSRBit(0) == 1 {
		fmt.Printf("C")
	} else {
		fmt.Printf("-")
	}
	fmt.Printf(" | $%08X  | ", instructionCounter)
	fmt.Printf(" $%08X  | ", cycleCounter)
	if timeSpent == 0 {
		fmt.Printf("%v\t\t|\n", timeSpent)
	} else {
		fmt.Printf("%v\t|\n", timeSpent)

	}
	// Move cursor back to beginning of previous line
	// Comment this line out to get full disassembly and machine state
	fmt.Printf("\033[1A")
}
func disassembleOpcode() {
	if *disassemble {
		fmt.Printf("%s\n", disassembledInstruction)
	}
}

func handleState(amount int) {
	if *stateMonitor {
		printMachineState()
	}
	incPC(amount)
	// If amount is 0, then we are in a branch instruction and we don't want to increment the instruction counter
	if amount != 0 {
		instructionCounter++
	}
	if irq {
		handleIRQ()
	}
	if nmi {
		handleNMI()
	}
	if reset {
		handleRESET()
	}
}
