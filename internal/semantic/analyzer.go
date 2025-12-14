package semantic

import (
	"fmt"
	"restr/internal/lexer"
)

type Analyzer struct {
	scope *Scope
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{
		scope: NewScope(nil),
	}
}

func (a *Analyzer) AnalyzeProgram(p lexer.Program) error {
	for _, stmt := range p.Statements {
		if err := a.analyzeStmt(stmt); err != nil {
			return err
		}
	}
	return nil
}

func (a *Analyzer) analyzeStmt(stmt lexer.Statement) error {
	switch s := stmt.(type) {
	case *lexer.VarDeclaration:
		if err := a.analyzeExpr(s.Value); err != nil {
			return err
		}
		return a.scope.Define(s.Name)
	case *lexer.PrintStatement:
		return a.analyzeExpr(s.Expr)
	default:
		return fmt.Errorf("unknown statement")
	}
}

func (a *Analyzer) analyzeExpr(expr lexer.Expression) error {
	switch e := expr.(type) {
	case *lexer.StringLiteral:
		return nil
	case *lexer.Identifier:
		if !a.scope.Resolve(e.Name) {
			return fmt.Errorf("undefined variable '%s'", e.Name)
		}
		return nil
	default:
		return fmt.Errorf("unknown expression: %v", expr)
	}
}
