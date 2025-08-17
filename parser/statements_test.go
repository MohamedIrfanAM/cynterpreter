package parser

import (
	"testing"

	"github.com/mohamedirfanam/cynterpreter/lexer/token"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

func TestExpressionStatements(t *testing.T) {
	input := `
	a+b;
	x*y+z;
	(a+b)*c;
	5*(x+y);
	(a+b)*(c-d);
	((a+b)*c)-d;
	7+(x*(y+z));
	count+value;
	5*count+value;
	(count+value)*factor;
	5*(count+value);
	(count+value)*(factor-constant);
	((count+value)*factor)-constant;
	7+(count*(value+factor));
	-5;
	+10;
	-x+y;
	-(a+b);
	-5*3;
	3*-5;
	-a*b+c;
	3.14+2.5;
	'a'+'b';
	"hello"+"world";
	5.5*2.0;
	-3.14;
	+2.71;
	count+-5;
	-value*factor;
	"str"+variable;
	'x'*2;
	-(3.14+2.5);
	add();
	func(5);
	calc(x, y);
	print("hello");
	multiply(2, 3, 4);
	nested(func(x), y);
	add(5) + subtract(3);
	multiply(add(2, 3), 4);
	func(1, 2.5, "test");
	process(count + value);
	`
	expected := []string{
		"(a + b)",
		"((x * y) + z)",
		"((a + b) * c)",
		"(5 * (x + y))",
		"((a + b) * (c - d))",
		"(((a + b) * c) - d)",
		"(7 + (x * (y + z)))",
		"(count + value)",
		"((5 * count) + value)",
		"((count + value) * factor)",
		"(5 * (count + value))",
		"((count + value) * (factor - constant))",
		"(((count + value) * factor) - constant)",
		"(7 + (count * (value + factor)))",
		"(-5)",
		"(+10)",
		"((-x) + y)",
		"(-(a + b))",
		"((-5) * 3)",
		"(3 * (-5))",
		"(((-a) * b) + c)",
		"(3.14 + 2.5)",
		"('a' + 'b')",
		"(\"hello\" + \"world\")",
		"(5.5 * 2.0)",
		"(-3.14)",
		"(+2.71)",
		"(count + (-5))",
		"((-value) * factor)",
		"(\"str\" + variable)",
		"('x' * 2)",
		"(-(3.14 + 2.5))",
		"add()",
		"func(5)",
		"calc(x, y)",
		"print(\"hello\")",
		"multiply(2, 3, 4)",
		"nested(func(x), y)",
		"(add(5) + subtract(3))",
		"multiply(add(2, 3), 4)",
		"func(1, 2.5, \"test\")",
		"process((count + value))",
	}

	p := New(input)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			t.Errorf("Parser Error: %s\n", err.Error())
		}
		t.Fatal("Exiting now!")
	}

	if len(program.Statements) != len(expected) {
		t.Fatalf("Expected %d statement, got %d", len(expected), len(program.Statements))
	}

	for i, statement := range program.Statements {
		stmt, ok := statement.(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Statement is not of type ast.ExpressionStatement, got %T", statement)
		}

		var expr ast.Expression = stmt.Expression

		if expr.String() != expected[i] {
			t.Errorf("Expression mismatch at index %d, expected %s, got %s", i, expected[i], expr.String())
		}
	}
}

