package main

import (
	"fmt"
	"os"
)

var (
	A            byte        //Accumulator
	X            byte        //X register
	Y            byte        //Y register
	byte_counter uint16      //Program Counter
	SR           byte        //Status Register
	SP           byte        //Stack Pointer
	memory       [65536]byte //Memory
	operand1     byte
	operand2     byte
	lowbyte      int
	highbyte     int
)

func main() {
	//go run . ./c64kernal.bin
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <target_filename>\n", os.Args[0])
		os.Exit(0)
	}

	//Read file
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	disassemble(string(file))
}

func disassemble(file string) {
	for byte_counter := 0; byte_counter < len(file); byte_counter++ {
		//1 byte instructions with no operands
		switch file[byte_counter] {
		case 0x00:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("BRK\n")
		case 0x02:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("CLE\n")
		case 0x03:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("SEE\n")
		case 0x08:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("PHP\n")
		case 0x0A:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("ASL\n")
		case 0x0B:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("TSY\n")
		case 0x18:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("CLC\n")
		case 0x1A:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("INC\n")
		case 0x1B:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("INZ\n")
		case 0x28:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("PLP\n")
		case 0x2A:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("ROL\n")
		case 0x2B:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("TYS\n")
		case 0x38:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("SEC\n")
		case 0x3A:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("DEC\n")
		case 0x3B:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("DEZ\n")
		case 0x40:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("RTI\n")
		case 0x42:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("NEG\n")
		case 0x43:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("ASR\n")
		case 0x48:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("PHA\n")
		case 0x4A:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("LSR\n")
		case 0x4B:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("TAZ\n")
		case 0x58:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("CLI\n")
		case 0x5A:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("PHY\n")
		case 0x5B:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("TAB\n")
		case 0x60:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("RTS\n")
		case 0x68:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("PLA\n")
		case 0x6A:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("ROR\n")
		case 0x6B:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("TZA\n")
		case 0x78:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("SEI\n")
		case 0x7A:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("PLY\n")
		case 0x7B:
			fmt.Printf("$%04x\t$%02x\t\t(TBA - Absolute,Y)\t\t\t", byte_counter, file[byte_counter])
			fmt.Printf("TBA\n")
		case 0x88:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("DEY\n")
		case 0x8A:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("TXA\n")
		case 0x98:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("TYA\n")
		case 0x9A:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("TXS\n")
		case 0xA8:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("TAY\n")
		case 0xAA:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("TAX\n")
		case 0xB8:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("CLV\n")
		case 0xBA:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("TSX\n")
		case 0xC8:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("INY\n")
		case 0xCA:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("DEX\n")
		case 0xD8:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("CLD\n")
		case 0xDA:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("PHX\n")
		case 0xDB:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("PHZ\n")
		case 0xE8:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("INX\n")
		case 0xEA:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("NOP\n")
		case 0xF8:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("SED\n")
		case 0xFA:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("PLX\n")
		case 0xFB:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", byte_counter, file[byte_counter])
			fmt.Printf("PLZ\n")
		}

		//2 byte instructions with 1 operand
		switch file[byte_counter] {
		case 0x05:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("ORA $%02x\n", file[byte_counter+1])
			byte_counter++
		case 0x06:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("ASL $%02x\n", file[byte_counter+1])
			byte_counter++
		case 0x07:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("RMB0 $%02x\n", file[byte_counter+1])
			byte_counter++
		case 0x09:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("ORA #$%02x\n", file[byte_counter+1])
			byte_counter++
		case 0x10:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Relative)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("BPL $%02X\n", (byte_counter+2+int(file[byte_counter+1]))&0xFF)
			byte_counter += 2
		case 0x11:
			fmt.Printf("$%04x\t$%02x $%02x\t\t((0Page),Indirect Y)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("ORA ($%02x),Y\n", file[byte_counter+1])
			byte_counter++
		case 0x12:
			fmt.Printf("$%04x\t$%02x $%02x\t\t((0Page),Indirect Z)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("ORA ($%02x),Z\n", file[byte_counter+1])
			byte_counter++
		case 0x14:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("TRB $%02x\n", file[byte_counter+1])
			byte_counter++
		case 0x15:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page,X)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("ORA $%02x,X\n", file[byte_counter+1])
			byte_counter++
		case 0x16:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ASL - ZeroPage,X)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("ASL $%02X,X\n", file[byte_counter+1])
			byte_counter++
		case 0x17:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("RMB1 $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x21:
			fmt.Printf("$%04x\t$%02x $%02x\t\t((0Page,Indirect))\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("AND ($%02X,X)\n", file[byte_counter+1])
			byte_counter++
		case 0x24:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("BIT $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x25:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("AND $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x26:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("ROL $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x27:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("RMB2 $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x29:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("AND #$%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x30:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Relative)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("BMI $%02X\n", (byte_counter + 2 + int(file[byte_counter+1])))
			byte_counter++
		case 0x31:
			fmt.Printf("$%04x\t$%02x $%02x\t\t((0Page Indirect),Y)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("AND ($%02X),Y\n", file[byte_counter+1])
			byte_counter++
		case 0x32:
			fmt.Printf("$%04x\t$%02x $%02x\t\t((Indirect),Z)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("AND ($%02X),Z\n", file[byte_counter+1])
			byte_counter++
		case 0x34:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(X ZeroPage)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("BIT $%02X,X\n", file[byte_counter+1])
			byte_counter++
		case 0x35:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(X ZeroPage)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("AND $%02X,X\n", file[byte_counter+1])
			byte_counter++
		case 0x36:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ROL - X ZeroPage)\n", byte_counter, file[byte_counter], file[byte_counter+1])
			byte_counter++
		case 0x37:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("RMB3 $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x41:
			fmt.Printf("$%04x\t$%02x $%02x\t\t((X ZeroPage, Indirect))\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("EOR ($%02X,X)\n", file[byte_counter+1])
			byte_counter++
		case 0x44:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("ASR $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x45:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("EOR $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x46:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("LSR $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x47:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("RMB4 $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x49:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("EOR #$%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x50:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BVC - Relative)\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("** UNIMPLEMENTED OPCODE **\n")
			byte_counter++
		case 0x51:
			fmt.Printf("$%04x\t$%02x $%02x\t\t((0Page Indirect),Y)\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("EOR ($%02X),Y\n", file[byte_counter+1])
			byte_counter++
		case 0x52:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR - (Indirect),Z)\t", byte_counter, file[byte_counter], file[byte_counter+1])
			//fmt.Printf("EOR ($%02X),Z\n", file[byte_counter+1])
			fmt.Printf("** UNIMPLEMENTED OPCODE **\n")
			byte_counter++
		case 0x54:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ASR - X ZeroPage)\n", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("** UNIMPLEMENTED OPCODE **\n")
			byte_counter++
		case 0x55:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(X ZeroPage)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("EOR $%02X,X\n", file[byte_counter+1])
			byte_counter++
		case 0x56:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(X ZeroPage)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("LSR $%02X,X\n", file[byte_counter+1])
			byte_counter++
		case 0x57:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(RMB5 - 0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("** UNIMPLEMENTED OPCODE **\n")
			byte_counter++
		case 0x61:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(X ZeroPage Indirect)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("ADC ($%02X,X)\n", file[byte_counter+1])
			byte_counter++
		case 0x62:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			//fmt.Printf("RTN #$%02X\n", file[byte_counter+1])
			fmt.Printf("** UNIMPLEMENTED OPCODE **\n")
			byte_counter++
		case 0x64:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("STZ $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x65:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("ADC $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x66:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("ROR $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x67:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("RMB6 $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x69:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("ADC #$%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x70:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Relative)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("BVS $%04X\n", byte_counter+2+int(file[byte_counter+1]))
			byte_counter++
		case 0x71:
			fmt.Printf("$%04x\t$%02x $%02x\t\t((0Page Indirect),Y)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("ADC ($%02X),Y\n", file[byte_counter+1])
			byte_counter++
		case 0x72:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Indirect,Z)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			//fmt.Printf("ADC ($%02X),Z\n", file[byte_counter+1])
			fmt.Printf("** UNIMPLEMENTED OPCODE **\n")
			byte_counter++
		case 0x74:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STZ - ZeroPage,X)\n", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("** UNIMPLEMENTED OPCODE **\n")
			byte_counter++
		case 0x75:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page,X)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("ADC $%02X,X\n", file[byte_counter+1])
			byte_counter++
		case 0x76:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page,X)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("ROR $%02X,X\n", file[byte_counter+1])
			byte_counter++
		case 0x77:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("RMB7 $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x80:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BRA - Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("** UNIMPLEMENTED OPCODE **\n")
			byte_counter++
		case 0x81:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(X 0Page Indirect)\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("STA ($%02X,X)\n", file[byte_counter+1])
			byte_counter++
		case 0x82:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA - Stack Relative Indirect, Y)\n", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("** UNIMPLEMENTED OPCODE **\n")
			byte_counter++
		case 0x84:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("STY $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x85:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("STA $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x86:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("STX $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x87:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("SMB0 $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x89:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("BIT #$%02X\n", file[byte_counter+1])
			byte_counter++
		case 0x90:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Relative)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("BCC $%02X\n", byte_counter+2+int((file[byte_counter+1])))
			byte_counter++
		case 0x91:
			fmt.Printf("$%04x\t$%02x $%02x\t\t((0Page Indirect),Y)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("STA ($%02X),Y\n", file[byte_counter+1])
			byte_counter++
		case 0x92:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA - Indirect,Z)\n", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("** UNIMPLEMENTED OPCODE **\n")
			byte_counter++
		case 0x94:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page,X)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("STY $%02X,X\n", file[byte_counter+1])
			byte_counter++
		case 0x95:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page,X)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("STA $%02X,X\n", file[byte_counter+1])
			byte_counter++
		case 0x96:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page,Y)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("STX $%02X,Y\n", file[byte_counter+1])
			byte_counter++
		case 0x97:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("SMB1 $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xA0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("LDY #$%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xA1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(X 0Page Indirect)\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("LDA ($%02X,X)\n", file[byte_counter+1])
			byte_counter++
		case 0xA2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("LDX #$%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xA3:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("LDZ #$%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xA4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("LDY $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xA5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("LDA $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xA6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("LDX $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xA7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("SMB2 $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xA9:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("LDA #$%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xB0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Relative)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("BCS $%04X\n", byte_counter+2+int(file[byte_counter+1]))
			byte_counter++
		case 0xB1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t((0Page Indirect),Y)\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("LDA ($%02X),Y\n", file[byte_counter+1])
			byte_counter++
		case 0xB2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA - 0Page,Indirect)\n", byte_counter, file[byte_counter], file[byte_counter+1])
			//fmt.Printf("LDA ($%02X)\n", file[byte_counter+1])
			fmt.Printf("** UNIMPLEMENTED OPCODE **\n")
			byte_counter++
		case 0xB4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page,X)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("LDY $%02X,X\n", file[byte_counter+1])
			byte_counter++
		case 0xB5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page,X)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("LDA $%02X,X\n", file[byte_counter+1])
			byte_counter++
		case 0xB6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page,Y)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("LDX $%02X,Y\n", file[byte_counter+1])
			byte_counter++
		case 0xB7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("SMB3 $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xC0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("CPY #$%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xC1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(X 0Page Indirect)\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("CMP ($%02X,X)\n", file[byte_counter+1])
			byte_counter++
		case 0xC2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("CPZ #$%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xC3:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("DEW $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xC4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("CPY $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xC5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("CMP $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xC6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("DEC $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xC7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("SMB4 $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xC9:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("CMP #$%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xD0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BNE - Relative)\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("** UNIMPLEMENTED OPCODE **\n")
			byte_counter++
		case 0xD1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t((0Page Indirect),Y)\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("CMP ($%02X),Y\n", file[byte_counter+1])
			byte_counter++
		case 0xD2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP - (Indirect) Z)\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("** UNIMPLEMENTED OPCODE **\n")
			byte_counter++
		case 0xD4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			//fmt.Printf("CPZ $%02x\n", file[byte_counter+1])
			fmt.Printf("** UNIMPLEMENTED OPCODE **\n")
			byte_counter++
		case 0xD5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page,X)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("CMP $%02X,X\n", file[byte_counter+1])
			byte_counter++
		case 0xD6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page,X)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("DEC $%02X,X\n", file[byte_counter+1])
			byte_counter++
		case 0xD7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("SMB5 $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xE0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("CPX #$%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xE1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(X 0Page Indirect)\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("SBC ($%02X,X)\n", file[byte_counter+1])
			byte_counter++
		case 0xE2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("LDA #$%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xE3:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("INW $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xE4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("CPX $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xE5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("SBC $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xE6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("INC $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xE7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("SMB6 $%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xE9:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("SBC #$%02X\n", file[byte_counter+1])
			byte_counter++
		case 0xF0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BEQ - Relative)\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("BEQ $%02X\n", (byte_counter+2+int(file[byte_counter+1]))&0xFF)
			byte_counter++
		case 0xF1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t((0Page Indirect),Y)\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("SBC ($%02X),Y\n", file[byte_counter+1])
			byte_counter++
		case 0xF2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t((Indirect) Z)\t", byte_counter, file[byte_counter], file[byte_counter+1])
			//fmt.Printf("SBC ($%02X)\n", file[byte_counter+1])
			fmt.Printf("** UNIMPLEMENTED OPCODE **\n")
			byte_counter++
		case 0xF5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page,X)\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("SBC $%02X,X\n", file[byte_counter+1])
			byte_counter++
		case 0xF6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page,X)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("INC $%02X,X\n", file[byte_counter+1])
			byte_counter++
		case 0xF7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1])
			fmt.Printf("SMB7 $%02X\n", file[byte_counter+1])
			byte_counter++
		}

		//3 byte instructions with 2 operands
		switch file[byte_counter] {
		case 0x0C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(Absolute)\t\t", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			fmt.Printf("TSB $%02x%02x\n", file[byte_counter+2], file[byte_counter+1])
			byte_counter += 2
		case 0x0D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ORA - Absolute)\t", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			fmt.Printf("ORA $%02x%02x\n", file[byte_counter+2], file[byte_counter+1])
			byte_counter += 2
		case 0x0E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(Absolute)\t\t", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			fmt.Printf("ASL $%02x%02x\n", file[byte_counter+2], file[byte_counter+1])
			byte_counter += 2
		case 0x0F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR0 - ZeroPage, Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x13:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BPL - Relative (word))\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x19:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ORA - Absolute,Y)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x1C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(Absolute)\t\t", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			fmt.Printf("TRB $%02x%02x\n", file[byte_counter+2], file[byte_counter+1])
			byte_counter += 2
		case 0x1D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ORA - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x1E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ASL - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x1F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR1 - ZeroPage, Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x20:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(Absolute)\t\t", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			fmt.Printf("JSR $%02X%02X\n", file[byte_counter+2], file[byte_counter+1])
			byte_counter += 2
		case 0x21:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND - (0Page Indirect,X))\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x22:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JSR - (Indirect) Z)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x23:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JSR - Absolute X Indirect)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x2C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BIT - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x2D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(Absolute)\t\t", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			fmt.Printf("AND $%02X%02X\n", file[byte_counter+2], file[byte_counter+1])
			byte_counter += 2
		case 0x2E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROL - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x2F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR2 - ZeroPage, Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x33:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BMI - Relative (word))\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x34:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BIT - ZeroPage,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x35:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND - ZeroPage,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x36:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROL - ZeroPage,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x37:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(0Page)\t\t", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			fmt.Printf("RMB3 $%02X\n", file[byte_counter+1])
			byte_counter += 2
		case 0x39:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND - Absolute,Y)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x3C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BIT - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x3D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x3E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROL - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x3F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR3 - ZeroPage, Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x4C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(Absolute)\t\t", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			fmt.Printf("JMP $%02X%02X\n", file[byte_counter+2], file[byte_counter+1])
			byte_counter += 2
		case 0x4D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(EOR - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x4E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(Absolute)\t\t", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			fmt.Printf("LSR $%02X%02X\n", file[byte_counter+2], file[byte_counter+1])
			byte_counter += 2
		case 0x4F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR4 - ZeroPage, Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x53:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BVC - Relative (word))\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x59:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(EOR - Absolute,Y)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x5D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(EOR - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x5E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LSR - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x5F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR5 - ZeroPage, Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x63:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BSR - Relative (word))\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x6C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JMP - Absolute Indirect)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x6D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ADC - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x6E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(Absolute)\t\t", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			fmt.Printf("ROR $%02X%02X\n", file[byte_counter+2], file[byte_counter+1])
			byte_counter += 2
		case 0x6F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR6 - ZeroPage, Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x73:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BVS - Relative (word))\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x79:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ADC - Absolute,Y)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x7C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JMP - Absolute,X Indirect)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x7D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ADC - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x7E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROR - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x7F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR7 - ZeroPage, Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x83:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BRA - Relative (word))\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x8B:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STY - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x8C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STY - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x8D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(Absolute)\t\t", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			fmt.Printf("STA $%04x\n", file[byte_counter+1]|(file[byte_counter+2]<<8))
			byte_counter += 2
		case 0x8E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STX - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x8F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS0- ZeroPage, Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x93:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BCC - Relative (word))\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x99:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STA - Absolute,Y)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x9B:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STX - Absolute,Y)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x9C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STZ - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x9D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STA - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x9E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STZ - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0x9F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS1 - ZeroPage, Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xAB:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDZ - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xAC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDY - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xAD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDA - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xAE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDX - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xAF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS2 - ZeroPage, Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xB3:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BCS - Relative (word))\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xB9:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDA - Absolute,Y)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xBC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDY - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xBD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDA - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xBE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDX - Absolute,Y)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xBF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS3 - ZeroPage, Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xCB:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ASW - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xCC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CPY - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xCD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CMP - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xCE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(DEC - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xCF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS4 - ZeroPage, Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xD3:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BNE - Relative (word))\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xD9:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CMP - Absolute,Y)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xDC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CPZ - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xDD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CMP - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xDE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(DEC - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xDF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS5 - ZeroPage, Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xEB:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROW - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xEC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CPX - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xED:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SBC - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xEE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(INC - Absolute)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xEF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS6 - ZeroPage, Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xF3:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BEQ - Relative (word))\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xF4:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(PHW - Immediate (word))\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xF9:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SBC - Absolute,Y)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xFC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(Absolute)\t\t", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			fmt.Printf("PHW #$%02X%02X\n", file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xFD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SBC - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xFE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(INC - Absolute,X)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		case 0xFF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS7 - ZeroPage, Relative)\n", byte_counter, file[byte_counter], file[byte_counter+1], file[byte_counter+2])
			byte_counter += 2
		}
	}
}
