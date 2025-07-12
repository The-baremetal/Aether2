package parser_test

import (
  "aether/src/lexer"
  "aether/src/parser"
  "fmt"
  "testing"
)

func TestParseTupleDestructuring(t *testing.T) {
  input := "a, b = 1, 2"
  l := lexer.NewLexer(input)
  
  fmt.Println("=== LEXER OUTPUT ===")
  for {
    tok := l.NextToken()
    fmt.Printf("Token: %+v\n", tok)
    if tok.Type == lexer.EOF {
      break
    }
  }
  
  l = lexer.NewLexer(input)
  p := parser.NewParser(l)
  ast := p.Parse()
  if p.Errors.Len() > 0 {
		t.Fatalf("parser errors: %+v", p.Errors.ToMessages())
	}
  assign, ok := ast.Statements[0].(*parser.Assignment)
  if !ok {
    t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
  }
  if len(assign.Names) != 2 {
    t.Errorf("expected 2 names in assignment, got %d", len(assign.Names))
  }
} 