func TestDeclarationStatements(t *testing.T) {
	input := `
	int x = 10;
	char l = 'a';
	float pi = 3.14;
	bool flag = true;
	string name = "hello";
	int zero = 0;
	char newline = '\n';
	float negative = -2.5;
	bool falseBool = false;
	int sum = 10 + 5;
	int diff = a - b;
	float product = 3.14 * 2;
	int complex = (x + y) * z;
	bool comparison = a > b;
	int funcCall = add(5, 10);
	float nested = calculate(x + y, z);
	int x;
	char y;
	`
	expected := []struct {
		tokenType  token.TokenType
		identifier string
		literal    string
	}{
		{token.INT, "x", "10"},
		{token.CHAR, "l", "'a'"},
		{token.FLOAT, "pi", "3.14"},
		{token.BOOL, "flag", "true"},
		{token.STRING, "name", "\"hello\""},
		{token.INT, "zero", "0"},
		{token.CHAR, "newline", "'\\n'"},
		{token.FLOAT, "negative", "(-2.5)"},
		{token.BOOL, "falseBool", "false"},
		{token.INT, "sum", "(10 + 5)"},
		{token.INT, "diff", "(a - b)"},
		{token.FLOAT, "product", "(3.14 * 2)"},
		{token.INT, "complex", "((x + y) * z)"},
		{token.BOOL, "comparison", "(a > b)"},
		{token.INT, "funcCall", "add(5, 10)"},
		{token.FLOAT, "nested", "calculate((x + y), z)"},
		{token.INT, "x", ""},
		{token.CHAR, "y", ""},
	}

	p := New(input)
	statements := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			t.Errorf("Parser Error: %s\n", err.Error())
		}
		t.Fatal("Exiting now!")
	}

	if len(statements.Statements) != len(expected) {
		t.Fatalf("Number of statements not valid, expected %d, got %d", len(expected), len(statements.Statements))
	}

	for i, statement := range statements.Statements {

		stmnt, ok := statement.(*ast.DeclarationStatement)
		if !ok {
			t.Fatalf("[%d] - Not valid statement, expected ast.DeclarationStatement got %T", i, statement)
		}

		if stmnt.Type != expected[i].tokenType {
			t.Errorf("[%d] - Declaration type not valid, expected %s, got %s", i, expected[i].tokenType, stmnt.Type)
		}

		if stmnt.Identifier.String() != expected[i].identifier {
			t.Errorf("[%d] - Declaration Identifier name not correct, expected %s, got %s", i, expected[i].identifier, stmnt.Identifier.String())
		}

		if stmnt.Literal != nil && stmnt.Literal.String() != expected[i].literal {
			t.Errorf("[%d] - Declaration Literal not correct, expected %s, got %s", i, expected[i].literal, stmnt.Literal.String())
		}
	}
}

func TestAssignmentStatements(t *testing.T) {
	input := `
	x = 10;
	l = 'a';
	pi = 3.14;
	flag = true;
	name = "hello";
	zero = 0;
	newline = '\n';
	negative = -2.5;
	falseBool = false;
	sum = 10 + 5;
	diff = a - b;
	product = 3.14 * 2;
	complex = (x + y) * z;
	comparison = a > b;
	funcCall = add(5, 10);
	nested = calculate(x + y, z);
	x += 5;
	count -= 10;
	value *= 2;
	total /= 4;
	score += func(a, b);
	result -= (x + y);
	product *= calculate(z);
	average /= count + 1;
	`
	expected := []struct {
		identifier string
		literal    string
	}{
		{"x", "10"},
		{"l", "'a'"},
		{"pi", "3.14"},
		{"flag", "true"},
		{"name", "\"hello\""},
		{"zero", "0"},
		{"newline", "'\\n'"},
		{"negative", "(-2.5)"},
		{"falseBool", "false"},
		{"sum", "(10 + 5)"},
		{"diff", "(a - b)"},
		{"product", "(3.14 * 2)"},
		{"complex", "((x + y) * z)"},
		{"comparison", "(a > b)"},
		{"funcCall", "add(5, 10)"},
		{"nested", "calculate((x + y), z)"},
		{"x", "(x + 5)"},
		{"count", "(count - 10)"},
		{"value", "(value * 2)"},
		{"total", "(total / 4)"},
		{"score", "(score + func(a, b))"},
		{"result", "(result - (x + y))"},
		{"product", "(product * calculate(z))"},
		{"average", "(average / (count + 1))"},
	}

	p := New(input)
	statements := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			t.Errorf("Parser Error: %s\n", err.Error())
		}
		t.Fatal("Exiting now!")
	}

	if len(statements.Statements) != len(expected) {
		t.Fatalf("Number of statements not valid, expected %d, got %d", len(expected), len(statements.Statements))
	}

	for i, statement := range statements.Statements {
		stmnt, ok := statement.(*ast.AssignmentStatement)
		if !ok {
			t.Fatalf("[%d] - Not valid statement, expected ast.AssignmentStatement got %T", i, statement)
		}

		if stmnt.Identifier.String() != expected[i].identifier {
			t.Errorf("[%d] - Assignment Identifier name not correct, expected %s, got %s", i, expected[i].identifier, stmnt.Identifier.String())
		}

		if stmnt.Literal != nil && stmnt.Literal.String() != expected[i].literal {
			t.Errorf("[%d] - Assignment Value not correct, expected %s, got %s", i, expected[i].literal, stmnt.Literal.String())
		}
	}
}

