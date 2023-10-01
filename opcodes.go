package main

import (
	"fmt"
	"os"
)

// Instructions with multiple addressing modes
func LDA(addressingMode string) {
	setFlags := func() {
		// If A is zero, set the SR Zero flag to 1 else set SR Zero flag to 0
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If bit 7 of accumulator is 1, set the SR negative flag to 1 else set the SR negative flag to 0
		//if getABit(7) == 1 {
		if A&0x80 != 0 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE: // Immediate
		A = operand1()
		setFlags()
		updateCycleCounter(2)
		handleState(2)
	case ZEROPAGE: // Zero Page
		// Get address
		address := operand1()
		// Get value from memory at address
		value := readMemory(uint16(address))
		// Set accumulator to value
		A = value
		setFlags()
		updateCycleCounter(3)
		handleState(2)
	case ZEROPAGEX: // Zero Page, X
		// Get address
		address := operand1() + X
		value := readMemory(uint16(address))
		// Set accumulator to value
		A = value
		setFlags()
		updateCycleCounter(4)
		handleState(2)
	case ABSOLUTE: // Absolute
		// Get 16 bit address from operand 1 and operand 2
		address := int(operand2())<<8 | int(operand1())
		value := readMemory(uint16(address))
		// Set accumulator to value
		A = value
		setFlags()
		updateCycleCounter(4)
		handleState(3)
	case ABSOLUTEX: // Absolute, X
		// Get the 16bit X indexed absolute memory address
		address := (int(operand2())<<8 | int(operand1())) + int(X)
		value := readMemory(uint16(address))
		// Set accumulator to value
		A = value
		setFlags()
		updateCycleCounter(4)
		handleState(3)
	case ABSOLUTEY: // Absolute, Y
		// Get 16 bit address from operand 1 and operand 2
		address := (int(operand2())<<8 | int(operand1())) + int(Y)
		value := readMemory(uint16(address))
		// Set accumulator to value
		A = value
		setFlags()
		updateCycleCounter(4)
		handleState(3)
	case INDIRECTX: // Indirect, X
		// Get the 16bit X indexed zero page indirect address
		indirectAddress := uint16(int(operand1()) + int(X)&0xFF)
		// Get the value at the indirect address
		indirectValue := readMemory(indirectAddress)
		// Get the value at the indirect address + 1
		indirectValue2 := readMemory(indirectAddress + 1)
		// Corrected line: Combine the two values to get the address
		indirectAddress = uint16(int(indirectValue2)<<8 + int(indirectValue))
		// Get the value at the address
		value := readMemory(indirectAddress)
		// Set the accumulator to the value
		A = value
		setFlags()
		updateCycleCounter(6)
		handleState(2)
	case INDIRECTY:
		zeroPageAddress := operand1()
		address := uint16(readMemory(uint16(zeroPageAddress))) | uint16(readMemory((uint16(zeroPageAddress)+1)&0xFF))<<8
		finalAddress := (address + uint16(Y)) & 0xFFFF
		value := readMemory(finalAddress)
		A = value
		setFlags()
		updateCycleCounter(5)
		handleState(2)
	}
}
func LDX(addressingMode string) {
	setFlags := func() {
		// If bit 7 of X is set, set the SR negative flag else reset it to 0
		if getXBit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If X is zero, set the SR zero flag else reset it to 0
		if X == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE: // Immediate
		// Load the value of the operand1() into the X register.
		X = operand1()
		setFlags()
		updateCycleCounter(2)
		handleState(2)
	case ZEROPAGE: // Zero Page
		// Get address
		address := operand1()
		value := readMemory(uint16(address))
		// Load the value at the address into X
		X = value
		setFlags()
		updateCycleCounter(3)
		handleState(2)
	case ZEROPAGEY: // Zero Page, Y
		// Get Y indexed Zero Page address
		address := operand1() + Y
		value := readMemory(uint16(address))
		// Load the X register with the Y indexed value in the operand
		X = value
		setFlags()
		updateCycleCounter(4)
		handleState(2)
	case ABSOLUTE: // Absolute
		// Get 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		value := readMemory(address)
		// Update X with the value stored at the address in the operands
		X = value
		setFlags()
		updateCycleCounter(4)
		handleState(3)
	case ABSOLUTEY: // Absolute, Y
		// Get 16 bit Y indexed address from operands
		address := int(operand2())<<8 | int(operand1()) + int(Y)
		value := readMemory(uint16(address))
		X = value
		setFlags()
		updateCycleCounter(4)
		handleState(3)
	}
}
func LDY(addressingMode string) {
	setFlags := func() {
		// If bit 7 of Y is set, set the SR negative flag else reset it to 0
		if getYBit(7) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If Y is zero, set the SR zero flag else reset it to 0
		if Y == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE: // Immediate
		// Load the value of the operand1() into the Y register.
		Y = operand1()
		setFlags()
		updateCycleCounter(2)
		handleState(2)
	case ZEROPAGE: // Zero Page
		// Get address
		address := operand1()
		value := readMemory(uint16(address))
		// Load the value at the address into Y
		Y = value
		setFlags()
		updateCycleCounter(3)
		handleState(2)
	case ZEROPAGEX: // Zero Page, X
		// Get the X indexed address
		address := operand1() + X
		value := readMemory(uint16(address))
		// Load the Y register with the X indexed value in the operand
		Y = value
		setFlags()
		updateCycleCounter(4)
		handleState(2)
	case ABSOLUTE: // Absolute
		// Get 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		value := readMemory(address)
		// Update Y with the value stored at the address in the operands
		Y = value
		setFlags()
		updateCycleCounter(4)
		handleState(3)
	case ABSOLUTEX: // Absolute, X
		// Get the 16bit X indexed absolute memory address
		address := (int(operand2())<<8 | int(operand1())) + int(X)
		value := readMemory(uint16(address))
		// Update Y with the value stored at the address
		Y = value
		setFlags()
		updateCycleCounter(4)
		handleState(3)
	}
}
func STA(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE:
		address := operand1()
		writeMemory(uint16(address), A)
		updateCycleCounter(3)
		handleState(2)
	case ZEROPAGEX:
		address := (operand1() + X) & 0xFF // Ensure wraparound in Zero Page
		writeMemory(uint16(address), A)
		updateCycleCounter(4)
		handleState(2)
	case ABSOLUTE:
		address := uint16(operand2())<<8 | uint16(operand1())
		writeMemory(address, A)
		updateCycleCounter(4)
		handleState(3)
	case ABSOLUTEX:
		address := (uint16(operand2())<<8 | uint16(operand1())) + uint16(X)
		writeMemory(address, A)
		updateCycleCounter(5)
		handleState(3)
	case ABSOLUTEY:
		address := (uint16(operand2())<<8 | uint16(operand1())) + uint16(Y)
		writeMemory(address, A)
		updateCycleCounter(5)
		handleState(3)
	case INDIRECTX:
		zeroPageAddress := (operand1() + X) & 0xFF
		address := uint16(readMemory(uint16(zeroPageAddress))) | uint16(readMemory((uint16(zeroPageAddress)+1)&0xFF))<<8
		writeMemory(address, A)
		updateCycleCounter(6)
		handleState(2)
	case INDIRECTY:
		zeroPageAddress := operand1()
		address := uint16(readMemory(uint16(zeroPageAddress))) | uint16(readMemory((uint16(zeroPageAddress)+1)&0xFF))<<8
		finalAddress := (address + uint16(Y)) & 0xFFFF
		writeMemory(finalAddress, A)
		updateCycleCounter(6)
		handleState(2)
	}
}
func STX(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE:
		address := operand1()
		writeMemory(uint16(address), X)
		updateCycleCounter(3)
		handleState(2)
	case ZEROPAGEY:
		address := (operand1() + Y) & 0xFF
		writeMemory(uint16(address), X)
		updateCycleCounter(4)
		handleState(2)
	case ABSOLUTE:
		address := uint16(operand2())<<8 | uint16(operand1())
		writeMemory(address, X)
		updateCycleCounter(4)
		handleState(3)
	}
}
func STY(addressingMode string) {
	switch addressingMode {
	case ZEROPAGE:
		address := operand1()
		writeMemory(uint16(address), Y)
		updateCycleCounter(3)
		handleState(2)
	case ZEROPAGEX:
		address := (operand1() + X) & 0xFF
		writeMemory(uint16(address), Y)
		updateCycleCounter(4)
		handleState(2)
	case ABSOLUTE:
		address := uint16(operand2())<<8 | uint16(operand1())
		writeMemory(address, Y)
		updateCycleCounter(4)
		handleState(3)
	}
}
func CMP(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// Subtract the value from the accumulator
		result = A - value
		// If the result is 0, set the zero flag
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		if readBit(7, result) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If the value is less than or equal to the accumulator, set the carry flag, else reset it
		if value <= A {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		if addressingMode == IMMEDIATE || addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX || addressingMode == INDIRECTX || addressingMode == INDIRECTY {
			updateCycleCounter(2)
			handleState(2)
		} else {
			updateCycleCounter(3)
			handleState(3)
		}
	}
	switch addressingMode {
	case IMMEDIATE: // Immediate
		// Get value from operand1()
		value = operand1()
		setFlags()
	case ZEROPAGE: // Zero Page
		// Get address
		address := operand1()
		// Subtract the operand from the accumulator
		value = readMemory(uint16(address))
		setFlags()
	case ZEROPAGEX: // Zero Page, X
		// Get address
		address := operand1() + X
		// Get value at address
		value = readMemory(uint16(address))
		setFlags()
	case ABSOLUTE: // Absolute
		// Get 16bit absolute address
		address := int(operand2())<<8 | int(operand1())
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case ABSOLUTEX: // Absolute, X
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Get value at address
		value = readMemory(address)
		setFlags()
	case ABSOLUTEY: // Absolute, Y
		// Get address
		address := int(operand2())<<8 | int(operand1()) + int(Y)
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case INDIRECTX: // Indirect, X
		// Get the address of the operand
		address := int(operand1()) + int(X)
		// Get the value of the operand
		value = readMemory(uint16(address))
		setFlags()
	case INDIRECTY: // Indirect, Y
		// Get the address of the operand
		zeroPageAddress := operand1()
		address := uint16(readMemory(uint16(zeroPageAddress))) | uint16(readMemory((uint16(zeroPageAddress)+1)&0xFF))<<8
		finalAddress := (address + uint16(Y)) & 0xFFFF
		value = readMemory(finalAddress)
		setFlags()
	}
}
func JMP(addressingMode string) {
	previousPC = PC
	previousOpcode = opcode()
	previousOperand1 = operand1()
	previousOperand2 = operand2()
	updateCycleCounter(3)
	handleState(0)
	switch addressingMode {
	case ABSOLUTE:
		// Get the 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		// Set the program counter to the absolute address
		setPC(int(address))
	case INDIRECT:
		// Get the 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		// Handle 6502 page boundary bug
		loByteAddress := address
		hiByteAddress := (address & 0xFF00) | ((address + 1) & 0xFF) // Ensure it wraps within the page
		indirectAddress := uint16(readMemory(hiByteAddress))<<8 | uint16(readMemory(loByteAddress))
		// Set the program counter to the indirect address
		setPC(int(indirectAddress))
	}
	if *klausd && PC == KlausDInfiniteLoopAddress {
		if readMemory(0x02) == 0xDE && readMemory(0x03) == 0xB0 {
			fmt.Println("All tests passed!")
			os.Exit(0)
		}
	}
}
func AND(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// If the result is 0, set the zero flag
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		if readBit(7, result) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = operand1()
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(2)
		handleState(2)
	case ZEROPAGE:
		// Get the address from the operand
		address := operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(3)
		handleState(2)
	case ZEROPAGEX:
		// Get address
		address := operand1() + X
		// Get value at address
		value = readMemory(uint16(address))
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(4)
		handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = readMemory(address)
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(4)
		handleState(3)
	case ABSOLUTEX:
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Get value at address
		value = readMemory(address)
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(4)
		handleState(3)
	case ABSOLUTEY:
		// Get the address
		address := int(operand2())<<8 | int(operand1()) + int(Y)
		// Get the value at the address
		value = readMemory(uint16(address))
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(4)
		handleState(3)
	case INDIRECTX:
		// Get the address
		indirectAddress := int(operand1()) + int(X)
		address := int(readMemory(uint16(indirectAddress))) + int(readMemory(uint16(indirectAddress+1)))<<8
		// Get the value from the address
		value = readMemory(uint16(address))
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(6)
		handleState(2)
	case INDIRECTY:
		// Get the 16bit address
		address := uint16(int(operand1()))
		// Get the indirect address
		indirectAddress1 := readMemory(address)
		indirectAddress2 := readMemory(address + 1)
		//indirectAddress := uint16(int(indirectAddress1)+int(indirectAddress2)<<8) + uint16(Y)
		indirectAddress := uint16(int(indirectAddress1)+int(indirectAddress2)<<8) + uint16(Y)
		// Get the value at the address
		value = readMemory(indirectAddress)
		// AND the value with the accumulator
		result = A & value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(5)
		handleState(2)
	}
}
func EOR(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// If the result is 0, set the zero flag
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		//if readBit(7, result) == 1 {
		if result&0x80 != 0 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = operand1()
		// XOR the value with the accumulator
		result = A ^ value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(2)
		handleState(2)
	case ZEROPAGE:
		// Get the address from the operand
		address := operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// XOR the value with the accumulator
		result = A ^ value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(3)
		handleState(2)
	case ZEROPAGEX:
		// Get address
		address := operand1() + X
		// Get value at address
		value = readMemory(uint16(address))
		// XOR the value with the accumulator
		result = A ^ value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(4)
		handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = readMemory(address)
		// XOR the value with the accumulator
		result = A ^ value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(4)
		handleState(3)
	case ABSOLUTEX:
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Get value at address
		value = readMemory(address)
		// XOR the value with the accumulator
		result = A ^ value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(4)
		handleState(3)
	case ABSOLUTEY:
		// Get the address
		address := int(operand2())<<8 | int(operand1()) + int(Y)
		// Get the value at the address
		value = readMemory(uint16(address))
		// XOR the value with the accumulator
		result = A ^ value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(4)
		handleState(3)
	case INDIRECTX:
		// Get the address
		indirectAddress := int(operand1()) + int(X)
		address := int(readMemory(uint16(indirectAddress))) + int(readMemory(uint16(indirectAddress+1)))<<8
		// Get the value from the address
		value = readMemory(uint16(address))
		// XOR the value with the accumulator
		result = A ^ value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(6)
		handleState(2)
	case INDIRECTY:
		// Get the 16bit address
		address := uint16(int(operand1()))
		// Get the indirect address
		indirectAddress1 := readMemory(address)
		indirectAddress2 := readMemory(address + 1)
		indirectAddress := uint16(int(indirectAddress1)+int(indirectAddress2)<<8) + uint16(Y)
		// Get the value at the address
		value = readMemory(indirectAddress)
		// XOR the value with the accumulator
		result = A ^ value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(5)
		handleState(2)
	}
}
func ORA(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// If the result is 0, set the zero flag
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If bit 7 of the result is set, set the negative flag
		if readBit(7, result) == 1 {
			setNegativeFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = operand1()
		// OR the value with the accumulator
		result = A | value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(2)
		handleState(2)
	case ZEROPAGE:
		// Get the address from the operand
		address := operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// OR the value with the accumulator
		result = A | value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(3)
		handleState(2)
	case ZEROPAGEX:
		// Get address
		address := operand1() + X
		// Get value at address
		value = readMemory(uint16(address))
		// OR the value with the accumulator
		result = A | value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(4)
		handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = readMemory(address)
		// OR the value with the accumulator
		result = A | value
		// Set the accumulator to the result
		A = result
		setFlags()
		updateCycleCounter(4)
		handleState(3)
	case ABSOLUTEX:
		address := (uint16(operand1()) + uint16(X)) | uint16(operand2())<<8
		value = readMemory(address)
		A |= value
		setFlags()
		updateCycleCounter(4)
		handleState(3)
	case ABSOLUTEY:
		address := (uint16(operand1()) + uint16(Y)) | uint16(operand2())<<8
		value = readMemory(address)
		A |= value
		setFlags()
		updateCycleCounter(4)
		handleState(3)
	case INDIRECTX:
		zeroPageAddress := (operand1() + X) & 0xFF
		effectiveAddrLo := readMemory(uint16(zeroPageAddress))
		effectiveAddrHi := readMemory(uint16((zeroPageAddress + 1) & 0xFF))
		address := uint16(effectiveAddrHi)<<8 | uint16(effectiveAddrLo)
		value = readMemory(address)
		A |= value
		setFlags()
		updateCycleCounter(6)
		handleState(2)
	case INDIRECTY:
		zeroPageAddress := operand1()
		effectiveAddrLo := readMemory(uint16(zeroPageAddress))
		effectiveAddrHi := readMemory(uint16((zeroPageAddress + 1) & 0xFF))
		address := (uint16(effectiveAddrHi)<<8 | uint16(effectiveAddrLo)) + uint16(Y)
		value = readMemory(address)
		A |= value
		setFlags()
		updateCycleCounter(5)
		handleState(2)
	}
}
func BIT(addressingMode string) {
	var value, result byte
	setFlags := func() {
		// Set Negative flag to bit 7 of the value
		if readBit(7, value) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// Set Overflow flag to bit 6 of the value
		if readBit(6, value) == 1 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}
		// If the result is 0, set the zero flag
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
	}
	switch addressingMode {
	case ZEROPAGE:
		// Get the address from the operand
		address := operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// AND the value with the accumulator
		result = A & value
		setFlags()
		updateCycleCounter(3)
		handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = readMemory(address)
		// AND the value with the accumulator
		result = A & value
		setFlags()
		updateCycleCounter(4)
		handleState(3)
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
		if readBit(7, result) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If the result is 0, set the zero flag
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
	}
	switch addressingMode {
	case ZEROPAGE:
		// Get the address from the operand
		address = uint16(operand1())
		setFlags()
		updateCycleCounter(5)
		handleState(2)
	case ZEROPAGEX:
		// Get the address from the operand with X offset
		address = uint16(operand1() + X)
		setFlags()
		updateCycleCounter(6)
		handleState(2)
	case ABSOLUTE:
		// Get 16-bit address from operand1 and operand2
		address = uint16(operand2())<<8 | uint16(operand1())
		setFlags()
		updateCycleCounter(6)
		handleState(3)
	case ABSOLUTEX:
		// Get 16-bit address from operand1 and operand2 with X offset
		address = (uint16(operand2())<<8 | uint16(operand1())) + uint16(X)
		setFlags()
		updateCycleCounter(7)
		handleState(3)
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
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If the result is 0, set the zero flag
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
	}
	switch addressingMode {
	case ZEROPAGE:
		// Get the address from the operand
		address = uint16(operand1())
		setFlags()
		updateCycleCounter(5)
		handleState(2)
	case ZEROPAGEX:
		// Get the address from the operand with X offset
		address = uint16(operand1() + X)
		setFlags()
		updateCycleCounter(6)
		handleState(2)
	case ABSOLUTE:
		// Get 16-bit address from operand1 and operand2
		address = uint16(operand2())<<8 | uint16(operand1())
		setFlags()
		updateCycleCounter(6)
		handleState(3)
	case ABSOLUTEX:
		// Get 16-bit address from operand1 and operand2 with X offset
		address = (uint16(operand2())<<8 | uint16(operand1())) + uint16(X)
		setFlags()
		updateCycleCounter(7)
		handleState(3)
	}
}
func ADC(addressingMode string) {
	var value byte
	var result int

	setFlags := func() {
		// Binary mode is the default
		tmpResult := int(A) + int(value)
		if getSRBit(0) == 1 {
			tmpResult++
		}

		if getSRBit(3) == 1 { // BCD mode
			temp := (A & 0x0F) + (value & 0x0F) + getSRBit(0)
			if temp > 9 {
				temp += 6
			}

			result = int((A & 0xF0) + (value & 0xF0) + (temp & 0x0F))

			if result > 0x99 {
				result += 0x60
			}

		}

		// Set or unset the C flag
		if tmpResult > 0xFF {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}

		// Handle V (overflow) flag
		if (int(A)^int(value))&0x80 == 0 && (int(A)^tmpResult)&0x80 != 0 {
			setOverflowFlag()
		} else {
			unsetOverflowFlag()
		}

		result = tmpResult & 0xFF // Store the result in 8 bits

		// Handle N (negative) and Z (zero) flags
		if result&0x80 != 0 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}

		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}

		A = byte(result)

		// Your addressing mode cycle counts remain the same
		if addressingMode == IMMEDIATE || addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX || addressingMode == INDIRECTX || addressingMode == INDIRECTY {
			updateCycleCounter(2)
			handleState(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX || addressingMode == ABSOLUTEY {
			updateCycleCounter(3)
			handleState(3)
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = operand1()
		setFlags()
	case ZEROPAGE:
		// Get the address from the operand
		address := operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case ZEROPAGEX:
		// Get the address from the operand
		address := operand1() + X
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = readMemory(address)
		setFlags()
	case ABSOLUTEX:
		// Get 16 bit address from operand1 and operand2
		address := (uint16(operand2())<<8 | uint16(operand1())) + uint16(X)
		// Get value at address
		value = readMemory(address)
		setFlags()
	case ABSOLUTEY:
		// Get 16 bit address from operand1 and operand2
		address := (uint16(operand2())<<8 | uint16(operand1())) + uint16(Y)
		// Get value at address
		value = readMemory(address)
		setFlags()
	case INDIRECTX:
		// Get the indirect address from the operand
		indirectAddress := operand1() + X
		// Get the address from the indirect address
		address := uint16(readMemory(uint16(indirectAddress+1)))<<8 | uint16(readMemory(uint16(indirectAddress)))
		// Get the value at the address
		value = readMemory(address)
		setFlags()
	case INDIRECTY:
		// Get the indirect address from the operand
		indirectAddress := operand1()
		// Get the address from the indirect address
		address := uint16(readMemory(uint16(indirectAddress+1)))<<8 | uint16(readMemory(uint16(indirectAddress))) + uint16(Y)
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
		if getSRBit(3) == 1 { // BCD mode

			temp := (A & 0x0F) - (value & 0x0F) - (getSRBit(0) ^ 1)
			if temp < 0 {
				temp -= 6
			}

			result = int((A & 0xF0) - (value & 0xF0) - (temp & 0x0F))

			if result < 0 {
				result -= 0x60
			}

		} else {
			// Binary mode
			result = int(A) - int(value)
			if getSRBit(0) == 0 {
				result--
			}
			if int(A) >= int(value) {
				setCarryFlag()
			} else {
				unsetCarryFlag()
			}
		}

		// Negative, and Zero flag checks remain the same
		if readBit(7, byte(result)) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		if result == 0 {
			setZeroFlag()
		}
		A = byte(result)

		// Your addressing mode cycle counts remain the same
		if addressingMode == IMMEDIATE || addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX || addressingMode == INDIRECTX || addressingMode == INDIRECTY {
			updateCycleCounter(2)
			handleState(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX || addressingMode == ABSOLUTEY {
			updateCycleCounter(3)
			handleState(3)
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get the value from the operand
		value = operand1()
		setFlags()
	case ZEROPAGE:
		// Get the address from the operand
		address := operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case ZEROPAGEX:
		// Get the address from the operand
		address := operand1() + X
		// Get the value at the address
		value = readMemory(uint16(address))
		setFlags()
	case ABSOLUTE:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = readMemory(address)
		setFlags()
	case ABSOLUTEX:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Get value at address
		value = readMemory(address)
		setFlags()
	case ABSOLUTEY:
		// Get 16 bit address from operand1 and operand2
		address := uint16(operand2())<<8 | uint16(operand1()) + uint16(Y)
		// Get value at address
		value = readMemory(address)
		setFlags()
	case INDIRECTX:
		// Get the indirect address from the operand
		indirectAddress := operand1() + X
		// Get the address from the indirect address
		address := uint16(readMemory(uint16(indirectAddress+1)))<<8 | uint16(readMemory(uint16(indirectAddress)))
		// Get the value at the address
		value = readMemory(address)
		setFlags()
	case INDIRECTY:
		// Get the indirect address from the operand
		indirectAddress := operand1()
		// Get the address from the indirect address
		address := uint16(readMemory(uint16(indirectAddress+1)))<<8 | uint16(readMemory(uint16(indirectAddress))) + uint16(Y)
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
		if getSRBit(0) == 1 {
			result |= 0x80
			setNegativeFlag()
		} else {
			result &= 0x7F
			unsetNegativeFlag()
		}
		// Set carry flag to bit 0 of value
		if readBit(0, value) == 1 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		if addressingMode == ACCUMULATOR {
			// Store the result in the accumulator
			A = result
			updateCycleCounter(2)
			handleState(1)
		}
		if addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX {
			// Store the value back into memory
			writeMemory(uint16(address), result)
			updateCycleCounter(5)
			handleState(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX {
			// Store the value back into memory
			writeMemory(address16, result)
			updateCycleCounter(6)
			handleState(3)
		}
	}
	switch addressingMode {
	case ACCUMULATOR:
		// Get value from accumulator
		value = A
		// Rotate right one bit
		result = value >> 1
		setFlags()
	case ZEROPAGE:
		// Get address
		address = operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value right 1 bit
		result = value >> 1
		setFlags()
	case ZEROPAGEX:
		// Get X indexed zero page address
		address = operand1() + X
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value right 1 bit
		result = value >> 1
		setFlags()
	case ABSOLUTE:
		// Get 16 bit address from operands
		address16 = uint16(operand2())<<8 | uint16(operand1())
		// Get the value stored at the address in the operands
		value = readMemory(address16)
		// Shift the value right 1 bit
		result = value >> 1
		setFlags()
	case ABSOLUTEX:
		// Get 16 bit address
		address16 = uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
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
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}

		// Update the zero flag
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}

		// Update the negative flag
		if readBit(7, A) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		if addressingMode == ACCUMULATOR {
			// Store the result in the accumulator
			A = result
			updateCycleCounter(2)
			handleState(1)
		}
		if addressingMode == ZEROPAGE || addressingMode == ZEROPAGEX {
			// Store the value back into memory
			writeMemory(uint16(address), result)
			updateCycleCounter(5)
			handleState(2)
		}
		if addressingMode == ABSOLUTE || addressingMode == ABSOLUTEX {
			// Store the value back into memory
			writeMemory(address16, result)
			updateCycleCounter(6)
			handleState(3)
		}
	}
	switch addressingMode {
	case ACCUMULATOR:
		// Get the value of the accumulator
		value = A
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | getSRBit(0)
		setFlags()
	case ZEROPAGE:
		// Get address
		address = operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | getSRBit(0)
		setFlags()
	case ZEROPAGEX:
		// Get X indexed zero page address
		address = operand1() + X
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | getSRBit(0)
		setFlags()
	case ABSOLUTE:
		// Get 16 bit address from operands
		address16 = uint16(operand2())<<8 | uint16(operand1())
		// Get the value stored at the address in the operands
		value = readMemory(address16)
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | getSRBit(0)
		setFlags()
	case ABSOLUTEX:
		// Get 16bit X indexed absolute memory address
		address16 = uint16(operand2())<<8 | uint16(operand1()) + uint16(X)
		// Get the value stored at the address
		value = readMemory(address16)
		// Shift the value left 1 bit
		result = value << 1
		// Update bit 0 of result with the value of the carry flag
		result = (result & 0xFE) | getSRBit(0)
		setFlags()
	}
}
func LSR(addressingMode string) {
	var value, result byte

	setFlags := func() {
		// Reset the SR negative flag
		unsetNegativeFlag()
		// If A is 0 then set SR zero flag else reset it
		// Update the zero flag
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
		// If bit 0 of value is 1 then set SR carry flag else reset it
		if readBit(0, value) == 1 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
	}

	switch addressingMode {
	case ACCUMULATOR:
		// Get the value of the accumulator
		value = A
		// Shift the value right 1 bit
		result = value >> 1
		// Store the result back into the accumulator
		A = result
		setFlags()
		updateCycleCounter(2)
		handleState(1)
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value right 1 bit
		result = value >> 1
		// Store the value back into memory
		writeMemory(uint16(address), result)
		setFlags()
		updateCycleCounter(5)
		handleState(2)
	case ZEROPAGEX:
		// Get the X indexed address
		address := operand1() + X
		// Get the value at the X indexed address
		value = readMemory(uint16(address))
		// Shift the value right 1 bit
		result = value >> 1
		// Store the shifted value in memory
		writeMemory(uint16(address), result)
		setFlags()
		updateCycleCounter(6)
		handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get the value stored at the address in the operands
		value = readMemory(address)
		// Shift the value right 1 bit
		result = value >> 1
		// Store the shifted value back in memory
		writeMemory(address, result)
		setFlags()
		updateCycleCounter(6)
		handleState(3)
	case ABSOLUTEX:
		// Get the 16bit X indexed absolute memory address
		address := uint16(operand2())<<8 | uint16(operand1())
		address += uint16(X)
		// Get the value stored at the address
		value = readMemory(address)
		// Shift the value right 1 bit
		result = value >> 1
		// Store the shifted value back in memory
		writeMemory(address, result)
		setFlags()
		updateCycleCounter(7)
		handleState(3)
	}
}
func ASL(addressingMode string) {
	var value, result byte

	setFlags := func() {
		// Update the zero flag
		if A == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}

		// Update the negative flag
		if readBit(7, A) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}

		// Set the Carry flag based on the original value's bit 7 before the shift operation
		if readBit(7, value) == 1 {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
	}
	switch addressingMode {
	case ACCUMULATOR:
		// Set value to accumulator
		value = A
		// Shift the value left 1 bit
		result = value << 1
		// Update the accumulator with the result
		A = result
		setFlags()
		updateCycleCounter(2)
		handleState(1)
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get the value at the address
		value = readMemory(uint16(address))
		// Shift the value left 1 bit
		result = value << 1
		// Store the value back into memory
		writeMemory(uint16(address), result)
		setFlags()
		updateCycleCounter(5)
		handleState(2)
	case ZEROPAGEX:
		// Get the X indexed address
		address := operand1() + X
		// Get the value at the X indexed address
		value = readMemory(uint16(address))
		// Shift the value left 1 bit
		result = value << 1
		// Store the shifted value in memory
		writeMemory(uint16(address), result)
		setFlags()
		updateCycleCounter(6)
		handleState(2)
	case ABSOLUTE:
		// Get 16 bit address from operands
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get the value stored at the address in the operands
		value = readMemory(address)
		// Shift the value left 1 bit
		result = value << 1
		// Store the shifted value back in memory
		writeMemory(address, result)
		setFlags()
		updateCycleCounter(6)
		handleState(3)
	case ABSOLUTEX:
		// Get the 16bit X indexed absolute memory address
		address := int(operand2())<<8 | int(operand1()) + int(X)
		// Get the value stored at the address
		value = readMemory(uint16(address))
		// Shift the value left 1 bit
		result = value << 1
		// Store the shifted value back in memory
		writeMemory(uint16(address), result)
		setFlags()
		updateCycleCounter(7)
		handleState(3)
	}
}
func CPX(addressingMode string) {
	var value, result byte

	setFlags := func() {
		// If X >= value then set carry flag bit 0 to 1 set carry flag bit 0 to 0
		if X >= value {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		// If value> X then reset carry flag
		if value > X {
			unsetCarryFlag()
		}
		// If bit 7 of result is 1 then set negative flag else unset negative flag
		//if readBit(7, result) == 1 {
		if result&0x80 != 0 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If value == X then set zero flag else unset zero flag
		//if value == X {
		if result == 0 {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get value from operand1
		value = operand1()
		// Compare X with value
		result = X - value
		setFlags()
		updateCycleCounter(2)
		handleState(2)
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get value at address
		value = readMemory(uint16(address))
		// Store result of X-memory stored at operand1() in result variable
		result = X - value
		setFlags()
		updateCycleCounter(3)
		handleState(2)
	case ABSOLUTE:
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = readMemory(address)
		setFlags()
		updateCycleCounter(3)
		handleState(3)
	}
}
func CPY(addressingMode string) {
	var value, result byte

	setFlags := func() {
		// If Y>value then set carry flag to 1 else set carry flag to 0
		if Y >= value {
			setCarryFlag()
		} else {
			unsetCarryFlag()
		}
		// If bit 7 of result is set, set N flag to 1 else reset it
		if readBit(7, result) == 1 {
			setNegativeFlag()
		} else {
			unsetNegativeFlag()
		}
		// If Y==value then set Z flag to 1 else reset it
		if Y == value {
			setZeroFlag()
		} else {
			unsetZeroFlag()
		}
	}
	switch addressingMode {
	case IMMEDIATE:
		// Get value from operand1
		value = operand1()
		// Subtract operand from Y
		result = Y - operand1()
		setFlags()
		updateCycleCounter(2)
		handleState(2)
	case ZEROPAGE:
		// Get address
		address := operand1()
		// Get value at address
		value = readMemory(uint16(address))
		// Store result of Y-memory stored at operand1() in result variable
		result = Y - value
		setFlags()
		updateCycleCounter(3)
		handleState(2)
	case ABSOLUTE:
		// Get address
		address := uint16(operand2())<<8 | uint16(operand1())
		// Get value at address
		value = readMemory(address)
		setFlags()
		updateCycleCounter(3)
		handleState(3)
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
		fmt.Printf("Test failed at PC: %04X\t", PC)
		// print opcode and disassembledInstruction at PC
		fmt.Printf("Opcode: %02X\t", readMemory(uint16(PC)))
		fmt.Printf("Disassembled Instruction: %s\n", disassembledInstruction)
	}

	disassembledInstruction = fmt.Sprintf("BRK\t")
	disassembleOpcode()
	previousPC = PC
	previousOpcode = opcode()
	// Increment PC
	//incPC(1)

	// Decrement SP and Push high byte of (PC+1) onto stack
	decSP()
	updateStack(byte((PC + 1) >> 8))

	// Decrement SP and Push low byte of (PC+1) onto stack
	decSP()
	updateStack(byte((PC + 1) & 0xFF))

	// Set a modified SR with the B flag for the pushed value
	modifiedSR := SR | 0x10
	// Decrement SP and Store modified SR on stack
	setBreakFlag()
	decSP()
	updateStack(modifiedSR)

	// Decrement SP and Store SR on stack
	decSP()
	updateStack(SR)

	// Set SR interrupt disable bit to 1
	setInterruptFlag()

	// Set PC to interrupt vector address
	setPC(int((uint16(readMemory(IRQVectorAddress+1)) << 8) | uint16(readMemory(IRQVectorAddress))))
	updateCycleCounter(7)
	handleState(0)
}
func CLC() {
	/*
		CLC - Clear Carry Flag
	*/
	disassembledInstruction = fmt.Sprintf("CLC\t")
	disassembleOpcode()
	// Set SR carry flag bit 0 to 0
	unsetCarryFlag()
	updateCycleCounter(2)
	handleState(1)
}
func CLD() {
	/*
		CLD - Clear Decimal Mode
	*/
	disassembledInstruction = fmt.Sprintf("CLD\t")
	disassembleOpcode()
	unsetDecimalFlag()
	updateCycleCounter(2)
	handleState(1)
}
func CLI() {
	/*
		CLI - Clear Interrupt Disable
	*/
	disassembledInstruction = fmt.Sprintf("CLI\t")
	disassembleOpcode()
	// Set SR interrupt disable bit 2 to 0
	unsetInterruptFlag()
	updateCycleCounter(2)
	handleState(1)
}
func CLV() {
	/*
		CLV - Clear Overflow Flag
	*/
	disassembledInstruction = fmt.Sprintf("CLV\t")
	disassembleOpcode()
	// Set SR overflow flag bit 6 to 0
	unsetOverflowFlag()
	updateCycleCounter(2)
	handleState(1)
}
func DEX() {
	// DEX - Decrement Index Register X By One
	disassembledInstruction = fmt.Sprintf("DEX\t")
	disassembleOpcode()

	// Decrement the X register by 1
	X--

	// Update the Negative Flag based on the new value of X
	if X&0x80 != 0 {
		setNegativeFlag()
	} else {
		unsetNegativeFlag()
	}

	// Update the Zero Flag based on the new value of X
	if X == 0 {
		setZeroFlag()
	} else {
		unsetZeroFlag()
	}
	updateCycleCounter(2)
	handleState(1)
}
func DEY() {
	/*
		DEY - Decrement Index Register Y By One
	*/
	disassembledInstruction = fmt.Sprintf("DEY\t")
	disassembleOpcode()

	// Decrement the  Y register by 1
	Y--
	// If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
	if getYBit(7) == 1 {
		setNegativeFlag()
	} else {
		unsetNegativeFlag()
	}
	// If Y==0 then set the Zero flag
	if Y == 0 {
		setZeroFlag()
	} else {
		unsetZeroFlag()
	}
	updateCycleCounter(2)
	handleState(1)
}
func INX() {
	/*
		INX - Increment Index Register X By One
	*/
	disassembledInstruction = fmt.Sprintf("INX\t")
	disassembleOpcode()

	// Increment the X register by 1
	X++
	// If X register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
	if getXBit(7) == 1 {
		setNegativeFlag()
	} else {
		unsetNegativeFlag()
	}
	// If X register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
	if X == 0 {
		setZeroFlag()
	} else {
		unsetZeroFlag()
	}
	updateCycleCounter(2)
	handleState(1)
}
func INY() {
	/*
		INY - Increment Index Register Y By One
	*/
	disassembledInstruction = fmt.Sprintf("INY\t")
	disassembleOpcode()

	// Increment the  Y register by 1
	Y++
	// If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
	if getYBit(7) == 1 {
		setNegativeFlag()
	} else {
		unsetNegativeFlag()
	}
	// If Y register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
	if Y == 0 {
		setZeroFlag()
	} else {
		unsetZeroFlag()
	}
	updateCycleCounter(2)
	handleState(1)
}
func NOP() {
	/*
		NOP - No Operation
	*/
	disassembledInstruction = fmt.Sprintf("NOP\t")
	disassembleOpcode()
	updateCycleCounter(2)
	handleState(1)
}
func PHA() {
	/*
		PHA - Push Accumulator On Stack
	*/
	disassembledInstruction = fmt.Sprintf("PHA\t")
	disassembleOpcode()

	// Update memory address pointed to by SP with value stored in accumulator
	updateStack(A)
	decSP()
	updateCycleCounter(3)
	handleState(1)
}
func PHP() {
	/*
	   PHP - Push Processor Status On Stack
	*/
	disassembledInstruction = fmt.Sprintf("PHP\t")
	disassembleOpcode()

	// Set the break flag and the unused bit before pushing
	SR |= 1 << 4 // Set break flag
	SR |= 1 << 5 // Set unused bit

	// Push the SR onto the stack
	updateStack(SR)

	// Decrement the stack pointer
	decSP()
	updateCycleCounter(3)
	handleState(1)
}
func PLA() {
	/*
	   PLA - Pull Accumulator From Stack
	*/
	disassembledInstruction = fmt.Sprintf("PLA\t")
	disassembleOpcode()

	// Increment the stack pointer first
	incSP()

	// Ensure all arithmetic is done in uint16
	expectedAddress := SPBaseAddress + SP

	// Now, update accumulator with value stored in memory address pointed to by SP
	A = readMemory(expectedAddress)

	// If bit 7 of accumulator is set, set negative SR flag else set negative SR flag to 0
	if getABit(7) == 1 {
		setNegativeFlag()
	} else {
		unsetNegativeFlag()
	}

	// If accumulator is 0, set zero SR flag else set zero SR flag to 0
	if A == 0 {
		setZeroFlag()
	} else {
		unsetZeroFlag()
	}
	updateCycleCounter(4)
	handleState(1)
}
func PLP() {
	/*
		PLP - Pull Processor Status From Stack
	*/
	disassembledInstruction = fmt.Sprintf("PLP\t")
	disassembleOpcode()

	// Update SR with the value stored at the address pointed to by SP
	SR = readStack()
	incSP()
	updateCycleCounter(4)
	handleState(1)
}
func RTI() {
	/*
	   RTI - Return From Interrupt
	*/

	disassembledInstruction = fmt.Sprintf("RTI\t")
	disassembleOpcode()

	SR = readStack() & 0xCF
	incSP()

	// Increment the stack pointer to get low byte of PC
	incSP()

	// Get low byte of PC
	low := uint16(readStack())

	// Increment the stack pointer to get high byte of PC
	incSP()

	// Get high byte of PC
	high := uint16(readStack())

	previousPC = PC
	previousOpcode = opcode()
	updateCycleCounter(6)
	handleState(0)
	// Update PC with the value stored in memory at the address pointed to by SP
	setPC(int((high << 8) | low))
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
	incSP()
	//Get high byte of new PC
	high := uint16(readStack())
	previousPC = PC
	previousOpcode = opcode()
	//Update PC with the value stored in memory at the address pointed to by SP
	setPC(int((high<<8)|low) + 1)
	updateCycleCounter(6)
	handleState(0)
}
func SEC() {
	/*
		SEC - Set Carry Flag
	*/
	disassembledInstruction = fmt.Sprintf("SEC\t")
	disassembleOpcode()

	// Set SR carry flag bit 0 to 1
	setCarryFlag()
	updateCycleCounter(2)
	handleState(1)
}
func SED() {
	/*
		SED - Set Decimal Mode
	*/
	disassembledInstruction = fmt.Sprintf("SED\t")
	disassembleOpcode()

	// Set SR decimal mode flag to 1
	setDecimalFlag()
	updateCycleCounter(2)
	handleState(1)
}
func SEI() {
	/*
		SEI - Set Interrupt Disable
	*/
	disassembledInstruction = fmt.Sprintf("SEI\t")
	disassembleOpcode()

	// Set SR interrupt disable bit 2 to 1
	setInterruptFlag()
	updateCycleCounter(2)
	handleState(1)
}
func TAX() {
	/*
		TAX - Transfer Accumulator To Index X
	*/
	disassembledInstruction = fmt.Sprintf("TAX\t")
	disassembleOpcode()

	// Update X with the value of A
	X = A
	// If X register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
	if getXBit(7) == 1 {
		setNegativeFlag()
	} else {
		unsetNegativeFlag()
	}
	// If X register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
	if X == 0 {
		setZeroFlag()
	} else {
		unsetZeroFlag()
	}
	updateCycleCounter(2)
	handleState(1)
}
func TAY() {
	/*
		TAY - Transfer Accumulator To Index Y
	*/
	disassembledInstruction = fmt.Sprintf("TAY\t")
	disassembleOpcode()

	// Set Y register to the value of the accumulator
	Y = A
	// If Y register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
	if getYBit(7) == 1 {
		setNegativeFlag()
	} else {
		unsetNegativeFlag()
	}
	// If Y register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
	if A == 0 {
		setZeroFlag()
	} else {
		unsetZeroFlag()
	}
	updateCycleCounter(2)
	handleState(1)
}
func TSX() {
	/*
		TSX - Transfer Stack Pointer To Index X
	*/
	disassembledInstruction = fmt.Sprintf("TSX\t")
	disassembleOpcode()

	// Update X with the SP
	X = byte(SP)
	// If X register bit 7 is 1, set the SR negative flag bit 7 to 1 else set SR negative flag bit 7 to 0
	if getXBit(7) == 1 {
		setNegativeFlag()
	} else {
		unsetNegativeFlag()
	}
	// If X register is 0, set the SR zero flag bit 1 to 1 else set SR zero flag bit 1 to 0
	if X == 0 {
		setZeroFlag()
	} else {
		unsetZeroFlag()
	}
	updateCycleCounter(2)
	handleState(1)
}
func TXA() {
	/*
		TXA - Transfer Index X To Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("TXA\t")
	disassembleOpcode()

	// Set accumulator to value of X register
	A = X
	// If bit 7 of accumulator is set, set negative SR flag else set negative SR flag to 0
	if getABit(7) == 1 {
		setNegativeFlag()
	} else {
		unsetNegativeFlag()
	}
	// If accumulator is 0, set zero SR flag else set zero SR flag to 0
	if A == 0 {
		setZeroFlag()
	} else {
		unsetZeroFlag()
	}
	updateCycleCounter(2)
	handleState(1)
}
func TXS() {
	/*
		TXS - Transfer Index X To Stack Pointer
	*/
	disassembledInstruction = fmt.Sprintf("TXS\t")
	disassembleOpcode()

	// Set stack pointer to value of X register
	SP = uint16(X)
	updateCycleCounter(2)
	handleState(1)
}
func TYA() {
	/*
		TYA - Transfer Index Y To Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("TYA\t")
	disassembleOpcode()

	// Set accumulator to value of Y register
	A = Y
	// If bit 7 of accumulator is set, set negative SR flag else set negative SR flag to 0
	if getABit(7) == 1 {
		setNegativeFlag()
	} else {
		unsetNegativeFlag()
	}
	// If accumulator is 0, set zero SR flag else set zero SR flag to 0
	if A == 0 {
		setZeroFlag()
	} else {
		unsetZeroFlag()
	}
	updateCycleCounter(2)
	handleState(1)
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
	disassembledInstruction = fmt.Sprintf("ADC #$%02X", operand1())
	disassembleOpcode()

	ADC("immediate")
}
func AND_I() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("AND #$%02X", operand1())
	disassembleOpcode()
	AND("immediate")
}
func CMP_I() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("CMP #$%02X", operand1())
	disassembleOpcode()
	CMP("immediate")
}
func CPX_I() {
	/*
		CPX - Compare Index Register X To Memory
	*/
	disassembledInstruction = fmt.Sprintf("CPX #$%02X", operand1())
	disassembleOpcode()

	CPX("immediate")
}
func CPY_I() {
	/*
		CPY - Compare Index Register Y To Memory
	*/
	disassembledInstruction = fmt.Sprintf("CPY #$%02X", operand1())
	disassembleOpcode()
	CPY("immediate")
}
func EOR_I() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("EOR #$%02X", operand1())
	disassembleOpcode()
	EOR("immediate")
}
func LDA_I() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDA #$%02X", operand1())
	disassembleOpcode()
	LDA("immediate")
}
func LDX_I() {
	/*
		LDX - Load Index Register X From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDX #$%02X", operand1())
	disassembleOpcode()
	LDX("immediate")
}
func LDY_I() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDY #$%02X", operand1())
	disassembleOpcode()
	LDY("immediate")
}
func ORA_I() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("ORA #$%02X", operand1())
	disassembleOpcode()
	ORA("immediate")
}
func SBC_I() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembledInstruction = fmt.Sprintf("SBC #$%02X", operand1())
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
	disassembledInstruction = fmt.Sprintf("ADC $%02X", operand1())
	disassembleOpcode()
	ADC("zeropage")
}
func AND_Z() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("AND $%02X", operand1())
	disassembleOpcode()
	AND("zeropage")
}
func ASL_Z() {
	/*
		ASL - Arithmetic Shift Left
	*/
	disassembledInstruction = fmt.Sprintf("ASL $%02X", operand1())
	disassembleOpcode()

	ASL("zeropage")
}
func BIT_Z() {
	/*
		BIT - Test Bits in Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("BIT $%02X", operand1())
	disassembleOpcode()
	BIT("zeropage")
}
func CMP_Z() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("CMP $%02X", operand1())
	disassembleOpcode()
	CMP("zeropage")
}
func CPX_Z() {
	/*
		CPX - Compare Index Register X To Memory
	*/
	disassembledInstruction = fmt.Sprintf("CPX $%02X", operand1())
	disassembleOpcode()
	CPX("zeropage")
}
func CPY_Z() {
	/*
		CPY - Compare Index Register Y To Memory
	*/
	disassembledInstruction = fmt.Sprintf("CPY $%02X", operand1())
	disassembleOpcode()
	CPY("zeropage")
}
func DEC_Z() {
	/*
		DEC - Decrement Memory By One
	*/
	disassembledInstruction = fmt.Sprintf("DEC $%02X", operand1())
	disassembleOpcode()
	DEC("zeropage")
}
func EOR_Z() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("EOR $%02X", operand1())
	disassembleOpcode()
	EOR("zeropage")
}
func INC_Z() {
	/*
		INC - Increment Memory By One
	*/
	disassembledInstruction = fmt.Sprintf("INC $%02X", operand1())
	disassembleOpcode()
	INC("zeropage")
}
func LDA_Z() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDA $%02X", operand1())
	disassembleOpcode()
	LDA("zeropage")
}
func LDX_Z() {
	/*
		LDX - Load Index Register X From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDX $%02X", operand1())
	disassembleOpcode()
	LDX("zeropage")
}
func LDY_Z() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDY $%02X", operand1())
	disassembleOpcode()
	LDY("zeropage")
}
func LSR_Z() {
	/*
		LSR - Logical Shift Right
	*/
	disassembledInstruction = fmt.Sprintf("LSR $%02X", operand1())
	disassembleOpcode()
	LSR("zeropage")
}
func ORA_Z() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("ORA $%02X", operand1())
	disassembleOpcode()
	ORA("zeropage")
}
func ROL_Z() {
	/*
		ROL - Rotate Left
	*/
	disassembledInstruction = fmt.Sprintf("ROL $%02X", operand1())
	disassembleOpcode()
	ROL("zeropage")
}
func ROR_Z() {
	/*
		ROR - Rotate Right
	*/
	disassembledInstruction = fmt.Sprintf("ROR $%02X", operand1())
	disassembleOpcode()

	ROR("zeropage")
}
func SBC_Z() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembledInstruction = fmt.Sprintf("SBC $%02X", operand1())
	disassembleOpcode()
	SBC("zeropage")
}
func STA_Z() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembledInstruction = fmt.Sprintf("STA $%02X", operand1())
	disassembleOpcode()
	STA("zeropage")
}
func STX_Z() {
	/*
		STX - Store Index Register X In Memory
	*/
	disassembledInstruction = fmt.Sprintf("STX $%02X", operand1())
	disassembleOpcode()
	STX("zeropage")
}
func STY_Z() {
	/*
		STY - Store Index Register Y In Memory
	*/
	disassembledInstruction = fmt.Sprintf("STY $%02X", operand1())
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
	disassembledInstruction = fmt.Sprintf("ADC $%02X,X", operand1())
	disassembleOpcode()
	ADC("zeropagex")
}
func AND_ZX() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("AND $%02X,X", operand1())
	disassembleOpcode()
	AND("zeropagex")
}
func ASL_ZX() {
	/*
		ASL - Arithmetic Shift Left
	*/
	disassembledInstruction = fmt.Sprintf("ASL $%02X,X", operand1())
	disassembleOpcode()
	ASL("zeropagex")
}
func CMP_ZX() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("CMP $%02X,X", operand1())
	disassembleOpcode()
	CMP("zeropagex")
}
func DEC_ZX() {
	/*
		DEC - Decrement Memory By One
	*/
	disassembledInstruction = fmt.Sprintf("DEC $%02X,X", operand1())
	disassembleOpcode()
	DEC("zeropagex")
}
func LDA_ZX() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDA $%02X,X", operand1())
	disassembleOpcode()
	LDA("zeropagex")
}
func LDY_ZX() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDY $%02X,X", operand1())
	disassembleOpcode()
	LDY("zeropagex")
}
func LSR_ZX() {
	/*
		LSR - Logical Shift Right
	*/
	disassembledInstruction = fmt.Sprintf("LSR $%02X,X", operand1())
	disassembleOpcode()
	LSR("zeropagex")
}
func ORA_ZX() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("ORA $%02X,X", operand1())
	disassembleOpcode()
	ORA("zeropagex")
}
func ROL_ZX() {
	/*
		ROL - Rotate Left
	*/
	disassembledInstruction = fmt.Sprintf("ROL $%02X,X", operand1())
	disassembleOpcode()
	ROL("zeropagex")
}
func ROR_ZX() {
	/*
		ROR - Rotate Right
	*/
	disassembledInstruction = fmt.Sprintf("ROR $%02X,X", operand1())
	disassembleOpcode()
	ROR("zeropagex")
}
func EOR_ZX() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("EOR $%02X,X", operand1())
	disassembleOpcode()
	EOR("zeropagex")
}
func INC_ZX() {
	/*
		INC - Increment Memory By One
	*/
	disassembledInstruction = fmt.Sprintf("INC $%02X,X", operand1())
	disassembleOpcode()
	INC("zeropagex")
}
func SBC_ZX() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembledInstruction = fmt.Sprintf("SBC $%02X,X", operand1())
	disassembleOpcode()
	SBC("zeropagex")
}
func STA_ZX() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembledInstruction = fmt.Sprintf("STA $%02X,X", operand1())
	disassembleOpcode()
	STA("zeropagex")
}
func STY_ZX() {
	/*
		STY - Store Index Register Y In Memory
	*/
	disassembledInstruction = fmt.Sprintf("STY $%02X,X", operand1())
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
	disassembledInstruction = fmt.Sprintf("LDX $%02X,Y", operand1())
	disassembleOpcode()
	LDX("zeropagey")
}
func STX_ZY() {
	/*
		STX - Store Index Register X In Memory
	*/
	disassembledInstruction = fmt.Sprintf("STX $%02X,Y", operand1())
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
	disassembledInstruction = fmt.Sprintf("ADC ($%02X,X)", operand1())
	disassembleOpcode()
	ADC("indirectx")
}
func AND_IX() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("AND ($%02X,X)", operand1())
	disassembleOpcode()
	AND("indirectx")
}
func CMP_IX() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("CMP ($%02X,X)", operand1())
	disassembleOpcode()
	CMP("indirectx")
}
func EOR_IX() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("EOR ($%02X,X)", operand1())
	disassembleOpcode()
	EOR("indirectx")
}
func LDA_IX() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDA ($%02X,X)", operand1())
	disassembleOpcode()
	LDA("indirectx")
}
func ORA_IX() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("ORA ($%02X,X)", operand1())
	disassembleOpcode()
	ORA("indirectx")
}
func SBC_IX() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembledInstruction = fmt.Sprintf("SBC ($%02X,X)", operand1())
	disassembleOpcode()
	SBC("indirectx")
}
func STA_IX() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembledInstruction = fmt.Sprintf("STA ($%02X,X)", operand1())
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
	disassembledInstruction = fmt.Sprintf("ADC ($%02X),Y", operand1())
	disassembleOpcode()
	ADC("indirecty")
}
func AND_IY() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("AND ($%02X),Y", operand1())
	disassembleOpcode()
	AND("indirecty")
}
func CMP_IY() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("CMP ($%02X),Y", operand1())
	disassembleOpcode()
	CMP("indirecty")
}
func EOR_IY() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("EOR ($%02X),Y", operand1())
	disassembleOpcode()
	EOR("indirecty")
}
func LDA_IY() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDA ($%02X),Y", operand1())
	disassembleOpcode()
	LDA("indirecty")
}
func ORA_IY() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("ORA ($%02X),Y", operand1())
	disassembleOpcode()
	ORA("indirecty")
}
func SBC_IY() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembledInstruction = fmt.Sprintf("SBC ($%02X),Y", operand1())
	disassembleOpcode()
	SBC("indirecty")
}
func STA_IY() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembledInstruction = fmt.Sprintf("STA ($%02X),Y", operand1())
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
	disassembledInstruction = fmt.Sprintf("BPL $%02X", (PC+2+int(operand1()))&0xFF)
	disassembleOpcode()
	offset := operand1()
	signedOffset := int8(offset)
	// Calculate the branch target address
	targetAddress := PC + 2 + int(signedOffset)
	// If N flag is not set, branch to address
	if getSRBit(7) == 0 {
		// Branch
		setPC(targetAddress)
		updateCycleCounter(1)
		//handleState(0)
		instructionCounter++
	} else {
		// Don't branch
		// Increment the instruction counter by 2
		updateCycleCounter(1)
		handleState(2)
	}
}
func BMI_R() {
	/*
		BMI - Branch on Result Minus
	*/
	disassembledInstruction = fmt.Sprintf("BMI $%02X", (PC+2+int(int8(operand1())))&0xFFFF)
	disassembleOpcode()
	// Get offset from operand
	offset := int8(operand1())
	// If N flag is set, branch to address
	if getSRBit(7) == 1 {
		// Branch
		// Add offset to PC (already incremented by 2)
		setPC(PC + 2 + int(offset))
	} else {
		// Don't branch
		setPC(PC + 2)
	}
}
func BVC_R() {
	/*
		BVC - Branch on Overflow Clear
	*/
	disassembledInstruction = fmt.Sprintf("BVC $%02X", PC+2+int(operand1()))
	disassembleOpcode()
	// Get offset from operand
	offset := operand1()
	// If overflow flag is not set, branch to address
	if getSRBit(6) == 0 {
		updateCycleCounter(1)
		handleState(0)
		// Branch
		// Add offset to lower 8bits of PC
		setPC(PC + 3 + int(offset)&0xFF)
		// If the offset is negative, decrement the PC by 1
		// If bit 7 is unset then it's negative
		if readBit(7, offset) == 0 {
			decPC(1)
		}
	} else {
		// Don't branch
		updateCycleCounter(1)
		handleState(2)
	}
}
func BVS_R() {
	/*
		BVS - Branch on Overflow Set
	*/
	disassembledInstruction = fmt.Sprintf("BVS $%02X", PC+2+int(operand1()))
	disassembleOpcode()
	// Get offset from operand
	offset := operand1()
	// If overflow flag is set, branch to address
	if getSRBit(6) == 1 {
		updateCycleCounter(1)
		handleState(0)
		// Branch
		// Add offset to lower 8bits of PC
		setPC(PC + 3 + int(offset)&0xFF)
		// If the offset is negative, decrement the PC by 1
		// If bit 7 is unset then it's negative
		if readBit(7, offset) == 0 {
			decPC(1)
		}
	} else {
		// Don't branch
		updateCycleCounter(1)
		handleState(2)
	}
}
func BCC_R() {
	/*
		BCC - Branch on Carry Clear
	*/
	disassembledInstruction = fmt.Sprintf("BCC $%02X", PC+2+int(operand1()))
	disassembleOpcode()
	previousPC = PC
	previousOpcode = opcode()
	previousOperand1 = operand1()
	// Get offset from operand
	offset := int8(operand1())
	target := PC + 2 + int(offset)
	if getSRBit(0) == 0 {
		setPC(target)
	} else {
		// Don't branch
		updateCycleCounter(1)
		handleState(2)
	}
}
func BCS_R() {
	/*
		BCS - Branch on Carry Set
	*/
	disassembledInstruction = fmt.Sprintf("BCS $%02X", (PC+2+int(operand1()))&0xFF)
	disassembleOpcode()
	// Get offset from operand
	offset := operand1()
	// If carry flag is set, branch to address
	if getSRBit(0) == 1 {
		updateCycleCounter(1)
		handleState(0)
		// Branch
		// Add offset to lower 8bits of PC
		setPC(PC + 3 + int(offset)&0xFF)
		// If the offset is negative, decrement the PC by 1
		// If bit 7 is unset then it's negative
		if readBit(7, offset) == 0 {
			decPC(1)
		}
	} else {
		// Don't branch
		updateCycleCounter(1)
		handleState(2)
	}
}
func BNE_R() {
	/*
		BNE - Branch on Result Not Zero
	*/
	disassembledInstruction = fmt.Sprintf("BNE $%04X", PC+2+int(int8(operand1())))
	disassembleOpcode()
	// Fetch offset from operand
	offset := operand1()
	// Check Z flag to determine if branching is needed
	if getSRBit(1) == 0 {
		// Calculate the branch target address
		targetAddr := PC + 2 + int(int8(offset))
		// Check if the branch crosses a page boundary
		if (PC & 0xFF00) != (targetAddr & 0xFF00) {
			updateCycleCounter(1)
			handleState(2)
		} else {
			updateCycleCounter(1)
			handleState(1)
		}
		// Update the program counter
		setPC(targetAddr & 0xFFFF)
	} else {
		// If Z flag is set, don't branch
		updateCycleCounter(1)
		handleState(2)
	}
}
func BEQ_R() {
	/*
	   BEQ - Branch on Result Zero
	*/
	disassembledInstruction = fmt.Sprintf("BEQ $%04X", PC+2+int(int8(operand1())))
	disassembleOpcode()

	// Get offset from address in operand
	offset := int8(operand1()) // Cast to signed 8-bit integer to handle negative offsets

	// If Z flag is set, branch to address
	if getSRBit(1) == 1 {
		updateCycleCounter(1)
		handleState(0)
		// Add 2 to PC (1 for opcode, 1 for operand) and then add offset
		setPC(PC + 2 + int(offset))
	} else {
		// Don't branch
		updateCycleCounter(1)
		handleState(2)
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
	disassembledInstruction = fmt.Sprintf("ADC $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	ADC("absolute")
}
func AND_ABS() {
	/*
		AND - "AND" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("AND $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	AND("absolute")
}
func ASL_ABS() {
	/*
		ASL - Arithmetic Shift Left
	*/
	disassembledInstruction = fmt.Sprintf("ASL $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	ASL("absolute")
}
func BIT_ABS() {
	/*
		BIT - Test Bits in Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("BIT $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	BIT("absolute")
}
func CMP_ABS() {
	/*
		CMP - Compare Memory and Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("CMP $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	CMP("absolute")
}
func CPX_ABS() {
	/*
		CPX - Compare Index Register X To Memory
	*/
	disassembledInstruction = fmt.Sprintf("CPX $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	CPX("absolute")
}
func CPY_ABS() {
	/*
		CPY - Compare Index Register Y To Memory
	*/
	disassembledInstruction = fmt.Sprintf("CPY $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	CPY("absolute")
}
func DEC_ABS() {
	/*
		DEC - Decrement Memory By One
	*/
	disassembledInstruction = fmt.Sprintf("DEC $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	DEC("absolute")
}
func EOR_ABS() {
	/*
		EOR - "Exclusive OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("EOR $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	EOR("absolute")
}
func INC_ABS() {
	/*
		INC - Increment Memory By One
	*/
	disassembledInstruction = fmt.Sprintf("INC $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	INC("absolute")
}
func JMP_ABS() {
	/*
		JMP - JMP Absolute
	*/
	disassembledInstruction = fmt.Sprintf("JMP $%04X", int(operand2())<<8|int(operand1()))
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
	disassembledInstruction = fmt.Sprintf("JSR $%04X", int(operand2())<<8|int(operand1()))
	disassembleOpcode()
	// First, push the high byte
	decSP()
	updateStack(byte(PC >> 8))
	decSP()
	updateStack(byte((PC)&0xFF) + 2)

	previousPC = PC
	previousOpcode = opcode()
	previousOperand1 = operand1()
	previousOperand2 = operand2()
	// Now, jump to the subroutine address specified by the operands
	setPC(int(operand2())<<8 | int(operand1()))
	updateCycleCounter(1)
	handleState(0)
}
func LDA_ABS() {
	/*
		LDA - Load Accumulator with Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDA $%04X", uint16(operand2())<<8|uint16(operand1()))
	disassembleOpcode()
	LDA("absolute")
}
func LDX_ABS() {
	/*
		LDX - Load Index Register X From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDX $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	LDX("absolute")
}
func LDY_ABS() {
	/*
		LDY - Load Index Register Y From Memory
	*/
	disassembledInstruction = fmt.Sprintf("LDY $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	LDY("absolute")
}
func LSR_ABS() {
	/*
		LSR - Logical Shift Right
	*/
	disassembledInstruction = fmt.Sprintf("LSR $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	LSR("absolute")
}
func ORA_ABS() {
	/*
		ORA - "OR" Memory with Accumulator
	*/
	disassembledInstruction = fmt.Sprintf("ORA $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	ORA("absolute")
}
func ROL_ABS() {
	/*
		ROL - Rotate Left
	*/
	disassembledInstruction = fmt.Sprintf("ROL $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	ROL("absolute")
}
func ROR_ABS() {
	/*
		ROR - Rotate Right
	*/
	disassembledInstruction = fmt.Sprintf("ROR $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	ROR("absolute")
}
func SBC_ABS() {
	/*
		SBC - Subtract Memory from Accumulator with Borrow
	*/
	disassembledInstruction = fmt.Sprintf("SBC $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	SBC("absolute")
}
func STA_ABS() {
	/*
		STA - Store Accumulator in Memory
	*/
	disassembledInstruction = fmt.Sprintf("STA $%04X", uint16(operand2())<<8|uint16(operand1()))
	disassembleOpcode()
	STA("absolute")
}
func STX_ABS() {
	/*
		STX - Store Index Register X In Memory
	*/
	disassembledInstruction = fmt.Sprintf("STX $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	STX("absolute")
}
func STY_ABS() {
	/*
		STY - Store Index Register Y In Memory
	*/
	disassembledInstruction = fmt.Sprintf("STY $%02X%02X", operand2(), operand1())
	disassembleOpcode()
	STY("absolute")
}
