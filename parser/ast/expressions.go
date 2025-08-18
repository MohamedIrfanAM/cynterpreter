package ast

import (
	"fmt"
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

// Bool Literal
type BoolLiteral struct {
	Token token.Token
	Value bool
}

func (bl *BoolLiteral) TokenLexeme() string {
	return bl.Token.Lexeme
}

func (bl *BoolLiteral) expressionNode() {}

func (bl *BoolLiteral) String() string {
	return bl.Token.Lexeme
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

// Prefix Expression Node
type PrefixExpression struct {
	Token token.Token
	Exp   Expression
	Op    string
}

func (prefixExp *PrefixExpression) TokenLexeme() string {
	return prefixExp.Token.Lexeme
}

func (prefixExp *PrefixExpression) expressionNode() {}

func (prefixExp *PrefixExpression) String() string {
	var str strings.Builder

	str.WriteString("(")
	str.WriteString(prefixExp.Op)
	str.WriteString(prefixExp.Exp.String())
	str.WriteString(")")

	return str.String()
}

// Call Expression Node
type CallExpression struct {
	Token    token.Token
	Function Expression
	Args     []Expression
}

func (ce *CallExpression) TokenLexeme() string {
	return ce.Token.Lexeme
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) String() string {
	var str strings.Builder

	str.WriteString(ce.Function.TokenLexeme() + "(")
	for i, exp := range ce.Args {
		str.WriteString(exp.String())
		if i < len(ce.Args)-1 {
			str.WriteString(", ")
		}
	}
	str.WriteByte(')')
	return str.String()
}

// Function Literal Node
type FunctionLiteral struct {
	Token    token.Token
	Function *IdentifierExpression
	Params   []*Parameter
	Block    *Block
}

func (fl *FunctionLiteral) TokenLexeme() string {
	return fl.Token.Lexeme
}

func (fl *FunctionLiteral) expressionNode() {}

func (fl *FunctionLiteral) String() string {
	var str strings.Builder
	str.WriteString(fl.Function.String() + "(")
	for i, param := range fl.Params {
		str.WriteString(param.String())
		if i < len(fl.Params)-1 {
			str.WriteString(",")
		}
	}
	str.WriteString(")")
	str.WriteString(fl.Block.String())
	return str.String()
}

type Parameter struct {
	Token      token.Token
	Type       token.TokenType
	Identifier *IdentifierExpression
}

func (param Parameter) TokenLexeme() string {
	return param.Token.Lexeme
}
func (param Parameter) String() string {
	return param.TokenLexeme() + " " + param.Identifier.String()
}

// Array Declaration Node
type ArrayDeclaration struct {
	Token     token.Token
	Type      token.TokenType
	Identifer IdentifierExpression
	Length    int
	Literal   []Expression
}

func (arr ArrayDeclaration) TokenLexeme() string {
	return arr.Token.Lexeme
}
func (arr ArrayDeclaration) expressionNode() {}
func (arr ArrayDeclaration) String() string {
	var str strings.Builder
	str.WriteString(arr.Identifer.Value + "[" + fmt.Sprint(arr.Length) + "]")
	str.WriteString(" =  {")
	for i, lit := range arr.Literal {
		str.WriteString(lit.String())
		if i < len(arr.Literal)-1 {
			str.WriteString(",")
		}
	}
	str.WriteString("}")
	return str.String()
}

// Array index Node
type ArrayExpression struct {
	Token     token.Token
	Identifer *IdentifierExpression
	Index     Expression
}

func (arr ArrayExpression) TokenLexeme() string {
	return arr.Token.Lexeme
}
func (arr ArrayExpression) expressionNode() {}
func (arr ArrayExpression) String() string {
	var str strings.Builder
	str.WriteString(arr.Identifer.Value + "[" + fmt.Sprint(arr.Index) + "]")
	return str.String()
}
