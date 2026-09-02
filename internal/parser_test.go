package assembler

import (
	"reflect"
	"testing"
)

func TestParseInstruction(t *testing.T) {
	tb := []struct {
		text        string
		shouldPanic bool
		param1      ParamFragment
		param2      ParamFragment
	}{
		{
			"MOV r0, r0",
			false,
			ParamFragment{t: ParameterTypeRegister, val: 0},
			ParamFragment{t: ParameterTypeRegister, val: 0},
		},
		{
			"MOV r1, r3",
			false,
			ParamFragment{t: ParameterTypeRegister, val: 1},
			ParamFragment{t: ParameterTypeRegister, val: 3},
		},
		{
			"MOV r1, 1020",
			false,
			ParamFragment{t: ParameterTypeRegister, val: 1},
			ParamFragment{t: ParameterTypeLiteral, val: 1020},
		},
		{
			"MOV r1, 0x11",
			false,
			ParamFragment{t: ParameterTypeRegister, val: 1},
			ParamFragment{t: ParameterTypeLiteral, val: 17},
		},
		{
			"MOV out1, r0",
			false,
			ParamFragment{t: ParameterTypeOutput, val: 1},
			ParamFragment{t: ParameterTypeRegister, val: 0},
		},
		{
			"MOV out1, 0b1000",
			false,
			ParamFragment{t: ParameterTypeOutput, val: 1},
			ParamFragment{t: ParameterTypeLiteral, val: 8},
		},
		{
			"LOAD r0, 0x1000",
			false,
			ParamFragment{t: ParameterTypeRegister, val: 0},
			ParamFragment{t: ParameterTypeLiteral, val: 4096},
		},
		{
			"LOAD r0, 0xABC",
			false,
			ParamFragment{t: ParameterTypeRegister, val: 0},
			ParamFragment{t: ParameterTypeLiteral, val: 2748},
		},
		{
			"RET",
			false,
			ParamFragment{t: ParameterTypeNone},
			ParamFragment{t: ParameterTypeNone},
		},
		{
			"CALL label",
			false,
			ParamFragment{t: ParameterTypeLabel, label: "label"},
			ParamFragment{t: ParameterTypeNone},
		},
	}

	for _, test := range tb {
		t.Run(test.text, func(t *testing.T) {
			parser := NewParser(test.text)
			f := parser.Line()

			if f.Type != FragmentTypeInstruction {
				t.Errorf("wrong fragment type received")
			}

			if !reflect.DeepEqual(f.Param1, test.param1) {
				t.Errorf("params 1 are not equal. want: %v got: %v", test.param1, f.Param1)
			}

			if !reflect.DeepEqual(f.Param2, test.param2) {
				t.Errorf("params 2 are not equal. want: %v got: %v", test.param2, f.Param2)
			}

			defer func() {
				v := recover()
				if v != nil && !test.shouldPanic {
					t.Errorf("should not have panicked: %v", v)
				} else if v == nil && test.shouldPanic {
					t.Errorf("should have panicked: %v", v)
				}
			}()
		})
	}
}
