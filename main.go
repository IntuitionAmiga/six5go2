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

func initRAM() {
	//Initialise first 256 bytes of memory with the 6502 opcode table
	for i := 0; i < 256; i++ {
		memory[i] = byte(i)
	}
}

func disassemble(file string) {
	//Run the program counter from zero to end of table and decode the fetched opcode
	for PC := 0; PC < len(file); PC++ {
		//1 byte instructions with no operands
		switch file[PC] {
		case 0x00:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(BRK)\n", PC, file[PC])
		case 0x02:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(KIL)\n", PC, file[PC])
		case 0x03:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(SLO (Indirect,X))\n", PC, file[PC])
		case 0x08:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(PHP)\n", PC, file[PC])
		case 0x0A:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(ASL Accumulator)\n", PC, file[PC])
		case 0x0B:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(ANC Immediate)\n", PC, file[PC])
		case 0x18:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(CLC)\n", PC, file[PC])
		case 0x1A:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(NOP)\n", PC, file[PC])
		case 0x1B:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(SLO Absolute,Y)\n", PC, file[PC])
		case 0x28:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(PLP)\n", PC, file[PC])
		case 0x2A:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(ROL Accumulator)\n", PC, file[PC])
		case 0x2B:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(ANC Immediate)\n", PC, file[PC])
		case 0x38:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(SEC)\n", PC, file[PC])
		case 0x3A:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(NOP)\n", PC, file[PC])
		case 0x3B:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(RLA Absolute,Y)\n", PC, file[PC])
		case 0x40:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(RTI)\n", PC, file[PC])
		case 0x42:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(KIL)\n", PC, file[PC])
		case 0x43:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(SRE (Indirect,X))\n", PC, file[PC])
		case 0x48:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(PHA)\n", PC, file[PC])
		case 0x4A:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(LSR Accumulator)\n", PC, file[PC])
		case 0x4B:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(ALR Immediate)\n", PC, file[PC])
		case 0x58:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(CLI)\n", PC, file[PC])
		case 0x5A:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(NOP)\n", PC, file[PC])
		case 0x5B:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(SRE Absolute,Y)\n", PC, file[PC])
		case 0x60:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(RTS)\n", PC, file[PC])
		case 0x68:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(PLA)\n", PC, file[PC])
		case 0x6A:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(ROR Accumulator)\n", PC, file[PC])
		case 0x6B:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(ARR Immediate)\n", PC, file[PC])
		case 0x78:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(SEI)\n", PC, file[PC])
		case 0x7A:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(NOP)\n", PC, file[PC])
		case 0x7B:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(RRA Absolute,Y)\n", PC, file[PC])
		case 0x88:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(DEY)\n", PC, file[PC])
		case 0x8A:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(TXA)\n", PC, file[PC])
		case 0x98:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(TYA)\n", PC, file[PC])
		case 0x9A:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(TXS)\n", PC, file[PC])
		case 0xA8:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(TAY)\n", PC, file[PC])
		case 0xAA:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(TAX)\n", PC, file[PC])
		case 0xB8:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(CLV)\n", PC, file[PC])
		case 0xBA:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(TSX)\n", PC, file[PC])
		case 0xC8:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(INY)\n", PC, file[PC])
		case 0xCA:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(DEX)\n", PC, file[PC])
		case 0xD8:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(CLD)\n", PC, file[PC])
		case 0xDA:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(PHX)\n", PC, file[PC])
		case 0xDB:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(PHZ)\n", PC, file[PC])
		case 0xE8:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(INX)\n", PC, file[PC])
		case 0xEA:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(NOP)\n", PC, file[PC])
		case 0xF8:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(SED)\n", PC, file[PC])
		case 0xFA:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(PLX)\n", PC, file[PC])
		case 0xFB:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x\t\t\t\t(PLZ)\n", PC, file[PC])
		}

		//2 byte instructions with 1 operand
		switch file[PC] {
		case 0x05:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ORA Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x06:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ASL Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x07:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(SLO Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x09:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ORA Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x10:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(BPL)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x11:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ORA Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x12:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ORA Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x14:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(TRB Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x15:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ORA Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x16:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ASL Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x17:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(SLO Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x21:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(AND Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x24:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(BIT Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x25:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(AND Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x26:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ROL Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x27:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(RLA Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x29:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(AND Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x30:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(BMI)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x31:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(AND Zero Page,Indirect,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x32:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(AND Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x34:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(BIT Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x35:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(AND Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x36:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ROL Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x37:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(RLA Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x41:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(EOR Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x44:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(BSR)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x45:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(EOR Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x46:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LSR Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x47:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(SRE Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x49:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(EOR Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x50:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(BVC)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x51:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(EOR Zero Page,Indirect,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x52:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(EOR Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x54:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(BSR)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x55:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(EOR Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x56:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LSR Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x57:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(SRE Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x61:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ADC Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x62:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(PER)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x64:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(STZ Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x65:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ADC Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x66:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ROR Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x67:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(RRA Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x69:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ADC Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x70:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(BVS)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x71:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ADC Zero Page,Indirect,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x72:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ADC Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x74:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(STZ Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x75:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ADC Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x76:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ROR Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x77:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(RRA Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x80:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(BRA)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x81:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(STA Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x82:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(BRL)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x84:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(STY Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x85:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(STA Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x86:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(STX Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x87:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(SAX Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x89:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(BIT Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x90:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(BCC)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x91:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(STA Zero Page,Indirect,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x92:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(STA Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x94:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(STY Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x95:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(STA Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x96:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(STX Zero Page,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0x97:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(SAX Zero Page,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA0:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LDY Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA1:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LDA Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA2:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LDX Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA3:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LAX Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA4:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LDY Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA5:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LDA Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA6:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LDX Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA7:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LAX Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xA9:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LDA Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xB0:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(BCS)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xB1:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LDA Zero Page,Indirect,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xB2:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LDA Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xB4:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LDY Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xB5:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LDA Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xB6:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LDX Zero Page,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xB7:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(LAX Zero Page,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC0:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(CPY Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC1:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(CMP Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC2:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(REP Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC3:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(DCP Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC4:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(CPY Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC5:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(CMP Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC6:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(DEC Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC7:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(DCP Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xC9:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(CMP Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xD0:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(BNE)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xD1:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(CMP Zero Page,Indirect,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xD2:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(CMP Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xD4:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(NOP Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xD5:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(CMP Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xD6:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(DEC Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xD7:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(DCP Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE0:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(CPX Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE1:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(SBC Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE2:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(SEP Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE3:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ISC Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE4:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(CPX Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE5:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(SBC Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE6:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(INC Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE7:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ISC Zero Page)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xE9:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(SBC Immediate)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xF0:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(BEQ)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xF1:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(SBC Zero Page,Indirect,Y)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xF2:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(SBC Zero Page,Indirect)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xF5:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(SBC Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xF6:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(INC Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		case 0xF7:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x\t\t(ISC Zero Page,X)\n", PC, file[PC], file[PC]+1)
			PC++
		}

		//3 byte instructions with 2 operands
		switch file[PC] {
		case 0x0C:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (NOP Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x0D:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (ORA Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x0E:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (ASL Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x0F:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (SLO Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x13:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (SLO (Indirect),Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x19:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (ORA Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x1C:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (NOP Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x1D:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (ORA Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x1E:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (ASL Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x1F:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (SLO Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x20:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (JSR)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x21:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (AND (Indirect,X))\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x22:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (KIL)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x23:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (RLA (Indirect,X))\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x2C:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (BIT Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x2D:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (AND Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x2E:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (ROL Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x2F:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (RLA Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x33:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (RLA (Indirect),Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x34:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (NOP Zero Page,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x35:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (AND Zero Page,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x36:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (ROL Zero Page,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x37:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (RLA Zero Page,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x39:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (AND Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x3C:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (NOP Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x3D:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (AND Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x3E:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (ROL Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x3F:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (RLA Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x4C:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (JMP Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x4D:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (EOR Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x4E:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (LSR Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x4F:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (SRE Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x53:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (SRE (Indirect),Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x59:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (EOR Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x5D:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (EOR Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x5E:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (LSR Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x5F:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (SRE Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x63:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (RRA (Indirect,X))\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x6C:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (JMP (Indirect))\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x6D:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (ADC Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x6E:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (ROR Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x6F:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (RRA Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x73:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (RRA (Indirect),Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x79:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (ADC Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x7C:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (NOP Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x7D:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (ADC Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x7E:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (ROR Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x7F:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (RRA Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x83:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (SAX (Indirect,X))\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x8B:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (XAA Immediate)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x8C:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (STY Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x8D:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (STA Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x8E:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (STX Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x8F:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (SAX Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x93:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (AHX (Indirect),Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x99:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (STA Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x9B:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (TAS Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x9C:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (SHY Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x9D:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (STA Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x9E:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (SHX Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0x9F:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (AHX Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xAB:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (LAX Immediate)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xAC:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (LDY Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xAD:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (LDA Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xAE:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (LDX Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xAF:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (LAX Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xB3:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (LAX (Indirect),Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xB9:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (LDA Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xBC:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (LDY Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xBD:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (LDA Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xBE:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (LDX Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xBF:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (LAX Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xCB:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (AXS Immediate)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xCC:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (CPY Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xCD:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (CMP Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xCE:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (DEC Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xCF:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (DCP Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xD3:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (DCP (Indirect),Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xD9:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (CMP Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xDC:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (NOP Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xDD:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (CMP Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xDE:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (DEC Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xDF:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (DCP Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xEB:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (SBC Immediate)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xEC:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (CPX Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xED:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (SBC Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xEE:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (INC Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xEF:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (ISC Absolute)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xF3:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (ISC (Indirect),Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xF4:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (NOP Zero Page,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xF9:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (SBC Absolute,Y)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xFC:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (NOP Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xFD:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (SBC Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xFE:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (INC Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		case 0xFF:
			fmt.Printf("Memory Address: $%04x\tOpcode: $%02x Operand1: $%02x Operand2: $%02x (ISC Absolute,X)\n", PC, file[PC], file[PC]+1, file[PC]+2)
			PC += 2
		}
	}
}
