package main

import (
	"aether/src/analysis"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	libraryFlags struct {
		create      bool
		analyze     bool
		bind        bool
		output      string
		libType     string
		libName     string
		libVersion  string
		description string
		url         string
		requires    string
		conflicts   string
		provides    string
		exportAll   bool
		generatePC  bool
		verbose     bool
		quiet       bool
	}
)

func doLibrary(args []string) {
	if len(args) == 0 {
		fmt.Println("Error: No library specified")
		fmt.Println("Usage: aether library [command] [library-name]")
		os.Exit(1)
	}

	libName := args[0]

	if libraryFlags.create {
		createLibraryCommand(libName)
	} else if libraryFlags.analyze {
		analyzeLibrary(libName)
	} else if libraryFlags.bind {
		bindLibrary(libName)
	} else {
		fmt.Println("Error: No operation specified")
		fmt.Println("Use --create, --analyze, or --bind")
		os.Exit(1)
	}
}

func createLibraryCommand(libName string) {
	if !libraryFlags.quiet {
		fmt.Printf("Creating library: %s\n", libName)
	}

	// Find source files
	sourceDir := "src"
	if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
		fmt.Printf("Error: Source directory '%s' not found\n", sourceDir)
		os.Exit(1)
	}

	// Build the library using the build command
	buildArgs := []string{
		"--create-library",
		"--library-type", libraryFlags.libType,
		"--library-name", libName,
		"--output", filepath.Join("lib", libName),
	}

	if libraryFlags.libVersion != "" {
		buildArgs = append(buildArgs, "--library-version", libraryFlags.libVersion)
	}

	if libraryFlags.description != "" {
		buildArgs = append(buildArgs, "--library-description", libraryFlags.description)
	}

	if libraryFlags.url != "" {
		buildArgs = append(buildArgs, "--library-url", libraryFlags.url)
	}

	if libraryFlags.requires != "" {
		buildArgs = append(buildArgs, "--library-requires", libraryFlags.requires)
	}

	if libraryFlags.conflicts != "" {
		buildArgs = append(buildArgs, "--library-conflicts", libraryFlags.conflicts)
	}

	if libraryFlags.provides != "" {
		buildArgs = append(buildArgs, "--library-provides", libraryFlags.provides)
	}

	if libraryFlags.exportAll {
		buildArgs = append(buildArgs, "--export-symbols")
	}

	if libraryFlags.generatePC {
		buildArgs = append(buildArgs, "--generate-pkg-config")
	}

	if libraryFlags.verbose {
		buildArgs = append(buildArgs, "--verbose")
	}

	if libraryFlags.quiet {
		buildArgs = append(buildArgs, "--quiet")
	}

	// Add source files
	buildArgs = append(buildArgs, sourceDir)

	// Execute build command
	if err := executeBuildCommand(buildArgs); err != nil {
		fmt.Printf("Error creating library: %v\n", err)
		os.Exit(1)
	}

	if !libraryFlags.quiet {
		fmt.Printf("Library '%s' created successfully!\n", libName)
	}
}

