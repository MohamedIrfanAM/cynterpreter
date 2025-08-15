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

		results := Eval(stmnt)

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
