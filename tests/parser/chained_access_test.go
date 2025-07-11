package parser_test

import (
  "aether/src/lexer"
  "aether/src/parser"
  "testing"
)

func TestParseChainedPropertyAccessAndCalls(t *testing.T) {
  input := "foo.bar().baz[0](x)"
  l := lexer.NewLexer(input)
  p := parser.NewParser(l)
  ast := p.Parse()
  if len(ast.Statements) == 0 {
    t.Fatalf("expected at least one statement")
  }
  call, ok := ast.Statements[0].(*parser.Call)
  _ = call
  if !ok {
    t.Fatalf("expected *Call node, got %T", ast.Statements[0])
  }
} 