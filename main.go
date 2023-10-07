package main

func main() {
	boilerPlate()
	loadROMs()
	//time.Sleep(2 * time.Second) //Uncomment when screenrecording
	cpu.resetCPU()
	ted.resetTED()
	go cpu.startCPU()
	userInterface()
}
