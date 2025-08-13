package obj

import "fmt"

type ObjType string

const (
	ERROR_OBJ   ObjType = "ERROR_OBJ"
	NULL_OBJ    ObjType = "NULL_OBJ"
	INTEGER_OBJ ObjType = "INTEGER_OBJ"
	BOOLEAN_OBJ ObjType = "BOOLEAN_OBJ"
	CHAR_OBJ    ObjType = "CHAR_OBJ"
	STRING_OBJ  ObjType = "STRING_OBJ"
	FLOAT_OBJ   ObjType = "FLOAT_OBJ"
)

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
	Value string
}

func (e *ErrorObject) Type() ObjType {
	return ERROR_OBJ
}

func (e *ErrorObject) String() string {
	return fmt.Sprintf("Error: %s", e.Value)
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

func (b *BooleanObject) String() ObjType {
	return ObjType(fmt.Sprintf("%t", b.Value))
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
