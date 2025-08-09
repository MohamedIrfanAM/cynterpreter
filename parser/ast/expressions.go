package ast

import (
	"strings"

	"github.com/mohamedirfanam/cynterpreter/lexer/token"
)

// Integer literal
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) TokenLexeme() string {
	return il.Token.Lexeme
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) String() string {
	return il.Token.Lexeme
}

// Identifier expression
type IdentifierExpression struct {
	Token token.Token
	Value string
}

func (il *IdentifierExpression) TokenLexeme() string {
	return il.Token.Lexeme
}
func (il *IdentifierExpression) expressionNode() {}

func (il *IdentifierExpression) String() string {
	return il.Value
}

// Infix Expression Node
type InfixExpression struct {
	Token    token.Token
	LeftExp  Expression
	Op       string
	RightExp Expression
}

func (infExp *InfixExpression) TokenLexeme() string {
	return infExp.Token.Lexeme
}

func (infExp *InfixExpression) expressionNode() {}

func (infExp *InfixExpression) String() string {
	var str strings.Builder

	str.WriteString("(")
	str.WriteString(infExp.LeftExp.String())
	str.WriteString(" ")
	str.WriteString(infExp.Op)
	str.WriteString(" ")
	str.WriteString(infExp.RightExp.String())
	str.WriteString(")")

	return str.String()
}
