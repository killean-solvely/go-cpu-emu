package cpu

import "testing"

func TestInstructionEncoding(t *testing.T) {
	inst := Instruction{Opcode: OP_LOAD, Operand1: 1, Operand2: 42}

	if inst.Opcode != OP_LOAD {
		t.Errorf("Expected Opcode to be OP_LOAD, got %d", inst.Opcode)
	}

	if inst.Operand1 != 1 {
		t.Errorf("Expected Operand1 to be 1, got %d", inst.Operand1)
	}

	if inst.Operand2 != 42 {
		t.Errorf("Expected Operand2 to be 42, got %d", inst.Operand2)
	}
}
