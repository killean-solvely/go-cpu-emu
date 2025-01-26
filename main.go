package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"cpu/cpu"
)

func main() {
	// Define a flag for the file name
	fileName := flag.String("f", "", "Path to the file to be loaded")
	toCompile := flag.Bool("c", false, "Compile the file")
	outputFileName := flag.String("o", "", "Output file name")
	toRun := flag.Bool("r", false, "Run the compiled file")

	// Parse the flags
	flag.Parse()

	if !*toCompile && !*toRun {
		log.Fatal("Please provide a flag to either compile or run the file")
	}

	if *toCompile && *toRun {
		log.Fatal("Please provide only one flag to either compile or run the file")
	}

	if *fileName == "" {
		log.Fatal("Please provide a file name using the -f flag")
	}

	if *toCompile {
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
			log.Fatalf("Failed to assemble code: %v", err)
		}

		outputFileName := *outputFileName
		if outputFileName == "" {
			outputFileName = *fileName + ".bin"
		}

		err = os.WriteFile(outputFileName, bytecode, 0644)
		if err != nil {
			log.Fatalf("Failed to write file: %v", err)
		}

		log.Printf("File compiled successfully: %s", outputFileName)

		return
	}

	if *toRun {
		bytecode, err := os.ReadFile(*fileName)
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}

		cpuInstance := cpu.NewCPU()
		memory := cpu.NewMemory()
		memory.LoadCode(bytecode)

		cpuInstance.Execute(memory)

		return
	}
}
