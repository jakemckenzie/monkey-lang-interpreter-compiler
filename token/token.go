package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   TokenType = "ILLEGAL"
	EOF       TokenType = "EOF"
	IDENT     TokenType = "IDENT"
	INT       TokenType = "INT"
	ASSIGN    TokenType = "ASSIGN"
	PLUS      TokenType = "PLUS"
	COMMA     TokenType = "COMMA"
	SEMICOLON TokenType = "SEMICOLON"
	LPAREN    TokenType = "LPAREN"
	RPAREN    TokenType = "RPAREN"
	LBRACE    TokenType = "LBRACE"
	RBRACE    TokenType = "RBRACE"
	LET       TokenType = "LET"
	FN        TokenType = "FN"
	TRUE      TokenType = "TRUE"
	FALSE     TokenType = "FALSE"
	IF        TokenType = "IF"
	ELSE      TokenType = "ELSE"
	RETURN    TokenType = "RETURN"
	EQ        TokenType = "=="
	NOT_EQ    TokenType = "!="
	MINUS     TokenType = "-"
	SLASH     TokenType = "/"
	ASTERISK  TokenType = "*"
	LT        TokenType = "<"
	GT        TokenType = ">"
	BANG      TokenType = "!"
)

var keywords = map[string]TokenType{
	"let":    LET,
	"fn":     FN,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
