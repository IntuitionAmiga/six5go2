package main

import (
	"fmt"
	"os"
)

// Instructions with multiple addressing modes
func LDA(addressingMode string) {
	setFlags := func() {
		// If A is zero, set the SR Zero flag to 1 else set SR Zero flag to 0
		if cpu.A == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		// If bit 7 of accumulator is 1, set the SR negative flag to 1 else set the SR negative flag to 0
		//if getABit(7) == 1 {
		if cpu.A&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE: // Immediate
		cpu.A = cpu.operand1()
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE: // Zero Page
		// Get address
		address := cpu.operand1()
		// Get value from memory at address
		value := readMemory(uint16(address))
		// Set accumulator to value
		cpu.A = value
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX: // Zero Page, X
		// Get address
		address := cpu.operand1() + cpu.X
		value := readMemory(uint16(address))
		// Set accumulator to value
		cpu.A = value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE: // Absolute
		// Get 16 bit address from operand 1 and operand 2
		address := int(cpu.operand2())<<8 | int(cpu.operand1())
		value := readMemory(uint16(address))
		// Set accumulator to value
		cpu.A = value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX: // Absolute, X
		// Get the 16bit X indexed absolute memory address
		address := (int(cpu.operand2())<<8 | int(cpu.operand1())) + int(cpu.X)
		value := readMemory(uint16(address))
		// Set accumulator to value
		cpu.A = value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEY: // Absolute, Y
		// Get 16 bit address from operand 1 and operand 2
		address := (int(cpu.operand2())<<8 | int(cpu.operand1())) + int(cpu.Y)
		value := readMemory(uint16(address))
		// Set accumulator to value
		cpu.A = value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case INDIRECTX: // Indirect, X
		// Get the 16bit X indexed zero page indirect address
		indirectAddress := uint16(int(cpu.operand1()) + int(cpu.X)&0xFF)
		// Get the value at the indirect address
		indirectValue := readMemory(indirectAddress)
		// Get the value at the indirect address + 1
		indirectValue2 := readMemory(indirectAddress + 1)
		// Corrected line: Combine the two values to get the address
		indirectAddress = uint16(int(indirectValue2)<<8 + int(indirectValue))
		// Get the value at the address
		value := readMemory(indirectAddress)
		// Set the accumulator to the value
		cpu.A = value
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case INDIRECTY:
		zeroPageAddress := cpu.operand1()
		address := uint16(readMemory(uint16(zeroPageAddress))) | uint16(readMemory((uint16(zeroPageAddress)+1)&0xFF))<<8
		finalAddress := (address + uint16(cpu.Y)) & 0xFFFF
		value := readMemory(finalAddress)
		cpu.A = value
		setFlags()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	}
}
func LDX(addressingMode string) {
	setFlags := func() {
		// If bit 7 of X is set, set the SR negative flag else reset it to 0
		if cpu.getXBit(7) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		// If X is zero, set the SR zero flag else reset it to 0
		if cpu.X == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE: // Immediate
		// Load the value of the cpu.operand1() into the X register.
		cpu.X = cpu.operand1()
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE: // Zero Page
		// Get address
		address := cpu.operand1()
		value := readMemory(uint16(address))
		// Load the value at the address into X
		cpu.X = value
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEY: // Zero Page, Y
		// Get Y indexed Zero Page address
		address := cpu.operand1() + cpu.Y
		value := readMemory(uint16(address))
		// Load the X register with the Y indexed value in the operand
		cpu.X = value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE: // Absolute
		// Get 16 bit address from operands
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		value := readMemory(address)
		// Update X with the value stored at the address in the operands
		cpu.X = value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEY: // Absolute, Y
		// Get 16 bit Y indexed address from operands
		address := int(cpu.operand2())<<8 | int(cpu.operand1()) + int(cpu.Y)
		value := readMemory(uint16(address))
		cpu.X = value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	}
}
func LDY(addressingMode string) {
	setFlags := func() {
		// If bit 7 of Y is set, set the SR negative flag else reset it to 0
		if cpu.getYBit(7) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		// If Y is zero, set the SR zero flag else reset it to 0
		if cpu.Y == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE: // Immediate
		// Load the value of the cpu.operand1() into the Y register.
		cpu.Y = cpu.operand1()
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE: // Zero Page
		// Get address
		address := cpu.operand1()
		value := readMemory(uint16(address))
		// Load the value at the address into Y
		cpu.Y = value
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX: // Zero Page, X
		// Get the X indexed address
		address := cpu.operand1() + cpu.X
		value := readMemory(uint16(address))
		// Load the Y register with the X indexed value in the operand
		cpu.Y = value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE: // Absolute
		// Get 16 bit address from operands
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		value := readMemory(address)
		// Update Y with the value stored at the address in the operands
		cpu.Y = value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX: // Absolute, X
		// Get the 16bit X indexed absolute memory address
		address := (int(cpu.operand2())<<8 | int(cpu.operand1())) + int(cpu.X)
		value := readMemory(uint16(address))
		// Update Y with the value stored at the address
		cpu.Y = value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	}
}
func STA(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE:
		address := cpu.operand1()
		writeMemory(uint16(address), cpu.A)
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX:
		address := (cpu.operand1() + cpu.X) & 0xFF // Ensure wraparound in Zero Page
		writeMemory(uint16(address), cpu.A)
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		writeMemory(address, cpu.A)
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX:
		address := (uint16(cpu.operand2())<<8 | uint16(cpu.operand1())) + uint16(cpu.X)
		writeMemory(address, cpu.A)
		cpu.updateCycleCounter(5)
		cpu.handleState(3)
	case ABSOLUTEY:
		address := (uint16(cpu.operand2())<<8 | uint16(cpu.operand1())) + uint16(cpu.Y)
		writeMemory(address, cpu.A)
		cpu.updateCycleCounter(5)
		cpu.handleState(3)
	case INDIRECTX:
		zeroPageAddress := (cpu.operand1() + cpu.X) & 0xFF
		address := uint16(readMemory(uint16(zeroPageAddress))) | uint16(readMemory((uint16(zeroPageAddress)+1)&0xFF))<<8
		writeMemory(address, cpu.A)
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case INDIRECTY:
		zeroPageAddress := cpu.operand1()
		address := uint16(readMemory(uint16(zeroPageAddress))) | uint16(readMemory((uint16(zeroPageAddress)+1)&0xFF))<<8
		finalAddress := (address + uint16(cpu.Y)) & 0xFFFF
		writeMemory(finalAddress, cpu.A)
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	}
}
func STX(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE:
		address := cpu.operand1()
		writeMemory(uint16(address), cpu.X)
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEY:
		address := (cpu.operand1() + cpu.Y) & 0xFF
		writeMemory(uint16(address), cpu.X)
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		writeMemory(address, cpu.X)
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	}
}
func STY(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE:
		address := cpu.operand1()
		writeMemory(uint16(address), cpu.Y)
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX:
		address := (cpu.operand1() + cpu.X) & 0xFF
		writeMemory(uint16(address), cpu.Y)
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		writeMemory(address, cpu.Y)
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	}
}
func CMP(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// Subtract the value from the accumulator
		result = cpu.A - value
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		if readBit(7, result) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		// If the value is less than or equal to the accumulator, set the carry flag, else reset it
		if value <= cpu.A {
			cpu.setCarryFlag()
		} else {
			cpu.unsetCarryFlag()
		}
		if addressingMode == IMMEDIATE || addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX || addressingMode == INDIRECTX || addressingMode == INDIRECTY {
			cpu.updateCycleCounter(2)
			cpu.handleState(2)
		} else {
			cpu.updateCycleCounter(3)
			cpu.handleState(3)
		}
	}
	switch addressingMode {
	case IMMEDIATE: // Immediate
		// Get value from cpu.operand1()
		value = cpu.operand1()
		setFlags()
	case ZEROPAGE: // Zero Page
		// Get address
		address := cpu.operand1()
		// Subtract the operand from the accumulator
		value = readMemory(uint16(address))
		setFlags()
	case ZEROPAGEX: // Zero Page, X
		// Get address
		address := cpu.operand1() + cpu.X
		// Get value at address
		value = readMemory(uint16(address))
		setFlags()
	case ABSOLUTE: // Absolute
		// Get 16bit absolute address
		address := int(cpu.operand2())<<8 | int(cpu.operand1())
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case ABSOLUTEX: // Absolute, X
		// Get address
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1()) + uint16(cpu.X)
		// Get value at address
		value = readMemory(address)
		setFlags()
	case ABSOLUTEY: // Absolute, Y
		// Get address
		address := int(cpu.operand2())<<8 | int(cpu.operand1()) + int(cpu.Y)
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case INDIRECTX: // Indirect, X
		// Get the address of the operand
		address := int(cpu.operand1()) + int(cpu.X)
		// Get the value of the operand
		value = readMemory(uint16(address))
		setFlags()
	case INDIRECTY: // Indirect, Y
		// Get the address of the operand
		zeroPageAddress := cpu.operand1()
		address := uint16(readMemory(uint16(zeroPageAddress))) | uint16(readMemory((uint16(zeroPageAddress)+1)&0xFF))<<8
		finalAddress := (address + uint16(cpu.Y)) & 0xFFFF
		value = readMemory(finalAddress)
		setFlags()
	}
}
func JMP(addressingMode string) {
	cpu.previousPC = cpu.PC
	cpu.previousOpcode = cpu.opcode()
	cpu.previousOperand1 = cpu.operand1()
	cpu.previousOperand2 = cpu.operand2()
	cpu.updateCycleCounter(3)
	cpu.handleState(0)
	switch addressingMode {
	case ABSOLUTE:
		// Get the 16 bit address from operands
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Set the program counter to the absolute address
		setPC(address)
	case INDIRECT:
		// Get the 16 bit address from operands
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Handle 6502 page boundary bug
		loByteAddress := address
		//hiByteAddress := (address & 0xFF00) | ((address + 1) & 0xFF) // Ensure it wraps within the page
		hiByteAddress := (address & 0xFF00) | (address & 0x00FF) + 1
		//indirectAddress := uint16(readMemory(hiByteAddress))<<8 | uint16(readMemory(loByteAddress))
		indirectAddress := uint16(readMemory(loByteAddress)) | uint16(readMemory(hiByteAddress))<<8
		// Set the program counter to the indirect address
		setPC(indirectAddress)
	}
	if *klausd && cpu.PC == KlausDInfiniteLoopAddress {
		if readMemory(0x02) == 0xDE && readMemory(0x03) == 0xB0 {
			//fmt.println("All tests passed!")
			os.Exit(0)
		}
	}
}
func AND(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		if readBit(7, result) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = cpu.operand1()
		// AND the value with the accumulator
		result = cpu.A & value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		// Get the address from the operand
		address := cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// AND the value with the accumulator
		result = cpu.A & value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX:
		// Get address
		address := cpu.operand1() + cpu.X
		// Get value at address
		value = readMemory(uint16(address))
		// AND the value with the accumulator
		result = cpu.A & value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get value at address
		value = readMemory(address)
		// AND the value with the accumulator
		result = cpu.A & value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX:
		// Get address
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1()) + uint16(cpu.X)
		// Get value at address
		value = readMemory(address)
		// AND the value with the accumulator
		result = cpu.A & value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEY:
		// Get the address
		address := int(cpu.operand2())<<8 | int(cpu.operand1()) + int(cpu.Y)
		// Get the value at the address
		value = readMemory(uint16(address))
		// AND the value with the accumulator
		result = cpu.A & value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case INDIRECTX:
		// Get the address
		indirectAddress := int(cpu.operand1()) + int(cpu.X)
		address := int(readMemory(uint16(indirectAddress))) + int(readMemory(uint16(indirectAddress+1)))<<8
		// Get the value from the address
		value = readMemory(uint16(address))
		// AND the value with the accumulator
		result = cpu.A & value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case INDIRECTY:
		// Get the 16bit address
		address := uint16(int(cpu.operand1()))
		// Get the indirect address
		indirectAddress1 := readMemory(address)
		indirectAddress2 := readMemory(address + 1)
		//indirectAddress := uint16(int(indirectAddress1)+int(indirectAddress2)<<8) + uint16(cpu.Y)
		indirectAddress := uint16(int(indirectAddress1)+int(indirectAddress2)<<8) + uint16(cpu.Y)
		// Get the value at the address
		value = readMemory(indirectAddress)
		// AND the value with the accumulator
		result = cpu.A & value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	}
}
func EOR(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		//if readBit(7, result) == 1 {
		if result&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = cpu.operand1()
		// XOR the value with the accumulator
		result = cpu.A ^ value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		// Get the address from the operand
		address := cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// XOR the value with the accumulator
		result = cpu.A ^ value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX:
		// Get address
		address := cpu.operand1() + cpu.X
		// Get value at address
		value = readMemory(uint16(address))
		// XOR the value with the accumulator
		result = cpu.A ^ value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get value at address
		value = readMemory(address)
		// XOR the value with the accumulator
		result = cpu.A ^ value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX:
		// Get address
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1()) + uint16(cpu.X)
		// Get value at address
		value = readMemory(address)
		// XOR the value with the accumulator
		result = cpu.A ^ value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEY:
		// Get the address
		address := int(cpu.operand2())<<8 | int(cpu.operand1()) + int(cpu.Y)
		// Get the value at the address
		value = readMemory(uint16(address))
		// XOR the value with the accumulator
		result = cpu.A ^ value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case INDIRECTX:
		// Get the address
		indirectAddress := int(cpu.operand1()) + int(cpu.X)
		address := int(readMemory(uint16(indirectAddress))) + int(readMemory(uint16(indirectAddress+1)))<<8
		// Get the value from the address
		value = readMemory(uint16(address))
		// XOR the value with the accumulator
		result = cpu.A ^ value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case INDIRECTY:
		// Get the 16bit address
		address := uint16(int(cpu.operand1()))
		// Get the indirect address
		indirectAddress1 := readMemory(address)
		indirectAddress2 := readMemory(address + 1)
		indirectAddress := uint16(int(indirectAddress1)+int(indirectAddress2)<<8) + uint16(cpu.Y)
		// Get the value at the address
		value = readMemory(indirectAddress)
		// XOR the value with the accumulator
		result = cpu.A ^ value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	}
}
func ORA(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		if readBit(7, result) == 1 {
			cpu.setNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = cpu.operand1()
		// OR the value with the accumulator
		result = cpu.A | value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		// Get the address from the operand
		address := cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// OR the value with the accumulator
		result = cpu.A | value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ZEROPAGEX:
		// Get address
		address := cpu.operand1() + cpu.X
		// Get value at address
		value = readMemory(uint16(address))
		// OR the value with the accumulator
		result = cpu.A | value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get value at address
		value = readMemory(address)
		// OR the value with the accumulator
		result = cpu.A | value
		// Set the accumulator to the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEX:
		address := (uint16(cpu.operand1()) + uint16(cpu.X)) | uint16(cpu.operand2())<<8
		value = readMemory(address)
		cpu.A |= value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case ABSOLUTEY:
		address := (uint16(cpu.operand1()) + uint16(cpu.Y)) | uint16(cpu.operand2())<<8
		value = readMemory(address)
		cpu.A |= value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	case INDIRECTX:
		zeroPageAddress := (cpu.operand1() + cpu.X) & 0xFF
		effectiveAddrLo := readMemory(uint16(zeroPageAddress))
		effectiveAddrHi := readMemory(uint16((zeroPageAddress + 1) & 0xFF))
		address := uint16(effectiveAddrHi)<<8 | uint16(effectiveAddrLo)
		value = readMemory(address)
		cpu.A |= value
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case INDIRECTY:
		zeroPageAddress := cpu.operand1()
		effectiveAddrLo := readMemory(uint16(zeroPageAddress))
		effectiveAddrHi := readMemory(uint16((zeroPageAddress + 1) & 0xFF))
		address := (uint16(effectiveAddrHi)<<8 | uint16(effectiveAddrLo)) + uint16(cpu.Y)
		value = readMemory(address)
		cpu.A |= value
		setFlags()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	}
}
func BIT(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// Set Negative flag to bit 7 of the value
		if readBit(7, value) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		// Set Overflow flag to bit 6 of the value
		if readBit(6, value) == 1 {
			cpu.setOverflowFlag()
		} else {
			cpu.unsetOverflowFlag()
		}
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
	}
	switch addressingMode {
	case ZEROPAGE:
		// Get the address from the operand
		address := cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// AND the value with the accumulator
		result = cpu.A & value
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get value at address
		value = readMemory(address)
		// AND the value with the accumulator
		result = cpu.A & value
		setFlags()
		cpu.updateCycleCounter(4)
		cpu.handleState(3)
	}
}
func INC(addressingMode string) {
	var address uint16
	var result byte

	setFlags := func() {
		// Fetch the value from the address
		value := readMemory(address)
		// Increment the value (wrapping around for 8-bit values)
		result = value + 1
		// Write the result back to memory
		writeMemory(address, result)

		// Update status flags
		// If bit 7 of the result is set, set the negative flag
		if result&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
	}
	switch addressingMode {
	case ZEROPAGE:
		// Get the address from the operand
		address = uint16(cpu.operand1())
		setFlags()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	case ZEROPAGEX:
		// Get the address from the operand with X offset
		address = uint16(cpu.operand1() + cpu.X)
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get 16-bit address from operand1 and operand2
		address = uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(3)
	case ABSOLUTEX:
		// Get 16-bit address from operand1 and operand2 with X offset
		address = (uint16(cpu.operand2())<<8 | uint16(cpu.operand1())) + uint16(cpu.X)
		setFlags()
		cpu.updateCycleCounter(7)
		cpu.handleState(3)
	}
}
func DEC(addressingMode string) {
	var address uint16
	var result byte

	setFlags := func() {
		// Fetch the value from the address
		value := readMemory(address)
		// Decrement the value (wrapping around for 8-bit values)
		result = value - 1
		// Write the result back to memory
		writeMemory(address, result)

		// Update status flags
		// If bit 7 of the result is set, set the negative flag
		//if readBit(7, result) == 1 {
		if result&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		// If the result is 0, set the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
	}
	switch addressingMode {
	case ZEROPAGE:
		// Get the address from the operand
		address = uint16(cpu.operand1())
		setFlags()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	case ZEROPAGEX:
		// Get the address from the operand with X offset
		address = uint16(cpu.operand1() + cpu.X)
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get 16-bit address from operand1 and operand2
		address = uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(3)
	case ABSOLUTEX:
		// Get 16-bit address from operand1 and operand2 with X offset
		address = (uint16(cpu.operand2())<<8 | uint16(cpu.operand1())) + uint16(cpu.X)
		setFlags()
		cpu.updateCycleCounter(7)
		cpu.handleState(3)
	}
}
func ADC(addressingMode string) {
	var value byte
	var result int

	setFlags := func() {
		// Binary mode is the default
		tmpResult := int(cpu.A) + int(value)
		if cpu.getSRBit(0) == 1 {
			tmpResult++
		}

		if cpu.getSRBit(3) == 1 { // BCD mode
			temp := (cpu.A & 0x0F) + (value & 0x0F) + cpu.getSRBit(0)
			if temp > 9 {
				temp += 6
			}

			result = int((cpu.A & 0xF0) + (value & 0xF0) + (temp & 0x0F))

			if result > 0x99 {
				result += 0x60
			}

		}

		// Set or unset the C flag
		if tmpResult > 0xFF {
			cpu.setCarryFlag()
		} else {
			cpu.unsetCarryFlag()
		}

		// Handle V (overflow) flag
		if (int(cpu.A)^int(value))&0x80 == 0 && (int(cpu.A)^tmpResult)&0x80 != 0 {
			cpu.setOverflowFlag()
		} else {
			cpu.unsetOverflowFlag()
		}

		result = tmpResult & 0xFF // Store the result in 8 bits

		// Handle N (negative) and Z (zero) flags
		if result&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}

		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}

		cpu.A = byte(result)

		// Your addressing mode cycle counts remain the same
		if addressingMode == IMMEDIATE || addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX || addressingMode == INDIRECTX || addressingMode == INDIRECTY {
			cpu.updateCycleCounter(2)
			cpu.handleState(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX || addressingMode == ABSOLUTEY {
			cpu.updateCycleCounter(3)
			cpu.handleState(3)
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = cpu.operand1()
		setFlags()
	case ZEROPAGE:
		// Get the address from the operand
		address := cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case ZEROPAGEX:
		// Get the address from the operand
		address := cpu.operand1() + cpu.X
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get value at address
		value = readMemory(address)
		setFlags()
	case ABSOLUTEX:
		// Get 16 bit address from operand1 and operand2
		address := (uint16(cpu.operand2())<<8 | uint16(cpu.operand1())) + uint16(cpu.X)
		// Get value at address
		value = readMemory(address)
		setFlags()
	case ABSOLUTEY:
		// Get 16 bit address from operand1 and operand2
		address := (uint16(cpu.operand2())<<8 | uint16(cpu.operand1())) + uint16(cpu.Y)
		// Get value at address
		value = readMemory(address)
		setFlags()
	case INDIRECTX:
		// Get the indirect address from the operand
		indirectAddress := cpu.operand1() + cpu.X
		// Get the address from the indirect address
		address := uint16(readMemory(uint16(indirectAddress+1)))<<8 | uint16(readMemory(uint16(indirectAddress)))
		// Get the value at the address
		value = readMemory(address)
		setFlags()
	case INDIRECTY:
		// Get the indirect address from the operand
		indirectAddress := cpu.operand1()
		// Get the address from the indirect address
		address := uint16(readMemory(uint16(indirectAddress+1)))<<8 | uint16(readMemory(uint16(indirectAddress))) + uint16(cpu.Y)
		// Get the value at the address
		value = readMemory(address)
		setFlags()
	}
}
func SBC(addressingMode string) {
	var value byte
	var result int

	setFlags := func() {
		// Check for BCD mode (Decimal flag set)
		if cpu.getSRBit(3) == 1 { // BCD mode

			temp := (cpu.A & 0x0F) - (value & 0x0F) - (cpu.getSRBit(0) ^ 1)
			if temp < 0 {
				temp -= 6
			}

			result = int((cpu.A & 0xF0) - (value & 0xF0) - (temp & 0x0F))

			if result < 0 {
				result -= 0x60
			}

		} else {
			// Binary mode
			result = int(cpu.A) - int(value)
			if cpu.getSRBit(0) == 0 {
				result--
			}
			if int(cpu.A) >= int(value) {
				cpu.setCarryFlag()
			} else {
				cpu.unsetCarryFlag()
			}
		}

		// Negative, and Zero flag checks remain the same
		if readBit(7, byte(result)) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		if result == 0 {
			cpu.setZeroFlag()
		}
		cpu.A = byte(result)

		// Your addressing mode cycle counts remain the same
		if addressingMode == IMMEDIATE || addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX || addressingMode == INDIRECTX || addressingMode == INDIRECTY {
			cpu.updateCycleCounter(2)
			cpu.handleState(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX || addressingMode == ABSOLUTEY {
			cpu.updateCycleCounter(3)
			cpu.handleState(3)
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = cpu.operand1()
		setFlags()
	case ZEROPAGE:
		// Get the address from the operand
		address := cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case ZEROPAGEX:
		// Get the address from the operand
		address := cpu.operand1() + cpu.X
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get value at address
		value = readMemory(address)
		setFlags()
	case ABSOLUTEX:
		// Get 16 bit address from operand1 and operand2
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1()) + uint16(cpu.X)
		// Get value at address
		value = readMemory(address)
		setFlags()
	case ABSOLUTEY:
		// Get 16 bit address from operand1 and operand2
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1()) + uint16(cpu.Y)
		// Get value at address
		value = readMemory(address)
		setFlags()
	case INDIRECTX:
		// Get the indirect address from the operand
		indirectAddress := cpu.operand1() + cpu.X
		// Get the address from the indirect address
		address := uint16(readMemory(uint16(indirectAddress+1)))<<8 | uint16(readMemory(uint16(indirectAddress)))
		// Get the value at the address
		value = readMemory(address)
		setFlags()
	case INDIRECTY:
		// Get the indirect address from the operand
		indirectAddress := cpu.operand1()
		// Get the address from the indirect address
		address := uint16(readMemory(uint16(indirectAddress+1)))<<8 | uint16(readMemory(uint16(indirectAddress))) + uint16(cpu.Y)
		// Get the value at the address
		value = readMemory(address)
		setFlags()
	}
}
func ROR(addressingMode string) {
	var address, value, result byte
	var address16 uint16

	setFlags := func() {
		// Set bit 7 of result and negative flag to the carry flag
		if cpu.getSRBit(0) == 1 {
			result |= 0x80
			cpu.setNegativeFlag()
		} else {
			result &= 0x7F
			cpu.unsetNegativeFlag()
		}
		// Set carry flag to bit 0 of value
		if readBit(0, value) == 1 {
			cpu.setCarryFlag()
		} else {
			cpu.unsetCarryFlag()
		}
		if cpu.A == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		if addressingMode == ACCUMULATOR {
			// Store the result in the accumulator
			cpu.A = result
			cpu.updateCycleCounter(2)
			cpu.handleState(1)
		}
		if addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX {
			// Store the value back into memory
			writeMemory(uint16(address), result)
			cpu.updateCycleCounter(5)
			cpu.handleState(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX {
			// Store the value back into memory
			writeMemory(address16, result)
			cpu.updateCycleCounter(6)
			cpu.handleState(3)
		}
	}
	switch addressingMode {
	case ACCUMULATOR:
		// Get value from accumulator
		value = cpu.A
		// Rotate right one bit
		result = value >> 1
		setFlags()
	case ZEROPAGE:
		// Get address
		address = cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value right 1 bit
		result = value >> 1
		setFlags()
	case ZEROPAGEX:
		// Get X indexed zero page address
		address = cpu.operand1() + cpu.X
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value right 1 bit
		result = value >> 1
		setFlags()
	case ABSOLUTE:
		// Get 16 bit address from operands
		address16 = uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get the value stored at the address in the operands
		value = readMemory(address16)
		// Shift the value right 1 bit
		result = value >> 1
		setFlags()
	case ABSOLUTEX:
		// Get 16 bit address
		address16 = uint16(cpu.operand2())<<8 | uint16(cpu.operand1()) + uint16(cpu.X)
		// Get value stored at address
		value = readMemory(address16)
		// Shift right the value by 1 bit
		result = value >> 1
		setFlags()
	}
}
func ROL(addressingMode string) {
	var address, value, result byte
	var address16 uint16

	setFlags := func() {
		// Set SR carry flag to bit 7 of value
		if readBit(7, value) == 1 {
			cpu.setCarryFlag()
		} else {
			cpu.unsetCarryFlag()
		}

		// Update the zero flag
		if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}

		// Update the negative flag
		if readBit(7, result) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		if addressingMode == ACCUMULATOR {
			// Store the result in the accumulator
			cpu.A = result
			cpu.updateCycleCounter(2)
			cpu.handleState(1)
		}
		if addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX {
			// Store the value back into memory
			writeMemory(uint16(address), result)
			cpu.updateCycleCounter(5)
			cpu.handleState(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX {
			// Store the value back into memory
			writeMemory(address16, result)
			cpu.updateCycleCounter(6)
			cpu.handleState(3)
		}
	}
	switch addressingMode {
	case ACCUMULATOR:
		// Get the value of the accumulator
		value = cpu.A
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | cpu.getSRBit(0)
		setFlags()
	case ZEROPAGE:
		// Get address
		address = cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | cpu.getSRBit(0)
		setFlags()
	case ZEROPAGEX:
		// Get X indexed zero page address
		address = cpu.operand1() + cpu.X
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | cpu.getSRBit(0)
		setFlags()
	case ABSOLUTE:
		// Get 16 bit address from operands
		address16 = uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get the value stored at the address in the operands
		value = readMemory(address16)
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | cpu.getSRBit(0)
		setFlags()
	case ABSOLUTEX:
		// Get 16bit X indexed absolute memory address
		address16 = uint16(cpu.operand2())<<8 | uint16(cpu.operand1()) + uint16(cpu.X)
		// Get the value stored at the address
		value = readMemory(address16)
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | cpu.getSRBit(0)
		setFlags()
	}
}
func LSR(addressingMode string) {
	var value, result byte

	setFlags := func() {
		// Reset the SR negative flag
		cpu.unsetNegativeFlag()
		// If A is 0 then set SR zero flag else reset it
		// Update the zero flag
		if cpu.A == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
		// If bit 0 of value is 1 then set SR carry flag else reset it
		if readBit(0, value) == 1 {
			cpu.setCarryFlag()
		} else {
			cpu.unsetCarryFlag()
		}
	}

	switch addressingMode {
	case ACCUMULATOR:
		// Get the value of the accumulator
		value = cpu.A
		// Shift the value right 1 bit
		result = value >> 1
		// Store the result back into the accumulator
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(1)
	case ZEROPAGE:
		// Get address
		address := cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value right 1 bit
		result = value >> 1
		// Store the value back into memory
		writeMemory(uint16(address), result)
		setFlags()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	case ZEROPAGEX:
		// Get the X indexed address
		address := cpu.operand1() + cpu.X
		// Get the value at the X indexed address
		value = readMemory(uint16(address))
		// Shift the value right 1 bit
		result = value >> 1
		// Store the shifted value in memory
		writeMemory(uint16(address), result)
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operands
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get the value stored at the address in the operands
		value = readMemory(address)
		// Shift the value right 1 bit
		result = value >> 1
		// Store the shifted value back in memory
		writeMemory(address, result)
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(3)
	case ABSOLUTEX:
		// Get the 16bit X indexed absolute memory address
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		address += uint16(cpu.X)
		// Get the value stored at the address
		value = readMemory(address)
		// Shift the value right 1 bit
		result = value >> 1
		// Store the shifted value back in memory
		writeMemory(address, result)
		setFlags()
		cpu.updateCycleCounter(7)
		cpu.handleState(3)
	}
}
func ASL(addressingMode string) {
	var value, result byte

	setFlags := func() {
		// Update the zero flag
		if cpu.A == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}

		// Update the negative flag
		if readBit(7, cpu.A) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}

		// Set the Carry flag based on the original value's bit 7 before the shift operation
		if readBit(7, value) == 1 {
			cpu.setCarryFlag()
		} else {
			cpu.unsetCarryFlag()
		}
	}
	switch addressingMode {
	case ACCUMULATOR:
		// Set value to accumulator
		value = cpu.A
		// Shift the value left 1 bit
		result = value << 1
		// Update the accumulator with the result
		cpu.A = result
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(1)
	case ZEROPAGE:
		// Get address
		address := cpu.operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value left 1 bit
		result = value << 1
		// Store the value back into memory
		writeMemory(uint16(address), result)
		setFlags()
		cpu.updateCycleCounter(5)
		cpu.handleState(2)
	case ZEROPAGEX:
		// Get the X indexed address
		address := cpu.operand1() + cpu.X
		// Get the value at the X indexed address
		value = readMemory(uint16(address))
		// Shift the value left 1 bit
		result = value << 1
		// Store the shifted value in memory
		writeMemory(uint16(address), result)
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operands
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get the value stored at the address in the operands
		value = readMemory(address)
		// Shift the value left 1 bit
		result = value << 1
		// Store the shifted value back in memory
		writeMemory(address, result)
		setFlags()
		cpu.updateCycleCounter(6)
		cpu.handleState(3)
	case ABSOLUTEX:
		// Get the 16bit X indexed absolute memory address
		address := int(cpu.operand2())<<8 | int(cpu.operand1()) + int(cpu.X)
		// Get the value stored at the address
		value = readMemory(uint16(address))
		// Shift the value left 1 bit
		result = value << 1
		// Store the shifted value back in memory
		writeMemory(uint16(address), result)
		setFlags()
		cpu.updateCycleCounter(7)
		cpu.handleState(3)
	}
}
func CPX(addressingMode string) {
	var value, result byte

	setFlags := func() {
		// if cpu.X >= value then set carry flag bit 0 to 1 set carry flag bit 0 to 0
		if cpu.X >= value {
			cpu.setCarryFlag()
		}
		// If value> X then reset carry flag
		if value > cpu.X {
			cpu.unsetCarryFlag()
		}
		// If bit 7 of result is 1 then set negative flag else unset negative flag
		//if readBit(7, result) == 1 {
		if result&0x80 != 0 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		// If value == cpu.X then set zero flag else unset zero flag
		if value == cpu.X {
			//if result == 0 {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get value from operand1
		value = cpu.operand1()
		// Compare X with value
		result = cpu.X - value
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		// Get address
		address := cpu.operand1()
		// Get value at address
		value = readMemory(uint16(address))
		// Store result of X-memory stored at cpu.operand1() in result variable
		result = cpu.X - value
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get address
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get value at address
		value = readMemory(address)
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(3)
	}
}
func CPY(addressingMode string) {
	var value, result byte

	setFlags := func() {
		// If Y>value then set carry flag to 1 else set carry flag to 0
		if cpu.Y >= value {
			cpu.setCarryFlag()
		}
		// If bit 7 of result is set, set N flag to 1 else reset it
		if readBit(7, result) == 1 {
			cpu.setNegativeFlag()
		} else {
			cpu.unsetNegativeFlag()
		}
		// If Y==value then set Z flag to 1 else reset it
		if cpu.Y == value {
			cpu.setZeroFlag()
		} else {
			cpu.unsetZeroFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get value from operand1
		value = cpu.operand1()
		// Subtract operand from Y
		result = cpu.Y - cpu.operand1()
		setFlags()
		cpu.updateCycleCounter(2)
		cpu.handleState(2)
	case ZEROPAGE:
		// Get address
		address := cpu.operand1()
		// Get value at address
		value = readMemory(uint16(address))
		// Store result of Y-memory stored at cpu.operand1() in result variable
		result = cpu.Y - value
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(2)
	case ABSOLUTE:
		// Get address
		address := uint16(cpu.operand2())<<8 | uint16(cpu.operand1())
		// Get value at address
		value = readMemory(address)
		setFlags()
		cpu.updateCycleCounter(3)
		cpu.handleState(3)
	}
}

// 1 byte instructions with no operands

// Implied addressing mode instructions
/*
	In the implied addressing mode, the address containing the operand is implicitly stated in the operation code of the instruction.

	Bytes: 1
*/
func BRK() {
	/*
		BRK - Break Command
	*/
	BRKtrue = true
	if *klausd {
		//fmt.Printf("Test failed at PC: %04X\t", cpu.PC)
		// print opcode and disassembledInstruction at PC
		fmt.Printf("Opcode: %02X\t", readMemory(cpu.PC))
		fmt.Printf("Disassembled Instruction: %s\n", disassembledInstruction)
	}

	disassembledInstruction = fmt.Sprintf("BRK\t")
	disassembleOpcode()
	cpu.previousPC = cpu.PC
	cpu.previousOpcode = cpu.opcode()
	// Increment PC
	//incPC(1)

	// Decrement SP and Push high byte of (PC+1) onto stack
	cpu.decSP()
	updateStack(byte((cpu.PC + 1) >> 8))

	// Decrement SP and Push low byte of (PC+1) onto stack
	cpu.decSP()
	updateStack(byte((cpu.PC + 1) & 0xFF))

	// Set a modified SR with the B flag for the pushed value
	modifiedSR := cpu.SR | 0x10
	// Decrement SP and Store modified SR on stack
	cpu.setBreakFlag()
	cpu.decSP()
	updateStack(modifiedSR)

	// Decrement SP and Store SR on stack
	cpu.decSP()
	updateStack(cpu.SR)

	// Set SR interrupt disable bit to 1
	cpu.setInterruptFlag()

	// Set PC to interrupt vector address
	setPC((uint16(readMemory(IRQVectorAddressHigh)) << 8) | uint16(readMemory(IRQVectorAddressLow)))
	cpu.updateCycleCounter(7)
	cpu.handleState(0)
}
func CLC() {
	/*
		CLC - Clear Carry Flag
	*/
	disassembledInstruction = fmt.Sprintf("CLC\t")
	disassembleOpcode()
	// Set SR carry flag bit 0 to 0
	cpu.unsetCarryFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func CLD() {
	/*
		CLD - Clear Decimal Mode
	*/
	disassembledInstruction = fmt.Sprintf("CLD\t")
	disassembleOpcode()
	cpu.unsetDecimalFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func CLI() {
	/*
		CLI - Clear Interrupt Disable
	*/
	disassembledInstruction = fmt.Sprintf("CLI\t")
	disassembleOpcode()
	// Set SR interrupt disable bit 2 to 0
	cpu.unsetInterruptFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func CLV() {
	/*
		CLV - Clear Overflow Flag
	*/
	disassembledInstruction = fmt.Sprintf("CLV\t")
	disassembleOpcode()
	// Set SR overflow flag bit 6 to 0
	cpu.unsetOverflowFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func DEX() {
	// DEX - Decrement Index Register X By One
	disassembledInstruction = fmt.Sprintf("DEX\t")
	disassembleOpcode()

	// Decrement the X register by 1
	cpu.X--

	// Update the Negative Flag based on the new value of X
	if cpu.X&0x80 != 0 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}

	// Update the Zero Flag based on the new value of X
	if cpu.X == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func DEY() {
	/*
		DEY - Decrement Index Register Y By One
	*/
	disassembledInstruction = fmt.Sprintf("DEY\t")
	disassembleOpcode()

	// Decrement the  Y register by 1
	cpu.Y--
	// If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
	if cpu.getYBit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	// If Y==0 then set the Zero flag
	if cpu.Y == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func INX() {
	/*
		INX - Increment Index Register X By One
	*/
	disassembledInstruction = fmt.Sprintf("INX\t")
	disassembleOpcode()

	// Increment the X register by 1
	cpu.X++
	// If X register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
	if cpu.getXBit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	// If X register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
	if cpu.X == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func INY() {
	/*
		INY - Increment Index Register Y By One
	*/
	disassembledInstruction = fmt.Sprintf("INY\t")
	disassembleOpcode()

	// Increment the  Y register by 1
	cpu.Y++
	// If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
	if cpu.getYBit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	// If Y register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
	if cpu.Y == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func NOP() {
	/*
		NOP - No Operation
	*/
	disassembledInstruction = fmt.Sprintf("NOP\t")
	disassembleOpcode()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func PHA() {
	/*
		PHA - Push Accumulator On Stack
	*/
	disassembledInstruction = fmt.Sprintf("PHA\t")
	disassembleOpcode()

	// Update memory address pointed to by SP with value stored in accumulator
	updateStack(cpu.A)
	cpu.decSP()
	cpu.updateCycleCounter(3)
	cpu.handleState(1)
}
func PHP() {
	/*
	   PHP - Push Processor Status On Stack
	*/
	disassembledInstruction = fmt.Sprintf("PHP\t")
	disassembleOpcode()

	// Set the break flag and the unused bit before pushing
	cpu.SR |= 1 << 4 // Set break flag
	cpu.SR |= 1 << 5 // Set unused bit

	// Push the SR onto the stack
	updateStack(cpu.SR)

	// Decrement the stack pointer
	cpu.decSP()
	cpu.updateCycleCounter(3)
	cpu.handleState(1)
}
func PLA() {
	/*
	   PLA - Pull Accumulator From Stack
	*/
	disassembledInstruction = fmt.Sprintf("PLA\t")
	disassembleOpcode()

	// Increment the stack pointer first
	cpu.incSP()

	// Ensure all arithmetic is done in uint16
	expectedAddress := SPBaseAddress + cpu.SP

	// Now, update accumulator with value stored in memory address pointed to by SP
	cpu.A = readMemory(expectedAddress)

	// If bit 7 of accumulator is set, set negative SR flag else set negative SR flag to 0
	if cpu.getABit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}

	// If accumulator is 0, set zero SR flag else set zero SR flag to 0
	if cpu.A == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(4)
	cpu.handleState(1)
}
func PLP() {
	/*
		PLP - Pull Processor Status From Stack
	*/
	disassembledInstruction = fmt.Sprintf("PLP\t")
	disassembleOpcode()

	// Update SR with the value stored at the address pointed to by SP
	cpu.SR = readStack()
	cpu.incSP()
	cpu.updateCycleCounter(4)
	cpu.handleState(1)
}
func RTI() {
	/*
	   RTI - Return From Interrupt
	*/

	disassembledInstruction = fmt.Sprintf("RTI\t")
	disassembleOpcode()

	cpu.SR = readStack() & 0xCF
	cpu.incSP()

	// Increment the stack pointer to get low byte of PC
	cpu.incSP()

	// Get low byte of PC
	low := uint16(readStack())

	// Increment the stack pointer to get high byte of PC
	cpu.incSP()

	// Get high byte of PC
	high := uint16(readStack())

	cpu.previousPC = cpu.PC
	cpu.previousOpcode = cpu.opcode()
	cpu.updateCycleCounter(6)
	cpu.handleState(0)
	// Update PC with the value stored in memory at the address pointed to by SP
	setPC((high << 8) | low)
}
func RTS() {
	/*
		RTS - Return From Subroutine
	*/
	disassembledInstruction = fmt.Sprintf("RTS\t")
	disassembleOpcode()
	//Get low byte of new PC
	low := uint16(readStack())
	// Increment the stack pointer
	cpu.incSP()
	//Get high byte of new PC
	high := uint16(readStack())
	cpu.previousPC = cpu.PC
	cpu.previousOpcode = cpu.opcode()
	//Update PC with the value stored in memory at the address pointed to by SP
	setPC((high << 8) | low + 1)
	cpu.updateCycleCounter(6)
	cpu.handleState(0)
}
func SEC() {
	/*
		SEC - Set Carry Flag
	*/
	disassembledInstruction = fmt.Sprintf("SEC\t")
	disassembleOpcode()

	// Set SR carry flag bit 0 to 1
	cpu.setCarryFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func SED() {
	/*
		SED - Set Decimal Mode
	*/
	disassembledInstruction = fmt.Sprintf("SED\t")
	disassembleOpcode()

	// Set SR decimal mode flag to 1
	cpu.setDecimalFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func SEI() {
	/*
		SEI - Set Interrupt Disable
	*/
	disassembledInstruction = fmt.Sprintf("SEI\t")
	disassembleOpcode()

	// Set SR interrupt disable bit 2 to 1
	cpu.setInterruptFlag()
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func TAX() {
	/*
		TAX - Transfer Accumulator To Index X
	*/
	disassembledInstruction = fmt.Sprintf("TAX\t")
	disassembleOpcode()

	// Update X with the value of A
	cpu.X = cpu.A
	// If X register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
	if cpu.getXBit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	// If X register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
	if cpu.X == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func TAY() {
	/*
		TAY - Transfer Accumulator To Index Y
	*/
	disassembledInstruction = fmt.Sprintf("TAY\t")
	disassembleOpcode()

	// Set Y register to the value of the accumulator
	cpu.Y = cpu.A
	// If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
	if cpu.getYBit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	// If Y register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
	if cpu.A == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func TSX() {
	/*
		TSX - Transfer Stack Pointer To Index X
	*/
	disassembledInstruction = fmt.Sprintf("TSX\t")
	disassembleOpcode()

	// Update X with the SP
	cpu.X = byte(cpu.SP)
	// If X register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
	if cpu.getXBit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	// If X register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
	if cpu.X == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func TXA() {
	/*
		TXA - Transfer Index X To Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("TXA\t")
	disassembleOpcode()

	// Set accumulator to value of X register
	cpu.A = cpu.X
	// If bit 7 of accumulator is set, set negative SR flag else set negative SR flag to 0
	if cpu.getABit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	// If accumulator is 0, set zero SR flag else set zero SR flag to 0
	if cpu.A == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func TXS() {
	/*
		TXS - Transfer Index X To Stack Pointer
	*/
	disassembledInstruction = fmt.Sprintf("TXS\t")
	disassembleOpcode()

	// Set stack pointer to value of X register
	cpu.SP = uint16(cpu.X)
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}
func TYA() {
	/*
		TYA - Transfer Index Y To Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("TYA\t")
	disassembleOpcode()

	// Set accumulator to value of Y register
	cpu.A = cpu.Y
	// If bit 7 of accumulator is set, set negative SR flag else set negative SR flag to 0
	if cpu.getABit(7) == 1 {
		cpu.setNegativeFlag()
	} else {
		cpu.unsetNegativeFlag()
	}
	// If accumulator is 0, set zero SR flag else set zero SR flag to 0
	if cpu.A == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.unsetZeroFlag()
	}
	cpu.updateCycleCounter(2)
	cpu.handleState(1)
}

// Accumulator instructions
/*
	A

	This form of addressing is represented with a one byte instruction, implying an operation on the accumulator.

	Bytes: 1
*/
func ASL_A() {
	/*
		ASL - Arithmetic Shift Left
	*/
	disassembledInstruction = fmt.Sprintf("ASL\t")
	disassembleOpcode()
	ASL("accumulator")
}
func LSR_A() {
	/*
		LSR - Logical Shift Right
	*/
	disassembledInstruction = fmt.Sprintf("LSR\t")
	disassembleOpcode()
	LSR("accumulator")
}
func ROL_A() {
	/*
		ROL - Rotate Left
	*/
	disassembledInstruction = fmt.Sprintf("ROL\t")
	disassembleOpcode()
	ROL("accumulator")
}
func ROR_A() {
	/*
		ROR - Rotate Right
	*/
	disassembledInstruction = fmt.Sprintf("ROR\t")
	disassembleOpcode()
	ROR("accumulator")
}

// 2 byte instructions with 1 operand
// Immediate addressing mode instructions
/*
	#$nn

	In immediate addressing, the operand is contained in the second byte of the instruction, with no further memory addressing required.

	Bytes: 2
*/
func ADC_I() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	disassembledInstruction = fmt.Sprintf("ADC #$%02X", cpu.operand1())
	disassembleOpcode()

	ADC("immediate")
}
func AND_I() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("AND #$%02X", cpu.operand1())
	disassembleOpcode()
	AND("immediate")
}
func CMP_I() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("CMP #$%02X", cpu.operand1())
	disassembleOpcode()
	CMP("immediate")
}
func CPX_I() {
	/*
		CPX - Compare Index Register X To Memory
	*/
	disassembledInstruction = fmt.Sprintf("CPX #$%02X", cpu.operand1())
	disassembleOpcode()

	CPX("immediate")
}
func CPY_I() {
	/*
		CPY - Compare Index Register Y To Memory
	*/
	disassembledInstruction = fmt.Sprintf("CPY #$%02X", cpu.operand1())
	disassembleOpcode()
	CPY("immediate")
}
func EOR_I() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("EOR #$%02X", cpu.operand1())
	disassembleOpcode()
	EOR("immediate")
}
func LDA_I() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDA #$%02X", cpu.operand1())
	disassembleOpcode()
	LDA("immediate")
}
func LDX_I() {
	/*
		LDX - Load Index Register X From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDX #$%02X", cpu.operand1())
	disassembleOpcode()
	LDX("immediate")
}
func LDY_I() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDY #$%02X", cpu.operand1())
	disassembleOpcode()
	LDY("immediate")
}
func ORA_I() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("ORA #$%02X", cpu.operand1())
	disassembleOpcode()
	ORA("immediate")
}
func SBC_I() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembledInstruction = fmt.Sprintf("SBC #$%02X", cpu.operand1())
	disassembleOpcode()
	SBC("immediate")
}

// Zero Page addressing mode instructions
/*
	$nn

	The zero page instructions allow for shorter code and execution times by only fetching the second byte of the instruction and assuming a zero low address byte. Careful use of the zero page can result in significant increase in code efficiency.

	Bytes: 2
*/
func ADC_Z() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	disassembledInstruction = fmt.Sprintf("ADC $%02X", cpu.operand1())
	disassembleOpcode()
	ADC("zeropage")
}
func AND_Z() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("AND $%02X", cpu.operand1())
	disassembleOpcode()
	AND("zeropage")
}
func ASL_Z() {
	/*
		ASL - Arithmetic Shift Left
	*/
	disassembledInstruction = fmt.Sprintf("ASL $%02X", cpu.operand1())
	disassembleOpcode()

	ASL("zeropage")
}
func BIT_Z() {
	/*
		BIT - Test Bits in Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("BIT $%02X", cpu.operand1())
	disassembleOpcode()
	BIT("zeropage")
}
func CMP_Z() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("CMP $%02X", cpu.operand1())
	disassembleOpcode()
	CMP("zeropage")
}
func CPX_Z() {
	/*
		CPX - Compare Index Register X To Memory
	*/
	disassembledInstruction = fmt.Sprintf("CPX $%02X", cpu.operand1())
	disassembleOpcode()
	CPX("zeropage")
}
func CPY_Z() {
	/*
		CPY - Compare Index Register Y To Memory
	*/
	disassembledInstruction = fmt.Sprintf("CPY $%02X", cpu.operand1())
	disassembleOpcode()
	CPY("zeropage")
}
func DEC_Z() {
	/*
		DEC - Decrement Memory By One
	*/
	disassembledInstruction = fmt.Sprintf("DEC $%02X", cpu.operand1())
	disassembleOpcode()
	DEC("zeropage")
}
func EOR_Z() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("EOR $%02X", cpu.operand1())
	disassembleOpcode()
	EOR("zeropage")
}
func INC_Z() {
	/*
		INC - Increment Memory By One
	*/
	disassembledInstruction = fmt.Sprintf("INC $%02X", cpu.operand1())
	disassembleOpcode()
	INC("zeropage")
}
func LDA_Z() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDA $%02X", cpu.operand1())
	disassembleOpcode()
	LDA("zeropage")
}
func LDX_Z() {
	/*
		LDX - Load Index Register X From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDX $%02X", cpu.operand1())
	disassembleOpcode()
	LDX("zeropage")
}
func LDY_Z() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDY $%02X", cpu.operand1())
	disassembleOpcode()
	LDY("zeropage")
}
func LSR_Z() {
	/*
		LSR - Logical Shift Right
	*/
	disassembledInstruction = fmt.Sprintf("LSR $%02X", cpu.operand1())
	disassembleOpcode()
	LSR("zeropage")
}
func ORA_Z() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("ORA $%02X", cpu.operand1())
	disassembleOpcode()
	ORA("zeropage")
}
func ROL_Z() {
	/*
		ROL - Rotate Left
	*/
	disassembledInstruction = fmt.Sprintf("ROL $%02X", cpu.operand1())
	disassembleOpcode()
	ROL("zeropage")
}
func ROR_Z() {
	/*
		ROR - Rotate Right
	*/
	disassembledInstruction = fmt.Sprintf("ROR $%02X", cpu.operand1())
	disassembleOpcode()

	ROR("zeropage")
}
func SBC_Z() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembledInstruction = fmt.Sprintf("SBC $%02X", cpu.operand1())
	disassembleOpcode()
	SBC("zeropage")
}
func STA_Z() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembledInstruction = fmt.Sprintf("STA $%02X", cpu.operand1())
	disassembleOpcode()
	STA("zeropage")
}
func STX_Z() {
	/*
		STX - Store Index Register X In Memory
	*/
	disassembledInstruction = fmt.Sprintf("STX $%02X", cpu.operand1())
	disassembleOpcode()
	STX("zeropage")
}
func STY_Z() {
	/*
		STY - Store Index Register Y In Memory
	*/
	disassembledInstruction = fmt.Sprintf("STY $%02X", cpu.operand1())
	disassembleOpcode()
	STY("zeropage")
}

// X Indexed Zero Page addressing mode instructions
/*
	$nn,X

	This form of addressing is used in conjunction with the X index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the low order 8 bits of memory and crossing of page boundaries does not occur.

	Bytes: 2
*/
func ADC_ZX() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	disassembledInstruction = fmt.Sprintf("ADC $%02X,X", cpu.operand1())
	disassembleOpcode()
	ADC("zeropagex")
}
func AND_ZX() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("AND $%02X,X", cpu.operand1())
	disassembleOpcode()
	AND("zeropagex")
}
func ASL_ZX() {
	/*
		ASL - Arithmetic Shift Left
	*/
	disassembledInstruction = fmt.Sprintf("ASL $%02X,X", cpu.operand1())
	disassembleOpcode()
	ASL("zeropagex")
}
func CMP_ZX() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("CMP $%02X,X", cpu.operand1())
	disassembleOpcode()
	CMP("zeropagex")
}
func DEC_ZX() {
	/*
		DEC - Decrement Memory By One
	*/
	disassembledInstruction = fmt.Sprintf("DEC $%02X,X", cpu.operand1())
	disassembleOpcode()
	DEC("zeropagex")
}
func LDA_ZX() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDA $%02X,X", cpu.operand1())
	disassembleOpcode()
	LDA("zeropagex")
}
func LDY_ZX() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDY $%02X,X", cpu.operand1())
	disassembleOpcode()
	LDY("zeropagex")
}
func LSR_ZX() {
	/*
		LSR - Logical Shift Right
	*/
	disassembledInstruction = fmt.Sprintf("LSR $%02X,X", cpu.operand1())
	disassembleOpcode()
	LSR("zeropagex")
}
func ORA_ZX() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("ORA $%02X,X", cpu.operand1())
	disassembleOpcode()
	ORA("zeropagex")
}
func ROL_ZX() {
	/*
		ROL - Rotate Left
	*/
	disassembledInstruction = fmt.Sprintf("ROL $%02X,X", cpu.operand1())
	disassembleOpcode()
	ROL("zeropagex")
}
func ROR_ZX() {
	/*
		ROR - Rotate Right
	*/
	disassembledInstruction = fmt.Sprintf("ROR $%02X,X", cpu.operand1())
	disassembleOpcode()
	ROR("zeropagex")
}
func EOR_ZX() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("EOR $%02X,X", cpu.operand1())
	disassembleOpcode()
	EOR("zeropagex")
}
func INC_ZX() {
	/*
		INC - Increment Memory By One
	*/
	disassembledInstruction = fmt.Sprintf("INC $%02X,X", cpu.operand1())
	disassembleOpcode()
	INC("zeropagex")
}
func SBC_ZX() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembledInstruction = fmt.Sprintf("SBC $%02X,X", cpu.operand1())
	disassembleOpcode()
	SBC("zeropagex")
}
func STA_ZX() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembledInstruction = fmt.Sprintf("STA $%02X,X", cpu.operand1())
	disassembleOpcode()
	STA("zeropagex")
}
func STY_ZX() {
	/*
		STY - Store Index Register Y In Memory
	*/
	disassembledInstruction = fmt.Sprintf("STY $%02X,X", cpu.operand1())
	disassembleOpcode()
	STY("zeropagex")
}

