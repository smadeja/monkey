package token

type TokenType string

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	ASSIGN    = "="
	PLUS      = "+"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	COMMA     = ","
	SEMICOLON = ";"
)

type Token struct {
	Type    TokenType
	Literal string
}
