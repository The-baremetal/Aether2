package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseStatements(t *testing.T) {
	input := `x = 1
print(2)
{ y = 3 }`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) != 3 {
		t.Fatalf("expected 3 statements, got %d", len(ast.Statements))
	}
	if _, ok := ast.Statements[0].(*parser.Assignment); !ok {
		t.Errorf("expected first statement to be assignment, got %T", ast.Statements[0])
	}
	if _, ok := ast.Statements[1].(*parser.Call); !ok {
		t.Errorf("expected second statement to be call, got %T", ast.Statements[1])
	}
	if _, ok := ast.Statements[2].(*parser.Block); !ok {
		t.Errorf("expected third statement to be block, got %T", ast.Statements[2])
	}
}
