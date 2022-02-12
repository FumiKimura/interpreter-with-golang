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
