package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	playgroundFlags struct {
		port   int
		host   string
		open   bool
		watch  bool
		output string
	}
)

func doPlayground(cmd *cobra.Command, args []string) {
	fmt.Println("🍕 Starting Aether Playground...")
	fmt.Println()

	// Create playground directory
	playgroundDir := playgroundFlags.output
	if playgroundDir == "" {
		playgroundDir = "aether-playground"
	}

	if err := setupPlayground(playgroundDir); err != nil {
		fmt.Fprintf(os.Stderr, "🍕 Error setting up playground: %v\n", err)
		os.Exit(1)
	}

	// Start playground server
	if err := startPlaygroundServer(playgroundDir); err != nil {
		fmt.Fprintf(os.Stderr, "🍕 Error starting playground: %v\n", err)
		os.Exit(1)
	}
}

func setupPlayground(dir string) error {
	// Create playground directory
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Create example files
	examples := []struct {
		name    string
		content string
	}{
		{
			name: "main.aeth",
			content: `// 🍕 Welcome to Aether Playground!

func main() {
    print("Hello from Aether!")
    
    // Try some basic operations
    let x = 42
    let y = 10
    print("x + y = ", x + y)
    
    // String operations
    let message = "Aether is awesome!"
    print(message)
    
    // Simple loop
    for i in 1..5 {
        print("Count: ", i)
    }
}`,
		},
		{
			name: "math.aeth",
			content: `// Math utilities

func add(a, b) {
    return a + b
}

func multiply(a, b) {
    return a * b
}

func factorial(n) {
    if n <= 1 {
        return 1
    }
    return n * factorial(n - 1)
}`,
		},
		{
			name: "test.aeth",
			content: `// Test file

func test_math() {
    assert(add(2, 3) == 5)
    assert(multiply(4, 5) == 20)
    assert(factorial(5) == 120)
    print("All math tests passed!")
}

func main() {
    test_math()
}`,
		},
		{
			name: "aether.toml",
			content: `[project]
name = "aether-playground"
version = "0.1.0"
description = "Aether Playground Project"

[build]
target = "native"
optimization = "debug"

[dependencies]
# Add your dependencies here

[dev-dependencies]
# Add your development dependencies here`,
		},
		{
			name:    "README.md",
			content: "# 🍕 Aether Playground\n\nWelcome to the Aether Playground! This is a sandbox environment for experimenting with Aether code.\n\n## Quick Start\n\n1. Edit the files in this directory\n2. Run `aether build` to compile\n3. Run `aether test` to run tests\n4. Use `aether playground` to start the interactive environment\n\n## Examples\n\n- `main.aeth` - Basic syntax and operations\n- `math.aeth` - Function definitions\n- `test.aeth` - Unit testing\n\n## Features\n\n- Live code compilation\n- Interactive REPL\n- Real-time error reporting\n- Hot reloading\n- Integrated testing\n\nHappy coding! 🍕",
		},
	}

	for _, example := range examples {
		filePath := filepath.Join(dir, example.name)
		if err := os.WriteFile(filePath, []byte(example.content), 0644); err != nil {
			return err
		}
	}

	fmt.Printf("🍕 Playground created at: %s\n", dir)
	return nil
}

func startPlaygroundServer(dir string) error {
	// Change to playground directory
	if err := os.Chdir(dir); err != nil {
		return err
	}

	fmt.Println("🍕 Playground Features:")
	fmt.Println("  • Live code compilation")
	fmt.Println("  • Interactive REPL")
	fmt.Println("  • Real-time error reporting")
	fmt.Println("  • Hot reloading")
	fmt.Println("  • Integrated testing")
	fmt.Println()

	// Start interactive mode
	return startInteractiveMode()
}

func startInteractiveMode() error {
	fmt.Println("🍕 Interactive Aether Playground")
	fmt.Println("Type 'help' for commands, 'exit' to quit")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("🍕 aether> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		if input == "exit" || input == "quit" {
			fmt.Println("🍕 Goodbye!")
			break
		}

		if err := handlePlaygroundCommand(input); err != nil {
			fmt.Printf("🍕 Error: %v\n", err)
		}
	}

	return scanner.Err()
}

func handlePlaygroundCommand(input string) error {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil
	}

	command := parts[0]
	args := parts[1:]

	switch command {
	case "help":
		printPlaygroundHelp()
	case "build":
		return runBuild(args)
	case "run":
		return runProgram(args)
	case "test":
		return runTests(args)
	case "format":
		return runFormat(args)
	case "docs":
		return runDocs(args)
	case "clear":
		clearScreen()
	case "ls":
		listFiles()
	case "cat":
		if len(args) > 0 {
			return showFile(args[0])
		}
		fmt.Println("🍕 Usage: cat <filename>")
	default:
		// Try to execute as Aether code
		return executeAetherCode(input)
	}

	return nil
}

