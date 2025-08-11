package parser

import (
	"strconv"
	"testing"

	"github.com/mohamedirfanam/cynterpreter/lexer/token"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

func TestIntegerLiterals(t *testing.T) {
	input := `
	123;
	232;
	`
	tests := []int{123, 232}

	p := New(input)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			t.Errorf("Parser Error: %s\n", err.Error())
		}
		t.Fatal("Exiting now!")
	}

	statements := program.Statements

	if len(statements) != len(tests) {
		t.Fatalf("Parser Error: Lenght of statements not correct, expected %v, got %v", len(tests), len(statements))
	}

	for i, statement := range statements {
		stmnt, ok := statement.(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("Statement type not matching, expected ast.ExpressionStatement, got %T", stmnt)
		}

		testIntegralLiteral(t, stmnt.Expression, tests[i])
	}
}

func TestIdentifierExpressios(t *testing.T) {
	input := `
	abc;
	count;
	name;
	`
	tests := []string{"abc", "count", "name"}

	p := New(input)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			t.Errorf("Parser Error: %s\n", err.Error())
		}
		t.Fatal("Exiting now!")
	}

	statements := program.Statements

	if len(statements) != len(tests) {
		t.Fatalf("Parser Error: Lenght of statements not correct, expected %v, got %v", len(tests), len(statements))
	}

	for i, statement := range statements {
		stmnt, ok := statement.(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("Statement type not matching, expected ast.ExpressionStatement, got %T", stmnt)
		}

		testIdentifierExpression(t, stmnt.Expression, tests[i])
	}

}

func TestCharLiterals(t *testing.T) {
	input := `
	'a';
	'Z';
	'5';
	'@';
	' ';
	'\n';
	'\t';
	'\r';
	'\\';
	`
	tests := []byte{
		'a',
		'Z',
		'5',
		'@',
		' ',
		'\n',
		'\t',
		'\r',
		'\\',
	}

	p := New(input)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			t.Errorf("Parser Error: %s\n", err.Error())
		}
		t.Fatal("Exiting now!")
	}

	statements := program.Statements

	if len(statements) != len(tests) {
		t.Fatalf("Parser Error: Length of statements not correct, expected %v, got %v", len(tests), len(statements))
	}

	for i, statement := range statements {
		stmnt, ok := statement.(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("Statement type not matching, expected ast.ExpressionStatement, got %T", stmnt)
		}

		testCharLiteral(t, stmnt.Expression, tests[i])
	}
}

func TestFloatLiterals(t *testing.T) {
	input := `
	3.14;
	0.5;
	123.456;
	0.0;
	1.0;
	999.999;
	`
	tests := []float64{3.14, 0.5, 123.456, 0.0, 1.0, 999.999}

	p := New(input)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			t.Errorf("Parser Error: %s\n", err.Error())
		}
		t.Fatal("Exiting now!")
	}

	statements := program.Statements

	if len(statements) != len(tests) {
		t.Fatalf("Parser Error: Length of statements not correct, expected %v, got %v", len(tests), len(statements))
	}

	for i, statement := range statements {
		stmnt, ok := statement.(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("Statement type not matching, expected ast.ExpressionStatement, got %T", stmnt)
		}

		testFloatLiteral(t, stmnt.Expression, tests[i])
	}
}

func TestStringLiterals(t *testing.T) {
	input := `
	"hello";
	"world";
	"Hello, World!";
	"";
	"line1\nline2";
	"tab\tseparated";
	"quote\"inside";
	"path\\to\\file";
	`
	tests := []string{
		"hello",
		"world",
		"Hello, World!",
		"",
		"line1\nline2",
		"tab\tseparated",
		"quote\"inside",
		"path\\to\\file",
	}

	p := New(input)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			t.Errorf("Parser Error: %s\n", err.Error())
		}
		t.Fatal("Exiting now!")
	}

	statements := program.Statements

	if len(statements) != len(tests) {
		t.Fatalf("Parser Error: Length of statements not correct, expected %v, got %v", len(tests), len(statements))
	}

	for i, statement := range statements {
		stmnt, ok := statement.(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("Statement type not matching, expected ast.ExpressionStatement, got %T", stmnt)
		}

		testStringLiteral(t, stmnt.Expression, tests[i])
	}
}

func TestBoolLiterals(t *testing.T) {
	input := `
	true;
	false;
	true;
	false;
	`
	tests := []bool{true, false, true, false}

	p := New(input)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			t.Errorf("Parser Error: %s\n", err.Error())
		}
		t.Fatal("Exiting now!")
	}

	statements := program.Statements

	if len(statements) != len(tests) {
		t.Fatalf("Parser Error: Length of statements not correct, expected %v, got %v", len(tests), len(statements))
	}

	for i, statement := range statements {
		stmnt, ok := statement.(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("Statement type not matching, expected ast.ExpressionStatement, got %T", stmnt)
		}

		testBoolLiteral(t, stmnt.Expression, tests[i])
	}
}

