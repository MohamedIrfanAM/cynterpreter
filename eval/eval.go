package eval

import (
	"github.com/mohamedirfanam/cynterpreter/eval/obj"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

func Eval(node ast.Node) obj.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IfStatement:
		return evalIfStatement(node)
	case *ast.IntegerLiteral:
		return &obj.IntegerObject{Value: node.Value}
	case *ast.BoolLiteral:
		if node.Value {
			return obj.TRUE
		} else {
			return obj.FALSE
		}
	case *ast.CharLiteral:
		return &obj.CharObject{Value: node.Value}
	case *ast.StringLiteral:
		return &obj.StringObject{Value: node.Value}
	case *ast.FloatLiteral:
		return &obj.FloatObject{Value: node.Value}
	case *ast.InfixExpression:
		return evalInfixExpression(node)
	case *ast.PrefixExpression:
		return evalPrefixExpression(node)

	}
	return obj.NULL
}

func evalProgram(statements []ast.Statement) obj.Object {
	var result obj.Object
	for _, stmnt := range statements {
		result = Eval(stmnt)
		if result.Type() == obj.ERROR_OBJ {
			return result
		}
	}
	return result
}
