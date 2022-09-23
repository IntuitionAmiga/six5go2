package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	printHex     bool
	file         []byte
	fileposition = 0 //Byte position counter

	//CPURegisters and RAM
	A      byte        = 0x0000     //Accumulator
	X      byte        = 0x0000     //X register
	Y      byte        = 0x0000     //Y register		(76543210) SR Bit 5 is always set
	SR     byte        = 0b10100100 //Status Register	(NVEBDIZC)
	SP                 = 0x0100     //Stack Pointer
	PC                 = 0x0000     //Program Counter
	memory [65536]byte              //Memory
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

	// Read file
	file, _ = os.ReadFile(os.Args[1])

	fmt.Printf("USAGE   - six5go2 <target_filename> <entry_point> (Hex memory address) <hex> (Print hex values above each instruction) \n")
	fmt.Printf("EXAMPLE - six5go2 cbmbasic35.rom 0800 hex\n\n")
	fmt.Printf("Length of file %s is %v ($%04X) bytes\n\n", os.Args[1], len(file), len(file))

	fmt.Printf("Size of addressable memory is %v ($%04X) bytes\n\n", len(memory), len(memory))

	copy(memory[:], file)
	printMachineState()
	execute(string(file))
}

func currentByte() byte {
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
}
func printMachineState() {
	fmt.Printf("A=$%02X X=$%02X Y=$%02X SR=%08b (NVEBDIZC) SP=$%08X PC=$%04X\n\n", A, X, Y, SR, SP, PC)
}

