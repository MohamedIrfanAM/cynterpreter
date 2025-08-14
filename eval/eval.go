package eval

import (
	"github.com/mohamedirfanam/cynterpreter/eval/obj"
	"github.com/mohamedirfanam/cynterpreter/lexer/token"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

func Eval(node ast.Node) obj.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
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
	}
	return result
}

func evalPrefixExpression(expr *ast.PrefixExpression) obj.Object {
	val := Eval(expr.Exp)
	switch expr.Token.TokenType {
	case token.MINUS:
		return evalPrefixMinusOp(val)
	case token.NOT:
		return evalPrefixNotOp(val)
	default:
		return &obj.ErrorObject{Value: "Not a valid prefix operator"}
	}
}

func evalPrefixMinusOp(val obj.Object) obj.Object {
	switch val := val.(type) {
	case *obj.IntegerObject:
		return &obj.IntegerObject{Value: -1 * val.Value}
	case *obj.FloatObject:
		return &obj.FloatObject{Value: -1 * val.Value}
	default:
		return &obj.ErrorObject{Value: "Not a valid type, for the operator -"}
	}
}

func evalPrefixNotOp(val obj.Object) obj.Object {
	switch val {
	case obj.TRUE:
		return obj.FALSE
	case obj.FALSE:
		return obj.TRUE
	case obj.NULL:
		return obj.FALSE
	default:
		intVal, ok := val.(*obj.IntegerObject)
		if ok {
			if intVal.Value == 0 {
				return obj.TRUE
			} else {
				return obj.FALSE
			}
		}
		floatVal, ok := val.(*obj.FloatObject)
		if ok {
			if floatVal.Value == 0.0 {
				return obj.TRUE
			} else {
				return obj.FALSE
			}
		}
		return &obj.ErrorObject{Value: "Not a valid type, for the operator !"}
	}
}

func evalInfixExpression(expr *ast.InfixExpression) obj.Object {
	rightVal := Eval(expr.RightExp)
	leftVal := Eval(expr.LeftExp)
	switch expr.Token.TokenType {
	case token.PLUS:
		return evalInfixPlusOp(leftVal, rightVal)
	case token.MINUS:
		return evalInfixMinusOp(leftVal, rightVal)
	default:
		return &obj.ErrorObject{Value: "Not a valid infix operator"}
	}
}

func evalInfixPlusOp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	if leftVal.Type() == obj.STRING_OBJ && rightVal.Type() == obj.STRING_OBJ {
		lval, _ := leftVal.(*obj.StringObject)
		rval, _ := rightVal.(*obj.StringObject)
		return &obj.StringObject{Value: lval.Value + rval.Value}
	}

	lNum, lIsNum := getNumericValue(leftVal)
	rNum, rIsNum := getNumericValue(rightVal)
	if lIsNum && rIsNum {
		if isFloat(leftVal) || isFloat(rightVal) {
			return &obj.FloatObject{Value: lNum + rNum}
		}
		return &obj.IntegerObject{Value: int64(lNum) + int64(rNum)}
	}

	return &obj.ErrorObject{Value: "Invalid types for the operator + "}
}

func evalInfixMinusOp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	lNum, lIsNum := getNumericValue(leftVal)
	rNum, rIsNum := getNumericValue(rightVal)
	if lIsNum && rIsNum {
		if isFloat(leftVal) || isFloat(rightVal) {
			return &obj.FloatObject{Value: lNum - rNum}
		}
		return &obj.IntegerObject{Value: int64(lNum) - int64(rNum)}
	}

	return &obj.ErrorObject{Value: "Invalid types for the operator - "}
}

func getNumericValue(val obj.Object) (float64, bool) {
	switch v := val.(type) {
	case *obj.IntegerObject:
		return float64(v.Value), true
	case *obj.FloatObject:
		return v.Value, true
	default:
		return 0, false
	}
}

func isFloat(val obj.Object) bool {
	_, ok := val.(*obj.FloatObject)
	return ok
}
