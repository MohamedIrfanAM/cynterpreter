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
		{"1 + 1;", 2},
		{"2 + 3;", 5},
		{"-1 + 1;", 0},
		{"0 + 100;", 100},
		{"1.5 + 2.5;", 4.0},
		{"1 + 2.5;", 3.5},
		{"2.5 + 1;", 3.5},
		{"0.0 + 0.0;", 0.0},
		{`"hello" + " world";`, "hello world"},
		{`"a" + "b";`, "ab"},
		{"1 - 1;", 0},
		{"5 - 3;", 2},
		{"0 - 1;", -1},
		{"100 - 50;", 50},
		{"3.5 - 1.5;", 2.0},
		{"5 - 2.5;", 2.5},
		{"2.5 - 1;", 1.5},
		{"0.0 - 0.0;", 0.0},
		{"-5 - 3;", -8},
		{"10 - (-5);", 15},
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
		case string:
			testStringObject(t, object, val)
		}
	}
}
