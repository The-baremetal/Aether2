package analysis

import (
	"aether/src/lexer"
	"aether/src/parser"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"aether/lib/utils"

	"github.com/BurntSushi/toml"
)

func AnalyzeProject(projectPath string) *AnalysisResult {
	result := &AnalysisResult{
		Valid:        true,
		Errors:       []utils.ParseError{},
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
		CIncludes:    []CInclude{},
	}

	files, err := findAetherFiles(projectPath)
	if err != nil {
		result.Valid = false
		result.Errors = append(result.Errors, utils.ParseError{Message: fmt.Sprintf("Failed to find Aether files: %v", err)})
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

func AnalyzeAST(ast *parser.ASTNode) *AnalysisResult {
	result := &AnalysisResult{
		Valid:        true,
		Errors:       []utils.ParseError{},
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
		CIncludes:    []CInclude{},
	}

	extractCIncludes(ast, result)
	return result
}

func extractCIncludes(node *parser.ASTNode, result *AnalysisResult) {
	if node == nil {
		return
	}

	if node.NodeKind == parser.CCommentKind {
		if content, ok := node.Value.(string); ok {
			includes := parseCIncludes(content)
			result.CIncludes = append(result.CIncludes, includes...)
		}
	}

	for _, inner := range node.Inner {
		extractCIncludes(inner, result)
	}

	if node.Left != nil {
		extractCIncludes(node.Left, result)
	}

	if node.Right != nil {
		extractCIncludes(node.Right, result)
	}

	if node.Body != nil {
		extractCIncludes(node.Body, result)
	}

	for _, param := range node.Params {
		extractCIncludes(param, result)
	}
}

func AnalyzeDependencies(projectPath string) *DependencyAnalysis {
	result := &DependencyAnalysis{
		Valid:        true,
		Errors:       []utils.ParseError{},
		Warnings:     []string{},
		ResolvedDeps: make(map[string]string),
		MissingDeps:  []string{},
		UnusedDeps:   []string{},
	}

	// Read aether.toml
	configPath := filepath.Join(projectPath, "aether.toml")
	configData, err := os.ReadFile(configPath)
	if err != nil {
		result.Valid = false
		// Read the actual file content for the snippet
		content, _ := os.ReadFile(filepath.Join(projectPath, "aether.toml"))
		snippet := ""
		if len(content) > 0 {
			lines := strings.Split(string(content), "\n")
			if len(lines) > 0 {
				snippet = lines[0] // Show first line as context
			}
		}
		result.Errors = append(result.Errors, utils.ParseError{
			Message: fmt.Sprintf("Failed to read aether.toml: %v", err),
			File:    filepath.Join(projectPath, "aether.toml"),
			Snippet: snippet,
			Caret:   1,
			Fix:     "Check if aether.toml exists in the project root",
		})
		return result
	}

	var config struct {
		Dependencies map[string]string `toml:"dependencies"`
	}

	if err := toml.Unmarshal(configData, &config); err != nil {
		result.Valid = false
		// Read the actual file content for the snippet
		lines := strings.Split(string(configData), "\n")
		snippet := ""
		if len(lines) > 0 {
			snippet = lines[0] // Show first line as context
		}
		result.Errors = append(result.Errors, utils.ParseError{
			Message: fmt.Sprintf("Failed to parse aether.toml: %v", err),
			File:    filepath.Join(projectPath, "aether.toml"),
			Snippet: snippet,
			Caret:   1,
			Fix:     "Check if aether.toml has valid TOML syntax",
		})
		return result
	}

	// Read lock file if it exists
	lockPath := filepath.Join(projectPath, "aether.lock")
	var lockFile LockFile
	if lockData, err := os.ReadFile(lockPath); err == nil {
		if err := toml.Unmarshal(lockData, &lockFile); err != nil {
			result.Warnings = append(result.Warnings, "Failed to parse aether.lock, will regenerate")
		}
	}

	// Analyze source files for imports and function usage
	files, err := findAetherFiles(filepath.Join(projectPath, "src"))
	if err != nil {
		result.Valid = false
		result.Errors = append(result.Errors, utils.ParseError{
			Message: fmt.Sprintf("Failed to find source files: %v", err),
			File:    filepath.Join(projectPath, "src"),
			Fix:     "Create a 'src' directory with .ae files",
		})
		return result
	}

	importedModules := make(map[string]bool)
	moduleFunctions := make(map[string]map[string]bool)

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			result.Errors = append(result.Errors, utils.ParseError{
				Message: fmt.Sprintf("Failed to read %s: %v", file, err),
				File:    file,
				Fix:     "Check file permissions and ensure the file exists",
			})
			continue
		}

		l := lexer.NewLexer(string(content))
		p := parser.NewParser(l)
		ast := p.Parse()

		if len(p.Errors.Errors) > 0 {
			for _, err := range p.Errors.Errors {
				// Copy the rich error context from the parser
				richError := utils.ParseError{
					Kind:          err.Kind,
					Message:       err.Message,
					Line:          err.Line,
					Column:        err.Column,
					File:          file,
					Snippet:       err.Snippet,
					Caret:         err.Caret,
					Fix:           err.Fix,
					SpecReference: err.SpecReference,
					CodemodPrompt: err.CodemodPrompt,
				}
				result.Errors = append(result.Errors, richError)
			}
			continue
		}

		// Extract imports and functions from AST
		currentModuleName := strings.TrimSuffix(filepath.Base(file), ".ae")
		moduleFunctions[currentModuleName] = make(map[string]bool)

		for _, stmt := range ast.Statements {
			if importStmt, ok := stmt.(*parser.Import); ok {
				importedModuleName := importStmt.Name.Value
				if importStmt.As != nil && importStmt.As.Value != "" {
					importedModuleName = importStmt.As.Value
				}
				importedModules[importedModuleName] = true
			}

			if assign, ok := stmt.(*parser.Assignment); ok {
				if len(assign.Names) > 0 {
					moduleFunctions[currentModuleName][assign.Names[0].Value] = true
				}
			}

			if funcDecl, ok := stmt.(*parser.Function); ok {
				moduleFunctions[currentModuleName][funcDecl.Name.Value] = true
			}
		}
	}

	// Validate imports against declared dependencies
	for moduleName := range importedModules {
		if depPath, exists := config.Dependencies[moduleName]; exists {
			// Check if the dependency file exists
			fullDepPath := filepath.Join(projectPath, depPath)
			if _, err := os.Stat(fullDepPath); os.IsNotExist(err) {
				result.MissingDeps = append(result.MissingDeps, moduleName)
				result.Valid = false
				result.Errors = append(result.Errors, utils.ParseError{
					Message: fmt.Sprintf("Dependency %s not found at %s", moduleName, depPath),
					File:    filepath.Join(projectPath, "aether.toml"),
					Fix:     fmt.Sprintf("Add '%s = \"%s\"' to dependencies in aether.toml", moduleName, depPath),
				})
			} else {
				result.ResolvedDeps[moduleName] = fullDepPath
			}
		} else {
			result.Valid = false
			// Read the actual aether.toml content for the snippet
			content, _ := os.ReadFile(filepath.Join(projectPath, "aether.toml"))
			snippet := ""
			if len(content) > 0 {
				lines := strings.Split(string(content), "\n")
				// Find the dependencies section
				for _, line := range lines {
					if strings.Contains(line, "dependencies") {
						snippet = line
						break
					}
				}
			}
			result.Errors = append(result.Errors, utils.ParseError{
				Message:       fmt.Sprintf("Import '%s' not declared in dependencies", moduleName),
				File:          filepath.Join(projectPath, "aether.toml"),
				Snippet:       snippet,
				Caret:         len(snippet) + 1, // Caret at end of line
				Fix:           fmt.Sprintf("Add '%s = \"path/to/module\"' to dependencies in aether.toml", moduleName),
				CodemodPrompt: fmt.Sprintf("Do you want to add '%s' to dependencies? (y/n)", moduleName),
			})
		}
	}

	// Validate function calls against available functions
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		l := lexer.NewLexer(string(content))
		p := parser.NewParser(l)
		ast := p.Parse()

		if len(p.Errors.ToMessages()) > 0 {
			continue
		}

		// Check all function calls in this file
		validateFunctionCalls(ast, moduleFunctions, result)
	}

	// Check for unused dependencies
	for depName := range config.Dependencies {
		if !importedModules[depName] {
			result.UnusedDeps = append(result.UnusedDeps, depName)
			result.Warnings = append(result.Warnings, fmt.Sprintf("Unused dependency: %s", depName))
		}
	}

	// Validate lock file consistency
	if len(lockFile.Dependencies) > 0 {
		for depName, lockInfo := range lockFile.Dependencies {
			if configPath, exists := config.Dependencies[depName]; exists {
				if lockInfo.Path != configPath {
					result.Warnings = append(result.Warnings, fmt.Sprintf("Lock file mismatch for %s: expected %s, got %s", depName, configPath, lockInfo.Path))
				}
			} else {
				result.Warnings = append(result.Warnings, fmt.Sprintf("Lock file contains undeclared dependency: %s", depName))
			}
		}
	}

	return result
}

