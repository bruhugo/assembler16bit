package assembler

import (
	"fmt"
	"testing"
)

func TestIsHex(t *testing.T) {
	tb := []struct {
		digit rune
		isHex bool
	}{
		{'a', true},
		{'A', true},
		{'b', true},
		{'B', true},
		{'c', true},
		{'C', true},
		{'d', true},
		{'D', true},
		{'e', true},
		{'E', true},
		{'f', true},
		{'F', true},
		{'g', false},
		{'Z', false},
		{' ', false},
		{'-', false},
		{'=', false},
		{'1', true},
		{'2', true},
	}

	for _, test := range tb {
		t.Run(fmt.Sprintf("%v", test.digit), func(t *testing.T) {
			if isHex := isHexDigit(test.digit); (!isHex && test.isHex) ||
				(isHex && !test.isHex) {
				t.Errorf("expected %v but got %v", test.isHex, isHex)
			}
		})
	}
}
