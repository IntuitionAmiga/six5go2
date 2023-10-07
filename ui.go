package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
	"time"
)

func userInterface() {
	app := tview.NewApplication()

	// Create CPU State Panel
	cpuState := tview.NewTextView()
	go func() {
		ticker := time.NewTicker(time.Millisecond) // Update every 500ms
		for range ticker.C {
			app.QueueUpdateDraw(func() {
				formattedText := fmt.Sprintf("A:$%02X X:$%02X Y:$%02X SP:$%04X", cpu.A, cpu.X, cpu.Y, SPBaseAddress+cpu.SP)
				cpuState.SetText(formattedText)

			})
		}
	}()
	cpuState.SetBorder(true).SetTitle(" CPU State ")

	// Create Trace Panel
	trace := tview.NewTextView().SetText("Execution and disassembly")
	go func() {
		ticker := time.NewTicker(time.Millisecond) // Update every 500ms
		for range ticker.C {
			app.QueueUpdateDraw(func() {
				traceLine := executionTrace()
				formattedText := fmt.Sprintf("%s", traceLine)
				trace.SetText(formattedText)

			})
		}
	}()
	trace.SetBorder(true).SetTitle(" Trace ")

	// Create Emulated Display Area
	display := tview.NewTextView().SetText("\n\t\tCOMMODORE BASIC V3.5 60671 BYTES FREE\n\t\t3-PLUS-1 ON KEY F1\n\n\t\tREADY.\n\n")
	display.SetBorder(true).SetTitle(" Plus/4 Display ")
	display.SetBackgroundColor(tcell.NewRGBColor(181, 155, 255))

	// Create Toolbar Panel
	toolbar := tview.NewTextView().SetText("[Q]uit  [R]eset [S]ettings")
	toolbar.SetBorder(true).SetTitle(" six5go2 - (c) Zayn Otley ")

	// Create SRFlags Panel
	srFlags := tview.NewTextView().SetText("SR and Stack Contents")
	go func() {
		ticker := time.NewTicker(time.Millisecond) // Update every 500ms
		for range ticker.C {
			app.QueueUpdateDraw(func() {
				statusLine := statusFlags()
				stackLine := fmt.Sprintf("\nStack: $%04X", readStack())
				formattedText := fmt.Sprintf("%s\n", statusLine+stackLine)
				srFlags.SetText(formattedText)

			})
		}
	}()
	srFlags.SetBorder(true).SetTitle(" Status Register & Stack ")

	// Create Counters Panel
	counters := tview.NewTextView().SetText("Cycle and Instruction Counters")
	go func() {
		ticker := time.NewTicker(time.Millisecond) // Update every 500ms
		for range ticker.C {
			app.QueueUpdateDraw(func() {
				instructionsLine := instructionCount()
				cyclesLine := fmt.Sprintf("\nCPU Cycles:\t\t$%08X ", cpu.cycleCounter)
				timeSpentLine := fmt.Sprintf("\nInstr Time:\t %v ", cpu.cpuTimeSpent)
				formattedText := fmt.Sprintf("%s\n", instructionsLine+cyclesLine+timeSpentLine)
				counters.SetText(formattedText)
			})
		}
	}()
	counters.SetBorder(true).SetTitle(" Counters ")

	// Create Disassembly Panel
	disassembly := tview.NewTextView().SetText("Disassembly")
	go func() {
		ticker := time.NewTicker(time.Millisecond) // Update every 500ms
		for range ticker.C {
			app.QueueUpdateDraw(func() {
				disassembly.SetText(disassembledInstruction)
			})
		}
	}()
	disassembly.SetBorder(true).SetTitle(" Disassembly ")

	// Create layout
	grid := tview.NewGrid().
		SetRows(3, 4, 5, 3, 3, 3, 0). // Set -1 for display to take remaining space
		SetColumns(28, 0).
		AddItem(cpuState, 0, 0, 1, 1, 0, 0, false).
		AddItem(srFlags, 1, 0, 1, 1, 0, 0, false).
		AddItem(counters, 2, 0, 1, 1, 0, 0, false).
		AddItem(trace, 3, 0, 1, 1, 0, 0, false).
		AddItem(disassembly, 4, 0, 1, 1, 0, 0, false).
		AddItem(toolbar, 9, 0, 1, 1, 0, 0, false).
		AddItem(display, 0, 1, 10, 1, 0, 0, false)

	// Handle input
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			app.Stop()
			os.Exit(0)
		case 'r':
			fmt.Printf("Resetting...\n")
			// Reset logic here
		}
		return event
	})

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
