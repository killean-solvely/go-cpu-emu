package cpu

import "testing"

func TestAssembler(t *testing.T) {
	program := []string{
		"LOAD 0 42",
		"LOAD 1 10",
		"ADD 0 1",
		"HLT",
	}

	bytecode, err := Assemble(program)
	if err != nil {
		t.Errorf("Assemble failed: %s", err.Error())
	}

	expected := []uint8{
		OP_LOAD, 0, 42,
		OP_LOAD, 1, 10,
		OP_ADD, 0, 1,
		OP_HLT,
	}

	for i, b := range bytecode {
		if b != expected[i] {
			t.Errorf("Expected bytecode at %d to be %d, got %d", i, expected[i], b)
		}
	}
}