func execute(file string) {
	PC += fileposition
	if printHex {
		fmt.Printf(" * = $%04X\n\n", PC)
	}
	for fileposition = 0; fileposition < len(file); {
		//PC += fileposition
		// 1 byte instructions with no operands
		switch currentByte() {
		case 0x00:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
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
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("CLE\n")

			//Set bit 5 of SR to 0
			SR |= 0 << 5

			incCount(1)
			printMachineState()
		case 0x03:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
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
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}

			//Push SR to stack
			memory[SP] = SR
			//Decrement the stack pointer by 1 byte
			SP--

			fmt.Printf("PHP\n")
			incCount(1)
			printMachineState()
		case 0x0A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, currentByte())
			}
			fmt.Printf("ASL\n")
			incCount(1)
		case 0x0B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
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
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("CLC\n")

			//Set SR carry flag bit 0 to 0
			SR |= 0 << 0
			incCount(1)
			printMachineState()
		case 0x1A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, currentByte())
			}
			fmt.Printf("INC\n")
			incCount(1)
		case 0x1B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("INZ\n")
			incCount(1)
		case 0x28:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("PLP\n")
			incCount(1)
		case 0x2A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, currentByte())
			}
			fmt.Printf("ROL\n")
			incCount(1)
		case 0x2B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
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
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, currentByte())
			}
			fmt.Printf("SEC\n")

			//Set SR carry flag bit 0 to 1
			SR |= 1 << 0
			incCount(1)
			printMachineState()
		case 0x3A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, currentByte())
			}
			fmt.Printf("DEC\n")
			incCount(1)
		case 0x3B:
			//NOP

			incCount(1)
			printMachineState()
		case 0x40:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("RTI\n")
			incCount(1)
		case 0x42:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, currentByte())
			}
			fmt.Printf("NEG\n")
			incCount(1)
		case 0x43:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, currentByte())
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
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("PHA\n")

			//Update memory address pointed to by SP with value stored in accumulator
			memory[SP] = A
			//Decrement the stack pointer by 1 byte
			SP--

			incCount(1)
			printMachineState()
		case 0x4A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, currentByte())
			}
			fmt.Printf("LSR\n")
			incCount(1)
		case 0x4B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
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
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("CLI\n")

			//Set SR interrupt disable bit 2 to 0
			SR |= 0 << 2

			incCount(1)
			printMachineState()
		case 0x5A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("PHY\n")
			incCount(1)
		case 0x5B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
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
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("AUG\n")

			incCount(1)
			printMachineState()
		case 0x60:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("RTS\n")
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
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("PLA\n")

			//Increment the stack pointer by 1 byte
			SP++
			//Update accumulator with value stored in memory address pointed to by SP
			A = memory[SP]
			//If bit 7 of accumulator is set, set negative SR flag else set negative SR flag to 0
			if A&1<<7 == 1 {
				SR |= 1 << 7
			} else {
				SR |= 1 << 7
			}
			//If accumulator is 0, set zero SR flag else set zero SR flag to 0
			if A == 0 {
				SR |= 1 << 1
			} else {
				SR |= 0 << 1
			}

			incCount(1)
			printMachineState()
		case 0x6A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, currentByte())
			}
			fmt.Printf("ROR\n")
			incCount(1)
		case 0x6B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("TZA\n")
			incCount(1)
		case 0x78:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("SEI\n")
			incCount(1)
		case 0x7A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("PLY\n")
			incCount(1)
		case 0x7B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(TBA - Absolute,Y)\n", PC, currentByte())
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
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("DEY\n")

			//Decrement the  Y register by 1
			Y--
			//If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if Y&1<<7 == 1 {
				//Set bit 7 of SR to 1
				SR = SR | 1<<7
			} else {
				//Set bit 7 of SR to 0
				SR ^= 1 << 7
			}

			incCount(1)
			printMachineState()
		case 0x8A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("TXA\n")
			incCount(1)
		case 0x98:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("TYA\n")
			incCount(1)
		case 0x9A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("TXS\n")
			incCount(1)
		case 0xA8:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("TAY\n")
			incCount(1)
		case 0xAA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("TAX\n")
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
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("CLV\n")

			//Set SR overflow flag bit 6 to 0
			SR |= 0 << 6

			incCount(1)
			printMachineState()
		case 0xBA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("TSX\n")
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
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("INY\n")

			//Increment the  Y register by 1
			Y++
			//If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if Y&1<<7 == 1 {
				//Set bit 7 of SR to 1
				SR = SR | 1<<7
			} else {
				//Set bit 7 of SR to 0
				SR ^= 1 << 7
			}

			//If Y register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
			if Y == 0 {
				//Set bit 1 of SR to 1
				SR = SR | 1<<1
			} else {
				//Set bit 1 of SR to 0
				SR ^= 1 << 1
			}

			incCount(1)
			printMachineState()
		case 0xCA:
			/*
				DEX - Decrement Index Register X By One
				Operation: X - 1 → X

				This instruction subtracts one from the current value of the index register X and stores the result
				in the index register X.

				DEX does not affect the carry or overflow flag, it sets the N flag if it has bit 7 on as a result
				of the decrement, otherwise it resets the N flag;
				sets the Z flag if X is a 0 as a result of the decrement, otherwise it resets the Z flag.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("DEX\n")

			//Decrement the X register by 1
			X--
			//If X register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if X&1<<7 == 1 {
				//Set bit 7 of SR to 1
				SR = SR | 1<<7
			} else {
				//Set bit 7 of SR to 0
				SR ^= 1 << 7
			}

			incCount(1)
			printMachineState()
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
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("CLD\n")

			//Set SR decimal mode flag to 0

			incCount(1)
			printMachineState()
		case 0xDA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("PHX\n")
			incCount(1)
		case 0xDB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
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
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("INX\n")

			//If X==0xFF then X=0 else increment the X by 1
			if X == 0xFF {
				X = 0
				//Set SR zero flag bit 1 to 1
				SR |= 1 << 1
			} else {
				X++
				//Set SR zero flag bit 1 to 0
				SR ^= 1 << 1
			}

			//If X bit 7 is 1, set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if X&1<<7 == 1 {
				//Set bit 7 of SR to 1
				SR = SR | 1<<7
			} else {
				//Set bit 7 of SR to 0
				SR ^= 1 << 7
			}

			incCount(1)
			printMachineState()
		case 0xEA:
			/*
				NOP - No Operation
				Operation: No operation
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("NOP\n")

			incCount(1)
			printMachineState()
		case 0xF8:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("SED\n")
			incCount(1)
		case 0xFA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("PLX\n")
			incCount(1)
		case 0xFB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, currentByte())
			}
			fmt.Printf("PLZ\n")
			incCount(1)
		}

		//2 byte instructions with 1 operand
		switch currentByte() {
		case 0x01:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page Indirect))\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ORA ($%02x,X)\n", operand1())
			incCount(2)
		case 0x04:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("TSB $%02x\n", operand1())
			incCount(2)
		case 0x05:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ORA $%02x\n", operand1())
			incCount(2)
		case 0x06:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ASL $%02x\n", operand1())
			incCount(2)
		case 0x07:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("RMB0 $%02x\n", operand1())
			incCount(2)
		case 0x09:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ORA #$%02x\n", operand1())
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
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("BPL $%02X\n", (fileposition+2+int(operand1()))&0xFF)

			//If SR negative bit 7 is 0, then branch
			if SR&7 == 1 {

				fileposition = (fileposition + 2 + int(operand1())) & 0xFF
				PC += fileposition
			}

			incCount(0)
			printMachineState()
		case 0x11:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page),Indirect Y)\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ORA ($%02x),Y\n", operand1())
			incCount(2)
		case 0x12:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page),Indirect Z)\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ORA ($%02x),Z\n", operand1())
			incCount(2)
		case 0x14:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("TRB $%02x\n", operand1())
			incCount(2)
		case 0x15:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ORA $%02x,X\n", operand1())
			incCount(2)
		case 0x16:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(ASL - Zero Page,X)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ASL $%02X,X\n", operand1())
			incCount(2)
		case 0x17:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("RMB1 $%02X\n", operand1())
			incCount(2)
		case 0x21:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page Indirect))\n", PC, currentByte(), operand1())
			}
			fmt.Printf("AND ($%02X,X)\n", operand1())
			incCount(2)
		case 0x24:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("BIT $%02X\n", operand1())
			incCount(2)
		case 0x25:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("AND $%02X\n", operand1())
			incCount(2)
		case 0x26:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ROL $%02X\n", operand1())
			incCount(2)
		case 0x27:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("RMB2 $%02X\n", operand1())
			incCount(2)
		case 0x29:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("AND #$%02X\n", operand1())
			incCount(2)
		case 0x30:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("BMI $%02X\n", (fileposition+2+int(operand1()))&0xFF)
			incCount(2)
		case 0x31:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, currentByte(), operand1())
			}
			fmt.Printf("AND ($%02X),Y\n", operand1())
			incCount(2)
		case 0x32:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect),Z)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("AND ($%02X),Z\n", operand1())
			incCount(2)
		case 0x34:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("BIT $%02X,X\n", operand1())
			incCount(2)
		case 0x35:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("AND $%02X,X\n", operand1())
			incCount(2)
		case 0x36:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ROL $%02X,X\n", operand1())
			incCount(2)
		case 0x37:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\n", PC, currentByte(), operand1())
			}
			fmt.Printf("RMB3 $%02X\n", operand1())
			incCount(2)
		case 0x41:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page, Indirect))\n", PC, currentByte(), operand1())
			}
			fmt.Printf("EOR ($%02X,X)\n", operand1())
			incCount(2)
		case 0x44:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ASR $%02X\n", operand1())
			incCount(2)
		case 0x45:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("EOR $%02X\n", operand1())
			incCount(2)
		case 0x46:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("LSR $%02X\n", operand1())
			incCount(2)
		case 0x47:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("RMB4 $%02X\n", operand1())
			incCount(2)
		case 0x49:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("EOR #$%02X\n", operand1())
			incCount(2)
		case 0x50:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("BVC $%02X\n", fileposition+2+int(operand1()))
			incCount(2)
		case 0x51:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, currentByte(), operand1())
			}
			fmt.Printf("EOR ($%02X),Y\n", operand1())
			incCount(2)
		case 0x52:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect),Z)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("EOR ($%02X),Z\n", operand1())
			incCount(2)
		case 0x54:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ASR $%02X,X\n", operand1())
			incCount(2)
		case 0x55:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("EOR $%02X,X\n", operand1())
			incCount(2)
		case 0x56:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("LSR $%02X,X\n", operand1())
			incCount(2)
		case 0x57:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("RMB5 $%02X\n", operand1())
			incCount(2)
		case 0x61:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ADC ($%02X,X)\n", operand1())
			incCount(2)
		case 0x62:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("RTN #$%02X\n", operand1())
			incCount(2)
		case 0x64:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("STZ $%02X\n", operand1())
			incCount(2)
		case 0x65:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ADC $%02X\n", operand1())
			incCount(2)
		case 0x66:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ROR $%02X\n", operand1())
			incCount(2)
		case 0x67:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
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
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ADC #$%02X\n", operand1())

			//Add operand 1 to accumulator
			A += operand1()
			//If accumulator>255 then set SR carry flag bit 0 to 1
			if A > 255 {
				SR |= 1 << 0
			} else {
				SR |= 0 << 0
			}
			//If accumulator>127 or accumulator<128 then set SR overflow flag bit 6 to 1 else set SR overflow flag bit 6 to 0
			if A > 127 || A < 128 {
				SR |= 1 << 6
			} else {
				SR |= 0 << 6
			}
			//If accumulator bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if A&1 == 1 {
				SR |= 1 << 7
			} else {
				SR |= 0 << 7
			}
			incCount(2)
			printMachineState()
		case 0x70:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("BVS $%04X\n", fileposition+2+int(operand1()))
			incCount(2)
		case 0x71:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ADC ($%02X),Y\n", operand1())
			incCount(2)
		case 0x72:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Indirect,Z)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ADC ($%02X),Z\n", operand1())
			incCount(2)
		case 0x74:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("STZ $%02X,X\n", operand1())
			incCount(2)
		case 0x75:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ADC $%02X,X\n", operand1())
			incCount(2)
		case 0x76:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("ROR $%02X,X\n", operand1())
			incCount(2)
		case 0x77:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("RMB7 $%02X\n", operand1())
			incCount(2)
		case 0x80:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("BRA $%04X\n", fileposition+2+int(operand1()))
			incCount(2)
		case 0x81:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, currentByte(), operand1())
			}
			fmt.Printf("STA ($%02X,X)\n", operand1())
			incCount(2)
		case 0x82:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Stack Relative Indirect,Y)\n", PC, currentByte(), operand1())
			}
			fmt.Printf("STA ($%02X,S),Y\n", operand1())
			incCount(2)
		case 0x84:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
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
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("STA $%02X\n", operand1())

			//Store contents of Accumulator in memory
			memory[operand1()] = A
			//fmt.Printf("Address[$%02X] = $%02X\n", operand1(), memory[operand1()])
			incCount(2)
			printMachineState()
		case 0x86:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("STX $%02X\n", operand1())
			incCount(2)
		case 0x87:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("SMB0 $%02X\n", operand1())
			incCount(2)
		case 0x89:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, currentByte(), operand1())
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
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("BCC $%02X\n", (fileposition+2+int(operand1()))&0xFF)

			//If carry flag/bit zero of the status register is clear, then branch to the address specified by the operand.
			if SR&1 == 0 {
				fileposition = (fileposition + 2) + int(operand1()) // & 0xFFFF
				PC += fileposition
			}

			incCount(0)
			printMachineState()
		case 0x91:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, currentByte(), operand1())
			}
			fmt.Printf("STA ($%02X),Y\n", operand1())
			incCount(2)
		case 0x92:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Indirect,Z)\t\n", PC, currentByte(), operand1())
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
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("STY $%02X,X\n", operand1())

			//Store contents of Y register in X indexed memory address
			memory[operand1()+X] = Y

			incCount(2)
			printMachineState()
		case 0x95:
			/*
				STA - Store Accumulator in Memory
				Operation: A → M

				This instruction transfers the contents of the accumulator to memory.

				This instruction affects none of the flags in the processor status register and does not affect
				the accumulator.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("STA $%02X,X\n", operand1())

			//Store contents of Accumulator in X indexed memory
			memory[operand1()+X] = A

			incCount(2)
			printMachineState()
		case 0x96:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Y)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("STX $%02X,Y\n", operand1())
			incCount(2)
		case 0x97:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("SMB1 $%02X\n", operand1())
			incCount(2)
		case 0xA0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("LDY #$%02X\n", operand1())
			incCount(2)
		case 0xA1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, currentByte(), operand1())
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
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("LDX #$%02X\n", operand1())

			//Load the value of the operand1() into the X register.
			X = operand1()

			//If accumulator is zero, set the SR zero flag
			if X == 0 {
				SR |= 1 << 1
			} else {
				SR |= 0 << 1
			}
			//If bit 7 of the accumulator is set, set the SR negative flag
			if A&1 == 1 {
				SR |= 1 << 7
			} else {
				SR |= 0 << 7
			}

			incCount(2)
			printMachineState()
		case 0xA3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("LDZ #$%02X\n", operand1())
			incCount(2)
		case 0xA4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
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
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("LDA $%02X\n", operand1())

			//Load the accumulator with the value in the operand
			A = operand1()
			//If A is zero, set the zero flag else reset the zero flag
			if A == 0 {
				//Set bit 1 to 1
				SR |= 1 << 1
			} else {
				//SR = 0b00100000
				//Set SR bit 1 to 0
				SR ^= 0 << 1
				//SR |= 0 << 1
			}

			//If bit 7 of A is 1, set the negative flag else reset the negative flag
			if A&1 == 1 {
				SR |= 1 << 7
			} else {
				SR |= 0 << 7
			}

			incCount(2)
			printMachineState()
		case 0xA6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("LDX $%02X\n", operand1())
			incCount(2)
		case 0xA7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("SMB2 $%02X\n", operand1())
			incCount(2)
		case 0xA9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("LDA #$%02X\n", operand1())
			incCount(2)
		case 0xB0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("BCS $%02X\n", (fileposition+2+int(operand1()))&0xFF)
			incCount(2)
		case 0xB1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, currentByte(), operand1())
			}
			fmt.Printf("LDA ($%02X),Y\n", operand1())
			incCount(2)
		case 0xB2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Indirect)\n", PC, currentByte(), operand1())
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
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("LDY $%02X,X\n", operand1())

			//Load the Y register with the X indexed value in the operand
			Y = memory[fileposition+1+int(X)]
			//If bit 7 of Y is 1, set the SR negative flag bit 7 else reset the SR negative flag
			if Y&1 == 1 {
				SR |= 1 << 7
			} else {
				SR |= 0 << 7
			}

			//If Y is zero, set the SR zero flag else reset the SR zero flag
			if Y == 0 {
				SR |= 1 << 1
			} else {
				SR |= 0 << 1
			}
			incCount(2)
			printMachineState()
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
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("LDA $%02X,X\n", operand1())

			//Load the accumulator with the X indexed value in the operand
			A = memory[fileposition+1+int(X)]

			//If A is zero, set the zero flag else reset the zero flag
			if A == 0 {
				//SR = 0b00100010
				//Set bit 1 to 1
				SR |= 1 << 1
			} else {
				//SR = 0b00100000
				//Set bit 1 to 0
				SR |= 0 << 1
			}

			//If bit 7 of A is 1, set the negative flag else reset the negative flag
			if A&0b10000000 != 0 {
				//SR = 0b00100001
				//Set bit 7 to 1
				SR |= 1 << 7
			} else {
				//SR = 0b00100000
				//Set bit 7 to 0
				SR |= 0 << 7
			}

			incCount(2)
			printMachineState()
		case 0xB6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Y)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("LDX $%02X,Y\n", operand1())
			incCount(2)
		case 0xB7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("SMB3 $%02X\n", operand1())
			incCount(2)
		case 0xC0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("CPY #$%02X\n", operand1())
			incCount(2)
		case 0xC1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, currentByte(), operand1())
			}
			fmt.Printf("CMP ($%02X,X)\n", operand1())
			incCount(2)
		case 0xC2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("CPZ #$%02X\n", operand1())
			incCount(2)
		case 0xC3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("DEW $%02X\n", operand1())
			incCount(2)
		case 0xC4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("CPY $%02X\n", operand1())
			incCount(2)
		case 0xC5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("CMP $%02X\n", operand1())
			incCount(2)
		case 0xC6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("DEC $%02X\n", operand1())
			incCount(2)
		case 0xC7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
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
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("CMP #$%02X\n", operand1())

			//Compare memory and accumulator
			if operand1() == A {
				//Set Z flag and negative flag to true
				//Set bit 1 to true
				SR |= 0 << 1
				//SR = 0b00100010
			}
			//Set carry flag to true if A is greater than or equal to operand
			if A >= operand1() {
				//Set bit 0 to true
				SR |= 1 << 0
				//SR = 0b00100001
			}
			//Set carry flag to false if A is less than operand
			if A < operand1() {
				//Set bit zero to false
				SR |= 0 << 0
				//SR = 0b00100000
			}
			//Set Z flag to false if A is not equal to operand
			if A != operand1() {
				//Set bit 1 to false
				SR |= 0 << 1
				//SR = 0b00100000
			}
			//Set N flag to true if A minus operand results in most significant bit being set
			if (A-operand1())&0b10000000 == 0b10000000 {
				//Set bit 7 to true
				SR |= 1 << 7
				//SR = 0b10100000
			}
			//Set N flag to false if A minus operand results in most significant bit being unset
			if (A-operand1())&0b10000000 == 0b00000000 {
				//Set bit 7 to false
				SR |= 0 << 7
				//SR = 0b00100000
			}
			incCount(2)
			printMachineState()
		case 0xD0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("BNE $%02X\n", (fileposition+2+int(operand1()))&0xFF)
			incCount(2)
		case 0xD1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, currentByte(), operand1())
			}
			fmt.Printf("CMP ($%02X),Y\n", operand1())
			incCount(2)
		case 0xD2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect) Z)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("CMP ($%02X)\n", operand1())
			incCount(2)
		case 0xD4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("CPZ $%02x\n", operand1())
			incCount(2)
		case 0xD5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("CMP $%02X,X\n", operand1())
			incCount(2)
		case 0xD6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("DEC $%02X,X\n", operand1())
			incCount(2)
		case 0xD7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("SMB5 $%02X\n", operand1())
			incCount(2)
		case 0xE0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("CPX #$%02X\n", operand1())
			incCount(2)
		case 0xE1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, currentByte(), operand1())
			}
			fmt.Printf("SBC ($%02X,X)\n", operand1())
			incCount(2)
		case 0xE2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("LDA #$%02X\n", operand1())
			incCount(2)
		case 0xE3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("INW $%02X\n", operand1())
			incCount(2)
		case 0xE4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("CPX $%02X\n", operand1())
			incCount(2)
		case 0xE5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("SBC $%02X\n", operand1())
			incCount(2)
		case 0xE6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("INC $%02X\n", operand1())
			incCount(2)
		case 0xE7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
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
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("SBC #$%02X\n", operand1())

			//Update the accumulator
			A = A - operand1() - (1 - SR&1)
			//Set carry flag bit 0 if result is greater than or equal to 1
			if A >= 1 {
				SR |= 1 << 0
			} else {
				SR |= 0 << 0
			}
			//Set overflow flag bit 6 if accumulator is greater than 127 or less than -127
			if int(A) > 127 || int(A) < -127 {
				SR |= 1 << 6
			} else {
				SR |= 0 << 6
			}
			//If accumulator bit 7 is 1 then set SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
			if A&1 == 1 {
				SR |= 1 << 7
			} else {
				SR |= 0 << 7
			}
			//Set Z flag bit 1 if accumulator is 0 else set Z flag bit 1 to 0
			if A == 0 {
				SR |= 1 << 1
			} else {
				SR |= 0 << 1
			}
			incCount(2)
			printMachineState()
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
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("BEQ $%02X\n", (fileposition+2+int(operand1()))&0xFF)

			//If SR zero flag 1 is set or the previous result is equal to 0, then branch to the address specified by the operand.
			if SR&1 == 1 || A == 0 {
				fileposition = (fileposition + 2 + int(operand1())) & 0xFF
				PC += fileposition
				fmt.Printf("XXXXX")
			} else {
				incCount(2)
				printMachineState()
			}

		case 0xF1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, currentByte(), operand1())
			}
			fmt.Printf("SBC ($%02X),Y\n", operand1())
			incCount(2)
		case 0xF2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect) Z)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("SBC ($%02X)\n", operand1())
			incCount(2)
		case 0xF5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("SBC $%02X,X\n", operand1())
			incCount(2)
		case 0xF6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("INC $%02X,X\n", operand1())
			incCount(2)
		case 0xF7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, currentByte(), operand1())
			}
			fmt.Printf("SMB7 $%02X\n", operand1())
			incCount(2)
		}

		//3 byte instructions with 2 operands
		switch currentByte() {
		case 0x0C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("TSB $%02x%02x\n", operand2(), operand1())
			incCount(3)
		case 0x0D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("ORA $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x0E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("ASL $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x0F:
			/*
				BBR0 - Branch on Bit 0 Reset
				Operation: Branch on M0 = 0

				This instruction tests the specified zero page location and branches if bit 0 is clear.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BBR0 $%02X, $%02X\n", operand1(), operand2())

			if memory[operand1()]&1 == 0 {
				fileposition = (fileposition + 3 + int(operand2())) & 0xFF
				PC = fileposition
			}

			incCount(3)
			printMachineState()
		case 0x13:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BPL $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x19:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("ORA $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0x1C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("TRB $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x1D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("ORA $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x1E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("ASL $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x1F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, currentByte(), operand1(), operand2())
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
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("JSR $%02X%02X\n", operand2(), operand1())
			//Store PC at address stored in SP
			memory[SP] = byte(PC)
			//Set PC to address stored in operand 1 and operand 2
			PC = int(operand1()) + int(operand2())<<8

			incCount(3)
			printMachineState()
		case 0x22:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t((Indirect) Z)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("JSR ($%02X%02X)\n", operand2(), operand1())
			incCount(0)
		case 0x23:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute X Indirect)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("JSR ($%02X%02X,X)\n", operand2(), operand1())
			incCount(3)
		case 0x2C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BIT $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x2D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("AND $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x2E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("ROL $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x2F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BBR2 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0x33:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BMI $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x34:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BIT $%02X,X\n", operand1())
			incCount(3)
		case 0x35:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("AND $%02X,X\n", operand1())
			incCount(3)
		case 0x36:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("ROL $%02X,X\n", operand1())
			incCount(3)
		case 0x39:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("AND $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0x3C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("AND $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x3D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("AND $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x3E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("ROL $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x3F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BBR3 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0x4C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("JMP $%02X%02X\n", operand2(), operand1())

			incCount(3)
		case 0x4D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("EOR $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x4E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("LSR $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x4F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BBR4 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0x53:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BVC $%02X\n", operand1())
			incCount(3)
		case 0x59:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("EOE $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0x5D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("EOR $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x5E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("LSR $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x5F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BBR5 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0x63:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BSR $%02X\n", operand1())
			incCount(3)
		case 0x6C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute Indirect)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("JMP ($%02X%02X)\n", operand2(), operand1())
			incCount(3)
		case 0x6D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("ADC $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x6E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("ROR $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x6F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BBR6 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0x73:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BVS $%02X\n", operand1())
			incCount(3)
		case 0x79:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("ADC $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0x7C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X Indirect)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("JMP ($%02X%02X,X)\n", operand2(), operand1())
			incCount(3)
		case 0x7D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("ADC $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x7E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("ROR $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x7F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BBR7 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0x83:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BRA $%02X\n", operand1())
			incCount(3)
		case 0x8B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("STY $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x8C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("STY $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x8D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("STA $%04X\n", operand1()|uint8(int(operand2())<<8))
			incCount(3)
		case 0x8E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("STX $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x8F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(BBS0- Zero Page, Relative)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BBS0 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0x93:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BCC $%02X\n", operand1())
			incCount(3)
		case 0x99:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("STA $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0x9B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("STX $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0x9C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("STZ $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0x9D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("STA $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x9E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("STZ $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0x9F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BBS1 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0xAB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("LDZ $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0xAC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("LDY $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xAD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("LDA $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xAE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("LDX $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xAF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BBS2 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0xB3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BCS $%02X\n", operand1())
			incCount(3)
		case 0xB9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("LDA $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0xBB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
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
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("LDY $%02X%02X,X\n", operand2(), operand1())

			// Set X to Operand 2 and Y to the X indexed value stored in operand 1
			X = operand2()
			Y = memory[int(operand1())+int(X)]

			//If bit 7 of Y is 1 then set bit 7 of SR to 1
			//else set bit 7 of SR to 0
			if Y&1 == 1 {
				SR |= 1 << 7
			} else {
				SR |= 0 << 7
			}

			//If value loaded to Y is 0 set bit 1 of SR to 0
			if Y == 0 {
				SR |= 0 << 1
			}
			incCount(3)
			printMachineState()
		case 0xBD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("LDA $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0xBE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("LDX $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0xBF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BBS3 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0xCB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("ASW $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xCC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("CPY $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xCD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("CMP $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xCE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("DEC $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xCF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BBS4 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0xD3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BNE $%02X\n", operand1())
			incCount(3)
		case 0xD9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("CMP $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0xDC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("CPZ $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0xDD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("CMP $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0xDE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("DEC $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0xDF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BBS5 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0xEB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("ROW $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xEC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("CPX $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xED:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("SBC $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xEE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("INC $%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xEF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BBS6 $%02X, $%02X\n", operand1(), operand2())
			incCount(3)
		case 0xF3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BEQ $%02X\n", operand1())
			incCount(3)
		case 0xF4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Immediate (word))\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("PHW #$%02X%02X\n", operand2(), operand1())
			incCount(3)
		case 0xF9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("SBC $%02X%02X,Y\n", operand2(), operand1())
			incCount(3)
		case 0xFC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("PHW #$%02X%02X\n", operand1(), operand2())
			incCount(3)
		case 0xFD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("SBC $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0xFE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("INC $%02X%02X,X\n", operand2(), operand1())
			incCount(3)
		case 0xFF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, currentByte(), operand1(), operand2())
			}
			fmt.Printf("BBS7 $%02X, $%02X\n", operand1(), operand2())
			//incCount(3)
			incCount(3)
		}
	}
}
