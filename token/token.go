package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"
	IDENT   TokenType = "IDENT"
	INT     TokenType = "INT"
	LET     TokenType = "LET"
	FN      TokenType = "FUNCTION"
	TRUE    TokenType = "TRUE"
	FALSE   TokenType = "FALSE"
	IF      TokenType = "IF"
	ELSE    TokenType = "ELSE"
	RETURN  TokenType = "RETURN"
	EQ      TokenType = "EQUALS"
	NOT_EQ  TokenType = "NOTEQUALS"

	MINUS     TokenType = "-"
	SLASH     TokenType = "/"
	ASTERISK  TokenType = "*"
	LT        TokenType = "<"
	GT        TokenType = ">"
	BANG      TokenType = "!"
	ASSIGN    TokenType = "="
	PLUS      TokenType = "+"
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"
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
