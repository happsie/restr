package runtime

import (
	"fmt"
	"restr/internal/lexer"
)

type Interpreter struct {
	env *Env
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		env: NewEnv(),
	}
}

func (i *Interpreter) Run(p lexer.Program) {
	for _, stmt := range p.Statements {
		i.execStmt(stmt)
	}
}

func (i *Interpreter) execStmt(stmt lexer.Statement) {
	switch s := stmt.(type) {
	case *lexer.VarDeclaration:
		val := i.evalExpr(s.Value)
		i.env.Define(s.Name, val)
	case *lexer.PrintStatement:
		val := i.evalExpr(s.Expr)
		fmt.Println(val)
	}
}

func (i *Interpreter) evalExpr(expr lexer.Expression) any {
	switch e := expr.(type) {
	case *lexer.StringLiteral:
		return e.Value
	case *lexer.Identifier:
		return i.env.Get(e.Name)
	default:
		panic("unknown expression")
	}
}
