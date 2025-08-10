package ast

import (
	"strings"

	"github.com/mohamedirfanam/cynterpreter/lexer/token"
)

// Expression Statement Node
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) TokenLexeme() string {
	return es.Token.Lexeme
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) String() string {
	return es.Expression.String()
}

type DeclarationStatement struct {
	Token      token.Token
	Type       token.TokenType
	Identifier *IdentifierExpression
	Literal    Expression
}

func (ds *DeclarationStatement) TokenLexeme() string {
	return ds.Token.Lexeme
}

func (ds *DeclarationStatement) statementNode() {}

func (ds *DeclarationStatement) String() string {
	var str strings.Builder
	str.WriteString(ds.TokenLexeme() + " ")
	str.WriteString(ds.Identifier.Value)

	if ds.Literal != nil {
		str.WriteString(" = " + ds.String())
	}
	return str.String()
}