// Y Indexed Zero Page addressing mode instructions
/*
	$nn,Y

	This form of addressing is used in conjunction with the Y index register. The effective address is calculated by adding the second byte to the contents of the index register. Since this is a form of "Zero Page" addressing, the content of the second byte references a location in page zero. Additionally, due to the “Zero Page" addressing nature of this mode, no carry is added to the low order 8 bits of memory and crossing of page boundaries does not occur.

	Bytes: 2
*/
func LDX_ZY() {
	/*
		LDX - Load Index Register X From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDX $%02X,Y", cpu.operand1())
	disassembleOpcode()
	LDX("zeropagey")
}
func STX_ZY() {
	/*
		STX - Store Index Register X In Memory
	*/
	disassembledInstruction = fmt.Sprintf("STX $%02X,Y", cpu.operand1())
	disassembleOpcode()
	STX("zeropagey")
}

// X Indexed Zero Page Indirect addressing mode instructions
/*
	($nn,X)

	In indexed indirect addressing, the second byte of the instruction is added to the contents of the X index register, discarding the carry. The result of this addition points to a memory location on page zero whose contents is the high order eight bits of the effective address. The next memory location in page zero contains the low order eight bits of the effective address. Both memory locations specifying the low and high order bytes of the effective address must be in page zero.

	Bytes: 2
*/
func ADC_IX() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	disassembledInstruction = fmt.Sprintf("ADC ($%02X,X)", cpu.operand1())
	disassembleOpcode()
	ADC("indirectx")
}
func AND_IX() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("AND ($%02X,X)", cpu.operand1())
	disassembleOpcode()
	AND("indirectx")
}
func CMP_IX() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("CMP ($%02X,X)", cpu.operand1())
	disassembleOpcode()
	CMP("indirectx")
}
func EOR_IX() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("EOR ($%02X,X)", cpu.operand1())
	disassembleOpcode()
	EOR("indirectx")
}
func LDA_IX() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDA ($%02X,X)", cpu.operand1())
	disassembleOpcode()
	LDA("indirectx")
}
func ORA_IX() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("ORA ($%02X,X)", cpu.operand1())
	disassembleOpcode()
	ORA("indirectx")
}
func SBC_IX() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembledInstruction = fmt.Sprintf("SBC ($%02X,X)", cpu.operand1())
	disassembleOpcode()
	SBC("indirectx")
}
func STA_IX() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembledInstruction = fmt.Sprintf("STA ($%02X,X)", cpu.operand1())
	disassembleOpcode()
	STA("indirectx")
}

