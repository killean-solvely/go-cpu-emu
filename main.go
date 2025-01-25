package main

import (
	"fmt"

	"cpu/cpu"
)

func example(a int, b int) int {
	return a + b
}

func main() {
	// var a uint8 = 10
	// var b uint8 = 27
	// var c uint8 = 38
	// var d uint8 = 8
	// var e uint8 = 123
	// var f uint8 = a + b + c - d/e
	// fmt.Println(f)

	program := []string{
		"LOAD R0 8",
		"LOAD R1 123",

		"DIV R0 R1",
		"LOAD R1 10",
		"LOAD R2 27",
		"ADD R1 R2",
		"LOAD R2 38",
		"ADD R1 R2",
		"SUB R1 R0",

		"CMP_REG_VAL R1 75",
		"JE success",
		"JMP fail",

		"success:",
		"LOAD R0 1",
		"PRINT_REG R0",
		"JMP end",

		"fail:",
		"LOAD R0 0",
		"PRINT_REG R0",
		"JMP end",

		"end:",
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
