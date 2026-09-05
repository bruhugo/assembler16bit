package main

import (
	"flag"
	"io"
	"os"

	assembler "assembler/internal"
)

var (
	inputFilename  = flag.String("i", "", "input file")
	outputFilename = flag.String("o", "", "output file")
	verbose        = flag.Bool("v", false, "verbose output")
)

func main() {
	flag.Parse()

	if inputFilename == nil || *inputFilename == "" {
		panic("Input file must be provided.")
	}

	if outputFilename == nil || *outputFilename == "" {
		panic("Output file must be provided.")
	}

	inputFile, err := os.OpenFile(*inputFilename, os.O_RDONLY, 0o600)
	if err != nil {
		panic(err)
	}

	program, err := io.ReadAll(inputFile)
	if err != nil {
		panic(err)
	}

	outputFile, err := os.OpenFile(*outputFilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
	if err != nil {
		panic(err)
	}

	as := assembler.NewAssembler(string(program))
	as.Assemble(outputFile, assembler.AssembleOptionLogisim, *verbose)
}
