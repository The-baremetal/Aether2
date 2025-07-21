package main

import (
  "fmt"
  "os"
  "os/exec"
  "runtime"
  "path/filepath"
  "strings"
)

func getLibraryName(outputBase string) string {
  if buildFlags.libraryName != "" {
    return buildFlags.libraryName
  }
  return filepath.Base(outputBase)
}

func getSharedLibraryPath(libName string) string {
  ext := getSharedLibraryExtension()
  return fmt.Sprintf("lib%s.%s", libName, ext)
}

func getDefaultLinker() string {
  if runtime.GOOS == "windows" {
    return "lld"
  }
  return "mold"
}

func getInstallPrefix() string {
  return "/usr/local"
}

func getSharedLibraryExtension() string {
  switch buildFlags.targetOS {
  case "windows":
    return "dll"
  case "darwin":
    return "dylib"
  default:
    return "so"
  }
}

func createSharedLibrary(objectFiles []string, outputBase string) {
  libName := getLibraryName(outputBase)
  outputFile := getSharedLibraryPath(libName)
  args := append(objectFiles, "-o", outputFile)
  args = append(args, "-shared")
  args = append(args, "-fPIC")
  if buildFlags.soname != "" {
    args = append(args, "-soname", buildFlags.soname)
  } else {
    args = append(args, "-soname", libName)
  }
  if buildFlags.exportSymbols {
    args = append(args, "-export-dynamic")
  }
  if buildFlags.fuseLd != "" {
    args = append([]string{"-fuse-ld=" + buildFlags.fuseLd}, args...)
  }
  cmd := exec.Command(buildFlags.linker, args...)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  must(cmd.Run())
  if !buildFlags.quiet {
    fmt.Printf("Created shared library: %s\n", outputFile)
  }
}

func createStaticLibrary(objectFiles []string, outputBase string) {
  libName := getLibraryName(outputBase)
  outputFile := getStaticLibraryPath(libName)
  args := append([]string{"rcs", outputFile}, objectFiles...)
  cmd := exec.Command("ar", args...)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  must(cmd.Run())
  if !buildFlags.quiet {
    fmt.Printf("Created static library: %s\n", outputFile)
  }
}

func generatePkgConfigContent(libName string) string {
  version := buildFlags.libraryVersion
  if version == "" {
    version = "1.0.0"
  }
  description := buildFlags.libraryDescription
  if description == "" {
    description = fmt.Sprintf("Aether library %s", libName)
  }
  url := buildFlags.libraryURL
  if url == "" {
    url = "https://github.com/aether-lang"
  }
  requires := buildFlags.libraryRequires
  conflicts := buildFlags.libraryConflicts
  provides := buildFlags.libraryProvides
  content := fmt.Sprintf(`prefix=%s
exec_prefix=${prefix}
libdir=${exec_prefix}/lib
includedir=${prefix}/include

Name: %s
Description: %s
Version: %s
URL: %s
`, getInstallPrefix(), libName, description, version, url)
  if requires != "" {
    content += fmt.Sprintf("Requires: %s\n", requires)
  }
  if conflicts != "" {
    content += fmt.Sprintf("Conflicts: %s\n", conflicts)
  }
  if provides != "" {
    content += fmt.Sprintf("Provides: %s\n", provides)
  }
  content += fmt.Sprintf(`
Libs: -L${libdir} -l%s
Cflags: -I${includedir}
`, libName)
  return content
}

func getStaticLibraryPath(libName string) string {
  ext := getStaticLibraryExtension()
  return fmt.Sprintf("lib%s.%s", libName, ext)
}

func getStaticLibraryExtension() string {
  switch buildFlags.targetOS {
  case "windows":
    return "lib"
  default:
    return "a"
  }
}

