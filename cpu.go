package main

import (
	"time"
)

const (
	ACCUMULATOR = "accumulator"
	IMMEDIATE   = "immediate"
	ZEROPAGE    = "zeropage"
	ZEROPAGEX   = "zeropagex"
	ZEROPAGEY   = "zeropagey"
	ABSOLUTE    = "absolute"
	ABSOLUTEX   = "absolutex"
	ABSOLUTEY   = "absolutey"
	INDIRECT    = "indirect"
	INDIRECTX   = "indirectx"
	INDIRECTY   = "indirecty"

	c64basicROMAddress  = 0xA000
	c64kernalROMAddress = 0xE000
	c64charROMAddress   = 0xD000

	plus4basicROMAddress  = 0x8000
	plus4kernalROMAddress = 0xC000
	plus4charROMAddress   = 0xC000
	threePlus1ROMAddress  = 0x8000

	AllSuiteAROMAddress       = 0x4000
	KlausDTestROMAddress      = 0x0000
	KlausDInfiniteLoopAddress = 0x062B
	RuudBTestROMAddress       = 0xE000

	SPBaseAddress uint16 = 0x0100
)
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
	// CPURegisters and RAM
	A                byte        = 0x0  // Accumulator
	X                byte        = 0x0  // X register
	Y                byte        = 0x0  // Y register		(76543210) SR Bit 5 is always set
	SR               byte               // Status Register	(NVEBDIZC)
	SP               uint16      = 0xFF // Stack Pointer
	PC               int                // Program Counter
	memory           [65536]byte        // Memory
	previousPC       int
	previousOpcode   byte
	previousOperand1 byte
	previousOperand2 byte
	irq              bool
	nmi              bool
	reset            bool
	BRKtrue          bool = false

	cycleCounter   uint64        = 0
	cpuSpeedHz     uint64        = 985248                                  // 985248 Hz for a standard 6502
	cycleTime      time.Duration = time.Second / time.Duration(cpuSpeedHz) // time per cycle in nanoseconds
	cycleStartTime time.Time                                               // High-resolution timer
	timeSpent      time.Duration                                           // Time spent executing instructions

)

func resetCPU() {
	cycleCounter = 0
	SP = SPBaseAddress
	// Set SR to 0b00110100
	SR = 0b00110110
	if *klausd {
		setPC(0x400)
	} else {
		// Set PC to value stored at reset vector address
		setPC(int(readMemory(RESETVectorAddress)) + int(readMemory(RESETVectorAddress+1))*256)
	}
}

func opcode() byte {
	return readMemory(uint16(PC))
}
func operand1() byte {
	return readMemory(uint16(PC + 1))
}
func operand2() byte {
	return readMemory(uint16(PC + 2))
}

func incSP() {
	if SP == 0xFF {
		// Wrap around from 0xFF to 0x00
		SP = 0x00
	} else {
		SP++
	}
}
func decSP() {
	if SP == 0x00 {
		// Wrap around from 0x00 to 0xFF
		SP = 0xFF
	} else {
		SP--
	}
}

func incPC(amount int) {
	PC += amount
	if PC > 0xFFFF {
		PC = 0x0000 + (PC & 0xFFFF)
	}
}
func decPC(amount int) {
	PC -= amount
	if PC < 0 {
		PC = 0xFFFF + (PC & 0xFFFF)
	}
}
func setPC(newAddress int) {
	PC = newAddress & 0xFFFF
}

