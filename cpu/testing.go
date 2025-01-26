package cpu

func prepCpuAndMem(code []string) (cpu *CPU, mem *Memory, err error) {
	asm := NewAssembler(code)
	bytecode, err := asm.Assemble()
	if err != nil {
		return nil, nil, err
	}

	cpu = NewCPU()
	mem = NewMemory()
	mem.LoadCode(bytecode)

	return cpu, mem, nil
}
