package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

const (
	Version   = "0.4.0-nightly"
	BuildDate = "2025-20-7"
	Format = "yy-mm-dd"
	Commit    = "development"
)

var (
	// Global flags
	verbose    bool
	quiet      bool
	noColor    bool
	configFile string

	// CLI state
	stdlibEnabled = true
)

func must(err error) {
	if err != nil {
		if !quiet {
			fmt.Fprintf(os.Stderr, "🍕 Error: %v\n", err)
		}
		os.Exit(1)
	}
}

func printVersion() {
	fmt.Printf("🍕 Aether Compiler v%s\n", Version)
	fmt.Printf("   Build Date: %s\n", BuildDate)
	fmt.Printf("   Commit: %s\n", Commit)
	fmt.Printf("   Go Version: %s\n", runtime.Version())
	fmt.Printf("   OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("   LLVM: Linked\n")
	fmt.Printf("   Mold: Linked\n")
}

func printCompletionScript(cmd *cobra.Command, args []string) error {
	shell := args[0]
	switch shell {
	case "bash":
		return cmd.Root().GenBashCompletion(os.Stdout)
	case "zsh":
		return cmd.Root().GenZshCompletion(os.Stdout)
	case "fish":
		return cmd.Root().GenFishCompletion(os.Stdout, true)
	case "powershell":
		return cmd.Root().GenPowerShellCompletion(os.Stdout)
	default:
		return fmt.Errorf("unsupported shell: %s", shell)
	}
}

func setupCompletion(cmd *cobra.Command) {
	completionCmd := &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate completion script",
		Long: `Generate shell completion script for Aether.

To load completions:

Bash:
  $ source <(aether completion bash)

Zsh:
  $ source <(aether completion zsh)

Fish:
  $ aether completion fish | source

PowerShell:
  PS> aether completion powershell | Out-String | Invoke-Expression
`,
		ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
		Args:      cobra.ExactValidArgs(1),
		RunE:      printCompletionScript,
	}
	cmd.AddCommand(completionCmd)
}

func setupRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "aether",
		Version: Version,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			return fmt.Errorf("unknown command: %s", args[0])
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Set up global flags
			if quiet {
				// Suppress all output
			}
			if noColor {
				// Disable colors
			}
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	// Global flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "suppress all output")
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable colored output")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default: aether.toml)")

	// Add commands
	rootCmd.AddCommand(BuildCmd)
	rootCmd.AddCommand(InitCmd)
	rootCmd.AddCommand(CleanCmd)
	rootCmd.AddCommand(LintCmd)
	rootCmd.AddCommand(UpdateCmd)
	rootCmd.AddCommand(VersionCmd)
	rootCmd.AddCommand(FormatCmd)
	rootCmd.AddCommand(TestCmd)
	rootCmd.AddCommand(DocsCmd)
	rootCmd.AddCommand(PackageCmd)
	rootCmd.AddCommand(CodemodCmd)

	// Setup completion
	setupCompletion(rootCmd)

	// Custom help template
	rootCmd.SetHelpTemplate(`🍕 Aether Language CLI

{{.Short}}

{{if .Long}}{{.Long}}{{end}}

Usage:
  {{.UseLine}}{{if .HasAvailableSubCommands}} [command]{{end}}{{if .HasAvailableFlags}} [flags]{{end}}

{{if .HasAvailableSubCommands}}Available Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}

{{if .HasAvailableLocalFlags}}Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}

{{if .HasAvailableInheritedFlags}}Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}

{{if .HasExample}}Examples:
{{.Example}}{{end}}

{{if .HasHelpSubCommands}}Additional help topics:{{range .Commands}}{{if .IsHelpCommand}}
  {{rpad .CommandPath .CommandPathPadding }} {{.Short}}{{end}}{{end}}{{end}}

{{if .HasAvailableSubCommands}}Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`)

	return rootCmd
}

func printSmartError(err error, context string) {
	if err == nil {
		return
	}
	msg := err.Error()
	if strings.Contains(msg, "unknown command") {
		fmt.Fprintf(os.Stderr, "🦄🍕 Unknown command: %s\nDid you mean one of these? Try 'aether --help' for a list of commands!\n", context)
	} else if strings.Contains(msg, "not found") || strings.Contains(msg, "No such file") {
		fmt.Fprintf(os.Stderr, "🍕 Oops! %s\nDid you forget to create or specify the file?\n", msg)
	} else if strings.Contains(msg, "permission denied") {
		fmt.Fprintf(os.Stderr, "🍕 Permission denied! Try running as admin or check your file permissions.\n")
	} else {
		fmt.Fprintf(os.Stderr, "🍕 Error: %s\n", msg)
	}
	os.Exit(1)
}

func main() {
	rootCmd := setupRootCmd()

	// Handle errors gracefully
	if err := rootCmd.Execute(); err != nil {
		if !quiet {
			printSmartError(err, "")
		}
		os.Exit(1)
	}
}
