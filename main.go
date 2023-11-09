package main

func main() {
	boilerPlate()
	loadROMs()
	//time.Sleep(time.Second / 4) //Uncomment when screenrecording
	cpu.resetCPU()
	ted.resetTED()
	go userInterface()
	cpu.startCPU()
}
