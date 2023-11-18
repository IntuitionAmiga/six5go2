package main

import (
	"fmt"
	"os"
	"sync"
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
)

const (
	// Choose video standard: 35200 for PAL, 29833 for NTSC
	VBLANK_CYCLES = 35200 // For PAL
	// const VBLANK_CYCLES = 29833 // For NTSC

	//cpuSpeedHz uint64 = 1790000 // 1790000 Hz for a NTSC 7501/8501
	//cpuSpeedHz uint64 = 1760000 // 1760000 Hz for a PAL 7501/8501
	//cpuSpeedHz uint64 = 7500 // Run it slow whilst debugging!
	cpuSpeedHz uint64 = 4000000000000 // 4 GHz to speed up debugging! :)
)

type CPU struct {
	A             byte   // Accumulator
	X             byte   // X register
	Y             byte   // Y register
	PC            uint16 // Program Counter
	SP            uint16 // Stack Pointer
	SR            byte   // Status Register
	preOpPC       uint16
	preOpOpcode   byte
	preOpOperand1 byte
	preOpOperand2 byte
	preOpSP       uint16
	preOpA        byte
	preOpX        byte
	preOpY        byte
	preOpSR       byte

	cycleCounter            uint64
	cycleStartTime          time.Time     // High-resolution timer
	cpuTimeSpent            time.Duration // Time spent executing instructions
	vblankCycleCounter      uint64
	traceLine               string
	nextTraceLine           string
	irq                     bool
	nmi                     bool
	reset                   bool
	traceMutex              sync.Mutex
	instructionCounter      uint32
	disassembledInstruction string

	cpuQuit bool
}

var (
	cpu              CPU
	cycleTime        = time.Second / time.Duration(cpuSpeedHz) // time per cycle in nanoseconds
	breakpointsMutex sync.Mutex
	breakpoints      = make(map[uint16]bool)
	breakpointHit    = make(chan uint16, 1)
	breakpointHalted = false
)

