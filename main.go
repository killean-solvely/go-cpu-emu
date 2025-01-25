package main

import (
	"fmt"

	"cpu/cpu"
)

func main() {
	program := []string{
		"LOAD R0 250",
		"LOAD R1 0",
		"start:",
		"CMP_REG_REG R0 R1",
		"JE end",
		"PRINT_REG R0",
		"DEC R0",
		"JMP start",
		"end:",
		"PRINT_REG R1",
		"HLT",
	}

	bytecode, err := cpu.Assemble(program)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cpuInstance := cpu.NewCPU()
	memory := cpu.NewMemory()
	memory.LoadCode(bytecode)

	cpuInstance.Execute(memory)
}
