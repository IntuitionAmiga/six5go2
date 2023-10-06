package main

import (
	"fmt"
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

type CPU struct {
	A                byte   // Accumulator
	X                byte   // X register
	Y                byte   // Y register
	PC               uint16 // Program Counter
	SP               uint16 // Stack Pointer
	SR               byte   // Status Register
	previousPC       uint16
	previousOpcode   byte
	previousOperand1 byte
	previousOperand2 byte
	cycleCounter     uint64
	cycleStartTime   time.Time     // High-resolution timer
	cpuTimeSpent     time.Duration // Time spent executing instructions

	irq   bool
	nmi   bool
	reset bool
}

var (
	cpu CPU

	cpuSpeedHz uint64        = 985248                                  // 985248 Hz for a standard 6502
	cycleTime  time.Duration = time.Second / time.Duration(cpuSpeedHz) // time per cycle in nanoseconds

	memory [65536]byte // Memory

	BRKtrue bool = false
)

func (cpu *CPU) opcode() byte {
	return readMemory(cpu.PC)
}
func (cpu *CPU) operand1() byte {
	return readMemory(cpu.PC + 1)
}
func (cpu *CPU) operand2() byte {
	return readMemory(cpu.PC + 2)
}

func (cpu *CPU) incSP() {
	if cpu.SP == 0xFF {
		// Wrap around from 0xFF to 0x00
		cpu.SP = 0x00
	} else {
		cpu.SP++
	}
}
func (cpu *CPU) decSP() {
	if cpu.SP == 0x00 {
		// Wrap around from 0x00 to 0xFF
		cpu.SP = 0xFF
	} else {
		cpu.SP--
	}
}

func incPC(amount int) {
	cpu.PC += uint16(amount)
	if cpu.PC > 0xFFFF {
		cpu.PC = 0x0000 + (cpu.PC & 0xFFFF)
	}
}
func decPC(amount int) {
	cpu.PC -= uint16(amount)
	if cpu.PC < 0 {
		cpu.PC = 0xFFFF + (cpu.PC & 0xFFFF)
	}
}
func setPC(newAddress uint16) {
	cpu.PC = uint16(newAddress) & 0xFFFF
}

func (cpu *CPU) handleIRQ() {
	fmt.Println("Debug: Entering handleIRQ()")
	if cpu.getSRBit(2) == 1 {
		fmt.Println("Debug: Interrupt disabled. Exiting handleIRQ()")
		return
	}
	fmt.Println("Debug: Interrupt enabled. Continuing...")
	// Push PC onto stack
	updateStack(byte(cpu.PC >> 8)) // high byte
	cpu.decSP()
	updateStack(byte(cpu.PC & 0xFF)) // low byte
	cpu.decSP()
	fmt.Printf("Debug: PC pushed to stack. SP: %X\n", cpu.SP)
	// Push SR onto stack
	updateStack(cpu.SR)
	cpu.decSP()
	fmt.Printf("Debug: PC pushed to stack. SP: %X\n", cpu.SP)
	// Set interrupt flag
	cpu.setInterruptFlag()
	fmt.Printf("Debug: PC pushed to stack. SP: %X\n", cpu.SP)
	// Set PC to IRQ Service Routine address
	setPC(uint16(readMemory(IRQVectorAddress)) | uint16(readMemory(IRQVectorAddress+1))<<8)
	fmt.Printf("Debug: Jumping to IRQ Service Routine at %X\n", cpu.PC)
	cpu.irq = false
	fmt.Println("Debug: Exiting handleIRQ()")
}
func (cpu *CPU) handleNMI() {
	// Push PC onto stack
	updateStack(byte(cpu.PC >> 8)) // high byte
	cpu.decSP()
	updateStack(byte(cpu.PC & 0xFF)) // low byte
	cpu.decSP()
	// Push SR onto stack
	updateStack(cpu.SR)
	cpu.decSP()
	// Set PC to NMI Service Routine address
	setPC(uint16(readMemory(NMIVectorAddress)) | uint16(readMemory(NMIVectorAddress+1))<<8)
	cpu.nmi = false // Clear the NMI flag
}
func (cpu *CPU) handleRESET() {
	cpu.resetCPU()
	cpu.reset = false // Clear the RESET flag
}
func (cpu *CPU) handleState(amount int) {
	if *stateMonitor {
		printMachineState()
	}
	incPC(amount)
	// If amount is 0, then we are in a branch instruction and we don't want to increment the instruction counter
	if amount != 0 {
		instructionCounter++
	}
	if cpu.irq {
		cpu.handleIRQ()
	}
	if cpu.nmi {
		cpu.handleNMI()
	}
	if cpu.reset {
		cpu.handleRESET()
	}
}

func (cpu *CPU) updateCycleCounter(amount uint64) {
	cpu.cycleCounter += amount
}
func (cpu *CPU) cycleStart() {
	cpu.cycleStartTime = time.Now() // High-resolution timer
}
func (cpu *CPU) cycleEnd() {
	// Calculate the time we should wait
	elapsedTime := time.Since(cpu.cycleStartTime)
	expectedTime := time.Duration(cpu.cycleCounter) * cycleTime
	remainingTime := expectedTime - elapsedTime

	// Wait for the remaining time if needed
	if remainingTime > 0 {
		time.Sleep(remainingTime)
	}
	cpu.cpuTimeSpent = time.Now().Sub(cpu.cycleStartTime)
}

func (cpu *CPU) resetCPU() {
	cpu.cycleCounter = 0
	cpu.SP = SPBaseAddress
	// Set SR to 0b00110100
	cpu.SR = 0b00110110
	if *klausd {
		setPC(0x400)
	} else {
		// Set PC to value stored at reset vector address
		setPC(uint16(readMemory(RESETVectorAddress)) + uint16(readMemory(RESETVectorAddress+1))*256)
	}
}
func (cpu *CPU) startCPU() {
	for uint(cpu.PC) < 0xFFFF {
		//for cpu.PC {
		//  1 byte instructions with no operands
		switch cpu.opcode() {
		// Implied addressing mode instructions
		/*
			In the implied addressing mode, the address containing the operand is implicitly stated in the operation code of the instruction.

			Bytes: 1
		*/
		case BRK_OPCODE:
			cpu.cycleStart()
			BRK()
			cpu.cycleEnd()
		case CLC_OPCODE:
			cpu.cycleStart()
			CLC()
			cpu.cycleEnd()
		case CLD_OPCODE:
			cpu.cycleStart()
			CLD()
			cpu.cycleEnd()
		case CLI_OPCODE:
			cpu.cycleStart()
			CLI()
			cpu.cycleEnd()
		case CLV_OPCODE:
			cpu.cycleStart()
			CLV()
			cpu.cycleEnd()
		case DEX_OPCODE:
			cpu.cycleStart()
			DEX()
			cpu.cycleEnd()
		case DEY_OPCODE:
			cpu.cycleStart()
			DEY()
			cpu.cycleEnd()
		case INX_OPCODE:
			cpu.cycleStart()
			INX()
			cpu.cycleEnd()
		case INY_OPCODE:
			cpu.cycleStart()
			INY()
			cpu.cycleEnd()
		case NOP_OPCODE:
			cpu.cycleStart()
			NOP()
			cpu.cycleEnd()
		case PHA_OPCODE:
			cpu.cycleStart()
			PHA()
			cpu.cycleEnd()
		case PHP_OPCODE:
			cpu.cycleStart()
			PHP()
			cpu.cycleEnd()
		case PLA_OPCODE:
			cpu.cycleStart()
			PLA()
			cpu.cycleEnd()
		case PLP_OPCODE:
			cpu.cycleStart()
			PLP()
			cpu.cycleEnd()
		case RTI_OPCODE:
			cpu.cycleStart()
			RTI()
			cpu.cycleEnd()
		case RTS_OPCODE:
			cpu.cycleStart()
			RTS()
			cpu.cycleEnd()
		case SEC_OPCODE:
			cpu.cycleStart()
			SEC()
			cpu.cycleEnd()
		case SED_OPCODE:
			cpu.cycleStart()
			SED()
			cpu.cycleEnd()
		case SEI_OPCODE:
			cpu.cycleStart()
			SEI()
			cpu.cycleEnd()
		case TAX_OPCODE:
			cpu.cycleStart()
			TAX()
			cpu.cycleEnd()
		case TAY_OPCODE:
			cpu.cycleStart()
			TAY()
			cpu.cycleEnd()
		case TSX_OPCODE:
			cpu.cycleStart()
			TSX()
			cpu.cycleEnd()
		case TXA_OPCODE:
			cpu.cycleStart()
			TXA()
			cpu.cycleEnd()
		case TXS_OPCODE:
			cpu.cycleStart()
			TXS()
			cpu.cycleEnd()
		case TYA_OPCODE:
			cpu.cycleStart()
			TYA()
			cpu.cycleEnd()

		// Accumulator instructions
		/*
			A

			This form of addressing is represented with a one byte instruction, implying an operation on the accumulator.

			Bytes: 1
		*/
		case ASL_ACCUMULATOR_OPCODE:
			cpu.cycleStart()
			ASL_A()
			cpu.cycleEnd()
		case LSR_ACCUMULATOR_OPCODE:
			cpu.cycleStart()
			LSR_A()
			cpu.cycleEnd()
		case ROL_ACCUMULATOR_OPCODE:
			cpu.cycleStart()
			ROL_A()
			cpu.cycleEnd()
		case ROR_ACCUMULATOR_OPCODE:
			cpu.cycleStart()
			ROR_A()
			cpu.cycleEnd()
		}

		// 2 byte instructions with 1 operand
		switch cpu.opcode() {
		// Immediate addressing mode instructions
		/*
			#$nn

			In immediate addressing, the operand is contained in the second byte of the instruction, with no further memory addressing required.

			Bytes: 2
		*/
		case ADC_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			ADC_I()
			cpu.cycleEnd()
		case AND_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			AND_I()
			cpu.cycleEnd()
		case CMP_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			CMP_I()
			cpu.cycleEnd()
		case CPX_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			CPX_I()
			cpu.cycleEnd()
		case CPY_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			CPY_I()
			cpu.cycleEnd()
		case EOR_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			EOR_I()
			cpu.cycleEnd()
		case LDA_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			LDA_I()
			cpu.cycleEnd()
		case LDX_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			LDX_I()
			cpu.cycleEnd()
		case LDY_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			LDY_I()
			cpu.cycleEnd()
		case ORA_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			ORA_I()
			cpu.cycleEnd()
		case SBC_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			SBC_I()
			cpu.cycleEnd()

		// Zero Page addressing mode instructions
		/*
			$nn

			The zero page instructions allow for shorter code and execution times by only fetching the second byte of the instruction and assuming a zero low address byte. Careful use of the zero page can result in significant increase in code efficiency.

			Bytes: 2
		*/
		case ADC_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			ADC_Z()
			cpu.cycleEnd()
		case AND_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			AND_Z()
			cpu.cycleEnd()
		case ASL_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			ASL_Z()
			cpu.cycleEnd()
		case BIT_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			BIT_Z()
			cpu.cycleEnd()
		case CMP_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			CMP_Z()
			cpu.cycleEnd()
		case CPX_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			CPX_Z()
			cpu.cycleEnd()
		case CPY_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			CPY_Z()
			cpu.cycleEnd()
		case DEC_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			DEC_Z()
			cpu.cycleEnd()
		case EOR_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			EOR_Z()
			cpu.cycleEnd()
		case INC_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			INC_Z()
			cpu.cycleEnd()
		case LDA_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			LDA_Z()
			cpu.cycleEnd()
		case LDX_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			LDX_Z()
			cpu.cycleEnd()
		case LDY_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			LDY_Z()
			cpu.cycleEnd()
		case LSR_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			LSR_Z()
			cpu.cycleEnd()
		case ORA_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			ORA_Z()
			cpu.cycleEnd()
		case ROL_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			ROL_Z()
			cpu.cycleEnd()
		case ROR_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			ROR_Z()
			cpu.cycleEnd()
		case SBC_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			SBC_Z()
			cpu.cycleEnd()
		case STA_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			STA_Z()
			cpu.cycleEnd()
		case STX_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			STX_Z()
			cpu.cycleEnd()
		case STY_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			STY_Z()
			cpu.cycleEnd()

		// X Indexed Zero Page addressing mode instructions
		/*
			$nn,X

			This form of addressing is used in conjunction with the X index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the low order 8 bits of memory and crossing of page boundaries does not occur.

			Bytes: 2
		*/
		case ADC_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			ADC_ZX()
			cpu.cycleEnd()
		case AND_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			AND_ZX()
			cpu.cycleEnd()
		case ASL_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			ASL_ZX()
			cpu.cycleEnd()
		case CMP_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			CMP_ZX()
			cpu.cycleEnd()
		case DEC_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			DEC_ZX()
			cpu.cycleEnd()
		case LDA_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			LDA_ZX()
			cpu.cycleEnd()
		case LDY_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			LDY_ZX()
			cpu.cycleEnd()
		case LSR_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			LSR_ZX()
			cpu.cycleEnd()
		case ORA_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			ORA_ZX()
			cpu.cycleEnd()
		case ROL_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			ROL_ZX()
			cpu.cycleEnd()
		case ROR_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			ROR_ZX()
			cpu.cycleEnd()
		case EOR_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			EOR_ZX()
			cpu.cycleEnd()
		case INC_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			INC_ZX()
			cpu.cycleEnd()
		case SBC_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			SBC_ZX()
			cpu.cycleEnd()
		case STA_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			STA_ZX()
			cpu.cycleEnd()
		case STY_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			STY_ZX()
			cpu.cycleEnd()

		// Y Indexed Zero Page addressing mode instructions
		/*
			$nn,Y

			This form of addressing is used in conjunction with the Y index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the low order 8 bits of memory and crossing of page boundaries does not occur.

			Bytes: 2
		*/
		case LDX_ZERO_PAGE_Y_OPCODE:
			cpu.cycleStart()
			LDX_ZY()
			cpu.cycleEnd()
		case STX_ZERO_PAGE_Y_OPCODE:
			cpu.cycleStart()
			STX_ZY()
			cpu.cycleEnd()

		// X Indexed Zero Page Indirect addressing mode instructions
		/*
			($nn,X)

			In indexed indirect addressing, the second byte of the instruction is added to the contents of the X index register, discarding the carry. The result of this addition points to a memory location on page zero whose contents is the high order eight bits of the effective address. The next memory location in page zero contains the low order eight bits of the effective address. Both memory locations specifying the low and high order bytes of the effective address must be in page zero.

			Bytes: 2
		*/
		case ADC_INDIRECT_X_OPCODE:
			cpu.cycleStart()
			ADC_IX()
			cpu.cycleEnd()
		case AND_INDIRECT_X_OPCODE:
			cpu.cycleStart()
			AND_IX()
			cpu.cycleEnd()
		case CMP_INDIRECT_X_OPCODE:
			cpu.cycleStart()
			CMP_IX()
			cpu.cycleEnd()
		case EOR_INDIRECT_X_OPCODE:
			cpu.cycleStart()
			EOR_IX()
			cpu.cycleEnd()
		case LDA_INDIRECT_X_OPCODE:
			cpu.cycleStart()
			LDA_IX()
			cpu.cycleEnd()
		case ORA_INDIRECT_X_OPCODE:
			cpu.cycleStart()
			ORA_IX()
			cpu.cycleEnd()
		case SBC_INDIRECT_X_OPCODE:
			cpu.cycleStart()
			SBC_IX()
			cpu.cycleEnd()
		case STA_INDIRECT_X_OPCODE:
			cpu.cycleStart()
			STA_IX()
			cpu.cycleEnd()

		// Zero Page Indirect Y Indexed addressing mode instructions
		/*
			($nn),Y

			In indirect indexed addressing, the second byte of the instruction points to a memory location in page zero. The contents of this memory location is added to the contents of the Y index register, the result being the high order eight bits of the effective address. The carry from this addition is added to the contents of the next page zero memory location, the result being the low order eight bits of the effective address.

			Bytes: 2
		*/
		case ADC_INDIRECT_Y_OPCODE:
			cpu.cycleStart()
			ADC_IY()
			cpu.cycleEnd()
		case AND_INDIRECT_Y_OPCODE:
			cpu.cycleStart()
			AND_IY()
			cpu.cycleEnd()
		case CMP_INDIRECT_Y_OPCODE:
			cpu.cycleStart()
			CMP_IY()
			cpu.cycleEnd()
		case EOR_INDIRECT_Y_OPCODE:
			cpu.cycleStart()
			EOR_IY()
			cpu.cycleEnd()
		case LDA_INDIRECT_Y_OPCODE:
			cpu.cycleStart()
			LDA_IY()
			cpu.cycleEnd()
		case ORA_INDIRECT_Y_OPCODE:
			cpu.cycleStart()
			ORA_IY()
			cpu.cycleEnd()
		case SBC_INDIRECT_Y_OPCODE:
			cpu.cycleStart()
			SBC_IY()
			cpu.cycleEnd()
		case STA_INDIRECT_Y_OPCODE:
			cpu.cycleStart()
			STA_IY()
			cpu.cycleEnd()

		// Relative addressing mode instructions
		/*
			$nnnn

			Relative addressing is used only with branch instructions and establishes a destination for the conditional branch.

			The second byte of-the instruction becomes the operand which is an “Offset" added to the contents of the lower eight bits of the program counter when the counter is set at the next instruction. The range of the offset is —128 to +127 bytes from the next instruction.

			Bytes: 2
		*/
		case BPL_RELATIVE_OPCODE:
			cpu.cycleStart()
			BPL_R()
			cpu.cycleEnd()
		case BMI_RELATIVE_OPCODE:
			cpu.cycleStart()
			BMI_R()
			cpu.cycleEnd()
		case BVC_RELATIVE_OPCODE:
			cpu.cycleStart()
			BVC_R()
			cpu.cycleEnd()
		case BVS_RELATIVE_OPCODE:
			cpu.cycleStart()
			BVS_R()
			cpu.cycleEnd()
		case BCC_RELATIVE_OPCODE:
			cpu.cycleStart()
			BCC_R()
			cpu.cycleEnd()
		case BCS_RELATIVE_OPCODE:
			cpu.cycleStart()
			BCS_R()
			cpu.cycleEnd()
		case BNE_RELATIVE_OPCODE:
			cpu.cycleStart()
			BNE_R()
			cpu.cycleEnd()
		case BEQ_RELATIVE_OPCODE:
			cpu.cycleStart()
			BEQ_R()
			cpu.cycleEnd()
		}

		// 3 byte instructions with 2 operands
		switch cpu.opcode() {
		// Absolute addressing mode instructions
		/*
			$nnnn

			In absolute addressing, the second byte of the instruction specifies the eight high order bits of the effective address while the third byte specifies the eight low order bits. Thus, the absolute addressing mode allows access to the entire 65 K bytes of addressable memory.

			Bytes: 3
		*/
		case ADC_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			ADC_ABS()
			cpu.cycleEnd()
		case AND_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			AND_ABS()
			cpu.cycleEnd()
		case ASL_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			ASL_ABS()
			cpu.cycleEnd()
		case BIT_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			BIT_ABS()
			cpu.cycleEnd()
		case CMP_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			CMP_ABS()
			cpu.cycleEnd()
		case CPX_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			CPX_ABS()
			cpu.cycleEnd()
		case CPY_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			CPY_ABS()
			cpu.cycleEnd()
		case DEC_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			DEC_ABS()
			cpu.cycleEnd()
		case EOR_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			EOR_ABS()
			cpu.cycleEnd()
		case INC_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			INC_ABS()
			cpu.cycleEnd()
		case JMP_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			JMP_ABS()
			cpu.cycleEnd()
		case JSR_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			JSR_ABS()
			cpu.cycleEnd()
		case LDA_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			LDA_ABS()
			cpu.cycleEnd()
		case LDX_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			LDX_ABS()
			cpu.cycleEnd()
		case LDY_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			LDY_ABS()
			cpu.cycleEnd()
		case LSR_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			LSR_ABS()
			cpu.cycleEnd()
		case ORA_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			ORA_ABS()
			cpu.cycleEnd()
		case ROL_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			ROL_ABS()
			cpu.cycleEnd()
		case ROR_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			ROR_ABS()
			cpu.cycleEnd()
		case SBC_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			SBC_ABS()
			cpu.cycleEnd()
		case STA_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			STA_ABS()
			cpu.cycleEnd()
		case STX_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			STX_ABS()
			cpu.cycleEnd()
		case STY_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			STY_ABS()
			cpu.cycleEnd()

		// X Indexed Absolute addressing mode instructions
		/*
			$nnnn,X

			This form of addressing is used in conjunction with the X index register. The effective address is formed by adding the contents of X to the address contained in the second and third bytes of the instruction. This mode allows the index register to contain the index or count value and the instruction to contain the base address. This type of indexing allows any location referencing and the index to modify multiple fields resulting in reduced coding and execution time.

			Note on the MOS 6502:

			The value at the specified address, ignoring the the addressing mode's X offset, is read (and discarded) before the final address is read. This may cause side effects in I/O registers.


			Bytes: 3
		*/
		case ADC_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			ADC_ABX()
			cpu.cycleEnd()
		case AND_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			AND_ABX()
			cpu.cycleEnd()
		case ASL_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			ASL_ABX()
			cpu.cycleEnd()
		case CMP_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			CMP_ABX()
			cpu.cycleEnd()
		case DEC_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			DEC_ABX()
			cpu.cycleEnd()
		case EOR_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			EOR_ABX()
			cpu.cycleEnd()
		case INC_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			INC_ABX()
			cpu.cycleEnd()
		case LDA_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			LDA_ABX()
			cpu.cycleEnd()
		case LDY_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			LDY_ABX()
			cpu.cycleEnd()
		case LSR_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			LSR_ABX()
			cpu.cycleEnd()
		case ORA_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			ORA_ABX()
			cpu.cycleEnd()
		case ROL_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			ROL_ABX()
			cpu.cycleEnd()
		case ROR_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			ROR_ABX()
			cpu.cycleEnd()
		case SBC_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			SBC_ABX()
			cpu.cycleEnd()
		case STA_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			STA_ABX()
			cpu.cycleEnd()

		// Y Indexed Absolute addressing mode instructions
		/*
			$nnnn,Y

			This form of addressing is used in conjunction with the Y index register. The effective address is formed by adding the contents of Y to the address contained in the second and third bytes of the instruction. This mode allows the index register to contain the index or count value and the instruction to contain the base address. This type of indexing allows any location referencing and the index to modify multiple fields resulting in reduced coding and execution time.

			Note on the MOS 6502:

			The value at the specified address, ignoring the the addressing mode's Y offset, is read (and discarded) before the final address is read. This may cause side effects in I/O registers.

			Bytes: 3
		*/
		case ADC_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			ADC_ABY()
			cpu.cycleEnd()
		case AND_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			AND_ABY()
			cpu.cycleEnd()
		case CMP_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			CMP_ABY()
			cpu.cycleEnd()
		case EOR_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			EOR_ABY()
			cpu.cycleEnd()
		case LDA_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			LDA_ABY()
			cpu.cycleEnd()
		case LDX_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			LDX_ABY()
			cpu.cycleEnd()
		case ORA_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			ORA_ABY()
			cpu.cycleEnd()
		case SBC_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			SBC_ABY()
			cpu.cycleEnd()
		case STA_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			STA_ABY()
			cpu.cycleEnd()

		// Absolute Indirect addressing mode instructions
		case JMP_INDIRECT_OPCODE:
			cpu.cycleStart()
			JMP_IND()
			cpu.cycleEnd()
		}
		if *plus4 {
			plus4KernalRoutines()
		}
	}
}

func (cpu *CPU) getSRBit(x byte) byte {
	return (cpu.SR >> x) & 1
}
func (cpu *CPU) setSRBitOn(x byte) {
	cpu.SR |= 1 << x
}
func (cpu *CPU) setSRBitOff(x byte) {
	cpu.SR &= ^(1 << x)
}
func (cpu *CPU) getABit(x byte) byte {
	return (cpu.A >> x) & 1
}
func (cpu *CPU) getXBit(x byte) byte {
	return (cpu.X >> x) & 1
}
func (cpu *CPU) getYBit(x byte) byte {
	return (cpu.Y >> x) & 1
}

func (cpu *CPU) setNegativeFlag() {
	cpu.setSRBitOn(7)
}
func (cpu *CPU) unsetNegativeFlag() {
	cpu.setSRBitOff(7)
}
func (cpu *CPU) setOverflowFlag() {
	cpu.setSRBitOn(6)
}
func (cpu *CPU) unsetOverflowFlag() {
	cpu.setSRBitOff(6)
}
func (cpu *CPU) setBreakFlag() {
	cpu.setSRBitOn(4)
}
func (cpu *CPU) setDecimalFlag() {
	cpu.setSRBitOn(3)
}
func (cpu *CPU) unsetDecimalFlag() {
	cpu.setSRBitOff(3)
}
func (cpu *CPU) setInterruptFlag() {
	cpu.setSRBitOn(2)
}
func (cpu *CPU) unsetInterruptFlag() {
	cpu.setSRBitOff(2)
}
func (cpu *CPU) setZeroFlag() {
	cpu.setSRBitOn(1)
}
func (cpu *CPU) unsetZeroFlag() {
	cpu.setSRBitOff(1)
}
func (cpu *CPU) setCarryFlag() {
	cpu.setSRBitOn(0)
}
func (cpu *CPU) unsetCarryFlag() {
	cpu.setSRBitOff(0)
}
