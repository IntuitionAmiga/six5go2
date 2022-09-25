package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	printHex     bool
	file         []byte
	fileposition = 0 //  Byte position counter

	// CPURegisters and RAM
	A      byte        = 0x0000     // Accumulator
	X      byte        = 0x0000     // X register
	Y      byte        = 0x0000     // Y register		(76543210) SR Bit 5 is always set
	SR     byte        = 0b00100010 // Status Register	(NVEBDIZC)
	SP                 = 0x01ff     // Stack Pointer
	PC                 = 0x0000     // Program Counter
	memory [65536]byte              // Memory
)

func main() {
	fmt.Printf("Six5go2 - 6502 Emulator and Disassembler in Golang (c) 2022 Zayn Otley\n\n")

	if len(os.Args) <= 2 {
		fmt.Printf("USAGE - %s <target_filename> <entry_point_address> <hex>\n", os.Args[0])
		os.Exit(0)
	}
	if len(os.Args) > 2 {
		parseUint, _ := strconv.ParseUint(os.Args[2], 16, 16)
		PC = int(parseUint)
	}
	if len(os.Args) > 3 && os.Args[3] == "hex" {
		printHex = true
	}

	//  Read file
	file, _ = os.ReadFile(os.Args[1])

	fmt.Printf("USAGE   - six5go2 <target_filename> <entry_point> (Hex memory address) <hex> (Print hex values above each instruction) \n")
	fmt.Printf("EXAMPLE - six5go2 cbmbasic35.rom 0800 hex\n\n")
	fmt.Printf("Length of file %s is %v ($%04X) bytes\n\n", os.Args[1], len(file), len(file))

	fmt.Printf("Size of addressable memory is %v ($%04X) bytes\n\n", len(memory), len(memory))

	//  Copy file into memory
	copy(memory[:], file)
	// Start emulation
	printMachineState()
	execute(string(file))
}
func opcode() byte {
	return memory[fileposition]
}
func operand1() byte {
	return memory[fileposition+1]
}
func operand2() byte {
	return memory[fileposition+2]
}
func incCount(amount int) {
	if fileposition+amount < len(file)-1 && amount != 0 {
		fileposition += amount
	}
	PC += amount
	printMachineState()
}
func printMachineState() {
	fmt.Printf("A=$%02X X=$%02X Y=$%02X SR=%08b (NVEBDIZC) SP=$%08X PC=$%04X\n\n", A, X, Y, SR, SP, PC)
}
func getSRBit(x byte) byte {
	return (SR >> x) & 1
}
func setSRBitOn(x byte) {
	SR |= 1 << x
}
func setSRBitOff(x byte) {
	SR &= ^(1 << x)
}
func getABit(x byte) byte {
	return (A >> x) & 1
}
func getXBit(x byte) byte {
	return (X >> x) & 1
}
func getYBit(x byte) byte {
	return (Y >> x) & 1
}
func setABitOn(x byte) {
	A |= 1 << x
}
func setABitOff(x byte) {
	A &= ^(1 << x)
}

