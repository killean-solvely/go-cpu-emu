package cpu

import (
	"fmt"
	"strconv"
	"strings"
)

func Assemble(program []string) ([]uint8, error) {
	var bytecode []uint8

	for _, line := range program {
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "LOAD":
			if len(parts) != 3 {
				return nil, fmt.Errorf("LOAD instruction must have 2 operands: %s", line)
			}
			reg, err := strconv.Atoi(parts[1])
			if err != nil || !validRegister(reg) {
				return nil, fmt.Errorf("Invalid register in LOAD instruction: %s", line)
			}
			value, err := strconv.Atoi(parts[2])
			if err != nil || !validValue(value) {
				return nil, fmt.Errorf("Invalid value in LOAD instruction: %s", line)
			}
			bytecode = append(bytecode, OP_LOAD, uint8(reg), uint8(value))

		case "STORE_VAL":
			if len(parts) != 3 {
				return nil, fmt.Errorf("STORE_VAL instruction must have 2 operands: %s", line)
			}
			address, err := strconv.Atoi(parts[1])
			if err != nil || !validStoredMemoryAddress(address) {
				return nil, fmt.Errorf("Invalid address in STORE_VAL instruction: %s", line)
			}
			value, err := strconv.Atoi(parts[2])
			if err != nil || !validValue(value) {
				return nil, fmt.Errorf("Invalid value in STORE_VAL instruction: %s", line)
			}
			bytecode = append(bytecode, OP_STORE_VAL, uint8(address), uint8(value))

		case "STORE_REG":
			if len(parts) != 3 {
				return nil, fmt.Errorf("STORE_REG instruction must have 2 operands: %s", line)
			}
			address, err := strconv.Atoi(parts[1])
			if err != nil || !validStoredMemoryAddress(address) {
				return nil, fmt.Errorf("Invalid address in STORE_REG instruction: %s", line)
			}
			reg, err := strconv.Atoi(parts[2])
			if err != nil || !validRegister(reg) {
				return nil, fmt.Errorf("Invalid register in STORE_REG instruction: %s", line)
			}
			bytecode = append(bytecode, OP_STORE_REG, uint8(address), uint8(reg))

		case "LOAD_MEM":
			if len(parts) != 3 {
				return nil, fmt.Errorf("LOAD_MEM instruction must have 2 operands: %s", line)
			}
			reg, err := strconv.Atoi(parts[1])
			if err != nil || !validRegister(reg) {
				return nil, fmt.Errorf("Invalid register in LOAD_MEM instruction: %s", line)
			}
			address, err := strconv.Atoi(parts[2])
			if err != nil || !validStoredMemoryAddress(address) {
				return nil, fmt.Errorf("Invalid address in LOAD_MEM instruction: %s", line)
			}
			bytecode = append(bytecode, OP_LOAD_MEM, uint8(reg), uint8(address))

		case "ADD":
			if len(parts) != 3 {
				return nil, fmt.Errorf("ADD instruction must have 2 operands: %s", line)
			}
			reg1, err := strconv.Atoi(parts[1])
			reg2, err := strconv.Atoi(parts[2])
			if err != nil || !validRegister(reg1) || !validRegister(reg2) {
				return nil, fmt.Errorf("Invalid register in ADD instruction: %s", line)
			}
			bytecode = append(bytecode, OP_ADD, uint8(reg1), uint8(reg2))

		case "JMP":
			if len(parts) != 2 {
				return nil, fmt.Errorf("JMP instruction must have 1 operand: %s", line)
			}
			address, err := strconv.Atoi(parts[1])
			if err != nil || !validCodeAddress(address) {
				return nil, fmt.Errorf("Invalid address in JMP instruction: %s", line)
			}
			bytecode = append(bytecode, OP_JMP, uint8(address))

		case "PUSH":
			if len(parts) != 2 {
				return nil, fmt.Errorf("PUSH instruction must have 1 operand: %s", line)
			}
			value, err := strconv.Atoi(parts[1])
			if err != nil || !validValue(value) {
				return nil, fmt.Errorf("Invalid value in PUSH instruction: %s", line)
			}
			bytecode = append(bytecode, OP_PUSH, uint8(value))

		case "PUSH_REG":
			if len(parts) != 2 {
				return nil, fmt.Errorf("PUSH_REG instruction must have 1 operand: %s", line)
			}
			reg, err := strconv.Atoi(parts[1])
			if err != nil || !validRegister(reg) {
				return nil, fmt.Errorf("Invalid register in PUSH_REG instruction: %s", line)
			}
			bytecode = append(bytecode, OP_PUSH_REG, uint8(reg))

		case "POP":
			bytecode = append(bytecode, OP_POP)

		case "POP_REG":
			if len(parts) != 2 {
				return nil, fmt.Errorf("POP_REG instruction must have 1 operand: %s", line)
			}
			reg, err := strconv.Atoi(parts[1])
			if err != nil || !validRegister(reg) {
				return nil, fmt.Errorf("Invalid register in POP_REG instruction: %s", line)
			}
			bytecode = append(bytecode, OP_POP_REG, uint8(reg))

		case "CMP_REG_VAL":
			if len(parts) != 3 {
				return nil, fmt.Errorf("CMP_REG_VAL instruction must have 2 operands: %s", line)
			}
			reg, err := strconv.Atoi(parts[1])
			if err != nil || !validRegister(reg) {
				return nil, fmt.Errorf("Invalid register in CMP_REG_VAL instruction: %s", line)
			}
			value, err := strconv.Atoi(parts[2])
			if err != nil || !validValue(value) {
				return nil, fmt.Errorf("Invalid value in CMP_REG_VAL instruction: %s", line)
			}
			bytecode = append(bytecode, OP_CMP_REG_VAL, uint8(reg), uint8(value))

		case "CMP_REG_REG":
			if len(parts) != 3 {
				return nil, fmt.Errorf("CMP_REG_REG instruction must have 2 operands: %s", line)
			}
			reg1, err := strconv.Atoi(parts[1])
			reg2, err := strconv.Atoi(parts[2])
			if err != nil || !validRegister(reg1) || !validRegister(reg2) {
				return nil, fmt.Errorf("Invalid register in CMP_REG_REG instruction: %s", line)
			}
			bytecode = append(bytecode, OP_CMP_REG_REG, uint8(reg1), uint8(reg2))

		case "JE":
			if len(parts) != 2 {
				return nil, fmt.Errorf("JE instruction must have 1 operand: %s", line)
			}
			address, err := strconv.Atoi(parts[1])
			if err != nil || !validCodeAddress(address) {
				return nil, fmt.Errorf("Invalid address in JE instruction: %s", line)
			}
			bytecode = append(bytecode, OP_JE, uint8(address))

		case "JNE":
			if len(parts) != 2 {
				return nil, fmt.Errorf("JNE instruction must have 1 operand: %s", line)
			}
			address, err := strconv.Atoi(parts[1])
			if err != nil || !validCodeAddress(address) {
				return nil, fmt.Errorf("Invalid address in JNE instruction: %s", line)
			}
			bytecode = append(bytecode, OP_JNE, uint8(address))

		case "JG":
			if len(parts) != 2 {
				return nil, fmt.Errorf("JG instruction must have 1 operand: %s", line)
			}
			address, err := strconv.Atoi(parts[1])
			if err != nil || !validCodeAddress(address) {
				return nil, fmt.Errorf("Invalid address in JG instruction: %s", line)
			}
			bytecode = append(bytecode, OP_JG, uint8(address))

		case "JGE":
			if len(parts) != 2 {
				return nil, fmt.Errorf("JGE instruction must have 1 operand: %s", line)
			}
			address, err := strconv.Atoi(parts[1])
			if err != nil || !validCodeAddress(address) {
				return nil, fmt.Errorf("Invalid address in JGE instruction: %s", line)
			}
			bytecode = append(bytecode, OP_JGE, uint8(address))

		case "JL":
			if len(parts) != 2 {
				return nil, fmt.Errorf("JL instruction must have 1 operand: %s", line)
			}
			address, err := strconv.Atoi(parts[1])
			if err != nil || !validCodeAddress(address) {
				return nil, fmt.Errorf("Invalid address in JL instruction: %s", line)
			}
			bytecode = append(bytecode, OP_JL, uint8(address))

		case "JLE":
			if len(parts) != 2 {
				return nil, fmt.Errorf("JLE instruction must have 1 operand: %s", line)
			}
			address, err := strconv.Atoi(parts[1])
			if err != nil || !validCodeAddress(address) {
				return nil, fmt.Errorf("Invalid address in JLE instruction: %s", line)
			}
			bytecode = append(bytecode, OP_JLE, uint8(address))

		case "HLT":
			bytecode = append(bytecode, OP_HLT)

		default:
			return nil, fmt.Errorf("Unknown instruction: %s", parts[0])
		}
	}

	return bytecode, nil
}

func validRegister(reg int) bool {
	return reg >= 0 && reg < RegisterCount
}

func validValue(value int) bool {
	return value >= 0 && value <= 255
}

func validStoredMemoryAddress(address int) bool {
	return address >= 0 && address < StoredMemorySize
}

func validCodeAddress(address int) bool {
	return address >= CodeMemoryStart && address < TotalMemorySize
}
