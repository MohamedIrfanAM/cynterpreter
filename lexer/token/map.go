package token

var TokenMap map[TokenType]string = map[TokenType]string{
	1: "LPAREN",
	2: "RPAREN",
}

var PunctuatorMap map[byte]TokenType = map[byte]TokenType{
	'(': LPAREN,
	')': RPAREN,
}
