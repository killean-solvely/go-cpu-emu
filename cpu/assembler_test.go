package cpu

import (
	"testing"
)

func TestInstructionSizes(t *testing.T) {
	// RR
	asm := NewAssembler([]string{"ADD R1 R2"})
	bytecode, err := asm.Assemble()
	if err != nil {
		t.Fatalf("Assemble failed: %s", err.Error())
	}

	if len(bytecode) != 3 {
		t.Errorf("RR: ADD R1 R2 xpected 3 bytes, got %d", len(bytecode))
	}

	// RV
	asm = NewAssembler([]string{"LOAD R1 1"})
	bytecode, err = asm.Assemble()
	if err != nil {
		t.Fatalf("Assemble failed: %s", err.Error())
	}

	if len(bytecode) != 4 {
		t.Errorf("RV: LOAD R1 1 Expected 4 bytes, got %d", len(bytecode))
	}

	// RA
	asm = NewAssembler([]string{"LOADM R1 1"})
	bytecode, err = asm.Assemble()
	if err != nil {
		t.Fatalf("Assemble failed: %s", err.Error())
	}

	if len(bytecode) != 4 {
		t.Errorf("RA: LOADM R1 1 Expected 4 bytes, got %d", len(bytecode))
	}

	// AV
	asm = NewAssembler([]string{"STORE 1 2"})
	bytecode, err = asm.Assemble()
	if err != nil {
		t.Fatalf("Assemble failed: %s", err.Error())
	}

	if len(bytecode) != 5 {
		t.Errorf("AV: STORE 1 2 Expected 5 bytes, got %d", len(bytecode))
	}

	// AL
	asm = NewAssembler([]string{"label:", "JMP label"})
	bytecode, err = asm.Assemble()
	if err != nil {
		t.Fatalf("Assemble failed: %s", err.Error())
	}

	if len(bytecode) != 3 {
		t.Errorf("AL: JMP label Expected 3 bytes, got %d", len(bytecode))
	}

	// A
	asm = NewAssembler([]string{"JMP 5"})
	bytecode, err = asm.Assemble()
	if err != nil {
		t.Fatalf("Assemble failed: %s", err.Error())
	}

	if len(bytecode) != 3 {
		t.Errorf("A: JMP 5 Expected 3 bytes, got %d", len(bytecode))
	}

	// V
	asm = NewAssembler([]string{"PUSH 1"})
	bytecode, err = asm.Assemble()
	if err != nil {
		t.Fatalf("Assemble failed: %s", err.Error())
	}

	if len(bytecode) != 3 {
		t.Errorf("V: PUSH 1 Expected 3 bytes, got %d", len(bytecode))
	}

	// R
	asm = NewAssembler([]string{"POP R1"})
	bytecode, err = asm.Assemble()
	if err != nil {
		t.Fatalf("Assemble failed: %s", err.Error())
	}

	if len(bytecode) != 2 {
		t.Errorf("R: POP R1 Expected 2 bytes, got %d", len(bytecode))
	}

	// NONE
	asm = NewAssembler([]string{"POP"})
	bytecode, err = asm.Assemble()
	if err != nil {
		t.Fatalf("Assemble failed: %s", err.Error())
	}

	if len(bytecode) != 1 {
		t.Errorf("NONE: POP Expected 1 byte, got %d", len(bytecode))
	}
}

func TestPushRBytes(t *testing.T) {
	asm := NewAssembler([]string{
		"DEC R1",
		"PUSH R1",
		"PUSH 1",
	})
	bytecode, err := asm.Assemble()
	if err != nil {
		t.Fatalf("Assemble failed: %s", err.Error())
	}

	if bytecode[0] != uint8(OP_DEC_R) {
		t.Errorf("Expected DEC R1 to be %d, got %d", OP_DEC_R, bytecode[0])
	}

	if bytecode[1] != 1 {
		t.Errorf("Expected DEC R1 to be 1, got %d", bytecode[1])
	}

	if bytecode[2] != uint8(OP_PUSH_R) {
		t.Errorf("Expected PUSH R1 to be %d, got %d", OP_PUSH_R, bytecode[2])
	}

	if bytecode[3] != 1 {
		t.Errorf("Expected PUSH R1 to be 1, got %d", bytecode[3])
	}

	if bytecode[4] != uint8(OP_PUSH_V) {
		t.Errorf("Expected PUSH 1 to be %d, got %d", OP_PUSH_V, bytecode[4])
	}

	if bytecode[5] != 0 {
		t.Errorf("Expected PUSH 1 byte 1 to be 1, got %d", bytecode[5])
	}

	if bytecode[6] != 1 {
		t.Errorf("Expected PUSH 1 byte 2 to be 1, got %d", bytecode[6])
	}
}

