package main

import (
  "os"
  "path/filepath"
  "fmt"
)

func findProjectRoot(start string) string {
  dir := start
  for {
    configPath := filepath.Join(dir, "aether.toml")
    if _, err := os.Stat(configPath); err == nil {
      if buildFlags.verbose {
        fmt.Printf("Found project root at: %s\n", dir)
      }
      return dir
    }
    parent := filepath.Dir(dir)
    if parent == dir {
      break
    }
    dir = parent
  }
  if buildFlags.verbose {
    fmt.Println("No project root found, using .")
  }
  return "."
}

func getMapKeys(m interface{}) []string {
  var keys []string
  switch mm := m.(type) {
  case map[string]func():
    for k := range mm {
      keys = append(keys, k)
    }
  case map[string][]string:
    for k := range mm {
      keys = append(keys, k)
    }
  }
  return keys
} 