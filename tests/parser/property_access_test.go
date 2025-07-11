package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"encoding/json"
	"os"
	"testing"
)

func TestParsePropertyAccess(t *testing.T) {
	input := `users = [
  { name: "bob", age: 20 },
  { name: "alice", age: 22 }
]
print(users[0].name)`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if p.Errors.Len() > 0 {
		t.Fatalf("parser errors: %+v", p.Errors.ToMessages())
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(ast)
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	if len(ast.Statements) < 2 {
		t.Fatalf("expected at least 2 statements, got %d", len(ast.Statements))
	}
	callNode, ok := ast.Statements[1].(*parser.Call)
	if !ok {
		t.Fatalf("expected *Call node, got %T", ast.Statements[1])
	}
	if callNode.Function == nil {
		t.Errorf("expected non-nil function expression")
	}
}

func TestParseNestedPropertyAccess(t *testing.T) {
	input := `matrix = [
  [1, 2],
  [3, 4]
]
print(matrix[1][0])`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if p.Errors.Len() > 0 {
		t.Fatalf("parser errors: %+v", p.Errors.ToMessages())
	}
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	if len(ast.Statements) < 2 {
		t.Fatalf("expected at least 2 statements, got %d", len(ast.Statements))
	}
	callNode, ok := ast.Statements[1].(*parser.Call)
	if !ok {
		t.Fatalf("expected *Call node, got %T", ast.Statements[1])
	}
	if callNode.Function == nil {
		t.Errorf("expected non-nil function expression")
	}
}

func TestParseStructPropertyAccess(t *testing.T) {
	input := `struct Point {
  x: int
  y: int
}
p = Point { x: 5, y: 10 }
print(p.x)
print(p.y)`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	if len(ast.Statements) < 4 {
		t.Fatalf("expected at least 4 statements, got %d", len(ast.Statements))
	}
}

func TestParseChainedPropertyAccess(t *testing.T) {
	input := `data = {
  user: {
    profile: {
      name: "john",
      age: 30
    }
  }
}
print(data.user.profile.name)`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 {
		t.Fatalf("expected at least one statement")
	}
	if len(ast.Statements) < 2 {
		t.Fatalf("expected at least 2 statements, got %d", len(ast.Statements))
	}
	callNode, ok := ast.Statements[1].(*parser.Call)
	if !ok {
		t.Fatalf("expected *Call node, got %T", ast.Statements[1])
	}
	if callNode.Function == nil {
		t.Errorf("expected non-nil function expression")
	}
}

func TestParsePropertyAssignment(t *testing.T) {
	input := `point = { x: 0, y: 0 }
point.x = 10
point.y = 20`
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

func TestParseArrayPropertyAccess(t *testing.T) {
	input := `numbers = [1, 2, 3, 4, 5]
print(numbers[2])
print(numbers.length)`
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