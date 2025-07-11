package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParsePartialApplication(t *testing.T) {
	input := `func bind(funcName, returnType, ...paramTypes) {
  return func(...args) {
    return c_call(funcName, returnType, paramTypes, args)
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
	if fn.Name.Value != "bind" {
		t.Errorf("expected function name 'bind', got %s", fn.Name.Value)
	}
	if len(fn.Params) != 3 {
		t.Errorf("expected 3 parameters, got %d", len(fn.Params))
	}
}

func TestParsePartialApplicationWithVarargs(t *testing.T) {
	input := `func createPrinter(prefix) {
  return func(...args) {
    print(prefix, ...args)
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
	if fn.Name.Value != "createPrinter" {
		t.Errorf("expected function name 'createPrinter', got %s", fn.Name.Value)
	}
	if len(fn.Params) != 1 {
		t.Errorf("expected 1 parameter, got %d", len(fn.Params))
	}
}

func TestParsePartialApplicationComplex(t *testing.T) {
	input := `func curry(fn, ...args) {
  return func(...moreArgs) {
    return fn(...args, ...moreArgs)
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
	if fn.Name.Value != "curry" {
		t.Errorf("expected function name 'curry', got %s", fn.Name.Value)
	}
	if len(fn.Params) != 2 {
		t.Errorf("expected 2 parameters, got %d", len(fn.Params))
	}
}

func TestParsePartialApplicationWithReturn(t *testing.T) {
	input := `func add(a, b) {
  return a + b
}
addFive = func(b) {
  return add(5, b)
}`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	if len(ast.Statements) != 2 {
		t.Fatalf("expected 2 statements, got %d", len(ast.Statements))
	}
}

func TestParsePartialApplicationChained(t *testing.T) {
	input := `func compose(f, g) {
  return func(x) {
    return f(g(x))
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
	if fn.Name.Value != "compose" {
		t.Errorf("expected function name 'compose', got %s", fn.Name.Value)
	}
	if len(fn.Params) != 2 {
		t.Errorf("expected 2 parameters, got %d", len(fn.Params))
	}
} 