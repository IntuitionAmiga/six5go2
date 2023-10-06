package main

import "fmt"

// Constants for MMIO address ranges and TED specific addresses
const (
	NMIVectorAddress   = 0xFFFA
	RESETVectorAddress = 0xFFFC

	TED_REG_START = 0xFD00
	TED_REG_END   = 0xFD19
)

const IRQVectorAddress uint16 = 0xFFFE

func readMemory(address uint16) byte {
	if address >= TED_REG_START && address <= TED_REG_END {
		return ted.readTEDReg(address)
	}
	return memory[address]
}
func writeMemory(address uint16, value byte) {
	fmt.Printf("IRQVectorAddress: %04X\n", IRQVectorAddress)
	fmt.Printf("Content of memory at IRQVectorAddress: %04X\n", memory[IRQVectorAddress])
	fmt.Printf("Content of memory at IRQVectorAddress+1: %04X\n", memory[IRQVectorAddress+1])

	if address == IRQVectorAddress {
		fmt.Println("Interrupt vector %04X written to with value: %04X", address, value)
		cpu.irq = true
		fmt.Println("IRQ request!")
		// or breakpoint()
	}
	if address >= TED_REG_START && address <= TED_REG_END {
		ted.writeTEDReg(address, value)
	} else {
		memory[address] = value
	}
	// Existing special address checks
	if address == IRQVectorAddress {
		cpu.irq = true // Signal an IRQ
	}
	if address == NMIVectorAddress {
		cpu.nmi = true // Signal an NMI
	}
	if address == RESETVectorAddress {
		cpu.reset = true // Signal a RESET
	}
	memory[address] = value
}

func readBit(bit byte, value byte) int {
	// Read bit from value and return it
	return int((value >> bit) & 1)
}
func readStack() byte {
	return readMemory(SPBaseAddress + cpu.SP)
}
func updateStack(value byte) {
	writeMemory(SPBaseAddress+cpu.SP, value)
}
