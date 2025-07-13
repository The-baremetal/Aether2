package codemod

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CodemodType int

const (
	CodemodSemicolonRemoval CodemodType = iota
	CodemodImportFix
	CodemodFunctionDeclaration
	CodemodAutoFix
)

type CodemodResult struct {
	FileChanged bool
	Changes     []Change
	Errors      []error
}

type Change struct {
	Line     int
	Column   int
	OldText  string
	NewText  string
	Type     string
	Message  string
}

type CodemodEngine struct {
	interactive bool
	previewOnly bool
	autoFix     bool
	backupDir   string
}

func NewCodemodEngine() *CodemodEngine {
	return &CodemodEngine{
		interactive: false,
		previewOnly: false,
		autoFix:     false,
		backupDir:   ".aether-codemod-backup",
	}
}

func (ce *CodemodEngine) SetInteractive(interactive bool) {
	ce.interactive = interactive
}

func (ce *CodemodEngine) SetPreviewOnly(preview bool) {
	ce.previewOnly = preview
}

func (ce *CodemodEngine) SetAutoFix(autoFix bool) {
	ce.autoFix = autoFix
}

func (ce *CodemodEngine) SetBackupDir(dir string) {
	ce.backupDir = dir
}

func (ce *CodemodEngine) ExecuteCodemod(filePath string, codemodType CodemodType) (*CodemodResult, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %v", filePath, err)
	}

	lines := strings.Split(string(content), "\n")
	result := &CodemodResult{
		FileChanged: false,
		Changes:     []Change{},
		Errors:      []error{},
	}

	switch codemodType {
	case CodemodSemicolonRemoval:
		result = ce.removeSemicolons(lines, filePath)
	case CodemodImportFix:
		result = ce.fixImports(lines, filePath)
	case CodemodFunctionDeclaration:
		result = ce.fixFunctionDeclarations(lines, filePath)
	case CodemodAutoFix:
		result = ce.autoFixAll(lines, filePath)
	default:
		return nil, fmt.Errorf("unknown codemod type: %d", codemodType)
	}

	if result.FileChanged && !ce.previewOnly {
		if err := ce.applyChanges(filePath, result); err != nil {
			result.Errors = append(result.Errors, err)
		}
	}

	return result, nil
}

func (ce *CodemodEngine) removeSemicolons(lines []string, filePath string) *CodemodResult {
	result := &CodemodResult{
		FileChanged: false,
		Changes:     []Change{},
		Errors:      []error{},
	}

	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		if strings.HasSuffix(strings.TrimSpace(line), ";") {
			newLine := strings.TrimSuffix(strings.TrimSpace(line), ";")
			if newLine != strings.TrimSpace(line) {
				result.Changes = append(result.Changes, Change{
					Line:     i + 1,
					Column:   len(line),
					OldText:  line,
					NewText:  newLine,
					Type:     "semicolon_removal",
					Message:  "Removed unnecessary semicolon",
				})
				result.FileChanged = true
			}
		}
	}

	return result
}

func (ce *CodemodEngine) fixImports(lines []string, filePath string) *CodemodResult {
	result := &CodemodResult{
		FileChanged: false,
		Changes:     []Change{},
		Errors:      []error{},
	}

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "import") {
			if strings.Contains(trimmed, `"`) {
				newLine := strings.ReplaceAll(trimmed, `"`, "")
				if newLine != trimmed {
					result.Changes = append(result.Changes, Change{
						Line:     i + 1,
						Column:   len(line),
						OldText:  line,
						NewText:  newLine,
						Type:     "import_fix",
						Message:  "Removed quotes from import statement",
					})
					result.FileChanged = true
				}
			}
		}
	}

	return result
}

func (ce *CodemodEngine) fixFunctionDeclarations(lines []string, filePath string) *CodemodResult {
	result := &CodemodResult{
		FileChanged: false,
		Changes:     []Change{},
		Errors:      []error{},
	}

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "func") {
			if strings.Contains(trimmed, "let ") {
				newLine := strings.ReplaceAll(trimmed, "let ", "")
				if newLine != trimmed {
					result.Changes = append(result.Changes, Change{
						Line:     i + 1,
						Column:   len(line),
						OldText:  line,
						NewText:  newLine,
						Type:     "function_declaration_fix",
						Message:  "Removed 'let' keyword from function declaration",
					})
					result.FileChanged = true
				}
			}
		}
	}

	return result
}

func (ce *CodemodEngine) autoFixAll(lines []string, filePath string) *CodemodResult {
	result := &CodemodResult{
		FileChanged: false,
		Changes:     []Change{},
		Errors:      []error{},
	}

	semicolonResult := ce.removeSemicolons(lines, filePath)
	result.Changes = append(result.Changes, semicolonResult.Changes...)
	result.FileChanged = result.FileChanged || semicolonResult.FileChanged

	importResult := ce.fixImports(lines, filePath)
	result.Changes = append(result.Changes, importResult.Changes...)
	result.FileChanged = result.FileChanged || importResult.FileChanged

	functionResult := ce.fixFunctionDeclarations(lines, filePath)
	result.Changes = append(result.Changes, functionResult.Changes...)
	result.FileChanged = result.FileChanged || functionResult.FileChanged

	return result
}

func (ce *CodemodEngine) applyChanges(filePath string, result *CodemodResult) error {
	if ce.backupDir != "" {
		if err := ce.createBackup(filePath); err != nil {
			return fmt.Errorf("failed to create backup: %v", err)
		}
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file for changes: %v", err)
	}

	lines := strings.Split(string(content), "\n")

	for _, change := range result.Changes {
		if change.Line > 0 && change.Line <= len(lines) {
			lines[change.Line-1] = change.NewText
		}
	}

	newContent := strings.Join(lines, "\n")
	return os.WriteFile(filePath, []byte(newContent), 0644)
}

func (ce *CodemodEngine) createBackup(filePath string) error {
	if ce.backupDir == "" {
		return nil
	}

	if err := os.MkdirAll(ce.backupDir, 0755); err != nil {
		return err
	}

	backupPath := filepath.Join(ce.backupDir, filepath.Base(filePath))
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	return os.WriteFile(backupPath, content, 0644)
}

func (ce *CodemodEngine) GetPreview(changes []Change) string {
	var preview strings.Builder
	preview.WriteString("Codemod Preview:\n")
	preview.WriteString("================\n\n")

	for i, change := range changes {
		preview.WriteString(fmt.Sprintf("Change %d:\n", i+1))
		preview.WriteString(fmt.Sprintf("  Line: %d, Column: %d\n", change.Line, change.Column))
		preview.WriteString(fmt.Sprintf("  Type: %s\n", change.Type))
		preview.WriteString(fmt.Sprintf("  Message: %s\n", change.Message))
		preview.WriteString(fmt.Sprintf("  Old: %s\n", change.OldText))
		preview.WriteString(fmt.Sprintf("  New: %s\n", change.NewText))
		preview.WriteString("\n")
	}

	return preview.String()
} 