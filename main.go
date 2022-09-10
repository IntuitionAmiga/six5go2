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
	for i := 0; i < 256; i++ {
		memory[i] = byte(i)
	}

	//Run the program counter from zero to end of memory and decode the fetched opcode
	for PC := 0; PC < 256; PC++ {
		switch memory[PC] {
		case 0x00:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x00")
			fmt.Println(" Instruction: BRK")

		case 0x01:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x01")
			fmt.Println(" Instruction: ORA (Indirect,X)")

		case 0x02:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x02")
			fmt.Println(" Instruction: KIL")

		case 0x03:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x03")
			fmt.Println(" Instruction: SLO (Indirect,X)")

		case 0x04:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x04")
			fmt.Println(" Instruction: NOP Zero Page")

		case 0x05:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x05")
			fmt.Println(" Instruction: ORA Zero Page")

		case 0x06:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x06")
			fmt.Println(" Instruction: ASL Zero Page")

		case 0x07:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x07")
			fmt.Println(" Instruction: SLO Zero Page")

		case 0x08:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x08")
			fmt.Println(" Instruction: PHP")

		case 0x09:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x09")
			fmt.Println(" Instruction: ORA Immediate")

		case 0x0A:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x0A")
			fmt.Println(" Instruction: ASL Accumulator")

		case 0x0B:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x0B")
			fmt.Println(" Instruction: ANC Immediate")

		case 0x0C:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x0C")
			fmt.Println(" Instruction: NOP Absolute")

		case 0x0D:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x0D")
			fmt.Println(" Instruction: ORA Absolute")

		case 0x0E:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x0E")
			fmt.Println(" Instruction: ASL Absolute")

		case 0x0F:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x0F")
			fmt.Println(" Instruction: SLO Absolute")

		case 0x10:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x10")
			fmt.Println(" Instruction: BPL")

		case 0x11:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x11")
			fmt.Println(" Instruction: ORA (Indirect),Y")

		case 0x12:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x12")
			fmt.Println(" Instruction: KIL")

		case 0x13:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x13")
			fmt.Println(" Instruction: SLO (Indirect),Y")

		case 0x14:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x14")
			fmt.Println(" Instruction: NOP Zero Page,X")

		case 0x15:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x15")
			fmt.Println(" Instruction: ORA Zero Page,X")

		case 0x16:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x16")
			fmt.Println(" Instruction: ASL Zero Page,X")

		case 0x17:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x17")
			fmt.Println(" Instruction: SLO Zero Page,X")

		case 0x18:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x18")
			fmt.Println(" Instruction: CLC")

		case 0x19:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x19")
			fmt.Println(" Instruction: ORA Absolute,Y")

		case 0x1A:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x1A")
			fmt.Println(" Instruction: NOP")

		case 0x1B:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x1B")
			fmt.Println(" Instruction: SLO Absolute,Y")

		case 0x1C:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x1C")
			fmt.Println(" Instruction: NOP Absolute,X")

		case 0x1D:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x1D")
			fmt.Println(" Instruction: ORA Absolute,X")

		case 0x1E:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x1E")
			fmt.Println(" Instruction: ASL Absolute,X")

		case 0x1F:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x1F")
			fmt.Println(" Instruction: SLO Absolute,X")

		case 0x20:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x20")
			fmt.Println(" Instruction: JSR")

		case 0x21:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x21")
			fmt.Println(" Instruction: AND (Indirect,X)")

		case 0x22:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x22")
			fmt.Println(" Instruction: KIL")

		case 0x23:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x23")
			fmt.Println(" Instruction: RLA (Indirect,X)")

		case 0x24:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x24")
			fmt.Println(" Instruction: BIT Zero Page")

		case 0x25:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x25")
			fmt.Println(" Instruction: AND Zero Page")

		case 0x26:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x26")
			fmt.Println(" Instruction: ROL Zero Page")

		case 0x27:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x27")
			fmt.Println(" Instruction: RLA Zero Page")

		case 0x28:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x28")
			fmt.Println(" Instruction: PLP")

		case 0x29:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x29")
			fmt.Println(" Instruction: AND Immediate")

		case 0x2A:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x2A")
			fmt.Println(" Instruction: ROL Accumulator")

		case 0x2B:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x2B")
			fmt.Println(" Instruction: ANC Immediate")

		case 0x2C:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x2C")
			fmt.Println(" Instruction: BIT Absolute")

		case 0x2D:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x2D")
			fmt.Println(" Instruction: AND Absolute")

		case 0x2E:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x2E")
			fmt.Println(" Instruction: ROL Absolute")

		case 0x2F:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x2F")
			fmt.Println(" Instruction: RLA Absolute")

		case 0x30:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x30")
			fmt.Println(" Instruction: BMI")

		case 0x31:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x31")
			fmt.Println(" Instruction: AND (Indirect),Y")

		case 0x32:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x32")
			fmt.Println(" Instruction: KIL")

		case 0x33:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x33")
			fmt.Println(" Instruction: RLA (Indirect),Y")

		case 0x34:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x34")
			fmt.Println(" Instruction: NOP Zero Page,X")

		case 0x35:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x35")
			fmt.Println(" Instruction: AND Zero Page,X")

		case 0x36:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x36")
			fmt.Println(" Instruction: ROL Zero Page,X")

		case 0x37:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x37")
			fmt.Println(" Instruction: RLA Zero Page,X")

		case 0x38:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x38")
			fmt.Println(" Instruction: SEC")

		case 0x39:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x39")
			fmt.Println(" Instruction: AND Absolute,Y")

		case 0x3A:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x3A")
			fmt.Println(" Instruction: NOP")

		case 0x3B:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x3B")
			fmt.Println(" Instruction: RLA Absolute,Y")

		case 0x3C:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x3C")
			fmt.Println(" Instruction: NOP Absolute,X")

		case 0x3D:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x3D")
			fmt.Println(" Instruction: AND Absolute,X")

		case 0x3E:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x3E")
			fmt.Println(" Instruction: ROL Absolute,X")

		case 0x3F:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x3F")
			fmt.Println(" Instruction: RLA Absolute,X")

		case 0x40:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x40")
			fmt.Println(" Instruction: RTI")

		case 0x41:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x41")
			fmt.Println(" Instruction: EOR (Indirect,X)")

		case 0x42:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x42")
			fmt.Println(" Instruction: KIL")

		case 0x43:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x43")
			fmt.Println(" Instruction: SRE (Indirect,X)")

		case 0x44:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x44")
			fmt.Println(" Instruction: NOP Zero Page")

		case 0x45:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x45")
			fmt.Println(" Instruction: EOR Zero Page")

		case 0x46:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x46")
			fmt.Println(" Instruction: LSR Zero Page")

		case 0x47:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x47")
			fmt.Println(" Instruction: SRE Zero Page")

		case 0x48:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x48")
			fmt.Println(" Instruction: PHA")

		case 0x49:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x49")
			fmt.Println(" Instruction: EOR Immediate")

		case 0x4A:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x4A")
			fmt.Println(" Instruction: LSR Accumulator")

		case 0x4B:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x4B")
			fmt.Println(" Instruction: ALR Immediate")

		case 0x4C:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x4C")
			fmt.Println(" Instruction: JMP Absolute")

		case 0x4D:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x4D")
			fmt.Println(" Instruction: EOR Absolute")

		case 0x4E:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x4E")
			fmt.Println(" Instruction: LSR Absolute")

		case 0x4F:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x4F")
			fmt.Println(" Instruction: SRE Absolute")

		case 0x50:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x50")
			fmt.Println(" Instruction: BVC")

		case 0x51:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x51")
			fmt.Println(" Instruction: EOR (Indirect),Y")

		case 0x52:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x52")
			fmt.Println(" Instruction: KIL")

		case 0x53:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x53")
			fmt.Println(" Instruction: SRE (Indirect),Y")

		case 0x54:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x54")
			fmt.Println(" Instruction: NOP Zero Page,X")

		case 0x55:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x55")
			fmt.Println(" Instruction: EOR Zero Page,X")

		case 0x56:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x56")
			fmt.Println(" Instruction: LSR Zero Page,X")

		case 0x57:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x57")
			fmt.Println(" Instruction: SRE Zero Page,X")

		case 0x58:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x58")
			fmt.Println(" Instruction: CLI")

		case 0x59:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x59")
			fmt.Println(" Instruction: EOR Absolute,Y")

		case 0x5A:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x5A")
			fmt.Println(" Instruction: NOP")

		case 0x5B:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x5B")
			fmt.Println(" Instruction: SRE Absolute,Y")

		case 0x5C:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x5C")
			fmt.Println(" Instruction: NOP Absolute,X")

		case 0x5D:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x5D")
			fmt.Println(" Instruction: EOR Absolute,X")

		case 0x5E:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x5E")
			fmt.Println(" Instruction: LSR Absolute,X")

		case 0x5F:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x5F")
			fmt.Println(" Instruction: SRE Absolute,X")

		case 0x60:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x60")
			fmt.Println(" Instruction: RTS")

		case 0x61:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x61")
			fmt.Println(" Instruction: ADC (Indirect,X)")

		case 0x62:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x62")
			fmt.Println(" Instruction: KIL")

		case 0x63:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x63")
			fmt.Println(" Instruction: RRA (Indirect,X)")

		case 0x64:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x64")
			fmt.Println(" Instruction: NOP Zero Page")

		case 0x65:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x65")
			fmt.Println(" Instruction: ADC Zero Page")

		case 0x66:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x66")
			fmt.Println(" Instruction: ROR Zero Page")

		case 0x67:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x67")
			fmt.Println(" Instruction: RRA Zero Page")

		case 0x68:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x68")
			fmt.Println(" Instruction: PLA")

		case 0x69:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x69")
			fmt.Println(" Instruction: ADC Immediate")

		case 0x6A:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x6A")
			fmt.Println(" Instruction: ROR Accumulator")

		case 0x6B:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x6B")
			fmt.Println(" Instruction: ARR Immediate")

		case 0x6C:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x6C")
			fmt.Println(" Instruction: JMP (Indirect)")

		case 0x6D:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x6D")
			fmt.Println(" Instruction: ADC Absolute")

		case 0x6E:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x6E")
			fmt.Println(" Instruction: ROR Absolute")

		case 0x6F:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x6F")
			fmt.Println(" Instruction: RRA Absolute")

		case 0x70:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x70")
			fmt.Println(" Instruction: BVS")

		case 0x71:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x71")
			fmt.Println(" Instruction: ADC (Indirect),Y")

		case 0x72:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x72")
			fmt.Println(" Instruction: KIL")

		case 0x73:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x73")
			fmt.Println(" Instruction: RRA (Indirect),Y")

		case 0x74:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x74")
			fmt.Println(" Instruction: NOP Zero Page,X")

		case 0x75:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x75")
			fmt.Println(" Instruction: ADC Zero Page,X")

		case 0x76:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x76")
			fmt.Println(" Instruction: ROR Zero Page,X")

		case 0x77:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x77")
			fmt.Println(" Instruction: RRA Zero Page,X")

		case 0x78:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x78")
			fmt.Println(" Instruction: SEI")

		case 0x79:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x79")
			fmt.Println(" Instruction: ADC Absolute,Y")

		case 0x7A:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x7A")
			fmt.Println(" Instruction: NOP")

		case 0x7B:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x7B")
			fmt.Println(" Instruction: RRA Absolute,Y")

		case 0x7C:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x7C")
			fmt.Println(" Instruction: NOP Absolute,X")

		case 0x7D:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x7D")
			fmt.Println(" Instruction: ADC Absolute,X")

		case 0x7E:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x7E")
			fmt.Println(" Instruction: ROR Absolute,X")

		case 0x7F:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x7F")
			fmt.Println(" Instruction: RRA Absolute,X")

		case 0x80:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x80")
			fmt.Println(" Instruction: NOP Immediate")

		case 0x81:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x81")
			fmt.Println(" Instruction: STA (Indirect,X)")

		case 0x82:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x82")
			fmt.Println(" Instruction: NOP Immediate")

		case 0x83:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x83")
			fmt.Println(" Instruction: SAX (Indirect,X)")

		case 0x84:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x84")
			fmt.Println(" Instruction: STY Zero Page")

		case 0x85:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x85")
			fmt.Println(" Instruction: STA Zero Page")

		case 0x86:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x86")
			fmt.Println(" Instruction: STX Zero Page")

		case 0x87:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x87")
			fmt.Println(" Instruction: SAX Zero Page")

		case 0x88:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x88")
			fmt.Println(" Instruction: DEY")

		case 0x89:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x89")
			fmt.Println(" Instruction: NOP Immediate")

		case 0x8A:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x8A")
			fmt.Println(" Instruction: TXA")

		case 0x8B:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x8B")
			fmt.Println(" Instruction: XAA Immediate")

		case 0x8C:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x8C")
			fmt.Println(" Instruction: STY Absolute")

		case 0x8D:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x8D")
			fmt.Println(" Instruction: STA Absolute")

		case 0x8E:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x8E")
			fmt.Println(" Instruction: STX Absolute")

		case 0x8F:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x8F")
			fmt.Println(" Instruction: SAX Absolute")

		case 0x90:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x90")
			fmt.Println(" Instruction: BCC")

		case 0x91:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x91")
			fmt.Println(" Instruction: STA (Indirect),Y")

		case 0x92:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x92")
			fmt.Println(" Instruction: KIL")

		case 0x93:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x93")
			fmt.Println(" Instruction: AHX (Indirect),Y")

		case 0x94:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x94")
			fmt.Println(" Instruction: STY Zero Page,X")

		case 0x95:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x95")
			fmt.Println(" Instruction: STA Zero Page,X")

		case 0x96:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x96")
			fmt.Println(" Instruction: STX Zero Page,Y")

		case 0x97:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x97")
			fmt.Println(" Instruction: SAX Zero Page,Y")

		case 0x98:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x98")
			fmt.Println(" Instruction: TYA")

		case 0x99:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x99")
			fmt.Println(" Instruction: STA Absolute,Y")

		case 0x9A:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x9A")
			fmt.Println(" Instruction: TXS")

		case 0x9B:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x9B")
			fmt.Println(" Instruction: TAS Absolute,Y")

		case 0x9C:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x9C")
			fmt.Println(" Instruction: SHY Absolute,X")

		case 0x9D:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x9D")
			fmt.Println(" Instruction: STA Absolute,X")

		case 0x9E:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x9E")
			fmt.Println(" Instruction: SHX Absolute,Y")

		case 0x9F:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0x9F")
			fmt.Println(" Instruction: AHX Absolute,Y")

		case 0xA0:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xA0")
			fmt.Println(" Instruction: LDY Immediate")

		case 0xA1:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xA1")
			fmt.Println(" Instruction: LDA (Indirect,X)")

		case 0xA2:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xA2")
			fmt.Println(" Instruction: LDX Immediate")

		case 0xA3:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xA3")
			fmt.Println(" Instruction: LAX (Indirect,X)")

		case 0xA4:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xA4")
			fmt.Println(" Instruction: LDY Zero Page")

		case 0xA5:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xA5")
			fmt.Println(" Instruction: LDA Zero Page")

		case 0xA6:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xA6")
			fmt.Println(" Instruction: LDX Zero Page")

		case 0xA7:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xA7")
			fmt.Println(" Instruction: LAX Zero Page")

		case 0xA8:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xA8")
			fmt.Println(" Instruction: TAY")

		case 0xA9:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xA9")
			fmt.Println(" Instruction: LDA Immediate")

		case 0xAA:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xAA")
			fmt.Println(" Instruction: TAX")

		case 0xAB:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xAB")
			fmt.Println(" Instruction: LAX Immediate")

		case 0xAC:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xAC")
			fmt.Println(" Instruction: LDY Absolute")

		case 0xAD:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xAD")
			fmt.Println(" Instruction: LDA Absolute")

		case 0xAE:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xAE")
			fmt.Println(" Instruction: LDX Absolute")

		case 0xAF:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xAF")
			fmt.Println(" Instruction: LAX Absolute")

		case 0xB0:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xB0")
			fmt.Println(" Instruction: BCS")

		case 0xB1:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xB1")
			fmt.Println(" Instruction: LDA (Indirect),Y")

		case 0xB2:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xB2")
			fmt.Println(" Instruction: KIL")

		case 0xB3:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xB3")
			fmt.Println(" Instruction: LAX (Indirect),Y")

		case 0xB4:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xB4")
			fmt.Println(" Instruction: LDY Zero Page,X")

		case 0xB5:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xB5")
			fmt.Println(" Instruction: LDA Zero Page,X")

		case 0xB6:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xB6")
			fmt.Println(" Instruction: LDX Zero Page,Y")

		case 0xB7:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xB7")
			fmt.Println(" Instruction: LAX Zero Page,Y")

		case 0xB8:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xB8")
			fmt.Println(" Instruction: CLV")

		case 0xB9:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xB9")
			fmt.Println(" Instruction: LDA Absolute,Y")

		case 0xBA:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xBA")
			fmt.Println(" Instruction: TSX")

		case 0xBB:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xBB")
			fmt.Println(" Instruction: LAS Absolute,Y")

		case 0xBC:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xBC")
			fmt.Println(" Instruction: LDY Absolute,X")

		case 0xBD:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xBD")
			fmt.Println(" Instruction: LDA Absolute,X")

		case 0xBE:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xBE")
			fmt.Println(" Instruction: LDX Absolute,Y")

		case 0xBF:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xBF")
			fmt.Println(" Instruction: LAX Absolute,Y")

		case 0xC0:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xC0")
			fmt.Println(" Instruction: CPY Immediate")

		case 0xC1:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xC1")
			fmt.Println(" Instruction: CMP (Indirect,X)")

		case 0xC2:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xC2")
			fmt.Println(" Instruction: NOP Immediate")

		case 0xC3:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xC3")
			fmt.Println(" Instruction: DCP (Indirect,X)")

		case 0xC4:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xC4")
			fmt.Println(" Instruction: CPY Zero Page")

		case 0xC5:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xC5")
			fmt.Println(" Instruction: CMP Zero Page")

		case 0xC6:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xC6")
			fmt.Println(" Instruction: DEC Zero Page")

		case 0xC7:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xC7")
			fmt.Println(" Instruction: DCP Zero Page")

		case 0xC8:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xC8")
			fmt.Println(" Instruction: INY")

		case 0xC9:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xC9")
			fmt.Println(" Instruction: CMP Immediate")

		case 0xCA:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xCA")
			fmt.Println(" Instruction: DEX")

		case 0xCB:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xCB")
			fmt.Println(" Instruction: AXS Immediate")

		case 0xCC:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xCC")
			fmt.Println(" Instruction: CPY Absolute")

		case 0xCD:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xCD")
			fmt.Println(" Instruction: CMP Absolute")

		case 0xCE:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xCE")
			fmt.Println(" Instruction: DEC Absolute")

		case 0xCF:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xCF")
			fmt.Println(" Instruction: DCP Absolute")

		case 0xD0:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xD0")
			fmt.Println(" Instruction: BNE Relative")

		case 0xD1:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xD1")
			fmt.Println(" Instruction: CMP (Indirect),Y")

		case 0xD2:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xD2")
			fmt.Println(" Instruction: KIL")

		case 0xD3:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xD3")
			fmt.Println(" Instruction: DCP (Indirect),Y")

		case 0xD4:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xD4")
			fmt.Println(" Instruction: NOP Zero Page,X")

		case 0xD5:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xD5")
			fmt.Println(" Instruction: CMP Zero Page,X")

		case 0xD6:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xD6")
			fmt.Println(" Instruction: DEC Zero Page,X")

		case 0xD7:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xD7")
			fmt.Println(" Instruction: DCP Zero Page,X")

		case 0xD8:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xD8")
			fmt.Println(" Instruction: CLD")

		case 0xD9:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xD9")
			fmt.Println(" Instruction: CMP Absolute,Y")

		case 0xDA:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xDA")
			fmt.Println(" Instruction: NOP")

		case 0xDB:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xDB")
			fmt.Println(" Instruction: DCP Absolute,Y")

		case 0xDC:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xDC")
			fmt.Println(" Instruction: NOP Absolute,X")

		case 0xDD:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xDD")
			fmt.Println(" Instruction: CMP Absolute,X")

		case 0xDE:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xDE")
			fmt.Println(" Instruction: DEC Absolute,X")

		case 0xDF:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xDF")
			fmt.Println(" Instruction: DCP Absolute,X")

		case 0xE0:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xE0")
			fmt.Println(" Instruction: CPX Immediate")

		case 0xE1:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xE1")
			fmt.Println(" Instruction: SBC (Indirect,X)")

		case 0xE2:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xE2")
			fmt.Println(" Instruction: NOP Immediate")

		case 0xE3:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xE3")
			fmt.Println(" Instruction: ISC (Indirect,X)")

		case 0xE4:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xE4")
			fmt.Println(" Instruction: CPX Zero Page")

		case 0xE5:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xE5")
			fmt.Println(" Instruction: SBC Zero Page")

		case 0xE6:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xE6")
			fmt.Println(" Instruction: INC Zero Page")

		case 0xE7:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xE7")
			fmt.Println(" Instruction: ISC Zero Page")

		case 0xE8:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xE8")
			fmt.Println(" Instruction: INX")

		case 0xE9:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xE9")
			fmt.Println(" Instruction: SBC Immediate")

		case 0xEA:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xEA")
			fmt.Println(" Instruction: NOP")

		case 0xEB:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xEB")
			fmt.Println(" Instruction: SBC Immediate")

		case 0xEC:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xEC")
			fmt.Println(" Instruction: CPX Absolute")

		case 0xED:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xED")
			fmt.Println(" Instruction: SBC Absolute")

		case 0xEE:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xEE")
			fmt.Println(" Instruction: INC Absolute")

		case 0xEF:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xEF")
			fmt.Println(" Instruction: ISC Absolute")

		case 0xF0:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xF0")
			fmt.Println(" Instruction: BEQ")

		case 0xF1:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xF1")
			fmt.Println(" Instruction: SBC (Indirect),Y")

		case 0xF2:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xF2")
			fmt.Println(" Instruction: KIL")

		case 0xF3:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xF3")
			fmt.Println(" Instruction: ISC (Indirect),Y")

		case 0xF4:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xF4")
			fmt.Println(" Instruction: NOP Zero Page,X")

		case 0xF5:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xF5")
			fmt.Println(" Instruction: SBC Zero Page,X")

		case 0xF6:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xF6")
			fmt.Println(" Instruction: INC Zero Page,X")

		case 0xF7:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xF7")
			fmt.Println(" Instruction: ISC Zero Page,X")

		case 0xF8:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xF8")
			fmt.Println(" Instruction: SED")

		case 0xF9:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xF9")
			fmt.Println(" Instruction: SBC Absolute,Y")

		case 0xFA:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xFA")
			fmt.Println(" Instruction: NOP")

		case 0xFB:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xFB")
			fmt.Println(" Instruction: ISC Absolute,Y")

		case 0xFC:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xFC")
			fmt.Println(" Instruction: NOP Absolute,X")

		case 0xFD:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xFD")
			fmt.Println(" Instruction: SBC Absolute,X")

		case 0xFE:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xFE")
			fmt.Println(" Instruction: INC Absolute,X")

		case 0xFF:
			fmt.Printf("Memory Address: $%02x", PC)
			fmt.Printf(" Opcode: 0xFF")
			fmt.Println(" Instruction: ISC Absolute,X")

		default:
			panic(fmt.Sprintf("Unhandled instruction: %v", memory[PC]))
		}
	}
}
