package utils

import "fmt"

type ErrorKind int

const (
	UnknownError ErrorKind = iota
	UnexpectedToken
	UnexpectedEOF
	InvalidSyntax
	UnterminatedString
	InvalidNumber
	UnexpectedSemicolon // New error kind for semicolons
	UndefinedReference // New error kind for undefined references
)

type ParseError struct {
	Kind          ErrorKind
	Message       string
	Line          int
	Column        int
	File          string
	Snippet       string // The line of code where the error occurred
	Caret         int    // The column for the caret (if different from Column)
	Fix           string // Suggested fix, if any
	CodemodPrompt string // Prompt for codemod, if any
	SpecReference string // Reference to the spec section, if any
}

type ParseErrorList struct {
	Errors []ParseError
}

func (l *ParseErrorList) Add(err ParseError) {
	l.Errors = append(l.Errors, err)
}

func (l *ParseErrorList) Len() int {
	return len(l.Errors)
}

func (l *ParseErrorList) ToMessages() []string {
	msgs := make([]string, 0, len(l.Errors))
	for _, err := range l.Errors {
		msgs = append(msgs, ErrorMessage(err))
	}
	return msgs
}

func ErrorMessage(err ParseError) string {
	return FormatErrorWithContext(err)
}

func getErrorType(kind ErrorKind) string {
	switch kind {
	case UnexpectedToken:
		return "SyntaxError"
	case UnexpectedEOF:
		return "SyntaxError"
	case InvalidSyntax:
		return "SyntaxError"
	case UnterminatedString:
		return "SyntaxError"
	case InvalidNumber:
		return "SyntaxError"
	case UnexpectedSemicolon:
		return "SyntaxError"
	case UndefinedReference:
		return "UndefinedReference"
	default:
		return "Error"
	}
}

func UserFriendlyTokenName(tokenType string, literal string) string {
	switch tokenType {
	case "RPAREN":
		return ")"
	case "LPAREN":
		return "("
	case "LBRACE":
		return "{"
	case "RBRACE":
		return "}"
	case "LBRACKET":
		return "["
	case "RBRACKET":
		return "]"
	case "IDENT":
		return literal
	case "SEMICOLON":
		return ";"
	case "COMMA":
		return ","
	case "DOT":
		return "."
	case "COLON":
		return ":"
	case "ASSIGN":
		return "="
	case "PLUS":
		return "+"
	case "MINUS":
		return "-"
	case "ASTERISK":
		return "*"
	case "SLASH":
		return "/"
	case "PERCENT":
		return "%"
	case "CARET":
		return "^"
	case "BANG":
		return "!"
	case "LT":
		return "<"
	case "GT":
		return ">"
	case "EQ":
		return "=="
	case "NOT_EQ":
		return "!="
	case "LT_EQ":
		return "<="
	case "GT_EQ":
		return ">="
	case "AND":
		return "&&"
	case "OR":
		return "||"
	case "VARARG":
		return "..."
	default:
		return tokenType
	}
}

func FormatTokenError(expected string, got string, gotLiteral string) string {
	expectedFriendly := UserFriendlyTokenName(expected, "")
	gotFriendly := UserFriendlyTokenName(got, gotLiteral)
	return fmt.Sprintf("expected %s, got %s", expectedFriendly, gotFriendly)
}
