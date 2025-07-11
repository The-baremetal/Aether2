package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

type Diagnostic struct {
	Name     string
	Status   string
	Message  string
	Fix      string
	Critical bool
}

func doDoctor(cmd *cobra.Command, args []string) {
	fmt.Println("🍕 Aether Doctor - Diagnosing your installation...")
	fmt.Println()

	diagnostics := []Diagnostic{
		checkAetherCompiler(),
		checkLLVM(),
		checkMold(),
		checkGo(),
		checkTargets(),
		checkPermissions(),
		checkConfig(),
	}

	// Run diagnostics
	allGood := true
	criticalIssues := 0

	for _, d := range diagnostics {
		fmt.Printf("  %s: %s\n", d.Name, d.Status)
		if d.Message != "" {
			fmt.Printf("    %s\n", d.Message)
		}
		if d.Fix != "" {
			fmt.Printf("    Fix: %s\n", d.Fix)
		}
		fmt.Println()

		if d.Status == "FAIL" {
			allGood = false
			if d.Critical {
				criticalIssues++
			}
		}
	}

	// Summary
	fmt.Println("🍕 Diagnosis Summary:")
	if allGood {
		fmt.Println("  ✅ All checks passed! Your Aether installation is healthy.")
	} else {
		fmt.Printf("  ❌ Found %d issue(s), %d critical\n", len(diagnostics)-countPassed(diagnostics), criticalIssues)
		if criticalIssues > 0 {
			fmt.Println("  Critical issues must be fixed before using Aether.")
		}
	}
}

func checkAetherCompiler() Diagnostic {
	// Check if aether compiler is available
	cmd := exec.Command("aether", "--version")
	if err := cmd.Run(); err != nil {
		return Diagnostic{
			Name:     "Aether Compiler",
			Status:   "FAIL",
			Message:  "Aether compiler not found in PATH",
			Fix:      "Add aether to your PATH or install Aether",
			Critical: true,
		}
	}

	return Diagnostic{
		Name:   "Aether Compiler",
		Status: "PASS",
	}
}

func checkLLVM() Diagnostic {
	// Check LLVM installation
	cmd := exec.Command("llc", "--version")
	if err := cmd.Run(); err != nil {
		return Diagnostic{
			Name:     "LLVM",
			Status:   "FAIL",
			Message:  "LLVM not found or not working",
			Fix:      "Install LLVM: https://llvm.org/docs/GettingStarted.html",
			Critical: true,
		}
	}

	return Diagnostic{
		Name:   "LLVM",
		Status: "PASS",
	}
}

func checkMold() Diagnostic {
	// Check Mold linker
	cmd := exec.Command("mold", "--version")
	if err := cmd.Run(); err != nil {
		return Diagnostic{
			Name:    "Mold Linker",
			Status:  "FAIL",
			Message: "Mold linker not found",
			Fix:     "Install Mold: https://github.com/rui314/mold",
		}
	}

	return Diagnostic{
		Name:   "Mold Linker",
		Status: "PASS",
	}
}

func checkGo() Diagnostic {
	// Check Go installation
	cmd := exec.Command("go", "version")
	if err := cmd.Run(); err != nil {
		return Diagnostic{
			Name:     "Go Runtime",
			Status:   "FAIL",
			Message:  "Go not found",
			Fix:      "Install Go: https://golang.org/doc/install",
			Critical: true,
		}
	}

	return Diagnostic{
		Name:   "Go Runtime",
		Status: "PASS",
	}
}

func checkTargets() Diagnostic {
	// Check if we can compile for current target
	testCode := `func main() { print("test") }`

	// Write test file
	testFile := "doctor_test.aeth"
	if err := os.WriteFile(testFile, []byte(testCode), 0644); err != nil {
		return Diagnostic{
			Name:    "Target Compilation",
			Status:  "FAIL",
			Message: "Cannot write test file",
			Fix:     "Check directory permissions",
		}
	}
	defer os.Remove(testFile)

	// Try to compile
	cmd := exec.Command("aether", "build", testFile)
	if err := cmd.Run(); err != nil {
		return Diagnostic{
			Name:    "Target Compilation",
			Status:  "FAIL",
			Message: "Cannot compile for current target",
			Fix:     "Check LLVM installation and target support",
		}
	}

	return Diagnostic{
		Name:   "Target Compilation",
		Status: "PASS",
	}
}

func checkPermissions() Diagnostic {
	// Check if we can write to current directory
	testFile := "doctor_perms_test"
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		return Diagnostic{
			Name:    "Directory Permissions",
			Status:  "FAIL",
			Message: "Cannot write to current directory",
			Fix:     "Check directory permissions or use different directory",
		}
	}
	os.Remove(testFile)

	return Diagnostic{
		Name:   "Directory Permissions",
		Status: "PASS",
	}
}

func checkConfig() Diagnostic {
	// Check for aether.toml
	if _, err := os.Stat("aether.toml"); os.IsNotExist(err) {
		return Diagnostic{
			Name:    "Project Configuration",
			Status:  "WARN",
			Message: "No aether.toml found in current directory",
			Fix:     "Run 'aether init' to create a new project",
		}
	}

	return Diagnostic{
		Name:   "Project Configuration",
		Status: "PASS",
	}
}

func countPassed(diagnostics []Diagnostic) int {
	count := 0
	for _, d := range diagnostics {
		if d.Status == "PASS" {
			count++
		}
	}
	return count
}

var DoctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Diagnose Aether installation issues",
	Long: `Check your Aether installation for common issues.

This command will:
  • Verify Aether compiler is working
  • Check LLVM installation
  • Verify Mold linker
  • Test Go runtime
  • Check target compilation
  • Verify permissions
  • Check project configuration

Examples:
  aether doctor              # Run all diagnostics
  aether doctor --verbose    # Show detailed output`,
	Run: doDoctor,
}
