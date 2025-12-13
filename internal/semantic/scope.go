package semantic

import "fmt"

type Scope struct {
	parent  *Scope
	symbols map[string]bool
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		parent:  parent,
		symbols: map[string]bool{},
	}
}

func (s *Scope) Define(name string) error {
	if s.symbols[name] {
		return fmt.Errorf("variable '%s' is already defined", name)
	}
	s.symbols[name] = true
	return nil
}

func (s *Scope) Resolve(name string) bool {
	if s.symbols[name] {
		return true
	}
	if s.parent != nil {
		return s.parent.Resolve(name)
	}
	return false
}
