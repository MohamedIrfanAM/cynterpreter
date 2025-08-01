package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mohamedirfanam/cynterpreter/lexer"
)

func main() {

	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	fmt.Print(">> ")

	for scanner.Scan() {

		var l = lexer.New(scanner.Text())

		for i := range scanner.Text() {
			token := l.NextToken()
			fmt.Println(i+1, " - Token = ", token.TokenType)
		}

		fmt.Print(">> ")
	}
}
