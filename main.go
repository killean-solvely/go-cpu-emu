package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"cpu/cpu"
)

func main() {
	// Define a flag for the file name
	fileName := flag.String("file", "", "Path to the file to be loaded")

	// Parse the flags
	flag.Parse()

	// Check if the file flag was provided
	if *fileName == "" {
		log.Fatal("Please provide a file name using the -file flag")
	}

	// Open and read the file
	data, err := os.ReadFile(*fileName)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Split the file contents by new line
	lines := strings.Split(string(data), "\n")

	asm := cpu.NewAssembler(lines)

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
