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

	token, err := token.GetPunctuatorToken(l.ch)
	if !err {
		return token
	}

	return token
}
