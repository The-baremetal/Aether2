package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseWhitespace(t *testing.T) {
	input := "  x   =   1  \n\n   y = 2   "
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) != 2 {
		t.Fatalf("expected 2 statements, got %d", len(ast.Statements))
	}
}