func GenerateLockFile(projectPath string) error {
	configPath := filepath.Join(projectPath, "aether.toml")
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read aether.toml: %v", err)
	}

	var config struct {
		Dependencies map[string]string `toml:"dependencies"`
	}

	if err := toml.Unmarshal(configData, &config); err != nil {
		return fmt.Errorf("failed to parse aether.toml: %v", err)
	}

	lockFile := LockFile{
		Dependencies: make(map[string]DependencyInfo),
	}

	for depName, depPath := range config.Dependencies {
		lockFile.Dependencies[depName] = DependencyInfo{
			Path: depPath,
		}
	}

	lockData, err := toml.Marshal(lockFile)
	if err != nil {
		return fmt.Errorf("failed to marshal lock file: %v", err)
	}

	lockPath := filepath.Join(projectPath, "aether.lock")
	return os.WriteFile(lockPath, lockData, 0644)
}

func analyzeFile(filePath string, result *AnalysisResult) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		result.Errors = append(result.Errors, utils.ParseError{Message: fmt.Sprintf("Failed to read file %s: %v", filePath, err)})
		return
	}

	// Parse the file
	lexer := lexer.NewLexer(string(content))
	parser := parser.NewParser(lexer)
	ast := parser.Parse()

	if len(parser.Errors.Errors) > 0 {
		for _, err := range parser.Errors.Errors {
			result.Errors = append(result.Errors, utils.ParseError{Message: fmt.Sprintf("%s: %s", filePath, utils.FormatErrorWithContext(err))})
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
	case *parser.CComment:
		analyzeCComment(s, filePath, result)
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
		result.Errors = append(result.Errors, utils.ParseError{Message: fmt.Sprintf("%s: Invalid import path '%s'", filePath, importPath)})
	}

	// First try to resolve using dependency configuration
	projectRoot := findProjectRoot(filepath.Dir(filePath))
	configPath := filepath.Join(projectRoot, "aether.toml")
	
	if data, err := os.ReadFile(configPath); err == nil {
		var config struct {
			Dependencies map[string]string `toml:"dependencies"`
		}
		if err := toml.Unmarshal(data, &config); err == nil {
			if depPath, exists := config.Dependencies[importPath]; exists {
				fullDepPath := filepath.Join(projectRoot, depPath)
				if _, err := os.Stat(fullDepPath); err == nil {
					importInfo.Exists = true
					importInfo.Resolved = fullDepPath
					result.Imports[importPath] = importInfo
					return
				}
			}
		}
	}

	// Fallback to direct file resolution
	resolvedPath := resolveImportPath(importPath, filepath.Dir(filePath))
	if resolvedPath != "" {
		if _, err := os.Stat(resolvedPath); err == nil {
			importInfo.Exists = true
			importInfo.Resolved = resolvedPath
		} else {
			importInfo.Errors = append(importInfo.Errors, "Imported file does not exist")
			result.Errors = append(result.Errors, utils.ParseError{Message: fmt.Sprintf("%s: Imported file '%s' does not exist", filePath, importPath)})
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
		Name:     assign.Names[0].Value,
		Type:     assign.Names[0].Type,
		Defined:  true,
		Used:     false,
		Exported: isExported(assign.Names[0].Value),
		Scope:    "local",
	}

	result.Variables[assign.Names[0].Value] = varInfo
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
		typeInfo.Fields[field.Name.Value] = field.Type
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
			// Enhanced error: function not found in any imported binding/module
			result.Undefined = append(result.Undefined, fmt.Sprintf("function '%s'", funcName))
			result.Errors = append(result.Errors, utils.ParseError{
				Kind:    utils.UndefinedReference,
				Message: "undefined reference to function '" + funcName + "' in imported binding or file. (Aether imports are not modules unless HMR is enabled.) Check spelling or update your binding.",
				File:    filePath,
				Fix:     "Check the spelling or update your binding.",
			})
		}
	}

	for _, arg := range call.Args {
		analyzeExpression(arg, filePath, result)
	}
}

