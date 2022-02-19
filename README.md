# interpreter-with-golang

## Chapter 1 - Lexing

Source Code -> Tokens -> Abstract Syntax Tree
Process of converting from Source Code to Tokens is called Lexing.

For example,

```
let x = 5 + 5;
```

will turn into

```
[
LET,
     IDENTIFIER("x"),
     EQUAL_SIGN,
     INTEGER(5),
     PLUS_SIGN,
     INTEGER(5),
     SEMICOLON
]
```

REPL - Read Eval Print Loop
-> I can think of this as what you see when you run "node" in terminal

## Chapter 2 - Parser

Parser converts token created by lexer into certain data structure.<br>
AST - abstract syntax tree -> data structure of token created by lexer<br>
Parse generator -> It is tools that, when fed with a formal description of a language, produce parsers as their output. This output is code that can then be compiled/interpreted and itself fed with source code as input to produce a syntax tree.<br>

### Difference between top-down and bottom-up parsers

top-down parser -> starts with constructing root node of the AST and then descends while the bottom-up parser does it other way around.<br>

Expression produces value, but statements do not.

### Some new terminology

Prefix operator - is an operator "in front of" its operand. <br>

Example: --5 <br>

Postfix operator - is an operator "after" its operand. <br>

Example: foobar++ <br>

Infix operators - is an operator "sits" between its operands <br>

5 \* 8 <br>


