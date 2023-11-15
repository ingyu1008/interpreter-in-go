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

	//Comparators
	EQ  = "=="
	NEQ = "!="
	GT  = ">"
	GE  = ">="
	LT  = "<"
	LE  = "<="

	//Logical
	AND = "&&"
	OR  = "||"

	//Bitwise
	BAND = "&"
	BOR  = "|"
	BXOR = "^"
	BANG = "!"

	//Seperators
	COMMA     = ","
	SEMICOLON = ";"

	//Brackets
	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	//keywords
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"var":    VAR,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
