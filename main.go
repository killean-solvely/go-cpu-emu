package main

import (
	"fmt"

	"cpu/cpu"
)

func main() {
	program := []string{
		"STORE_VAL 0 H",
		"STORE_VAL 1 E",
		"STORE_VAL 2 L",
		"STORE_VAL 3 L",
		"STORE_VAL 4 O",
		"STORE_VAL 5 W",
		"STORE_VAL 6 O",
		"STORE_VAL 7 R",
		"STORE_VAL 8 L",
		"STORE_VAL 9 D",
		"STORE_VAL 10 0",
		"PRINT_MEM 0",
		"HLT",
	}

	asm := cpu.NewAssembler(program)

	bytecode, err := asm.Assemble()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cpuInstance := cpu.NewCPU()
	memory := cpu.NewMemory()
	memory.LoadCode(bytecode)

	cpuInstance.Execute(memory)
}
