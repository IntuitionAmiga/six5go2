package main

import (
	"fmt"
	"os"
)

var (
	A        byte        //Accumulator
	X        byte        //X register
	Y        byte        //Y register
	PC       uint16      //Program Counter
	SR       byte        //Status Register
	SP       byte        //Stack Pointer
	memory   [65536]byte //Memory
	operand1 byte
	operand2 byte
	lowbyte  int
	highbyte int
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
	for PC := 0; PC < len(file); PC++ {
		//1 byte instructions with no operands
		switch file[PC] {
		case 0x00:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("BRK\n")
		case 0x02:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("CLE\n")
		case 0x03:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("SEE\n")
		case 0x08:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("PHP\n")
		case 0x0A:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", PC, file[PC])
			fmt.Printf("ASL\n")
		case 0x0B:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("TSY\n")
		case 0x18:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("CLC\n")
		case 0x1A:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", PC, file[PC])
			fmt.Printf("INC\n")
		case 0x1B:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("INZ\n")
		case 0x28:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("PLP\n")
		case 0x2A:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", PC, file[PC])
			fmt.Printf("ROL\n")
		case 0x2B:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("TYS\n")
		case 0x38:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", PC, file[PC])
			fmt.Printf("SEC\n")
		case 0x3A:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", PC, file[PC])
			fmt.Printf("DEC\n")
		case 0x3B:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("DEZ\n")
		case 0x40:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("RTI\n")
		case 0x42:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", PC, file[PC])
			fmt.Printf("NEG\n")
		case 0x43:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", PC, file[PC])
			fmt.Printf("ASR\n")
		case 0x48:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("PHA\n")
		case 0x4A:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", PC, file[PC])
			fmt.Printf("LSR\n")
		case 0x4B:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("TAZ\n")
		case 0x58:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("CLI\n")
		case 0x5A:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("PHY\n")
		case 0x5B:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("TAB\n")
		case 0x60:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("RTS\n")
		case 0x68:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("PLA\n")
		case 0x6A:
			fmt.Printf("$%04x\t$%02x\t\t(Accumulator)\t\t", PC, file[PC])
			fmt.Printf("ROR\n")
		case 0x6B:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("TZA\n")
		case 0x78:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("SEI\n")
		case 0x7A:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("PLY\n")
		case 0x7B:
			fmt.Printf("$%04x\t$%02x\t\t(TBA - Absolute,Y)\t\t\t", PC, file[PC])
			fmt.Printf("TBA\n")
		case 0x88:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("DEY\n")
		case 0x8A:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("TXA\n")
		case 0x98:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("TYA\n")
		case 0x9A:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("TXS\n")
		case 0xA8:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("TAY\n")
		case 0xAA:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("TAX\n")
		case 0xB8:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("CLV\n")
		case 0xBA:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("TSX\n")
		case 0xC8:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("INY\n")
		case 0xCA:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("DEX\n")
		case 0xD8:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("CLD\n")
		case 0xDA:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("PHX\n")
		case 0xDB:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("PHZ\n")
		case 0xE8:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("INX\n")
		case 0xEA:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("NOP\n")
		case 0xF8:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("SED\n")
		case 0xFA:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("PLX\n")
		case 0xFB:
			fmt.Printf("$%04x\t$%02x\t\t(Implied)\t\t", PC, file[PC])
			fmt.Printf("PLZ\n")
		}

		//2 byte instructions with 1 operand
		switch file[PC] {
		case 0x05:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("ORA $%02x\n", file[PC+1])
			PC++
		case 0x06:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("ASL $%02x\n", file[PC+1])
			PC++
		case 0x07:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("RMB0 $%02x\n", file[PC+1])
			PC++
		case 0x09:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("ORA #$%02x\n", file[PC+1])
			PC++
		case 0x10:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BPL - Relative)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x11:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ORA - (ZeroPage),Indirect Y)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x12:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ORA - (ZeroPage),Indirect Z)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x14:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("TRB $%02x\n", file[PC+1])
			PC++
		case 0x15:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ORA - ZeroPage,X)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("ORA $%02x\n", file[PC+1])
			PC++
		case 0x16:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ASL - ZeroPage,X)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("ASL $%02x\n", file[PC+1])
			PC++
		case 0x17:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("RMB1 $%02x\n", file[PC+1])
			PC++
		case 0x21:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(AND - (ZeroPage,Indirect))\n", PC, file[PC], file[PC+1])
			PC++
		case 0x24:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("BIT $%02x\n", file[PC+1])
			PC++
		case 0x25:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("AND $%02x\n", file[PC+1])
			PC++
		case 0x26:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("ROL $%02x\n", file[PC+1])
			PC++
		case 0x27:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("RMB2 $%02x\n", file[PC+1])
			PC++
		case 0x29:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("AND #$%02x\n", file[PC+1])
			PC++
		case 0x30:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BMI - Relative)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x31:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(AND (ZeroPage Indirect),Y)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x32:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(AND (Indirect),Z)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x34:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BIT - X ZeroPage)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x35:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(AND - X ZeroPage)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x36:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ROL - X ZeroPage)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x37:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("RMB3 $%02x\n", file[PC+1])
			PC++
		case 0x41:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR - (X ZeroPage, Indirect))\n", PC, file[PC], file[PC+1])
			PC++
		case 0x44:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("ASR $%02x\n", file[PC+1])
			PC++
		case 0x45:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("EOR $%02x\n", file[PC+1])
			PC++
		case 0x46:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("LSR $%02x\n", file[PC+1])
			PC++
		case 0x47:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("RMB4 $%02x\n", file[PC+1])
			PC++
		case 0x49:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("EOR #$%02x\n", file[PC+1])
			PC++
		case 0x50:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BVC - Relative)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x51:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR (ZeroPage Indirect),Y)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x52:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR - (Indirect),Z)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x54:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ASR - X ZeroPage)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x55:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR - X ZeroPage)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x56:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LSR - X ZeroPage)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x57:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			PC++
		case 0x61:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ADC - X ZeroPage Indirect)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x62:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("RTN #$%02x\n", file[PC+1])
			PC++
		case 0x64:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("STZ $%02x\n", file[PC+1])
			PC++
		case 0x65:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("ADC $%02x\n", file[PC+1])
			PC++
		case 0x66:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("ROR $%02x\n", file[PC+1])
			PC++
		case 0x67:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("RMB6 $%02x\n", file[PC+1])
			PC++
		case 0x69:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("ADC #$%02x\n", file[PC+1])
			PC++
		case 0x70:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BVS - Relative)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x71:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ADC (ZeroPage Indirect),Y)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x72:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ADC - Indirect,Z)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x74:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STZ - ZeroPage,X)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x75:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ADC - ZeroPage,X)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x76:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ROR - ZeroPage,X)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x77:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(RMB7 - ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("RMB7 $%02x\n", file[PC+1])
			PC++
		case 0x80:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BRA - Relative)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x81:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA - X ZeroPage Indirect)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x82:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA - Stack Relative Indirect, Y)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x84:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("STY $%02x\n", file[PC+1])
			PC++
		case 0x85:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("STA $%02X\n", file[PC+1])
			PC++
		case 0x86:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("STX $%02x\n", file[PC+1])
			PC++
		case 0x87:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("SMB0 $%02x\n", file[PC+1])
			PC++
		case 0x89:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("BIT #$%02X\n", file[PC+1])
			PC++
		case 0x90:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Relative)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("BCC $%02X\n", PC+2+int((file[PC+1])))
			PC++
		case 0x91:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA - (ZeroPage Indirect),Y)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x92:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA - Indirect,Z)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x94:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STY - ZeroPage,X)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x95:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA - ZeroPage,X)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x96:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STX - ZeroPage,Y)\n", PC, file[PC], file[PC+1])
			PC++
		case 0x97:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("SMB1 $%02x\n", file[PC+1])
			PC++
		case 0xA0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("LDY #$%02X\n", file[PC+1])
			PC++
		case 0xA1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA - X ZeroPage Indirect)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xA2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("LDX #$%02X\n", file[PC+1])
			PC++
		case 0xA3:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("LDZ #$%02X\n", file[PC+1])
			PC++
		case 0xA4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("LDY $%02x\n", file[PC+1])
			PC++
		case 0xA5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("LDA $%02X\n", file[PC+1])
			PC++
		case 0xA6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("LDX $%02x\n", file[PC+1])
			PC++
		case 0xA7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("SMB2 $%02x\n", file[PC+1])
			PC++
		case 0xA9:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("LDA #$%02X\n", file[PC+1])
			PC++
		case 0xB0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BCS - Relative)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xB1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA - (ZeroPage Indirect),Y)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xB2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA - ZeroPage,Indirect)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xB4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDY - ZeroPage,X)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xB5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA - ZeroPage,X)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xB6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDX - ZeroPage,Y)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xB7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("SMB3 $%02x\n", file[PC+1])
			PC++
		case 0xC0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("CPY #$%02X\n", file[PC+1])
			PC++
		case 0xC1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP - X ZeroPage Indirect)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xC2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("CPZ #$%02X\n", file[PC+1])
			PC++
		case 0xC3:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("DEW $%02x\n", file[PC+1])
			PC++
		case 0xC4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("CPY $%02x\n", file[PC+1])
			PC++
		case 0xC5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("CMP $%02x\n", file[PC+1])
			PC++
		case 0xC6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("DEC $%02x\n", file[PC+1])
			PC++
		case 0xC7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("SMB4 $%02x\n", file[PC+1])
			PC++
		case 0xC9:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("CMP #$%02X\n", file[PC+1])
			PC++
		case 0xD0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BNE - Relative)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xD1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP - (ZeroPage Indirect),Y)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xD2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP - (Indirect) Z)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xD4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("CPZ $%02x\n", file[PC+1])
			PC++
		case 0xD5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP - ZeroPage,X)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xD6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(DEC - ZeroPage,X)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xD7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("SMB5 $%02x\n", file[PC+1])
			PC++
		case 0xE0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("CPX #$%02X\n", file[PC+1])
			PC++
		case 0xE1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SBC - X ZeroPage Indirect)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xE2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("LDA #$%02X\n", file[PC+1])
			PC++
		case 0xE3:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("INW $%02x\n", file[PC+1])
			PC++
		case 0xE4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("CPX $%02x\n", file[PC+1])
			PC++
		case 0xE5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("SBC $%02x\n", file[PC+1])
			PC++
		case 0xE6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("INC $%02x\n", file[PC+1])
			PC++
		case 0xE7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("SMB6 $%02x\n", file[PC+1])
			PC++
		case 0xE9:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(Immediate)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("SBC #$%02X\n", file[PC+1])
			PC++
		case 0xF0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BEQ - Relative)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xF1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SBC - (ZeroPage Indirect),Y)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xF2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SBC - (Indirect) Z)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xF5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SBC - ZeroPage,X)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xF6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(INC - ZeroPage,X)\n", PC, file[PC], file[PC+1])
			PC++
		case 0xF7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ZeroPage)\t\t", PC, file[PC], file[PC+1])
			fmt.Printf("SMB7 $%02x\n", file[PC+1])
			PC++
		}

		//3 byte instructions with 2 operands
		switch file[PC] {
		case 0x0C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(TSB - Absolute)\t", PC, file[PC], file[PC+1], file[PC+2])
			fmt.Printf("TSB $%02x%02x\n", file[PC+2], file[PC+1])
			PC += 2
		case 0x0D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ORA - Absolute)\t", PC, file[PC], file[PC+1], file[PC+2])
			fmt.Printf("ORA $%02x%02x\n", file[PC+2], file[PC+1])
			PC += 2
		case 0x0E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			fmt.Printf("ASL $%02x%02x\n", file[PC+2], file[PC+1])
			PC += 2
		case 0x0F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR0 - ZeroPage, Relative)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x13:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BPL - Relative (word))\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x19:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ORA - Absolute,Y)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x1C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(Absolute)\t", PC, file[PC], file[PC+1], file[PC+2])
			fmt.Printf("TRB $%02x%02x\n", file[PC+2], file[PC+1])
			PC += 2
		case 0x1D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ORA - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x1E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ASL - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x1F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR1 - ZeroPage, Relative)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x20:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(Absolute)\t\t", PC, file[PC], file[PC+1], file[PC+2])
			fmt.Printf("JSR $%02X%02X\n", file[PC+2], file[PC+1])
			PC += 2
		case 0x21:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND - (ZeroPage Indirect,X))\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x22:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JSR - (Indirect) Z)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x23:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JSR - Absolute X Indirect)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x2C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BIT - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x2D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x2E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROL - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x2F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR2 - ZeroPage, Relative)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x33:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BMI - Relative (word))\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x34:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BIT - ZeroPage,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x35:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND - ZeroPage,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x36:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROL - ZeroPage,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x37:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ZeroPage)\t\t", PC, file[PC], file[PC+1], file[PC+2])
			fmt.Printf("RMB3 $%02X\n", file[PC+1])
			PC += 2
		case 0x39:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND - Absolute,Y)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x3C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BIT - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x3D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x3E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROL - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x3F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR3 - ZeroPage, Relative)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x4C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JMP - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x4D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(EOR - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x4E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LSR - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x4F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR4 - ZeroPage, Relative)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x53:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BVC - Relative (word))\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x59:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(EOR - Absolute,Y)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x5D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(EOR - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x5E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LSR - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x5F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR5 - ZeroPage, Relative)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x63:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BSR - Relative (word))\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x6C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JMP - Absolute Indirect)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x6D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ADC - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x6E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROR - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x6F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR6 - ZeroPage, Relative)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x73:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BVS - Relative (word))\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x79:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ADC - Absolute,Y)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x7C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JMP - Absolute,X Indirect)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x7D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ADC - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x7E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROR - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x7F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR7 - ZeroPage, Relative)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x83:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BRA - Relative (word))\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x8B:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STY - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x8C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STY - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x8D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(Absolute)\t", PC, file[PC], file[PC+1], file[PC+2])
			fmt.Printf("STA $%04x\n", file[PC+1]|(file[PC+2]<<8))
			PC += 2
		case 0x8E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STX - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x8F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS0- ZeroPage, Relative)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x93:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BCC - Relative (word))\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x99:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STA - Absolute,Y)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x9B:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STX - Absolute,Y)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x9C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STZ - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x9D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STA - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x9E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STZ - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0x9F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS1 - ZeroPage, Relative)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xAB:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDZ - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xAC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDY - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xAD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDA - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xAE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDX - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xAF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS2 - ZeroPage, Relative)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xB3:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BCS - Relative (word))\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xB9:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDA - Absolute,Y)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xBC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDY - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xBD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDA - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xBE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDX - Absolute,Y)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xBF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS3 - ZeroPage, Relative)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xCB:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ASW - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xCC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CPY - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xCD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CMP - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xCE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(DEC - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xCF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS4 - ZeroPage, Relative)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xD3:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BNE - Relative (word))\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xD9:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CMP - Absolute,Y)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xDC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CPZ - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xDD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CMP - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xDE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(DEC - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xDF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS5 - ZeroPage, Relative)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xEB:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROW - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xEC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CPX - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xED:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SBC - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xEE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(INC - Absolute)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xEF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS6 - ZeroPage, Relative)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xF3:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BEQ - Relative (word))\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xF4:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(PHW - Immediate (word))\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xF9:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SBC - Absolute,Y)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xFC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(PHW - Absolute\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xFD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SBC - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xFE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(INC - Absolute,X)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		case 0xFF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS7 - ZeroPage, Relative)\n", PC, file[PC], file[PC+1], file[PC+2])
			PC += 2
		}
	}
}
