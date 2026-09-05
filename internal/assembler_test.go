package assembler

import (
	"bytes"
	"testing"
)

func TestLogisimOptionHasRightHeader(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	as := NewAssembler("")
	as.Assemble(buffer, AssembleOptionLogisim)

	if buffer.String() != LogisimHeader {
		t.Errorf("expected header '%s', but got '%s'", LogisimHeader, buffer.String())
	}
}
