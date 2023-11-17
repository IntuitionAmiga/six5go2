package main

import (
	"fmt"
	"os"
)

var opcodeMnemonics = map[byte]string{
	0x69: "ADC", 0x65: "ADC", 0x75: "ADC", 0x6D: "ADC", 0x7D: "ADC", 0x79: "ADC", 0x61: "ADC", 0x71: "ADC",
	0x29: "AND", 0x25: "AND", 0x35: "AND", 0x2D: "AND", 0x3D: "AND", 0x39: "AND", 0x21: "AND", 0x31: "AND",
	0x0A: "ASL", 0x06: "ASL", 0x16: "ASL", 0x0E: "ASL", 0x1E: "ASL",
	0x90: "BCC",
	0xB0: "BCS",
	0xF0: "BEQ",
	0x24: "BIT", 0x2C: "BIT",
	0x30: "BMI",
	0xD0: "BNE",
	0x10: "BPL",
	0x00: "BRK",
	0x50: "BVC",
	0x70: "BVS",
	0x18: "CLC",
	0xD8: "CLD",
	0x58: "CLI",
	0xB8: "CLV",
	0xC9: "CMP", 0xC5: "CMP", 0xD5: "CMP", 0xCD: "CMP", 0xDD: "CMP", 0xD9: "CMP", 0xC1: "CMP", 0xD1: "CMP",
	0xE0: "CPX", 0xE4: "CPX", 0xEC: "CPX",
	0xC0: "CPY", 0xC4: "CPY", 0xCC: "CPY",
	0xC6: "DEC", 0xD6: "DEC", 0xCE: "DEC", 0xDE: "DEC",
	0xCA: "DEX",
	0x88: "DEY",
	0x49: "EOR", 0x45: "EOR", 0x55: "EOR", 0x4D: "EOR", 0x5D: "EOR", 0x59: "EOR", 0x41: "EOR", 0x51: "EOR",
	0xE6: "INC", 0xF6: "INC", 0xEE: "INC", 0xFE: "INC",
	0xE8: "INX",
	0xC8: "INY",
	0x4C: "JMP", 0x6C: "JMP",
	0x20: "JSR",
	0xA9: "LDA", 0xA5: "LDA", 0xB5: "LDA", 0xAD: "LDA", 0xBD: "LDA", 0xB9: "LDA", 0xA1: "LDA", 0xB1: "LDA",
	0xA2: "LDX", 0xA6: "LDX", 0xB6: "LDX", 0xAE: "LDX", 0xBE: "LDX",
	0xA0: "LDY", 0xA4: "LDY", 0xB4: "LDY", 0xAC: "LDY", 0xBC: "LDY",
	0x4A: "LSR", 0x46: "LSR", 0x56: "LSR", 0x4E: "LSR", 0x5E: "LSR",
	0xEA: "NOP",
	0x09: "ORA", 0x05: "ORA", 0x15: "ORA", 0x0D: "ORA", 0x1D: "ORA", 0x19: "ORA", 0x01: "ORA", 0x11: "ORA",
	0x48: "PHA",
	0x08: "PHP",
	0x68: "PLA",
	0x28: "PLP",
	0x2A: "ROL", 0x26: "ROL", 0x36: "ROL", 0x2E: "ROL", 0x3E: "ROL",
	0x6A: "ROR", 0x66: "ROR", 0x76: "ROR", 0x6E: "ROR", 0x7E: "ROR",
	0x40: "RTI",
	0x60: "RTS",
	0xE9: "SBC", 0xE5: "SBC", 0xF5: "SBC", 0xED: "SBC", 0xFD: "SBC", 0xF9: "SBC", 0xE1: "SBC", 0xF1: "SBC",
	0x38: "SEC",
	0xF8: "SED",
	0x78: "SEI",
	0x85: "STA", 0x95: "STA", 0x8D: "STA", 0x9D: "STA", 0x99: "STA", 0x81: "STA", 0x91: "STA",
	0x86: "STX", 0x96: "STX", 0x8E: "STX",
	0x84: "STY", 0x94: "STY", 0x8C: "STY",
	0xAA: "TAX",
	0xA8: "TAY",
	0xBA: "TSX",
	0x8A: "TXA",
	0x9A: "TXS",
	0x98: "TYA",
}

// Constant names for opcodes
const (
	LDA_IMMEDIATE_OPCODE   = 0xA9
	LDA_ZERO_PAGE_OPCODE   = 0xA5
	LDA_ZERO_PAGE_X_OPCODE = 0xB5
	LDA_ABSOLUTE_OPCODE    = 0xAD
	LDA_ABSOLUTE_X_OPCODE  = 0xBD
	LDA_ABSOLUTE_Y_OPCODE  = 0xB9
	LDA_INDIRECT_X_OPCODE  = 0xA1
	LDA_INDIRECT_Y_OPCODE  = 0xB1

	LDX_IMMEDIATE_OPCODE   = 0xA2
	LDX_ZERO_PAGE_OPCODE   = 0xA6
	LDX_ZERO_PAGE_Y_OPCODE = 0xB6
	LDX_ABSOLUTE_OPCODE    = 0xAE
	LDX_ABSOLUTE_Y_OPCODE  = 0xBE

	LDY_IMMEDIATE_OPCODE   = 0xA0
	LDY_ZERO_PAGE_OPCODE   = 0xA4
	LDY_ZERO_PAGE_X_OPCODE = 0xB4
	LDY_ABSOLUTE_OPCODE    = 0xAC
	LDY_ABSOLUTE_X_OPCODE  = 0xBC

	STA_ZERO_PAGE_OPCODE   = 0x85
	STA_ZERO_PAGE_X_OPCODE = 0x95
	STA_ABSOLUTE_OPCODE    = 0x8D
	STA_ABSOLUTE_X_OPCODE  = 0x9D
	STA_ABSOLUTE_Y_OPCODE  = 0x99
	STA_INDIRECT_X_OPCODE  = 0x81
	STA_INDIRECT_Y_OPCODE  = 0x91

	STX_ZERO_PAGE_OPCODE   = 0x86
	STX_ZERO_PAGE_Y_OPCODE = 0x96
	STX_ABSOLUTE_OPCODE    = 0x8E

	STY_ZERO_PAGE_OPCODE   = 0x84
	STY_ZERO_PAGE_X_OPCODE = 0x94
	STY_ABSOLUTE_OPCODE    = 0x8C

	CMP_IMMEDIATE_OPCODE   = 0xC9
	CMP_ZERO_PAGE_OPCODE   = 0xC5
	CMP_ZERO_PAGE_X_OPCODE = 0xD5
	CMP_ABSOLUTE_OPCODE    = 0xCD
	CMP_ABSOLUTE_X_OPCODE  = 0xDD
	CMP_ABSOLUTE_Y_OPCODE  = 0xD9
	CMP_INDIRECT_X_OPCODE  = 0xC1
	CMP_INDIRECT_Y_OPCODE  = 0xD1

	CPX_IMMEDIATE_OPCODE = 0xE0
	CPX_ZERO_PAGE_OPCODE = 0xE4
	CPX_ABSOLUTE_OPCODE  = 0xEC

	CPY_IMMEDIATE_OPCODE = 0xC0
	CPY_ZERO_PAGE_OPCODE = 0xC4
	CPY_ABSOLUTE_OPCODE  = 0xCC

	DEC_ZERO_PAGE_OPCODE   = 0xC6
	DEC_ZERO_PAGE_X_OPCODE = 0xD6
	DEC_ABSOLUTE_OPCODE    = 0xCE
	DEC_ABSOLUTE_X_OPCODE  = 0xDE

	INC_ZERO_PAGE_OPCODE   = 0xE6
	INC_ZERO_PAGE_X_OPCODE = 0xF6
	INC_ABSOLUTE_OPCODE    = 0xEE
	INC_ABSOLUTE_X_OPCODE  = 0xFE

	JMP_ABSOLUTE_OPCODE = 0x4C
	JMP_INDIRECT_OPCODE = 0x6C

	JSR_ABSOLUTE_OPCODE = 0x20

	AND_IMMEDIATE_OPCODE   = 0x29
	AND_ZERO_PAGE_OPCODE   = 0x25
	AND_ZERO_PAGE_X_OPCODE = 0x35
	AND_ABSOLUTE_OPCODE    = 0x2D
	AND_ABSOLUTE_X_OPCODE  = 0x3D
	AND_ABSOLUTE_Y_OPCODE  = 0x39
	AND_INDIRECT_X_OPCODE  = 0x21
	AND_INDIRECT_Y_OPCODE  = 0x31

	EOR_IMMEDIATE_OPCODE   = 0x49
	EOR_ZERO_PAGE_OPCODE   = 0x45
	EOR_ZERO_PAGE_X_OPCODE = 0x55
	EOR_ABSOLUTE_OPCODE    = 0x4D
	EOR_ABSOLUTE_X_OPCODE  = 0x5D
	EOR_ABSOLUTE_Y_OPCODE  = 0x59
	EOR_INDIRECT_X_OPCODE  = 0x41
	EOR_INDIRECT_Y_OPCODE  = 0x51

	ORA_IMMEDIATE_OPCODE   = 0x09
	ORA_ZERO_PAGE_OPCODE   = 0x05
	ORA_ZERO_PAGE_X_OPCODE = 0x15
	ORA_ABSOLUTE_OPCODE    = 0x0D
	ORA_ABSOLUTE_X_OPCODE  = 0x1D
	ORA_ABSOLUTE_Y_OPCODE  = 0x19
	ORA_INDIRECT_X_OPCODE  = 0x01
	ORA_INDIRECT_Y_OPCODE  = 0x11

	ADC_IMMEDIATE_OPCODE   = 0x69
	ADC_ZERO_PAGE_OPCODE   = 0x65
	ADC_ZERO_PAGE_X_OPCODE = 0x75
	ADC_ABSOLUTE_OPCODE    = 0x6D
	ADC_ABSOLUTE_X_OPCODE  = 0x7D
	ADC_ABSOLUTE_Y_OPCODE  = 0x79
	ADC_INDIRECT_X_OPCODE  = 0x61
	ADC_INDIRECT_Y_OPCODE  = 0x71

	SBC_IMMEDIATE_OPCODE   = 0xE9
	SBC_ZERO_PAGE_OPCODE   = 0xE5
	SBC_ZERO_PAGE_X_OPCODE = 0xF5
	SBC_ABSOLUTE_OPCODE    = 0xED
	SBC_ABSOLUTE_X_OPCODE  = 0xFD
	SBC_ABSOLUTE_Y_OPCODE  = 0xF9
	SBC_INDIRECT_X_OPCODE  = 0xE1
	SBC_INDIRECT_Y_OPCODE  = 0xF1

	BIT_ZERO_PAGE_OPCODE = 0x24
	BIT_ABSOLUTE_OPCODE  = 0x2C

	ROL_ACCUMULATOR_OPCODE = 0x2A
	ROL_ZERO_PAGE_OPCODE   = 0x26
	ROL_ZERO_PAGE_X_OPCODE = 0x36
	ROL_ABSOLUTE_OPCODE    = 0x2E
	ROL_ABSOLUTE_X_OPCODE  = 0x3E

	ROR_ACCUMULATOR_OPCODE = 0x6A
	ROR_ZERO_PAGE_OPCODE   = 0x66
	ROR_ZERO_PAGE_X_OPCODE = 0x76
	ROR_ABSOLUTE_OPCODE    = 0x6E
	ROR_ABSOLUTE_X_OPCODE  = 0x7E

	LSR_ACCUMULATOR_OPCODE = 0x4A
	LSR_ZERO_PAGE_OPCODE   = 0x46
	LSR_ZERO_PAGE_X_OPCODE = 0x56
	LSR_ABSOLUTE_OPCODE    = 0x4E
	LSR_ABSOLUTE_X_OPCODE  = 0x5E

	ASL_ACCUMULATOR_OPCODE = 0x0A
	ASL_ZERO_PAGE_OPCODE   = 0x06
	ASL_ZERO_PAGE_X_OPCODE = 0x16
	ASL_ABSOLUTE_OPCODE    = 0x0E
	ASL_ABSOLUTE_X_OPCODE  = 0x1E

	BCC_RELATIVE_OPCODE = 0x90
	BCS_RELATIVE_OPCODE = 0xB0
	BEQ_RELATIVE_OPCODE = 0xF0
	BNE_RELATIVE_OPCODE = 0xD0
	BMI_RELATIVE_OPCODE = 0x30
	BPL_RELATIVE_OPCODE = 0x10
	BVC_RELATIVE_OPCODE = 0x50
	BVS_RELATIVE_OPCODE = 0x70

	CLC_OPCODE = 0x18
	CLD_OPCODE = 0xD8
	CLI_OPCODE = 0x58
	CLV_OPCODE = 0xB8
	DEX_OPCODE = 0xCA
	DEY_OPCODE = 0x88
	INX_OPCODE = 0xE8
	INY_OPCODE = 0xC8
	NOP_OPCODE = 0xEA
	PHA_OPCODE = 0x48
	PHP_OPCODE = 0x08
	PLA_OPCODE = 0x68
	PLP_OPCODE = 0x28
	RTI_OPCODE = 0x40
	RTS_OPCODE = 0x60
	SEC_OPCODE = 0x38
	SED_OPCODE = 0xF8
	SEI_OPCODE = 0x78
	TAX_OPCODE = 0xAA
	TAY_OPCODE = 0xA8
	TSX_OPCODE = 0xBA
	TXA_OPCODE = 0x8A
	TXS_OPCODE = 0x9A
	TYA_OPCODE = 0x98

	BRK_OPCODE = 0x00
)

