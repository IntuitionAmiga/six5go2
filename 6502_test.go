package main

import (
	"testing"
)

func TestLDAImmediate(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, LDA_IMMEDIATE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the accumulator has the expected value
	if cpu.A != 0x10 {
		t.Errorf("LDA Immediate failed: got %02X, want %02X", cpu.A, 0x10)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for LDA immediate
		t.Errorf("LDA immediate failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestLDAZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, LDA_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x20)
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the accumulator has the expected value
	if cpu.A != 0x20 {
		t.Errorf("LDA Zero Page failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for LDA immediate
		t.Errorf("LDA Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestLDAZeroPageX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, LDA_ZERO_PAGE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x20)
	cpu.preOpX = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the accumulator has the expected value
	if cpu.A != 0x20 {
		t.Errorf("LDA Zero Page X failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for LDA immediate
		t.Errorf("LDA Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestLDAAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, LDA_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(cpu.PC+2, 0x00)
	cpu.writeMemory(0x0010, 0x20)
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the accumulator has the expected value
	if cpu.A != 0x20 {
		t.Errorf("LDA Absolute failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for LDA immediate
		t.Errorf("LDA Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestLDAAbsoluteX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, LDA_ABSOLUTE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(cpu.PC+2, 0x00)
	cpu.writeMemory(0x0011, 0x20)
	cpu.X = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the accumulator has the expected value
	if cpu.A != 0x20 {
		t.Errorf("LDA Absolute X failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for LDA immediate
		t.Errorf("LDA Absolute X failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestLDAAbsoluteY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, LDA_ABSOLUTE_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(cpu.PC+2, 0x00)
	cpu.writeMemory(0x0011, 0x20)
	cpu.Y = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the accumulator has the expected value
	if cpu.A != 0x20 {
		t.Errorf("LDA Absolute Y failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for LDA immediate
		t.Errorf("LDA Absolute Y failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestLDAIndirectX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, LDA_INDIRECT_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x20)
	cpu.writeMemory(0x0011, 0x00)
	cpu.writeMemory(0x0020, 0x30)
	cpu.preOpX = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the accumulator has the expected value
	if cpu.A != 0x30 {
		t.Errorf("LDA Indirect X failed: got %02X, want %02X", cpu.A, 0x30)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for LDA immediate
		t.Errorf("LDA Indirect X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestLDAIndirectY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, LDA_INDIRECT_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x20)
	cpu.writeMemory(0x0011, 0x00)
	cpu.writeMemory(0x0021, 0x30)
	cpu.Y = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the accumulator has the expected value
	if cpu.A != 0x30 {
		t.Errorf("LDA Indirect Y failed: got %02X, want %02X", cpu.A, 0x30)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for LDA immediate
		t.Errorf("LDA Indirect Y failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}

func TestLDXImmediate(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// LDX #10
	cpu.writeMemory(cpu.PC, LDX_IMMEDIATE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the X register has the expected value
	if cpu.X != 0x10 {
		t.Errorf("LDX Immediate failed: got %02X, want %02X", cpu.X, 0x10)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for LDX immediate
		t.Errorf("LDX immediate failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestLDXZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// LDX $10
	cpu.writeMemory(cpu.PC, LDX_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x20)
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the X register has the expected value
	if cpu.X != 0x20 {
		t.Errorf("LDX Zero Page failed: got %02X, want %02X", cpu.X, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for LDX immediate
		t.Errorf("LDX Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestLDXZeroPageY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// LDX $10,Y
	cpu.writeMemory(cpu.PC, LDX_ZERO_PAGE_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0011, 0x20)
	cpu.Y = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the X register has the expected value
	if cpu.X != 0x20 {
		t.Errorf("LDX Zero Page Y failed: got %02X, want %02X", cpu.X, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for LDX immediate
		t.Errorf("LDX Zero Page Y failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestLDXAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// LDX $1000
	cpu.writeMemory(cpu.PC, LDX_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1000, 0x20)
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the X register has the expected value
	if cpu.X != 0x20 {
		t.Errorf("LDX Absolute failed: got %02X, want %02X", cpu.X, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for LDX immediate
		t.Errorf("LDX Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestLDXAbsoluteY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// LDX $1000,Y
	cpu.writeMemory(cpu.PC, LDX_ABSOLUTE_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1001, 0x20)
	cpu.Y = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the X register has the expected value
	if cpu.X != 0x20 {
		t.Errorf("LDX Absolute Y failed: got %02X, want %02X", cpu.X, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for LDX immediate
		t.Errorf("LDX Absolute Y failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}

func TestLDYImmediate(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// LDY #10
	cpu.writeMemory(cpu.PC, LDY_IMMEDIATE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the Y register has the expected value
	if cpu.Y != 0x10 {
		t.Errorf("LDY Immediate failed: got %02X, want %02X", cpu.Y, 0x10)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for LDY immediate
		t.Errorf("LDY immediate failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestLDYZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// LDY $10
	cpu.writeMemory(cpu.PC, LDY_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x20)
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the Y register has the expected value
	if cpu.Y != 0x20 {
		t.Errorf("LDY Zero Page failed: got %02X, want %02X", cpu.Y, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for LDY immediate
		t.Errorf("LDY Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestLDYZeroPageX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// LDY $10,X
	cpu.writeMemory(cpu.PC, LDY_ZERO_PAGE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0011, 0x20)
	cpu.X = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the Y register has the expected value
	if cpu.Y != 0x20 {
		t.Errorf("LDY Zero Page X failed: got %02X, want %02X", cpu.Y, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for LDY immediate
		t.Errorf("LDY Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestLDYAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// LDY $1000
	cpu.writeMemory(cpu.PC, LDY_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1000, 0x20)
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the Y register has the expected value
	if cpu.Y != 0x20 {
		t.Errorf("LDY Absolute failed: got %02X, want %02X", cpu.Y, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for LDY immediate
		t.Errorf("LDY Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestLDYAbsoluteX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// LDY $1000,X
	cpu.writeMemory(cpu.PC, LDY_ABSOLUTE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1001, 0x20)
	cpu.X = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the Y register has the expected value
	if cpu.Y != 0x20 {
		t.Errorf("LDY Absolute X failed: got %02X, want %02X", cpu.Y, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for LDY immediate
		t.Errorf("LDY Absolute X failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}

func TestSTAZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// STA $10
	cpu.writeMemory(cpu.PC, STA_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if memory has the expected value
	if cpu.readMemory(0x0010) != 0x20 {
		t.Errorf("STA Zero Page failed: got %02X, want %02X", cpu.readMemory(0x0010), 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for STA immediate
		t.Errorf("STA Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestSTAZeroPageX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// STA $10,X
	cpu.writeMemory(cpu.PC, STA_ZERO_PAGE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.X = 0x01
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if memory has the expected value
	if cpu.readMemory(0x0011) != 0x20 {
		t.Errorf("STA Zero Page X failed: got %02X, want %02X", cpu.readMemory(0x0011), 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for STA immediate
		t.Errorf("STA Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestSTAAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// STA $1000
	cpu.writeMemory(cpu.PC, STA_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if memory has the expected value
	if cpu.readMemory(0x1000) != 0x20 {
		t.Errorf("STA Absolute failed: got %02X, want %02X", cpu.readMemory(0x1000), 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for STA immediate
		t.Errorf("STA Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestSTAAbsoluteX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// STA $1000,X
	cpu.writeMemory(cpu.PC, STA_ABSOLUTE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.A = 0x20
	cpu.X = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if memory has the expected value
	if cpu.readMemory(0x1001) != 0x20 {
		t.Errorf("STA Absolute X failed: got %02X, want %02X", cpu.readMemory(0x1001), 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for STA immediate
		t.Errorf("STA Absolute X failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestSTAAbsoluteY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// STA $1000,Y
	cpu.writeMemory(cpu.PC, STA_ABSOLUTE_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.A = 0x20
	cpu.Y = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if memory has the expected value
	if cpu.readMemory(0x1001) != 0x20 {
		t.Errorf("STA Absolute Y failed: got %02X, want %02X", cpu.readMemory(0x1001), 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for STA immediate
		t.Errorf("STA Absolute Y failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestSTAIndirectX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// STA ($10,X)
	cpu.writeMemory(cpu.PC, STA_INDIRECT_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x00)
	cpu.writeMemory(0x0011, 0x10)
	cpu.A = 0x20
	cpu.X = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if memory has the expected value
	if cpu.readMemory(0x1000) != 0x20 {
		t.Errorf("STA Indirect X failed: got %02X, want %02X", cpu.readMemory(0x1000), 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for STA immediate
		t.Errorf("STA Indirect X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestSTAIndirectY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// STA ($10),Y
	cpu.writeMemory(cpu.PC, STA_INDIRECT_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x00)
	cpu.writeMemory(0x0011, 0x10)
	cpu.A = 0x20
	cpu.Y = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if memory has the expected value
	if cpu.readMemory(0x1001) != 0x20 {
		t.Errorf("STA Indirect Y failed: got %02X, want %02X", cpu.readMemory(0x1001), 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for STA immediate
		t.Errorf("STA Indirect Y failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}

func TestSTXZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// STX $10
	cpu.writeMemory(cpu.PC, STX_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.X = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if memory has the expected value
	if cpu.readMemory(0x0010) != 0x20 {
		t.Errorf("STX Zero Page failed: got %02X, want %02X", cpu.readMemory(0x0010), 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for STX immediate
		t.Errorf("STX Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestSTXZeroPageY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// STX $10,Y
	cpu.writeMemory(cpu.PC, STX_ZERO_PAGE_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.Y = 0x01
	cpu.X = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if memory has the expected value
	if cpu.readMemory(0x0011) != 0x20 {
		t.Errorf("STX Zero Page Y failed: got %02X, want %02X", cpu.readMemory(0x0011), 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for STX immediate
		t.Errorf("STX Zero Page Y failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestSTXAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// STX $1000
	cpu.writeMemory(cpu.PC, STX_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.X = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if memory has the expected value
	if cpu.readMemory(0x1000) != 0x20 {
		t.Errorf("STX Absolute failed: got %02X, want %02X", cpu.readMemory(0x1000), 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for STX immediate
		t.Errorf("STX Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}

func TestSTYZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// STY $10
	cpu.writeMemory(cpu.PC, STY_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.Y = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if memory has the expected value
	if cpu.readMemory(0x0010) != 0x20 {
		t.Errorf("STY Zero Page failed: got %02X, want %02X", cpu.readMemory(0x0010), 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for STY immediate
		t.Errorf("STY Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestSTYZeroPageX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// STY $10,X
	cpu.writeMemory(cpu.PC, STY_ZERO_PAGE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.X = 0x01
	cpu.Y = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if memory has the expected value
	if cpu.readMemory(0x0011) != 0x20 {
		t.Errorf("STY Zero Page X failed: got %02X, want %02X", cpu.readMemory(0x0011), 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for STY immediate
		t.Errorf("STY Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestSTYAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// STY $1000
	cpu.writeMemory(cpu.PC, STY_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.Y = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if memory has the expected value
	if cpu.readMemory(0x1000) != 0x20 {
		t.Errorf("STY Absolute failed: got %02X, want %02X", cpu.readMemory(0x1000), 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for STY immediate
		t.Errorf("STY Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}

func TestCMPImmediate(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// CMP #10
	cpu.writeMemory(cpu.PC, CMP_IMMEDIATE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Immediate value to compare
	cpu.A = 0x20                    // Set accumulator for comparison

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the zero flag is set correctly
	if cpu.getSRBit(1) != 0 {
		t.Errorf("CMP Immediate failed: Zero flag not set correctly")
	}

	// Check if the negative flag is set correctly
	if cpu.getSRBit(7) != 0 {
		t.Errorf("CMP Immediate failed: Negative flag not set correctly")
	}

	// Check if the carry flag is set correctly (no borrow)
	if cpu.getSRBit(0) != 1 {
		t.Errorf("CMP Immediate failed: Carry flag not set correctly")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for CMP Immediate
		t.Errorf("CMP Immediate failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestCMPZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// CMP $10
	cpu.writeMemory(cpu.PC, CMP_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero Page address to compare
	cpu.writeMemory(0x0010, 0x10)   // Value at Zero Page address to compare
	cpu.A = 0x20                    // Set accumulator for comparison

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the zero flag is set correctly
	if cpu.getSRBit(1) != 0 {
		t.Errorf("CMP Zero Page failed: Zero flag not set correctly")
	}

	// Check if the negative flag is set correctly
	if cpu.getSRBit(7) != 0 {
		t.Errorf("CMP Zero Page failed: Negative flag not set correctly")
	}

	// Check if the carry flag is set correctly (no borrow)
	if cpu.getSRBit(0) != 1 {
		t.Errorf("CMP Zero Page failed: Carry flag not set correctly")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for CMP Zero Page
		t.Errorf("CMP Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestCMPZeroPageX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// CMP $10,X
	cpu.writeMemory(cpu.PC, CMP_ZERO_PAGE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero Page address to compare
	cpu.writeMemory(0x0011, 0x10)   // Value at Zero Page address to compare
	cpu.X = 0x01                    // Set X register for Zero Page addressing
	cpu.A = 0x20                    // Set accumulator for comparison

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the zero flag is set correctly
	if cpu.getSRBit(1) != 0 {
		t.Errorf("CMP Zero Page X failed: Zero flag not set correctly")
	}

	// Check if the negative flag is set correctly
	if cpu.getSRBit(7) != 0 {
		t.Errorf("CMP Zero Page X failed: Negative flag not set correctly")
	}

	// Check if the carry flag is set correctly (no borrow)
	if cpu.getSRBit(0) != 1 {
		t.Errorf("CMP Zero Page X failed: Carry flag not set correctly")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for CMP Zero Page
		t.Errorf("CMP Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestCMPAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// CMP $1000
	cpu.writeMemory(cpu.PC, CMP_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address to compare
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1000, 0x10) // Value at Absolute address to compare
	cpu.A = 0x20                  // Set accumulator for comparison

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the zero flag is set correctly
	if cpu.getSRBit(1) != 0 {
		t.Errorf("CMP Absolute failed: Zero flag not set correctly")
	}

	// Check if the negative flag is set correctly
	if cpu.getSRBit(7) != 0 {
		t.Errorf("CMP Absolute failed: Negative flag not set correctly")
	}

	// Check if the carry flag is set correctly (no borrow)
	if cpu.getSRBit(0) != 1 {
		t.Errorf("CMP Absolute failed: Carry flag not set correctly")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for CMP Absolute
		t.Errorf("CMP Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestCMPAbsoluteX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// CMP $1000,X
	cpu.writeMemory(cpu.PC, CMP_ABSOLUTE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address to compare
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1001, 0x10) // Value at Absolute address to compare
	cpu.X = 0x01                  // Set X register for Absolute addressing
	cpu.A = 0x20                  // Set accumulator for comparison

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the zero flag is set correctly
	if cpu.getSRBit(1) != 0 {
		t.Errorf("CMP Absolute X failed: Zero flag not set correctly")
	}

	// Check if the negative flag is set correctly
	if cpu.getSRBit(7) != 0 {
		t.Errorf("CMP Absolute X failed: Negative flag not set correctly")
	}

	// Check if the carry flag is set correctly (no borrow)
	if cpu.getSRBit(0) != 1 {
		t.Errorf("CMP Absolute X failed: Carry flag not set correctly")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for CMP Absolute
		t.Errorf("CMP Absolute X failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestCMPAbsoluteY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// CMP $1000,Y
	cpu.writeMemory(cpu.PC, CMP_ABSOLUTE_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address to compare
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1001, 0x10) // Value at Absolute address to compare
	cpu.Y = 0x01                  // Set Y register for Absolute addressing
	cpu.A = 0x20                  // Set accumulator for comparison

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the zero flag is set correctly
	if cpu.getSRBit(1) != 0 {
		t.Errorf("CMP Absolute Y failed: Zero flag not set correctly")
	}

	// Check if the negative flag is set correctly
	if cpu.getSRBit(7) != 0 {
		t.Errorf("CMP Absolute Y failed: Negative flag not set correctly")
	}

	// Check if the carry flag is set correctly (no borrow)
	if cpu.getSRBit(0) != 1 {
		t.Errorf("CMP Absolute Y failed: Carry flag not set correctly")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for CMP Absolute
		t.Errorf("CMP Absolute Y failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestCMPIndirectX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// CMP ($10,X)
	cpu.writeMemory(cpu.PC, CMP_INDIRECT_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero Page address to compare
	cpu.writeMemory(0x0010, 0x00)   // Value at Zero Page address to compare
	cpu.writeMemory(0x0011, 0x10)   // Value at Zero Page address to compare
	cpu.writeMemory(0x1000, 0x10)   // Value at Zero Page address to compare
	cpu.X = 0x01                    // Set X register for Zero Page addressing
	cpu.A = 0x20                    // Set accumulator for comparison

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the zero flag is set correctly
	if cpu.getSRBit(1) != 0 {
		t.Errorf("CMP Indirect X failed: Zero flag not set correctly")
	}

	// Check if the negative flag is set correctly
	if cpu.getSRBit(7) != 0 {
		t.Errorf("CMP Indirect X failed: Negative flag not set correctly")
	}

	// Check if the carry flag is set correctly (no borrow)
	if cpu.getSRBit(0) != 1 {
		t.Errorf("CMP Indirect X failed: Carry flag not set correctly")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for CMP Zero Page
		t.Errorf("CMP Indirect X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestCMPIndirectY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// CMP ($10),Y
	cpu.writeMemory(cpu.PC, CMP_INDIRECT_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero Page address to compare
	cpu.writeMemory(0x0010, 0x00)   // Value at Zero Page address to compare
	cpu.writeMemory(0x0011, 0x10)   // Value at Zero Page address to compare
	cpu.writeMemory(0x1001, 0x10)   // Value at Zero Page address to compare
	cpu.Y = 0x01                    // Set Y register for Zero Page addressing
	cpu.A = 0x20                    // Set accumulator for comparison

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the zero flag is set correctly
	if cpu.getSRBit(1) != 0 {
		t.Errorf("CMP Indirect Y failed: Zero flag not set correctly")
	}

	// Check if the negative flag is set correctly
	if cpu.getSRBit(7) != 0 {
		t.Errorf("CMP Indirect Y failed: Negative flag not set correctly")
	}

	// Check if the carry flag is set correctly (no borrow)
	if cpu.getSRBit(0) != 1 {
		t.Errorf("CMP Indirect Y failed: Carry flag not set correctly")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for CMP Zero Page
		t.Errorf("CMP Indirect Y failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}

func TestJMPAbsolute(t *testing.T) {
	var cpu CPU
	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, JMP_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x34)
	cpu.writeMemory(cpu.PC+2, 0x12)

	cpu.cpuQuit = true
	cpu.startCPU()

	// Check if the program counter is set correctly
	if cpu.PC != 0x1234 {
		t.Errorf("JMP Absolute failed: expected PC = %04X, got %04X", 0x1234, cpu.PC)
	}
}
func TestJMPIndirect(t *testing.T) {
	var cpu CPU
	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, JMP_INDIRECT_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x34)
	cpu.writeMemory(cpu.PC+2, 0x12)
	cpu.writeMemory(0x1234, 0x56)
	cpu.writeMemory(0x1235, 0x78)

	cpu.cpuQuit = true
	cpu.startCPU()

	// Check if the program counter is set correctly
	if cpu.PC != 0x7856 {
		t.Errorf("JMP Indirect failed: expected PC = %04X, got %04X", 0x7856, cpu.PC)
	}
}

func TestANDImmediate(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// AND #10
	cpu.writeMemory(cpu.PC, AND_IMMEDIATE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if A has the expected value
	if cpu.A != 0x00 {
		t.Errorf("AND Immediate failed: got %02X, want %02X", cpu.A, 0x00)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for AND immediate
		t.Errorf("AND Immediate failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestANDZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// AND $10
	cpu.writeMemory(cpu.PC, AND_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x20)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("AND Zero Page failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for AND immediate
		t.Errorf("AND Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestANDZeroPageX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// AND $10,X
	cpu.writeMemory(cpu.PC, AND_ZERO_PAGE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0011, 0x20)
	cpu.X = 0x01
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("AND Zero Page X failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for AND immediate
		t.Errorf("AND Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestANDAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// AND $1000
	cpu.writeMemory(cpu.PC, AND_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1000, 0x20)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("AND Absolute failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for AND immediate
		t.Errorf("AND Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestANDAbsoluteX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// AND $1000,X
	cpu.writeMemory(cpu.PC, AND_ABSOLUTE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1001, 0x20)
	cpu.X = 0x01
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("AND Absolute X failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { //
	}
}
func TestANDAbsoluteY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// AND $1000,Y
	cpu.writeMemory(cpu.PC, AND_ABSOLUTE_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1001, 0x20)
	cpu.Y = 0x01
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("AND Absolute Y failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for AND immediate
		t.Errorf("AND Absolute Y failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestANDIndirectX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// AND ($10,X)
	cpu.writeMemory(cpu.PC, AND_INDIRECT_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)

	// Adjusted memory setup for effective address calculation
	cpu.writeMemory(0x0011, 0x00) // Low byte of effective address
	cpu.writeMemory(0x0012, 0x10) // High byte of effective address (was missing in your setup)

	cpu.writeMemory(0x1000, 0x20) // Value to AND with
	cpu.A = 0x20
	cpu.X = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("AND Indirect X failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for AND immediate
		t.Errorf("AND Indirect X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestANDIndirectY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// AND ($10),Y
	cpu.writeMemory(cpu.PC, AND_INDIRECT_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x00)
	cpu.writeMemory(0x0011, 0x10)
	cpu.writeMemory(0x1001, 0x20)
	cpu.A = 0x20
	cpu.Y = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("AND Indirect Y failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for AND immediate
		t.Errorf("AND Indirect Y failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestEORImmediate(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// EOR #10
	cpu.writeMemory(cpu.PC, EOR_IMMEDIATE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if A has the expected value
	if cpu.A != (0x20 ^ 0x10) {
		t.Errorf("EOR Immediate failed: got %02X, want %02X", cpu.A, 0x20^0x10)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for EOR immediate
		t.Errorf("EOR Immediate failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestEORZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// EOR $10
	cpu.writeMemory(cpu.PC, EOR_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x20)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x00 {
		t.Errorf("EOR Zero Page failed: got %02X, want %02X", cpu.A, 0x00)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for EOR immediate
		t.Errorf("EOR Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestEORZeroPageX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// EOR $10,X
	cpu.writeMemory(cpu.PC, EOR_ZERO_PAGE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0011, 0x20)
	cpu.X = 0x01
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x00 {
		t.Errorf("EOR Zero Page X failed: got %02X, want %02X", cpu.A, 0x00)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for EOR immediate
		t.Errorf("EOR Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestEORAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// EOR $1000
	cpu.writeMemory(cpu.PC, EOR_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1000, 0x20)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x00 {
		t.Errorf("EOR Absolute failed: got %02X, want %02X", cpu.A, 0x00)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for EOR immediate
		t.Errorf("EOR Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestEORAbsoluteX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// EOR $1000,X
	cpu.writeMemory(cpu.PC, EOR_ABSOLUTE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1001, 0x20)
	cpu.X = 0x01
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x00 {
		t.Errorf("EOR Absolute X failed: got %02X, want %02X", cpu.A, 0x00)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for EOR immediate
		t.Errorf("EOR Absolute X failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestEORAbsoluteY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// EOR $1000,Y
	cpu.writeMemory(cpu.PC, EOR_ABSOLUTE_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1001, 0x20)
	cpu.Y = 0x01
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x00 {
		t.Errorf("EOR Absolute Y failed: got %02X, want %02X", cpu.A, 0x00)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for EOR immediate
		t.Errorf("EOR Absolute Y failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestEORIndirectX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// EOR ($10,X)
	cpu.writeMemory(cpu.PC, EOR_INDIRECT_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)

	// Corrected memory setup for effective address calculation
	cpu.writeMemory(0x0011, 0x00) // Low byte of effective address
	cpu.writeMemory(0x0012, 0x10) // High byte of effective address

	cpu.writeMemory(0x1000, 0x20) // Value to EOR with
	cpu.A = 0x20
	cpu.X = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x00 {
		t.Errorf("EOR Indirect X failed: got %02X, want %02X", cpu.A, 0x00)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for EOR immediate
		t.Errorf("EOR Indirect X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestEORIndirectY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// EOR ($10),Y
	cpu.writeMemory(cpu.PC, EOR_INDIRECT_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x00)
	cpu.writeMemory(0x0011, 0x10)
	cpu.writeMemory(0x1001, 0x20)
	cpu.A = 0x20
	cpu.Y = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x00 {
		t.Errorf("EOR Indirect Y failed: got %02X, want %02X", cpu.A, 0x00)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for EOR immediate
		t.Errorf("EOR Indirect Y failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestORAImmediate(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ORA #10
	cpu.writeMemory(cpu.PC, ORA_IMMEDIATE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x30 {
		t.Errorf("ORA Immediate failed: got %02X, want %02X", cpu.A, 0x30)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for ORA immediate
		t.Errorf("ORA Immediate failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestORAZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// ORA $10
	cpu.writeMemory(cpu.PC, ORA_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x20)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("ORA Zero Page failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for ORA immediate
		t.Errorf("ORA Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestORAZeroPageX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// ORA $10,X
	cpu.writeMemory(cpu.PC, ORA_ZERO_PAGE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0011, 0x20)
	cpu.X = 0x01
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("ORA Zero Page X failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for ORA immediate
		t.Errorf("ORA Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestORAAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// ORA $1000
	cpu.writeMemory(cpu.PC, ORA_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1000, 0x20)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("ORA Absolute failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for ORA immediate
		t.Errorf("ORA Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestORAAbsoluteX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// ORA $1000,X
	cpu.writeMemory(cpu.PC, ORA_ABSOLUTE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1001, 0x20)
	cpu.X = 0x01
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("ORA Absolute X failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for ORA immediate
		t.Errorf("ORA Absolute X failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestORAAbsoluteY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// ORA $1000,Y
	cpu.writeMemory(cpu.PC, ORA_ABSOLUTE_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1001, 0x20)
	cpu.Y = 0x01
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("ORA Absolute Y failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for ORA immediate
		t.Errorf("ORA Absolute Y failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestORAIndirectX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// ORA ($10,X)
	cpu.writeMemory(cpu.PC, ORA_INDIRECT_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x00)
	cpu.writeMemory(0x0011, 0x10)
	cpu.writeMemory(0x1000, 0x20)
	cpu.A = 0x20
	cpu.X = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("ORA Indirect X failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes
	}
}
func TestORAIndirectY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// ORA ($10),Y
	cpu.writeMemory(cpu.PC, ORA_INDIRECT_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x00)
	cpu.writeMemory(0x0011, 0x10)
	cpu.writeMemory(0x1001, 0x20)
	cpu.A = 0x20
	cpu.Y = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("ORA Indirect Y failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes
	}
}
func TestBITZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// Setup for BIT instruction (Zero Page addressing mode)
	// BIT $10
	cpu.writeMemory(cpu.PC, BIT_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Address in Zero Page
	testValue := byte(0xC0)         // Test value at Zero Page address (0xC0 chosen to set negative and overflow flags)
	cpu.writeMemory(0x0010, testValue)
	cpu.A = 0x80 // Set accumulator to a value that will affect the zero flag

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if Zero Flag is set correctly
	if cpu.getSRBit(1) != 0 { // Zero flag is bit 1 of the status register
		t.Errorf("BIT Zero Page failed: Zero Flag should be set")
	}

	// Check if Negative Flag is set correctly
	if cpu.getSRBit(7) != (testValue>>7)&1 {
		t.Errorf("BIT Zero Page failed: Negative Flag should match bit 7 of test value")
	}

	// Check if Overflow Flag is set correctly
	if cpu.getSRBit(6) != (testValue>>6)&1 {
		t.Errorf("BIT Zero Page failed: Overflow Flag should match bit 6 of test value")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for BIT Zero Page
		t.Errorf("BIT Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestBITAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// Setup for BIT instruction (Absolute addressing mode)
	// BIT $1000
	cpu.writeMemory(cpu.PC, BIT_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10) // Address in Zero Page
	testValue := byte(0xC0)         // Test value at Zero Page address (0xC0 chosen to set negative and overflow flags)
	cpu.writeMemory(0x1000, testValue)
	cpu.A = 0x80 // Set accumulator to a value that will affect the zero flag

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if Zero Flag is set correctly
	if cpu.getSRBit(1) != 0 { // Zero flag is bit 1 of the status register
		t.Errorf("BIT Absolute failed: Zero Flag should be set")
	}

	// Check if Negative Flag is set correctly
	if cpu.getSRBit(7) != (testValue>>7)&1 {
		t.Errorf("BIT Absolute failed: Negative Flag should match bit 7 of test value")
	}

	// Check if Overflow Flag is set correctly
	if cpu.getSRBit(6) != (testValue>>6)&1 {
		t.Errorf("BIT Absolute failed: Overflow Flag should match bit 6 of test value")
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for BIT Absolute
		t.Errorf("BIT Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}

func TestINCZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// INC $10 (Zero Page Addressing)
	cpu.writeMemory(cpu.PC, INC_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero page address
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x10, initialMemoryValue) // Set initial value at zero page address

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue + 1 // Expected value after increment

	// Check if memory at zero page address has the expected incremented value
	if cpu.readMemory(0x10) != expectedValue {
		t.Errorf("INC Zero Page failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x10))
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("INC Zero Page failed: Negative flag not set correctly")
		}
	} else {
		if cpu.getSRBit(7) != 0 {
			t.Errorf("INC Zero Page failed: Negative flag should not be set")
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("INC Zero Page failed: Zero flag not set correctly")
		}
	} else {
		if cpu.getSRBit(1) != 0 {
			t.Errorf("INC Zero Page failed: Zero flag should not be set")
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for INC Zero Page
		t.Errorf("INC Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestINCZeroPageX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// INC $10,X (Zero Page Addressing)
	cpu.writeMemory(cpu.PC, INC_ZERO_PAGE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero page address
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x11, initialMemoryValue) // Set initial value at zero page address
	cpu.X = 0x01                              // Set X register for Zero Page addressing

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue + 1 // Expected value after increment

	// Check if memory at zero page address has the expected incremented value
	if cpu.readMemory(0x11) != expectedValue {
		t.Errorf("INC Zero Page X failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x11))
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("INC Zero Page X failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("INC Zero Page X failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("INC Zero Page X failed: Zero flag not set correctly")
		}
	} else {
		if cpu.getSRBit(1) != 0 {
			t.Errorf("INC Zero Page X failed: Zero flag should not be set")
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for INC Zero Page X
		t.Errorf("INC Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestINCAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// INC $1000 (Absolute Addressing)
	cpu.writeMemory(cpu.PC, INC_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address
	cpu.writeMemory(cpu.PC+2, 0x10)
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x1000, initialMemoryValue) // Set initial value at absolute address

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue + 1 // Expected value after increment

	// Check if memory at absolute address has the expected incremented value
	if cpu.readMemory(0x1000) != expectedValue {
		t.Errorf("INC Absolute failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x1000))
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("INC Absolute failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("INC Absolute failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("INC Absolute failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("INC Absolute failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for INC Absolute
		t.Errorf("INC Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestINCAbsoluteX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// INC $1000,X (Absolute Addressing)
	cpu.writeMemory(cpu.PC, INC_ABSOLUTE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address
	cpu.writeMemory(cpu.PC+2, 0x10)
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x1001, initialMemoryValue) // Set initial value at absolute address
	cpu.X = 0x01                                // Set X register for Absolute addressing

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue + 1 // Expected value after increment

	// Check if memory at absolute address has the expected incremented value
	if cpu.readMemory(0x1001) != expectedValue {
		t.Errorf("INC Absolute X failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x1001))
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("INC Absolute X failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("INC Absolute X failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("INC Absolute X failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("INC Absolute X failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for INC Absolute X
		t.Errorf("INC Absolute X failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}

func TestDECZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// DEC $10 (Zero Page Addressing)
	cpu.writeMemory(cpu.PC, DEC_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero page address
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x10, initialMemoryValue) // Set initial value at zero page address

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue - 1 // Expected value after decrement

	// Check if memory at zero page address has the expected decremented value
	if cpu.readMemory(0x10) != expectedValue {
		t.Errorf("DEC Zero Page failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x10))
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("DEC Zero Page failed: Negative flag not set correctly")
		}
	} else {
		if cpu.getSRBit(7) != 0 {
			t.Errorf("DEC Zero Page failed: Negative flag should not be set")
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("DEC Zero Page failed: Zero flag not set correctly")
		}
	} else {
		if cpu.getSRBit(1) != 0 {
			t.Errorf("DEC Zero Page failed: Zero flag should not be set")
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for DEC Zero Page
		t.Errorf("DEC Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestDECZeroPageX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// DEC $10,X (Zero Page Addressing)
	cpu.writeMemory(cpu.PC, DEC_ZERO_PAGE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero page address
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x11, initialMemoryValue) // Set initial value at zero page address
	cpu.X = 0x01                              // Set X register for Zero Page addressing

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue - 1 // Expected value after decrement

	// Check if memory at zero page address has the expected decremented value
	if cpu.readMemory(0x11) != expectedValue {
		t.Errorf("DEC Zero Page X failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x11))
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("DEC Zero Page X failed: Negative flag not set correctly")
		}
	} else {
		if cpu.getSRBit(7) != 0 {
			t.Errorf("DEC Zero Page X failed: Negative flag should not be set")
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("DEC Zero Page X failed: Zero flag not set correctly")
		}
	} else {
		if cpu.getSRBit(1) != 0 {
			t.Errorf("DEC Zero Page X failed: Zero flag should not be set")
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for DEC Zero Page X
		t.Errorf("DEC Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestDECAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// DEC $1000 (Absolute Addressing)
	cpu.writeMemory(cpu.PC, DEC_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address
	cpu.writeMemory(cpu.PC+2, 0x10)
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x1000, initialMemoryValue) // Set initial value at absolute address

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue - 1 // Expected value after decrement

	// Check if memory at absolute address has the expected decremented value
	if cpu.readMemory(0x1000) != expectedValue {
		t.Errorf("DEC Absolute failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x1000))
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("DEC Absolute failed: Negative flag not set correctly")
		}
	} else {
		if cpu.getSRBit(7) != 0 {
			t.Errorf("DEC Absolute failed: Negative flag should not be set")
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("DEC Absolute failed: Zero flag not set correctly")
		}
	} else {
		if cpu.getSRBit(1) != 0 {
			t.Errorf("DEC Absolute failed: Zero flag should not be set")
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for DEC Absolute
		t.Errorf("DEC Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestDECAbsoluteX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// DEC $1000,X (Absolute Addressing)
	cpu.writeMemory(cpu.PC, DEC_ABSOLUTE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address
	cpu.writeMemory(cpu.PC+2, 0x10)
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x1001, initialMemoryValue) // Set initial value at absolute address
	cpu.X = 0x01                                // Set X register for Absolute addressing

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue - 1 // Expected value after decrement

	// Check if memory at absolute address has the expected decremented value
	if cpu.readMemory(0x1001) != expectedValue {
		t.Errorf("DEC Absolute X failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x1001))
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("DEC Absolute X failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("DEC Absolute X failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("DEC Absolute X failed: Zero flag not set correctly")
		}
	} else {
		if cpu.getSRBit(1) != 0 {
			t.Errorf("DEC Absolute X failed: Zero flag should not be set")
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for DEC Absolute X
		t.Errorf("DEC Absolute X failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}

func TestADCImmediate(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ADC #10
	cpu.writeMemory(cpu.PC, ADC_IMMEDIATE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x30 {
		t.Errorf("ADC Immediate failed: got %02X, want %02X", cpu.A, 0x30)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for ADC immediate
		t.Errorf("ADC Immediate failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestADCZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ADC $10
	cpu.writeMemory(cpu.PC, ADC_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x10)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x30 {
		t.Errorf("ADC Zero Page failed: got %02X, want %02X", cpu.A, 0x30)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for ADC immediate
		t.Errorf("ADC Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestADCZeroPageX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ADC $10,X
	cpu.writeMemory(cpu.PC, ADC_ZERO_PAGE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0011, 0x10)
	cpu.X = 0x01
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x30 {
		t.Errorf("ADC Zero Page X failed: got %02X, want %02X", cpu.A, 0x30)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for ADC immediate
		t.Errorf("ADC Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestADCAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ADC $1000
	cpu.writeMemory(cpu.PC, ADC_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1000, 0x10)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value
	if cpu.A != 0x30 {
		t.Errorf("ADC Absolute failed: got %02X, want %02X", cpu.A, 0x30)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for ADC immediate
		t.Errorf("ADC Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestADCAbsoluteX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ADC $1000,X
	cpu.writeMemory(cpu.PC, ADC_ABSOLUTE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1001, 0x10)
	cpu.X = 0x01
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state
	// Check if A has the expected value

	if cpu.A != 0x30 {
		t.Errorf("ADC Absolute X failed: got %02X, want %02X", cpu.A, 0x30)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for ADC immediate
		t.Errorf("ADC Absolute X failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestADCAbsoluteY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ADC $1000,Y
	cpu.writeMemory(cpu.PC, ADC_ABSOLUTE_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00)
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1001, 0x10)
	cpu.Y = 0x01
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if A has the expected value
	if cpu.A != 0x30 {
		t.Errorf("ADC Absolute Y failed: got %02X, want %02X", cpu.A, 0x30)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for ADC immediate
		t.Errorf("ADC Absolute Y failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestADCIndirectX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ADC ($10,X)
	cpu.writeMemory(cpu.PC, ADC_INDIRECT_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)

	// Corrected memory setup for effective address calculation
	cpu.writeMemory(0x0011, 0x00) // Low byte of effective address
	cpu.writeMemory(0x0012, 0x10) // High byte of effective address

	cpu.writeMemory(0x1000, 0x10) // Value to add
	cpu.A = 0x20
	cpu.X = 0x01

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if A has the expected value
	if cpu.A != 0x30 {
		t.Errorf("ADC Indirect X failed: got %02X, want %02X", cpu.A, 0x30)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for ADC immediate
		t.Errorf("ADC Indirect X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestADCIndirectY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ADC ($10),Y
	cpu.writeMemory(cpu.PC, ADC_INDIRECT_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10)
	cpu.writeMemory(0x0010, 0x00)
	cpu.writeMemory(0x0011, 0x10)
	cpu.writeMemory(0x1001, 0x10)
	cpu.A = 0x20
	cpu.Y = 0x01
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if A has the expected value
	if cpu.A != 0x30 {
		t.Errorf("ADC Indirect Y failed: got %02X, want %02X", cpu.A, 0x30)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for ADC immediate
		t.Errorf("ADC Indirect Y failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}

func TestSBCImmediate(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// SBC #10
	cpu.writeMemory(cpu.PC, SBC_IMMEDIATE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Immediate value to subtract
	cpu.A = 0x20                    // Set accumulator for subtraction
	cpu.setCarryFlag()              // Ensure the carry flag is set before subtraction

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := cpu.preOpA - 0x10 - (1 - cpu.getSRBit(0)) // Adjusted expected value after subtraction

	// Check if accumulator has the expected value
	if cpu.A != expectedValue {
		t.Errorf("SBC Immediate failed: got %02X, want %02X", cpu.A, expectedValue)
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 && cpu.getSRBit(1) != 1 {
		t.Errorf("SBC Immediate failed: Zero flag not set correctly")
	}

	// Check if the negative flag is set correctly
	if (expectedValue&0x80) != 0 && cpu.getSRBit(7) != 1 {
		t.Errorf("SBC Immediate failed: Negative flag not set correctly")
	}

	// Check if the carry flag is set correctly (no borrow)
	if expectedValue >= 0 && cpu.getSRBit(0) != 1 {
		t.Errorf("SBC Immediate failed: Carry flag not set correctly")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for SBC Immediate
		t.Errorf("SBC Immediate failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}

func TestSBCZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// SBC $10
	cpu.writeMemory(cpu.PC, SBC_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero Page address to subtract
	cpu.writeMemory(0x0010, 0x10)   // Value at Zero Page address to subtract
	cpu.A = 0x20                    // Set accumulator for subtraction
	cpu.setCarryFlag()              // Ensure the carry flag is set before subtraction

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := cpu.preOpA - 0x10 - (1 - cpu.getSRBit(0)) // Adjusted expected value after subtraction

	// Check if accumulator has the expected value
	if cpu.A != expectedValue {
		t.Errorf("SBC Zero Page failed: got %02X, want %02X", cpu.A, expectedValue)
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 && cpu.getSRBit(1) != 1 {
		t.Errorf("SBC Zero Page failed: Zero flag not set correctly")
	}

	// Check if the negative flag is set correctly
	if (expectedValue&0x80) != 0 && cpu.getSRBit(7) != 1 {
		t.Errorf("SBC Zero Page failed: Negative flag not set correctly")
	}

	// Check if the carry flag is set correctly (no borrow)
	if expectedValue >= 0 && cpu.getSRBit(0) != 1 {
		t.Errorf("SBC Zero Page failed: Carry flag not set correctly")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for SBC Zero Page
		t.Errorf("SBC Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}

func TestSBCZeroPageX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// SBC $10,X
	cpu.writeMemory(cpu.PC, SBC_ZERO_PAGE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero Page address to subtract
	cpu.writeMemory(0x0011, 0x10)   // Value at Zero Page address to subtract
	cpu.X = 0x01                    // Set X register for Zero Page addressing
	cpu.A = 0x20                    // Set accumulator for subtraction
	cpu.setCarryFlag()              // Ensure the carry flag is set before subtraction

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := cpu.preOpA - 0x10 - (1 - cpu.getSRBit(0)) // Adjusted expected value after subtraction

	// Check if accumulator has the expected value
	if cpu.A != expectedValue {
		t.Errorf("SBC Zero Page X failed: got %02X, want %02X", cpu.A, expectedValue)
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 && cpu.getSRBit(1) != 1 {
		t.Errorf("SBC Zero Page X failed: Zero flag not set correctly")
	}

	// Check if the negative flag is set correctly
	if (expectedValue&0x80) != 0 && cpu.getSRBit(7) != 1 {
		t.Errorf("SBC Zero Page X failed: Negative flag not set correctly")
	}

	// Check if the carry flag is set correctly (no borrow)
	if expectedValue >= 0 && cpu.getSRBit(0) != 1 {
		t.Errorf("SBC Zero Page X failed: Carry flag not set correctly")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for SBC Zero Page
		t.Errorf("SBC Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}

func TestSBCAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// SBC $1000
	cpu.writeMemory(cpu.PC, SBC_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address to subtract
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1000, 0x10) // Value at Absolute address to subtract
	cpu.A = 0x20                  // Set accumulator for subtraction
	cpu.setCarryFlag()            // Ensure the carry flag is set before subtraction

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := cpu.preOpA - 0x10 - (1 - cpu.getSRBit(0)) // Adjusted expected value after subtraction

	// Check if accumulator has the expected value
	if cpu.A != expectedValue {
		t.Errorf("SBC Absolute failed: got %02X, want %02X", cpu.A, expectedValue)
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 && cpu.getSRBit(1) != 1 {
		t.Errorf("SBC Absolute failed: Zero flag not set correctly")
	}

	// Check if the negative flag is set correctly
	if (expectedValue&0x80) != 0 && cpu.getSRBit(7) != 1 {
		t.Errorf("SBC Absolute failed: Negative flag not set correctly")
	}

	// Check if the carry flag is set correctly (no borrow)
	if expectedValue >= 0 && cpu.getSRBit(0) != 1 {
		t.Errorf("SBC Absolute failed: Carry flag not set correctly")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for SBC Absolute
		t.Errorf("SBC Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}

func TestSBCAbsoluteX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// SBC $1000,X
	cpu.writeMemory(cpu.PC, SBC_ABSOLUTE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address to subtract
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1001, 0x10) // Value at Absolute address to subtract
	cpu.X = 0x01                  // Set X register for Absolute addressing
	cpu.A = 0x20                  // Set accumulator for subtraction
	cpu.setCarryFlag()            // Ensure the carry flag is set before subtraction

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := cpu.preOpA - 0x10 - (1 - cpu.getSRBit(0)) // Adjusted expected value after subtraction

	// Check if accumulator has the expected value
	if cpu.A != expectedValue {
		t.Errorf("SBC Absolute X failed: got %02X, want %02X", cpu.A, expectedValue)
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 && cpu.getSRBit(1) != 1 {
		t.Errorf("SBC Absolute X failed: Zero flag not set correctly")
	}

	// Check if the negative flag is set correctly
	if (expectedValue&0x80) != 0 && cpu.getSRBit(7) != 1 {
		t.Errorf("SBC Absolute X failed: Negative flag not set correctly")
	}

	// Check if the carry flag is set correctly (no borrow)
	if expectedValue >= 0 && cpu.getSRBit(0) != 1 {
		t.Errorf("SBC Absolute X failed: Carry flag not set correctly")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for SBC Absolute
		t.Errorf("SBC Absolute X failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}

func TestSBCAbsoluteY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// SBC $1000,Y
	cpu.writeMemory(cpu.PC, SBC_ABSOLUTE_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address to subtract
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1001, 0x10) // Value at Absolute address to subtract
	cpu.Y = 0x01                  // Set Y register for Absolute addressing
	cpu.A = 0x20                  // Set accumulator for subtraction
	cpu.setCarryFlag()            // Ensure the carry flag is set before subtraction

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := cpu.preOpA - 0x10 - (1 - cpu.getSRBit(0)) // Adjusted expected value after subtraction

	// Check if accumulator has the expected value
	if cpu.A != expectedValue {
		t.Errorf("SBC Absolute Y failed: got %02X, want %02X", cpu.A, expectedValue)
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 && cpu.getSRBit(1) != 1 {
		t.Errorf("SBC Absolute Y failed: Zero flag not set correctly")
	}

	// Check if the negative flag is set correctly
	if (expectedValue&0x80) != 0 && cpu.getSRBit(7) != 1 {
		t.Errorf("SBC Absolute Y failed: Negative flag not set correctly")
	}

	// Check if the carry flag is set correctly (no borrow)
	if expectedValue >= 0 && cpu.getSRBit(0) != 1 {
		t.Errorf("SBC Absolute Y failed: Carry flag not set correctly")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for SBC Absolute
		t.Errorf("SBC Absolute Y failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}

func TestSBCIndirectX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// SBC ($10,X)
	cpu.writeMemory(cpu.PC, SBC_INDIRECT_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero Page address to subtract

	// Set up the Indirect X addressing
	cpu.writeMemory(0x0011, 0x00) // Low byte of effective address at $11 ($10 + X)
	cpu.writeMemory(0x0012, 0x10) // High byte of effective address at $12 ($11 + 1)

	// Set up value at the effective address $1000
	cpu.writeMemory(0x1000, 0x10) // Value at $1000 to subtract

	cpu.X = 0x01       // Set X register for Zero Page addressing
	cpu.A = 0x20       // Set accumulator for subtraction
	cpu.setCarryFlag() // Ensure the carry flag is set before subtraction

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := cpu.preOpA - 0x10 - (1 - cpu.getSRBit(0)) // Adjusted expected value after subtraction

	// Check if accumulator has the expected value
	if cpu.A != expectedValue {
		t.Errorf("SBC Indirect X failed: got %02X, want %02X", cpu.A, expectedValue)
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 && cpu.getSRBit(1) != 1 {
		t.Errorf("SBC Indirect X failed: Zero flag not set correctly")
	}

	// Check if the negative flag is set correctly
	if (expectedValue&0x80) != 0 && cpu.getSRBit(7) != 1 {
		t.Errorf("SBC Indirect X failed: Negative flag not set correctly")
	}

	// Check if the carry flag is set correctly (no borrow)
	if expectedValue >= 0 && cpu.getSRBit(0) != 1 {
		t.Errorf("SBC Indirect X failed: Carry flag not set correctly")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for SBC Zero Page
		t.Errorf("SBC Indirect X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}

func TestSBCIndirectY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// SBC ($10),Y
	cpu.writeMemory(cpu.PC, SBC_INDIRECT_Y_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero Page address to subtract
	cpu.writeMemory(0x0010, 0x00)   // Value at Zero Page address to subtract
	cpu.writeMemory(0x0011, 0x10)   // Value at Zero Page address to subtract
	cpu.writeMemory(0x1001, 0x10)   // Value at Zero Page address to subtract
	cpu.Y = 0x01                    // Set Y register for Zero Page addressing
	cpu.A = 0x20                    // Set accumulator for subtraction
	cpu.setCarryFlag()              // Ensure the carry flag is set before subtraction

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := cpu.preOpA - 0x10 - (1 - cpu.getSRBit(0)) // Adjusted expected value after subtraction

	// Check if accumulator has the expected value
	if cpu.A != expectedValue {
		t.Errorf("SBC Indirect Y failed: got %02X, want %02X", cpu.A, expectedValue)
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 && cpu.getSRBit(1) != 1 {
		t.Errorf("SBC Indirect Y failed: Zero flag not set correctly")
	}

	// Check if the negative flag is set correctly
	if (expectedValue&0x80) != 0 && cpu.getSRBit(7) != 1 {
		t.Errorf("SBC Indirect Y failed: Negative flag should be set but is not")
	} else if (expectedValue&0x80) == 0 && cpu.getSRBit(7) != 0 {
		t.Errorf("SBC Indirect Y failed: Negative flag should not be set but is")
	}

	// Check if the carry flag is set correctly (no borrow)
	if expectedValue >= 0 && cpu.getSRBit(0) != 1 {
		t.Errorf("SBC Indirect Y failed: Carry flag not set correctly")
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for SBC Zero Page
		t.Errorf("SBC Indirect Y failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}

func TestRORAccumulator(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// ROR A (Accumulator Addressing)
	cpu.writeMemory(cpu.PC, ROR_ACCUMULATOR_OPCODE)
	initialA := byte(0x20)
	cpu.A = initialA // Set initial value of A register

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialA >> 1 // Expected value after shift right
	if cpu.getSRBit(0) == 1 {      // Check if carry flag is set
		expectedValue |= 0x80 // Set bit 7 if carry flag is set
	}

	// Check if A register has the expected shifted value
	if cpu.A != expectedValue {
		t.Errorf("ROR Accumulator failed: expected A = %02X, got %02X", expectedValue, cpu.A)
	}

	// Check if the carry flag is set correctly
	if initialA&0x01 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("ROR Accumulator failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("ROR Accumulator failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("ROR Accumulator failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("ROR Accumulator failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("ROR Accumulator failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("ROR Accumulator failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for ROR Accumulator
		t.Errorf("ROR Accumulator failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}
func TestRORZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// ROR $10 (Zero Page Addressing)
	cpu.writeMemory(cpu.PC, ROR_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero page address
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x10, initialMemoryValue) // Set initial value at zero page address

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue >> 1 // Expected value after shift right
	if cpu.getSRBit(0) == 1 {                // Check if carry flag is set
		expectedValue |= 0x80 // Set bit 7 if carry flag is set
	}

	// Check if memory at zero page address has the expected shifted value
	if cpu.readMemory(0x10) != expectedValue {
		t.Errorf("ROR Zero Page failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x10))
	}

	// Check if the carry flag is set correctly
	if initialMemoryValue&0x01 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("ROR Zero Page failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("ROR Zero Page failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("ROR Zero Page failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("ROR Zero Page failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("ROR Zero Page failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("ROR Zero Page failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for ROR Zero Page
		t.Errorf("ROR Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestRORZeroPageX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ROR $10,X (Zero Page Addressing)
	cpu.writeMemory(cpu.PC, ROR_ZERO_PAGE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero page address
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x11, initialMemoryValue) // Set initial value at zero page address
	cpu.X = 0x01                              // Set X register for Zero Page addressing

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue >> 1 // Expected value after shift right
	if cpu.getSRBit(0) == 1 {                // Check if carry flag is set
		expectedValue |= 0x80 // Set bit 7 if carry flag is set
	}

	// Check if memory at zero page address has the expected shifted value
	if cpu.readMemory(0x11) != expectedValue {
		t.Errorf("ROR Zero Page X failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x11))
	}

	// Check if the carry flag is set correctly
	if initialMemoryValue&0x01 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("ROR Zero Page X failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("ROR Zero Page X failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("ROR Zero Page X failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("ROR Zero Page X failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("ROR Zero Page X failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("ROR Zero Page X failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for ROR Zero Page X
		t.Errorf("ROR Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestRORAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ROR $1000 (Absolute Addressing)
	cpu.writeMemory(cpu.PC, ROR_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address
	cpu.writeMemory(cpu.PC+2, 0x10)
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x1000, initialMemoryValue) // Set initial value at absolute address

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue >> 1 // Expected value after shift right
	if cpu.getSRBit(0) == 1 {                // Check if carry flag is set
		expectedValue |= 0x80 // Set bit 7 if carry flag is set
	}

	// Check if memory at absolute address has the expected shifted value
	if cpu.readMemory(0x1000) != expectedValue {
		t.Errorf("ROR Absolute failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x1000))
	}

	// Check if the carry flag is set correctly
	if initialMemoryValue&0x01 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("ROR Absolute failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("ROR Absolute failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("ROR Absolute failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("ROR Absolute failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("ROR Absolute failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("ROR Absolute failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for ROR Absolute
		t.Errorf("ROR Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestRORAbsoluteX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ROR $1000,X (Absolute Addressing)
	cpu.writeMemory(cpu.PC, ROR_ABSOLUTE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address
	cpu.writeMemory(cpu.PC+2, 0x10)
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x1001, initialMemoryValue) // Set initial value at absolute address
	cpu.X = 0x01                                // Set X register for Absolute addressing

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue >> 1 // Expected value after shift right
	if cpu.getSRBit(0) == 1 {                // Check if carry flag is set
		expectedValue |= 0x80 // Set bit 7 if carry flag is set
	}

	// Check if memory at absolute address has the expected shifted value
	if cpu.readMemory(0x1001) != expectedValue {
		t.Errorf("ROR Absolute X failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x1001))
	}

	// Check if the carry flag is set correctly
	if initialMemoryValue&0x01 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("ROR Absolute X failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("ROR Absolute X failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("ROR Absolute X failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("ROR Absolute X failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("ROR Absolute X failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("ROR Absolute X failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for ROR Absolute X
		t.Errorf("ROR Absolute X failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}

func TestROLAccumulator(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// ROL A (Accumulator Addressing)
	cpu.writeMemory(cpu.PC, ROL_ACCUMULATOR_OPCODE)
	initialA := byte(0x20)
	cpu.A = initialA // Set initial value of A register

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialA << 1 // Expected value after shift left
	if cpu.getSRBit(0) == 1 {      // Check if carry flag is set
		expectedValue |= 0x01 // Set bit 0 if carry flag is set
	}

	// Check if A register has the expected shifted value
	if cpu.A != expectedValue {
		t.Errorf("ROL Accumulator failed: expected A = %02X, got %02X", expectedValue, cpu.A)
	}

	// Check if the carry flag is set correctly
	if initialA&0x80 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("ROL Accumulator failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("ROL Accumulator failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("ROL Accumulator failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("ROL Accumulator failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("ROL Accumulator failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("ROL Accumulator failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for ROL Accumulator
		t.Errorf("ROL Accumulator failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}
func TestROLZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// ROL $10 (Zero Page Addressing)
	cpu.writeMemory(cpu.PC, ROL_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero page address
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x10, initialMemoryValue) // Set initial value at zero page address

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue << 1 // Expected value after shift left
	if cpu.getSRBit(0) == 1 {                // Check if carry flag is set
		expectedValue |= 0x01 // Set bit 0 if carry flag is set
	}

	// Check if memory at zero page address has the expected shifted value
	if cpu.readMemory(0x10) != expectedValue {
		t.Errorf("ROL Zero Page failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x10))
	}

	// Check if the carry flag is set correctly
	if initialMemoryValue&0x80 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("ROL Zero Page failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("ROL Zero Page failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("ROL Zero Page failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("ROL Zero Page failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("ROL Zero Page failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("ROL Zero Page failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for ROL Zero Page
		t.Errorf("ROL Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestROLZeroPageX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// ROL $10,X (Zero Page Addressing)
	cpu.writeMemory(cpu.PC, ROL_ZERO_PAGE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero page address
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x11, initialMemoryValue) // Set initial value at zero page address
	cpu.X = 0x01                              // Set X register for Zero Page addressing

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue << 1 // Expected value after shift left
	if cpu.getSRBit(0) == 1 {                // Check if carry flag is set
		expectedValue |= 0x01 // Set bit 0 if carry flag is set
	}

	// Check if memory at zero page address has the expected shifted value
	if cpu.readMemory(0x11) != expectedValue {
		t.Errorf("ROL Zero Page X failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x11))
	}

	// Check if the carry flag is set correctly
	if initialMemoryValue&0x80 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("ROL Zero Page X failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("ROL Zero Page X failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("ROL Zero Page X failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("ROL Zero Page X failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("ROL Zero Page X failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("ROL Zero Page X failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for ROL Zero Page X
		t.Errorf("ROL Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestROLAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// ROL $1000 (Absolute Addressing)
	cpu.writeMemory(cpu.PC, ROL_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address
	cpu.writeMemory(cpu.PC+2, 0x10)
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x1000, initialMemoryValue) // Set initial value at absolute address

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue << 1 // Expected value after shift left
	if cpu.getSRBit(0) == 1 {                // Check if carry flag is set
		expectedValue |= 0x01 // Set bit 0 if carry flag is set
	}

	// Check if memory at absolute address has the expected shifted value
	if cpu.readMemory(0x1000) != expectedValue {
		t.Errorf("ROL Absolute failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x1000))
	}

	// Check if the carry flag is set correctly
	if initialMemoryValue&0x80 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("ROL Absolute failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("ROL Absolute failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("ROL Absolute failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("ROL Absolute failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("ROL Absolute failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("ROL Absolute failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for ROL Absolute
		t.Errorf("ROL Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestROLAbsoluteX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// ROL $1000,X (Absolute Addressing)
	cpu.writeMemory(cpu.PC, ROL_ABSOLUTE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address
	cpu.writeMemory(cpu.PC+2, 0x10)
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x1001, initialMemoryValue) // Set initial value at absolute address
	cpu.X = 0x01                                // Set X register for Absolute addressing

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue << 1 // Expected value after shift left
	if cpu.getSRBit(0) == 1 {                // Check if carry flag is set
		expectedValue |= 0x01 // Set bit 0 if carry flag is set
	}

	// Check if memory at absolute address has the expected shifted value
	if cpu.readMemory(0x1001) != expectedValue {
		t.Errorf("ROL Absolute X failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x1001))
	}

	// Check if the carry flag is set correctly
	if initialMemoryValue&0x80 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("ROL Absolute X failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("ROL Absolute X failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("ROL Absolute X failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("ROL Absolute X failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("ROL Absolute X failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("ROL Absolute X failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for ROL Absolute X
		t.Errorf("ROL Absolute X failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}

func TestLSRAccumulator(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// LSR A (Accumulator Addressing)
	cpu.writeMemory(cpu.PC, LSR_ACCUMULATOR_OPCODE)
	initialA := byte(0x20)
	cpu.A = initialA // Set initial value of A register

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialA >> 1 // Expected value after shift right

	// Check if A register has the expected shifted value
	if cpu.A != expectedValue {
		t.Errorf("LSR Accumulator failed: expected A = %02X, got %02X", expectedValue, cpu.A)
	}

	// Check if the carry flag is set correctly
	if initialA&0x01 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("LSR Accumulator failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("LSR Accumulator failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("LSR Accumulator failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("LSR Accumulator failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("LSR Accumulator failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("LSR Accumulator failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for LSR Accumulator
		t.Errorf("LSR Accumulator failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}
func TestLSRZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// LSR $10 (Zero Page Addressing)
	cpu.writeMemory(cpu.PC, LSR_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero page address
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x10, initialMemoryValue) // Set initial value at zero page address

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue >> 1 // Expected value after shift right

	// Check if memory at zero page address has the expected shifted value
	if cpu.readMemory(0x10) != expectedValue {
		t.Errorf("LSR Zero Page failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x10))
	}

	// Check if the carry flag is set correctly
	if initialMemoryValue&0x01 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("LSR Zero Page failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("LSR Zero Page failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("LSR Zero Page failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("LSR Zero Page failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("LSR Zero Page failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("LSR Zero Page failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for LSR Zero Page
		t.Errorf("LSR Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestLSRZeroPageX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// LSR $10,X (Zero Page Addressing)
	cpu.writeMemory(cpu.PC, LSR_ZERO_PAGE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero page address
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x11, initialMemoryValue) // Set initial value at zero page address
	cpu.X = 0x01                              // Set X register for Zero Page addressing

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue >> 1 // Expected value after shift right

	// Check if memory at zero page address has the expected shifted value
	if cpu.readMemory(0x11) != expectedValue {
		t.Errorf("LSR Zero Page X failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x11))
	}

	// Check if the carry flag is set correctly
	if initialMemoryValue&0x01 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("LSR Zero Page X failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("LSR Zero Page X failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("LSR Zero Page X failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("LSR Zero Page X failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("LSR Zero Page X failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("LSR Zero Page X failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for LSR Zero Page X
		t.Errorf("LSR Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestLSRAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// LSR $1000 (Absolute Addressing)
	cpu.writeMemory(cpu.PC, LSR_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address
	cpu.writeMemory(cpu.PC+2, 0x10)
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x1000, initialMemoryValue) // Set initial value at absolute address

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue >> 1 // Expected value after shift right

	// Check if memory at absolute address has the expected shifted value
	if cpu.readMemory(0x1000) != expectedValue {
		t.Errorf("LSR Absolute failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x1000))
	}

	// Check if the carry flag is set correctly
	if initialMemoryValue&0x01 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("LSR Absolute failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("LSR Absolute failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("LSR Absolute failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("LSR Absolute failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("LSR Absolute failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("LSR Absolute failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for LSR Absolute
		t.Errorf("LSR Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}

func TestLSRAbsoluteX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// LSR $1000,X (Absolute Addressing)
	cpu.writeMemory(cpu.PC, LSR_ABSOLUTE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address
	cpu.writeMemory(cpu.PC+2, 0x10)
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x1001, initialMemoryValue) // Set initial value at absolute address
	cpu.X = 0x01                                // Set X register for Absolute addressing

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue >> 1 // Expected value after shift right

	// Check if memory at absolute address has the expected shifted value
	if cpu.readMemory(0x1001) != expectedValue {
		t.Errorf("LSR Absolute X failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x1001))
	}

	// Check if the carry flag is set correctly
	if initialMemoryValue&0x01 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("LSR Absolute X failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("LSR Absolute X failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("LSR Absolute X failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("LSR Absolute X failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("LSR Absolute X failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("LSR Absolute X failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for LSR Absolute X
		t.Errorf("LSR Absolute X failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}

func TestASLAccumulator(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ASL A (Accumulator Addressing)
	cpu.writeMemory(cpu.PC, ASL_ACCUMULATOR_OPCODE)
	initialA := byte(0x20)
	cpu.A = initialA // Set initial value of A register

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialA << 1 // Expected value after shift left

	// Check if A register has the expected shifted value
	if cpu.A != expectedValue {
		t.Errorf("ASL Accumulator failed: expected A = %02X, got %02X", expectedValue, cpu.A)
	}

	// Check if the carry flag is set correctly
	if initialA&0x80 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("ASL Accumulator failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("ASL Accumulator failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("ASL Accumulator failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("ASL Accumulator failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("ASL Accumulator failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("ASL Accumulator failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for ASL Accumulator
		t.Errorf("ASL Accumulator failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}
func TestASLZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ASL $10 (Zero Page Addressing)
	cpu.writeMemory(cpu.PC, ASL_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero page address
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x10, initialMemoryValue) // Set initial value at zero page address

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue << 1 // Expected value after shift left

	// Check if memory at zero page address has the expected shifted value
	if cpu.readMemory(0x10) != expectedValue {
		t.Errorf("ASL Zero Page failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x10))
	}

	// Check if the carry flag is set correctly
	if initialMemoryValue&0x80 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("ASL Zero Page failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("ASL Zero Page failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("ASL Zero Page failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("ASL Zero Page failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("ASL Zero Page failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("ASL Zero Page failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for ASL Zero Page
		t.Errorf("ASL Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestASLZeroPageX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ASL $10,X (Zero Page Addressing)
	cpu.writeMemory(cpu.PC, ASL_ZERO_PAGE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero page address
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x11, initialMemoryValue) // Set initial value at zero page address
	cpu.X = 0x01                              // Set X register for Zero Page addressing

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue << 1 // Expected value after shift left

	// Check if memory at zero page address has the expected shifted value
	if cpu.readMemory(0x11) != expectedValue {
		t.Errorf("ASL Zero Page X failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x11))
	}

	// Check if the carry flag is set correctly
	if initialMemoryValue&0x80 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("ASL Zero Page X failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("ASL Zero Page X failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("ASL Zero Page X failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("ASL Zero Page X failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("ASL Zero Page X failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("ASL Zero Page X failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for ASL Zero Page X
		t.Errorf("ASL Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestASLAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ASL $1000 (Absolute Addressing)
	cpu.writeMemory(cpu.PC, ASL_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address
	cpu.writeMemory(cpu.PC+2, 0x10)
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x1000, initialMemoryValue) // Set initial value at absolute address

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue << 1 // Expected value after shift left

	// Check if memory at absolute address has the expected shifted value
	if cpu.readMemory(0x1000) != expectedValue {
		t.Errorf("ASL Absolute failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x1000))
	}

	// Check if the carry flag is set correctly
	if initialMemoryValue&0x80 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("ASL Absolute failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("ASL Absolute failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("ASL Absolute failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("ASL Absolute failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("ASL Absolute failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("ASL Absolute failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for ASL Absolute
		t.Errorf("ASL Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestASLAbsoluteX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// ASL $1000,X (Absolute Addressing)
	cpu.writeMemory(cpu.PC, ASL_ABSOLUTE_X_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address
	cpu.writeMemory(cpu.PC+2, 0x10)
	initialMemoryValue := byte(0x20)
	cpu.writeMemory(0x1001, initialMemoryValue) // Set initial value at absolute address
	cpu.X = 0x01                                // Set X register for Absolute addressing

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialMemoryValue << 1 // Expected value after shift left

	// Check if memory at absolute address has the expected shifted value
	if cpu.readMemory(0x1001) != expectedValue {
		t.Errorf("ASL Absolute X failed: expected memory value = %02X, got %02X", expectedValue, cpu.readMemory(0x1001))
	}

	// Check if the carry flag is set correctly
	if initialMemoryValue&0x80 != 0 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("ASL Absolute X failed: Carry flag not set correctly")
		} else {
			if cpu.getSRBit(0) != 0 {
				t.Errorf("ASL Absolute X failed: Carry flag should not be set")
			}
		}
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("ASL Absolute X failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("ASL Absolute X failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("ASL Absolute X failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("ASL Absolute X failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for ASL Absolute X
		t.Errorf("ASL Absolute X failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}

func TestCPXImmediate(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// CPX #10
	cpu.writeMemory(cpu.PC, CPX_IMMEDIATE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Immediate value for comparison
	cpu.X = 0x20                    // Set X register for comparison

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the carry flag is set correctly (X >= immediate)
	if cpu.X >= 0x10 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("CPX Immediate failed: Carry flag not set correctly")
		}
	} else {
		if cpu.getSRBit(0) != 0 {
			t.Errorf("CPX Immediate failed: Carry flag should not be set")
		}
	}

	// Check if the zero flag is set correctly (X == immediate)
	if cpu.X == 0x10 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("CPX Immediate failed: Zero flag not set correctly")
		}
	} else {
		if cpu.getSRBit(1) != 0 {
			t.Errorf("CPX Immediate failed: Zero flag should not be set")
		}
	}

	// Check if the negative flag is set correctly
	if (cpu.X-0x10)&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("CPX Immediate failed: Negative flag not set correctly")
		}
	} else {
		if cpu.getSRBit(7) != 0 {
			t.Errorf("CPX Immediate failed: Negative flag should not be set")
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for CPX Immediate
		t.Errorf("CPX Immediate failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestCPXZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// CPX $10 (Zero Page Addressing)
	cpu.writeMemory(cpu.PC, CPX_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero page address
	cpu.writeMemory(0x10, 0x15)     // Value at the zero page address for comparison
	cpu.X = 0x20                    // Set X register for comparison

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the carry flag is set correctly (X >= memory)
	if cpu.X >= 0x15 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("CPX Zero Page failed: Carry flag not set correctly")
		}
	} else {
		if cpu.getSRBit(0) != 0 {
			t.Errorf("CPX Zero Page failed: Carry flag should not be set")
		}
	}

	// Check if the zero flag is set correctly (X == memory)
	if cpu.X == 0x15 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("CPX Zero Page failed: Zero flag not set correctly")
		}
	} else {
		if cpu.getSRBit(1) != 0 {
			t.Errorf("CPX Zero Page failed: Zero flag should not be set")
		}
	}

	// Check if the negative flag is set correctly
	if (cpu.X-0x15)&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("CPX Zero Page failed: Negative flag not set correctly")
		}
	} else {
		if cpu.getSRBit(7) != 0 {
			t.Errorf("CPX Zero Page failed: Negative flag should not be set")
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for CPX Zero Page
		t.Errorf("CPX Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestCPXAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// CPX $1000 (Absolute Addressing)
	cpu.writeMemory(cpu.PC, CPX_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1000, 0x15) // Value at the absolute address for comparison
	cpu.X = 0x20                  // Set X register for comparison

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the carry flag is set correctly (X >= memory)
	if cpu.X >= 0x15 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("CPX Absolute failed: Carry flag not set correctly")
		}
	} else {
		if cpu.getSRBit(0) != 0 {
			t.Errorf("CPX Absolute failed: Carry flag should not be set")
		}
	}

	// Check if the zero flag is set correctly (X == memory)
	if cpu.X == 0x15 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("CPX Absolute failed: Zero flag not set correctly")
		}
	} else {
		if cpu.getSRBit(1) != 0 {
			t.Errorf("CPX Absolute failed: Zero flag should not be set")
		}
	}

	// Check if the negative flag is set correctly
	if (cpu.X-0x15)&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("CPX Absolute failed: Negative flag not set correctly")
		}
	} else {
		if cpu.getSRBit(7) != 0 {
			t.Errorf("CPX Absolute failed: Negative flag should not be set")
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for CPX Absolute
		t.Errorf("CPX Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}
func TestCPYImmediate(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// CPY #10
	cpu.writeMemory(cpu.PC, CPY_IMMEDIATE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Immediate value for comparison
	cpu.Y = 0x20                    // Set Y register for comparison

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the carry flag is set correctly (Y >= immediate)
	if cpu.Y >= 0x10 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("CPY Immediate failed: Carry flag not set correctly")
		}
	} else {
		if cpu.getSRBit(0) != 0 {
			t.Errorf("CPY Immediate failed: Carry flag should not be set")
		}
	}

	// Check if the zero flag is set correctly (Y == immediate)
	if cpu.Y == 0x10 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("CPY Immediate failed: Zero flag not set correctly")
		}
	} else {
		if cpu.getSRBit(1) != 0 {
			t.Errorf("CPY Immediate failed: Zero flag should not be set")
		}
	}

	// Check if the negative flag is set correctly
	if (cpu.Y-0x10)&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("CPY Immediate failed: Negative flag not set correctly")
		}
	} else {
		if cpu.getSRBit(7) != 0 {
			t.Errorf("CPY Immediate failed: Negative flag should not be set")
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for CPY Immediate
		t.Errorf("CPY Immediate failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestCPYZeroPage(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// CPY $10 (Zero Page Addressing)
	cpu.writeMemory(cpu.PC, CPY_ZERO_PAGE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x10) // Zero page address
	cpu.writeMemory(0x10, 0x15)     // Value at the zero page address for comparison
	cpu.Y = 0x20                    // Set Y register for comparison

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the carry flag is set correctly (Y >= memory)
	if cpu.Y >= 0x15 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("CPY Zero Page failed: Carry flag not set correctly")
		}
	} else {
		if cpu.getSRBit(0) != 0 {
			t.Errorf("CPY Zero Page failed: Carry flag should not be set")
		}
	}

	// Check if the zero flag is set correctly (Y == memory)
	if cpu.Y == 0x15 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("CPY Zero Page failed: Zero flag not set correctly")
		}
	} else {
		if cpu.getSRBit(1) != 0 {
			t.Errorf("CPY Zero Page failed: Zero flag should not be set")
		}
	}

	// Check if the negative flag is set correctly
	if (cpu.Y-0x15)&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("CPY Zero Page failed: Negative flag not set correctly")
		}
	} else {
		if cpu.getSRBit(7) != 0 {
			t.Errorf("CPY Zero Page failed: Negative flag should not be set")
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for CPY Zero Page
		t.Errorf("CPY Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
	}
}
func TestCPYAbsolute(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// CPY $1000 (Absolute Addressing)
	cpu.writeMemory(cpu.PC, CPY_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address
	cpu.writeMemory(cpu.PC+2, 0x10)
	cpu.writeMemory(0x1000, 0x15) // Value at the absolute address for comparison
	cpu.Y = 0x20                  // Set Y register for comparison

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the carry flag is set correctly (Y >= memory)
	if cpu.Y >= 0x15 {
		if cpu.getSRBit(0) != 1 {
			t.Errorf("CPY Absolute failed: Carry flag not set correctly")
		}
	}

	// Check if the zero flag is set correctly (Y == memory)
	if cpu.Y == 0x15 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("CPY Absolute failed: Zero flag not set correctly")
		}
	} else {
		if cpu.getSRBit(1) != 0 {
			t.Errorf("CPY Absolute failed: Zero flag should not be set")
		}
	}

	// Check if the negative flag is set correctly
	if (cpu.Y-0x15)&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("CPY Absolute failed: Negative flag not set correctly")
		}
	} else {
		if cpu.getSRBit(7) != 0 {
			t.Errorf("CPY Absolute failed: Negative flag should not be set")
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for CPY Absolute
		t.Errorf("CPY Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
	}
}

func TestBRK(t *testing.T) {
	var cpu CPU
	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, BRK_OPCODE)
	// Update Reset Vector address with address for BRK
	cpu.writeMemory(0xFFFE, 0x00)
	cpu.writeMemory(0xFFFF, 0x10)

	cpu.cpuQuit = true
	cpu.startCPU()

	// Read the pushed status register
	pushedSR := cpu.readMemory(SPBaseAddress + cpu.SP)

	// Check if the break flag is set correctly in the pushed status register
	if pushedSR&0x10 == 0 {
		t.Errorf("BRK failed: Break flag not set correctly in pushed SR")
	}

	// Check if the interrupt disable flag is set correctly
	if cpu.getSRBit(2) != 1 {
		t.Errorf("BRK failed: Interrupt disable flag not set correctly")
	}

	// Check if the Program Counter is incremented correctly
	if cpu.PC != 0x1000 {
		t.Errorf("BRK failed: expected PC = %04X, got %04X", 0x1000, cpu.PC)
	}
}
func TestCLC(t *testing.T) {
	var cpu CPU
	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, CLC_OPCODE)

	cpu.cpuQuit = true
	cpu.startCPU()

	// Check if the carry flag is cleared correctly
	if cpu.getSRBit(0) != 0 {
		t.Errorf("CLC failed: Carry flag should not be set")
	}
}
func TestCLD(t *testing.T) {
	var cpu CPU
	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, CLD_OPCODE)

	cpu.cpuQuit = true
	cpu.startCPU()

	// Check if the decimal flag is cleared correctly
	if cpu.getSRBit(3) != 0 {
		t.Errorf("CLD failed: Decimal flag should not be set")
	}
}
func TestCLI(t *testing.T) {
	var cpu CPU
	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, CLI_OPCODE)

	cpu.cpuQuit = true
	cpu.startCPU()

	// Check if the interrupt disable flag is cleared correctly
	if cpu.getSRBit(2) != 0 {
		t.Errorf("CLI failed: Interrupt disable flag should not be set")
	}
}
func TestCLV(t *testing.T) {
	var cpu CPU
	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, CLV_OPCODE)

	cpu.cpuQuit = true
	cpu.startCPU()

	// Check if the overflow flag is cleared correctly
	if cpu.getSRBit(6) != 0 {
		t.Errorf("CLV failed: Overflow flag should not be set")
	}
}
func TestDEX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// DEX
	cpu.writeMemory(cpu.PC, DEX_OPCODE) // Opcode for DEX
	initialX := byte(0x20)
	cpu.X = initialX // Set initial value of X register

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialX - 1 // Expected value after decrement

	// Check if X register has the expected decremented value
	if cpu.X != expectedValue {
		t.Errorf("DEX failed: expected X = %02X, got %02X", expectedValue, cpu.X)
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("DEX failed: Negative flag not set correctly")
		}
	} else {
		if cpu.getSRBit(7) != 0 {
			t.Errorf("DEX failed: Negative flag should not be set")
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("DEX failed: Zero flag not set correctly")
		}
	} else {
		if cpu.getSRBit(1) != 0 {
			t.Errorf("DEX failed: Zero flag should not be set")
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for DEX
		t.Errorf("DEX failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}
func TestDEY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)

	// DEY
	cpu.writeMemory(cpu.PC, DEY_OPCODE) // Opcode for DEY
	initialY := byte(0x20)
	cpu.Y = initialY // Set initial value of Y register

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialY - 1 // Expected value after decrement

	// Check if Y register has the expected decremented value
	if cpu.Y != expectedValue {
		t.Errorf("DEY failed: expected Y = %02X, got %02X", expectedValue, cpu.Y)
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("DEY failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("DEY failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("DEY failed: Zero flag not set correctly")
		}
	} else {
		if cpu.getSRBit(1) != 0 {
			t.Errorf("DEY failed: Zero flag should not be set")
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for DEY
		t.Errorf("DEY failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}
func TestINX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// INX
	cpu.writeMemory(cpu.PC, INX_OPCODE) // Opcode for INX
	initialX := byte(0x20)
	cpu.X = initialX // Set initial value of X register

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialX + 1 // Expected value after increment

	// Check if X register has the expected incremented value
	if cpu.X != expectedValue {
		t.Errorf("INX failed: expected X = %02X, got %02X", expectedValue, cpu.X)
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("INX failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("INX failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("INX failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("INX failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for INX
		t.Errorf("INX failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}
func TestINY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// INY
	cpu.writeMemory(cpu.PC, INY_OPCODE) // Opcode for INY
	initialY := byte(0x20)
	cpu.Y = initialY // Set initial value of Y register

	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	expectedValue := initialY + 1 // Expected value after increment

	// Check if Y register has the expected incremented value
	if cpu.Y != expectedValue {
		t.Errorf("INY failed: expected Y = %02X, got %02X", expectedValue, cpu.Y)
	}

	// Check if the negative flag is set correctly
	if expectedValue&0x80 != 0 {
		if cpu.getSRBit(7) != 1 {
			t.Errorf("INY failed: Negative flag not set correctly")
		} else {
			if cpu.getSRBit(7) != 0 {
				t.Errorf("INY failed: Negative flag should not be set")
			}
		}
	}

	// Check if the zero flag is set correctly
	if expectedValue == 0 {
		if cpu.getSRBit(1) != 1 {
			t.Errorf("INY failed: Zero flag not set correctly")
		} else {
			if cpu.getSRBit(1) != 0 {
				t.Errorf("INY failed: Zero flag should not be set")
			}
		}
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for INY
		t.Errorf("INY failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}
func TestNOP(t *testing.T) {
	var cpu CPU
	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, NOP_OPCODE)

	cpu.cpuQuit = true
	cpu.startCPU()

	// Check if the program counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 {
		t.Errorf("NOP failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}
func TestPHA(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.preOpSP = 0xFD
	// PHA
	cpu.writeMemory(cpu.PC, PHA_OPCODE)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if memory has the expected value
	if cpu.readMemory(SPBaseAddress+cpu.preOpSP) != 0x20 {
		t.Errorf("PHA failed: got %02X, want %02X", cpu.readMemory(SPBaseAddress+cpu.preOpSP), 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for PHA
		t.Errorf("PHA failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}
func TestPHP(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	initialSP := cpu.SP // Store the initial value of the stack pointer

	// Set some flags in the status register to test
	cpu.SR = 0b11010101 // Set the status register

	// PHP opcode
	cpu.writeMemory(cpu.PC, PHP_OPCODE)
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the status register is correctly pushed onto the stack
	expectedStatus := byte(0b11010101 | 0x10 | 0x20) // Set both the B flag and the unused bit

	actualStatus := cpu.readMemory(SPBaseAddress + cpu.SP + 1) // +1 because SP points to next free space

	if actualStatus != expectedStatus {
		t.Errorf("PHP failed: got status %08b, want %08b", actualStatus, expectedStatus)
	}

	// Check if the stack pointer is decremented correctly
	if cpu.SP != initialSP-1 {
		t.Errorf("PHP failed: expected SP = %02X, got %02X", initialSP-1, cpu.SP)
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for PHP
		t.Errorf("PHP failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}

func TestPLA(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	initialSP := uint16(0xFD) // Initial stack pointer value
	cpu.SP = initialSP        // Setting SP to initial value
	// PLA
	cpu.writeMemory(cpu.PC, PLA_OPCODE)
	cpu.writeMemory(SPBaseAddress+cpu.SP+1, 0x20) // Push 0x20 onto the stack
	cpu.cpuQuit = true                            // Stop the CPU after one execution cycle
	cpu.startCPU()                                // Initialize the CPU state

	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("PLA failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for PLA
		t.Errorf("PLA failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}

func TestPLP(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	initialSP := cpu.SP // Store the initial value of the stack pointer

	// Set some flags in the status register to test, including the break flag and the unused bit
	cpu.SR = 0b11010101 // Set the status register, including break flag and unused bit

	// Push the status register onto the stack
	cpu.decSP()
	cpu.updateStack(cpu.SR)

	// PLP opcode
	cpu.writeMemory(cpu.PC, PLP_OPCODE)
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if the status register is correctly restored from the stack
	expectedStatus := byte(0b11010101) // The expected status register value after PLP

	if cpu.SR != expectedStatus {
		t.Errorf("PLP failed: got status %08b, want %08b", cpu.SR, expectedStatus)
	}

	// Check if the stack pointer is incremented correctly
	if cpu.SP != initialSP {
		t.Errorf("PLP failed: expected SP = %02X, got %02X", initialSP, cpu.SP)
	}

	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for PLP
		t.Errorf("PLP failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}

func TestRTI(t *testing.T) {
	var cpu CPU
	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, RTI_OPCODE)

	returnAddress := uint16(0x1234)
	processorStatus := byte(0x20)
	cpu.SP = 0xFD // Initial stack pointer

	// Push return address and processor status onto the stack in reverse order
	cpu.updateStack(byte(returnAddress >> 8)) // High byte of return address
	cpu.decSP()
	cpu.updateStack(byte(returnAddress & 0xFF)) // Low byte of return address
	cpu.decSP()
	cpu.updateStack(processorStatus) // Processor status
	cpu.decSP()

	cpu.cpuQuit = true
	cpu.startCPU()

	// Verify the program counter and processor status
	if cpu.PC != returnAddress {
		t.Errorf("RTI failed: expected PC = %04X, got %04X", returnAddress, cpu.PC)
	}
	if cpu.SR != processorStatus {
		t.Errorf("RTI failed: expected SR = %02X, got %02X", processorStatus, cpu.SR)
	}
	expectedSP := uint16(0xFD)
	if cpu.SP != expectedSP {
		t.Errorf("RTI failed: expected SP = %02X, got %02X", expectedSP, cpu.SP)
	}
}

func TestRTS(t *testing.T) {
	var cpu CPU
	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, RTS_OPCODE)

	returnAddress := uint16(0x1234)
	cpu.SP = 0xFF // Set stack pointer to initial value

	// Push the return address onto the stack: high byte, then low byte
	cpu.decSP()
	cpu.updateStack(byte(returnAddress >> 8)) // High byte of return address
	cpu.decSP()
	cpu.updateStack(byte(returnAddress & 0xFF)) // Low byte of return address

	cpu.cpuQuit = true
	cpu.startCPU()

	// Check if the program counter is set correctly
	if cpu.PC != returnAddress+1 {
		t.Errorf("RTS failed: expected PC = %04X, got %04X", returnAddress+1, cpu.PC)
	}

	// Check if the stack pointer is updated correctly
	expectedSP := uint16(0xFF) // Stack pointer should be back to its initial value
	if cpu.SP != expectedSP {
		t.Errorf("RTS failed: expected SP = %02X, got %02X", expectedSP, cpu.SP)
	}
}

func TestSEC(t *testing.T) {
	var cpu CPU
	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, SEC_OPCODE)

	cpu.cpuQuit = true
	cpu.startCPU()

	// Check if the carry flag is set correctly
	if cpu.getSRBit(0) != 1 {
		t.Errorf("SEC failed: Carry flag not set correctly")
	}
}
func TestSED(t *testing.T) {
	var cpu CPU
	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, SED_OPCODE)

	cpu.cpuQuit = true
	cpu.startCPU()

	// Check if the decimal flag is set correctly
	if cpu.getSRBit(3) != 1 {
		t.Errorf("SED failed: Decimal flag not set correctly")
	}
}
func TestSEI(t *testing.T) {
	var cpu CPU
	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.writeMemory(cpu.PC, SEI_OPCODE)

	cpu.cpuQuit = true
	cpu.startCPU()

	// Check if the interrupt disable flag is set correctly
	if cpu.getSRBit(2) != 1 {
		t.Errorf("SEI failed: Interrupt disable flag not set correctly")
	}
}
func TestTAX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// TAX
	cpu.writeMemory(cpu.PC, TAX_OPCODE)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if X has the expected value
	if cpu.X != 0x20 {
		t.Errorf("TAX failed: got %02X, want %02X", cpu.X, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for TAX
		t.Errorf("TAX failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}
func TestTAY(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// TAY
	cpu.writeMemory(cpu.PC, TAY_OPCODE)
	cpu.A = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if Y has the expected value
	if cpu.Y != 0x20 {
		t.Errorf("TAY failed: got %02X, want %02X", cpu.Y, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for TAY
		t.Errorf("TAY failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}
func TestTSX(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// TSX
	cpu.writeMemory(cpu.PC, TSX_OPCODE)
	cpu.SP = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if X has the expected value
	if cpu.X != 0x20 {
		t.Errorf("TSX failed: got %02X, want %02X", cpu.X, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for TSX
		t.Errorf("TSX failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}
func TestTXA(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// TXA
	cpu.writeMemory(cpu.PC, TXA_OPCODE)
	cpu.X = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("TXA failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for TXA
		t.Errorf("TXA failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}
func TestTXS(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// TXS
	cpu.writeMemory(cpu.PC, TXS_OPCODE)
	cpu.X = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if SP has the expected value
	if cpu.SP != 0x20 {
		t.Errorf("TXS failed: got %02X, want %02X", cpu.SP, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for TXS
		t.Errorf("TXS failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}

func TestTYA(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	// TYA
	cpu.writeMemory(cpu.PC, TYA_OPCODE)
	cpu.Y = 0x20
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("TYA failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for TYA
		t.Errorf("TYA failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
	}
}

func TestJSR(t *testing.T) {
	var cpu CPU
	cpu.resetCPU()
	cpu.setPC(0x0000)

	// Write JSR instruction and its operand (address of subroutine)
	cpu.writeMemory(cpu.PC, JSR_ABSOLUTE_OPCODE)
	cpu.writeMemory(cpu.PC+1, 0x34) // Low byte of subroutine address
	cpu.writeMemory(cpu.PC+2, 0x12) // High byte of subroutine address

	cpu.cpuQuit = true
	cpu.startCPU()

	// Check stack values
	expectedHighByte := byte(0x00)
	expectedLowByte := byte(0x02) // PC was 0x0000, next instruction at 0x0003, pushed value is 0x0002
	actualHighByte := cpu.readMemory(SPBaseAddress + cpu.SP + 1)
	actualLowByte := cpu.readMemory(SPBaseAddress + cpu.SP)

	if expectedHighByte != actualHighByte || expectedLowByte != actualLowByte {
		t.Errorf("JSR failed: expected stack value = %02X%02X, got %02X%02X", expectedHighByte, expectedLowByte, actualHighByte, actualLowByte)
	}

	// Check stack pointer
	expectedSP := uint16(0xFD) // Two bytes pushed to the stack
	if cpu.SP != expectedSP {
		t.Errorf("JSR failed: expected SP = %02X, got %02X", expectedSP, cpu.SP)
	}

	// Check program counter (should be set to the subroutine address)
	expectedPC := uint16(0x1234) // Address of the subroutine
	if cpu.PC != expectedPC {
		t.Errorf("JSR failed: expected PC = %04X, got %04X", expectedPC, cpu.PC)
	}
}
