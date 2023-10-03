package main

// Constants for MMIO address ranges and TED specific addresses
const (
	NMIVectorAddress   = 0xFFFA
	RESETVectorAddress = 0xFFFC

	MMIO_START = 0xFD00
	MMIO_END   = 0xFFFF

	TED_REG_START = 0xFD00
	TED_REG_END   = 0xFD19
)

var IRQVectorAddress uint16 = 0xFFFE

func readMemory(address uint16) byte {
	if address >= MMIO_START && address <= MMIO_END {
		//fmt.Printf("MMIO Read from: %X\n", address) // Log MMIO read
		return readMMIO(address)
	}
	return memory[address]
}
func writeMemory(address uint16, value byte) {
	if address >= MMIO_START && address <= MMIO_END {
		// Just log for now
		//fmt.Printf("MMIO Write to: %X\n", address)
		writeMMIO(address, value)
		//fmt.Printf("\nMMIO Write to: $%04X\n", address)
		return
	}
	// Existing special address checks
	if address == IRQVectorAddress {
		irq = true // Signal an IRQ
	}
	if address == NMIVectorAddress {
		nmi = true // Signal an NMI
	}
	if address == RESETVectorAddress {
		reset = true // Signal a RESET
	}
	memory[address] = value
}

func readMMIO(address uint16) byte {
	switch {
	case address == RESETVectorAddress, address == 0xFFFD, address == IRQVectorAddress, address == 0xFFFF:
		return memory[address]
	case address >= TED_REG_START && address <= TED_REG_END:
		return readTEDReg(address)
	default:
		return 0
	}
}

func writeMMIO(address uint16, value byte) {
	switch {
	case address >= TED_REG_START && address <= TED_REG_END:
		writeTEDReg(address, value)
	default:
		// Default case, if needed
	}
}

func readBit(bit byte, value byte) int {
	// Read bit from value and return it
	return int((value >> bit) & 1)
}
func readStack() byte {
	return readMemory(SPBaseAddress + SP)
}
func updateStack(value byte) {
	writeMemory(SPBaseAddress+SP, value)
}
