package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseIf(t *testing.T) {
	input := "if x > 0 { print(1) }"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	ifNode, ok := ast.Statements[0].(*parser.If)
	if !ok {
		t.Fatalf("expected *If node, got %T", ast.Statements[0])
	}
	if ifNode.Consequence == nil {
		t.Errorf("expected non-nil consequence block")
	}
}

func TestParseIfElse(t *testing.T) {
	input := "if x > 0 { print(1) } else { print(0) }"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	ifNode, ok := ast.Statements[0].(*parser.If)
	if !ok {
		t.Fatalf("expected *If node, got %T", ast.Statements[0])
	}
	if ifNode.Alternative == nil {
		t.Errorf("expected non-nil alternative block")
	}
}
