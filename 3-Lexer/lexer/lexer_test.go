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

func TestMoreTokenTypes(t *testing.T) {
	input := `var f = !(true)
	var illegal = ?
	
	fn eq(x, y){
		if((x > y) || (x < y)) {
			return false
		}
		return true
	}

	fn eq2(x, y){
		if((x >= y) && (x <= y)) {
			return true
		}
		return false
	}

	var a = true != false

	var b = ((6 & 4) | 11)
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENT, "f"},
		{token.ASSIGN, "="},
		{token.BANG, "!"},
		{token.LPAREN, "("},
		{token.TRUE, "true"},
		{token.RPAREN, ")"},
		{token.VAR, "var"},
		{token.IDENT, "illegal"},
		{token.ASSIGN, "="},
		{token.ILLEGAL, "?"},

		{token.FUNCTION, "fn"},
		{token.IDENT, "eq"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.GT, ">"},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.OR, "||"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.LT, "<"},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.RBRACE, "}"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.RBRACE, "}"},

		{token.FUNCTION, "fn"},
		{token.IDENT, "eq2"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.GE, ">="},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.AND, "&&"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.LE, "<="},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.RBRACE, "}"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.RBRACE, "}"},

		{token.VAR, "var"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.TRUE, "true"},
		{token.NEQ, "!="},
		{token.FALSE, "false"},

		{token.VAR, "var"},
		{token.IDENT, "b"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.LPAREN, "("},
		{token.INT, "6"},
		{token.BAND, "&"},
		{token.INT, "4"},
		{token.RPAREN, ")"},
		{token.BOR, "|"},
		{token.INT, "11"},
		{token.RPAREN, ")"},
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
