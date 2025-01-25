package main

import (
	"fmt"

	"cpu/cpu"
)

func main() {
	program := []string{
		"LOAD R0 5",
		"LOAD R1 7",
		"ADD R0 R1",
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