func analyzeAssignment(assign *parser.Assignment, filePath string, result *AnalysisResult) {
	if len(assign.Names) > 0 {
		if varInfo, exists := result.Variables[assign.Names[0].Value]; exists {
			varInfo.Used = true
			result.Variables[assign.Names[0].Value] = varInfo
		} else {
			result.Undefined = append(result.Undefined, fmt.Sprintf("variable '%s'", assign.Names[0].Value))
		}

		analyzeExpression(assign.Value, filePath, result)
	}
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

func analyzeCComment(comment *parser.CComment, filePath string, result *AnalysisResult) {
	includes := parseCIncludes(comment.Content)
	result.CIncludes = append(result.CIncludes, includes...)
}

func parseCIncludes(content string) []CInclude {
	var includes []CInclude

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "//") {
			line = strings.TrimSpace(line[2:])
		}

		if strings.HasPrefix(line, "#include") {
			include := parseIncludeDirective(line)
			if include != nil {
				includes = append(includes, *include)
			}
		}
	}

	return includes
}

func parseIncludeDirective(line string) *CInclude {
	re := regexp.MustCompile(`#include\s*[<"]([^>"]+)[>"]`)
	matches := re.FindStringSubmatch(line)
	if len(matches) < 2 {
		return nil
	}

	header := matches[1]
	isSystem := strings.Contains(line, "<") && strings.Contains(line, ">")

	return &CInclude{
		Header:   header,
		IsSystem: isSystem,
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
				result.Errors = append(result.Errors, utils.ParseError{Message: "Circular import detected"})
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
			result.Errors = append(result.Errors, utils.ParseError{Message: fmt.Sprintf("Undefined reference: %s", item)})
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
		return importPath + ".aeth"
	}

	if strings.HasPrefix(importPath, ".") {
		return filepath.Join(currentDir, importPath+".aeth")
	}

	return filepath.Join(currentDir, importPath+".aeth")
}

