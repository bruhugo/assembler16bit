package assembler

type FragmentType int

const (
	FragmentTypeInstruction FragmentType = iota
	FragmentTypeLabel
)

func (f FragmentType) String() string {
	switch f {
	case FragmentTypeInstruction:
		return "FragmentTypeInstruction"
	case FragmentTypeLabel:
		return "FragmentTypeLabel"
	default:
		return ""
	}
}

// Fragment is an intermidiate data structure
// made for the compiler to consume and generate the
// appropriate binary
type Fragment struct {
	Next *Fragment
	Type FragmentType

	Label string

	Base uint16
	Dest ParamFragment
	Src0 ParamFragment
	Src1 ParamFragment
}

type ParamFragment struct {
	t            ParameterType
	val          uint16
	label        string
	outputOffset OutputOffset
}

func NewFragmentLabel(label string) *Fragment {
	return &Fragment{
		Label: label,
		Type:  FragmentTypeLabel,
	}
}

func NewFragmentInstruction(base uint16, src0, src1, dest ParamFragment) *Fragment {
	return &Fragment{
		Type: FragmentTypeInstruction,
		Base: base,
		Dest: dest,
		Src0: src0,
		Src1: src1,
	}
}
