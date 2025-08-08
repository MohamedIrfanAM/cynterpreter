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

	for i, stmnt := range statements {
		testIntegralLiteral(t, stmnt, tests[i])
	}
}

func testIntegralLiteral(t *testing.T, statement ast.Statement, exptectedValue int) {
	stmnt, ok := statement.(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("Statement type not matching, expected ast.ExpressionStatement, got %T", stmnt)
	}

	expr, ok := stmnt.Expression.(*ast.IntegerLiteral)
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
