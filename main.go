package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	A        byte        //Accumulator
	X        byte        //X register
	Y        byte        //Y register
	address  uint16      //Program Counter
	SR       byte        //Status Register
	SP       byte        //Stack Pointer
	memory   [65536]byte //Memory
	operand1 byte
	operand2 byte
)

func main() {
	//go run . ./c64kernal.bin 12
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <target_filename> <entry_point_address\n", os.Args[0])
		os.Exit(0)
	}

	//Read file
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		//command := os.Args[1]
		disassemble(string(file))
	} else {
		fmt.Println("Entry point parameter is missing!")
		fmt.Printf("USAGE : %s <target_filename> <entry_point_address>\n", os.Args[0])
	}
}

func disassemble(file string) {
	entrypoint, _ := strconv.ParseUint(os.Args[2], 16, 8)
	fmt.Printf("Entrypoint Hex: $%x Decimal: %v\n\n", entrypoint, entrypoint)
	entrypoint_decimal := fmt.Sprintf("%d", entrypoint)

	for address, _ := strconv.Atoi(entrypoint_decimal); address < len(file); address++ {
		//1 byte instructions with no operands
		switch file[address] {
		case 0x00:
			fmt.Printf("$%04x\t$%02x\t\t(BRK - Implied)\n", address, file[address])
		case 0x02:
			fmt.Printf("$%04x\t$%02x\t\t(CLE - Implied)\n", address, file[address])
		case 0x03:
			fmt.Printf("$%04x\t$%02x\t\t(SEE - Implied)\n", address, file[address])
		case 0x08:
			fmt.Printf("$%04x\t$%02x\t\t(PHP - Implied)\n", address, file[address])
		case 0x0A:
			fmt.Printf("$%04x\t$%02x\t\t(ASL - Accumulator)\n", address, file[address])
		case 0x0B:
			fmt.Printf("$%04x\t$%02x\t\t(TSY - Implied)\n", address, file[address])
		case 0x18:
			fmt.Printf("$%04x\t$%02x\t\t(CLC - Implied)\n", address, file[address])
		case 0x1A:
			fmt.Printf("$%04x\t$%02x\t\t(INC - Accumulator)\n", address, file[address])
		case 0x1B:
			fmt.Printf("$%04x\t$%02x\t\t(INZ - Implied)\n", address, file[address])
		case 0x28:
			fmt.Printf("$%04x\t$%02x\t\t(PLP - Implied)\n", address, file[address])
		case 0x2A:
			fmt.Printf("$%04x\t$%02x\t\t(ROL - Accumulator)\n", address, file[address])
		case 0x2B:
			fmt.Printf("$%04x\t$%02x\t\t(TYS - Implied)\n", address, file[address])
		case 0x38:
			fmt.Printf("$%04x\t$%02x\t\t(SEC - Accumulator)\n", address, file[address])
		case 0x3A:
			fmt.Printf("$%04x\t$%02x\t\t(DEC - Accumulator)\n", address, file[address])
		case 0x3B:
			fmt.Printf("$%04x\t$%02x\t\t(DEZ - Implied)\n", address, file[address])
		case 0x40:
			fmt.Printf("$%04x\t$%02x\t\t(RTI - Implied)\n", address, file[address])
		case 0x42:
			fmt.Printf("$%04x\t$%02x\t\t(NEG - Accumulator)\n", address, file[address])
		case 0x43:
			fmt.Printf("$%04x\t$%02x\t\t(ASR - Accumulator)\n", address, file[address])
		case 0x48:
			fmt.Printf("$%04x\t$%02x\t\t(PHA - Implied)\n", address, file[address])
		case 0x4A:
			fmt.Printf("$%04x\t$%02x\t\t(LSR - Accumulator)\n", address, file[address])
		case 0x4B:
			fmt.Printf("$%04x\t$%02x\t\t(TAZ - Implied)\n", address, file[address])
		case 0x58:
			fmt.Printf("$%04x\t$%02x\t\t(CLI - Implied)\n", address, file[address])
		case 0x5A:
			fmt.Printf("$%04x\t$%02x\t\t(PHY - Implied)\n", address, file[address])
		case 0x5B:
			fmt.Printf("$%04x\t$%02x\t\t(TAB - Implied)\n", address, file[address])
		case 0x60:
			fmt.Printf("$%04x\t$%02x\t\t(RTS - Implied)\n", address, file[address])
		case 0x68:
			fmt.Printf("$%04x\t$%02x\t\t(PLA - Implied)\n", address, file[address])
		case 0x6A:
			fmt.Printf("$%04x\t$%02x\t\t(ROR - Accumulator)\n", address, file[address])
		case 0x6B:
			fmt.Printf("$%04x\t$%02x\t\t(TZA - Implied)\n", address, file[address])
		case 0x78:
			fmt.Printf("$%04x\t$%02x\t\t(SEI - Implied)\n", address, file[address])
		case 0x7A:
			fmt.Printf("$%04x\t$%02x\t\t(PLY - Implied)\n", address, file[address])
		case 0x7B:
			fmt.Printf("$%04x\t$%02x\t\t(TBA - Absolute,Y)\n", address, file[address])
		case 0x88:
			fmt.Printf("$%04x\t$%02x\t\t(DEY - Implied)\n", address, file[address])
		case 0x8A:
			fmt.Printf("$%04x\t$%02x\t\t(TXA - Implied)\n", address, file[address])
		case 0x98:
			fmt.Printf("$%04x\t$%02x\t\t(TYA - Implied)\n", address, file[address])
		case 0x9A:
			fmt.Printf("$%04x\t$%02x\t\t(TXS - Implied)\n", address, file[address])
		case 0xA8:
			fmt.Printf("$%04x\t$%02x\t\t(TAY - Implied)\n", address, file[address])
		case 0xAA:
			fmt.Printf("$%04x\t$%02x\t\t(TAX - Implied)\n", address, file[address])
		case 0xB8:
			fmt.Printf("$%04x\t$%02x\t\t(CLV - Implied)\n", address, file[address])
		case 0xBA:
			fmt.Printf("$%04x\t$%02x\t\t(TSX - Implied)\n", address, file[address])
		case 0xC8:
			fmt.Printf("$%04x\t$%02x\t\t(INY - Implied)\n", address, file[address])
		case 0xCA:
			fmt.Printf("$%04x\t$%02x\t\t(DEX - Implied)\n", address, file[address])
		case 0xD8:
			fmt.Printf("$%04x\t$%02x\t\t(CLD - Implied)\n", address, file[address])
		case 0xDA:
			fmt.Printf("$%04x\t$%02x\t\t(PHX - Implied)\n", address, file[address])
		case 0xDB:
			fmt.Printf("$%04x\t$%02x\t\t(PHZ - Implied)\n", address, file[address])
		case 0xE8:
			fmt.Printf("$%04x\t$%02x\t\t(INX - Implied)\n", address, file[address])
		case 0xEA:
			fmt.Printf("$%04x\t$%02x\t\t(NOP - Implied)\n", address, file[address])
		case 0xF8:
			fmt.Printf("$%04x\t$%02x\t\t(SED - Implied)\n", address, file[address])
		case 0xFA:
			fmt.Printf("$%04x\t$%02x\t\t(PLX - Implied)\n", address, file[address])
		case 0xFB:
			fmt.Printf("$%04x\t$%02x\t\t(PLZ - Implied)\n", address, file[address])
		}

		//2 byte instructions with 1 operand
		switch file[address] {
		case 0x05:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ORA - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x06:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ASL - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x07:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(RMB0 - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x09:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ORA - Immediate)\n", address, file[address], file[address]+1)
			address++
		case 0x10:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BPL - Relative)\n", address, file[address], file[address]+1)
			address++
		case 0x11:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ORA - (Zero Page),Indirect Y)\n", address, file[address], file[address]+1)
			address++
		case 0x12:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ORA - (Zero Page),Indirect Z)\n", address, file[address], file[address]+1)
			address++
		case 0x14:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(TRB - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x15:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ORA - Zero Page,X)\n", address, file[address], file[address]+1)
			address++
		case 0x16:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ASL - Zero Page,X)\n", address, file[address], file[address]+1)
			address++
		case 0x17:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(RMB1 - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x21:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(AND - (Zero Page,Indirect))\n", address, file[address], file[address]+1)
			address++
		case 0x24:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BIT - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x25:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(AND - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x26:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ROL - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x27:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(RMB2 - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x29:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(AND - Immediate)\n", address, file[address], file[address]+1)
			address++
		case 0x30:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BMI - Relative)\n", address, file[address], file[address]+1)
			address++
		case 0x31:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(AND (Zero Page Indirect),Y)\n", address, file[address], file[address]+1)
			address++
		case 0x32:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(AND (Indirect),Z)\n", address, file[address], file[address]+1)
			address++
		case 0x34:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BIT - X Indexed Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x35:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(AND - X Indexed Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x36:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ROL - X Indexed Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x37:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(RMB3 - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x41:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR - (X Indexed Zero Page, Indirect))\n", address, file[address], file[address]+1)
			address++
		case 0x44:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ASR - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x45:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x46:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LSR - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x47:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(RMB4 - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x49:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR - Immediate)\n", address, file[address], file[address]+1)
			address++
		case 0x50:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BVC - Relative)\n", address, file[address], file[address]+1)
			address++
		case 0x51:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR (Zero Page Indirect),Y Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0x52:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR - (Indirect),Z Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0x54:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ASR - X Indexed Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x55:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(EOR - X Indexed Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x56:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LSR - X Indexed Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x57:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(RMB5 - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x61:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ADC - X Indexed Zero Page Indirect)\n", address, file[address], file[address]+1)
			address++
		case 0x62:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(RTN - Immediate)\n", address, file[address], file[address]+1)
			address++
		case 0x64:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STZ - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x65:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ADC - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x66:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ROR - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x67:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(RMB6 - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x69:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ADC - Immediate)\n", address, file[address], file[address]+1)
			address++
		case 0x70:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BVS - Relative)\n", address, file[address], file[address]+1)
			address++
		case 0x71:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ADC (Zero Page Indirect),Y Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0x72:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ADC - Indirect,Z Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0x74:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STZ - Zero Page,X Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0x75:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ADC - Zero Page,X Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0x76:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(ROR - Zero Page,X Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0x77:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(RMB7 - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x80:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BRA - Relative)\n", address, file[address], file[address]+1)
			address++
		case 0x81:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA - X Indexed Zero Page Indirect)\n", address, file[address], file[address]+1)
			address++
		case 0x82:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA - Stack Relative Indirect, Y Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0x84:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STY - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x85:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x86:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STX - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x87:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SMB0 - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0x89:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BIT - Immediate)\n", address, file[address], file[address]+1)
			address++
		case 0x90:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BCC - Relative)\n", address, file[address], file[address]+1)
			address++
		case 0x91:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA - (Zero Page Indirect),Y Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0x92:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA - Indirect,Z Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0x94:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STY - Zero Page,X Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0x95:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STA - Zero Page,X Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0x96:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(STX - Zero Page,Y Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0x97:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SMB1 - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xA0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDY - Immediate)\n", address, file[address], file[address]+1)
			address++
		case 0xA1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA - X Indexed Zero Page Indirect)\n", address, file[address], file[address]+1)
			address++
		case 0xA2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDX - Immediate)\n", address, file[address], file[address]+1)
			address++
		case 0xA3:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDZ - Immediate)\n", address, file[address], file[address]+1)
			address++
		case 0xA4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDY - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xA5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xA6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDX - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xA7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SMB2 - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xA9:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA - Immediate)\n", address, file[address], file[address]+1)
			address++
		case 0xB0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BCS - Relative)\n", address, file[address], file[address]+1)
			address++
		case 0xB1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA - (Zero Page Indirect),Y Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0xB2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA - Zero Page,Indirect)\n", address, file[address], file[address]+1)
			address++
		case 0xB4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDY - Zero Page,X Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0xB5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA - Zero Page,X Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0xB6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDX - Zero Page,Y Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0xB7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SMB3 - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xC0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CPY - Immediate)\n", address, file[address], file[address]+1)
			address++
		case 0xC1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP - X Indexed Zero Page Indirect)\n", address, file[address], file[address]+1)
			address++
		case 0xC2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CPZ - Immediate)\n", address, file[address], file[address]+1)
			address++
		case 0xC3:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(DEW - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xC4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CPY - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xC5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xC6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(DEC - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xC7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SMB4 - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xC9:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP - Immediate)\n", address, file[address], file[address]+1)
			address++
		case 0xD0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BNE - Relative)\n", address, file[address], file[address]+1)
			address++
		case 0xD1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP - (Zero Page Indirect),Y Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0xD2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP - (Indirect) Z Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0xD4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CPZ - Zero Page\n", address, file[address], file[address]+1)
			address++
		case 0xD5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CMP - Zero Page,X Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0xD6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(DEC - Zero Page,X Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0xD7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SMB5 - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xE0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CPX - Immediate)\n", address, file[address], file[address]+1)
			address++
		case 0xE1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SBC - X Indexed Zero Page Indirect)\n", address, file[address], file[address]+1)
			address++
		case 0xE2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(LDA - Immediate)\n", address, file[address], file[address]+1)
			address++
		case 0xE3:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(INW - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xE4:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(CPX - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xE5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SBC - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xE6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(INC - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xE7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SMB6 - Zero Page)\n", address, file[address], file[address]+1)
			address++
		case 0xE9:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SBC - Immediate)\n", address, file[address], file[address]+1)
			address++
		case 0xF0:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(BEQ - Relative)\n", address, file[address], file[address]+1)
			address++
		case 0xF1:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SBC - (Zero Page Indirect),Y Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0xF2:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SBC - (Indirect) Z Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0xF5:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SBC - Zero Page,X Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0xF6:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(INC - Zero Page,X Indexed)\n", address, file[address], file[address]+1)
			address++
		case 0xF7:
			fmt.Printf("$%04x\t$%02x $%02x\t\t(SMB7 - Zero Page)\n", address, file[address], file[address]+1)
			address++
		}

		//3 byte instructions with 2 operands
		switch file[address] {
		case 0x0C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(TSB - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x0D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ORA - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x0E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ASL - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x0F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR0 - Zero Page, Relative)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x13:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BPL - Relative (word))\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x19:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ORA - Absolute,Y Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x1C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(TRB - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x1D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ORA - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x1E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ASL - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x1F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR1 - Zero Page, Relative)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x20:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JSR - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x21:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND - (Zero Page Indirect,X Indexed))\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x22:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JSR - (Indirect) Z Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x23:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JSR - Absolute X Indexed Indirect)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x2C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BIT - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x2D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x2E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROL - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x2F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR2 - Zero Page, Relative)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x33:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BMI - Relative (word))\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x34:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BIT - Zero Page,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x35:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND - Zero Page,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x36:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROL - Zero Page,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x37:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(RMB3 - Zero Page)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x39:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND - Absolute,Y Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x3C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BIT - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x3D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(AND - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x3E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROL - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x3F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR3 - Zero Page, Relative)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x4C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JMP - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x4D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(EOR - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x4E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LSR - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x4F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR4 - Zero Page, Relative)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x53:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BVC - Relative (word))\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x59:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(EOR - Absolute,Y Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x5D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(EOR - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x5E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LSR - Absolute,X)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x5F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR5 - Zero Page, Relative)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x63:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BSR - Relative (word))\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x6C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JMP - Absolute Indirect)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x6D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ADC - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x6E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROR - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x6F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR6 - Zero Page, Relative)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x73:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BVS - Relative (word))\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x79:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ADC - Absolute,Y Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x7C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(JMP - Absolute,X Indexed Indirect)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x7D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ADC - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x7E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROR - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x7F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBR7 - Zero Page, Relative)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x83:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BRA - Relative (word))\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x8B:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STY - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x8C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STY - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x8D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STA - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x8E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STX - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x8F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS0- Zero Page, Relative)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x93:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BCC - Relative (word))\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x99:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STA - Absolute,Y Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x9B:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STX - Absolute,Y Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x9C:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STZ - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x9D:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STA - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x9E:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(STZ - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0x9F:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS1 - Zero Page, Relative)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xAB:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDZ - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xAC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDY - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xAD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDA - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xAE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDX - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xAF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS2 - Zero Page, Relative)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xB3:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BCS - Relative (word))\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xB9:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDA - Absolute,Y Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xBC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDY - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xBD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDA - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xBE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(LDX - Absolute,Y Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xBF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS3 - Zero Page, Relative)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xCB:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ASW - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xCC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CPY - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xCD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CMP - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xCE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(DEC - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xCF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS4 - Zero Page, Relative)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xD3:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BNE - Relative (word))\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xD9:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CMP - Absolute,Y Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xDC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CPZ - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xDD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CMP - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xDE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(DEC - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xDF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS5 - Zero Page, Relative)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xEB:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(ROW - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xEC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(CPX - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xED:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SBC - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xEE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(INC - Absolute)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xEF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS6 - Zero Page, Relative)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xF3:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BEQ - Relative (word))\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xF4:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(PHW - Immediate (word))\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xF9:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SBC - Absolute,Y Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xFC:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(PHW - Absolute\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xFD:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(SBC - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xFE:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(INC - Absolute,X Indexed)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		case 0xFF:
			fmt.Printf("$%04x\t$%02x $%02x $%02x\t(BBS7 - Zero Page, Relative)\n", address, file[address], file[address]+1, file[address]+2)
			address += 2
		}
	}
}