// Zero Page Indirect Y Indexed addressing mode instructions
/*
	($nn),Y

	In indirect indexed addressing, the second byte of the instruction points to a memory location in page zero. The contents of this memory location is added to the contents of the Y index register, the result being the high order eight bits of the effective address. The carry from this addition is added to the contents of the next page zero memory location, the result being the low order eight bits of the effective address.

	Bytes: 2
*/
func ADC_IY() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	disassembledInstruction = fmt.Sprintf("ADC ($%02X),Y", cpu.operand1())
	disassembleOpcode()
	ADC("indirecty")
}
func AND_IY() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("AND ($%02X),Y", cpu.operand1())
	disassembleOpcode()
	AND("indirecty")
}
func CMP_IY() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("CMP ($%02X),Y", cpu.operand1())
	disassembleOpcode()
	CMP("indirecty")
}
func EOR_IY() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("EOR ($%02X),Y", cpu.operand1())
	disassembleOpcode()
	EOR("indirecty")
}
func LDA_IY() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDA ($%02X),Y", cpu.operand1())
	disassembleOpcode()
	LDA("indirecty")
}
func ORA_IY() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("ORA ($%02X),Y", cpu.operand1())
	disassembleOpcode()
	ORA("indirecty")
}
func SBC_IY() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembledInstruction = fmt.Sprintf("SBC ($%02X),Y", cpu.operand1())
	disassembleOpcode()
	SBC("indirecty")
}
func STA_IY() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembledInstruction = fmt.Sprintf("STA ($%02X),Y", cpu.operand1())
	disassembleOpcode()
	STA("indirecty")
}

