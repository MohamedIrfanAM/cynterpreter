package lexer

import (
	"testing"

	"github.com/mohamedirfanam/cynterpreter/lexer/token"
)

func TestPunctuatorTokens(t *testing.T) {
	var input = `( ) [ ] { } , ; * = # . ~`

	expectedTokens := []token.Token{
		{TokenType: token.LPAREN, Lexeme: "("},
		{TokenType: token.RPAREN, Lexeme: ")"},
		{TokenType: token.LBRACK, Lexeme: "["},
		{TokenType: token.RBRACK, Lexeme: "]"},
		{TokenType: token.LBRACE, Lexeme: "{"},
		{TokenType: token.RBRACE, Lexeme: "}"},
		{TokenType: token.COMMA, Lexeme: ","},
		{TokenType: token.SEMCOL, Lexeme: ";"},
		{TokenType: token.ASTER, Lexeme: "*"},
		{TokenType: token.ASSIGN, Lexeme: "="},
		{TokenType: token.PREPROC, Lexeme: "#"},
		{TokenType: token.DOT, Lexeme: "."},
		{TokenType: token.TILDE, Lexeme: "~"},
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
