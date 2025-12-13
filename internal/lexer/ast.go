package lexer

type Program struct {
	Statements []Statement
}

type Node interface {
	node()
}

type Statement interface {
	Node
	stmt()
}

type Expression interface {
	Node
	expr()
}

type VarDeclaration struct {
	Name  string
	Value Expression
}

type StringLiteral struct {
	Value string
}

type PrintStatement struct {
	Expr Expression
}

type Identifier struct {
	Name string
}

func (*Program) node() {}

func (*VarDeclaration) node() {

}

func (*VarDeclaration) stmt() {

}

func (*PrintStatement) node() {

}

func (*PrintStatement) stmt() {

}

func (*StringLiteral) node() {

}

func (*StringLiteral) expr() {

}

func (*Identifier) node() {

}

func (*Identifier) expr() {

}
