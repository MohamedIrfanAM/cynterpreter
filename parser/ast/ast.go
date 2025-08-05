package ast

import (
	"bytes"
)

type Node interface {
	TokenLexeme() string
	String() string
}

type Expression interface {
	Node
	expressionNode()
}

type Statement interface {
	Node
	statementNode()
}

// Root Node
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLexeme() string {
	if len(p.Statements) >= 1 {
		return p.Statements[0].String()
	}
	return ""
}

func (p *Program) String() string {
	var buf bytes.Buffer
	for _, statement := range p.Statements {
		buf.WriteString(statement.String())
		buf.WriteString("\n")
	}
	return buf.String()
}
