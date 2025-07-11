package analysis

import (
	"aether/src/parser"
	"regexp"
	"strings"
)

func ExtractCIncludes(node *parser.ASTNode, result *AnalysisResult) {
	if node == nil {
		return
	}

	if node.NodeKind == parser.CCommentKind {
		if content, ok := node.Value.(string); ok {
			includes := ParseCIncludes(content)
			result.CIncludes = append(result.CIncludes, includes...)
		}
	}

	for _, inner := range node.Inner {
		ExtractCIncludes(inner, result)
	}

	if node.Left != nil {
		ExtractCIncludes(node.Left, result)
	}

	if node.Right != nil {
		ExtractCIncludes(node.Right, result)
	}

	if node.Body != nil {
		ExtractCIncludes(node.Body, result)
	}

	for _, param := range node.Params {
		ExtractCIncludes(param, result)
	}
}

func ParseCIncludes(content string) []CInclude {
	var includes []CInclude

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "//") {
			line = strings.TrimSpace(line[2:])
		}

		if strings.HasPrefix(line, "#include") {
			include := ParseIncludeDirective(line)
			if include != nil {
				includes = append(includes, *include)
			}
		}
	}

	return includes
}

func ParseIncludeDirective(line string) *CInclude {
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

func AnalyzeCComment(comment *parser.CComment, filePath string, result *AnalysisResult) {
	includes := ParseCIncludes(comment.Content)
	result.CIncludes = append(result.CIncludes, includes...)
}
