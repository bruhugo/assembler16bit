package main

import (
	"io"
	"os"

	assembler "assembler/internal"
)

func main() {
	// TODO: make commands better
	inputFilename := os.Args[1]
	outputFilename := os.Args[2]

	inputFile, err := os.OpenFile(inputFilename, os.O_RDONLY, 0o600)
	if err != nil {
		panic(err)
	}

	program, err := io.ReadAll(inputFile)
	if err != nil {
		panic(err)
	}

	outputFile, err := os.OpenFile(outputFilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
	if err != nil {
		panic(err)
	}

	as := assembler.NewAssembler(string(program))
	as.Assemble(outputFile, assembler.AssembleOptionLogisim)
}
