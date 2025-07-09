package compiler

import (
	"aether/src/lexer"
	"aether/src/parser"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type AnalysisResult struct {
	Valid        bool
	Errors       []string
	Warnings     []string
	Imports      map[string]ImportInfo
	Functions    map[string]FunctionInfo
	Variables    map[string]VariableInfo
	Types        map[string]TypeInfo
	Constants    map[string]ConstantInfo
	Dependencies map[string][]string
	Cycles       [][]string
	Unused       []string
	Undefined    []string
}

type ImportInfo struct {
	Path     string
	Valid    bool
	Exists   bool
	Resolved string
	Errors   []string
}

type FunctionInfo struct {
	Name       string
	Parameters []ParameterInfo
	ReturnType string
	Defined    bool
	Used       bool
	Exported   bool
}

type VariableInfo struct {
	Name     string
	Type     string
	Defined  bool
	Used     bool
	Exported bool
	Scope    string
}

type TypeInfo struct {
	Name     string
	Defined  bool
	Used     bool
	Exported bool
	Fields   map[string]string
}

type ConstantInfo struct {
	Name     string
	Type     string
	Value    interface{}
	Defined  bool
	Used     bool
	Exported bool
}

type ParameterInfo struct {
	Name string
	Type string
}

func AnalyzeProject(projectPath string) *AnalysisResult {
	result := &AnalysisResult{
		Valid:        true,
		Errors:       []string{},
		Warnings:     []string{},
		Imports:      make(map[string]ImportInfo),
		Functions:    make(map[string]FunctionInfo),
		Variables:    make(map[string]VariableInfo),
		Types:        make(map[string]TypeInfo),
		Constants:    make(map[string]ConstantInfo),
		Dependencies: make(map[string][]string),
		Cycles:       [][]string{},
		Unused:       []string{},
		Undefined:    []string{},
	}

	files, err := findAetherFiles(projectPath)
	if err != nil {
		result.Valid = false
		result.Errors = append(result.Errors, fmt.Sprintf("Failed to find Aether files: %v", err))
		return result
	}

	for _, file := range files {
		analyzeFile(file, result)
	}

	validateImports(result)
	checkDependencies(result)
	detectCycles(result)
	checkUnusedDeclarations(result)
	checkUndefinedReferences(result)

	return result
}

func analyzeFile(filePath string, result *AnalysisResult) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("Failed to read file %s: %v", filePath, err))
		return
	}

	// Parse the file
	lexer := lexer.NewLexer(string(content))
	parser := parser.NewParser(lexer)
	ast := parser.Parse()

	if len(parser.Errors.ToMessages()) > 0 {
		for _, msg := range parser.Errors.ToMessages() {
			result.Errors = append(result.Errors, fmt.Sprintf("%s: %s", filePath, msg))
		}
		return
	}

	analyzeAST(ast, filePath, result)
}

func analyzeAST(ast *parser.Program, filePath string, result *AnalysisResult) {
	for _, stmt := range ast.Statements {
		analyzeStatement(stmt, filePath, result)
	}
}

func analyzeStatement(stmt parser.Statement, filePath string, result *AnalysisResult) {
	switch s := stmt.(type) {
	case *parser.Import:
		analyzeImportStatement(s, filePath, result)
	case *parser.Function:
		analyzeFunctionDeclaration(s, filePath, result)
	case *parser.Assignment:
		analyzeAssignment(s, filePath, result)
	case *parser.StructDef:
		analyzeTypeDeclaration(s, filePath, result)
	case *parser.Call:
		analyzeFunctionCall(s, filePath, result)
	case *parser.If:
		analyzeIfStatement(s, filePath, result)
	case *parser.While:
		analyzeWhileStatement(s, filePath, result)
	case *parser.For:
		analyzeForStatement(s, filePath, result)
	case *parser.Return:
		analyzeReturnStatement(s, filePath, result)
	case *parser.Block:
		analyzeBlock(s, filePath, result)
	}
}

