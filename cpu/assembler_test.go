package cpu

import "testing"

// func TestAssembler(t *testing.T) {
// 	program := []string{
// 		"LOAD R0 42",
// 		"LOAD R1 10",
// 		"ADD R0 R1",
// 		"HLT",
// 	}
//
// 	asm := NewAssembler(program)
//
// 	bytecode, err := asm.Assemble()
// 	if err != nil {
// 		t.Errorf("Assemble failed: %s", err.Error())
// 	}
//
// 	expected := []uint8{
// 		uint8(OP_LOAD), 0, 42,
// 		uint8(OP_LOAD), 1, 10,
// 		uint8(OP_ADD), 0, 1,
// 		uint8(OP_HLT),
// 	}
//
// 	for i, b := range bytecode {
// 		if b != expected[i] {
// 			t.Errorf("Expected bytecode at %d to be %d, got %d", i, expected[i], b)
// 		}
// 	}
// }

func TestAssemblerLabels(t *testing.T) {
	program := []string{
		"LOAD R0 42",
		"LOAD R1 10",
		"CMP_REG_REG R0 R1",
		"JE success",
		"JMP fail",
		"success:",
		"LOAD R0 1",
		"PRINT_REG R0",
		"HLT",
		"fail:",
		"LOAD R0 0",
		"PRINT_REG R0",
		"HLT",
	}

	asm := NewAssembler(program)

	bytecode, err := asm.Assemble()
	if err != nil {
		t.Errorf("Assemble failed: %s", err.Error())
	}

	expected := []uint8{
		uint8(OP_LOAD), 0, 42,
		uint8(OP_LOAD), 1, 10,
		uint8(OP_CMP_REG_REG), 0, 1,
		uint8(OP_JE), 68,
		uint8(OP_JMP), 74,
		uint8(OP_LOAD), 0, 1,
		uint8(OP_PRINT_REG), 0,
		uint8(OP_HLT),
		uint8(OP_LOAD), 0, 0,
		uint8(OP_PRINT_REG), 0,
		uint8(OP_HLT),
	}

	for i, b := range bytecode {
		if b != expected[i] {
			t.Errorf("Expected bytecode at %d to be %d, got %d", i, expected[i], b)
		}
	}
}
