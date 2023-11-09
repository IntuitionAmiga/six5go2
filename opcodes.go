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
func LDA(addressingMode string) {
	setFlags := func() {
		// If A is zero, set the SR Zero flag to 1 else set SR Zero flag to 0
		if cpu.A == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		if cpu.A&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
	}

	switch addressingMode {
	case IMMEDIATE: // Immediate
		cpu.A = cpu.operand1()
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE: // Zero Page
		// Get address
		address := cpu.operand1()
		// Get value from memory at address
		value := readMemory(uint16(address))
		// Set accumulator to value
		cpu.A = value
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX: // Zero Page, X
		// Get address
		address := cpu.operand1() + cpu.X
		value := readMemory(uint16(address))
		// Set accumulator to value
		cpu.A = value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE: // Absolute
		// Get 16 bit address from operand 1 and operand 2
		address := int(cpu.operand2())<<8 | int(cpu.operand1())
		value := readMemory(uint16(address))
		// Set accumulator to value
		cpu.A = value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX:
		baseAddress := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		address := baseAddress + uint16(cpu.X)

		// Dummy read: occurs if adding X to the low byte of the address crosses a page boundary
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			readMemory(baseAddress&0xFF00 | (address & 0x00FF))
		}
		value := readMemory(address)
		cpu.A = value
		setFlags()
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			cpu.updateCycleCounter(5)
		} else {
			cpu.updateCycleCounter(4)
		}
		cpu.handleState(3)
	case ABSOLUTEY:
		baseAddress := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		address := baseAddress + uint16(cpu.Y)

		// Dummy read: occurs if adding Y to the low byte of the address crosses a page boundary
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			readMemory(baseAddress&0xFF00 | (address & 0x00FF))
		}
		value := readMemory(address)
		cpu.A = value
		setFlags()
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			cpu.updateCycleCounter(5)
		} else {
			cpu.updateCycleCounter(4)
		}
		cpu.handleState(3)
	case INDIRECTX:
		zeroPageAddress := uint16(cpu.operand1()+cpu.X) & 0xFF // zero page wraparound
		lo := readMemory(zeroPageAddress)
		hi := readMemory((zeroPageAddress + 1) & 0xFF) // zero page wraparound
		address := uint16(hi)<<8 | uint16(lo)
		value := readMemory(address)
		cpu.A = value
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case INDIRECTY:
		zeroPageAddress := cpu.operand1()
		address := uint16(readMemory(uint16(zeroPageAddress))) | uint16(readMemory((uint16(zeroPageAddress)+1)&0xFF))<<8
		finalAddress := (address + uint16(cpu.Y)) & 0xFFFF
		value := readMemory(finalAddress)
		cpu.A = value
		setFlags()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	}
}
func LDX(addressingMode string) {
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
	case IMMEDIATE: // Immediate
		// Load the value of the cpu.operand1() into the X register.
		cpu.X = cpu.operand1()
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE: // Zero Page
		// Get address
		address := cpu.operand1()
		value := readMemory(uint16(address))
		// Load the value at the address into X
		cpu.X = value
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEY: // Zero Page, Y
		// Get Y indexed Zero Page address
		address := cpu.operand1() + cpu.Y
		value := readMemory(uint16(address))
		// Load the X register with the Y indexed value in the operand
		cpu.X = value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE: // Absolute
		// Get 16 bit address from operands
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		value := readMemory(address)
		// Update X with the value stored at the address in the operands
		cpu.X = value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEY:
		baseAddress := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		address := baseAddress + uint16(cpu.Y)

		// Dummy read: occurs if adding Y to the low byte of the address crosses a page boundary
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			readMemory(baseAddress&0xFF00 | (address & 0x00FF))
		}
		value := readMemory(address)
		cpu.X = value
		setFlags()
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			cpu.updateCycleCounter(5)
		} else {
			cpu.updateCycleCounter(4)
		}
		cpu.handleState(3)
	}
}
func LDY(addressingMode string) {
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
	case IMMEDIATE: // Immediate
		// Load the value of the cpu.operand1() into the Y register.
		cpu.Y = cpu.operand1()
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE: // Zero Page
		// Get address
		address := cpu.operand1()
		value := readMemory(uint16(address))
		// Load the value at the address into Y
		cpu.Y = value
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX: // Zero Page, X
		// Get the X indexed address
		address := cpu.operand1() + cpu.X
		value := readMemory(uint16(address))
		// Load the Y register with the X indexed value in the operand
		cpu.Y = value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE: // Absolute
		// Get 16 bit address from operands
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		value := readMemory(address)
		// Update Y with the value stored at the address in the operands
		cpu.Y = value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX:
		baseAddress := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		address := baseAddress + uint16(cpu.X)

		// Dummy read: occurs if adding X to the low byte of the address crosses a page boundary
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			readMemory(baseAddress&0xFF00 | (address & 0x00FF))
		}
		value := readMemory(address)
		cpu.Y = value
		setFlags()
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			cpu.updateCycleCounter(5)
		} else {
			cpu.updateCycleCounter(4)
		}
		cpu.handleState(3)
	}
}
func STA(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE:
		address := cpu.operand1()
		writeMemory(uint16(address), cpu.A)
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX:
		address := (cpu.operand1() + cpu.X) & 0xFF // Ensure wraparound in Zero Page
		writeMemory(uint16(address), cpu.A)
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		writeMemory(address, cpu.A)
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX:
		baseAddress := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		address := baseAddress + uint16(cpu.X)
		// Perform dummy read if page boundary is crossed
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			_ = readMemory(baseAddress&0xFF00 | (address & 0x00FF))
		}
		writeMemory(address, cpu.A)
		cpu.updateCycleCounter(5)
		cpu.handleState(3)
	case ABSOLUTEY:
		baseAddress := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		address := baseAddress + uint16(cpu.Y)
		// Perform dummy read if page boundary is crossed
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			_ = readMemory(baseAddress&0xFF00 | (address & 0x00FF))
		}
		writeMemory(address, cpu.A)
		cpu.updateCycleCounter(5)
		cpu.handleState(3)
	case INDIRECTX:
		zeroPageAddress := (cpu.operand1() + cpu.X) & 0xFF
		address := uint16(readMemory(uint16(zeroPageAddress))) | uint16(readMemory((uint16(zeroPageAddress)+1)&0xFF))<<8
		writeMemory(address, cpu.A)
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case INDIRECTY:
		zeroPageAddress := cpu.operand1()
		address := uint16(readMemory(uint16(zeroPageAddress))) | uint16(readMemory((uint16(zeroPageAddress)+1)&0xFF))<<8
		finalAddress := (address + uint16(cpu.Y)) & 0xFFFF
		writeMemory(finalAddress, cpu.A)
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	}
}
func STX(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE:
		address := cpu.operand1()
		writeMemory(uint16(address), cpu.X)
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEY:
		address := (cpu.operand1() + cpu.Y) & 0xFF
		writeMemory(uint16(address), cpu.X)
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		writeMemory(address, cpu.X)
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	}
}
func STY(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE:
		address := cpu.operand1()
		writeMemory(uint16(address), cpu.Y)
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX:
		address := (cpu.operand1() + cpu.X) & 0xFF
		writeMemory(uint16(address), cpu.Y)
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		writeMemory(address, cpu.Y)
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	}
}
func CMP(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// Subtract the value from the accumulator
		result = cpu.A - value
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		if readBit(7, result) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		// If the value is less than or equal to the accumulator, set the carry flag, else reset it
		if value <= cpu.A {
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
	case IMMEDIATE: // Immediate
		// Get value from cpu.operand1()
		value = cpu.operand1()
		setFlags()
	case ZEROPAGE: // Zero Page
		// Get address
		address := cpu.operand1()
		// Subtract the operand from the accumulator
		value = readMemory(uint16(address))
		setFlags()
	case ZEROPAGEX: // Zero Page, X
		// Get address
		address := cpu.operand1() + cpu.X
		// Get value at address
		value = readMemory(uint16(address))
		setFlags()
	case ABSOLUTE: // Absolute
		// Get 16bit absolute address
		address := int(cpu.operand2())<<8 | int(cpu.operand1())
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case ABSOLUTEX: // Absolute, X
		// Get address
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1()) + uint16(cpu.X)
		// Get value at address
		value = readMemory(address)
		setFlags()
	case ABSOLUTEY: // Absolute, Y
		// Get address
		address := int(cpu.operand2())<<8 | int(cpu.operand1()) + int(cpu.Y)
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case INDIRECTX: // Indirect, X
		// Get the address of the operand
		address := int(cpu.operand1()) + int(cpu.X)
		// Get the value of the operand
		value = readMemory(uint16(address))
		setFlags()
	case INDIRECTY: // Indirect, Y
		// Get the zero page address from the operand
		zeroPageAddress := cpu.operand1()
		// Fetch the low byte of the address
		lowByte := readMemory(uint16(zeroPageAddress))
		// Fetch the high byte of the address from the next page address, wrapping within the zero page
		highByte := readMemory(uint16(zeroPageAddress+1) & 0x00FF)
		// Combine the high and low bytes to form the 16-bit address
		address := uint16(highByte)<<8 | uint16(lowByte)
		// Add the Y register to the address with proper wrapping around the 16-bit address space
		finalAddress := (address + uint16(cpu.Y)) & 0xFFFF
		// Get the value at the final address
		value = readMemory(finalAddress)
		setFlags()

	}
}
func JMP(addressingMode string) {
	cpu.previousPC = cpu.PC
	cpu.previousOpcode = cpu.opcode()
	cpu.previousOperand1 = cpu.operand1()
	cpu.previousOperand2 = cpu.operand2()
	cpu.updateCycleCounter(3)
	cpu.handleState(0)
	switch addressingMode {
	case ABSOLUTE:
		// Get the 16 bit address from operands
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Set the program counter to the absolute address
		cpu.setPC(address)
	case INDIRECT:
		// Get the 16 bit address from operands
		loByteAddress := uint16(cpu.operand1()) | uint16(cpu.operand2())<<8
		// Correctly emulate the 6502 JMP indirect page boundary hardware bug
		hiByteAddress := (loByteAddress & 0xFF00) | uint16((loByteAddress+1)&0x00FF)
		// Fetch the address to jump to, considering the page boundary bug
		indirectAddress := uint16(readMemory(loByteAddress)) | uint16(readMemory(hiByteAddress))<<8
		// Set the program counter to the indirect address
		cpu.setPC(indirectAddress)
	}
	if *klausd {
		//if readMemory(0x02) == 0xDE && readMemory(0x03) == 0xB0 {
		//	//fmt.println("All tests passed!")
		//	os.Exit(0)
		//}
		//fmt.Fprintf(os.Stderr, "Klaus D. loop detected at address 0x%04X\n", cpu.PC)
		if cpu.PC == KlausDInfiniteLoopAddress {
			fmt.Printf("Klaus D. loop detected at address $%04X\n", KlausDInfiniteLoopAddress)
			os.Exit(0)
		}
	}
}
func AND(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		if readBit(7, result) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = cpu.operand1()
		// AND the value with the accumulator
		result = cpu.A & value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		// Get the address from the operand
		address := cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// AND the value with the accumulator
		result = cpu.A & value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX:
		// Get address
		address := cpu.operand1() + cpu.X
		// Get value at address
		value = readMemory(uint16(address))
		// AND the value with the accumulator
		result = cpu.A & value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get value at address
		value = readMemory(address)
		// AND the value with the accumulator
		result = cpu.A & value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX:
		// Get address
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1()) + uint16(cpu.X)
		// Get value at address
		value = readMemory(address)
		// AND the value with the accumulator
		result = cpu.A & value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEY:
		// Get the address
		address := int(cpu.operand2())<<8 | int(cpu.operand1()) + int(cpu.Y)
		// Get the value at the address
		value = readMemory(uint16(address))
		// AND the value with the accumulator
		result = cpu.A & value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case INDIRECTX:
		// Get the address
		indirectAddress := int(cpu.operand1()) + int(cpu.X)
		address := int(readMemory(uint16(indirectAddress))) + int(readMemory(uint16(indirectAddress+1)))<<8
		// Get the value from the address
		value = readMemory(uint16(address))
		// AND the value with the accumulator
		result = cpu.A & value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case INDIRECTY:
		// Get the 16bit address
		address := uint16(int(cpu.operand1()))
		// Get the indirect address
		indirectAddress1 := readMemory(address)
		indirectAddress2 := readMemory(address + 1)
		//indirectAddress := uint16(int(indirectAddress1)+int(indirectAddress2)<<8) + uint16(cpu.Y)
		indirectAddress := uint16(int(indirectAddress1)+int(indirectAddress2)<<8) + uint16(cpu.Y)
		// Get the value at the address
		value = readMemory(indirectAddress)
		// AND the value with the accumulator
		result = cpu.A & value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	}
}
func EOR(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		//if readBit(7, result) == 1 {
		if result&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = cpu.operand1()
		// XOR the value with the accumulator
		result = cpu.A ^ value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		// Get the address from the operand
		address := cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// XOR the value with the accumulator
		result = cpu.A ^ value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX:
		// Get address
		address := cpu.operand1() + cpu.X
		// Get value at address
		value = readMemory(uint16(address))
		// XOR the value with the accumulator
		result = cpu.A ^ value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get value at address
		value = readMemory(address)
		// XOR the value with the accumulator
		result = cpu.A ^ value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX:
		// Get address
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1()) + uint16(cpu.X)
		// Get value at address
		value = readMemory(address)
		// XOR the value with the accumulator
		result = cpu.A ^ value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEY:
		// Get the address
		address := int(cpu.operand2())<<8 | int(cpu.operand1()) + int(cpu.Y)
		// Get the value at the address
		value = readMemory(uint16(address))
		// XOR the value with the accumulator
		result = cpu.A ^ value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case INDIRECTX:
		// Get the zero page address to use as a pointer
		zpAddress := cpu.operand1()
		// Calculate the effective address using X register to get the LSB of the address
		effAddrLsb := readMemory(uint16(zpAddress+cpu.X) & 0x00FF)
		// Get the MSB of the address from the next zero page location
		effAddrMsb := readMemory(uint16(zpAddress+cpu.X+1) & 0x00FF)
		address := uint16(effAddrMsb)<<8 | uint16(effAddrLsb)
		// Get the value at the effective address
		value = readMemory(address)
		result = cpu.A ^ value
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case INDIRECTY:
		// Get the zero page address to use as a pointer
		zpAddress := cpu.operand1()
		// Get the LSB of the effective address from the zero page
		effAddrLsb := readMemory(uint16(zpAddress))
		// Get the MSB of the effective address from the next zero page location
		effAddrMsb := readMemory(uint16(zpAddress+1) & 0x00FF)
		// Calculate the final address using the Y register
		finalAddress := uint16(effAddrMsb)<<8 | uint16(effAddrLsb) + uint16(cpu.Y)
		// Get the value at the final address
		value = readMemory(finalAddress)
		result = cpu.A ^ value
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	}
}
func ORA(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		if readBit(7, result) == 1 {
			cpu.setNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = cpu.operand1()
		// OR the value with the accumulator
		result = cpu.A | value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		// Get the address from the operand
		address := cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// OR the value with the accumulator
		result = cpu.A | value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX:
		// Get address
		address := cpu.operand1() + cpu.X
		// Get value at address
		value = readMemory(uint16(address))
		// OR the value with the accumulator
		result = cpu.A | value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get value at address
		value = readMemory(address)
		// OR the value with the accumulator
		result = cpu.A | value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX:
		address := (uint16(cpu.operand1()) + uint16(cpu.X)) | uint16(cpu.operand2())<<8
		value = readMemory(address)
		cpu.A |= value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEY:
		address := (uint16(cpu.operand1()) + uint16(cpu.Y)) | uint16(cpu.operand2())<<8
		value = readMemory(address)
		cpu.A |= value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case INDIRECTX:
		// Get the zero page address to use as a pointer
		zpAddress := cpu.operand1()
		// Calculate the effective address using X register to get the LSB of the address
		effAddrLsb := readMemory(uint16(zpAddress+cpu.X) & 0x00FF)
		// Get the MSB of the address from the next zero page location
		effAddrMsb := readMemory(uint16(zpAddress+cpu.X+1) & 0x00FF)
		address := uint16(effAddrMsb)<<8 | uint16(effAddrLsb)
		// Get the value at the effective address
		value = readMemory(address)
		result = cpu.A | value
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case INDIRECTY:
		// Get the zero page address to use as a pointer
		zpAddress := cpu.operand1()
		// Get the LSB of the effective address from the zero page
		effAddrLsb := readMemory(uint16(zpAddress))
		// Get the MSB of the effective address from the next zero page location
		effAddrMsb := readMemory(uint16(zpAddress+1) & 0x00FF)
		// Calculate the final address using the Y register
		finalAddress := uint16(effAddrMsb)<<8 | uint16(effAddrLsb) + uint16(cpu.Y)
		// Get the value at the final address
		value = readMemory(finalAddress)
		result = cpu.A | value
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	}
}
func BIT(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// Set Negative flag to bit 7 of the value
		if readBit(7, value) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		// Set Overflow flag to bit 6 of the value
		if readBit(6, value) == 1 {
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
		// Get the address from the operand
		address := cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// AND the value with the accumulator
		result = cpu.A & value
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get value at address
		value = readMemory(address)
		// AND the value with the accumulator
		result = cpu.A & value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	}
}
func INC(addressingMode string) {
	var address uint16
	var result byte

	setFlags := func() {
		// Fetch the value from the address
		value := readMemory(address)
		// Increment the value (wrapping around for 8-bit values)
		result = value + 1
		// Write the result back to memory
		writeMemory(address, result)

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
		// Get the address from the operand
		address = uint16(cpu.operand1())
		setFlags()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	case ZEROPAGEX:
		// Get the address from the operand with X offset
		address = uint16(cpu.operand1() + cpu.X)
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get 16-bit address from operand1 and operand2
		address = uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(3)
	case ABSOLUTEX:
		// Get 16-bit address from operand1 and operand2 with X offset
		address = (uint16(cpu.operand2())<<8 | uint16(cpu.operand1())) + uint16(cpu.X)
		setFlags()
		cpu.updateCycleCounter(7)
		cpu.handleState(3)
	}
}
func DEC(addressingMode string) {
	var address uint16
	var result byte

	setFlags := func() {
		// Fetch the value from the address
		value := readMemory(address)
		// Decrement the value (wrapping around for 8-bit values)
		result = value - 1
		// Write the result back to memory
		writeMemory(address, result)

		// Update status flags
		// If bit 7 of the result is set, set the negative flag
		//if readBit(7, result) == 1 {
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
		// Get the address from the operand
		address = uint16(cpu.operand1())
		setFlags()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	case ZEROPAGEX:
		// Get the address from the operand with X offset
		address = uint16(cpu.operand1() + cpu.X)
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get 16-bit address from operand1 and operand2
		address = uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(3)
	case ABSOLUTEX:
		// Get 16-bit address from operand1 and operand2 with X offset
		address = (uint16(cpu.operand2())<<8 | uint16(cpu.operand1())) + uint16(cpu.X)
		setFlags()
		cpu.updateCycleCounter(7)
		cpu.handleState(3)
	}
}
func ADC(addressingMode string) {
	var value byte
	var result int

	setFlags := func() {
		// Binary mode is the default
		tmpResult := int(cpu.A) + int(value)
		if cpu.getSRBit(0) == 1 {
			tmpResult++
		}

		if cpu.getSRBit(3) == 1 { // BCD mode
			temp := (cpu.A & 0x0F) + (value & 0x0F) + cpu.getSRBit(0)
			if temp > 9 {
				temp += 6
			}

			result = int((cpu.A & 0xF0) + (value & 0xF0) + (temp & 0x0F))

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
		if (int(cpu.A)^int(value))&0x80 == 0 && (int(cpu.A)^tmpResult)&0x80 != 0 {
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

		// Your addressing mode cycle counts remain the same
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
		// Get the value from the operand
		value = cpu.operand1()
		setFlags()
	case ZEROPAGE:
		// Get the address from the operand
		address := cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case ZEROPAGEX:
		// Get the address from the operand
		address := cpu.operand1() + cpu.X
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get value at address
		value = readMemory(address)
		setFlags()
	case ABSOLUTEX:
		// Get 16 bit address from operand1 and operand2
		baseAddress := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Perform a dummy read at the base address
		_ = readMemory(baseAddress)
		// Add the X register to the base address for the actual read
		address := baseAddress + uint16(cpu.X)
		// Get value at address
		value = readMemory(address)
		setFlags()
	case ABSOLUTEY:
		// Get 16 bit address from operand1 and operand2
		baseAddress := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Perform a dummy read at the base address
		_ = readMemory(baseAddress)
		// Add the Y register to the base address for the actual read
		address := baseAddress + uint16(cpu.Y)
		// Get value at address
		value = readMemory(address)
		setFlags()
	case INDIRECTX:
		// Get the indirect address from the operand
		indirectAddress := cpu.operand1() + cpu.X
		// Get the address from the indirect address
		address := uint16(readMemory(uint16(indirectAddress+1)))<<8 | uint16(readMemory(uint16(indirectAddress)))
		// Get the value at the address
		value = readMemory(address)
		setFlags()
	case INDIRECTY:
		// Get the indirect address from the operand
		indirectAddress := cpu.operand1()
		// Get the address from the indirect address
		address := uint16(readMemory(uint16(indirectAddress+1)))<<8 | uint16(readMemory(uint16(indirectAddress))) + uint16(cpu.Y)
		// Get the value at the address
		value = readMemory(address)
		setFlags()
	}
}
func SBC(addressingMode string) {
	var value byte
	var result int

	setFlags := func() {
		// Check for BCD mode (Decimal flag set)
		if cpu.getSRBit(3) == 1 { // BCD mode
			lowNibble := (cpu.A & 0x0F) - (value & 0x0F) - (cpu.getSRBit(0) ^ 1)
			if lowNibble < 0 {
				lowNibble = (lowNibble - 6) & 0x0F // Only keep the low nibble
			}

			highNibble := (cpu.A & 0xF0) - (value & 0xF0) - (lowNibble & 0x10)
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
			result = int(cpu.A) - int(value)
			if cpu.getSRBit(0) == 0 {
				result--
			}

			// The carry flag in binary mode is the inverse of the borrow indicator
			if result >= 0 {
				cpu.setCarryFlag()
			} else {
				cpu.unsetCarryFlag()
			}
		}

		// Negative, and Zero flag checks remain the same
		if readBit(7, byte(result)) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		if result == 0 {
			cpu.setZeroFlag()
		}
		cpu.A = byte(result)

		// Your addressing mode cycle counts remain the same
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
		// Get the value from the operand
		value = cpu.operand1()
		setFlags()
	case ZEROPAGE:
		// Get the address from the operand
		address := cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case ZEROPAGEX:
		// Get the address from the operand
		address := cpu.operand1() + cpu.X
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get value at address
		value = readMemory(address)
		setFlags()
	case ABSOLUTEX:
		// Get 16 bit address from operand1 and operand2
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1()) + uint16(cpu.X)
		// Get value at address
		value = readMemory(address)
		setFlags()
	case ABSOLUTEY:
		// Get 16 bit address from operand1 and operand2
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1()) + uint16(cpu.Y)
		// Get value at address
		value = readMemory(address)
		setFlags()
	case INDIRECTX:
		// Get the indirect address from the operand
		indirectAddress := cpu.operand1() + cpu.X
		// Get the address from the indirect address
		address := uint16(readMemory(uint16(indirectAddress+1)))<<8 | uint16(readMemory(uint16(indirectAddress)))
		// Get the value at the address
		value = readMemory(address)
		setFlags()
	case INDIRECTY:
		// Get the indirect address from the operand
		indirectAddress := cpu.operand1()
		// Get the address from the indirect address
		address := uint16(readMemory(uint16(indirectAddress+1)))<<8 | uint16(readMemory(uint16(indirectAddress))) + uint16(cpu.Y)
		// Get the value at the address
		value = readMemory(address)
		setFlags()
	}
}
func ROR(addressingMode string) {
	var address, value, result byte
	var address16 uint16

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
			writeMemory(uint16(address), result)
			cpu.updateCycleCounter(5)
			cpu.handleState(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX {
			// Store the value back into memory
			writeMemory(address16, result)
			cpu.updateCycleCounter(6)
			cpu.handleState(3)
		}
	}
	switch addressingMode {
	case ACCUMULATOR:
		// Get value from accumulator
		value = cpu.A
		// Rotate right one bit
		result = value >> 1
		setFlags()
	case ZEROPAGE:
		// Get address
		address = cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value right 1 bit
		result = value >> 1
		setFlags()
	case ZEROPAGEX:
		// Get X indexed zero page address
		address = cpu.operand1() + cpu.X
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value right 1 bit
		result = value >> 1
		setFlags()
	case ABSOLUTE:
		// Get 16 bit address from operands
		address16 = uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get the value stored at the address in the operands
		value = readMemory(address16)
		// Shift the value right 1 bit
		result = value >> 1
		setFlags()
	case ABSOLUTEX:
		// Get 16 bit address
		address16 = uint16(cpu.operand2())<<8 | uint16(cpu.operand1()) + uint16(cpu.X)
		// Get value stored at address
		value = readMemory(address16)
		// Shift right the value by 1 bit
		result = value >> 1
		setFlags()
	}
}
func ROL(addressingMode string) {
	var address, value, result byte
	var address16 uint16

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
			writeMemory(uint16(address), result)
			cpu.updateCycleCounter(5)
			cpu.handleState(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX {
			// Store the value back into memory
			writeMemory(address16, result)
			cpu.updateCycleCounter(6)
			cpu.handleState(3)
		}
	}
	switch addressingMode {
	case ACCUMULATOR:
		// Get the value of the accumulator
		value = cpu.A
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | cpu.getSRBit(0)
		setFlags()
	case ZEROPAGE:
		// Get address
		address = cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | cpu.getSRBit(0)
		setFlags()
	case ZEROPAGEX:
		// Get X indexed zero page address
		address = cpu.operand1() + cpu.X
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | cpu.getSRBit(0)
		setFlags()
	case ABSOLUTE:
		// Get 16 bit address from operands
		address16 = uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get the value stored at the address in the operands
		value = readMemory(address16)
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | cpu.getSRBit(0)
		setFlags()
	case ABSOLUTEX:
		// Get 16bit X indexed absolute memory address
		address16 = uint16(cpu.operand2())<<8 | uint16(cpu.operand1()) + uint16(cpu.X)
		// Get the value stored at the address
		value = readMemory(address16)
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | cpu.getSRBit(0)
		setFlags()
	}
}
func LSR(addressingMode string) {
	var value, result byte

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
		value = cpu.A
		// Shift the value right 1 bit
		result = value >> 1
		// Store the result back into the accumulator
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(1)
	case ZEROPAGE:
		// Get address
		address := cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value right 1 bit
		result = value >> 1
		// Store the value back into memory
		writeMemory(uint16(address), result)
		setFlags()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	case ZEROPAGEX:
		// Get the X indexed address
		address := cpu.operand1() + cpu.X
		// Get the value at the X indexed address
		value = readMemory(uint16(address))
		// Shift the value right 1 bit
		result = value >> 1
		// Store the shifted value in memory
		writeMemory(uint16(address), result)
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operands
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get the value stored at the address in the operands
		value = readMemory(address)
		// Shift the value right 1 bit
		result = value >> 1
		// Store the shifted value back in memory
		writeMemory(address, result)
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(3)
	case ABSOLUTEX:
		// Dummy read for Absolute,X addressing mode if page boundary is crossed
		baseAddress := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		address := baseAddress + uint16(cpu.X)
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			// Perform a dummy read if a page boundary is crossed
			readMemory(baseAddress&0xFF00 | (address & 0x00FF))
		}
		value = readMemory(address)
		result = value >> 1
		// Write the result back to the calculated effective address
		writeMemory(address, result)
		setFlags()
		// Update cycle counter depending on page boundary crossing
		if (baseAddress & 0xFF00) != (address & 0xFF00) {
			cpu.updateCycleCounter(7) // Add 1 if page boundary is crossed
		} else {
			cpu.updateCycleCounter(6)
		}
		cpu.handleState(3)

	}
}
func ASL(addressingMode string) {
	var value, result byte

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
		value = cpu.A
		// Shift the value left 1 bit
		result = value << 1
		// Update the accumulator with the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(1)
	case ZEROPAGE:
		// Get address
		address := cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value left 1 bit
		result = value << 1
		// Store the value back into memory
		writeMemory(uint16(address), result)
		setFlags()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	case ZEROPAGEX:
		// Get the X indexed address
		address := cpu.operand1() + cpu.X
		// Get the value at the X indexed address
		value = readMemory(uint16(address))
		// Shift the value left 1 bit
		result = value << 1
		// Store the shifted value in memory
		writeMemory(uint16(address), result)
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operands
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get the value stored at the address in the operands
		value = readMemory(address)
		// Shift the value left 1 bit
		result = value << 1
		// Store the shifted value back in memory
		writeMemory(address, result)
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(3)
	case ABSOLUTEX:
		// Get the 16bit X indexed absolute memory address
		address := int(cpu.operand2())<<8 | int(cpu.operand1()) + int(cpu.X)
		// Get the value stored at the address
		value = readMemory(uint16(address))
		// Shift the value left 1 bit
		result = value << 1
		// Store the shifted value back in memory
		writeMemory(uint16(address), result)
		setFlags()
		cpu.updateCycleCounter(7)
		cpu.handleState(3)
	}
}
func CPX(addressingMode string) {
	var value, result byte

	setFlags := func() {
		// Set or clear the carry flag
		if cpu.X >= value {
			cpu.setCarryFlag()
		} else {
			cpu.unsetCarryFlag()
		}

		// Set or clear the zero flag
		if cpu.X == value {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}

		// Set or clear the negative flag based on the high bit of the subtraction result
		result = cpu.X - value
		if result&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get value from operand1
		value = cpu.operand1()
		// Compare X with value
		result = cpu.X - value
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		// Get address
		address := cpu.operand1()
		// Get value at address
		value = readMemory(uint16(address))
		// Store result of X-memory stored at cpu.operand1() in result variable
		result = cpu.X - value
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get address
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get value at address
		value = readMemory(address)
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(3)
	}
}
func CPY(addressingMode string) {
	var value, result byte

	setFlags := func() {
		// Set or clear the carry flag
		if cpu.Y >= value {
			cpu.setCarryFlag()
		} else {
			cpu.unsetCarryFlag()
		}

		// Set or clear the zero flag
		if cpu.Y == value {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}

		// Set or clear the negative flag based on the high bit of the subtraction result
		result = cpu.Y - value
		if result&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get value from operand1
		value = cpu.operand1()
		// Subtract operand from Y
		result = cpu.Y - cpu.operand1()
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		// Get address
		address := cpu.operand1()
		// Get value at address
		value = readMemory(uint16(address))
		// Store result of Y-memory stored at cpu.operand1() in result variable
		result = cpu.Y - value
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get address
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get value at address
		value = readMemory(address)
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(3)
	}
}

