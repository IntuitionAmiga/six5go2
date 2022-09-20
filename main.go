package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	printHex bool
	file     []byte
	counter  = 0 //Byte position Counter

	// CPURegisters and RAM
	A      byte                     //Accumulator
	X      byte                     //X register
	Y      byte                     //Y register		(76543210) SR Bit 5 is always set
	SR     byte        = 0b00100000 //Status Register	(NVEBDIZC)
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
	//disassemble(string(file))
	execute(string(file))

}

func incCounter(amount int) {
	if counter+amount < len(file)-1 {
		counter += amount
	} else {
		os.Exit(0)
	}
	PC++
}
func printMachineState() {
	fmt.Printf("A=$%02X X=$%02X Y=$%02X SR=%08b (NVEBDIZC) SP=$%08X PC=$%04X\n\n", A, X, Y, SR, SP, PC)
}

/*
	func disassemble(file string) {
		PC += counter
		if printHex {
			fmt.Printf(" * = $%04X\n\n", PC)
		}
		for counter = 0; counter < len(file); PC++ {

			// 1 byte instructions with no operands
			switch file[counter] {
			case 0x00:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("BRK\n")
				incCounter(1)
			case 0x02:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("CLE\n")
				incCounter(1)
			case 0x03:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("SEE\n")
				incCounter(1)
			case 0x08:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("PHP\n")
				incCounter(1)
			case 0x0A:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, file[counter])
				}
				fmt.Printf("ASL\n")
				incCounter(1)
			case 0x0B:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("TSY\n")
				incCounter(1)
			case 0x18:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("CLC\n")
				incCounter(1)
			case 0x1A:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, file[counter])
				}
				fmt.Printf("INC\n")
				incCounter(1)
			case 0x1B:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("INZ\n")
				incCounter(1)
			case 0x28:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("PLP\n")
				incCounter(1)
			case 0x2A:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, file[counter])
				}
				fmt.Printf("ROL\n")
				incCounter(1)
			case 0x2B:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("TYS\n")
				incCounter(1)
			case 0x38:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, file[counter])
				}
				fmt.Printf("SEC\n")
				incCounter(1)
			case 0x3A:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, file[counter])
				}
				fmt.Printf("DEC\n")
				incCounter(1)
			case 0x3B:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("DEZ\n")
				incCounter(1)
			case 0x40:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("RTI\n")
				incCounter(1)
			case 0x42:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, file[counter])
				}
				fmt.Printf("NEG\n")
				incCounter(1)
			case 0x43:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, file[counter])
				}
				fmt.Printf("ASR\n")
				incCounter(1)
			case 0x48:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("PHA\n")
				//incCounter(1)
				incCounter(1)
			case 0x4A:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, file[counter])
				}
				fmt.Printf("LSR\n")
				incCounter(1)
			case 0x4B:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("TAZ\n")
				incCounter(1)
			case 0x58:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("CLI\n")
				incCounter(1)
			case 0x5A:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("PHY\n")
				incCounter(1)
			case 0x5B:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("TAB\n")
				incCounter(1)
			case 0x60:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("RTS\n")
				incCounter(1)
			case 0x68:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("PLA\n")
				incCounter(1)
			case 0x6A:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, file[counter])
				}
				fmt.Printf("ROR\n")
				incCounter(1)
			case 0x6B:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("TZA\n")
				incCounter(1)
			case 0x78:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("SEI\n")
				incCounter(1)
			case 0x7A:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("PLY\n")
				incCounter(1)
			case 0x7B:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(TBA - Absolute,Y)\n", PC, file[counter])
				}
				fmt.Printf("TBA\n")
				incCounter(1)
			case 0x88:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("DEY\n")
				incCounter(1)
			case 0x8A:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("TXA\n")
				incCounter(1)
			case 0x98:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("TYA\n")
				incCounter(1)
			case 0x9A:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("TXS\n")
				incCounter(1)
			case 0xA8:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("TAY\n")
				incCounter(1)
			case 0xAA:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("TAX\n")
				incCounter(1)
			case 0xB8:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("CLV\n")
				incCounter(1)
			case 0xBA:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("TSX\n")
				incCounter(1)
			case 0xC8:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("INY\n")
				incCounter(1)
			case 0xCA:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("DEX\n")
				incCounter(1)
			case 0xD8:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("CLD\n")
				incCounter(1)
			case 0xDA:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("PHX\n")
				incCounter(1)
			case 0xDB:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("PHZ\n")
				incCounter(1)
			case 0xE8:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("INX\n")
				incCounter(1)
			case 0xEA:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("NOP\n")
				incCounter(1)
			case 0xF8:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("SED\n")
				incCounter(1)
			case 0xFA:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("PLX\n")
				incCounter(1)
			case 0xFB:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, file[counter])
				}
				fmt.Printf("PLZ\n")
				incCounter(1)
			}

			//2 byte instructions with 1 operand
			switch file[counter] {
			case 0x01:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page Indirect))\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ORA ($%02x,X)\n", file[counter+1])
				incCounter(2)
			case 0x04:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("TSB $%02x\n", file[counter+1])
				incCounter(2)
			case 0x05:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ORA $%02x\n", file[counter+1])
				incCounter(2)
			case 0x06:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ASL $%02x\n", file[counter+1])
				incCounter(2)
			case 0x07:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("RMB0 $%02x\n", file[counter+1])
				incCounter(2)
			case 0x09:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ORA #$%02x\n", file[counter+1])
				incCounter(2)
			case 0x10:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("BPL $%02X\n", (counter+2+int(file[counter+1]))&0xFF)
				incCounter(2)
			case 0x11:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page),Indirect Y)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ORA ($%02x),Y\n", file[counter+1])
				incCounter(2)
			case 0x12:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page),Indirect Z)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ORA ($%02x),Z\n", file[counter+1])
				incCounter(2)
			case 0x14:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("TRB $%02x\n", file[counter+1])
				incCounter(2)
			case 0x15:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ORA $%02x,X\n", file[counter+1])
				incCounter(2)
			case 0x16:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(ASL - Zero Page,X)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ASL $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0x17:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("RMB1 $%02X\n", file[counter+1])
				incCounter(2)
			case 0x21:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page Indirect))\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("AND ($%02X,X)\n", file[counter+1])
				incCounter(2)
			case 0x24:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("BIT $%02X\n", file[counter+1])
				incCounter(2)
			case 0x25:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("AND $%02X\n", file[counter+1])
				incCounter(2)
			case 0x26:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ROL $%02X\n", file[counter+1])
				incCounter(2)
			case 0x27:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("RMB2 $%02X\n", file[counter+1])
				incCounter(2)
			case 0x29:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("AND #$%02X\n", file[counter+1])
				incCounter(2)
			case 0x30:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("BMI $%02X\n", (counter+2+int(file[counter+1]))&0xFF)
				incCounter(2)
			case 0x31:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("AND ($%02X),Y\n", file[counter+1])
				incCounter(2)
			case 0x32:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect),Z)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("AND ($%02X),Z\n", file[counter+1])
				incCounter(2)
			case 0x34:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("BIT $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0x35:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("AND $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0x36:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ROL $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0x37:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("RMB3 $%02X\n", file[counter+1])
				incCounter(2)
			case 0x41:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page, Indirect))\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("EOR ($%02X,X)\n", file[counter+1])
				incCounter(2)
			case 0x44:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ASR $%02X\n", file[counter+1])
				incCounter(2)
			case 0x45:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("EOR $%02X\n", file[counter+1])
				incCounter(2)
			case 0x46:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("LSR $%02X\n", file[counter+1])
				incCounter(2)
			case 0x47:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("RMB4 $%02X\n", file[counter+1])
				incCounter(2)
			case 0x49:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("EOR #$%02X\n", file[counter+1])
				incCounter(2)
			case 0x50:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("BVC $%02X\n", counter+2+int(file[counter+1]))
				incCounter(2)
			case 0x51:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("EOR ($%02X),Y\n", file[counter+1])
				incCounter(2)
			case 0x52:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect),Z)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("EOR ($%02X),Z\n", file[counter+1])
				incCounter(2)
			case 0x54:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ASR $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0x55:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("EOR $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0x56:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("LSR $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0x57:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("RMB5 $%02X\n", file[counter+1])
				incCounter(2)
			case 0x61:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ADC ($%02X,X)\n", file[counter+1])
				incCounter(2)
			case 0x62:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("RTN #$%02X\n", file[counter+1])
				incCounter(2)
			case 0x64:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("STZ $%02X\n", file[counter+1])
				incCounter(2)
			case 0x65:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ADC $%02X\n", file[counter+1])
				incCounter(2)
			case 0x66:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ROR $%02X\n", file[counter+1])
				incCounter(2)
			case 0x67:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("RMB6 $%02X\n", file[counter+1])
				incCounter(2)
			case 0x69:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ADC #$%02X\n", file[counter+1])
				incCounter(2)
			case 0x70:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("BVS $%04X\n", counter+2+int(file[counter+1]))
				incCounter(2)
			case 0x71:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ADC ($%02X),Y\n", file[counter+1])
				incCounter(2)
			case 0x72:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Indirect,Z)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ADC ($%02X),Z\n", file[counter+1])
				incCounter(2)
			case 0x74:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("STZ $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0x75:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ADC $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0x76:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("ROR $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0x77:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("RMB7 $%02X\n", file[counter+1])
				incCounter(2)
			case 0x80:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("BRA $%04X\n", counter+2+int(file[counter+1]))
				incCounter(2)
			case 0x81:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("STA ($%02X,X)\n", file[counter+1])
				incCounter(2)
			case 0x82:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Stack Relative Indirect,Y)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("STA ($%02X,S),Y\n", file[counter+1])
				incCounter(2)
			case 0x84:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("STY $%02X\n", file[counter+1])
				incCounter(2)
			case 0x85:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("STA $%02X\n", file[counter+1])
				incCounter(2)
			case 0x86:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("STX $%02X\n", file[counter+1])
				incCounter(2)
			case 0x87:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("SMB0 $%02X\n", file[counter+1])
				incCounter(2)
			case 0x89:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("BIT #$%02X\n", file[counter+1])
				incCounter(2)
			case 0x90:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("BCC $%02X\n", (counter+2+int(file[counter+1]))&0xFF)
				incCounter(2)
			case 0x91:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("STA ($%02X),Y\n", file[counter+1])
				incCounter(2)
			case 0x92:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Indirect,Z)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("STA ($%02X)\n", file[counter+1])
				incCounter(2)
			case 0x94:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("STY $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0x95:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("STA $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0x96:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Y)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("STX $%02X,Y\n", file[counter+1])
				incCounter(2)
			case 0x97:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("SMB1 $%02X\n", file[counter+1])
				incCounter(2)
			case 0xA0:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("LDY #$%02X\n", file[counter+1])
				incCounter(2)
			case 0xA1:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("LDA ($%02X,X)\n", file[counter+1])
				incCounter(2)
			case 0xA2:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("LDX #$%02X\n", file[counter+1])
				incCounter(2)
			case 0xA3:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("LDZ #$%02X\n", file[counter+1])
				incCounter(2)
			case 0xA4:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("LDY $%02X\n", file[counter+1])
				incCounter(2)
			case 0xA5:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("LDA $%02X\n", file[counter+1])
				incCounter(2)
			case 0xA6:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("LDX $%02X\n", file[counter+1])
				incCounter(2)
			case 0xA7:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("SMB2 $%02X\n", file[counter+1])
				incCounter(2)
			case 0xA9:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("LDA #$%02X\n", file[counter+1])
				incCounter(2)
			case 0xB0:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("BCS $%02X\n", (counter+2+int(file[counter+1]))&0xFF)
				incCounter(2)
			case 0xB1:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("LDA ($%02X),Y\n", file[counter+1])
				incCounter(2)
			case 0xB2:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Indirect)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("LDA ($%02X)\n", file[counter+1])
				incCounter(2)
			case 0xB4:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("LDY $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0xB5:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("LDA $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0xB6:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Y)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("LDX $%02X,Y\n", file[counter+1])
				incCounter(2)
			case 0xB7:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("SMB3 $%02X\n", file[counter+1])
				incCounter(2)
			case 0xC0:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("CPY #$%02X\n", file[counter+1])
				incCounter(2)
			case 0xC1:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("CMP ($%02X,X)\n", file[counter+1])
				incCounter(2)
			case 0xC2:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("CPZ #$%02X\n", file[counter+1])
				incCounter(2)
			case 0xC3:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("DEW $%02X\n", file[counter+1])
				incCounter(2)
			case 0xC4:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("CPY $%02X\n", file[counter+1])
				incCounter(2)
			case 0xC5:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("CMP $%02X\n", file[counter+1])
				incCounter(2)
			case 0xC6:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("DEC $%02X\n", file[counter+1])
				incCounter(2)
			case 0xC7:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("SMB4 $%02X\n", file[counter+1])
				incCounter(2)
			case 0xC9:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("CMP #$%02X\n", file[counter+1])
				incCounter(2)
			case 0xD0:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("BNE $%02X\n", (counter+2+int(file[counter+1]))&0xFF)
				incCounter(2)
			case 0xD1:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("CMP ($%02X),Y\n", file[counter+1])
				incCounter(2)
			case 0xD2:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect) Z)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("CMP ($%02X)\n", file[counter+1])
				incCounter(2)
			case 0xD4:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("CPZ $%02x\n", file[counter+1])
				incCounter(2)
			case 0xD5:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("CMP $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0xD6:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("DEC $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0xD7:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("SMB5 $%02X\n", file[counter+1])
				incCounter(2)
			case 0xE0:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("CPX #$%02X\n", file[counter+1])
				incCounter(2)
			case 0xE1:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("SBC ($%02X,X)\n", file[counter+1])
				incCounter(2)
			case 0xE2:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("LDA #$%02X\n", file[counter+1])
				incCounter(2)
			case 0xE3:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("INW $%02X\n", file[counter+1])
				incCounter(2)
			case 0xE4:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("CPX $%02X\n", file[counter+1])
				incCounter(2)
			case 0xE5:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("SBC $%02X\n", file[counter+1])
				incCounter(2)
			case 0xE6:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("INC $%02X\n", file[counter+1])
				incCounter(2)
			case 0xE7:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("SMB6 $%02X\n", file[counter+1])
				incCounter(2)
			case 0xE9:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("SBC #$%02X\n", file[counter+1])
				incCounter(2)
			case 0xF0:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("BEQ $%02X\n", (counter+2+int(file[counter+1]))&0xFF)
				incCounter(2)
			case 0xF1:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("SBC ($%02X),Y\n", file[counter+1])
				incCounter(2)
			case 0xF2:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect) Z)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("SBC ($%02X)\n", file[counter+1])
				incCounter(2)
			case 0xF5:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("SBC $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0xF6:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("INC $%02X,X\n", file[counter+1])
				incCounter(2)
			case 0xF7:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, file[counter], file[counter+1])
				}
				fmt.Printf("SMB7 $%02X\n", file[counter+1])
				incCounter(2)
			}

			//3 byte instructions with 2 operands
			switch file[counter] {
			case 0x0C:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("TSB $%02x%02x\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x0D:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("ORA $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x0E:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("ASL $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x0F:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BBR0 $%02X, $%02X\n", file[counter+1], file[counter+2])
				incCounter(3)
			case 0x13:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BPL $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x19:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("ORA $%02X%02X,Y\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x1C:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("TRB $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x1D:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("ORA $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x1E:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("ASL $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x1F:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BBR1 $%02X, $%02X\n", file[counter+1], file[counter+2])
				incCounter(3)
			case 0x20:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("JSR $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x22:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t((Indirect) Z)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("JSR ($%02X%02X)\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x23:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute X Indirect)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("JSR ($%02X%02X,X)\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x2C:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BIT $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x2D:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("AND $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x2E:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("ROL $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x2F:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BBR2 $%02X, $%02X\n", file[counter+1], file[counter+2])
				incCounter(3)
			case 0x33:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BMI $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x34:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BIT $%02X,X\n", file[counter+1])
				incCounter(3)
			case 0x35:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("AND $%02X,X\n", file[counter+1])
				incCounter(3)
			case 0x36:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("ROL $%02X,X\n", file[counter+1])
				incCounter(3)
			case 0x39:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("AND $%02X%02X,Y\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x3C:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("AND $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x3D:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("AND $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x3E:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("ROL $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x3F:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BBR3 $%02X, $%02X\n", file[counter+1], file[counter+2])
				incCounter(3)
			case 0x4C:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("JMP $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x4D:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("EOR $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x4E:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("LSR $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x4F:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BBR4 $%02X, $%02X\n", file[counter+1], file[counter+2])
				incCounter(3)
			case 0x53:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BVC $%02X\n", file[counter+1])
				incCounter(3)
			case 0x59:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("EOE $%02X%02X,Y\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x5D:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("EOR $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x5E:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("LSR $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x5F:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BBR5 $%02X, $%02X\n", file[counter+1], file[counter+2])
				incCounter(3)
			case 0x63:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BSR $%02X\n", file[counter+1])
				incCounter(3)
			case 0x6C:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute Indirect)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("JMP ($%02X%02X)\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x6D:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("ADC $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x6E:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("ROR $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x6F:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BBR6 $%02X, $%02X\n", file[counter+1], file[counter+2])
				incCounter(3)
			case 0x73:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BVS $%02X\n", file[counter+1])
				incCounter(3)
			case 0x79:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("ADC $%02X%02X,Y\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x7C:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X Indirect)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("JMP ($%02X%02X,X)\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x7D:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("ADC $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x7E:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("ROR $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x7F:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BBR7 $%02X, $%02X\n", file[counter+1], file[counter+2])
				incCounter(3)
			case 0x83:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BRA $%02X\n", file[counter+1])
				incCounter(3)
			case 0x8B:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("STY $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x8C:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("STY $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x8D:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("STA $%04X\n", file[counter+1]|uint8(int(file[counter+2])<<8))
				incCounter(3)
			case 0x8E:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("STX $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x8F:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(BBS0- Zero Page, Relative)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BBS0 $%02X, $%02X\n", file[counter+1], file[counter+2])
				incCounter(3)
			case 0x93:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BCC $%02X\n", file[counter+1])
				incCounter(3)
			case 0x99:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("STA $%02X%02X,Y\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x9B:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("STX $%02X%02X,Y\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x9C:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("STZ $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x9D:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("STA $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x9E:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("STZ $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0x9F:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BBS1 $%02X, $%02X\n", file[counter+1], file[counter+2])
				incCounter(3)
			case 0xAB:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("LDZ $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xAC:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("LDY $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xAD:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("LDA $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xAE:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("LDX $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xAF:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BBS2 $%02X, $%02X\n", file[counter+1], file[counter+2])
				incCounter(3)
			case 0xB3:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BCS $%02X\n", file[counter+1])
				incCounter(3)
			case 0xB9:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("LDA $%02X%02X,Y\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xBB:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("LDZ $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(1)
			case 0xBC:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("LDY $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xBD:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("LDA $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xBE:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("LDX $%02X%02X,Y\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xBF:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BBS3 $%02X, $%02X\n", file[counter+1], file[counter+2])
				incCounter(3)
			case 0xCB:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("ASW $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xCC:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("CPY $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xCD:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("CMP $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xCE:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("DEC $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xCF:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BBS4 $%02X, $%02X\n", file[counter+1], file[counter+2])
				incCounter(3)
			case 0xD3:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BNE $%02X\n", file[counter+1])
				incCounter(3)
			case 0xD9:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("CMP $%02X%02X,Y\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xDC:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("CPZ $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xDD:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("CMP $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xDE:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("DEC $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xDF:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BBS5 $%02X, $%02X\n", file[counter+1], file[counter+2])
				incCounter(3)
			case 0xEB:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("ROW $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xEC:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("CPX $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xED:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("SBC $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xEE:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("INC $%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xEF:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BBS6 $%02X, $%02X\n", file[counter+1], file[counter+2])
				incCounter(3)
			case 0xF3:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BEQ $%02X\n", file[counter+1])
				incCounter(3)
			case 0xF4:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Immediate (word))\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("PHW #$%02X%02X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xF9:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("SBC $%02X%02X,Y\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xFC:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("PHW #$%02X%02X\n", file[counter+1], file[counter+2])
				incCounter(3)
			case 0xFD:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("SBC $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xFE:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("INC $%02X%02X,X\n", file[counter+2], file[counter+1])
				incCounter(3)
			case 0xFF:
				if printHex {
					fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, file[counter], file[counter+1], file[counter+2])
				}
				fmt.Printf("BBS7 $%02X, $%02X\n", file[counter+1], file[counter+2])
				//incCounter(3)
				incCounter(3)
			}
		}
	}
*/

