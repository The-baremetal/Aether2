package main

import (
	"crypto/sha256"
	"encoding/json"
	"io"
	"io/ioutil"
	"aether/src/analysis"
	compiler_pkg "aether/src/compiler"
	"aether/src/lexer"
	"aether/src/parser"
	"aether/src/scheduler"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"aether/lib/utils"
	"aether/src/buildcache"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

// BuildCacheEntry stores metadata for a single file
type BuildCacheEntry struct {
	Hash         string            `json:"hash"`
	Output       string            `json:"output"`
	Deps         []string          `json:"deps"`
	DepHashes    map[string]string `json:"dep_hashes"`
	LastBuild    int64             `json:"last_build"`
}

// BuildCache stores the build cache for the project
type BuildCache struct {
	Files map[string]BuildCacheEntry `json:"files"`
}

func LoadCache(path string) (*BuildCache, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return &BuildCache{Files: make(map[string]BuildCacheEntry)}, nil
	}
	var cache BuildCache
	if err := json.Unmarshal(data, &cache); err != nil {
		return &BuildCache{Files: make(map[string]BuildCacheEntry)}, nil
	}
	return &cache, nil
}

func SaveCache(path string, cache *BuildCache) error {
	data, err := json.MarshalIndent(cache, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0644)
}

func fileHash(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		f.Close()
		return "", err
	}
	f.Close()
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// Project configuration structure

type ProjectConfig struct {
	Project struct {
		Name        string `toml:"name"`
		Version     string `toml:"version"`
		Author      string `toml:"author,omitempty"`
		Description string `toml:"description,omitempty"`
	} `toml:"project"`

	Build struct {
		SourceDirectories []string `toml:"source_directories,omitempty"`
		OutputDirectory   string   `toml:"output_directory,omitempty"`
		Target            string   `toml:"target,omitempty"`
		Optimization      string   `toml:"optimization,omitempty"`
		Linker            string   `toml:"linker,omitempty"`
		CreateLibrary     bool     `toml:"create_library,omitempty"`
		LibraryType       string   `toml:"library_type,omitempty"`
		CompilerFlags     compiler_pkg.CompilerFlags `toml:"compiler_flags,omitempty"`
		Targets           map[string]compiler_pkg.TargetConfig `toml:"target,omitempty"`
	} `toml:"build"`

	Dependencies map[string]string `toml:"dependencies,omitempty"`

	DevDependencies map[string]string `toml:"dev-dependencies,omitempty"`
}

func loadProjectConfig(projectRoot string) ProjectConfig {
	configPath := filepath.Join(projectRoot, "aether.toml")

	var config ProjectConfig

	// Set defaults
	config.Project.Name = "aether-project"
	config.Project.Version = "0.1.0"
	config.Build.SourceDirectories = []string{"src", "."}
	config.Build.OutputDirectory = "bin"
	config.Build.Target = "native"
	config.Build.Optimization = "debug"
	config.Build.Linker = "mold"

	// Try to read aether.toml
	if data, err := os.ReadFile(configPath); err == nil {
		if err := toml.Unmarshal(data, &config); err != nil {
			if !buildFlags.quiet {
				fmt.Printf("Warning: Failed to parse aether.toml: %v\n", err)
			}
		}
	} else {
		if !buildFlags.quiet {
			fmt.Printf("Warning: No aether.toml found, using defaults\n")
		}
	}

	return config
}



