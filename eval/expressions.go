package eval

import (
	"fmt"

	"github.com/mohamedirfanam/cynterpreter/eval/obj"
	"github.com/mohamedirfanam/cynterpreter/lexer/token"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

func evalPrefixExpression(expr *ast.PrefixExpression) obj.Object {
	val := Eval(expr.Exp)
	switch expr.Token.TokenType {
	case token.MINUS:
		return evalPrefixMinusOp(val)
	case token.NOT:
		return evalPrefixNotOp(val)
	default:
		return obj.NewError(fmt.Errorf("operator error: Not a valid operator, got %s", expr.Token.TokenType))
	}
}

func evalPrefixMinusOp(val obj.Object) obj.Object {
	switch val := val.(type) {
	case *obj.IntegerObject:
		return &obj.IntegerObject{Value: -1 * val.Value}
	case *obj.FloatObject:
		return &obj.FloatObject{Value: -1 * val.Value}
	default:
		return obj.NewError(fmt.Errorf("type error: Invalid operand type for unary minus operator, expected number but got %s", val.Type()))
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
		return obj.NewError(fmt.Errorf("type error: Invalid operand type for logical NOT operator, expected boolean or number but got %s", val.Type()))
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
	case token.ASTER:
		return evalInfixMultOp(leftVal, rightVal)
	case token.SLASH:
		return evalInfixDevideOp(leftVal, rightVal)
	case token.GT:
		return evalInfixGTOp(leftVal, rightVal)
	case token.LT:
		return evalInfixLTOp(leftVal, rightVal)
	case token.LE:
		return evalInfixLEOp(leftVal, rightVal)
	case token.GE:
		return evalInfixGEOp(leftVal, rightVal)
	default:
		return obj.NewError(fmt.Errorf("operator error: Unsupported infix operator '%s'", expr.Token.TokenType))
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

	return obj.NewError(fmt.Errorf("type error: Invalid operand types for addition operator, expected number+number or string+string but got %s + %s", leftVal.Type(), rightVal.Type()))
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

	return obj.NewError(fmt.Errorf("type error: Invalid operand types for subtraction operator, expected number-number but got %s - %s", leftVal.Type(), rightVal.Type()))
}

func evalInfixMultOp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	lNum, lIsNum := getNumericValue(leftVal)
	rNum, rIsNum := getNumericValue(rightVal)
	if lIsNum && rIsNum {
		if isFloat(leftVal) || isFloat(rightVal) {
			return &obj.FloatObject{Value: lNum * rNum}
		}
		return &obj.IntegerObject{Value: int64(lNum) * int64(rNum)}
	}

	return obj.NewError(fmt.Errorf("type error: Invalid operand types for product operator, expected number*number but got %s * %s", leftVal.Type(), rightVal.Type()))
}

func evalInfixDevideOp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	lNum, lIsNum := getNumericValue(leftVal)
	rNum, rIsNum := getNumericValue(rightVal)
	if rNum == 0 {
		return obj.NewError(fmt.Errorf("runtime error: devide by zero "))
	}
	if lIsNum && rIsNum {
		if isFloat(leftVal) || isFloat(rightVal) {
			return &obj.FloatObject{Value: lNum / rNum}
		}
		return &obj.IntegerObject{Value: int64(lNum) / int64(rNum)}
	}

	return obj.NewError(fmt.Errorf("type error: Invalid operand types for devide operator, expected number/number but got %s / %s", leftVal.Type(), rightVal.Type()))
}

func evalInfixGTOp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	lNum, lIsNum := getNumericValue(leftVal)
	rNum, rIsNum := getNumericValue(rightVal)
	if lIsNum && rIsNum {
		if isFloat(leftVal) || isFloat(rightVal) {
			return &obj.BooleanObject{Value: lNum > rNum}
		}
		return &obj.BooleanObject{Value: int64(lNum) > int64(rNum)}
	}

	return obj.NewError(fmt.Errorf("type error: Invalid operand types for Greater Than operator, expected number>number but got %s > %s", leftVal.Type(), rightVal.Type()))
}

func evalInfixLTOp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	lNum, lIsNum := getNumericValue(leftVal)
	rNum, rIsNum := getNumericValue(rightVal)
	if lIsNum && rIsNum {
		if isFloat(leftVal) || isFloat(rightVal) {
			return &obj.BooleanObject{Value: lNum < rNum}
		}
		return &obj.BooleanObject{Value: int64(lNum) < int64(rNum)}
	}

	return obj.NewError(fmt.Errorf("type error: Invalid operand types for Less Than operator, expected number<number but got %s < %s", leftVal.Type(), rightVal.Type()))
}

func evalInfixLEOp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	lNum, lIsNum := getNumericValue(leftVal)
	rNum, rIsNum := getNumericValue(rightVal)
	if lIsNum && rIsNum {
		if isFloat(leftVal) || isFloat(rightVal) {
			return &obj.BooleanObject{Value: lNum <= rNum}
		}
		return &obj.BooleanObject{Value: int64(lNum) <= int64(rNum)}
	}

	return obj.NewError(fmt.Errorf("type error: Invalid operand types for Less Than or Equal operator, expected number<=number but got %s <= %s", leftVal.Type(), rightVal.Type()))
}

func evalInfixGEOp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	lNum, lIsNum := getNumericValue(leftVal)
	rNum, rIsNum := getNumericValue(rightVal)
	if lIsNum && rIsNum {
		if isFloat(leftVal) || isFloat(rightVal) {
			return &obj.BooleanObject{Value: lNum >= rNum}
		}
		return &obj.BooleanObject{Value: int64(lNum) >= int64(rNum)}
	}

	return obj.NewError(fmt.Errorf("type error: Invalid operand types for Greater Than or Equal operator, expected number>=number but got %s >= %s", leftVal.Type(), rightVal.Type()))
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
