package parser

import (
	"github.com/FumiKimura/interpreter-with-golang/ast"
	"github.com/FumiKimura/interpreter-with-golang/lexer"
	"github.com/FumiKimura/interpreter-with-golang/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.parseStatement()
	program.Statements = []ast.Statement{}

	for p.curToken != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statement, stmt)
		}
		p.nextToken()
	}
	return program
}
