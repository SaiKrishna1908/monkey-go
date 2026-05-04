package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/saikrishna1908/monkey/lexer"
	"github.com/saikrishna1908/monkey/parser"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		parse(l)
	}
}

// TODO: this is just for testing REPL to check generated AST, revert it back
func parse(l *lexer.Lexer) {
	parser := parser.New(l)
	program := parser.ParseProgram()

	errors := parser.Errors()

	if len(errors) > 0 {
		fmt.Printf("parser has %d errors", len(errors))
		for _, msg := range errors {
			fmt.Printf("parser error: %q", msg)
		}

	}

	fmt.Println(program.String())
}
