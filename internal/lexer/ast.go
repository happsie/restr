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

type ReqExpression struct {
	Method string
	URL    Expression
	Blocks []ReqBlock
}

type ReqBlock interface {
	reqBlock()
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

func (*ReqExpression) node() {

}

func (*ReqExpression) expr() {

}

type HeadersBlock struct {
	Pairs map[string]Expression
}

func (*HeadersBlock) reqBlock() {}

type JsonBlock struct {
	Value map[string]Expression
}

func (*JsonBlock) reqBlock() {}

type BlockStatement struct {
	Statements []Statement
}

func (*BlockStatement) stmt() {}
func (*BlockStatement) node() {}
