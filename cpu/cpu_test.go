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

	mem.Write(CodeMemoryStart+0, OP_LOAD)
	mem.Write(CodeMemoryStart+1, 0)
	mem.Write(CodeMemoryStart+2, 42)

	mem.Write(CodeMemoryStart+3, OP_LOAD)
	mem.Write(CodeMemoryStart+4, 1)
	mem.Write(CodeMemoryStart+5, 10)

	mem.Write(CodeMemoryStart+6, OP_ADD)
	mem.Write(CodeMemoryStart+7, 0)
	mem.Write(CodeMemoryStart+8, 1)

	mem.Write(CodeMemoryStart+9, OP_HLT)

	cpu.Execute(mem)

	if cpu.Registers[0] != 52 {
		t.Errorf("Expected register 0 to be 52, got %d", cpu.Registers[0])
	}

	if cpu.Registers[1] != 10 {
		t.Errorf("Expected register 1 to be 10, got %d", cpu.Registers[1])
	}
}
