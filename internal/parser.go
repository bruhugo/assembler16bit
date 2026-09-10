package assembler

type Parser struct {
	t *Tokenizer
}

func NewParser(text string) *Parser {
	t := NewTokenizer(text)
	return &Parser{
		t: t,
	}
}

func (p *Parser) Program() *Fragment {
	dummy := &Fragment{}
	f := dummy
	for p.t.CurType() != TokenTypeEOF {
		f.Next = p.Line()
		f = f.Next
	}

	return dummy.Next
}

func (p *Parser) Line() *Fragment {
	for p.t.CurType() == TokenTypeNewLine {
		p.t.AssertAndNext(TokenTypeNewLine)
	}

	curType := p.t.CurType()
	var f *Fragment
	switch curType {
	case TokenTypeInstruction:
		f = p.Instruction()
	case TokenTypeLabel:
		f = p.Label()
	case TokenTypeEOF:
		return nil
	default:
		panic(p.t.errMsg("expected either an instruction or a label"))
	}

	if p.t.CurType() == TokenTypeEOF {
		return f
	}

	p.t.AssertAndNext(TokenTypeNewLine)
	return f
}

func (p *Parser) Label() *Fragment {
	l := p.t.AssertAndNext(TokenTypeLabel)
	lVal := l.(*TokenLabel).label
	p.t.AssertAndNext(TokenTypeColon)

	return NewFragmentLabel(lVal)
}

func (p *Parser) Instruction() *Fragment {
	var dest, src0, src1 ParamFragment
	token := p.t.AssertAndNext(TokenTypeInstruction).(*TokenInstruction)

	if p.t.CurType() != TokenTypeParameter && p.t.CurType() != TokenTypeLabel {
		return NewFragmentInstruction(token.Base, src0, src1, dest)
	}

	params := []ParamFragment{
		p.Parameter(),
	}

	for p.t.CurType() == TokenTypeComa {
		p.t.AssertAndNext(TokenTypeComa)
		params = append(params, p.Parameter())
	}

	paramsIndex := len(params) - 1

	// the number of parameters is greater or less than the required number
	if paramsIndex >= len(token.Params) || token.Params[paramsIndex] == nil {
		panic(p.t.errMsg("unexpected number of parameters"))
	}

	paramTokens := token.Params[paramsIndex]
	for i, t := range paramTokens {
		curParam := params[i]
		_, ok := t.allowedTypes[curParam.t]
		if !ok {
			panic(p.t.errMsg("parameter of invalid type"))
		}

		offsets := t.outputOffsets
		if len(offsets) == 0 {
			offsets = []OutputOffset{t.outputOffset}
		}
		for _, offset := range offsets {
			switch offset {
			case OutputOffsetDest:
				dest = curParam
			case OutputOffsetSrc0:
				src0 = curParam
			case OutputOffsetSrc1:
				src1 = curParam
			}
		}
	}

	return NewFragmentInstruction(token.Base, src0, src1, dest)
}

func (p *Parser) Parameter() ParamFragment {
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
