package eval

import (
	"fmt"

	"github.com/mohamedirfanam/cynterpreter/eval/obj"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

func evalIfStatement(ifs *ast.IfStatement, env *obj.Environment) obj.Object {
	result := Eval(ifs.Condition, env)
	condition := IsTrue(result)
	if condition {
		return evalBlock(ifs.Block, env)
	} else if !condition && ifs.ElseBlock != nil {
		return evalBlock(ifs.ElseBlock, env)
	}
	return obj.NULL
}

func evalBlock(blk *ast.Block, env *obj.Environment) obj.Object {
	var results []obj.Object
	for _, stmnt := range blk.Statements {
		result := Eval(stmnt, env)
		if result.Type() == obj.ERROR_OBJ || result.Type() == obj.RETURN_OBJ {
			return result
		}
		results = append(results, result)
	}
	return &obj.ResultsObject{Results: results}
}

func evalDeclarationStatement(ls *ast.DeclarationStatement, env *obj.Environment) obj.Object {
	if _, ok := env.GetVar(ls.Identifier.Value); ok {
		return obj.NewError(fmt.Errorf("variable redeclaration error: variable %s already declared before", ls.Identifier))
	}
	if ls.Literal == nil {
		env.SetVar(ls.Identifier.Value, obj.GetDefaultVal(ls.Type))
		return obj.NULL
	}
	val := Eval(ls.Literal, env)
	if val.Type() == obj.ERROR_OBJ {
		return val
	}
	if fl, ok := ls.Literal.(*ast.FunctionLiteral); ok {
		returnType := obj.GetObjectType(ls.Type)
		functionObj := obj.GetFunctionObject(returnType, fl)
		env.SetVar(ls.Identifier.Value, functionObj)
		return obj.NULL
	} else if obj.GetObjectType(ls.Type) == val.Type() {
		env.SetVar(ls.Identifier.Value, val)
		return obj.NULL
	}
	return obj.NewError(fmt.Errorf("type error: invalid declaration type cannot assign %s to %s", val.Type(), ls.Type))
}

func evalAssignmentStatement(ls *ast.AssignmentStatement, env *obj.Environment) obj.Object {
	val := Eval(ls.Literal, env)
	if val.Type() == obj.ERROR_OBJ {
		return val
	}
	varObj, ok := env.GetVar(ls.Identifier.Value)
	if !ok {
		return obj.NewError(fmt.Errorf("variable not declared: variable %s not declared before, for assigment", ls.Identifier))
	}
	if varObj.Type() != val.Type() {
		return obj.NewError(fmt.Errorf("type error: invalid assigment type cannot assign %s to %s", val.Type(), varObj.Type()))
	}
	env.SetVar(ls.Identifier.Value, val)
	return obj.NULL
}

func evalReturnStatement(rs *ast.ReturnStatement, env *obj.Environment) obj.Object {
	expr := Eval(rs.Expression, env)
	return &obj.ReturnObject{
		Return: expr,
	}
}
