package aetherc

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func generateAetherBindings(binding *CBinding, outputPath string) error {
	tmpl, err := template.New("aether_binding").Parse(aetherBindingTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer file.Close()

	return tmpl.Execute(file, binding)
}

func generateFunctionWrapper(function CFunction) string {
	var params []string
	var paramNames []string

	for _, param := range function.Parameters {
		aetherType := mapCTypeToAether(param.Type)
		params = append(params, fmt.Sprintf("%s %s", param.Name, aetherType))
		paramNames = append(paramNames, param.Name)
	}

	returnType := mapCTypeToAether(function.ReturnType)

	return fmt.Sprintf(`func %s(%s) %s {
	return C.%s(%s)
}`, function.Name, strings.Join(params, ", "), returnType, function.Name, strings.Join(paramNames, ", "))
}

func mapCTypeToAether(cType string) string {
	switch strings.ToLower(cType) {
	case "int", "int32_t":
		return "int"
	case "long", "int64_t":
		return "long"
	case "float":
		return "float"
	case "double":
		return "double"
	case "char*", "const char*":
		return "string"
	case "void":
		return ""
	default:
		return "any"
	}
}

const aetherBindingTemplate = `package {{.PackageName}}

import "C"

{{range .Functions}}
{{if .Comment}}
{{.Comment}}
{{end}}
{{generateFunctionWrapper .}}

{{end}}

{{range .Constants}}
const {{.Name}} = {{.Value}}
{{end}}

{{range .Types}}
type {{.Name}} struct {
{{range .Fields}}
	{{.Name}} {{mapCTypeToAether .Type}}
{{end}}
}
{{end}}
`
