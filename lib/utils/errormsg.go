package utils

import (
	"fmt"
	"strings"
)

func UserErrorMessage(err ParseError) string {
	var b strings.Builder
	switch err.Kind {
	case UnexpectedToken:
		b.WriteString(fmt.Sprintf("🍕 SyntaxError: Unexpected token at line %d, column %d: %s\n", err.Line, err.Column, err.Message))
	case UnexpectedEOF:
		b.WriteString(fmt.Sprintf("🍕 SyntaxError: Unexpected end of file at line %d, column %d.\n", err.Line, err.Column))
	case InvalidSyntax:
		b.WriteString(fmt.Sprintf("🍕 SyntaxError: Invalid syntax at line %d, column %d: %s\n", err.Line, err.Column, err.Message))
	case UnterminatedString:
		b.WriteString(fmt.Sprintf("🍕 SyntaxError: Unterminated string at line %d, column %d.\n", err.Line, err.Column))
	case InvalidNumber:
		b.WriteString(fmt.Sprintf("🍕 SyntaxError: Invalid number at line %d, column %d.\n", err.Line, err.Column))
	case UnexpectedSemicolon:
		b.WriteString(fmt.Sprintf("🍕 SyntaxError: Unexpected `;` at line %d\n", err.Line))
	default:
		b.WriteString(fmt.Sprintf("🍕 Error at line %d, column %d: %s\n", err.Line, err.Column, err.Message))
	}
	if err.Snippet != "" {
		b.WriteString("    " + err.Snippet + "\n")
		caretPos := err.Caret
		if caretPos == 0 {
			caretPos = err.Column
		}
		if caretPos > 0 {
			b.WriteString("    " + strings.Repeat(" ", caretPos-1) + "^\n")
		}
	}
	if err.Kind == UnexpectedSemicolon {
		b.WriteString("    Fix: Remove the semicolon(s)\n")
		b.WriteString("    Do you want to apply the codemod to remove the semicolon? (y/n)\n")
	} else {
		if err.Fix != "" {
			b.WriteString("    Fix: " + err.Fix + "\n")
		}
		if err.CodemodPrompt != "" {
			b.WriteString("    " + err.CodemodPrompt + "\n")
		}
	}
	return b.String()
}
