package lexer

import "github.com/mohamedirfanam/cynterpreter/lexer/token"

func isOperatorSymbol(ch byte) bool {
	for _, v := range token.OpSymbols {
		if v == ch {
			return true
		}
	}
	return false
}
