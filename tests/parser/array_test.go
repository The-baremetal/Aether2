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
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok {
		t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
	}
	arr, ok := assign.Value.(*parser.Array)
	if !ok {
		t.Fatalf("expected *Array value, got %T", assign.Value)
	}
	if len(arr.Elements) != 5 {
		t.Errorf("expected 5 elements, got %d", len(arr.Elements))
	}
}

func TestParseNestedArray(t *testing.T) {
	input := "matrix = [[1, 2], [3, 4]]"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok {
		t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
	}
	arr, ok := assign.Value.(*parser.Array)
	if !ok {
		t.Fatalf("expected *Array value, got %T", assign.Value)
	}
	if len(arr.Elements) != 2 {
		t.Errorf("expected 2 elements, got %d", len(arr.Elements))
	}
}
