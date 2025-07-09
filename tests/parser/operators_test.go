package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseMathOperator(t *testing.T) {
	input := "x = 1 + 2"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok {
		t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
	}
	call, ok := assign.Value.(*parser.Call)
	if !ok {
		t.Fatalf("expected *Call node for operator, got %T", assign.Value)
	}
	_ = call // You can further inspect the call if needed
}

func TestParseComparisonOperator(t *testing.T) {
	input := "x = a == b"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok {
		t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
	}
	call, ok := assign.Value.(*parser.Call)
	if !ok {
		t.Fatalf("expected *Call node for operator, got %T", assign.Value)
	}
	_ = call // You can further inspect the call if needed
}
