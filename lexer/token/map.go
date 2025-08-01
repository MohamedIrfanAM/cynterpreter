package token

var TokenMap map[TokenType]string = map[TokenType]string{
	0:  "ILLEGAL",
	1:  "LPAREN",
	2:  "RPAREN",
	3:  "LBRACK",
	4:  "RBRACK",
	5:  "LBRACE",
	6:  "RBRACE",
	7:  "COMMA",
	8:  "SEMCOL",
	9:  "ASTER",
	10: "ASSIGN",
	11: "PREPROC",
	12: "DOT",
	14: "TILDE",
}

var PunctuatorMap map[byte]TokenType = map[byte]TokenType{
	'(': LPAREN,
	')': RPAREN,
	'[': LBRACK,
	']': RBRACK,
	'{': LBRACE,
	'}': RBRACE,
	',': COMMA,
	';': SEMCOL,
	'*': ASTER,
	'=': ASSIGN,
	'#': PREPROC,
	'.': DOT,
	'~': TILDE,
}
