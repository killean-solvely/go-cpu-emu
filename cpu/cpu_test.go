package cpu

import (
	"testing"
)

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

	mem.WriteByte(CodeMemoryStart+0, uint8(OP_LOAD_RV))
	mem.WriteByte(CodeMemoryStart+1, 0)
	mem.WriteByte(CodeMemoryStart+2, 0)
	mem.WriteByte(CodeMemoryStart+3, 42)

	mem.WriteByte(CodeMemoryStart+4, uint8(OP_LOAD_RV))
	mem.WriteByte(CodeMemoryStart+5, 1)
	mem.WriteByte(CodeMemoryStart+6, 0)
	mem.WriteByte(CodeMemoryStart+7, 10)

	mem.WriteByte(CodeMemoryStart+8, uint8(OP_ADD_RR))
	mem.WriteByte(CodeMemoryStart+9, 0)
	mem.WriteByte(CodeMemoryStart+10, 1)

	mem.WriteByte(CodeMemoryStart+11, uint8(OP_HLT_NONE))

	cpu.Execute(mem)

	if cpu.Registers[0] != 52 {
		t.Errorf("Expected register 0 to be 52, got %d", cpu.Registers[0])
	}

	if cpu.Registers[1] != 10 {
		t.Errorf("Expected register 1 to be 10, got %d", cpu.Registers[1])
	}
}

func TestLoadRV(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}
	cpu.Execute(mem)

	if cpu.Registers[0] != 42 {
		t.Errorf("Expected register 0 to be 42, got %d", cpu.Registers[0])
	}
}

func TestLoadRR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"LOAD R1 R0",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}
	cpu.Execute(mem)

	if cpu.Registers[0] != 42 {
		t.Errorf("Expected register 0 to be 42, got %d", cpu.Registers[0])
	}

	if cpu.Registers[1] != 42 {
		t.Errorf("Expected register 1 to be 42, got %d", cpu.Registers[1])
	}
}

func TestLoadMRA(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOADM R0 0",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	mem.WriteWord(0, 42)
	cpu.Execute(mem)

	if cpu.Registers[0] != 42 {
		t.Errorf("Expected register 0 to be 42, got %d", cpu.Registers[0])
	}
}

func TestStoreRA(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"STORE R0 0",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if mem.ReadWord(0) != 42 {
		t.Errorf("Expected memory address 0 to be 42, got %d", mem.ReadByte(0))
	}
}

func TestStoreAV(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"STORE 0 42",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if mem.ReadWord(0) != 42 {
		t.Errorf("Expected memory address 0 to be 42, got %d", mem.ReadByte(0))
	}
}

func TestStoreRR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 0",
		"LOAD R1 42",
		"STORE R0 R1",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if mem.ReadWord(0) != 42 {
		t.Errorf("Expected memory address 0 to be 42, got %d", mem.ReadByte(0))
	}
}

func TestAddRR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"LOAD R1 10",
		"ADD R0 R1",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 52 {
		t.Errorf("Expected register 0 to be 52, got %d", cpu.Registers[0])
	}
}

func TestAddRV(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"ADD R0 10",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 52 {
		t.Errorf("Expected register 0 to be 52, got %d", cpu.Registers[0])
	}
}

func TestSubRR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"LOAD R1 10",
		"SUB R0 R1",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 32 {
		t.Errorf("Expected register 0 to be 32, got %d", cpu.Registers[0])
	}
}

func TestSubRV(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"SUB R0 10",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 32 {
		t.Errorf("Expected register 0 to be 32, got %d", cpu.Registers[0])
	}
}

func TestMulRR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 5",
		"LOAD R1 3",
		"MUL R0 R1",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 15 {
		t.Errorf("Expected register 0 to be 420, got %d", cpu.Registers[0])
	}
}

func TestMulRV(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 5",
		"MUL R0 3",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 15 {
		t.Errorf("Expected register 0 to be 420, got %d", cpu.Registers[0])
	}
}

func TestDivRR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 15",
		"LOAD R1 3",
		"DIV R0 R1",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 5 {
		t.Errorf("Expected register 0 to be 5, got %d", cpu.Registers[0])
	}
}

func TestDivRV(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 15",
		"DIV R0 3",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 5 {
		t.Errorf("Expected register 0 to be 5, got %d", cpu.Registers[0])
	}
}

func TestModRR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 14",
		"LOAD R1 3",
		"MOD R0 R1",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 2 {
		t.Errorf("Expected register 0 to be 0, got %d", cpu.Registers[0])
	}
}

func TestModRV(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 14",
		"MOD R0 3",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 2 {
		t.Errorf("Expected register 0 to be 0, got %d", cpu.Registers[0])
	}
}

