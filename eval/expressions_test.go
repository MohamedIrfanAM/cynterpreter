package eval

import (
	"testing"

	"github.com/mohamedirfanam/cynterpreter/eval/obj"
	"github.com/mohamedirfanam/cynterpreter/parser"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

func TestLiteralExpression(t *testing.T) {
	input := `
	123;
	53.32;
	true;
	false;
	'a';
	"Hello world!";
	`
	expected := []any{123, 53.32, true, false, 'a', "Hello world!"}
	p := parser.New(input)
	program := p.ParseProgram()

	for i, statement := range program.Statements {
		stmnt, ok := statement.(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("[%d] - Not valid statement, expected *ast.ExpressionStatement got %T", i, stmnt)
		}
		object := Eval(statement)
		switch val := expected[i].(type) {
		case int:
			testIntegerObject(t, object, val)
		case float64:
			testFloatObject(t, object, val)
		case bool:
			testBooleanObject(t, object, val)
		case rune:
			testCharObject(t, object, byte(val))
		case string:
			testStringObject(t, object, val)
		}
	}
}

func testIntegerObject(t *testing.T, object obj.Object, value int) {
	intObject, ok := object.(*obj.IntegerObject)
	if !ok {
		t.Errorf("Object not valid, expected integer object got %T", object)
		return
	}

	if object.Type() != obj.INTEGER_OBJ {
		t.Errorf("Object type not correct, expected INTEGER_OBJ, got %s", object.Type())
	}

	if intObject.Value != int64(value) {
		t.Errorf("Value Mismatch, expected %d, got %d", value, intObject.Value)
	}
}

func testFloatObject(t *testing.T, object obj.Object, value float64) {
	floatObject, ok := object.(*obj.FloatObject)
	if !ok {
		t.Errorf("Object not valid, expected float object got %T", object)
		return
	}

	if object.Type() != obj.FLOAT_OBJ {
		t.Errorf("Object type not correct, expected FLOAT_OBJ, got %s", object.Type())
	}

	if floatObject.Value != value {
		t.Errorf("Value Mismatch, expected %f, got %f", value, floatObject.Value)
	}
}

func testBooleanObject(t *testing.T, object obj.Object, value bool) {
	boolObject, ok := object.(*obj.BooleanObject)
	if !ok {
		t.Errorf("Object not valid, expected boolean object got %T", object)
		return
	}

	if object.Type() != obj.BOOLEAN_OBJ {
		t.Errorf("Object type not correct, expected BOOLEAN_OBJ, got %s", object.Type())
	}

	if boolObject.Value != value {
		t.Errorf("Value Mismatch, expected %t, got %t", value, boolObject.Value)
	}
}

func testCharObject(t *testing.T, object obj.Object, value byte) {
	charObject, ok := object.(*obj.CharObject)
	if !ok {
		t.Errorf("Object not valid, expected char object got %T", object)
		return
	}

	if object.Type() != obj.CHAR_OBJ {
		t.Errorf("Object type not correct, expected CHAR_OBJ, got %s", object.Type())
	}

	if charObject.Value != value {
		t.Errorf("Value Mismatch, expected %c, got %c", value, charObject.Value)
	}
}

func testStringObject(t *testing.T, object obj.Object, value string) {
	stringObject, ok := object.(*obj.StringObject)
	if !ok {
		t.Errorf("Object not valid, expected string object got %T", object)
		return
	}

	if object.Type() != obj.STRING_OBJ {
		t.Errorf("Object type not correct, expected STRING_OBJ, got %s", object.Type())
	}

	if stringObject.Value != value {
		t.Errorf("Value Mismatch, expected %s, got %s", value, stringObject.Value)
	}
}

func TestPrefixNot(t *testing.T) {
	input := `
	!121;
	!true;
	!false;
	!!true;
	!!false;
	!0;
	!1;
	!0.0;
	!1.5;
	!(-2.3);
	`
	expected := []bool{false, false, true, true, false, true, false, true, false, false}
	p := parser.New(input)
	program := p.ParseProgram()

	for i, statement := range program.Statements {
		stmnt, ok := statement.(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("[%d] - Not valid statement, expected *ast.ExpressionStatement got %T", i, stmnt)
		}
		object := Eval(statement)
		testBooleanObject(t, object, expected[i])
	}
}

func TestPrefixMinus(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"-1;", -1},
		{"-0;", 0},
		{"-15;", -15},
		{"-(-5);", 5},
		{"-1.5;", -1.5},
		{"-0.0;", 0.0},
		{"-(-2.3);", 2.3},
		{"-123.456;", -123.456},
	}

	for i, tt := range tests {
		p := parser.New(tt.input)
		program := p.ParseProgram()

		if len(program.Statements) != 1 {
			t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
		}

		stmnt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("[%d] - Not valid statement, expected *ast.ExpressionStatement got %T", i, stmnt)
		}

		object := Eval(stmnt)

		switch val := tt.expected.(type) {
		case int:
			testIntegerObject(t, object, val)
		case float64:
			testFloatObject(t, object, val)
		}
	}
}

