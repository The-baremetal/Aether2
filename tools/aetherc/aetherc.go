package aetherc

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]
	switch command {
	case "generate":
		if len(os.Args) < 4 {
			fmt.Println("Usage: aetherc generate <header_file> <output.ae>")
			fmt.Println("Example: aetherc generate /usr/include/stdio.h packages/c/src/stdio.ae")
			return
		}
		generateBindings(os.Args[2], os.Args[3])
	case "parse":
		if len(os.Args) < 3 {
			fmt.Println("Usage: aetherc parse <header_file>")
			return
		}
		parseHeader(os.Args[2])
	case "list":
		if len(os.Args) < 3 {
			fmt.Println("Usage: aetherc list <header_file>")
			return
		}
		listFunctions(os.Args[2])
	case "scan":
		if len(os.Args) < 3 {
			fmt.Println("Usage: aetherc scan <directory>")
			return
		}
		scanDirectory(os.Args[2])
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}
}

func printUsage() {
	fmt.Println("🍕 aetherc - Dynamic C Binding Generator for Aether")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  generate <header> <output.ae>  Parse C header and generate Aether bindings")
	fmt.Println("  parse <header>                  Parse and display C header contents")
	fmt.Println("  list <header>                   List functions in C header")
	fmt.Println("  scan <directory>                Scan directory for C headers")
	fmt.Println("  help                            Show this help")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  aetherc generate /usr/include/stdio.h packages/c/src/stdio.ae")
	fmt.Println("  aetherc generate /usr/include/math.h packages/c/src/math.ae")
	fmt.Println("  aetherc generate /usr/include/llvm-c/Core.h packages/llvm/src/core.ae")
	fmt.Println("  aetherc scan /usr/include")
}

func generateBindings(headerPath, outputPath string) {
	fmt.Printf("🍕 Parsing C header: %s\n", headerPath)

	header, err := parseCHeader(headerPath)
	if err != nil {
		fmt.Printf("❌ Failed to parse header: %s - %v\n", headerPath, err)
		os.Exit(1)
	}

	packageName := extractPackageName(outputPath)
	binding := CBinding{
		PackageName: packageName,
		Headers:     []CHeader{*header},
		Functions:   header.Functions,
		Constants:   header.Constants,
		Types:       header.Types,
		Libraries:   header.Libraries,
		Includes:    []string{header.Name},
	}

	fmt.Printf("📦 Generating bindings for package: %s\n", packageName)
	fmt.Printf("🔧 Found %d functions, %d constants, %d types\n",
		len(header.Functions), len(header.Constants), len(header.Types))

	generateBindingFile(binding, outputPath)
	fmt.Printf("✅ Generated bindings: %s\n", outputPath)
}

func parseLibraryDependency(line string) string {
	// Match library dependencies in comments or pragmas
	if strings.Contains(line, "#pragma comment(lib,") {
		libRegex := regexp.MustCompile(`#pragma comment\(lib,\s*"([^"]+)"\)`)
		matches := libRegex.FindStringSubmatch(line)
		if len(matches) >= 2 {
			return strings.TrimSpace(matches[1])
		}
	}

	return ""
}

func extractComment(line string) string {
	// Extract comments from the line
	if idx := strings.Index(line, "//"); idx != -1 {
		return strings.TrimSpace(line[idx+2:])
	}
	return ""
}

func extractPackageName(outputPath string) string {
	// Extract package name from output path
	// e.g., "packages/c/src/stdio.ae" -> "c"
	parts := strings.Split(outputPath, string(os.PathSeparator))
	for i, part := range parts {
		if part == "packages" && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return "c"
}

func generateBindingFile(binding CBinding, outputPath string) {
	// Create directory if it doesn't exist
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("❌ Failed to create directory: %v\n", err)
		return
	}

	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("❌ Failed to create output file: %v\n", err)
		return
	}
	defer file.Close()

	tmpl := template.Must(template.New("binding").Parse(bindingTemplate))
	if err := tmpl.Execute(file, binding); err != nil {
		fmt.Printf("❌ Failed to generate binding file: %v\n", err)
		return
	}
}

func parseHeader(headerPath string) {
	header, err := parseCHeader(headerPath)
	if err != nil {
		fmt.Printf("❌ Failed to parse header: %s - %v\n", headerPath, err)
		return
	}

	fmt.Printf("📋 Header: %s\n", header.Name)
	fmt.Printf("📁 Path: %s\n", header.Path)
	fmt.Printf("🔧 Functions: %d\n", len(header.Functions))
	fmt.Printf("📊 Constants: %d\n", len(header.Constants))
	fmt.Printf("🏗️  Types: %d\n", len(header.Types))
	fmt.Printf("📚 Libraries: %v\n", header.Libraries)

	if len(header.Functions) > 0 {
		fmt.Println("\n🔧 Functions:")
		for _, fn := range header.Functions {
			fmt.Printf("  %s %s(%s)\n", fn.ReturnType, fn.Name, formatParameters(fn.Parameters))
		}
	}

	if len(header.Constants) > 0 {
		fmt.Println("\n📊 Constants:")
		for _, c := range header.Constants {
			fmt.Printf("  %s = %s\n", c.Name, c.Value)
		}
	}
}

func listFunctions(headerPath string) {
	header, err := parseCHeader(headerPath)
	if err != nil {
		fmt.Printf("❌ Failed to parse header: %s - %v\n", headerPath, err)
		return
	}

	fmt.Printf("🔧 Functions in %s:\n", header.Name)
	for _, fn := range header.Functions {
		fmt.Printf("  %s %s(%s)\n", fn.ReturnType, fn.Name, formatParameters(fn.Parameters))
	}
}

func scanDirectory(dirPath string) {
	fmt.Printf("🔍 Scanning directory: %s\n", dirPath)

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".h") {
			header, err := parseCHeader(path)
			if err == nil && len(header.Functions) > 0 {
				fmt.Printf("📦 %s: %d functions\n", filepath.Base(path), len(header.Functions))
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("❌ Error scanning directory: %v\n", err)
	}
}

func formatParameters(params []CParameter) string {
	if len(params) == 0 {
		return "void"
	}

	parts := make([]string, len(params))
	for i, param := range params {
		if param.Name == "..." {
			parts[i] = "..."
		} else {
			parts[i] = fmt.Sprintf("%s %s", param.Type, param.Name)
		}
	}
	return strings.Join(parts, ", ")
}

const bindingTemplate = `// Aether C Bindings
// Generated by aetherc from {{range .Headers}}{{.Name}} {{end}}
// Package: {{.PackageName}}

{{range .Headers}}
// #include <{{.Name}}>
{{end}}

{{range .Functions}}
func {{.Name}}({{range $i, $param := .Parameters}}{{if $i}}, {{end}}{{$param.Name}}{{end}}) {
    return c_{{.Name}}({{range $i, $param := .Parameters}}{{if $i}}, {{end}}{{$param.Name}}{{end}})
}
{{end}}

{{range .Constants}}
const {{.Name}} = {{.Value}}
{{end}}

{{range .Types}}
struct {{.Name}} {
    {{range .Fields}}
    {{.Type}} {{.Name}}
    {{end}}
}
{{end}}

{{if .Libraries}}
// Libraries: {{range .Libraries}}-l{{.}} {{end}}
{{end}}
`
