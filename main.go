package main

import (
	"encoding/json"
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
	data, _ := json.MarshalIndent(tokens, "", "    ")
	fmt.Println(string(data))
}
