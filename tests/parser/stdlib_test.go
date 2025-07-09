package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseStdlibPrint(t *testing.T) {
	input := "print(42)"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	call, ok := ast.Statements[0].(*parser.Call)
	if !ok {
		t.Fatalf("expected *Call node, got %T", ast.Statements[0])
	}
	id, ok := call.Function.(*parser.Identifier)
	if !ok || id.Value != "print" {
		t.Errorf("expected function 'print', got %v", call.Function)
	}
}

func TestParseStdlibSqrt(t *testing.T) {
	input := "sqrt(16)"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	call, ok := ast.Statements[0].(*parser.Call)
	if !ok {
		t.Fatalf("expected *Call node, got %T", ast.Statements[0])
	}
	id, ok := call.Function.(*parser.Identifier)
	if !ok || id.Value != "sqrt" {
		t.Errorf("expected function 'sqrt', got %v", call.Function)
	}
}
