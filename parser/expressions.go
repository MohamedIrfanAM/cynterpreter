package parser

import (
	"fmt"
	"strconv"

	"github.com/mohamedirfanam/cynterpreter/lexer/token"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

const (
	LOWEST        = iota
	OR            // ||
	AND           // &&
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
	token.AND:     AND,
	token.OR:      OR,
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

func (p *Parser) parseBoolLiteral() ast.Expression {
	var val bool
	if p.curToken.Lexeme == "true" {
		val = true
	} else {
		val = false
	}
	return &ast.BoolLiteral{
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

func getOpInfixExpression(ident ast.Expression, expr ast.Expression, opType token.TokenType) ast.Expression {
	tkn := token.AssignmentOpMap[opType]
	exp := &ast.InfixExpression{
		Token:    tkn,
		LeftExp:  ident,
		Op:       tkn.Lexeme,
		RightExp: expr,
	}
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

func (p *Parser) parseFunctionLiteral(funcIdentifier *ast.IdentifierExpression) *ast.FunctionLiteral {
	expr := &ast.FunctionLiteral{
		Token:    p.curToken,
		Function: funcIdentifier,
	}
	p.nextToken()
	expr.Params = p.parseFunctionParams()
	p.expectPeekToken(token.LBRACE)
	expr.Block = p.parseBlockStatement()
	return expr
}

func (p *Parser) parseFunctionParams() []*ast.Parameter {
	var params []*ast.Parameter
	if p.curTokenIs(token.RPAREN) {
		return params
	}
	params = append(params, p.parseFunctionParam())
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		params = append(params, p.parseFunctionParam())
	}
	if !p.expectPeekToken(token.RPAREN) {
		return nil
	}
	return params
}

func (p *Parser) parseFunctionParam() *ast.Parameter {
	if !token.IsDatatype(p.curToken.TokenType) {
		p.errors = append(p.errors, fmt.Errorf("not valid datatype token for function parameter,got %s", p.curToken.TokenType))
		return nil
	}
	param := &ast.Parameter{
		Token: p.curToken,
		Type:  p.curToken.TokenType,
	}
	p.nextToken()
	ident := p.parseIdentifierExpression()
	param.Identifier = ident.(*ast.IdentifierExpression)
	return param
}

func (p *Parser) parseArrayDeclaration(tkn token.Token, arrIdentifier *ast.IdentifierExpression) *ast.ArrayDeclaration {
	expr := &ast.ArrayDeclaration{
		Token:     tkn,
		Type:      tkn.TokenType,
		Identifer: *arrIdentifier,
		Length:    -1,
	}

	p.nextToken()
	if p.curTokenIs(token.INT_LITERAL) {
		len, _ := strconv.ParseInt(p.curToken.Lexeme, 10, 32)
		expr.Length = int(len)
		p.nextToken()
	} else if !p.curTokenIs(token.RBRACK) {
		p.errors = append(p.errors, fmt.Errorf("invalid token as array length found"))
		return nil
	}

	if p.peekTokenIs(token.SEMCOL) {
		if expr.Length == -1 {
			p.errors = append(p.errors, fmt.Errorf("array declaration without specifying length or assigning array literal"))
			return nil
		}
		p.nextToken()
		return expr
	}
	p.expectPeekToken(token.ASSIGN)
	p.expectPeekToken(token.LBRACE)
	p.nextToken()

	vals := p.parseArrayLiteral()
	if expr.Length != -1 && expr.Length != len(vals) {
		p.errors = append(p.errors, fmt.Errorf("length of the array declared, mismatch. Declared %d, assigned %d", expr.Length, len(vals)))
		return nil
	}
	expr.Length = len(vals)
	expr.Literal = vals
	p.expectPeekToken(token.SEMCOL)
	return expr
}

func (p *Parser) parseArrayLiteral() []ast.Expression {
	var vals []ast.Expression
	if p.curTokenIs(token.RBRACK) {
		return vals
	}
	vals = append(vals, p.parseExpression(LOWEST))
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		vals = append(vals, p.parseExpression(LOWEST))
	}
	if !p.expectPeekToken(token.RBRACE) {
		p.errors = append(p.errors, fmt.Errorf("missing closing bracket ] for array declaration"))
		return nil
	}
	return vals
}
