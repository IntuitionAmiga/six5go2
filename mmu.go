package main

func readMemory(address uint16) byte {
	return memory[address]
}
func writeMemory(address uint16, value byte) {
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
