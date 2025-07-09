package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var CrossCmd = &cobra.Command{
	Use:   "cross <target1> [target2] [target3]...",
	Short: "Cross-compile for multiple targets",
	Run: func(cmd *cobra.Command, args []string) {
		doCross(args)
	},
}

func doCross(targets []string) {
	if len(targets) < 1 {
		fmt.Println("Usage: aether2 cross <target1> [target2] [target3]...")
		fmt.Println("Targets: linux-amd64, linux-arm64, darwin-amd64, darwin-arm64, windows-amd64")
		os.Exit(1)
	}

	fmt.Printf("Cross-compiling for targets: %v\n", targets)

	for _, target := range targets {
		parts := strings.Split(target, "-")
		if len(parts) != 2 {
			fmt.Printf("Invalid target format: %s (use format: os-arch)\n", target)
			continue
		}

		targetOS, arch := parts[0], parts[1]
		outputName := fmt.Sprintf("bin/aether-%s-%s", targetOS, arch)
		if targetOS == "windows" {
			outputName += ".exe"
		}

		fmt.Printf("\nBuilding for %s/%s...\n", targetOS, arch)

		args := []string{"build", "--target-os", targetOS, "--target-arch", arch, "--output", outputName}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Printf("Failed to build for %s/%s: %v\n", targetOS, arch, err)
		} else {
			fmt.Printf("Successfully built for %s/%s: %s\n", targetOS, arch, outputName)
		}
	}

	fmt.Println("\nCross-compilation complete!")
}
