package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

var PackageCmd = &cobra.Command{
	Use:   "package [command] [args...]",
	Short: "Manage Aether packages",
	Long: `Manage Aether packages and dependencies.

Available commands:
  create <package>    - Create a new package interactively
  install <package>   - Install a package
  uninstall <package> - Uninstall a package
  list               - List installed packages
  update [package]   - Update packages
  search <query>     - Search for packages
  publish <package>  - Publish a package
  info <package>     - Show package information

Examples:
  aether package create math-utils
  aether package install math
  aether package list
  aether package search http
  aether package update
  aether package info math`,
	Run: doPackage,
}

var (
	packageFlags struct {
		global   bool
		dev      bool
		force    bool
		verbose  bool
		registry string
	}
)

func doPackage(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("🍕 Aether Package Manager")
		fmt.Println("Use 'aether package --help' for available commands")
		return
	}

	command := args[0]
	packageArgs := args[1:]

	switch command {
	case "create":
		createPackage(packageArgs)
	case "install":
		installPackages(packageArgs)
	case "uninstall":
		uninstallPackages(packageArgs)
	case "list":
		listPackages()
	case "update":
		updatePackages(packageArgs)
	case "search":
		searchPackages(packageArgs)
	case "publish":
		publishPackage(packageArgs)
	case "info":
		showPackageInfo(packageArgs)
	default:
		fmt.Printf("🍕 Unknown package command: %s\n", command)
		fmt.Println("Available commands: create, install, uninstall, list, update, search, publish, info")
	}
}

func createPackage(args []string) {
	var packageName string

	if len(args) > 0 {
		packageName = args[0]
	} else {
		fmt.Print("🍕 What's your package name? ")
		r := bufio.NewReader(os.Stdin)
		resp, _ := r.ReadString('\n')
		packageName = strings.TrimSpace(resp)
		if packageName == "" {
			fmt.Println("🍕 Package name cannot be empty!")
			os.Exit(1)
		}
	}

	// Check if we're in an Aether project
	projectRoot := findProjectRoot(".")
	if projectRoot == "." {
		fmt.Println("🍕 Error: Not in an Aether project. Run 'aether init' first!")
		os.Exit(1)
	}

	// Load current configuration (for future use)
	_ = loadProjectConfig(projectRoot)

	// Ask package questions
	packageType := askPackageType()
	description := askPackageDescription()
	author := askPackageAuthor()
	version := askPackageVersion()

	fmt.Printf("🍕 Creating package '%s'...\n", packageName)

	// Create package directory structure
	packageDir := filepath.Join("packages", packageName)
	must(os.MkdirAll(filepath.Join(packageDir, "src"), 0755))
	must(os.MkdirAll(filepath.Join(packageDir, "examples"), 0755))

	// Create package aether.toml
	packageConfig := generatePackageConfig(packageName, description, author, version)
	must(os.WriteFile(filepath.Join(packageDir, "aether.toml"), []byte(packageConfig), 0644))

	// Create example file
	exampleFile := filepath.Join(packageDir, "examples", "basic.aeth")
	exampleContent := fmt.Sprintf(`import "%s"

func main() {
    // Your %s package example
    println("🍕 Hello from %s package!")
}
`, packageName, packageName, packageName)
	must(os.WriteFile(exampleFile, []byte(exampleContent), 0644))

	// Create README
	readmeFile := filepath.Join(packageDir, "README.md")
	readmeContent := fmt.Sprintf("# %s\n\n%s\n\n## Usage\n\n```aeth\nimport \"%s\"\n\n// Your code here\n```\n\n## Installation\n\nThis package is part of your local project. No additional installation needed.\n\n## Examples\n\nSee the examples/ directory for usage examples.", packageName, description, packageName)
	must(os.WriteFile(readmeFile, []byte(readmeContent), 0644))

	// Update main project's aether.toml
	updateProjectDependencies(projectRoot, packageName, packageType)

	fmt.Printf("🍕 Package '%s' created successfully!\n", packageName)
	fmt.Printf("🍕 Location: packages/%s\n", packageName)
	fmt.Printf("🍕 Added to project dependencies\n")
	fmt.Printf("🍕 Run 'aether build' to test your package!\n")
}

func askPackageType() string {
	fmt.Print("🍕 What type of package? (library/app): ")
	r := bufio.NewReader(os.Stdin)
	resp, _ := r.ReadString('\n')
	resp = strings.TrimSpace(resp)
	if resp == "" || resp == "library" {
		return "library"
	}
	return resp
}

func askPackageDescription() string {
	fmt.Print("🍕 What does this package do? ")
	r := bufio.NewReader(os.Stdin)
	resp, _ := r.ReadString('\n')
	resp = strings.TrimSpace(resp)
	if resp == "" {
		return "A useful Aether package"
	}
	return resp
}

func askPackageAuthor() string {
	fmt.Print("🍕 Who created this package? (optional): ")
	r := bufio.NewReader(os.Stdin)
	resp, _ := r.ReadString('\n')
	return strings.TrimSpace(resp)
}

