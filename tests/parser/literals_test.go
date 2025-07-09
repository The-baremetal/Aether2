package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseIntLiteral(t *testing.T) {
	input := "x = 42"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok {
		t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
	}
	lit, ok := assign.Value.(*parser.Literal)
	if !ok || lit.Value != "42" {
		t.Errorf("expected int literal '42', got %v", assign.Value)
	}
}

func TestParseStringLiteral(t *testing.T) {
	input := "msg = \"hello\""
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok {
		t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
	}
	lit, ok := assign.Value.(*parser.Literal)
	if !ok || lit.Value != "hello" {
		t.Errorf("expected string literal 'hello', got %v", assign.Value)
	}
}
