package lexer

import (
	"interpreter/3-Lexer/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `var a = 1
	var b = true
	
	fn gcd(x, y){
		if(x == 0){
			return y
		} else {
			return gcd(y%x, x)
		}
	}

	var k = ((a + b) - a) * (3 / 1);
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
		{token.VAR, "var"},
		{token.IDENT, "b"},
		{token.ASSIGN, "="},
		{token.TRUE, "true"},
		{token.FUNCTION, "fn"},
		{token.IDENT, "gcd"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.EQ, "=="},
		{token.INT, "0"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "y"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "gcd"},
		{token.LPAREN, "("},
		{token.IDENT, "y"},
		{token.PERC, "%"},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.VAR, "var"},
		{token.IDENT, "k"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.PLUS, "+"},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.MINUS, "-"},
		{token.IDENT, "a"},
		{token.RPAREN, ")"},
		{token.AST, "*"},
		{token.LPAREN, "("},
		{token.INT, "3"},
		{token.SLASH, "/"},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
