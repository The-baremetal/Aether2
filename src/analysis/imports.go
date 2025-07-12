package analysis

import (
	"aether/src/lexer"
	"aether/src/parser"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
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

		if !info.IsDir() && strings.HasSuffix(path, ".aeth") {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

func ResolveImportPathsToFiles(imports map[string][]string, projectRoot string) ([]string, error) {
	var resolvedFiles []string
	
	// Read aether.toml to get dependency configuration
	configPath := filepath.Join(projectRoot, "aether.toml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	
	var config struct {
		Dependencies map[string]string `toml:"dependencies"`
	}
	if err := toml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	
	// Resolve each import path to actual file path
	for _, importPaths := range imports {
		for _, importPath := range importPaths {
			if depPath, exists := config.Dependencies[importPath]; exists {
				fullDepPath := filepath.Join(projectRoot, depPath)
				if _, err := os.Stat(fullDepPath); err == nil {
					resolvedFiles = append(resolvedFiles, fullDepPath)
				}
			}
		}
	}
	
	return resolvedFiles, nil
}

func IsExported(name string) bool {
	return len(name) > 0 && name[0] >= 'A' && name[0] <= 'Z'
}

func IsStdlibFunction(name string) bool {
	return false
}
