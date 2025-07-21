package main

import (
  "os"
  "fmt"
  "path/filepath"
  "strings"
  "aether/src/parser"
  "aether/src/buildcache"
)

// isStale closure must be set up in build_cmd.go, but the function body is here for modularity
func makeIsStale(cache map[string]buildcache.BuildCacheEntry, projectRoot string, projectConfig ProjectConfig, resolvedImports map[string][]string, stale map[string]bool, reasons map[string]string) func(file string) bool {
  return func(file string) bool {
    entry, ok := cache[file]
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
    for _, dep := range resolvedImports[file] {
      if makeIsStale(cache, projectRoot, projectConfig, resolvedImports, stale, reasons)(dep) {
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
}

func extractModuleSymbols(prog *parser.Program) map[string]interface{} {
  symbols := make(map[string]interface{})
  for _, stmt := range prog.Statements {
    if assign, ok := stmt.(*parser.Assignment); ok {
      if len(assign.Names) > 0 && assign.Names[0].Value[0] >= 'A' && assign.Names[0].Value[0] <= 'Z' {
        symbols[assign.Names[0].Value] = assign.Value
      }
    }
  }
  return symbols
} 