package assembler

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestLogisimOptionHasRightHeader(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	as := NewAssembler("")
	as.Assemble(buffer, AssembleOptionLogisim, false)

	if buffer.String() != LogisimHeader {
		t.Errorf("expected header '%s', but got '%s'", LogisimHeader, buffer.String())
	}
}

func BenchmarkCompleteProgram(b *testing.B) {
	b.ReportAllocs()

	programFile, ok := os.LookupEnv("TEST_PROGRAM")
	if !ok {
		b.Skip("Must set the TEST_PROGRAM env variable to run this benchmark")
	}

	content := readProgramTestFile(programFile)

	for b.Loop() {
		assembler := NewAssembler(content)
		assembler.Assemble(io.Discard, AssembleOptionLogisim, false)
	}
}

func readProgramTestFile(filename string) string {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0o600)
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return string(content)
}
