package parser

import (
	"fmt"

	"github.com/mohamedirfanam/cynterpreter/lexer/token"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

func (p *Parser) ParseStatement() ast.Statement {
	switch p.curToken.TokenType {
	case token.INT, token.CHAR, token.FLOAT, token.VOID, token.BOOL, token.STRING:
		return p.parseDeclarationStatement()
	case token.IF:
		return p.parseIfStatement()
	case token.IDENTIFIER:
		return p.parseIdentifierStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.WHILE:
		return p.parseWhileStatement()
	case token.FOR:
		return p.parseForStatement()
	default:
		return p.parseExpressionStatement(p.curToken, nil, true)
	}
}

func (p *Parser) parseIdentifierStatement() ast.Statement {
	tkn := p.curToken
	ident := p.parseExpression(LOWEST)
	if token.IsAssignmentOp(p.peekToken.TokenType) {
		return p.parseAssignmentStatement(tkn, ident, true, false)
	} else {
		return p.parseExpressionStatement(tkn, ident, false)
	}
}

func (p *Parser) parseExpressionStatement(tkn token.Token, exp ast.Expression, parse bool) *ast.ExpressionStatement {
	if parse {
		tkn = p.curToken
		exp = p.parseExpression(LOWEST)
	}
	p.expectPeekToken(token.SEMCOL)
	return &ast.ExpressionStatement{
		Token:      tkn,
		Expression: exp,
	}
}

func (p *Parser) parseDeclarationStatement() *ast.DeclarationStatement {
	tkn := p.curToken
	p.nextToken()
	ident := p.parseIdentifierExpression()
	var stmnt = &ast.DeclarationStatement{
		Token:      tkn,
		Type:       tkn.TokenType,
		Identifier: ident.(*ast.IdentifierExpression),
	}
	p.nextToken()
	if p.curTokenIs(token.SEMCOL) {
		return stmnt
	} else if p.curTokenIs(token.LPAREN) {
		stmnt.Literal = p.parseFunctionLiteral(stmnt.Identifier)
	} else if p.curTokenIs(token.LBRACK) {
		stmnt.Literal = p.parseArrayDeclaration(tkn, stmnt.Identifier)
	} else {
		if p.curToken.TokenType != token.ASSIGN {
			p.errors = append(p.errors, fmt.Errorf("expected '=' Sign for assigment in declaration, Got - %s", p.curToken.TokenType))
		}
		p.nextToken()
		stmnt.Literal = p.parseExpression(LOWEST)
		p.expectPeekToken(token.SEMCOL)

	}
	return stmnt
}

func (p *Parser) parseAssignmentStatement(tkn token.Token, ident ast.Expression, semCol bool, parse bool) *ast.AssignmentStatement {
	if parse {
		tkn = p.curToken
		ident = p.parseExpression(LOWEST)
	}
	var stmnt = &ast.AssignmentStatement{
		Token:      tkn,
		Identifier: ident.(ast.IdentifierNode),
	}
	p.nextToken()
	if p.curTokenIs(token.SEMCOL) {
		return stmnt
	} else {
		if !token.IsAssignmentOp(p.curToken.TokenType) {
			p.errors = append(p.errors, fmt.Errorf("expected assignment op for assigment statement , Got - %s", p.curToken.TokenType))
		}
		opTkn := p.curToken
		p.nextToken()
		exp := p.parseExpression(LOWEST)
		if opTkn.TokenType != token.ASSIGN {
			stmnt.Literal = getOpInfixExpression(ident, exp, opTkn.TokenType)
		} else {
			stmnt.Literal = exp
		}
		if semCol {
			p.expectPeekToken(token.SEMCOL)
		}
	}
	return stmnt
}

func (p *Parser) parseBlockStatement() *ast.Block {
	p.nextToken()
	blk := &ast.Block{}
	for !p.curTokenIs(token.RBRACE) {
		statement := p.ParseStatement()
		blk.Statements = append(blk.Statements, statement)
		p.nextToken()
	}
	return blk
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmnt := &ast.ReturnStatement{
		Token: p.curToken,
	}
	p.nextToken()
	if p.curTokenIs(token.SEMCOL) {
		return stmnt
	}
	expr := p.parseExpression(LOWEST)
	stmnt.Expression = expr
	p.expectPeekToken(token.SEMCOL)
	return stmnt
}

func (p *Parser) parseIfStatement() *ast.IfStatement {
	stmnt := &ast.IfStatement{
		Token: p.curToken,
	}
	p.expectPeekToken(token.LPAREN)
	stmnt.Condition = p.parseExpression(LOWEST)
	p.expectPeekToken(token.LBRACE)
	stmnt.Block = p.parseBlockStatement()
	if p.peekTokenIs(token.ELSE) {
		p.nextToken()
		p.nextToken()
		stmnt.ElseBlock = p.parseBlockStatement()
	}
	return stmnt
}

func (p *Parser) parseWhileStatement() *ast.WhileStatement {
	stmnt := &ast.WhileStatement{
		Token: p.curToken,
	}
	p.expectPeekToken(token.LPAREN)
	stmnt.Condition = p.parseExpression(LOWEST)
	p.expectPeekToken(token.LBRACE)
	stmnt.Block = p.parseBlockStatement()
	return stmnt
}

func (p *Parser) parseForStatement() *ast.ForStatement {
	stmnt := &ast.ForStatement{
		Token: p.curToken,
	}
	p.expectPeekToken(token.LPAREN)
	p.nextToken()
	if p.curTokenIs(token.IDENTIFIER) {
		stmnt.InitializationStatement = p.parseAssignmentStatement(p.curToken, nil, true, true)
	} else if token.IsDatatype(p.curToken.TokenType) {
		stmnt.InitializationStatement = p.parseDeclarationStatement()
	} else if p.curTokenIs(token.SEMCOL) {
	} else {
		p.errors = append(p.errors, fmt.Errorf("not valid statement in initializationStatement of for loop, got %s", p.curToken.TokenType))
		return nil
	}
	p.nextToken()
	stmnt.Condition = p.parseExpression(LOWEST)
	p.expectPeekToken(token.SEMCOL)
	p.nextToken()
	stmnt.Increment = p.parseAssignmentStatement(p.curToken, nil, false, true)
	p.expectPeekToken(token.RPAREN)
	p.expectPeekToken(token.LBRACE)
	stmnt.Block = p.parseBlockStatement()
	return stmnt
}
