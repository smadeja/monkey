package lexer

import (
	"github.com/smadeja/monkey/token"
)

type Lexer struct {
	input        string
	readPosition int
	char         byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.char {
	case '=':
		tok = token.New(token.ASSIGN, string(l.char))
	case '+':
		tok = token.New(token.PLUS, string(l.char))
	case '(':
		tok = token.New(token.LPAREN, string(l.char))
	case ')':
		tok = token.New(token.RPAREN, string(l.char))
	case '{':
		tok = token.New(token.LBRACE, string(l.char))
	case '}':
		tok = token.New(token.RBRACE, string(l.char))
	case ',':
		tok = token.New(token.COMMA, string(l.char))
	case ';':
		tok = token.New(token.SEMICOLON, string(l.char))
	case 0:
		tok = token.New(token.EOF, "")
	}

	l.readChar()
	return tok
}
