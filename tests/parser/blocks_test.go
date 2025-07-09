package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseBlock(t *testing.T) {
	input := "{ x = 1 y = 2 }"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	block, ok := ast.Statements[0].(*parser.Block)
	if !ok {
		t.Fatalf("expected *Block node, got %T", ast.Statements[0])
	}
	if len(block.Statements) != 2 {
		t.Errorf("expected 2 statements in block, got %d", len(block.Statements))
	}
}

func TestBlockAsExpression(t *testing.T) {
	input := "x = { y = 2 }"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok {
		t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
	}
	if _, ok := assign.Value.(*parser.Block); !ok {
		t.Errorf("expected *Block value, got %T", assign.Value)
	}
}
