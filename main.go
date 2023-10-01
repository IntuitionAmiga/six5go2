package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Six5go2 v2.0 - 6502 Emulator and Disassembler in Golang (c) 2022-2023 Zayn Otley\n\n")
	fmt.Printf("https://github.com/intuitionamiga/six5go2/tree/v2\n\n")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Printf("Size of addressable memory is %v ($%04X) bytes\n\n", len(memory), len(memory)-1)

	// Start emulation
	loadROMs()
	resetCPU()
	startCPU()
}
