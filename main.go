package main

import (
	"fmt"
	"os"
)

func main() {
	b, err := os.ReadFile("lexer_test.rtr")
	if err != nil {
		panic(err)
	}

	s := Scanner{}
	tokens, err := s.Tokenize(string(b))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", tokens)
}
