package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseBreakInForLoop(t *testing.T) {
	input := `for i in [1, 2, 3, 4, 5, 6, 7, 8] {
  if i == 5 {
    break
  }
  print(i)
}`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	forNode, ok := ast.Statements[0].(*parser.For)
	if !ok {
		t.Fatalf("expected *For node, got %T", ast.Statements[0])
	}
	if forNode.Body == nil {
		t.Errorf("expected non-nil body block")
	}
}

func TestParseContinueInForLoop(t *testing.T) {
	input := `for i in [1, 2, 3, 4, 5, 6, 7, 8] {
  if i % 2 == 0 {
    continue
  }
  print(i)
}`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	forNode, ok := ast.Statements[0].(*parser.For)
	if !ok {
		t.Fatalf("expected *For node, got %T", ast.Statements[0])
	}
	if forNode.Body == nil {
		t.Errorf("expected non-nil body block")
	}
}

func TestParseBreakInWhileLoop(t *testing.T) {
	input := `while x < 10 {
  if x == 5 {
    break
  }
  x = x + 1
}`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	whileNode, ok := ast.Statements[0].(*parser.While)
	if !ok {
		t.Fatalf("expected *While node, got %T", ast.Statements[0])
	}
	if whileNode.Body == nil {
		t.Errorf("expected non-nil body block")
	}
}

func TestParseContinueInWhileLoop(t *testing.T) {
	input := `while x < 10 {
  if x % 2 == 0 {
    x = x + 1
    continue
  }
  print(x)
  x = x + 1
}`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	whileNode, ok := ast.Statements[0].(*parser.While)
	if !ok {
		t.Fatalf("expected *While node, got %T", ast.Statements[0])
	}
	if whileNode.Body == nil {
		t.Errorf("expected non-nil body block")
	}
}

func TestParseBreakInRepeatLoop(t *testing.T) {
	input := `repeat(10) {
  if counter == 5 {
    break
  }
  counter = counter + 1
}`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	repeatNode, ok := ast.Statements[0].(*parser.Repeat)
	if !ok {
		t.Fatalf("expected *Repeat node, got %T", ast.Statements[0])
	}
	if repeatNode.Body == nil {
		t.Errorf("expected non-nil body block")
	}
}

func TestParseBreakContinueCombined(t *testing.T) {
	input := `for i in [1, 2, 3, 4, 5, 6, 7, 8] {
  if i == 5 {
    break
  }
  if i % 2 == 0 {
    continue
  }
  print(i)
}`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	forNode, ok := ast.Statements[0].(*parser.For)
	if !ok {
		t.Fatalf("expected *For node, got %T", ast.Statements[0])
	}
	if forNode.Body == nil {
		t.Errorf("expected non-nil body block")
	}
} 