var (
	oneByteInstructions = map[byte]bool{
		CLC_OPCODE: true, CLD_OPCODE: true, CLI_OPCODE: true, CLV_OPCODE: true,
		DEX_OPCODE: true, DEY_OPCODE: true, INX_OPCODE: true, INY_OPCODE: true,
		NOP_OPCODE: true,
		PHA_OPCODE: true, PHP_OPCODE: true,
		PLA_OPCODE: true, PLP_OPCODE: true,
		SEC_OPCODE: true, SED_OPCODE: true, SEI_OPCODE: true,
		TAX_OPCODE: true, TAY_OPCODE: true,
		TSX_OPCODE: true,
		TXA_OPCODE: true, TXS_OPCODE: true, TYA_OPCODE: true, BRK_OPCODE: true, LDX_ZERO_PAGE_OPCODE: true,
		STA_ZERO_PAGE_OPCODE: true, STX_ZERO_PAGE_OPCODE: true, STY_ZERO_PAGE_OPCODE: true, RTS_OPCODE: true, RTI_OPCODE: true,
		ROL_ACCUMULATOR_OPCODE: true, ROR_ACCUMULATOR_OPCODE: true, LSR_ACCUMULATOR_OPCODE: true, ASL_ACCUMULATOR_OPCODE: true,
	}
	twoByteInstructions = map[byte]bool{
		LDA_IMMEDIATE_OPCODE: true, LDX_IMMEDIATE_OPCODE: true, LDY_IMMEDIATE_OPCODE: true, LDA_ZERO_PAGE_OPCODE: true,
		LDX_ZERO_PAGE_OPCODE: true, LDY_ZERO_PAGE_OPCODE: true, LDA_ZERO_PAGE_X_OPCODE: true, LDX_ZERO_PAGE_Y_OPCODE: true,
		LDY_ZERO_PAGE_X_OPCODE: true,
		CMP_IMMEDIATE_OPCODE:   true,
		CPX_IMMEDIATE_OPCODE:   true, CPY_IMMEDIATE_OPCODE: true,
		AND_IMMEDIATE_OPCODE: true,
		EOR_IMMEDIATE_OPCODE: true,
		ORA_IMMEDIATE_OPCODE: true,
		ADC_IMMEDIATE_OPCODE: true, SBC_IMMEDIATE_OPCODE: true,
		BIT_ZERO_PAGE_OPCODE: true,
		STX_ZERO_PAGE_OPCODE: true, STY_ZERO_PAGE_OPCODE: true, STA_ZERO_PAGE_OPCODE: true, STA_ZERO_PAGE_X_OPCODE: true,
		DEC_ZERO_PAGE_OPCODE: true, INC_ZERO_PAGE_OPCODE: true,
		ROL_ZERO_PAGE_OPCODE: true, ROR_ZERO_PAGE_OPCODE: true,
		LSR_ZERO_PAGE_OPCODE: true, ASL_ZERO_PAGE_OPCODE: true,
		BCC_RELATIVE_OPCODE: true, BCS_RELATIVE_OPCODE: true,
		BEQ_RELATIVE_OPCODE: true, BNE_RELATIVE_OPCODE: true,
		BMI_RELATIVE_OPCODE: true, BPL_RELATIVE_OPCODE: true,
		BVC_RELATIVE_OPCODE: true, BVS_RELATIVE_OPCODE: true,

		LDA_INDIRECT_X_OPCODE: true, LDA_INDIRECT_Y_OPCODE: true,
		CMP_INDIRECT_X_OPCODE: true, CMP_INDIRECT_Y_OPCODE: true,
		AND_INDIRECT_X_OPCODE: true, AND_INDIRECT_Y_OPCODE: true,
		EOR_INDIRECT_X_OPCODE: true, EOR_INDIRECT_Y_OPCODE: true,
		ORA_INDIRECT_X_OPCODE: true, ORA_INDIRECT_Y_OPCODE: true,
		ADC_INDIRECT_X_OPCODE: true, ADC_INDIRECT_Y_OPCODE: true,
		SBC_INDIRECT_X_OPCODE: true, SBC_INDIRECT_Y_OPCODE: true,
		STY_ZERO_PAGE_X_OPCODE: true, STX_ZERO_PAGE_Y_OPCODE: true,
		ORA_ZERO_PAGE_OPCODE: true, ORA_ZERO_PAGE_X_OPCODE: true,
		ASL_ZERO_PAGE_X_OPCODE: true,
		AND_ZERO_PAGE_OPCODE:   true, AND_ZERO_PAGE_X_OPCODE: true,
		ROL_ZERO_PAGE_X_OPCODE: true,
		EOR_ZERO_PAGE_OPCODE:   true, EOR_ZERO_PAGE_X_OPCODE: true,
		LSR_ZERO_PAGE_X_OPCODE: true,
		ADC_ZERO_PAGE_OPCODE:   true, ADC_ZERO_PAGE_X_OPCODE: true,
		ROR_ZERO_PAGE_X_OPCODE: true,
		CPY_ZERO_PAGE_OPCODE:   true,
		CMP_ZERO_PAGE_OPCODE:   true, CMP_ZERO_PAGE_X_OPCODE: true,
		DEC_ZERO_PAGE_X_OPCODE: true,
		CPX_ZERO_PAGE_OPCODE:   true,
		SBC_ZERO_PAGE_OPCODE:   true, SBC_ZERO_PAGE_X_OPCODE: true,
		INC_ZERO_PAGE_X_OPCODE: true,
	}
	threeByteInstructions = map[byte]bool{
		ADC_ABSOLUTE_OPCODE: true, ADC_ABSOLUTE_X_OPCODE: true, ADC_ABSOLUTE_Y_OPCODE: true,
		AND_ABSOLUTE_OPCODE: true, AND_ABSOLUTE_X_OPCODE: true, AND_ABSOLUTE_Y_OPCODE: true,
		ASL_ABSOLUTE_OPCODE: true, ASL_ABSOLUTE_X_OPCODE: true,
		BIT_ABSOLUTE_OPCODE: true,
		CMP_ABSOLUTE_OPCODE: true, CMP_ABSOLUTE_X_OPCODE: true, CMP_ABSOLUTE_Y_OPCODE: true,
		CPX_ABSOLUTE_OPCODE: true, CPY_ABSOLUTE_OPCODE: true,
		DEC_ABSOLUTE_OPCODE: true, DEC_ABSOLUTE_X_OPCODE: true,
		EOR_ABSOLUTE_OPCODE: true, EOR_ABSOLUTE_X_OPCODE: true, EOR_ABSOLUTE_Y_OPCODE: true,
		INC_ABSOLUTE_OPCODE: true, INC_ABSOLUTE_X_OPCODE: true,
		LDA_ABSOLUTE_OPCODE: true, LDA_ABSOLUTE_X_OPCODE: true, LDA_ABSOLUTE_Y_OPCODE: true,
		LDX_ABSOLUTE_OPCODE: true, LDX_ABSOLUTE_Y_OPCODE: true, LDY_ABSOLUTE_OPCODE: true, LDY_ABSOLUTE_X_OPCODE: true,
		LSR_ABSOLUTE_OPCODE: true, LSR_ABSOLUTE_X_OPCODE: true,
		ORA_ABSOLUTE_OPCODE: true, ORA_ABSOLUTE_X_OPCODE: true, ORA_ABSOLUTE_Y_OPCODE: true,
		ROL_ABSOLUTE_OPCODE: true, ROL_ABSOLUTE_X_OPCODE: true,
		ROR_ABSOLUTE_OPCODE: true, ROR_ABSOLUTE_X_OPCODE: true,
		SBC_ABSOLUTE_OPCODE: true, SBC_ABSOLUTE_X_OPCODE: true, SBC_ABSOLUTE_Y_OPCODE: true,
		STA_ABSOLUTE_OPCODE: true, STA_ABSOLUTE_X_OPCODE: true, STA_ABSOLUTE_Y_OPCODE: true,
		STX_ABSOLUTE_OPCODE: true, STY_ABSOLUTE_OPCODE: true, JMP_ABSOLUTE_OPCODE: true, JSR_ABSOLUTE_OPCODE: true, JMP_INDIRECT_OPCODE: true,
		STA_INDIRECT_X_OPCODE: true, STA_INDIRECT_Y_OPCODE: true,
	}
)

