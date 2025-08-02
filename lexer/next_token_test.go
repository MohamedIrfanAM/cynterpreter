package lexer

import (
	"testing"

	"github.com/mohamedirfanam/cynterpreter/lexer/token"
)

func TestPunctuatorTokens(t *testing.T) {
	var input = `( ) [ ] { } , ; # . ~`

	expectedTokens := []token.Token{
		{TokenType: token.LPAREN, Lexeme: "("},
		{TokenType: token.RPAREN, Lexeme: ")"},
		{TokenType: token.LBRACK, Lexeme: "["},
		{TokenType: token.RBRACK, Lexeme: "]"},
		{TokenType: token.LBRACE, Lexeme: "{"},
		{TokenType: token.RBRACE, Lexeme: "}"},
		{TokenType: token.COMMA, Lexeme: ","},
		{TokenType: token.SEMCOL, Lexeme: ";"},
		{TokenType: token.PREPROC, Lexeme: "#"},
		{TokenType: token.DOT, Lexeme: "."},
		{TokenType: token.TILDE, Lexeme: "~"},
		{TokenType: token.EOF, Lexeme: ""},
	}

	var l = New(input)

	for i, expectedToken := range expectedTokens {
		var tkn token.Token = l.NextToken()

		if tkn.TokenType != expectedToken.TokenType {
			t.Errorf("[%d] - Wrong TokenType, Expected - %s, got - %s", i, expectedToken.TokenType, tkn.TokenType)
		}

		if tkn.Lexeme != expectedToken.Lexeme {
			t.Fatalf("[%d] - Wrong Lexeme, Expected - %s, got - %s", i, expectedToken.Lexeme, tkn.Lexeme)
		}
	}
}

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
			t.Errorf("[%d] - Wrong TokenType, Expected - %s, got - %s", i, expectedToken.TokenType, tkn.TokenType)
		}

		if tkn.Lexeme != expectedToken.Lexeme {
			t.Fatalf("[%d] - Wrong Lexeme, Expected - %s, got - %s", i, expectedToken.Lexeme, tkn.Lexeme)
		}
	}
}

func TestKeywordTokens(t *testing.T) {
	var input = `auto break case char const continue default do double else enum extern float for goto if inline int long register restrict return short signed sizeof static struct switch typedef union unsigned void volatile while`

	expectedTokens := []token.Token{
		{TokenType: token.AUTO, Lexeme: "auto"},
		{TokenType: token.BREAK, Lexeme: "break"},
		{TokenType: token.CASE, Lexeme: "case"},
		{TokenType: token.CHAR, Lexeme: "char"},
		{TokenType: token.CONST, Lexeme: "const"},
		{TokenType: token.CONTINUE, Lexeme: "continue"},
		{TokenType: token.DEFAULT, Lexeme: "default"},
		{TokenType: token.DO, Lexeme: "do"},
		{TokenType: token.DOUBLE, Lexeme: "double"},
		{TokenType: token.ELSE, Lexeme: "else"},
		{TokenType: token.ENUM, Lexeme: "enum"},
		{TokenType: token.EXTERN, Lexeme: "extern"},
		{TokenType: token.FLOAT, Lexeme: "float"},
		{TokenType: token.FOR, Lexeme: "for"},
		{TokenType: token.GOTO, Lexeme: "goto"},
		{TokenType: token.IF, Lexeme: "if"},
		{TokenType: token.INLINE, Lexeme: "inline"},
		{TokenType: token.INT, Lexeme: "int"},
		{TokenType: token.LONG, Lexeme: "long"},
		{TokenType: token.REGISTER, Lexeme: "register"},
		{TokenType: token.RESTRICT, Lexeme: "restrict"},
		{TokenType: token.RETURN, Lexeme: "return"},
		{TokenType: token.SHORT, Lexeme: "short"},
		{TokenType: token.SIGNED, Lexeme: "signed"},
		{TokenType: token.SIZEOF, Lexeme: "sizeof"},
		{TokenType: token.STATIC, Lexeme: "static"},
		{TokenType: token.STRUCT, Lexeme: "struct"},
		{TokenType: token.SWITCH, Lexeme: "switch"},
		{TokenType: token.TYPEDEF, Lexeme: "typedef"},
		{TokenType: token.UNION, Lexeme: "union"},
		{TokenType: token.UNSIGNED, Lexeme: "unsigned"},
		{TokenType: token.VOID, Lexeme: "void"},
		{TokenType: token.VOLATILE, Lexeme: "volatile"},
		{TokenType: token.WHILE, Lexeme: "while"},
		{TokenType: token.EOF, Lexeme: ""},
	}

	var l = New(input)

	for i, expectedToken := range expectedTokens {
		var tkn token.Token = l.NextToken()

		if tkn.TokenType != expectedToken.TokenType {
			t.Errorf("[%d] - Wrong TokenType, Expected - %s, got - %s", i, expectedToken.TokenType, tkn.TokenType)
		}

		if tkn.Lexeme != expectedToken.Lexeme {
			t.Fatalf("[%d] - Wrong Lexeme, Expected - %s, got - %s", i, expectedToken.Lexeme, tkn.Lexeme)
		}
	}
}

