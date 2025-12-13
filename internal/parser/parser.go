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
	tok := p.current()

	switch tok.Type {
	case lexer.STRING:
		p.advance()
		str, ok := tok.Literal.(string)
		if !ok {
			panic("could not parse string expression" + tok.Lexeme)
		}
		return &lexer.StringLiteral{
			Value: str,
		}

	case lexer.IDENTIFIER:
		p.advance()
		return &lexer.Identifier{
			Name: tok.Lexeme,
		}

	default:
		panic(fmt.Sprintf(
			"unexpected token in expression: %s",
			tok.Type,
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
	default:
		panic("could not parse statement" + p.current().Type)
	}
}

func (p *Parser) ParseProgram() lexer.Program {
	program := lexer.Program{}

	for p.current().Type != lexer.EOF {
		stmt := p.parseStatement()
		program.Statements = append(program.Statements, stmt)
	}

	return program
}
