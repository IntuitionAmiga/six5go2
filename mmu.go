package main

import (
	"fmt"
	"io"
	"os"
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
