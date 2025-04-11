package main

import (
	"fmt"
	"monkey/repl"
	"os"
)

func main() {
	fmt.Println("Welcome to the Monkey REPL!")
	fmt.Println("This is my attempt at constructing an interpretter for Monkey language (type 'exit' to quit):")
	repl.Start(os.Stdin, os.Stdout)
}
