package analysis

import (
	"fmt"
	"strings"

	"aether/lib/utils"
)

func TopoSort(files []string, imports map[string][]string) ([]string, error) {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)
	var result []string

	for _, file := range files {
		if !visited[file] {
			if !topoSortDFS(file, imports, visited, recStack, &result) {
				return nil, fmt.Errorf("circular dependency detected")
			}
		}
	}

	return result, nil
}

func topoSortDFS(file string, imports map[string][]string, visited, recStack map[string]bool, result *[]string) bool {
	visited[file] = true
	recStack[file] = true

	for _, dep := range imports[file] {
		if !visited[dep] {
			if !topoSortDFS(dep, imports, visited, recStack, result) {
				return false
			}
		} else if recStack[dep] {
			return false
		}
	}

	recStack[file] = false
	*result = append(*result, file)
	return true
}

func DetectCycles(imports map[string][]string) [][]string {
	var cycles [][]string
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	for file := range imports {
		if !visited[file] {
			pathStack := []string{}
			if detectCycle(file, imports, visited, recStack, pathStack) {
				cycles = append(cycles, pathStack)
			}
		}
	}

	return cycles
}

func detectCycle(path string, deps map[string][]string, visited, recStack map[string]bool, pathStack []string) bool {
	visited[path] = true
	recStack[path] = true
	pathStack = append(pathStack, path)

	for _, dep := range deps[path] {
		if !visited[dep] {
			if detectCycle(dep, deps, visited, recStack, pathStack) {
				return true
			}
		} else if recStack[dep] {
			return true
		}
	}

	recStack[path] = false
	return false
}

func ValidateImports(result *AnalysisResult) {
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

func CheckDependencies(result *AnalysisResult) {
	for path, importInfo := range result.Imports {
		if importInfo.Exists && importInfo.Resolved != "" {
			result.Dependencies[path] = []string{importInfo.Resolved}
		}
	}
}

func DetectCyclesInResult(result *AnalysisResult) {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	for path := range result.Dependencies {
		if !visited[path] {
			if detectCycle(path, result.Dependencies, visited, recStack, []string{}) {
				result.Cycles = append(result.Cycles, []string{})
				result.Valid = false
				result.Errors = append(result.Errors, utils.ParseError{Message: "Circular import detected"})
			}
		}
	}
}

func IsValidImportPath(path string) bool {
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

func ResolveImportPath(importPath, currentDir string) string {
	if strings.HasPrefix(importPath, "/") {
		return importPath + ".ae"
	}

	if strings.HasPrefix(importPath, ".") {
		return currentDir + "/" + importPath + ".ae"
	}

	return currentDir + "/" + importPath + ".ae"
}
