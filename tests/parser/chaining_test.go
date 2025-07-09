package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseChaining(t *testing.T) {
	input := "pizza().addCheese().bake()"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	call, ok := ast.Statements[0].(*parser.Call)
	if !ok {
		t.Fatalf("expected *Call node, got %T", ast.Statements[0])
	}
	prop, ok := call.Function.(*parser.PropertyAccess)
	if !ok {
		t.Fatalf("expected PropertyAccess as function, got %T", call.Function)
	}
	if prop.Property.Value != "bake" {
		t.Errorf("expected property 'bake', got %s", prop.Property.Value)
	}
	innerProp, ok := prop.Object.(*parser.Call)
	if !ok {
		t.Fatalf("expected Call as object of PropertyAccess, got %T", prop.Object)
	}
	_ = innerProp // You can further inspect innerProp if needed
}