// Instructions with multiple addressing modes
func (cpu *CPU) LDA(addressingMode string) {
	var address uint16
	var value byte
	setFlags := func() {
		// If A is zero, set the SR Zero flag to 1 else set SR Zero flag to 0
		if cpu.preOpA == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		if cpu.preOpA&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		value = cpu.ImmediateAddressing()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX:
		address = cpu.ZeroPageXAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE: // Absolute
		address = cpu.AbsoluteAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX:
		baseAddress := uint16(cpu.preOpOperand2)<<8 | uint16(cpu.preOpOperand1)
		address = cpu.AbsoluteXAddressing()
		// Dummy read: occurs if adding X to the low byte of the address crosses a page boundary
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			cpu.readMemory(baseAddress&0xFF00 | (address & 0x00FF))
		}
		value = cpu.readMemory(address)
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			cpu.updateCycleCounter(5)
		} else {
			cpu.updateCycleCounter(4)
		}
		cpu.handleState(3)
	case ABSOLUTEY:
		baseAddress := uint16(cpu.preOpOperand2)<<8 | uint16(cpu.preOpOperand1)
		address = cpu.AbsoluteYAddressing()
		setFlags()
		// Add an extra cycle if page boundary is crossed
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			cpu.updateCycleCounter(5)
		} else {
			cpu.updateCycleCounter(4)
		}
		cpu.handleState(3)
	case INDIRECTX:
		address = cpu.IndirectXAddressing()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case INDIRECTY:
		address = cpu.IndirectYAddressing()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	}
	if addressingMode != IMMEDIATE {
		value = cpu.readMemory(address)
	}
	cpu.A = value
	setFlags()
}
func (cpu *CPU) LDX(addressingMode string) {
	var address uint16
	var value byte
	setFlags := func() {
		// If bit 7 of X is set, set the SR negative flag else reset it to 0
		if cpu.getXBit(7) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		// If X is zero, set the SR zero flag else reset it to 0
		if cpu.X == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		value = cpu.ImmediateAddressing()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEY:
		address = cpu.ZeroPageYAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEY:
		baseAddress := uint16(cpu.preOpOperand2)<<8 | uint16(cpu.preOpOperand1)
		//address := baseAddress + uint16(cpu.Y)
		address = cpu.AbsoluteYAddressing()
		// Dummy read: occurs if adding Y to the low byte of the address crosses a page boundary
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			cpu.readMemory(baseAddress&0xFF00 | (address & 0x00FF))
		}
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			cpu.updateCycleCounter(5)
		} else {
			cpu.updateCycleCounter(4)
		}
		cpu.handleState(3)
	}
	if addressingMode != IMMEDIATE {
		value = cpu.readMemory(address)
	}
	cpu.X = value
	setFlags()
}
func (cpu *CPU) LDY(addressingMode string) {
	var address uint16
	var value byte
	setFlags := func() {
		// If bit 7 of Y is set, set the SR negative flag else reset it to 0
		if cpu.getYBit(7) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		// If Y is zero, set the SR zero flag else reset it to 0
		if cpu.Y == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		value = cpu.ImmediateAddressing()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE: // Zero Page
		address = cpu.ZeroPageAddressing()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX: // Zero Page, X
		address = cpu.ZeroPageXAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE: // Absolute
		address = cpu.AbsoluteAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX:
		baseAddress := uint16(cpu.preOpOperand2)<<8 | uint16(cpu.preOpOperand1)
		address = cpu.AbsoluteXAddressing()
		// Dummy read: occurs if adding X to the low byte of the address crosses a page boundary
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			cpu.readMemory(baseAddress&0xFF00 | (address & 0x00FF))
		}
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			cpu.updateCycleCounter(5)
		} else {
			cpu.updateCycleCounter(4)
		}
		cpu.handleState(3)
	}
	if addressingMode != IMMEDIATE {
		value = cpu.readMemory(address)
	}
	cpu.Y = value
	setFlags()
}
func (cpu *CPU) STA(addressingMode string) {
	var address uint16
	switch addressingMode {
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX:
		address = cpu.ZeroPageXAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX:
		baseAddress := uint16(cpu.preOpOperand2)<<8 | uint16(cpu.preOpOperand1)
		address = cpu.AbsoluteXAddressing()
		// Perform dummy read if page boundary is crossed
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			_ = cpu.readMemory(baseAddress&0xFF00 | (address & 0x00FF))
		}
		cpu.updateCycleCounter(5)
		cpu.handleState(3)
	case ABSOLUTEY:
		baseAddress := uint16(cpu.preOpOperand2)<<8 | uint16(cpu.preOpOperand1)
		address = cpu.AbsoluteYAddressing()
		// Perform dummy read if page boundary is crossed
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			_ = cpu.readMemory(baseAddress&0xFF00 | (address & 0x00FF))
		}
		cpu.updateCycleCounter(5)
		cpu.handleState(3)
	case INDIRECTX:
		address = cpu.IndirectXAddressing()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case INDIRECTY:
		address = cpu.IndirectYAddressing()
		//finalAddress := (address + uint16(cpu.preOpY)) & 0xFFFF
		//cpu.writeMemory(finalAddress, cpu.preOpA)
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	}
	cpu.writeMemory(address, cpu.preOpA)
}
func (cpu *CPU) STX(addressingMode string) {
	var address uint16
	switch addressingMode {
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEY:
		address = cpu.ZeroPageYAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	}
	cpu.writeMemory(address, cpu.preOpX)
}
func (cpu *CPU) STY(addressingMode string) {
	var address uint16
	switch addressingMode {
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX:
		address = cpu.ZeroPageXAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	}
	cpu.writeMemory(address, cpu.preOpY)
}
func (cpu *CPU) CMP(addressingMode string) {
	var value, result byte
	var address uint16
	setFlags := func() {
		// Subtract the value from the accumulator
		result = cpu.preOpA - value
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		if cpu.readBit(7, result) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		// If the value is less than or equal to the accumulator, set the carry flag, else reset it
		if value <= cpu.preOpA {
			cpu.setCarryFlag()
		} else {
			cpu.unsetCarryFlag()
		}
		if addressingMode == IMMEDIATE || addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX || addressingMode == INDIRECTX || addressingMode == INDIRECTY {
			cpu.updateCycleCounter(2)
			cpu.handleState(2)
		} else {
			cpu.updateCycleCounter(3)
			cpu.handleState(3)
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		value = cpu.ImmediateAddressing()
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
	case ZEROPAGEX:
		address = cpu.ZeroPageXAddressing()
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
	case ABSOLUTEX:
		address = cpu.AbsoluteXAddressing()
	case ABSOLUTEY:
		address = cpu.AbsoluteYAddressing()
	case INDIRECTX: // Indirect, X
		address = cpu.IndirectXAddressing()
	case INDIRECTY:
		address = cpu.IndirectYAddressing()
	}
	if addressingMode != IMMEDIATE {
		value = cpu.readMemory(address)
	}
	setFlags()
}
func (cpu *CPU) JMP(addressingMode string) {
	var address uint16
	switch addressingMode {
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
	case INDIRECT:
		address = cpu.IndirectAddressing()
	}
	if *klausd {
		//if cpu.readMemory(0x02) == 0xDE && cpu.readMemory(0x03) == 0xB0 {
		//	//fmt.println("All tests passed!")
		//	os.Exit(0)
		//}
		//fmt.Fprintf(os.Stderr, "Klaus D. loop detected at address 0x%04X\n", cpu.PC)
		if cpu.PC == KlausDInfiniteLoopAddress {
			fmt.Printf("Klaus D. loop detected at address $%04X\n", KlausDInfiniteLoopAddress)
			os.Exit(0)
		}
	}
	cpu.setPC(address)
	cpu.updateCycleCounter(3)
	cpu.handleState(0)
}
func (cpu *CPU) AND(addressingMode string) {
	var value, result byte
	var address uint16

	setFlags := func() {
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		if cpu.readBit(7, result) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		value = cpu.ImmediateAddressing()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX:
		address = cpu.ZeroPageXAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX:
		address = cpu.AbsoluteXAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEY:
		address = cpu.AbsoluteYAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case INDIRECTX:
		address = cpu.IndirectXAddressing()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case INDIRECTY:
		address = cpu.IndirectYAddressing()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	}
	if addressingMode != IMMEDIATE {
		value = cpu.readMemory(address)
	}
	// AND the value with the accumulator
	result = cpu.A & value
	cpu.A = result
	setFlags()
}
func (cpu *CPU) EOR(addressingMode string) {
	var value, result byte
	var address uint16
	setFlags := func() {
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		//if cpu.readBit(7, result) == 1 {
		if result&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		value = cpu.ImmediateAddressing()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX:
		address = cpu.ZeroPageXAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX:
		address = cpu.AbsoluteXAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEY:
		address = cpu.AbsoluteYAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case INDIRECTX:
		address = cpu.IndirectXAddressing()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case INDIRECTY:
		address = cpu.IndirectYAddressing()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	}
	if addressingMode != IMMEDIATE {
		value = cpu.readMemory(address)
	}
	// XOR the value with the accumulator
	result = cpu.A ^ value
	cpu.A = result
	setFlags()
}
func (cpu *CPU) ORA(addressingMode string) {
	var value, result byte
	var address uint16
	setFlags := func() {
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		if cpu.readBit(7, result) == 1 {
			cpu.setNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		value = cpu.ImmediateAddressing()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX:
		address = cpu.ZeroPageXAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX:
		address = cpu.AbsoluteXAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEY:
		address = cpu.AbsoluteYAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case INDIRECTX:
		address = cpu.IndirectXAddressing()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case INDIRECTY:
		address = cpu.IndirectYAddressing()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	}
	if addressingMode != IMMEDIATE {
		value = cpu.readMemory(address)
	}
	// OR the value with the accumulator
	result = cpu.preOpA | value
	cpu.A = result
	setFlags()
}
func (cpu *CPU) BIT(addressingMode string) {
	var value, result byte
	var address uint16
	setFlags := func() {
		// Set Negative flag to bit 7 of the value
		if cpu.readBit(7, value) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		// Set Overflow flag to bit 6 of the value
		if cpu.readBit(6, value) == 1 {
			cpu.setOverflowFlag()
		} else {
			cpu.unsetOverflowFlag()
		}
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
	}
	switch addressingMode {
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	}
	value = cpu.readMemory(address)
	// AND the value with the accumulator
	result = cpu.preOpA & value
	cpu.A = result
	setFlags()
}
func (cpu *CPU) INC(addressingMode string) {
	var address uint16
	var result, value byte

	setFlags := func() {
		// Fetch the value from the address
		value = cpu.readMemory(address)
		// Increment the value (wrapping around for 8-bit values)
		result = value + 1
		// Write the result back to memory
		cpu.writeMemory(address, result)

		// Update status flags
		// If bit 7 of the result is set, set the negative flag
		if result&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
	}
	switch addressingMode {
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	case ZEROPAGEX:
		address = cpu.ZeroPageXAddressing()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
		cpu.updateCycleCounter(6)
		cpu.handleState(3)
	case ABSOLUTEX:
		address = cpu.AbsoluteXAddressing()
		cpu.updateCycleCounter(7)
		cpu.handleState(3)
	}
	setFlags()
}
func (cpu *CPU) DEC(addressingMode string) {
	var address uint16
	var result, value byte

	setFlags := func() {
		// Fetch the value from the address
		value = cpu.readMemory(address)
		// Decrement the value (wrapping around for 8-bit values)
		result = value - 1
		// Write the result back to memory
		cpu.writeMemory(address, result)

		// Update status flags
		// If bit 7 of the result is set, set the negative flag
		//if cpu.readBit(7, result) == 1 {
		if result&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
	}
	switch addressingMode {
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	case ZEROPAGEX:
		address = cpu.ZeroPageXAddressing()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
		cpu.updateCycleCounter(6)
		cpu.handleState(3)
	case ABSOLUTEX:
		address = cpu.AbsoluteXAddressing()
		cpu.updateCycleCounter(7)
		cpu.handleState(3)
	}
	setFlags()
}
func (cpu *CPU) ADC(addressingMode string) {
	var value byte
	var result int
	var address uint16

	setFlags := func() {
		// Binary mode is the default
		tmpResult := int(cpu.preOpA) + int(value)
		if cpu.getSRBit(0) == 1 {
			tmpResult++
		}

		if cpu.getSRBit(3) == 1 { // BCD mode
			temp := (cpu.preOpA & 0x0F) + (value & 0x0F) + cpu.getSRBit(0)
			if temp > 9 {
				temp += 6
			}

			result = int((cpu.preOpA & 0xF0) + (value & 0xF0) + (temp & 0x0F))

			if result > 0x99 {
				result += 0x60
			}

		}

		// Set or unset the C flag
		if tmpResult > 0xFF {
			cpu.setCarryFlag()
		} else {
			cpu.unsetCarryFlag()
		}

		// Handle V (overflow) flag
		if (int(cpu.preOpA)^int(value))&0x80 == 0 && (int(cpu.preOpA)^tmpResult)&0x80 != 0 {
			cpu.setOverflowFlag()
		} else {
			cpu.unsetOverflowFlag()
		}

		result = tmpResult & 0xFF // Store the result in 8 bits

		// Handle N (negative) and Z (zero) flags
		if result&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}

		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}

		cpu.A = byte(result)

		if addressingMode == IMMEDIATE || addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX || addressingMode == INDIRECTX || addressingMode == INDIRECTY {
			cpu.updateCycleCounter(2)
			cpu.handleState(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX || addressingMode == ABSOLUTEY {
			cpu.updateCycleCounter(3)
			cpu.handleState(3)
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		value = cpu.ImmediateAddressing()
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
	case ZEROPAGEX:
		address = cpu.ZeroPageXAddressing()
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
	case ABSOLUTEX:
		address = cpu.AbsoluteXAddressing()
		//Perform a dummy read of the address then fetch value
		_ = cpu.readMemory(address)
	case ABSOLUTEY:
		address = cpu.AbsoluteYAddressing()
		// Perform a dummy read of the address then fetch value
		_ = cpu.readMemory(address)
	case INDIRECTX:
		address = cpu.IndirectXAddressing()
	case INDIRECTY:
		address = cpu.IndirectYAddressing()
	}
	if addressingMode != IMMEDIATE {
		value = cpu.readMemory(address)
	}
	setFlags()
}
func (cpu *CPU) SBC(addressingMode string) {
	var value byte
	var result int
	var address uint16

	setFlags := func() {

		// Check for BCD mode (Decimal flag set)
		if cpu.getSRBit(3) == 1 { // BCD mode
			lowNibble := (cpu.preOpA & 0x0F) - (value & 0x0F) - (cpu.getSRBit(0) ^ 1)
			if lowNibble < 0 {
				lowNibble = (lowNibble - 6) & 0x0F // Only keep the low nibble
			}

			highNibble := (cpu.preOpA & 0xF0) - (value & 0xF0) - (lowNibble & 0x10)
			if highNibble < 0 {
				highNibble -= 0x60
			}

			result = int(highNibble) | int(lowNibble) // Convert to int before combining

			if result >= 0 {
				cpu.setCarryFlag() // No borrow occurred, set carry
			} else {
				cpu.unsetCarryFlag() // Borrow occurred, clear carry
			}

			result &= 0xFF // Ensure result is only 8 bits
		} else {
			// Binary mode
			result = int(cpu.preOpA) - int(value)
			if cpu.getSRBit(0) == 0 {
				result--
			}

			// The carry flag in binary mode is the inverse of the borrow indicator
			if result >= 0 {
				cpu.setCarryFlag()
			} else {
				cpu.unsetCarryFlag()
				result = 0xFF + result + 1 // Adjust result for negative values
			}
		}

		cpu.A = byte(result)

		if cpu.readBit(7, cpu.A) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		if result == 0 {
			cpu.setZeroFlag()
		}

		if addressingMode == IMMEDIATE || addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX || addressingMode == INDIRECTX || addressingMode == INDIRECTY {
			cpu.updateCycleCounter(2)
			cpu.handleState(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX || addressingMode == ABSOLUTEY {
			cpu.updateCycleCounter(3)
			cpu.handleState(3)
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		value = cpu.ImmediateAddressing()
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
	case ZEROPAGEX:
		address = cpu.ZeroPageXAddressing()
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
	case ABSOLUTEX:
		address = cpu.AbsoluteXAddressing()
	case ABSOLUTEY:
		address = cpu.AbsoluteYAddressing()
	case INDIRECTX:
		address = cpu.IndirectXAddressing()
	case INDIRECTY:
		address = cpu.IndirectYAddressing()
	}
	if addressingMode != IMMEDIATE {
		value = cpu.readMemory(address)
	}
	setFlags()
}
func (cpu *CPU) ROR(addressingMode string) {
	var value, result byte
	var address uint16

	setFlags := func() {
		// Store the carry in a temp variable
		oldCarry := cpu.getSRBit(0)

		// Set carry flag to bit 0 of value
		if (value & 0x01) != 0 {
			cpu.setCarryFlag()
		} else {
			cpu.unsetCarryFlag()
		}

		// Rotate right one bit
		result = (value >> 1) | (oldCarry << 7)

		// Set zero flag based on the result
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}

		// Set negative flag based on bit 7 of the result
		if (result & 0x80) != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		if addressingMode == ACCUMULATOR {
			// Store the result in the accumulator
			cpu.A = result
			cpu.updateCycleCounter(2)
			cpu.handleState(1)
		}
		if addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX {
			// Store the value back into memory
			cpu.writeMemory(address, result)
			cpu.updateCycleCounter(5)
			cpu.handleState(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX {
			// Store the value back into memory
			cpu.writeMemory(address, result)
			cpu.updateCycleCounter(6)
			cpu.handleState(3)
		}
	}
	switch addressingMode {
	case ACCUMULATOR:
		value = cpu.preOpA
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
	case ZEROPAGEX:
		address = cpu.ZeroPageXAddressing()
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
	case ABSOLUTEX:
		address = cpu.AbsoluteXAddressing()
	}
	if addressingMode != ACCUMULATOR {
		value = cpu.readMemory(address)
	}
	result = (value >> 1)
	setFlags()
}
func (cpu *CPU) ROL(addressingMode string) {
	var value, result byte
	var address uint16

	setFlags := func() {
		// Store the carry in a temp variable
		oldCarry := cpu.getSRBit(0)

		// Set SR carry flag to bit 7 of value
		if (value & 0x80) != 0 {
			cpu.setCarryFlag()
		} else {
			cpu.unsetCarryFlag()
		}

		// Rotate left one bit
		result = (value << 1) | oldCarry

		// Set zero flag based on the result
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}

		// Set negative flag based on bit 7 of the result
		if (result & 0x80) != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		if addressingMode == ACCUMULATOR {
			// Store the result in the accumulator
			cpu.A = result
			cpu.updateCycleCounter(2)
			cpu.handleState(1)
		}
		if addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX {
			// Store the value back into memory
			cpu.writeMemory(address, result)
			cpu.updateCycleCounter(5)
			cpu.handleState(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX {
			// Store the value back into memory
			cpu.writeMemory(address, result)
			cpu.updateCycleCounter(6)
			cpu.handleState(3)
		}
	}
	switch addressingMode {
	case ACCUMULATOR:
		// Get the value of the accumulator
		value = cpu.preOpA
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
	case ZEROPAGEX:
		address = cpu.ZeroPageXAddressing()
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
	case ABSOLUTEX:
		address = cpu.AbsoluteXAddressing()
	}
	if addressingMode != ACCUMULATOR {
		value = cpu.readMemory(address)
	}
	result = value << 1
	result = (result & 0xFE) | cpu.getSRBit(0)
	setFlags()
}
func (cpu *CPU) LSR(addressingMode string) {
	var value, result byte
	var address uint16

	setFlags := func() {
		// Reset the SR negative flag
		cpu.unsetNegativeFlag()

		// Update the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}

		// If bit 0 of value is 1 then set SR carry flag else reset it
		if (value & 0x01) != 0 {
			cpu.setCarryFlag()
		} else {
			cpu.unsetCarryFlag()
		}
	}

	switch addressingMode {
	case ACCUMULATOR:
		// Get the value of the accumulator
		value = cpu.preOpA
		// Shift the value right 1 bit
		result = value >> 1
		// Store the result back into the accumulator
		cpu.A = result
		cpu.updateCycleCounter(2)
		cpu.handleState(1)
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	case ZEROPAGEX:
		address = cpu.ZeroPageXAddressing()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
		cpu.updateCycleCounter(6)
		cpu.handleState(3)
	case ABSOLUTEX:
		// Dummy read for Absolute,X addressing mode if page boundary is crossed
		baseAddress := uint16(cpu.preOpOperand2)<<8 | uint16(cpu.preOpOperand1)
		address = cpu.AbsoluteXAddressing()
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			// Perform a dummy read if a page boundary is crossed
			cpu.readMemory(baseAddress&0xFF00 | (address & 0x00FF))
		}
		// Update cycle counter depending on page boundary crossing
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			cpu.updateCycleCounter(7) // Add 1 if page boundary is crossed
		} else {
			cpu.updateCycleCounter(6)
		}
		cpu.handleState(3)
	}
	if addressingMode != ACCUMULATOR {
		value = cpu.readMemory(address)
		result = value >> 1
		cpu.writeMemory(address, result)
	}
	setFlags()
}
func (cpu *CPU) ASL(addressingMode string) {
	var value, result byte
	var address uint16

	setFlags := func() {
		// Update the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}

		// Update the negative flag
		if (result & 0x80) != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}

		// Set the Carry flag based on the original value's bit 7 before the shift operation
		if (value & 0x80) != 0 {
			cpu.setCarryFlag()
		} else {
			cpu.unsetCarryFlag()
		}
	}
	switch addressingMode {
	case ACCUMULATOR:
		// Set value to accumulator
		value = cpu.preOpA
		// Shift the value left 1 bit
		result = value << 1
		// Update the accumulator with the result
		cpu.A = result
		cpu.updateCycleCounter(2)
		cpu.handleState(1)
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	case ZEROPAGEX:
		address = cpu.ZeroPageXAddressing()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
		cpu.updateCycleCounter(6)
		cpu.handleState(3)
	case ABSOLUTEX:
		address = cpu.AbsoluteXAddressing()
		cpu.updateCycleCounter(7)
		cpu.handleState(3)
	}
	if addressingMode != ACCUMULATOR {
		value = cpu.readMemory(address)
		result = value << 1
		cpu.writeMemory(address, result)
	}
	setFlags()
}
func (cpu *CPU) CPX(addressingMode string) {
	var value, result byte
	var address uint16

	setFlags := func() {
		// Set or clear the carry flag
		if cpu.preOpX >= value {
			cpu.setCarryFlag()
		} else {
			cpu.unsetCarryFlag()
		}

		// Set or clear the zero flag
		if cpu.preOpX == value {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}

		// Set or clear the negative flag based on the high bit of the subtraction result
		result = cpu.preOpX - value
		if result&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		value = cpu.ImmediateAddressing()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
		cpu.updateCycleCounter(3)
		cpu.handleState(3)
	}
	if addressingMode != IMMEDIATE {
		value = cpu.readMemory(address)
	}
	setFlags()
}
func (cpu *CPU) CPY(addressingMode string) {
	var value, result byte
	var address uint16

	setFlags := func() {
		// Set or clear the carry flag
		if cpu.preOpY >= value {
			cpu.setCarryFlag()
		} else {
			cpu.unsetCarryFlag()
		}

		// Set or clear the zero flag
		if cpu.preOpY == value {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}

		// Set or clear the negative flag based on the high bit of the subtraction result
		result = cpu.preOpY - value
		if result&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		value = cpu.ImmediateAddressing()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		address = cpu.ZeroPageAddressing()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ABSOLUTE:
		address = cpu.AbsoluteAddressing()
		cpu.updateCycleCounter(3)
		cpu.handleState(3)
	}
	if addressingMode != IMMEDIATE {
		value = cpu.readMemory(address)
	}
	setFlags()
}

// 1 byte instructions with no operands

// Implied addressing mode instructions
/*
	In the implied addressing mode, the address containing the operand is implicitly stated in the operation code of the instruction.

	Bytes: 1
*/
func (cpu *CPU) BRK() {
	// BRK - Break Command

	disassembleOpcode()

	// Increment PC by 1 as BRK is a 2-byte instruction (opcode + padding byte)
	nextPC := cpu.preOpPC + 1

	// Push high byte of next PC onto stack
	cpu.decSP()
	cpu.updateStack(byte(nextPC >> 8))

	// Push low byte of next PC onto stack
	cpu.decSP()
	cpu.updateStack(byte(nextPC & 0xFF))

	// Modify status register: set break flag and push onto stack
	modifiedSR := cpu.preOpSR | 0x10 // Set break flag
	cpu.decSP()
	cpu.updateStack(modifiedSR)

	// Set interrupt disable flag to prevent further interrupts
	cpu.setInterruptFlag()

	// Load the IRQ/BRK vector into PC to jump to the interrupt service routine
	irqVectorLow := cpu.readMemory(IRQVectorAddressLow)
	irqVectorHigh := cpu.readMemory(IRQVectorAddressHigh)
	cpu.setPC((uint16(irqVectorHigh) << 8) | uint16(irqVectorLow))

	// Update cycle counter and state
	cpu.updateCycleCounter(7)
	cpu.handleState(0)
}
func (cpu *CPU) CLC() {
	/*
		CLC - Clear Carry Flag
	*/
	disassembleOpcode()
	cpu.unsetCarryFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) CLD() {
	/*
		CLD - Clear Decimal Mode
	*/
	disassembleOpcode()
	cpu.unsetDecimalFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) CLI() {
	/*
		CLI - Clear Interrupt Disable
	*/
	disassembleOpcode()
	cpu.unsetInterruptFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) CLV() {
	/*
		CLV - Clear Overflow Flag
	*/
	disassembleOpcode()
	cpu.unsetOverflowFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) DEX() {
	/*
		DEX - Decrement Index Register X By One
	*/
	disassembleOpcode()
	// Decrement the X register by 1, with underflow handling
	cpu.X = (cpu.preOpX - 1) & 0xFF
	// Update the Negative Flag based on the new value of X
	if cpu.X&0x80 != 0 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	// Update the Zero Flag based on the new value of X
	if cpu.X == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) DEY() {
	/*
	   DEY - Decrement Index Register Y By One
	*/
	disassembleOpcode()
	// Decrement the X register by 1, with underflow handling
	cpu.Y = (cpu.preOpY - 1) & 0xFF
	// Update the Negative Flag based on the new value of Y
	if cpu.Y&0x80 != 0 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	// Update the Zero Flag based on the new value of Y
	if cpu.Y == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) INX() {
	/*
		INX - Increment Index Register X By One
	*/
	disassembleOpcode()
	cpu.X++
	if cpu.X == 0 { // Check for overflow
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	if cpu.X&0x80 != 0 { // Check the negative flag
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) INY() {
	/*
		INY - Increment Index Register Y By One
	*/
	disassembleOpcode()
	cpu.Y++
	if cpu.Y == 0 { // Check for overflow
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	if cpu.Y&0x80 != 0 { // Check the negative flag
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) NOP() {
	/*
		NOP - No Operation
	*/
	disassembleOpcode()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) PHA() {
	/*
		PHA - Push Accumulator On Stack
	*/
	disassembleOpcode()
	cpu.updateStack(cpu.A)
	cpu.decSP()
	cpu.updateCycleCounter(3)
	cpu.handleState(1)
}
func (cpu *CPU) PHP() {
	/*
	   PHP - Push Processor Status On Stack
	*/
	disassembleOpcode()
	// Set the break flag and the unused bit only for the push operation
	status := cpu.SR | (1 << 4) // Set break flag
	status |= (1 << 5)          // Set unused bit
	// Push the status onto the stack
	cpu.updateStack(status)
	// Decrement the stack pointer
	cpu.decSP()
	cpu.updateCycleCounter(3)
	cpu.handleState(1)
}
func (cpu *CPU) PLA() {
	/*
	   PLA - Pull Accumulator From Stack
	*/
	disassembleOpcode()
	// Increment the stack pointer first
	cpu.incSP()
	// Read the value from the stack into the accumulator
	cpu.A = cpu.readStack()
	// No flag updates should occur here
	cpu.updateCycleCounter(4)
	cpu.handleState(1)
}
func (cpu *CPU) PLP() {
	/*
		PLP - Pull Processor Status From Stack
	*/
	disassembleOpcode()
	// Read the status from the stack
	newStatus := cpu.readStack()
	// Preserve break flag and unused bit from current status
	//newStatus = (newStatus & 0xCF) | (cpu.SR & 0x30)
	newStatus = (newStatus & 0x3F) | (cpu.SR & 0xC0)

	// Update SR with the new status
	cpu.SR = newStatus
	// Increment the stack pointer after the operation
	cpu.incSP()
	cpu.updateCycleCounter(4)
	cpu.handleState(1)
}
func (cpu *CPU) RTI() {
	disassembleOpcode()
	// Increment the stack pointer and read processor status
	cpu.incSP()
	cpu.SR = cpu.readStack()
	// Increment the stack pointer and read low byte of PC
	cpu.incSP()
	low := uint16(cpu.readStack())
	// Increment the stack pointer and read high byte of PC
	cpu.incSP()
	high := uint16(cpu.readStack())
	// Combine high and low bytes and set the program counter
	cpu.setPC((high << 8) | low)
	cpu.updateCycleCounter(6)
	cpu.handleState(0)
}
func (cpu *CPU) RTS() {
	// Get low byte of new PC
	low := uint16(cpu.readStack())
	cpu.incSP() // Increment the stack pointer
	// Get high byte of new PC
	high := uint16(cpu.readStack())
	cpu.incSP() // Increment the stack pointer again
	// Update PC with the value stored in memory at the address pointed to by SP
	cpu.setPC(((high << 8) | low) + 1)
	cpu.updateCycleCounter(6)
	cpu.handleState(0)
}
func (cpu *CPU) SEC() {
	/*
		SEC - Set Carry Flag
	*/
	disassembleOpcode()
	cpu.setCarryFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) SED() {
	/*
		SED - Set Decimal Mode
	*/
	disassembleOpcode()
	cpu.setDecimalFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) SEI() {
	/*
		SEI - Set Interrupt Disable
	*/
	disassembleOpcode()
	cpu.setInterruptFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) TAX() {
	/*
		TAX - Transfer Accumulator To Index X
	*/
	disassembleOpcode()
	// Update X with the value of A
	cpu.X = cpu.preOpA
	if cpu.getXBit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	if cpu.X == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) TAY() {
	/*
		TAY - Transfer Accumulator To Index Y
	*/
	disassembleOpcode()
	// Set Y register to the value of the accumulator
	cpu.Y = cpu.preOpA
	if cpu.getYBit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	if cpu.Y == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) TSX() {
	/*
		TSX - Transfer Stack Pointer To Index X
	*/
	disassembleOpcode()
	// Update X with the SP
	cpu.X = byte(cpu.preOpSP)
	// If X register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
	if cpu.getXBit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	// If X register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
	if cpu.X == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) TXA() {
	/*
		TXA - Transfer Index X To Accumulator
	*/
	disassembleOpcode()
	// Set accumulator to value of X register
	cpu.A = cpu.preOpX
	// If bit 7 of accumulator is set, set negative SR flag else set negative SR flag to 0
	if cpu.getABit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	// If accumulator is 0, set zero SR flag else set zero SR flag to 0
	if cpu.A == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) TXS() {
	/*
		TXS - Transfer Index X To Stack Pointer
	*/
	disassembleOpcode()
	// Set stack pointer to value of X register
	cpu.SP = uint16(cpu.X)
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func (cpu *CPU) TYA() {
	/*
		TYA - Transfer Index Y To Accumulator
	*/
	disassembleOpcode()
	// Set accumulator to value of Y register
	cpu.A = cpu.Y
	// If bit 7 of accumulator is set, set negative SR flag else set negative SR flag to 0
	if cpu.getABit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	// If accumulator is 0, set zero SR flag else set zero SR flag to 0
	if cpu.A == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}

// Accumulator instructions
/*
	A

	This form of addressing is represented with a one byte instruction, implying an operation on the accumulator.

	Bytes: 1
*/
func (cpu *CPU) ASL_A() {
	/*
		ASL - Arithmetic Shift Left
	*/
	disassembleOpcode()
	cpu.ASL("accumulator")
}
func (cpu *CPU) LSR_A() {
	/*
		LSR - Logical Shift Right
	*/
	disassembleOpcode()
	cpu.LSR("accumulator")
}
func (cpu *CPU) ROL_A() {
	/*
		ROL - Rotate Left
	*/
	disassembleOpcode()
	cpu.ROL("accumulator")
}
func (cpu *CPU) ROR_A() {
	/*
		ROR - Rotate Right
	*/
	disassembleOpcode()
	cpu.ROR("accumulator")
}

// 2 byte instructions with 1 operand
// Immediate addressing mode instructions
/*
	#$nn

	In immediate addressing, the operand is contained in the second byte of the instruction, with no further memory addressing required.

	Bytes: 2
*/
func (cpu *CPU) ADC_I() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	disassembleOpcode()

	cpu.ADC("immediate")
}
func (cpu *CPU) AND_I() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.AND("immediate")
}
func (cpu *CPU) CMP_I() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembleOpcode()
	cpu.CMP("immediate")
}
func (cpu *CPU) CPX_I() {
	/*
		CPX - Compare Index Register X To Memory
	*/
	disassembleOpcode()

	cpu.CPX("immediate")
}
func (cpu *CPU) CPY_I() {
	/*
		CPY - Compare Index Register Y To Memory
	*/
	disassembleOpcode()
	cpu.CPY("immediate")
}
func (cpu *CPU) EOR_I() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.EOR("immediate")
}
func (cpu *CPU) LDA_I() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembleOpcode()
	cpu.LDA("immediate")
}
func (cpu *CPU) LDX_I() {
	/*
		LDX - Load Index Register X From Memory
	*/
	disassembleOpcode()
	cpu.LDX("immediate")
}
func (cpu *CPU) LDY_I() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	disassembleOpcode()
	cpu.LDY("immediate")
}
func (cpu *CPU) ORA_I() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.ORA("immediate")
}
func (cpu *CPU) SBC_I() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembleOpcode()
	cpu.SBC("immediate")
}

// Zero Page addressing mode instructions
/*
	$nn

	The zero page instructions allow for shorter code and execution times by only fetching the second byte of the instruction and assuming a zero low address byte. Careful use of the zero page can result in significant increase in code efficiency.

	Bytes: 2
*/
func (cpu *CPU) ADC_Z() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	disassembleOpcode()
	cpu.ADC("zeropage")
}
func (cpu *CPU) AND_Z() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.AND("zeropage")
}
func (cpu *CPU) ASL_Z() {
	/*
		ASL - Arithmetic Shift Left
	*/
	disassembleOpcode()

	cpu.ASL("zeropage")
}
func (cpu *CPU) BIT_Z() {
	/*
		BIT - Test Bits in Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.BIT("zeropage")
}
func (cpu *CPU) CMP_Z() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembleOpcode()
	cpu.CMP("zeropage")
}
func (cpu *CPU) CPX_Z() {
	/*
		CPX - Compare Index Register X To Memory
	*/
	disassembleOpcode()
	cpu.CPX("zeropage")
}
func (cpu *CPU) CPY_Z() {
	/*
		CPY - Compare Index Register Y To Memory
	*/
	disassembleOpcode()
	cpu.CPY("zeropage")
}
func (cpu *CPU) DEC_Z() {
	/*
		DEC - Decrement Memory By One
	*/
	disassembleOpcode()
	cpu.DEC("zeropage")
}
func (cpu *CPU) EOR_Z() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.EOR("zeropage")
}
func (cpu *CPU) INC_Z() {
	/*
		INC - Increment Memory By One
	*/
	disassembleOpcode()
	cpu.INC("zeropage")
}
func (cpu *CPU) LDA_Z() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembleOpcode()
	cpu.LDA("zeropage")
}
func (cpu *CPU) LDX_Z() {
	/*
		LDX - Load Index Register X From Memory
	*/
	disassembleOpcode()
	cpu.LDX("zeropage")
}
func (cpu *CPU) LDY_Z() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	disassembleOpcode()
	cpu.LDY("zeropage")
}
func (cpu *CPU) LSR_Z() {
	/*
		LSR - Logical Shift Right
	*/
	disassembleOpcode()
	cpu.LSR("zeropage")
}
func (cpu *CPU) ORA_Z() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.ORA("zeropage")
}
func (cpu *CPU) ROL_Z() {
	/*
		ROL - Rotate Left
	*/
	disassembleOpcode()
	cpu.ROL("zeropage")
}
func (cpu *CPU) ROR_Z() {
	/*
		ROR - Rotate Right
	*/
	disassembleOpcode()
	cpu.ROR("zeropage")
}
func (cpu *CPU) SBC_Z() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembleOpcode()
	cpu.SBC("zeropage")
}
func (cpu *CPU) STA_Z() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembleOpcode()
	cpu.STA("zeropage")
}
func (cpu *CPU) STX_Z() {
	/*
		STX - Store Index Register X In Memory
	*/
	disassembleOpcode()
	cpu.STX("zeropage")
}
func (cpu *CPU) STY_Z() {
	/*
		STY - Store Index Register Y In Memory
	*/
	disassembleOpcode()
	cpu.STY("zeropage")
}

