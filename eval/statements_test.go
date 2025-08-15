package eval

import (
	"testing"

	"github.com/mohamedirfanam/cynterpreter/eval/obj"
	"github.com/mohamedirfanam/cynterpreter/parser"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

func TestIfStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{"if(true){10;20;}", []int{10, 20}},
		{"if(false){10;20;}", []int{}},
		{"if(false){10;20;}else{30;40;}", []int{30, 40}},
		// Numeric conditions
		{"if(1){42;}", []int{42}},
		{"if(0){10;}else{20;}", []int{20}},
		{"if(-1){15;}", []int{15}},
		// Comparison conditions
		{"if(5 > 3){100;}", []int{100}},
		{"if(3 > 5){10;}else{200;}", []int{200}},
		{"if(5 == 5){50;60;}", []int{50, 60}},
		{"if(5 != 3){70;}", []int{70}},
		{"if(10 >= 10){80;}", []int{80}},
		{"if(5 <= 3){90;}else{110;}", []int{110}},
		// Multiple statements in blocks
		{"if(true){1;2;3;4;5;}", []int{1, 2, 3, 4, 5}},
		{"if(false){1;2;}else{6;7;8;9;}", []int{6, 7, 8, 9}},
		// Empty blocks
		{"if(true){}", []int{}},
		{"if(false){}else{}", []int{}},
		{"if(false){10;}else{}", []int{}},
		// Single statement
		{"if(true){999;}", []int{999}},
		{"if(false){888;}else{777;}", []int{777}},
	}
	env := obj.NewEnv()
	for i, tt := range tests {
		p := parser.New(tt.input)
		program := p.ParseProgram()

		if len(program.Statements) != 1 {
			t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
		}

		stmnt, ok := program.Statements[0].(*ast.IfStatement)
		if !ok {
			t.Fatalf("[%d] - Not valid statement, expected *ast.IfStatement got %T", i, stmnt)
		}

		results := Eval(stmnt, env)

		if len(tt.expected) == 0 {
			if results == nil || results.Type() == obj.NULL_OBJ {
				continue
			}
			if resultsObj, ok := results.(*obj.ResultsObject); ok && len(resultsObj.Results) == 0 {
				continue
			}
			t.Fatalf("[%d] - Expected no results, but got: %v", i, results)
		}

		objects := results.(*obj.ResultsObject).Results
		if len(objects) != len(tt.expected) {
			t.Fatalf("[%d] - Expected %d results, got %d", i, len(tt.expected), len(objects))
		}

		for j, object := range objects {
			testIntegerObject(t, object, tt.expected[j])
		}
	}
}

func TestDeclarationStatement(t *testing.T) {
	tests := []struct {
		input       string
		identifier  string
		expectedVal interface{}
	}{
		{"int x = 10;", "x", int64(10)},
		{"char l = 'a';", "l", byte('a')},
		{"float pi = 3.14;", "pi", 3.14},
		{"bool flag = true;", "flag", true},
		{"string name = \"hello\";", "name", "hello"},
		{"int zero = 0;", "zero", int64(0)},
		{"bool falseBool = false;", "falseBool", false},
		{"float negative = -2.5;", "negative", -2.5},
	}

	for i, tt := range tests {
		env := obj.NewEnv()
		p := parser.New(tt.input)
		program := p.ParseProgram()

		if len(program.Statements) != 1 {
			t.Fatalf("[%d] Expected 1 statement, got %d", i, len(program.Statements))
		}

		stmnt, ok := program.Statements[0].(*ast.DeclarationStatement)
		if !ok {
			t.Fatalf("[%d] - Not valid statement, expected *ast.DeclarationStatement got %T", i, stmnt)
		}

		result := Eval(stmnt, env)
		if result.Type() == obj.ERROR_OBJ {
			t.Fatalf("[%d] - Evaluation error: %s", i, result.String())
		}

		storedObj, exists := env.GetVar(tt.identifier)
		if !exists {
			t.Fatalf("[%d] - Variable %s not found in environment", i, tt.identifier)
		}

		switch expected := tt.expectedVal.(type) {
		case int64:
			intObj, ok := storedObj.(*obj.IntegerObject)
			if !ok {
				t.Fatalf("[%d] - Expected IntegerObject, got %T", i, storedObj)
			}
			if intObj.Value != expected {
				t.Errorf("[%d] - Expected %d, got %d", i, expected, intObj.Value)
			}
		case byte:
			charObj, ok := storedObj.(*obj.CharObject)
			if !ok {
				t.Fatalf("[%d] - Expected CharObject, got %T", i, storedObj)
			}
			if charObj.Value != expected {
				t.Errorf("[%d] - Expected %c, got %c", i, expected, charObj.Value)
			}
		case float64:
			floatObj, ok := storedObj.(*obj.FloatObject)
			if !ok {
				t.Fatalf("[%d] - Expected FloatObject, got %T", i, storedObj)
			}
			if floatObj.Value != expected {
				t.Errorf("[%d] - Expected %f, got %f", i, expected, floatObj.Value)
			}
		case bool:
			boolObj, ok := storedObj.(*obj.BooleanObject)
			if !ok {
				t.Fatalf("[%d] - Expected BooleanObject, got %T", i, storedObj)
			}
			if boolObj.Value != expected {
				t.Errorf("[%d] - Expected %t, got %t", i, expected, boolObj.Value)
			}
		case string:
			stringObj, ok := storedObj.(*obj.StringObject)
			if !ok {
				t.Fatalf("[%d] - Expected StringObject, got %T", i, storedObj)
			}
			if stringObj.Value != expected {
				t.Errorf("[%d] - Expected %s, got %s", i, expected, stringObj.Value)
			}
		}
	}

	env := obj.NewEnv()
	input := "int x = 10; int x = 20;"
	p := parser.New(input)
	program := p.ParseProgram()

	result1 := Eval(program.Statements[0], env)
	if result1.Type() == obj.ERROR_OBJ {
		t.Fatalf("First declaration should succeed: %s", result1.String())
	}

	result2 := Eval(program.Statements[1], env)
	if result2.Type() != obj.ERROR_OBJ {
		t.Fatalf("Expected redeclaration error, got %T", result2)
	}
}

