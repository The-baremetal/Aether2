package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var DepsCmd = &cobra.Command{
	Use:   "deps",
	Short: "Resolve and update dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		doDeps()
	},
}

func doDeps() {
	fmt.Println("Resolving dependencies...")
	data, err := os.ReadFile("aether.toml")
	must(err)
	deps := make(map[string]string)
	inDeps := false
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "[dependencies]" {
			inDeps = true
			continue
		}
		if strings.HasPrefix(line, "[") && line != "[dependencies]" {
			inDeps = false
		}
		if inDeps && line != "" && !strings.HasPrefix(line, "#") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				name := strings.TrimSpace(parts[0])
				path := strings.Trim(strings.TrimSpace(parts[1]), "\"")
				absPath, err := filepath.Abs(path)
				must(err)
				if _, err := os.Stat(absPath); os.IsNotExist(err) {
					fmt.Printf("Dependency '%s' not found at '%s'\n", name, absPath)
					continue
				}
				deps[name] = absPath
			}
		}
	}
	lock := "# aether.lock - DO NOT EDIT\n[dependencies]\n"
	for name, path := range deps {
		lock += fmt.Sprintf("%s = \"%s\"\n", name, path)
	}
	must(os.WriteFile("aether.lock", []byte(lock), 0644))
	fmt.Println("Dependencies resolved and aether.lock updated!")
}