func handleIRQ() {
	if getSRBit(2) == 1 {
		return
	}
	// Push PC onto stack
	updateStack(byte(PC >> 8)) // high byte
	decSP()
	updateStack(byte(PC & 0xFF)) // low byte
	decSP()
	// Push SR onto stack
	updateStack(SR)
	decSP()
	// Set interrupt flag
	setInterruptFlag()
	// Set PC to IRQ Service Routine address
	setPC(int(readMemory(IRQVectorAddress)) | int(readMemory(IRQVectorAddress+1))<<8)
	irq = false
}
func handleNMI() {
	// Push PC onto stack
	updateStack(byte(PC >> 8)) // high byte
	decSP()
	updateStack(byte(PC & 0xFF)) // low byte
	decSP()
	// Push SR onto stack
	updateStack(SR)
	decSP()
	// Set PC to NMI Service Routine address
	setPC(int(readMemory(NMIVectorAddress)) | int(readMemory(NMIVectorAddress+1))<<8)
	nmi = false // Clear the NMI flag
}
func handleRESET() {
	resetCPU()
	reset = false // Clear the RESET flag
}

func updateCycleCounter(amount uint64) {
	cycleCounter += amount
}
func cycleStart() {
	cycleStartTime = time.Now() // High-resolution timer
}
func cycleEnd() {
	// Calculate the time we should wait
	elapsedTime := time.Since(cycleStartTime)
	expectedTime := time.Duration(cycleCounter) * cycleTime
	remainingTime := expectedTime - elapsedTime

	// Wait for the remaining time if needed
	if remainingTime > 0 {
		time.Sleep(remainingTime)
	}
	timeSpent = time.Now().Sub(cycleStartTime)
}

