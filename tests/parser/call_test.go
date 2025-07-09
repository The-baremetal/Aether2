package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseCall(t *testing.T) {
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
	if len(call.Args) != 1 {
		t.Errorf("expected 1 argument, got %d", len(call.Args))
	}
}

func TestParseCallMultipleArgs(t *testing.T) {
	input := "add(1, 2, 3)"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	call, ok := ast.Statements[0].(*parser.Call)
	if !ok {
		t.Fatalf("expected *Call node, got %T", ast.Statements[0])
	}
	if len(call.Args) != 3 {
		t.Errorf("expected 3 arguments, got %d", len(call.Args))
	}
}