func analyzeImportStatement(importStmt *parser.Import, filePath string, result *AnalysisResult) {
	importPath := importStmt.Name.Value
	importInfo := ImportInfo{
		Path:   importPath,
		Valid:  true,
		Exists: false,
	}

	// Check if import path is valid
	if !isValidImportPath(importPath) {
		importInfo.Valid = false
		importInfo.Errors = append(importInfo.Errors, "Invalid import path format")
		result.Errors = append(result.Errors, fmt.Sprintf("%s: Invalid import path '%s'", filePath, importPath))
	}

	// Check if imported file exists
	resolvedPath := resolveImportPath(importPath, filepath.Dir(filePath))
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

func analyzeFunctionDeclaration(funcDecl *parser.Function, filePath string, result *AnalysisResult) {
	funcInfo := FunctionInfo{
		Name:       funcDecl.Name.Value,
		Parameters: []ParameterInfo{},
		ReturnType: "void",
		Defined:    true,
		Used:       false,
		Exported:   isExported(funcDecl.Name.Value),
	}

	for _, param := range funcDecl.Params {
		funcInfo.Parameters = append(funcInfo.Parameters, ParameterInfo{
			Name: param.Value,
			Type: param.Type,
		})
	}

	result.Functions[funcDecl.Name.Value] = funcInfo
}

func analyzeVariableDeclaration(assign *parser.Assignment, filePath string, result *AnalysisResult) {
	varInfo := VariableInfo{
		Name:     assign.Name.Value,
		Type:     assign.Name.Type,
		Defined:  true,
		Used:     false,
		Exported: isExported(assign.Name.Value),
		Scope:    "local",
	}

	result.Variables[assign.Name.Value] = varInfo
}

func analyzeTypeDeclaration(structDef *parser.StructDef, filePath string, result *AnalysisResult) {
	typeInfo := TypeInfo{
		Name:     structDef.Name.Value,
		Defined:  true,
		Used:     false,
		Exported: isExported(structDef.Name.Value),
		Fields:   make(map[string]string),
	}

	for _, field := range structDef.Fields {
		typeInfo.Fields[field.Value] = field.Type
	}

	result.Types[structDef.Name.Value] = typeInfo
}

func analyzeFunctionCall(call *parser.Call, filePath string, result *AnalysisResult) {
	if ident, ok := call.Function.(*parser.Identifier); ok {
		funcName := ident.Value

		if funcInfo, exists := result.Functions[funcName]; exists {
			funcInfo.Used = true
			result.Functions[funcName] = funcInfo
		} else if !isStdlibFunction(funcName) {
			result.Undefined = append(result.Undefined, fmt.Sprintf("function '%s'", funcName))
		}
	}

	for _, arg := range call.Args {
		analyzeExpression(arg, filePath, result)
	}
}

func analyzeAssignment(assign *parser.Assignment, filePath string, result *AnalysisResult) {
	if varInfo, exists := result.Variables[assign.Name.Value]; exists {
		varInfo.Used = true
		result.Variables[assign.Name.Value] = varInfo
	} else {
		result.Undefined = append(result.Undefined, fmt.Sprintf("variable '%s'", assign.Name.Value))
	}

	analyzeExpression(assign.Value, filePath, result)
}

func analyzeIfStatement(ifstmt *parser.If, filePath string, result *AnalysisResult) {
	analyzeExpression(ifstmt.Condition, filePath, result)
	analyzeStatement(ifstmt.Consequence, filePath, result)
	if ifstmt.Alternative != nil {
		analyzeStatement(ifstmt.Alternative, filePath, result)
	}
}

func analyzeWhileStatement(while *parser.While, filePath string, result *AnalysisResult) {
	analyzeExpression(while.Condition, filePath, result)
	analyzeStatement(while.Body, filePath, result)
}

func analyzeForStatement(forstmt *parser.For, filePath string, result *AnalysisResult) {
	analyzeExpression(forstmt.Iterable, filePath, result)
	analyzeStatement(forstmt.Body, filePath, result)
}

func analyzeReturnStatement(ret *parser.Return, filePath string, result *AnalysisResult) {
	if ret.Value != nil {
		analyzeExpression(ret.Value, filePath, result)
	}
}

func analyzeBlock(block *parser.Block, filePath string, result *AnalysisResult) {
	for _, stmt := range block.Statements {
		analyzeStatement(stmt, filePath, result)
	}
}

func analyzeExpression(expr parser.Expression, filePath string, result *AnalysisResult) {
	switch e := expr.(type) {
	case *parser.Identifier:
		analyzeIdentifier(e, filePath, result)
	case *parser.Call:
		analyzeFunctionCall(e, filePath, result)
	case *parser.Literal:
		// Literals don't need analysis
	case *parser.Array:
		for _, elem := range e.Elements {
			analyzeExpression(elem, filePath, result)
		}
	case *parser.PropertyAccess:
		analyzeExpression(e.Object, filePath, result)
	case *parser.PartialApplication:
		analyzeExpression(e.Function, filePath, result)
		for _, arg := range e.Args {
			analyzeExpression(arg, filePath, result)
		}
	}
}

func analyzeIdentifier(ident *parser.Identifier, filePath string, result *AnalysisResult) {
	name := ident.Value

	if varInfo, exists := result.Variables[name]; exists {
		varInfo.Used = true
		result.Variables[name] = varInfo
	} else if funcInfo, exists := result.Functions[name]; exists {
		funcInfo.Used = true
		result.Functions[name] = funcInfo
	} else if constInfo, exists := result.Constants[name]; exists {
		constInfo.Used = true
		result.Constants[name] = constInfo
	} else if typeInfo, exists := result.Types[name]; exists {
		typeInfo.Used = true
		result.Types[name] = typeInfo
	} else if !isStdlibFunction(name) {
		result.Undefined = append(result.Undefined, fmt.Sprintf("identifier '%s'", name))
	}
}

func validateImports(result *AnalysisResult) {
	for path, importInfo := range result.Imports {
		_ = path
		if !importInfo.Valid {
			result.Valid = false
		}
		if !importInfo.Exists {
			result.Valid = false
		}
	}
}

func checkDependencies(result *AnalysisResult) {
	// Build dependency graph based on imports
	for path, importInfo := range result.Imports {
		if importInfo.Exists && importInfo.Resolved != "" {
			result.Dependencies[path] = []string{importInfo.Resolved}
		}
	}
}

func detectCycles(result *AnalysisResult) {
	// Use topological sort to detect cycles
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	for path := range result.Dependencies {
		if !visited[path] {
			if hasCycle(path, result.Dependencies, visited, recStack, []string{}) {
				result.Cycles = append(result.Cycles, []string{})
				result.Valid = false
				result.Errors = append(result.Errors, "Circular import detected")
			}
		}
	}
}

func hasCycle(path string, deps map[string][]string, visited, recStack map[string]bool, pathStack []string) bool {
	visited[path] = true
	recStack[path] = true
	pathStack = append(pathStack, path)

	for _, dep := range deps[path] {
		if !visited[dep] {
			if hasCycle(dep, deps, visited, recStack, pathStack) {
				return true
			}
		} else if recStack[dep] {
			return true
		}
	}

	recStack[path] = false
	return false
}

func checkUnusedDeclarations(result *AnalysisResult) {
	for name, funcInfo := range result.Functions {
		if funcInfo.Defined && !funcInfo.Used && !funcInfo.Exported {
			result.Unused = append(result.Unused, fmt.Sprintf("function '%s'", name))
			result.Warnings = append(result.Warnings, fmt.Sprintf("Unused function '%s'", name))
		}
	}

	for name, varInfo := range result.Variables {
		if varInfo.Defined && !varInfo.Used && !varInfo.Exported {
			result.Unused = append(result.Unused, fmt.Sprintf("variable '%s'", name))
			result.Warnings = append(result.Warnings, fmt.Sprintf("Unused variable '%s'", name))
		}
	}

	for name, constInfo := range result.Constants {
		if constInfo.Defined && !constInfo.Used && !constInfo.Exported {
			result.Unused = append(result.Unused, fmt.Sprintf("constant '%s'", name))
			result.Warnings = append(result.Warnings, fmt.Sprintf("Unused constant '%s'", name))
		}
	}

	for name, typeInfo := range result.Types {
		if typeInfo.Defined && !typeInfo.Used && !typeInfo.Exported {
			result.Unused = append(result.Unused, fmt.Sprintf("type '%s'", name))
			result.Warnings = append(result.Warnings, fmt.Sprintf("Unused type '%s'", name))
		}
	}
}

func checkUndefinedReferences(result *AnalysisResult) {
	// Remove duplicates from undefined list
	seen := make(map[string]bool)
	var unique []string
	for _, item := range result.Undefined {
		if !seen[item] {
			seen[item] = true
			unique = append(unique, item)
		}
	}
	result.Undefined = unique

	if len(result.Undefined) > 0 {
		result.Valid = false
		for _, item := range result.Undefined {
			result.Errors = append(result.Errors, fmt.Sprintf("Undefined reference: %s", item))
		}
	}
}

func isValidImportPath(path string) bool {
	if path == "" {
		return false
	}

	if strings.HasPrefix(path, ".") || strings.HasPrefix(path, "/") {
		return true
	}

	if strings.Contains(path, "..") {
		return false
	}

	return true
}

func resolveImportPath(importPath, currentDir string) string {
	if strings.HasPrefix(importPath, "/") {
		return importPath + ".ae"
	}

	if strings.HasPrefix(importPath, ".") {
		return filepath.Join(currentDir, importPath+".ae")
	}

	return filepath.Join(currentDir, importPath+".ae")
}

func isExported(name string) bool {
	return len(name) > 0 && name[0] >= 'A' && name[0] <= 'Z'
}

func isStdlibFunction(name string) bool {
	stdlibFuncs := []string{"print", "println", "len", "cap", "append", "make", "new", "delete", "close", "panic", "recover"}
	for _, funcName := range stdlibFuncs {
		if name == funcName {
			return true
		}
	}
	return false
}

func findAetherFiles(root string) ([]string, error) {
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
