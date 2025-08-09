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

func testIntegralLiteral(t *testing.T, expression ast.Expression, exptectedValue int) {

	expr, ok := expression.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("Expression is not of exptected type ast.IntegerLiteral, got %T", expr)
	}

	if expr.Token.TokenType != token.INT_LITERAL {
		t.Errorf("Token type is not INT_LITERAL")
	}

	if expr.TokenLexeme() != strconv.Itoa(exptectedValue) || expr.Value != int64(exptectedValue) {
		t.Errorf("Value not correct, Exptected Lexeme - %s, Got Lexeme - %s, Expected Value - %d, Got value - %d", strconv.Itoa(exptectedValue), expr.TokenLexeme(), exptectedValue, expr.Value)
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
		expr, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("Expression is not of type ast.InfixExpression, got %T", stmt.Expression)
		}

		if expr.String() != expected[i] {
			t.Errorf("Expression mismatch, expected %s, got %s", expected[i], expr.String())
		}
	}
}
