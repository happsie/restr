package main

import (
	"encoding/json"
	"fmt"
	"os"
	"restr/internal/lexer"
	"restr/internal/parser"
	"restr/internal/runtime"
	"restr/internal/semantic"
)

func main() {
	fmt.Println("======= TOKEN VERIFICATION =======")
	tv := tokens("rtr/lexer_test.rtr")
	printPrettyJson(tv)

	fmt.Println("======= PARSER VERIFICATION =======")

	tp := tokens("rtr/parser_test.rtr")
	p := parser.New(tp)
	stmts := p.ParseProgram()
	printPrettyJson(stmts)

	analyzer := semantic.NewAnalyzer()
	err := analyzer.AnalyzeProgram(stmts)
	if err != nil {
		fmt.Println(err)
	}

	i := runtime.NewInterpreter()
	i.Run(stmts)
}

func tokens(file string) []lexer.Token {
	b, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	s := lexer.Scanner{}
	tokens, err := s.FindTokens(string(b))
	if err != nil {
		panic(err)
	}
	return tokens
}

func printPrettyJson(v any) {
	data, _ := json.MarshalIndent(v, "", "    ")
	fmt.Println(string(data))
}