func TestInfixOps(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		// Addition tests
		{"1 + 1;", 2},
		{"2 + 3;", 5},
		{"-1 + 1;", 0},
		{"1.5 + 2.5;", 4.0},
		{"1 + 2.5;", 3.5},
		{`"hello" + " world";`, "hello world"},

		// Subtraction tests
		{"5 - 3;", 2},
		{"0 - 1;", -1},
		{"3.5 - 1.5;", 2.0},
		{"5 - 2.5;", 2.5},
		{"10 - (-5);", 15},

		// Multiplication tests
		{"2 * 3;", 6},
		{"0 * 5;", 0},
		{"-2 * 3;", -6},
		{"2.5 * 4;", 10.0},
		{"3 * 2.5;", 7.5},

		// Division tests
		{"6 / 2;", 3},
		{"-6 / 2;", -3},
		{"7.5 / 2.5;", 3.0},
		{"10 / 2.5;", 4.0},

		// Comparison tests - Greater Than
		{"5 > 3;", true},
		{"3 > 5;", false},
		{"5 > 5;", false},
		{"3.5 > 2.1;", true},
		{"5 > 2.5;", true},

		// Comparison tests - Less Than
		{"3 < 5;", true},
		{"5 < 3;", false},
		{"5 < 5;", false},
		{"2.1 < 3.5;", true},

		// Comparison tests - Greater Than or Equal
		{"5 >= 3;", true},
		{"5 >= 5;", true},
		{"3 >= 5;", false},
		{"3.5 >= 2.1;", true},

		// Comparison tests - Less Than or Equal
		{"3 <= 5;", true},
		{"5 <= 5;", true},
		{"5 <= 3;", false},
		{"2.1 <= 3.5;", true},

		// Equality tests
		{"5 == 5;", true},
		{"5 == 3;", false},
		{"5.5 == 5.5;", true},
		{"5 == 5.0;", true},
		{"true == true;", true},
		{"true == false;", false},
		{`"hello" == "hello";`, true},
		{`"hello" == "world";`, false},
		{"'a' == 'a';", true},
		{"'a' == 'b';", false},

		// Inequality tests
		{"5 != 3;", true},
		{"5 != 5;", false},
		{"5.5 != 3.2;", true},
		{"5 != 5.0;", false},
		{"true != false;", true},
		{"true != true;", false},
		{`"hello" != "world";`, true},
		{`"hello" != "hello";`, false},

		// Complex expressions
		{"2 + 3 * 4;", 14},
		{"(2 + 3) * 4;", 20},
		{"10 / 2 - 8 / 4;", 3},
		{"(5 + 3) * (2 - 1);", 8},
		{"2.5 * 3.5 + 1.5;", 10.25},
		{"(3 + 1.5) * 2;", 9.0},
	}

	for i, tt := range tests {
		p := parser.New(tt.input)
		program := p.ParseProgram()

		if len(program.Statements) != 1 {
			t.Fatalf("Test [%d]: Expected 1 statement, got %d", i, len(program.Statements))
		}

		stmnt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Test [%d]: Not valid statement, expected *ast.ExpressionStatement got %T", i, stmnt)
		}

		object := Eval(stmnt)

		switch val := tt.expected.(type) {
		case int:
			testIntegerObject(t, object, val)
		case float64:
			testFloatObject(t, object, val)
		case bool:
			testBooleanObject(t, object, val)
		case string:
			testStringObject(t, object, val)
		default:
			t.Fatalf("Test [%d]: Unsupported expected type %T", i, val)
		}
	}
}
