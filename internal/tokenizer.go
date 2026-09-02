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
	}

	r := t.curRune()
	switch {
	case isDigit(r):
		return &TokenParameter{
			val:       uint16(t.getNumber()),
			paramType: ParameterTypeLiteral,
		}
	case r == ',':
		t.i++
		return &SimpleToken{
			t: TokenTypeComa,
		}
	case r == '.':
		t.i++
		return &SimpleToken{
			t: TokenTypeDot,
		}
	case isAlnum(r):
		return t.getWord()
	case r == '\n':
		return &SimpleToken{
			t: TokenTypeNewLine,
		}
	default:
		panic(fmt.Sprintf("unexpected token in line %d and column %d", t.line, t.column))
	}
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
	return len(t.r) > t.i
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
	var base uint16 = 10

	if t.isHex() {
		t.i += 2
		base = 16
	} else if t.isBin() {
		t.i += 2
		base = 2
	}

	// in case the user types "0x" or "0b" alone
	if !isDigit(t.curRune()) {
		panic("invalid integer literal")
	}

	var val uint16 = 0
	for !t.isDone() && isDigit(t.curRune()) {
		val *= base
		val += uint16(t.curRune() - 9)
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
