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
			if err != nil || reg < 0 || reg > RegisterCount-1 {
				return nil, fmt.Errorf("Invalid register in LOAD instruction: %s", line)
			}
			value, err := strconv.Atoi(parts[2])
			if err != nil || value < 0 || value > 255 {
				return nil, fmt.Errorf("Invalid value in LOAD instruction: %s", line)
			}
			bytecode = append(bytecode, OP_LOAD, uint8(reg), uint8(value))

		case "STORE_VAL":
			if len(parts) != 3 {
				return nil, fmt.Errorf("STORE_VAL instruction must have 2 operands: %s", line)
			}
			address, err := strconv.Atoi(parts[1])
			if err != nil || address < 0 || address > StoredMemorySize {
				return nil, fmt.Errorf("Invalid address in STORE_VAL instruction: %s", line)
			}
			value, err := strconv.Atoi(parts[2])
			if err != nil || value < 0 || value > 255 {
				return nil, fmt.Errorf("Invalid value in STORE_VAL instruction: %s", line)
			}
			bytecode = append(bytecode, OP_STORE_VAL, uint8(address), uint8(value))

		case "STORE_REG":
			if len(parts) != 3 {
				return nil, fmt.Errorf("STORE_REG instruction must have 2 operands: %s", line)
			}
			address, err := strconv.Atoi(parts[1])
			if err != nil || address < 0 || address > StoredMemorySize {
				return nil, fmt.Errorf("Invalid address in STORE_REG instruction: %s", line)
			}
			reg, err := strconv.Atoi(parts[2])
			if err != nil || reg < 0 || reg > RegisterCount-1 {
				return nil, fmt.Errorf("Invalid register in STORE_REG instruction: %s", line)
			}
			bytecode = append(bytecode, OP_STORE_REG, uint8(address), uint8(reg))

		case "LOAD_MEM":
			if len(parts) != 3 {
				return nil, fmt.Errorf("LOAD_MEM instruction must have 2 operands: %s", line)
			}
			reg, err := strconv.Atoi(parts[1])
			if err != nil || reg < 0 || reg > RegisterCount-1 {
				return nil, fmt.Errorf("Invalid register in LOAD_MEM instruction: %s", line)
			}
			address, err := strconv.Atoi(parts[2])
			if err != nil || address < 0 || address > StoredMemorySize {
				return nil, fmt.Errorf("Invalid address in LOAD_MEM instruction: %s", line)
			}
			bytecode = append(bytecode, OP_LOAD_MEM, uint8(reg), uint8(address))

		case "ADD":
			if len(parts) != 3 {
				return nil, fmt.Errorf("ADD instruction must have 2 operands: %s", line)
			}
			reg1, err := strconv.Atoi(parts[1])
			reg2, err := strconv.Atoi(parts[2])
			if err != nil || reg1 < 0 || reg1 > RegisterCount-1 || reg2 < 0 ||
				reg2 > RegisterCount-1 {
				return nil, fmt.Errorf("Invalid register in ADD instruction: %s", line)
			}
			bytecode = append(bytecode, OP_ADD, uint8(reg1), uint8(reg2))

		case "JMP":
			if len(parts) != 2 {
				return nil, fmt.Errorf("JMP instruction must have 1 operand: %s", line)
			}
			address, err := strconv.Atoi(parts[1])
			if err != nil || address < 0 || address > TotalMemorySize {
				return nil, fmt.Errorf("Invalid address in JMP instruction: %s", line)
			}
			bytecode = append(bytecode, OP_JMP, uint8(address))

		case "PUSH":
			if len(parts) != 2 {
				return nil, fmt.Errorf("PUSH instruction must have 1 operand: %s", line)
			}
			value, err := strconv.Atoi(parts[1])
			if err != nil || value < 0 || value > 255 {
				return nil, fmt.Errorf("Invalid value in PUSH instruction: %s", line)
			}
			bytecode = append(bytecode, OP_PUSH, uint8(value))

		case "PUSH_REG":
			if len(parts) != 2 {
				return nil, fmt.Errorf("PUSH_REG instruction must have 1 operand: %s", line)
			}
			reg, err := strconv.Atoi(parts[1])
			if err != nil || reg < 0 || reg > RegisterCount-1 {
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
			if err != nil || reg < 0 || reg > RegisterCount-1 {
				return nil, fmt.Errorf("Invalid register in POP_REG instruction: %s", line)
			}
			bytecode = append(bytecode, OP_POP_REG, uint8(reg))

		case "HLT":
			bytecode = append(bytecode, OP_HLT)

		default:
			return nil, fmt.Errorf("Unknown instruction: %s", parts[0])
		}
	}

	return bytecode, nil
}