func testIntegralLiteral(t *testing.T, expression ast.Expression, expectedValue int) {
	expr, ok := expression.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("Expression is not of expected type *ast.IntegerLiteral, got %T", expression)
		return
	}

	if expr.Token.TokenType != token.INT_LITERAL {
		t.Errorf("Token type is not INT_LITERAL")
	}

	if expr.TokenLexeme() != strconv.Itoa(expectedValue) || expr.Value != int64(expectedValue) {
		t.Errorf("Value not correct, Expected Lexeme - %s, Got Lexeme - %s, Expected Value - %d, Got value - %d", strconv.Itoa(expectedValue), expr.TokenLexeme(), expectedValue, expr.Value)
	}
}

func testIdentifierExpression(t *testing.T, expression ast.Expression, expectedValue string) {

	expr, ok := expression.(*ast.IdentifierExpression)
	if !ok {
		t.Errorf("Expression is not of exptected type ast.IdentifierExpression, got %T", expr)
	}

	if expr.Token.TokenType != token.IDENTIFIER {
		t.Errorf("Token type is not IDENTIFIER")
	}

	if expr.TokenLexeme() != expectedValue || expr.Value != expectedValue {
		t.Errorf("Value not correct, Exptected Lexeme - %s, Got Lexeme - %s, Expected Value - %s, Got value - %s", expectedValue, expr.TokenLexeme(), expectedValue, expr.Value)
	}
}

func testCharLiteral(t *testing.T, expression ast.Expression, expectedValue byte) {
	expr, ok := expression.(*ast.CharLiteral)
	if !ok {
		t.Errorf("Expression is not of expected type ast.CharLiteral, got %T", expr)
	}

	if expr.Token.TokenType != token.CHAR_LITERAL {
		t.Errorf("Token type is not CHAR_LITERAL")
	}

	if expr.Value != expectedValue {
		t.Errorf("Value not correct, Expected Value - %c, Got value - %c", expectedValue, expr.Value)
	}
}

func testFloatLiteral(t *testing.T, expression ast.Expression, expectedValue float64) {
	expr, ok := expression.(*ast.FloatLiteral)
	if !ok {
		t.Errorf("Expression is not of expected type ast.FloatLiteral, got %T", expr)
	}

	if expr.Token.TokenType != token.FLOAT_LITERAL {
		t.Errorf("Token type is not FLOAT_LITERAL")
	}

	if expr.Value != expectedValue {
		t.Errorf("Value not correct, Expected Value - %f, Got value - %f", expectedValue, expr.Value)
	}
}

func testStringLiteral(t *testing.T, expression ast.Expression, expectedValue string) {
	expr, ok := expression.(*ast.StringLiteral)
	if !ok {
		t.Errorf("Expression is not of expected type ast.StringLiteral, got %T", expr)
	}

	if expr.Token.TokenType != token.STRING_LITERAL {
		t.Errorf("Token type is not STRING_LITERAL")
	}

	if expr.Value != expectedValue {
		t.Errorf("Value not correct, Expected Value - %s, Got value - %s", expectedValue, expr.Value)
	}
}

func testBoolLiteral(t *testing.T, expression ast.Expression, expectedValue bool) {
	expr, ok := expression.(*ast.BoolLiteral)
	if !ok {
		t.Errorf("Expression is not of expected type ast.BoolLiteral, got %T", expr)
	}

	if expr.Token.TokenType != token.BOOL_LITERAL {
		t.Errorf("Token type is not TRUE or FALSE")
	}

	if expr.Value != expectedValue {
		t.Errorf("Value not correct, Expected Value - %t, Got value - %t", expectedValue, expr.Value)
	}
}

func TestCallExpression(t *testing.T) {
	tests := []struct {
		input        string
		functionName string
		argCount     int
		args         []interface{}
	}{
		{"add();", "add", 0, []interface{}{}},
		{"add(5);", "add", 1, []interface{}{5}},
		{"add(5, 10);", "add", 2, []interface{}{5, 10}},
		{"multiply(2, 3, 4);", "multiply", 3, []interface{}{2, 3, 4}},
		{"print(\"hello\");", "print", 1, []interface{}{"hello"}},
		{"calc(x, y);", "calc", 2, []interface{}{"x", "y"}},
		{"func(1, 2.5);", "func", 2, []interface{}{1, 2.5}},
	}

	for _, tt := range tests {
		p := New(tt.input)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			for _, err := range p.Errors() {
				t.Errorf("Parser Error: %s\n", err.Error())
			}
			t.Fatal("Exiting now!")
		}

		if len(program.Statements) != 1 {
			t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Statement is not of type ast.ExpressionStatement, got %T", program.Statements[0])
		}

		testCallExpression(t, stmt.Expression, tt.functionName, tt.argCount, tt.args)
	}
}