// X Indexed Zero Page addressing mode instructions
/*
	$nn,X

	This form of addressing is used in conjunction with the X index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the low order 8 bits of memory and crossing of page boundaries does not occur.

	Bytes: 2
*/
func (cpu *CPU) ADC_ZX() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	disassembleOpcode()
	cpu.ADC("zeropagex")
}
func (cpu *CPU) AND_ZX() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.AND("zeropagex")
}
func (cpu *CPU) ASL_ZX() {
	/*
		ASL - Arithmetic Shift Left
	*/
	disassembleOpcode()
	cpu.ASL("zeropagex")
}
func (cpu *CPU) CMP_ZX() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembleOpcode()
	cpu.CMP("zeropagex")
}
func (cpu *CPU) DEC_ZX() {
	/*
		DEC - Decrement Memory By One
	*/
	disassembleOpcode()
	cpu.DEC("zeropagex")
}
func (cpu *CPU) LDA_ZX() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembleOpcode()
	cpu.LDA("zeropagex")
}
func (cpu *CPU) LDY_ZX() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	disassembleOpcode()
	cpu.LDY("zeropagex")
}
func (cpu *CPU) LSR_ZX() {
	/*
		LSR - Logical Shift Right
	*/
	disassembleOpcode()
	cpu.LSR("zeropagex")
}
func (cpu *CPU) ORA_ZX() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.ORA("zeropagex")
}
func (cpu *CPU) ROL_ZX() {
	/*
		ROL - Rotate Left
	*/
	disassembleOpcode()
	cpu.ROL("zeropagex")
}
func (cpu *CPU) ROR_ZX() {
	/*
		ROR - Rotate Right
	*/
	disassembleOpcode()
	cpu.ROR("zeropagex")
}
func (cpu *CPU) EOR_ZX() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.EOR("zeropagex")
}
func (cpu *CPU) INC_ZX() {
	/*
		INC - Increment Memory By One
	*/
	disassembleOpcode()
	cpu.INC("zeropagex")
}
func (cpu *CPU) SBC_ZX() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembleOpcode()
	cpu.SBC("zeropagex")
}
func (cpu *CPU) STA_ZX() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembleOpcode()
	cpu.STA("zeropagex")
}
func (cpu *CPU) STY_ZX() {
	/*
		STY - Store Index Register Y In Memory
	*/
	disassembleOpcode()
	cpu.STY("zeropagex")
}

