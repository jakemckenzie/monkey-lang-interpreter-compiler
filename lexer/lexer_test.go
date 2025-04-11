package lexer

import (
	"monkey/token"
	"testing"
)

func TestCodeChunk(t *testing.T) {
	input := `
let five = 5;
let ten = 10;
let fn(ten, five);
let add = fn(x, y) {
	x + y;
};
let result = add(five, ten);
!-/*5;
5 < 10 > 5;
if (5 < 10) {
	return true;
} else {
	return false;
};

`

	expected := []token.Token{
		// let five = 5;
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "five"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},

		// let ten = 10;
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "ten"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "10"},
		{Type: token.SEMICOLON, Literal: ";"},

		//let fn(ten, five);
		{Type: token.LET, Literal: "let"},
		{Type: token.FN, Literal: "fn"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IDENT, Literal: "ten"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENT, Literal: "five"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.SEMICOLON, Literal: ";"},

		// let add = fn(x, y) {
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "add"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.FN, Literal: "fn"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IDENT, Literal: "x"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENT, Literal: "y"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},

		// // x + y;
		{Type: token.IDENT, Literal: "x"},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.IDENT, Literal: "y"},
		{Type: token.SEMICOLON, Literal: ";"},

		// // };
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.SEMICOLON, Literal: ";"},

		// let result = add(five, ten);
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "result"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.IDENT, Literal: "add"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IDENT, Literal: "five"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENT, Literal: "ten"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.SEMICOLON, Literal: ";"},

		// !-/*5;
		{Type: token.BANG, Literal: "!"},
		{Type: token.MINUS, Literal: "-"},
		{Type: token.SLASH, Literal: "/"},
		{Type: token.ASTERISK, Literal: "*"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},

		// 5 < 10 > 5;
		{Type: token.INT, Literal: "5"},
		{Type: token.LT, Literal: "<"},
		{Type: token.INT, Literal: "10"},
		{Type: token.GT, Literal: ">"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},

		// if (5 < 10) {
		{Type: token.IF, Literal: "if"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.INT, Literal: "5"},
		{Type: token.LT, Literal: "<"},
		{Type: token.INT, Literal: "10"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},

		// return true;
		{Type: token.RETURN, Literal: "return"},
		{Type: token.TRUE, Literal: "true"},
		{Type: token.SEMICOLON, Literal: ";"},

		// } else {
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.ELSE, Literal: "else"},
		{Type: token.LBRACE, Literal: "{"},

		// return false;
		{Type: token.RETURN, Literal: "return"},
		{Type: token.FALSE, Literal: "false"},
		{Type: token.SEMICOLON, Literal: ";"},

		// };
		{Type: token.RBRACE, Literal: "}"},

		// // 10 == 10;
		// {Type: token.INT, Literal: "10"},
		// {Type: token.EQ, Literal: "=="},
		// {Type: token.INT, Literal: "10"},
		// {Type: token.SEMICOLON, Literal: ";"},

		// // 10 != 9;
		// {Type: token.INT, Literal: "10"},
		// {Type: token.NOT_EQ, Literal: "!="},
		// {Type: token.INT, Literal: "9"},
		// {Type: token.SEMICOLON, Literal: ";"},
	}

	l := New(input)
	for i, expectedToken := range expected {
		tok := l.NextToken()
		if tok.Type != expectedToken.Type {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, expectedToken.Type, tok.Type)
		}
		if tok.Literal != expectedToken.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, expectedToken.Literal, tok.Literal)
		}
	}
}
