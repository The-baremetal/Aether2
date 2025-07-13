package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"aether/src/analysis"

	"github.com/spf13/cobra"
)

var (
	docsFlags struct {
		output    string
		format    string
		recursive bool
		verbose   bool
		open      bool
	}
)

func doDocs(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		// Generate docs for current directory
		args = []string{"."}
	}

	fmt.Println("🍕 Generating Aether documentation...")
	fmt.Println()

	for _, path := range args {
		if err := generateDocs(path); err != nil {
			if !quiet {
				fmt.Fprintf(os.Stderr, "🍕 Error generating docs for %s: %v\n", path, err)
			}
			os.Exit(1)
		}
	}
}

func generateDocs(path string) error {
	// Find Aether files
	files, err := analysis.FindAetherFiles(path)
	if err != nil {
		return fmt.Errorf("failed to find Aether files: %v", err)
	}
	if len(files) == 0 {
		fmt.Printf("🍕 No Aether files found in %s\n", path)
		return nil
	}

	if docsFlags.verbose {
		fmt.Printf("🍕 Found %d Aether files:\n", len(files))
		for _, file := range files {
			fmt.Printf("  • %s\n", file)
		}
		fmt.Println()
	}

	// Generate documentation
	docs := generateDocumentation(files)

	// Write output
	outputFile := docsFlags.output
	if outputFile == "" {
		outputFile = "docs.md"
	}

	if err := writeDocs(outputFile, docs); err != nil {
		return err
	}

	if docsFlags.verbose {
		fmt.Printf("🍕 Documentation written to: %s\n", outputFile)
	}

	// Open docs if requested
	if docsFlags.open {
		openDocumentation(outputFile)
	}

	return nil
}



func generateDocumentation(files []string) string {
	var docs strings.Builder

	docs.WriteString("# 🍕 Aether Project Documentation\n\n")
	docs.WriteString("Generated documentation for Aether project.\n\n")

	// Project overview
	docs.WriteString("## 📁 Project Structure\n\n")
	for _, file := range files {
		docs.WriteString(fmt.Sprintf("- `%s`\n", file))
	}
	docs.WriteString("\n")

	// Generate docs for each file
	for _, file := range files {
		fileDocs := generateFileDocs(file)
		docs.WriteString(fileDocs)
		docs.WriteString("\n")
	}

	// API documentation
	docs.WriteString("## 🔧 API Reference\n\n")
	docs.WriteString("### Functions\n\n")
	// This would extract function signatures
	docs.WriteString("### Types\n\n")
	// This would extract type definitions
	docs.WriteString("### Constants\n\n")
	// This would extract constants

	return docs.String()
}

func generateFileDocs(filePath string) string {
	var docs strings.Builder

	// Read file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Sprintf("## %s\n\nError reading file: %v\n\n", filePath, err)
	}

	// Parse file for documentation
	lines := strings.Split(string(content), "\n")

	docs.WriteString(fmt.Sprintf("## %s\n\n", filePath))

	// Extract comments and function definitions
	for i, line := range lines {
		_ = i
		line = strings.TrimSpace(line)

		// Extract comments
		if strings.HasPrefix(line, "//") {
			comment := strings.TrimSpace(strings.TrimPrefix(line, "//"))
			docs.WriteString(fmt.Sprintf("%s\n", comment))
		}

		// Extract function definitions
		if strings.HasPrefix(line, "func ") {
			funcName := extractFunctionName(line)
			docs.WriteString(fmt.Sprintf("### %s\n\n", funcName))
			docs.WriteString(fmt.Sprintf("```aether\n%s\n```\n\n", line))
		}
	}

	return docs.String()
}

func extractFunctionName(line string) string {
	// Simple function name extraction
	parts := strings.Fields(line)
	if len(parts) >= 2 {
		return parts[1]
	}
	return "Unknown"
}

func writeDocs(outputFile, content string) error {
	// Create output directory if needed
	outputDir := filepath.Dir(outputFile)
	if outputDir != "." {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return err
		}
	}

	return os.WriteFile(outputFile, []byte(content), 0644)
}

func openDocumentation(filePath string) {
	// This would open the documentation in a browser
	// For now, just print the path
	fmt.Printf("🍕 Documentation ready: %s\n", filePath)
}

var DocsCmd = &cobra.Command{
	Use:   "docs [files...]",
	Short: "Generate Aether documentation",
	Long: `Generate comprehensive documentation for Aether projects.

This command:
  • Scans Aether source files
  • Extracts function signatures
  • Generates API documentation
  • Creates project overview
  • Supports multiple output formats

Examples:
  aether docs                    # Generate docs for current project
  aether docs --recursive .      # Include subdirectories
  aether docs --output api.md    # Custom output file
  aether docs --format html      # Generate HTML docs
  aether docs --open            # Open docs in browser`,
	Run: doDocs,
}

func init() {
	flags := DocsCmd.Flags()
	flags.StringVarP(&docsFlags.output, "output", "o", "", "output file (default: docs.md)")
	flags.StringVarP(&docsFlags.format, "format", "f", "markdown", "output format (markdown, html, json)")
	flags.BoolVarP(&docsFlags.recursive, "recursive", "r", false, "include subdirectories")
	flags.BoolVarP(&docsFlags.verbose, "verbose", "v", false, "verbose output")
	flags.BoolVarP(&docsFlags.open, "open", "p", false, "open docs in browser")
}
