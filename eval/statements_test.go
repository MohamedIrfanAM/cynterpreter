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
		{"int x = 10;", "x += 5;", "x", int64(15)},
		{"int y = 20;", "y -= 8;", "y", int64(12)},
		{"int z = 6;", "z *= 4;", "z", int64(24)},
		{"int w = 15;", "w /= 3;", "w", int64(5)},
		{"float a = 2.5;", "a += 1.5;", "a", 4.0},
		{"float b = 10.0;", "b -= 3.5;", "b", 6.5},
		{"float c = 3.0;", "c *= 2.0;", "c", 6.0},
		{"float d = 8.0;", "d /= 2.0;", "d", 4.0},
		{"string str = \"Hello\";", "str += \" World\";", "str", "Hello World"},
		{"int neg = 5;", "neg -= 10;", "neg", int64(-5)},
		{"float zero = 5.0;", "zero *= 0.0;", "zero", 0.0},
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

func TestVariableUsage(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		// C-style variable declarations and usage
		{"int x = 5; x;", 5},
		{"float y = 3.14; y;", 3.14},
		{"bool flag = true; flag;", true},
		{"string name = \"John\"; name;", "John"},
		{"char ch = 'A'; ch;", byte('A')},

		// Variable operations
		{"int a = 10; int b = 20; a + b;", 30},
		{"int x = 5; int y = 2; x * y;", 10},
		{"float pi = 3.14; float radius = 2.0; pi * radius;", 6.28},
		{"string firstName = \"Hello\"; string lastName = \" World\"; firstName + lastName;", "Hello World"},
		{"bool isActive = true; bool isValid = false; isActive && isValid;", false},
		{"int score1 = 85; int score2 = 92; score1 > score2;", false},
		{"float temp = 25.5; float threshold = 30.0; temp < threshold;", true},

		// Variable reassignment
		{"int count = 5; count = count + 1; count;", 6},
		{"int count; count = 10; count;", 10},
		{"int value = 10; value = value * 2; value;", 20},
		{"string message = \"Hello\"; message = message + \"!\"; message;", "Hello!"},

		// Mixed type operations
		{"int intVal = 5; float floatVal = 2.5; intVal + floatVal;", 7.5},
		{"bool result = false; int num = 0; result || num;", false},
		{"char letter = 'Z'; string word = \"oo\"; letter;", byte('Z')},
	}

	for i, tt := range tests {
		env := obj.NewEnv()
		p := parser.New(tt.input)
		program := p.ParseProgram()

		var result obj.Object
		for _, stmt := range program.Statements {
			result = Eval(stmt, env)
		}

		switch val := tt.expected.(type) {
		case int:
			testIntegerObject(t, result, val)
		case float64:
			testFloatObject(t, result, val)
		case bool:
			testBooleanObject(t, result, val)
		case string:
			testStringObject(t, result, val)
		case byte:
			testCharObject(t, result, val)
		default:
			t.Fatalf("Test [%d]: Unsupported expected type %T", i, val)
		}
	}
}

