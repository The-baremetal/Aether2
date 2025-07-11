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
	if p.Errors.Len() > 0 {
		t.Fatalf("parser errors: %+v", p.Errors.ToMessages())
	}
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok || assign == nil {
		t.Fatalf("expected non-nil *Assignment node, got %T", ast.Statements[0])
	}
	if len(assign.Names) != 1 || assign.Names[0].Value != "x" {
		t.Errorf("expected assignment to 'x', got %v", assign.Names)
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
	if p.Errors.Len() > 0 {
		t.Fatalf("parser errors: %+v", p.Errors.ToMessages())
	}
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok || assign == nil {
		t.Fatalf("expected non-nil *Assignment node, got %T", ast.Statements[0])
	}
	if len(assign.Names) != 1 || assign.Names[0].Value != "arr" {
		t.Errorf("expected assignment to 'arr', got %v", assign.Names)
	}
	if _, ok := assign.Value.(*parser.Array); !ok {
		t.Errorf("expected *Array value, got %T", assign.Value)
	}
}
