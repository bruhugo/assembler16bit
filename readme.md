# 16 bit assembler written in Golang

I'm building a 16-bit computer in a simulator called Logisim, and in order to write programs faster, I made an assembler written in Go that translates common assembly instructions into machine code instructions.

## Assembler structure

The assembler is divided in three main parts, each performing one specific task. I'm no expert on compilers and interpreters, so it might not follow the best structure for an assembler, but I works, and that's what matters :)

### Tokenizer

This component is responsible for transforming the plain text program file into well defined units that the program can later consume. Those units are usually called tokens.

### Parser

The parser is responsible for consuming the tokens and making sure that it follows the appropriate structure defined by my assembly language grammar. It's basically a simple state machine. After ensuring its structure, the parser transforms the stream of tokens into an intermediate data structure designed to be easily consumed by the assembler component. This data structure is called a fragment and is similar to what the GCC assembler does.

### Assembler

This final component consumes the fragments and transforms them into binary following the instruction structure. It also has a very simple symbol table (hash map) of labels. Its a two-pass assembler, so it first scans the fragments for label symbols, and then iterates them again to populate it when it's used.

<img width="1549" height="687" alt="image" src="https://github.com/user-attachments/assets/e8c31554-0363-4e95-a9aa-bcfd614be2ab" />

## Why Go?

I'm too lazy to write it in C

## Instruction format

Each instruction is composed of 16 bits, but it can be extended to 32 bits in order to allow addresses and numbers to be passed to the instruction. The structure goes as follows:

- 0-1:      operations (IO and extend)
- 2-4:      destination register
- 5-7:      source 1 register
- 8-10:     source 0 register
- 11-15:    the instruction number

## Instruction set

- 0: MOV
- 1: LOAD
- 2: STR
- 3: JMP
- 4: JMPZ
- 5: JMPN
- 6: ADD
- 7: SUB
- 8: INC
- 9: DEC
- 10: AND
- 11: OR
- 12: NOT
- 13: PUSH
- 14: POP
- 15: SHL
- 16: SHR
- 17: ROT
- 18: MUL
- 19: DIV
- 20: CALL
- 21: RET
- 22: HLT
- 23: CMP

## Parameters

The instructions might take 0, 1, 2 or even three parameters. For example:

- RET
- PUSH r0
- MOV r0, r1
- ADD r0, r1, r2

Notice that some instructions that take three parameters can also take only two. In those cases, the source 0 offset will be the same as destination. Example:

ADD r0, r0, r1  is the same as  ADD r0, r1

## How to use the assembler

Download the binary located in the release page of this repository and download the one for your OS and architecture. After that, you can run the executable as following (change "assembler" for whatever name your executable has):

./assembler -i inputFile -o outputFile

and it should output the binary file. (use -v for dumping the instructions in stdout, usefull for debugging). After that, load the output file into the instruction memory box in Logisim and start your program!
