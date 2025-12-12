package main

import (
	"fmt"
	"strings"
)

type TokenType string

const (
	// One or two character tokens
	EQUAL       TokenType = "EQUAL"
	EQUAL_EQUAL TokenType = "EQUAL_EQUAL"
	RIGHT_BRACE TokenType = "RIGHT_BRACE"
	LEFT_BRACE  TokenType = "LEFT_BRACE"
	RIGHT_PAREN TokenType = "RIGHT_PAREN"
	LEFT_PAREN  TokenType = "LEFT_PAREN"
	COLON       TokenType = "COLON"

	// Literals
	STRING TokenType = "STRING"
	NUMBER TokenType = "NUMBER"

	// Keywords
	VAR     TokenType = "VAR"
	REQ     TokenType = "REQ"
	HEADERS TokenType = "HEADERS"
	BODY    TokenType = "BODY"
	POST    TokenType = "POST"
	GET     TokenType = "GET"
	PUT     TokenType = "PUT"
	DELETE  TokenType = "DELETE"
	PATCH   TokenType = "PATCH"

	EOF TokenType = "EOF"
)

type Token struct {
	Type    TokenType
	Literal any
	Line    int
	Lexeme  string
}

type Scanner struct {
	line    int
	start   int
	current int
	source  string
	tokens  []Token
}

func (s *Scanner) Tokenize(code string) ([]Token, error) {
	for i, row := range strings.Split(code, "\n") {
		fmt.Printf("Converting line %d to tokens\n", i)
		s.source = row
		s.line = i
		s.current = 0

		for !s.atEnd() {
			s.start = s.current
			ignore, token, err := s.scanToken()
			if err != nil {
				return nil, err
			}
			if ignore {
				continue
			}
			s.tokens = append(s.tokens, token)
		}
	}
	s.tokens = append(s.tokens, s.addToken(EOF))
	return s.tokens, nil
}

func (s *Scanner) scanToken() (bool, Token, error) {
	ch := s.advance()
	switch ch {
	case "=":
		if s.match("=") {
			return false, s.addToken(EQUAL_EQUAL), nil
		} else {
			return false, s.addToken(EQUAL), nil
		}
	case "(":
		return false, s.addToken(LEFT_PAREN), nil
	case ")":
		return false, s.addToken(RIGHT_PAREN), nil
	case "{":
		return false, s.addToken(LEFT_BRACE), nil
	case "}":
		return false, s.addToken(RIGHT_BRACE), nil
	case ":":
		return false, s.addToken(COLON), nil
	case "/":
		if s.match("/") { // A comment, ignore it as token.
			for s.peek() != "\n" && !s.atEnd() {
				s.advance()
			}
		}
		// TODO: If we want to handle division etc, add an else here to create that token instead of an ignored one.
		return true, Token{}, nil
	case "'":
		for s.peek() != "'" && !s.atEnd() {
			s.advance()
		}
		if s.atEnd() {
			return false, Token{}, fmt.Errorf("unterminated string")
		}
		s.advance()

		value := s.source[s.start+1 : s.current-1]
		return false, s.addTokenWithLiteral(STRING, value), nil
	case " ":
		return true, Token{}, nil
	case "\r":
		return true, Token{}, nil
	case "\t":
		return true, Token{}, nil
	default:
		return false, Token{}, fmt.Errorf("unexpected character at line: %d position: %d, %s", s.line, s.current, string(ch))
	}
}

func (s *Scanner) advance() string {
	ch := s.source[s.current]
	s.current++
	return string(ch)
}

func (s *Scanner) peek() string {
	if s.atEnd() {
		return ""
	}
	return string(s.source[s.current])
}

func (s *Scanner) match(expected string) bool {
	if len(expected) > 1 {
		panic("can only match next with 1 character")
	}
	if s.atEnd() {
		return false
	}
	if string(s.source[s.current]) != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) atEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) addToken(token TokenType) Token {
	return Token{
		Type:    token,
		Literal: nil,
		Line:    s.line,
		Lexeme:  "",
	}
}

func (s *Scanner) addTokenWithLiteral(token TokenType, literal any) Token {
	return Token{
		Type:    token,
		Literal: literal,
		Line:    s.line,
		Lexeme:  "",
	}
}
