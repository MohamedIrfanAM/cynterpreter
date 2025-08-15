package eval

import (
	"github.com/mohamedirfanam/cynterpreter/eval/obj"
	"github.com/mohamedirfanam/cynterpreter/parser/ast"
)

func evalIfStatement(ifs *ast.IfStatement) obj.Object {
	result := Eval(ifs.Condition)
	condition := IsTrue(result)
	if condition {
		return evalBlock(ifs.Block)
	} else if !condition && ifs.ElseBlock != nil {
		return evalBlock(ifs.ElseBlock)
	}
	return obj.NULL
}

func evalBlock(blk *ast.Block) obj.Object {
	var results []obj.Object
	for _, stmnt := range blk.Statements {
		result := Eval(stmnt)
		if result.Type() == obj.ERROR_OBJ {
			return result
		}
		results = append(results, result)
	}
	return &obj.ResultsObject{Results: results}
}
