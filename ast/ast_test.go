package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	// transformed Ball's original test
	// Ball's original TestString function
	// was fairly ugly, he was probably running
	// out of space on the page
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{
					Type:    token.LET,
					Literal: "let",
				},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}

	const expected = "let myVar = anotherVar;"
	got := program.String()

	if got != expected {
		t.Errorf("program.String() wrong. expected=%q, got=%q", expected, got)
	}
}
