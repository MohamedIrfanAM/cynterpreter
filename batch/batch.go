package batch

import (
	"fmt"
	"os"

	"github.com/mohamedirfanam/cynterpreter/eval"
	"github.com/mohamedirfanam/cynterpreter/eval/obj"
	"github.com/mohamedirfanam/cynterpreter/lexer/token"
	"github.com/mohamedirfanam/cynterpreter/parser"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

func HandleFile(filName string) {

	data, err := os.ReadFile(filName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var p = parser.New(string(data))

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			fmt.Printf("Parser Error: %s\n", err.Error())
		}
		return
	}
	env := obj.NewEnv()
	eval.Eval(program, env)

	_, ok := env.GetVar("main")
	if !ok {
		fmt.Printf("Program error: No main function found")
	}

	result := eval.Eval(getMainCall(), env)
	errorObj, ok := result.(*obj.ErrorObject)
	if ok {
		fmt.Print(errorObj.Error.Error())
	}
}

func getMainCall() *ast.CallExpression {
	identTkn := token.GetIdentifierToken("main")
	funcIdent := &ast.IdentifierExpression{
		Token: identTkn,
		Value: "main",
	}
	tkn, _ := token.GetPunctuatorToken('(')
	return &ast.CallExpression{
		Token:    tkn,
		Function: funcIdent,
	}
}
