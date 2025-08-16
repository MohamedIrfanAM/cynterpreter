package obj

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
