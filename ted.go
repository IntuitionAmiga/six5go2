package main

import "fmt"

type TED struct {
	// TED registers
	TED_CTRL_REG_0     byte
	TED_CTRL_REG_1     byte
	TED_INTR_STAT_REG  byte
	TED_INTR_MASK_REG  byte
	TED_RASTER_LOW     byte
	TED_RASTER_HIGH    byte
	TED_CLOCK_FREQ     byte
	TED_BORDER_COLOR   byte
	TED_BG_COLOR_0     byte
	TED_BG_COLOR_1     byte
	TED_BG_COLOR_2     byte
	TED_BG_COLOR_3     byte
	TED_FG_COLOR_0     byte
	TED_FG_COLOR_1     byte
	TED_FG_COLOR_2     byte
	TED_FG_COLOR_3     byte
	TED_KEYBOARD_LINE0 byte
	TED_KEYBOARD_LINE1 byte
	TED_KEYBOARD_LINE2 byte
	TED_KEYBOARD_LINE3 byte
	TED_KEYBOARD_LINE4 byte
	TED_KEYBOARD_LINE5 byte
	TED_KEYBOARD_LINE6 byte
	TED_KEYBOARD_LINE7 byte
	TED_KEYBOARD_LINE8 byte
	TED_KEYBOARD_LINE9 byte
}

var ted TED

// Constants for TED registers
const (
	TED_CTRL_REG_0     = 0xFD00
	TED_CTRL_REG_1     = 0xFD01
	TED_INTR_STAT_REG  = 0xFD02
	TED_INTR_MASK_REG  = 0xFD03
	TED_RASTER_LOW     = 0xFD04
	TED_RASTER_HIGH    = 0xFD05
	TED_CLOCK_FREQ     = 0xFD06
	TED_BORDER_COLOR   = 0xFD07
	TED_BG_COLOR_0     = 0xFD08
	TED_BG_COLOR_1     = 0xFD09
	TED_BG_COLOR_2     = 0xFD0A
	TED_BG_COLOR_3     = 0xFD0B
	TED_FG_COLOR_0     = 0xFD0C
	TED_FG_COLOR_1     = 0xFD0D
	TED_FG_COLOR_2     = 0xFD0E
	TED_FG_COLOR_3     = 0xFD0F
	TED_KEYBOARD_LINE0 = 0xFD10
	TED_KEYBOARD_LINE1 = 0xFD11
	TED_KEYBOARD_LINE2 = 0xFD12
	TED_KEYBOARD_LINE3 = 0xFD13
	TED_KEYBOARD_LINE4 = 0xFD14
	TED_KEYBOARD_LINE5 = 0xFD15
	TED_KEYBOARD_LINE6 = 0xFD16
	TED_KEYBOARD_LINE7 = 0xFD17
	TED_KEYBOARD_LINE8 = 0xFD18
	TED_KEYBOARD_LINE9 = 0xFD19
)

func (ted *TED) resetTED() {
	ted.TED_CTRL_REG_0 = 0x00
	ted.TED_CTRL_REG_1 = 0x00
	ted.TED_INTR_STAT_REG = 0x00
	ted.TED_INTR_MASK_REG = 0x00
	ted.TED_RASTER_LOW = 0x00
	ted.TED_RASTER_HIGH = 0x00
	ted.TED_CLOCK_FREQ = 0x00
	ted.TED_BORDER_COLOR = 0x00
	ted.TED_BG_COLOR_0 = 0x00
	ted.TED_BG_COLOR_1 = 0x00
	ted.TED_BG_COLOR_2 = 0x00
	ted.TED_BG_COLOR_3 = 0x00
	ted.TED_FG_COLOR_0 = 0x00
	ted.TED_FG_COLOR_1 = 0x00
	ted.TED_FG_COLOR_2 = 0x00
	ted.TED_FG_COLOR_3 = 0x00
	ted.TED_KEYBOARD_LINE0 = 0x00
	ted.TED_KEYBOARD_LINE1 = 0x00
	ted.TED_KEYBOARD_LINE2 = 0x00
	ted.TED_KEYBOARD_LINE3 = 0x00
	ted.TED_KEYBOARD_LINE4 = 0x00
	ted.TED_KEYBOARD_LINE5 = 0x00
	ted.TED_KEYBOARD_LINE6 = 0x00
	ted.TED_KEYBOARD_LINE7 = 0x00
	ted.TED_KEYBOARD_LINE8 = 0x00
	ted.TED_KEYBOARD_LINE9 = 0x00
}