func TestFunctionDeclarationStatement(t *testing.T) {
	tests := []struct {
		input      string
		varName    string
		paramCount int
		blockLen   int
	}{
		{
			input: `int add(int a, int b){
				return a+b;
			}`,
			varName:    "add",
			paramCount: 2,
			blockLen:   1,
		},
		{
			input: `float multiply(float x, float y, float z){
				float result = x * y * z;
				return result;
			}`,
			varName:    "multiply",
			paramCount: 3,
			blockLen:   2,
		},
		{
			input: `bool isEmpty(){
				return true;
			}`,
			varName:    "isEmpty",
			paramCount: 0,
			blockLen:   1,
		},
		{
			input: `string greet(string name){
				string message = "Hello, ";
				message = message + name;
				return message;
			}`,
			varName:    "greet",
			paramCount: 1,
			blockLen:   3,
		},
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

		storedObj, exists := env.GetVar(tt.varName)
		if !exists {
			t.Fatalf("[%d] - Function %s not found in environment", i, tt.varName)
		}

		functionObj, ok := storedObj.(*obj.FunctionObject)
		if !ok {
			t.Fatalf("[%d] - Expected FunctionObject, got %T", i, storedObj)
		}

		if len(functionObj.Params) != tt.paramCount {
			t.Errorf("[%d] - Expected %d parameters, got %d", i, tt.paramCount, len(functionObj.Params))
		}

		if len(functionObj.Block.Statements) != tt.blockLen {
			t.Errorf("[%d] - Expected %d statements in block, got %d", i, tt.blockLen, len(functionObj.Block.Statements))
		}

		// Verify return type matches declaration type
		expectedReturnType := obj.GetObjectType(stmnt.Type)
		if functionObj.ReturnType != expectedReturnType {
			t.Errorf("[%d] - Expected return type %s, got %s", i, expectedReturnType, functionObj.ReturnType)
		}
	}

	// Test function redeclaration error
	env := obj.NewEnv()
	input := `int test(){return 1;} int test(){return 2;}`
	p := parser.New(input)
	program := p.ParseProgram()

	result1 := Eval(program.Statements[0], env)
	if result1.Type() == obj.ERROR_OBJ {
		t.Fatalf("First function declaration should succeed: %s", result1.String())
	}

	result2 := Eval(program.Statements[1], env)
	if result2.Type() != obj.ERROR_OBJ {
		t.Fatalf("Expected redeclaration error for function, got %T", result2)
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 1;
	return 2+2;
	return 3*7;
	`
	expected := []int{1, 4, 21}

	env := obj.NewEnv()
	p := parser.New(input)
	program := p.ParseProgram()

	if len(program.Statements) != len(expected) {
		t.Fatalf("Expected %d statements, got %d", len(expected), len(program.Statements))
	}

	for i, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("[%d] - Expected *ast.ReturnStatement, got %T", i, stmt)
		}

		result := Eval(returnStmt, env)

		returnObj, ok := result.(*obj.ReturnObject)
		if !ok {
			t.Fatalf("[%d] - Expected *obj.ReturnObject, got %T", i, result)
		}

		testIntegerObject(t, returnObj.Return, expected[i])
	}
}

func TestFunctionCall(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			input: `int add(int a, int b){
				return a+b;
			}
			add(10,30);`,
			expected: 40,
		},
		{
			input: `float multiply(float x, float y){
				return x * y;
			}
			multiply(2.5, 4.0);`,
			expected: 10.0,
		},
		{
			input: `bool isEmpty(){
				return true;
			}
			isEmpty();`,
			expected: true,
		},
		{
			input: `string greet(string name){
				string message = "Hello, ";
				message = message + name;
				return message;
			}
			greet("World");`,
			expected: "Hello, World",
		},
		{
			input: `int factorial(int n){
				if(n <= 1){
					return 1;
				}
				return n * factorial(n - 1);
			}
			factorial(5);`,
			expected: 120,
		},
	}

	for i, tt := range tests {
		env := obj.NewEnv()
		p := parser.New(tt.input)
		program := p.ParseProgram()

		var result obj.Object
		for _, stmt := range program.Statements {
			result = Eval(stmt, env)
			if result.Type() == obj.ERROR_OBJ {
				t.Fatalf("[%d] - Evaluation error: %s", i, result.String())
			}
		}

		if returnObj, ok := result.(*obj.ReturnObject); ok {
			result = returnObj.Return
		}

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, result, expected)
		case float64:
			testFloatObject(t, result, expected)
		case bool:
			testBooleanObject(t, result, expected)
		case string:
			testStringObject(t, result, expected)
		default:
			t.Fatalf("[%d] - Unsupported expected type %T", i, expected)
		}
	}

	env := obj.NewEnv()
	input := `int add(int a, int b){return a+b;} add(10);`
	p := parser.New(input)
	program := p.ParseProgram()

	Eval(program.Statements[0], env)           // Declare function
	result := Eval(program.Statements[1], env) // Call with wrong args
	if result.Type() != obj.ERROR_OBJ {
		t.Fatalf("Expected error for wrong number of arguments, got %T", result)
	}

	env = obj.NewEnv()
	input = `undefinedFunction();`
	p = parser.New(input)
	program = p.ParseProgram()
	result = Eval(program.Statements[0], env)
	if result.Type() != obj.ERROR_OBJ {
		t.Fatalf("Expected error for undefined function call, got %T", result)
	}
}

func TestWhileStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{"int i = 0; while(i < 3){i = i + 1; i;}", []int{1, 2, 3}},
		{"int count = 5; while(count > 0){count; count = count - 1;}", []int{5, 4, 3, 2, 1}},
		{"bool flag = true; int x = 1; while(flag){x; flag = false;}", []int{1}},
		{"bool flag = false; while(flag){42;}", []int{}},
		{"int i = 0; while(i < 2){i = i + 1; i * 10; i;}", []int{10, 1, 20, 2}},
		{"int val = 10; while(val >= 8){val; val = val - 1;}", []int{10, 9, 8}},
		{"int num = 1; while(num <= 3){num * 5; num = num + 1;}", []int{5, 10, 15}},
		{"int x = 0; while(x != 2){x = x + 1; x;}", []int{1, 2}},
		{"int i = 0; while(i < 2){i = i + 1;}", []int{}},
		{"int once = 0; while(once == 0){999; once = 1;}", []int{999}},
		{"while(false){123;}", []int{}},
		{"int x = 5; while(x < 0){x;}", []int{}},
		{"int a = 2; int b = 3; while(a * b < 10){a * b; a = a + 1;}", []int{6, 9}},
	}

	for i, tt := range tests {
		env := obj.NewEnv()
		p := parser.New(tt.input)
		program := p.ParseProgram()

		var result obj.Object
		for _, stmt := range program.Statements {
			result = Eval(stmt, env)
			if result.Type() == obj.ERROR_OBJ {
				t.Fatalf("[%d] - Evaluation error: %s", i, result.String())
			}
		}

		var whileResult obj.Object
		if result != nil && result.Type() == obj.RESULTS_OBJ {
			whileResult = result
		} else {
			for j := len(program.Statements) - 1; j >= 0; j-- {
				if _, ok := program.Statements[j].(*ast.WhileStatement); ok {
					whileResult = Eval(program.Statements[j], env)
					break
				}
			}
		}

		if len(tt.expected) == 0 {
			if whileResult == nil || whileResult.Type() == obj.NULL_OBJ {
				continue
			}
			if resultsObj, ok := whileResult.(*obj.ResultsObject); ok {
				var nonNullResults []obj.Object
				for _, object := range resultsObj.Results {
					if object.Type() != obj.NULL_OBJ {
						nonNullResults = append(nonNullResults, object)
					}
				}
				if len(nonNullResults) == 0 {
					continue
				}
			}
			t.Fatalf("[%d] - Expected no results, but got: %v", i, whileResult)
		}

		resultsObj, ok := whileResult.(*obj.ResultsObject)
		if !ok {
			t.Fatalf("[%d] - Expected *obj.ResultsObject, got %T", i, whileResult)
		}

		var nonNullResults []obj.Object
		for _, object := range resultsObj.Results {
			if object.Type() != obj.NULL_OBJ {
				nonNullResults = append(nonNullResults, object)
			}
		}

		if len(nonNullResults) != len(tt.expected) {
			t.Fatalf("[%d] - Expected %d results, got %d", i, len(tt.expected), len(nonNullResults))
		}

		for j, object := range nonNullResults {
			testIntegerObject(t, object, tt.expected[j])
		}
	}
}

func TestForStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{"for(int i = 0; i < 3; i = i + 1){i;}", []int{0, 1, 2}},
		{"for(int count = 5; count > 2; count = count - 1){count;}", []int{5, 4, 3}},
		{"for(int x = 1; x <= 3; x = x + 1){x * 10;}", []int{10, 20, 30}},
		{"for(bool flag = true; flag; flag = false){42;}", []int{42}},
		{"for(int val = 8; val <= 10; val = val + 1){val;}", []int{8, 9, 10}},
		{"for(int num = 1; num <= 3; num = num + 1){num * 2; num;}", []int{2, 1, 4, 2, 6, 3}},
		{"for(int x = 0; x != 2; x = x + 1){x + 1;}", []int{1, 2}},
		{"for(int i = 0; i < 2; i = i + 1){}", []int{}},
		{"for(int once = 0; once == 0; once = 1){999;}", []int{999}},
		{"for(bool done = false; done; done = true){123;}", []int{}},
		{"for(int x = 5; x < 0; x = x + 1){x;}", []int{}},
	}

	for i, tt := range tests {
		env := obj.NewEnv()
		p := parser.New(tt.input)
		program := p.ParseProgram()

		var result obj.Object
		for _, stmt := range program.Statements {
			result = Eval(stmt, env)
			if result.Type() == obj.ERROR_OBJ {
				t.Fatalf("[%d] - Evaluation error: %s", i, result.String())
			}
		}

		var forResult obj.Object
		if result != nil && result.Type() == obj.RESULTS_OBJ {
			forResult = result
		} else {
			for j := len(program.Statements) - 1; j >= 0; j-- {
				if _, ok := program.Statements[j].(*ast.ForStatement); ok {
					forResult = Eval(program.Statements[j], env)
					break
				}
			}
		}

		if len(tt.expected) == 0 {
			if forResult == nil || forResult.Type() == obj.NULL_OBJ {
				continue
			}
			if resultsObj, ok := forResult.(*obj.ResultsObject); ok {
				var nonNullResults []obj.Object
				for _, object := range resultsObj.Results {
					if object.Type() != obj.NULL_OBJ {
						nonNullResults = append(nonNullResults, object)
					}
				}
				if len(nonNullResults) == 0 {
					continue
				}
			}
			t.Fatalf("[%d] - Expected no results, but got: %v", i, forResult)
		}

		resultsObj, ok := forResult.(*obj.ResultsObject)
		if !ok {
			t.Fatalf("[%d] - Expected *obj.ResultsObject, got %T", i, forResult)
		}

		var nonNullResults []obj.Object
		for _, object := range resultsObj.Results {
			if object.Type() != obj.NULL_OBJ {
				nonNullResults = append(nonNullResults, object)
			}
		}

		if len(nonNullResults) != len(tt.expected) {
			t.Fatalf("[%d] - Expected %d results, got %d", i, len(tt.expected), len(nonNullResults))
		}

		for j, object := range nonNullResults {
			testIntegerObject(t, object, tt.expected[j])
		}
	}
}
