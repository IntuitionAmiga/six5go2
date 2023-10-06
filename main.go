package main

func main() {
	boilerPlate()
	loadROMs()
	cpu.resetCPU()
	ted.resetTED()
	cpu.startCPU()
}
