package parser

import (
	"github.com/mohamedirfanam/cynterpreter/lexer"
	"github.com/mohamedirfanam/cynterpreter/lexer/token"
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

func (p *Parser) ParseProgram() {

}
