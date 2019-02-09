package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"tamago/compiler"
	"tamago/lexer"
	"tamago/parser"
)

func main() {
	start(os.Stdin, os.Stdout)
}

func start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) > 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		c := compiler.New()
		val, err := c.Compile(program)
		if err != nil {
			io.WriteString(out, err.Error())
			io.WriteString(out, "\n")
			continue
		}
		io.WriteString(out, val)
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "parse errors:\n")
	for _, msg := range errors {
		io.WriteString(out, fmt.Sprintf("\t%s\n", msg))
	}
}