func analyzeLibrary(libPath string) {
	if !libraryFlags.quiet {
		fmt.Printf("Analyzing library: %s\n", libPath)
	}

	// Find the library
	libPath, err := analysis.FindLibrary(libPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Analyze the library
	lib, err := analysis.AnalyzeExternalLibrary(libPath)
	if err != nil {
		fmt.Printf("Error analyzing library: %v\n", err)
		os.Exit(1)
	}

	// Display analysis results
	fmt.Printf("Library Analysis: %s\n", lib.Name)
	fmt.Printf("  Path: %s\n", lib.Path)
	fmt.Printf("  Type: %s\n", lib.Type)
	fmt.Printf("  Symbols: %d\n", len(lib.Symbols))
	fmt.Printf("  Functions: %d\n", len(lib.Functions))
	fmt.Printf("  Dependencies: %d\n", len(lib.Dependencies))

	if libraryFlags.verbose {
		fmt.Println("\nSymbols:")
		for _, symbol := range lib.Symbols {
			fmt.Printf("  %s\n", symbol)
		}

		fmt.Println("\nFunctions:")
		for _, function := range lib.Functions {
			fmt.Printf("  %s\n", function.Name)
		}

		fmt.Println("\nDependencies:")
		for _, dep := range lib.Dependencies {
			fmt.Printf("  %s\n", dep)
		}
	}
}

func bindLibrary(libName string) {
	if !libraryFlags.quiet {
		fmt.Printf("Generating binding for library: %s\n", libName)
	}

	// Find the library
	libPath, err := analysis.FindLibrary(libName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Analyze the library
	lib, err := analysis.AnalyzeExternalLibrary(libPath)
	if err != nil {
		fmt.Printf("Error analyzing library: %v\n", err)
		os.Exit(1)
	}

	// Generate binding
	binding, err := analysis.GenerateLibraryBinding(lib)
	if err != nil {
		fmt.Printf("Error generating binding: %v\n", err)
		os.Exit(1)
	}

	// Validate binding
	errors := analysis.ValidateLibraryBinding(binding)
	if len(errors) > 0 {
		fmt.Println("Binding validation errors:")
		for _, err := range errors {
			fmt.Printf("  %s\n", err)
		}
		os.Exit(1)
	}

	// Generate Aether binding code
	code := analysis.GenerateAetherBinding(binding)

	// Write to file
	outputFile := libraryFlags.output
	if outputFile == "" {
		outputFile = fmt.Sprintf("%s_binding.ae", libName)
	}

	if err := os.WriteFile(outputFile, []byte(code), 0644); err != nil {
		fmt.Printf("Error writing binding file: %v\n", err)
		os.Exit(1)
	}

	if !libraryFlags.quiet {
		fmt.Printf("Binding generated: %s\n", outputFile)
	}
}

func executeBuildCommand(args []string) error {
	// This would call the build command with the given arguments
	// For now, we'll simulate it
	fmt.Printf("Executing: aether build %s\n", strings.Join(args, " "))
	return nil
}

var LibraryCmd = &cobra.Command{
	Use:   "library [library-name]",
	Short: "Manage Aether libraries",
	Long: `Create, analyze, and bind libraries in Aether.

Examples:
  aether library --create mylib                    # Create a new library
  aether library --analyze libc                    # Analyze system library
  aether library --bind openssl                    # Generate binding for OpenSSL
  aether library --create mylib --lib-type both    # Create shared and static libraries`,
	Run: func(cmd *cobra.Command, args []string) {
		doLibrary(args)
	},
}

func init() {
	flags := LibraryCmd.Flags()

	// Operation flags
	flags.BoolVar(&libraryFlags.create, "create", false, "create a new library")
	flags.BoolVar(&libraryFlags.analyze, "analyze", false, "analyze an existing library")
	flags.BoolVar(&libraryFlags.bind, "bind", false, "generate binding for external library")

	// Library creation flags
	flags.StringVar(&libraryFlags.libType, "lib-type", "shared", "type of library (shared, static, both)")
	flags.StringVar(&libraryFlags.libVersion, "version", "", "library version")
	flags.StringVar(&libraryFlags.description, "description", "", "library description")
	flags.StringVar(&libraryFlags.url, "url", "", "library URL")
	flags.StringVar(&libraryFlags.requires, "requires", "", "required libraries")
	flags.StringVar(&libraryFlags.conflicts, "conflicts", "", "conflicting libraries")
	flags.StringVar(&libraryFlags.provides, "provides", "", "provided libraries")
	flags.BoolVar(&libraryFlags.exportAll, "export-all", false, "export all symbols")
	flags.BoolVar(&libraryFlags.generatePC, "generate-pc", false, "generate pkg-config file")

	// Output flags
	flags.StringVar(&libraryFlags.output, "output", "", "output file name")

	// Verbosity flags
	flags.BoolVar(&libraryFlags.verbose, "verbose", false, "verbose output")
	flags.BoolVar(&libraryFlags.quiet, "quiet", false, "quiet output")
}