// Handle TED specific registers
func (ted *TED) readTEDReg(address uint16) byte {
	fmt.Printf("Reading TED register %04X\n", address)
	switch address {
	case TED_CTRL_REG_0:
		fmt.Printf("Reading TED register %04X TED_CTRL_REG_0\n", address)
		return ted.TED_CTRL_REG_0
	case TED_CTRL_REG_1:
		fmt.Printf("Reading TED register %04X TED_CTRL_REG_1\n", address)
		return ted.TED_CTRL_REG_1
	case TED_INTR_STAT_REG:
		fmt.Printf("Reading TED register %04X TED_INTR_STAT_REG\n", address)
		// Clear the interrupt flag when this register is read
		cpu.unsetInterruptFlag()
		return ted.TED_INTR_STAT_REG
	case TED_INTR_MASK_REG:
		fmt.Printf("Reading TED register %04X TED_INTR_MASK_REG\n", address)
		return ted.TED_INTR_MASK_REG
	case TED_RASTER_LOW:
		fmt.Printf("Reading TED register %04X TED_RASTER_LOW\n", address)
		return ted.TED_RASTER_LOW
	case TED_RASTER_HIGH:
		fmt.Printf("Reading TED register %04X TED_RASTER_HIGH\n", address)
		return ted.TED_RASTER_HIGH
	case TED_CLOCK_FREQ:
		fmt.Printf("Reading TED register %04X TED_CLOCK_FREQ\n", address)
		return ted.TED_CLOCK_FREQ
	case TED_BORDER_COLOR:
		fmt.Printf("Reading TED register %04X TED_BORDER_COLOR\n", address)
		return ted.TED_BORDER_COLOR
	case TED_BG_COLOR_0:
		fmt.Printf("Reading TED register %04X TED_BG_COLOR_0\n", address)
		return ted.TED_BG_COLOR_0
	case TED_BG_COLOR_1:
		fmt.Printf("Reading TED register %04X TED_BG_COLOR_1\n", address)
		return ted.TED_BG_COLOR_1
	case TED_BG_COLOR_2:
		fmt.Printf("Reading TED register %04X TED_BG_COLOR_2\n", address)
		return ted.TED_BG_COLOR_2
	case TED_BG_COLOR_3:
		fmt.Printf("Reading TED register %04X TED_BG_COLOR_3\n", address)
		return ted.TED_BG_COLOR_3
	case TED_FG_COLOR_0:
		fmt.Printf("Reading TED register %04X TED_FG_COLOR_0\n", address)
		return ted.TED_FG_COLOR_0
	case TED_FG_COLOR_1:
		fmt.Printf("Reading TED register %04X TED_FG_COLOR_1\n", address)
		return ted.TED_FG_COLOR_1
	case TED_FG_COLOR_2:
		fmt.Printf("Reading TED register %04X TED_FG_COLOR_2\n", address)
		return ted.TED_FG_COLOR_2
	case TED_FG_COLOR_3:
		fmt.Printf("Reading TED register %04X TED_FG_COLOR_3\n", address)
		return ted.TED_FG_COLOR_3
	case TED_KEYBOARD_LINE0:
		fmt.Printf("Reading TED register %04X TED_KEYBOARD_LINE0\n", address)
		return ted.TED_KEYBOARD_LINE0
	case TED_KEYBOARD_LINE1:
		fmt.Printf("Reading TED register %04X TED_KEYBOARD_LINE1\n", address)
		return ted.TED_KEYBOARD_LINE1
	case TED_KEYBOARD_LINE2:
		fmt.Printf("Reading TED register %04X TED_KEYBOARD_LINE2\n", address)
		return ted.TED_KEYBOARD_LINE2
	case TED_KEYBOARD_LINE3:
		fmt.Printf("Reading TED register %04X TED_KEYBOARD_LINE3\n", address)
		return ted.TED_KEYBOARD_LINE3
	case TED_KEYBOARD_LINE4:
		fmt.Printf("Reading TED register %04X TED_KEYBOARD_LINE4\n", address)
		return ted.TED_KEYBOARD_LINE4
	case TED_KEYBOARD_LINE5:
		fmt.Printf("Reading TED register %04X TED_KEYBOARD_LINE5\n", address)
		return ted.TED_KEYBOARD_LINE5
	case TED_KEYBOARD_LINE6:
		fmt.Printf("Reading TED register %04X TED_KEYBOARD_LINE6\n", address)
		return ted.TED_KEYBOARD_LINE6
	case TED_KEYBOARD_LINE7:
		fmt.Printf("Reading TED register %04X TED_KEYBOARD_LINE7\n", address)
		return ted.TED_KEYBOARD_LINE7
	case TED_KEYBOARD_LINE8:
		fmt.Printf("Reading TED register %04X TED_KEYBOARD_LINE8\n", address)
		return ted.TED_KEYBOARD_LINE8
	case TED_KEYBOARD_LINE9:
		fmt.Printf("Reading TED register %04X TED_KEYBOARD_LINE9\n", address)
		return ted.TED_KEYBOARD_LINE9
	default:
		return 0
	}
}
func (ted *TED) writeTEDReg(address uint16, value byte) {
	switch address {
	case TED_CTRL_REG_0:
		fmt.Printf("Writing %02X to TED register %04X TED_CTRL_REG_0\n", value, address)
		ted.TED_CTRL_REG_0 = value
	case TED_CTRL_REG_1:
		fmt.Printf("Writing %02X to TED register %04X TED_CTRL_REG_1\n", value, address)
		ted.TED_CTRL_REG_1 = value
		if readBit(0, value) == 1 {
			fmt.Println("IRQ request!")
			cpu.irq = true
		} else {
			fmt.Println("IRQ request cleared!")
			cpu.irq = false
		}
		break
	case TED_INTR_STAT_REG:
		fmt.Printf("Writing %02X to TED register %04X TED_INTR_STAT_REG\n", value, address)
		// Writing a 1 to a bit clears the corresponding interrupt flag
		ted.TED_INTR_STAT_REG &= ^value
	case TED_INTR_MASK_REG:
		fmt.Printf("Writing %02X to TED register %04X TED_INTR_MASK_REG\n", value, address)
		ted.TED_INTR_MASK_REG = value
		// Check if we need to trigger an interrupt based on the mask
		if ted.TED_INTR_STAT_REG&ted.TED_INTR_MASK_REG != 0 {
			cpu.irq = true
		}
	case TED_RASTER_LOW:
		fmt.Printf("Writing %02X to TED register %04X TED_RASTER_LOW\n", value, address)
		ted.TED_RASTER_LOW = value
	case TED_RASTER_HIGH:
		fmt.Printf("Writing %02X to TED register %04X TED_RASTER_HIGH\n", value, address)
		ted.TED_RASTER_HIGH = value
	case TED_CLOCK_FREQ:
		fmt.Printf("Writing %02X to TED register %04X TED_CLOCK_FREQ\n", value, address)
		ted.TED_CLOCK_FREQ = value
	case TED_BORDER_COLOR:
		fmt.Printf("Writing %02X to TED register %04X TED_BORDER_COLOR\n", value, address)
		ted.TED_BORDER_COLOR = value
	case TED_BG_COLOR_0:
		fmt.Printf("Writing %02X to TED register %04X TED_BG_COLOR_0\n", value, address)
		ted.TED_BG_COLOR_0 = value
	case TED_BG_COLOR_1:
		fmt.Printf("Writing %02X to TED register %04X TED_BG_COLOR_1\n", value, address)
		ted.TED_BG_COLOR_1 = value
	case TED_BG_COLOR_2:
		fmt.Printf("Writing %02X to TED register %04X TED_BG_COLOR_2\n", value, address)
		ted.TED_BG_COLOR_2 = value
	case TED_BG_COLOR_3:
		fmt.Printf("Writing %02X to TED register %04X TED_BG_COLOR_3\n", value, address)
		ted.TED_BG_COLOR_3 = value
	case TED_FG_COLOR_0:
		fmt.Printf("Writing %02X to TED register %04X TED_FG_COLOR_0\n", value, address)
		ted.TED_FG_COLOR_0 = value
	case TED_FG_COLOR_1:
		fmt.Printf("Writing %02X to TED register %04X TED_FG_COLOR_1\n", value, address)
		ted.TED_FG_COLOR_1 = value
	case TED_FG_COLOR_2:
		fmt.Printf("Writing %02X to TED register %04X TED_FG_COLOR_2\n", value, address)
		ted.TED_FG_COLOR_2 = value
	case TED_FG_COLOR_3:
		fmt.Printf("Writing %02X to TED register %04X TED_FG_COLOR_3\n", value, address)
		ted.TED_FG_COLOR_3 = value
	case TED_KEYBOARD_LINE0:
		fmt.Printf("Writing %02X to TED register %04X TED_KEYBOARD_LINE0\n", value, address)
		ted.TED_KEYBOARD_LINE0 = value
		break
	case TED_KEYBOARD_LINE1:
		fmt.Printf("Writing %02X to TED register %04X TED_KEYBOARD_LINE1\n", value, address)
		ted.TED_KEYBOARD_LINE1 = value
	case TED_KEYBOARD_LINE2:
		fmt.Printf("Writing %02X to TED register %04X TED_KEYBOARD_LINE2\n", value, address)
		ted.TED_KEYBOARD_LINE2 = value
	case TED_KEYBOARD_LINE3:
		fmt.Printf("Writing %02X to TED register %04X TED_KEYBOARD_LINE3\n", value, address)
		ted.TED_KEYBOARD_LINE3 = value
	case TED_KEYBOARD_LINE4:
		fmt.Printf("Writing %02X to TED register %04X TED_KEYBOARD_LINE4\n", value, address)
		ted.TED_KEYBOARD_LINE4 = value
	case TED_KEYBOARD_LINE5:
		fmt.Printf("Writing %02X to TED register %04X TED_KEYBOARD_LINE5\n", value, address)
		ted.TED_KEYBOARD_LINE5 = value
	case TED_KEYBOARD_LINE6:
		fmt.Printf("Writing %02X to TED register %04X TED_KEYBOARD_LINE6\n", value, address)
		ted.TED_KEYBOARD_LINE6 = value
	case TED_KEYBOARD_LINE7:
		fmt.Printf("Writing %02X to TED register %04X TED_KEYBOARD_LINE7\n", value, address)
		ted.TED_KEYBOARD_LINE7 = value
	case TED_KEYBOARD_LINE8:
		fmt.Printf("Writing %02X to TED register %04X TED_KEYBOARD_LINE8\n", value, address)
		ted.TED_KEYBOARD_LINE8 = value
	case TED_KEYBOARD_LINE9:
		fmt.Printf("Writing %02X to TED register %04X TED_KEYBOARD_LINE9\n", value, address)
		ted.TED_KEYBOARD_LINE9 = value
	default:
		return
	}
}
