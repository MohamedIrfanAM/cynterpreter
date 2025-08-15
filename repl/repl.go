package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/mohamedirfanam/cynterpreter/eval"
	"github.com/mohamedirfanam/cynterpreter/eval/obj"
	"github.com/mohamedirfanam/cynterpreter/parser"
)

func REPL(in io.Reader, out io.Writer) {

	var scanner *bufio.Scanner = bufio.NewScanner(in)
	fmt.Fprint(out, ">> ")

	var input strings.Builder
	for scanner.Scan() {

		input.WriteString(scanner.Text())
		if !isBalanced(input.String()) {
			fmt.Fprint(out, ">>> ")
			continue
		}
		var p = parser.New(input.String())
		input.Reset()

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			for _, err := range p.Errors() {
				fmt.Fprintf(out, "Parser Error: %s\n", err.Error())
			}
			fmt.Fprint(out, ">>  ")
			continue
		}

		result := eval.Eval(program.Statements[0])
		if result.Type() != obj.NULL_OBJ {
			fmt.Println(result.String())
		}

		fmt.Fprint(out, ">> ")
	}
}

func isBalanced(input string) bool {
	lParentCount := 0
	for _, c := range input {
		switch c {
		case '{':
			lParentCount++
		case '}':
			lParentCount--
		}
	}
	return lParentCount == 0
}
