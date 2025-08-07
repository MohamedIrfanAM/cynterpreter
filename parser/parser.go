package parser

import (
	"fmt"

	"github.com/mohamedirfanam/cynterpreter/lexer"
	"github.com/mohamedirfanam/cynterpreter/lexer/token"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []error
}

func New(input string) *Parser {
	p := Parser{
		l: lexer.New(input),
	}

	p.nextToken()
	p.nextToken()

	return &p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.TokenType == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.TokenType == t
}

func (p *Parser) expectPeekToken(t token.TokenType) {
	if p.peekTokenIs(t) {
		p.nextToken()
	}
	p.errors = append(p.errors, fmt.Errorf("Parser Error, Exptected Token - %s, Got - %s", t, p.peekToken.TokenType))
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	for !p.peekTokenIs(token.EOF) {
		statement := p.ParseStatement()
		program.Statements = append(program.Statements, statement)
		p.nextToken()
	}
	return program
}
