package main

import (
	"aether/src/lexer"
	"aether/src/parser"
	"fmt"
)

func main() {
	input := "a, b = 1, 2"
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	
	fmt.Println("=== Parser Token Advancement ===")
	fmt.Printf("Input: %q\n", input)
	
	fmt.Printf("Initial curToken: %s '%s' (line %d, col %d)\n", p.curToken.Type, p.curToken.Literal, p.curToken.Line, p.curToken.Column)
	fmt.Printf("Initial peekToken: %s '%s' (line %d, col %d)\n", p.peekToken.Type, p.peekToken.Literal, p.peekToken.Line, p.peekToken.Column)
	
	for i := 0; i < 5; i++ {
		fmt.Printf("\n--- Call %d ---\n", i+1)
		p.nextToken()
		fmt.Printf("After nextToken():\n")
		fmt.Printf("  curToken: %s '%s' (line %d, col %d)\n", p.curToken.Type, p.curToken.Literal, p.curToken.Line, p.curToken.Column)
		fmt.Printf("  peekToken: %s '%s' (line %d, col %d)\n", p.peekToken.Type, p.peekToken.Literal, p.peekToken.Line, p.peekToken.Column)
	}
} 