func TestIdentifierTokens(t *testing.T) {
	var input = `variable_name myFunction _underscore identifier123 main argc argv`

	expectedTokens := []token.Token{
		{TokenType: token.IDENTIFIER, Lexeme: "variable_name"},
		{TokenType: token.IDENTIFIER, Lexeme: "myFunction"},
		{TokenType: token.IDENTIFIER, Lexeme: "_underscore"},
		{TokenType: token.IDENTIFIER, Lexeme: "identifier123"},
		{TokenType: token.IDENTIFIER, Lexeme: "main"},
		{TokenType: token.IDENTIFIER, Lexeme: "argc"},
		{TokenType: token.IDENTIFIER, Lexeme: "argv"},
		{TokenType: token.EOF, Lexeme: ""},
	}

	var l = New(input)

	for i, expectedToken := range expectedTokens {
		var tkn token.Token = l.NextToken()

		if tkn.TokenType != expectedToken.TokenType {
			t.Errorf("[%d] - Wrong TokenType, Expected - %s, got - %s", i, expectedToken.TokenType, tkn.TokenType)
		}

		if tkn.Lexeme != expectedToken.Lexeme {
			t.Fatalf("[%d] - Wrong Lexeme, Expected - %s, got - %s", i, expectedToken.Lexeme, tkn.Lexeme)
		}
	}
}

func TestLiterals(t *testing.T) {
	var input = `42 123 3.14 2.5f 'a' 'Z' '\n' "hello" "world" "hello world"`

	expectedTokens := []token.Token{
		{TokenType: token.INT_LITERAL, Lexeme: "42"},
		{TokenType: token.INT_LITERAL, Lexeme: "123"},
		{TokenType: token.FLOAT_LITERAL, Lexeme: "3.14"},
		{TokenType: token.FLOAT_LITERAL, Lexeme: "2.5f"},
		{TokenType: token.CHAR_LITERAL, Lexeme: "'a'"},
		{TokenType: token.CHAR_LITERAL, Lexeme: "'Z'"},
		{TokenType: token.CHAR_LITERAL, Lexeme: "'\\n'"},
		{TokenType: token.STRING_LITERAL, Lexeme: "\"hello\""},
		{TokenType: token.STRING_LITERAL, Lexeme: "\"world\""},
		{TokenType: token.STRING_LITERAL, Lexeme: "\"hello world\""},
		{TokenType: token.EOF, Lexeme: ""},
	}

	var l = New(input)

	for i, expectedToken := range expectedTokens {
		var tkn token.Token = l.NextToken()

		if tkn.TokenType != expectedToken.TokenType {
			t.Errorf("[%d] - Wrong TokenType, Expected - %s, got - %s", i, expectedToken.TokenType, tkn.TokenType)
		}

		if tkn.Lexeme != expectedToken.Lexeme {
			t.Errorf("[%d] - Wrong Lexeme, Expected - %s, got - %s", i, expectedToken.Lexeme, tkn.Lexeme)
		}
	}
}
