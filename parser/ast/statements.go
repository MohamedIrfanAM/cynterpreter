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

func (ds *DeclarationStatement) statementNode()           {}
func (ds *DeclarationStatement) initializationStatement() {}

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

// Assignment Statement
type IdentifierNode interface {
	Expression
	identifierNode()
}
type AssignmentStatement struct {
	Token      token.Token
	Identifier IdentifierNode
	Literal    Expression
}

func (as *AssignmentStatement) TokenLexeme() string {
	return as.Token.Lexeme
}

func (as *AssignmentStatement) statementNode()           {}
func (as *AssignmentStatement) initializationStatement() {}

func (as *AssignmentStatement) String() string {
	var str strings.Builder

	str.WriteString(as.Identifier.String())
	if as.Literal != nil {
		str.WriteString(" = " + as.Literal.String())
	}
	return str.String()
}

// While loop
type WhileStatement struct {
	Token     token.Token
	Condition Expression
	Block     *Block
}

func (wl *WhileStatement) TokenLexeme() string {
	return wl.Token.Lexeme
}

func (wl *WhileStatement) statementNode() {}

func (wl *WhileStatement) String() string {
	var str strings.Builder
	str.WriteString("while (")
	str.WriteString(wl.Condition.String() + ")")
	str.WriteString(wl.Block.String())
	return str.String()
}

// For loop
type InitializationStatement interface {
	Statement
	initializationStatement()
}
type ForStatement struct {
	Token token.Token
	InitializationStatement
	Condition Expression
	Increment *AssignmentStatement
	Block     *Block
}

func (fs *ForStatement) TokenLexeme() string {
	return fs.Token.Lexeme
}

func (fs *ForStatement) statementNode() {}

func (fs *ForStatement) String() string {
	var str strings.Builder
	str.WriteString("for (")
	str.WriteString(fs.InitializationStatement.String() + "; ")
	str.WriteString(fs.Condition.String() + ";")
	str.WriteString(fs.Increment.String() + ")")
	str.WriteString(fs.Block.String())
	return str.String()
}
