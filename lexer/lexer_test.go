package lexer

import (
	"testing"

	"github.com/smadeja/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	expected := []token.Token{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, exp := range expected {
		act := l.NextToken()

		if act != exp {
			t.Fatalf("wrong token at index %d: got %#v, expected %#v", i, act, exp)
		}
	}
}
