package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
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

	IDENTIFIER TokenType = "IDENTIFIER"

	// Keywords
	VAR     TokenType = "VAR"
	REQ     TokenType = "REQ"
	HEADERS TokenType = "HEADERS"
	JSON    TokenType = "JSON"
	POST    TokenType = "POST"
	GET     TokenType = "GET"
	PUT     TokenType = "PUT"
	DELETE  TokenType = "DELETE"
	PATCH   TokenType = "PATCH"
	PRINT   TokenType = "print"

	EOF TokenType = "EOF"
)

var keywords = map[string]TokenType{
	"var":     VAR,
	"req":     REQ,
	"headers": HEADERS,
	"json":    JSON,
	"POST":    POST,
	"GET":     GET,
	"PUT":     PUT,
	"DELETE":  DELETE,
	"PATCH":   PATCH,
	"print":   PRINT,
}

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
		s.start = 0

		for !s.atEnd() {
			s.start = s.current
			err := s.scanToken()
			if err != nil {
				return nil, err
			}
		}
	}
	s.addToken(EOF)
	return s.tokens, nil
}

func (s *Scanner) scanToken() error {
	ch := s.advance()
	switch ch {
	case "=":
		if s.match("=") {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case "(":
		s.addToken(LEFT_PAREN)
	case ")":
		s.addToken(RIGHT_PAREN)
	case "{":
		s.addToken(LEFT_BRACE)
	case "}":
		s.addToken(RIGHT_BRACE)
	case ":":
		s.addToken(COLON)
	case "/":
		if s.match("/") { // A comment, ignore it as token.
			for s.peek() != "\n" && !s.atEnd() {
				s.advance()
			}
		}
		// TODO: If we want to handle division etc, add an else here to create that token instead of an ignored one.
	case "'":
		for s.peek() != "'" && !s.atEnd() {
			s.advance()
		}
		if s.atEnd() {
			return fmt.Errorf("unterminated string")
		}
		s.advance()

		value := s.source[s.start+1 : s.current-1]
		s.addTokenWithLiteral(STRING, value)
	case " ":
	case "\r":
	case "\t":
		break
	default:
		if s.isDigit(ch) {
			err := s.parseNumber()
			if err != nil {
				return err
			}
		} else if s.isAlpha(ch) {
			for s.isAlphaNumeric(s.peek()) {
				s.advance()
			}

			text := s.source[s.start:s.current]
			tokenType := keywords[text]
			if tokenType == "" {
				s.addToken(IDENTIFIER)
				break
			}
			s.addToken(tokenType)
		} else {
			return fmt.Errorf("unexpected character at line: %d position: %d, %s", s.line, s.current, s.source[s.start:s.current])
		}
	}
	return nil
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

func (s *Scanner) peekNext() string {
	if s.atEnd() {
		return ""
	}
	return string(s.source[s.current+1])
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

func (s *Scanner) addToken(token TokenType) {
	t := Token{
		Type:    token,
		Literal: nil,
		Line:    s.line,
		Lexeme:  s.lexeme(),
	}
	s.tokens = append(s.tokens, t)
}

func (s *Scanner) addTokenWithLiteral(token TokenType, literal any) {
	t := Token{
		Type:    token,
		Literal: literal,
		Line:    s.line,
		Lexeme:  s.lexeme(),
	}
	s.tokens = append(s.tokens, t)
}

func (s *Scanner) lexeme() string {
	return s.source[s.start:s.current]
}

func (s *Scanner) isDigit(ch string) bool {
	r, _ := utf8.DecodeRuneInString(ch)
	return unicode.IsDigit(r)
}

func (s *Scanner) isAlpha(ch string) bool {
	r, _ := utf8.DecodeRuneInString(ch)
	return unicode.IsLetter(r)
}

func (s *Scanner) isAlphaNumeric(ch string) bool {
	return s.isAlpha(ch) || s.isDigit(ch)
}

func (s *Scanner) parseNumber() error {
	for s.isDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == "." && s.isDigit(s.peekNext()) {
		s.advance()
		for s.isDigit(s.peek()) {
			s.advance()
		}
	}

	f, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		return fmt.Errorf("unexpected character at line: %d position: %d, %s", s.line, s.current, s.source[s.start:s.current])
	}
	s.addTokenWithLiteral(NUMBER, f)
	return nil
}
