package cpu

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Assembler struct {
	Program        []string
	OpcodeCount    int
	LabelAddresses map[string]int // label name to label address
	JumpAddresses  map[int]string // jump location to label name
	ParseMap       map[string]func(int, []string, string, Opcode) ([]uint8, error)
}

func NewAssembler(program []string) *Assembler {
	labelAddresses := make(map[string]int)
	jumpAddresses := make(map[int]string)
	asm := &Assembler{
		Program:        program,
		OpcodeCount:    0,
		LabelAddresses: labelAddresses,
		JumpAddresses:  jumpAddresses,
		ParseMap:       make(map[string]func(int, []string, string, Opcode) ([]uint8, error)),
	}

	var instructionTypeToParseFunc = map[InstructionType]func(int, []string, string, Opcode) ([]uint8, error){
		INST_R_R:  asm.parseRR,
		INST_R_V:  asm.parseRV,
		INST_R_A:  asm.parseRA,
		INST_R_L:  asm.parseRL,
		INST_A_V:  asm.parseAV,
		INST_A_L:  asm.parseAL,
		INST_A:    asm.parseA,
		INST_V:    asm.parseV,
		INST_R:    asm.parseR,
		INST_NONE: asm.parseNone,
	}

	for opcodeString, instruction := range OpcodeMap {
		asm.ParseMap[opcodeString] = instructionTypeToParseFunc[instruction.Type]
	}

	return asm
}

func (a *Assembler) Assemble() ([]uint8, error) {
	a.firstPass()
	bytecode, err := a.secondPass()
	if err != nil {
		return nil, err
	}
	return bytecode, nil
}

// First pass goes through and fills out the labels
func (a *Assembler) firstPass() {
	opcodeCount := 0
	for _, line := range a.Program {
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		// If the line is a label, add it to the label map
		if len(parts) == 1 && parts[0][len(parts[0])-1] == ':' {
			labelName := parts[0][:len(parts[0])-1]
			a.LabelAddresses[labelName] = opcodeCount + StoredMemorySize
			continue
		}

		opcodeName := parts[0]

		jmpRegisters := []string{"JMP_REG", "JE", "JNE", "JG", "JGE", "JL", "JLE"}
		if slices.Contains(jmpRegisters, opcodeName) {
			a.JumpAddresses[opcodeCount+1] = parts[1]
		}

		opcodeCount += InstructionSizeMap[OpcodeMap[opcodeName].Type]
	}

	a.OpcodeCount = opcodeCount
}

// Second pass goes through and fills out the instructions, returning the bytecode
func (a Assembler) secondPass() ([]uint8, error) {
	var bytecode []uint8

	for i, line := range a.Program {
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		// If the line is a label, skip it
		if len(parts) == 1 && parts[0][len(parts[0])-1] == ':' {
			continue
		}

		opcodeName := parts[0]
		instruction, ok := OpcodeMap[opcodeName]
		if !ok {
			return nil, NewAssemblerError(INVALID_OPCODE, i, 0, opcodeName, "Invalid opcode")
		}

		bytes, err := a.ParseMap[opcodeName](i, parts, opcodeName, instruction.Opcode)
		if err != nil {
			return nil, err
		}

		bytecode = append(bytecode, bytes...)
	}

	for jumpAddress, labelName := range a.JumpAddresses {
		labelAddress, ok := a.LabelAddresses[labelName]
		if !ok {
			return nil, NewAssemblerError(INVALID_LABEL, 0, 0, labelName, "Invalid label")
		}

		bytecode[jumpAddress] = uint8(labelAddress)
	}

	return bytecode, nil
}

func (a *Assembler) parseRR(
	line int,
	parts []string,
	opcodeName string,
	opcode Opcode,
) ([]uint8, error) {
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
	return []uint8{uint8(opcode), uint8(RegisterMap[parts[1]]), uint8(RegisterMap[parts[2]])}, nil
}

func (a *Assembler) parseRV(
	line int,
	parts []string,
	opcodeName string,
	opcode Opcode,
) ([]uint8, error) {
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
	return []uint8{uint8(opcode), uint8(RegisterMap[parts[1]]), uint8(value)}, nil
}

