package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseArithmeticExpression(t *testing.T) {
	input := "x = 1 + 2 * 3"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok {
		t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
	}
	call, ok := assign.Value.(*parser.Call)
	if !ok {
		t.Fatalf("expected *Call node for expression, got %T", assign.Value)
	}
	if call.Function == nil {
		t.Errorf("expected function/operator in call, got nil")
	}
}

func TestParseParenthesizedExpression(t *testing.T) {
	input := "x = (1 + 2) * 3"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok {
		t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
	}
	call, ok := assign.Value.(*parser.Call)
	if !ok {
		t.Fatalf("expected *Call node for expression, got %T", assign.Value)
	}
	_ = call
}
