package obj

import (
	"fmt"
	"strings"

	"github.com/mohamedirfanam/cynterpreter/lexer/token"
)

type ObjType string

const (
	ERROR_OBJ   ObjType = "ERROR_OBJ"
	NULL_OBJ    ObjType = "NULL_OBJ"
	INTEGER_OBJ ObjType = "INTEGER_OBJ"
	BOOLEAN_OBJ ObjType = "BOOLEAN_OBJ"
	CHAR_OBJ    ObjType = "CHAR_OBJ"
	STRING_OBJ  ObjType = "STRING_OBJ"
	FLOAT_OBJ   ObjType = "FLOAT_OBJ"

	RESULTS_OBJ ObjType = "RESULTS_OBJ"
)

var (
	TRUE  = &BooleanObject{Value: true}
	FALSE = &BooleanObject{Value: false}
	NULL  = &NullObject{}
)

func GetObjectType(tknType token.TokenType) ObjType {
	switch tknType {
	case token.INT:
		return INTEGER_OBJ
	case token.CHAR:
		return CHAR_OBJ
	case token.BOOL:
		return BOOLEAN_OBJ
	case token.STRING:
		return STRING_OBJ
	case token.FLOAT:
		return FLOAT_OBJ
	case token.DOUBLE:
		return FLOAT_OBJ
	default:
		return ERROR_OBJ
	}
}

type Object interface {
	Type() ObjType
	String() string
}

// NULL object
type NullObject struct {
}

func (n *NullObject) Type() ObjType {
	return NULL_OBJ
}
func (n *NullObject) String() string {
	return "null"
}

// Error Object
type ErrorObject struct {
	Error error
}

func (e *ErrorObject) Type() ObjType {
	return ERROR_OBJ
}

func (e *ErrorObject) String() string {
	return e.Error.Error()
}

func NewError(err error) *ErrorObject {
	return &ErrorObject{
		Error: err,
	}
}

// Integer Object
type IntegerObject struct {
	Value int64
}

func (i *IntegerObject) Type() ObjType {
	return INTEGER_OBJ
}

func (i *IntegerObject) String() string {
	return fmt.Sprintf("%d", i.Value)
}

// Float Objtect
type FloatObject struct {
	Value float64
}

func (f *FloatObject) Type() ObjType {
	return FLOAT_OBJ
}
func (f *FloatObject) String() string {
	return fmt.Sprintf("%f", f.Value)
}

// Boolean Object
type BooleanObject struct {
	Value bool
}

func (b *BooleanObject) Type() ObjType {
	return BOOLEAN_OBJ
}

func (b *BooleanObject) String() string {
	return fmt.Sprintf("%t", b.Value)
}

func GetBoolean(state bool) *BooleanObject {
	if state {
		return TRUE
	}
	return FALSE
}

// Char Object
type CharObject struct {
	Value byte
}

func (c *CharObject) Type() ObjType {
	return CHAR_OBJ
}

func (c *CharObject) String() string {
	return fmt.Sprintf("%c", c.Value)
}

// String Object
type StringObject struct {
	Value string
}

func (s *StringObject) Type() ObjType {
	return STRING_OBJ
}

func (s *StringObject) String() string {
	return s.Value
}

// Results Object
type ResultsObject struct {
	Results []Object
}

func (r *ResultsObject) Type() ObjType {
	return RESULTS_OBJ
}

func (r *ResultsObject) String() string {
	var str strings.Builder
	for i, result := range r.Results {
		str.WriteString(result.String())
		if i < len(r.Results)-1 {
			str.WriteRune('\n')
		}
	}
	return str.String()
}
