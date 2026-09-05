package assembler

import (
	"fmt"
	"testing"
)

func TestTokenTypes(t *testing.T) {
	tb := []struct {
		program     string
		types       []TokenType
		shouldPanic bool
	}{
		{
			"MOV",
			[]TokenType{TokenTypeInstruction},
			false,
		},
		{
			";",
			[]TokenType{TokenTypeColon},
			false,
		},
		{
			".",
			[]TokenType{TokenTypeDot},
			false,
		},
		{
			",",
			[]TokenType{TokenTypeComa},
			false,
		},
		{
			"customlabel",
			[]TokenType{TokenTypeLabel},
			false,
		},
		{
			"r0",
			[]TokenType{TokenTypeParameter},
			false,
		},
		{
			"in0",
			[]TokenType{TokenTypeParameter},
			false,
		},
		{
			"out0",
			[]TokenType{TokenTypeParameter},
			false,
		},
		{
			"r9",
			[]TokenType{TokenTypeLabel},
			false,
		},
		{
			"r9",
			[]TokenType{TokenTypeLabel},
			false,
		},
		{
			"MOV r0, r0",
			[]TokenType{TokenTypeInstruction, TokenTypeParameter, TokenTypeComa, TokenTypeParameter},
			false,
		},
		{
			"-",
			[]TokenType{},
			true,
		},
		{
			"; comment",
			[]TokenType{},
			false,
		},
		{
			"MOV ; comment",
			[]TokenType{TokenTypeInstruction},
			false,
		},
		{
			"; comment MOV",
			[]TokenType{},
			false,
		},
		{
			"; comment \n MOV",
			[]TokenType{TokenTypeNewLine, TokenTypeInstruction},
			false,
		},
		{
			"MOV",
			[]TokenType{TokenTypeNewLine},
			true,
		},
	}

	for _, test := range tb {
		t.Run(fmt.Sprintf("Testing text: %s", test.program), func(t *testing.T) {
			t.Parallel()

			defer testPanic(t, test.shouldPanic)

			tokenizer := NewTokenizer(test.program)
			for i := 0; tokenizer.CurType() != TokenTypeEOF; i++ {
				tokenizer.AssertAndNext(test.types[i])
			}
		})
	}
}
