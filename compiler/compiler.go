package compiler

import (
	"fmt"
	llvm "github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"tamago/ast"
)

type Compiler struct {
	Module       *llvm.Module
	CurrentFunc  *llvm.Func
	CurrentBlock *llvm.Block
}

func New() *Compiler {
	m := llvm.NewModule()
	return &Compiler{
		Module:      m,
		CurrentFunc: nil,
	}
}

func (c *Compiler) Compile(node ast.Node) (string, error) {
	program, ok := node.(*ast.Program)
	if !ok {
		return "", fmt.Errorf("node must be *ast.Program")
	}

	if err := c.compileProgram(program); err != nil {
		return "", err
	}

	return c.Module.String(), nil
}

func (c *Compiler) compileProgram(node *ast.Program) error {
	f := c.Module.NewFunc("main", types.I64)
	b := f.NewBlock("entry")

	c.CurrentFunc = f
	c.CurrentBlock = b

	var retVal value.Value
	for _, stmt := range node.Statements {
		var err error
		retVal, err = c.compileStatement(stmt)
		if err != nil {
			return err
		}
	}

	b.NewRet(retVal)

	return nil
}

func (c *Compiler) compileStatement(node ast.Statement) (value.Value, error) {
	switch node := node.(type) {
	case *ast.ExpressionStatement:
		return c.compileExpression(node.Expression)
	}
	return nil, nil
}

func (c *Compiler) compileExpression(node ast.Expression) (value.Value, error) {
	switch node := node.(type) {
	case *ast.InfixExpression:
		left, err := c.compileExpression(node.Left)
		if err != nil {
			return nil, err
		}

		right, err := c.compileExpression(node.Right)
		if err != nil {
			return nil, err
		}

		return c.compileInfixExpression(node, left, right)
	case *ast.IntegerLiteral:
		return c.compileIntegerLiteral(node)
	}
	return nil, nil
}

func (c *Compiler) compileInfixExpression(node *ast.InfixExpression, left value.Value, right value.Value) (value.Value, error) {
	switch (node.Operator) {
	case "+":
		return c.CurrentBlock.NewAdd(left, right), nil
	case "-":
		return c.CurrentBlock.NewSub(left, right), nil
	case "*":
		return c.CurrentBlock.NewMul(left, right), nil
	case "/":
		return c.CurrentBlock.NewSDiv(left, right), nil
	}
	return nil, fmt.Errorf("invalid operator %s", node.Operator)
}

func (c *Compiler) compileIntegerLiteral(node *ast.IntegerLiteral) (value.Value, error) {
	return constant.NewInt(types.I64, node.Value), nil
}