func TestAndRR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 170",
		"LOAD R1 204",
		"AND R0 R1",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 136 {
		t.Errorf("Expected register 0 to be 136, got %d", cpu.Registers[0])
	}
}

func TestAndRV(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 170",
		"AND R0 204",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 136 {
		t.Errorf("Expected register 0 to be 136, got %d", cpu.Registers[0])
	}
}

func TestOrRR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 170",
		"LOAD R1 204",
		"OR R0 R1",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 238 {
		t.Errorf("Expected register 0 to be 238, got %d", cpu.Registers[0])
	}
}

func TestOrRV(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 170",
		"OR R0 204",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 238 {
		t.Errorf("Expected register 0 to be 238, got %d", cpu.Registers[0])
	}
}

func TestXorRR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 170",
		"LOAD R1 204",
		"XOR R0 R1",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 102 {
		t.Errorf("Expected register 0 to be 102, got %d", cpu.Registers[0])
	}
}

func TestXorRV(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 170",
		"XOR R0 204",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 102 {
		t.Errorf("Expected register 0 to be 102, got %d", cpu.Registers[0])
	}
}

func TestNotR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 170",
		"NOT R0",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 65365 {
		t.Errorf("Expected register 0 to be 65365, got %d", cpu.Registers[0])
	}
}

func TestShlR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 170",
		"SHL R0",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 340 {
		t.Errorf("Expected register 0 to be 340, got %d", cpu.Registers[0])
	}
}

func TestShrR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 170",
		"SHR R0",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 85 {
		t.Errorf("Expected register 0 to be 85, got %d", cpu.Registers[0])
	}
}

func TestIncR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 170",
		"INC R0",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 171 {
		t.Errorf("Expected register 0 to be 171, got %d", cpu.Registers[0])
	}
}

func TestDecR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 170",
		"DEC R0",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 169 {
		t.Errorf("Expected register 0 to be 169, got %d", cpu.Registers[0])
	}
}

func TestPushR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"PUSH R0",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if len(cpu.Stack.Data) != 1 {
		t.Errorf("Expected stack to have 1 item, got %d", len(cpu.Stack.Data))
	}

	if cpu.Stack.Pop() != 42 {
		t.Errorf("Expected stack to pop 42, got %d", cpu.Stack.Pop())
	}
}

func TestPushV(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"PUSH 42",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if len(cpu.Stack.Data) != 1 {
		t.Errorf("Expected stack to have 1 item, got %d", len(cpu.Stack.Data))
	}

	if cpu.Stack.Pop() != 42 {
		t.Errorf("Expected stack to pop 42, got %d", cpu.Stack.Pop())
	}
}

func TestPopNone(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"PUSH 42",
		"POP",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if len(cpu.Stack.Data) != 0 {
		t.Errorf("Expected stack to have 0 items, got %d", len(cpu.Stack.Data))
	}
}

func TestPopR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"PUSH 42",
		"POP R0",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Registers[0] != 42 {
		t.Errorf("Expected register 0 to be 42, got %d", cpu.Registers[0])
	}
}

func TestCmpRR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"LOAD R1 42",
		"CMP R0 R1",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Flags.Equal == 0 {
		t.Errorf("Expected registers to be equal")
	}

	if cpu.Flags.Greater == 1 {
		t.Errorf("Expected registers to not be greater")
	}

	if cpu.Flags.Less == 1 {
		t.Errorf("Expected registers to not be less")
	}

	cpu, mem, err = prepCpuAndMem([]string{
		"LOAD R0 42",
		"LOAD R1 43",
		"CMP R0 R1",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Flags.Equal != 0 {
		t.Errorf("Expected registers to not be equal")
	}

	if cpu.Flags.Greater != 0 {
		t.Errorf("Expected registers to not be greater")
	}

	if cpu.Flags.Less != 1 {
		t.Errorf("Expected registers to be less")
	}

	cpu, mem, err = prepCpuAndMem([]string{
		"LOAD R0 43",
		"LOAD R1 42",
		"CMP R0 R1",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Flags.Equal != 0 {
		t.Errorf("Expected registers to not be equal")
	}

	if cpu.Flags.Greater != 1 {
		t.Errorf("Expected registers to be greater")
	}

	if cpu.Flags.Less != 0 {
		t.Errorf("Expected registers to not be less")
	}

	cpu, mem, err = prepCpuAndMem([]string{
		"LOAD R0 42",
		"LOAD R1 42",
		"CMP R0 R1",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)
}

func TestCmpRV(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"CMP R0 42",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Flags.Equal == 0 {
		t.Errorf("Expected registers to be equal")
	}

	if cpu.Flags.Greater == 1 {
		t.Errorf("Expected registers to not be greater")
	}

	if cpu.Flags.Less == 1 {
		t.Errorf("Expected registers to not be less")
	}

	cpu, mem, err = prepCpuAndMem([]string{
		"LOAD R0 42",
		"CMP R0 43",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Flags.Equal != 0 {
		t.Errorf("Expected registers to not be equal")
	}

	if cpu.Flags.Greater != 0 {
		t.Errorf("Expected registers to not be greater")
	}

	if cpu.Flags.Less != 1 {
		t.Errorf("Expected registers to be less")
	}

	cpu, mem, err = prepCpuAndMem([]string{
		"LOAD R0 43",
		"CMP R0 42",
		"HLT",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.Execute(mem)

	if cpu.Flags.Equal != 0 {
		t.Errorf("Expected registers to not be equal")
	}

	if cpu.Flags.Greater != 1 {
		t.Errorf("Expected registers to be greater")
	}

	if cpu.Flags.Less != 0 {
		t.Errorf("Expected registers to not be less")
	}
}

func TestJmpA(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"JMP 5",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}
}

func TestJmpR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 5",
		"JMP R0",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}
}

func TestJeA(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"LOAD R1 42",
		"CMP R0 R1",
		"JE 5",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}
}

func TestJeR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"LOAD R1 42",
		"CMP R0 R1",
		"LOAD R2 5",
		"JE R2",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}
}

func TestJneA(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"LOAD R1 43",
		"CMP R0 R1",
		"JNE 5",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}
}

func TestJneR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"LOAD R1 43",
		"CMP R0 R1",
		"LOAD R2 5",
		"JNE R2",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}
}

