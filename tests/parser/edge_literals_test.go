package parser_test

import (
  "aether/src/lexer"
  "aether/src/parser"
  "testing"
)

func TestParseEmptyStringLiteral(t *testing.T) {
  input := "x = \"\""
  l := lexer.NewLexer(input)
  p := parser.NewParser(l)
  ast := p.Parse()
  assign, ok := ast.Statements[0].(*parser.Assignment)
  _ = assign
  if !ok {
    t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
  }
}

func TestParseHugeNumberLiteral(t *testing.T) {
  input := "x = 12345678901234567890"
  l := lexer.NewLexer(input)
  p := parser.NewParser(l)
  ast := p.Parse()
  assign, ok := ast.Statements[0].(*parser.Assignment)
  _ = assign
  if !ok {
    t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
  }
}

func TestParseUnicodeStringLiteral(t *testing.T) {
  input := "x = \"こんにちは\""
  l := lexer.NewLexer(input)
  p := parser.NewParser(l)
  ast := p.Parse()
  assign, ok := ast.Statements[0].(*parser.Assignment)
  _ = assign
  if !ok {
    t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
  }
}

func TestParseEscapeSequenceStringLiteral(t *testing.T) {
  input := "x = \"line1\\nline2\""
  l := lexer.NewLexer(input)
  p := parser.NewParser(l)
  ast := p.Parse()
  assign, ok := ast.Statements[0].(*parser.Assignment)
  _ = assign
  if !ok {
    t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
  }
} 