// Relative addressing mode instructions
/*
	$nnnn

	Relative addressing is used only with branch instructions and establishes a destination for the conditional branch.

	The second byte of-the instruction becomes the operand which is an “Offset" added to the contents of the lower eight bits of the program counter when the counter is set at the next instruction. The range of the offset is —128 to +127 bytes from the next instruction.

	Bytes: 2
*/
func BPL_R() {
	/*
		BPL - Branch on Result Plus
	*/
	disassembledInstruction = fmt.Sprintf("BPL $%02X", (cpu.PC+2+uint16(cpu.operand1()))&0xFF)
	disassembleOpcode()
	offset := cpu.operand1()
	signedOffset := int8(offset)
	// Calculate the branch target address
	targetAddress := cpu.PC + 2 + uint16(signedOffset)
	// If N flag is not set, branch to address
	if cpu.getSRBit(7) == 0 {
		// Branch
		setPC(targetAddress)
		cpu.updateCycleCounter(1)
		//cpu.handleState(0)
		instructionCounter++
	} else {
		// Don't branch
		// Increment the instruction counter by 2
		cpu.updateCycleCounter(1)
		cpu.handleState(2)
	}
}
func BMI_R() {
	/*
		BMI - Branch on Result Minus
	*/
	disassembledInstruction = fmt.Sprintf("BMI $%02X", (cpu.PC+2+uint16(cpu.operand1()))&0xFFFF)
	disassembleOpcode()
	// Get offset from operand
	offset := int8(cpu.operand1())
	// If N flag is set, branch to address
	if cpu.getSRBit(7) == 1 {
		// Branch
		// Add offset to PC (already incremented by 2)
		setPC(cpu.PC + 2 + uint16(offset))
	} else {
		// Don't branch
		setPC(cpu.PC + 2)
	}
}
func BVC_R() {
	/*
		BVC - Branch on Overflow Clear
	*/
	disassembledInstruction = fmt.Sprintf("BVC $%02X", cpu.PC+2+uint16(cpu.operand1()))
	disassembleOpcode()
	// Get offset from operand
	offset := cpu.operand1()
	// If overflow flag is not set, branch to address
	if cpu.getSRBit(6) == 0 {
		cpu.updateCycleCounter(1)
		cpu.handleState(0)
		// Branch
		// Add offset to lower 8bits of PC
		setPC(cpu.PC + 3 + uint16(offset)&0xFF)
		// If the offset is negative, decrement the PC by 1
		// If bit 7 is unset then it's negative
		if readBit(7, offset) == 0 {
			decPC(1)
		}
	} else {
		// Don't branch
		cpu.updateCycleCounter(1)
		cpu.handleState(2)
	}
}
func BVS_R() {
	/*
		BVS - Branch on Overflow Set
	*/
	disassembledInstruction = fmt.Sprintf("BVS $%02X", cpu.PC+2+uint16(cpu.operand1()))
	disassembleOpcode()
	// Get offset from operand
	offset := cpu.operand1()
	// If overflow flag is set, branch to address
	if cpu.getSRBit(6) == 1 {
		cpu.updateCycleCounter(1)
		cpu.handleState(0)
		// Branch
		// Add offset to lower 8bits of PC
		setPC(cpu.PC + 3 + uint16(offset)&0xFF)
		// If the offset is negative, decrement the PC by 1
		// If bit 7 is unset then it's negative
		if readBit(7, offset) == 0 {
			decPC(1)
		}
	} else {
		// Don't branch
		cpu.updateCycleCounter(1)
		cpu.handleState(2)
	}
}
func BCC_R() {
	/*
		BCC - Branch on Carry Clear
	*/
	disassembledInstruction = fmt.Sprintf("BCC $%02X", cpu.PC+2+uint16(cpu.operand1()))
	disassembleOpcode()
	cpu.previousPC = cpu.PC
	cpu.previousOpcode = cpu.opcode()
	cpu.previousOperand1 = cpu.operand1()
	// Get offset from operand
	offset := int8(cpu.operand1())
	target := cpu.PC + 2 + uint16(offset)
	if cpu.getSRBit(0) == 0 {
		setPC(target)
	} else {
		// Don't branch
		cpu.updateCycleCounter(1)
		cpu.handleState(2)
	}
}
func BCS_R() {
	/*
		BCS - Branch on Carry Set
	*/
	disassembledInstruction = fmt.Sprintf("BCS $%02X", (cpu.PC+2+uint16(cpu.operand1()))&0xFF)
	disassembleOpcode()
	// Get offset from operand
	offset := cpu.operand1()
	// If carry flag is set, branch to address
	if cpu.getSRBit(0) == 1 {
		cpu.updateCycleCounter(1)
		cpu.handleState(0)
		// Branch
		// Add offset to lower 8bits of PC
		setPC(cpu.PC + 3 + uint16(offset)&0xFF)
		// If the offset is negative, decrement the PC by 1
		// If bit 7 is unset then it's negative
		if readBit(7, offset) == 0 {
			decPC(1)
		}
	} else {
		// Don't branch
		cpu.updateCycleCounter(1)
		cpu.handleState(2)
	}
}
func BNE_R() {
	/*
		BNE - Branch on Result Not Zero
	*/
	disassembledInstruction = fmt.Sprintf("BNE $%04X", cpu.PC+2+uint16(cpu.operand1()))
	disassembleOpcode()
	// Fetch offset from operand
	offset := int8(cpu.operand1()) // Cast to signed 8-bit
	// Check Z flag to determine if branching is needed
	if cpu.getSRBit(1) == 0 {
		// Calculate the branch target address
		targetAddr := cpu.PC + 2 + uint16(offset)
		// Check if the branch crosses a page boundary
		if (cpu.PC & 0xFF00) != (targetAddr & 0xFF00) {
			cpu.updateCycleCounter(1)
			cpu.handleState(2)
		} else {
			cpu.updateCycleCounter(1)
			cpu.handleState(1)
		}
		// Update the program counter
		setPC(targetAddr & 0xFFFF)
	} else {
		// If Z flag is set, don't branch
		cpu.updateCycleCounter(1)
		cpu.handleState(2)
	}
}
func BEQ_R() {
	/*
	   BEQ - Branch on Result Zero
	*/
	disassembledInstruction = fmt.Sprintf("BEQ $%04X", cpu.PC+2+uint16(cpu.operand1()))
	disassembleOpcode()

	// Get offset from address in operand
	offset := int8(cpu.operand1()) // Cast to signed 8-bit integer to handle negative offsets

	// If Z flag is set, branch to address
	if cpu.getSRBit(1) == 1 {
		cpu.updateCycleCounter(1)
		cpu.handleState(0)
		// Add 2 to PC (1 for opcode, 1 for operand) and then add offset
		setPC(cpu.PC + 2 + uint16(offset))
	} else {
		// Don't branch
		cpu.updateCycleCounter(1)
		cpu.handleState(2)
	}
}

