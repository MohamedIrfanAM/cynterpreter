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

// Char Literal
type CharLiteral struct {
	Token token.Token
	Value byte
}

func (cl *CharLiteral) TokenLexeme() string {
	return cl.Token.Lexeme
}

func (cl *CharLiteral) expressionNode() {}

func (cl *CharLiteral) String() string {
	return cl.Token.Lexeme
}

// Float literal
type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (fl *FloatLiteral) TokenLexeme() string {
	return fl.Token.Lexeme
}

func (fl *FloatLiteral) expressionNode() {}

func (fl *FloatLiteral) String() string {
	return fl.Token.Lexeme
}

// String Literal
type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) TokenLexeme() string {
	return sl.Token.Lexeme
}

func (sl *StringLiteral) expressionNode() {}

func (sl *StringLiteral) String() string {
	return sl.Token.Lexeme
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
