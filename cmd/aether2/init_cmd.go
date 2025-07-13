package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init [project_name]",
	Short: "Initialize a new aether project",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		doInit(args)
	},
}

func doInit(args []string) {
	var projectName string

	if len(args) > 0 {
		projectName = args[0]
	} else {
		projectName = getCurrentDirName()
	}

	if _, err := os.Stat("aether.toml"); err == nil {
		fmt.Print("🍕 Project already initialized. Reinitialize? (y/n): ")
		r := bufio.NewReader(os.Stdin)
		resp, _ := r.ReadString('\n')
		resp = strings.TrimSpace(resp)
		if resp != "y" && resp != "Y" {
			fmt.Println("Aborted.")
			return
		}
	}

	fmt.Printf("🍕 Initializing Aether project '%s'...\n", projectName)

	// Interactive questions
	sourceDir := askSourceDirectory()
	outputDir := askOutputDirectory()
	author := askAuthor()
	description := askDescription()
	createMainFile := askCreateMainFile()

	fmt.Println("🍕 Creating project structure...")

	// Create directories
	must(os.MkdirAll(sourceDir, 0755))
	must(os.MkdirAll(outputDir, 0755))
	must(os.MkdirAll("build", 0755))
	must(os.MkdirAll("packages", 0755))

	// Create aether.toml with custom configuration
	toml := generateProjectConfig(projectName, sourceDir, outputDir, author, description)
	must(os.WriteFile("aether.toml", []byte(toml), 0644))

	// Create .gitignore
	gitignore := []byte(fmt.Sprintf("/%s\n/build\n*.o\n*.a\n*.so\n*.dll\n*.dylib\n", outputDir))
	must(os.WriteFile(".gitignore", gitignore, 0644))

	// Create main file if requested
	if createMainFile {
		mainFile := filepath.Join(sourceDir, "main.aeth")
		mainSrc := []byte(fmt.Sprintf(`import "stdio"

stdio.printf("🍕 Hello from %s!")
stdio.printf("Your Aether project is ready to build!")
`, projectName))
		must(os.WriteFile(mainFile, mainSrc, 0644))
	}

	fmt.Printf("🍕 Project '%s' initialized successfully!\n", projectName)
	fmt.Printf("🍕 Source directory: %s\n", sourceDir)
	fmt.Printf("🍕 Output directory: %s\n", outputDir)
	fmt.Println("🍕 Run 'aether build' to compile your project!")
}

func getCurrentDirName() string {
	dir, err := os.Getwd()
	if err != nil {
		return "aether-project"
	}
	return filepath.Base(dir)
}

func askSourceDirectory() string {
	fmt.Print("🍕 Where do you want your code to be? (src): ")
	r := bufio.NewReader(os.Stdin)
	resp, _ := r.ReadString('\n')
	resp = strings.TrimSpace(resp)
	if resp == "" {
		return "src"
	}
	return resp
}

func askOutputDirectory() string {
	fmt.Print("🍕 Where do you want build outputs? (bin): ")
	r := bufio.NewReader(os.Stdin)
	resp, _ := r.ReadString('\n')
	resp = strings.TrimSpace(resp)
	if resp == "" {
		return "bin"
	}
	return resp
}

func askAuthor() string {
	fmt.Print("🍕 Who are you? (optional): ")
	r := bufio.NewReader(os.Stdin)
	resp, _ := r.ReadString('\n')
	return strings.TrimSpace(resp)
}

func askDescription() string {
	fmt.Print("🍕 What does your project do? (optional): ")
	r := bufio.NewReader(os.Stdin)
	resp, _ := r.ReadString('\n')
	return strings.TrimSpace(resp)
}

func askCreateMainFile() bool {
	fmt.Print("🍕 Create a main.aeth file? (y/n): ")
	r := bufio.NewReader(os.Stdin)
	resp, _ := r.ReadString('\n')
	resp = strings.TrimSpace(resp)
	return resp == "y" || resp == "Y"
}

func generateProjectConfig(name, sourceDir, outputDir, author, description string) string {
	config := fmt.Sprintf(`[project]
name = "%s"
version = "0.1.0"`, name)

	if author != "" {
		config += fmt.Sprintf("\nauthor = \"%s\"", author)
	}

	if description != "" {
		config += fmt.Sprintf("\ndescription = \"%s\"", description)
	}

	config += fmt.Sprintf(`

[build]
source_directories = ["%s"]
output_directory = "%s"
target = "native"
optimization = "2"
linker = "mold"
create_library = false
library_type = "shared"

[dependencies]
# Add your dependencies here
# stdlib = "packages/stdlib"
# utils = "packages/utils"

[dev-dependencies]
# Add development dependencies here
# linter = "packages/linter"
`, sourceDir, outputDir)

	return config
}
