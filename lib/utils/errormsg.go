package utils

import (
	"fmt"
	"sort"
	"strings"
)

type ErrorGroup struct {
	File   string
	Errors []ParseError
}

type ErrorSummary struct {
	TotalFiles    int
	TotalErrors   int
	FilesAffected []string
	Groups        []ErrorGroup
}

func GroupErrorsByFile(errors []ParseError) *ErrorSummary {
	fileGroups := make(map[string][]ParseError)
	filesAffected := make(map[string]bool)

	for _, err := range errors {
		fileGroups[err.File] = append(fileGroups[err.File], err)
		filesAffected[err.File] = true
	}

	var groups []ErrorGroup
	var files []string
	for file, fileErrors := range fileGroups {
		groups = append(groups, ErrorGroup{
			File:   file,
			Errors: DeduplicateErrors(fileErrors),
		})
		files = append(files, file)
	}

	sort.Slice(groups, func(i, j int) bool {
		return groups[i].File < groups[j].File
	})
	sort.Strings(files)

	return &ErrorSummary{
		TotalFiles:    len(files),
		TotalErrors:   len(errors),
		FilesAffected: files,
		Groups:        groups,
	}
}

func DeduplicateErrors(errors []ParseError) []ParseError {
	if len(errors) == 0 {
		return errors
	}

	// Sort by line, then column
	sort.Slice(errors, func(i, j int) bool {
		if errors[i].Line != errors[j].Line {
			return errors[i].Line < errors[j].Line
		}
		return errors[i].Column < errors[j].Column
	})

	var deduplicated []ParseError
	seenLines := make(map[int]bool)

	for _, err := range errors {
		if !seenLines[err.Line] {
			deduplicated = append(deduplicated, err)
			seenLines[err.Line] = true
		}
	}

	// Limit to max 5 errors per file
	if len(deduplicated) > 5 {
		deduplicated = deduplicated[:5]
	}

	return deduplicated
}

func FormatErrorSummary(summary *ErrorSummary) string {
	var b strings.Builder

	if summary.TotalErrors == 0 {
		return "🍕 Build successful! No errors found."
	}

	b.WriteString(fmt.Sprintf("🍕 Build failed! %d files with errors, %d total errors.\n", summary.TotalFiles, summary.TotalErrors))
	b.WriteString(fmt.Sprintf("Files affected: %s\n\n", strings.Join(summary.FilesAffected, ", ")))

	for _, group := range summary.Groups {
		b.WriteString(fmt.Sprintf("%s\n", group.File))
		b.WriteString(strings.Repeat("─", len(group.File)) + "\n")

		for _, err := range group.Errors {
			b.WriteString(FormatErrorWithContext(err))
			b.WriteString("\n")
		}

		if len(group.Errors) >= 5 {
			b.WriteString("... (too many errors, stopping here)\n")
		}
		b.WriteString("\n")
	}

	return b.String()
}

func FormatErrorWithContext(err ParseError) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("🍕 %s: %s at line %d, column %d\n", getErrorType(err.Kind), err.Message, err.Line, err.Column))

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

	if err.Fix != "" {
		b.WriteString("    Fix: " + err.Fix + "\n")
	}

	if err.SpecReference != "" {
		b.WriteString("    See: " + err.SpecReference + "\n")
	}

	if err.CodemodPrompt != "" {
		b.WriteString("    " + err.CodemodPrompt + "\n")
	}

	return b.String()
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
	default:
		return "Error"
	}
}
