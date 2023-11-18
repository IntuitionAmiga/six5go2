package main

func main() {
	boilerPlate()
	// initialise every byte in memory array to 0
	for i := 0; i < len(memory); i++ {
		memory[i] = 0
	}
	loadROMs()
	//dumpMemoryToFile(plus4basicROMAddress, len(PLUS4BASICROM))
	cpu.resetCPU()
	ted.resetTED()
	go userInterface()
	cpu.startCPU()
}
