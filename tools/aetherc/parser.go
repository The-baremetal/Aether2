package aetherc

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parseCHeader(filePath string) (*CHeader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open header file: %v", err)
	}
	defer file.Close()

	header := &CHeader{
		Name:      extractHeaderName(filePath),
		Path:      filePath,
		Functions: []CFunction{},
		Constants: []CConstant{},
		Types:     []CType{},
		Libraries: []string{},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if function := parseFunctionDeclaration(line); function != nil {
			header.Functions = append(header.Functions, *function)
		}

		if constant := parseConstantDeclaration(line); constant != nil {
			header.Constants = append(header.Constants, *constant)
		}

		if ctype := parseTypeDeclaration(line); ctype != nil {
			header.Types = append(header.Types, *ctype)
		}
	}

	return header, scanner.Err()
}

func extractHeaderName(filePath string) string {
	parts := strings.Split(filePath, "/")
	if len(parts) > 0 {
		return strings.TrimSuffix(parts[len(parts)-1], ".h")
	}
	return "unknown"
}

func parseFunctionDeclaration(line string) *CFunction {
	functionRegex := regexp.MustCompile(`^(\w+)\s+(\w+)\s*\(([^)]*)\)\s*;?\s*(//\s*(.+))?$`)
	matches := functionRegex.FindStringSubmatch(line)

	if len(matches) < 3 {
		return nil
	}

	function := &CFunction{
		ReturnType: matches[1],
		Name:       matches[2],
		Parameters: parseParameters(matches[3]),
	}

	if len(matches) > 5 {
		function.Comment = matches[5]
	}

	return function
}

func parseParameters(paramStr string) []CParameter {
	if strings.TrimSpace(paramStr) == "void" || paramStr == "" {
		return []CParameter{}
	}

	params := []CParameter{}
	paramList := strings.Split(paramStr, ",")

	for _, param := range paramList {
		parts := strings.Fields(strings.TrimSpace(param))
		if len(parts) >= 2 {
			paramType := strings.Join(parts[:len(parts)-1], " ")
			paramName := parts[len(parts)-1]

			params = append(params, CParameter{
				Name: paramName,
				Type: paramType,
			})
		}
	}

	return params
}

func parseConstantDeclaration(line string) *CConstant {
	constantRegex := regexp.MustCompile(`^#define\s+(\w+)\s+(.+)$`)
	matches := constantRegex.FindStringSubmatch(line)

	if len(matches) < 3 {
		return nil
	}

	return &CConstant{
		Name:  matches[1],
		Value: strings.TrimSpace(matches[2]),
		Type:  "int",
	}
}

func parseTypeDeclaration(line string) *CType {
	typedefRegex := regexp.MustCompile(`^typedef\s+struct\s+(\w+)\s*\{`)
	matches := typedefRegex.FindStringSubmatch(line)

	if len(matches) < 2 {
		return nil
	}

	return &CType{
		Name:   matches[1],
		Type:   "struct",
		Fields: []CField{},
	}
}
