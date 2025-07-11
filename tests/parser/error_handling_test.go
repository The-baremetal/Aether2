package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseExplicitErrorHandling(t *testing.T) {
	input := `
result, err = risky_stuff()
if err != nil {
  print("oops!", err)
}
`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	// Check that the first statement is an assignment
	if _, ok := ast.Statements[0].(*parser.Assignment); !ok {
		t.Fatalf("expected first statement to be an assignment, got %T", ast.Statements[0])
	}
	// Check that the second statement is an call statement
	if _, ok := ast.Statements[1].(*parser.Call); !ok {
		t.Fatalf("expected second statement to be an call statement, got %T", ast.Statements[1])
	}
}
