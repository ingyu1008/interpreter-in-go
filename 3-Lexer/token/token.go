package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // x, y, z, add, mult
	INT   = "INT"   // 1, 2, 3, ...

	//Arithmatic Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	AST    = "*"
	SLASH  = "/"
	PERC   = "%"

	//Seperators
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
