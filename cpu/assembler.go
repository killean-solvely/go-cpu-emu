package cpu

import (
	"fmt"
	"strconv"
	"strings"
)

func Assemble(program []string) ([]uint8, error) {
	var bytecode []uint8

	bookmarks := make(map[string]int)
	lookback := make(map[string]int)
	opcodeCount := 0
	for i, line := range program {
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		if len(parts) == 1 && parts[0][len(parts[0])-1] == ':' {
			bookmarks[parts[0][:len(parts[0])-1]] = opcodeCount + StoredMemorySize
			continue
		}

		switch parts[0] {
		case "LOAD":
			bytes, err := parseRV(i, parts, "LOAD", OP_LOAD)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "STORE_VAL":
			bytes, err := parseAV(i, parts, "STORE_VAL", OP_STORE_VAL)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "STORE_REG":
			bytes, err := parseAR(i, parts, "STORE_REG", OP_STORE_REG)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "LOAD_MEM":
			bytes, err := parseRA(i, parts, "LOAD_MEM", OP_LOAD_MEM)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "ADD":
			bytes, err := parseRR(i, parts, "ADD", OP_ADD)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "SUB":
			bytes, err := parseRR(i, parts, "SUB", OP_SUB)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "MUL":
			bytes, err := parseRR(i, parts, "MUL", OP_MUL)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "DIV":
			bytes, err := parseRR(i, parts, "DIV", OP_DIV)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "MOD":
			bytes, err := parseRR(i, parts, "MOD", OP_MOD)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "AND":
			bytes, err := parseRR(i, parts, "AND", OP_AND)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "OR":
			bytes, err := parseRR(i, parts, "OR", OP_OR)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "XOR":
			bytes, err := parseRR(i, parts, "XOR", OP_XOR)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "NOT":
			bytes, err := parseR(i, parts, "NOT", OP_NOT)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "SHL":
			bytes, err := parseR(i, parts, "SHL", OP_SHL)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "SHR":
			bytes, err := parseR(i, parts, "SHR", OP_SHR)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "INC":
			bytes, err := parseR(i, parts, "INC", OP_INC)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "DEC":
			bytes, err := parseR(i, parts, "DEC", OP_DEC)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "JMP":
			bytes, err := parseAL(i, parts, "JMP", OP_JMP, bookmarks)
			if err != nil {
				if err.(*AssemblerError).Type == INVALID_LABEL {
					lookback[parts[1]] = opcodeCount + len(bytes)
				} else {
					return nil, err
				}
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "JMP_REG":
			bytes, err := parseRL(i, parts, "JMP_REG", OP_JMP_REG, bookmarks)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "PUSH":
			bytes, err := parseV(i, parts, "PUSH", OP_PUSH)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "PUSH_REG":
			bytes, err := parseR(i, parts, "PUSH_REG", OP_PUSH_REG)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "POP":
			bytes, err := parseNone(i, parts, "POP", OP_POP)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "POP_REG":
			bytes, err := parseR(i, parts, "POP_REG", OP_POP_REG)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "CMP_REG_VAL":
			bytes, err := parseRV(i, parts, "CMP_REG_VAL", OP_CMP_REG_VAL)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "CMP_REG_REG":
			bytes, err := parseRR(i, parts, "CMP_REG_REG", OP_CMP_REG_REG)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "JE":
			bytes, err := parseAL(i, parts, "JE", OP_JE, bookmarks)
			if err != nil {
				if err.(*AssemblerError).Type == INVALID_LABEL {
					lookback[parts[1]] = opcodeCount + len(bytes)
				} else {
					return nil, err
				}
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "JNE":
			bytes, err := parseAL(i, parts, "JNE", OP_JNE, bookmarks)
			if err != nil {
				if err.(*AssemblerError).Type == INVALID_LABEL {
					lookback[parts[1]] = opcodeCount + len(bytes)
				} else {
					return nil, err
				}
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "JG":
			bytes, err := parseAL(i, parts, "JG", OP_JG, bookmarks)
			if err != nil {
				if err.(*AssemblerError).Type == INVALID_LABEL {
					lookback[parts[1]] = opcodeCount + len(bytes)
				} else {
					return nil, err
				}
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "JGE":
			bytes, err := parseAL(i, parts, "JGE", OP_JGE, bookmarks)
			if err != nil {
				if err.(*AssemblerError).Type == INVALID_LABEL {
					lookback[parts[1]] = opcodeCount + len(bytes)
				} else {
					return nil, err
				}
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "JL":
			bytes, err := parseAL(i, parts, "JL", OP_JL, bookmarks)
			if err != nil {
				if err.(*AssemblerError).Type == INVALID_LABEL {
					lookback[parts[1]] = opcodeCount + len(bytes)
				} else {
					return nil, err
				}
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "JLE":
			bytes, err := parseAL(i, parts, "JLE", OP_JLE, bookmarks)
			if err != nil {
				if err.(*AssemblerError).Type == INVALID_LABEL {
					lookback[parts[1]] = opcodeCount + len(bytes)
				} else {
					return nil, err
				}
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "PRINT":
			bytes, err := parseV(i, parts, "PRINT", OP_PRINT)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "PRINT_REG":
			bytes, err := parseR(i, parts, "PRINT_REG", OP_PRINT_REG)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		case "HLT":
			bytes, err := parseNone(i, parts, "HLT", OP_HLT)
			if err != nil {
				return nil, err
			}
			opcodeCount += len(bytes)
			bytecode = append(bytecode, bytes...)

		default:
			return nil, fmt.Errorf("Unknown instruction: %s", parts[0])
		}
	}

	for label, address := range lookback {
		if _, ok := bookmarks[label]; !ok {
			return nil, fmt.Errorf("Invalid label in instruction: %s", label)
		}
		bytecode[address-1] = uint8(bookmarks[label])
	}

	return bytecode, nil
}

func validRegister(reg string) bool {
	if _, ok := RegisterMap[reg]; !ok {
		return false
	}
	return true
}

func validValue(value int) bool {
	return value >= 0 && value <= 255
}

func validAddress(address int) bool {
	return address >= 0 && address < TotalMemorySize
}

func parseRR(line int, parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 3 {
		return nil, NewAssemblerError(
			INVALID_OPERAND_COUNT,
			line,
			opcode,
			opcodeName,
			"Instruction must have 2 operands",
		)
	}
	if !validRegister(parts[1]) || !validRegister(parts[2]) {
		return nil, NewAssemblerError(
			INVALID_REGISTER,
			line,
			opcode,
			opcodeName,
			"Invalid register",
		)
	}
	return []uint8{opcode, uint8(RegisterMap[parts[1]]), uint8(RegisterMap[parts[2]])}, nil
}

func parseRV(line int, parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 3 {
		return nil, NewAssemblerError(
			INVALID_OPERAND_COUNT,
			line,
			opcode,
			opcodeName,
			"Instruction must have 2 operands",
		)
	}
	if !validRegister(parts[1]) {
		return nil, NewAssemblerError(
			INVALID_REGISTER,
			line,
			opcode,
			opcodeName,
			"Invalid register",
		)
	}
	value, err := strconv.Atoi(parts[2])
	if err != nil || !validValue(value) {
		return nil, NewAssemblerError(INVALID_VALUE, line, opcode, opcodeName, "Invalid value")
	}
	return []uint8{opcode, uint8(RegisterMap[parts[1]]), uint8(value)}, nil
}

func parseRA(line int, parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 3 {
		return nil, NewAssemblerError(
			INVALID_OPERAND_COUNT,
			line,
			opcode,
			opcodeName,
			"Instruction must have 2 operands",
		)
	}
	if !validRegister(parts[1]) {
		return nil, NewAssemblerError(
			INVALID_REGISTER,
			line,
			opcode,
			opcodeName,
			"Invalid register",
		)
	}
	address, err := strconv.Atoi(parts[2])
	if err != nil || !validAddress(address) {
		return nil, NewAssemblerError(INVALID_ADDRESS, line, opcode, opcodeName, "Invalid address")
	}
	return []uint8{opcode, uint8(RegisterMap[parts[1]]), uint8(address)}, nil
}

func parseRL(
	line int,
	parts []string,
	opcodeName string,
	opcode uint8,
	bookmarks map[string]int,
) ([]uint8, error) {
	if len(parts) != 2 {
		return nil, NewAssemblerError(
			INVALID_OPERAND_COUNT,
			line,
			opcode,
			opcodeName,
			"Instruction must have 1 operand",
		)
	}
	if _, ok := bookmarks[parts[1]]; !ok {
		return []uint8{
				opcode,
				0,
			}, NewAssemblerError(
				INVALID_LABEL,
				line,
				opcode,
				opcodeName,
				"Invalid label",
			)
	}
	return []uint8{opcode, uint8(bookmarks[parts[1]])}, nil
}

func parseAR(line int, parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 3 {
		return nil, NewAssemblerError(
			INVALID_OPERAND_COUNT,
			line,
			opcode,
			opcodeName,
			"Instruction must have 2 operands",
		)
	}
	address, err := strconv.Atoi(parts[1])
	if err != nil || !validAddress(address) {
		return nil, NewAssemblerError(INVALID_ADDRESS, line, opcode, opcodeName, "Invalid address")
	}
	if !validRegister(parts[2]) {
		return nil, NewAssemblerError(
			INVALID_REGISTER,
			line,
			opcode,
			opcodeName,
			"Invalid register",
		)
	}
	return []uint8{opcode, uint8(address), uint8(RegisterMap[parts[2]])}, nil
}

func parseAV(line int, parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 3 {
		return nil, NewAssemblerError(
			INVALID_OPERAND_COUNT,
			line,
			opcode,
			opcodeName,
			"Instruction must have 2 operands",
		)
	}
	address, err := strconv.Atoi(parts[1])
	if err != nil || !validAddress(address) {
		return nil, NewAssemblerError(INVALID_ADDRESS, line, opcode, opcodeName, "Invalid address")
	}
	value, err := strconv.Atoi(parts[2])
	if err != nil || !validValue(value) {
		return nil, NewAssemblerError(INVALID_VALUE, line, opcode, opcodeName, "Invalid value")
	}
	return []uint8{opcode, uint8(address), uint8(value)}, nil
}

func parseAL(
	line int,
	parts []string,
	opcodeName string,
	opcode uint8,
	bookmarks map[string]int,
) ([]uint8, error) {
	if len(parts) != 2 {
		return nil, NewAssemblerError(
			INVALID_OPERAND_COUNT,
			line,
			opcode,
			opcodeName,
			"Instruction must have 1 operand",
		)
	}
	if _, ok := bookmarks[parts[1]]; !ok {
		return []uint8{
				opcode,
				0,
			}, NewAssemblerError(
				INVALID_LABEL,
				line,
				opcode,
				opcodeName,
				"Invalid label",
			)
	}
	return []uint8{opcode, uint8(bookmarks[parts[1]])}, nil
}

func parseA(line int, parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 2 {
		return nil, NewAssemblerError(
			INVALID_OPERAND_COUNT,
			line,
			opcode,
			opcodeName,
			"Instruction must have 1 operand",
		)
	}
	address, err := strconv.Atoi(parts[1])
	if err != nil || !validAddress(address) {
		return nil, NewAssemblerError(INVALID_ADDRESS, line, opcode, opcodeName, "Invalid address")
	}
	return []uint8{opcode, uint8(address)}, nil
}

func parseV(line int, parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 2 {
		return nil, NewAssemblerError(
			INVALID_OPERAND_COUNT,
			line,
			opcode,
			opcodeName,
			"Instruction must have 1 operand",
		)
	}
	value, err := strconv.Atoi(parts[1])
	if err != nil || !validValue(value) {
		return nil, NewAssemblerError(INVALID_VALUE, line, opcode, opcodeName, "Invalid value")
	}
	return []uint8{opcode, uint8(value)}, nil
}

func parseR(line int, parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 2 {
		return nil, NewAssemblerError(
			INVALID_OPERAND_COUNT,
			line,
			opcode,
			opcodeName,
			"Instruction must have 1 operand",
		)
	}
	if !validRegister(parts[1]) {
		return nil, NewAssemblerError(
			INVALID_REGISTER,
			line,
			opcode,
			opcodeName,
			"Invalid register",
		)
	}
	return []uint8{opcode, uint8(RegisterMap[parts[1]])}, nil
}

func parseNone(line int, parts []string, opcodeName string, opcode uint8) ([]uint8, error) {
	if len(parts) != 1 {
		return nil, NewAssemblerError(
			INVALID_OPERAND_COUNT,
			line,
			opcode,
			opcodeName,
			"Instruction must have 0 operands",
		)
	}
	return []uint8{opcode}, nil
}