func execute(file string) {
	//PC += counter
	if printHex {
		fmt.Printf(" * = $%04X\n\n", PC)
	}
	for counter = 0; counter < len(file); PC++ {

		// 1 byte instructions with no operands
		switch memory[counter] {
		case 0x00:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("BRK\n")
			incCounter(1)
		case 0x02:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("CLE\n")
			incCounter(1)
		case 0x03:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("SEE\n")
			incCounter(1)
		case 0x08:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("PHP\n")
			incCounter(1)
		case 0x0A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, memory[counter])
			}
			fmt.Printf("ASL\n")
			incCounter(1)
		case 0x0B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("TSY\n")
			incCounter(1)
		case 0x18:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("CLC\n")
			incCounter(1)
		case 0x1A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, memory[counter])
			}
			fmt.Printf("INC\n")
			incCounter(1)
		case 0x1B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("INZ\n")
			incCounter(1)
		case 0x28:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("PLP\n")
			incCounter(1)
		case 0x2A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, memory[counter])
			}
			fmt.Printf("ROL\n")
			incCounter(1)
		case 0x2B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("TYS\n")
			incCounter(1)
		case 0x38:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, memory[counter])
			}
			fmt.Printf("SEC\n")
			incCounter(1)
		case 0x3A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, memory[counter])
			}
			fmt.Printf("DEC\n")
			incCounter(1)
		case 0x3B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("DEZ\n")
			incCounter(1)
		case 0x40:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("RTI\n")
			incCounter(1)
		case 0x42:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, memory[counter])
			}
			fmt.Printf("NEG\n")
			incCounter(1)
		case 0x43:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, memory[counter])
			}
			fmt.Printf("ASR\n")
			incCounter(1)
		case 0x48:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("PHA\n")
			//incCounter(1)
			incCounter(1)
		case 0x4A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, memory[counter])
			}
			fmt.Printf("LSR\n")
			incCounter(1)
		case 0x4B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("TAZ\n")
			incCounter(1)
		case 0x58:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("CLI\n")
			incCounter(1)
		case 0x5A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("PHY\n")
			incCounter(1)
		case 0x5B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("TAB\n")
			incCounter(1)
		case 0x60:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("RTS\n")
			incCounter(1)
		case 0x68:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("PLA\n")
			incCounter(1)
		case 0x6A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", PC, memory[counter])
			}
			fmt.Printf("ROR\n")
			incCounter(1)
		case 0x6B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("TZA\n")
			incCounter(1)
		case 0x78:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("SEI\n")
			incCounter(1)
		case 0x7A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("PLY\n")
			incCounter(1)
		case 0x7B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(TBA - Absolute,Y)\n", PC, memory[counter])
			}
			fmt.Printf("TBA\n")
			incCounter(1)
		case 0x88:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("DEY\n")
			incCounter(1)
		case 0x8A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("TXA\n")
			incCounter(1)
		case 0x98:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("TYA\n")
			incCounter(1)
		case 0x9A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("TXS\n")
			incCounter(1)
		case 0xA8:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("TAY\n")
			incCounter(1)
		case 0xAA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("TAX\n")
			incCounter(1)
		case 0xB8:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("CLV\n")
			incCounter(1)
		case 0xBA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("TSX\n")
			incCounter(1)
		case 0xC8:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("INY\n")
			incCounter(1)
		case 0xCA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("DEX\n")
			incCounter(1)
		case 0xD8:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("CLD\n")
			incCounter(1)
		case 0xDA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("PHX\n")
			incCounter(1)
		case 0xDB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("PHZ\n")
			incCounter(1)
		case 0xE8:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("INX\n")
			incCounter(1)
		case 0xEA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("NOP\n")
			incCounter(1)
		case 0xF8:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("SED\n")
			incCounter(1)
		case 0xFA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("PLX\n")
			incCounter(1)
		case 0xFB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", PC, memory[counter])
			}
			fmt.Printf("PLZ\n")
			incCounter(1)
		}

		//2 byte instructions with 1 operand
		switch memory[counter] {
		case 0x01:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page Indirect))\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ORA ($%02x,X)\n", memory[counter+1])
			incCounter(2)
		case 0x04:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("TSB $%02x\n", memory[counter+1])
			incCounter(2)
		case 0x05:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ORA $%02x\n", memory[counter+1])
			incCounter(2)
		case 0x06:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ASL $%02x\n", memory[counter+1])
			incCounter(2)
		case 0x07:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("RMB0 $%02x\n", memory[counter+1])
			incCounter(2)
		case 0x09:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ORA #$%02x\n", memory[counter+1])
			incCounter(2)
		case 0x10:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("BPL $%02X\n", (counter+2+int(memory[counter+1]))&0xFF)
			incCounter(2)
		case 0x11:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page),Indirect Y)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ORA ($%02x),Y\n", memory[counter+1])
			incCounter(2)
		case 0x12:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page),Indirect Z)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ORA ($%02x),Z\n", memory[counter+1])
			incCounter(2)
		case 0x14:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("TRB $%02x\n", memory[counter+1])
			incCounter(2)
		case 0x15:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ORA $%02x,X\n", memory[counter+1])
			incCounter(2)
		case 0x16:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(ASL - Zero Page,X)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ASL $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0x17:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("RMB1 $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x21:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page Indirect))\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("AND ($%02X,X)\n", memory[counter+1])
			incCounter(2)
		case 0x24:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("BIT $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x25:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("AND $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x26:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ROL $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x27:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("RMB2 $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x29:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("AND #$%02X\n", memory[counter+1])
			incCounter(2)
		case 0x30:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("BMI $%02X\n", (counter+2+int(memory[counter+1]))&0xFF)
			incCounter(2)
		case 0x31:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("AND ($%02X),Y\n", memory[counter+1])
			incCounter(2)
		case 0x32:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect),Z)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("AND ($%02X),Z\n", memory[counter+1])
			incCounter(2)
		case 0x34:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("BIT $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0x35:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("AND $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0x36:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ROL $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0x37:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("RMB3 $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x41:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page, Indirect))\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("EOR ($%02X,X)\n", memory[counter+1])
			incCounter(2)
		case 0x44:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ASR $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x45:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("EOR $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x46:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("LSR $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x47:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("RMB4 $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x49:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("EOR #$%02X\n", memory[counter+1])
			incCounter(2)
		case 0x50:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("BVC $%02X\n", counter+2+int(memory[counter+1]))
			incCounter(2)
		case 0x51:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("EOR ($%02X),Y\n", memory[counter+1])
			incCounter(2)
		case 0x52:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect),Z)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("EOR ($%02X),Z\n", memory[counter+1])
			incCounter(2)
		case 0x54:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ASR $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0x55:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("EOR $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0x56:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("LSR $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0x57:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("RMB5 $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x61:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ADC ($%02X,X)\n", memory[counter+1])
			incCounter(2)
		case 0x62:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("RTN #$%02X\n", memory[counter+1])
			incCounter(2)
		case 0x64:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("STZ $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x65:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ADC $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x66:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ROR $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x67:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("RMB6 $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x69:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ADC #$%02X\n", memory[counter+1])
			incCounter(2)
		case 0x70:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("BVS $%04X\n", counter+2+int(memory[counter+1]))
			incCounter(2)
		case 0x71:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ADC ($%02X),Y\n", memory[counter+1])
			incCounter(2)
		case 0x72:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Indirect,Z)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ADC ($%02X),Z\n", memory[counter+1])
			incCounter(2)
		case 0x74:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("STZ $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0x75:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ADC $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0x76:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("ROR $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0x77:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("RMB7 $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x80:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("BRA $%04X\n", counter+2+int(memory[counter+1]))
			incCounter(2)
		case 0x81:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("STA ($%02X,X)\n", memory[counter+1])
			incCounter(2)
		case 0x82:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Stack Relative Indirect,Y)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("STA ($%02X,S),Y\n", memory[counter+1])
			incCounter(2)
		case 0x84:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("STY $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x85:
			/*
				STA - Store Accumulator in Memory

				Operation: A → M

				This instruction transfers the contents of the accumulator to memory.

				This instruction affects none of the flags in the processor status register and does not affect the accumulator.
			*/

			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("STA $%02X\n", memory[counter+1])

			//Store contents of Accumulator in memory
			memory[memory[counter+1]] = A
			fmt.Printf("Address[$%02X] = $%02X\n", memory[counter+1], memory[memory[counter+1]])
			printMachineState()

			incCounter(2)
		case 0x86:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("STX $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x87:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("SMB0 $%02X\n", memory[counter+1])
			incCounter(2)
		case 0x89:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("BIT #$%02X\n", memory[counter+1])
			incCounter(2)
		case 0x90:
			/*
				BCC - Branch on Carry Clear
				Operation: Branch on C = 0

				This instruction tests the state of the carry bit and takes a conditional branch if the carry bit is reset.

				It affects no flags or registers other than the program counter and then only if the C flag is not on.
			*/

			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("BCC $%02X\n", (counter+2+int(memory[counter+1]))&0xFF)

			//If carry flag/bit zero of the status register is clear, then branch to the address specified by the operand.
			if SR&1 == 0 {
				PC = int(memory[counter+1])
			}
			printMachineState()

			incCounter(2)
		case 0x91:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("STA ($%02X),Y\n", memory[counter+1])
			incCounter(2)
		case 0x92:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Indirect,Z)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("STA ($%02X)\n", memory[counter+1])
			incCounter(2)
		case 0x94:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("STY $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0x95:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("STA $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0x96:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Y)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("STX $%02X,Y\n", memory[counter+1])
			incCounter(2)
		case 0x97:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("SMB1 $%02X\n", memory[counter+1])
			incCounter(2)
		case 0xA0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("LDY #$%02X\n", memory[counter+1])
			incCounter(2)
		case 0xA1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("LDA ($%02X,X)\n", memory[counter+1])
			incCounter(2)
		case 0xA2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("LDX #$%02X\n", memory[counter+1])
			incCounter(2)
		case 0xA3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("LDZ #$%02X\n", memory[counter+1])
			incCounter(2)
		case 0xA4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("LDY $%02X\n", memory[counter+1])
			incCounter(2)
		case 0xA5:
			/*
				LDA - Load Accumulator with Memory
				Operation: M → A

				When instruction LDA is executed by the microprocessor, data is transferred from memory to the accumulator and stored in the accumulator.

				LDA affects the contents of the accumulator, does not affect the carry or overflow flags;
				sets the zero flag if the accumulator is zero as a result of the LDA, otherwise resets the zero flag;
				sets the negative flag if bit 7 of the accumulator is a 1, otherwise resets the negative flag.
			*/
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("LDA $%02X\n", memory[counter+1])

			//Load the accumulator with the value in the operand
			A = memory[counter+1]
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

			printMachineState()

			incCounter(2)
		case 0xA6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("LDX $%02X\n", memory[counter+1])
			incCounter(2)
		case 0xA7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("SMB2 $%02X\n", memory[counter+1])
			incCounter(2)
		case 0xA9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("LDA #$%02X\n", memory[counter+1])
			incCounter(2)
		case 0xB0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("BCS $%02X\n", (counter+2+int(memory[counter+1]))&0xFF)
			incCounter(2)
		case 0xB1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("LDA ($%02X),Y\n", memory[counter+1])
			incCounter(2)
		case 0xB2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Indirect)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("LDA ($%02X)\n", memory[counter+1])
			incCounter(2)
		case 0xB4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("LDY $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0xB5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("LDA $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0xB6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Y)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("LDX $%02X,Y\n", memory[counter+1])
			incCounter(2)
		case 0xB7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("SMB3 $%02X\n", memory[counter+1])
			incCounter(2)
		case 0xC0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("CPY #$%02X\n", memory[counter+1])
			incCounter(2)
		case 0xC1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("CMP ($%02X,X)\n", memory[counter+1])
			incCounter(2)
		case 0xC2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("CPZ #$%02X\n", memory[counter+1])
			incCounter(2)
		case 0xC3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("DEW $%02X\n", memory[counter+1])
			incCounter(2)
		case 0xC4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("CPY $%02X\n", memory[counter+1])
			incCounter(2)
		case 0xC5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("CMP $%02X\n", memory[counter+1])
			incCounter(2)
		case 0xC6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("DEC $%02X\n", memory[counter+1])
			incCounter(2)
		case 0xC7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("SMB4 $%02X\n", memory[counter+1])
			incCounter(2)
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
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("CMP #$%02X\n", memory[counter+1])

			//Compare memory and accumulator
			if memory[counter+1] == A {
				//Set Z flag and negative flag to true
				//Set bit 1 to true
				SR |= 0 << 1
				//SR = 0b00100010
			}
			//Set carry flag to true if A is greater than or equal to operand
			if A >= memory[counter+1] {
				//Set bit 0 to true
				SR |= 1 << 0
				//SR = 0b00100001
			}
			//Set carry flag to false if A is less than operand
			if A < memory[counter+1] {
				//Set bit zero to false
				SR |= 0 << 0
				//SR = 0b00100000
			}
			//Set Z flag to false if A is not equal to operand
			if A != memory[counter+1] {
				//Set bit 1 to false
				SR |= 0 << 1
				//SR = 0b00100000
			}
			//Set N flag to true if A minus operand results in most significant bit being set
			if (A-memory[counter+1])&0b10000000 == 0b10000000 {
				//Set bit 7 to true
				SR |= 1 << 7
				//SR = 0b10100000
			}
			//Set N flag to false if A minus operand results in most significant bit being unset
			if (A-memory[counter+1])&0b10000000 == 0b00000000 {
				//Set bit 7 to false
				SR |= 0 << 7
				//SR = 0b00100000
			}
			printMachineState()

			incCounter(2)
		case 0xD0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("BNE $%02X\n", (counter+2+int(memory[counter+1]))&0xFF)
			incCounter(2)
		case 0xD1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("CMP ($%02X),Y\n", memory[counter+1])
			incCounter(2)
		case 0xD2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect) Z)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("CMP ($%02X)\n", memory[counter+1])
			incCounter(2)
		case 0xD4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("CPZ $%02x\n", memory[counter+1])
			incCounter(2)
		case 0xD5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("CMP $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0xD6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("DEC $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0xD7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("SMB5 $%02X\n", memory[counter+1])
			incCounter(2)
		case 0xE0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("CPX #$%02X\n", memory[counter+1])
			incCounter(2)
		case 0xE1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("SBC ($%02X,X)\n", memory[counter+1])
			incCounter(2)
		case 0xE2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("LDA #$%02X\n", memory[counter+1])
			incCounter(2)
		case 0xE3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("INW $%02X\n", memory[counter+1])
			incCounter(2)
		case 0xE4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("CPX $%02X\n", memory[counter+1])
			incCounter(2)
		case 0xE5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("SBC $%02X\n", memory[counter+1])
			incCounter(2)
		case 0xE6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("INC $%02X\n", memory[counter+1])
			incCounter(2)
		case 0xE7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("SMB6 $%02X\n", memory[counter+1])
			incCounter(2)
		case 0xE9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("SBC #$%02X\n", memory[counter+1])
			incCounter(2)
		case 0xF0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("BEQ $%02X\n", (counter+2+int(memory[counter+1]))&0xFF)
			incCounter(2)
		case 0xF1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("SBC ($%02X),Y\n", memory[counter+1])
			incCounter(2)
		case 0xF2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect) Z)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("SBC ($%02X)\n", memory[counter+1])
			incCounter(2)
		case 0xF5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("SBC $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0xF6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("INC $%02X,X\n", memory[counter+1])
			incCounter(2)
		case 0xF7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", PC, memory[counter], memory[counter+1])
			}
			fmt.Printf("SMB7 $%02X\n", memory[counter+1])
			incCounter(2)
		}

		//3 byte instructions with 2 operands
		switch memory[counter] {
		case 0x0C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("TSB $%02x%02x\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x0D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("ORA $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x0E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("ASL $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x0F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BBR0 $%02X, $%02X\n", memory[counter+1], memory[counter+2])
			incCounter(3)
		case 0x13:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BPL $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x19:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("ORA $%02X%02X,Y\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x1C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("TRB $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x1D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("ORA $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x1E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("ASL $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x1F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BBR1 $%02X, $%02X\n", memory[counter+1], memory[counter+2])
			incCounter(3)
		case 0x20:
			/*
				JSR - Jump To Subroutine
				Operation: PC + 2↓, [PC + 1] → PCL, [PC + 2] → PCH

				This instruction transfers control of the program counter to a subroutine location but leaves a
				return pointer on the stack to allow the user to return to perform the next instruction in the
				main program after the subroutine is complete.

				To accomplish this, JSR instruction stores the program counter address which points to the last byte of the
				jump instruction onto the stack using the stack pointer. The stack byte contains the program count high first,
				followed by program count low. The JSR then transfers the addresses following the jump instruction to the
				program counter low and the program counter high, thereby directing the program to begin at that new address.

				The JSR instruction affects no flags, causes the stack pointer to be decremented by 2 and substitutes new values into the program counter low and the program counter high.
			*/

			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("JSR $%02X%02X\n", memory[counter+2], memory[counter+1])

			//Store PC at address stored in SP
			memory[SP] = byte(PC)
			//Set PC to address stored in operand 1 and operand 2
			PC = int(memory[counter+1]) + int(memory[counter+2])<<8
			printMachineState()

			incCounter(3)
		case 0x22:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t((Indirect) Z)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("JSR ($%02X%02X)\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x23:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute X Indirect)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("JSR ($%02X%02X,X)\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x2C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BIT $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x2D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("AND $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x2E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("ROL $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x2F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BBR2 $%02X, $%02X\n", memory[counter+1], memory[counter+2])
			incCounter(3)
		case 0x33:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BMI $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x34:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BIT $%02X,X\n", memory[counter+1])
			incCounter(3)
		case 0x35:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("AND $%02X,X\n", memory[counter+1])
			incCounter(3)
		case 0x36:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("ROL $%02X,X\n", memory[counter+1])
			incCounter(3)
		case 0x39:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("AND $%02X%02X,Y\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x3C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("AND $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x3D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("AND $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x3E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("ROL $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x3F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BBR3 $%02X, $%02X\n", memory[counter+1], memory[counter+2])
			incCounter(3)
		case 0x4C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("JMP $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x4D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("EOR $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x4E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("LSR $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x4F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BBR4 $%02X, $%02X\n", memory[counter+1], memory[counter+2])
			incCounter(3)
		case 0x53:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BVC $%02X\n", memory[counter+1])
			incCounter(3)
		case 0x59:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("EOE $%02X%02X,Y\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x5D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("EOR $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x5E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("LSR $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x5F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BBR5 $%02X, $%02X\n", memory[counter+1], memory[counter+2])
			incCounter(3)
		case 0x63:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BSR $%02X\n", memory[counter+1])
			incCounter(3)
		case 0x6C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute Indirect)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("JMP ($%02X%02X)\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x6D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("ADC $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x6E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("ROR $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x6F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BBR6 $%02X, $%02X\n", memory[counter+1], memory[counter+2])
			incCounter(3)
		case 0x73:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BVS $%02X\n", memory[counter+1])
			incCounter(3)
		case 0x79:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("ADC $%02X%02X,Y\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x7C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X Indirect)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("JMP ($%02X%02X,X)\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x7D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("ADC $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x7E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("ROR $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x7F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BBR7 $%02X, $%02X\n", memory[counter+1], memory[counter+2])
			incCounter(3)
		case 0x83:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BRA $%02X\n", memory[counter+1])
			incCounter(3)
		case 0x8B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("STY $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x8C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("STY $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x8D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("STA $%04X\n", memory[counter+1]|uint8(int(memory[counter+2])<<8))
			incCounter(3)
		case 0x8E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("STX $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x8F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(BBS0- Zero Page, Relative)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BBS0 $%02X, $%02X\n", memory[counter+1], memory[counter+2])
			incCounter(3)
		case 0x93:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BCC $%02X\n", memory[counter+1])
			incCounter(3)
		case 0x99:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("STA $%02X%02X,Y\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x9B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("STX $%02X%02X,Y\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x9C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("STZ $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x9D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("STA $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x9E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("STZ $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0x9F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BBS1 $%02X, $%02X\n", memory[counter+1], memory[counter+2])
			incCounter(3)
		case 0xAB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("LDZ $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xAC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("LDY $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xAD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("LDA $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xAE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("LDX $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xAF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BBS2 $%02X, $%02X\n", memory[counter+1], memory[counter+2])
			incCounter(3)
		case 0xB3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BCS $%02X\n", memory[counter+1])
			incCounter(3)
		case 0xB9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("LDA $%02X%02X,Y\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xBB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("LDZ $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(1)
		case 0xBC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("LDY $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xBD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("LDA $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xBE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("LDX $%02X%02X,Y\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xBF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BBS3 $%02X, $%02X\n", memory[counter+1], memory[counter+2])
			incCounter(3)
		case 0xCB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("ASW $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xCC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("CPY $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xCD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("CMP $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xCE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("DEC $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xCF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BBS4 $%02X, $%02X\n", memory[counter+1], memory[counter+2])
			incCounter(3)
		case 0xD3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BNE $%02X\n", memory[counter+1])
			incCounter(3)
		case 0xD9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("CMP $%02X%02X,Y\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xDC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("CPZ $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xDD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("CMP $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xDE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("DEC $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xDF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BBS5 $%02X, $%02X\n", memory[counter+1], memory[counter+2])
			incCounter(3)
		case 0xEB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("ROW $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xEC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("CPX $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xED:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("SBC $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xEE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("INC $%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xEF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BBS6 $%02X, $%02X\n", memory[counter+1], memory[counter+2])
			incCounter(3)
		case 0xF3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Relative (word))\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BEQ $%02X\n", memory[counter+1])
			incCounter(3)
		case 0xF4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Immediate (word))\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("PHW #$%02X%02X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xF9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("SBC $%02X%02X,Y\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xFC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("PHW #$%02X%02X\n", memory[counter+1], memory[counter+2])
			incCounter(3)
		case 0xFD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("SBC $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xFE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("INC $%02X%02X,X\n", memory[counter+2], memory[counter+1])
			incCounter(3)
		case 0xFF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", PC, memory[counter], memory[counter+1], memory[counter+2])
			}
			fmt.Printf("BBS7 $%02X, $%02X\n", memory[counter+1], memory[counter+2])
			//incCounter(3)
			incCounter(3)
		}
	}
}
