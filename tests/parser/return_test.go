package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseReturnLiteral(t *testing.T) {
	input := "return 42"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	ret, ok := ast.Statements[0].(*parser.Return)
	if !ok {
		t.Fatalf("expected *Return node, got %T", ast.Statements[0])
	}
	lit, ok := ret.Value.(*parser.Literal)
	if !ok || lit.Value != "42" {
		t.Errorf("expected literal value '42', got %v", ret.Value)
	}
}

func TestParseReturnExpression(t *testing.T) {
	input := "return x + 1"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	ret, ok := ast.Statements[0].(*parser.Return)
	if !ok {
		t.Fatalf("expected *Return node, got %T", ast.Statements[0])
	}
	if _, ok := ret.Value.(*parser.Call); !ok {
		t.Errorf("expected *Call value, got %T", ret.Value)
	}
}
