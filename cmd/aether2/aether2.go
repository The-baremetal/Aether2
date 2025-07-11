package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

const (
	Version   = "0.2.0-nightly"
	BuildDate = "2024-01-01"
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
		Short: "🍕 Aether Language - Fast, Simple, Delicious",
		Long: `Aether is a modern programming language that's fast, simple, and delicious.

Features:
  • LLVM-powered compilation for maximum performance
  • Mold linker for lightning-fast builds
  • Simple, consistent syntax
  • Built-in library support
  • Cross-platform compilation
  • Professional CMake integration

Examples:
  aether build                    # Build current project
  aether build --create-library   # Create shared library
  aether library --analyze libc   # Analyze system library
  aether cross --target linux     # Cross-compile for Linux

Documentation: https://aether-lang.org
Repository:   https://github.com/aether-lang/aether`,
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
	rootCmd.AddCommand(CrossCmd)
	rootCmd.AddCommand(InitCmd)
	rootCmd.AddCommand(CleanCmd)
	rootCmd.AddCommand(InfoCmd)
	rootCmd.AddCommand(DepsCmd)
	rootCmd.AddCommand(ScaffoldCmd)
	rootCmd.AddCommand(LintCmd)
	rootCmd.AddCommand(UpdateCmd)
	rootCmd.AddCommand(LibraryCmd)
	rootCmd.AddCommand(VersionCmd)
	rootCmd.AddCommand(DoctorCmd)
	rootCmd.AddCommand(FormatCmd)
	rootCmd.AddCommand(TestCmd)
	rootCmd.AddCommand(DocsCmd)
	rootCmd.AddCommand(PlaygroundCmd)
	rootCmd.AddCommand(PackageCmd)
	rootCmd.AddCommand(NewCmd)

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

func main() {
	rootCmd := setupRootCmd()

	// Handle errors gracefully
	if err := rootCmd.Execute(); err != nil {
		if !quiet {
			// Print error with context
			if strings.Contains(err.Error(), "unknown command") {
				fmt.Fprintf(os.Stderr, "🍕 Unknown command. Use 'aether --help' for available commands.\n")
			} else {
				fmt.Fprintf(os.Stderr, "🍕 Error: %v\n", err)
			}
		}
		os.Exit(1)
	}
}
