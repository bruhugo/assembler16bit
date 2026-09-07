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

## Why Go?

I'm too lazy to write it in C