func TestFunctionDeclarationStatemenet(t *testing.T) {
	input := `
	int testFunc(int a, char b, float c, string d, bool e){
		int x = a;
		int y = x;
		int g;
	}
	`

	p := New(input)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			t.Errorf("Parser Error: %s\n", err.Error())
		}
		t.Fatal("Exiting now!")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("Number of statements not valid, expected 1, got %d", len(program.Statements))
	}

	stmnt, ok := program.Statements[0].(*ast.DeclarationStatement)

	if !ok {
		t.Fatalf("Statement is not of type ast.DeclarationStatement, got %T", program.Statements[0])
	}

	if stmnt.Type != token.INT {
		t.Errorf("Function return type not correct, expected %s, got %s", token.INT, stmnt.Type)
	}

	if stmnt.Identifier.String() != "testFunc" {
		t.Errorf("Function name not correct, expected %s, got %s", "testFunc", stmnt.Identifier.String())
	}

	funcLiteral, ok := stmnt.Literal.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("Literal is not of type ast.FunctionLiteral, got %T", stmnt.Literal)
	}

	expectedParams := []struct {
		paramType token.TokenType
		paramName string
	}{
		{token.INT, "a"},
		{token.CHAR, "b"},
		{token.FLOAT, "c"},
		{token.STRING, "d"},
		{token.BOOL, "e"},
	}

	if len(funcLiteral.Params) != len(expectedParams) {
		t.Fatalf("Number of parameters not correct, expected %d, got %d", len(expectedParams), len(funcLiteral.Params))
	}

	for i, param := range funcLiteral.Params {
		if param.Type != expectedParams[i].paramType {
			t.Errorf("Parameter %d type not correct, expected %s, got %s", i, expectedParams[i].paramType, param.Type)
		}
		if param.Identifier.String() != expectedParams[i].paramName {
			t.Errorf("Parameter %d name not correct, expected %s, got %s", i, expectedParams[i].paramName, param.Identifier.String())
		}
	}

	if funcLiteral.Block == nil {
		t.Fatal("Function block is nil")
	}

	if len(funcLiteral.Block.Statements) != 3 {
		t.Fatalf("Number of statements in function body not correct, expected 3, got %d", len(funcLiteral.Block.Statements))
	}

	for i, stmt := range funcLiteral.Block.Statements {
		_, ok := stmt.(*ast.DeclarationStatement)
		if !ok {
			t.Errorf("Statement %d in function body is not a DeclarationStatement, got %T", i, stmt)
		}
	}
}
func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return x;
	return a + b;
	return multiply(2, 3);
	return (x + y) * z;
	return "hello world";
	return 'a';
	return 3.14;
	return true;
	return false;
	return -5;
	return !flag;
	return calculate(x, y) + z;
	return nested(func(a), b);
	`
	expected := []string{
		"return 5",
		"return x",
		"return (a + b)",
		"return multiply(2, 3)",
		"return ((x + y) * z)",
		"return \"hello world\"",
		"return 'a'",
		"return 3.14",
		"return true",
		"return false",
		"return (-5)",
		"return (!flag)",
		"return (calculate(x, y) + z)",
		"return nested(func(a), b)",
	}

	p := New(input)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			t.Errorf("Parser Error: %s\n", err.Error())
		}
		t.Fatal("Exiting now!")
	}

	if len(program.Statements) != len(expected) {
		t.Fatalf("Expected %d statements, got %d", len(expected), len(program.Statements))
	}

	for i, statement := range program.Statements {
		stmt, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("Statement %d is not of type ast.ReturnStatement, got %T", i, statement)
		}

		if stmt.String() != expected[i] {
			t.Errorf("Return statement mismatch at index %d, expected %s, got %s", i, expected[i], stmt.String())
		}
	}
}

func TestIfStatement(t *testing.T) {
	input := `
	if(count > 0 && condition == true){
		int x = a;
		int y = x;
		int g;
	}
	else{
		int y = x;
		int g;
	}
	`

	p := New(input)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			t.Errorf("Parser Error: %s\n", err.Error())
		}
		t.Fatal("Exiting now!")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("Number of statements not valid, expected 1, got %d", len(program.Statements))
	}

	stmnt, ok := program.Statements[0].(*ast.IfStatement)
	if !ok {
		t.Fatalf("Statement is not of type ast.IfStatement, got %T", program.Statements[0])
	}

	expectedCondition := "((count > 0) && (condition == true))"
	if stmnt.Condition.String() != expectedCondition {
		t.Errorf("If condition not correct, expected %s, got %s", expectedCondition, stmnt.Condition.String())
	}

	if stmnt.Block == nil {
		t.Fatal("If block is nil")
	}

	if len(stmnt.Block.Statements) != 3 {
		t.Fatalf("Number of statements in if body not correct, expected 3, got %d", len(stmnt.Block.Statements))
	}

	for i, stmt := range stmnt.Block.Statements {
		_, ok := stmt.(*ast.DeclarationStatement)
		if !ok {
			t.Errorf("Statement %d in if body is not a DeclarationStatement, got %T", i, stmt)
		}
	}

	if stmnt.ElseBlock == nil {
		t.Fatal("ElseBlock should not be nil for this test case")
	}

	if len(stmnt.ElseBlock.Statements) != 2 {
		t.Fatalf("Number of statements in else body not correct, expected 2, got %d", len(stmnt.ElseBlock.Statements))
	}

	for i, stmt := range stmnt.ElseBlock.Statements {
		_, ok := stmt.(*ast.DeclarationStatement)
		if !ok {
			t.Errorf("Statement %d in else body is not a DeclarationStatement, got %T", i, stmt)
		}
	}
}

func TestWhileLoopStatement(t *testing.T) {
	input := `
	while(count < 10 && flag == true){
		int x = count;
		count = count + 1;
		int result = x * 2;
	}
	`

	p := New(input)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			t.Errorf("Parser Error: %s\n", err.Error())
		}
		t.Fatal("Exiting now!")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("Number of statements not valid, expected 1, got %d", len(program.Statements))
	}

	stmnt, ok := program.Statements[0].(*ast.WhileStatement)
	if !ok {
		t.Fatalf("Statement is not of type ast.WhileStatement, got %T", program.Statements[0])
	}

	expectedCondition := "((count < 10) && (flag == true))"
	if stmnt.Condition.String() != expectedCondition {
		t.Errorf("While condition not correct, expected %s, got %s", expectedCondition, stmnt.Condition.String())
	}

	if stmnt.Block == nil {
		t.Fatal("While block is nil")
	}

	if len(stmnt.Block.Statements) != 3 {
		t.Fatalf("Number of statements in while body not correct, expected 3, got %d", len(stmnt.Block.Statements))
	}

	// Check first statement is a declaration
	_, ok = stmnt.Block.Statements[0].(*ast.DeclarationStatement)
	if !ok {
		t.Errorf("Statement 0 in while body is not a DeclarationStatement, got %T", stmnt.Block.Statements[0])
	}

	// Check second statement is an assignment
	_, ok = stmnt.Block.Statements[1].(*ast.AssignmentStatement)
	if !ok {
		t.Errorf("Statement 1 in while body is not an AssignmentStatement, got %T", stmnt.Block.Statements[1])
	}

	// Check third statement is a declaration
	_, ok = stmnt.Block.Statements[2].(*ast.DeclarationStatement)
	if !ok {
		t.Errorf("Statement 2 in while body is not a DeclarationStatement, got %T", stmnt.Block.Statements[2])
	}
}

func TestForStatement(t *testing.T) {
	input := `
	for(int i = 0; i < 10; i = i + 1){
		int x = i;
		i = i + 1;
		int result = x * 2;
	}
	`

	p := New(input)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			t.Errorf("Parser Error: %s\n", err.Error())
		}
		t.Fatal("Exiting now!")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("Number of statements not valid, expected 1, got %d", len(program.Statements))
	}

	stmnt, ok := program.Statements[0].(*ast.ForStatement)
	if !ok {
		t.Fatalf("Statement is not of type ast.ForStatement, got %T", program.Statements[0])
	}

	initStmt, ok := stmnt.InitializationStatement.(*ast.DeclarationStatement)
	if !ok {
		t.Fatalf("Initialization statement is not of type ast.DeclarationStatement, got %T", stmnt.InitializationStatement)
	}

	if initStmt.Type != token.INT {
		t.Errorf("Initialization statement type not correct, expected %s, got %s", token.INT, initStmt.Type)
	}

	if initStmt.Identifier.String() != "i" {
		t.Errorf("Initialization statement identifier not correct, expected %s, got %s", "i", initStmt.Identifier.String())
	}

	if initStmt.Literal.String() != "0" {
		t.Errorf("Initialization statement literal not correct, expected %s, got %s", "0", initStmt.Literal.String())
	}

	expectedCondition := "(i < 10)"
	if stmnt.Condition.String() != expectedCondition {
		t.Errorf("For condition not correct, expected %s, got %s", expectedCondition, stmnt.Condition.String())
	}

	if stmnt.Increment == nil {
		t.Fatal("Increment statement is nil")
	}

	if stmnt.Increment.Identifier.String() != "i" {
		t.Errorf("Increment statement identifier not correct, expected %s, got %s", "i", stmnt.Increment.Identifier.String())
	}

	if stmnt.Increment.Literal.String() != "(i + 1)" {
		t.Errorf("Increment statement literal not correct, expected %s, got %s", "(i + 1)", stmnt.Increment.Literal.String())
	}

	if stmnt.Block == nil {
		t.Fatal("For block is nil")
	}

	if len(stmnt.Block.Statements) != 3 {
		t.Fatalf("Number of statements in for body not correct, expected 3, got %d", len(stmnt.Block.Statements))
	}

	_, ok = stmnt.Block.Statements[0].(*ast.DeclarationStatement)
	if !ok {
		t.Errorf("Statement 0 in for body is not a DeclarationStatement, got %T", stmnt.Block.Statements[0])
	}

	_, ok = stmnt.Block.Statements[1].(*ast.AssignmentStatement)
	if !ok {
		t.Errorf("Statement 1 in for body is not an AssignmentStatement, got %T", stmnt.Block.Statements[1])
	}

	_, ok = stmnt.Block.Statements[2].(*ast.DeclarationStatement)
	if !ok {
		t.Errorf("Statement 2 in for body is not a DeclarationStatement, got %T", stmnt.Block.Statements[2])
	}
}
