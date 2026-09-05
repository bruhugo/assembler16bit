package assembler

import "testing"

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isAlnum(r rune) bool {
	return isDigit(r) || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isHexDigit(r rune) bool {
	if r >= 'a' && r <= 'z' {
		r -= 'a' - 'A'
	}
	return isDigit(r) ||
		r == 'A' ||
		r == 'B' ||
		r == 'C' ||
		r == 'D' ||
		r == 'E' ||
		r == 'F'
}

func testPanic(t *testing.T, shouldPanic bool) {
	v := recover()
	if v != nil && !shouldPanic {
		t.Errorf("should not have panicked: %v", v)
	} else if v == nil && shouldPanic {
		t.Errorf("should have panicked: %v", v)
	}
}
