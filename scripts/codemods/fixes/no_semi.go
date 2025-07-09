package main

import (
	"aether/src/lexer"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: no_semi <source.ae>")
		os.Exit(1)
	}
	filename := os.Args[1]
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	code := string(src)
	l := lexer.NewLexer(code)
	tokens := l.Tokenize()
	var b strings.Builder
	for _, tok := range tokens {
		// Only skip semicolons that are statement tokens
		if tok.Type == lexer.ILLEGAL && tok.Literal == ";" {
			continue
		}
		b.WriteString(tok.Literal)
		// Add a space after certain tokens for readability (optional)
		if tok.Type == lexer.IDENT || tok.Type == lexer.NUMBER || tok.Type == lexer.STRING {
			b.WriteString(" ")
		}
	}
	fmt.Print(b.String())
}
