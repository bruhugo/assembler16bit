package assembler

var reservedKeywords = map[string]Token{
	"MOV": &TokenInstruction{
		Base: 0,
		Params: [][]*ParameterFormat{
			nil,
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister, ParameterTypeOutput),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeInput, ParameterTypeLiteral),
			},
		},
	},
	"LOAD": &TokenInstruction{
		Base: 1,
		Params: [][]*ParameterFormat{
			nil,
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeAddress, ParameterTypeLiteral, ParameterTypeLabel),
			},
		},
	},
	"STR": &TokenInstruction{
		Base: 2,
		Params: [][]*ParameterFormat{
			nil,
			{
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeAddress, ParameterTypeLiteral, ParameterTypeLabel),
				NewParameterFormat(OutputOffsetSrc0, ParameterTypeRegister),
			},
		},
	},
	"JMP": &TokenInstruction{
		Base: 3,
		Params: [][]*ParameterFormat{
			{
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeLiteral, ParameterTypeLabel),
			},
		},
	},
	"JMPZ": &TokenInstruction{
		Base: 4,
		Params: [][]*ParameterFormat{
			{
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeLiteral, ParameterTypeLabel),
			},
		},
	},
	"JMPN": &TokenInstruction{
		Base: 5,
		Params: [][]*ParameterFormat{
			{
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeLiteral, ParameterTypeLabel),
			},
		},
	},
	"ADD": &TokenInstruction{
		Base: 6,
		Params: [][]*ParameterFormat{
			nil,
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister).WithOffset(OutputOffsetSrc0),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc0, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
		},
	},
	"SUB": &TokenInstruction{
		Base: 7,
		Params: [][]*ParameterFormat{
			nil,
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister).WithOffset(OutputOffsetSrc0),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc0, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
		},
	},
	"INC": &TokenInstruction{
		Base: 8,
		Params: [][]*ParameterFormat{
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister),
			},
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister).WithOffset(OutputOffsetSrc0),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
		},
	},
	"DEC": &TokenInstruction{
		Base: 9,
		Params: [][]*ParameterFormat{
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister),
			},
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister).WithOffset(OutputOffsetSrc0),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
		},
	},
	"AND": &TokenInstruction{
		Base: 10,
		Params: [][]*ParameterFormat{
			nil,
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister).WithOffset(OutputOffsetSrc0),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc0, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
		},
	},
	"OR": &TokenInstruction{
		Base: 11,
		Params: [][]*ParameterFormat{
			nil,
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister).WithOffset(OutputOffsetSrc0),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc0, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
		},
	},
	"NOT": &TokenInstruction{
		Base: 12,
		Params: [][]*ParameterFormat{
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister),
			},
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
		},
	},
	"PUSH": &TokenInstruction{
		Base: 13,
		Params: [][]*ParameterFormat{
			{
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
		},
	},
	"POP": &TokenInstruction{
		Base: 14,
		Params: [][]*ParameterFormat{
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister),
			},
		},
	},
	"SHL": &TokenInstruction{
		Base: 15,
		Params: [][]*ParameterFormat{
			nil,
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister).WithOffset(OutputOffsetSrc0),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc0, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
		},
	},
	"SHR": &TokenInstruction{
		Base: 16,
		Params: [][]*ParameterFormat{
			nil,
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister).WithOffset(OutputOffsetSrc0),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc0, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
		},
	},
	"ROT": &TokenInstruction{
		Base: 17,
		Params: [][]*ParameterFormat{
			nil,
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister).WithOffset(OutputOffsetSrc0),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc0, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
		},
	},
	"MUL": &TokenInstruction{
		Base: 18,
		Params: [][]*ParameterFormat{
			nil,
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister).WithOffset(OutputOffsetSrc0),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc0, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
		},
	},
	"DIV": &TokenInstruction{
		Base: 19,
		Params: [][]*ParameterFormat{
			nil,
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister).WithOffset(OutputOffsetSrc0),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
			{
				NewParameterFormat(OutputOffsetDest, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc0, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
		},
	},
	"CALL": &TokenInstruction{
		Base: 20,
		Params: [][]*ParameterFormat{
			{
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeLiteral, ParameterTypeLabel),
			},
		},
	},
	"RET": &TokenInstruction{
		Base:   21,
		Params: nil,
	},
	"HLT": &TokenInstruction{
		Base:   22,
		Params: nil,
	},
	"CMP": &TokenInstruction{
		Base: 23,
		Params: [][]*ParameterFormat{
			nil,
			{
				NewParameterFormat(OutputOffsetSrc0, ParameterTypeRegister),
				NewParameterFormat(OutputOffsetSrc1, ParameterTypeRegister, ParameterTypeLiteral),
			},
		},
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