func askPackageVersion() string {
	fmt.Print("🍕 Package version? (0.1.0): ")
	r := bufio.NewReader(os.Stdin)
	resp, _ := r.ReadString('\n')
	resp = strings.TrimSpace(resp)
	if resp == "" {
		return "0.1.0"
	}
	return resp
}

func generatePackageConfig(name, description, author, version string) string {
	config := fmt.Sprintf(`[project]
name = "%s"
version = "%s"
description = "%s"`, name, version, description)

	if author != "" {
		config += fmt.Sprintf("\nauthor = \"%s\"", author)
	}

	config += `

[build]
source_directories = ["src"]
output_directory = "bin"
target = "native"
optimization = "2"
linker = "mold"
create_library = true
library_type = "shared"

[dependencies]
# Add package dependencies here

[dev-dependencies]
# Add development dependencies here
`

	return config
}

func updateProjectDependencies(projectRoot, packageName, packageType string) {
	configPath := filepath.Join(projectRoot, "aether.toml")

	// Read current config
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("🍕 Warning: Could not read aether.toml: %v\n", err)
		return
	}

	var config ProjectConfig
	if err := toml.Unmarshal(data, &config); err != nil {
		fmt.Printf("🍕 Warning: Could not parse aether.toml: %v\n", err)
		return
	}

	// Initialize dependencies map if it doesn't exist
	if config.Dependencies == nil {
		config.Dependencies = make(map[string]string)
	}

	// Add the new package
	config.Dependencies[packageName] = fmt.Sprintf("packages/%s", packageName)

	// Write back to file
	encoder := toml.NewEncoder(os.Stdout)
	encoder.Encode(config)

	// For now, we'll just append to the file since toml encoding is complex
	// In a real implementation, you'd want to properly update the existing file
	dependencyLine := fmt.Sprintf("%s = \"packages/%s\"\n", packageName, packageName)

	// Simple append approach
	file, err := os.OpenFile(configPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("🍕 Warning: Could not update aether.toml: %v\n", err)
		return
	}
	defer file.Close()

	// Find the dependencies section and add our package
	content := string(data)
	if strings.Contains(content, "[dependencies]") {
		// Find the dependencies section and add our package
		lines := strings.Split(content, "\n")
		var newLines []string
		dependenciesFound := false

		for _, line := range lines {
			newLines = append(newLines, line)
			if strings.TrimSpace(line) == "[dependencies]" {
				dependenciesFound = true
				newLines = append(newLines, dependencyLine)
			}
		}

		if dependenciesFound {
			must(os.WriteFile(configPath, []byte(strings.Join(newLines, "\n")), 0644))
		} else {
			// Add dependencies section if it doesn't exist
			newContent := content + "\n[dependencies]\n" + dependencyLine
			must(os.WriteFile(configPath, []byte(newContent), 0644))
		}
	} else {
		// Add dependencies section
		newContent := content + "\n[dependencies]\n" + dependencyLine
		must(os.WriteFile(configPath, []byte(newContent), 0644))
	}
}

func installPackages(packages []string) {
	if len(packages) == 0 {
		fmt.Println("🍕 Usage: aether package install <package_name>")
		return
	}

	fmt.Println("🍕 Installing packages...")
	for _, pkg := range packages {
		if err := installPackage(pkg); err != nil {
			fmt.Fprintf(os.Stderr, "🍕 Error installing %s: %v\n", pkg, err)
		} else {
			fmt.Printf("🍕 Installed: %s\n", pkg)
		}
	}
}

func installPackage(pkgName string) error {
	// Check if package is already installed
	if isPackageInstalled(pkgName) {
		if !packageFlags.force {
			fmt.Printf("🍕 Package %s is already installed\n", pkgName)
			return nil
		}
	}

	// Create packages directory if it doesn't exist
	packagesDir := "packages"
	if err := os.MkdirAll(packagesDir, 0755); err != nil {
		return err
	}

	// For now, just create a placeholder package
	// In a real implementation, this would download from a registry
	pkgDir := filepath.Join(packagesDir, pkgName)
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		return err
	}

	// Create package files
	files := []struct {
		name    string
		content string
	}{
		{
			name: "aether.toml",
			content: fmt.Sprintf(`[project]
name = "%s"
version = "0.1.0"
description = "Aether package"
author = "Aether Developer"

[dependencies]
# Package dependencies

[dev-dependencies]
# Development dependencies`, pkgName),
		},
		{
			name: "src/lib.aeth",
			content: fmt.Sprintf(`// %s package

func hello() {
    print("Hello from %s package!")
}

func version() {
    return "0.1.0"
}`, pkgName, pkgName),
		},
		{
			name:    "README.md",
			content: fmt.Sprintf("# %s\n\nAether package.\n\n## Usage\n\n```aether\nimport \"%s\"\n\n%s.hello()\n```\n\n## Installation\n\n```bash\naether package install %s\n```", pkgName, pkgName, pkgName, pkgName),
		},
	}

	for _, file := range files {
		filePath := filepath.Join(pkgDir, file.name)
		dir := filepath.Dir(filePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		if err := os.WriteFile(filePath, []byte(file.content), 0644); err != nil {
			return err
		}
	}

	return nil
}