func TestJgR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 43",
		"LOAD R1 42",
		"CMP R0 R1",
		"LOAD R2 5",
		"JG R2",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}
}

func TestJgA(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 43",
		"LOAD R1 42",
		"CMP R0 R1",
		"JG 5",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}
}

func TestJgeR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"LOAD R1 42",
		"CMP R0 R1",
		"LOAD R2 5",
		"JGE R2",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}

	cpu, mem, err = prepCpuAndMem([]string{
		"LOAD R0 43",
		"LOAD R1 42",
		"CMP R0 R1",
		"LOAD R2 5",
		"JGE R2",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}
}

func TestJgeA(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"LOAD R1 42",
		"CMP R0 R1",
		"JGE 5",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}

	cpu, mem, err = prepCpuAndMem([]string{
		"LOAD R0 43",
		"LOAD R1 42",
		"CMP R0 R1",
		"JGE 5",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}
}

func TestJlR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"LOAD R1 43",
		"CMP R0 R1",
		"LOAD R2 5",
		"JL R2",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}
}

func TestJlA(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"LOAD R1 43",
		"CMP R0 R1",
		"JL 5",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}
}

func TestJleR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"LOAD R1 42",
		"CMP R0 R1",
		"LOAD R2 5",
		"JLE R2",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}

	cpu, mem, err = prepCpuAndMem([]string{
		"LOAD R0 41",
		"LOAD R1 42",
		"CMP R0 R1",
		"LOAD R2 5",
		"JLE R2",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}
}

func TestJleA(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 42",
		"LOAD R1 42",
		"CMP R0 R1",
		"JLE 5",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}

	cpu, mem, err = prepCpuAndMem([]string{
		"LOAD R0 41",
		"LOAD R1 42",
		"CMP R0 R1",
		"JLE 5",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}
}

func TestCallA(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"CALL 5",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}

	if len(cpu.Stack.Data) != 1 {
		t.Errorf("Expected stack to have 1 item, got %d", len(cpu.Stack.Data))
	}

	if cpu.Stack.Pop() != CodeMemoryStart+3 {
		t.Errorf("Expected return address to be pushed to stack")
	}
}

func TestCallR(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"LOAD R0 5",
		"CALL R0",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != 5 {
		t.Errorf("Expected program counter to be 5, got %d", cpu.ProgramCounter)
	}

	if len(cpu.Stack.Data) != 1 {
		t.Errorf("Expected stack to have 1 item, got %d", len(cpu.Stack.Data))
	}

	if cpu.Stack.Pop() != CodeMemoryStart+6 {
		t.Errorf("Expected return address to be pushed to stack")
	}
}

func TestRet(t *testing.T) {
	cpu, mem, err := prepCpuAndMem([]string{
		"CALL 2051",
		"RET",
	})
	if err != nil {
		t.Fatalf("Error preparing CPU and memory: %s", err)
	}

	cpu.ProgramCounter = CodeMemoryStart
	cpu.executeNext(mem)
	cpu.executeNext(mem)

	if cpu.ProgramCounter != CodeMemoryStart+3 {
		t.Errorf("Expected program counter to be %d, got %d", CodeMemoryStart+3, cpu.ProgramCounter)
	}
}
