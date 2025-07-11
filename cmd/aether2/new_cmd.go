package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var NewCmd = &cobra.Command{
	Use:   "new <package_name>",
	Short: "Create a new package with one command",
	Long: `Create a new Aether package instantly!

This command creates a new package in the packages/ directory with a simple template.
No questions, no complex setup - just one file!

Examples:
  aether new math      # Creates packages/math.aeth
  aether new utils     # Creates packages/utils.aeth
  aether new http      # Creates packages/http.aeth`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		packageName := args[0]
		createSimplePackage(packageName)
	},
}

func createSimplePackage(packageName string) {
	// Validate package name
	if !isValidPackageName(packageName) {
		fmt.Printf("🍕 Error: Invalid package name '%s'\n", packageName)
		fmt.Println("🍕 Package names must be lowercase letters, numbers, and hyphens only")
		os.Exit(1)
	}

	// Check if we're in an Aether project
	projectRoot := findProjectRoot(".")
	if projectRoot == "." {
		fmt.Println("🍕 Error: Not in an Aether project. Run 'aether init' first!")
		os.Exit(1)
	}

	// Create packages directory if it doesn't exist
	packagesDir := filepath.Join(projectRoot, "packages")
	if err := os.MkdirAll(packagesDir, 0755); err != nil {
		fmt.Printf("🍕 Error: Could not create packages directory: %v\n", err)
		os.Exit(1)
	}

	// Check if package already exists
	packagePath := filepath.Join(packagesDir, packageName+".aeth")
	if _, err := os.Stat(packagePath); err == nil {
		fmt.Printf("🍕 Error: Package '%s' already exists at %s\n", packageName, packagePath)
		os.Exit(1)
	}

	// Create the package file
	packageContent := generatePackageTemplate(packageName)
	if err := os.WriteFile(packagePath, []byte(packageContent), 0644); err != nil {
		fmt.Printf("🍕 Error: Could not create package file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("🍕 Created package '%s' at %s\n", packageName, packagePath)
	fmt.Printf("🍕 Package is ready to use! Import it with: import \"%s\"\n", packageName)
	fmt.Printf("🍕 Run 'aether build' to test your package!\n")
}

func isValidPackageName(name string) bool {
	if name == "" {
		return false
	}

	// Check for valid characters: lowercase letters, numbers, hyphens
	for _, char := range name {
		if !((char >= 'a' && char <= 'z') ||
			(char >= '0' && char <= '9') ||
			char == '-') {
			return false
		}
	}

	// Check for valid start/end
	if name[0] == '-' || name[len(name)-1] == '-' {
		return false
	}

	// Check for consecutive hyphens
	if strings.Contains(name, "--") {
		return false
	}

	return true
}

func generatePackageTemplate(packageName string) string {
	// Convert package name to title case for function names
	titleCase := strings.Title(strings.ReplaceAll(packageName, "-", "_"))

	return fmt.Sprintf(`package %s

# %s package
# Created with: aether new %s

# Example function - you can delete this
func example_function() {
    return "Hello from %s package!"
}

# Add your functions here:
# func add(a, b) {
#     return a + b
# }
#
# func multiply(a, b) {
#     return a * b
# }

# Remember: Capitalized function names are exported (can be imported)
# Lowercase function names are internal (cannot be imported)
`, packageName, titleCase, packageName, titleCase)
}
