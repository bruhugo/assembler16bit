package assembler

type Parser struct {
	t *Tokenizer
}

func NewParser(text string) *Parser {
	t := NewTokenizer(text)
	t.NewToken()
	return &Parser{
		t: t,
	}
}

func (p *Parser) Program() *Fragment {
	f := &Fragment{}
	for p.t.CurType() != TokenTypeEOF {
		f.Next = p.Line()
		f = f.Next
	}

	return f.Next
}

func (p *Parser) Line() *Fragment {
	curType := p.t.CurType()
	var f *Fragment
	switch curType {
	case TokenTypeInstruction:
		f = p.Instruction()
	case TokenTypeDot:
		f = p.Label()
	default:
		panic(p.t.errMsg("expected either an instruction or a label"))
	}

	p.t.AssertAndNext(TokenTypeNewLine)

	return f
}

func (p *Parser) Label() *Fragment {
	p.t.AssertAndNext(TokenTypeDot)
	l := p.t.AssertAndNext(TokenTypeLabel)
	lVal := l.(*TokenLabel).label
	p.t.AssertAndNext(TokenTypeComa)

	return NewFragmentLabel(lVal)
}

func (p *Parser) Instruction() *Fragment {
	token := p.t.AssertAndNext(TokenTypeInstruction).(*TokenInstruction)

	param1 := p.Parameter(token.Param1)
	param2 := p.Parameter(token.Param2)

	return NewFragmentInstruction(token.Base, param1, param2)
}

func (p *Parser) Parameter(f ParameterFormat) ParamFragment {
	if !f.isRequired {
		return ParamFragment{t: ParameterTypeNone}
	}

	param := p.GetParameter()
	if _, ok := f.allowedTypes[param.t]; !ok {
		panic(p.t.errMsg("invalid parameter type"))
	}

	return param
}

func (p *Parser) GetParameter() ParamFragment {
	var paramFragment ParamFragment
	switch p.t.CurType() {
	case TokenTypeLabel:
		t := p.t.AssertAndNext(TokenTypeLabel)
		paramFragment.t = ParameterTypeLabel
		paramFragment.label = t.(*TokenLabel).label

	case TokenTypeLeftBrackets:
		p.t.AssertAndNext(TokenTypeLeftBrackets)
		t := p.t.AssertAndNext(TokenTypeParameter)
		tParam := t.(*TokenParameter)
		if tParam.paramType != ParameterTypeRegister {
			panic(p.t.errMsg("expected a register inside the address brackets"))
		}
		p.t.AssertAndNext(TokenTypeRightBrackets)

		paramFragment.t = ParameterTypeAddress
		paramFragment.val = tParam.val

	case TokenTypeParameter:
		t := p.t.AssertAndNext(TokenTypeParameter)
		tp := t.(*TokenParameter)

		paramFragment.t = tp.paramType
		paramFragment.val = tp.val

	default:
		panic(p.t.errMsg("invalid intruction parameter"))
	}

	return paramFragment
}
