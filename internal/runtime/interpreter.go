package runtime

import (
	"fmt"
	"io"
	"net/http"
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
	case *lexer.ReqExpression:
		return i.evalReq(e)
	default:
		panic(fmt.Errorf("unknown expression: %v", expr))
	}
}

type HttpResponse struct {
	Body    string
	Status  int
	Headers map[string]string
}

func (i *Interpreter) evalReq(req *lexer.ReqExpression) HttpResponse {
	url := i.evalExpr(req.URL).(string)

	client := http.DefaultClient

	r, err := http.NewRequest(req.Method, url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(body)
	}
	return HttpResponse{
		Body:   string(body),
		Status: resp.StatusCode,
		Headers: map[string]string{
			"test": "test",
		},
	}
}
