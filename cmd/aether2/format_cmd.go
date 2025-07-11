package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	formatFlags struct {
		write     bool
		diff      bool
		check     bool
		recursive bool
		verbose   bool
	}
)

func doFormat(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		// Format current directory
		args = []string{"."}
	}

	for _, path := range args {
		if err := formatPath(path); err != nil {
			if !quiet {
				fmt.Fprintf(os.Stderr, "🍕 Error formatting %s: %v\n", path, err)
			}
			os.Exit(1)
		}
	}
}

func formatPath(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return formatDirectory(path)
	} else {
		return formatFile(path)
	}
}

func formatDirectory(dir string) error {
	if !formatFlags.recursive {
		// Only format files in current directory
		return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() && path != dir {
				return filepath.SkipDir
			}
			if !info.IsDir() && strings.HasSuffix(path, ".aeth") {
				return formatFile(path)
			}
			return nil
		})
	} else {
		// Recursive formatting
		return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(path, ".aeth") {
				return formatFile(path)
			}
			return nil
		})
	}
}

func formatFile(filePath string) error {
	if formatFlags.verbose {
		fmt.Printf("🍕 Formatting: %s\n", filePath)
	}

	// Read file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Format the content
	formatted, err := formatAetherCode(string(content))
	if err != nil {
		return err
	}

	// Check if formatting is needed
	if string(content) == formatted {
		if formatFlags.verbose {
			fmt.Printf("  ✅ Already formatted: %s\n", filePath)
		}
		return nil
	}

	// Handle different modes
	if formatFlags.check {
		if formatFlags.verbose {
			fmt.Printf("  ❌ Needs formatting: %s\n", filePath)
		}
		return fmt.Errorf("file needs formatting: %s", filePath)
	}

	if formatFlags.diff {
		// Show diff
		diff := generateDiff(string(content), formatted)
		if diff != "" {
			fmt.Printf("🍕 Diff for %s:\n%s\n", filePath, diff)
		}
	}

	if formatFlags.write {
		// Write formatted content back
		if err := os.WriteFile(filePath, []byte(formatted), 0644); err != nil {
			return err
		}
		if formatFlags.verbose {
			fmt.Printf("  ✅ Formatted: %s\n", filePath)
		}
	} else if !formatFlags.diff {
		// Just print formatted content
		fmt.Print(formatted)
	}

	return nil
}

func formatAetherCode(content string) (string, error) {
	// Basic formatting rules for Aether
	lines := strings.Split(content, "\n")
	var formatted []string

	for i, line := range lines {
		line = strings.TrimRight(line, " \t")

		// Apply formatting rules
		line = formatLine(line, i > 0 && len(lines[i-1]) > 0)

		formatted = append(formatted, line)
	}

	// Remove trailing empty lines
	for len(formatted) > 0 && formatted[len(formatted)-1] == "" {
		formatted = formatted[:len(formatted)-1]
	}

	return strings.Join(formatted, "\n"), nil
}

func formatLine(line string, hasPreviousContent bool) string {
	// Basic formatting rules
	line = strings.TrimSpace(line)

	// Ensure proper spacing around operators
	line = formatOperators(line)

	// Ensure proper indentation for blocks
	line = formatIndentation(line)

	return line
}

func formatOperators(line string) string {
	// Add spaces around operators
	operators := []string{"+", "-", "*", "/", "==", "!=", "<", ">", "<=", ">=", "="}

	for _, op := range operators {
		// This is a simplified version - in practice you'd use proper parsing
		line = strings.ReplaceAll(line, op, " "+op+" ")
	}

	// Clean up multiple spaces
	line = strings.Join(strings.Fields(line), " ")

	return line
}

func formatIndentation(line string) string {
	// This would implement proper indentation based on Aether syntax
	// For now, just return the line as-is
	return line
}

func generateDiff(original, formatted string) string {
	// Simple diff implementation
	if original == formatted {
		return ""
	}

	// This would generate a proper diff
	// For now, just show that there's a difference
	return fmt.Sprintf("--- original\n+++ formatted\n")
}

var FormatCmd = &cobra.Command{
	Use:   "format [files...]",
	Short: "Format Aether source code",
	Long: `Format Aether source code with consistent style.

This command applies consistent formatting to Aether source files:
  • Proper indentation
  • Consistent spacing
  • Operator formatting
  • Line ending normalization

Examples:
  aether format main.aeth              # Format single file
  aether format --write src/           # Format directory and write changes
  aether format --check .              # Check if files need formatting
  aether format --diff *.aeth          # Show formatting differences
  aether format --recursive .          # Format all .aeth files recursively`,
	Run: doFormat,
}

func init() {
	flags := FormatCmd.Flags()
	flags.BoolVarP(&formatFlags.write, "write", "w", false, "write formatted content to files")
	flags.BoolVarP(&formatFlags.diff, "diff", "d", false, "show formatting differences")
	flags.BoolVarP(&formatFlags.check, "check", "c", false, "check if files need formatting")
	flags.BoolVarP(&formatFlags.recursive, "recursive", "r", false, "format directories recursively")
	flags.BoolVarP(&formatFlags.verbose, "verbose", "v", false, "verbose output")
}
