package parser

import (
	"fmt"
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

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		lefValue   interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 * 5", 5, "*", 5},
		{"5 / 5", 5, "/", 5},
	}

	for _, tt := range infixTests {
		t.Run(tt.input, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)
			if len(program.Statements) != 1 {
				t.Fatalf("len(program.Statements) = %v, want %v", len(program.Statements), 1)
			}

			stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("program.Statements[0] isn't *ast.ExpressionStatement. got=%T", program.Statements[0])
			}

			if !testInfixExpression(t, stmt.Expression, tt.lefValue, tt.operator, tt.rightValue) {
				return
			}
		})
	}
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {

	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is %T(%s). want ast.InfixExpression. got=", exp, exp)
		return false
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("exp.Operator is %q. want %s", opExp.Operator, operator)
		return false

	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}

	return true
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	}

	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value is %v, want %v", integ.Value, value)
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral() is %v, want %v", integ.TokenLiteral(), value)
		return false

	}

	return true
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
