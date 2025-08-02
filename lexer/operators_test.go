package lexer

import (
	"testing"

	"github.com/mohamedirfanam/cynterpreter/lexer/token"
)

func TestOperatorTokens(t *testing.T) {
	var input = `-> ++ -- + - * = ! / % << >> < <= > >= == != & ^ | && || ? : += -= *= /= %= &= ^= |= <<= >>=`

	expectedTokens := []token.Token{
		{TokenType: token.ARROW, Lexeme: "->"},
		{TokenType: token.INCR, Lexeme: "++"},
		{TokenType: token.DECR, Lexeme: "--"},
		{TokenType: token.PLUS, Lexeme: "+"},
		{TokenType: token.MINUS, Lexeme: "-"},
		{TokenType: token.ASTER, Lexeme: "*"},
		{TokenType: token.ASSIGN, Lexeme: "="},
		{TokenType: token.NOT, Lexeme: "!"},
		{TokenType: token.SLASH, Lexeme: "/"},
		{TokenType: token.PERCENT, Lexeme: "%"},
		{TokenType: token.LSHIFT, Lexeme: "<<"},
		{TokenType: token.RSHIFT, Lexeme: ">>"},
		{TokenType: token.LT, Lexeme: "<"},
		{TokenType: token.LE, Lexeme: "<="},
		{TokenType: token.GT, Lexeme: ">"},
		{TokenType: token.GE, Lexeme: ">="},
		{TokenType: token.EQ, Lexeme: "=="},
		{TokenType: token.NE, Lexeme: "!="},
		{TokenType: token.AMP, Lexeme: "&"},
		{TokenType: token.XOR, Lexeme: "^"},
		{TokenType: token.PIPE, Lexeme: "|"},
		{TokenType: token.AND, Lexeme: "&&"},
		{TokenType: token.OR, Lexeme: "||"},
		{TokenType: token.QUESTION, Lexeme: "?"},
		{TokenType: token.COLON, Lexeme: ":"},
		{TokenType: token.PLUSEQ, Lexeme: "+="},
		{TokenType: token.MINUSEQ, Lexeme: "-="},
		{TokenType: token.ASTEREQ, Lexeme: "*="},
		{TokenType: token.SLASHEQ, Lexeme: "/="},
		{TokenType: token.PERCENTEQ, Lexeme: "%="},
		{TokenType: token.AMPEQ, Lexeme: "&="},
		{TokenType: token.XOREQ, Lexeme: "^="},
		{TokenType: token.PIPEEQ, Lexeme: "|="},
		{TokenType: token.LSHIFTEQ, Lexeme: "<<="},
		{TokenType: token.RSHIFTEQ, Lexeme: ">>="},
		{TokenType: token.EOF, Lexeme: ""},
	}

	var l = New(input)

	for i, expectedToken := range expectedTokens {
		var tkn token.Token = l.NextToken()

		if tkn.TokenType != expectedToken.TokenType {
			t.Errorf("[%d] - Wrong TokenType, Expected - %s, got - %s", i, tkn.TokenType, expectedToken.TokenType)
		}

		if tkn.Lexeme != expectedToken.Lexeme {
			t.Fatalf("[%d] - Wrong Lexeme, Expected - %s, got - %s", i, tkn.Lexeme, expectedToken.Lexeme)
		}
	}
}
