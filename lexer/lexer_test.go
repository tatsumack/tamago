package lexer

import (
	"tamago/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
7
773
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "7"},
		{token.INT, "773"},

		{token.EOF, ""},
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
