package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/mohamedirfanam/cynterpreter/parser"
)

func REPL(in io.Reader, out io.Writer) {

	var scanner *bufio.Scanner = bufio.NewScanner(in)
	fmt.Fprint(out, ">> ")

	for scanner.Scan() {

		var p = parser.New(scanner.Text())

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			for _, err := range p.Errors() {
				fmt.Fprintf(out, "Parser Error: %s\n", err.Error())
			}
			continue
		}

		fmt.Println(program.Statements[0])

		fmt.Fprint(out, ">> ")
	}
}
