package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseFunc(t *testing.T) {
	input := "func greet { print(\"hi\") }"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	fn, ok := ast.Statements[0].(*parser.Function)
	if !ok {
		t.Fatalf("expected *Function node, got %T", ast.Statements[0])
	}
	if fn.Name.Value != "greet" {
		t.Errorf("expected function name 'greet', got %s", fn.Name.Value)
	}
}

func TestParseFuncWithParams(t *testing.T) {
	input := "func add(a, b) { return a + b }"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	fn, ok := ast.Statements[0].(*parser.Function)
	if !ok {
		t.Fatalf("expected *Function node, got %T", ast.Statements[0])
	}
	if len(fn.Params) != 2 {
		t.Errorf("expected 2 params, got %d", len(fn.Params))
	}
}
