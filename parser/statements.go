package parser

import (
	"github.com/mohamedirfanam/cynterpreter/lexer/token"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

func (p *Parser) ParseStatement() ast.Statement {
	switch p.curToken {
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	tkn := p.curToken
	exp := p.parseExpression(LOWEST)
	p.expectPeekToken(token.SEMCOL)
	return &ast.ExpressionStatement{
		Token:      tkn,
		Expression: exp,
	}
}
