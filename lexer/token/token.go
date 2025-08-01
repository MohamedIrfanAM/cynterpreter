package token

type TokenType int

type Token struct {
	TokenType TokenType
	Literal   string
}

const (
	ILLEGAL TokenType = 0
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
	SIZEOF    TokenType = 48 // sizeof
)

func (t TokenType) String() string {
	return TokenMap[t]
}

func GetPunctuatorToken(ch byte) (Token, bool) {
	tokenType, ok := PunctuatorMap[ch]
	if ok {
		return Token{
			TokenType: tokenType,
		}, ok
	}
	return Token{}, ok
}

func GetOperatorToken(op string) (Token, bool) {
	tokenType, ok := OperatorMap[op]
	if ok {
		return Token{
			TokenType: tokenType,
		}, ok
	}
	return Token{}, ok
}
