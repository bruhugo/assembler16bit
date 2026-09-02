package assembler

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isAlnum(r rune) bool {
	return isDigit(r) || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'z')
}
