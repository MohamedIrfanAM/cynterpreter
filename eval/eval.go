package eval

import (
	"github.com/mohamedirfanam/cynterpreter/eval/obj"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

func Eval(node ast.Node, env *obj.Environment) obj.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.IfStatement:
		return evalIfStatement(node, env)
	case *ast.DeclarationStatement:
		return evalDeclarationStatement(node, env)
	case *ast.AssignmentStatement:
		return evalAssignmentStatement(node, env)
	case *ast.IdentifierExpression:
		return evalIdentifierExpression(node, env)
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
		return evalInfixExpression(node, env)
	case *ast.PrefixExpression:
		return evalPrefixExpression(node, env)

	}
	return obj.NULL
}

func evalProgram(statements []ast.Statement, env *obj.Environment) obj.Object {
	var result obj.Object
	for _, stmnt := range statements {
		result = Eval(stmnt, env)
		if result.Type() == obj.ERROR_OBJ {
			return result
		}
	}
	return result
}
