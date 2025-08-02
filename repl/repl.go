package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/mohamedirfanam/cynterpreter/lexer"
	"github.com/mohamedirfanam/cynterpreter/lexer/token"
)

func REPL(in io.Reader, out io.Writer) {

	var scanner *bufio.Scanner = bufio.NewScanner(in)
	fmt.Fprint(out, ">> ")

	for scanner.Scan() {

		var l = lexer.New(scanner.Text())

		for tkn := l.NextToken(); tkn.TokenType != token.EOF; tkn = l.NextToken() {
			fmt.Fprintln(out, "Token = ", tkn.TokenType, ", Lexeme = ", tkn.Lexeme)
		}

		fmt.Fprint(out, ">> ")
	}
}
