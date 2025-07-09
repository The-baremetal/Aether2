package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseImport(t *testing.T) {
	input := "import math"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	imp, ok := ast.Statements[0].(*parser.Import)
	if !ok {
		t.Fatalf("expected *Import node, got %T", ast.Statements[0])
	}
	if imp.Name.Value != "math" {
		t.Errorf("expected import name 'math', got %s", imp.Name.Value)
	}
}

func TestParseImportAs(t *testing.T) {
	input := "import math as m"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	imp, ok := ast.Statements[0].(*parser.Import)
	if !ok {
		t.Fatalf("expected *Import node, got %T", ast.Statements[0])
	}
	if imp.As == nil || imp.As.Value != "m" {
		t.Errorf("expected import as 'm', got %v", imp.As)
	}
}
