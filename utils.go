package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	allsuitea = flag.Bool("allsuitea", false, "AllSuiteA ROM")
	klausd    = flag.Bool("klausd", false, "Klaus Dormann's 6502 functional test ROM")
	plus4     = flag.Bool("plus4", false, "Plus/4 ROMs")
	ruudb     = flag.Bool("ruudb", false, "RuudB's 8K Test ROM")
	c64       = flag.Bool("c64", false, "C64 ROMs")
	traceLog  = flag.Bool("trace", false, "State monitor mode (Optional)")
	diag      = flag.Bool("diag", false, "Diag264 ROM")

	plus4ROMSize = 16384 // 8KB

	C64BASICROM  = make([]byte, 8192)
	C64KERNALROM = make([]byte, 8192)
	C64CHARROM   = make([]byte, 4096)

	PLUS4BASICROM  = make([]byte, plus4ROMSize)
	PLUS4KERNALROM = make([]byte, plus4ROMSize)
	PLUS4CHARROM   = make([]byte, 4096)
	THREEPLUS1ROM  = make([]byte, 16384)

	AllSuiteAROM  = make([]byte, 16384)
	KlausDTestROM = make([]byte, 65536)
	RuudBTestROM  = make([]byte, 8192)

	c64basicROMAddress  = 0xA000
	c64kernalROMAddress = 0xE000
	c64charROMAddress   = 0xD000

	plus4basicROMAddress     = 0x8000
	plus4kernalROMAddress    = 0xC000
	plus4charROMAddress      = 0xF800
	threePlus1ROMAddressLow  = 0x8000
	threePlus1ROMAddressHigh = 0xA000

	AllSuiteAROMAddress              = 0x4000
	KlausDTestROMAddress             = 0x0000
	KlausDInfiniteLoopAddress uint16 = 0x3469
	RuudBTestROMAddress              = 0xE000
)

