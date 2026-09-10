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

func TestInstructionTwoParametersDestAndSrc0(t *testing.T) {
	p := NewParser("ADD r2, r1")
	frag := p.Instruction()
	if frag.Dest.val != 2 || frag.Src0.val != 2 || frag.Src1.val != 1 {
		t.Errorf("expected dest=2, src0=2, src1=1, got dest=%d, src0=%d, src1=%d", frag.Dest.val, frag.Src0.val, frag.Src1.val)
	}

	p2 := NewParser("ADD r2, r0, r1")
	frag2 := p2.Instruction()
	if frag2.Dest.val != 2 || frag2.Src0.val != 0 || frag2.Src1.val != 1 {
		t.Errorf("expected dest=2, src0=0, src1=1, got dest=%d, src0=%d, src1=%d", frag2.Dest.val, frag2.Src0.val, frag2.Src1.val)
	}

	p3 := NewParser("MOV r2, r1")
	frag3 := p3.Instruction()
	if frag3.Dest.val != 2 || frag3.Src0.val != 0 || frag3.Src1.val != 1 {
		t.Errorf("expected dest=2, src0=0, src1=1, got dest=%d, src0=%d, src1=%d", frag3.Dest.val, frag3.Src0.val, frag3.Src1.val)
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
