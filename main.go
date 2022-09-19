package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	printHex      bool
	counter       = 0
	memoryAddress = 0x0000 // Add entry point command line parameter later
	file          []byte
)

func main() {
	fmt.Printf("Six5go2 - 6502 Emulator and Disassembler in Golang (c) 2022 Zayn Otley\n\n")

	if len(os.Args) <= 2 {
		fmt.Printf("USAGE - %s <target_filename> <entry_point_address> <hex>\n", os.Args[0])
		os.Exit(0)
	}

	if len(os.Args) > 2 {
		parseUint, _ := strconv.ParseUint(os.Args[2], 16, 16)
		memoryAddress = int(parseUint)
	}

	if len(os.Args) > 3 && os.Args[3] == "hex" {
		printHex = true
	}

	// Read file
	file, _ = os.ReadFile(os.Args[1])

	fmt.Printf("USAGE   - six5go2 <target_filename> <entry_point> (Hex memory address) <hex> (Print hex values above each instruction) \n")
	fmt.Printf("EXAMPLE - six5go2 cbmbasic35.rom 0800 hex\n\n")
	fmt.Printf("Length of file %s is %v ($%04X) bytes\n\n", os.Args[1], len(file), len(file))

	disassemble(string(file))
}

func incCounter(amount int) {
	if counter+amount < len(file)-1 {
		counter += amount
	} else {
		os.Exit(0)
	}
}
func disassemble(file string) {
	memoryAddress += counter
	if printHex {
		fmt.Printf(" * = $%04X\n\n", memoryAddress)
	}
	for counter = 0; counter < len(file); memoryAddress++ {

		// 1 byte instructions with no operands
		switch file[counter] {
		case 0x00:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("BRK\n")
			incCounter(1)
		case 0x02:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("CLE\n")
			incCounter(1)
		case 0x03:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("SEE\n")
			incCounter(1)
		case 0x08:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("PHP\n")
			incCounter(1)
		case 0x0A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("ASL\n")
			incCounter(1)
		case 0x0B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("TSY\n")
			incCounter(1)
		case 0x18:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("CLC\n")
			incCounter(1)
		case 0x1A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("INC\n")
			incCounter(1)
		case 0x1B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("INZ\n")
			incCounter(1)
		case 0x28:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("PLP\n")
			incCounter(1)
		case 0x2A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("ROL\n")
			incCounter(1)
		case 0x2B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("TYS\n")
			incCounter(1)
		case 0x38:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("SEC\n")
			incCounter(1)
		case 0x3A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("DEC\n")
			incCounter(1)
		case 0x3B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("DEZ\n")
			incCounter(1)
		case 0x40:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("RTI\n")
			incCounter(1)
		case 0x42:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("NEG\n")
			incCounter(1)
		case 0x43:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("ASR\n")
			incCounter(1)
		case 0x48:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("PHA\n")
			//incCounter(1)
			incCounter(1)
		case 0x4A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("LSR\n")
			incCounter(1)
		case 0x4B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("TAZ\n")
			incCounter(1)
		case 0x58:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("CLI\n")
			incCounter(1)
		case 0x5A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("PHY\n")
			incCounter(1)
		case 0x5B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("TAB\n")
			incCounter(1)
		case 0x60:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("RTS\n")
			incCounter(1)
		case 0x68:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("PLA\n")
			incCounter(1)
		case 0x6A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Accumulator)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("ROR\n")
			incCounter(1)
		case 0x6B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("TZA\n")
			incCounter(1)
		case 0x78:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("SEI\n")
			incCounter(1)
		case 0x7A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("PLY\n")
			incCounter(1)
		case 0x7B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(TBA - Absolute,Y)\n", memoryAddress, file[counter])
			}
			fmt.Printf("TBA\n")
			incCounter(1)
		case 0x88:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("DEY\n")
			incCounter(1)
		case 0x8A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("TXA\n")
			incCounter(1)
		case 0x98:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("TYA\n")
			incCounter(1)
		case 0x9A:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("TXS\n")
			incCounter(1)
		case 0xA8:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("TAY\n")
			incCounter(1)
		case 0xAA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("TAX\n")
			incCounter(1)
		case 0xB8:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("CLV\n")
			incCounter(1)
		case 0xBA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("TSX\n")
			incCounter(1)
		case 0xC8:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("INY\n")
			incCounter(1)
		case 0xCA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("DEX\n")
			incCounter(1)
		case 0xD8:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("CLD\n")
			incCounter(1)
		case 0xDA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("PHX\n")
			incCounter(1)
		case 0xDB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("PHZ\n")
			incCounter(1)
		case 0xE8:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("INX\n")
			incCounter(1)
		case 0xEA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("NOP\n")
			incCounter(1)
		case 0xF8:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("SED\n")
			incCounter(1)
		case 0xFA:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("PLX\n")
			incCounter(1)
		case 0xFB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x\t\t(Implied)\t\n", memoryAddress, file[counter])
			}
			fmt.Printf("PLZ\n")
			incCounter(1)
		}

		//2 byte instructions with 1 operand
		switch file[counter] {
		case 0x01:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page Indirect))\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ORA ($%02x,X)\n", file[counter+1])
			incCounter(2)
		case 0x04:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("TSB $%02x\n", file[counter+1])
			incCounter(2)
		case 0x05:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ORA $%02x\n", file[counter+1])
			incCounter(2)
		case 0x06:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ASL $%02x\n", file[counter+1])
			incCounter(2)
		case 0x07:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("RMB0 $%02x\n", file[counter+1])
			incCounter(2)
		case 0x09:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ORA #$%02x\n", file[counter+1])
			incCounter(2)
		case 0x10:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("BPL $%02X\n", (counter+2+int(file[counter+1]))&0xFF)
			incCounter(2)
		case 0x11:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page),Indirect Y)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ORA ($%02x),Y\n", file[counter+1])
			incCounter(2)
		case 0x12:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page),Indirect Z)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ORA ($%02x),Z\n", file[counter+1])
			incCounter(2)
		case 0x14:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("TRB $%02x\n", file[counter+1])
			incCounter(2)
		case 0x15:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ORA $%02x,X\n", file[counter+1])
			incCounter(2)
		case 0x16:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(ASL - Zero Page,X)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ASL $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0x17:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("RMB1 $%02X\n", file[counter+1])
			incCounter(2)
		case 0x21:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page Indirect))\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("AND ($%02X,X)\n", file[counter+1])
			incCounter(2)
		case 0x24:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("BIT $%02X\n", file[counter+1])
			incCounter(2)
		case 0x25:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("AND $%02X\n", file[counter+1])
			incCounter(2)
		case 0x26:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ROL $%02X\n", file[counter+1])
			incCounter(2)
		case 0x27:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("RMB2 $%02X\n", file[counter+1])
			incCounter(2)
		case 0x29:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("AND #$%02X\n", file[counter+1])
			incCounter(2)
		case 0x30:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("BMI $%02X\n", (counter+2+int(file[counter+1]))&0xFF)
			incCounter(2)
		case 0x31:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("AND ($%02X),Y\n", file[counter+1])
			incCounter(2)
		case 0x32:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect),Z)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("AND ($%02X),Z\n", file[counter+1])
			incCounter(2)
		case 0x34:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("BIT $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0x35:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("AND $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0x36:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ROL $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0x37:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("RMB3 $%02X\n", file[counter+1])
			incCounter(2)
		case 0x41:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((X Zero Page, Indirect))\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("EOR ($%02X,X)\n", file[counter+1])
			incCounter(2)
		case 0x44:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ASR $%02X\n", file[counter+1])
			incCounter(2)
		case 0x45:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("EOR $%02X\n", file[counter+1])
			incCounter(2)
		case 0x46:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("LSR $%02X\n", file[counter+1])
			incCounter(2)
		case 0x47:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("RMB4 $%02X\n", file[counter+1])
			incCounter(2)
		case 0x49:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("EOR #$%02X\n", file[counter+1])
			incCounter(2)
		case 0x50:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("BVC $%02X\n", counter+2+int(file[counter+1]))
			incCounter(2)
		case 0x51:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("EOR ($%02X),Y\n", file[counter+1])
			incCounter(2)
		case 0x52:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect),Z)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("EOR ($%02X),Z\n", file[counter+1])
			incCounter(2)
		case 0x54:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ASR $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0x55:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("EOR $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0x56:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("LSR $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0x57:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("RMB5 $%02X\n", file[counter+1])
			incCounter(2)
		case 0x61:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ADC ($%02X,X)\n", file[counter+1])
			incCounter(2)
		case 0x62:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("RTN #$%02X\n", file[counter+1])
			incCounter(2)
		case 0x64:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("STZ $%02X\n", file[counter+1])
			incCounter(2)
		case 0x65:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ADC $%02X\n", file[counter+1])
			incCounter(2)
		case 0x66:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ROR $%02X\n", file[counter+1])
			incCounter(2)
		case 0x67:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("RMB6 $%02X\n", file[counter+1])
			incCounter(2)
		case 0x69:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ADC #$%02X\n", file[counter+1])
			incCounter(2)
		case 0x70:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("BVS $%04X\n", counter+2+int(file[counter+1]))
			incCounter(2)
		case 0x71:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ADC ($%02X),Y\n", file[counter+1])
			incCounter(2)
		case 0x72:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Indirect,Z)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ADC ($%02X),Z\n", file[counter+1])
			incCounter(2)
		case 0x74:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("STZ $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0x75:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ADC $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0x76:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("ROR $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0x77:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("RMB7 $%02X\n", file[counter+1])
			incCounter(2)
		case 0x80:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("BRA $%04X\n", counter+2+int(file[counter+1]))
			incCounter(2)
		case 0x81:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("STA ($%02X,X)\n", file[counter+1])
			incCounter(2)
		case 0x82:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Stack Realtive Indirect,Y)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("STA ($%02X,S),Y\n", file[counter+1])
			incCounter(2)
		case 0x84:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("STY $%02X\n", file[counter+1])
			incCounter(2)
		case 0x85:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("STA $%02X\n", file[counter+1])
			incCounter(2)
		case 0x86:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("STX $%02X\n", file[counter+1])
			incCounter(2)
		case 0x87:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("SMB0 $%02X\n", file[counter+1])
			incCounter(2)
		case 0x89:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("BIT #$%02X\n", file[counter+1])
			incCounter(2)
		case 0x90:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("BCC $%02X\n", (counter+2+int(file[counter+1]))&0xFF)
			incCounter(2)
		case 0x91:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("STA ($%02X),Y\n", file[counter+1])
			incCounter(2)
		case 0x92:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Indirect,Z)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("STA ($%02X)\n", file[counter+1])
			incCounter(2)
		case 0x94:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("STY $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0x95:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("STA $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0x96:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Y)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("STX $%02X,Y\n", file[counter+1])
			incCounter(2)
		case 0x97:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("SMB1 $%02X\n", file[counter+1])
			incCounter(2)
		case 0xA0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("LDY #$%02X\n", file[counter+1])
			incCounter(2)
		case 0xA1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("LDA ($%02X,X)\n", file[counter+1])
			incCounter(2)
		case 0xA2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("LDX #$%02X\n", file[counter+1])
			incCounter(2)
		case 0xA3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("LDZ #$%02X\n", file[counter+1])
			incCounter(2)
		case 0xA4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("LDY $%02X\n", file[counter+1])
			incCounter(2)
		case 0xA5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("LDA $%02X\n", file[counter+1])
			incCounter(2)
		case 0xA6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("LDX $%02X\n", file[counter+1])
			incCounter(2)
		case 0xA7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("SMB2 $%02X\n", file[counter+1])
			incCounter(2)
		case 0xA9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("LDA #$%02X\n", file[counter+1])
			incCounter(2)
		case 0xB0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("BCS $%02X\n", (counter+2+int(file[counter+1]))&0xFF)
			incCounter(2)
		case 0xB1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("LDA ($%02X),Y\n", file[counter+1])
			incCounter(2)
		case 0xB2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Indirect)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("LDA ($%02X)\n", file[counter+1])
			incCounter(2)
		case 0xB4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("LDY $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0xB5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("LDA $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0xB6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,Y)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("LDX $%02X,Y\n", file[counter+1])
			incCounter(2)
		case 0xB7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("SMB3 $%02X\n", file[counter+1])
			incCounter(2)
		case 0xC0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("CPY #$%02X\n", file[counter+1])
			incCounter(2)
		case 0xC1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("CMP ($%02X,X)\n", file[counter+1])
			incCounter(2)
		case 0xC2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("CPZ #$%02X\n", file[counter+1])
			incCounter(2)
		case 0xC3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("DEW $%02X\n", file[counter+1])
			incCounter(2)
		case 0xC4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("CPY $%02X\n", file[counter+1])
			incCounter(2)
		case 0xC5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("CMP $%02X\n", file[counter+1])
			incCounter(2)
		case 0xC6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("DEC $%02X\n", file[counter+1])
			incCounter(2)
		case 0xC7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("SMB4 $%02X\n", file[counter+1])
			incCounter(2)
		case 0xC9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("CMP #$%02X\n", file[counter+1])
			incCounter(2)
		case 0xD0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("BNE $%02X\n", (counter+2+int(file[counter+1]))&0xFF)
			incCounter(2)
		case 0xD1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("CMP ($%02X),Y\n", file[counter+1])
			incCounter(2)
		case 0xD2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect) Z)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("CMP ($%02X)\n", file[counter+1])
			incCounter(2)
		case 0xD4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("CPZ $%02x\n", file[counter+1])
			incCounter(2)
		case 0xD5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("CMP $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0xD6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("DEC $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0xD7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("SMB5 $%02X\n", file[counter+1])
			incCounter(2)
		case 0xE0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("CPX #$%02X\n", file[counter+1])
			incCounter(2)
		case 0xE1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(X Zero Page Indirect)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("SBC ($%02X,X)\n", file[counter+1])
			incCounter(2)
		case 0xE2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("LDA #$%02X\n", file[counter+1])
			incCounter(2)
		case 0xE3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("INW $%02X\n", file[counter+1])
			incCounter(2)
		case 0xE4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("CPX $%02X\n", file[counter+1])
			incCounter(2)
		case 0xE5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("SBC $%02X\n", file[counter+1])
			incCounter(2)
		case 0xE6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("INC $%02X\n", file[counter+1])
			incCounter(2)
		case 0xE7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("SMB6 $%02X\n", file[counter+1])
			incCounter(2)
		case 0xE9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Immediate)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("SBC #$%02X\n", file[counter+1])
			incCounter(2)
		case 0xF0:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Relative)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("BEQ $%02X\n", (counter+2+int(file[counter+1]))&0xFF)
			incCounter(2)
		case 0xF1:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Zero Page Indirect),Y)\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("SBC ($%02X),Y\n", file[counter+1])
			incCounter(2)
		case 0xF2:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t((Indirect) Z)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("SBC ($%02X)\n", file[counter+1])
			incCounter(2)
		case 0xF5:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("SBC $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0xF6:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page,X)\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("INC $%02X,X\n", file[counter+1])
			incCounter(2)
		case 0xF7:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x\t\t(Zero Page)\t\t\n", memoryAddress, file[counter], file[counter+1])
			}
			fmt.Printf("SMB7 $%02X\n", file[counter+1])
			incCounter(2)
		}

		//3 byte instructions with 2 operands
		switch file[counter] {
		case 0x0C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("TSB $%02x%02x\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x0D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("ORA $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x0E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("ASL $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x0F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BBR0 $%02X, $%02X\n", file[counter+1], file[counter+2])
			incCounter(3)
		case 0x13:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Realtive (word))\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BPL $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x19:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("ORA $%02X%02X,Y\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x1C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("TRB $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x1D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("ORA $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x1E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("ASL $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x1F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BBR1 $%02X, $%02X\n", file[counter+1], file[counter+2])
			incCounter(3)
		case 0x20:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("JSR $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x22:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t((Indirect) Z)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("JSR ($%02X%02X)\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x23:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute X Indirect)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("JSR ($%02X%02X,X)\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x2C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BIT $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x2D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("AND $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x2E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("ROL $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x2F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BBR2 $%02X, $%02X\n", file[counter+1], file[counter+2])
			incCounter(3)
		case 0x33:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Realtive (word))\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BMI $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x34:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BIT $%02X,X\n", file[counter+1])
			incCounter(3)
		case 0x35:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("AND $%02X,X\n", file[counter+1])
			incCounter(3)
		case 0x36:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("ROL $%02X,X\n", file[counter+1])
			incCounter(3)
		case 0x39:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("AND $%02X%02X,Y\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x3C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("AND $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x3D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("AND $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x3E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("ROL $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x3F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BBR3 $%02X, $%02X\n", file[counter+1], file[counter+2])
			incCounter(3)
		case 0x4C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("JMP $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x4D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("EOR $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x4E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("LSR $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x4F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BBR4 $%02X, $%02X\n", file[counter+1], file[counter+2])
			incCounter(3)
		case 0x53:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Realtive (word))\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BVC $%02X\n", file[counter+1])
			incCounter(3)
		case 0x59:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("EOE $%02X%02X,Y\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x5D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("EOR $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x5E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("LSR $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x5F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BBR5 $%02X, $%02X\n", file[counter+1], file[counter+2])
			incCounter(3)
		case 0x63:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Realtive (word))\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BSR $%02X\n", file[counter+1])
			incCounter(3)
		case 0x6C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute Indirect)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("JMP ($%02X%02X)\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x6D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("ADC $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x6E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("ROR $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x6F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BBR6 $%02X, $%02X\n", file[counter+1], file[counter+2])
			incCounter(3)
		case 0x73:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Realtive (word))\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BVS $%02X\n", file[counter+1])
			incCounter(3)
		case 0x79:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("ADC $%02X%02X,Y\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x7C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X Indirect)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("JMP ($%02X%02X,X)\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x7D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("ADC $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x7E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("ROR $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x7F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BBR7 $%02X, $%02X\n", file[counter+1], file[counter+2])
			incCounter(3)
		case 0x83:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Realtive (word))\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BRA $%02X\n", file[counter+1])
			incCounter(3)
		case 0x8B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("STY $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x8C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("STY $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x8D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("STA $%04X\n", file[counter+1]|uint8(int(file[counter+2])<<8))
			incCounter(3)
		case 0x8E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("STX $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x8F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(BBS0- Zero Page, Relative)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BBS0 $%02X, $%02X\n", file[counter+1], file[counter+2])
			incCounter(3)
		case 0x93:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Realtive (word))\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BCC $%02X\n", file[counter+1])
			incCounter(3)
		case 0x99:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("STA $%02X%02X,Y\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x9B:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("STX $%02X%02X,Y\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x9C:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("STZ $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x9D:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("STA $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x9E:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("STZ $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0x9F:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BBS1 $%02X, $%02X\n", file[counter+1], file[counter+2])
			incCounter(3)
		case 0xAB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("LDZ $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xAC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("LDY $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xAD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("LDA $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xAE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("LDX $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xAF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BBS2 $%02X, $%02X\n", file[counter+1], file[counter+2])
			incCounter(3)
		case 0xB3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Realtive (word))\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BCS $%02X\n", file[counter+1])
			incCounter(3)
		case 0xB9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("LDA $%02X%02X,Y\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xBB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("LDZ $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(1)
		case 0xBC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("LDY $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xBD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("LDA $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xBE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("LDX $%02X%02X,Y\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xBF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BBS3 $%02X, $%02X\n", file[counter+1], file[counter+2])
			incCounter(3)
		case 0xCB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("ASW $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xCC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("CPY $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xCD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("CMP $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xCE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("DEC $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xCF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BBS4 $%02X, $%02X\n", file[counter+1], file[counter+2])
			incCounter(3)
		case 0xD3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Realtive (word))\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BNE $%02X\n", file[counter+1])
			incCounter(3)
		case 0xD9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("CMP $%02X%02X,Y\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xDC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("CPZ $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xDD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("CMP $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xDE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("DEC $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xDF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BBS5 $%02X, $%02X\n", file[counter+1], file[counter+2])
			incCounter(3)
		case 0xEB:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("ROW $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xEC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("CPX $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xED:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("SBC $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xEE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("INC $%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xEF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BBS6 $%02X, $%02X\n", file[counter+1], file[counter+2])
			incCounter(3)
		case 0xF3:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Realtive (word))\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BEQ $%02X\n", file[counter+1])
			incCounter(3)
		case 0xF4:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Immediate (word))\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("PHW #$%02X%02X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xF9:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,Y)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("SBC $%02X%02X,Y\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xFC:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("PHW #$%02X%02X\n", file[counter+1], file[counter+2])
			incCounter(3)
		case 0xFD:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("SBC $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xFE:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Absolute,X)\t\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("INC $%02X%02X,X\n", file[counter+2], file[counter+1])
			incCounter(3)
		case 0xFF:
			if printHex {
				fmt.Printf(";; $%04x\t$%02x $%02x $%02x\t(Zero Page, Relative)\n", memoryAddress, file[counter], file[counter+1], file[counter+2])
			}
			fmt.Printf("BBS7 $%02X, $%02X\n", file[counter+1], file[counter+2])
			//incCounter(3)
			incCounter(3)
		}
	}
}
