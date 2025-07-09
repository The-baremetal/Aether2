package main

import (
	"aether/src/compiler"
	"aether/src/lexer"
	"aether/src/parser"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"aether/lib/utils"

	"github.com/spf13/cobra"
)

var (
	buildFlags struct {
		noStdlib       bool
		targetOS       string
		targetArch     string
		linker         string
		outputName     string
		fuseLd         string
		optimization   string
		debugInfo      bool
		debugSymbols   bool
		verbose        bool
		quiet          bool
		noOptimize     bool
		noInline       bool
		noVectorize    bool
		noUnroll       bool
		stackProtector string
		relocModel     string
		codeModel      string
		cpu            string
		features       string
		emitIR         bool
		emitASM        bool
		emitBitcode    bool
		emitLLVM       bool
		emitObj        bool
		emitExe        bool
		emitTokens     bool
		checkImports   bool
		analyzeOnly    bool
		parallel       bool
		threads        int
		timeCompile    bool
		stats          bool
		profile        bool
		sanitize       string
		strip          bool
		pie            bool
		static         bool
		shared         bool
		rdynamic       bool
		exportDynamic  bool
		noStartFiles   bool
		noDefaultLibs  bool
		nostdlib       bool
		nodefaultlibs  bool
		nostartfiles   bool
		wholeArchive   bool
		noWholeArchive bool
		asNeeded       bool
		noAsNeeded     bool
		buildID        string
		hashStyle      string
		ehFrameHdr     bool
		noEhFrameHdr   bool
		excludeLibs    string
		excludeLibsAll string
		libraryPath    string
		library        string
		framework      string
		frameworkPath  string
		rpath          string
		rpathLink      string
		soname         string
		versionScript  string
		dynamicList    string
		init           string
		fini           string
		preload        string
		wrap           string
		demangle       bool
		help           bool
		version        bool
	}
)