func detectToolchain() string {
  // Check for MSVC
  if os.Getenv("VSINSTALLDIR") != "" || os.Getenv("VCINSTALLDIR") != "" {
    return "msvc"
  }
  if _, err := exec.LookPath("cl.exe"); err == nil {
    return "msvc"
  }
  // Check for MinGW
  if _, err := exec.LookPath("gcc.exe"); err == nil {
    return "mingw"
  }
  if _, err := exec.LookPath("g++.exe"); err == nil {
    return "mingw"
  }
  if _, err := exec.LookPath("mingw32-gcc.exe"); err == nil {
    return "mingw"
  }
  return "unknown"
}

func getMSVCLibPath() string {
  // Try common env vars
  if lib := os.Getenv("LIB"); lib != "" {
    paths := strings.Split(lib, ";")
    if len(paths) > 0 {
      return paths[0]
    }
  }
  // Fallback: try Program Files
  pf := os.Getenv("ProgramFiles(x86)")
  if pf == "" {
    pf = os.Getenv("ProgramFiles")
  }
  if pf != "" {
    return filepath.Join(pf, "Microsoft Visual Studio", "2022", "Community", "VC", "Tools", "MSVC")
  }
  return ""
}

func getMinGWLibPath() string {
  if mingw := os.Getenv("MINGW_HOME"); mingw != "" {
    return filepath.Join(mingw, "lib")
  }
  if path, err := exec.LookPath("gcc.exe"); err == nil {
    base := filepath.Dir(filepath.Dir(path))
    return filepath.Join(base, "lib")
  }
  return ""
}

func linkObjectFiles(objectFiles []string, output string) {
  var args []string
  if runtime.GOOS == "windows" {
    toolchain := detectToolchain()
    if buildFlags.verbose {
      fmt.Printf("🍕 Detected toolchain: %s\n", toolchain)
    }
    // Always use MSVC flags if linker is lld-link or link.exe
    if buildFlags.linker == "lld-link" || buildFlags.linker == "link.exe" {
      if toolchain == "unknown" {
        wslHint := ""
        if runtime.GOOS == "windows" {
          if _, err := exec.LookPath("wsl.exe"); err == nil {
            wslHint = "\n🦄🍕 Pssst! I see you have WSL installed. You might have better luck building in WSL! Try opening a WSL shell and running 'aether build' there for a smoother experience. 🍕🐧"
          }
        }
        fmt.Println("🦄🍕 Oops! I can't find MSVC or MinGW.\nTo build on Windows, please install MinGW (recommended) or Visual Studio Build Tools.\nIf you want to use MinGW, set your linker to ld.lld or clang, not lld-link!\nIf you want to use MSVC, install Visual Studio and run from the Developer Command Prompt.\nNo pizza for you until you feed me a toolchain! 🥲" + wslHint + "\nI am ANGRY! Exiting with error! 😡")
        os.Exit(1)
      }
      args = append(args, "/OUT:"+output)
      args = append(args, objectFiles...)
      libPath := getMSVCLibPath()
      if libPath != "" {
        args = append(args, "/LIBPATH:"+libPath)
      }
      args = append(args, "kernel32.lib", "msvcrt.lib", "ucrt.lib", "oldnames.lib")
      if buildFlags.debugSymbols && !buildFlags.strip {
        args = append(args, "/DEBUG")
      }
      if buildFlags.optimization != "0" {
        args = append(args, "/OPT:REF")
      }
    } else if toolchain == "mingw" {
      args = append(objectFiles, "-o", output)
      libPath := getMinGWLibPath()
      if libPath != "" {
        args = append(args, "-L"+libPath)
      }
      args = append(args, "-lkernel32", "-lmsvcrt", "-lucrt", "-loldnames")
      if buildFlags.debugSymbols && !buildFlags.strip {
        args = append(args, "-g")
      }
      if buildFlags.optimization != "0" {
        args = append(args, "-O"+buildFlags.optimization)
      }
      if buildFlags.strip {
        args = append(args, "-s")
      }
    } else {
      fmt.Println("🍕 Warning: Could not detect MSVC or MinGW toolchain. Defaulting to MinGW-style flags.\nIf you see linker errors, install MinGW and make sure gcc.exe is in your PATH! 🦄")
      args = append(objectFiles, "-o", output)
      args = append(args, "-lkernel32", "-lmsvcrt", "-lucrt", "-loldnames")
    }
  } else {
    args = append(objectFiles, "-o", output)
    if buildFlags.fuseLd != "" {
      args = append([]string{"-fuse-ld=" + buildFlags.fuseLd}, args...)
    }
    addTargetLibraries(&args)
    // Always link libc for mold/unix linkers
    args = append(args, "-lc")
    if buildFlags.optimization != "0" {
      args = append(args, "-O"+buildFlags.optimization)
    }
    if buildFlags.debugSymbols && !buildFlags.strip {
      args = append(args, "-g")
    }
    if buildFlags.strip {
      args = append(args, "-s")
    }
  }
  // Ensure output directory exists
  outDir := filepath.Dir(output)
  if outDir != "." && outDir != "" {
    _ = os.MkdirAll(outDir, 0755)
  }
  cmd := exec.Command(buildFlags.linker, args...)
  if buildFlags.verbose {
    fmt.Println("[Aether-DEBUG] Linker command:", buildFlags.linker, strings.Join(args, " "))
  }
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  must(cmd.Run())
}

