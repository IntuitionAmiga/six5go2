package main

func main() {
	boilerPlate()
	// initialise every byte in memory array to 0
	for i := 0; i < len(memory); i++ {
		memory[i] = 0
	}
	//fmt.Fprintf(os.Stderr, "% X\n", memory)
	loadROMs()
	cpu.resetCPU()
	ted.resetTED()
	go userInterface()
	cpu.startCPU()
}
