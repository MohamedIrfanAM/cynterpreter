package eval

import (
	"fmt"

	"github.com/mohamedirfanam/cynterpreter/eval/obj"
	"github.com/mohamedirfanam/cynterpreter/lexer/token"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

func evalIdentifierExpression(ident *ast.IdentifierExpression, env *obj.Environment) obj.Object {
	val, ok := env.GetVar(ident.Value)
	if !ok {
		return obj.NewError(fmt.Errorf("variable error: variable %s not declared in this scope", ident))
	}
	return val
}

func evalPrefixExpression(expr *ast.PrefixExpression, env *obj.Environment) obj.Object {
	val := Eval(expr.Exp, env)
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

func evalInfixExpression(expr *ast.InfixExpression, env *obj.Environment) obj.Object {
	rightVal := Eval(expr.RightExp, env)
	leftVal := Eval(expr.LeftExp, env)
	switch expr.Token.TokenType {
	case token.PLUS:
		return evalInfixPlusOp(leftVal, rightVal)
	case token.MINUS:
		return evalInfixMinusOp(leftVal, rightVal)
	case token.ASTER:
		return evalInfixMultOp(leftVal, rightVal)
	case token.SLASH:
		return evalInfixDevideOp(leftVal, rightVal)
	case token.PERCENT:
		return evalInfixModOp(leftVal, rightVal)
	case token.GT:
		return evalInfixGTOp(leftVal, rightVal)
	case token.LT:
		return evalInfixLTOp(leftVal, rightVal)
	case token.LE:
		return evalInfixLEOp(leftVal, rightVal)
	case token.GE:
		return evalInfixGEOp(leftVal, rightVal)
	case token.EQ:
		return evalInfixEQOp(leftVal, rightVal)
	case token.NE:
		return evalInfixNEOp(leftVal, rightVal)
	case token.AND:
		return evalInfixANDOp(leftVal, rightVal)
	case token.OR:
		return evalInfixOROp(leftVal, rightVal)
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

func evalInfixModOp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	rVal, ok1 := rightVal.(*obj.IntegerObject)
	lVal, ok2 := leftVal.(*obj.IntegerObject)

	if !ok1 || !ok2 {
		return obj.NewError(fmt.Errorf("type error: Invalid operand types for devide operator, expected number %% number but got %s %% %s", leftVal.Type(), rightVal.Type()))
	}
	if rVal.Value == 0 {
		return obj.NewError(fmt.Errorf("runtime error: devide by zero "))
	}
	return &obj.IntegerObject{Value: int64(lVal.Value) % int64(rVal.Value)}

}

func evalInfixEQOp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	if leftVal.Type() == obj.STRING_OBJ && rightVal.Type() == obj.STRING_OBJ {
		lval, _ := leftVal.(*obj.StringObject)
		rval, _ := rightVal.(*obj.StringObject)
		return obj.GetBoolean(lval.Value == rval.Value)
	} else if leftVal.Type() == obj.CHAR_OBJ && rightVal.Type() == obj.CHAR_OBJ {
		lval, _ := leftVal.(*obj.CharObject)
		rval, _ := rightVal.(*obj.CharObject)
		return obj.GetBoolean(lval.Value == rval.Value)
	} else if leftVal.Type() == obj.BOOLEAN_OBJ && rightVal.Type() == obj.BOOLEAN_OBJ {
		lval, _ := leftVal.(*obj.BooleanObject)
		rval, _ := rightVal.(*obj.BooleanObject)
		return obj.GetBoolean(lval.Value == rval.Value)
	}

	lNum, lIsNum := getNumericValue(leftVal)
	rNum, rIsNum := getNumericValue(rightVal)
	if lIsNum && rIsNum {
		if isFloat(leftVal) || isFloat(rightVal) {
			return obj.GetBoolean(lNum == rNum)
		}
		return obj.GetBoolean(int64(lNum) == int64(rNum))
	}

	return obj.NewError(fmt.Errorf("type error: Invalid operand types for equal operator, expected number==number or string==string or char == char but got %s + %s", leftVal.Type(), rightVal.Type()))
}

func evalInfixNEOp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	if leftVal.Type() == obj.STRING_OBJ && rightVal.Type() == obj.STRING_OBJ {
		lval, _ := leftVal.(*obj.StringObject)
		rval, _ := rightVal.(*obj.StringObject)
		return obj.GetBoolean(lval.Value != rval.Value)
	} else if leftVal.Type() == obj.CHAR_OBJ && rightVal.Type() == obj.CHAR_OBJ {
		lval, _ := leftVal.(*obj.CharObject)
		rval, _ := rightVal.(*obj.CharObject)
		return obj.GetBoolean(lval.Value != rval.Value)
	} else if leftVal.Type() == obj.BOOLEAN_OBJ && rightVal.Type() == obj.BOOLEAN_OBJ {
		lval, _ := leftVal.(*obj.BooleanObject)
		rval, _ := rightVal.(*obj.BooleanObject)
		return obj.GetBoolean(lval.Value != rval.Value)
	}

	lNum, lIsNum := getNumericValue(leftVal)
	rNum, rIsNum := getNumericValue(rightVal)
	if lIsNum && rIsNum {
		if isFloat(leftVal) || isFloat(rightVal) {
			return obj.GetBoolean(lNum != rNum)
		}
		return obj.GetBoolean(int64(lNum) != int64(rNum))
	}

	return obj.NewError(fmt.Errorf("type error: Invalid operand types for equal operator, expected number!=number or string==string or char != char but got %s + %s", leftVal.Type(), rightVal.Type()))
}

func evalInfixGTOp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	lNum, lIsNum := getNumericValue(leftVal)
	rNum, rIsNum := getNumericValue(rightVal)
	if lIsNum && rIsNum {
		if isFloat(leftVal) || isFloat(rightVal) {
			return obj.GetBoolean(lNum > rNum)
		}
		return obj.GetBoolean(int64(lNum) > int64(rNum))
	}

	return obj.NewError(fmt.Errorf("type error: Invalid operand types for Greater Than operator, expected number>number but got %s > %s", leftVal.Type(), rightVal.Type()))
}

