package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
	"strconv"
	"strings"
	"time"
)

var lastFoundPosition = -1
var lastSearchString = ""
var lastReverseFoundPosition = -1

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
			modal := tview.NewTextView().SetText("Memory Monitor").
				SetDynamicColors(true).
				SetTextAlign(tview.AlignLeft).
				SetScrollable(true) // Make the TextView scrollable
			modal.SetBorder(true).SetTitle(" Memory Monitor - (Ctrl-G to Goto/Edit Memory Address / Ctrl-F to Search) ")

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
				case tcell.KeyCtrlG:
					var form *tview.Form
					form = tview.NewForm().
						AddInputField("Go to address: ", "", 10, nil, func(text string) {
							// Convert entered address to integer
							address, err := strconv.ParseInt(text, 16, 32)
							if err != nil {
								// Handle error
								return
							}
							// Calculate the line number to jump to
							lineNumber := int(address) / 16
							modal.ScrollTo(lineNumber, 0)
							// Populate the value field
							form.GetFormItem(1).(*tview.InputField).SetText(fmt.Sprintf("%02X", memory[address]))
						}).
						AddInputField("Current Value: ", "", 2, nil, nil).
						AddButton("Update", func() {
							addressField := form.GetFormItem(0).(*tview.InputField).GetText()
							valueField := form.GetFormItem(1).(*tview.InputField).GetText()
							address, err := strconv.ParseInt(addressField, 16, 32)
							if err != nil {
								return
							}
							value, err := strconv.ParseInt(valueField, 16, 8)
							if err != nil {
								return
							}
							memory[address] = byte(value)
						})
					form.SetBorder(true).SetTitle(" Enter address ").SetTitleAlign(tview.AlignLeft)
					form.AddButton("Exit", func() {
						pages.RemovePage("gotoAddressForm")
						app.SetFocus(modal)
					})
					// existing code for formGrid, etc.

					// Create a new grid layout for the form, specifying its size
					formGrid := tview.NewGrid().
						SetRows(10).
						SetColumns(30).
						AddItem(form, 0, 0, 1, 1, 0, 0, true)

					// Add this new grid layout as a new page
					pages.AddPage("gotoAddressForm", formGrid, true, true)

					// Set focus to the new form grid layout
					app.SetFocus(formGrid)
				case tcell.KeyCtrlF:
					form := tview.NewForm().
						AddInputField(" Find string: ", "", 10, nil, func(text string) {
							// Search for the string in the memory and find the corresponding line
							lineIndex := findStringInMemory(text, memory[:], true)
							if lineIndex != -1 {
								modal.ScrollTo(lineIndex, 0)
								lastFoundPosition = lineIndex * 16 // Update the last found position
								lastSearchString = text            // Update the last search string
							}
						})
					form.SetBorder(true).SetTitle(" Enter string ").SetTitleAlign(tview.AlignLeft)
					form.AddButton("Next", func() {
						if lastSearchString != "" {
							// Use KMP search starting from lastFoundPosition
							newLineIndex := findStringInMemory(lastSearchString, memory[:], true)
							if newLineIndex != -1 {
								modal.ScrollTo(newLineIndex, 0)
							}
						}
					})
					form.AddButton("Previous", func() {
						if lastSearchString != "" {
							// Use KMP search starting from lastFoundPosition
							newLineIndex := findStringInMemory(lastSearchString, memory[:], false)
							if newLineIndex != -1 {
								modal.ScrollTo(newLineIndex, 0)
							}
						}
					})

					form.AddButton("Exit", func() {
						pages.RemovePage("findStringForm")
						app.SetFocus(modal)
					})

					// Create a new grid layout for the form, specifying its size (same as for Ctrl-G)
					formGrid := tview.NewGrid().
						SetRows(8).
						SetColumns(34).
						AddItem(form, 0, 0, 1, 1, 0, 0, true)

					// Add this new grid layout as a new page
					pages.AddPage("findStringForm", formGrid, true, true)

					// Set focus to the new form grid layout
					app.SetFocus(formGrid)

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

// KMP Search algorithm
func KMPSearch(pat string, txt string) bool {
	if pat == "" {
		return false
	}

	m := len(pat)
	n := len(txt)

	// create lps[] that will hold the longest prefix suffix values for pattern
	var lps []int
	lps = make([]int, m)
	var j = 0 // index for pat[]

	// Preprocess the pattern (calculate lps[] array)
	computeLPSArray(pat, m, lps)

	i := 0 // index for txt[]
	for i < n {
		if pat[j] == txt[i] {
			j++
			i++
		}

		if j == m {
			return true
			//j = lps[j-1] // For use when enabling Find Next Occurrence feature
		}

		// mismatch after j matches
		if i < n && pat[j] != txt[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i = i + 1
			}
		}
	}
	return false
}

// Fills lps[] for given patttern pat[0..M-1]
func computeLPSArray(pat string, M int, lps []int) {
	length := 0 // length of the previous longest prefix suffix

	lps[0] = 0 // lps[0] is always 0
	i := 1

	// the loop calculates lps[i] for i = 1 to M-1
	for i < M {
		if pat[i] == pat[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
}

// Convert string to its hexadecimal representation with spaces
func stringToHexSpace(str string) string {
	var hexStr string
	for i := 0; i < len(str); i++ {
		hexStr += fmt.Sprintf("%02X ", str[i])
	}
	return hexStr
}

func findStringInMemory(search string, memory []byte, searchForward bool) int {
	if search == "" {
		return -1
	}

	var line string
	hexSearch := stringToHexSpace(search)

	// Determine starting position
	var i int
	if search == lastSearchString {
		if searchForward && lastFoundPosition < len(memory) {
			i = lastFoundPosition + 16 // Start after the last found position
		} else if !searchForward && lastReverseFoundPosition > 0 {
			i = lastReverseFoundPosition - 16 // Start before the last reverse found position
		} else if !searchForward {
			i = len(memory) - 16 // Start from the end if searching backward
		}
	} else {
		i = 0 // Start from the beginning if searching forward
	}

	if searchForward {
		for ; i < len(memory); i += 16 {
			// Create the line string from the memory slice
			for j := 0; j < 16; j++ {
				line += fmt.Sprintf("%02X ", memory[i+j])
			}
			// Check if the line contains the search string using KMP
			if KMPSearch(hexSearch, line) {
				lastFoundPosition = i
				lastReverseFoundPosition = i // Update the last reverse found position too
				lastSearchString = search
				return i / 16
			}
			line = ""
		}
	} else {
		for ; i >= 0; i -= 16 {
			// Create the line string from the memory slice
			for j := 0; j < 16; j++ {
				line += fmt.Sprintf("%02X ", memory[i+j])
			}
			// Check if the line contains the search string using KMP
			if KMPSearch(hexSearch, line) {
				lastReverseFoundPosition = i
				lastFoundPosition = i // Update the last found position too
				lastSearchString = search
				return i / 16
			}
			line = ""
		}
	}
	return -1
}