func TestAssignmentStatement(t *testing.T) {
	tests := []struct {
		declaration string
		assignment  string
		identifier  string
		expectedVal interface{}
	}{
		{"int x = 5;", "x = 10;", "x", int64(10)},
		{"char c = 'a';", "c = 'z';", "c", byte('z')},
		{"float f = 1.5;", "f = 2.7;", "f", 2.7},
		{"bool b = false;", "b = true;", "b", true},
		{"string s = \"old\";", "s = \"new\";", "s", "new"},
		{"int zero = 100;", "zero = 0;", "zero", int64(0)},
		{"bool flag = true;", "flag = false;", "flag", false},
		{"float negative = 1.0;", "negative = -3.14;", "negative", -3.14},
	}

	for i, tt := range tests {
		env := obj.NewEnv()

		// First declare the variable
		declP := parser.New(tt.declaration)
		declProgram := declP.ParseProgram()

		if len(declProgram.Statements) != 1 {
			t.Fatalf("[%d] Expected 1 declaration statement, got %d", i, len(declProgram.Statements))
		}

		declResult := Eval(declProgram.Statements[0], env)
		if declResult.Type() == obj.ERROR_OBJ {
			t.Fatalf("[%d] - Declaration error: %s", i, declResult.String())
		}

		// Then perform the assignment
		assignP := parser.New(tt.assignment)
		assignProgram := assignP.ParseProgram()

		if len(assignProgram.Statements) != 1 {
			t.Fatalf("[%d] Expected 1 assignment statement, got %d", i, len(assignProgram.Statements))
		}

		stmnt, ok := assignProgram.Statements[0].(*ast.AssignmentStatement)
		if !ok {
			t.Fatalf("[%d] - Not valid statement, expected *ast.AssignmentStatement got %T", i, stmnt)
		}

		result := Eval(stmnt, env)
		if result.Type() == obj.ERROR_OBJ {
			t.Fatalf("[%d] - Assignment error: %s", i, result.String())
		}

		storedObj, exists := env.GetVar(tt.identifier)
		if !exists {
			t.Fatalf("[%d] - Variable %s not found in environment", i, tt.identifier)
		}

		switch expected := tt.expectedVal.(type) {
		case int64:
			intObj, ok := storedObj.(*obj.IntegerObject)
			if !ok {
				t.Fatalf("[%d] - Expected IntegerObject, got %T", i, storedObj)
			}
			if intObj.Value != expected {
				t.Errorf("[%d] - Expected %d, got %d", i, expected, intObj.Value)
			}
		case byte:
			charObj, ok := storedObj.(*obj.CharObject)
			if !ok {
				t.Fatalf("[%d] - Expected CharObject, got %T", i, storedObj)
			}
			if charObj.Value != expected {
				t.Errorf("[%d] - Expected %c, got %c", i, expected, charObj.Value)
			}
		case float64:
			floatObj, ok := storedObj.(*obj.FloatObject)
			if !ok {
				t.Fatalf("[%d] - Expected FloatObject, got %T", i, storedObj)
			}
			if floatObj.Value != expected {
				t.Errorf("[%d] - Expected %f, got %f", i, expected, floatObj.Value)
			}
		case bool:
			boolObj, ok := storedObj.(*obj.BooleanObject)
			if !ok {
				t.Fatalf("[%d] - Expected BooleanObject, got %T", i, storedObj)
			}
			if boolObj.Value != expected {
				t.Errorf("[%d] - Expected %t, got %t", i, expected, boolObj.Value)
			}
		case string:
			stringObj, ok := storedObj.(*obj.StringObject)
			if !ok {
				t.Fatalf("[%d] - Expected StringObject, got %T", i, storedObj)
			}
			if stringObj.Value != expected {
				t.Errorf("[%d] - Expected %s, got %s", i, expected, stringObj.Value)
			}
		}
	}

	env := obj.NewEnv()
	p := parser.New("x = 10;")
	program := p.ParseProgram()

	result := Eval(program.Statements[0], env)
	if result.Type() != obj.ERROR_OBJ {
		t.Fatalf("Expected error for undeclared variable assignment, got %T", result)
	}

	env = obj.NewEnv()
	declP := parser.New("int x = 5;")
	declProgram := declP.ParseProgram()
	Eval(declProgram.Statements[0], env)

	assignP := parser.New("x = \"string\";")
	assignProgram := assignP.ParseProgram()
	result = Eval(assignProgram.Statements[0], env)
	if result.Type() != obj.ERROR_OBJ {
		t.Fatalf("Expected type mismatch error, got %T", result)
	}
}
