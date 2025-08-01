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
