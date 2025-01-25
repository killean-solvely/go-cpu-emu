package cpu

import "testing"

func TestAssemblerLabels(t *testing.T) {
	program := []string{
		"LOAD R0 42",
		"LOAD R1 10",
		"CMP R0 R1",
		"JE success",
		"JMP fail",
		"success:",
		"LOAD R0 1",
		"PRINT R0",
		"HLT",
		"fail:",
		"LOAD R0 0",
		"PRINT R0",
		"HLT",
	}

	asm := NewAssembler(program)

	bytecode, err := asm.Assemble()
	if err != nil {
		t.Errorf("Assemble failed: %s", err.Error())
	}

	expected := []uint8{
		uint8(OP_LOAD_RV), 0, 42,
		uint8(OP_LOAD_RV), 1, 10,
		uint8(OP_CMP_RR), 0, 1,
		uint8(OP_JE_A), 68,
		uint8(OP_JMP_A), 74,
		uint8(OP_LOAD_RV), 0, 1,
		uint8(OP_PRINT_R), 0,
		uint8(OP_HLT_NONE),
		uint8(OP_LOAD_RV), 0, 0,
		uint8(OP_PRINT_R), 0,
		uint8(OP_HLT_NONE),
	}

	for i, b := range bytecode {
		if b != expected[i] {
			t.Errorf("Expected bytecode at %d to be %d, got %d", i, expected[i], b)
		}
	}
}
