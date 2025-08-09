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

	prefixParseFuncs map[token.TokenType]prefixParseFunc
	infixParseFuncs  map[token.TokenType]infixParseFunc

	errors []error
}

type (
	prefixParseFunc func() ast.Expression
	infixParseFunc  func(ast.Expression) ast.Expression
)

func New(input string) *Parser {
	p := Parser{
		l: lexer.New(input),
	}

	p.nextToken()
	p.nextToken()

	p.prefixParseFuncs = make(map[token.TokenType]prefixParseFunc)
	p.infixParseFuncs = make(map[token.TokenType]infixParseFunc)

	p.registerPrefixFunc(token.INT_LITERAL, p.parseIntegerLiteral)
	p.registerPrefixFunc(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefixFunc(token.IDENTIFIER, p.parseIdentifierExpression)
	p.registerPrefixFunc(token.CHAR_LITERAL, p.parseCharLiteral)
	p.registerPrefixFunc(token.FLOAT_LITERAL, p.parseFloatLiteral)
	p.registerPrefixFunc(token.STRING_LITERAL, p.parseStringLiteral)
	p.registerPrefixFunc(token.PLUS, p.parsePrefixExpression)
	p.registerPrefixFunc(token.MINUS, p.parsePrefixExpression)
	p.registerPrefixFunc(token.NOT, p.parsePrefixExpression)

	p.registerInfixFunc(token.PLUS, p.parseInfixExpression)
	p.registerInfixFunc(token.MINUS, p.parseInfixExpression)
	p.registerInfixFunc(token.ASTER, p.parseInfixExpression)
	p.registerInfixFunc(token.SLASH, p.parseInfixExpression)
	p.registerInfixFunc(token.PERCENT, p.parseInfixExpression)
	p.registerInfixFunc(token.EQ, p.parseInfixExpression)
	p.registerInfixFunc(token.NE, p.parseInfixExpression)
	p.registerInfixFunc(token.LT, p.parseInfixExpression)
	p.registerInfixFunc(token.LE, p.parseInfixExpression)
	p.registerInfixFunc(token.GT, p.parseInfixExpression)
	p.registerInfixFunc(token.GE, p.parseInfixExpression)
	p.registerInfixFunc(token.LPAREN, p.parseInfixExpression)

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

func (p *Parser) expectPeekToken(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.errors = append(p.errors, fmt.Errorf("Parser Error, Exptected Token - %s, Got - %s", t, p.peekToken.TokenType))
	return false
}

func (p *Parser) peekPrecedence() int {
	precedence, ok := precedences[p.peekToken.TokenType]
	if ok {
		return precedence
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	precedence, ok := precedences[p.curToken.TokenType]
	if ok {
		return precedence
	}
	return LOWEST
}

func (p *Parser) registerPrefixFunc(t token.TokenType, fn prefixParseFunc) {
	p.prefixParseFuncs[t] = fn
}

func (p *Parser) registerInfixFunc(t token.TokenType, fn infixParseFunc) {
	p.infixParseFuncs[t] = fn
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	for !p.curTokenIs(token.EOF) {
		statement := p.ParseStatement()
		program.Statements = append(program.Statements, statement)
		p.nextToken()
	}
	return program
}