// 1 byte instructions with no operands

// Implied addressing mode instructions
/*
	In the implied addressing mode, the address containing the operand is implicitly stated in the operation code of the instruction.

	Bytes: 1
*/
func BRK() {
	/*
		BRK - Break Command
	*/
	cpu.BRKtrue = true

	disassembleOpcode()
	cpu.previousPC = cpu.PC
	cpu.previousOpcode = cpu.opcode()

	// Decrement SP and Push high byte of (PC+1) onto stack
	cpu.decSP()
	updateStack(byte((cpu.PC + 1) >> 8))

	// Decrement SP and Push low byte of (PC+1) onto stack
	cpu.decSP()
	updateStack(byte((cpu.PC + 1) & 0xFF))

	// Set a modified SR with the B flag for the pushed value
	modifiedSR := cpu.SR | 0x10
	// Decrement SP and Store modified SR on stack
	cpu.setBreakFlag()
	cpu.decSP()
	updateStack(modifiedSR)

	// Decrement SP and Store SR on stack
	cpu.decSP()
	updateStack(cpu.SR)

	// Set SR interrupt disable bit to 1
	cpu.setInterruptFlag()

	// Set PC to interrupt vector address
	cpu.setPC((uint16(readMemory(IRQVectorAddressHigh)) << 8) | uint16(readMemory(IRQVectorAddressLow)))
	cpu.updateCycleCounter(7)
	cpu.handleState(0)
}
func CLC() {
	/*
		CLC - Clear Carry Flag
	*/
	//cpu.mnemonic = fmt.Sprintf("CLC\t")
	disassembleOpcode()
	// Set SR carry flag bit 0 to 0
	cpu.unsetCarryFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func CLD() {
	/*
		CLD - Clear Decimal Mode
	*/
	//cpu.mnemonic = fmt.Sprintf("CLD\t")
	disassembleOpcode()
	cpu.unsetDecimalFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func CLI() {
	/*
		CLI - Clear Interrupt Disable
	*/
	//cpu.mnemonic = fmt.Sprintf("CLI\t")
	disassembleOpcode()
	// Set SR interrupt disable bit 2 to 0
	cpu.unsetInterruptFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func CLV() {
	/*
		CLV - Clear Overflow Flag
	*/
	//cpu.mnemonic = fmt.Sprintf("CLV\t")
	disassembleOpcode()
	// Set SR overflow flag bit 6 to 0
	cpu.unsetOverflowFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func DEX() {
	// DEX - Decrement Index Register X By One
	//cpu.mnemonic = fmt.Sprintf("DEX\t")
	disassembleOpcode()

	// Decrement the X register by 1
	cpu.X--

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
func DEY() {
	/*
		DEY - Decrement Index Register Y By One
	*/
	//cpu.mnemonic = fmt.Sprintf("DEY\t")
	disassembleOpcode()

	// Decrement the  Y register by 1
	cpu.Y--
	// If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
	if cpu.getYBit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	// If Y==0 then set the Zero flag
	if cpu.Y == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func INX() {
	/*
		INX - Increment Index Register X By One
	*/
	//cpu.mnemonic = fmt.Sprintf("INX\t")
	disassembleOpcode()

	// Increment the X register by 1
	cpu.X++
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
func INY() {
	/*
		INY - Increment Index Register Y By One
	*/
	//cpu.mnemonic = fmt.Sprintf("INY\t")
	disassembleOpcode()

	// Increment the  Y register by 1
	cpu.Y++
	// If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
	if cpu.getYBit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	// If Y register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
	if cpu.Y == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func NOP() {
	/*
		NOP - No Operation
	*/
	//cpu.mnemonic = fmt.Sprintf("NOP\t")
	disassembleOpcode()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func PHA() {
	/*
		PHA - Push Accumulator On Stack
	*/
	//cpu.mnemonic = fmt.Sprintf("PHA\t")
	disassembleOpcode()

	// Update memory address pointed to by SP with value stored in accumulator
	updateStack(cpu.A)
	cpu.decSP()
	cpu.updateCycleCounter(3)
	cpu.handleState(1)
}
func PHP() {
	/*
	   PHP - Push Processor Status On Stack
	*/
	disassembleOpcode()

	// Set the break flag and the unused bit only for the push operation
	status := cpu.SR | (1 << 4) // Set break flag
	status |= (1 << 5)          // Set unused bit

	// Push the status onto the stack
	updateStack(status)

	// Decrement the stack pointer
	cpu.decSP()
	cpu.updateCycleCounter(3)
	cpu.handleState(1)
}
func PLA() {
	/*
	   PLA - Pull Accumulator From Stack
	*/
	disassembleOpcode()

	// Increment the stack pointer first
	cpu.incSP()

	// Read the value from the stack into the accumulator
	cpu.A = readStack()

	// No flag updates should occur here

	cpu.updateCycleCounter(4)
	cpu.handleState(1)
}
func PLP() {
	/*
		PLP - Pull Processor Status From Stack
	*/
	disassembleOpcode()

	// Read the status from the stack
	newStatus := readStack()

	// Preserve break flag and unused bit from current status
	newStatus = (newStatus & 0xCF) | (cpu.SR & 0x30)

	// Update SR with the new status
	cpu.SR = newStatus

	// Increment the stack pointer after the operation
	cpu.incSP()

	cpu.updateCycleCounter(4)
	cpu.handleState(1)
}
func RTI() {
	/*
	   RTI - Return From Interrupt
	*/

	//cpu.mnemonic = fmt.Sprintf("RTI\t")
	disassembleOpcode()

	cpu.SR = readStack() & 0xCF
	cpu.incSP()

	// Increment the stack pointer to get low byte of PC
	cpu.incSP()

	// Get low byte of PC
	low := uint16(readStack())

	// Increment the stack pointer to get high byte of PC
	cpu.incSP()

	// Get high byte of PC
	high := uint16(readStack())

	cpu.previousPC = cpu.PC
	cpu.previousOpcode = cpu.opcode()
	cpu.updateCycleCounter(6)
	cpu.handleState(0)
	// Update PC with the value stored in memory at the address pointed to by SP
	cpu.setPC((high << 8) | low)
}
func RTS() {
	/*
		RTS - Return From Subroutine
	*/
	//cpu.mnemonic = fmt.Sprintf("RTS\t")
	disassembleOpcode()
	//Get low byte of new PC
	low := uint16(readStack())
	// Increment the stack pointer
	cpu.incSP()
	//Get high byte of new PC
	high := uint16(readStack())
	cpu.previousPC = cpu.PC
	cpu.previousOpcode = cpu.opcode()
	//Update PC with the value stored in memory at the address pointed to by SP
	cpu.setPC((high << 8) | low + 1)
	cpu.updateCycleCounter(6)
	cpu.handleState(0)
}
func SEC() {
	/*
		SEC - Set Carry Flag
	*/
	//cpu.mnemonic = fmt.Sprintf("SEC\t")
	disassembleOpcode()

	// Set SR carry flag bit 0 to 1
	cpu.setCarryFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func SED() {
	/*
		SED - Set Decimal Mode
	*/
	//cpu.mnemonic = fmt.Sprintf("SED\t")
	disassembleOpcode()

	// Set SR decimal mode flag to 1
	cpu.setDecimalFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func SEI() {
	/*
		SEI - Set Interrupt Disable
	*/
	//cpu.mnemonic = fmt.Sprintf("SEI\t")
	disassembleOpcode()

	// Set SR interrupt disable bit 2 to 1
	cpu.setInterruptFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func TAX() {
	/*
		TAX - Transfer Accumulator To Index X
	*/
	//cpu.mnemonic = fmt.Sprintf("TAX\t")
	disassembleOpcode()

	// Update X with the value of A
	cpu.X = cpu.A
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
func TAY() {
	/*
		TAY - Transfer Accumulator To Index Y
	*/
	//cpu.mnemonic = fmt.Sprintf("TAY\t")
	disassembleOpcode()

	// Set Y register to the value of the accumulator
	cpu.Y = cpu.A
	// If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
	if cpu.getYBit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	// If Y register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
	if cpu.A == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func TSX() {
	/*
		TSX - Transfer Stack Pointer To Index X
	*/
	//cpu.mnemonic = fmt.Sprintf("TSX\t")
	disassembleOpcode()

	// Update X with the SP
	cpu.X = byte(cpu.SP)
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
func TXA() {
	/*
		TXA - Transfer Index X To Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("TXA\t")
	disassembleOpcode()

	// Set accumulator to value of X register
	cpu.A = cpu.X
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
func TXS() {
	/*
		TXS - Transfer Index X To Stack Pointer
	*/
	//cpu.mnemonic = fmt.Sprintf("TXS\t")
	disassembleOpcode()

	// Set stack pointer to value of X register
	cpu.SP = uint16(cpu.X)
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func TYA() {
	/*
		TYA - Transfer Index Y To Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("TYA\t")
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
func ASL_A() {
	/*
		ASL - Arithmetic Shift Left
	*/
	//cpu.mnemonic = fmt.Sprintf("ASL\t")
	disassembleOpcode()
	ASL("accumulator")
}
func LSR_A() {
	/*
		LSR - Logical Shift Right
	*/
	//cpu.mnemonic = fmt.Sprintf("LSR\t")
	disassembleOpcode()
	LSR("accumulator")
}
func ROL_A() {
	/*
		ROL - Rotate Left
	*/
	//cpu.mnemonic = fmt.Sprintf("ROL\t")
	disassembleOpcode()
	ROL("accumulator")
}
func ROR_A() {
	/*
		ROR - Rotate Right
	*/
	//cpu.mnemonic = fmt.Sprintf("ROR\t")
	disassembleOpcode()
	ROR("accumulator")
}

// 2 byte instructions with 1 operand
// Immediate addressing mode instructions
/*
	#$nn

	In immediate addressing, the operand is contained in the second byte of the instruction, with no further memory addressing required.

	Bytes: 2
*/
func ADC_I() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	//cpu.mnemonic = fmt.Sprintf("ADC #$%02X", cpu.operand1())
	disassembleOpcode()

	ADC("immediate")
}
func AND_I() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("AND #$%02X", cpu.operand1())
	disassembleOpcode()
	AND("immediate")
}
func CMP_I() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("CMP #$%02X", cpu.operand1())
	disassembleOpcode()
	CMP("immediate")
}
func CPX_I() {
	/*
		CPX - Compare Index Register X To Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("CPX #$%02X", cpu.operand1())
	disassembleOpcode()

	CPX("immediate")
}
func CPY_I() {
	/*
		CPY - Compare Index Register Y To Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("CPY #$%02X", cpu.operand1())
	disassembleOpcode()
	CPY("immediate")
}
func EOR_I() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("EOR #$%02X", cpu.operand1())
	disassembleOpcode()
	EOR("immediate")
}
func LDA_I() {
	/*
		LDA - Load Accumulator with Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDA #$%02X", cpu.operand1())
	disassembleOpcode()
	LDA("immediate")
}
func LDX_I() {
	/*
		LDX - Load Index Register X From Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDX #$%02X", cpu.operand1())
	disassembleOpcode()
	LDX("immediate")
}
func LDY_I() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDY #$%02X", cpu.operand1())
	disassembleOpcode()
	LDY("immediate")
}
func ORA_I() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("ORA #$%02X", cpu.operand1())
	disassembleOpcode()
	ORA("immediate")
}
func SBC_I() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	//cpu.mnemonic = fmt.Sprintf("SBC #$%02X", cpu.operand1())
	disassembleOpcode()
	SBC("immediate")
}

// Zero Page addressing mode instructions
/*
	$nn

	The zero page instructions allow for shorter code and execution times by only fetching the second byte of the instruction and assuming a zero low address byte. Careful use of the zero page can result in significant increase in code efficiency.

	Bytes: 2
*/
func ADC_Z() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	//cpu.mnemonic = fmt.Sprintf("ADC $%02X", cpu.operand1())
	disassembleOpcode()
	ADC("zeropage")
}
func AND_Z() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("AND $%02X", cpu.operand1())
	disassembleOpcode()
	AND("zeropage")
}
func ASL_Z() {
	/*
		ASL - Arithmetic Shift Left
	*/
	//cpu.mnemonic = fmt.Sprintf("ASL $%02X", cpu.operand1())
	disassembleOpcode()

	ASL("zeropage")
}
func BIT_Z() {
	/*
		BIT - Test Bits in Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("BIT $%02X", cpu.operand1())
	disassembleOpcode()
	BIT("zeropage")
}
func CMP_Z() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("CMP $%02X", cpu.operand1())
	disassembleOpcode()
	CMP("zeropage")
}
func CPX_Z() {
	/*
		CPX - Compare Index Register X To Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("CPX $%02X", cpu.operand1())
	disassembleOpcode()
	CPX("zeropage")
}
func CPY_Z() {
	/*
		CPY - Compare Index Register Y To Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("CPY $%02X", cpu.operand1())
	disassembleOpcode()
	CPY("zeropage")
}
func DEC_Z() {
	/*
		DEC - Decrement Memory By One
	*/
	//cpu.mnemonic = fmt.Sprintf("DEC $%02X", cpu.operand1())
	disassembleOpcode()
	DEC("zeropage")
}
func EOR_Z() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("EOR $%02X", cpu.operand1())
	disassembleOpcode()
	EOR("zeropage")
}
func INC_Z() {
	/*
		INC - Increment Memory By One
	*/
	//cpu.mnemonic = fmt.Sprintf("INC $%02X", cpu.operand1())
	disassembleOpcode()
	INC("zeropage")
}
func LDA_Z() {
	/*
		LDA - Load Accumulator with Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDA $%02X", cpu.operand1())
	disassembleOpcode()
	LDA("zeropage")
}
func LDX_Z() {
	/*
		LDX - Load Index Register X From Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDX $%02X", cpu.operand1())
	disassembleOpcode()
	LDX("zeropage")
}
func LDY_Z() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDY $%02X", cpu.operand1())
	disassembleOpcode()
	LDY("zeropage")
}
func LSR_Z() {
	/*
		LSR - Logical Shift Right
	*/
	//cpu.mnemonic = fmt.Sprintf("LSR $%02X", cpu.operand1())
	disassembleOpcode()
	LSR("zeropage")
}
func ORA_Z() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("ORA $%02X", cpu.operand1())
	disassembleOpcode()
	ORA("zeropage")
}
func ROL_Z() {
	/*
		ROL - Rotate Left
	*/
	//cpu.mnemonic = fmt.Sprintf("ROL $%02X", cpu.operand1())
	disassembleOpcode()
	ROL("zeropage")
}
func ROR_Z() {
	/*
		ROR - Rotate Right
	*/
	//cpu.mnemonic = fmt.Sprintf("ROR $%02X", cpu.operand1())
	disassembleOpcode()

	ROR("zeropage")
}
func SBC_Z() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	//cpu.mnemonic = fmt.Sprintf("SBC $%02X", cpu.operand1())
	disassembleOpcode()
	SBC("zeropage")
}
func STA_Z() {
	/*
		STA - Store Accumulator in Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("STA $%02X", cpu.operand1())
	disassembleOpcode()
	STA("zeropage")
}
func STX_Z() {
	/*
		STX - Store Index Register X In Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("STX $%02X", cpu.operand1())
	disassembleOpcode()
	STX("zeropage")
}
func STY_Z() {
	/*
		STY - Store Index Register Y In Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("STY $%02X", cpu.operand1())
	disassembleOpcode()
	STY("zeropage")
}

// X Indexed Zero Page addressing mode instructions
/*
	$nn,X

	This form of addressing is used in conjunction with the X index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the low order 8 bits of memory and crossing of page boundaries does not occur.

	Bytes: 2
*/
func ADC_ZX() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	//cpu.mnemonic = fmt.Sprintf("ADC $%02X,X", cpu.operand1())
	disassembleOpcode()
	ADC("zeropagex")
}
func AND_ZX() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("AND $%02X,X", cpu.operand1())
	disassembleOpcode()
	AND("zeropagex")
}
func ASL_ZX() {
	/*
		ASL - Arithmetic Shift Left
	*/
	//cpu.mnemonic = fmt.Sprintf("ASL $%02X,X", cpu.operand1())
	disassembleOpcode()
	ASL("zeropagex")
}
func CMP_ZX() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("CMP $%02X,X", cpu.operand1())
	disassembleOpcode()
	CMP("zeropagex")
}
func DEC_ZX() {
	/*
		DEC - Decrement Memory By One
	*/
	//cpu.mnemonic = fmt.Sprintf("DEC $%02X,X", cpu.operand1())
	disassembleOpcode()
	DEC("zeropagex")
}
func LDA_ZX() {
	/*
		LDA - Load Accumulator with Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDA $%02X,X", cpu.operand1())
	disassembleOpcode()
	LDA("zeropagex")
}
func LDY_ZX() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDY $%02X,X", cpu.operand1())
	disassembleOpcode()
	LDY("zeropagex")
}
func LSR_ZX() {
	/*
		LSR - Logical Shift Right
	*/
	//cpu.mnemonic = fmt.Sprintf("LSR $%02X,X", cpu.operand1())
	disassembleOpcode()
	LSR("zeropagex")
}
func ORA_ZX() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("ORA $%02X,X", cpu.operand1())
	disassembleOpcode()
	ORA("zeropagex")
}
func ROL_ZX() {
	/*
		ROL - Rotate Left
	*/
	//cpu.mnemonic = fmt.Sprintf("ROL $%02X,X", cpu.operand1())
	disassembleOpcode()
	ROL("zeropagex")
}
func ROR_ZX() {
	/*
		ROR - Rotate Right
	*/
	//cpu.mnemonic = fmt.Sprintf("ROR $%02X,X", cpu.operand1())
	disassembleOpcode()
	ROR("zeropagex")
}
func EOR_ZX() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("EOR $%02X,X", cpu.operand1())
	disassembleOpcode()
	EOR("zeropagex")
}
func INC_ZX() {
	/*
		INC - Increment Memory By One
	*/
	//cpu.mnemonic = fmt.Sprintf("INC $%02X,X", cpu.operand1())
	disassembleOpcode()
	INC("zeropagex")
}
func SBC_ZX() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	//cpu.mnemonic = fmt.Sprintf("SBC $%02X,X", cpu.operand1())
	disassembleOpcode()
	SBC("zeropagex")
}
func STA_ZX() {
	/*
		STA - Store Accumulator in Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("STA $%02X,X", cpu.operand1())
	disassembleOpcode()
	STA("zeropagex")
}
func STY_ZX() {
	/*
		STY - Store Index Register Y In Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("STY $%02X,X", cpu.operand1())
	disassembleOpcode()
	STY("zeropagex")
}

// Y Indexed Zero Page addressing mode instructions
/*
	$nn,Y

	This form of addressing is used in conjunction with the Y index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the low order 8 bits of memory and crossing of page boundaries does not occur.

	Bytes: 2
*/
func LDX_ZY() {
	/*
		LDX - Load Index Register X From Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDX $%02X,Y", cpu.operand1())
	disassembleOpcode()
	LDX("zeropagey")
}
func STX_ZY() {
	/*
		STX - Store Index Register X In Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("STX $%02X,Y", cpu.operand1())
	disassembleOpcode()
	STX("zeropagey")
}

// X Indexed Zero Page Indirect addressing mode instructions
/*
	($nn,X)

	In indexed indirect addressing, the second byte of the instruction is added to the contents of the X index register, discarding the carry. The result of this addition points to a memory location on page zero whose contents is the high order eight bits of the effective address. The next memory location in page zero contains the low order eight bits of the effective address. Both memory locations specifying the low and high order bytes of the effective address must be in page zero.

	Bytes: 2
*/
func ADC_IX() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	//cpu.mnemonic = fmt.Sprintf("ADC ($%02X,X)", cpu.operand1())
	disassembleOpcode()
	ADC("indirectx")
}
func AND_IX() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("AND ($%02X,X)", cpu.operand1())
	disassembleOpcode()
	AND("indirectx")
}
func CMP_IX() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("CMP ($%02X,X)", cpu.operand1())
	disassembleOpcode()
	CMP("indirectx")
}
func EOR_IX() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("EOR ($%02X,X)", cpu.operand1())
	disassembleOpcode()
	EOR("indirectx")
}
func LDA_IX() {
	/*
		LDA - Load Accumulator with Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDA ($%02X,X)", cpu.operand1())
	disassembleOpcode()
	LDA("indirectx")
}
func ORA_IX() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("ORA ($%02X,X)", cpu.operand1())
	disassembleOpcode()
	ORA("indirectx")
}
func SBC_IX() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	//cpu.mnemonic = fmt.Sprintf("SBC ($%02X,X)", cpu.operand1())
	disassembleOpcode()
	SBC("indirectx")
}
func STA_IX() {
	/*
		STA - Store Accumulator in Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("STA ($%02X,X)", cpu.operand1())
	disassembleOpcode()
	STA("indirectx")
}

// Zero Page Indirect Y Indexed addressing mode instructions
/*
	($nn),Y

	In indirect indexed addressing, the second byte of the instruction points to a memory location in page zero. The contents of this memory location is added to the contents of the Y index register, the result being the high order eight bits of the effective address. The carry from this addition is added to the contents of the next page zero memory location, the result being the low order eight bits of the effective address.

	Bytes: 2
*/
func ADC_IY() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	//cpu.mnemonic = fmt.Sprintf("ADC ($%02X),Y", cpu.operand1())
	disassembleOpcode()
	ADC("indirecty")
}
func AND_IY() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("AND ($%02X),Y", cpu.operand1())
	disassembleOpcode()
	AND("indirecty")
}
func CMP_IY() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("CMP ($%02X),Y", cpu.operand1())
	disassembleOpcode()
	CMP("indirecty")
}
func EOR_IY() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("EOR ($%02X),Y", cpu.operand1())
	disassembleOpcode()
	EOR("indirecty")
}
func LDA_IY() {
	/*
		LDA - Load Accumulator with Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDA ($%02X),Y", cpu.operand1())
	disassembleOpcode()
	LDA("indirecty")
}
func ORA_IY() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("ORA ($%02X),Y", cpu.operand1())
	disassembleOpcode()
	ORA("indirecty")
}
func SBC_IY() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	//cpu.mnemonic = fmt.Sprintf("SBC ($%02X),Y", cpu.operand1())
	disassembleOpcode()
	SBC("indirecty")
}
func STA_IY() {
	/*
		STA - Store Accumulator in Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("STA ($%02X),Y", cpu.operand1())
	disassembleOpcode()
	STA("indirecty")
}

// Relative addressing mode instructions
/*
	$nnnn

	Relative addressing is used only with branch instructions and establishes a destination for the conditional branch.

	The second byte of-the instruction becomes the operand which is an “Offset" added to the contents of the lower eight bits of the program counter when the counter is set at the next instruction. The range of the offset is —128 to +127 bytes from the next instruction.

	Bytes: 2
*/
func BPL_R() {
	/*
		BPL - Branch on Result Plus
	*/
	disassembleOpcode()
	offset := cpu.operand1()
	signedOffset := int8(offset)
	// Calculate the branch target address
	targetAddress := cpu.PC + 2 + uint16(signedOffset)
	// If N flag is not set, branch to address
	if cpu.getSRBit(7) == 0 {
		oldPC := cpu.PC
		cpu.PC = targetAddress
		cpu.updateCycleCounter(1) // Every branch takes at least one cycle
		cpu.incrementCycleCountForBranch(oldPC)
		cpu.instructionCounter++ // Increment the instruction counter if branch is taken
	} else {
		// Don't branch
		cpu.updateCycleCounter(2) // Two cycles when branch is not taken
		cpu.handleState(2)
	}
}
func BMI_R() {
	/*
	   BMI - Branch on Result Minus
	*/
	disassembleOpcode()
	offset := int8(cpu.operand1()) // Get offset from operand

	// If N flag is set, branch to address
	if cpu.getSRBit(7) == 1 {
		oldPC := cpu.PC
		targetAddress := uint16(int16(cpu.PC) + 2 + int16(offset)) // Calculate the target address

		cpu.PC = targetAddress
		cpu.updateCycleCounter(1) // Every branch takes at least one cycle

		// Increment the cycle count for branch and page crossing
		cpu.incrementCycleCountForBranch(oldPC)
	} else {
		// Don't branch, skip to the next instruction
		cpu.incPC(2)
		cpu.updateCycleCounter(2) // Two cycles when branch is not taken
	}
}
func BVC_R() {
	/*
	   BVC - Branch on Overflow Clear
	*/
	disassembleOpcode()
	offset := int8(cpu.operand1()) // Get offset from operand

	// If overflow flag is not set, branch to address
	if cpu.getSRBit(6) == 0 {
		oldPC := cpu.PC
		targetAddress := uint16(int16(cpu.PC) + 2 + int16(offset)) // Calculate the target address

		cpu.PC = targetAddress
		cpu.updateCycleCounter(1) // Every branch takes at least one cycle

		// Increment the cycle count for branch and page crossing
		cpu.incrementCycleCountForBranch(oldPC)
	} else {
		// Don't branch, skip to the next instruction
		cpu.incPC(2)
		cpu.updateCycleCounter(2) // Two cycles when branch is not taken
	}
}
func BVS_R() {
	/*
	   BVS - Branch on Overflow Set
	*/
	disassembleOpcode()
	offset := int8(cpu.operand1()) // Get offset from operand

	// If overflow flag is set, branch to address
	if cpu.getSRBit(6) == 1 {
		oldPC := cpu.PC
		targetAddress := uint16(int16(cpu.PC) + 2 + int16(offset)) // Calculate the target address

		cpu.PC = targetAddress
		cpu.updateCycleCounter(1) // Every branch takes at least one cycle

		// Increment the cycle count for branch and page crossing
		cpu.incrementCycleCountForBranch(oldPC)
	} else {
		// Don't branch, skip to the next instruction
		cpu.incPC(2)
		cpu.updateCycleCounter(2) // Two cycles when branch is not taken
	}
}
func BCC_R() {
	/*
	   BCC - Branch on Carry Clear
	*/
	disassembleOpcode()

	offset := int8(cpu.operand1())                             // Get offset from operand
	targetAddress := uint16(int16(cpu.PC) + 2 + int16(offset)) // Calculate the target address

	// If carry flag is clear, branch to address
	if cpu.getSRBit(0) == 0 {
		oldPC := cpu.PC
		cpu.PC = targetAddress

		cpu.updateCycleCounter(1) // Every branch takes at least one cycle

		// Increment the cycle count for branch and page crossing
		cpu.incrementCycleCountForBranch(oldPC)
	} else {
		// Don't branch, skip to the next instruction
		cpu.incPC(2)
		cpu.updateCycleCounter(2) // Two cycles when branch is not taken
	}
}
func BCS_R() {
	/*
	   BCS - Branch on Carry Set
	*/
	disassembleOpcode()

	offset := int8(cpu.operand1())                             // Get offset as signed 8-bit integer
	targetAddress := uint16(int16(cpu.PC) + 2 + int16(offset)) // Calculate the target address

	// If carry flag is set, branch to address
	if cpu.getSRBit(0) == 1 {
		oldPC := cpu.PC
		cpu.PC = targetAddress

		cpu.updateCycleCounter(1) // Every branch takes at least one cycle

		// Increment the cycle count for branch and page crossing
		cpu.incrementCycleCountForBranch(oldPC)
	} else {
		// Don't branch, skip to the next instruction
		cpu.incPC(2)
		cpu.updateCycleCounter(2) // Two cycles when branch is not taken
	}
}
func BNE_R() {
	/*
	   BNE - Branch on Result Not Zero
	*/
	disassembleOpcode()

	offset := int8(cpu.operand1())                             // Cast to signed 8-bit to handle negative offsets
	targetAddress := uint16(int16(cpu.PC) + 2 + int16(offset)) // Calculate the target address

	// Check Z flag to determine if branching is needed
	if cpu.getSRBit(1) == 0 {
		oldPC := cpu.PC
		cpu.PC = targetAddress

		cpu.updateCycleCounter(1) // Every branch takes at least one cycle

		// Increment the cycle count for branch and page crossing
		cpu.incrementCycleCountForBranch(oldPC)
	} else {
		// If Z flag is set, don't branch, skip to the next instruction
		cpu.incPC(2)
		cpu.updateCycleCounter(2) // Two cycles when branch is not taken
	}
}
func BEQ_R() {
	/*
	   BEQ - Branch on Result Zero
	*/
	disassembleOpcode()

	offset := int8(cpu.operand1())                             // Cast to signed 8-bit to handle negative offsets
	targetAddress := uint16(int16(cpu.PC) + 2 + int16(offset)) // Calculate the target address

	// If Z flag is set, branch to address
	if cpu.getSRBit(1) == 1 {
		oldPC := cpu.PC
		cpu.PC = targetAddress

		cpu.updateCycleCounter(1) // Every branch takes at least one cycle

		// Increment the cycle count for branch and page crossing
		cpu.incrementCycleCountForBranch(oldPC)
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
func ADC_ABS() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	//cpu.mnemonic = fmt.Sprintf("ADC $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ADC("absolute")
}
func AND_ABS() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("AND $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	AND("absolute")
}
func ASL_ABS() {
	/*
		ASL - Arithmetic Shift Left
	*/
	//cpu.mnemonic = fmt.Sprintf("ASL $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ASL("absolute")
}
func BIT_ABS() {
	/*
		BIT - Test Bits in Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("BIT $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	BIT("absolute")
}
func CMP_ABS() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("CMP $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	CMP("absolute")
}
func CPX_ABS() {
	/*
		CPX - Compare Index Register X To Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("CPX $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	CPX("absolute")
}
func CPY_ABS() {
	/*
		CPY - Compare Index Register Y To Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("CPY $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	CPY("absolute")
}
func DEC_ABS() {
	/*
		DEC - Decrement Memory By One
	*/
	//cpu.mnemonic = fmt.Sprintf("DEC $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	DEC("absolute")
}
func EOR_ABS() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("EOR $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	EOR("absolute")
}
func INC_ABS() {
	/*
		INC - Increment Memory By One
	*/
	//cpu.mnemonic = fmt.Sprintf("INC $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	INC("absolute")
}
func JMP_ABS() {
	/*
		JMP - JMP Absolute
	*/
	//cpu.mnemonic = fmt.Sprintf("JMP $%04X", int(cpu.operand2())<<8|int(cpu.operand1()))
	disassembleOpcode()
	// For AllSuiteA.bin 6502 opcode test suite
	if *allsuitea && readMemory(0x210) == 0xFF {
		fmt.Printf("\n\u001B[32;5mMemory address $210 == $%02X. All opcodes succesfully tested and passed!\u001B[0m\n", readMemory(0x210))
		os.Exit(0)
	}
	JMP("absolute")
}
func JSR_ABS() {
	/*
		JSR - Jump To Subroutine
	*/
	//cpu.mnemonic = fmt.Sprintf("JSR $%04X", int(cpu.operand2())<<8|int(cpu.operand1()))
	disassembleOpcode()
	// First, push the high byte
	cpu.decSP()
	updateStack(byte(cpu.PC >> 8))
	cpu.decSP()
	updateStack(byte((cpu.PC)&0xFF) + 2)

	cpu.previousPC = cpu.PC
	cpu.previousOpcode = cpu.opcode()
	cpu.previousOperand1 = cpu.operand1()
	cpu.previousOperand2 = cpu.operand2()
	// Now, jump to the subroutine address specified by the operands
	cpu.setPC(uint16(cpu.operand2())<<8 | uint16(cpu.operand1()))
	cpu.updateCycleCounter(1)
	cpu.handleState(0)
}
func LDA_ABS() {
	/*
		LDA - Load Accumulator with Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDA $%04X", uint16(cpu.operand2())<<8|uint16(cpu.operand1()))
	disassembleOpcode()
	LDA("absolute")
}
func LDX_ABS() {
	/*
		LDX - Load Index Register X From Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDX $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	LDX("absolute")
}
func LDY_ABS() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDY $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	LDY("absolute")
}
func LSR_ABS() {
	/*
		LSR - Logical Shift Right
	*/
	//cpu.mnemonic = fmt.Sprintf("LSR $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	LSR("absolute")
}
func ORA_ABS() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("ORA $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ORA("absolute")
}
func ROL_ABS() {
	/*
		ROL - Rotate Left
	*/
	//cpu.mnemonic = fmt.Sprintf("ROL $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ROL("absolute")
}
func ROR_ABS() {
	/*
		ROR - Rotate Right
	*/
	//cpu.mnemonic = fmt.Sprintf("ROR $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ROR("absolute")
}
func SBC_ABS() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	//cpu.mnemonic = fmt.Sprintf("SBC $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	SBC("absolute")
}
func STA_ABS() {
	/*
		STA - Store Accumulator in Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("STA $%04X", uint16(cpu.operand2())<<8|uint16(cpu.operand1()))
	disassembleOpcode()
	STA("absolute")
}
func STX_ABS() {
	/*
		STX - Store Index Register X In Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("STX $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	STX("absolute")
}
func STY_ABS() {
	/*
		STY - Store Index Register Y In Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("STY $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	STY("absolute")
}

// X Indexed Absolute addressing mode instructions
/*
	$nnnn,X

	This form of addressing is used in conjunction with the X index register. The effective address is formed by adding the contents of X to the address contained in the second and third bytes of the instruction. This mode allows the index register to contain the index or count value and the instruction to contain the base address. This type of indexing allows any location referencing and the index to modify multiple fields resulting in reduced coding and execution time.

	Note on the MOS 6502:

	The value at the specified address, ignoring the the addressing mode's X offset, is read (and discarded) before the final address is read. This may cause side effects in I/O registers.


	Bytes: 3
*/
func ADC_ABX() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	//cpu.mnemonic = fmt.Sprintf("ADC $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ADC("absolutex")
}
func AND_ABX() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("AND $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	AND("absolutex")
}
func ASL_ABX() {
	/*
		ASL - Arithmetic Shift Left
	*/
	//cpu.mnemonic = fmt.Sprintf("ASL $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ASL("absolutex")
}
func CMP_ABX() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("CMP $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	CMP("absolutex")
}
func DEC_ABX() {
	/*
		DEC - Decrement Memory By One
	*/
	//cpu.mnemonic = fmt.Sprintf("DEC $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	DEC("absolutex")
}
func EOR_ABX() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("EOR $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	EOR("absolutex")
}
func INC_ABX() {
	/*
		INC - Increment Memory By One
	*/
	//cpu.mnemonic = fmt.Sprintf("INC $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	INC("absolutex")
}
func LDA_ABX() {
	/*
		LDA - Load Accumulator with Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDA $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	LDA("absolutex")
}
func LDY_ABX() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDY $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	LDY("absolutex")
}
func LSR_ABX() {
	/*
		LSR - Logical Shift Right
	*/
	//cpu.mnemonic = fmt.Sprintf("LSR $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	LSR("absolutex")
}
func ORA_ABX() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("ORA $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ORA("absolutex")
}
func ROL_ABX() {
	/*
	 */
	//cpu.mnemonic = fmt.Sprintf("ROL $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ROL("absolutex")
}
func ROR_ABX() {
	/*
		ROR - Rotate Right
	*/
	//cpu.mnemonic = fmt.Sprintf("ROR $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ROR("absolutex")
}
func SBC_ABX() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	//cpu.mnemonic = fmt.Sprintf("SBC $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	SBC("absolutex")
}
func STA_ABX() {
	/*
		STA - Store Accumulator in Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("STA $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	STA("absolutex")
}

// Y Indexed Absolute addressing mode instructions
/*
	$nnnn,Y

	This form of addressing is used in conjunction with the Y index register. The effective address is formed by adding the contents of Y to the address contained in the second and third bytes of the instruction. This mode allows the index register to contain the index or count value and the instruction to contain the base address. This type of indexing allows any location referencing and the index to modify multiple fields resulting in reduced coding and execution time.

	Note on the MOS 6502:

	The value at the specified address, ignoring the the addressing mode's Y offset, is read (and discarded) before the final address is read. This may cause side effects in I/O registers.

	Bytes: 3
*/
func ADC_ABY() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	//cpu.mnemonic = fmt.Sprintf("ADC $%04X,Y", int(cpu.operand2())<<8|int(cpu.operand1()))
	disassembleOpcode()
	ADC("absolutey")
}
func AND_ABY() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("AND $%02X%02X,Y", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	AND("absolutey")
}
func CMP_ABY() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("CMP $%02X%02X,Y", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	CMP("absolutey")
}
func EOR_ABY() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("EOR $%02X%02X,Y", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	EOR("absolutey")
}
func LDA_ABY() {
	/*
		LDA - Load Accumulator with Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDA $%02X%02X,Y", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	LDA("absolutey")
}
func LDX_ABY() {
	/*
		LDX - Load Index Register X From Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("LDX $%02X%02X,Y", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	LDX("absolutey")
}
func ORA_ABY() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	//cpu.mnemonic = fmt.Sprintf("ORA $%02X%02X,Y", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ORA("absolutey")
}
func SBC_ABY() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	//cpu.mnemonic = fmt.Sprintf("SBC $%02X%02X,Y", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	SBC("absolutey")
}
func STA_ABY() {
	/*
		STA - Store Accumulator in Memory
	*/
	//cpu.mnemonic = fmt.Sprintf("STA $%02X%02X,Y", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	STA("absolutey")
}

// Absolute Indirect addressing mode instructions
func JMP_IND() {
	/*
		JMP - JMP Indirect
	*/
	//cpu.mnemonic = fmt.Sprintf("JMP ($%02X%02X)", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	JMP("indirect")
}
