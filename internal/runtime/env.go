package runtime

type Env struct {
	values map[string]any
}

func NewEnv() *Env {
	return &Env{
		values: map[string]any{},
	}
}

func (e *Env) Define(name string, val any) {
	e.values[name] = val
}

func (e *Env) Get(name string) any {
	return e.values[name]
}
