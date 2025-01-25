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
			bytes, err := parseRV(parts, "LOAD", OP_LOAD)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "STORE_VAL":
			bytes, err := parseAV(parts, "STORE_VAL", OP_STORE_VAL)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "STORE_REG":
			bytes, err := parseAR(parts, "STORE_REG", OP_STORE_REG)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "LOAD_MEM":
			bytes, err := parseRA(parts, "LOAD_MEM", OP_LOAD_MEM)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "ADD":
			bytes, err := parseRR(parts, "ADD", OP_ADD)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "SUB":
			bytes, err := parseRR(parts, "SUB", OP_SUB)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "MUL":
			bytes, err := parseRR(parts, "MUL", OP_MUL)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "DIV":
			bytes, err := parseRR(parts, "DIV", OP_DIV)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "MOD":
			bytes, err := parseRR(parts, "MOD", OP_MOD)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "AND":
			bytes, err := parseRR(parts, "AND", OP_AND)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "OR":
			bytes, err := parseRR(parts, "OR", OP_OR)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "XOR":
			bytes, err := parseRR(parts, "XOR", OP_XOR)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "NOT":
			bytes, err := parseR(parts, "NOT", OP_NOT)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "SHL":
			bytes, err := parseR(parts, "SHL", OP_SHL)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "SHR":
			bytes, err := parseR(parts, "SHR", OP_SHR)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "INC":
			bytes, err := parseR(parts, "INC", OP_INC)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "DEC":
			bytes, err := parseR(parts, "DEC", OP_DEC)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "JMP":
			bytes, err := parseA(parts, "JMP", OP_JMP)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "JMP_REG":
			bytes, err := parseR(parts, "JMP_REG", OP_JMP_REG)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "PUSH":
			bytes, err := parseV(parts, "PUSH", OP_PUSH)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "PUSH_REG":
			bytes, err := parseR(parts, "PUSH_REG", OP_PUSH_REG)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "POP":
			bytes, err := parseNone(parts, "POP", OP_POP)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "POP_REG":
			bytes, err := parseR(parts, "POP_REG", OP_POP_REG)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "CMP_REG_VAL":
			bytes, err := parseRV(parts, "CMP_REG_VAL", OP_CMP_REG_VAL)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "CMP_REG_REG":
			bytes, err := parseRR(parts, "CMP_REG_REG", OP_CMP_REG_REG)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "JE":
			bytes, err := parseA(parts, "JE", OP_JE)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "JNE":
			bytes, err := parseA(parts, "JNE", OP_JNE)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "JG":
			bytes, err := parseA(parts, "JG", OP_JG)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "JGE":
			bytes, err := parseA(parts, "JGE", OP_JGE)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "JL":
			bytes, err := parseA(parts, "JL", OP_JL)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "JLE":
			bytes, err := parseA(parts, "JLE", OP_JLE)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		case "HLT":
			bytes, err := parseNone(parts, "HLT", OP_HLT)
			if err != nil {
				return nil, err
			}
			bytecode = append(bytecode, bytes...)

		default:
			return nil, fmt.Errorf("Unknown instruction: %s", parts[0])
		}
	}

	return bytecode, nil
}

func validRegister(reg string) bool {
	if reg != "R0" && reg != "R1" && reg != "R2" && reg != "R3" {
		return false
	}
	return true
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

func parseRR(parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 3 {
		return nil, fmt.Errorf("%s instruction must have 2 operands", opcodeName)
	}
	if !validRegister(parts[1]) || !validRegister(parts[2]) {
		return nil, fmt.Errorf("Invalid register in %s instruction", opcodeName)
	}
	return []uint8{opcode, uint8(RegisterMap[parts[1]]), uint8(RegisterMap[parts[2]])}, nil
}

func parseRV(parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 3 {
		return nil, fmt.Errorf("%s instruction must have 2 operands", opcodeName)
	}
	if !validRegister(parts[1]) {
		return nil, fmt.Errorf("Invalid register in %s instruction", opcodeName)
	}
	value, err := strconv.Atoi(parts[2])
	if err != nil || !validValue(value) {
		return nil, fmt.Errorf("Invalid value in %s instruction", opcodeName)
	}
	return []uint8{opcode, uint8(RegisterMap[parts[1]]), uint8(value)}, nil
}

func parseRA(parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 3 {
		return nil, fmt.Errorf("%s instruction must have 2 operands", opcodeName)
	}
	if !validRegister(parts[1]) {
		return nil, fmt.Errorf("Invalid register in %s instruction", opcodeName)
	}
	address, err := strconv.Atoi(parts[2])
	if err != nil || !validStoredMemoryAddress(address) {
		return nil, fmt.Errorf("Invalid address in %s instruction", opcodeName)
	}
	return []uint8{opcode, uint8(RegisterMap[parts[1]]), uint8(address)}, nil
}

func parseAR(parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 3 {
		return nil, fmt.Errorf("%s instruction must have 2 operands", opcodeName)
	}
	address, err := strconv.Atoi(parts[1])
	if err != nil || !validStoredMemoryAddress(address) {
		return nil, fmt.Errorf("Invalid address in %s instruction", opcodeName)
	}
	if !validRegister(parts[2]) {
		return nil, fmt.Errorf("Invalid register in %s instruction", opcodeName)
	}
	return []uint8{opcode, uint8(address), uint8(RegisterMap[parts[2]])}, nil
}

func parseAV(parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 3 {
		return nil, fmt.Errorf("%s instruction must have 2 operands", opcodeName)
	}
	address, err := strconv.Atoi(parts[1])
	if err != nil || !validStoredMemoryAddress(address) {
		return nil, fmt.Errorf("Invalid address in %s instruction", opcodeName)
	}
	value, err := strconv.Atoi(parts[2])
	if err != nil || !validValue(value) {
		return nil, fmt.Errorf("Invalid value in %s instruction", opcodeName)
	}
	return []uint8{opcode, uint8(address), uint8(value)}, nil
}

func parseA(parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 2 {
		return nil, fmt.Errorf("%s instruction must have 1 operand", opcodeName)
	}
	address, err := strconv.Atoi(parts[1])
	if err != nil || !validStoredMemoryAddress(address) {
		return nil, fmt.Errorf("Invalid address in %s instruction", opcodeName)
	}
	return []uint8{opcode, uint8(address)}, nil
}

func parseV(parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 2 {
		return nil, fmt.Errorf("%s instruction must have 1 operand", opcodeName)
	}
	value, err := strconv.Atoi(parts[1])
	if err != nil || !validValue(value) {
		return nil, fmt.Errorf("Invalid value in %s instruction", opcodeName)
	}
	return []uint8{opcode, uint8(value)}, nil
}

func parseR(parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 2 {
		return nil, fmt.Errorf("%s instruction must have 1 operand", opcodeName)
	}
	if !validRegister(parts[1]) {
		return nil, fmt.Errorf("Invalid register in %s instruction", opcodeName)
	}
	return []uint8{opcode, uint8(RegisterMap[parts[1]])}, nil
}

func parseNone(parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 1 {
		return nil, fmt.Errorf("%s instruction must have 0 operands", opcodeName)
	}
	return []uint8{opcode}, nil
}