func parseTarget(target string) (string, string) {
	switch target {
	case "native":
		return runtime.GOOS, runtime.GOARCH
	case "linux":
		return "linux", runtime.GOARCH
	case "windows":
		return "windows", runtime.GOARCH
	case "darwin":
		return "darwin", runtime.GOARCH
	case "linux-amd64":
		return "linux", "amd64"
	case "linux-arm64":
		return "linux", "arm64"
	case "windows-amd64":
		return "windows", "amd64"
	case "darwin-amd64":
		return "darwin", "amd64"
	case "darwin-arm64":
		return "darwin", "arm64"
	default:
		// Try to parse target like "linux-amd64"
		if strings.Contains(target, "-") {
			parts := strings.Split(target, "-")
			if len(parts) == 2 {
				return parts[0], parts[1]
			}
		}
		return runtime.GOOS, runtime.GOARCH
	}
}

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
		// New library-specific flags
		createLibrary      bool
		libraryType        string // "shared", "static", "both"
		libraryName        string
		libraryVersion     string
		exportSymbols      bool
		generatePkgConfig  bool
		libraryDescription string
		libraryURL         string
		libraryRequires    string
		libraryConflicts   string
		libraryProvides    string
		forceRebuild       bool
		libcType           string
	}

	projectConfig ProjectConfig
)

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
		projectRoot = findProjectRoot(".")

		// Load project configuration from aether.toml
		config := loadProjectConfig(projectRoot)

		// Try to find Aether files based on configuration
		var files []string
		var err error

		// Use configured source directory or fall back to defaults
		sourceDirs := config.Build.SourceDirectories
		if len(sourceDirs) == 0 {
			sourceDirs = []string{"src", "."}
		}

		for _, sourceDir := range sourceDirs {
			files, err = analysis.FindAetherFiles(sourceDir)
			if err == nil && len(files) > 0 {
				if !buildFlags.quiet {
					fmt.Printf("Found %d files in %s\n", len(files), sourceDir)
				}
				break
			}
		}
		if err != nil {
			printSmartError(err, "build")
		}
		filesToBuild = files

		// Apply configuration overrides
		if config.Build.Target != "" {
			buildFlags.targetOS, buildFlags.targetArch = parseTarget(config.Build.Target)
		}
		if config.Build.Optimization != "" {
			buildFlags.optimization = config.Build.Optimization
		}
		if config.Build.Linker != "" {
			buildFlags.linker = config.Build.Linker
		}
		if config.Build.CreateLibrary {
			buildFlags.createLibrary = true
		}
		if config.Build.LibraryType != "" {
			buildFlags.libraryType = config.Build.LibraryType
		}

		// Store config for later use
		projectConfig = config

		// After loading projectConfig, detect Windows and set linker/fuseLd accordingly
		if runtime.GOOS == "windows" {
			buildFlags.linker = "lld-link"
			buildFlags.fuseLd = "lld"
			projectConfig.Build.Linker = "lld-link"
		}

		// Initialize flag merger for compiler flags
		flagMerger := compiler_pkg.NewFlagMerger()
		flagMerger.SetConfigFlags(config.Build.CompilerFlags, config.Build.Targets)
		flagMerger.SetTargetOS(buildFlags.targetOS)
		flagMerger.SetOptimization(buildFlags.optimization)
		flagMerger.SetDebugInfo(buildFlags.debugInfo)

		// Validate compiler flags
		if err := flagMerger.ValidateFlags(); err != nil {
			fmt.Printf("Warning: Invalid compiler flags: %v\n", err)
		}

		if !buildFlags.quiet {
			fmt.Printf("Building aether project '%s' v%s\n", config.Project.Name, config.Project.Version)
			if buildFlags.verbose {
				fmt.Print(flagMerger.GetFlagSummary())
			}
		}
	} else {
		// Arguments provided - check if they're files or directories
		for _, arg := range args {
			info, err := os.Stat(arg)
			if err != nil {
				printSmartError(err, arg)
			}

			if info.IsDir() {
				// Directory - find all .ae files in it
				files, err := analysis.FindAetherFiles(arg)
				if err != nil {
					printSmartError(err, arg)
				}
				filesToBuild = append(filesToBuild, files...)
				buildDir = arg
				projectRoot = findProjectRoot(arg)
				if !buildFlags.quiet {
					fmt.Printf("Building aether project in directory: %s\n", arg)
				}
			} else {
				// File - add it to the list
				if !strings.HasSuffix(arg, ".aeth") {
					printSmartError(fmt.Errorf("File '%s' is not an Aether file (.aeth)", arg), arg)
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
		printSmartError(fmt.Errorf("No Aether files found to build."), "build")
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
		depAnalysis := analysis.AnalyzeDependencies(projectRoot)
		if !depAnalysis.Valid {
			printSmartError(fmt.Errorf("Dependency analysis failed"), "dependency analysis")
		}

		if len(depAnalysis.Warnings) > 0 {
			fmt.Println("Dependency warnings:")
			for _, warning := range depAnalysis.Warnings {
				fmt.Println("  Warning:", warning)
			}
		}

		// Generate lock file if needed
		if err := analysis.GenerateLockFile(projectRoot); err != nil {
			printSmartError(fmt.Errorf("Failed to generate lock file: %v", err), "lock file")
		}

		imports, err := analysis.AnalyzeImports(filesToBuild)
		if err != nil {
			printSmartError(err, "import analysis")
		}

		if len(scheduler.DetectCycles(imports)) > 0 {
			printSmartError(fmt.Errorf("Circular imports detected"), "import analysis")
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

	imports, err := analysis.AnalyzeImports(filesToBuild)
	if err != nil {
		printSmartError(err, "import analysis")
	}

	// Resolve import paths to actual files
	importedFiles, err := analysis.ResolveImportPathsToFiles(imports, projectRoot)
	if err != nil {
		printSmartError(err, "import resolution")
	}

	// Create a mapping from import names to resolved file paths
	importNameToPath := make(map[string]string)
	for _, importPath := range importedFiles {
		importName := filepath.Base(importPath)
		importName = strings.TrimSuffix(importName, ".ae")
		importNameToPath[importName] = importPath
	}

	// Update imports map to use resolved file paths instead of import names
	resolvedImports := make(map[string][]string)
	for sourceFile, importNames := range imports {
		var resolvedPaths []string
		for _, importName := range importNames {
			if resolvedPath, exists := importNameToPath[importName]; exists {
				resolvedPaths = append(resolvedPaths, resolvedPath)
			}
		}
		resolvedImports[sourceFile] = resolvedPaths
	}

	// Ensure every file in jobs will be a key in resolvedImports
	for _, f := range filesToBuild {
		if _, ok := resolvedImports[f]; !ok {
			resolvedImports[f] = []string{}
		}
	}
	for _, f := range importedFiles {
		if _, ok := resolvedImports[f]; !ok {
			resolvedImports[f] = []string{}
		}
	}

	// Combine source files with imported files, ensuring uniqueness
	fileSet := make(map[string]struct{})
	for _, f := range filesToBuild {
		fileSet[f] = struct{}{}
	}
	for _, f := range importedFiles {
		fileSet[f] = struct{}{}
	}
	var allFiles []string
	for f := range fileSet {
		allFiles = append(allFiles, f)
	}

	sortedFiles, err := scheduler.TopoSort(resolvedImports)
	if err != nil {
		printSmartError(err, "topological sort")
	}

	if !buildFlags.quiet {
		fmt.Println("Compiling", len(sortedFiles), "files...")
	}

	var objectFiles []string
	var allParseErrors []utils.ParseError
	var moduleSymbols map[string]map[string]interface{} = make(map[string]map[string]interface{})
	entryFile := ""
	if len(filesToBuild) > 0 {
		entryFile = filesToBuild[0]
	}

	cachePath := filepath.Join(projectRoot, ".aetherbuildcache.json")
	cache, _ := buildcache.LoadCache(cachePath)
	cacheMu := &sync.Mutex{}

	stale := make(map[string]bool)
	reasons := make(map[string]string)

	// Helper to check if a file is stale
	var isStale func(file string) bool
	isStale = func(file string) bool {
		entry, ok := cache.Files[file]
		hash, err := buildcache.FileHash(file)
		if err != nil {
			stale[file] = true
			reasons[file] = "missing or unreadable"
			return true
		}
		if !ok || entry.Hash != hash {
			stale[file] = true
			reasons[file] = "changed"
			return true
		}
		// Check output exists
		outputDir := filepath.Join(projectRoot, projectConfig.Build.OutputDirectory)
		if outputDir == "" {
			outputDir = "bin"
		}
		baseName := strings.TrimSuffix(filepath.Base(file), ".ae")
		output := filepath.Join(outputDir, baseName+".o")
		if _, err := os.Stat(output); err != nil {
			stale[file] = true
			reasons[file] = "missing output"
			return true
		}
		// Check dependencies
		for _, dep := range resolvedImports[file] {
			if isStale(dep) {
				stale[file] = true
				reasons[file] = fmt.Sprintf("dependency %s changed", dep)
				return true
			}
			depHash, _ := buildcache.FileHash(dep)
			if entry.DepHashes == nil || entry.DepHashes[dep] != depHash {
				stale[file] = true
				reasons[file] = fmt.Sprintf("dependency %s changed", dep)
				return true
			}
		}
		return false
	}

	if buildFlags.forceRebuild {
		for _, file := range sortedFiles {
			stale[file] = true
			reasons[file] = "forced rebuild"
		}
	} else {
		for _, file := range sortedFiles {
			isStale(file)
		}
	}

	jobs := make(map[string]func())
	objectFilesMu := &sync.Mutex{}
	parseErrorsMu := &sync.Mutex{}

	for _, file := range sortedFiles {
		if !stale[file] {
			if buildFlags.verbose {
				fmt.Printf("Up to date: %s\n", file)
			}
			continue
		}
		f := file
		jobs[f] = func() {
			if buildFlags.verbose {
				fmt.Printf("SUBMITTED JOB FOR: %s\n", f)
			}
			content, err := os.ReadFile(f)
			if err != nil {
				printSmartError(err, f)
			}
			l := lexer.NewLexer(string(content))
			if buildFlags.emitTokens {
				fmt.Printf("=== Tokens for %s ===\n", f)
				tokens := l.Tokenize()
				for i, tok := range tokens {
					fmt.Printf("%3d: %-12s '%s' (line %d, col %d)\n", i, tok.Type, tok.Literal, tok.Line, tok.Column)
				}
				fmt.Println()
			}
			p := parser.NewParser(l)
			p.SetFile(f)
			if f == entryFile {
				p.IsEntryFile = true
			} else {
				p.IsEntryFile = false
			}
			ast := p.Parse()
			if len(p.Errors.Errors) > 0 {
				parseErrorsMu.Lock()
				for _, err := range p.Errors.Errors {
					allParseErrors = append(allParseErrors, err)
				}
				parseErrorsMu.Unlock()
				return
			}
			moduleName := strings.TrimSuffix(filepath.Base(f), ".ae")
			moduleSymbols[moduleName] = extractModuleSymbols(ast)
			// When calling CompileWithOptionsAndModules, determine the function name robustly:
			// If this is the entry module, use 'main'. Otherwise, use '__module_' + moduleName.
			// Pass this as the funcName argument.
			funcName := "main"
			if f != entryFile {
				funcName = "__module_" + moduleName
			}
			ir := compiler_pkg.CompileWithOptionsAndModules(ast, moduleName, moduleSymbols, buildFlags.verbose, funcName)
			outputDir := filepath.Join(projectRoot, projectConfig.Build.OutputDirectory)
			fmt.Println(outputDir)
			if outputDir == "" {
				outputDir = "bin"
			}
			_ = os.MkdirAll(outputDir, 0755)
			baseName := strings.TrimSuffix(filepath.Base(f), ".ae")
			basePath := filepath.Join(outputDir, baseName)
			if buildFlags.emitIR || buildFlags.emitLLVM {
				llFile := basePath + ".ll"
				if err := os.WriteFile(llFile, []byte(ir), 0644); err != nil {
					printSmartError(err, llFile)
				}
				if buildFlags.verbose {
					fmt.Printf("    Generated IR: %s\n", llFile)
				}
			}
			if buildFlags.emitASM {
				asmFile := basePath + ".s"
				generateAssembly(ir, asmFile)
				if buildFlags.verbose {
					fmt.Printf("    Generated ASM: %s\n", asmFile)
				}
			}
			if buildFlags.emitBitcode {
				bcFile := basePath + ".bc"
				generateBitcode(ir, bcFile)
				if buildFlags.verbose {
					fmt.Printf("    Generated Bitcode: %s\n", bcFile)
				}
			}
			if buildFlags.emitObj || buildFlags.emitExe {
				objFile := basePath + ".o"
				objectFilesMu.Lock()
				objectFiles = append(objectFiles, objFile)
				objectFilesMu.Unlock()
				generateObjectFile(ir, objFile)
				if buildFlags.verbose {
					fmt.Printf("    Generated Object: %s\n", objFile)
				}
			}
			// Update cache after build
			fileHashVal, _ := buildcache.FileHash(f)
			depHashes := make(map[string]string)
			for _, dep := range resolvedImports[f] {
				dh, _ := buildcache.FileHash(dep)
				depHashes[dep] = dh
			}
			cacheMu.Lock()
			cache.Files[f] = buildcache.BuildCacheEntry{
				Hash:      fileHashVal,
				Output:    basePath + ".o",
				Deps:      resolvedImports[f],
				DepHashes: depHashes,
				LastBuild: time.Now().Unix(),
			}
			cacheMu.Unlock()
		}
		if reason, ok := reasons[file]; ok {
			fmt.Printf("Rebuilding %s (%s)\n", file, reason)
		}
	}

	// Before scheduling jobs
	if buildFlags.verbose {
		fmt.Printf("JOBS KEYS: %v\n", getMapKeys(jobs))
		fmt.Printf("RESOLVED IMPORTS KEYS: %v\n", getMapKeys(resolvedImports))
	}

	maxWorkers := buildFlags.threads
	pool := scheduler.NewPool(maxWorkers)
	scheduler.RunBatchesDebug(jobs, resolvedImports, pool, buildFlags.verbose)

	if len(allParseErrors) > 0 {
		summary := utils.GroupErrorsByFile(allParseErrors)
		fmt.Print(utils.FormatErrorSummary(summary))
		os.Exit(1)
	}

	_ = buildcache.SaveCache(cachePath, cache)

	// Linking phase
	if buildFlags.emitExe && len(objectFiles) > 0 {
		if !buildFlags.quiet {
			fmt.Println("Linking object files...")
		}

		output := buildFlags.outputName
		if output == "bin/aether.out" && projectConfig.Build.OutputDirectory != "" {
			// Use configured output directory
			output = filepath.Join(projectConfig.Build.OutputDirectory, "aether.out")
		}
		linkObjectFiles(objectFiles, output)

		if !buildFlags.quiet {
			fmt.Println("Build complete! Executable at:", output)
		}
	}

	// Library creation phase
	if buildFlags.createLibrary && len(objectFiles) > 0 {
		if !buildFlags.quiet {
			fmt.Println("Creating library...")
		}

		outputBase := buildFlags.outputName
		if outputBase == "bin/aether.out" {
			if projectConfig.Build.OutputDirectory != "" {
				outputBase = filepath.Join(projectConfig.Build.OutputDirectory, "aether")
			} else {
				outputBase = "lib/aether"
			}
		}

		// Create the library
		createLibrary(objectFiles, outputBase, buildFlags.libraryType)

		// Generate pkg-config file if requested
		libName := getLibraryName(outputBase)
		generatePkgConfigFile(libName, outputBase)

		if !buildFlags.quiet {
			fmt.Println("Library creation complete!")
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
	flags.StringVarP(&buildFlags.outputName, "output", "o", "bin/aether.out", "output executable name")
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

	// New library-specific flags
	flags.BoolVar(&buildFlags.createLibrary, "create-library", false, "create a new library (shared, static, or both)")
	flags.StringVar(&buildFlags.libraryType, "library-type", "shared", "type of library to create (shared, static, both)")
	flags.StringVar(&buildFlags.libraryName, "library-name", "", "name of the library to create")
	flags.StringVar(&buildFlags.libraryVersion, "library-version", "", "version of the library to create")
	flags.BoolVar(&buildFlags.exportSymbols, "export-symbols", false, "export all symbols from the library")
	flags.BoolVar(&buildFlags.generatePkgConfig, "generate-pkg-config", false, "generate a pkg-config file for the library")
	flags.StringVar(&buildFlags.libraryDescription, "library-description", "", "description for the library")
	flags.StringVar(&buildFlags.libraryURL, "library-url", "", "URL for the library")
	flags.StringVar(&buildFlags.libraryRequires, "library-requires", "", "libraries required by the library")
	flags.StringVar(&buildFlags.libraryConflicts, "library-conflicts", "", "libraries conflicting with the library")
	flags.StringVar(&buildFlags.libraryProvides, "library-provides", "", "libraries provided by the library")
	flags.BoolVar(&buildFlags.forceRebuild, "force-rebuild", false, "force rebuild all files, ignoring cache")
	flags.StringVar(&buildFlags.libcType, "libc-type", "glibc", "type of C library to link against (glibc, musl)")
}
