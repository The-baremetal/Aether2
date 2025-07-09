package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var LintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Lint all .ae files for casing consistency",
	Run: func(cmd *cobra.Command, args []string) {
		lintProject()
	},
}

func lintProject() {
	files := []string{}
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() && strings.HasSuffix(path, ".ae") {
			files = append(files, path)
		}
		return nil
	})
	for _, file := range files {
		checkFileCasing(file)
		checkIdentifierCasing(file)
	}
}

func checkFileCasing(path string) {
	base := filepath.Base(path)
	if !regexp.MustCompile(`^[a-z0-9_]+\.ae$`).MatchString(base) {
		fmt.Println("🍕 Warning: File '", base, "' should be snake_case (e.g., 'main.ae')")
	}
}

func checkIdentifierCasing(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "func ") {
			fn := strings.Fields(line)[1]
			fn = strings.Split(fn, "(")[0]
			if !regexp.MustCompile(`^[a-z0-9_]+$`).MatchString(fn) {
				fmt.Println("🍕 Warning: Function '", fn, "' in ", path, " should be snake_case")
			}
		}
		if strings.Contains(line, " = ") {
			parts := strings.Split(line, " = ")
			varName := strings.Fields(parts[0])[0]
			if !regexp.MustCompile(`^[a-z0-9_]+$`).MatchString(varName) {
				fmt.Println("🍕 Warning: Variable '", varName, "' in ", path, " should be snake_case")
			}
		}
		if strings.HasPrefix(line, "type ") {
			typeName := strings.Fields(line)[1]
			if !regexp.MustCompile(`^[A-Z][A-Za-z0-9]*$`).MatchString(typeName) {
				fmt.Println("🍕 Warning: Type '", typeName, "' in ", path, " should be PascalCase")
			}
		}
		if strings.HasPrefix(line, "const ") {
			constName := strings.Fields(line)[1]
			if !regexp.MustCompile(`^[A-Z0-9_]+$`).MatchString(constName) {
				fmt.Println("🍕 Warning: Constant '", constName, "' in ", path, " should be UPPER_CASE")
			}
		}
	}
}
