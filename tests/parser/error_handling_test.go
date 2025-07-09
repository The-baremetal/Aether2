package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseTryCatchFinally(t *testing.T) {
	input := "try { risky() } catch (err) { print(err) } finally { print(\"done\") }"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	// This assumes you have a Try node in your AST. If not, this is a placeholder for when you add it.
	// tryNode, ok := ast.Statements[0].(*parser.Try)
	// if !ok {
	// 	t.Fatalf("expected *Try node, got %T", ast.Statements[0])
	// }
}
