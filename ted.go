package main

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

// Handle TED specific registers
func readTEDReg(address uint16) byte {
	switch address {
	case TED_CTRL_REG_0:
		return 0
	case TED_CTRL_REG_1:
		return 0
	case TED_INTR_STAT_REG:
		return 0
	case TED_INTR_MASK_REG:
		return 0
	case TED_RASTER_LOW:
		return 0
	case TED_RASTER_HIGH:
		return 0
	case TED_CLOCK_FREQ:
		return 0
	case TED_BORDER_COLOR:
		return 0
	case TED_BG_COLOR_0:
		return 0
	case TED_BG_COLOR_1:
		return 0
	case TED_BG_COLOR_2:
		return 0
	case TED_BG_COLOR_3:
		return 0
	case TED_FG_COLOR_0:
		return 0
	case TED_FG_COLOR_1:
		return 0
	case TED_FG_COLOR_2:
		return 0
	case TED_FG_COLOR_3:
		return 0
	case TED_KEYBOARD_LINE0:
		return 0
	case TED_KEYBOARD_LINE1:
		return 0
	case TED_KEYBOARD_LINE2:
		return 0
	case TED_KEYBOARD_LINE3:
		return 0
	case TED_KEYBOARD_LINE4:
		return 0
	case TED_KEYBOARD_LINE5:
		return 0
	case TED_KEYBOARD_LINE6:
		return 0
	case TED_KEYBOARD_LINE7:
		return 0
	case TED_KEYBOARD_LINE8:
		return 0
	case TED_KEYBOARD_LINE9:
		return 0
	default:
		return 0
	}
}

func writeTEDReg(address uint16, value byte) {
	switch address {
	case TED_CTRL_REG_0:
	case TED_CTRL_REG_1:
	case TED_INTR_STAT_REG:
	case TED_INTR_MASK_REG:
	case TED_RASTER_LOW:
	case TED_RASTER_HIGH:
	case TED_CLOCK_FREQ:
	case TED_BORDER_COLOR:
	case TED_BG_COLOR_0:
	case TED_BG_COLOR_1:
	case TED_BG_COLOR_2:
	case TED_BG_COLOR_3:
	case TED_FG_COLOR_0:
	case TED_FG_COLOR_1:
	case TED_FG_COLOR_2:
	case TED_FG_COLOR_3:
	case TED_KEYBOARD_LINE0:
	case TED_KEYBOARD_LINE1:
	case TED_KEYBOARD_LINE2:
	case TED_KEYBOARD_LINE3:
	case TED_KEYBOARD_LINE4:
	case TED_KEYBOARD_LINE5:
	case TED_KEYBOARD_LINE6:
	case TED_KEYBOARD_LINE7:
	case TED_KEYBOARD_LINE8:
	case TED_KEYBOARD_LINE9:
	default:
		// Default case
	}
}
