package parser_test

import (
  "aether/src/lexer"
  "aether/src/parser"
  "testing"
)

func TestParseMissingBrace(t *testing.T) {
  input := "if x > 0 { print(1) " // missing closing brace
  l := lexer.NewLexer(input)
  p := parser.NewParser(l)
  _ = p.Parse()
  if p.Errors.Len() == 0 {
    t.Errorf("expected parser errors for missing brace")
  }
}

func TestParseGarbageInput(t *testing.T) {
  input := "@#%$^&*()"
  l := lexer.NewLexer(input)
  p := parser.NewParser(l)
  _ = p.Parse()
  if p.Errors.Len() == 0 {
    t.Errorf("expected parser errors for garbage input")
  }
} 