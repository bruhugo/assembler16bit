package assembler

var reservedKeywords = map[string]Token{
	"MOV": &TokenInstruction{
		Base:   0,
		Param1: NewParameterFormat(true, ParameterTypeOutput, ParameterTypeRegister),
		Param2: NewParameterFormat(true, ParameterTypeInput, ParameterTypeRegister, ParameterTypeLiteral),
	},
	"LOAD": &TokenInstruction{
		Base:   1,
		Param1: NewParameterFormat(true, ParameterTypeRegister),
		Param2: NewParameterFormat(true, ParameterTypeAddress, ParameterTypeLiteral),
	},
	"STR": &TokenInstruction{
		Base:   2,
		Param1: NewParameterFormat(true, ParameterTypeAddress, ParameterTypeLiteral),
		Param2: NewParameterFormat(true, ParameterTypeLiteral, ParameterTypeRegister),
	},
	"JMP": &TokenInstruction{
		Base:   3,
		Param1: NewParameterFormat(true, ParameterTypeLiteral),
		Param2: NewParameterFormat(false),
	},
	"JMPZ": &TokenInstruction{
		Base:   4,
		Param1: NewParameterFormat(true, ParameterTypeLiteral),
		Param2: NewParameterFormat(false),
	},
	"JMPN": &TokenInstruction{
		Base:   5,
		Param1: NewParameterFormat(true, ParameterTypeLiteral),
		Param2: NewParameterFormat(false),
	},
	"ADD": &TokenInstruction{
		Base:   6,
		Param1: NewParameterFormat(true, ParameterTypeRegister),
		Param2: NewParameterFormat(true, ParameterTypeLiteral, ParameterTypeRegister),
	},
	"SUB": &TokenInstruction{
		Base:   7,
		Param1: NewParameterFormat(true, ParameterTypeRegister),
		Param2: NewParameterFormat(true, ParameterTypeLiteral, ParameterTypeRegister),
	},
	"INC": &TokenInstruction{
		Base:   8,
		Param1: NewParameterFormat(true, ParameterTypeRegister),
		Param2: NewParameterFormat(false),
	},
	"DEC": &TokenInstruction{
		Base:   9,
		Param1: NewParameterFormat(true, ParameterTypeRegister),
		Param2: NewParameterFormat(false),
	},
	"AND": &TokenInstruction{
		Base:   10,
		Param1: NewParameterFormat(true, ParameterTypeRegister),
		Param2: NewParameterFormat(true, ParameterTypeLiteral, ParameterTypeRegister),
	},
	"OR": &TokenInstruction{
		Base:   11,
		Param1: NewParameterFormat(true, ParameterTypeRegister),
		Param2: NewParameterFormat(true, ParameterTypeLiteral, ParameterTypeRegister),
	},
	"NOT": &TokenInstruction{
		Base:   12,
		Param1: NewParameterFormat(true, ParameterTypeRegister),
		Param2: NewParameterFormat(true, ParameterTypeLiteral, ParameterTypeRegister),
	},
	"PUSH": &TokenInstruction{
		Base:   13,
		Param1: NewParameterFormat(true, ParameterTypeRegister, ParameterTypeLiteral),
		Param2: NewParameterFormat(false),
	},
	"POP": &TokenInstruction{
		Base:   14,
		Param1: NewParameterFormat(true, ParameterTypeRegister),
		Param2: NewParameterFormat(false),
	},
	"SHL": &TokenInstruction{
		Base:   15,
		Param1: NewParameterFormat(true, ParameterTypeRegister),
		Param2: NewParameterFormat(true, ParameterTypeLiteral, ParameterTypeRegister),
	},
	"SHR": &TokenInstruction{
		Base:   16,
		Param1: NewParameterFormat(true, ParameterTypeRegister),
		Param2: NewParameterFormat(true, ParameterTypeLiteral, ParameterTypeRegister),
	},
	"ROT": &TokenInstruction{
		Base:   17,
		Param1: NewParameterFormat(true, ParameterTypeRegister),
		Param2: NewParameterFormat(true, ParameterTypeLiteral, ParameterTypeRegister),
	},
	"MUL": &TokenInstruction{
		Base:   18,
		Param1: NewParameterFormat(true, ParameterTypeRegister),
		Param2: NewParameterFormat(true, ParameterTypeLiteral, ParameterTypeRegister),
	},
	"DIV": &TokenInstruction{
		Base:   19,
		Param1: NewParameterFormat(true, ParameterTypeRegister),
		Param2: NewParameterFormat(true, ParameterTypeLiteral, ParameterTypeRegister),
	},
	"CALL": &TokenInstruction{
		Base:   20,
		Param1: NewParameterFormat(true, ParameterTypeLiteral, ParameterTypeLabel),
		Param2: NewParameterFormat(false),
	},
	"RET": &TokenInstruction{
		Base:   21,
		Param1: NewParameterFormat(false),
		Param2: NewParameterFormat(false),
	},
	"HLT": &TokenInstruction{
		Base:   22,
		Param1: NewParameterFormat(false),
		Param2: NewParameterFormat(false),
	},
	"CMP": &TokenInstruction{
		Base:   22,
		Param1: NewParameterFormat(true, ParameterTypeRegister),
		Param2: NewParameterFormat(true, ParameterTypeRegister, ParameterTypeLiteral),
	},
	"r0": &TokenParameter{
		paramType: ParameterTypeRegister,
		val:       0,
	},
	"r1": &TokenParameter{
		paramType: ParameterTypeRegister,
		val:       1,
	},
	"r2": &TokenParameter{
		paramType: ParameterTypeRegister,
		val:       2,
	},
	"r3": &TokenParameter{
		paramType: ParameterTypeRegister,
		val:       3,
	},
	"r4": &TokenParameter{
		paramType: ParameterTypeRegister,
		val:       4,
	},
	"r5": &TokenParameter{
		paramType: ParameterTypeRegister,
		val:       5,
	},
	"r6": &TokenParameter{
		paramType: ParameterTypeRegister,
		val:       6,
	},
	"r7": &TokenParameter{
		paramType: ParameterTypeRegister,
		val:       7,
	},
	"r8": &TokenParameter{
		paramType: ParameterTypeRegister,
		val:       8,
	},
	"out0": &TokenParameter{
		paramType: ParameterTypeOutput,
		val:       0,
	},
	"out1": &TokenParameter{
		paramType: ParameterTypeOutput,
		val:       1,
	},
	"out2": &TokenParameter{
		paramType: ParameterTypeOutput,
		val:       2,
	},
	"out3": &TokenParameter{
		paramType: ParameterTypeOutput,
		val:       3,
	},
	"out4": &TokenParameter{
		paramType: ParameterTypeOutput,
		val:       4,
	},
	"out5": &TokenParameter{
		paramType: ParameterTypeOutput,
		val:       5,
	},
	"out6": &TokenParameter{
		paramType: ParameterTypeOutput,
		val:       6,
	},
	"out7": &TokenParameter{
		paramType: ParameterTypeOutput,
		val:       7,
	},
	"in0": &TokenParameter{
		paramType: ParameterTypeInput,
		val:       0,
	},
	"in1": &TokenParameter{
		paramType: ParameterTypeInput,
		val:       1,
	},
	"in2": &TokenParameter{
		paramType: ParameterTypeInput,
		val:       2,
	},
	"in3": &TokenParameter{
		paramType: ParameterTypeInput,
		val:       3,
	},
	"in4": &TokenParameter{
		paramType: ParameterTypeInput,
		val:       4,
	},
	"in5": &TokenParameter{
		paramType: ParameterTypeInput,
		val:       5,
	},
	"in6": &TokenParameter{
		paramType: ParameterTypeInput,
		val:       6,
	},
	"in7": &TokenParameter{
		paramType: ParameterTypeInput,
		val:       7,
	},
}
