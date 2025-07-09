package main

import (
	"aether/src/lexer"
	"aether/src/parser"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var stdlibEnabled = true

func must(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func findAeFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".ae") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func analyzeImports(files []string) (map[string][]string, error) {
	imports := make(map[string][]string)
	for _, file := range files {
		src, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		l := lexer.NewLexer(string(src))
		p := parser.NewParser(l)
		ast := p.Parse()
		_ = ast // TODO: analyze imports from ast
	}
	return imports, nil
}

func detectCycles(imports map[string][]string) bool {
	// TODO: implement cycle detection
	return false
}

func topoSort(files []string, imports map[string][]string) ([]string, error) {
	// TODO: implement topological sort
	return files, nil
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "aether2",
		Short: "Aether language CLI",
	}

	rootCmd.AddCommand(BuildCmd)
	rootCmd.AddCommand(CrossCmd)
	rootCmd.AddCommand(InitCmd)
	rootCmd.AddCommand(CleanCmd)
	rootCmd.AddCommand(InfoCmd)
	rootCmd.AddCommand(DepsCmd)
	rootCmd.AddCommand(ScaffoldCmd)
	rootCmd.AddCommand(LintCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
