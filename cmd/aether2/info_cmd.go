package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show project info from aether.toml",
	Run: func(cmd *cobra.Command, args []string) {
		doInfo()
	},
}

func doInfo() {
	data, err := os.ReadFile("aether.toml")
	if err != nil {
		fmt.Println("No aether.toml found. Not an aether project?")
		os.Exit(1)
	}
	fmt.Println("Project info:")
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "[") || strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		fmt.Println(line)
	}
}
