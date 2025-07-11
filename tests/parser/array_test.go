package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseArrayLiteral(t *testing.T) {
	input := "nums = [1, 2, 3, 4, 5]"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if p.Errors.Len() > 0 {
		t.Fatalf("parser errors: %+v", p.Errors.ToMessages())
	}
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
}
