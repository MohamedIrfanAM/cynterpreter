package ast

import "github.com/mohamedirfanam/cynterpreter/lexer/token"

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