func (cpu *CPU) opcode() byte {
	return cpu.readMemory(cpu.PC)
}
func (cpu *CPU) operand1() byte {
	return cpu.readMemory(cpu.PC + 1)
}
func (cpu *CPU) operand2() byte {
	return cpu.readMemory(cpu.PC + 2)
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

func (cpu *CPU) incPC(amount int) {
	cpu.PC += uint16(amount)
	if cpu.PC > 0xFFFF {
		cpu.PC = 0x0000 + (cpu.PC & 0xFFFF)
	}
}
func (cpu *CPU) decPC(amount int) {
	cpu.PC -= uint16(amount)
	if cpu.PC < 0 {
		cpu.PC = 0xFFFF + (cpu.PC & 0xFFFF)
	}
}
func (cpu *CPU) setPC(newAddress uint16) {
	cpu.PC = uint16(newAddress) & 0xFFFF
}

func (cpu *CPU) handleIRQ() {
	//fmt.Fprintf(os.Stderr, "Debug: Entering handleIRQ() at PC: $%04X, cycleCounter is %d\n", cpu.PC, cpu.cycleCounter)
	if cpu.getSRBit(2) == 1 {
		//fmt.println("Debug: Interrupt disabled. Exiting handleIRQ()")
		return
	}
	//fmt.println("Debug: Interrupt enabled. Continuing...")
	// Push PC onto stack
	cpu.updateStack(byte(cpu.PC >> 8)) // high byte
	cpu.decSP()
	cpu.updateStack(byte(cpu.PC & 0xFF)) // low byte
	cpu.decSP()
	//fmt.Printf("Debug: PC pushed to stack. SP: %X\n", cpu.SP)
	// Push SR onto stack
	cpu.updateStack(cpu.SR)
	cpu.decSP()
	//fmt.Printf("Debug: PC pushed to stack. SP: %X\n", cpu.SP)
	// Set interrupt flag
	cpu.setInterruptFlag()
	//fmt.Printf("Debug: PC pushed to stack. SP: %X\n", cpu.SP)
	// Set PC to IRQ Service Routine address
	lowByte := cpu.readMemory(IRQVectorAddressLow)
	highByte := cpu.readMemory(IRQVectorAddressHigh)
	cpu.setPC(uint16(lowByte) | uint16(highByte)<<8)
	//fmt.Printf("Debug: Jumping to IRQ Service Routine at %X\n", cpu.PC)
	cpu.irq = false
	//fmt.println("Debug: Exiting handleIRQ()")
}
func (cpu *CPU) handleNMI() {
	// Push PC onto stack
	cpu.updateStack(byte(cpu.PC >> 8)) // high byte
	cpu.decSP()
	cpu.updateStack(byte(cpu.PC & 0xFF)) // low byte
	cpu.decSP()
	// Push SR onto stack
	cpu.updateStack(cpu.SR)
	cpu.decSP()
	// Set PC to NMI Service Routine address
	lowByte := cpu.readMemory(NMIVectorAddressLow)
	highByte := cpu.readMemory(NMIVectorAddressHigh)
	cpu.setPC(uint16(lowByte) | uint16(highByte)<<8)
	cpu.nmi = false // Clear the NMI flag
}
func (cpu *CPU) handleRESET() {
	cpu.resetCPU()
	cpu.reset = false // Clear the RESET flag
}
func (cpu *CPU) handleState(amount int) {
	cpu.incPC(amount)
	// If amount is 0, then we are in a branch instruction and we don't want to increment the instruction counter
	if amount != 0 {
		cpu.instructionCounter++
	}
	plus4SystemInterrupts()
	if cpu.irq && cpu.getSRBit(2) == 0 {
		cpu.handleIRQ()
	}
	if cpu.nmi {
		cpu.handleNMI()
	}
	if cpu.reset {
		cpu.handleRESET()
	}
}

func plus4SystemInterrupts() {
	// Check for VBLANK
	if cpu.vblankCycleCounter >= VBLANK_CYCLES {
		//fmt.Printf("VBLANK triggered. Resetting vblankCycleCounter.\n")
		cpu.irq = true             // Trigger IRQ
		cpu.vblankCycleCounter = 0 // Reset cycle counter
	}
}

func (cpu *CPU) updateCycleCounter(amount uint64) {
	cpu.cycleCounter += amount
	cpu.vblankCycleCounter += amount
	//fmt.Fprintf(os.Stderr, "Cycle count: %d, VBLANK cycle count: %d\n", cpu.cycleCounter, cpu.vblankCycleCounter)
}
func (cpu *CPU) incrementCycleCountForBranch(oldPC uint16) {
	// Add 1 more cycle if branch is taken
	cpu.updateCycleCounter(1)
	// Check for page boundary crossing and add another cycle if necessary
	if (oldPC & 0xFF00) != (cpu.PC & 0xFF00) {
		cpu.updateCycleCounter(1)
	}
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
	cpu.SR = 0b01100010

	if *klausd {
		cpu.setPC(0x400)
	} else {
		// Set PC to value stored at reset vector address
		lowByte := cpu.readMemory(RESETVectorAddressLow)
		highByte := cpu.readMemory(RESETVectorAddressHigh)
		cpu.setPC(uint16(lowByte) | uint16(highByte)<<8)
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

func (cpu *CPU) startCPU() {
	fmt.Fprintf(os.Stderr, "Before any instruction is executed, readMemory(0x8009): %02X\n", cpu.readMemory(0x8009))

	for uint(cpu.PC) < 0xFFFF {
		cpu.preOpPC = cpu.PC
		cpu.preOpOpcode = cpu.opcode()
		cpu.preOpOperand1 = cpu.operand1()
		cpu.preOpOperand2 = cpu.operand2()
		cpu.preOpSP = cpu.SP
		cpu.preOpA = cpu.A
		cpu.preOpX = cpu.X
		cpu.preOpY = cpu.Y
		cpu.preOpSR = cpu.SR

		breakpointsMutex.Lock()
		_, exists := breakpoints[cpu.preOpPC]
		breakpointsMutex.Unlock()
		if exists {
			breakpointHalted = true      // Indicate that we're halted at a breakpoint
			breakpointHit <- cpu.preOpPC // Signal that a breakpoint has been hit
			<-breakpointHit              // Wait here until we receive a signal to continue
			breakpointHalted = false     // Reset the flag as we're continuing
		}

		//  1 byte instructions with no operands
		switch cpu.opcode() {
		// Implied addressing mode instructions
		/*
			In the implied addressing mode, the address containing the operand is implicitly stated in the operation code of the instruction.

			Bytes: 1
		*/
		case BRK_OPCODE:
			cpu.cycleStart()
			cpu.BRK()
			cpu.cycleEnd()
		case CLC_OPCODE:
			cpu.cycleStart()
			cpu.CLC()
			cpu.cycleEnd()
		case CLD_OPCODE:
			cpu.cycleStart()
			cpu.CLD()
			cpu.cycleEnd()
		case CLI_OPCODE:
			cpu.cycleStart()
			cpu.CLI()
			cpu.cycleEnd()
		case CLV_OPCODE:
			cpu.cycleStart()
			cpu.CLV()
			cpu.cycleEnd()
		case DEX_OPCODE:
			cpu.cycleStart()
			cpu.DEX()
			cpu.cycleEnd()
		case DEY_OPCODE:
			cpu.cycleStart()
			cpu.DEY()
			cpu.cycleEnd()
		case INX_OPCODE:
			cpu.cycleStart()
			cpu.INX()
			cpu.cycleEnd()
		case INY_OPCODE:
			cpu.cycleStart()
			cpu.INY()
			cpu.cycleEnd()
		case NOP_OPCODE:
			cpu.cycleStart()
			cpu.NOP()
			cpu.cycleEnd()
		case PHA_OPCODE:
			cpu.cycleStart()
			cpu.PHA()
			cpu.cycleEnd()
		case PHP_OPCODE:
			cpu.cycleStart()
			cpu.PHP()
			cpu.cycleEnd()
		case PLA_OPCODE:
			cpu.cycleStart()
			cpu.PLA()
			cpu.cycleEnd()
		case PLP_OPCODE:
			cpu.cycleStart()
			cpu.PLP()
			cpu.cycleEnd()
		case RTI_OPCODE:
			cpu.cycleStart()
			cpu.RTI()
			cpu.cycleEnd()
		case RTS_OPCODE:
			cpu.cycleStart()
			cpu.RTS()
			cpu.cycleEnd()
		case SEC_OPCODE:
			cpu.cycleStart()
			cpu.SEC()
			cpu.cycleEnd()
		case SED_OPCODE:
			cpu.cycleStart()
			cpu.SED()
			cpu.cycleEnd()
		case SEI_OPCODE:
			cpu.cycleStart()
			cpu.SEI()
			cpu.cycleEnd()
		case TAX_OPCODE:
			cpu.cycleStart()
			cpu.TAX()
			cpu.cycleEnd()
		case TAY_OPCODE:
			cpu.cycleStart()
			cpu.TAY()
			cpu.cycleEnd()
		case TSX_OPCODE:
			cpu.cycleStart()
			cpu.TSX()
			cpu.cycleEnd()
		case TXA_OPCODE:
			cpu.cycleStart()
			cpu.TXA()
			cpu.cycleEnd()
		case TXS_OPCODE:
			cpu.cycleStart()
			cpu.TXS()
			cpu.cycleEnd()
		case TYA_OPCODE:
			cpu.cycleStart()
			cpu.TYA()
			cpu.cycleEnd()

		// Accumulator instructions
		/*
			A

			This form of addressing is represented with a one byte instruction, implying an operation on the accumulator.

			Bytes: 1
		*/
		case ASL_ACCUMULATOR_OPCODE:
			cpu.cycleStart()
			cpu.ASL_A()
			cpu.cycleEnd()
		case LSR_ACCUMULATOR_OPCODE:
			cpu.cycleStart()
			cpu.LSR_A()
			cpu.cycleEnd()
		case ROL_ACCUMULATOR_OPCODE:
			cpu.cycleStart()
			cpu.ROL_A()
			cpu.cycleEnd()
		case ROR_ACCUMULATOR_OPCODE:
			cpu.cycleStart()
			cpu.ROR_A()
			cpu.cycleEnd()
		//}

		// 2 byte instructions with 1 operand
		//switch cpu.opcode() {
		// Immediate addressing mode instructions
		/*
			#$nn

			In immediate addressing, the operand is contained in the second byte of the instruction, with no further memory addressing required.

			Bytes: 2
		*/
		case ADC_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			cpu.ADC_I()
			cpu.cycleEnd()
		case AND_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			cpu.AND_I()
			cpu.cycleEnd()
		case CMP_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			cpu.CMP_I()
			cpu.cycleEnd()
		case CPX_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			cpu.CPX_I()
			cpu.cycleEnd()
		case CPY_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			cpu.CPY_I()
			cpu.cycleEnd()
		case EOR_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			cpu.EOR_I()
			cpu.cycleEnd()
		case LDA_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			cpu.LDA_I()
			cpu.cycleEnd()
		case LDX_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			cpu.LDX_I()
			cpu.cycleEnd()
		case LDY_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			cpu.LDY_I()
			cpu.cycleEnd()
		case ORA_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			cpu.ORA_I()
			cpu.cycleEnd()
		case SBC_IMMEDIATE_OPCODE:
			cpu.cycleStart()
			cpu.SBC_I()
			cpu.cycleEnd()

		// Zero Page addressing mode instructions
		/*
			$nn

			The zero page instructions allow for shorter code and execution times by only fetching the second byte of the instruction and assuming a zero low address byte. Careful use of the zero page can result in significant increase in code efficiency.

			Bytes: 2
		*/
		case ADC_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.ADC_Z()
			cpu.cycleEnd()
		case AND_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.AND_Z()
			cpu.cycleEnd()
		case ASL_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.ASL_Z()
			cpu.cycleEnd()
		case BIT_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.BIT_Z()
			cpu.cycleEnd()
		case CMP_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.CMP_Z()
			cpu.cycleEnd()
		case CPX_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.CPX_Z()
			cpu.cycleEnd()
		case CPY_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.CPY_Z()
			cpu.cycleEnd()
		case DEC_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.DEC_Z()
			cpu.cycleEnd()
		case EOR_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.EOR_Z()
			cpu.cycleEnd()
		case INC_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.INC_Z()
			cpu.cycleEnd()
		case LDA_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.LDA_Z()
			cpu.cycleEnd()
		case LDX_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.LDX_Z()
			cpu.cycleEnd()
		case LDY_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.LDY_Z()
			cpu.cycleEnd()
		case LSR_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.LSR_Z()
			cpu.cycleEnd()
		case ORA_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.ORA_Z()
			cpu.cycleEnd()
		case ROL_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.ROL_Z()
			cpu.cycleEnd()
		case ROR_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.ROR_Z()
			cpu.cycleEnd()
		case SBC_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.SBC_Z()
			cpu.cycleEnd()
		case STA_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.STA_Z()
			cpu.cycleEnd()
		case STX_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.STX_Z()
			cpu.cycleEnd()
		case STY_ZERO_PAGE_OPCODE:
			cpu.cycleStart()
			cpu.STY_Z()
			cpu.cycleEnd()

		// X Indexed Zero Page addressing mode instructions
		/*
			$nn,X

			This form of addressing is used in conjunction with the X index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the low order 8 bits of memory and crossing of page boundaries does not occur.

			Bytes: 2
		*/
		case ADC_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			cpu.ADC_ZX()
			cpu.cycleEnd()
		case AND_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			cpu.AND_ZX()
			cpu.cycleEnd()
		case ASL_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			cpu.ASL_ZX()
			cpu.cycleEnd()
		case CMP_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			cpu.CMP_ZX()
			cpu.cycleEnd()
		case DEC_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			cpu.DEC_ZX()
			cpu.cycleEnd()
		case LDA_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			cpu.LDA_ZX()
			cpu.cycleEnd()
		case LDY_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			cpu.LDY_ZX()
			cpu.cycleEnd()
		case LSR_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			cpu.LSR_ZX()
			cpu.cycleEnd()
		case ORA_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			cpu.ORA_ZX()
			cpu.cycleEnd()
		case ROL_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			cpu.ROL_ZX()
			cpu.cycleEnd()
		case ROR_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			cpu.ROR_ZX()
			cpu.cycleEnd()
		case EOR_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			cpu.EOR_ZX()
			cpu.cycleEnd()
		case INC_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			cpu.INC_ZX()
			cpu.cycleEnd()
		case SBC_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			cpu.SBC_ZX()
			cpu.cycleEnd()
		case STA_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			cpu.STA_ZX()
			cpu.cycleEnd()
		case STY_ZERO_PAGE_X_OPCODE:
			cpu.cycleStart()
			cpu.STY_ZX()
			cpu.cycleEnd()

		// Y Indexed Zero Page addressing mode instructions
		/*
			$nn,Y

			This form of addressing is used in conjunction with the Y index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the low order 8 bits of memory and crossing of page boundaries does not occur.

			Bytes: 2
		*/
		case LDX_ZERO_PAGE_Y_OPCODE:
			cpu.cycleStart()
			cpu.LDX_ZY()
			cpu.cycleEnd()
		case STX_ZERO_PAGE_Y_OPCODE:
			cpu.cycleStart()
			cpu.STX_ZY()
			cpu.cycleEnd()

		// X Indexed Zero Page Indirect addressing mode instructions
		/*
			($nn,X)

			In indexed indirect addressing, the second byte of the instruction is added to the contents of the X index register, discarding the carry. The result of this addition points to a memory location on page zero whose contents is the high order eight bits of the effective address. The next memory location in page zero contains the low order eight bits of the effective address. Both memory locations specifying the low and high order bytes of the effective address must be in page zero.

			Bytes: 2
		*/
		case ADC_INDIRECT_X_OPCODE:
			cpu.cycleStart()
			cpu.ADC_IX()
			cpu.cycleEnd()
		case AND_INDIRECT_X_OPCODE:
			cpu.cycleStart()
			cpu.AND_IX()
			cpu.cycleEnd()
		case CMP_INDIRECT_X_OPCODE:
			cpu.cycleStart()
			cpu.CMP_IX()
			cpu.cycleEnd()
		case EOR_INDIRECT_X_OPCODE:
			cpu.cycleStart()
			cpu.EOR_IX()
			cpu.cycleEnd()
		case LDA_INDIRECT_X_OPCODE:
			cpu.cycleStart()
			cpu.LDA_IX()
			cpu.cycleEnd()
		case ORA_INDIRECT_X_OPCODE:
			cpu.cycleStart()
			cpu.ORA_IX()
			cpu.cycleEnd()
		case SBC_INDIRECT_X_OPCODE:
			cpu.cycleStart()
			cpu.SBC_IX()
			cpu.cycleEnd()
		case STA_INDIRECT_X_OPCODE:
			cpu.cycleStart()
			cpu.STA_IX()
			cpu.cycleEnd()

		// Zero Page Indirect Y Indexed addressing mode instructions
		/*
			($nn),Y

			In indirect indexed addressing, the second byte of the instruction points to a memory location in page zero. The contents of this memory location is added to the contents of the Y index register, the result being the high order eight bits of the effective address. The carry from this addition is added to the contents of the next page zero memory location, the result being the low order eight bits of the effective address.

			Bytes: 2
		*/
		case ADC_INDIRECT_Y_OPCODE:
			cpu.cycleStart()
			cpu.ADC_IY()
			cpu.cycleEnd()
		case AND_INDIRECT_Y_OPCODE:
			cpu.cycleStart()
			cpu.AND_IY()
			cpu.cycleEnd()
		case CMP_INDIRECT_Y_OPCODE:
			cpu.cycleStart()
			cpu.CMP_IY()
			cpu.cycleEnd()
		case EOR_INDIRECT_Y_OPCODE:
			cpu.cycleStart()
			cpu.EOR_IY()
			cpu.cycleEnd()
		case LDA_INDIRECT_Y_OPCODE:
			cpu.cycleStart()
			cpu.LDA_IY()
			cpu.cycleEnd()
		case ORA_INDIRECT_Y_OPCODE:
			cpu.cycleStart()
			cpu.ORA_IY()
			cpu.cycleEnd()
		case SBC_INDIRECT_Y_OPCODE:
			cpu.cycleStart()
			cpu.SBC_IY()
			cpu.cycleEnd()
		case STA_INDIRECT_Y_OPCODE:
			cpu.cycleStart()
			cpu.STA_IY()
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
			cpu.BPL_R()
			cpu.cycleEnd()
		case BMI_RELATIVE_OPCODE:
			cpu.cycleStart()
			cpu.BMI_R()
			cpu.cycleEnd()
		case BVC_RELATIVE_OPCODE:
			cpu.cycleStart()
			cpu.BVC_R()
			cpu.cycleEnd()
		case BVS_RELATIVE_OPCODE:
			cpu.cycleStart()
			cpu.BVS_R()
			cpu.cycleEnd()
		case BCC_RELATIVE_OPCODE:
			cpu.cycleStart()
			cpu.BCC_R()
			cpu.cycleEnd()
		case BCS_RELATIVE_OPCODE:
			cpu.cycleStart()
			cpu.BCS_R()
			cpu.cycleEnd()
		case BNE_RELATIVE_OPCODE:
			cpu.cycleStart()
			cpu.BNE_R()
			cpu.cycleEnd()
		case BEQ_RELATIVE_OPCODE:
			cpu.cycleStart()
			cpu.BEQ_R()
			cpu.cycleEnd()
		//}

		// 3 byte instructions with 2 operands
		//switch cpu.opcode() {
		// Absolute addressing mode instructions
		/*
			$nnnn

			In absolute addressing, the second byte of the instruction specifies the eight high order bits of the effective address while the third byte specifies the eight low order bits. Thus, the absolute addressing mode allows access to the entire 65 K bytes of addressable memory.

			Bytes: 3
		*/
		case ADC_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.ADC_ABS()
			cpu.cycleEnd()
		case AND_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.AND_ABS()
			cpu.cycleEnd()
		case ASL_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.ASL_ABS()
			cpu.cycleEnd()
		case BIT_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.BIT_ABS()
			cpu.cycleEnd()
		case CMP_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.CMP_ABS()
			cpu.cycleEnd()
		case CPX_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.CPX_ABS()
			cpu.cycleEnd()
		case CPY_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.CPY_ABS()
			cpu.cycleEnd()
		case DEC_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.DEC_ABS()
			cpu.cycleEnd()
		case EOR_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.EOR_ABS()
			cpu.cycleEnd()
		case INC_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.INC_ABS()
			cpu.cycleEnd()
		case JMP_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.JMP_ABS()
			cpu.cycleEnd()
		case JSR_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.JSR_ABS()
			cpu.cycleEnd()
		case LDA_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.LDA_ABS()
			cpu.cycleEnd()
		case LDX_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.LDX_ABS()
			cpu.cycleEnd()
		case LDY_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.LDY_ABS()
			cpu.cycleEnd()
		case LSR_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.LSR_ABS()
			cpu.cycleEnd()
		case ORA_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.ORA_ABS()
			cpu.cycleEnd()
		case ROL_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.ROL_ABS()
			cpu.cycleEnd()
		case ROR_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.ROR_ABS()
			cpu.cycleEnd()
		case SBC_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.SBC_ABS()
			cpu.cycleEnd()
		case STA_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.STA_ABS()
			cpu.cycleEnd()
		case STX_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.STX_ABS()
			cpu.cycleEnd()
		case STY_ABSOLUTE_OPCODE:
			cpu.cycleStart()
			cpu.STY_ABS()
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
			cpu.ADC_ABX()
			cpu.cycleEnd()
		case AND_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			cpu.AND_ABX()
			cpu.cycleEnd()
		case ASL_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			cpu.ASL_ABX()
			cpu.cycleEnd()
		case CMP_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			cpu.CMP_ABX()
			cpu.cycleEnd()
		case DEC_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			cpu.DEC_ABX()
			cpu.cycleEnd()
		case EOR_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			cpu.EOR_ABX()
			cpu.cycleEnd()
		case INC_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			cpu.INC_ABX()
			cpu.cycleEnd()
		case LDA_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			cpu.LDA_ABX()
			cpu.cycleEnd()
		case LDY_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			cpu.LDY_ABX()
			cpu.cycleEnd()
		case LSR_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			cpu.LSR_ABX()
			cpu.cycleEnd()
		case ORA_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			cpu.ORA_ABX()
			cpu.cycleEnd()
		case ROL_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			cpu.ROL_ABX()
			cpu.cycleEnd()
		case ROR_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			cpu.ROR_ABX()
			cpu.cycleEnd()
		case SBC_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			cpu.SBC_ABX()
			cpu.cycleEnd()
		case STA_ABSOLUTE_X_OPCODE:
			cpu.cycleStart()
			cpu.STA_ABX()
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
			cpu.ADC_ABY()
			cpu.cycleEnd()
		case AND_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			cpu.AND_ABY()
			cpu.cycleEnd()
		case CMP_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			cpu.CMP_ABY()
			cpu.cycleEnd()
		case EOR_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			cpu.EOR_ABY()
			cpu.cycleEnd()
		case LDA_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			cpu.LDA_ABY()
			cpu.cycleEnd()
		case LDX_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			cpu.LDX_ABY()
			cpu.cycleEnd()
		case ORA_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			cpu.ORA_ABY()
			cpu.cycleEnd()
		case SBC_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			cpu.SBC_ABY()
			cpu.cycleEnd()
		case STA_ABSOLUTE_Y_OPCODE:
			cpu.cycleStart()
			cpu.STA_ABY()
			cpu.cycleEnd()

		// Absolute Indirect addressing mode instructions
		case JMP_INDIRECT_OPCODE:
			cpu.cycleStart()
			cpu.JMP_IND()
			cpu.cycleEnd()
		}
		if *plus4 {
			plus4KernalRoutines()
			ted.Timer1Counter++
			ted.Timer2Counter++
			ted.Timer3Counter++
		}
		//fmt.Fprintf(os.Stderr, "Debug: At PC $%04X/preOpPC $%04X memory[0x0314] is %04X\n", cpu.PC, cpu.preOpPC, memory[0x0314])

		executionTrace()
		if cpu.cpuQuit {
			break
		}
		fmt.Fprintf(os.Stderr, "After instruction %s, readMemory(0x8009): %02X\n", cpu.disassembledInstruction, cpu.readMemory(0x8009))
		// For AllSuiteA.bin 6502 opcode test suite
		if *allsuitea && cpu.readMemory(0x210) == 0xFF {
			fmt.Printf("\n\u001B[32;5mMemory address $210 == $%02X. All opcodes succesfully tested and passed!\u001B[0m\n", cpu.readMemory(0x210))
			os.Exit(0)
		}
	}
}
