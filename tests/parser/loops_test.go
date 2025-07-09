package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseRepeat(t *testing.T) {
	input := "repeat(5) { print(\"hi\") }"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	repeat, ok := ast.Statements[0].(*parser.Repeat)
	if !ok {
		t.Fatalf("expected *Repeat node, got %T", ast.Statements[0])
	}
	if repeat.Body == nil {
		t.Errorf("expected non-nil body block")
	}
}

func TestParseWhile(t *testing.T) {
	input := "while x < 10 { x = x + 1 }"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	whileNode, ok := ast.Statements[0].(*parser.While)
	if !ok {
		t.Fatalf("expected *While node, got %T", ast.Statements[0])
	}
	if whileNode.Body == nil {
		t.Errorf("expected non-nil body block")
	}
}

func TestParseFor(t *testing.T) {
	input := "for i, v in [1,2,3] { print(v) }"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	forNode, ok := ast.Statements[0].(*parser.For)
	if !ok {
		t.Fatalf("expected *For node, got %T", ast.Statements[0])
	}
	if forNode.Body == nil {
		t.Errorf("expected non-nil body block")
	}
}