func testCallExpression(t *testing.T, expression ast.Expression, expectedFunctionName string, expectedArgCount int, expectedArgs []interface{}) {
	expr, ok := expression.(*ast.CallExpression)
	if !ok {
		t.Errorf("Expression is not of expected type ast.CallExpression, got %T", expression)
		return
	}

	if expr.Token.TokenType != token.LPAREN {
		t.Errorf("Token type is not LPAREN")
	}

	funcIdent, ok := expr.Function.(*ast.IdentifierExpression)
	if !ok {
		t.Errorf("Function is not an identifier, got %T", expr.Function)
		return
	}

	if funcIdent.Value != expectedFunctionName {
		t.Errorf("Function name mismatch, expected %s, got %s", expectedFunctionName, funcIdent.Value)
	}

	if len(expr.Args) != expectedArgCount {
		t.Errorf("Argument count mismatch, expected %d, got %d", expectedArgCount, len(expr.Args))
		return
	}

	for i, expectedArg := range expectedArgs {
		switch v := expectedArg.(type) {
		case int:
			testIntegralLiteral(t, expr.Args[i], v)
		case float64:
			testFloatLiteral(t, expr.Args[i], v)
		case string:
			if _, ok := expr.Args[i].(*ast.IdentifierExpression); ok {
				testIdentifierExpression(t, expr.Args[i], v)
			} else {
				testStringLiteral(t, expr.Args[i], v)
			}
		default:
			t.Errorf("Unsupported argument type: %T", v)
		}
	}
}

func TestInfixExpressions(t *testing.T) {
	tests := []struct {
		input      string
		operator   string
		leftValue  int
		rightValue int
	}{
		{"5 + 3;", "+", 5, 3},
		{"5 - 3;", "-", 5, 3},
		{"5 * 3;", "*", 5, 3},
		{"5 / 3;", "/", 5, 3},
		{"5 % 3;", "%", 5, 3},
		{"5 == 3;", "==", 5, 3},
		{"5 != 3;", "!=", 5, 3},
		{"5 < 3;", "<", 5, 3},
		{"5 <= 3;", "<=", 5, 3},
		{"5 > 3;", ">", 5, 3},
		{"5 >= 3;", ">=", 5, 3},
	}

	for _, tt := range tests {
		p := New(tt.input)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			for _, err := range p.Errors() {
				t.Errorf("Parser Error: %s\n", err.Error())
			}
			t.Fatal("Exiting now!")
		}

		if len(program.Statements) != 1 {
			t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Statement is not of type ast.ExpressionStatement, got %T", program.Statements[0])
		}

		expr, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("Expression is not of type ast.InfixExpression, got %T", stmt.Expression)
		}

		if expr.Op != tt.operator {
			t.Errorf("Operator mismatch, expected %s, got %s", tt.operator, expr.Op)
		}

		testIntegralLiteral(t, expr.LeftExp, tt.leftValue)
		testIntegralLiteral(t, expr.RightExp, tt.rightValue)
	}
}

func TestPrefixExpressions(t *testing.T) {
	tests := []struct {
		input    string
		operator string
		value    int
	}{
		{"-5;", "-", 5},
		{"+5;", "+", 5},
		{"-15;", "-", 15},
		{"+15;", "+", 15},
		{"-0;", "-", 0},
		{"+0;", "+", 0},
		{"-999;", "-", 999},
		{"+999;", "+", 999},
	}

	for _, tt := range tests {
		p := New(tt.input)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			for _, err := range p.Errors() {
				t.Errorf("Parser Error: %s\n", err.Error())
			}
			t.Fatal("Exiting now!")
		}

		if len(program.Statements) != 1 {
			t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Statement is not of type ast.ExpressionStatement, got %T", program.Statements[0])
		}

		expr, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("Expression is not of type ast.PrefixExpression, got %T", stmt.Expression)
		}

		if expr.Op != tt.operator {
			t.Errorf("Operator mismatch, expected %s, got %s", tt.operator, expr.Op)
		}

		testIntegralLiteral(t, expr.Exp, tt.value)
	}
}

func TestExpressions(t *testing.T) {
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
	5 == 3;
	x != y;
	a > b;
	count < limit;
	value >= threshold;
	score <= max;
	x == y && a > b;
	condition1 || condition2;
	a > b && c < d;
	x == y || a != b;
	(a > b) && (c <= d);
	x + y > z && result;
	a == b || c > d && e < f;
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
		"(5 == 3)",
		"(x != y)",
		"(a > b)",
		"(count < limit)",
		"(value >= threshold)",
		"(score <= max)",
		"((x == y) && (a > b))",
		"(condition1 || condition2)",
		"((a > b) && (c < d))",
		"((x == y) || (a != b))",
		"((a > b) && (c <= d))",
		"(((x + y) > z) && result)",
		"((a == b) || ((c > d) && (e < f)))",
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
			t.Fatalf("Statement is not of type ast.ExpressionStatement, got %T", program.Statements[0])
		}

		var expr ast.Expression = stmt.Expression

		if expr.String() != expected[i] {
			t.Errorf("Expression mismatch at index %d, expected %s, got %s", i, expected[i], expr.String())
		}
	}
}
