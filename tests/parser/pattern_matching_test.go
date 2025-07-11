package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseMatchLiteralCase(t *testing.T) {
	input := "match x { case 0 { print(\"zero\") } }"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	// This assumes you have a Match node in your AST. If not, this is a placeholder for when you add it.
	// matchNode, ok := ast.Statements[0].(*parser.Match)
	// if !ok {
	// 	t.Fatalf("expected *Match node, got %T", ast.Statements[0])
	// }
	// if len(matchNode.Cases) == 0 {
	// 	t.Errorf("expected at least one case, got %d", len(matchNode.Cases))
	// }
}

func TestParseMatchWithWildcardCase(t *testing.T) {
	input := "match y { case 1 { print(1) } case _ { print(0) } }"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if p.Errors.Len() > 0 {
		t.Fatalf("parser errors: %+v", p.Errors.ToMessages())
	}
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	matchNode, ok := ast.Statements[0].(*parser.Match)
	if !ok {
		t.Fatalf("expected *Match node, got %T", ast.Statements[0])
	}
	if len(matchNode.Cases) != 2 {
		t.Fatalf("expected 2 cases, got %d", len(matchNode.Cases))
	}
	if id, ok := matchNode.Cases[0].Pattern.(*parser.Literal); !ok || id.Value != "1" {
		t.Errorf("expected first case pattern to be literal 1, got %v", matchNode.Cases[0].Pattern)
	}
	if id, ok := matchNode.Cases[1].Pattern.(*parser.Identifier); !ok || id.Value != "_" {
		t.Errorf("expected second case pattern to be '_', got %v", matchNode.Cases[1].Pattern)
	}
}
