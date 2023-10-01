package main

func getSRBit(x byte) byte {
	return (SR >> x) & 1
}
func setSRBitOn(x byte) {
	SR |= 1 << x
}
func setSRBitOff(x byte) {
	SR &= ^(1 << x)
}
func getABit(x byte) byte {
	return (A >> x) & 1
}
func getXBit(x byte) byte {
	return (X >> x) & 1
}
func getYBit(x byte) byte {
	return (Y >> x) & 1
}

func setNegativeFlag() {
	setSRBitOn(7)
}
func unsetNegativeFlag() {
	setSRBitOff(7)
}
func setOverflowFlag() {
	setSRBitOn(6)
}
func unsetOverflowFlag() {
	setSRBitOff(6)
}
func setBreakFlag() {
	setSRBitOn(4)
}
func setDecimalFlag() {
	setSRBitOn(3)
}
func unsetDecimalFlag() {
	setSRBitOff(3)
}
func setInterruptFlag() {
	setSRBitOn(2)
}
func unsetInterruptFlag() {
	setSRBitOff(2)
}
func setZeroFlag() {
	setSRBitOn(1)
}
func unsetZeroFlag() {
	setSRBitOff(1)
}
func setCarryFlag() {
	setSRBitOn(0)
}
func unsetCarryFlag() {
	setSRBitOff(0)
}
