package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseStruct(t *testing.T) {
	input := "struct Point { x y }"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	strct, ok := ast.Statements[0].(*parser.StructDef)
	if !ok {
		t.Fatalf("expected *StructDef node, got %T", ast.Statements[0])
	}
	if strct.Name.Value != "Point" {
		t.Errorf("expected struct name 'Point', got %s", strct.Name.Value)
	}
	if len(strct.Fields) != 2 {
		t.Errorf("expected 2 fields, got %d", len(strct.Fields))
	}
}