func addTargetLibraries(args *[]string) {
  libcType := buildFlags.libcType
  if libcType == "" {
    libcType = "glibc"
  }
  switch buildFlags.targetOS {
  case "linux":
    switch buildFlags.targetArch {
    case "amd64":
      if libcType == "musl" {
        *args = append(*args,
          "-L/usr/lib/musl",
          "/usr/lib/musl/crt1.o",
          "/usr/lib/musl/crti.o",
          "-lc",
          "/usr/lib/musl/crtn.o",
        )
      } else {
        *args = append(*args,
          "-L/usr/lib/x86_64-linux-gnu",
          "-L/usr/lib",
          "/usr/lib/x86_64-linux-gnu/crt1.o",
          "/usr/lib/x86_64-linux-gnu/crti.o",
          "-lc",
          "/usr/lib/x86_64-linux-gnu/crtn.o",
        )
      }
    case "arm64":
      if libcType == "musl" {
        *args = append(*args,
          "-L/usr/lib/musl",
          "/usr/lib/musl/crt1.o",
          "/usr/lib/musl/crti.o",
          "-lc",
          "/usr/lib/musl/crtn.o",
        )
      } else {
        *args = append(*args,
          "-L/usr/lib/aarch64-linux-gnu",
          "-L/usr/lib",
          "/usr/lib/aarch64-linux-gnu/crt1.o",
          "/usr/lib/aarch64-linux-gnu/crti.o",
          "-lc",
          "/usr/lib/aarch64-linux-gnu/crtn.o",
        )
      }
    default:
      *args = append(*args, "-lc")
    }
  case "darwin":
    *args = append(*args, "-L/usr/lib", "-lc")
  case "windows":
    *args = append(*args, "-lkernel32", "-lmsvcrt", "-lucrt", "-loldnames")
  }
}

func createLibrary(objectFiles []string, outputBase string, libType string) {
  switch libType {
  case "shared":
    createSharedLibrary(objectFiles, outputBase)
  case "static":
    createStaticLibrary(objectFiles, outputBase)
  case "both":
    createSharedLibrary(objectFiles, outputBase)
    createStaticLibrary(objectFiles, outputBase)
  default:
    fmt.Printf("Error: Unknown library type '%s'\n", libType)
    os.Exit(1)
  }
}

func generatePkgConfigFile(libName string, outputBase string) {
  if !buildFlags.generatePkgConfig {
    return
  }
  pcContent := generatePkgConfigContent(libName)
  pcFile := getPkgConfigPath(libName)
  must(os.WriteFile(pcFile, []byte(pcContent), 0644))
  if !buildFlags.quiet {
    fmt.Printf("Generated pkg-config file: %s\n", pcFile)
  }
}

func getPkgConfigPath(libName string) string {
  return fmt.Sprintf("%s.pc", libName)
} 