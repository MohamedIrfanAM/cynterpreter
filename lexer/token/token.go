package token

import (
	"fmt"
	"strings"
)

type TokenType int

type Token struct {
	TokenType TokenType
	Lexeme    string
}

const (
	EOF     TokenType = 0
	ILLEGAL TokenType = 13
	// Punctuators
	LPAREN  TokenType = 1  // (
	RPAREN  TokenType = 2  // )
	LBRACK  TokenType = 3  // [
	RBRACK  TokenType = 4  // ]
	LBRACE  TokenType = 5  // {
	RBRACE  TokenType = 6  // }
	COMMA   TokenType = 7  // ,
	SEMCOL  TokenType = 8  // ;
	ASTER   TokenType = 9  // *
	ASSIGN  TokenType = 10 // =
	PREPROC TokenType = 11 // #
	DOT     TokenType = 12 // .
	TILDE   TokenType = 14 // ~
	// Operators
	ARROW     TokenType = 15 // ->
	INCR      TokenType = 16 // ++
	DECR      TokenType = 17 // --
	PLUS      TokenType = 18 // +
	MINUS     TokenType = 19 // -
	NOT       TokenType = 20 // !
	SLASH     TokenType = 21 // /
	PERCENT   TokenType = 22 // %
	LSHIFT    TokenType = 23 // <<
	RSHIFT    TokenType = 24 // >>
	LT        TokenType = 25 // <
	LE        TokenType = 26 // <=
	GT        TokenType = 27 // >
	GE        TokenType = 28 // >=
	EQ        TokenType = 29 // ==
	NE        TokenType = 30 // !=
	AMP       TokenType = 31 // &
	XOR       TokenType = 32 // ^
	PIPE      TokenType = 33 // |
	AND       TokenType = 34 // &&
	OR        TokenType = 35 // ||
	QUESTION  TokenType = 36 // ?
	COLON     TokenType = 37 // :
	PLUSEQ    TokenType = 38 // +=
	MINUSEQ   TokenType = 39 // -=
	ASTEREQ   TokenType = 40 // *=
	SLASHEQ   TokenType = 41 // /=
	PERCENTEQ TokenType = 42 // %=
	AMPEQ     TokenType = 43 // &=
	XOREQ     TokenType = 44 // ^=
	PIPEEQ    TokenType = 45 // |=
	LSHIFTEQ  TokenType = 46 // <<=
	RSHIFTEQ  TokenType = 47 // >>=
	//Keywords
	AUTO     TokenType = 48 // auto
	BREAK    TokenType = 49 // break
	CASE     TokenType = 50 // case
	CHAR     TokenType = 51 // char
	CONST    TokenType = 52 // const
	CONTINUE TokenType = 53 // continue
	DEFAULT  TokenType = 54 // default
	DO       TokenType = 55 // do
	DOUBLE   TokenType = 56 // double
	ELSE     TokenType = 57 // else
	ENUM     TokenType = 58 // enum
	EXTERN   TokenType = 59 // extern
	FLOAT    TokenType = 60 // float
	FOR      TokenType = 61 // for
	GOTO     TokenType = 62 // goto
	IF       TokenType = 63 // if
	INLINE   TokenType = 64 // inline
	INT      TokenType = 65 // int
	LONG     TokenType = 66 // long
	REGISTER TokenType = 67 // register
	RESTRICT TokenType = 68 // restrict
	RETURN   TokenType = 69 // return
	SHORT    TokenType = 70 // short
	SIGNED   TokenType = 71 // signed
	SIZEOF   TokenType = 72 // sizeof
	STATIC   TokenType = 73 // static
	STRUCT   TokenType = 74 // struct
	SWITCH   TokenType = 75 // switch
	TYPEDEF  TokenType = 76 // typedef
	UNION    TokenType = 77 // union
	UNSIGNED TokenType = 78 // unsigned
	VOID     TokenType = 79 // void
	VOLATILE TokenType = 80 // volatile
	WHILE    TokenType = 81 // while

	// Identifiers and Literals
	IDENTIFIER     TokenType = 92 // user-defined names
	INT_LITERAL    TokenType = 93 // integer constants
	FLOAT_LITERAL  TokenType = 94 // floating-point constants
	CHAR_LITERAL   TokenType = 95 // character constants
	STRING_LITERAL TokenType = 96 // string literals
)

