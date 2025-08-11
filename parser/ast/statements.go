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

// Declaration Statement
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

	if _, ok := ds.Literal.(*FunctionLiteral); ok {
		str.WriteString(ds.Literal.String())
	} else if ds.Literal != nil {
		str.WriteString(ds.Identifier.Value)
		str.WriteString(" = " + ds.Literal.String())
	} else {
		str.WriteString(ds.Identifier.Value)
	}
	return str.String()
}

// Block statement
type Block struct {
	Statements []Statement
}

func (blk Block) String() string {
	var str strings.Builder
	str.WriteString("{\n")
	for _, stmnt := range blk.Statements {
		str.WriteString("\t" + stmnt.String() + ";\n")
	}
	str.WriteString("}\n")
	return str.String()
}

// Return statement
type ReturnStatement struct {
	Token      token.Token
	Expression Expression
}

func (rs *ReturnStatement) TokenLexeme() string {
	return rs.Token.Lexeme
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) String() string {
	return "return " + rs.Expression.String()
}

// If statement
type IfStatement struct {
	Token     token.Token
	Condition Expression
	Block     *Block
	ElseBlock *Block
}

func (ifs *IfStatement) TokenLexeme() string {
	return ifs.Token.Lexeme
}

func (ifs *IfStatement) statementNode() {}

func (ifs *IfStatement) String() string {
	var str strings.Builder
	str.WriteString("if (")
	str.WriteString(ifs.Condition.String() + ")")
	str.WriteString(ifs.Block.String())
	if ifs.ElseBlock != nil {
		str.WriteString("else ")
		str.WriteString(ifs.ElseBlock.String())
	}
	return str.String()
}
