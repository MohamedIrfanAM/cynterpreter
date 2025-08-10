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

		if stmnt.Literal.String() != expected[i].literal {
			t.Errorf("[%d] - Declaration Literal not correct, expected %s, got %s", i, expected[i].literal, stmnt.Literal.String())
		}
	}
}