func isExported(name string) bool {
	return len(name) > 0 && name[0] >= 'A' && name[0] <= 'Z'
}

func isStdlibFunction(name string) bool {
	// Stdlib functions are now defined in actual Aether files
	// This function is kept for backward compatibility but always returns false
	// The real stdlib functions are in lib/core/*.ae files
	return false
}

func findProjectRoot(start string) string {
	dir := start
	for {
		configPath := filepath.Join(dir, "aether.toml")
		if _, err := os.Stat(configPath); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "."
}

func findAetherFiles(root string) ([]string, error) {
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

func validateFunctionCalls(ast *parser.Program, moduleFunctions map[string]map[string]bool, result *DependencyAnalysis) {
	for _, stmt := range ast.Statements {
		validateStatementFunctionCalls(stmt, moduleFunctions, result)
	}
}

func validateStatementFunctionCalls(stmt parser.Statement, moduleFunctions map[string]map[string]bool, result *DependencyAnalysis) {
	switch s := stmt.(type) {
	case *parser.Block:
		for _, subStmt := range s.Statements {
			validateStatementFunctionCalls(subStmt, moduleFunctions, result)
		}
	case *parser.If:
		validateStatementFunctionCalls(s.Consequence, moduleFunctions, result)
		if s.Alternative != nil {
			validateStatementFunctionCalls(s.Alternative, moduleFunctions, result)
		}
	case *parser.While:
		validateStatementFunctionCalls(s.Body, moduleFunctions, result)
	case *parser.For:
		validateStatementFunctionCalls(s.Body, moduleFunctions, result)
	case *parser.Function:
		validateStatementFunctionCalls(s.Body, moduleFunctions, result)
	}

	// Check for function calls in expressions
	validateExpressionFunctionCalls(stmt, moduleFunctions, result)
}

func validateExpressionFunctionCalls(expr interface{}, moduleFunctions map[string]map[string]bool, result *DependencyAnalysis) {
	switch e := expr.(type) {
	case *parser.Call:
		if ident, ok := e.Function.(*parser.Identifier); ok {
			if !isStdlibFunction(ident.Value) {
				// Check if function exists in any module
				found := false
				for moduleName, functions := range moduleFunctions {
					_ = moduleName
					if functions[ident.Value] {
						found = true
						break
					}
				}
				if !found {
					result.Valid = false
					result.Errors = append(result.Errors, utils.ParseError{
						Kind:    utils.UndefinedReference,
						Message: "undefined reference to function '" + ident.Value + "' in imported binding or file. (Aether imports are not modules unless HMR is enabled.) Check spelling or update your binding.",
						Fix:     "Check the spelling or update your binding.",
					})
				}
			}
		}
	case *parser.PropertyAccess:
		if moduleIdent, ok := e.Object.(*parser.Identifier); ok {
			// Check if module.property exists
			if functions, exists := moduleFunctions[moduleIdent.Value]; exists {
				if !functions[e.Property.Value] {
					result.Valid = false
					result.Errors = append(result.Errors, utils.ParseError{
						Kind:    utils.UndefinedReference,
						Message: "undefined reference to function '" + e.Property.Value + "' in binding or file '" + moduleIdent.Value + "'. (Aether imports are not modules unless HMR is enabled.) Check spelling or update your binding.",
						Fix:     "Check the spelling or update your binding.",
					})
				}
			} else {
				result.Valid = false
				result.Errors = append(result.Errors, utils.ParseError{
					Kind:    utils.UndefinedReference,
					Message: "undefined reference to module '" + moduleIdent.Value + "'. (Aether imports are not modules unless HMR is enabled.) Check import or binding.",
					Fix:     "Check the import or binding for the module.",
				})
			}
		}
	}
}
