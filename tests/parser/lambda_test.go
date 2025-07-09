package parser_test

import (
	"aether/src/lexer"
	"aether/src/parser"
	"testing"
)

func TestParseLambdaAssignment(t *testing.T) {
	input := "myLambda = { print(\"hi\") }"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	ast := p.Parse()
	assign, ok := ast.Statements[0].(*parser.Assignment)
	if !ok {
		t.Fatalf("expected *Assignment node, got %T", ast.Statements[0])
	}
	if _, ok := assign.Value.(*parser.Block); !ok {
		t.Errorf("expected *Block value, got %T", assign.Value)
	}
}