// Y Indexed Zero Page addressing mode instructions
/*
	$nn,Y

	This form of addressing is used in conjunction with the Y index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the low order 8 bits of memory and crossing of page boundaries does not occur.

	Bytes: 2
*/
func (cpu *CPU) LDX_ZY() {
	/*
		LDX - Load Index Register X From Memory
	*/
	disassembleOpcode()
	cpu.LDX("zeropagey")
}
func (cpu *CPU) STX_ZY() {
	/*
		STX - Store Index Register X In Memory
	*/
	disassembleOpcode()
	cpu.STX("zeropagey")
}

// X Indexed Zero Page Indirect addressing mode instructions
/*
	($nn,X)

	In indexed indirect addressing, the second byte of the instruction is added to the contents of the X index register, discarding the carry. The result of this addition points to a memory location on page zero whose contents is the high order eight bits of the effective address. The next memory location in page zero contains the low order eight bits of the effective address. Both memory locations specifying the low and high order bytes of the effective address must be in page zero.

	Bytes: 2
*/
func (cpu *CPU) ADC_IX() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	disassembleOpcode()
	cpu.ADC("indirectx")
}
func (cpu *CPU) AND_IX() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.AND("indirectx")
}
func (cpu *CPU) CMP_IX() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembleOpcode()
	cpu.CMP("indirectx")
}
func (cpu *CPU) EOR_IX() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.EOR("indirectx")
}
func (cpu *CPU) LDA_IX() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembleOpcode()
	cpu.LDA("indirectx")
}
func (cpu *CPU) ORA_IX() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.ORA("indirectx")
}
func (cpu *CPU) SBC_IX() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembleOpcode()
	cpu.SBC("indirectx")
}
func (cpu *CPU) STA_IX() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembleOpcode()
	cpu.STA("indirectx")
}