func (a *Assembler) parseRA(
	line int,
	parts []string,
	opcodeName string,
	opcode Opcode,
) ([]uint8, error) {
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
	return []uint8{uint8(opcode), uint8(RegisterMap[parts[1]]), uint8(address)}, nil
}

func (a *Assembler) parseRL(
	line int,
	parts []string,
	opcodeName string,
	opcode Opcode,
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
	if _, ok := a.LabelAddresses[parts[1]]; !ok {
		return []uint8{
				uint8(opcode),
				0,
			}, NewAssemblerError(
				INVALID_LABEL,
				line,
				opcode,
				opcodeName,
				"Invalid label",
			)
	}
	return []uint8{uint8(opcode), uint8(a.LabelAddresses[parts[1]])}, nil
}

func (a *Assembler) parseAV(
	line int,
	parts []string,
	opcodeName string,
	opcode Opcode,
) ([]uint8, error) {
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
		if convErr, ok := err.(*strconv.NumError); ok {
			if len(convErr.Num) == 1 && parts[1][0] >= 0 && parts[1][0] <= 255 {
				value = int(parts[1][0])
			} else {
				return nil, NewAssemblerError(INVALID_VALUE, line, opcode, opcodeName, "Invalid value")
			}
		} else {
			return nil, NewAssemblerError(INVALID_VALUE, line, opcode, opcodeName, "Invalid value")
		}
	}

	return []uint8{uint8(opcode), uint8(address), uint8(value)}, nil
}

func (a *Assembler) parseAL(
	line int,
	parts []string,
	opcodeName string,
	opcode Opcode,
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
	if _, ok := a.LabelAddresses[parts[1]]; !ok {
		return []uint8{
				uint8(opcode),
				0,
			}, NewAssemblerError(
				INVALID_LABEL,
				line,
				opcode,
				opcodeName,
				"Invalid label",
			)
	}
	return []uint8{uint8(opcode), uint8(a.LabelAddresses[parts[1]])}, nil
}

func (a *Assembler) parseA(
	line int,
	parts []string,
	opcodeName string,
	opcode Opcode,
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
	address, err := strconv.Atoi(parts[1])
	if err != nil || !validAddress(address) {
		return nil, NewAssemblerError(INVALID_ADDRESS, line, opcode, opcodeName, "Invalid address")
	}
	return []uint8{uint8(opcode), uint8(address)}, nil
}

func (a *Assembler) parseV(
	line int,
	parts []string,
	opcodeName string,
	opcode Opcode,
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

	value, err := strconv.Atoi(parts[1])
	if err != nil || !validValue(value) {
		if convErr, ok := err.(*strconv.NumError); ok {
			if len(convErr.Num) == 1 && parts[1][0] >= 0 && parts[1][0] <= 255 {
				value = int(parts[1][0])
			} else {
				return nil, NewAssemblerError(INVALID_VALUE, line, opcode, opcodeName, "Invalid value")
			}
		} else {
			return nil, NewAssemblerError(INVALID_VALUE, line, opcode, opcodeName, "Invalid value")
		}
	}

	return []uint8{uint8(opcode), uint8(value)}, nil
}

func (a *Assembler) parseR(
	line int,
	parts []string,
	opcodeName string,
	opcode Opcode,
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
	if !validRegister(parts[1]) {
		return nil, NewAssemblerError(
			INVALID_REGISTER,
			line,
			opcode,
			opcodeName,
			"Invalid register",
		)
	}
	return []uint8{uint8(opcode), uint8(RegisterMap[parts[1]])}, nil
}

func (a *Assembler) parseNone(
	line int,
	parts []string,
	opcodeName string,
	opcode Opcode,
) ([]uint8, error) {
	if len(parts) != 1 {
		return nil, NewAssemblerError(
			INVALID_OPERAND_COUNT,
			line,
			opcode,
			opcodeName,
			"Instruction must have 0 operands",
		)
	}
	return []uint8{uint8(opcode)}, nil
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
