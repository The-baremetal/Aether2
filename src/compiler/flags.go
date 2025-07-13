package compiler

import (
	"fmt"
	"strings"
)

type CompilerFlags struct {
	Global  []string `toml:"global,omitempty"`
	Debug   []string `toml:"debug,omitempty"`
	Release []string `toml:"release,omitempty"`
}

type TargetConfig struct {
	CompilerFlags []string `toml:"compiler_flags,omitempty"`
}

type FlagMerger struct {
	configFlags []string
	cliFlags    []string
	targetOS    string
	optimization string
	debugInfo   bool
}

func NewFlagMerger() *FlagMerger {
	return &FlagMerger{
		configFlags: []string{},
		cliFlags:    []string{},
	}
}

func (fm *FlagMerger) SetConfigFlags(flags CompilerFlags, targets map[string]TargetConfig) {
	fm.configFlags = fm.mergeConfigFlags(flags, targets)
}

func (fm *FlagMerger) SetCLIFlags(flags []string) {
	fm.cliFlags = flags
}

func (fm *FlagMerger) SetTargetOS(os string) {
	fm.targetOS = os
}

func (fm *FlagMerger) SetOptimization(opt string) {
	fm.optimization = opt
}

func (fm *FlagMerger) SetDebugInfo(debug bool) {
	fm.debugInfo = debug
}

func (fm *FlagMerger) mergeConfigFlags(flags CompilerFlags, targets map[string]TargetConfig) []string {
	var merged []string
	
	if len(flags.Global) > 0 {
		merged = append(merged, flags.Global...)
	}
	
	if fm.optimization == "0" || fm.debugInfo {
		if len(flags.Debug) > 0 {
			merged = append(merged, flags.Debug...)
		}
	} else if fm.optimization == "3" {
		if len(flags.Release) > 0 {
			merged = append(merged, flags.Release...)
		}
	}
	
	targetKey := fmt.Sprintf("cfg(target_os = \"%s\")", fm.targetOS)
	if targetConfig, exists := targets[targetKey]; exists {
		if len(targetConfig.CompilerFlags) > 0 {
			merged = append(merged, targetConfig.CompilerFlags...)
		}
	}
	
	return merged
}

func (fm *FlagMerger) GetMergedFlags() []string {
	var result []string
	
	result = append(result, fm.configFlags...)
	result = append(result, fm.cliFlags...)
	
	return result
}

func (fm *FlagMerger) ValidateFlags() error {
	for _, flag := range fm.configFlags {
		if err := fm.validateFlag(flag); err != nil {
			return fmt.Errorf("invalid config flag '%s': %v", flag, err)
		}
	}
	
	for _, flag := range fm.cliFlags {
		if err := fm.validateFlag(flag); err != nil {
			return fmt.Errorf("invalid CLI flag '%s': %v", flag, err)
		}
	}
	
	return nil
}

func (fm *FlagMerger) validateFlag(flag string) error {
	if strings.HasPrefix(flag, "-") {
		return nil
	}
	
	if strings.HasPrefix(flag, "--") {
		return nil
	}
	
	if strings.HasPrefix(flag, "-fuse-ld=") {
		return nil
	}
	
	if strings.HasPrefix(flag, "-O") {
		return nil
	}
	
	if strings.HasPrefix(flag, "-g") {
		return nil
	}
	
	if strings.HasPrefix(flag, "-D") {
		return nil
	}
	
	if strings.HasPrefix(flag, "-I") {
		return nil
	}
	
	if strings.HasPrefix(flag, "-L") {
		return nil
	}
	
	if strings.HasPrefix(flag, "-l") {
		return nil
	}
	
	return fmt.Errorf("unknown flag format: %s", flag)
}

func (fm *FlagMerger) GetFlagSummary() string {
	var summary strings.Builder
	
	summary.WriteString("Compiler Flags Summary:\n")
	
	if len(fm.configFlags) > 0 {
		summary.WriteString("  Config flags: ")
		summary.WriteString(strings.Join(fm.configFlags, " "))
		summary.WriteString("\n")
	}
	
	if len(fm.cliFlags) > 0 {
		summary.WriteString("  CLI flags: ")
		summary.WriteString(strings.Join(fm.cliFlags, " "))
		summary.WriteString("\n")
	}
	
	summary.WriteString("  Target OS: ")
	summary.WriteString(fm.targetOS)
	summary.WriteString("\n")
	
	summary.WriteString("  Optimization: ")
	summary.WriteString(fm.optimization)
	summary.WriteString("\n")
	
	return summary.String()
} 