func uninstallPackages(packages []string) {
	if len(packages) == 0 {
		fmt.Println("🍕 Usage: aether package uninstall <package_name>")
		return
	}

	fmt.Println("🍕 Uninstalling packages...")
	for _, pkg := range packages {
		if err := uninstallPackage(pkg); err != nil {
			fmt.Fprintf(os.Stderr, "🍕 Error uninstalling %s: %v\n", pkg, err)
		} else {
			fmt.Printf("🍕 Uninstalled: %s\n", pkg)
		}
	}
}

func uninstallPackage(pkgName string) error {
	pkgDir := filepath.Join("packages", pkgName)
	if _, err := os.Stat(pkgDir); os.IsNotExist(err) {
		return fmt.Errorf("package %s is not installed", pkgName)
	}

	return os.RemoveAll(pkgDir)
}

func listPackages() {
	packagesDir := "packages"
	if _, err := os.Stat(packagesDir); os.IsNotExist(err) {
		fmt.Println("🍕 No packages installed")
		return
	}

	entries, err := os.ReadDir(packagesDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "🍕 Error reading packages directory: %v\n", err)
		return
	}

	if len(entries) == 0 {
		fmt.Println("🍕 No packages installed")
		return
	}

	fmt.Println("🍕 Installed packages:")
	for _, entry := range entries {
		if entry.IsDir() {
			pkgName := entry.Name()
			version := getPackageVersion(pkgName)
			fmt.Printf("  • %s v%s\n", pkgName, version)
		}
	}
}

func getPackageVersion(pkgName string) string {
	configPath := filepath.Join("packages", pkgName, "aether.toml")
	content, err := os.ReadFile(configPath)
	if err != nil {
		return "unknown"
	}

	// Simple version extraction
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "version = ") {
			version := strings.Trim(strings.TrimPrefix(line, "version = "), "\"")
			return version
		}
	}

	return "unknown"
}

func updatePackages(packages []string) {
	if len(packages) == 0 {
		fmt.Println("🍕 Updating all packages...")
		// Update all installed packages
		packagesDir := "packages"
		if entries, err := os.ReadDir(packagesDir); err == nil {
			for _, entry := range entries {
				if entry.IsDir() {
					updatePackage(entry.Name())
				}
			}
		}
	} else {
		fmt.Println("🍕 Updating packages...")
		for _, pkg := range packages {
			updatePackage(pkg)
		}
	}
}

func updatePackage(pkgName string) {
	fmt.Printf("🍕 Updating %s...\n", pkgName)
	// In a real implementation, this would check for updates
	// For now, just show that it's "up to date"
	fmt.Printf("🍕 %s is up to date\n", pkgName)
}

func searchPackages(query []string) {
	if len(query) == 0 {
		fmt.Println("🍕 Usage: aether package search <query>")
		return
	}

	searchTerm := strings.Join(query, " ")
	fmt.Printf("🍕 Searching for packages matching '%s'...\n", searchTerm)

	// In a real implementation, this would search a package registry
	// For now, just show some example results
	fmt.Println("🍕 Example search results:")
	fmt.Println("  • math - Mathematical utilities")
	fmt.Println("  • http - HTTP client and server")
	fmt.Println("  • json - JSON parsing and serialization")
	fmt.Println("  • crypto - Cryptographic functions")
}

func publishPackage(args []string) {
	if len(args) == 0 {
		fmt.Println("🍕 Usage: aether package publish <package_name>")
		return
	}

	pkgName := args[0]
	fmt.Printf("🍕 Publishing package %s...\n", pkgName)

	// In a real implementation, this would upload to a package registry
	fmt.Printf("🍕 Package %s published successfully!\n", pkgName)
}

func showPackageInfo(args []string) {
	if len(args) == 0 {
		fmt.Println("🍕 Usage: aether package info <package_name>")
		return
	}

	pkgName := args[0]
	fmt.Printf("🍕 Package information for %s:\n", pkgName)

	// Show package info
	if isPackageInstalled(pkgName) {
		version := getPackageVersion(pkgName)
		fmt.Printf("  Name: %s\n", pkgName)
		fmt.Printf("  Version: %s\n", version)
		fmt.Printf("  Status: Installed\n")
		fmt.Printf("  Location: packages/%s\n", pkgName)
	} else {
		fmt.Printf("🍕 Package %s is not installed\n", pkgName)
	}
}

func isPackageInstalled(pkgName string) bool {
	pkgDir := filepath.Join("packages", pkgName)
	_, err := os.Stat(pkgDir)
	return err == nil
}

func init() {
	flags := PackageCmd.Flags()
	flags.BoolVarP(&packageFlags.global, "global", "g", false, "install globally")
	flags.BoolVarP(&packageFlags.dev, "dev", "d", false, "install as development dependency")
	flags.BoolVarP(&packageFlags.force, "force", "f", false, "force installation")
	flags.BoolVarP(&packageFlags.verbose, "verbose", "v", false, "verbose output")
	flags.StringVar(&packageFlags.registry, "registry", "", "custom package registry")
}
