package analysis

import (
	"aether/src/lexer"
	"aether/src/parser"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func AnalyzeImports(files []string) (map[string][]string, error) {
	imports := make(map[string][]string)

	for _, file := range files {
		src, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		l := lexer.NewLexer(string(src))
		p := parser.NewParser(l)
		ast := p.Parse()

		fileImports := extractImportsFromAST(ast)
		imports[file] = fileImports
	}

	return imports, nil
}

func extractImportsFromAST(ast *parser.Program) []string {
	var imports []string

	for _, stmt := range ast.Statements {
		if importStmt, ok := stmt.(*parser.Import); ok {
			importPath := importStmt.Name.Value
			if importStmt.As != nil && importStmt.As.Value != "" {
				importPath = importStmt.As.Value
			}
			imports = append(imports, importPath)
		}
	}

	return imports
}

func AnalyzeImportStatement(importStmt *parser.Import, filePath string, result *AnalysisResult) {
	importPath := importStmt.Name.Value
	importInfo := ImportInfo{
		Path:   importPath,
		Valid:  true,
		Exists: false,
	}

	if !IsValidImportPath(importPath) {
		importInfo.Valid = false
		importInfo.Errors = append(importInfo.Errors, "Invalid import path format")
		result.Errors = append(result.Errors, fmt.Sprintf("%s: Invalid import path '%s'", filePath, importPath))
	}

	resolvedPath := ResolveImportPath(importPath, filepath.Dir(filePath))
	if resolvedPath != "" {
		if _, err := os.Stat(resolvedPath); err == nil {
			importInfo.Exists = true
			importInfo.Resolved = resolvedPath
		} else {
			importInfo.Errors = append(importInfo.Errors, "Imported file does not exist")
			result.Errors = append(result.Errors, fmt.Sprintf("%s: Imported file '%s' does not exist", filePath, importPath))
		}
	}

	result.Imports[importPath] = importInfo
}

func FindAetherFiles(root string) ([]string, error) {
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

func IsExported(name string) bool {
	return len(name) > 0 && name[0] >= 'A' && name[0] <= 'Z'
}

func IsStdlibFunction(name string) bool {
	return false
}
