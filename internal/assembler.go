// Package assembler implements an assembler for the 16 bit machine
// im working on :)
package assembler

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const LogisimHeader = "v2.0 raw\n"

type AssembleOption int

const (
	AssembleOptionLogisim AssembleOption = iota
	AssembleOptionRaw
)

type Assembler struct {
	parser *Parser
}

func NewAssembler(text string) *Assembler {
	return &Assembler{
		parser: NewParser(text),
	}
}

func (a *Assembler) Assemble(w io.Writer, option AssembleOption, verbose bool) {
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
		err := writeSimulideOption(w, inst, verbose)
		if err != nil {
			panic(err)
		}
	}
}

func writeSimulideOption(writer io.Writer, inst []uint16, verbose bool) error {
	w := bufio.NewWriter(writer)
	_, err := w.Write([]byte(LogisimHeader))
	if err != nil {
		return fmt.Errorf("error writing to output file: %w", err)
	}
	if verbose {
		printInstructionHeader()
	}
	for i, instruction := range inst {
		_, err := w.WriteString(strconv.FormatInt(int64(instruction), 16))
		if err != nil {
			return fmt.Errorf("error writing to output file: %w", err)
		}

		if verbose {
			printInstruction(instruction)
		}

		_, err = w.WriteString(" ")
		if err != nil {
			return fmt.Errorf("error writing to output file: %w", err)
		}
		if i%10 == 9 {
			_, err = w.WriteString("\n")
			if err != nil {
				return fmt.Errorf("error writing to output file: %w", err)
			}

		}
	}
	err = w.Flush()
	if err != nil {
		return fmt.Errorf("error flushing buffer: %w", err)
	}
	return nil
}

func getLiteralValue(f *Fragment, symbols map[string]uint16) (uint16, bool) {
	switch f.Src1.t {
	case ParameterTypeLabel:
		val, ok := symbols[f.Src1.label]
		if !ok {
			return 0, false
		}
		return val, true

	case ParameterTypeLiteral:
		return f.Src1.val, true
	}
	return 0, false
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
	inst |= f.Base << OutputOffsetInst
	inst |= f.Src0.val << OutputOffsetSrc0
	if f.Src1.t != ParameterTypeLiteral {
		inst |= f.Src1.val << OutputOffsetSrc1
	}
	inst |= f.Dest.val << OutputOffsetDest
	inst |= getOpsValue(f)

	return inst
}

func getOpsValue(f *Fragment) uint16 {
	if f.Src1.t.IsLiteralType() {
		return 0b11
	}

	if f.Src1.t == ParameterTypeInput {
		return 0b01
	}

	if f.Dest.t == ParameterTypeOutput {
		return 0b10
	}

	return 0
}

func isLiteral(f *Fragment) bool {
	return f.Src1.t.IsLiteralType()
}

func printInstructionHeader() {
	fmt.Printf("%5s %3s %3s %3s %2s\n\n", "BASE", "SR0", "SR1", "DST", "OP")
}

func printInstruction(inst uint16) {
	base := inst >> OutputOffsetInst & 0b11111
	src0 := inst >> OutputOffsetSrc0 & 0b111
	src1 := inst >> OutputOffsetSrc1 & 0b111
	dest := inst >> OutputOffsetDest & 0b111
	op := inst >> OutputOffsetOps & 0b11

	fmt.Printf("%05s %03s %03s %03s %02s\n",
		toB(base),
		toB(src0),
		toB(src1),
		toB(dest),
		toB(op),
	)
}

func toB(i uint16) string {
	return strconv.FormatUint(uint64(i), 2)
}
