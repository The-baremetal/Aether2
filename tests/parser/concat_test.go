package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseConcat(t *testing.T) {
	input := "greeting = [\"h\"] .. [\"i\"]"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	if len(ast.Statements) == 0 || ast.Statements[0] == nil {
		t.Fatalf("expected at least one statement, got none or nil. Parser errors: %+v", p.Errors)
	}
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok {
		t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
	}
	if assign.Value == nil {
		t.Fatalf("expected assignment value, got nil")
	}
	call, ok := assign.Value.(*parser.Call)
	if !ok {
		t.Fatalf("expected *Call node for concat, got %T (value: %#v)", assign.Value, assign.Value)
	}
	if call == nil {
		t.Fatalf("call is nil after type assertion")
	}
	if call.Function == nil {
		t.Fatalf("expected call.Function to be non-nil")
	}
	if call.Args == nil {
		t.Fatalf("expected call.Args to be non-nil")
	}
}
