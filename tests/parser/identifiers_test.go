package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseIdentifier(t *testing.T) {
	input := "foo = 1"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok {
		t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
	}
	if assign.Name.Value != "foo" {
		t.Errorf("expected identifier 'foo', got %s", assign.Name.Value)
	}
}

func TestParseIdentifierInExpression(t *testing.T) {
	input := "x = foo + 2"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok {
		t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
	}
	// The left argument of the addition should be an identifier 'foo'
	call, ok := assign.Value.(*parser.Call)
	if !ok {
		t.Fatalf("expected *Call node for addition, got %T", assign.Value)
	}
	id, ok := call.Args[0].(*parser.Identifier)
	if !ok || id.Value != "foo" {
		t.Errorf("expected identifier 'foo' as left arg, got %v", call.Args[0])
	}
}