// Zero Page Indirect Y Indexed addressing mode instructions
/*
	($nn),Y

	In indirect indexed addressing, the second byte of the instruction points to a memory location in page zero. The contents of this memory location is added to the contents of the Y index register, the result being the high order eight bits of the effective address. The carry from this addition is added to the contents of the next page zero memory location, the result being the low order eight bits of the effective address.

	Bytes: 2
*/
func (cpu *CPU) ADC_IY() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	disassembleOpcode()
	cpu.ADC("indirecty")
}
func (cpu *CPU) AND_IY() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.AND("indirecty")
}
func (cpu *CPU) CMP_IY() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembleOpcode()
	cpu.CMP("indirecty")
}
func (cpu *CPU) EOR_IY() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.EOR("indirecty")
}
func (cpu *CPU) LDA_IY() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembleOpcode()
	cpu.LDA("indirecty")
}
func (cpu *CPU) ORA_IY() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.ORA("indirecty")
}
func (cpu *CPU) SBC_IY() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembleOpcode()
	cpu.SBC("indirecty")
}
func (cpu *CPU) STA_IY() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembleOpcode()
	cpu.STA("indirecty")
}

// Relative addressing mode instructions
/*
	$nnnn

	Relative addressing is used only with branch instructions and establishes a destination for the conditional branch.

	The second byte of-the instruction becomes the operand which is an “Offset" added to the contents of the lower eight bits of the program counter when the counter is set at the next instruction. The range of the offset is —128 to +127 bytes from the next instruction.

	Bytes: 2
*/
func (cpu *CPU) BPL_R() {
	/*
	   BPL - Branch on Result Plus
	*/
	disassembleOpcode()
	offset := cpu.preOpOperand1
	signedOffset := int8(offset) // Cast to int8 to handle signed offset

	// The branch target address should be calculated from the address of the next instruction
	// The next instruction's address is current PC + size of this instruction (2 bytes: 1 for opcode, 1 for operand)
	nextInstructionAddress := cpu.preOpPC + 2
	// If the Negative flag is clear, then branch
	if cpu.getSRBit(7) == 0 {
		targetAddress := nextInstructionAddress + uint16(signedOffset) // Add signed offset to the next instruction's address
		cpu.setPC(targetAddress)
		// Increment the cycle counter based on branching
		cpu.updateCycleCounter(1) // Minimum cycle count for branching
		if (nextInstructionAddress & 0xFF00) != (targetAddress & 0xFF00) {
			cpu.updateCycleCounter(1) // Add an extra cycle for crossing a page boundary
		}
		cpu.handleState(0)
	} else {
		// Don't branch
		cpu.updateCycleCounter(2) // Two cycles when branch is not taken
		cpu.handleState(2)
	}
}
func (cpu *CPU) BMI_R() {
	/*
	   BMI - Branch on Result Minus
	*/
	disassembleOpcode()
	offset := int8(cpu.preOpOperand1) // Get offset from operand

	// If N flag is set, branch to address
	if cpu.getSRBit(7) == 1 {
		targetAddress := uint16(int16(cpu.preOpPC) + 2 + int16(offset)) // Calculate the target address

		cpu.setPC(targetAddress)
		cpu.updateCycleCounter(1) // Every branch takes at least one cycle

		// Increment the cycle count for branch and page crossing
		cpu.incrementCycleCountForBranch(cpu.preOpPC)
	} else {
		// Don't branch, skip to the next instruction
		cpu.incPC(2)
		cpu.updateCycleCounter(2) // Two cycles when branch is not taken
	}
}
func (cpu *CPU) BVC_R() {
	/*
	   BVC - Branch on Overflow Clear
	*/
	disassembleOpcode()
	offset := int8(cpu.preOpOperand1) // Get offset from operand

	// If overflow flag is not set, branch to address
	if cpu.getSRBit(6) == 0 {
		targetAddress := uint16(int16(cpu.preOpPC) + 2 + int16(offset)) // Calculate the target address

		cpu.setPC(targetAddress)
		cpu.updateCycleCounter(1) // Every branch takes at least one cycle

		// Increment the cycle count for branch and page crossing
		cpu.incrementCycleCountForBranch(cpu.preOpPC)
	} else {
		// Don't branch, skip to the next instruction
		cpu.incPC(2)
		cpu.updateCycleCounter(2) // Two cycles when branch is not taken
	}
}
func (cpu *CPU) BVS_R() {
	/*
	   BVS - Branch on Overflow Set
	*/
	disassembleOpcode()
	offset := int8(cpu.preOpOperand1) // Get offset from operand

	// If overflow flag is set, branch to address
	if cpu.getSRBit(6) == 1 {
		targetAddress := uint16(int16(cpu.preOpPC) + 2 + int16(offset)) // Calculate the target address

		cpu.setPC(targetAddress)
		cpu.updateCycleCounter(1) // Every branch takes at least one cycle

		// Increment the cycle count for branch and page crossing
		cpu.incrementCycleCountForBranch(cpu.preOpPC)
	} else {
		// Don't branch, skip to the next instruction
		cpu.incPC(2)
		cpu.updateCycleCounter(2) // Two cycles when branch is not taken
	}
}
func (cpu *CPU) BCC_R() {
	/*
	   BCC - Branch on Carry Clear
	*/
	disassembleOpcode()

	offset := int8(cpu.preOpOperand1)                               // Get offset from operand
	targetAddress := uint16(int16(cpu.preOpPC) + 2 + int16(offset)) // Calculate the target address

	// If carry flag is clear, branch to address
	if cpu.getSRBit(0) == 0 {
		cpu.setPC(targetAddress)
		cpu.updateCycleCounter(1) // Every branch takes at least one cycle
		// Increment the cycle count for branch and page crossing
		cpu.incrementCycleCountForBranch(cpu.preOpPC)
	} else {
		// Don't branch, skip to the next instruction
		cpu.incPC(2)
		cpu.updateCycleCounter(2) // Two cycles when branch is not taken
	}
}
func (cpu *CPU) BCS_R() {
	/*
	   BCS - Branch on Carry Set
	*/
	disassembleOpcode()

	offset := int8(cpu.preOpOperand1)                               // Get offset as signed 8-bit integer
	targetAddress := uint16(int16(cpu.preOpPC) + 2 + int16(offset)) // Calculate the target address

	// If carry flag is set, branch to address
	if cpu.getSRBit(0) == 1 {
		cpu.setPC(targetAddress)
		cpu.updateCycleCounter(1) // Every branch takes at least one cycle
		// Increment the cycle count for branch and page crossing
		cpu.incrementCycleCountForBranch(cpu.preOpPC)
	} else {
		// Don't branch, skip to the next instruction
		cpu.incPC(2)
		cpu.updateCycleCounter(2) // Two cycles when branch is not taken
	}
}
func (cpu *CPU) BNE_R() {
	/*
	   BNE - Branch on Result Not Zero
	*/
	disassembleOpcode()

	offset := int8(cpu.preOpOperand1)                               // Cast to signed 8-bit to handle negative offsets
	targetAddress := uint16(int16(cpu.preOpPC) + 2 + int16(offset)) // Calculate the target address

	// Check Z flag to determine if branching is needed
	if cpu.getSRBit(1) == 0 {
		cpu.setPC(targetAddress)
		cpu.updateCycleCounter(1) // Every branch takes at least one cycle
		// Increment the cycle count for branch and page crossing
		cpu.incrementCycleCountForBranch(cpu.preOpPC)
	} else {
		// If Z flag is set, don't branch, skip to the next instruction
		cpu.incPC(2)
		cpu.updateCycleCounter(2) // Two cycles when branch is not taken
	}
}
func (cpu *CPU) BEQ_R() {
	/*
	   BEQ - Branch on Result Zero
	*/
	disassembleOpcode()

	offset := int8(cpu.preOpOperand1)                               // Cast to signed 8-bit to handle negative offsets
	targetAddress := uint16(int16(cpu.preOpPC) + 2 + int16(offset)) // Calculate the target address

	// If Z flag is set, branch to address
	if cpu.getSRBit(1) == 1 {
		cpu.setPC(targetAddress)
		cpu.updateCycleCounter(1) // Every branch takes at least one cycle
		// Increment the cycle count for branch and page crossing
		cpu.incrementCycleCountForBranch(cpu.preOpPC)
	} else {
		// If Z flag is not set, don't branch, skip to the next instruction
		cpu.incPC(2)
		cpu.updateCycleCounter(2) // Two cycles when branch is not taken
	}
}

