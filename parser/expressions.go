package parser

import (
	"fmt"
	"strconv"

	"github.com/mohamedirfanam/cynterpreter/lexer/token"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

const (
	LOWEST        = iota
	EQUALS        // ==
	LESSGREATER   // < > <= >=
	SUMSUB        // + -
	PRODUCTDEVIDE // * /
	PREFIX        // -x
	CALL          // function calls
)

var precedences = map[token.TokenType]int{
	token.PLUS:    SUMSUB,
	token.MINUS:   SUMSUB,
	token.ASTER:   PRODUCTDEVIDE,
	token.SLASH:   PRODUCTDEVIDE,
	token.PERCENT: PRODUCTDEVIDE,
	token.EQ:      EQUALS,
	token.NE:      EQUALS,
	token.LT:      LESSGREATER,
	token.LE:      LESSGREATER,
	token.GT:      LESSGREATER,
	token.GE:      LESSGREATER,
	token.LPAREN:  CALL,
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefixFunc, ok := p.prefixParseFuncs[p.curToken.TokenType]
	if !ok {
		p.errors = append(p.errors, fmt.Errorf("No valid prefix parsing function found for token %s", p.curToken.TokenType))
		return nil
	}
	leftExp := prefixFunc()

	for !p.peekTokenIs(token.SEMCOL) && precedence < p.peekPrecedence() {
		infixFunc, ok := p.infixParseFuncs[p.curToken.TokenType]
		if !ok {
			p.errors = append(p.errors, fmt.Errorf("No valid infix parsing function found for token %s", p.curToken.TokenType))
			return nil
		}

		p.nextToken()
		leftExp = infixFunc(leftExp)
	}
	return leftExp
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	val, err := strconv.ParseInt(p.curToken.Lexeme, 10, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("Parser Error: Error parsring integer literal %s", p.curToken.Lexeme))
		return nil
	}
	return &ast.IntegerLiteral{
		Token: p.curToken,
		Value: val,
	}
}
