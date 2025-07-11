package analysis

import "aether/lib/utils"

type AnalysisResult struct {
	Valid        bool
	Errors       []string
	Warnings     []string
	Imports      map[string]ImportInfo
	Functions    map[string]FunctionInfo
	Variables    map[string]VariableInfo
	Types        map[string]TypeInfo
	Constants    map[string]ConstantInfo
	Dependencies map[string][]string
	Cycles       [][]string
	Unused       []string
	Undefined    []string
	CIncludes    []CInclude
}

type ImportInfo struct {
	Path     string
	Valid    bool
	Exists   bool
	Resolved string
	Errors   []string
}

type FunctionInfo struct {
	Name       string
	Parameters []ParameterInfo
	ReturnType string
	Defined    bool
	Used       bool
	Exported   bool
}

type VariableInfo struct {
	Name     string
	Type     string
	Defined  bool
	Used     bool
	Exported bool
	Scope    string
}

type TypeInfo struct {
	Name     string
	Defined  bool
	Used     bool
	Exported bool
	Fields   map[string]string
}

type ConstantInfo struct {
	Name     string
	Type     string
	Value    interface{}
	Defined  bool
	Used     bool
	Exported bool
}

type ParameterInfo struct {
	Name string
	Type string
}

type DependencyInfo struct {
	Path    string `toml:"path"`
	Version string `toml:"version,omitempty"`
}

type LockFile struct {
	Dependencies map[string]DependencyInfo `toml:"dependencies"`
}

type DependencyAnalysis struct {
	Valid        bool
	Errors       []utils.ParseError
	Warnings     []string
	ResolvedDeps map[string]string
	MissingDeps  []string
	UnusedDeps   []string
}

type CInclude struct {
	Header   string
	IsSystem bool
}
