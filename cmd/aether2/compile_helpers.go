package main

import (
  "os"
  "os/exec"
  "fmt"
  "strings"
)

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
  llFile := strings.TrimSuffix(outputFile, ".s") + ".ll"
  must(os.WriteFile(llFile, []byte(ir), 0644))
  if buildFlags.verbose {
    fmt.Printf("Generating assembly: %s\n", outputFile)
  }
  cmd := exec.Command("llc", "-filetype=asm", llFile, "-o", outputFile)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  must(cmd.Run())
}

func generateBitcode(ir string, outputFile string) {
  llFile := strings.TrimSuffix(outputFile, ".bc") + ".ll"
  must(os.WriteFile(llFile, []byte(ir), 0644))
  cmd := exec.Command("llvm-as", llFile, "-o", outputFile)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  must(cmd.Run())
}

func generateObjectFile(ir string, outputFile string) {
  llFile := strings.TrimSuffix(outputFile, ".o") + ".ll"
  must(os.WriteFile(llFile, []byte(ir), 0644))
  cmd := exec.Command("llc", "-filetype=obj", llFile, "-o", outputFile)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  must(cmd.Run())
} 