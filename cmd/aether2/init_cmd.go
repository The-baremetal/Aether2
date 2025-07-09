package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init [project_name]",
	Short: "Initialize a new aether project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		doInit(args)
	},
}

func doInit(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: aether2 init <project_name>")
		os.Exit(1)
	}
	name := args[0]
	if _, err := os.Stat("aether.toml"); err == nil {
		fmt.Print("Project already initialized. Reinitialize? (y/n): ")
		r := bufio.NewReader(os.Stdin)
		resp, _ := r.ReadString('\n')
		resp = strings.TrimSpace(resp)
		if resp != "y" && resp != "Y" {
			fmt.Println("Aborted.")
			return
		}
	}
	fmt.Println("Initializing aether project '", name, "' in current directory...")
	must(os.MkdirAll("src", 0755))
	must(os.MkdirAll("bin", 0755))
	must(os.MkdirAll("build", 0755))
	toml := []byte(fmt.Sprintf("[project]\nname = \"%s\"\nversion = \"0.1.0\"\n\n[dependencies]\n# Only local dependencies are supported for now\n", name))
	must(os.WriteFile("aether.toml", toml, 0644))
	gitignore := []byte("/bin\n/build\n*.o\n")
	must(os.WriteFile(".gitignore", gitignore, 0644))
	mainFile := "src/main.ae"
	mainSrc := []byte("print(\"Hello, Aether!\")\n")
	must(os.WriteFile(mainFile, mainSrc, 0644))
	fmt.Println("Project '", name, "' initialized!")
}
