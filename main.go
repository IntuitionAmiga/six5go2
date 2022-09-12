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
)

func main() {
	//Use apple1basic.bin as the program to disassemble for now
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <target_filename> \n", os.Args[0])
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
	//Run the program counter from zero to end of table and decode the fetched opcode
	for PC := 0; PC < len(file); PC++ {
		//1 byte instructions with no operands
		switch file[PC] {
		case 0x00:
			fmt.Printf("$%04x\t$%02x\t\t(BRK)\n", PC, file[PC])
		case 0x02:
			fmt.Printf("$%04x\t$%02x\t\t(KIL)\n", PC, file[PC])
		case 0x03:
			fmt.Printf("$%04x\t$%02x\t\t(SLO (Indirect,X))\n", PC, file[PC])
		case 0x08:
			fmt.Printf("$%04x\t$%02x\t\t(PHP)\n", PC, file[PC])
		case 0x0A:
			fmt.Printf("$%04x\t$%02x\t\t(ASL Accumulator)\n", PC, file[PC])
		case 0x0B:
			fmt.Printf("$%04x\t$%02x\t\t(ANC Immediate)\n", PC, file[PC])
		case 0x18:
			fmt.Printf("$%04x\t$%02x\t\t(CLC)\n", PC, file[PC])
		case 0x1A:
			fmt.Printf("$%04x\t$%02x\t\t(NOP)\n", PC, file[PC])
		case 0x1B:
			fmt.Printf("$%04x\t$%02x\t\t(SLO Absolute,Y)\n", PC, file[PC])
		case 0x28:
			fmt.Printf("$%04x\t$%02x\t\t(PLP)\n", PC, file[PC])
		case 0x2A:
			fmt.Printf("$%04x\t$%02x\t\t(ROL Accumulator)\n", PC, file[PC])
		case 0x2B:
			fmt.Printf("$%04x\t$%02x\t\t(ANC Immediate)\n", PC, file[PC])
		case 0x38:
			fmt.Printf("$%04x\t$%02x\t\t(SEC)\n", PC, file[PC])
		case 0x3A:
			fmt.Printf("$%04x\t$%02x\t\t(NOP)\n", PC, file[PC])
		case 0x3B:
			fmt.Printf("$%04x\t$%02x\t\t(RLA Absolute,Y)\n", PC, file[PC])
		case 0x40:
			fmt.Printf("$%04x\t$%02x\t\t(RTI)\n", PC, file[PC])
		case 0x42:
			fmt.Printf("$%04x\t$%02x\t\t(KIL)\n", PC, file[PC])
		case 0x43:
			fmt.Printf("$%04x\t$%02x\t\t(SRE (Indirect,X))\n", PC, file[PC])
		case 0x48:
			fmt.Printf("$%04x\t$%02x\t\t(PHA)\n", PC, file[PC])
		case 0x4A:
			fmt.Printf("$%04x\t$%02x\t\t(LSR Accumulator)\n", PC, file[PC])
		case 0x4B:
			fmt.Printf("$%04x\t$%02x\t\t(ALR Immediate)\n", PC, file[PC])
		case 0x58:
			fmt.Printf("$%04x\t$%02x\t\t(CLI)\n", PC, file[PC])
		case 0x5A:
			fmt.Printf("$%04x\t$%02x\t\t(NOP)\n", PC, file[PC])
		case 0x5B:
			fmt.Printf("$%04x\t$%02x\t\t(SRE Absolute,Y)\n", PC, file[PC])
		case 0x60:
			fmt.Printf("$%04x\t$%02x\t\t(RTS)\n", PC, file[PC])
		case 0x68:
			fmt.Printf("$%04x\t$%02x\t\t(PLA)\n", PC, file[PC])
		case 0x6A:
			fmt.Printf("$%04x\t$%02x\t\t(ROR Accumulator)\n", PC, file[PC])
		case 0x6B:
			fmt.Printf("$%04x\t$%02x\t\t(ARR Immediate)\n", PC, file[PC])
		case 0x78:
			fmt.Printf("$%04x\t$%02x\t\t(SEI)\n", PC, file[PC])
		case 0x7A:
			fmt.Printf("$%04x\t$%02x\t\t(NOP)\n", PC, file[PC])
		case 0x7B:
			fmt.Printf("$%04x\t$%02x\t\t(TBA Absolute,Y)\n", PC, file[PC])
		case 0x88:
			fmt.Printf("$%04x\t$%02x\t\t(DEY)\n", PC, file[PC])
		case 0x8A:
			fmt.Printf("$%04x\t$%02x\t\t(TXA)\n", PC, file[PC])
		case 0x98:
			fmt.Printf("$%04x\t$%02x\t\t(TYA)\n", PC, file[PC])
		case 0x9A:
			fmt.Printf("$%04x\t$%02x\t\t(TXS)\n", PC, file[PC])
		case 0xA8:
			fmt.Printf("$%04x\t$%02x\t\t(TAY)\n", PC, file[PC])
		case 0xAA:
			fmt.Printf("$%04x\t$%02x\t\t(TAX)\n", PC, file[PC])
		case 0xB8:
			fmt.Printf("$%04x\t$%02x\t\t(CLV)\n", PC, file[PC])
		case 0xBA:
			fmt.Printf("$%04x\t$%02x\t\t(TSX)\n", PC, file[PC])
		case 0xC8:
			fmt.Printf("$%04x\t$%02x\t\t(INY)\n", PC, file[PC])
		case 0xCA:
			fmt.Printf("$%04x\t$%02x\t\t(DEX)\n", PC, file[PC])
		case 0xD8:
			fmt.Printf("$%04x\t$%02x\t\t(CLD)\n", PC, file[PC])
		case 0xDA:
			fmt.Printf("$%04x\t$%02x\t\t(PHX)\n", PC, file[PC])
		case 0xDB:
			fmt.Printf("$%04x\t$%02x\t\t(PHZ)\n", PC, file[PC])
		case 0xE8:
			fmt.Printf("$%04x\t$%02x\t\t(INX)\n", PC, file[PC])
		case 0xEA:
			fmt.Printf("$%04x\t$%02x\t\t(NOP)\n", PC, file[PC])
		case 0xF8:
			fmt.Printf("$%04x\t$%02x\t\t(SED)\n", PC, file[PC])
		case 0xFA:
			fmt.Printf("$%04x\t$%02x\t\t(PLX)\n", PC, file[PC])
		case 0xFB:
			fmt.Printf("$%04x\t$%02x\t\t(PLZ)\n", PC, file[PC])
		}

		//2 byte instructions with 1 operand
		switch file[PC] {
		case 0x05:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ORA Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x06:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ASL Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x07:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SLO Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x09:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ORA Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x10:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BPL)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x11:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ORA Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x12:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ORA Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x14:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(TRB Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x15:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ORA Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x16:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ASL Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x17:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SLO Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x21:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(AND Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x24:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BIT Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x25:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(AND Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x26:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ROL Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x27:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(RLA Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x29:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(AND Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x30:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BMI)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x31:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(AND Zero Page,Indirect,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x32:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(AND Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x34:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BIT Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x35:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(AND Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x36:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ROL Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x37:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(RLA Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x41:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x44:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BSR)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x45:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x46:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LSR Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x47:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SRE Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x49:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x50:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BVC)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x51:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR Zero Page,Indirect,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x52:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x54:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BSR)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x55:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x56:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LSR Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x57:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SRE Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x61:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ADC Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x62:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(PER)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x64:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STZ Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x65:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ADC Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x66:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ROR Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x67:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(RRA Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x69:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ADC Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x70:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BVS)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x71:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ADC Zero Page,Indirect,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x72:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ADC Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x74:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STZ Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x75:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ADC Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x76:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ROR Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x77:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(RRA Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x80:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BRA)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x81:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x82:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BRL)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x84:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STY Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x85:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x86:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STX Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x87:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SAX Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x89:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BIT Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x90:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BCC)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x91:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA Zero Page,Indirect,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x92:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x94:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STY Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x95:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x96:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STX Zero Page,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x97:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SAX Zero Page,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDY Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDX Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA3:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LAX Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDY Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDX Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LAX Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA9:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xB0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BCS)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xB1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA Zero Page,Indirect,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xB2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xB4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDY Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xB5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xB6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDX Zero Page,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xB7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LAX Zero Page,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CPY Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(REP Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC3:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(DCP Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CPY Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(DEC Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(DCP Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC9:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xD0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BNE)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xD1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP Zero Page,Indirect,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xD2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xD4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(NOP Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xD5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xD6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(DEC Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xD7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(DCP Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CPX Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SBC Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SEP Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE3:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ISC Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CPX Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SBC Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(INC Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ISC Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE9:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SBC Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xF0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BEQ)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xF1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SBC Zero Page,Indirect,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xF2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SBC Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xF5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SBC Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xF6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(INC Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xF7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ISC Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		}

		//3 byte instructions with 2 operands
		switch file[PC] {
		case 0x0C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(NOP Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x0D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ORA Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x0E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ASL Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x0F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SLO Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x13:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SLO (Indirect),Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x19:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ORA Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x1C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(NOP Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x1D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ORA Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x1E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ASL Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x1F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SLO Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x20:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JSR)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x21:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND (Indirect,X))\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x22:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(KIL)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x23:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(RLA (Indirect,X))\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x2C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BIT Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x2D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x2E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROL Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x2F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(RLA Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x33:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(RLA (Indirect),Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x34:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(NOP Zero Page,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x35:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND Zero Page,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x36:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROL Zero Page,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x37:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(RLA Zero Page,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x39:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x3C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(NOP Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x3D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x3E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROL Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x3F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(RLA Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x4C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JMP Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x4D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(EOR Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x4E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LSR Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x4F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SRE Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x53:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SRE (Indirect),Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x59:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(EOR Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x5D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(EOR Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x5E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LSR Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x5F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SRE Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x63:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(RRA (Indirect,X))\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x6C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JMP (Indirect))\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x6D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ADC Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x6E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROR Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x6F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(RRA Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x73:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(RRA (Indirect),Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x79:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ADC Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x7C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(NOP Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x7D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ADC Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x7E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROR Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x7F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(RRA Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x83:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SAX (Indirect,X))\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x8B:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(XAA Immediate)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x8C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STY Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x8D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STA Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x8E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STX Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x8F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SAX Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x93:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AHX (Indirect),Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x99:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STA Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x9B:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(TAS Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x9C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SHY Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x9D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STA Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x9E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SHX Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x9F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AHX Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xAB:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LAX Immediate)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xAC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDY Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xAD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDA Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xAE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDX Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xAF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LAX Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xB3:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LAX (Indirect),Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xB9:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDA Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xBC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDY Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xBD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDA Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xBE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDX Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xBF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LAX Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xCB:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AXS Immediate)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xCC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CPY Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xCD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CMP Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xCE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(DEC Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xCF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(DCP Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xD3:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(DCP (Indirect),Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xD9:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CMP Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xDC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(NOP Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xDD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CMP Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xDE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(DEC Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xDF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(DCP Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xEB:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SBC Immediate)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xEC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CPX Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xED:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SBC Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xEE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(INC Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xEF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ISC Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xF3:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ISC (Indirect),Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xF4:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(NOP Zero Page,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xF9:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SBC Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xFC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(NOP Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xFD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SBC Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xFE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(INC Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xFF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ISC Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		}
	}
}
