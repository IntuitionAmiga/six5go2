package main

import (
	"fmt"
)

var (
	A      byte        //Accumulator
	X      byte        //X register
	Y      byte        //Y register
	PC     uint16      //Program Counter
	SR     byte        //Status Register
	SP     byte        //Stack Pointer
	memory [65536]byte //Memory
)

func main() {
	//Initialise first 256 bytes of memory with the 6502 opcode table
	for i := 0; i < 256; i++ {
		memory[i] = byte(i)
	}

	//Run the program counter from zero to end of table and decode the fetched opcode
	for PC := 0; PC < 256; PC++ {
		switch memory[PC] {
		case 0x00:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: BRK\n", PC, memory[PC])
		case 0x01:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ORA (Indirect,X)\n", PC, memory[PC])
		case 0x02:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: KIL\n", PC, memory[PC])
		case 0x03:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SLO (Indirect,X)\n", PC, memory[PC])
		case 0x04:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Zero Page\n", PC, memory[PC])
		case 0x05:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ORA Zero Page\n", PC, memory[PC])
		case 0x06:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ASL Zero Page\n", PC, memory[PC])
		case 0x07:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SLO Zero Page\n", PC, memory[PC])
		case 0x08:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: PHP\n", PC, memory[PC])
		case 0x09:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ORA Immediate\n", PC, memory[PC])
		case 0x0A:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ASL Accumulator\n", PC, memory[PC])
		case 0x0B:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ANC Immediate\n", PC, memory[PC])
		case 0x0C:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Absolute\n", PC, memory[PC])
		case 0x0D:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ORA Absolute\n", PC, memory[PC])
		case 0x0E:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ASL Absolute\n", PC, memory[PC])
		case 0x0F:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SLO Absolute\n", PC, memory[PC])
		case 0x10:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: BPL\n", PC, memory[PC])
		case 0x11:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ORA (Indirect),Y\n", PC, memory[PC])
		case 0x12:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: KIL\n", PC, memory[PC])
		case 0x13:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SLO (Indirect),Y\n", PC, memory[PC])
		case 0x14:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Zero Page,X\n", PC, memory[PC])
		case 0x15:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ORA Zero Page,X\n", PC, memory[PC])
		case 0x16:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ASL Zero Page,X\n", PC, memory[PC])
		case 0x17:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SLO Zero Page,X\n", PC, memory[PC])
		case 0x18:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CLC\n", PC, memory[PC])
		case 0x19:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ORA Absolute,Y\n", PC, memory[PC])
		case 0x1A:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP\n", PC, memory[PC])
		case 0x1B:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SLO Absolute,Y\n", PC, memory[PC])
		case 0x1C:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Absolute,X\n", PC, memory[PC])
		case 0x1D:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ORA Absolute,X\n", PC, memory[PC])
		case 0x1E:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ASL Absolute,X\n", PC, memory[PC])
		case 0x1F:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SLO Absolute,X\n", PC, memory[PC])
		case 0x20:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: JSR\n", PC, memory[PC])
		case 0x21:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: AND (Indirect,X)\n", PC, memory[PC])
		case 0x22:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: KIL\n", PC, memory[PC])
		case 0x23:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: RLA (Indirect,X)\n", PC, memory[PC])
		case 0x24:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: BIT Zero Page\n", PC, memory[PC])
		case 0x25:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: AND Zero Page\n", PC, memory[PC])
		case 0x26:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ROL Zero Page\n", PC, memory[PC])
		case 0x27:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: RLA Zero Page\n", PC, memory[PC])
		case 0x28:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: PLP\n", PC, memory[PC])
		case 0x29:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: AND Immediate\n", PC, memory[PC])
		case 0x2A:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ROL Accumulator\n", PC, memory[PC])
		case 0x2B:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ANC Immediate\n", PC, memory[PC])
		case 0x2C:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: BIT Absolute\n", PC, memory[PC])
		case 0x2D:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: AND Absolute\n", PC, memory[PC])
		case 0x2E:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ROL Absolute\n", PC, memory[PC])
		case 0x2F:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: RLA Absolute\n", PC, memory[PC])
		case 0x30:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: BMI\n", PC, memory[PC])
		case 0x31:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: AND (Indirect),Y\n", PC, memory[PC])
		case 0x32:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: KIL\n", PC, memory[PC])
		case 0x33:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: RLA (Indirect),Y\n", PC, memory[PC])
		case 0x34:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Zero Page,X\n", PC, memory[PC])
		case 0x35:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: AND Zero Page,X\n", PC, memory[PC])
		case 0x36:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ROL Zero Page,X\n", PC, memory[PC])
		case 0x37:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: RLA Zero Page,X\n", PC, memory[PC])
		case 0x38:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SEC\n", PC, memory[PC])
		case 0x39:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: AND Absolute,Y\n", PC, memory[PC])
		case 0x3A:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP\n", PC, memory[PC])
		case 0x3B:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: RLA Absolute,Y\n", PC, memory[PC])
		case 0x3C:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Absolute,X\n", PC, memory[PC])
		case 0x3D:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: AND Absolute,X\n", PC, memory[PC])
		case 0x3E:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ROL Absolute,X\n", PC, memory[PC])
		case 0x3F:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: RLA Absolute,X\n", PC, memory[PC])
		case 0x40:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: RTI\n", PC, memory[PC])
		case 0x41:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: EOR (Indirect,X)\n", PC, memory[PC])
		case 0x42:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: KIL\n", PC, memory[PC])
		case 0x43:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SRE (Indirect,X)\n", PC, memory[PC])
		case 0x44:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Zero Page\n", PC, memory[PC])
		case 0x45:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: EOR Zero Page\n", PC, memory[PC])
		case 0x46:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LSR Zero Page\n", PC, memory[PC])
		case 0x47:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SRE Zero Page\n", PC, memory[PC])
		case 0x48:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: PHA\n", PC, memory[PC])
		case 0x49:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: EOR Immediate\n", PC, memory[PC])
		case 0x4A:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LSR Accumulator\n", PC, memory[PC])
		case 0x4B:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ALR Immediate\n", PC, memory[PC])
		case 0x4C:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: JMP Absolute\n", PC, memory[PC])
		case 0x4D:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: EOR Absolute\n", PC, memory[PC])
		case 0x4E:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LSR Absolute\n", PC, memory[PC])
		case 0x4F:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SRE Absolute\n", PC, memory[PC])
		case 0x50:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: BVC\n", PC, memory[PC])
		case 0x51:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: EOR (Indirect),Y\n", PC, memory[PC])
		case 0x52:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: KIL\n", PC, memory[PC])
		case 0x53:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SRE (Indirect),Y\n", PC, memory[PC])
		case 0x54:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Zero Page,X\n", PC, memory[PC])
		case 0x55:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: EOR Zero Page,X\n", PC, memory[PC])
		case 0x56:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LSR Zero Page,X\n", PC, memory[PC])
		case 0x57:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SRE Zero Page,X\n", PC, memory[PC])
		case 0x58:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CLI\n", PC, memory[PC])
		case 0x59:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: EOR Absolute,Y\n", PC, memory[PC])
		case 0x5A:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP\n", PC, memory[PC])
		case 0x5B:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SRE Absolute,Y\n", PC, memory[PC])
		case 0x5C:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Absolute,X\n", PC, memory[PC])
		case 0x5D:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: EOR Absolute,X\n", PC, memory[PC])
		case 0x5E:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LSR Absolute,X\n", PC, memory[PC])
		case 0x5F:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SRE Absolute,X\n", PC, memory[PC])
		case 0x60:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: RTS\n", PC, memory[PC])
		case 0x61:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ADC (Indirect,X)\n", PC, memory[PC])
		case 0x62:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: KIL\n", PC, memory[PC])
		case 0x63:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: RRA (Indirect,X)\n", PC, memory[PC])
		case 0x64:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Zero Page\n", PC, memory[PC])
		case 0x65:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ADC Zero Page\n", PC, memory[PC])
		case 0x66:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ROR Zero Page\n", PC, memory[PC])
		case 0x67:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: RRA Zero Page\n", PC, memory[PC])
		case 0x68:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: PLA\n", PC, memory[PC])
		case 0x69:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ADC Immediate\n", PC, memory[PC])
		case 0x6A:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ROR Accumulator\n", PC, memory[PC])
		case 0x6B:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ARR Immediate\n", PC, memory[PC])
		case 0x6C:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: JMP (Indirect)\n", PC, memory[PC])
		case 0x6D:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ADC Absolute\n", PC, memory[PC])
		case 0x6E:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ROR Absolute\n", PC, memory[PC])
		case 0x6F:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: RRA Absolute\n", PC, memory[PC])
		case 0x70:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: BVS\n", PC, memory[PC])
		case 0x71:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ADC (Indirect),Y\n", PC, memory[PC])
		case 0x72:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: KIL\n", PC, memory[PC])
		case 0x73:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: RRA (Indirect),Y\n", PC, memory[PC])
		case 0x74:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Zero Page,X\n", PC, memory[PC])
		case 0x75:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ADC Zero Page,X\n", PC, memory[PC])
		case 0x76:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ROR Zero Page,X\n", PC, memory[PC])
		case 0x77:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: RRA Zero Page,X\n", PC, memory[PC])
		case 0x78:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SEI\n", PC, memory[PC])
		case 0x79:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ADC Absolute,Y\n", PC, memory[PC])
		case 0x7A:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP\n", PC, memory[PC])
		case 0x7B:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: RRA Absolute,Y\n", PC, memory[PC])
		case 0x7C:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Absolute,X\n", PC, memory[PC])
		case 0x7D:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ADC Absolute,X\n", PC, memory[PC])
		case 0x7E:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ROR Absolute,X\n", PC, memory[PC])
		case 0x7F:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: RRA Absolute,X\n", PC, memory[PC])
		case 0x80:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Immediate\n", PC, memory[PC])
		case 0x81:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: STA (Indirect,X)\n", PC, memory[PC])
		case 0x82:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP\n", PC, memory[PC])
		case 0x83:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SAX (Indirect,X)\n", PC, memory[PC])
		case 0x84:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: STY Zero Page\n", PC, memory[PC])
		case 0x85:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: STA Zero Page\n", PC, memory[PC])
		case 0x86:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: STX Zero Page\n", PC, memory[PC])
		case 0x87:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SAX Zero Page\n", PC, memory[PC])
		case 0x88:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: DEY\n", PC, memory[PC])
		case 0x89:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP\n", PC, memory[PC])
		case 0x8A:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: TXA\n", PC, memory[PC])
		case 0x8B:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: XAA Immediate\n", PC, memory[PC])
		case 0x8C:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: STY Absolute\n", PC, memory[PC])
		case 0x8D:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: STA Absolute\n", PC, memory[PC])
		case 0x8E:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: STX Absolute\n", PC, memory[PC])
		case 0x8F:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SAX Absolute\n", PC, memory[PC])
		case 0x90:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: BCC\n", PC, memory[PC])
		case 0x91:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: STA (Indirect),Y\n", PC, memory[PC])
		case 0x92:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: KIL\n", PC, memory[PC])
		case 0x93:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: AHX (Indirect),Y\n", PC, memory[PC])
		case 0x94:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: STY Zero Page,X\n", PC, memory[PC])
		case 0x95:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: STA Zero Page,X\n", PC, memory[PC])
		case 0x96:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: STX Zero Page,Y\n", PC, memory[PC])
		case 0x97:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SAX Zero Page,Y\n", PC, memory[PC])
		case 0x98:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: TYA\n", PC, memory[PC])
		case 0x99:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: STA Absolute,Y\n", PC, memory[PC])
		case 0x9A:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: TXS\n", PC, memory[PC])
		case 0x9B:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: TAS Absolute,Y\n", PC, memory[PC])
		case 0x9C:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SHY Absolute,X\n", PC, memory[PC])
		case 0x9D:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: STA Absolute,X\n", PC, memory[PC])
		case 0x9E:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SHX Absolute,Y\n", PC, memory[PC])
		case 0x9F:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: AHX Absolute,Y\n", PC, memory[PC])
		case 0xA0:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDY Immediate\n", PC, memory[PC])
		case 0xA1:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDA (Indirect,X)\n", PC, memory[PC])
		case 0xA2:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDX Immediate\n", PC, memory[PC])
		case 0xA3:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LAX (Indirect,X)\n", PC, memory[PC])
		case 0xA4:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDY Zero Page\n", PC, memory[PC])
		case 0xA5:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDA Zero Page\n", PC, memory[PC])
		case 0xA6:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDX Zero Page\n", PC, memory[PC])
		case 0xA7:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LAX Zero Page\n", PC, memory[PC])
		case 0xA8:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: TAY\n", PC, memory[PC])
		case 0xA9:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDA Immediate\n", PC, memory[PC])
		case 0xAA:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: TAX\n", PC, memory[PC])
		case 0xAB:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LAX Immediate\n", PC, memory[PC])
		case 0xAC:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDY Absolute\n", PC, memory[PC])
		case 0xAD:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDA Absolute\n", PC, memory[PC])
		case 0xAE:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDX Absolute\n", PC, memory[PC])
		case 0xAF:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LAX Absolute\n", PC, memory[PC])
		case 0xB0:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: BCS\n", PC, memory[PC])
		case 0xB1:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDA (Indirect),Y\n", PC, memory[PC])
		case 0xB2:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: KIL\n", PC, memory[PC])
		case 0xB3:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LAX (Indirect),Y\n", PC, memory[PC])
		case 0xB4:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDY Zero Page,X\n", PC, memory[PC])
		case 0xB5:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDA Zero Page,X\n", PC, memory[PC])
		case 0xB6:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDX Zero Page,Y\n", PC, memory[PC])
		case 0xB7:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LAX Zero Page,Y\n", PC, memory[PC])
		case 0xB8:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CLV\n", PC, memory[PC])
		case 0xB9:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDA Absolute,Y\n", PC, memory[PC])
		case 0xBA:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: TSX\n", PC, memory[PC])
		case 0xBB:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LAS Absolute,Y\n", PC, memory[PC])
		case 0xBC:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDY Absolute,X\n", PC, memory[PC])
		case 0xBD:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDA Absolute,X\n", PC, memory[PC])
		case 0xBE:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LDX Absolute,Y\n", PC, memory[PC])
		case 0xBF:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: LAX Absolute,Y\n", PC, memory[PC])
		case 0xC0:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CPY Immediate\n", PC, memory[PC])
		case 0xC1:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CMP (Indirect,X)\n", PC, memory[PC])
		case 0xC2:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Immediate\n", PC, memory[PC])
		case 0xC3:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: DCP (Indirect,X)\n", PC, memory[PC])
		case 0xC4:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CPY Zero Page\n", PC, memory[PC])
		case 0xC5:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CMP Zero Page\n", PC, memory[PC])
		case 0xC6:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: DEC Zero Page\n", PC, memory[PC])
		case 0xC7:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: DCP Zero Page\n", PC, memory[PC])
		case 0xC8:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: INY\n", PC, memory[PC])
		case 0xC9:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CMP Immediate\n", PC, memory[PC])
		case 0xCA:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: DEX\n", PC, memory[PC])
		case 0xCB:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: AXS Immediate\n", PC, memory[PC])
		case 0xCC:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CPY Absolute\n", PC, memory[PC])
		case 0xCD:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CMP Absolute\n", PC, memory[PC])
		case 0xCE:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: DEC Absolute\n", PC, memory[PC])
		case 0xCF:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: DCP Absolute\n", PC, memory[PC])
		case 0xD0:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: BNE\n", PC, memory[PC])
		case 0xD1:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CMP (Indirect),Y\n", PC, memory[PC])
		case 0xD2:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: KIL\n", PC, memory[PC])
		case 0xD3:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: DCP (Indirect),Y\n", PC, memory[PC])
		case 0xD4:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Zero Page,X\n", PC, memory[PC])
		case 0xD5:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CMP Zero Page,X\n", PC, memory[PC])
		case 0xD6:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: DEC Zero Page,X\n", PC, memory[PC])
		case 0xD7:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: DCP Zero Page,X\n", PC, memory[PC])
		case 0xD8:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CLD\n", PC, memory[PC])
		case 0xD9:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CMP Absolute,Y\n", PC, memory[PC])
		case 0xDA:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP\n", PC, memory[PC])
		case 0xDB:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: DCP Absolute,Y\n", PC, memory[PC])
		case 0xDC:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Absolute,X\n", PC, memory[PC])
		case 0xDD:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CMP Absolute,X\n", PC, memory[PC])
		case 0xDE:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: DEC Absolute,X\n", PC, memory[PC])
		case 0xDF:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: DCP Absolute,X\n", PC, memory[PC])
		case 0xE0:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CPX Immediate\n", PC, memory[PC])
		case 0xE1:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SBC (Indirect,X)\n", PC, memory[PC])
		case 0xE2:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Immediate\n", PC, memory[PC])
		case 0xE3:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ISC (Indirect,X)\n", PC, memory[PC])
		case 0xE4:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CPX Zero Page\n", PC, memory[PC])
		case 0xE5:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SBC Zero Page\n", PC, memory[PC])
		case 0xE6:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: INC Zero Page\n", PC, memory[PC])
		case 0xE7:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ISC Zero Page\n", PC, memory[PC])
		case 0xE8:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: INX\n", PC, memory[PC])
		case 0xE9:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SBC Immediate\n", PC, memory[PC])
		case 0xEA:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP\n", PC, memory[PC])
		case 0xEB:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SBC Immediate\n", PC, memory[PC])
		case 0xEC:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: CPX Absolute\n", PC, memory[PC])
		case 0xED:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SBC Absolute\n", PC, memory[PC])
		case 0xEE:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: INC Absolute\n", PC, memory[PC])
		case 0xEF:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ISC Absolute\n", PC, memory[PC])
		case 0xF0:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: BEQ\n", PC, memory[PC])
		case 0xF1:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SBC (Indirect),Y\n", PC, memory[PC])
		case 0xF2:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: KIL\n", PC, memory[PC])
		case 0xF3:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ISC (Indirect),Y\n", PC, memory[PC])
		case 0xF4:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Zero Page,X\n", PC, memory[PC])
		case 0xF5:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SBC Zero Page,X\n", PC, memory[PC])
		case 0xF6:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: INC Zero Page,X\n", PC, memory[PC])
		case 0xF7:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ISC Zero Page,X\n", PC, memory[PC])
		case 0xF8:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SED\n", PC, memory[PC])
		case 0xF9:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SBC Absolute,Y\n", PC, memory[PC])
		case 0xFA:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP\n", PC, memory[PC])
		case 0xFB:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ISC Absolute,Y\n", PC, memory[PC])
		case 0xFC:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: NOP Absolute,X\n", PC, memory[PC])
		case 0xFD:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: SBC Absolute,X\n", PC, memory[PC])
		case 0xFE:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: INC Absolute,X\n", PC, memory[PC])
		case 0xFF:
			fmt.Printf("Memory Address: $%04x Opcode: $%04x Instruction: ISC Absolute,X\n", PC, memory[PC])

		default:
			panic(fmt.Sprintf("Unhandled instruction: $%04x", memory[PC]))
		}
	}
}
