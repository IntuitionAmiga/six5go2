package main

import (
	"fmt"
	"os"
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

	SPBaseAddress      uint16 = 0x0100
	NMIVectorAddress          = 0xFFFA
	RESETVectorAddress        = 0xFFFC
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
	BRKtrue          bool          = false
	IRQVectorAddress uint16        = 0xFFFE
	cycleCounter     uint64        = 0
	cpuSpeedHz       uint64        = 985248                                  // 985248 Hz for a standard 6502
	cycleTime        time.Duration = time.Second / time.Duration(cpuSpeedHz) // time per cycle in nanoseconds
	cycleStartTime   time.Time                                               // High-resolution timer
	timeSpent        time.Duration                                           // Time spent executing instructions

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
		case 0x00: //BRK
			cycleStart()
			BRK()
			cycleEnd()
		case 0x18: //CLC
			cycleStart()
			CLC()
			cycleEnd()
		case 0xD8: //CLD
			cycleStart()
			CLD()
			cycleEnd()
		case 0x58: //CLI
			cycleStart()
			CLI()
			cycleEnd()
		case 0xB8: //CLV
			cycleStart()
			CLV()
			cycleEnd()
		case 0xCA: //DEX
			cycleStart()
			DEX()
			cycleEnd()
		case 0x88: //DEY
			cycleStart()
			DEY()
			cycleEnd()
		case 0xE8: //INX
			cycleStart()
			INX()
			cycleEnd()
		case 0xC8:
			cycleStart()
			INY()
			cycleEnd()
		case 0xEA: //NOP
			cycleStart()
			NOP()
			cycleEnd()
		case 0x48: //PHA
			cycleStart()
			PHA()
			cycleEnd()
		case 0x08: //PHP
			cycleStart()
			PHP()
			cycleEnd()
		case 0x68: //PLA
			cycleStart()
			PLA()
			cycleEnd()
		case 0x28: //PLP
			cycleStart()
			PLP()
			cycleEnd()
		case 0x40: //RTI
			cycleStart()
			RTI()
			cycleEnd()
		case 0x60: //RTS
			cycleStart()
			RTS()
			cycleEnd()
		case 0x38: //SEC
			cycleStart()
			SEC()
			cycleEnd()
		case 0xF8: //SED
			cycleStart()
			SED()
			cycleEnd()
		case 0x78: //SEI
			cycleStart()
			SEI()
			cycleEnd()
		case 0xAA: //TAX
			cycleStart()
			TAX()
			cycleEnd()
		case 0xA8: //TAY
			cycleStart()
			TAY()
			cycleEnd()
		case 0xBA: //TSX
			cycleStart()
			TSX()
			cycleEnd()
		case 0x8A: //TXA
			cycleStart()
			TXA()
			cycleEnd()
		case 0x9A: //TXS
			cycleStart()
			TXS()
			cycleEnd()
		case 0x98: //TYA
			cycleStart()
			TYA()
			cycleEnd()

		// Accumulator instructions
		/*
			A

			This form of addressing is represented with a one byte instruction, implying an operation on the accumulator.

			Bytes: 1
		*/
		case 0x0A: //ASL
			cycleStart()
			ASL_A()
			cycleEnd()
		case 0x4A: //LSR
			cycleStart()
			LSR_A()
			cycleEnd()
		case 0x2A: //ROL
			cycleStart()
			ROL_A()
			cycleEnd()
		case 0x6A: //ROR
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
		case 0x69: //ADC
			cycleStart()
			ADC_I()
			cycleEnd()
		case 0x29: //AND
			cycleStart()
			AND_I()
			cycleEnd()
		case 0xC9: //CMP
			cycleStart()
			CMP_I()
			cycleEnd()
		case 0xE0: //CPX
			cycleStart()
			CPX_I()
			cycleEnd()
		case 0xC0: //CPY
			cycleStart()
			CPY_I()
			cycleEnd()
		case 0x49: //EOR
			cycleStart()
			EOR_I()
			cycleEnd()
		case 0xA9: //LDA
			cycleStart()
			LDA_I()
			cycleEnd()
		case 0xA2: //LDX
			cycleStart()
			LDX_I()
			cycleEnd()
		case 0xA0: //LDY
			cycleStart()
			LDY_I()
			cycleEnd()
		case 0x09: //ORA
			cycleStart()
			ORA_I()
			cycleEnd()
		case 0xE9: //SBC
			cycleStart()
			SBC_I()
			cycleEnd()

		// Zero Page addressing mode instructions
		/*
			$nn

			The zero page instructions allow for shorter code and execution times by only fetching the second byte of the instruction and assuming a zero low address byte. Careful use of the zero page can result in significant increase in code efficiency.

			Bytes: 2
		*/
		case 0x65: //ADC
			cycleStart()
			ADC_Z()
			cycleEnd()
		case 0x25: //AND
			cycleStart()
			AND_Z()
			cycleEnd()
		case 0x06: //ASL
			cycleStart()
			ASL_Z()
			cycleEnd()
		case 0x24: //BIT
			cycleStart()
			BIT_Z()
			cycleEnd()
		case 0xC5: //CMP
			cycleStart()
			CMP_Z()
			cycleEnd()
		case 0xE4: //CPX
			cycleStart()
			CPX_Z()
			cycleEnd()
		case 0xC4: //CPY
			cycleStart()
			CPY_Z()
			cycleEnd()
		case 0xC6: //DEC
			cycleStart()
			DEC_Z()
			cycleEnd()
		case 0x45: //EOR
			cycleStart()
			EOR_Z()
			cycleEnd()
		case 0xE6: //INC
			cycleStart()
			INC_Z()
			cycleEnd()
		case 0xA5: //LDA
			cycleStart()
			LDA_Z()
			cycleEnd()
		case 0xA6: //LDX
			cycleStart()
			LDX_Z()
			cycleEnd()
		case 0xA4: //LDY
			cycleStart()
			LDY_Z()
			cycleEnd()
		case 0x46: //LSR
			cycleStart()
			LSR_Z()
			cycleEnd()
		case 0x05: //ORA
			cycleStart()
			ORA_Z()
			cycleEnd()
		case 0x26: //ROL
			cycleStart()
			ROL_Z()
			cycleEnd()
		case 0x66: //ROR
			cycleStart()
			ROR_Z()
			cycleEnd()
		case 0xE5: //SBC
			cycleStart()
			SBC_Z()
			cycleEnd()
		case 0x85: //STA
			cycleStart()
			STA_Z()
			cycleEnd()
		case 0x86: //STX
			cycleStart()
			STX_Z()
			cycleEnd()
		case 0x84: //STY
			cycleStart()
			STY_Z()
			cycleEnd()

		// X Indexed Zero Page addressing mode instructions
		/*
			$nn,X

			This form of addressing is used in conjunction with the X index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the low order 8 bits of memory and crossing of page boundaries does not occur.

			Bytes: 2
		*/
		case 0x75: //ADC
			cycleStart()
			ADC_ZX()
			cycleEnd()
		case 0x35: //AND
			cycleStart()
			AND_ZX()
			cycleEnd()
		case 0x16: //ASL
			cycleStart()
			ASL_ZX()
			cycleEnd()
		case 0xD5: //CMP
			cycleStart()
			CMP_ZX()
			cycleEnd()
		case 0xD6: //DEC
			cycleStart()
			DEC_ZX()
			cycleEnd()
		case 0xB5: //LDA
			cycleStart()
			LDA_ZX()
			cycleEnd()
		case 0xB4: //LDY
			cycleStart()
			LDY_ZX()
			cycleEnd()
		case 0x56: //LSR
			cycleStart()
			LSR_ZX()
			cycleEnd()
		case 0x15: //ORA
			cycleStart()
			ORA_ZX()
			cycleEnd()
		case 0x36: //ROL
			cycleStart()
			ROL_ZX()
			cycleEnd()
		case 0x76: //ROR
			cycleStart()
			ROR_ZX()
			cycleEnd()
		case 0xF5: //SBC
			cycleStart()
			SBC_ZX()
			cycleEnd()
		case 0x95: //STA
			cycleStart()
			STA_ZX()
			cycleEnd()
		case 0x94: //STY
			cycleStart()
			STY_ZX()
			cycleEnd()

		// Y Indexed Zero Page addressing mode instructions
		/*
			$nn,Y

			This form of addressing is used in conjunction with the Y index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the low order 8 bits of memory and crossing of page boundaries does not occur.

			Bytes: 2
		*/
		case 0xB6: //LDX
			cycleStart()
			LDX_ZY()
			cycleEnd()
		case 0x96: //STX
			cycleStart()
			STX_ZY()
			cycleEnd()

		// X Indexed Zero Page Indirect addressing mode instructions
		/*
			($nn,X)

			In indexed indirect addressing, the second byte of the instruction is added to the contents of the X index register, discarding the carry. The result of this addition points to a memory location on page zero whose contents is the high order eight bits of the effective address. The next memory location in page zero contains the low order eight bits of the effective address. Both memory locations specifying the low and high order bytes of the effective address must be in page zero.

			Bytes: 2
		*/
		case 0x61: //ADC
			cycleStart()
			ADC_IX()
			cycleEnd()
		case 0x21: //AND
			cycleStart()
			AND_IX()
			cycleEnd()
		case 0xC1: //CMP
			cycleStart()
			CMP_IX()
			cycleEnd()
		case 0x41: //EOR
			cycleStart()
			EOR_IX()
			cycleEnd()
		case 0xA1: //LDA
			cycleStart()
			LDA_IX()
			cycleEnd()
		case 0x01: //ORA
			cycleStart()
			ORA_IX()
			cycleEnd()
		case 0xE1: //SBC
			cycleStart()
			SBC_IX()
			cycleEnd()
		case 0x81: //STA
			cycleStart()
			STA_IX()
			cycleEnd()

		// Zero Page Indirect Y Indexed addressing mode instructions
		/*
			($nn),Y

			In indirect indexed addressing, the second byte of the instruction points to a memory location in page zero. The contents of this memory location is added to the contents of the Y index register, the result being the high order eight bits of the effective address. The carry from this addition is added to the contents of the next page zero memory location, the result being the low order eight bits of the effective address.

			Bytes: 2
		*/
		case 0x71: //ADC
			cycleStart()
			ADC_IY()
			cycleEnd()
		case 0x31: //AND
			cycleStart()
			AND_IY()
			cycleEnd()
		case 0xD1: //CMP
			cycleStart()
			CMP_IY()
			cycleEnd()
		case 0x51: //EOR
			cycleStart()
			EOR_IY()
			cycleEnd()
		case 0xB1: //LDA
			cycleStart()
			LDA_IY()
			cycleEnd()
		case 0x11: //ORA
			cycleStart()
			ORA_IY()
			cycleEnd()
		case 0xF1: //SBC
			cycleStart()
			SBC_IY()
			cycleEnd()
		case 0x91: //STA
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
		case 0x10:
			cycleStart()
			/*
				BPL - Branch on Result Plus
			*/
			disassembledInstruction = fmt.Sprintf("BPL $%02X", (PC+2+int(operand1()))&0xFF)
			disassembleOpcode()

			offset := operand1()
			signedOffset := int8(offset)
			// Calculate the branch target address
			targetAddress := PC + 2 + int(signedOffset)
			// If N flag is not set, branch to address
			if getSRBit(7) == 0 {
				// Branch
				setPC(targetAddress)
				updateCycleCounter(1)
				//handleState(0)
				instructionCounter++
			} else {
				// Don't branch
				// Increment the instruction counter by 2
				updateCycleCounter(1)
				handleState(2)
			}
		case 0x30:
			cycleStart()
			/*
				BMI - Branch on Result Minus
			*/
			disassembledInstruction = fmt.Sprintf("BMI $%02X", (PC+2+int(int8(operand1())))&0xFFFF)
			disassembleOpcode()

			// Get offset from operand
			offset := int8(operand1())
			// If N flag is set, branch to address
			if getSRBit(7) == 1 {
				// Branch
				// Add offset to PC (already incremented by 2)
				setPC(PC + 2 + int(offset))
			} else {
				// Don't branch
				setPC(PC + 2)
			}
			cycleEnd()
		case 0x50:
			cycleStart()
			/*
				BVC - Branch on Overflow Clear
			*/
			disassembledInstruction = fmt.Sprintf("BVC $%02X", PC+2+int(operand1()))
			disassembleOpcode()

			// Get offset from operand
			offset := operand1()
			// If overflow flag is not set, branch to address
			if getSRBit(6) == 0 {
				updateCycleCounter(1)
				handleState(0)
				// Branch
				// Add offset to lower 8bits of PC
				setPC(PC + 3 + int(offset)&0xFF)
				// If the offset is negative, decrement the PC by 1
				// If bit 7 is unset then it's negative
				if readBit(7, offset) == 0 {
					decPC(1)
				}
			} else {
				// Don't branch
				updateCycleCounter(1)
				handleState(2)
			}
			cycleEnd()
		case 0x55:
			cycleStart()
			/*
				EOR - "Exclusive OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("EOR $%02X,X", operand1())
			disassembleOpcode()

			EOR("zeropagex")
			cycleEnd()
		case 0x70:
			cycleStart()
			/*
				BVS - Branch on Overflow Set
			*/
			disassembledInstruction = fmt.Sprintf("BVS $%02X", PC+2+int(operand1()))
			disassembleOpcode()

			// Get offset from operand
			offset := operand1()
			// If overflow flag is set, branch to address
			if getSRBit(6) == 1 {
				updateCycleCounter(1)
				handleState(0)
				// Branch
				// Add offset to lower 8bits of PC
				setPC(PC + 3 + int(offset)&0xFF)
				// If the offset is negative, decrement the PC by 1
				// If bit 7 is unset then it's negative
				if readBit(7, offset) == 0 {
					decPC(1)
				}
			} else {
				// Don't branch
				updateCycleCounter(1)
				handleState(2)
			}
			cycleEnd()
		case 0x90:
			cycleStart()
			/*
				BCC - Branch on Carry Clear
			*/
			disassembledInstruction = fmt.Sprintf("BCC $%02X", PC+2+int(operand1()))
			disassembleOpcode()

			previousPC = PC
			previousOpcode = opcode()
			previousOperand1 = operand1()
			// Get offset from operand
			offset := int8(operand1())
			target := PC + 2 + int(offset)

			if getSRBit(0) == 0 {
				setPC(target)
			} else {
				// Don't branch
				updateCycleCounter(1)
				handleState(2)
			}
			cycleEnd()
		case 0xB0:
			cycleStart()
			/*
				BCS - Branch on Carry Set
			*/
			disassembledInstruction = fmt.Sprintf("BCS $%02X", (PC+2+int(operand1()))&0xFF)
			disassembleOpcode()
			// Get offset from operand
			offset := operand1()
			// If carry flag is set, branch to address
			if getSRBit(0) == 1 {
				updateCycleCounter(1)
				handleState(0)
				// Branch
				// Add offset to lower 8bits of PC
				setPC(PC + 3 + int(offset)&0xFF)
				// If the offset is negative, decrement the PC by 1
				// If bit 7 is unset then it's negative
				if readBit(7, offset) == 0 {
					decPC(1)
				}
			} else {
				// Don't branch
				updateCycleCounter(1)
				handleState(2)
			}
			cycleEnd()
		case 0xD0:
			cycleStart()
			/*
				BNE - Branch on Result Not Zero
			*/

			disassembledInstruction = fmt.Sprintf("BNE $%04X", PC+2+int(int8(operand1())))
			disassembleOpcode()

			// Fetch offset from operand
			offset := operand1()

			// Check Z flag to determine if branching is needed
			if getSRBit(1) == 0 {
				// Calculate the branch target address
				targetAddr := PC + 2 + int(int8(offset))
				// Check if the branch crosses a page boundary
				if (PC & 0xFF00) != (targetAddr & 0xFF00) {
					updateCycleCounter(1)
					handleState(2)
				} else {
					updateCycleCounter(1)
					handleState(1)
				}
				// Update the program counter
				setPC(targetAddr & 0xFFFF)
			} else {
				// If Z flag is set, don't branch
				updateCycleCounter(1)
				handleState(2)
			}
			cycleEnd()
		case 0xF0:
			cycleStart()
			/*
			   BEQ - Branch on Result Zero
			*/
			disassembledInstruction = fmt.Sprintf("BEQ $%04X", PC+2+int(int8(operand1())))
			disassembleOpcode()

			// Get offset from address in operand
			offset := int8(operand1()) // Cast to signed 8-bit integer to handle negative offsets

			// If Z flag is set, branch to address
			if getSRBit(1) == 1 {
				updateCycleCounter(1)
				handleState(0)
				// Add 2 to PC (1 for opcode, 1 for operand) and then add offset
				setPC(PC + 2 + int(offset))
			} else {
				// Don't branch
				updateCycleCounter(1)
				handleState(2)
			}
			cycleEnd()
		case 0xF6:
			cycleStart()
			/*
				INC - Increment Memory By One
			*/
			disassembledInstruction = fmt.Sprintf("INC $%02X,X", operand1())
			disassembleOpcode()

			INC("zeropagex")
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
		case 0x6D:
			cycleStart()
			/*
				ADC - Add Memory to Accumulator with Carry
			*/
			disassembledInstruction = fmt.Sprintf("ADC $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			ADC("absolute")
			cycleEnd()
		case 0x2D:
			cycleStart()
			/*
				AND - "AND" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("AND $%02X%02X", operand2(), operand1())
			disassembleOpcode()

			AND("absolute")
			cycleEnd()
		case 0x0E:
			cycleStart()
			/*
				ASL - Arithmetic Shift Left
			*/
			disassembledInstruction = fmt.Sprintf("ASL $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			ASL("absolute")
			cycleEnd()
		case 0x2C:
			cycleStart()
			/*
				BIT - Test Bits in Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("BIT $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			BIT("absolute")
			cycleEnd()
		case 0xCD:
			cycleStart()
			/*
				CMP - Compare Memory and Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("CMP $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			CMP("absolute")
			cycleEnd()
		case 0xEC:
			cycleStart()
			/*
				CPX - Compare Index Register X To Memory
			*/
			disassembledInstruction = fmt.Sprintf("CPX $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			CPX("absolute")
			cycleEnd()
		case 0xCC:
			cycleStart()
			/*
				CPY - Compare Index Register Y To Memory
			*/
			disassembledInstruction = fmt.Sprintf("CPY $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			CPY("absolute")
			cycleEnd()
		case 0xCE:
			cycleStart()
			/*
				DEC - Decrement Memory By One
			*/
			disassembledInstruction = fmt.Sprintf("DEC $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			DEC("absolute")
			cycleEnd()
		case 0x4D:
			cycleStart()
			/*
				EOR - "Exclusive OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("EOR $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			EOR("absolute")
			cycleEnd()
		case 0xEE:
			cycleStart()
			/*
				INC - Increment Memory By One
			*/
			disassembledInstruction = fmt.Sprintf("INC $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			INC("absolute")
			cycleEnd()
		case 0x4C:
			cycleStart()
			/*
				JMP - JMP Absolute
			*/
			disassembledInstruction = fmt.Sprintf("JMP $%04X", int(operand2())<<8|int(operand1()))
			disassembleOpcode()
			// For AllSuiteA.bin 6502 opcode test suite
			if *allsuitea && readMemory(0x210) == 0xFF {
				fmt.Printf("\n\u001B[32;5mMemory address $210 == $%02X. All opcodes succesfully tested and passed!\u001B[0m\n", readMemory(0x210))
				os.Exit(0)
			}
			JMP("absolute")
			cycleEnd()
		case 0x20:
			cycleStart()
			/*
				JSR - Jump To Subroutine
			*/
			disassembledInstruction = fmt.Sprintf("JSR $%04X", int(operand2())<<8|int(operand1()))
			disassembleOpcode()
			// First, push the high byte
			decSP()
			updateStack(byte(PC >> 8))
			decSP()
			updateStack(byte((PC)&0xFF) + 2)

			previousPC = PC
			previousOpcode = opcode()
			previousOperand1 = operand1()
			previousOperand2 = operand2()
			// Now, jump to the subroutine address specified by the operands
			setPC(int(operand2())<<8 | int(operand1()))
			updateCycleCounter(1)
			handleState(0)
			cycleEnd()
		case 0xAD:
			cycleStart()
			/*
				LDA - Load Accumulator with Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDA $%04X", uint16(operand2())<<8|uint16(operand1()))
			disassembleOpcode()
			LDA("absolute")
			cycleEnd()
		case 0xAE:
			cycleStart()
			/*
				LDX - Load Index Register X From Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDX $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			LDX("absolute")
			cycleEnd()
		case 0xAC:
			cycleStart()
			/*
				LDY - Load Index Register Y From Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDY $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			LDY("absolute")
			cycleEnd()
		case 0x4E:
			cycleStart()
			/*
				LSR - Logical Shift Right
			*/
			disassembledInstruction = fmt.Sprintf("LSR $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			LSR("absolute")
			cycleEnd()
		case 0x0D:
			cycleStart()
			/*
				ORA - "OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("ORA $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			ORA("absolute")
			cycleEnd()
		case 0x2E:
			cycleStart()
			/*
				ROL - Rotate Left
			*/
			disassembledInstruction = fmt.Sprintf("ROL $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			ROL("absolute")
			cycleEnd()
		case 0x6E:
			cycleStart()
			/*
				ROR - Rotate Right
			*/
			disassembledInstruction = fmt.Sprintf("ROR $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			ROR("absolute")
			cycleEnd()
		case 0xED:
			cycleStart()
			/*
				SBC - Subtract Memory from Accumulator with Borrow
			*/
			disassembledInstruction = fmt.Sprintf("SBC $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			SBC("absolute")
			cycleEnd()
		case 0x8D:
			cycleStart()
			/*
				STA - Store Accumulator in Memory
			*/
			disassembledInstruction = fmt.Sprintf("STA $%04X", uint16(operand2())<<8|uint16(operand1()))
			disassembleOpcode()
			STA("absolute")
			cycleEnd()
		case 0x8E:
			cycleStart()
			/*
				STX - Store Index Register X In Memory
			*/
			disassembledInstruction = fmt.Sprintf("STX $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			STX("absolute")
			cycleEnd()
		case 0x8C:
			cycleStart()
			/*
				STY - Store Index Register Y In Memory
			*/
			disassembledInstruction = fmt.Sprintf("STY $%02X%02X", operand2(), operand1())
			disassembleOpcode()
			STY("absolute")
			cycleEnd()

		// X Indexed Absolute addressing mode instructions
		/*
			$nnnn,X

			This form of addressing is used in conjunction with the X index register. The effective address is formed by adding the contents of X to the address contained in the second and third bytes of the instruction. This mode allows the index register to contain the index or count value and the instruction to contain the base address. This type of indexing allows any location referencing and the index to modify multiple fields resulting in reduced coding and execution time.

			Note on the MOS 6502:

			The value at the specified address, ignoring the the addressing mode's X offset, is read (and discarded) before the final address is read. This may cause side effects in I/O registers.


			Bytes: 3
		*/
		case 0x7D:
			cycleStart()
			/*
				ADC - Add Memory to Accumulator with Carry
			*/
			disassembledInstruction = fmt.Sprintf("ADC $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			ADC("absolutex")
			cycleEnd()
		case 0x3D:
			cycleStart()
			/*
				AND - "AND" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("AND $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			AND("absolutex")
			cycleEnd()
		case 0x1E:
			cycleStart()
			/*
				ASL - Arithmetic Shift Left
			*/
			disassembledInstruction = fmt.Sprintf("ASL $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			ASL("absolutex")
			cycleEnd()
		case 0xDD:
			cycleStart()
			/*
				CMP - Compare Memory and Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("CMP $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			CMP("absolutex")
			cycleEnd()
		case 0xDE:
			cycleStart()
			/*
				DEC - Decrement Memory By One
			*/
			disassembledInstruction = fmt.Sprintf("DEC $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()

			DEC("absolutex")
			cycleEnd()
		case 0x5D:
			cycleStart()
			/*
				EOR - "Exclusive OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("EOR $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			EOR("absolutex")
			cycleEnd()
		case 0xFE:
			cycleStart()
			/*
				INC - Increment Memory By One
			*/
			disassembledInstruction = fmt.Sprintf("INC $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			INC("absolutex")
			cycleEnd()
		case 0xBD:
			cycleStart()
			/*
				LDA - Load Accumulator with Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDA $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			LDA("absolutex")
			cycleEnd()
		case 0xBC:
			cycleStart()
			/*
				LDY - Load Index Register Y From Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDY $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			LDY("absolutex")
			cycleEnd()
		case 0x5E:
			cycleStart()
			/*
				LSR - Logical Shift Right
			*/
			disassembledInstruction = fmt.Sprintf("LSR $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			LSR("absolutex")
			cycleEnd()
		case 0x1D:
			cycleStart()
			/*
				ORA - "OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("ORA $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			ORA("absolutex")
			cycleEnd()
		case 0x3E:
			cycleStart()
			/*
			 */
			disassembledInstruction = fmt.Sprintf("ROL $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			ROL("absolutex")
			cycleEnd()
		case 0x7E:
			cycleStart()
			/*
				ROR - Rotate Right
			*/
			disassembledInstruction = fmt.Sprintf("ROR $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			ROR("absolutex")
			cycleEnd()
		case 0xFD:
			cycleStart()
			/*
				SBC - Subtract Memory from Accumulator with Borrow
			*/
			disassembledInstruction = fmt.Sprintf("SBC $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			SBC("absolutex")
			cycleEnd()
		case 0x9D:
			cycleStart()
			/*
				STA - Store Accumulator in Memory
			*/
			disassembledInstruction = fmt.Sprintf("STA $%02X%02X,X", operand2(), operand1())
			disassembleOpcode()
			STA("absolutex")
			cycleEnd()

		// Y Indexed Absolute addressing mode instructions
		/*
			$nnnn,Y

			This form of addressing is used in conjunction with the Y index register. The effective address is formed by adding the contents of Y to the address contained in the second and third bytes of the instruction. This mode allows the index register to contain the index or count value and the instruction to contain the base address. This type of indexing allows any location referencing and the index to modify multiple fields resulting in reduced coding and execution time.

			Note on the MOS 6502:

			The value at the specified address, ignoring the the addressing mode's Y offset, is read (and discarded) before the final address is read. This may cause side effects in I/O registers.

			Bytes: 3
		*/
		case 0x79:
			cycleStart()
			/*
				ADC - Add Memory to Accumulator with Carry
			*/
			disassembledInstruction = fmt.Sprintf("ADC $%04X,Y", int(operand2())<<8|int(operand1()))
			disassembleOpcode()
			ADC("absolutey")
			cycleEnd()
		case 0x39:
			cycleStart()
			/*
				AND - "AND" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("AND $%02X%02X,Y", operand2(), operand1())
			disassembleOpcode()
			AND("absolutey")
			cycleEnd()
		case 0xD9:
			cycleStart()
			/*
				CMP - Compare Memory and Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("CMP $%02X%02X,Y", operand2(), operand1())
			disassembleOpcode()
			CMP("absolutey")
			cycleEnd()
		case 0x59:
			cycleStart()
			/*
				EOR - "Exclusive OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("EOR $%02X%02X,Y", operand2(), operand1())
			disassembleOpcode()
			EOR("absolutey")
			cycleEnd()
		case 0xB9:
			cycleStart()
			/*
				LDA - Load Accumulator with Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDA $%02X%02X,Y", operand2(), operand1())
			disassembleOpcode()
			LDA("absolutey")
			cycleEnd()
		case 0xBE:
			cycleStart()
			/*
				LDX - Load Index Register X From Memory
			*/
			disassembledInstruction = fmt.Sprintf("LDX $%02X%02X,Y", operand2(), operand1())
			disassembleOpcode()
			LDX("absolutey")
			cycleEnd()
		case 0x19:
			cycleStart()
			/*
				ORA - "OR" Memory with Accumulator
			*/
			disassembledInstruction = fmt.Sprintf("ORA $%02X%02X,Y", operand2(), operand1())
			disassembleOpcode()
			ORA("absolutey")
			cycleEnd()
		case 0xF9:
			cycleStart()
			/*
				SBC - Subtract Memory from Accumulator with Borrow
			*/
			disassembledInstruction = fmt.Sprintf("SBC $%02X%02X,Y", operand2(), operand1())
			disassembleOpcode()
			SBC("absolutey")
			cycleEnd()
		case 0x99:
			cycleStart()
			/*
				STA - Store Accumulator in Memory
			*/
			disassembledInstruction = fmt.Sprintf("STA $%02X%02X,Y", operand2(), operand1())
			disassembleOpcode()
			STA("absolutey")
			cycleEnd()
		// Absolute Indirect addressing mode instructions
		case 0x6C:
			cycleStart()
			/*
				JMP - JMP Indirect
			*/
			disassembledInstruction = fmt.Sprintf("JMP ($%02X%02X)", operand2(), operand1())
			disassembleOpcode()
			JMP("indirect")
			cycleEnd()
		}
		if *plus4 {
			plus4KernalRoutines()
		}
	}
}