// 3 byte instructions with 2 operands
// Absolute addressing mode instructions
/*
	$nnnn

	In absolute addressing, the second byte of the instruction specifies the eight high order bits of the effective address while the third byte specifies the eight low order bits. Thus, the absolute addressing mode allows access to the entire 65 K bytes of addressable memory.

	Bytes: 3
*/
func (cpu *CPU) ADC_ABS() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	disassembleOpcode()
	cpu.ADC("absolute")
}
func (cpu *CPU) AND_ABS() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.AND("absolute")
}
func (cpu *CPU) ASL_ABS() {
	/*
		ASL - Arithmetic Shift Left
	*/
	disassembleOpcode()
	cpu.ASL("absolute")
}
func (cpu *CPU) BIT_ABS() {
	/*
		BIT - Test Bits in Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.BIT("absolute")
}
func (cpu *CPU) CMP_ABS() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembleOpcode()
	cpu.CMP("absolute")
}
func (cpu *CPU) CPX_ABS() {
	/*
		CPX - Compare Index Register X To Memory
	*/
	disassembleOpcode()
	cpu.CPX("absolute")
}
func (cpu *CPU) CPY_ABS() {
	/*
		CPY - Compare Index Register Y To Memory
	*/
	disassembleOpcode()
	cpu.CPY("absolute")
}
func (cpu *CPU) DEC_ABS() {
	/*
		DEC - Decrement Memory By One
	*/
	disassembleOpcode()
	cpu.DEC("absolute")
}
func (cpu *CPU) EOR_ABS() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.EOR("absolute")
}
func (cpu *CPU) INC_ABS() {
	/*
		INC - Increment Memory By One
	*/
	disassembleOpcode()
	cpu.INC("absolute")
}
func (cpu *CPU) JMP_ABS() {
	/*
		JMP - JMP Absolute
	*/
	disassembleOpcode()
	// For AllSuiteA.bin 6502 opcode test suite
	if *allsuitea && cpu.readMemory(0x210) == 0xFF {
		fmt.Printf("\n\u001B[32;5mMemory address $210 == $%02X. All opcodes succesfully tested and passed!\u001B[0m\n", cpu.readMemory(0x210))
		os.Exit(0)
	}
	cpu.JMP("absolute")
}
func (cpu *CPU) JSR_ABS() {
	// JSR - Jump To Subroutine
	disassembleOpcode()

	// Calculate the return address (the address of the next instruction minus one)
	returnAddr := cpu.preOpPC + 2 // JSR is 3 bytes, so next instruction is at PC + 3, return address is PC + 2

	// Push return address onto the stack
	cpu.decSP()
	cpu.updateStack(byte(returnAddr >> 8)) // High byte
	cpu.decSP()
	cpu.updateStack(byte(returnAddr & 0xFF)) // Low byte

	// Jump to the subroutine address specified by the operands
	cpu.setPC(uint16(cpu.preOpOperand2)<<8 | uint16(cpu.preOpOperand1))

	cpu.updateCycleCounter(6) // JSR takes 6 cycles
	cpu.handleState(0)
}
func (cpu *CPU) LDA_ABS() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembleOpcode()
	cpu.LDA("absolute")
}
func (cpu *CPU) LDX_ABS() {
	/*
		LDX - Load Index Register X From Memory
	*/
	disassembleOpcode()
	cpu.LDX("absolute")
}
func (cpu *CPU) LDY_ABS() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	disassembleOpcode()
	cpu.LDY("absolute")
}
func (cpu *CPU) LSR_ABS() {
	/*
		LSR - Logical Shift Right
	*/
	disassembleOpcode()
	cpu.LSR("absolute")
}
func (cpu *CPU) ORA_ABS() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.ORA("absolute")
}
func (cpu *CPU) ROL_ABS() {
	/*
		ROL - Rotate Left
	*/
	disassembleOpcode()
	cpu.ROL("absolute")
}
func (cpu *CPU) ROR_ABS() {
	/*
		ROR - Rotate Right
	*/
	disassembleOpcode()
	cpu.ROR("absolute")
}
func (cpu *CPU) SBC_ABS() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembleOpcode()
	cpu.SBC("absolute")
}
func (cpu *CPU) STA_ABS() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembleOpcode()
	cpu.STA("absolute")
}
func (cpu *CPU) STX_ABS() {
	/*
		STX - Store Index Register X In Memory
	*/
	disassembleOpcode()
	cpu.STX("absolute")
}
func (cpu *CPU) STY_ABS() {
	/*
		STY - Store Index Register Y In Memory
	*/
	disassembleOpcode()
	cpu.STY("absolute")
}

