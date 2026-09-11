package assembler

import "testing"

func TestParserInstructions(t *testing.T) {
	tb := []struct {
		text        string
		src0Type    ParameterType
		src1Type    ParameterType
		destType    ParameterType
		shouldPanic bool
	}{
		{
			"MOV r0, r1",
			ParameterTypeNone,
			ParameterTypeRegister,
			ParameterTypeRegister,
			false,
		},
		{
			"MOV r0, 0x001",
			ParameterTypeNone,
			ParameterTypeLiteral,
			ParameterTypeRegister,
			false,
		},
		{
			"ADD r0, r1",
			ParameterTypeRegister,
			ParameterTypeRegister,
			ParameterTypeRegister,
			false,
		},
		{
			"ADD r0, r0, r1",
			ParameterTypeRegister,
			ParameterTypeRegister,
			ParameterTypeRegister,
			false,
		},
		{
			"ADD r0, r0, 0x0",
			ParameterTypeRegister,
			ParameterTypeLiteral,
			ParameterTypeRegister,
			false,
		},
		{
			"RET",
			ParameterTypeNone,
			ParameterTypeNone,
			ParameterTypeNone,
			false,
		},
		{
			"PUSH r0",
			ParameterTypeNone,
			ParameterTypeRegister,
			ParameterTypeNone,
			false,
		},
		{
			"PUSH 0x01",
			ParameterTypeNone,
			ParameterTypeLiteral,
			ParameterTypeNone,
			false,
		},
		{
			"POP r0",
			ParameterTypeNone,
			ParameterTypeNone,
			ParameterTypeRegister,
			false,
		},
		{
			"ADD r0, r0, r0, r0",
			ParameterTypeNone,
			ParameterTypeNone,
			ParameterTypeNone,
			true,
		},
		{
			"RET r0",
			ParameterTypeNone,
			ParameterTypeNone,
			ParameterTypeNone,
			true,
		},
		{
			"PUSH [r0]",
			ParameterTypeNone,
			ParameterTypeNone,
			ParameterTypeNone,
			true,
		},
		{
			"PUSH r0, r0",
			ParameterTypeNone,
			ParameterTypeNone,
			ParameterTypeNone,
			true,
		},
		{
			"STR 0x00, 0x01010",
			ParameterTypeNone,
			ParameterTypeNone,
			ParameterTypeNone,
			true,
		},
	}

	for _, s := range tb {
		t.Run(s.text, func(t *testing.T) {
			defer testPanic(t, s.shouldPanic)

			parser := NewParser(s.text)
			fragment := parser.Instruction()

			if fragment.Src0.t != s.src0Type {
				t.Errorf("expected src0 parameter type of %v but got %v", fragment.Src0.t, s.src0Type)
			}

			if fragment.Src1.t != s.src1Type {
				t.Errorf("expected src1 parameter type of %v but got %v", fragment.Src1.t, s.src1Type)
			}

			if fragment.Dest.t != s.destType {
				t.Errorf("expected dest parameter type of %v but got %v", fragment.Dest.t, s.destType)
			}
		})
	}
}
