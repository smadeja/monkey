package token

type TokenType string

const (
	IDENT   = "IDENT"
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"

	// Literals
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Separators
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

func New(ttype TokenType, literal string) Token {
	return Token{ttype, literal}
}
