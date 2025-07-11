package parser_test

import (
  "aether/src/lexer"
  "aether/src/parser"
  "testing"
)

func TestParseDeeplyNestedIfWhileFunc(t *testing.T) {
  input := `func wow {
    while x < 10 {
      if y > 5 {
        z = [1, 2, { a = 3 }]
      }
    }
  }`
  l := lexer.NewLexer(input)
  p := parser.NewParser(l)
  ast := p.Parse()
  if len(ast.Statements) == 0 {
    t.Fatalf("expected at least one statement")
  }
  fn, ok := ast.Statements[0].(*parser.Function)
  if !ok {
    t.Fatalf("expected *Function node, got %T", ast.Statements[0])
  }
  if fn.Body == nil || len(fn.Body.Statements) == 0 {
    t.Fatalf("expected non-empty function body")
  }
}

func TestParseNestedBlocksAndArrays(t *testing.T) {
  input := `x = [{ y = { z = [1, 2, 3] } }]`
  l := lexer.NewLexer(input)
  p := parser.NewParser(l)
  ast := p.Parse()
  if len(ast.Statements) == 0 {
    t.Fatalf("expected at least one statement")
  }
  assign, ok := ast.Statements[0].(*parser.Assignment)
  _ = assign
  if !ok {
    t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
  }
} 