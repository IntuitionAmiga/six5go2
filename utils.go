package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	allsuitea   = flag.Bool("allsuitea", false, "AllSuiteA ROM")
	klausd      = flag.Bool("klausd", false, "Klaus Dormann's 6502 functional test ROM")
	plus4       = flag.Bool("plus4", false, "Plus/4 ROMs")
	ruudb       = flag.Bool("ruudb", false, "RuudB's 8K Test ROM")
	c64         = flag.Bool("c64", false, "C64 ROMs")
	disassemble = flag.Bool("dis", false, "Disassembler mode (Optional)")
	traceLog    = flag.Bool("trace", false, "State monitor mode (Optional)")

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
	if *traceLog {
		f, err := os.OpenFile("trace.txt", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(f, "|  PC  | OP |OPERANDS|    DISASSEMBLY   | \t REGISTERS\t  |  STACK  | SR FLAGS | INST COUNT | CYCLE COUNT | TIME SPENT  |\n")
		fmt.Fprintf(f, "|------|----|--------|------------------|-------------------------|---------|----------|------------|-------------|-------------|\n")
		f.Sync()
		f.Close()
	}
}
func dumpMemoryToFile(memory [65536]byte) {
	file, err := os.Create("memorydump.txt")
	if err != nil {
		//fmt.println("Error creating file:", err)
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
	//fmt.println("Memory dump completed!")
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

func executionTraceLog() {
	f, err := os.OpenFile("trace.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//cpuMutex.Lock()         // Lock before reading
	//defer cpuMutex.Unlock() // Unlock after done
	switch {
	case BRKtrue, cpu.previousOpcode != JSR_ABSOLUTE_OPCODE && cpu.previousOpcode != JMP_ABSOLUTE_OPCODE && cpu.previousOpcode != JMP_INDIRECT_OPCODE:
		switch {
		case cpu.previousOpcode == RTS_OPCODE, cpu.previousOpcode == BRK_OPCODE && BRKtrue, cpu.previousOpcode == RTI_OPCODE:
			fmt.Fprintf(f, "| %04X | ", cpu.previousPC)
			fmt.Fprintf(f, "%02X |        |", cpu.previousOpcode)
			cpu.previousOpcode = 0x00
			cpu.previousPC = 0x0000
			cpu.previousOperand1 = 0x00
			cpu.previousOperand2 = 0x00
			BRKtrue = false
		default:
			fmt.Fprintf(f, "| %04X | ", cpu.PC)
			switch cpu.opcode() {
			case CLC_OPCODE, CLD_OPCODE, CLI_OPCODE, CLV_OPCODE, DEX_OPCODE, DEY_OPCODE, INX_OPCODE, INY_OPCODE, NOP_OPCODE, PHA_OPCODE, PHP_OPCODE, PLA_OPCODE, PLP_OPCODE, RTI_OPCODE, RTS_OPCODE, SEC_OPCODE, SED_OPCODE, SEI_OPCODE, TAX_OPCODE, TAY_OPCODE, TSX_OPCODE, TXA_OPCODE, TXS_OPCODE, TYA_OPCODE, BRK_OPCODE:
				fmt.Fprintf(f, "%02X |        |", cpu.opcode())
			case ADC_IMMEDIATE_OPCODE, AND_IMMEDIATE_OPCODE, CMP_IMMEDIATE_OPCODE, CPX_IMMEDIATE_OPCODE, CPY_IMMEDIATE_OPCODE, EOR_IMMEDIATE_OPCODE, LDA_IMMEDIATE_OPCODE, LDX_IMMEDIATE_OPCODE, LDY_IMMEDIATE_OPCODE, ORA_IMMEDIATE_OPCODE, SBC_IMMEDIATE_OPCODE:
				fmt.Fprintf(f, "%02X | %02X     |", cpu.opcode(), cpu.operand1())
			default:
				fmt.Fprintf(f, "%02X | %02X %02X  |", cpu.opcode(), cpu.operand1(), cpu.operand2())
			}
		}
	case cpu.previousOpcode == JSR_ABSOLUTE_OPCODE, cpu.previousOpcode == JMP_ABSOLUTE_OPCODE, cpu.previousOpcode == JMP_INDIRECT_OPCODE:
		fmt.Fprintf(f, "| %04X | %02X | %02X %02X  |", cpu.previousPC, cpu.previousOpcode, cpu.previousOperand1, cpu.previousOperand2)
		cpu.previousOpcode = 0x00
		cpu.previousPC = 0x0000
		cpu.previousOperand1 = 0x00
		cpu.previousOperand2 = 0x00
	default:
		fmt.Fprintf(f, "| %04X | ", cpu.previousPC)
		fmt.Fprintf(f, "%02X | %02X %02X  |", cpu.previousOpcode, cpu.previousOperand1, cpu.previousOperand2)
	}

	// Print disassembled instruction
	fmt.Fprintf(f, "\t %s\t|", disassembledInstruction)
	// Print A,X,Y,SP as hex values
	fmt.Fprintf(f, " A:%02X X:%02X Y:%02X SP:$%04X |  $%04X  | ", cpu.A, cpu.X, cpu.Y, SPBaseAddress+cpu.SP, readStack())

	// Print full SR as binary digits with zero padding
	//fmt.Fprintf(f,"%08b | ", SR)

	// Print N if SR bit 7 is 1 else print -
	if cpu.getSRBit(7) == 1 {
		fmt.Fprintf(f, "N")
	} else {
		fmt.Fprintf(f, "-")
	}
	// Print V if SR bit 6 is 1 else print -
	if cpu.getSRBit(6) == 1 {
		fmt.Fprintf(f, "V")
	} else {
		fmt.Fprintf(f, "-")
	}
	// Print - for SR bit 5
	fmt.Fprintf(f, "-")
	// Print B if SR bit 4 is 1 else print -
	if cpu.getSRBit(4) == 1 {
		fmt.Fprintf(f, "B")
	} else {
		fmt.Fprintf(f, "-")
	}
	// Print D if SR bit 3 is 1 else print -
	if cpu.getSRBit(3) == 1 {
		fmt.Fprintf(f, "D")
	} else {
		fmt.Fprintf(f, "-")
	}
	// Print I if SR bit 2 is 1 else print -
	if cpu.getSRBit(2) == 1 {
		fmt.Fprintf(f, "I")
	} else {
		fmt.Fprintf(f, "-")
	}
	// Print Z if SR bit 1 is 1 else print -
	if cpu.getSRBit(1) == 1 {
		fmt.Fprintf(f, "Z")
	} else {
		fmt.Fprintf(f, "-")
	}
	// Print C if SR bit 0 is 1 else print -
	if cpu.getSRBit(0) == 1 {
		fmt.Fprintf(f, "C")
	} else {
		fmt.Fprintf(f, "-")
	}
	fmt.Fprintf(f, " | $%08X  | ", instructionCounter)
	fmt.Fprintf(f, " $%08X  | ", cpu.cycleCounter)
	if cpu.cpuTimeSpent == 0 {
		fmt.Fprintf(f, "%v\t\t|\n", cpu.cpuTimeSpent)
	} else {
		fmt.Fprintf(f, "%v\t|\n", cpu.cpuTimeSpent)
	}
	defer f.Close()
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

func executionTrace() string {
	var traceLine string
	//cpuMutex.Lock()         // Lock before reading
	//defer cpuMutex.Unlock() // Unlock after done
	switch {
	case BRKtrue, cpu.previousOpcode != JSR_ABSOLUTE_OPCODE && cpu.previousOpcode != JMP_ABSOLUTE_OPCODE && cpu.previousOpcode != JMP_INDIRECT_OPCODE:
		switch {
		case cpu.previousOpcode == RTS_OPCODE, cpu.previousOpcode == BRK_OPCODE && BRKtrue, cpu.previousOpcode == RTI_OPCODE:
			traceLine = fmt.Sprintf("$%04X ", cpu.previousPC)
			traceLine += fmt.Sprintf("%02X ", cpu.previousOpcode)
			cpu.previousOpcode = 0x00
			cpu.previousPC = 0x0000
			cpu.previousOperand1 = 0x00
			cpu.previousOperand2 = 0x00
			BRKtrue = false
		default:
			traceLine += fmt.Sprintf("$%04X ", cpu.PC)
			switch cpu.opcode() {
			case CLC_OPCODE, CLD_OPCODE, CLI_OPCODE, CLV_OPCODE, DEX_OPCODE, DEY_OPCODE, INX_OPCODE, INY_OPCODE, NOP_OPCODE, PHA_OPCODE, PHP_OPCODE, PLA_OPCODE, PLP_OPCODE, RTI_OPCODE, RTS_OPCODE, SEC_OPCODE, SED_OPCODE, SEI_OPCODE, TAX_OPCODE, TAY_OPCODE, TSX_OPCODE, TXA_OPCODE, TXS_OPCODE, TYA_OPCODE, BRK_OPCODE:
				traceLine += fmt.Sprintf("%02X ", cpu.opcode())
			case ADC_IMMEDIATE_OPCODE, AND_IMMEDIATE_OPCODE, CMP_IMMEDIATE_OPCODE, CPX_IMMEDIATE_OPCODE, CPY_IMMEDIATE_OPCODE, EOR_IMMEDIATE_OPCODE, LDA_IMMEDIATE_OPCODE, LDX_IMMEDIATE_OPCODE, LDY_IMMEDIATE_OPCODE, ORA_IMMEDIATE_OPCODE, SBC_IMMEDIATE_OPCODE:
				traceLine += fmt.Sprintf("%02X %02X ", cpu.opcode(), cpu.operand1())
			default:
				traceLine += fmt.Sprintf("%02X %02X %02X ", cpu.opcode(), cpu.operand1(), cpu.operand2())
			}
		}
	case cpu.previousOpcode == JSR_ABSOLUTE_OPCODE, cpu.previousOpcode == JMP_ABSOLUTE_OPCODE, cpu.previousOpcode == JMP_INDIRECT_OPCODE:
		traceLine += fmt.Sprintf("$%04X %02X %02X %02X ", cpu.previousPC, cpu.previousOpcode, cpu.previousOperand1, cpu.previousOperand2)
		cpu.previousOpcode = 0x00
		cpu.previousPC = 0x0000
		cpu.previousOperand1 = 0x00
		cpu.previousOperand2 = 0x00
	default:
		traceLine += fmt.Sprintf("%04X ", cpu.previousPC)
		traceLine += fmt.Sprintf("%02X %02X %02X ", cpu.previousOpcode, cpu.previousOperand1, cpu.previousOperand2)
	}
	if *traceLog {
		executionTraceLog()
	}
	return traceLine
}

func instructionCount() string {
	var instructionCountLine string
	instructionCountLine = fmt.Sprintf("Instructions:\t$%08X ", instructionCounter)
	return instructionCountLine
}

func statusFlags() string {
	var statusLine string
	// Print N if SR bit 7 is 1 else print -
	if cpu.getSRBit(7) == 1 {
		//fmt.Printf("N")
		statusLine += fmt.Sprintf("N")
	} else {
		//fmt.Printf("-")
		statusLine += fmt.Sprintf("-")
	}
	// Print V if SR bit 6 is 1 else print -
	if cpu.getSRBit(6) == 1 {
		//fmt.Printf("V")
		statusLine += fmt.Sprintf("V")
	} else {
		//fmt.Printf("-")
		statusLine += fmt.Sprintf("-")
	}
	// Print - for SR bit 5
	//fmt.Printf("-")
	statusLine += fmt.Sprintf("-")
	// Print B if SR bit 4 is 1 else print -
	if cpu.getSRBit(4) == 1 {
		//fmt.Printf("B")
		statusLine += fmt.Sprintf("B")
	} else {
		//fmt.Printf("-")
		statusLine += fmt.Sprintf("-")
	}
	// Print D if SR bit 3 is 1 else print -
	if cpu.getSRBit(3) == 1 {
		//fmt.Printf("D")
		statusLine += fmt.Sprintf("D")
	} else {
		//fmt.Printf("-")
		statusLine += fmt.Sprintf("-")
	}
	// Print I if SR bit 2 is 1 else print -
	if cpu.getSRBit(2) == 1 {
		//fmt.Printf("I")
		statusLine += fmt.Sprintf("I")
	} else {
		//fmt.Printf("-")
		statusLine += fmt.Sprintf("-")
	}
	// Print Z if SR bit 1 is 1 else print -
	if cpu.getSRBit(1) == 1 {
		//fmt.Printf("Z")
		statusLine += fmt.Sprintf("Z")
	} else {
		//fmt.Printf("-")
		statusLine += fmt.Sprintf("-")
	}
	// Print C if SR bit 0 is 1 else print -
	if cpu.getSRBit(0) == 1 {
		//fmt.Printf("C")
		statusLine += fmt.Sprintf("C")
	} else {
		//fmt.Printf("-")
		statusLine += fmt.Sprintf("-")
	}
	return statusLine
}
