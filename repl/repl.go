package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		io.WriteString(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if line == "exit" {
			fmt.Fprintln(out, "Thank you for using my Monkey interpreter")
			return
		}
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