func execute(file string) {
	PC += fileposition
	if printHex {
		fmt.Printf(" * = $%04X\n\n", PC)
	}
	for fileposition = 0; fileposition < len(file); {
		//  1 byte instructions with no operands
		switch opcode() {
		case 0x00:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("BRK\n")
			incCount(1)
		case 0x02:
			/*
				CLE - Clear Extend Disable Flag
				Operation: 0 → E

				This instruction initializes the extend disable to a 0. This sets the stack pointer to 16 bit mode.

				It affects no registers in the microprocessor and no flags other than the extend disable which is cleared.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("CLE\n")
			// Set bit 5 of SR to 0
			setSRBitOff(5)
			incCount(1)
		case 0x03:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("SEE\n")
			incCount(1)
		case 0x08:
			/*
				PHP - Push Processor Status On Stack
				Operation: P↓

				This instruction transfers the contents of the processor status register unchanged to the stack,
				as governed by the stack pointer.

				The PHP instruction affects no registers or flags in the microprocessor.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("PHP\n")

			// Push SR to stack
			memory[SP] = SR
			// Decrement the stack pointer by 1 byte
			SP--
			incCount(1)
		case 0x0A:
			/*
				ASL - Arithmetic Shift Left
				Operation: C ← /M7...M0/ ← 0

				The shift left instruction shifts either the accumulator or the address memory location 1 bit to
				the left, with the bit 0 always being set to 0 and the the input bit 7 being stored in the carry flag.

				ASL either shifts the accumulator left 1 bit or is a read/modify/write instruction that affects only memory.

				The instruction does not affect the overflow bit, sets N equal to the result bit 7 (bit 6 in the input),
				sets Z flag if the result is equal to 0, otherwise resets Z and stores the input bit 7 in the carry flag
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, opcode())
			}
			fmt.Printf("ASL\n")

			// Shift left the accumulator by 1 bit
			A <<= 1
			// Set SR negative flag if bit 7 of A is set
			if getABit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// Set SR zero flag if A is 0
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(1)
		case 0x0B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("TSY\n")
			incCount(1)
		case 0x18:
			/*
				CLC - Clear Carry Flag
				Operation: 0 → C

				This instruction initializes the carry flag to a 0. This operation should normally precede an ADC loop.
				It is also useful when used with a R0L instruction to clear a bit in memory.

				This instruction affects no registers in the microprocessor and no flags other than the carry flag which is reset.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("CLC\n")

			// Set SR carry flag bit 0 to 0
			setSRBitOff(0)
			incCount(1)
		case 0x1A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, opcode())
			}
			fmt.Printf("INC\n")
			incCount(1)
		case 0x1B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("INZ\n")
			incCount(1)
		case 0x28:
			/*
				PLP - Pull Processor Status From Stack
				Operation: P↑

				This instruction transfers the next value on the stack to the Processor Status register,
				thereby changing all of the flags and setting the mode switches to the values from the stack.

				The PLP instruction affects no registers in the processor other than the status register.

				This instruction could affect all flags in the status register.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("PLP\n")

			// Update SR with the value stored at the address pointed to by SP
			SR = memory[SP]
			incCount(1)
		case 0x2A:
			/*
				ROL - Rotate Left
				Operation: C ← /M7...M0/ ← C

				The rotate left instruction shifts either the accumulator or addressed memory left 1 bit, with
				the input carry being stored in bit 0 and with the input bit 7 being stored in the carry flags.

				The ROL instruction either shifts the accumulator left 1 bit and stores the carry in accumulator bit 0
				or does not affect the internal registers at all.
				The ROL instruction sets carry equal to the input bit 7,
				sets N equal to the input bit 6,
				sets the Z flag if the result of the rotate is 0,
				otherwise it resets Z and does not affect the overflow flag at all.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, opcode())
			}
			fmt.Printf("ROL\n")

			// Shift left the accumulator by 1 bit
			A <<= 1
			// Set SR carry flag bit 0 to the value of Accumulator bit 7
			if getABit(7) == 1 {
				setSRBitOn(0)
			} else {
				setSRBitOff(0)
			}

			// Set SR negative flag bit 7 to the value of Accumulator bit 6
			if getABit(6) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}

			// Set SR zero flag bit 1 to 1 if Accumulator is 0 else set SR zero flag to 0
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(1)
		case 0x2B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("TYS\n")
			incCount(1)
		case 0x38:
			/*
				SEC - Set Carry Flag
				Operation: 1 → C

				This instruction initializes the carry flag to a 1.
				This operation should normally precede an SBC loop.
				It is also useful when used with a ROL instruction to initialize a bit in memory to a 1.

				This instruction affects no registers in the microprocessor and no flags other than the carry
				flag which is set.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, opcode())
			}
			fmt.Printf("SEC\n")

			// Set SR carry flag bit 0 to 1
			setSRBitOn(0)
			incCount(1)
		case 0x3A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, opcode())
			}
			fmt.Printf("DEC\n")
			incCount(1)
		case 0x3B:
			// NOP
			// 65CE02 only - DEZ - Decrement Z Register
			incCount(1)
		case 0x40:
			/*
				RTI - Return From Interrupt
				Operation: P↑ PC↑

				This instruction transfers from the stack into the microprocessor the processor status and the
				program counter location for the instruction which was interrupted.

				By virtue of the interrupt having stored this data before executing the instruction and the fact
				that the RTI re-initialises the microprocessor to the same state as when it was interrupted, the
				combination of interrupt plus RTI allows truly reentrant coding.

				The RTI instruction re-initialises all flags to the position to the point they were at the time
				the interrupt was taken and sets the program counter back to its pre-interrupt state.

				It affects no other registers in the microprocessor.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("RTI\n")

			// Update SR with the value stored at the address pointed to by SP
			SR = memory[SP]
			// Update PC with the value stored at the address pointed to by SP+1
			PC = int(memory[SP] + 1)
			incCount(1)
		case 0x42:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, opcode())
			}
			fmt.Printf("NEG\n")
			incCount(1)
		case 0x43:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, opcode())
			}
			fmt.Printf("ASR\n")
			incCount(1)
		case 0x48:
			/*
				PHA - Push Accumulator On Stack
				Operation: A↓

				This instruction transfers the current value of the accumulator to the next location on the stack,
				automatically decrementing the stack to point to the next empty location.

				The Push A instruction only affects the stack pointer register which is decremented by 1 as a result of
				the operation. It affects no flags.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("PHA\n")

			// Update memory address pointed to by SP with value stored in accumulator
			memory[SP] = A
			// Decrement the stack pointer by 1 byte
			SP--
			incCount(1)
		case 0x4A:
			/*
				LSR - Logical Shift Right
				Operation: 0 → /M7...M0/ → C

				This instruction shifts either the accumulator or a specified memory location 1 bit to the right,
				with the higher bit of the result always being set to 0, and the low bit which is shifted out of
				the field being stored in the carry flag.

				The shift right instruction either affects the accumulator by shifting it right 1 or is a
				read/modify/write instruction which changes a specified memory location but does not affect
				any internal registers. The shift right does not affect the overflow flag.
				The N flag is always reset.
				The Z flag is set if the result of the shift is 0 and reset otherwise.
				The carry is set equal to bit 0 of the input.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, opcode())
			}
			fmt.Printf("LSR\n")

			// Shift A right 1 bit
			A >>= 1
			// Set SR negative flag bit 7 to 0
			setSRBitOff(7)
			// Set SR zero flag bit 1 to 1 if A is 0 else set it to 0
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			// Set SR carry flag bit 0 to bit 0 of A
			if getABit(0) == 1 {
				setSRBitOn(0)
			} else {
				setSRBitOff(0)
			}
			incCount(1)
		case 0x4B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("TAZ\n")
			incCount(1)
		case 0x58:
			/*
				CLI - Clear Interrupt Disable
				Operation: 0 → I

				This instruction initializes the interrupt disable to a 0.
				his allows the microprocessor to receive interrupts.

				It affects no registers in the microprocessor and no flags other than the interrupt disable
				which is cleared.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("CLI\n")

			// Set SR interrupt disable bit 2 to 0
			setSRBitOff(2)
			incCount(1)
		case 0x5A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("PHY\n")
			incCount(1)
		case 0x5B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("TAB\n")
			incCount(1)
		case 0x5C:
			/*
				AUG - Augment
				Operation: No Operation

				The AUG instruction is a 4-byte NOP, and reserved for future expansion.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("AUG\n")
			incCount(1)
		case 0x60:
			/*
				RTS - Return From Subroutine
				Operation: PC↑, PC + 1 → PC

				This instruction loads the program count low and program count high from the stack into the program
				counter and increments the program counter so that it points to the instruction following the JSR.

				The stack pointer is adjusted by incrementing it twice.

				The RTS instruction does not affect any flags and affects only PCL and PCH.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("RTS\n")

			// Update PC with the value stored at the address pointed to by SP+1
			PC = int(memory[SP] + 1)
			// Increment the stack pointer twice
			SP += 2
			incCount(1)
		case 0x68:
			/*
				PLA - Pull Accumulator From Stack
				Operation: A↑

				This instruction adds 1 to the current value of the stack pointer and uses it to address the stack
				and loads the contents of the stack into the A register.

				The PLA instruction does not affect the carry or overflow flags.
				It sets N if the bit 7 is on in accumulator A as a result of instructions, otherwise it is reset.
				If accumulator A is zero as a result of the PLA, then the Z flag is set, otherwise it is reset.

				The PLA instruction changes content of the accumulator A to the contents of the memory location at
				stack register plus 1 and also increments the stack register.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("PLA\n")

			// Increment the stack pointer by 1 byte
			SP++
			// Update accumulator with value stored in memory address pointed to by SP
			A = memory[SP]
			// If bit 7 of accumulator is set, set negative SR flag else set negative SR flag to 0
			if getABit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// If accumulator is 0, set zero SR flag else set zero SR flag to 0
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(1)
		case 0x6A:
			/*
				ROR - Rotate Right
				Operation: C → /M7...M0/ → C

				The rotate right instruction shifts either the accumulator or addressed memory right 1 bit with
				bit 0 shifted into the carry and carry shifted into bit 7.

				The ROR instruction either shifts the accumulator right 1 bit and stores the carry in accumulator
				bit 7 or does not affect the internal registers at all.
				The ROR instruction sets carry equal to input bit 0,
				sets N equal to the input carry and sets the Z flag if the result of the rotate is 0;
				It otherwise it resets Z and does not affect the overflow flag at all.

				(Available on Microprocessors after June, 1976)
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, opcode())
			}
			fmt.Printf("ROR\n")

			Atemp := A
			// Shift accumulator right 1 bit
			A >>= 1
			// Update accumulator bit 7 with bit 0 of Atemp
			if Atemp&1<<7 == 1 {
				setABitOn(7)
			} else {
				setABitOff(7)
			}
			// If carry is 1, set accumulator bit 0 to 1 else set accumulator bit 0 to 0
			if getSRBit(0) == 1 {
				setABitOn(0)
			} else {
				setABitOff(0)
			}
			// If accumulator bit 7 is 1, set carry to 1 else set carry to 0
			if getABit(7) == 1 {
				setSRBitOn(0)
			} else {
				setSRBitOff(0)
			}
			// If accumulator is 0, set zero SR flag else set zero SR flag to 0
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(1)
		case 0x6B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("TZA\n")
			incCount(1)
		case 0x78:
			/*
				SEI - Set Interrupt Disable
				Operation: 1 → I

				This instruction initializes the interrupt disable to a 1.
				It is used to mask interrupt requests during system reset operations and during interrupt commands.

				It affects no registers in the microprocessor and no flags other than the interrupt disable which is set.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("SEI\n")

			// Set SR interrupt disable bit 2 to 1
			setSRBitOn(2)
			incCount(1)
		case 0x7A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("PLY\n")
			incCount(1)
		case 0x7B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(TBA - Absolute,Y)\n", PC, opcode())
			}
			fmt.Printf("TBA\n")
			incCount(1)
		case 0x88:
			/*
				DEY - Decrement Index Register Y By One
				Operation: Y - 1 → Y

				This instruction subtracts one from the current value in the index register Y and stores the result
				into the index register Y. The result does not affect or consider carry so that the value in the index
				register Y is decremented to 0 and then through 0 to FF.

				Decrement Y does not affect the carry or overflow flags;
				if the Y register contains bit 7 on as a result of the decrement the N flag is set,
				otherwise the N flag is reset.
				If the Y register is 0 as a result of the decrement, the Z flag is set otherwise the Z flag is reset.
				This instruction only affects the index register Y.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("DEY\n")

			// Decrement the  Y register by 1
			Y--
			// If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getYBit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			incCount(1)
		case 0x8A:
			/*
				TXA - Transfer Index X To Accumulator
				Operation: X → A

				This instruction moves the value that is in the index register X to the accumulator A without disturbing
				the content of the index register X.

				TXA does not affect any register other than the accumulator and does not affect the carry or overflow flag.
				If the result in A has bit 7 on, then the N flag is set, otherwise it is reset.
				If the resultant value in the accumulator is 0, then the Z flag is set, otherwise it is reset.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("TXA\n")

			// Set accumulator to value of X register
			A = X
			// If bit 7 of accumulator is set, set negative SR flag else set negative SR flag to 0
			if getABit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// If accumulator is 0, set zero SR flag else set zero SR flag to 0
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(1)
		case 0x98:
			/*
				TYA - Transfer Index Y To Accumulator
				Operation: Y → A

				This instruction moves the value that is in the index register Y to accumulator A without disturbing
				the content of the register Y.

				TYA does not affect any other register other than the accumulator and does not affect the carry
				or overflow flag.
				If the result in the accumulator A has bit 7 on, the N flag is set, otherwise it is reset.
				If the resultant value in the accumulator A is 0, then the Z flag is set, otherwise it is reset.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("TYA\n")

			// Set accumulator to value of Y register
			A = Y
			// If bit 7 of accumulator is set, set negative SR flag else set negative SR flag to 0
			if getABit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// If accumulator is 0, set zero SR flag else set zero SR flag to 0
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(1)
		case 0x9A:
			/*
				TXS - Transfer Index X To Stack Pointer
				Operation: X → S

				This instruction transfers the value in the index register X to the stack pointer.

				TXS changes only the stack pointer, making it equal to the content of the index register X.
				It does not affect any of the flags.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("TXS\n")

			// Set stack pointer to value of X register
			SP = int(X)
			incCount(1)
		case 0xA8:
			/*
				TAY - Transfer Accumulator To Index Y
				Operation: A → Y

				This instruction moves the value of the accumulator into index register Y without affecting
				the accumulator.

				TAY instruction only affects the Y register and does not affect either the carry or overflow flags.
				If the index register Y has bit 7 on, then N is set, otherwise it is reset.
				If the content of the index register Y equals 0 as a result of the operation, Z is set on, otherwise it is reset.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("TAY\n")

			// Set Y register to the value of the accumulator
			Y = A
			// If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getYBit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// If Y register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(1)
		case 0xAA:
			/*
				TAX - Transfer Accumulator To Index X
				Operation: A → X

				This instruction takes the value from accumulator A and transfers or loads it into the index register X
				without disturbing the content of the accumulator A.

				TAX only affects the index register X, does not affect the carry or overflow flags.
				The N flag is set if the resultant value in the index register X has bit 7 on, otherwise N is reset.
				The Z bit is set if the content of the register X is 0 as a result of the operation, otherwise it is reset.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("TAX\n")

			// Update X with the value of A
			X = A
			// If X register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getXBit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// If X register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
			if X == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(1)
		case 0xB8:
			/*
				CLV - Clear Overflow Flag
				Operation: 0 → V

				This instruction clears the overflow flag to a 0. This command is used in conjunction with the
				set overflow pin which can change the state of the overflow flag with an external signal.

				CLV affects no registers in the microprocessor and no flags other than the overflow flag which
				is set to a 0.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("CLV\n")

			// Set SR overflow flag bit 6 to 0
			setSRBitOff(6)
			incCount(1)
		case 0xBA:
			/*
				TSX - Transfer Stack Pointer To Index X
				Operation: S → X

				This instruction transfers the value in the stack pointer to the index register X.

				TSX does not affect the carry or overflow flags.
				It sets N if bit 7 is on in index X as a result of the instruction, otherwise it is reset.
				If index X is zero as a result of the TSX, the Z flag is set, otherwise it is reset.
				TSX changes the value of index X, making it equal to the content of the stack pointer.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("TSX\n")

			// Update X with the value stored at the address pointed to by SP
			X = memory[SP]
			// If X register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getXBit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// If X register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
			if X == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(1)
		case 0xC8:
			/*
				INY - Increment Index Register Y By One
				Operation: Y + 1 → Y

				Increment Y increments or adds one to the current value in the Y register,
				storing the result in the Y register.

				As in the case of INX the primary application is to step thru a set of values using the Y register.

				The INY does not affect the carry or overflow flags, sets the N flag if the result of the increment
				has a one in bit 7, otherwise resets N,
				sets Z if as a result of the increment the Y register is zero otherwise resets the Z flag.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("INY\n")

			// Increment the  Y register by 1
			Y++
			// If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getYBit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// If Y register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
			if Y == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(1)
		case 0xCA:
			/*
				DEX - Decrement Index Register X By One
				Operation: X - 1 → X

				This instruction subtracts one from the current value of the index register X and stores the result
				in the index register X.

				DEX does not affect the carry or overflow flag, it
				sets the N flag if it has bit 7 on as a result of the decrement, otherwise it resets the N flag;
				sets the Z flag if X is a 0 as a result of the decrement, otherwise it resets the Z flag.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("DEX\n")

			// Decrement the X register by 1
			X--
			// If X register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getXBit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// If X register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
			if X == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(1)
		case 0xD8:
			/*
				CLD - Clear Decimal Mode
				Operation: 0 → D

				This instruction sets the decimal mode flag to a 0. This all subsequent ADC and SBC instructions
				to operate as simple operations.

				CLD affects no registers in the microprocessor and no flags other than the decimal mode flag which
				is set to a 0.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("CLD\n")

			setSRBitOff(3)
			incCount(1)
		case 0xDA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("PHX\n")
			incCount(1)
		case 0xDB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("PHZ\n")
			incCount(1)
		case 0xE8:
			/*
				INX - Increment Index Register X By One
				Operation: X + 1 → X

				Increment X adds 1 to the current value of the X register.
				This is an 8-bit increment which does not affect the carry operation, therefore,
				if the value of X before the increment was FF, the resulting value is 00.

				INX does not affect the carry or overflow flags;
				it sets the N flag if the result of the increment has a one in bit 7, otherwise resets N;
				sets the Z flag if the result of the increment is 0, otherwise it resets the Z flag.

				INX does not affect any other register other than the X register.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("INX\n")

			// Increment the X register by 1
			X++
			// If X register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getXBit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// If X register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
			if X == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}

			incCount(1)
		case 0xEA:
			/*
				NOP - No Operation
				Operation: No operation
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("NOP\n")
			incCount(1)
		case 0xF8:
			/*
				SED - Set Decimal Mode
				Operation: 1 → D

				This instruction sets the decimal mode flag D to a 1.
				This makes all subsequent ADC and SBC instructions operate as a decimal arithmetic operation.

				SED affects no registers in the microprocessor and no flags other than the decimal mode which
				is set to a 1.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("SED\n")

			// Set SR decimal mode flag to 1
			setSRBitOn(3)
			incCount(1)
		case 0xFA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("PLX\n")
			incCount(1)
		case 0xFB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, opcode())
			}
			fmt.Printf("PLZ\n")
			incCount(1)
		}

		// 2 byte instructions with 1 operand
		switch opcode() {
		case 0x01:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page Indirect))\n", PC, opcode(), operand1())
			}
			fmt.Printf("ORA ($%02x,X)\n", operand1())
			incCount(2)
		case 0x04:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("TSB $%02x\n", operand1())
			incCount(2)
		case 0x05:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("ORA $%02x\n", operand1())
			incCount(2)
		case 0x06:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("ASL $%02x\n", operand1())
			incCount(2)
		case 0x07:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("RMB0 $%02x\n", operand1())
			incCount(2)
		case 0x09:
			/*
				ORA - "OR" Memory with Accumulator
				Operation: A ∨ M → A

				The ORA instruction transfers the memory and the accumulator to the adder which performs a binary "OR"
				on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("ORA #$%02x\n", operand1())

			// OR the accumulator with the operand
			A |= operand1()
			// If A==0 then set SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			// If A bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getABit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			incCount(2)
		case 0x10:
			/*
				BPL - Branch on Result Plus
				Operation: Branch on N = 0

				This instruction is the complementary branch to branch on result minus.

				It is a conditional branch which takes the branch when the N bit is reset (0).

				BPL is used to test if the previous result bit 7 was off (0) and branch on result minus is used to
				determine if the previous result was minus or bit 7 was on (1).

				The instruction affects no flags or other registers other than the P counter and only affects the
				P counter when the N bit is reset.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("BPL $%02X\n", (fileposition+2+int(operand1()))&0xFF)

			// If SR negative flag bit 7 is 0 then branch
			if getSRBit(7) == 0 {
				// Branch
				fileposition += 2 + int(operand1())
				PC = fileposition
				incCount(0)
			} else {
				// Don't branch
				incCount(2)
			}
		case 0x11:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page),Indirect Y)\n", PC, opcode(), operand1())
			}
			fmt.Printf("ORA ($%02x),Y\n", operand1())
			incCount(2)
		case 0x12:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page),Indirect Z)\n", PC, opcode(), operand1())
			}
			fmt.Printf("ORA ($%02x),Z\n", operand1())
			incCount(2)
		case 0x14:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("TRB $%02x\n", operand1())
			incCount(2)
		case 0x15:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("ORA $%02x,X\n", operand1())
			incCount(2)
		case 0x16:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(ASL - Zero Page,X)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("ASL $%02X,X\n", operand1())
			incCount(2)
		case 0x17:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("RMB1 $%02X\n", operand1())
			incCount(2)
		case 0x21:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page Indirect))\n", PC, opcode(), operand1())
			}
			fmt.Printf("AND ($%02X,X)\n", operand1())
			incCount(2)
		case 0x24:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("BIT $%02X\n", operand1())
			incCount(2)
		case 0x25:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("AND $%02X\n", operand1())
			incCount(2)
		case 0x26:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("ROL $%02X\n", operand1())
			incCount(2)
		case 0x27:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("RMB2 $%02X\n", operand1())
			incCount(2)
		case 0x29:
			/*
				AND - "AND" Memory with Accumulator
				Operation: A ∧ M → A

				The AND instruction transfer the accumulator and memory to the adder which performs a bit-by-bit
				AND operation and stores the result back in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("AND #$%02X\n", operand1())

			// AND the accumulator with the operand
			A &= operand1()
			// If A==0, set zero flag
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			// If accumulator bit 7 is 1, set negative SR flag
			if getABit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			incCount(2)
		case 0x30:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("BMI $%02X\n", (fileposition+2+int(operand1()))&0xFF)
			incCount(2)
		case 0x31:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, opcode(), operand1())
			}
			fmt.Printf("AND ($%02X),Y\n", operand1())
			incCount(2)
		case 0x32:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect),Z)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("AND ($%02X),Z\n", operand1())
			incCount(2)
		case 0x34:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("BIT $%02X,X\n", operand1())
			incCount(2)
		case 0x35:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("AND $%02X,X\n", operand1())
			incCount(2)
		case 0x36:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("ROL $%02X,X\n", operand1())
			incCount(2)
		case 0x37:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\n", PC, opcode(), operand1())
			}
			fmt.Printf("RMB3 $%02X\n", operand1())
			incCount(2)
		case 0x41:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page, Indirect))\n", PC, opcode(), operand1())
			}
			fmt.Printf("EOR ($%02X,X)\n", operand1())
			incCount(2)
		case 0x44:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("ASR $%02X\n", operand1())
			incCount(2)
		case 0x45:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("EOR $%02X\n", operand1())
			incCount(2)
		case 0x46:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("LSR $%02X\n", operand1())
			incCount(2)
		case 0x47:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("RMB4 $%02X\n", operand1())
			incCount(2)
		case 0x49:
			/*
				EOR - "Exclusive OR" Memory with Accumulator
				Operation: A ⊻ M → A

				The EOR instruction transfers the memory and the accumulator to the adder which performs a binary
				"EXCLUSIVE OR" on a bit-by-bit basis and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("EOR #$%02X\n", operand1())

			// XOR the accumulator with the operand
			A ^= operand1()

			// If accumulator is 0the  set Zero flag to 1 else set Zero flag to 0
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			// If bit 7 of accumulator is set, set SR Negative flag to 1
			if getABit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			incCount(2)
		case 0x50:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("BVC $%02X\n", fileposition+2+int(operand1()))
			incCount(2)
		case 0x51:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, opcode(), operand1())
			}
			fmt.Printf("EOR ($%02X),Y\n", operand1())
			incCount(2)
		case 0x52:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect),Z)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("EOR ($%02X),Z\n", operand1())
			incCount(2)
		case 0x54:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("ASR $%02X,X\n", operand1())
			incCount(2)
		case 0x55:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("EOR $%02X,X\n", operand1())
			incCount(2)
		case 0x56:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("LSR $%02X,X\n", operand1())
			incCount(2)
		case 0x57:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("RMB5 $%02X\n", operand1())
			incCount(2)
		case 0x61:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, opcode(), operand1())
			}
			fmt.Printf("ADC ($%02X,X)\n", operand1())
			incCount(2)
		case 0x62:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("RTN #$%02X\n", operand1())
			incCount(2)
		case 0x64:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("STZ $%02X\n", operand1())
			incCount(2)
		case 0x65:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("ADC $%02X\n", operand1())
			incCount(2)
		case 0x66:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("ROR $%02X\n", operand1())
			incCount(2)
		case 0x67:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("RMB6 $%02X\n", operand1())
			incCount(2)
		case 0x69:
			/*
				ADC - Add Memory to Accumulator with Carry
				Operation: A + M + C → A, C

				This instruction adds the value of memory and carry from the previous operation to the value of the
				accumulator and stores the result in the accumulator.

				This instruction affects the accumulator; sets the carry flag when the sum of a binary add exceeds
				255 or when the sum of a decimal add exceeds 99, otherwise carry is reset.

				The overflow flag is set when the sign or bit 7 is changed due to the result exceeding +127 or -128,
				otherwise overflow is reset.

				The negative flag is set if the accumulator result contains bit 7 on, otherwise the negative flag is reset.

				The zero flag is set if the accumulator result is 0, otherwise the zero flag is reset.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("ADC #$%02X\n", operand1())

			// Add operand 1 to accumulator
			A += operand1()
			// If accumulator>255 then set SR carry flag bit 0 to 1
			if A > 255 {
				setSRBitOn(0)
			} else {
				setSRBitOff(0)
			}
			// If accumulator>127 or accumulator<128 then set SR overflow flag bit 6 to 1 else set SR overflow flag bit 6 to 0
			if A > 127 || A < 128 {
				setSRBitOn(6)
			} else {
				setSRBitOff(6)
			}
			// If accumulator bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getABit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// If accumulator is 0 then set SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(2)
		case 0x70:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("BVS $%04X\n", fileposition+2+int(operand1()))
			incCount(2)
		case 0x71:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, opcode(), operand1())
			}
			fmt.Printf("ADC ($%02X),Y\n", operand1())
			incCount(2)
		case 0x72:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Indirect,Z)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("ADC ($%02X),Z\n", operand1())
			incCount(2)
		case 0x74:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("STZ $%02X,X\n", operand1())
			incCount(2)
		case 0x75:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("ADC $%02X,X\n", operand1())
			incCount(2)
		case 0x76:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("ROR $%02X,X\n", operand1())
			incCount(2)
		case 0x77:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("RMB7 $%02X\n", operand1())
			incCount(2)
		case 0x80:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("BRA $%04X\n", fileposition+2+int(operand1()))
			incCount(2)
		case 0x81:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, opcode(), operand1())
			}
			fmt.Printf("STA ($%02X,X)\n", operand1())
			incCount(2)
		case 0x82:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Stack Relative Indirect,Y)\n", PC, opcode(), operand1())
			}
			fmt.Printf("STA ($%02X,S),Y\n", operand1())
			incCount(2)
		case 0x84:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("STY $%02X\n", operand1())
			incCount(2)
		case 0x85:
			/*
				STA - Store Accumulator in Memory

				Operation: A → M

				This instruction transfers the contents of the accumulator to memory.

				This instruction affects none of the flags in the processor status register and does not affect the accumulator.
			*/

			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("STA $%02X\n", operand1())

			// Store contents of Accumulator in memory
			memory[operand1()] = A
			incCount(2)
		case 0x86:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("STX $%02X\n", operand1())
			incCount(2)
		case 0x87:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("SMB0 $%02X\n", operand1())
			incCount(2)
		case 0x89:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("BIT #$%02X\n", operand1())
			incCount(2)
		case 0x90:
			/*
				BCC - Branch on Carry Clear
				Operation: Branch on C = 0

				This instruction tests the state of the carry bit and takes a conditional branch if the carry bit is reset.

				It affects no flags or registers other than the program fileposition and then only if the C flag is not on.
			*/

			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("BCC $%02X\n", (fileposition+2+int(operand1()))&0xFF)

			// If carry flag bit zero of the status register is clear, then branch to the address specified by the operand.
			if getSRBit(0) == 0 {
				fileposition = (fileposition + 2) + int(operand1())
				PC += fileposition
			}
			incCount(0)
		case 0x91:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, opcode(), operand1())
			}
			fmt.Printf("STA ($%02X),Y\n", operand1())
			incCount(2)
		case 0x92:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Indirect,Z)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("STA ($%02X)\n", operand1())
			incCount(2)
		case 0x94:
			/*
				STY - Store Index Register Y In Memory
				Operation: Y → M

				Transfer the value of the Y register to the addressed memory location.

				STY does not affect any flags or registers in the microprocessor.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("STY $%02X,X\n", operand1())

			// Store contents of Y register in X indexed memory address
			memory[(operand1())+X] = Y
			incCount(2)
		case 0x95:
			/*
				STA - Store Accumulator in Memory
				Operation: A → M

				This instruction transfers the contents of the accumulator to memory.

				This instruction affects none of the flags in the processor status register and does not affect
				the accumulator.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("STA $%02X,X\n", operand1())

			// Store contents of Accumulator in X indexed memory
			memory[(operand1())+X] = A
			incCount(2)
		case 0x96:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Y)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("STX $%02X,Y\n", operand1())
			incCount(2)
		case 0x97:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("SMB1 $%02X\n", operand1())
			incCount(2)
		case 0xA0:
			/*
				LDY - Load Index Register Y From Memory
				Operation: M → Y

				Load the index register Y from memory.

				LDY does not affect the C or V flags,
				sets the N flag if the value loaded in bit 7 is a 1, otherwise resets N,
				sets Z flag if the loaded value is zero otherwise resets Z and only affects the Y register.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("LDY #$%02X\n", operand1())

			// Load the value of the operand1() into the Y register.
			Y = operand1()
			// If bit 7 of Y is set, set the SR negative flag else reset it to 0
			if getYBit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// If Y is zero, set the SR zero flag else reset it to 0
			if Y == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(2)
		case 0xA1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, opcode(), operand1())
			}
			fmt.Printf("LDA ($%02X,X)\n", operand1())
			incCount(2)
		case 0xA2:
			/*
				LDX - Load Index Register X From Memory
				Operation: M → X

				Load the index register X from memory.

				LDX does not affect the C or V flags; sets Z if the value loaded was zero, otherwise resets it;
				sets N if the value loaded in bit 7 is a 1; otherwise N is reset, and affects only the X register.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("LDX #$%02X\n", operand1())
			// Load the value of the operand1() into the X register.
			X = operand1()
			// If bit 7 of X is set, set the SR negative flag else reset it to 0
			if getXBit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// If X is zero, set the SR zero flag else reset it to 0
			if X == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(2)
		case 0xA3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("LDZ #$%02X\n", operand1())
			incCount(2)
		case 0xA4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("LDY $%02X\n", operand1())
			incCount(2)
		case 0xA5:
			/*
				LDA - Load Accumulator with Memory
				Operation: M → A

				When instruction LDA is executed by the microprocessor, data is transferred from memory to the
				accumulator and stored in the accumulator.

				LDA affects the contents of the accumulator, does not affect the carry or overflow flags;
				sets the zero flag if the accumulator is zero as a result of the LDA, otherwise resets the zero flag;
				sets the negative flag if bit 7 of the accumulator is a 1, otherwise resets the negative flag.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("LDA $%02X\n", operand1())

			// Load the accumulator with the value in the operand
			A = operand1()
			// If A is zero, set the zero flag else reset the zero flag
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}

			// If bit 7 of A is 1, set the negative flag else reset the negative flag
			if getABit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			incCount(2)
		case 0xA6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("LDX $%02X\n", operand1())
			incCount(2)
		case 0xA7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("SMB2 $%02X\n", operand1())
			incCount(2)
		case 0xA9:
			/*
				LDA - Load Accumulator with Memory
				Operation: M → A

				When instruction LDA is executed by the microprocessor, data is transferred from memory to the
				accumulator and stored in the accumulator.

				LDA affects the contents of the accumulator, does not affect the carry or overflow flags;
				sets the zero flag if the accumulator is zero as a result of the LDA, otherwise resets the zero flag;
				sets the negative flag if bit 7 of the accumulator is a 1, otherwise resets the negative flag.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("LDA #$%02X\n", operand1())

			// Load the accumulator with the value in the operand
			A = operand1()
			// If A is zero, set the SR Zero flag to 1 else set SR Zero flag to 0
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			// If bit 7 of accumulator is 1, set the SR negative flag to 1 else set the SR negative flag to 0
			if getABit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			incCount(2)
		case 0xB0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("BCS $%02X\n", (fileposition+2+int(operand1()))&0xFF)
			incCount(2)
		case 0xB1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, opcode(), operand1())
			}
			fmt.Printf("LDA ($%02X),Y\n", operand1())
			incCount(2)
		case 0xB2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Indirect)\n", PC, opcode(), operand1())
			}
			fmt.Printf("LDA ($%02X)\n", operand1())
			incCount(2)
		case 0xB4:
			/*
				LDY - Load Index Register Y From Memory
				Operation: M → Y

				Load the index register Y from memory.

				LDY does not affect the C or V flags, sets the N flag if the value loaded in bit 7 is a 1, otherwise resets N,
				sets Z flag if the loaded value is zero otherwise resets Z and only affects the Y register.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("LDY $%02X,X\n", operand1())

			// Load the Y register with the X indexed value in the operand
			Y = memory[int(operand1())+int(X)]
			// If bit 7 of Y is 1, set the SR negative flag bit 7 else reset the SR negative flag
			if getYBit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}

			// If Y is zero, set the SR zero flag else reset the SR zero flag
			if Y == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(2)

		case 0xB5:
			/*
				LDA - Load Accumulator with Memory
				Operation: M → A

				When instruction LDA is executed by the microprocessor, data is transferred from memory to the accumulator
				and stored in the accumulator.

				LDA affects the contents of the accumulator, does not affect the carry or overflow flags;
				sets the zero flag if the accumulator is zero as a result of the LDA, otherwise resets the zero flag;
				sets the negative flag if bit 7 of the accumulator is a 1, otherwise resets the negative flag.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("LDA $%02X,X\n", operand1())

			// Load the accumulator with the X indexed value in the operand
			A = memory[fileposition+1+int(X)]

			// If A is zero, set the zero flag else reset the zero flag
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}

			// If bit 7 of A is 1, set the negative flag else reset the negative flag
			if getABit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			incCount(2)
		case 0xB6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Y)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("LDX $%02X,Y\n", operand1())
			incCount(2)
		case 0xB7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("SMB3 $%02X\n", operand1())
			incCount(2)
		case 0xC0:
			/*
				CPY - Compare Index Register Y To Memory
				Operation: Y - M

				This instruction performs a two's complement subtraction between the index register Y and the
				specified memory location. The results of the subtraction are not stored anywhere. The instruction is
				strictly used to set the flags.

				CPY affects no registers in the microprocessor and also does not affect the overflow flag.

				If the value in the index register Y is equal to or greater than the value in the memory,
				the carry flag will be set, otherwise it will be cleared.

				If the results of the subtraction contain bit 7 on the N bit will be set, otherwise it will be cleared.

				If the value in the index register Y and the value in the memory are equal, the zero flag will be set,
				otherwise it will be cleared.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("CPY #$%02X\n", operand1())

			// Subtract operand from Y
			TempResult := Y - operand1()
			// If operand is greater than Y, set carry flag to 0
			if operand1() > Y {
				setSRBitOff(0)
			}
			// If bit 7 of TempResult is set, set N flag to 1
			if TempResult&0b10000000 == 0b10000000 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// If operand is equal to Y, set Z flag to 1 else set Zero flag to 0
			if operand1() == Y {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(2)
		case 0xC1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, opcode(), operand1())
			}
			fmt.Printf("CMP ($%02X,X)\n", operand1())
			incCount(2)
		case 0xC2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("CPZ #$%02X\n", operand1())
			incCount(2)
		case 0xC3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("DEW $%02X\n", operand1())
			incCount(2)
		case 0xC4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("CPY $%02X\n", operand1())
			incCount(2)
		case 0xC5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("CMP $%02X\n", operand1())
			incCount(2)
		case 0xC6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("DEC $%02X\n", operand1())
			incCount(2)
		case 0xC7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("SMB4 $%02X\n", operand1())
			incCount(2)
		case 0xC9:
			/*
				CMP - Compare Memory and Accumulator
				Operation: A - M

				This instruction subtracts the contents of memory from the contents of the accumulator.

				The use of the CMP affects the following flags: Z flag is set on an equal comparison, reset otherwise;
				the N flag is set or reset by the result bit 7,
				the carry flag is set when the value in memory is less than or equal to the accumulator,
				reset when it is greater than the accumulator.

				The accumulator is not affected.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("CMP #$%02X\n", operand1())

			// Compare memory and accumulator
			if operand1() == A {
				setSRBitOn(1)
				setSRBitOn(7)
			}
			// Set carry flag to true if A is greater than or equal to operand
			if A >= operand1() {
				setSRBitOn(0)
			}
			// Set carry flag to false if A is less than operand
			if A < operand1() {
				setSRBitOff(0)
			}
			// Set Z flag to false if A is not equal to operand
			if A != operand1() {
				setSRBitOff(1)
			}
			// Set N flag to true if A minus operand results in most significant bit being set
			if (A-operand1())&0b10000000 == 0b10000000 {
				setSRBitOn(7)
			}
			// Set N flag to false if A minus operand results in most significant bit being unset
			if (A-operand1())&0b10000000 == 0b00000000 {
				setSRBitOff(7)
			}
			incCount(2)
		case 0xD0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("BNE $%02X\n", (fileposition+2+int(operand1()))&0xFF)
			incCount(2)
		case 0xD1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, opcode(), operand1())
			}
			fmt.Printf("CMP ($%02X),Y\n", operand1())
			incCount(2)
		case 0xD2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect) Z)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("CMP ($%02X)\n", operand1())
			incCount(2)
		case 0xD4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("CPZ $%02x\n", operand1())
			incCount(2)
		case 0xD5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("CMP $%02X,X\n", operand1())
			incCount(2)
		case 0xD6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("DEC $%02X,X\n", operand1())
			incCount(2)
		case 0xD7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("SMB5 $%02X\n", operand1())
			incCount(2)
		case 0xE0:
			/*
				CPX - Compare Index Register X To Memory
				Operation: X - M

				This instruction subtracts the value of the addressed memory location from the content of index
				register X using the adder but does not store the result;
				therefore, its only use is to set the N, Z and C flags to allow for comparison between the index
				register X and the value in memory.

				The CPX instruction does not affect any register in the machine; it also does not affect the overflow flag.
				It causes the carry to be set on if the absolute value of the index register X is equal to or greater
				than the data from memory.
				If the value of the memory is greater than the content of the index register X, carry is reset.
				If the results of the subtraction contain a bit 7, then the N flag is set, if not, it is reset.
				If the value in memory is equal to the value in index register X, the Z flag is set, otherwise it is reset.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("CPX #$%02X\n", operand1())

			// Subtract operand from X
			TempResult := X - operand1()
			// If operand is greater than X, set carry flag to 0
			if operand1() > X {
				setSRBitOff(0)
			}
			// If bit 7 of TempResult is set, set N flag to 1
			if TempResult&0b10000000 == 0b10000000 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// If operand is equal to X, set Z flag to 1 else set Zero flag to 0
			if operand1() == X {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(2)
		case 0xE1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, opcode(), operand1())
			}
			fmt.Printf("SBC ($%02X,X)\n", operand1())
			incCount(2)
		case 0xE2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("LDA #$%02X\n", operand1())
			incCount(2)
		case 0xE3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("INW $%02X\n", operand1())
			incCount(2)
		case 0xE4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("CPX $%02X\n", operand1())
			incCount(2)
		case 0xE5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("SBC $%02X\n", operand1())
			incCount(2)
		case 0xE6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("INC $%02X\n", operand1())
			incCount(2)
		case 0xE7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("SMB6 $%02X\n", operand1())
			incCount(2)
		case 0xE9:
			/*
				SBC - Subtract Memory from Accumulator with Borrow
				Operation: A - M - ~C → A

				This instruction subtracts the value of memory and borrow from the value of the accumulator,
				using two's complement arithmetic, and stores the result in the accumulator.

				Borrow is defined as the carry flag complemented; therefore, a resultant carry flag indicates
				that a borrow has not occurred.

				This instruction affects the accumulator.
				The carry flag is set if the result is greater than or equal to 0.
				The carry flag is reset when the result is less than 0, indicating a borrow.
				The overflow flag is set when the result exceeds +127 or -127, otherwise it is reset.
				The negative flag is set if the result in the accumulator has bit 7 on, otherwise it is reset.
				The Z flag is set if the result in the accumulator is 0, otherwise it is reset.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("SBC #$%02X\n", operand1())

			// Update the accumulator
			A = A - operand1() - (1 - SR&1)
			// Set carry flag bit 0 if result is greater than or equal to 1
			if A >= 1 {
				setSRBitOn(0)
			} else {
				setSRBitOff(0)
			}
			// Set overflow flag bit 6 if accumulator is greater than 127 or less than -127
			if int(A) > 127 || int(A) < -127 {
				setSRBitOn(6)
			} else {
				setSRBitOff(6)
			}
			// If accumulator bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if getABit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// Set Z flag bit 1 if accumulator is 0 else set Z flag bit 1 to 0
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(2)
		case 0xF0:
			/*
				BEQ - Branch on Result Zero
				Operation: Branch on Z = 1

				This instruction could also be called "Branch on Equal."

				It takes a conditional branch whenever the Z flag is on or the previous result is equal to 0.

				BEQ does not affect any of the flags or registers other than the program fileposition and only then
				when the Z flag is set.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("BEQ $%02X\n", (fileposition+2+int(operand1()))&0xFF) // If SR zero flag 1 is set or the previous result is equal to 0, then branch to the address specified by the operand.
			if getSRBit(1) == 1 || A == 0 {
				fileposition = (fileposition + 2 + int(operand1())) & 0xFF
				PC += fileposition
				fmt.Printf("XXXXX")
			} else {
				incCount(2)
			}
		case 0xF1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, opcode(), operand1())
			}
			fmt.Printf("SBC ($%02X),Y\n", operand1())
			incCount(2)
		case 0xF2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect) Z)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("SBC ($%02X)\n", operand1())
			incCount(2)
		case 0xF5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("SBC $%02X,X\n", operand1())
			incCount(2)
		case 0xF6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("INC $%02X,X\n", operand1())
			incCount(2)
		case 0xF7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, opcode(), operand1())
			}
			fmt.Printf("SMB7 $%02X\n", operand1())
			incCount(2)
		}

		// 3 byte instructions with 2 operands
		switch opcode() {
		case 0x0C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("TSB $%02x%02x\n", operand2(), operand1())
			incCount(3)
		case 0x0D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("ORA $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x0E:
			/*
				ASL - Arithmetic Shift Left
				Operation: C ← /M7...M0/ ← 0

				The shift left instruction shifts either the accumulator or the address memory location
				1 bit to the left, with the bit 0 always being set to 0 and the the input bit 7 being stored in
				the carry flag.

				ASL either shifts the accumulator left 1 bit or is a read/modify/write instruction that affects only memory.

				The instruction does not affect the overflow bit,
				sets N equal to the result bit 7 (bit 6 in the input),
				sets Z flag if the result is equal to 0, otherwise resets Z
				and stores the input bit 7 in the carry flag.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("ASL $%02X%02X\n", operand2(), operand1())
			// Update temp var with the values of Operands 1 and 2
			temp := (int(operand2()) << 8) + int(operand1())
			// Shift left 1 bit
			temp <<= 1
			// Set SR negative flag 7 to 1 if the result bit 7 is 1
			if temp&0b10000000 == 0b10000000 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// Set SR zero flag 1 to 1 if the result is equal to 0, otherwise reset Zero flag to 0 and store bit 7 of temp in SR carry flag
			if temp == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
				if temp&0b10000000 == 0b10000000 {
					setSRBitOn(0)
				} else {
					setSRBitOff(0)
				}
			}
			// Update memory[temp] with the new value
			memory[temp] = byte(temp)
			incCount(3)
		case 0x0F:
			/*
				BBR0 - Branch on Bit 0 Reset
				Operation: Branch on M0 = 0

				This instruction tests the specified zero page location and branches if bit 0 is clear.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BBR0 $%02X, $%02X\n", operand1(), operand2())

			if memory[operand1()]&1 == 0 {
				fileposition = (fileposition + 3 + int(operand2())) & 0xFF
				PC = fileposition
			}
			incCount(3)
		case 0x13:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BPL $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x19:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("ORA $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0x1C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("TRB $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x1D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("ORA $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x1E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("ASL $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x1F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BBR1 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0x20:
			/*
				JSR - Jump To Subroutine
				Operation: PC + 2↓, [PC + 1] → PCL, [PC + 2] → PCH

				This instruction transfers control of the program counter to a subroutine location but leaves a
				return pointer on the stack to allow the user to return to perform the next instruction in the
				main program after the subroutine is complete.

				To accomplish this, JSR instruction stores the program counter address which points to the last byte
				of the jump instruction onto the stack using the stack pointer. The stack byte contains the
				program count high first, followed by program count low. The JSR then transfers the addresses following
				the jump instruction to the	program counter low and the program counter high, thereby directing the
				program to begin at that new address.

				The JSR instruction affects no flags, causes the stack pointer to be decremented by 2 and substitutes
				new values into the program fileposition low and the program fileposition high.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("JSR $%02X%02X\n", operand2(), operand1())
			// Store PC at address stored in SP
			memory[SP] = byte(PC)
			// Set PC to address stored in operand 1 and operand 2
			PC = int(operand1()) + int(operand2())<<8
			incCount(3)
		case 0x22:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t((Indirect) Z)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("JSR ($%02X%02X)\n", operand2(), operand1())
			incCount(0)
		case 0x23:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute X Indirect)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("JSR ($%02X%02X,X)\n", operand2(), operand1())
			incCount(3)
		case 0x2C:
			/*
				BIT - Test Bits in Memory with Accumulator
				Operation: A ∧ M, M7 → N, M6 → V

				This instruction performs an AND between a memory location and the accumulator but does not store the
				result of the AND into the accumulator.

				The bit instruction affects the N flag with
				N being set to the value of bit 7 of the memory being tested, the V flag with
				V being set equal to bit 6 of the memory being tested and
				Z being set by the result of the AND operation between the accumulator and the memory if the
				result is Zero, Z is reset otherwise.

				It does not affect the accumulator.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BIT $%02X%02X\n", operand2(), operand1())

			// Store the result of the AND between the accumulator and the operands in a temp var
			temp := A & memory[int(operand1())+int(operand2())<<8]
			// Set the SR Negative flag bit 7 to the value of bit 7 of temp
			if temp&0b10000000 == 0b10000000 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// Set the SR Overflow flag bit 6 to the value of bit 6 of temp
			if temp&0b01000000 == 0b01000000 {
				setSRBitOn(6)
			} else {
				setSRBitOff(6)
			}
			// If temp==0 then set the SR Zero flag bit 1 to the result of temp else set SR negative flag to 0
			if temp == 0 {
				// If bit 7 of temp is 1 then set SR negative flag to 1
				if temp&0b10000000 == 0b10000000 {
					setSRBitOn(7)
				}
			} else {
				setSRBitOff(7)
			}
			incCount(3)
		case 0x2D:
			/*
				AND - "AND" Memory with Accumulator
				Operation: A ∧ M → A

				The AND instruction transfer the accumulator and memory to the adder which performs a bit-by-bit
				AND operation and stores the result back in the accumulator.

				This instruction affects the accumulator;
				sets the zero flag if the result in the accumulator is 0, otherwise resets the zero flag;
				sets the negative flag if the result in the accumulator has bit 7 on, otherwise resets the negative flag.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("AND $%02X%02X\n", operand2(), operand1())

			// AND the accumulator with the value stored at the address stored in operand 1 and operand 2
			A &= memory[int(operand1())+int(operand2())<<8]
			// If A==0 set the SR zero flag to 1
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			// If bit 7 of A is 1 then set the SR negative flag to 1 else set negative flag to 0
			if getABit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			incCount(3)
		case 0x2E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("ROL $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x2F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BBR2 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0x33:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BMI $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x34:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BIT $%02X,X\n", operand1())
			incCount(3)
		case 0x35:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("AND $%02X,X\n", operand1())
			incCount(3)
		case 0x36:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("ROL $%02X,X\n", operand1())
			incCount(3)
		case 0x39:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("AND $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0x3C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("AND $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x3D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("AND $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x3E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("ROL $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x3F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BBR3 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0x4C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("JMP $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x4D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("EOR $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x4E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("LSR $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x4F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BBR4 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0x53:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BVC $%02X\n", operand1())
			incCount(3)
		case 0x59:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("EOE $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0x5D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("EOR $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x5E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("LSR $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x5F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BBR5 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0x63:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BSR $%02X\n", operand1())
			incCount(3)
		case 0x6C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute Indirect)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("JMP ($%02X%02X)\n", operand2(), operand1())
			incCount(3)
		case 0x6D:
			/*
				ADC - Add Memory to Accumulator with Carry
				Operation: A + M + C → A, C

				This instruction adds the value of memory and carry from the previous operation to the value of the
				accumulator and stores the result in the accumulator.

				This instruction affects the accumulator;
				sets the carry flag when the sum of a binary add exceeds 255 or when the sum of a decimal add exceeds 99
				otherwise carry is reset.

				The overflow flag is set when the sign or bit 7 is changed due to the result exceeding +127 or -128,
				otherwise overflow is reset.

				The negative flag is set if the accumulator result contains bit 7 on, otherwise the negative flag is reset.
				The zero flag is set if the accumulator result is 0, otherwise the zero flag is reset.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("ADC $%02X%02X\n", operand2(), operand1())
			// Read bit 7 of the accumulator into a temp var
			temp := getABit(7)
			// If SR carry bit is set, add 1 to A
			if getSRBit(0) == 1 {
				A++
			}
			// Add the value of operand
			A += operand1()
			// If bit 7 of accumulator is not equal to bit 7 of temp then set SR overflow flag bit 6 to 1
			if getABit(7) != (temp & 0b10000000) {
				setSRBitOn(6)
			}
			// If bit 7 of accumulator is 1 then set negative flag
			if getABit(7) == 1 {
				setSRBitOn(7)
			}
			// If accumulator is 0 then set zero flag else set SR zero flag to 0
			if A == 0 {
				setSRBitOn(1)
			} else {
				setSRBitOff(1)
			}
			incCount(3)
		case 0x6E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("ROR $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x6F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BBR6 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0x73:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BVS $%02X\n", operand1())
			incCount(3)
		case 0x79:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("ADC $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0x7C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X Indirect)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("JMP ($%02X%02X,X)\n", operand2(), operand1())
			incCount(3)
		case 0x7D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("ADC $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x7E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("ROR $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x7F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BBR7 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0x83:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BRA $%02X\n", operand1())
			incCount(3)
		case 0x8B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("STY $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x8C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("STY $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x8D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("STA $%04X\n", operand1()|uint8(int(operand2())<<8))
			incCount(3)
		case 0x8E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("STX $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x8F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(BBS0- Zero Page, Relative)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BBS0 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0x93:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BCC $%02X\n", operand1())
			incCount(3)
		case 0x99:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("STA $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0x9B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("STX $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0x9C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("STZ $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x9D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("STA $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x9E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("STZ $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x9F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BBS1 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0xAB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("LDZ $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0xAC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("LDY $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xAD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("LDA $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xAE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("LDX $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xAF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BBS2 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0xB3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BCS $%02X\n", operand1())
			incCount(3)
		case 0xB9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("LDA $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0xBB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("LDZ $%02X%02X,X\n", operand2(), operand1())
			incCount(1)
		case 0xBC:
			/*
				LDY - Load Index Register Y From Memory
				Operation: M → Y

				Load the index register Y from memory.

				LDY does not affect the C or V flags, sets the N flag if the value loaded in bit 7 is a 1, otherwise resets N,
				sets Z flag if the loaded value is zero otherwise resets Z and only affects the Y register.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("LDY $%02X%02X,X\n", operand2(), operand1())

			//  Set X to Operand 2 and Y to the X indexed value stored in operand 1
			X = operand2()
			Y = memory[int(operand1())+int(X)]

			// If bit 7 of Y is 1 then set bit 7 of SR to 1
			// else set bit 7 of SR to 0
			if getYBit(7) == 1 {
				setSRBitOn(7)
			} else {
				setSRBitOff(7)
			}
			// If value loaded to Y is 0 set bit 1 of SR to 0
			if Y == 0 {
				setSRBitOff(1)
			}
			incCount(3)
		case 0xBD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("LDA $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0xBE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("LDX $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0xBF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BBS3 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0xCB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("ASW $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xCC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("CPY $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xCD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("CMP $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xCE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("DEC $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xCF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BBS4 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0xD3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BNE $%02X\n", operand1())
			incCount(3)
		case 0xD9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("CMP $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0xDC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("CPZ $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0xDD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("CMP $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0xDE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("DEC $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0xDF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BBS5 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0xEB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("ROW $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xEC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("CPX $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xED:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("SBC $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xEE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("INC $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xEF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BBS6 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0xF3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BEQ $%02X\n", operand1())
			incCount(3)
		case 0xF4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Immediate (word))\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("PHW #$%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xF9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("SBC $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0xFC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("PHW #$%02X%02X\n", operand1(), operand2())
			incCount(3)
		case 0xFD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("SBC $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0xFE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("INC $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0xFF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, opcode(), operand1(), operand2())
			}
			fmt.Printf("BBS7 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		}
	}
}
