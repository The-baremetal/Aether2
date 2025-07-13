package main

import (
	"fmt"
	"os"
	"strings"

	"aether/src/analysis"
	"aether/src/codemod"

	"github.com/spf13/cobra"
)

var codemodFlags struct {
	interactive bool
	previewOnly bool
	autoFix     bool
	backup      bool
	types       []string
}

var CodemodCmd = &cobra.Command{
	Use:   "codemod [files...]",
	Short: "Apply code modifications to Aether files",
	Long: `Apply automated code modifications to Aether source files.

Available codemod types:
  semicolon-removal    Remove unnecessary semicolons
  import-fix          Fix import statement syntax
  function-declaration Fix function declaration syntax
  auto-fix            Apply all available fixes

Examples:
  aether codemod src/main.aeth                    # Apply auto-fix to single file
  aether codemod --preview-only src/              # Preview changes without applying
  aether codemod --interactive src/               # Interactive mode with confirmations
  aether codemod --types semicolon-removal src/   # Apply specific codemod type`,
	Run: func(cmd *cobra.Command, args []string) {
		doCodemod(args)
	},
}

func init() {
	flags := CodemodCmd.Flags()
	flags.BoolVarP(&codemodFlags.interactive, "interactive", "i", false, "interactive mode with confirmations")
	flags.BoolVarP(&codemodFlags.previewOnly, "preview-only", "p", false, "show preview without applying changes")
	flags.BoolVarP(&codemodFlags.autoFix, "auto-fix", "a", false, "apply all available fixes")
	flags.BoolVarP(&codemodFlags.backup, "backup", "b", true, "create backup before making changes")
	flags.StringSliceVarP(&codemodFlags.types, "types", "t", []string{"auto-fix"}, "codemod types to apply")
}

func doCodemod(args []string) {
	if len(args) == 0 {
		fmt.Println("Error: No files or directories specified")
		fmt.Println("Usage: aether codemod [files...]")
		os.Exit(1)
	}

	engine := codemod.NewCodemodEngine()
	engine.SetInteractive(codemodFlags.interactive)
	engine.SetPreviewOnly(codemodFlags.previewOnly)
	engine.SetAutoFix(codemodFlags.autoFix)

	if !codemodFlags.backup {
		engine.SetBackupDir("")
	}

	var files []string
	for _, arg := range args {
		info, err := os.Stat(arg)
		if err != nil {
			fmt.Printf("Error: Cannot access '%s': %v\n", arg, err)
			continue
		}

		if info.IsDir() {
			dirFiles, err := analysis.FindAetherFiles(arg)
			if err != nil {
				fmt.Printf("Error: Failed to scan directory '%s': %v\n", arg, err)
				continue
			}
			files = append(files, dirFiles...)
		} else {
			if !strings.HasSuffix(arg, ".aeth") {
				fmt.Printf("Warning: Skipping non-Aether file '%s'\n", arg)
				continue
			}
			files = append(files, arg)
		}
	}

	if len(files) == 0 {
		fmt.Println("No Aether files found to process")
		os.Exit(1)
	}

	fmt.Printf("Processing %d files...\n", len(files))

	totalChanges := 0
	totalErrors := 0

	for _, file := range files {
		fmt.Printf("\nProcessing: %s\n", file)
		
		for _, codemodType := range codemodFlags.types {
			ct := parseCodemodType(codemodType)
			if ct == -1 {
				fmt.Printf("Warning: Unknown codemod type '%s', skipping\n", codemodType)
				continue
			}

			result, err := engine.ExecuteCodemod(file, ct)
			if err != nil {
				fmt.Printf("Error processing %s: %v\n", file, err)
				totalErrors++
				continue
			}

			if len(result.Changes) > 0 {
				if codemodFlags.previewOnly {
					fmt.Println(engine.GetPreview(result.Changes))
				} else {
					fmt.Printf("Applied %d changes to %s\n", len(result.Changes), file)
					totalChanges += len(result.Changes)
				}
			}

			if len(result.Errors) > 0 {
				for _, err := range result.Errors {
					fmt.Printf("Error: %v\n", err)
				}
				totalErrors += len(result.Errors)
			}
		}
	}

	fmt.Printf("\nSummary:\n")
	fmt.Printf("  Files processed: %d\n", len(files))
	fmt.Printf("  Total changes: %d\n", totalChanges)
	fmt.Printf("  Total errors: %d\n", totalErrors)

	if totalErrors > 0 {
		os.Exit(1)
	}
}

func parseCodemodType(typeStr string) codemod.CodemodType {
	switch strings.ToLower(typeStr) {
	case "semicolon-removal":
		return codemod.CodemodSemicolonRemoval
	case "import-fix":
		return codemod.CodemodImportFix
	case "function-declaration":
		return codemod.CodemodFunctionDeclaration
	case "auto-fix":
		return codemod.CodemodAutoFix
	default:
		return -1
	}
} 