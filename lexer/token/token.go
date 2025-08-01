package token

type TokenType int

type Token struct {
	TokenType TokenType
	Literal   string
}

const (
	ILLEGAL TokenType = 0
	LPAREN  TokenType = 1
	RPAREN  TokenType = 2
	LBRACK  TokenType = 3
	RBRACK  TokenType = 4
	LBRACE  TokenType = 5
	RBRACE  TokenType = 6
	COMMA   TokenType = 7
	SEMCOL  TokenType = 8
	ASTER   TokenType = 9
	ASSIGN  TokenType = 10
	PREPROC TokenType = 11
	DOT     TokenType = 12
	TILDE   TokenType = 14
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