func loadROMs() {
	if *plus4 {
		// Load BASIC ROM
		basicfile, _ := os.Open("roms/plus4/basic")
		_, _ = io.ReadFull(basicfile, PLUS4BASICROM)
		copy(memory[plus4basicROMAddress:], PLUS4BASICROM[:])
		fmt.Printf("Copying BASIC ROM into memory at $%04X\n", plus4basicROMAddress)
		_ = basicfile.Close()

		// Load KERNAL ROM
		kernelfile, _ := os.Open("roms/plus4/kernal")
		_, _ = io.ReadFull(kernelfile, PLUS4KERNALROM)
		copy(memory[plus4kernalROMAddress:], PLUS4KERNALROM[:])
		fmt.Printf("Copying KERNAL ROM into memory at $%04X\n", plus4kernalROMAddress)
		_ = kernelfile.Close()
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
		cpu.writeMemory(IRQVectorAddressLow, 0x00)  // Low byte of 0x4000
		cpu.writeMemory(IRQVectorAddressHigh, 0x40) // High byte of 0x4000
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
		cpu.writeMemory(IRQVectorAddressLow, 0x00)  // Low byte of 0x4000
		cpu.writeMemory(IRQVectorAddressHigh, 0x40) // High byte of 0x4000
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

	if *diag {
		kernelfile, _ := os.Open("roms/plus4/diag264_097_pal_kernal.bin")
		_, _ = io.ReadFull(kernelfile, PLUS4KERNALROM)
		copy(memory[plus4kernalROMAddress:], PLUS4KERNALROM[:])
		fmt.Printf("Copying 264Diag as KERNAL ROM into memory at $%04X\n", plus4kernalROMAddress)
		_ = kernelfile.Close()
	}
	if *traceLog {
		f, err := os.OpenFile("trace.txt", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(f, "|   PC  | OP |OPERANDS|   DISASSEMBLY   |        REGISTERS        |  STACK  | SR FLAGS | INST | CYCLE |  TIME   |\n")
		fmt.Fprintf(f, "|-------|----|--------|-----------------|-------------------------|---------|----------|------|-------|---------|\n")
		f.Sync()
		f.Close()
	}
}
func dumpMemoryToFile(startAddress, numBytes int) {
	file, err := os.Create("memorydump.txt")
	if err != nil {
		// Handle error
		return
	}
	defer file.Close()

	endAddress := startAddress + numBytes
	if endAddress > len(memory) {
		endAddress = len(memory)
	}

	for i := startAddress; i < endAddress; i++ {
		// Print the address at the start of each line
		if (i-startAddress)%16 == 0 {
			if i > startAddress {
				_, _ = fmt.Fprintln(file) // Newline at the end of a line
			}
			_, _ = fmt.Fprintf(file, "$%04X: ", i)
		}

		// Print the byte value in hex
		_, err := fmt.Fprintf(file, "%02X ", memory[i])
		if err != nil {
			// Handle error
			return
		}

		// Add ASCII representation at the end of every 16 bytes
		if (i-startAddress)%16 == 15 {
			_, _ = fmt.Fprint(file, " ")
			for j := i - 15; j <= i; j++ {
				asciiChar := petsciiToAscii(memory[j])
				if asciiChar >= 32 && asciiChar <= 126 {
					_, _ = fmt.Fprintf(file, "%c", asciiChar)
				} else {
					_, _ = fmt.Fprint(file, ".")
				}
			}
		}
	}

	// Handle last line if it's not a full line of 16 bytes
	remaining := (endAddress - startAddress) % 16
	if remaining != 0 {
		for i := 0; i < 16-remaining; i++ {
			_, _ = fmt.Fprint(file, "   ") // Align ASCII representation
		}
		_, _ = fmt.Fprint(file, " ")
		for i := endAddress - remaining; i < endAddress; i++ {
			asciiChar := petsciiToAscii(memory[i])
			if asciiChar >= 32 && asciiChar <= 126 {
				_, _ = fmt.Fprintf(file, "%c", asciiChar)
			} else {
				_, _ = fmt.Fprint(file, ".")
			}
		}
	}

	// Add a newline at the end of the file
	_, _ = fmt.Fprintln(file)
}
func petsciiToAscii(petscii uint8) uint8 {
	// PETSCII uppercase (A-Z) and lowercase (a-z) to ASCII
	if (petscii >= 65 && petscii <= 90) || (petscii >= 193 && petscii <= 218) {
		return petscii & 0x7F // Unset the 6th bit to convert to ASCII
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
	//fmt.Printf("\n\u001B[32;5mKernal ROM call address: $%04X\u001B[0m\n", PC)
}

func boilerPlate() {
	os.Remove("trace.txt")
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
	cpu.traceLine = fmt.Sprintf("| $%04X | %02X | ", cpu.preOpPC, cpu.preOpOpcode)
	cpu.nextTraceLine = fmt.Sprintf("| $%04X | %02X | ", cpu.PC, cpu.opcode())
	if oneByteInstructions[cpu.preOpOpcode] {
		cpu.disassembledInstruction = getMnemonic(cpu.preOpOpcode) + "\t" // Add tabs for formatting
	}
	if twoByteInstructions[cpu.preOpOpcode] {
		cpu.disassembledInstruction = getMnemonic(cpu.preOpOpcode) + "  " // Add tab for formatting
	}
	if threeByteInstructions[cpu.preOpOpcode] {
		cpu.disassembledInstruction = getMnemonic(cpu.preOpOpcode)
	}
	switch {
	case oneByteInstructions[cpu.preOpOpcode]:
		cpu.traceLine += "       |"
		cpu.nextTraceLine += "	   |"
	case twoByteInstructions[cpu.preOpOpcode]:
		cpu.traceLine += fmt.Sprintf("%02X     |", cpu.preOpOperand1)
		cpu.nextTraceLine += fmt.Sprintf("%02X     |", cpu.operand1())
	case threeByteInstructions[cpu.preOpOpcode]:
		cpu.traceLine += fmt.Sprintf("%02X %02X  |", cpu.preOpOperand1, cpu.preOpOperand2)
		cpu.nextTraceLine += fmt.Sprintf("%02X %02X  |", cpu.operand1(), cpu.operand2())

	}
	if *traceLog {
		writeTraceToFile(cpu.traceLine, cpu.disassembledInstruction, cpu.A, cpu.X, cpu.Y, SPBaseAddress+cpu.SP, cpu.readStack())
	}
	return cpu.traceLine
}
func writeTraceToFile(traceLine, disassembledInstruction string, A, X, Y byte, SP uint16, stackValue byte) {
	// Open the file for appending. Create it if it doesn't exist.
	f, err := os.OpenFile("trace.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open trace file: %v", err)
	}
	defer f.Close()
	// Create the full trace line with additional information including SR flags
	fullTraceLine := fmt.Sprintf("%s\t%s\t| A:%02X X:%02X Y:%02X SP:$%04X |  $%04X  | %s | %04X | %04X  | %v\t|\n",
		traceLine, disassembledInstruction, A, X, Y, SP, stackValue, getSRFlags(), cpu.instructionCounter, cpu.cycleCounter, cpu.cpuTimeSpent)
	// Write the full trace line to the file
	if _, err := f.WriteString(fullTraceLine); err != nil {
		log.Fatalf("Failed to write to trace file: %v", err)
	}
}
func getSRFlags() string {
	// Using a string builder to efficiently build the SR flags string
	var SRFlagsBuilder strings.Builder
	// Append the N flag or '-' accordingly
	if cpu.getSRBit(7) == 1 {
		SRFlagsBuilder.WriteString("N")
	} else {
		SRFlagsBuilder.WriteString("-")
	}
	// Append the V flag or '-' accordingly
	if cpu.getSRBit(6) == 1 {
		SRFlagsBuilder.WriteString("V")
	} else {
		SRFlagsBuilder.WriteString("-")
	}
	// Append '-' for bit 5 (which is not used)
	SRFlagsBuilder.WriteString("-")
	// Append the B flag or '-' accordingly
	if cpu.getSRBit(4) == 1 {
		SRFlagsBuilder.WriteString("B")
	} else {
		SRFlagsBuilder.WriteString("-")
	}
	// Append the D flag or '-' accordingly
	if cpu.getSRBit(3) == 1 {
		SRFlagsBuilder.WriteString("D")
	} else {
		SRFlagsBuilder.WriteString("-")
	}
	// Append the I flag or '-' accordingly
	if cpu.getSRBit(2) == 1 {
		SRFlagsBuilder.WriteString("I")
	} else {
		SRFlagsBuilder.WriteString("-")
	}
	// Append the Z flag or '-' accordingly
	if cpu.getSRBit(1) == 1 {
		SRFlagsBuilder.WriteString("Z")
	} else {
		SRFlagsBuilder.WriteString("-")
	}
	// Append the C flag or '-' accordingly
	if cpu.getSRBit(0) == 1 {
		SRFlagsBuilder.WriteString("C")
	} else {
		SRFlagsBuilder.WriteString("-")
	}
	return SRFlagsBuilder.String()
}
func instructionCount() string {
	var instructionCountLine string
	instructionCountLine = fmt.Sprintf("Instructions:\t$%08X ", cpu.instructionCounter)
	return instructionCountLine
}
func statusFlags() string {
	var statusLine string
	// Print N if SR bit 7 is 1 else print -
	if cpu.getSRBit(7) == 1 {
		statusLine += fmt.Sprintf("N")
	} else {
		statusLine += fmt.Sprintf("-")
	}
	// Print V if SR bit 6 is 1 else print -
	if cpu.getSRBit(6) == 1 {
		statusLine += fmt.Sprintf("V")
	} else {
		statusLine += fmt.Sprintf("-")
	}
	// Print - for SR bit 5
	statusLine += fmt.Sprintf("-")
	// Print B if SR bit 4 is 1 else print -
	if cpu.getSRBit(4) == 1 {
		statusLine += fmt.Sprintf("B")
	} else {
		statusLine += fmt.Sprintf("-")
	}
	// Print D if SR bit 3 is 1 else print -
	if cpu.getSRBit(3) == 1 {
		statusLine += fmt.Sprintf("D")
	} else {
		statusLine += fmt.Sprintf("-")
	}
	// Print I if SR bit 2 is 1 else print -
	if cpu.getSRBit(2) == 1 {
		statusLine += fmt.Sprintf("I")
	} else {
		statusLine += fmt.Sprintf("-")
	}
	// Print Z if SR bit 1 is 1 else print -
	if cpu.getSRBit(1) == 1 {
		statusLine += fmt.Sprintf("Z")
	} else {
		statusLine += fmt.Sprintf("-")
	}
	// Print C if SR bit 0 is 1 else print -
	if cpu.getSRBit(0) == 1 {
		statusLine += fmt.Sprintf("C")
	} else {
		statusLine += fmt.Sprintf("-")
	}
	return statusLine
}
func getMnemonic(opcode byte) string {
	mnemonic, exists := opcodeMnemonics[opcode]
	if !exists {
		return "???"
	}
	var address uint16
	// Determine the addressing mode and format the full disassembly accordingly
	switch opcode {
	// Immediate
	case 0x69, 0x29, 0xC9, 0xE0, 0xC0, 0x49, 0xA9, 0xA2, 0xA0, 0x09, 0xE9:
		return fmt.Sprintf("%s #$%02X", mnemonic, cpu.preOpOperand1)
	// Zero Page
	case 0x65, 0x25, 0x05, 0x24, 0xA5, 0xA6, 0xA4, 0x06, 0x46, 0xE6, 0xC6, 0x85, 0x86, 0x84, 0xC5:
		return fmt.Sprintf("%s $%02X", mnemonic, cpu.preOpOperand1)
	// Zero Page,X
	case 0x75, 0x35, 0x15, 0xB5, 0xB4, 0x16, 0x56, 0xF6, 0xD6, 0x95, 0x94:
		return fmt.Sprintf("%s $%02X,X", mnemonic, cpu.preOpOperand1)
	// Zero Page,Y
	case 0xB6, 0x96:
		return fmt.Sprintf("%s $%02X,Y", mnemonic, cpu.preOpOperand1)
	// Absolute
	case 0x6D, 0x2D, 0x0D, 0x2C, 0xAD, 0xAE, 0xAC, 0x0E, 0x4E, 0xEE, 0xCE, 0x8D, 0x8E, 0x8C, 0x4C, 0x20, 0xCD:
		address = uint16(cpu.preOpOperand2)<<8 | uint16(cpu.preOpOperand1)
		return fmt.Sprintf("%s $%04X", mnemonic, address)
	// Absolute,X
	case 0x7D, 0x3D, 0x1D, 0xBD, 0xBC, 0x1E, 0x5E, 0xFE, 0xDE, 0x9D, 0xDD:
		address = uint16(cpu.preOpOperand2)<<8 | uint16(cpu.preOpOperand1)
		return fmt.Sprintf("%s $%04X,X", mnemonic, address)
	// Absolute,Y
	case 0x79, 0x39, 0x19, 0xB9, 0xBE, 0x99, 0x91, 0xD9:
		address = uint16(cpu.preOpOperand2)<<8 | uint16(cpu.preOpOperand1)
		return fmt.Sprintf("%s $%04X,Y", mnemonic, address)
	// Indirect,X
	case 0x61, 0x21, 0x41, 0xA1, 0x01, 0xE1:
		return fmt.Sprintf("%s ($%02X,X)", mnemonic, cpu.preOpOperand1)
	// Indirect,Y
	case 0x71, 0x31, 0x51, 0xB1, 0x11, 0xF1:
		return fmt.Sprintf("%s ($%02X),Y", mnemonic, cpu.preOpOperand1)
	// Relative (for branch instructions)
	case 0x90, 0xB0, 0xF0, 0x30, 0xD0, 0x10, 0x50, 0x70:
		return fmt.Sprintf("%s $%02X", mnemonic, cpu.preOpOperand1)
	// Indirect (only JMP has this addressing mode)
	case 0x6C:
		address = uint16(cpu.preOpOperand2)<<8 | uint16(cpu.preOpOperand1)
		return fmt.Sprintf("%s ($%04X)", mnemonic, address)
	// Accumulator or Implied
	case 0x0A, 0x4A, 0x2A, 0x6A, 0xEA, 0x40, 0x60, 0x18, 0xD8, 0x58, 0xB8, 0xCA, 0x88, 0xE8, 0xC8, 0x48, 0x08, 0x68, 0x28, 0x78, 0xAA, 0xA8, 0xBA, 0x8A, 0x9A, 0x98:
		if opcode == 0x0A || opcode == 0x4A || opcode == 0x2A || opcode == 0x6A {
			return fmt.Sprintf("%s A", mnemonic)
		}
		return mnemonic // Implied
	default:
		return fmt.Sprintf("%s ???", mnemonic)
	}
}
