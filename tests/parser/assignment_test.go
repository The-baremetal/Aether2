package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseAssignment(t *testing.T) {
	input := "x = 10"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok {
		t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
	}
	if assign.Name.Value != "x" {
		t.Errorf("expected assignment to 'x', got %s", assign.Name.Value)
	}
	if lit, ok := assign.Value.(*parser.Literal); !ok || lit.Value != "10" {
		t.Errorf("expected literal value '10', got %v", assign.Value)
	}
}

func TestParseAssignmentArray(t *testing.T) {
	input := "arr = [1, 2, 3]"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok {
		t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
	}
	if assign.Name.Value != "arr" {
		t.Errorf("expected assignment to 'arr', got %s", assign.Name.Value)
	}
	if _, ok := assign.Value.(*parser.Array); !ok {
		t.Errorf("expected *Array value, got %T", assign.Value)
	}
}
