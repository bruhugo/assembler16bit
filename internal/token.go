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

func (p ParameterType) IsLiteralType() bool {
	return p == ParameterTypeLabel ||
		p == ParameterTypeLiteral
}

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
	TokenTypeColon
	TokenTypeEOF
)

type OutputOffset uint16

const (
	OutputOffsetInst = 11
	OutputOffsetSrc0 = 8
	OutputOffsetSrc1 = 5
	OutputOffsetDest = 2
	OutputOffsetOps  = 0
)

type Token interface {
	Type() TokenType
}

type ParameterFormat struct {
	allowedTypes  ParameterSet
	outputOffset  OutputOffset
	outputOffsets []OutputOffset
}

func NewParameterFormat(outputOffset OutputOffset, allowedTypes ...ParameterType) *ParameterFormat {
	set := make(ParameterSet)
	for _, t := range allowedTypes {
		set[t] = struct{}{}
	}
	return &ParameterFormat{
		allowedTypes:  set,
		outputOffset:  outputOffset,
		outputOffsets: []OutputOffset{outputOffset},
	}
}

func (p *ParameterFormat) WithOffset(offset OutputOffset) *ParameterFormat {
	p.outputOffsets = append(p.outputOffsets, offset)
	return p
}

type TokenInstruction struct {
	// each instruction can have up to three params,
	// and depending on the quantity, the semantic might change
	Params [][]*ParameterFormat
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