func (t TokenType) String() string {
	return TokenMap[t]
}

func GetEofToken() Token {
	return Token{EOF, ""}
}

func GetIllegalToken() Token {
	return Token{ILLEGAL, ""}
}

func GetPunctuatorToken(ch byte) (Token, bool) {
	tokenType, ok := PunctuatorMap[ch]
	if ok {
		return Token{
			TokenType: tokenType,
			Lexeme:    string(ch),
		}, ok
	}
	return Token{}, ok
}

func GetOperatorToken(op string) (Token, bool) {
	tokenType, ok := OperatorMap[op]
	if ok {
		return Token{
			TokenType: tokenType,
			Lexeme:    op,
		}, ok
	}
	return Token{}, ok
}

func GetKeywordToken(word string) (Token, bool) {
	tokenType, ok := KeywordMap[word]
	if ok {
		return Token{
			TokenType: tokenType,
			Lexeme:    word,
		}, ok
	}
	return Token{}, ok
}

func GetIdentifierToken(word string) Token {
	return Token{
		TokenType: IDENTIFIER,
		Lexeme:    word,
	}
}

func GetNumberToken(num string) Token {
	var suffix = ""
	for i := 0; i < len(num); i++ {
		l := strings.ToLower(string(num[i]))
		if !IsDigit(num[i]) && l == "f" || l == "u" || l == "l" {
			suffix += l
		} else if len(suffix) != 0 {
			return Token{TokenType: ILLEGAL}
		}
	}

	fmt.Println(suffix)

	if strings.Count(num, ".") == 0 && len(suffix) <= 3 && strings.Count(suffix, "f") == 0 && suffix != "lul" {
		return Token{
			TokenType: INT_LITERAL,
			Lexeme:    num,
		}
	} else if strings.Count(num, ".") == 1 && len(suffix) <= 1 && strings.Count(suffix, "u") == 0 {
		return Token{
			TokenType: FLOAT_LITERAL,
			Lexeme:    num,
		}
	}
	return Token{TokenType: ILLEGAL}
}

func GetCharToken(char string) Token {
	c := char[1 : len(char)-1]
	if (char[0] == '\'' && char[len(char)-1] == '\'') && len(c) == 1 || (len(c) == 2 && c[0] == '\\') {
		return Token{
			TokenType: CHAR_LITERAL,
			Lexeme:    char,
		}
	}
	return Token{TokenType: ILLEGAL}
}

func GetStringToken(str string) Token {
	if str[0] == '"' && str[len(str)-1] == '"' {
		return Token{
			TokenType: STRING_LITERAL,
			Lexeme:    str,
		}
	}
	return Token{TokenType: ILLEGAL}
}

func IsOperatorSymbol(ch byte) bool {
	for _, v := range OpSymbols {
		if v == ch {
			return true
		}
	}
	return false
}

func IsWhiteSpace(ch byte) bool {
	for _, v := range WhiteSpaceSymbols {
		if ch == v {
			return true
		}
	}
	return false
}

func IsLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

func IsDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func IsNumber(ch byte) bool {
	lower := strings.ToLower(string(ch))
	return IsDigit(ch) || ch == '.' || lower == "u" || lower == "f" || lower == "l"
}

func IsWordSymbol(ch byte) bool {
	return IsLetter(ch) || IsDigit(ch) || ch == '_'
}

func IsWordStartSymbol(ch byte) bool {
	return IsLetter(ch) || ch == '_'
}
