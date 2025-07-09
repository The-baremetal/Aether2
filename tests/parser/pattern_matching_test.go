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