func TestAssembler(t *testing.T) {
	program := []string{
		"LOAD R1 1",
		"prelabel:",
		"LOAD R1 R2",
		"LOADM R1 1",
		"STORE R1 1",
		"STORE 1 2",
		"STORE R1 R2",
		"ADD R1 R2",
		"ADD R1 1",
		"SUB R1 R2",
		"SUB R1 1",
		"MUL R1 R2",
		"MUL R1 1",
		"DIV R1 R2",
		"DIV R1 1",
		"MOD R1 R2",
		"MOD R1 1",
		"AND R1 R2",
		"AND R1 1",
		"OR R1 R2",
		"OR R1 1",
		"XOR R1 R2",
		"XOR R1 1",
		"NOT R1",
		"SHL R1",
		"SHR R1",
		"INC R1",
		"DEC R1",
		"PUSH R1",
		"PUSH 1",
		"POP",
		"POP R1",
		"CMP R1 R2",
		"CMP R1 1",
		"JMP prelabel",
		"JMP postlabel",
		"JMP 5",
		"JMP R1",
		"postlabel:",
		"JE prelabel",
		"JE 5",
		"JE R1",
		"JNE prelabel",
		"JNE 5",
		"JNE R1",
		"JG prelabel",
		"JG 5",
		"JG R1",
		"JGE prelabel",
		"JGE 5",
		"JGE R1",
		"JL prelabel",
		"JL 5",
		"JL R1",
		"JLE prelabel",
		"JLE 5",
		"JLE R1",
		"CALL prelabel",
		"CALL 5",
		"CALL R1",
		"RET",
		"PRINT 1",
		"PRINT R1",
		"PRINTS 10",
		"HLT",
	}

	asm := NewAssembler(program)

	bytecode, err := asm.Assemble()
	if err != nil {
		t.Errorf("Assemble failed: %s", err.Error())
	}

	prelabelAddress := uint16(4 + CodeMemoryStart)
	postlabelAddress := uint16(115 + CodeMemoryStart)
	prelabelAddressBytes := uint16ToBytes(prelabelAddress)
	postlabelAddressBytes := uint16ToBytes(postlabelAddress)

	expected := []uint8{
		uint8(OP_LOAD_RV), 1, 0, 1,
		uint8(OP_LOAD_RR), 1, 2,
		uint8(OP_LOADM_RA), 1, 0, 1,
		uint8(OP_STORE_RA), 1, 0, 1,
		uint8(OP_STORE_AV), 0, 1, 0, 2,
		uint8(OP_STORE_RR), 1, 2,
		uint8(OP_ADD_RR), 1, 2,
		uint8(OP_ADD_RV), 1, 0, 1,
		uint8(OP_SUB_RR), 1, 2,
		uint8(OP_SUB_RV), 1, 0, 1,
		uint8(OP_MUL_RR), 1, 2,
		uint8(OP_MUL_RV), 1, 0, 1,
		uint8(OP_DIV_RR), 1, 2,
		uint8(OP_DIV_RV), 1, 0, 1,
		uint8(OP_MOD_RR), 1, 2,
		uint8(OP_MOD_RV), 1, 0, 1,
		uint8(OP_AND_RR), 1, 2,
		uint8(OP_AND_RV), 1, 0, 1,
		uint8(OP_OR_RR), 1, 2,
		uint8(OP_OR_RV), 1, 0, 1,
		uint8(OP_XOR_RR), 1, 2,
		uint8(OP_XOR_RV), 1, 0, 1,
		uint8(OP_NOT_R), 1,
		uint8(OP_SHL_R), 1,
		uint8(OP_SHR_R), 1,
		uint8(OP_INC_R), 1,
		uint8(OP_DEC_R), 1,
		uint8(OP_PUSH_R), 1,
		uint8(OP_PUSH_V), 0, 1,
		uint8(OP_POP_NONE),
		uint8(OP_POP_R), 1,
		uint8(OP_CMP_RR), 1, 2,
		uint8(OP_CMP_RV), 1, 0, 1,
		uint8(OP_JMP_A), prelabelAddressBytes[0], prelabelAddressBytes[1],
		uint8(OP_JMP_A), postlabelAddressBytes[0], postlabelAddressBytes[1],
		uint8(OP_JMP_A), 0, 5,
		uint8(OP_JMP_R), 1,
		uint8(OP_JE_A), prelabelAddressBytes[0], prelabelAddressBytes[1],
		uint8(OP_JE_A), 0, 5,
		uint8(OP_JE_R), 1,
		uint8(OP_JNE_A), prelabelAddressBytes[0], prelabelAddressBytes[1],
		uint8(OP_JNE_A), 0, 5,
		uint8(OP_JNE_R), 1,
		uint8(OP_JG_A), prelabelAddressBytes[0], prelabelAddressBytes[1],
		uint8(OP_JG_A), 0, 5,
		uint8(OP_JG_R), 1,
		uint8(OP_JGE_A), prelabelAddressBytes[0], prelabelAddressBytes[1],
		uint8(OP_JGE_A), 0, 5,
		uint8(OP_JGE_R), 1,
		uint8(OP_JL_A), prelabelAddressBytes[0], prelabelAddressBytes[1],
		uint8(OP_JL_A), 0, 5,
		uint8(OP_JL_R), 1,
		uint8(OP_JLE_A), prelabelAddressBytes[0], prelabelAddressBytes[1],
		uint8(OP_JLE_A), 0, 5,
		uint8(OP_JLE_R), 1,
		uint8(OP_CALL_A), prelabelAddressBytes[0], prelabelAddressBytes[1],
		uint8(OP_CALL_A), 0, 5,
		uint8(OP_CALL_R), 1,
		uint8(OP_RET_NONE),
		uint8(OP_PRINT_V), 0, 1,
		uint8(OP_PRINT_R), 1,
		uint8(OP_PRINTS_A), 0, 10,
		uint8(OP_HLT_NONE),
	}

	for i := range bytecode {
		if bytecode[i] != expected[i] {
			t.Errorf("Expected bytecode at %d to be %d, got %d", i, expected[i], bytecode[i])
		}
	}
}