// X Indexed Absolute addressing mode instructions
/*
	$nnnn,X

	This form of addressing is used in conjunction with the X index register. The effective address is formed by adding the contents of X to the address contained in the second and third bytes of the instruction. This mode allows the index register to contain the index or count value and the instruction to contain the base address. This type of indexing allows any location referencing and the index to modify multiple fields resulting in reduced coding and execution time.

	Note on the MOS 6502:

	The value at the specified address, ignoring the the addressing mode's X offset, is read (and discarded) before the final address is read. This may cause side effects in I/O registers.


	Bytes: 3
*/
func (cpu *CPU) ADC_ABX() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	disassembleOpcode()
	cpu.ADC("absolutex")
}
func (cpu *CPU) AND_ABX() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.AND("absolutex")
}
func (cpu *CPU) ASL_ABX() {
	/*
		ASL - Arithmetic Shift Left
	*/
	disassembleOpcode()
	cpu.ASL("absolutex")
}
func (cpu *CPU) CMP_ABX() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembleOpcode()
	cpu.CMP("absolutex")
}
func (cpu *CPU) DEC_ABX() {
	/*
		DEC - Decrement Memory By One
	*/
	disassembleOpcode()
	cpu.DEC("absolutex")
}
func (cpu *CPU) EOR_ABX() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.EOR("absolutex")
}
func (cpu *CPU) INC_ABX() {
	/*
		INC - Increment Memory By One
	*/
	disassembleOpcode()
	cpu.INC("absolutex")
}
func (cpu *CPU) LDA_ABX() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembleOpcode()
	cpu.LDA("absolutex")
}
func (cpu *CPU) LDY_ABX() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	disassembleOpcode()
	cpu.LDY("absolutex")
}
func (cpu *CPU) LSR_ABX() {
	/*
		LSR - Logical Shift Right
	*/
	disassembleOpcode()
	cpu.LSR("absolutex")
}
func (cpu *CPU) ORA_ABX() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.ORA("absolutex")
}
func (cpu *CPU) ROL_ABX() {
	/*
	 */
	disassembleOpcode()
	cpu.ROL("absolutex")
}
func (cpu *CPU) ROR_ABX() {
	/*
		ROR - Rotate Right
	*/
	disassembleOpcode()
	cpu.ROR("absolutex")
}
func (cpu *CPU) SBC_ABX() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembleOpcode()
	cpu.SBC("absolutex")
}
func (cpu *CPU) STA_ABX() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembleOpcode()
	cpu.STA("absolutex")
}

// Y Indexed Absolute addressing mode instructions
/*
	$nnnn,Y

	This form of addressing is used in conjunction with the Y index register. The effective address is formed by adding the contents of Y to the address contained in the second and third bytes of the instruction. This mode allows the index register to contain the index or count value and the instruction to contain the base address. This type of indexing allows any location referencing and the index to modify multiple fields resulting in reduced coding and execution time.

	Note on the MOS 6502:

	The value at the specified address, ignoring the the addressing mode's Y offset, is read (and discarded) before the final address is read. This may cause side effects in I/O registers.

	Bytes: 3
*/
func (cpu *CPU) ADC_ABY() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	disassembleOpcode()
	cpu.ADC("absolutey")
}
func (cpu *CPU) AND_ABY() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.AND("absolutey")
}
func (cpu *CPU) CMP_ABY() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembleOpcode()
	cpu.CMP("absolutey")
}
func (cpu *CPU) EOR_ABY() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.EOR("absolutey")
}
func (cpu *CPU) LDA_ABY() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembleOpcode()
	cpu.LDA("absolutey")
}
func (cpu *CPU) LDX_ABY() {
	/*
		LDX - Load Index Register X From Memory
	*/
	disassembleOpcode()
	cpu.LDX("absolutey")
}
func (cpu *CPU) ORA_ABY() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembleOpcode()
	cpu.ORA("absolutey")
}
func (cpu *CPU) SBC_ABY() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembleOpcode()
	cpu.SBC("absolutey")
}
func (cpu *CPU) STA_ABY() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembleOpcode()
	cpu.STA("absolutey")
}

// Absolute Indirect addressing mode instructions
func (cpu *CPU) JMP_IND() {
	/*
		JMP - JMP Indirect
	*/
	disassembleOpcode()
	cpu.JMP("indirect")
}

// ZeroPageAddressing computes the effective address for the Zero Page addressing mode.
func (cpu *CPU) ZeroPageAddressing() uint16 {
	address := cpu.preOpOperand1
	return uint16(address)
}

// ZeroPageXAddressing computes the effective address for the Zero Page, X addressing mode.
func (cpu *CPU) ZeroPageXAddressing() uint16 {
	address := (cpu.preOpOperand1 + cpu.X) & 0xFF
	return uint16(address)
}

// AbsoluteAddressing computes the effective address for the Absolute addressing mode.
func (cpu *CPU) AbsoluteAddressing() uint16 {
	address := uint16(cpu.preOpOperand2)<<8 | uint16(cpu.preOpOperand1)
	return address
}

// AbsoluteXAddressing computes the effective address for the Absolute, X addressing mode.
func (cpu *CPU) AbsoluteXAddressing() uint16 {
	baseAddress := uint16(cpu.preOpOperand2)<<8 | uint16(cpu.preOpOperand1)
	address := baseAddress + uint16(cpu.X)
	return address
}

// AbsoluteYAddressing computes the effective address for the Absolute, Y addressing mode.
func (cpu *CPU) AbsoluteYAddressing() uint16 {
	baseAddress := uint16(cpu.preOpOperand2)<<8 | uint16(cpu.preOpOperand1)
	address := baseAddress + uint16(cpu.Y)
	return address
}

func (cpu *CPU) IndirectXAddressing() uint16 {
	// Calculate the zero page address including the X register offset
	zeroPageAddress := uint16(cpu.preOpOperand1+cpu.X) & 0xFF

	// Read the low byte of the address
	lo := cpu.readMemory(zeroPageAddress)

	// Read the high byte of the address. Note that the high byte is read from the next zero page location.
	// Wraparound needs to be handled correctly.
	hi := cpu.readMemory((zeroPageAddress + 1) & 0xFF)

	// Combine the high and low bytes to get the final address
	address := (uint16(hi) << 8) | uint16(lo)

	//print out the values of zeroPageAddress, lo, hi, and the final address
	return address
}

// IndirectYAddressing computes the effective address for the Post-indexed Indirect addressing mode.
func (cpu *CPU) IndirectYAddressing() uint16 {
	zeroPageAddress := cpu.preOpOperand1
	lo := cpu.readMemory(uint16(zeroPageAddress))
	hi := cpu.readMemory((uint16(zeroPageAddress) + 1) & 0xFF) // Zero page wraparound for high byte
	baseAddress := uint16(hi)<<8 | uint16(lo)
	address := baseAddress + uint16(cpu.Y)
	return address
}

// ImmediateAddressing simply returns the operand as the immediate value.
func (cpu *CPU) ImmediateAddressing() byte {
	return cpu.preOpOperand1 // Immediate mode uses the operand itself as the value
}
func (cpu *CPU) ZeroPageYAddressing() uint16 {
	// Add the Y register to the operand, but keep the address in the zero page (0x00FF)
	address := (cpu.preOpOperand1 + cpu.Y) & 0xFF
	return uint16(address)
}

func (cpu *CPU) IndirectAddressing() uint16 {
	operandAddress := uint16(cpu.preOpOperand2)<<8 | uint16(cpu.preOpOperand1)

	var effectiveAddress uint16
	if cpu.preOpOperand1 == 0xFF {
		// Handle page boundary bug: high byte comes from the start of the current page
		wrapAddress := operandAddress & 0xFF00
		effectiveAddress = uint16(cpu.readMemory(wrapAddress)) | uint16(cpu.readMemory(operandAddress))<<8
	} else {
		// Normal behavior: increment the address to get the high byte
		effectiveAddress = uint16(cpu.readMemory(operandAddress+1))<<8 | uint16(cpu.readMemory(operandAddress))
	}

	return effectiveAddress
}