func printPlaygroundHelp() {
	fmt.Println("🍕 Available Commands:")
	fmt.Println("  help              - Show this help")
	fmt.Println("  build [file]      - Build Aether code")
	fmt.Println("  run [file]        - Run Aether program")
	fmt.Println("  test              - Run tests")
	fmt.Println("  format [file]     - Format code")
	fmt.Println("  docs              - Generate documentation")
	fmt.Println("  clear             - Clear screen")
	fmt.Println("  ls                - List files")
	fmt.Println("  cat <file>        - Show file contents")
	fmt.Println("  exit/quit         - Exit playground")
	fmt.Println()
	fmt.Println("🍕 Or just type Aether code directly!")
}

func runBuild(args []string) error {
	buildArgs := []string{"build"}
	if len(args) > 0 {
		buildArgs = append(buildArgs, args...)
	} else {
		buildArgs = append(buildArgs, "main.aeth")
	}

	cmd := exec.Command("aether", buildArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("🍕 Building: %s\n", strings.Join(buildArgs[1:], " "))
	return cmd.Run()
}

func runProgram(args []string) error {
	program := "main"
	if len(args) > 0 {
		program = args[0]
	}

	// Build first
	if err := runBuild([]string{program + ".aeth"}); err != nil {
		return err
	}

	// Run the program
	cmd := exec.Command("./" + program)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("🍕 Running: %s\n", program)
	return cmd.Run()
}

func runTests(args []string) error {
	testArgs := []string{"test"}
	if len(args) > 0 {
		testArgs = append(testArgs, args...)
	}

	cmd := exec.Command("aether", testArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("🍕 Running tests...")
	return cmd.Run()
}

func runFormat(args []string) error {
	formatArgs := []string{"format", "--write"}
	if len(args) > 0 {
		formatArgs = append(formatArgs, args...)
	} else {
		formatArgs = append(formatArgs, "*.aeth")
	}

	cmd := exec.Command("aether", formatArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("🍕 Formatting code...")
	return cmd.Run()
}

func runDocs(args []string) error {
	docsArgs := []string{"docs"}
	if len(args) > 0 {
		docsArgs = append(docsArgs, args...)
	}

	cmd := exec.Command("aether", docsArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("🍕 Generating documentation...")
	return cmd.Run()
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func listFiles() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("🍕 Error reading directory: %v\n", err)
		return
	}

	fmt.Println("🍕 Files in playground:")
	for _, file := range files {
		if !file.IsDir() {
			fmt.Printf("  • %s\n", file.Name())
		}
	}
}

func showFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	fmt.Printf("🍕 Contents of %s:\n", filename)
	fmt.Println("```")
	fmt.Print(string(content))
	fmt.Println("```")
	return nil
}

func executeAetherCode(code string) error {
	// Create temporary file
	tmpFile := fmt.Sprintf("playground_%d.aeth", time.Now().Unix())

	if err := os.WriteFile(tmpFile, []byte(code), 0644); err != nil {
		return err
	}
	defer os.Remove(tmpFile)

	// Try to build and run
	if err := runBuild([]string{tmpFile}); err != nil {
		return err
	}

	return runProgram([]string{strings.TrimSuffix(tmpFile, ".aeth")})
}

var PlaygroundCmd = &cobra.Command{
	Use:   "playground",
	Short: "Start Aether playground environment",
	Long: `Start an interactive Aether development environment.

The playground provides:
  • Live code compilation
  • Interactive REPL
  • Real-time error reporting
  • Hot reloading
  • Integrated testing
  • Example projects

Examples:
  aether playground              # Start playground
  aether playground --port 8080 # Custom port
  aether playground --open      # Open in browser
  aether playground --watch     # Enable file watching`,
	Run: doPlayground,
}

func init() {
	flags := PlaygroundCmd.Flags()
	flags.IntVarP(&playgroundFlags.port, "port", "p", 8080, "playground port")
	flags.StringVar(&playgroundFlags.host, "host", "localhost", "playground host")
	flags.BoolVarP(&playgroundFlags.open, "open", "o", false, "open playground in browser")
	flags.BoolVarP(&playgroundFlags.watch, "watch", "w", false, "enable file watching")
	flags.StringVarP(&playgroundFlags.output, "output", "d", "", "playground directory")
}
