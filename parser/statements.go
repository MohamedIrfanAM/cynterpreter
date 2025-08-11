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
		return p.parseAssignmentStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.WHILE:
		return p.parseWhileStatement()
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

func (p *Parser) parseAssignmentStatement() *ast.AssignmentStatement {
	tkn := p.curToken
	ident := p.parseIdentifierExpression()
	var stmnt = &ast.AssignmentStatement{
		Token:      tkn,
		Identifier: ident.(*ast.IdentifierExpression),
	}
	p.nextToken()
	if p.curTokenIs(token.SEMCOL) {
		return stmnt
	} else {
		if p.curToken.TokenType != token.ASSIGN {
			p.errors = append(p.errors, fmt.Errorf("expected '=' Sign for assigment statement , Got - %s", p.curToken.TokenType))
		}
		p.nextToken()
		stmnt.Literal = p.parseExpression(LOWEST)
		p.expectPeekToken(token.SEMCOL)

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
	p.nextToken()
	if p.curTokenIs(token.ELSE) {
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
	p.nextToken()
	return stmnt
}
