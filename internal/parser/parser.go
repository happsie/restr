package parser

import (
	"fmt"
	"restr/internal/lexer"
)

type Parser struct {
	tokens []lexer.Token
	pos    int
}

func New(tokens []lexer.Token) *Parser {
	return &Parser{
		tokens: tokens,
		pos:    0,
	}
}

func (p *Parser) current() lexer.Token {
	return p.tokens[p.pos]
}

func (p *Parser) advance() lexer.Token {
	tok := p.current()
	p.pos++
	return tok
}

func (p *Parser) parseExpression() lexer.Expression {
	token := p.current()

	switch token.Type {
	case lexer.STRING:
		p.advance()
		str, ok := token.Literal.(string)
		if !ok {
			panic("could not parse string expression: " + token.Lexeme)
		}
		return &lexer.StringLiteral{
			Value: str,
		}

	case lexer.IDENTIFIER:
		p.advance()
		return &lexer.Identifier{
			Name: token.Lexeme,
		}
	case lexer.REQ:
		return p.parseReqExpression()

	default:
		panic(fmt.Sprintf(
			"unexpected token in expression: %s",
			token.Type,
		))
	}
}

func (p *Parser) parseVarDeclaration() lexer.Statement {
	p.advance()

	name := p.advance().Lexeme // TODO: Ska vi verkligen använda lexeme som värde för var?
	p.advance()

	value := p.parseExpression()

	return &lexer.VarDeclaration{
		Name:  name,
		Value: value,
	}
}

func (p *Parser) parsePrintStatement() lexer.Statement {
	p.advance()

	expr := p.parseExpression()

	return &lexer.PrintStatement{
		Expr: expr,
	}
}

func (p *Parser) parseStatement() lexer.Statement {
	switch p.current().Type {
	case lexer.VAR:
		return p.parseVarDeclaration()
	case lexer.PRINT:
		return p.parsePrintStatement()
	case lexer.LEFT_BRACE:
		return p.parseBlockStatement()
	default:
		panic("could not parse statement" + p.current().Type)
	}
}

func (p *Parser) parseReqExpression() lexer.Expression {
	p.expect(lexer.REQ)

	method := p.expect(lexer.GET).Lexeme // TODO: really lexeme?
	url := p.parseExpression()

	return &lexer.ReqExpression{
		Method: method,
		URL:    url,
		Blocks: []lexer.ReqBlock{},
	}
}

func (p *Parser) parseBlockStatement() lexer.Statement {
	p.expect(lexer.LEFT_BRACE)

	stmts := []lexer.Statement{}

	for p.current().Type != lexer.RIGHT_BRACE {
		stmts = append(stmts, p.parseStatement())
	}

	p.expect(lexer.RIGHT_BRACE)

	return &lexer.BlockStatement{
		Statements: stmts,
	}
}

func (p *Parser) expect(t lexer.TokenType) lexer.Token {
	if p.current().Type != t {
		panic("expected " + t + ", got " + p.current().Type)
	}
	return p.advance()
}

func (p *Parser) ParseProgram() lexer.Program {
	program := lexer.Program{}

	for p.current().Type != lexer.EOF {
		stmt := p.parseStatement()
		program.Statements = append(program.Statements, stmt)
	}

	return program
}
