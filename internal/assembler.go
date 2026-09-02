// Package assembler implements an assembler for the 16 bit machine
// im working on :)
package assembler

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type AssembleOption int

const (
	AssembleOptionLogisim AssembleOption = iota
	AssembleOptionRaw
)

const (
	OffsetInstructionID = 11
	OffsetSrc           = 8
	OffsetDest          = 5
	OffsetIO            = 3
	OffsetLiteral       = 2
)

type Assembler struct {
	parser *Parser
}

func NewAssembler(text string) *Assembler {
	return &Assembler{
		parser: NewParser(text),
	}
}

func (a *Assembler) Assemble(w io.Writer, option AssembleOption) {
	var count uint16 = 0
	symbolTable := make(map[string]uint16)

	head := a.parser.Program()
	cur := head
	for cur != nil {
		if cur.Type == FragmentTypeLabel {
			symbolTable[cur.Label] = count
		} else {
			count++
		}
		// if the instruction requires 16 more bits, increment one more
		if isLiteral(cur) {
			count++
		}
		cur = cur.Next
	}
	cur = head

	inst := make([]uint16, 0, count)
	for cur != nil {
		if cur.Type == FragmentTypeLabel {
			cur = cur.Next
			continue
		}
		inst = append(inst, getInstructionValue(cur))
		lit, ok := getLiteralValue(cur, symbolTable)
		if ok {
			inst = append(inst, lit)
		}
		cur = cur.Next
	}

	switch option {
	case AssembleOptionLogisim:
		err := writeSimulideOption(w, inst)
		if err != nil {
			panic(err)
		}
	}
}

func writeSimulideOption(writer io.Writer, inst []uint16) error {
	w := bufio.NewWriter(writer)
	_, err := w.Write([]byte("v2.0 raw\n"))
	if err != nil {
		return fmt.Errorf("error writing to output file: %w", err)
	}
	for _, instruction := range inst {
		_, err := w.WriteString(strconv.FormatInt(int64(instruction), 16))
		if err != nil {
			return fmt.Errorf("error writing to output file: %w", err)
		}

		_, err = w.WriteString(" ")
		if err != nil {
			return fmt.Errorf("error writing to output file: %w", err)
		}
	}
	err = w.Flush()
	if err != nil {
		return fmt.Errorf("error flushing buffer: %w", err)
	}
	return nil
}

func getLiteralValue(f *Fragment, symbols map[string]uint16) (uint16, bool) {
	switch {
	case f.Param1.t == ParameterTypeLabel:
		return getLiteralParam(f.Param1, symbols), true
	case f.Param2.t == ParameterTypeLabel:
		return getLiteralParam(f.Param2, symbols), true
	case f.Param1.t == ParameterTypeLiteral:
		return f.Param1.val, true
	case f.Param2.t == ParameterTypeLiteral:
		return f.Param2.val, true
	default:
		return 0, false
	}
}

func getLiteralParam(param ParamFragment, symbols map[string]uint16) uint16 {
	v, ok := symbols[param.label]
	if !ok {
		panic(fmt.Sprintf("symbol %s was not defined", param.label))
	}
	return v
}

func getInstructionValue(f *Fragment) uint16 {
	var inst uint16
	inst |= f.Base << OffsetInstructionID
	inst |= getSrcValue(f)
	inst |= getDestValue(f)
	inst |= getIOValue(f)
	if isLiteral(f) {
		inst |= 1 << OffsetLiteral
	}

	return inst
}

func getSrcValue(f *Fragment) uint16 {
	switch f.Param2.t {
	case ParameterTypeAddress, ParameterTypeRegister, ParameterTypeInput:
		return f.Param2.val << OffsetSrc
	default:
		return 0
	}
}

func getDestValue(f *Fragment) uint16 {
	switch f.Param2.t {
	case ParameterTypeAddress, ParameterTypeOutput, ParameterTypeRegister:
		return f.Param1.val << OffsetDest
	default:
		return 0
	}
}

func getIOValue(f *Fragment) uint16 {
	var io uint16
	if f.Param1.t == ParameterTypeOutput {
		io |= 0b10
	}
	if f.Param2.t == ParameterTypeInput {
		io |= 0b01
	}
	return io << OffsetIO
}

func isLiteral(f *Fragment) bool {
	return f.Param1.t.IsLiteralType() || f.Param2.t.IsLiteralType()
}
