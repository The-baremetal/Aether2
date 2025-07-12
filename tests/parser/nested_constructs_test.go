package parser_test

import (
  "aether/src/lexer"
  "aether/src/parser"
  "fmt"
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
  
  fmt.Printf("🍕 AST Statements count: %d\n", len(ast.Statements))
  for i, stmt := range ast.Statements {
    fmt.Printf("🍕 Statement %d: %T - %+v\n", i, stmt, stmt)
  }
  
  if len(ast.Statements) == 0 {
    t.Fatalf("expected at least one statement")
  }
  fn, ok := ast.Statements[0].(*parser.Function)
  if !ok {
    t.Fatalf("expected *Function node, got %T", ast.Statements[0])
  }
  fmt.Printf("🍕 Function: %+v\n", fn)
  fmt.Printf("🍕 Function Body: %+v\n", fn.Body)
  if fn.Body != nil {
    fmt.Printf("🍕 Function Body Statements count: %d\n", len(fn.Body.Statements))
    for i, stmt := range fn.Body.Statements {
      fmt.Printf("🍕 Body Statement %d: %T - %+v\n", i, stmt, stmt)
    }
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