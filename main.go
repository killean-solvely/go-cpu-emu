package main

import (
	"fmt"

	"cpu/cpu"
)

func main() {
	program := []string{
		"LOAD R0 0",
		"LOAD R1 10",
		"CMP_REG_REG R0 R1",
		"JE 72",
		"PRINT_REG R0",
		"INC R0",
		"JMP 61",
		"PRINT_REG R1",
		"HLT",
	}

	bytecode, err := cpu.Assemble(program)
	if err != nil {
		fmt.Println("Assembly failed:", err)
		return
	}

	cpuInstance := cpu.NewCPU()
	memory := cpu.NewMemory()
	memory.LoadCode(bytecode)

	cpuInstance.Execute(memory)

	fmt.Printf("Register 0 contains: %d\n", cpuInstance.Registers[0])
}
