package cpu

import "testing"

func TestNewCPU(t *testing.T) {
	cpu := NewCPU()

	for i, reg := range cpu.Registers {
		if reg != 0 {
			t.Errorf("Expected register %d to be 0, got %d", i, reg)
		}
	}

	if cpu.ProgramCounter != 0 {
		t.Errorf("Expected program counter to be 0, got %d", cpu.ProgramCounter)
	}
}

func TestCPUExecution(t *testing.T) {
	cpu := NewCPU()
	mem := NewMemory()

	mem.Write(CodeMemoryStart+0, uint8(OP_LOAD_RV))
	mem.Write(CodeMemoryStart+1, 0)
	mem.Write(CodeMemoryStart+2, 42)

	mem.Write(CodeMemoryStart+3, uint8(OP_LOAD_RV))
	mem.Write(CodeMemoryStart+4, 1)
	mem.Write(CodeMemoryStart+5, 10)

	mem.Write(CodeMemoryStart+6, uint8(OP_ADD_RR))
	mem.Write(CodeMemoryStart+7, 0)
	mem.Write(CodeMemoryStart+8, 1)

	mem.Write(CodeMemoryStart+9, uint8(OP_HLT_NONE))

	cpu.Execute(mem)

	if cpu.Registers[0] != 52 {
		t.Errorf("Expected register 0 to be 52, got %d", cpu.Registers[0])
	}

	if cpu.Registers[1] != 10 {
		t.Errorf("Expected register 1 to be 10, got %d", cpu.Registers[1])
	}
}

func TestPop(t *testing.T) {
	asm := NewAssembler([]string{
		"LOAD R0 1",
		"LOAD R1 2",
		"LOAD R2 3",
		"LOAD R3 4",

		"PUSH R0",
		"PUSH R1",
		"PUSH R2",
		"PUSH R3",

		"POP R3",
		"POP R2",
		"POP R1",
		"POP R0",

		"HLT",
	})
	bytecode, err := asm.Assemble()
	if err != nil {
		t.Fatalf("Failed to assemble code: %v", err)
	}

	cpu := NewCPU()
	mem := NewMemory()
	mem.LoadCode(bytecode)
	cpu.Execute(mem)

	if cpu.Registers[0] != 1 {
		t.Errorf("Expected register 0 to be 1, got %d", cpu.Registers[0])
	}

	if cpu.Registers[1] != 2 {
		t.Errorf("Expected register 1 to be 2, got %d", cpu.Registers[1])
	}

	if cpu.Registers[2] != 3 {
		t.Errorf("Expected register 2 to be 3, got %d", cpu.Registers[2])
	}

	if cpu.Registers[3] != 4 {
		t.Errorf("Expected register 3 to be 4, got %d", cpu.Registers[3])
	}
}