func evalInfixLTOp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	lNum, lIsNum := getNumericValue(leftVal)
	rNum, rIsNum := getNumericValue(rightVal)
	if lIsNum && rIsNum {
		if isFloat(leftVal) || isFloat(rightVal) {
			return obj.GetBoolean(lNum < rNum)
		}
		return obj.GetBoolean(int64(lNum) < int64(rNum))
	}

	return obj.NewError(fmt.Errorf("type error: Invalid operand types for Less Than operator, expected number<number but got %s < %s", leftVal.Type(), rightVal.Type()))
}

func evalInfixLEOp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	lNum, lIsNum := getNumericValue(leftVal)
	rNum, rIsNum := getNumericValue(rightVal)
	if lIsNum && rIsNum {
		if isFloat(leftVal) || isFloat(rightVal) {
			return obj.GetBoolean(lNum <= rNum)
		}
		return obj.GetBoolean(int64(lNum) <= int64(rNum))
	}

	return obj.NewError(fmt.Errorf("type error: Invalid operand types for Less Than or Equal operator, expected number<=number but got %s <= %s", leftVal.Type(), rightVal.Type()))
}

func evalInfixGEOp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	lNum, lIsNum := getNumericValue(leftVal)
	rNum, rIsNum := getNumericValue(rightVal)
	if lIsNum && rIsNum {
		if isFloat(leftVal) || isFloat(rightVal) {
			return obj.GetBoolean(lNum >= rNum)
		}
		return obj.GetBoolean(int64(lNum) >= int64(rNum))
	}

	return obj.NewError(fmt.Errorf("type error: Invalid operand types for Greater Than or Equal operator, expected number>=number but got %s >= %s", leftVal.Type(), rightVal.Type()))
}

func evalInfixANDOp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	return obj.GetBoolean(IsTrue(leftVal) && IsTrue(rightVal))
}

func evalInfixOROp(leftVal obj.Object, rightVal obj.Object) obj.Object {
	return obj.GetBoolean(IsTrue(leftVal) || IsTrue(rightVal))
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

func IsTrue(val obj.Object) bool {
	switch val := val.(type) {
	case *obj.IntegerObject:
		if val.Value != 0 {
			return true
		} else {
			return false
		}
	case *obj.BooleanObject:
		return val.Value
	case *obj.FloatObject:
		if val.Value != 0.0 {
			return true
		} else {
			return false
		}
	case *obj.StringObject:
		if val.Value != "" {
			return true
		} else {
			return false
		}
	case *obj.NullObject:
		return false
	case *obj.CharObject:
		if val.Value != 0 {
			return true
		} else {
			return false
		}
	default:
		return true
	}
}

func evalCallExpression(ce *ast.CallExpression, env *obj.Environment) obj.Object {
	object, ok := env.GetVar(ce.Function.String())
	if !ok {
		return obj.NewError(fmt.Errorf("error calling function %s, function not found", ce.Function.String()))
	}
	funcObj, ok := object.(*obj.FunctionObject)
	if !ok {
		return obj.NewError(fmt.Errorf("error calling function %s, identifier not associated with a function", ce.Function.String()))
	}
	if len(funcObj.Params) != len(ce.Args) {
		return obj.NewError(fmt.Errorf("error calling function %s, number of args and parameter mismatch, Parameters - %d, Args - %d", ce.Function.String(), len(funcObj.Params), len(ce.Args)))
	}
	newEnv := env.ExtendEnv()
	// validate parameter argument pairs and assign args to params
	for i, param := range funcObj.Params {
		arg := Eval(ce.Args[i], env)
		if obj.GetObjectType(param.Type) != arg.Type() {
			return obj.NewError(fmt.Errorf("error calling function %s, type of parameter %s mismatch, expected %s, got %s", ce.Function, param.Identifier, param.Type, arg.Type()))
		}
		newEnv.SetVar(param.Identifier.Value, arg)
	}
	returnObj := evalBlock(funcObj.Block, newEnv)
	if funcObj.ReturnType == obj.NULL_OBJ {
		if returnObj.Type() == obj.RESULTS_OBJ {
			return obj.NULL
		} else if returnObj.Type() == obj.RETURN_OBJ {
			return obj.NewError(fmt.Errorf("invalid return from a void function %s, no return expected", ce.Function))
		}
	}
	returnVal, ok := returnObj.(*obj.ReturnObject)
	if !ok {
		return obj.NewError(fmt.Errorf("error calling function %s, expected return value of type %s, got none", ce.Function, funcObj.ReturnType))
	}
	if returnVal.Return.Type() != funcObj.ReturnType {
		return obj.NewError(fmt.Errorf("error calling function %s, return value type mismatch, expected %s, got %s", ce.Function, funcObj.ReturnType, returnVal.Return.Type()))
	}
	return returnVal.Return
}
