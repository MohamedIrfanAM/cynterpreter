package parser

import (
	"testing"

	"github.com/mohamedirfanam/cynterpreter/lexer/token"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

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
