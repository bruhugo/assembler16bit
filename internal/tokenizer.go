package assembler

import (
	"fmt"
	"strings"
)

type Tokenizer struct {
	line     int
	column   int
	i        int
	r        []rune
	curToken Token
}

func NewTokenizer(program string) *Tokenizer {
	return &Tokenizer{
		r:    []rune(program),
		line: 1,
	}
}

// NewToken fetches the next token and sets int
// as the new current one
func (t *Tokenizer) NewToken() Token {
	t.skipSpaces()
	for !t.isDone() && t.curRune() == ';' {
		t.skipComment()
		t.skipSpaces()
	}

	if t.isDone() {
		t.curToken = &SimpleToken{
			t: TokenTypeEOF,
		}
		return t.curToken

	}

	var token Token
	r := t.curRune()
	switch {
	case isDigit(r):
		token = &TokenParameter{
			val:       uint16(t.getNumber()),
			paramType: ParameterTypeLiteral,
		}
	case r == ',':
		t.i++
		token = &SimpleToken{
			t: TokenTypeComa,
		}
	case r == '.':
		t.i++
		token = &SimpleToken{
			t: TokenTypeDot,
		}
	case isAlnum(r):
		token = t.getWord()
	case r == '\n':
		t.line++
		t.column = 0
		t.i++
		token = &SimpleToken{
			t: TokenTypeNewLine,
		}
	default:
		panic(fmt.Sprintf("unexpected token in line %d and column %d", t.line, t.column))
	}

	t.curToken = token
	return token
}

func (t *Tokenizer) AssertAndNext(tp TokenType) Token {
	if t.curToken.Type() != tp {
		panic(t.errMsg("invalid token"))
	}

	token := t.curToken
	t.NewToken()
	return token
}

func (t *Tokenizer) CurType() TokenType {
	return t.curToken.Type()
}

func (t *Tokenizer) skipSpaces() {
	for !t.isDone() && t.curRune() == ' ' {
		t.i++
	}
}

func (t *Tokenizer) isDone() bool {
	return len(t.r) <= t.i
}

func (t *Tokenizer) curRune() rune {
	return t.r[t.i]
}

func (t *Tokenizer) skipComment() {
	if t.curRune() != ';' {
		return
	}

	for !t.isDone() && t.curRune() != '\n' {
		t.i++
	}
}

func (t *Tokenizer) getNumber() uint16 {
	if t.isHex() {
		t.i += 2
		return t.getHex()
	} else if t.isBin() {
		t.i += 2
		return t.getBin()
	}

	var val uint16 = 0
	for !t.isDone() && isDigit(t.curRune()) {
		val *= 10
		val += uint16(t.curRune() - '0')
		t.i++
	}

	return val
}

func (t *Tokenizer) getHex() uint16 {
	var val uint16 = 0
	for !t.isDone() && isHexDigit(t.curRune()) {
		cur := t.curRune()
		val *= 16
		switch {
		case isDigit(cur):
			val += uint16(cur - '0')
		case cur == 'a' || cur == 'A':
			val += 10
		case cur == 'b' || cur == 'B':
			val += 11
		case cur == 'c' || cur == 'C':
			val += 12
		case cur == 'd' || cur == 'D':
			val += 13
		case cur == 'e' || cur == 'E':
			val += 14
		case cur == 'f' || cur == 'F':
			val += 15
		default:
			// should not be able to get here, but just to make sure
			panic(t.errMsg("invalid hex format"))
		}
		t.i++
	}
	return val
}

func (t *Tokenizer) getBin() uint16 {
	var val uint16

	for !t.isDone() && (t.curRune() == '1' || t.curRune() == '0') {
		val *= 2
		val += uint16(t.curRune() - '0')
		t.i++
	}

	return val
}

func (t *Tokenizer) isHex() bool {
	if len(t.r) <= t.i+1 {
		return false
	}

	return t.curRune() == '0' && t.r[t.i+1] == 'x'
}

func (t *Tokenizer) isBin() bool {
	if len(t.r) <= t.i+1 {
		return false
	}

	return t.curRune() == '0' && t.r[t.i+1] == 'b'
}

func (t *Tokenizer) getWord() Token {
	sb := strings.Builder{}
	for !t.isDone() && isAlnum(t.curRune()) {
		sb.WriteRune(t.curRune())
		t.i++
	}
	str := sb.String()

	token, ok := reservedKeywords[str]
	if ok {
		return token
	}

	return &TokenLabel{
		label: str,
	}
}

func (t *Tokenizer) errMsg(msg string) string {
	return fmt.Sprintf("line %d col %d: %s", t.line, t.column, msg)
}
