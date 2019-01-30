package parser

import (
	"tamago/ast"
	"tamago/lexer"
	"testing"
)

func TestIntegerLiteralExpression(t *testing.T) {
	input := `773;`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	t.Run("TestIntegerLiteralExpression::checkParseError", func(t *testing.T) {
		checkParserErrors(t, p)
	})

	if len(program.Statements) != 1 {
		println(program.Statements[0].String())
		//println(program.Statements[1].String())
		t.Fatalf("len(program.Statements) = %v, want %v", len(program.Statements), 1)
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] isn't *ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	integer, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp isn't *ast.IntegerLiteral. got=%T", stmt.Expression)
	}

	if integer.Value != 773 {
		t.Errorf("integer.Value is %v, want %v", integer.Value, 773)
	}
	if integer.TokenLiteral() != "773" {
		t.Errorf("integer.TokenLiteral is %v, want %v", integer.TokenLiteral(), "773")
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
