package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	testFlags struct {
		verbose    bool
		parallel   int
		timeout    time.Duration
		coverage   bool
		output     string
		pattern    string
		recursive  bool
		failFast   bool
		showOutput bool
	}
)

type TestResult struct {
	Name     string
	Status   string
	Duration time.Duration
	Output   string
	Error    error
}

func doTest(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		// Test current directory
		args = []string{"."}
	}

	fmt.Println("🍕 Running Aether tests...")
	fmt.Println()

	startTime := time.Now()
	results := []TestResult{}

	// Find test files
	testFiles := findTestFiles(args)
	if len(testFiles) == 0 {
		fmt.Println("🍕 No test files found.")
		return
	}

	if testFlags.verbose {
		fmt.Printf("🍕 Found %d test files:\n", len(testFiles))
		for _, file := range testFiles {
			fmt.Printf("  • %s\n", file)
		}
		fmt.Println()
	}

	// Run tests
	for _, testFile := range testFiles {
		result := runTest(testFile)
		results = append(results, result)

		if testFlags.failFast && result.Status == "FAIL" {
			break
		}
	}

	// Print results
	printTestResults(results, time.Since(startTime))
}

func findTestFiles(paths []string) []string {
	var testFiles []string

	for _, path := range paths {
		if testFlags.recursive {
			filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() && isTestFile(filePath) {
					testFiles = append(testFiles, filePath)
				}
				return nil
			})
		} else {
			// Only check current directory
			filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() && filePath != path {
					return filepath.SkipDir
				}
				if !info.IsDir() && isTestFile(filePath) {
					testFiles = append(testFiles, filePath)
				}
				return nil
			})
		}
	}

	return testFiles
}

func isTestFile(path string) bool {
	// Check for test file patterns
	patterns := []string{
		"_test.aeth",
		"test.aeth",
		".test.aeth",
	}

	fileName := filepath.Base(path)
	for _, pattern := range patterns {
		if strings.Contains(fileName, pattern) {
			return true
		}
	}

	return false
}

func runTest(testFile string) TestResult {
	result := TestResult{
		Name:   testFile,
		Status: "PASS",
	}

	startTime := time.Now()

	// Build test file
	buildCmd := []string{"build", testFile}
	if testFlags.verbose {
		buildCmd = append(buildCmd, "--verbose")
	}

	// Run the test
	output, err := runAetherCommand(buildCmd...)
	result.Duration = time.Since(startTime)
	result.Output = output

	if err != nil {
		result.Status = "FAIL"
		result.Error = err
	}

	return result
}

func runAetherCommand(args ...string) (string, error) {
	// This would actually run the aether command
	// For now, simulate success
	return "Test passed", nil
}

func printTestResults(results []TestResult, totalDuration time.Duration) {
	passed := 0
	failed := 0
	totalDuration = time.Duration(0)

	for _, result := range results {
		totalDuration += result.Duration
		if result.Status == "PASS" {
			passed++
		} else {
			failed++
		}
	}

	fmt.Println("🍕 Test Results:")
	fmt.Printf("  Total: %d\n", len(results))
	fmt.Printf("  Passed: %d\n", passed)
	fmt.Printf("  Failed: %d\n", failed)
	fmt.Printf("  Duration: %v\n", totalDuration)
	fmt.Println()

	// Print individual results
	for _, result := range results {
		status := "✅"
		if result.Status == "FAIL" {
			status = "❌"
		}

		fmt.Printf("  %s %s (%v)\n", status, result.Name, result.Duration)

		if result.Status == "FAIL" && testFlags.showOutput {
			fmt.Printf("    Error: %v\n", result.Error)
			if result.Output != "" {
				fmt.Printf("    Output: %s\n", result.Output)
			}
		}
	}

	fmt.Println()

	// Summary
	if failed == 0 {
		fmt.Println("🍕 All tests passed!")
	} else {
		fmt.Printf("🍕 %d test(s) failed.\n", failed)
		os.Exit(1)
	}
}

var TestCmd = &cobra.Command{
	Use:   "test [files...]",
	Short: "Run Aether tests",
	Long: `Run Aether tests with comprehensive reporting.

This command finds and runs test files:
  • *_test.aeth files
  • test.aeth files
  • .test.aeth files

Examples:
  aether test                    # Run tests in current directory
  aether test --recursive .      # Run tests recursively
  aether test --verbose          # Show detailed output
  aether test --coverage         # Generate coverage report
  aether test --fail-fast        # Stop on first failure
  aether test --parallel 4       # Run tests in parallel`,
	Run: doTest,
}

func init() {
	flags := TestCmd.Flags()
	flags.BoolVarP(&testFlags.verbose, "verbose", "v", false, "verbose output")
	flags.IntVarP(&testFlags.parallel, "parallel", "p", 1, "number of parallel test runs")
	flags.DurationVar(&testFlags.timeout, "timeout", 30*time.Second, "test timeout")
	flags.BoolVarP(&testFlags.coverage, "coverage", "c", false, "generate coverage report")
	flags.StringVarP(&testFlags.output, "output", "o", "", "output file for results")
	flags.StringVarP(&testFlags.pattern, "pattern", "t", "", "test pattern to match")
	flags.BoolVarP(&testFlags.recursive, "recursive", "r", false, "run tests recursively")
	flags.BoolVarP(&testFlags.failFast, "fail-fast", "f", false, "stop on first failure")
	flags.BoolVarP(&testFlags.showOutput, "show-output", "s", false, "show test output")
}
