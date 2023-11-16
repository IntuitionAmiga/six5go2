package main

// Constants for fixed addresses and TED specific addresses
const (
	SPBaseAddress          uint16 = 0xFF
	NMIVectorAddressLow    uint16 = 0xFFFA
	NMIVectorAddressHigh   uint16 = 0xFFFB
	RESETVectorAddressLow  uint16 = 0xFFFC
	RESETVectorAddressHigh uint16 = 0xFFFD
	IRQVectorAddressLow    uint16 = 0xFFFE
	IRQVectorAddressHigh   uint16 = 0xFFFF
	TED_REG_START                 = 0xFD00
	TED_REG_END                   = 0xFD19
)

var memory [65536]byte // Memory

func (cpu *CPU) readMemory(address uint16) byte {
	var value byte
	if address >= TED_REG_START && address <= TED_REG_END {
		value = ted.readTEDReg(address)
	} else {
		value = memory[address]
	}
	return value
}

func (cpu *CPU) writeMemory(address uint16, value byte) {
	if address >= TED_REG_START && address <= TED_REG_END {
		ted.writeTEDReg(address, value)
	} else {
		memory[address] = value
	}
	// Existing special address checks
	if address == IRQVectorAddressLow || address == IRQVectorAddressHigh {
		cpu.irq = true // Signal an IRQ
	}
	if address == NMIVectorAddressLow || address == NMIVectorAddressHigh {
		cpu.nmi = true // Signal an NMI
	}
	if address == RESETVectorAddressLow || address == RESETVectorAddressHigh {
		cpu.reset = true // Signal a RESET
	}
}

func (cpu *CPU) readBit(bit byte, value byte) int {
	// Read bit from value and return it
	return int((value >> bit) & 1)
}
func (cpu *CPU) readStack() byte {
	return cpu.readMemory(SPBaseAddress + cpu.SP)
}
func (cpu *CPU) updateStack(value byte) {
	cpu.writeMemory(SPBaseAddress+cpu.SP, value)
}
