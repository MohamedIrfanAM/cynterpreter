package main

import (
	"fmt"
	"os"

	"github.com/mohamedirfanam/cynterpreter/repl"
)

func main() {

	fmt.Println("Cynterpreter: A C interprer")
	fmt.Print("REPL Mode \n\n")

	repl.REPL(os.Stdin, os.Stdout)
}
