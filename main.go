package main

import (
	"fmt"
	"monkey/lexer"
	"monkey/token"
)

func main() {
	input := "let x = 5;"
	l := lexer.New(input)
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
