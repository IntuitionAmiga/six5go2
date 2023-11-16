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

func TestPLA(t *testing.T) {
	var cpu CPU // Create a new CPU instance for the test

	cpu.resetCPU()
	cpu.setPC(0x0000)
	cpu.preOpSP = 0xFD
	// PLA
	cpu.writeMemory(cpu.PC, PLA_OPCODE)
	cpu.writeMemory(SPBaseAddress+cpu.preOpSP, 0x20)
	cpu.SP--
	cpu.cpuQuit = true // Stop the CPU after one execution cycle
	cpu.startCPU()     // Initialize the CPU state

	// Check if A has the expected value
	if cpu.A != 0x20 {
		t.Errorf("PLA failed: got %02X, want %02X", cpu.A, 0x20)
	}
	// Check if Program Counter is incremented correctly
	if cpu.PC != cpu.preOpPC+1 { // 1 byte for PLA
		t.Errorf("PLA failed: expected PC = %04X, got %04X", cpu.preOpPC+1, cpu.PC)
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

//func TestANDIndirectX(t *testing.T) {
//	var cpu CPU // Create a new CPU instance for the test
//
//	cpu.resetCPU()
//	cpu.setPC(0x0000)
//
//	// AND ($10,X)
//	cpu.writeMemory(cpu.PC, AND_INDIRECT_X_OPCODE)
//	cpu.writeMemory(cpu.PC+1, 0x10)
//	cpu.writeMemory(0x0010, 0x00)
//	cpu.writeMemory(0x0011, 0x10)
//	cpu.writeMemory(0x1000, 0x20)
//	cpu.A = 0x20
//	cpu.X = 0x01
//	cpu.cpuQuit = true // Stop the CPU after one execution cycle
//	cpu.startCPU()     // Initialize the CPU state
//
//	// Check if A has the expected value
//	if cpu.A != 0x20 {
//		t.Errorf("AND Indirect X failed: got %02X, want %02X", cpu.A, 0x20)
//	}
//	// Check if Program Counter is incremented correctly
//	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for AND immediate
//		t.Errorf("AND Indirect X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
//	}
//}

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

//func TestEORIndirectX(t *testing.T) {
//	var cpu CPU // Create a new CPU instance for the test
//
//	cpu.resetCPU()
//	cpu.setPC(0x0000)
//
//	// EOR ($10,X)
//	cpu.writeMemory(cpu.PC, EOR_INDIRECT_X_OPCODE)
//	cpu.writeMemory(cpu.PC+1, 0x10)
//	cpu.writeMemory(0x0010, 0x00)
//	cpu.writeMemory(0x0011, 0x10)
//	cpu.writeMemory(0x1000, 0x20)
//	cpu.A = 0x20
//	cpu.X = 0x01
//	cpu.cpuQuit = true // Stop the CPU after one execution cycle
//	cpu.startCPU()     // Initialize the CPU state
//	// Check if A has the expected value
//	if cpu.A != 0x00 {
//		t.Errorf("EOR Indirect X failed: got %02X, want %02X", cpu.A, 0x00)
//	}
//	// Check if Program Counter is incremented correctly
//	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for EOR immediate
//		t.Errorf("EOR Indirect X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
//	}
//}

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

//func TestADCIndirectX(t *testing.T) {
//	var cpu CPU // Create a new CPU instance for the test
//
//	cpu.resetCPU()
//	cpu.setPC(0x0000)
//	// ADC ($10,X)
//	cpu.writeMemory(cpu.PC, ADC_INDIRECT_X_OPCODE)
//	cpu.writeMemory(cpu.PC+1, 0x10)
//	cpu.writeMemory(0x0010, 0x00)
//	cpu.writeMemory(0x0011, 0x10)
//	cpu.writeMemory(0x1000, 0x10)
//	cpu.A = 0x20
//	cpu.X = 0x01
//	cpu.cpuQuit = true // Stop the CPU after one execution cycle
//	cpu.startCPU()     // Initialize the CPU state
//
//	// Check if A has the expected value
//	if cpu.A != 0x30 {
//		t.Errorf("ADC Indirect X failed: got %02X, want %02X", cpu.A, 0x30)
//	}
//	// Check if Program Counter is incremented correctly
//	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for ADC immediate
//		t.Errorf("ADC Indirect X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
//	}
//}

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

//func TestSBCImmediate(t *testing.T) {
//	var cpu CPU // Create a new CPU instance for the test
//
//	cpu.resetCPU()
//	cpu.setPC(0x0000)
//
//	// SBC #10
//	cpu.writeMemory(cpu.PC, SBC_IMMEDIATE_OPCODE)
//	cpu.writeMemory(cpu.PC+1, 0x10) // Immediate value to subtract
//	cpu.A = 0x20                    // Set accumulator for subtraction
//	cpu.setCarryFlag()              // Ensure the carry flag is set before subtraction
//
//	cpu.cpuQuit = true // Stop the CPU after one execution cycle
//	cpu.startCPU()     // Initialize the CPU state
//
//	expectedValue := cpu.A - 0x10 - (1 - cpu.getSRBit(0)) // Adjusted expected value after subtraction
//
//	// Check if accumulator has the expected value
//	if cpu.A != expectedValue {
//		t.Errorf("SBC Immediate failed: got %02X, want %02X", cpu.A, expectedValue)
//	}
//
//	// Check if the zero flag is set correctly
//	if expectedValue == 0 && cpu.getSRBit(1) != 1 {
//		t.Errorf("SBC Immediate failed: Zero flag not set correctly")
//	}
//
//	// Check if the negative flag is set correctly
//	if (expectedValue&0x80) != 0 && cpu.getSRBit(7) != 1 {
//		t.Errorf("SBC Immediate failed: Negative flag not set correctly")
//	}
//
//	// Check if the carry flag is set correctly (no borrow)
//	if expectedValue >= 0 && cpu.getSRBit(0) != 1 {
//		t.Errorf("SBC Immediate failed: Carry flag not set correctly")
//	}
//
//	// Check if Program Counter is incremented correctly
//	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for SBC Immediate
//		t.Errorf("SBC Immediate failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
//	}
//}
//
//func TestSBCZeroPage(t *testing.T) {
//	var cpu CPU // Create a new CPU instance for the test
//
//	cpu.resetCPU()
//	cpu.setPC(0x0000)
//
//	// SBC $10
//	cpu.writeMemory(cpu.PC, SBC_ZERO_PAGE_OPCODE)
//	cpu.writeMemory(cpu.PC+1, 0x10) // Zero Page address to subtract
//	cpu.writeMemory(0x0010, 0x10)   // Value at Zero Page address to subtract
//	cpu.A = 0x20                    // Set accumulator for subtraction
//	cpu.setCarryFlag()              // Ensure the carry flag is set before subtraction
//
//	cpu.cpuQuit = true // Stop the CPU after one execution cycle
//	cpu.startCPU()     // Initialize the CPU state
//
//	expectedValue := cpu.A - 0x10 - (1 - cpu.getSRBit(0)) // Adjusted expected value after subtraction
//
//	// Check if accumulator has the expected value
//	if cpu.A != expectedValue {
//		t.Errorf("SBC Zero Page failed: got %02X, want %02X", cpu.A, expectedValue)
//	}
//
//	// Check if the zero flag is set correctly
//	if expectedValue == 0 && cpu.getSRBit(1) != 1 {
//		t.Errorf("SBC Zero Page failed: Zero flag not set correctly")
//	}
//
//	// Check if the negative flag is set correctly
//	if (expectedValue&0x80) != 0 && cpu.getSRBit(7) != 1 {
//		t.Errorf("SBC Zero Page failed: Negative flag not set correctly")
//	}
//
//	// Check if the carry flag is set correctly (no borrow)
//	if expectedValue >= 0 && cpu.getSRBit(0) != 1 {
//		t.Errorf("SBC Zero Page failed: Carry flag not set correctly")
//	}
//
//	// Check if Program Counter is incremented correctly
//	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for SBC Zero Page
//		t.Errorf("SBC Zero Page failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
//	}
//}
//
//func TestSBCZeroPageX(t *testing.T) {
//	var cpu CPU // Create a new CPU instance for the test
//
//	cpu.resetCPU()
//	cpu.setPC(0x0000)
//	// SBC $10,X
//	cpu.writeMemory(cpu.PC, SBC_ZERO_PAGE_X_OPCODE)
//	cpu.writeMemory(cpu.PC+1, 0x10) // Zero Page address to subtract
//	cpu.writeMemory(0x0011, 0x10)   // Value at Zero Page address to subtract
//	cpu.X = 0x01                    // Set X register for Zero Page addressing
//	cpu.A = 0x20                    // Set accumulator for subtraction
//	cpu.setCarryFlag()              // Ensure the carry flag is set before subtraction
//
//	cpu.cpuQuit = true // Stop the CPU after one execution cycle
//	cpu.startCPU()     // Initialize the CPU state
//
//	expectedValue := cpu.A - 0x10 - (1 - cpu.getSRBit(0)) // Adjusted expected value after subtraction
//
//	// Check if accumulator has the expected value
//	if cpu.A != expectedValue {
//		t.Errorf("SBC Zero Page X failed: got %02X, want %02X", cpu.A, expectedValue)
//	}
//
//	// Check if the zero flag is set correctly
//	if expectedValue == 0 && cpu.getSRBit(1) != 1 {
//		t.Errorf("SBC Zero Page X failed: Zero flag not set correctly")
//	}
//
//	// Check if the negative flag is set correctly
//	if (expectedValue&0x80) != 0 && cpu.getSRBit(7) != 1 {
//		t.Errorf("SBC Zero Page X failed: Negative flag not set correctly")
//	}
//
//	// Check if the carry flag is set correctly (no borrow)
//	if expectedValue >= 0 && cpu.getSRBit(0) != 1 {
//		t.Errorf("SBC Zero Page X failed: Carry flag not set correctly")
//	}
//
//	// Check if Program Counter is incremented correctly
//	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for SBC Zero Page
//		t.Errorf("SBC Zero Page X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
//	}
//}
//
//func TestSBCAbsolute(t *testing.T) {
//	var cpu CPU // Create a new CPU instance for the test
//
//	cpu.resetCPU()
//	cpu.setPC(0x0000)
//	// SBC $1000
//	cpu.writeMemory(cpu.PC, SBC_ABSOLUTE_OPCODE)
//	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address to subtract
//	cpu.writeMemory(cpu.PC+2, 0x10)
//	cpu.writeMemory(0x1000, 0x10) // Value at Absolute address to subtract
//	cpu.A = 0x20                  // Set accumulator for subtraction
//	cpu.setCarryFlag()            // Ensure the carry flag is set before subtraction
//
//	cpu.cpuQuit = true // Stop the CPU after one execution cycle
//	cpu.startCPU()     // Initialize the CPU state
//
//	expectedValue := cpu.A - 0x10 - (1 - cpu.getSRBit(0)) // Adjusted expected value after subtraction
//
//	// Check if accumulator has the expected value
//	if cpu.A != expectedValue {
//		t.Errorf("SBC Absolute failed: got %02X, want %02X", cpu.A, expectedValue)
//	}
//
//	// Check if the zero flag is set correctly
//	if expectedValue == 0 && cpu.getSRBit(1) != 1 {
//		t.Errorf("SBC Absolute failed: Zero flag not set correctly")
//	}
//
//	// Check if the negative flag is set correctly
//	if (expectedValue&0x80) != 0 && cpu.getSRBit(7) != 1 {
//		t.Errorf("SBC Absolute failed: Negative flag not set correctly")
//	}
//
//	// Check if the carry flag is set correctly (no borrow)
//	if expectedValue >= 0 && cpu.getSRBit(0) != 1 {
//		t.Errorf("SBC Absolute failed: Carry flag not set correctly")
//	}
//
//	// Check if Program Counter is incremented correctly
//	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for SBC Absolute
//		t.Errorf("SBC Absolute failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
//	}
//}
//
//func TestSBCAbsoluteX(t *testing.T) {
//	var cpu CPU // Create a new CPU instance for the test
//
//	cpu.resetCPU()
//	cpu.setPC(0x0000)
//	// SBC $1000,X
//	cpu.writeMemory(cpu.PC, SBC_ABSOLUTE_X_OPCODE)
//	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address to subtract
//	cpu.writeMemory(cpu.PC+2, 0x10)
//	cpu.writeMemory(0x1001, 0x10) // Value at Absolute address to subtract
//	cpu.X = 0x01                  // Set X register for Absolute addressing
//	cpu.A = 0x20                  // Set accumulator for subtraction
//	cpu.setCarryFlag()            // Ensure the carry flag is set before subtraction
//
//	cpu.cpuQuit = true // Stop the CPU after one execution cycle
//	cpu.startCPU()     // Initialize the CPU state
//
//	expectedValue := cpu.A - 0x10 - (1 - cpu.getSRBit(0)) // Adjusted expected value after subtraction
//
//	// Check if accumulator has the expected value
//	if cpu.A != expectedValue {
//		t.Errorf("SBC Absolute X failed: got %02X, want %02X", cpu.A, expectedValue)
//	}
//
//	// Check if the zero flag is set correctly
//	if expectedValue == 0 && cpu.getSRBit(1) != 1 {
//		t.Errorf("SBC Absolute X failed: Zero flag not set correctly")
//	}
//
//	// Check if the negative flag is set correctly
//	if (expectedValue&0x80) != 0 && cpu.getSRBit(7) != 1 {
//		t.Errorf("SBC Absolute X failed: Negative flag not set correctly")
//	}
//
//	// Check if the carry flag is set correctly (no borrow)
//	if expectedValue >= 0 && cpu.getSRBit(0) != 1 {
//		t.Errorf("SBC Absolute X failed: Carry flag not set correctly")
//	}
//
//	// Check if Program Counter is incremented correctly
//	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for SBC Absolute
//		t.Errorf("SBC Absolute X failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
//	}
//}
//
//func TestSBCAbsoluteY(t *testing.T) {
//	var cpu CPU // Create a new CPU instance for the test
//
//	cpu.resetCPU()
//	cpu.setPC(0x0000)
//	// SBC $1000,Y
//	cpu.writeMemory(cpu.PC, SBC_ABSOLUTE_Y_OPCODE)
//	cpu.writeMemory(cpu.PC+1, 0x00) // Absolute address to subtract
//	cpu.writeMemory(cpu.PC+2, 0x10)
//	cpu.writeMemory(0x1001, 0x10) // Value at Absolute address to subtract
//	cpu.Y = 0x01                  // Set Y register for Absolute addressing
//	cpu.A = 0x20                  // Set accumulator for subtraction
//	cpu.setCarryFlag()            // Ensure the carry flag is set before subtraction
//
//	cpu.cpuQuit = true // Stop the CPU after one execution cycle
//	cpu.startCPU()     // Initialize the CPU state
//
//	expectedValue := cpu.A - 0x10 - (1 - cpu.getSRBit(0)) // Adjusted expected value after subtraction
//
//	// Check if accumulator has the expected value
//	if cpu.A != expectedValue {
//		t.Errorf("SBC Absolute Y failed: got %02X, want %02X", cpu.A, expectedValue)
//	}
//
//	// Check if the zero flag is set correctly
//	if expectedValue == 0 && cpu.getSRBit(1) != 1 {
//		t.Errorf("SBC Absolute Y failed: Zero flag not set correctly")
//	}
//
//	// Check if the negative flag is set correctly
//	if (expectedValue&0x80) != 0 && cpu.getSRBit(7) != 1 {
//		t.Errorf("SBC Absolute Y failed: Negative flag not set correctly")
//	}
//
//	// Check if the carry flag is set correctly (no borrow)
//	if expectedValue >= 0 && cpu.getSRBit(0) != 1 {
//		t.Errorf("SBC Absolute Y failed: Carry flag not set correctly")
//	}
//
//	// Check if Program Counter is incremented correctly
//	if cpu.PC != cpu.preOpPC+3 { // 3 bytes for SBC Absolute
//		t.Errorf("SBC Absolute Y failed: expected PC = %04X, got %04X", cpu.preOpPC+3, cpu.PC)
//	}
//}
//
//func TestSBCIndirectX(t *testing.T) {
//	var cpu CPU // Create a new CPU instance for the test
//
//	cpu.resetCPU()
//	cpu.setPC(0x0000)
//	// SBC ($10,X)
//	cpu.writeMemory(cpu.PC, SBC_INDIRECT_X_OPCODE)
//	cpu.writeMemory(cpu.PC+1, 0x10) // Zero Page address to subtract
//	cpu.writeMemory(0x0010, 0x00)   // Value at Zero Page address to subtract
//	cpu.writeMemory(0x0011, 0x10)   // Value at Zero Page address to subtract
//	cpu.writeMemory(0x1000, 0x10)   // Value at Zero Page address to subtract
//	cpu.X = 0x01                    // Set X register for Zero Page addressing
//	cpu.A = 0x20                    // Set accumulator for subtraction
//	cpu.setCarryFlag()              // Ensure the carry flag is set before subtraction
//
//	cpu.cpuQuit = true // Stop the CPU after one execution cycle
//	cpu.startCPU()     // Initialize the CPU state
//
//	expectedValue := cpu.A - 0x10 - (1 - cpu.getSRBit(0)) // Adjusted expected value after subtraction
//
//	// Check if accumulator has the expected value
//	if cpu.A != expectedValue {
//		t.Errorf("SBC Indirect X failed: got %02X, want %02X", cpu.A, expectedValue)
//	}
//
//	// Check if the zero flag is set correctly
//	if expectedValue == 0 && cpu.getSRBit(1) != 1 {
//		t.Errorf("SBC Indirect X failed: Zero flag not set correctly")
//	}
//
//	// Check if the negative flag is set correctly
//	if (expectedValue&0x80) != 0 && cpu.getSRBit(7) != 1 {
//		t.Errorf("SBC Indirect X failed: Negative flag not set correctly")
//	}
//
//	// Check if the carry flag is set correctly (no borrow)
//	if expectedValue >= 0 && cpu.getSRBit(0) != 1 {
//		t.Errorf("SBC Indirect X failed: Carry flag not set correctly")
//	}
//
//	// Check if Program Counter is incremented correctly
//	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for SBC Zero Page
//		t.Errorf("SBC Indirect X failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
//	}
//}
//
//func TestSBCIndirectY(t *testing.T) {
//	var cpu CPU // Create a new CPU instance for the test
//
//	cpu.resetCPU()
//	cpu.setPC(0x0000)
//	// SBC ($10),Y
//	cpu.writeMemory(cpu.PC, SBC_INDIRECT_Y_OPCODE)
//	cpu.writeMemory(cpu.PC+1, 0x10) // Zero Page address to subtract
//	cpu.writeMemory(0x0010, 0x00)   // Value at Zero Page address to subtract
//	cpu.writeMemory(0x0011, 0x10)   // Value at Zero Page address to subtract
//	cpu.writeMemory(0x1001, 0x10)   // Value at Zero Page address to subtract
//	cpu.Y = 0x01                    // Set Y register for Zero Page addressing
//	cpu.A = 0x20                    // Set accumulator for subtraction
//	cpu.setCarryFlag()              // Ensure the carry flag is set before subtraction
//
//	cpu.cpuQuit = true // Stop the CPU after one execution cycle
//	cpu.startCPU()     // Initialize the CPU state
//
//	expectedValue := cpu.A - 0x10 - (1 - cpu.getSRBit(0)) // Adjusted expected value after subtraction
//
//	// Check if accumulator has the expected value
//	if cpu.A != expectedValue {
//		t.Errorf("SBC Indirect Y failed: got %02X, want %02X", cpu.A, expectedValue)
//	}
//
//	// Check if the zero flag is set correctly
//	if expectedValue == 0 && cpu.getSRBit(1) != 1 {
//		t.Errorf("SBC Indirect Y failed: Zero flag not set correctly")
//	}
//
//	// Check if the negative flag is set correctly
//	if (expectedValue&80) != 0 && cpu.getSRBit(7) != 1 {
//		t.Errorf("SBC Indirect Y failed: Negative flag not set correctly")
//	}
//
//	// Check if the carry flag is set correctly (no borrow)
//	if expectedValue >= 0 && cpu.getSRBit(0) != 1 {
//		t.Errorf("SBC Indirect Y failed: Carry flag not set correctly")
//	}
//
//	// Check if Program Counter is incremented correctly
//	if cpu.PC != cpu.preOpPC+2 { // 2 bytes for SBC Zero Page
//		t.Errorf("SBC Indirect Y failed: expected PC = %04X, got %04X", cpu.preOpPC+2, cpu.PC)
//	}
//}

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
