package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseStructWithTypes(t *testing.T) {
	input := `struct User {
  name: string
  age: int
}`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	structDef, ok := ast.Statements[0].(*parser.StructDef)
	if !ok {
		t.Fatalf("expected *StructDef node, got %T", ast.Statements[0])
	}
	if structDef.Name.Value != "User" {
		t.Errorf("expected struct name 'User', got %s", structDef.Name.Value)
	}
	if len(structDef.Fields) != 2 {
		t.Errorf("expected 2 fields, got %d", len(structDef.Fields))
	}
}

func TestParseFunctionWithTypeAnnotations(t *testing.T) {
	input := `func agePlusOne(user) {
  return user.age + 1
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
	if fn.Name.Value != "agePlusOne" {
		t.Errorf("expected function name 'agePlusOne', got %s", fn.Name.Value)
	}
	if len(fn.Params) != 1 {
		t.Errorf("expected 1 parameter, got %d", len(fn.Params))
	}
}

func TestParseFunctionWithParamTypes(t *testing.T) {
	input := `func add(a: int, b: int) {
  return a + b
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
	if fn.Name.Value != "add" {
		t.Errorf("expected function name 'add', got %s", fn.Name.Value)
	}
	if len(fn.Params) != 2 {
		t.Errorf("expected 2 parameters, got %d", len(fn.Params))
	}
}

func TestParseStructWithComplexTypes(t *testing.T) {
	input := `struct Point {
  x: int
  y: int
  name: string
  active: bool
}`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	structDef, ok := ast.Statements[0].(*parser.StructDef)
	if !ok {
		t.Fatalf("expected *StructDef node, got %T", ast.Statements[0])
	}
	if structDef.Name.Value != "Point" {
		t.Errorf("expected struct name 'Point', got %s", structDef.Name.Value)
	}
	if len(structDef.Fields) != 4 {
		t.Errorf("expected 4 fields, got %d", len(structDef.Fields))
	}
}

func TestParseFunctionWithReturnType(t *testing.T) {
	input := `func calculate(x: int, y: int): int {
  return x * y + 10
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
	if fn.Name.Value != "calculate" {
		t.Errorf("expected function name 'calculate', got %s", fn.Name.Value)
	}
	if len(fn.Params) != 2 {
		t.Errorf("expected 2 parameters, got %d", len(fn.Params))
	}
} 