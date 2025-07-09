package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseSingleLineComment(t *testing.T) {
	input := "// this is a comment\nx = 1"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(ast.Statements))
	}
}

func TestParseMultiLineComment(t *testing.T) {
	input := "/* multi\nline\ncomment */\ny = 2"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(ast.Statements))
	}
}

func TestParseDocComment(t *testing.T) {
	input := "/// doc comment\nz = 3"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(ast.Statements))
	}
}
