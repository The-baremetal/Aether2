package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

func doVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("🍕 Aether Compiler v%s\n", Version)
	fmt.Printf("   Build Date: %s\n", BuildDate)
	fmt.Printf("   Commit: %s\n", Commit)
	fmt.Printf("   Go Version: %s\n", runtime.Version())
	fmt.Printf("   OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)

	// Check LLVM
	if llvmVersion, err := getLLVMVersion(); err == nil {
		fmt.Printf("   LLVM: %s\n", llvmVersion)
	} else {
		fmt.Printf("   LLVM: Not found\n")
	}

	// Check Mold
	if moldVersion, err := getMoldVersion(); err == nil {
		fmt.Printf("   Mold: %s\n", moldVersion)
	} else {
		fmt.Printf("   Mold: Not found\n")
	}

	// Check Go
	fmt.Printf("   Go: %s\n", runtime.Version())

	// Check available targets
	fmt.Printf("   Targets: %s\n", getAvailableTargets())

	// Check stdlib
	if stdlibEnabled {
		fmt.Printf("   Stdlib: Enabled\n")
	} else {
		fmt.Printf("   Stdlib: Disabled\n")
	}
}

func getLLVMVersion() (string, error) {
	cmd := exec.Command("llc", "--version")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) > 0 {
		return strings.TrimSpace(lines[0]), nil
	}
	return "Unknown", nil
}

func getMoldVersion() (string, error) {
	cmd := exec.Command("mold", "--version")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) > 0 {
		return strings.TrimSpace(lines[0]), nil
	}
	return "Unknown", nil
}

func getAvailableTargets() string {
	targets := []string{}

	// Check common architectures
	archs := []string{"amd64", "arm64", "386", "arm"}
	for _, arch := range archs {
		if isTargetAvailable(arch) {
			targets = append(targets, arch)
		}
	}

	if len(targets) == 0 {
		return "None"
	}

	return strings.Join(targets, ", ")
}

func isTargetAvailable(arch string) bool {
	// This would check if LLVM supports the target
	// For now, return true for common targets
	commonTargets := map[string]bool{
		"amd64": true,
		"arm64": true,
		"386":   true,
		"arm":   true,
	}

	return commonTargets[arch]
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Aether version information",
	Long: `Display detailed version information about Aether.

This command shows:
  • Aether compiler version
  • Build information
  • System information
  • LLVM and Mold versions
  • Available targets
  • Configuration status`,
	Run: doVersion,
}
