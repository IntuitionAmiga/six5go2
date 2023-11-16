package main

func main() {
	boilerPlate()
	// initialise every byte in memory array to 0
	for i := 0; i < len(memory); i++ {
		memory[i] = 0
	}
	loadROMs()
	cpu.resetCPU()
	ted.resetTED()
	go userInterface()
	cpu.startCPU()
}
