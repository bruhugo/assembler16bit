package assembler

type FragmentType int

const (
	FragmentTypeInstruction FragmentType = iota
	FragmentTypeLabel
)

// Fragment is an intermidiate data structure
// made for the compiler to consume and generate the
// appropriate binary
type Fragment struct {
	Next *Fragment
	Type FragmentType

	Label string

	Base   uint16
	Param1 ParamFragment
	Param2 ParamFragment
}

type ParamFragment struct {
	t     ParameterType
	val   uint16
	label string
}

func NewFragmentLabel(label string) *Fragment {
	return &Fragment{
		Label: label,
		Type:  FragmentTypeLabel,
	}
}

func NewFragmentInstruction(base uint16, param1, param2 ParamFragment) *Fragment {
	return &Fragment{
		Type:   FragmentTypeInstruction,
		Base:   base,
		Param1: param1,
		Param2: param2,
	}
}
