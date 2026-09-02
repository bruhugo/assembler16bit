package assembler

type ParameterType int

const (
	ParameterTypeNone ParameterType = iota
	ParameterTypeInput
	ParameterTypeOutput
	ParameterTypeRegister
	ParameterTypeAddress
	ParameterTypeLiteral
	ParameterTypeLabel
)

type ParameterSet map[ParameterType]struct{}

type TokenType int

const (
	TokenTypeInstruction TokenType = iota
	TokenTypeLabel
	TokenTypeLiteral
	TokenTypeParameter
	TokenTypeNewLine
	TokenTypeComa
	TokenTypeDot
	TokenTypeLeftBrackets
	TokenTypeRightBrackets
	TokenTypeEOF
)

type Token interface {
	Type() TokenType
}

type ParameterFormat struct {
	isRequired   bool
	allowedTypes ParameterSet
}

func NewParameterFormat(required bool, allowedTypes ...ParameterType) ParameterFormat {
	set := make(ParameterSet)
	for _, t := range allowedTypes {
		set[t] = struct{}{}
	}
	return ParameterFormat{
		isRequired:   required,
		allowedTypes: set,
	}
}

type TokenInstruction struct {
	Param1 ParameterFormat
	Param2 ParameterFormat
	// the instruction is distinguished from the others
	// by a 5 bit field
	Base uint16
}

func (t *TokenInstruction) Type() TokenType {
	return TokenTypeInstruction
}

type TokenParameter struct {
	paramType ParameterType
	val       uint16
}

func (t *TokenParameter) Type() TokenType {
	return TokenTypeParameter
}

type TokenLabel struct {
	label string
}

func (t *TokenLabel) Type() TokenType {
	return TokenTypeLabel
}

type SimpleToken struct {
	t TokenType
}

func (s *SimpleToken) Type() TokenType {
	return s.t
}
