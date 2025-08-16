package main

import (
	"fmt"
	"os"

	"github.com/mohamedirfanam/cynterpreter/batch"
	"github.com/mohamedirfanam/cynterpreter/repl"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Cynterpreter: A C interprer")
		fmt.Print("REPL Mode \n\n")

		repl.REPL(os.Stdin, os.Stdout)
	}

	batch.HandleFile(os.Args[1])
}