func startCPU() {
	for PC < len(memory) {
		//  1 byte instructions with no operands
		switch opcode() {
		// Implied addressing mode instructions
		/*
			In the implied addressing mode, the address containing the operand is implicitly stated in the operation code of the instruction.

			Bytes: 1
		*/
		case BRK_OPCODE:
			cycleStart()
			BRK()
			cycleEnd()
		case CLC_OPCODE:
			cycleStart()
			CLC()
			cycleEnd()
		case CLD_OPCODE:
			cycleStart()
			CLD()
			cycleEnd()
		case CLI_OPCODE:
			cycleStart()
			CLI()
			cycleEnd()
		case CLV_OPCODE:
			cycleStart()
			CLV()
			cycleEnd()
		case DEX_OPCODE:
			cycleStart()
			DEX()
			cycleEnd()
		case DEY_OPCODE:
			cycleStart()
			DEY()
			cycleEnd()
		case INX_OPCODE:
			cycleStart()
			INX()
			cycleEnd()
		case INY_OPCODE:
			cycleStart()
			INY()
			cycleEnd()
		case NOP_OPCODE:
			cycleStart()
			NOP()
			cycleEnd()
		case PHA_OPCODE:
			cycleStart()
			PHA()
			cycleEnd()
		case PHP_OPCODE:
			cycleStart()
			PHP()
			cycleEnd()
		case PLA_OPCODE:
			cycleStart()
			PLA()
			cycleEnd()
		case PLP_OPCODE:
			cycleStart()
			PLP()
			cycleEnd()
		case RTI_OPCODE:
			cycleStart()
			RTI()
			cycleEnd()
		case RTS_OPCODE:
			cycleStart()
			RTS()
			cycleEnd()
		case SEC_OPCODE:
			cycleStart()
			SEC()
			cycleEnd()
		case SED_OPCODE:
			cycleStart()
			SED()
			cycleEnd()
		case SEI_OPCODE:
			cycleStart()
			SEI()
			cycleEnd()
		case TAX_OPCODE:
			cycleStart()
			TAX()
			cycleEnd()
		case TAY_OPCODE:
			cycleStart()
			TAY()
			cycleEnd()
		case TSX_OPCODE:
			cycleStart()
			TSX()
			cycleEnd()
		case TXA_OPCODE:
			cycleStart()
			TXA()
			cycleEnd()
		case TXS_OPCODE:
			cycleStart()
			TXS()
			cycleEnd()
		case TYA_OPCODE:
			cycleStart()
			TYA()
			cycleEnd()

		// Accumulator instructions
		/*
			A

			This form of addressing is represented with a one byte instruction, implying an operation on the accumulator.

			Bytes: 1
		*/
		case ASL_ACCUMULATOR_OPCODE:
			cycleStart()
			ASL_A()
			cycleEnd()
		case LSR_ACCUMULATOR_OPCODE:
			cycleStart()
			LSR_A()
			cycleEnd()
		case ROL_ACCUMULATOR_OPCODE:
			cycleStart()
			ROL_A()
			cycleEnd()
		case ROR_ACCUMULATOR_OPCODE:
			cycleStart()
			ROR_A()
			cycleEnd()
		}

		// 2 byte instructions with 1 operand
		switch opcode() {
		// Immediate addressing mode instructions
		/*
			#$nn

			In immediate addressing, the operand is contained in the second byte of the instruction, with no further memory addressing required.

			Bytes: 2
		*/
		case ADC_IMMEDIATE_OPCODE:
			cycleStart()
			ADC_I()
			cycleEnd()
		case AND_IMMEDIATE_OPCODE:
			cycleStart()
			AND_I()
			cycleEnd()
		case CMP_IMMEDIATE_OPCODE:
			cycleStart()
			CMP_I()
			cycleEnd()
		case CPX_IMMEDIATE_OPCODE:
			cycleStart()
			CPX_I()
			cycleEnd()
		case CPY_IMMEDIATE_OPCODE:
			cycleStart()
			CPY_I()
			cycleEnd()
		case EOR_IMMEDIATE_OPCODE:
			cycleStart()
			EOR_I()
			cycleEnd()
		case LDA_IMMEDIATE_OPCODE:
			cycleStart()
			LDA_I()
			cycleEnd()
		case LDX_IMMEDIATE_OPCODE:
			cycleStart()
			LDX_I()
			cycleEnd()
		case LDY_IMMEDIATE_OPCODE:
			cycleStart()
			LDY_I()
			cycleEnd()
		case ORA_IMMEDIATE_OPCODE:
			cycleStart()
			ORA_I()
			cycleEnd()
		case SBC_IMMEDIATE_OPCODE:
			cycleStart()
			SBC_I()
			cycleEnd()

		// Zero Page addressing mode instructions
		/*
			$nn

			The zero page instructions allow for shorter code and execution times by only fetching the second byte of the instruction and assuming a zero low address byte. Careful use of the zero page can result in significant increase in code efficiency.

			Bytes: 2
		*/
		case ADC_ZERO_PAGE_OPCODE:
			cycleStart()
			ADC_Z()
			cycleEnd()
		case AND_ZERO_PAGE_OPCODE:
			cycleStart()
			AND_Z()
			cycleEnd()
		case ASL_ZERO_PAGE_OPCODE:
			cycleStart()
			ASL_Z()
			cycleEnd()
		case BIT_ZERO_PAGE_OPCODE:
			cycleStart()
			BIT_Z()
			cycleEnd()
		case CMP_ZERO_PAGE_OPCODE:
			cycleStart()
			CMP_Z()
			cycleEnd()
		case CPX_ZERO_PAGE_OPCODE:
			cycleStart()
			CPX_Z()
			cycleEnd()
		case CPY_ZERO_PAGE_OPCODE:
			cycleStart()
			CPY_Z()
			cycleEnd()
		case DEC_ZERO_PAGE_OPCODE:
			cycleStart()
			DEC_Z()
			cycleEnd()
		case EOR_ZERO_PAGE_OPCODE:
			cycleStart()
			EOR_Z()
			cycleEnd()
		case INC_ZERO_PAGE_OPCODE:
			cycleStart()
			INC_Z()
			cycleEnd()
		case LDA_ZERO_PAGE_OPCODE:
			cycleStart()
			LDA_Z()
			cycleEnd()
		case LDX_ZERO_PAGE_OPCODE:
			cycleStart()
			LDX_Z()
			cycleEnd()
		case LDY_ZERO_PAGE_OPCODE:
			cycleStart()
			LDY_Z()
			cycleEnd()
		case LSR_ZERO_PAGE_OPCODE:
			cycleStart()
			LSR_Z()
			cycleEnd()
		case ORA_ZERO_PAGE_OPCODE:
			cycleStart()
			ORA_Z()
			cycleEnd()
		case ROL_ZERO_PAGE_OPCODE:
			cycleStart()
			ROL_Z()
			cycleEnd()
		case ROR_ZERO_PAGE_OPCODE:
			cycleStart()
			ROR_Z()
			cycleEnd()
		case SBC_ZERO_PAGE_OPCODE:
			cycleStart()
			SBC_Z()
			cycleEnd()
		case STA_ZERO_PAGE_OPCODE:
			cycleStart()
			STA_Z()
			cycleEnd()
		case STX_ZERO_PAGE_OPCODE:
			cycleStart()
			STX_Z()
			cycleEnd()
		case STY_ZERO_PAGE_OPCODE:
			cycleStart()
			STY_Z()
			cycleEnd()

		// X Indexed Zero Page addressing mode instructions
		/*
			$nn,X

			This form of addressing is used in conjunction with the X index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the low order 8 bits of memory and crossing of page boundaries does not occur.

			Bytes: 2
		*/
		case ADC_ZERO_PAGE_X_OPCODE:
			cycleStart()
			ADC_ZX()
			cycleEnd()
		case AND_ZERO_PAGE_X_OPCODE:
			cycleStart()
			AND_ZX()
			cycleEnd()
		case ASL_ZERO_PAGE_X_OPCODE:
			cycleStart()
			ASL_ZX()
			cycleEnd()
		case CMP_ZERO_PAGE_X_OPCODE:
			cycleStart()
			CMP_ZX()
			cycleEnd()
		case DEC_ZERO_PAGE_X_OPCODE:
			cycleStart()
			DEC_ZX()
			cycleEnd()
		case LDA_ZERO_PAGE_X_OPCODE:
			cycleStart()
			LDA_ZX()
			cycleEnd()
		case LDY_ZERO_PAGE_X_OPCODE:
			cycleStart()
			LDY_ZX()
			cycleEnd()
		case LSR_ZERO_PAGE_X_OPCODE:
			cycleStart()
			LSR_ZX()
			cycleEnd()
		case ORA_ZERO_PAGE_X_OPCODE:
			cycleStart()
			ORA_ZX()
			cycleEnd()
		case ROL_ZERO_PAGE_X_OPCODE:
			cycleStart()
			ROL_ZX()
			cycleEnd()
		case ROR_ZERO_PAGE_X_OPCODE:
			cycleStart()
			ROR_ZX()
			cycleEnd()
		case EOR_ZERO_PAGE_X_OPCODE:
			cycleStart()
			EOR_ZX()
			cycleEnd()
		case INC_ZERO_PAGE_X_OPCODE:
			cycleStart()
			INC_ZX()
			cycleEnd()
		case SBC_ZERO_PAGE_X_OPCODE:
			cycleStart()
			SBC_ZX()
			cycleEnd()
		case STA_ZERO_PAGE_X_OPCODE:
			cycleStart()
			STA_ZX()
			cycleEnd()
		case STY_ZERO_PAGE_X_OPCODE:
			cycleStart()
			STY_ZX()
			cycleEnd()

		// Y Indexed Zero Page addressing mode instructions
		/*
			$nn,Y

			This form of addressing is used in conjunction with the Y index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the low order 8 bits of memory and crossing of page boundaries does not occur.

			Bytes: 2
		*/
		case LDX_ZERO_PAGE_Y_OPCODE:
			cycleStart()
			LDX_ZY()
			cycleEnd()
		case STX_ZERO_PAGE_Y_OPCODE:
			cycleStart()
			STX_ZY()
			cycleEnd()

		// X Indexed Zero Page Indirect addressing mode instructions
		/*
			($nn,X)

			In indexed indirect addressing, the second byte of the instruction is added to the contents of the X index register, discarding the carry. The result of this addition points to a memory location on page zero whose contents is the high order eight bits of the effective address. The next memory location in page zero contains the low order eight bits of the effective address. Both memory locations specifying the low and high order bytes of the effective address must be in page zero.

			Bytes: 2
		*/
		case ADC_INDIRECT_X_OPCODE:
			cycleStart()
			ADC_IX()
			cycleEnd()
		case AND_INDIRECT_X_OPCODE:
			cycleStart()
			AND_IX()
			cycleEnd()
		case CMP_INDIRECT_X_OPCODE:
			cycleStart()
			CMP_IX()
			cycleEnd()
		case EOR_INDIRECT_X_OPCODE:
			cycleStart()
			EOR_IX()
			cycleEnd()
		case LDA_INDIRECT_X_OPCODE:
			cycleStart()
			LDA_IX()
			cycleEnd()
		case ORA_INDIRECT_X_OPCODE:
			cycleStart()
			ORA_IX()
			cycleEnd()
		case SBC_INDIRECT_X_OPCODE:
			cycleStart()
			SBC_IX()
			cycleEnd()
		case STA_INDIRECT_X_OPCODE:
			cycleStart()
			STA_IX()
			cycleEnd()

		// Zero Page Indirect Y Indexed addressing mode instructions
		/*
			($nn),Y

			In indirect indexed addressing, the second byte of the instruction points to a memory location in page zero. The contents of this memory location is added to the contents of the Y index register, the result being the high order eight bits of the effective address. The carry from this addition is added to the contents of the next page zero memory location, the result being the low order eight bits of the effective address.

			Bytes: 2
		*/
		case ADC_INDIRECT_Y_OPCODE:
			cycleStart()
			ADC_IY()
			cycleEnd()
		case AND_INDIRECT_Y_OPCODE:
			cycleStart()
			AND_IY()
			cycleEnd()
		case CMP_INDIRECT_Y_OPCODE:
			cycleStart()
			CMP_IY()
			cycleEnd()
		case EOR_INDIRECT_Y_OPCODE:
			cycleStart()
			EOR_IY()
			cycleEnd()
		case LDA_INDIRECT_Y_OPCODE:
			cycleStart()
			LDA_IY()
			cycleEnd()
		case ORA_INDIRECT_Y_OPCODE:
			cycleStart()
			ORA_IY()
			cycleEnd()
		case SBC_INDIRECT_Y_OPCODE:
			cycleStart()
			SBC_IY()
			cycleEnd()
		case STA_INDIRECT_Y_OPCODE:
			cycleStart()
			STA_IY()
			cycleEnd()

		// Relative addressing mode instructions
		/*
			$nnnn

			Relative addressing is used only with branch instructions and establishes a destination for the conditional branch.

			The second byte of-the instruction becomes the operand which is an “Offset" added to the contents of the lower eight bits of the program counter when the counter is set at the next instruction. The range of the offset is —128 to +127 bytes from the next instruction.

			Bytes: 2
		*/
		case BPL_RELATIVE_OPCODE:
			cycleStart()
			BPL_R()
			cycleEnd()
		case BMI_RELATIVE_OPCODE:
			cycleStart()
			BMI_R()
			cycleEnd()
		case BVC_RELATIVE_OPCODE:
			cycleStart()
			BVC_R()
			cycleEnd()
		case BVS_RELATIVE_OPCODE:
			cycleStart()
			BVS_R()
			cycleEnd()
		case BCC_RELATIVE_OPCODE:
			cycleStart()
			BCC_R()
			cycleEnd()
		case BCS_RELATIVE_OPCODE:
			cycleStart()
			BCS_R()
			cycleEnd()
		case BNE_RELATIVE_OPCODE:
			cycleStart()
			BNE_R()
			cycleEnd()
		case BEQ_RELATIVE_OPCODE:
			cycleStart()
			BEQ_R()
			cycleEnd()
		}

		// 3 byte instructions with 2 operands
		switch opcode() {
		// Absolute addressing mode instructions
		/*
			$nnnn

			In absolute addressing, the second byte of the instruction specifies the eight high order bits of the effective address while the third byte specifies the eight low order bits. Thus, the absolute addressing mode allows access to the entire 65 K bytes of addressable memory.

			Bytes: 3
		*/
		case ADC_ABSOLUTE_OPCODE:
			cycleStart()
			ADC_ABS()
			cycleEnd()
		case AND_ABSOLUTE_OPCODE:
			cycleStart()
			AND_ABS()
			cycleEnd()
		case ASL_ABSOLUTE_OPCODE:
			cycleStart()
			ASL_ABS()
			cycleEnd()
		case BIT_ABSOLUTE_OPCODE:
			cycleStart()
			BIT_ABS()
			cycleEnd()
		case CMP_ABSOLUTE_OPCODE:
			cycleStart()
			CMP_ABS()
			cycleEnd()
		case CPX_ABSOLUTE_OPCODE:
			cycleStart()
			CPX_ABS()
			cycleEnd()
		case CPY_ABSOLUTE_OPCODE:
			cycleStart()
			CPY_ABS()
			cycleEnd()
		case DEC_ABSOLUTE_OPCODE:
			cycleStart()
			DEC_ABS()
			cycleEnd()
		case EOR_ABSOLUTE_OPCODE:
			cycleStart()
			EOR_ABS()
			cycleEnd()
		case INC_ABSOLUTE_OPCODE:
			cycleStart()
			INC_ABS()
			cycleEnd()
		case JMP_ABSOLUTE_OPCODE:
			cycleStart()
			JMP_ABS()
			cycleEnd()
		case JSR_ABSOLUTE_OPCODE:
			cycleStart()
			JSR_ABS()
			cycleEnd()
		case LDA_ABSOLUTE_OPCODE:
			cycleStart()
			LDA_ABS()
			cycleEnd()
		case LDX_ABSOLUTE_OPCODE:
			cycleStart()
			LDX_ABS()
			cycleEnd()
		case LDY_ABSOLUTE_OPCODE:
			cycleStart()
			LDY_ABS()
			cycleEnd()
		case LSR_ABSOLUTE_OPCODE:
			cycleStart()
			LSR_ABS()
			cycleEnd()
		case ORA_ABSOLUTE_OPCODE:
			cycleStart()
			ORA_ABS()
			cycleEnd()
		case ROL_ABSOLUTE_OPCODE:
			cycleStart()
			ROL_ABS()
			cycleEnd()
		case ROR_ABSOLUTE_OPCODE:
			cycleStart()
			ROR_ABS()
			cycleEnd()
		case SBC_ABSOLUTE_OPCODE:
			cycleStart()
			SBC_ABS()
			cycleEnd()
		case STA_ABSOLUTE_OPCODE:
			cycleStart()
			STA_ABS()
			cycleEnd()
		case STX_ABSOLUTE_OPCODE:
			cycleStart()
			STX_ABS()
			cycleEnd()
		case STY_ABSOLUTE_OPCODE:
			cycleStart()
			STY_ABS()
			cycleEnd()

		// X Indexed Absolute addressing mode instructions
		/*
			$nnnn,X

			This form of addressing is used in conjunction with the X index register. The effective address is formed by adding the contents of X to the address contained in the second and third bytes of the instruction. This mode allows the index register to contain the index or count value and the instruction to contain the base address. This type of indexing allows any location referencing and the index to modify multiple fields resulting in reduced coding and execution time.

			Note on the MOS 6502:

			The value at the specified address, ignoring the the addressing mode's X offset, is read (and discarded) before the final address is read. This may cause side effects in I/O registers.


			Bytes: 3
		*/
		case ADC_ABSOLUTE_X_OPCODE:
			cycleStart()
			ADC_ABX()
			cycleEnd()
		case AND_ABSOLUTE_X_OPCODE:
			cycleStart()
			AND_ABX()
			cycleEnd()
		case ASL_ABSOLUTE_X_OPCODE:
			cycleStart()
			ASL_ABX()
			cycleEnd()
		case CMP_ABSOLUTE_X_OPCODE:
			cycleStart()
			CMP_ABX()
			cycleEnd()
		case DEC_ABSOLUTE_X_OPCODE:
			cycleStart()
			DEC_ABX()
			cycleEnd()
		case EOR_ABSOLUTE_X_OPCODE:
			cycleStart()
			EOR_ABX()
			cycleEnd()
		case INC_ABSOLUTE_X_OPCODE:
			cycleStart()
			INC_ABX()
			cycleEnd()
		case LDA_ABSOLUTE_X_OPCODE:
			cycleStart()
			LDA_ABX()
			cycleEnd()
		case LDY_ABSOLUTE_X_OPCODE:
			cycleStart()
			LDY_ABX()
			cycleEnd()
		case LSR_ABSOLUTE_X_OPCODE:
			cycleStart()
			LSR_ABX()
			cycleEnd()
		case ORA_ABSOLUTE_X_OPCODE:
			cycleStart()
			ORA_ABX()
			cycleEnd()
		case ROL_ABSOLUTE_X_OPCODE:
			cycleStart()
			ROL_ABX()
			cycleEnd()
		case ROR_ABSOLUTE_X_OPCODE:
			cycleStart()
			ROR_ABX()
			cycleEnd()
		case SBC_ABSOLUTE_X_OPCODE:
			cycleStart()
			SBC_ABX()
			cycleEnd()
		case STA_ABSOLUTE_X_OPCODE:
			cycleStart()
			STA_ABX()
			cycleEnd()

		// Y Indexed Absolute addressing mode instructions
		/*
			$nnnn,Y

			This form of addressing is used in conjunction with the Y index register. The effective address is formed by adding the contents of Y to the address contained in the second and third bytes of the instruction. This mode allows the index register to contain the index or count value and the instruction to contain the base address. This type of indexing allows any location referencing and the index to modify multiple fields resulting in reduced coding and execution time.

			Note on the MOS 6502:

			The value at the specified address, ignoring the the addressing mode's Y offset, is read (and discarded) before the final address is read. This may cause side effects in I/O registers.

			Bytes: 3
		*/
		case ADC_ABSOLUTE_Y_OPCODE:
			cycleStart()
			ADC_ABY()
			cycleEnd()
		case AND_ABSOLUTE_Y_OPCODE:
			cycleStart()
			AND_ABY()
			cycleEnd()
		case CMP_ABSOLUTE_Y_OPCODE:
			cycleStart()
			CMP_ABY()
			cycleEnd()
		case EOR_ABSOLUTE_Y_OPCODE:
			cycleStart()
			EOR_ABY()
			cycleEnd()
		case LDA_ABSOLUTE_Y_OPCODE:
			cycleStart()
			LDA_ABY()
			cycleEnd()
		case LDX_ABSOLUTE_Y_OPCODE:
			cycleStart()
			LDX_ABY()
			cycleEnd()
		case ORA_ABSOLUTE_Y_OPCODE:
			cycleStart()
			ORA_ABY()
			cycleEnd()
		case SBC_ABSOLUTE_Y_OPCODE:
			cycleStart()
			SBC_ABY()
			cycleEnd()
		case STA_ABSOLUTE_Y_OPCODE:
			cycleStart()
			STA_ABY()
			cycleEnd()

		// Absolute Indirect addressing mode instructions
		case JMP_INDIRECT_OPCODE:
			cycleStart()
			JMP_IND()
			cycleEnd()
		}
		if *plus4 {
			plus4KernalRoutines()
		}
	}
}
