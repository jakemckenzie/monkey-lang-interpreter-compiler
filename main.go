package main

import (
	"fmt"
	"monkey/lexer"
	"monkey/token"
)

func main() {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	}
	fn(ten, ten)`
	l := lexer.New(input)
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
