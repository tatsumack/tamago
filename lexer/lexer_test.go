package lexer

import (
	"tamago/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
7
773;

21 + 21
43 - 1
21 * 2
84 / 2

x := 1
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "7"},
		{token.INT, "773"},
		{token.SEMICOLON, ";"},

		{token.INT, "21"},
		{token.PLUS, "+"},
		{token.INT, "21"},

		{token.INT, "43"},
		{token.MINUS, "-"},
		{token.INT, "1"},

		{token.INT, "21"},
		{token.ASTERISK, "*"},
		{token.INT, "2"},

		{token.INT, "84"},
		{token.SLASH, "/"},
		{token.INT, "2"},

		{token.IDENT, "x"},
		{token.SHORTDEC, ":="},
		{token.INT, "1"},

		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
