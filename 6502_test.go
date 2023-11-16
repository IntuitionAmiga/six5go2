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
