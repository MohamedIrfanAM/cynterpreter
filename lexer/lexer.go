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

	// Check if it's int or float iteral
	if token.IsDigit(l.ch) {
		num := l.readNumber()
		return token.GetNumberToken(num)
	}

	// Check if it's a char literal
	if l.ch == '\'' {
		char := l.readCharLiteral()
		return token.GetCharToken(char)
	}

	if l.ch == '"' {
		str := l.readStringLiteral()
		return token.GetStringToken(str)
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

func (l *Lexer) readNumber() string {
	position := l.position
	for token.IsNumber(l.peekChar()) {
		l.readChar()
	}
	return l.input[position:l.pointer]
}

func (l *Lexer) readCharLiteral() string {
	position := l.position
	l.readChar()
	var escaped = false
	for ; l.ch != '\'' && l.ch != 0; l.readChar() {
		if l.ch == '\\' && !escaped {
			l.readChar()
			escaped = true
		} else {
			escaped = false
		}
	}
	l.readChar()
	char := l.input[position:l.pointer]
	return char
}

func (l *Lexer) readStringLiteral() string {
	position := l.position
	l.readChar()
	var escaped = false
	for ; l.ch != '"' && l.ch != 0; l.readChar() {
		if l.ch == '\\' && !escaped {
			l.readChar()
			escaped = true
		} else {
			escaped = false
		}
	}
	l.readChar()
	str := l.input[position:l.pointer]
	return str
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