// 3 byte instructions with 2 operands
// Absolute addressing mode instructions
/*
	$nnnn

	In absolute addressing, the second byte of the instruction specifies the eight high order bits of the effective address while the third byte specifies the eight low order bits. Thus, the absolute addressing mode allows access to the entire 65 K bytes of addressable memory.

	Bytes: 3
*/
func ADC_ABS() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	disassembledInstruction = fmt.Sprintf("ADC $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ADC("absolute")
}
func AND_ABS() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("AND $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	AND("absolute")
}
func ASL_ABS() {
	/*
		ASL - Arithmetic Shift Left
	*/
	disassembledInstruction = fmt.Sprintf("ASL $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ASL("absolute")
}
func BIT_ABS() {
	/*
		BIT - Test Bits in Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("BIT $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	BIT("absolute")
}
func CMP_ABS() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("CMP $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	CMP("absolute")
}
func CPX_ABS() {
	/*
		CPX - Compare Index Register X To Memory
	*/
	disassembledInstruction = fmt.Sprintf("CPX $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	CPX("absolute")
}
func CPY_ABS() {
	/*
		CPY - Compare Index Register Y To Memory
	*/
	disassembledInstruction = fmt.Sprintf("CPY $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	CPY("absolute")
}
func DEC_ABS() {
	/*
		DEC - Decrement Memory By One
	*/
	disassembledInstruction = fmt.Sprintf("DEC $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	DEC("absolute")
}
func EOR_ABS() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("EOR $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	EOR("absolute")
}
func INC_ABS() {
	/*
		INC - Increment Memory By One
	*/
	disassembledInstruction = fmt.Sprintf("INC $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	INC("absolute")
}
func JMP_ABS() {
	/*
		JMP - JMP Absolute
	*/
	disassembledInstruction = fmt.Sprintf("JMP $%04X", int(cpu.operand2())<<8|int(cpu.operand1()))
	disassembleOpcode()
	// For AllSuiteA.bin 6502 opcode test suite
	if *allsuitea && readMemory(0x210) == 0xFF {
		fmt.Printf("\n\u001B[32;5mMemory address $210 == $%02X. All opcodes succesfully tested and passed!\u001B[0m\n", readMemory(0x210))
		os.Exit(0)
	}
	JMP("absolute")
}
func JSR_ABS() {
	/*
		JSR - Jump To Subroutine
	*/
	disassembledInstruction = fmt.Sprintf("JSR $%04X", int(cpu.operand2())<<8|int(cpu.operand1()))
	disassembleOpcode()
	// First, push the high byte
	cpu.decSP()
	updateStack(byte(cpu.PC >> 8))
	cpu.decSP()
	updateStack(byte((cpu.PC)&0xFF) + 2)

	cpu.previousPC = cpu.PC
	cpu.previousOpcode = cpu.opcode()
	cpu.previousOperand1 = cpu.operand1()
	cpu.previousOperand2 = cpu.operand2()
	// Now, jump to the subroutine address specified by the operands
	setPC(uint16(cpu.operand2())<<8 | uint16(cpu.operand1()))
	cpu.updateCycleCounter(1)
	cpu.handleState(0)
}
func LDA_ABS() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDA $%04X", uint16(cpu.operand2())<<8|uint16(cpu.operand1()))
	disassembleOpcode()
	LDA("absolute")
}
func LDX_ABS() {
	/*
		LDX - Load Index Register X From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDX $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	LDX("absolute")
}
func LDY_ABS() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDY $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	LDY("absolute")
}
func LSR_ABS() {
	/*
		LSR - Logical Shift Right
	*/
	disassembledInstruction = fmt.Sprintf("LSR $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	LSR("absolute")
}
func ORA_ABS() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("ORA $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ORA("absolute")
}
func ROL_ABS() {
	/*
		ROL - Rotate Left
	*/
	disassembledInstruction = fmt.Sprintf("ROL $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ROL("absolute")
}
func ROR_ABS() {
	/*
		ROR - Rotate Right
	*/
	disassembledInstruction = fmt.Sprintf("ROR $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ROR("absolute")
}
func SBC_ABS() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembledInstruction = fmt.Sprintf("SBC $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	SBC("absolute")
}
func STA_ABS() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembledInstruction = fmt.Sprintf("STA $%04X", uint16(cpu.operand2())<<8|uint16(cpu.operand1()))
	disassembleOpcode()
	STA("absolute")
}
func STX_ABS() {
	/*
		STX - Store Index Register X In Memory
	*/
	disassembledInstruction = fmt.Sprintf("STX $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	STX("absolute")
}
func STY_ABS() {
	/*
		STY - Store Index Register Y In Memory
	*/
	disassembledInstruction = fmt.Sprintf("STY $%02X%02X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	STY("absolute")
}

// X Indexed Absolute addressing mode instructions
/*
	$nnnn,X

	This form of addressing is used in conjunction with the X index register. The effective address is formed by adding the contents of X to the address contained in the second and third bytes of the instruction. This mode allows the index register to contain the index or count value and the instruction to contain the base address. This type of indexing allows any location referencing and the index to modify multiple fields resulting in reduced coding and execution time.

	Note on the MOS 6502:

	The value at the specified address, ignoring the the addressing mode's X offset, is read (and discarded) before the final address is read. This may cause side effects in I/O registers.


	Bytes: 3
*/
func ADC_ABX() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	disassembledInstruction = fmt.Sprintf("ADC $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ADC("absolutex")
}
func AND_ABX() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("AND $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	AND("absolutex")
}
func ASL_ABX() {
	/*
		ASL - Arithmetic Shift Left
	*/
	disassembledInstruction = fmt.Sprintf("ASL $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ASL("absolutex")
}
func CMP_ABX() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("CMP $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	CMP("absolutex")
}
func DEC_ABX() {
	/*
		DEC - Decrement Memory By One
	*/
	disassembledInstruction = fmt.Sprintf("DEC $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	DEC("absolutex")
}
func EOR_ABX() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("EOR $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	EOR("absolutex")
}
func INC_ABX() {
	/*
		INC - Increment Memory By One
	*/
	disassembledInstruction = fmt.Sprintf("INC $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	INC("absolutex")
}
func LDA_ABX() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDA $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	LDA("absolutex")
}
func LDY_ABX() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDY $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	LDY("absolutex")
}
func LSR_ABX() {
	/*
		LSR - Logical Shift Right
	*/
	disassembledInstruction = fmt.Sprintf("LSR $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	LSR("absolutex")
}
func ORA_ABX() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("ORA $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ORA("absolutex")
}
func ROL_ABX() {
	/*
	 */
	disassembledInstruction = fmt.Sprintf("ROL $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ROL("absolutex")
}
func ROR_ABX() {
	/*
		ROR - Rotate Right
	*/
	disassembledInstruction = fmt.Sprintf("ROR $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ROR("absolutex")
}
func SBC_ABX() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembledInstruction = fmt.Sprintf("SBC $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	SBC("absolutex")
}
func STA_ABX() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembledInstruction = fmt.Sprintf("STA $%02X%02X,X", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	STA("absolutex")
}

// Y Indexed Absolute addressing mode instructions
/*
	$nnnn,Y

	This form of addressing is used in conjunction with the Y index register. The effective address is formed by adding the contents of Y to the address contained in the second and third bytes of the instruction. This mode allows the index register to contain the index or count value and the instruction to contain the base address. This type of indexing allows any location referencing and the index to modify multiple fields resulting in reduced coding and execution time.

	Note on the MOS 6502:

	The value at the specified address, ignoring the the addressing mode's Y offset, is read (and discarded) before the final address is read. This may cause side effects in I/O registers.

	Bytes: 3
*/
func ADC_ABY() {
	/*
		ADC - Add Memory to Accumulator with Carry
	*/
	disassembledInstruction = fmt.Sprintf("ADC $%04X,Y", int(cpu.operand2())<<8|int(cpu.operand1()))
	disassembleOpcode()
	ADC("absolutey")
}
func AND_ABY() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("AND $%02X%02X,Y", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	AND("absolutey")
}
func CMP_ABY() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("CMP $%02X%02X,Y", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	CMP("absolutey")
}
func EOR_ABY() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("EOR $%02X%02X,Y", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	EOR("absolutey")
}
func LDA_ABY() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDA $%02X%02X,Y", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	LDA("absolutey")
}
func LDX_ABY() {
	/*
		LDX - Load Index Register X From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDX $%02X%02X,Y", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	LDX("absolutey")
}
func ORA_ABY() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("ORA $%02X%02X,Y", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	ORA("absolutey")
}
func SBC_ABY() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembledInstruction = fmt.Sprintf("SBC $%02X%02X,Y", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	SBC("absolutey")
}
func STA_ABY() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembledInstruction = fmt.Sprintf("STA $%02X%02X,Y", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	STA("absolutey")
}

// Absolute Indirect addressing mode instructions
func JMP_IND() {
	/*
		JMP - JMP Indirect
	*/
	disassembledInstruction = fmt.Sprintf("JMP ($%02X%02X)", cpu.operand2(), cpu.operand1())
	disassembleOpcode()
	JMP("indirect")
}
