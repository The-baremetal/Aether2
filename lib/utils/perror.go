package utils

type ErrorKind int

const (
	UnknownError ErrorKind = iota
	UnexpectedToken
	UnexpectedEOF
	InvalidSyntax
	UnterminatedString
	InvalidNumber
	UnexpectedSemicolon // New error kind for semicolons
)

type ParseError struct {
	Kind          ErrorKind
	Message       string
	Line          int
	Column        int
	Snippet       string // The line of code where the error occurred
	Caret         int    // The column for the caret (if different from Column)
	Fix           string // Suggested fix, if any
	CodemodPrompt string // Prompt for codemod, if any
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
	return UserErrorMessage(err)
}
