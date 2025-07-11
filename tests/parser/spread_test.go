package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseSpreadInFunctionCall(t *testing.T) {
	input := `printAll(...args, "cheese", ...flags, "pepperoni")`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	callNode, ok := ast.Statements[0].(*parser.Call)
	if !ok {
		t.Fatalf("expected *Call node, got %T", ast.Statements[0])
	}
	if callNode.Function == nil {
		t.Errorf("expected non-nil function expression")
	}
}

func TestParseSpreadInArray(t *testing.T) {
	input := `hello = ["h", "e", "l", "l", "o"]
world = ["w", "o", "r", "l", "d"]
greeting = hello .. world`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	if len(ast.Statements) != 3 {
		t.Fatalf("expected 3 statements, got %d", len(ast.Statements))
	}
}

func TestParseSpreadInStringConcatenation(t *testing.T) {
	input := `str1 = "foo"
str2 = "bar"
result = str1 .. str2`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	if len(ast.Statements) != 3 {
		t.Fatalf("expected 3 statements, got %d", len(ast.Statements))
	}
}

func TestParseSpreadInFunctionParameters(t *testing.T) {
	input := `func printAll(...args) {
  repeat(args.length) {
    print(args[_])
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
	if fn.Name.Value != "printAll" {
		t.Errorf("expected function name 'printAll', got %s", fn.Name.Value)
	}
	if len(fn.Params) != 1 {
		t.Errorf("expected 1 parameter, got %d", len(fn.Params))
	}
}

func TestParseSpreadInVarargsPosition(t *testing.T) {
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

func TestParseSpreadInMixedArguments(t *testing.T) {
	input := `func processData(prefix, ...data, suffix) {
  print(prefix, ...data, suffix)
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
	if fn.Name.Value != "processData" {
		t.Errorf("expected function name 'processData', got %s", fn.Name.Value)
	}
	if len(fn.Params) != 3 {
		t.Errorf("expected 3 parameters, got %d", len(fn.Params))
	}
} 