// Add helper to find project root (where aether.toml lives)
func findProjectRoot(start string) string {
	dir := start
	for {
		configPath := filepath.Join(dir, "aether.toml")
		if _, err := os.Stat(configPath); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "."
}

func doBuild(args []string) {
	if buildFlags.help {
		fmt.Println("Aether Build Command")
		fmt.Println("Usage: aether build [flags] [files...]")
		fmt.Println("Use --help for more information")
		return
	}

	if buildFlags.version {
		fmt.Println("Aether Compiler v0.2.0-tinygo1")
		return
	}

	// Determine what to build based on arguments
	var filesToBuild []string
	var buildDir string
	var projectRoot string

	if len(args) == 0 {
		// No arguments - build current directory
		buildDir = "."
		files, err := findAeFiles("src")
		must(err)
		filesToBuild = files
		projectRoot = findProjectRoot(".")
		if !buildFlags.quiet {
			fmt.Println("Building aether project in current directory.")
		}
	} else {
		// Arguments provided - check if they're files or directories
		for _, arg := range args {
			info, err := os.Stat(arg)
			if err != nil {
				fmt.Printf("Error: Cannot access '%s': %v\n", arg, err)
				os.Exit(1)
			}

			if info.IsDir() {
				// Directory - find all .ae files in it
				files, err := findAeFiles(arg)
				must(err)
				filesToBuild = append(filesToBuild, files...)
				buildDir = arg
				projectRoot = findProjectRoot(arg)
				if !buildFlags.quiet {
					fmt.Printf("Building aether project in directory: %s\n", arg)
				}
			} else {
				// File - add it to the list
				if !strings.HasSuffix(arg, ".ae") {
					fmt.Printf("Error: File '%s' is not an Aether file (.ae)\n", arg)
					os.Exit(1)
				}
				filesToBuild = append(filesToBuild, arg)
				buildDir = filepath.Dir(arg)
				projectRoot = findProjectRoot(buildDir)
				if !buildFlags.quiet {
					fmt.Printf("Building specific file: %s\n", arg)
				}
			}
		}
	}

	if len(filesToBuild) == 0 {
		fmt.Println("No Aether files found to build.")
		os.Exit(1)
	}

	if buildFlags.quiet {
		// Suppress output
	} else {
		if buildFlags.verbose {
			fmt.Printf("Target: %s-%s\n", buildFlags.targetOS, buildFlags.targetArch)
			fmt.Printf("Optimization: %s\n", buildFlags.optimization)
			fmt.Printf("Linker: %s\n", buildFlags.linker)
		}
	}

	// Systematic analysis phase
	if buildFlags.checkImports || buildFlags.analyzeOnly {
		if !buildFlags.quiet {
			fmt.Println("Performing systematic analysis...")
		}

		// Analyze dependencies
		depAnalysis := compiler.AnalyzeDependencies(projectRoot)
		if !depAnalysis.Valid {
			fmt.Println("Dependency analysis failed:")
			summary := utils.GroupErrorsByFile(depAnalysis.Errors)
			fmt.Print(utils.FormatErrorSummary(summary))
			os.Exit(1)
		}

		if len(depAnalysis.Warnings) > 0 {
			fmt.Println("Dependency warnings:")
			for _, warning := range depAnalysis.Warnings {
				fmt.Println("  Warning:", warning)
			}
		}

		// Generate lock file if needed
		if err := compiler.GenerateLockFile(projectRoot); err != nil {
			fmt.Println("Failed to generate lock file:", err)
			os.Exit(1)
		}

		imports, err := analyzeImports(filesToBuild)
		must(err)

		if detectCycles(imports) {
			fmt.Println("Error: Circular imports detected")
			os.Exit(1)
		}

		if buildFlags.verbose {
			fmt.Println("Import analysis complete:")
			for file, deps := range imports {
				fmt.Printf("  %s -> %v\n", file, deps)
			}
		}

		if buildFlags.analyzeOnly {
			fmt.Println("Analysis complete. No compilation performed.")
			return
		}
	}

	// Compilation phase
	if !buildFlags.emitExe && !buildFlags.emitObj && !buildFlags.emitIR && !buildFlags.emitASM && !buildFlags.emitBitcode {
		buildFlags.emitExe = true
	}

	imports, err := analyzeImports(filesToBuild)
	must(err)

	sortedFiles, err := topoSort(filesToBuild, imports)
	must(err)

	if !buildFlags.quiet {
		fmt.Println("Compiling", len(sortedFiles), "files...")
	}

	var objectFiles []string
	var allParseErrors []utils.ParseError
	var moduleSymbols map[string]map[string]interface{} = make(map[string]map[string]interface{})

	for _, file := range sortedFiles {
		if buildFlags.verbose {
			fmt.Printf("  Compiling %s\n", file)
		}

		content, err := os.ReadFile(file)
		must(err)

		l := lexer.NewLexer(string(content))

		// Emit tokens if requested
		if buildFlags.emitTokens {
			fmt.Printf("=== Tokens for %s ===\n", file)
			tokens := l.Tokenize()
			for i, tok := range tokens {
				fmt.Printf("%3d: %-12s '%s' (line %d, col %d)\n",
					i, tok.Type, tok.Literal, tok.Line, tok.Column)
			}
			fmt.Println()
		}

		p := parser.NewParser(l)
		p.SetFile(file)
		ast := p.Parse()

		if len(p.Errors.Errors) > 0 {
			for _, err := range p.Errors.Errors {
				allParseErrors = append(allParseErrors, err)
			}
			continue
		}

		// Extract module symbols for linking
		moduleName := strings.TrimSuffix(filepath.Base(file), ".ae")
		moduleSymbols[moduleName] = extractModuleSymbols(ast)

		// Compile with enhanced options
		ir := compiler.CompileWithOptionsAndModules(ast, moduleName, moduleSymbols)

		baseName := strings.TrimSuffix(file, ".ae")

		// Generate different outputs based on flags
		if buildFlags.emitIR || buildFlags.emitLLVM {
			llFile := baseName + ".ll"
			must(os.WriteFile(llFile, []byte(ir), 0644))
			if buildFlags.verbose {
				fmt.Printf("    Generated IR: %s\n", llFile)
			}
		}

		if buildFlags.emitASM {
			asmFile := baseName + ".s"
			generateAssembly(ir, asmFile)
			if buildFlags.verbose {
				fmt.Printf("    Generated ASM: %s\n", asmFile)
			}
		}

		if buildFlags.emitBitcode {
			bcFile := baseName + ".bc"
			generateBitcode(ir, bcFile)
			if buildFlags.verbose {
				fmt.Printf("    Generated Bitcode: %s\n", bcFile)
			}
		}

		if buildFlags.emitObj || buildFlags.emitExe {
			objFile := baseName + ".o"
			objectFiles = append(objectFiles, objFile)
			generateObjectFile(ir, objFile)
			if buildFlags.verbose {
				fmt.Printf("    Generated Object: %s\n", objFile)
			}
		}
	}

	if len(allParseErrors) > 0 {
		summary := utils.GroupErrorsByFile(allParseErrors)
		fmt.Print(utils.FormatErrorSummary(summary))
		os.Exit(1)
	}

	// Linking phase
	if buildFlags.emitExe && len(objectFiles) > 0 {
		if !buildFlags.quiet {
			fmt.Println("Linking object files...")
		}

		output := buildFlags.outputName
		linkObjectFiles(objectFiles, output)

		if !buildFlags.quiet {
			fmt.Println("Build complete! Executable at:", output)
		}
	}
}

var BuildCmd = &cobra.Command{
	Use:   "build [flags] [files...]",
	Short: "Build the current aether project",
	Long: `Build the current aether project with comprehensive optimization and analysis options.

Examples:
  aether build                    # Build with default settings
  aether build -O2               # Build with optimization level 2
  aether build --debug-info      # Build with debug information
  aether build --target-os=linux --target-arch=arm64  # Cross-compile
  aether build --emit-ir         # Only generate LLVM IR
  aether build --analyze-only    # Only analyze, don't compile`,
	Run: func(cmd *cobra.Command, args []string) {
		doBuild(args)
	},
}

func init() {
	flags := BuildCmd.Flags()

	// Basic build flags
	flags.BoolVar(&buildFlags.noStdlib, "no-stdlib", false, "disable stdlib builtins")
	flags.StringVar(&buildFlags.targetOS, "target-os", runtime.GOOS, "target operating system (linux, darwin, windows)")
	flags.StringVar(&buildFlags.targetArch, "target-arch", runtime.GOARCH, "target architecture (amd64, arm64, 386, arm)")
	flags.StringVar(&buildFlags.linker, "linker", getDefaultLinker(), "linker to use (mold, ld, lld)")
	flags.StringVar(&buildFlags.outputName, "output", "bin/aether.out", "output executable name")
	flags.StringVar(&buildFlags.fuseLd, "fuse-ld", "", "linker to use (like clang -fuse-ld)")

	// Optimization flags
	flags.StringVar(&buildFlags.optimization, "O", "2", "optimization level (0, 1, 2, 3, s, z)")
	flags.BoolVar(&buildFlags.noOptimize, "no-optimize", false, "disable all optimizations")
	flags.BoolVar(&buildFlags.noInline, "no-inline", false, "disable function inlining")
	flags.BoolVar(&buildFlags.noVectorize, "no-vectorize", false, "disable vectorization")
	flags.BoolVar(&buildFlags.noUnroll, "no-unroll", false, "disable loop unrolling")

	// Debug flags
	flags.BoolVar(&buildFlags.debugInfo, "debug-info", false, "generate debug information")
	flags.BoolVar(&buildFlags.debugSymbols, "debug-symbols", false, "include debug symbols")
	flags.BoolVar(&buildFlags.strip, "strip", false, "strip debug symbols from output")

	// Output flags
	flags.BoolVar(&buildFlags.emitIR, "emit-ir", false, "emit LLVM IR (.ll)")
	flags.BoolVar(&buildFlags.emitASM, "emit-asm", false, "emit assembly (.s)")
	flags.BoolVar(&buildFlags.emitBitcode, "emit-bitcode", false, "emit bitcode (.bc)")
	flags.BoolVar(&buildFlags.emitLLVM, "emit-llvm", false, "emit LLVM IR (alias for --emit-ir)")
	flags.BoolVar(&buildFlags.emitObj, "emit-obj", false, "emit object files (.o)")
	flags.BoolVar(&buildFlags.emitExe, "emit-exe", true, "emit executable")
	flags.BoolVar(&buildFlags.emitTokens, "emit-tokens", false, "emit lexer tokens for debugging")

	// Analysis flags
	flags.BoolVar(&buildFlags.checkImports, "check-imports", true, "check import validity")
	flags.BoolVar(&buildFlags.analyzeOnly, "analyze-only", false, "only analyze, don't compile")

	// Performance flags
	flags.BoolVar(&buildFlags.parallel, "parallel", true, "enable parallel compilation")
	flags.IntVar(&buildFlags.threads, "threads", runtime.NumCPU(), "number of compilation threads")
	flags.BoolVar(&buildFlags.timeCompile, "time-compile", false, "time compilation phases")
	flags.BoolVar(&buildFlags.stats, "stats", false, "show compilation statistics")
	flags.BoolVar(&buildFlags.profile, "profile", false, "enable profiling")

	// Code generation flags
	flags.StringVar(&buildFlags.stackProtector, "stack-protector", "strong", "stack protector mode (none, basic, strong, all)")
	flags.StringVar(&buildFlags.relocModel, "reloc-model", "pic", "relocation model (static, pic, dynamic-no-pic)")
	flags.StringVar(&buildFlags.codeModel, "code-model", "small", "code model (tiny, small, kernel, medium, large)")
	flags.StringVar(&buildFlags.cpu, "cpu", "generic", "target CPU")
	flags.StringVar(&buildFlags.features, "features", "", "target features (e.g., +sse4.2)")

	// Sanitizer flags
	flags.StringVar(&buildFlags.sanitize, "sanitize", "", "sanitizer to use (address, thread, memory, undefined)")

	// Linking flags
	flags.BoolVar(&buildFlags.pie, "pie", false, "create position independent executable")
	flags.BoolVar(&buildFlags.static, "static", false, "create static executable")
	flags.BoolVar(&buildFlags.shared, "shared", false, "create shared library")
	flags.BoolVar(&buildFlags.rdynamic, "rdynamic", false, "add all symbols to dynamic symbol table")
	flags.BoolVar(&buildFlags.exportDynamic, "export-dynamic", false, "export all symbols")

	// Library flags
	flags.BoolVar(&buildFlags.noStartFiles, "no-start-files", false, "don't link startup files")
	flags.BoolVar(&buildFlags.noDefaultLibs, "no-default-libs", false, "don't link default libraries")
	flags.BoolVar(&buildFlags.nostdlib, "nostdlib", false, "don't link standard library")
	flags.BoolVar(&buildFlags.nodefaultlibs, "nodefaultlibs", false, "don't link default libraries")
	flags.BoolVar(&buildFlags.nostartfiles, "nostartfiles", false, "don't link startup files")

	// Archive flags
	flags.BoolVar(&buildFlags.wholeArchive, "whole-archive", false, "include all objects from archives")
	flags.BoolVar(&buildFlags.noWholeArchive, "no-whole-archive", false, "don't include all objects from archives")
	flags.BoolVar(&buildFlags.asNeeded, "as-needed", false, "link libraries only when needed")
	flags.BoolVar(&buildFlags.noAsNeeded, "no-as-needed", false, "always link libraries")

	// Advanced linking flags
	flags.StringVar(&buildFlags.buildID, "build-id", "", "generate build ID")
	flags.StringVar(&buildFlags.hashStyle, "hash-style", "", "hash style (sysv, gnu, both)")
	flags.BoolVar(&buildFlags.ehFrameHdr, "eh-frame-hdr", true, "generate .eh_frame_hdr section")
	flags.BoolVar(&buildFlags.noEhFrameHdr, "no-eh-frame-hdr", false, "don't generate .eh_frame_hdr section")

	// Library and framework flags
	flags.StringVar(&buildFlags.excludeLibs, "exclude-libs", "", "exclude libraries from linking")
	flags.StringVar(&buildFlags.excludeLibsAll, "exclude-libs-all", "", "exclude all libraries from linking")
	flags.StringVar(&buildFlags.libraryPath, "library-path", "", "add library search path")
	flags.StringVar(&buildFlags.library, "library", "", "link against library")
	flags.StringVar(&buildFlags.framework, "framework", "", "link against framework")
	flags.StringVar(&buildFlags.frameworkPath, "framework-path", "", "add framework search path")

	// Runtime flags
	flags.StringVar(&buildFlags.rpath, "rpath", "", "set runtime library search path")
	flags.StringVar(&buildFlags.rpathLink, "rpath-link", "", "set runtime library search path for dependencies")
	flags.StringVar(&buildFlags.soname, "soname", "", "set shared object name")
	flags.StringVar(&buildFlags.versionScript, "version-script", "", "version script file")
	flags.StringVar(&buildFlags.dynamicList, "dynamic-list", "", "dynamic list file")

	// Initialization flags
	flags.StringVar(&buildFlags.init, "init", "", "initialization function")
	flags.StringVar(&buildFlags.fini, "fini", "", "finalization function")
	flags.StringVar(&buildFlags.preload, "preload", "", "preload library")
	flags.StringVar(&buildFlags.wrap, "wrap", "", "wrap symbol")

	// Output control flags
	flags.BoolVar(&buildFlags.demangle, "demangle", false, "demangle symbol names")
	flags.BoolVar(&buildFlags.verbose, "verbose", false, "verbose output")
	flags.BoolVar(&buildFlags.quiet, "quiet", false, "suppress output")
	flags.BoolVar(&buildFlags.help, "help", false, "show help")
	flags.BoolVar(&buildFlags.version, "version", false, "show version")
}

func getOptimizationLevel() string {
	if buildFlags.noOptimize {
		return "default<O0>"
	}

	switch buildFlags.optimization {
	case "0":
		return "default<O0>"
	case "1":
		return "default<O1>"
	case "2":
		return "default<O2>"
	case "3":
		return "default<O3>"
	case "s":
		return "default<Os>"
	case "z":
		return "default<Oz>"
	default:
		return "default<O2>"
	}
}

func generateAssembly(ir string, outputFile string) {
	// Generate assembly from IR
	llFile := strings.TrimSuffix(outputFile, ".s") + ".ll"
	must(os.WriteFile(llFile, []byte(ir), 0644))

	cmd := exec.Command("llc", "-filetype=asm", llFile, "-o", outputFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	must(cmd.Run())
}

func generateBitcode(ir string, outputFile string) {
	// Generate bitcode from IR
	llFile := strings.TrimSuffix(outputFile, ".bc") + ".ll"
	must(os.WriteFile(llFile, []byte(ir), 0644))

	cmd := exec.Command("llvm-as", llFile, "-o", outputFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	must(cmd.Run())
}

func generateObjectFile(ir string, outputFile string) {
	// Generate object file from IR
	llFile := strings.TrimSuffix(outputFile, ".o") + ".ll"
	must(os.WriteFile(llFile, []byte(ir), 0644))

	cmd := exec.Command("llc", "-filetype=obj", llFile, "-o", outputFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	must(cmd.Run())
}

func linkObjectFiles(objectFiles []string, output string) {
	args := append(objectFiles, "-o", output)

	// Add linker-specific flags
	if buildFlags.fuseLd != "" {
		args = append([]string{"-fuse-ld=" + buildFlags.fuseLd}, args...)
	}

	// Add target-specific libraries
	addTargetLibraries(&args)

	// Add optimization flags
	if buildFlags.optimization != "0" {
		args = append(args, "-O"+buildFlags.optimization)
	}

	// Add debug flags
	if buildFlags.debugSymbols && !buildFlags.strip {
		args = append(args, "-g")
	}

	if buildFlags.strip {
		args = append(args, "-s")
	}

	// Add sanitizer flags
	if buildFlags.sanitize != "" {
		args = append(args, "-fsanitize="+buildFlags.sanitize)
	}

	// Add linking flags
	if buildFlags.static {
		args = append(args, "-static")
	}

	if buildFlags.shared {
		args = append(args, "-shared")
	}

	if buildFlags.pie {
		args = append(args, "-pie")
	}

	if buildFlags.rdynamic {
		args = append(args, "-rdynamic")
	}

	if buildFlags.exportDynamic {
		args = append(args, "-export-dynamic")
	}

	// Add library flags
	if buildFlags.nostdlib {
		args = append(args, "-nostdlib")
	}

	if buildFlags.nodefaultlibs {
		args = append(args, "-nodefaultlibs")
	}

	if buildFlags.nostartfiles {
		args = append(args, "-nostartfiles")
	}

	// Add archive flags
	if buildFlags.wholeArchive {
		args = append(args, "-whole-archive")
	}

	if buildFlags.noWholeArchive {
		args = append(args, "-no-whole-archive")
	}

	if buildFlags.asNeeded {
		args = append(args, "-as-needed")
	}

	if buildFlags.noAsNeeded {
		args = append(args, "-no-as-needed")
	}

	// Add runtime flags
	if buildFlags.rpath != "" {
		args = append(args, "-rpath", buildFlags.rpath)
	}

	if buildFlags.rpathLink != "" {
		args = append(args, "-rpath-link", buildFlags.rpathLink)
	}

	if buildFlags.soname != "" {
		args = append(args, "-soname", buildFlags.soname)
	}

	// Add library paths and libraries
	if buildFlags.libraryPath != "" {
		args = append(args, "-L"+buildFlags.libraryPath)
	}

	if buildFlags.library != "" {
		args = append(args, "-l"+buildFlags.library)
	}

	if buildFlags.framework != "" {
		args = append(args, "-framework", buildFlags.framework)
	}

	if buildFlags.frameworkPath != "" {
		args = append(args, "-F"+buildFlags.frameworkPath)
	}

	// Add initialization flags
	if buildFlags.init != "" {
		args = append(args, "-init", buildFlags.init)
	}

	if buildFlags.fini != "" {
		args = append(args, "-fini", buildFlags.fini)
	}

	if buildFlags.preload != "" {
		args = append(args, "-preload", buildFlags.preload)
	}

	if buildFlags.wrap != "" {
		args = append(args, "-wrap", buildFlags.wrap)
	}

	// Add output control flags
	if buildFlags.demangle {
		args = append(args, "--demangle")
	}

	cmd := exec.Command(buildFlags.linker, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	must(cmd.Run())
}

func addTargetLibraries(args *[]string) {
	switch buildFlags.targetOS {
	case "linux":
		switch buildFlags.targetArch {
		case "amd64":
			*args = append(*args, "-L/usr/lib/x86_64-linux-gnu", "-L/usr/lib", "/usr/lib/x86_64-linux-gnu/crt1.o", "/usr/lib/x86_64-linux-gnu/crti.o", "-lc", "/usr/lib/x86_64-linux-gnu/crtn.o")
		case "arm64":
			*args = append(*args, "-L/usr/lib/aarch64-linux-gnu", "-L/usr/lib", "/usr/lib/aarch64-linux-gnu/crt1.o", "/usr/lib/aarch64-linux-gnu/crti.o", "-lc", "/usr/lib/aarch64-linux-gnu/crtn.o")
		default:
			*args = append(*args, "-lc")
		}
	case "darwin":
		*args = append(*args, "-L/usr/lib", "-lc")
	case "windows":
		*args = append(*args, "-lkernel32", "-lmsvcrt", "-lucrt", "-loldnames")
	}
}

func getDefaultLinker() string {
	if runtime.GOOS == "windows" {
		return "lld"
	}
	return "mold"
}

func extractModuleSymbols(prog *parser.Program) map[string]interface{} {
	symbols := make(map[string]interface{})
	for _, stmt := range prog.Statements {
		if assign, ok := stmt.(*parser.Assignment); ok {
			if assign.Name.Value[0] >= 'A' && assign.Name.Value[0] <= 'Z' {
				// This is an exported symbol
				symbols[assign.Name.Value] = assign.Value
			}
		}
	}
	return symbols
}
