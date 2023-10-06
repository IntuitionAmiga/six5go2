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

	c64basicROMAddress  = 0xA000
	c64kernalROMAddress = 0xE000
	c64charROMAddress   = 0xD000

	plus4basicROMAddress  = 0x8000
	plus4kernalROMAddress = 0xC000
	plus4charROMAddress   = 0xC000
	threePlus1ROMAddress  = 0x8000

	AllSuiteAROMAddress              = 0x4000
	KlausDTestROMAddress             = 0x0000
	KlausDInfiniteLoopAddress uint16 = 0x062B
	RuudBTestROMAddress              = 0xE000
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
	switch cpu.PC {
	case 0xFFD2:
		// This is a CHROUT call
		fmt.Printf("Call to CHROUT!!!!\n")
		ch := petsciiToAscii(cpu.A) // Convert PETSCII to ASCII

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
		cpu.resetCPU()
	}
	// print "kernal rom call address"
	//fmt.Printf("\n\u001B[32;5mKernal ROM call address: $%04X\u001B[0m\n", PC)
}
func printMachineState() {
	if BRKtrue || (cpu.previousOpcode != JSR_ABSOLUTE_OPCODE && cpu.previousOpcode != JMP_ABSOLUTE_OPCODE && cpu.previousOpcode != JMP_INDIRECT_OPCODE) {
		if cpu.previousOpcode == RTS_OPCODE || (cpu.previousOpcode == BRK_OPCODE && BRKtrue || cpu.previousOpcode == RTI_OPCODE) {
			fmt.Printf("| %04X | ", cpu.previousPC)
			fmt.Printf("%02X |        |", cpu.previousOpcode)
			cpu.previousOpcode = 0x00
			cpu.previousPC = 0x0000
			cpu.previousOperand1 = 0x00
			cpu.previousOperand2 = 0x00
			BRKtrue = false
		} else {
			fmt.Printf("| %04X | ", cpu.PC)
			// If cpu.opcode() is a 1 byte instruction, print opcode
			if cpu.opcode() == CLC_OPCODE || cpu.opcode() == CLD_OPCODE || cpu.opcode() == CLI_OPCODE || cpu.opcode() == CLV_OPCODE || cpu.opcode() == DEX_OPCODE || cpu.opcode() == DEY_OPCODE || cpu.opcode() == INX_OPCODE || cpu.opcode() == INY_OPCODE || cpu.opcode() == NOP_OPCODE || cpu.opcode() == PHA_OPCODE || cpu.opcode() == PHP_OPCODE || cpu.opcode() == PLA_OPCODE || cpu.opcode() == PLP_OPCODE || cpu.opcode() == RTI_OPCODE || cpu.opcode() == RTS_OPCODE || cpu.opcode() == SEC_OPCODE || cpu.opcode() == SED_OPCODE || cpu.opcode() == SEI_OPCODE || cpu.opcode() == TAX_OPCODE || cpu.opcode() == TAY_OPCODE || cpu.opcode() == TSX_OPCODE || cpu.opcode() == TXA_OPCODE || cpu.opcode() == TXS_OPCODE || cpu.opcode() == TYA_OPCODE || cpu.opcode() == BRK_OPCODE {
				fmt.Printf("%02X |        |", cpu.opcode())
			}

			// If opcode() is a 2 byte instruction, print opcode and operand1
			// 		fmt.Printf("%02X %02X ", opcode(), operand1())
			// The 0x hex opcodes for the 2 byte instructions on the 6502 are
			if cpu.opcode() == ADC_IMMEDIATE_OPCODE || cpu.opcode() == AND_IMMEDIATE_OPCODE || cpu.opcode() == CMP_IMMEDIATE_OPCODE || cpu.opcode() == CPX_IMMEDIATE_OPCODE || cpu.opcode() == CPY_IMMEDIATE_OPCODE || cpu.opcode() == EOR_IMMEDIATE_OPCODE || cpu.opcode() == LDA_IMMEDIATE_OPCODE || cpu.opcode() == LDX_IMMEDIATE_OPCODE || cpu.opcode() == LDY_IMMEDIATE_OPCODE || cpu.opcode() == ORA_IMMEDIATE_OPCODE || cpu.opcode() == SBC_IMMEDIATE_OPCODE {
				fmt.Printf("%02X | %02X     |", cpu.opcode(), cpu.operand1())
			}

			// If opcode() is a 3 byte instruction, print opcode, operand1 and operand2
			// 			fmt.Printf("%02X %02X %02X ", opcode(), operand1(), operand2())
			if cpu.opcode() != CLC_OPCODE && cpu.opcode() != CLD_OPCODE && cpu.opcode() != CLI_OPCODE && cpu.opcode() != CLV_OPCODE && cpu.opcode() != DEX_OPCODE && cpu.opcode() != DEY_OPCODE && cpu.opcode() != INX_OPCODE && cpu.opcode() != INY_OPCODE && cpu.opcode() != NOP_OPCODE && cpu.opcode() != PHA_OPCODE && cpu.opcode() != PHP_OPCODE && cpu.opcode() != PLA_OPCODE && cpu.opcode() != PLP_OPCODE && cpu.opcode() != RTI_OPCODE && cpu.opcode() != RTS_OPCODE && cpu.opcode() != SEC_OPCODE && cpu.opcode() != SED_OPCODE && cpu.opcode() != SEI_OPCODE && cpu.opcode() != TAX_OPCODE && cpu.opcode() != TAY_OPCODE && cpu.opcode() != TSX_OPCODE && cpu.opcode() != TXA_OPCODE && cpu.opcode() != TXS_OPCODE && cpu.opcode() != TYA_OPCODE && cpu.opcode() != BRK_OPCODE && cpu.opcode() != ADC_IMMEDIATE_OPCODE && cpu.opcode() != AND_IMMEDIATE_OPCODE && cpu.opcode() != CMP_IMMEDIATE_OPCODE && cpu.opcode() != CPX_IMMEDIATE_OPCODE && cpu.opcode() != CPY_IMMEDIATE_OPCODE && cpu.opcode() != EOR_IMMEDIATE_OPCODE && cpu.opcode() != LDA_IMMEDIATE_OPCODE && cpu.opcode() != LDX_IMMEDIATE_OPCODE && cpu.opcode() != LDY_IMMEDIATE_OPCODE && cpu.opcode() != ORA_IMMEDIATE_OPCODE && cpu.opcode() != SBC_IMMEDIATE_OPCODE {
				fmt.Printf("%02X | %02X %02X  |", cpu.opcode(), cpu.operand1(), cpu.operand2())
			}
		}
	} else if cpu.previousOpcode == JSR_ABSOLUTE_OPCODE || cpu.previousOpcode == JMP_ABSOLUTE_OPCODE || cpu.previousOpcode == JMP_INDIRECT_OPCODE {
		fmt.Printf("| %04X | %02X | %02X %02X  |", cpu.previousPC, cpu.previousOpcode, cpu.previousOperand1, cpu.previousOperand2)
		cpu.previousOpcode = 0x00
		cpu.previousPC = 0x0000
		cpu.previousOperand1 = 0x00
		cpu.previousOperand2 = 0x00
	} else {
		fmt.Printf("| %04X | ", cpu.previousPC)
		fmt.Printf("%02X | %02X %02X  |", cpu.previousOpcode, cpu.previousOperand1, cpu.previousOperand2)

	}

	// Print disassembled instruction
	fmt.Printf("\t %s\t|", disassembledInstruction)
	// Print A,X,Y,SP as hex values
	fmt.Printf(" A:%02X X:%02X Y:%02X SP:$%04X |  $%04X  | ", cpu.A, cpu.X, cpu.Y, SPBaseAddress+cpu.SP, readStack())

	// Print full SR as binary digits with zero padding
	//fmt.Printf("%08b | ", SR)

	// Print N if SR bit 7 is 1 else print -
	if cpu.getSRBit(7) == 1 {
		fmt.Printf("N")
	} else {
		fmt.Printf("-")
	}
	// Print V if SR bit 6 is 1 else print -
	if cpu.getSRBit(6) == 1 {
		fmt.Printf("V")
	} else {
		fmt.Printf("-")
	}
	// Print - for SR bit 5
	fmt.Printf("-")
	// Print B if SR bit 4 is 1 else print -
	if cpu.getSRBit(4) == 1 {
		fmt.Printf("B")
	} else {
		fmt.Printf("-")
	}
	// Print D if SR bit 3 is 1 else print -
	if cpu.getSRBit(3) == 1 {
		fmt.Printf("D")
	} else {
		fmt.Printf("-")
	}
	// Print I if SR bit 2 is 1 else print -
	if cpu.getSRBit(2) == 1 {
		fmt.Printf("I")
	} else {
		fmt.Printf("-")
	}
	// Print Z if SR bit 1 is 1 else print -
	if cpu.getSRBit(1) == 1 {
		fmt.Printf("Z")
	} else {
		fmt.Printf("-")
	}
	// Print C if SR bit 0 is 1 else print -
	if cpu.getSRBit(0) == 1 {
		fmt.Printf("C")
	} else {
		fmt.Printf("-")
	}
	fmt.Printf(" | $%08X  | ", instructionCounter)
	fmt.Printf(" $%08X  | ", cpu.cycleCounter)
	if cpu.cpuTimeSpent == 0 {
		fmt.Printf("%v\t\t|\n", cpu.cpuTimeSpent)
	} else {
		fmt.Printf("%v\t|\n", cpu.cpuTimeSpent)

	}
	// Move cursor back to beginning of previous line
	// Comment this line out to get full disassembly and machine state
	//fmt.Printf("\033[1A")
}
func disassembleOpcode() {
	if *disassemble {
		fmt.Printf("%s\n", disassembledInstruction)
	}
}

func boilerPlate() {
	// Clear the screen and move cursor to top left
	fmt.Printf("\033[2J")
	fmt.Printf("\033[0;0H")

	fmt.Printf("Six5go2 v2.0 - 6502 Emulator and Disassembler in Golang (c) 2022-2023 Zayn Otley\n\n")
	fmt.Printf("https://github.com/intuitionamiga/six5go2/tree/v2\n\n")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Printf("Size of addressable memory is %v ($%04X) bytes\n\n", len(memory), len(memory)-1)
}
