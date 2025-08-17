package eval

import (
	"fmt"

	"github.com/mohamedirfanam/cynterpreter/eval/obj"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

type BuildInFunc func(args ...obj.Object) obj.Object

var BuiltInFuncMap = map[string]BuildInFunc{
	"print":  print,
	"printf": printf,
}

func ApplyBuiltInFunc(funcName string, args []ast.Expression, env *obj.Environment) (obj.Object, bool) {
	buildInfunc, ok := BuiltInFuncMap[funcName]
	if !ok {
		return nil, false
	}
	var argsObjs []obj.Object
	for _, arg := range args {
		argObj := Eval(arg, env)
		argsObjs = append(argsObjs, argObj)
	}
	result := buildInfunc(argsObjs...)
	return result, true
}

func print(args ...obj.Object) obj.Object {
	var vals []any
	for _, arg := range args {
		vals = append(vals, obj.ExtractVal(arg))
	}
	fmt.Print(vals...)
	return obj.NULL
}

func printf(args ...obj.Object) obj.Object {
	format, ok := args[0].(*obj.StringObject)
	if !ok {
		return obj.NewError(fmt.Errorf("format string expected for printf"))
	}
	var vals []any
	for _, arg := range args[1:] {
		vals = append(vals, obj.ExtractVal(arg))
	}
	fmt.Printf(format.Value, vals...)
	return obj.NULL
}
