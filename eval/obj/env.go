package obj

import (
	"fmt"
)

type Environment struct {
	memory map[string]Object
}

func NewEnv() *Environment {
	memory := make(map[string]Object)
	return &Environment{
		memory: memory,
	}
}

func (env *Environment) SetVar(varname string, val Object) {
	env.memory[varname] = val
}

func (env *Environment) GetIndexVar(varname string, index int) (Object, error) {
	object, ok := env.memory[varname]
	if !ok {
		return nil, fmt.Errorf("%s variable doesn't exist", varname)
	}
	switch val := object.(type) {
	case *ArrayObject:
		if index >= val.Length {
			return nil, fmt.Errorf("invalid index, index greater than lenghth of the string %d", val.Length)
		}
		return val.Vals[index], nil
	case *StringObject:
		if index >= len(val.Value) {
			return nil, fmt.Errorf("invalid index, index greater than lenghth of the string %d", len(val.Value))
		}
		return &CharObject{Value: val.Value[index]}, nil
	}
	return object, nil
}

func (env *Environment) SetIndexVar(varname string, index int, updateVal Object) error {
	object, ok := env.memory[varname]
	if !ok {
		return fmt.Errorf("%s variable doesn't exist", varname)
	}
	switch val := object.(type) {
	case *ArrayObject:
		if index >= val.Length {
			return fmt.Errorf("invalid index, index greater than lenghth of the string %d", val.Length)
		}
		if val.DataType != updateVal.Type() {
			return fmt.Errorf("type error,cannot assign %s to %s", object.Type(), val.DataType)
		}
		val.Vals[index] = updateVal
		return nil
	case *StringObject:
		if index >= len(val.Value) {
			return fmt.Errorf("invalid index, index greater than lenghth of the string %d", len(val.Value))
		}
		newChar, ok := updateVal.(*CharObject)
		if !ok {
			return fmt.Errorf("expectecd a char literal to assign to string indexed")
		}
		b := []byte(val.Value)
		b[index] = newChar.Value
		val.Value = string(b)
	}
	return nil
}

func (env *Environment) GetVar(varname string) (Object, bool) {
	val, ok := env.memory[varname]
	if ok {
		return val, true
	}
	return nil, false
}

func (env *Environment) ExtendEnv() *Environment {
	newEnv := NewEnv()
	for k, v := range env.memory {
		if v.Type() == FUNCTION_OBJ {
			newEnv.memory[k] = v
		}
	}
	return newEnv
}

func (env *Environment) CopyEnv() *Environment {
	newEnv := NewEnv()
	for k, v := range env.memory {
		newEnv.memory[k] = v
	}
	return newEnv
}

func (env *Environment) UpdateVals(newEnv *Environment) {
	for k := range env.memory {
		val, ok := newEnv.memory[k]
		if ok {
			env.memory[k] = val
		}
	}
}
