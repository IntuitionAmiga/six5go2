package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
	"strings"
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
	go func() {
		for {
			renderASCII(&ted, app, display)   // Assuming ted is an instance of your TED struct
			time.Sleep(16 * time.Millisecond) // Approx. 60 FPS
		}
	}()
	display.SetBorder(true).SetTitle(" Plus/4 Display ")
	display.SetBackgroundColor(tcell.NewRGBColor(181, 155, 255))

	// Create Toolbar Panel
	toolbar := tview.NewTextView().SetText("[Q]uit  [R]eset [M]emory")
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

	// Create a tview.Pages widget to manage layers
	pages := tview.NewPages()
	// Add your main grid to the pages widget
	pages.AddPage("main", grid, true, true)

	// Handle input
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			app.Stop()
			os.Exit(0)
		case 'r':
			// Resetting logic
			cpu.handleRESET()
			ted.resetTED()
			loadROMs()
		case 'm':
			// Create modal
			modal := tview.NewTextView().
				SetDynamicColors(true).
				SetTextAlign(tview.AlignLeft).
				SetScrollable(true) // Make the TextView scrollable

			modalGrid := tview.NewGrid().
				SetRows(40).
				SetColumns(0).
				AddItem(modal, 0, 0, 1, 1, 0, 0, true)

			// Add modal to the pages widget
			pages.AddPage("modalGrid", modalGrid, true, false) // Switch to the modal page
			pages.SwitchToPage("modalGrid")

			// Set input capture for modal to close it when Esc is pressed
			var scrollPosition int
			var stopUpdate = make(chan bool)
			modal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				switch event.Key() {
				case tcell.KeyEscape:
					pages.SwitchToPage("main")
					stopUpdate <- true
				case tcell.KeyUp:
					if scrollPosition > 0 {
						scrollPosition--
						modal.ScrollTo(scrollPosition, 0)
					}
				case tcell.KeyDown:
					scrollPosition++
					modal.ScrollTo(scrollPosition, 0)
				}
				return event
			})

			// Goroutine to update the TextView
			go func() {
				for {
					select {
					case <-stopUpdate:
						return
					default:
						var builder strings.Builder
						// Pre-allocate some memory
						builder.Grow(0xFFFF * 4)
						// Format the RAM data
						for i := 0; i < 0xFFFF; i += 16 {
							builder.WriteString(fmt.Sprintf("$%04X: ", i))
							for j := 0; j < 16; j++ {
								builder.WriteString(fmt.Sprintf("%02X ", memory[(i+j)]))
							}
							builder.WriteString("  ")
							for j := 0; j < 16; j++ {
								char := memory[(i + j)]
								if char >= 32 && char <= 126 {
									builder.WriteByte(char)
								} else {
									builder.WriteByte('.')
								}
							}
							builder.WriteByte('\n')
						}

						formattedRAM := builder.String()
						app.QueueUpdateDraw(func() {
							modal.SetText(formattedRAM)
						})
						time.Sleep(500 * time.Millisecond)
					}
				}
			}()
		}
		return event

	})
	// Set the pages widget as the root
	app.SetRoot(pages, true)

	if err := app.Run(); err != nil {
		panic(err)
	}
}

// Function to render ASCII-based graphics in the terminal
func renderASCII(t *TED, app *tview.Application, textView *tview.TextView) {
	var asciiArt string
	for y := 0; y < 200; y++ {
		for x := 0; x < 320; x++ {
			pixel := t.GetPixel(x, y)
			if pixel == 1 {
				asciiArt += "▓" // ASCII 254 or any suitable character to represent a pixel
			} else {
				asciiArt += " " // Space for empty pixel
			}
		}
		asciiArt += "\n" // New line at the end of each row
	}

	app.QueueUpdateDraw(func() {
		textView.SetText(asciiArt)
	})
}
