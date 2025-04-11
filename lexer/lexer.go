package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{
				Type:    token.EQ,
				Literal: string(ch) + string(l.ch),
			}
			l.readChar()
		} else {
			tok = newToken(token.ASSIGN, l.ch)
			l.readChar()
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: string(ch) + string(l.ch),
			}
			l.readChar()
		} else {
			tok = newToken(token.BANG, l.ch)
			l.readChar()
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
		l.readChar()
	case '-':
		tok = newToken(token.MINUS, l.ch)
		l.readChar()
	case '/':
		tok = newToken(token.SLASH, l.ch)
		l.readChar()
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
		l.readChar()
	case '<':
		tok = newToken(token.LT, l.ch)
		l.readChar()
	case '>':
		tok = newToken(token.GT, l.ch)
		l.readChar()
	case ',':
		tok = newToken(token.COMMA, l.ch)
		l.readChar()
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
		l.readChar()
	case '(':
		tok = newToken(token.LPAREN, l.ch)
		l.readChar()
	case ')':
		tok = newToken(token.RPAREN, l.ch)
		l.readChar()
	case '{':
		tok = newToken(token.LBRACE, l.ch)
		l.readChar()
	case '}':
		tok = newToken(token.RBRACE, l.ch)
		l.readChar()
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tokType := token.LookupIdent(literal)
			return token.Token{Type: tokType, Literal: literal} // Position advanced by readIdentifier
		} else if isDigit(l.ch) {
			return token.Token{Type: token.INT, Literal: l.readNumber()} // Position advanced by readNumber
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
			l.readChar() // Move past the illegal character
		}
	}
	return tok
}
