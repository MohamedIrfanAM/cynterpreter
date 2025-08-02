package lexer

import (
	"github.com/mohamedirfanam/cynterpreter/lexer/token"
)

type Lexer struct {
	input    string
	ch       byte
	position int
	pointer  int
}

func New(input string) *Lexer {
	var l Lexer = Lexer{
		input:    input,
		position: -1,
		pointer:  0,
	}
	return &l
}

func (l *Lexer) readChar() {
	if l.pointer >= len(l.input) {
		l.ch = 0
		return
	}
	l.ch = l.input[l.pointer]
	l.position = l.pointer
	l.pointer++
}

func (l *Lexer) NextToken() token.Token {
	l.readChar()
	l.skipWhiteSpace()

	if l.ch == 0 {
		return token.GetEofToken()
	}

	// Check if it's a Punctuator
	tkn, found := token.GetPunctuatorToken(l.ch)
	if found {
		return tkn
	}

	// Check if it's an Operator
	if token.IsOperatorSymbol(l.ch) {
		op := l.readOperator()
		tkn, found = token.GetOperatorToken(op)
		if found {
			return tkn
		}
	}

	// Check if it's a keyword or identifier
	if token.IsWordStartSymbol(l.ch) {
		word := l.readWord()

		tkn, found = token.GetKeywordToken(word)
		if found {
			return tkn
		}

		return token.GetIdentifierToken(word)
	}

	return token.GetIllegalToken()
}

func (l *Lexer) readOperator() string {
	position := l.position
	for token.IsOperatorSymbol(l.peekChar()) {
		l.readChar()
	}
	return l.input[position:l.pointer]
}

func (l *Lexer) readWord() string {
	position := l.position
	for token.IsWordSymbol(l.peekChar()) {
		l.readChar()
	}
	return l.input[position:l.pointer]
}

func (l *Lexer) peekChar() byte {
	if l.pointer >= len(l.input) {
		return 0
	}
	return l.input[l.pointer]
}

func (l *Lexer) skipWhiteSpace() {
	for token.IsWhiteSpace(l.ch) {
		l.readChar()
	}
}
