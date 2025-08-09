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
		p.errors = append(p.errors, fmt.Errorf("no valid prefix parsing function found for token %s", p.curToken.TokenType))
		return nil
	}
	leftExp := prefixFunc()

	for !p.peekTokenIs(token.SEMCOL) && precedence < p.peekPrecedence() {
		p.nextToken()
		infixFunc, ok := p.infixParseFuncs[p.curToken.TokenType]
		if !ok {
			p.errors = append(p.errors, fmt.Errorf("no valid infix parsing function found for token %s", p.curToken.TokenType))
			return nil
		}

		leftExp = infixFunc(leftExp)
	}
	return leftExp
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	val, err := strconv.ParseInt(p.curToken.Lexeme, 10, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("parser Error: Error parsring integer literal %s", p.curToken.Lexeme))
		return nil
	}
	return &ast.IntegerLiteral{
		Token: p.curToken,
		Value: val,
	}
}

func (p *Parser) parseCharLiteral() ast.Expression {
	val, err := strconv.Unquote(p.curToken.Lexeme)
	if err != nil || len(val) > 1 {
		p.errors = append(p.errors, fmt.Errorf("parser Error: Error parsring char literal %s", p.curToken.Lexeme))
		return nil
	}
	return &ast.CharLiteral{
		Token: p.curToken,
		Value: val[0],
	}
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	val, err := strconv.ParseFloat(p.curToken.Lexeme, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("parser Error: Error parsring float  literal %s", p.curToken.Lexeme))
		return nil
	}
	return &ast.FloatLiteral{
		Token: p.curToken,
		Value: val,
	}
}

func (p *Parser) parseStringLiteral() ast.Expression {
	val, err := strconv.Unquote(p.curToken.Lexeme)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("parser Error: Error parsing String literal %s", p.curToken.Lexeme))
		return nil
	}
	return &ast.StringLiteral{
		Token: p.curToken,
		Value: val,
	}
}

func (p *Parser) parseIdentifierExpression() ast.Expression {
	return &ast.IdentifierExpression{
		Token: p.curToken,
		Value: p.curToken.Lexeme,
	}
}

func (p *Parser) parseCallExpression(funcIdentifier ast.Expression) ast.Expression {
	expr := &ast.CallExpression{
		Token:    p.curToken,
		Function: funcIdentifier,
	}
	expr.Args = p.parseCallArgs()
	return expr
}

func (p *Parser) parseCallArgs() []ast.Expression {
	p.nextToken()
	var args []ast.Expression
	if p.curTokenIs(token.RPAREN) {
		return args
	}
	args = append(args, p.parseExpression(LOWEST))
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}
	if !p.expectPeekToken(token.RPAREN) {
		return nil
	}
	return args
}

func (p *Parser) parseInfixExpression(leftExp ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:   p.curToken,
		LeftExp: leftExp,
		Op:      p.curToken.Lexeme,
	}
	precedence := p.curPrecedence()
	p.nextToken()
	rightExp := p.parseExpression(precedence)
	exp.RightExp = rightExp

	return exp
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()
	exp := p.parseExpression(LOWEST)
	p.expectPeekToken(token.RPAREN)
	return exp
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	exp := &ast.PrefixExpression{
		Token: p.curToken,
		Op:    p.curToken.Lexeme,
	}
	p.nextToken()
	rexp := p.parseExpression(PREFIX)
	exp.Exp = rexp
	return exp
}
