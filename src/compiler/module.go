package compiler

import (
	"strings"

	"github.com/llir/llvm/ir"
)

type Module struct {
	irModule          *ir.Module
	optimizationLevel string
}

func NewModule(moduleName string) *Module {
	module := ir.NewModule()
	return &Module{
		irModule:          module,
		optimizationLevel: "default<O2>",
	}
}

func NewModuleWithOptimization(moduleName string, optimizationLevel string) *Module {
	module := ir.NewModule()
	return &Module{
		irModule:          module,
		optimizationLevel: optimizationLevel,
	}
}

func (m *Module) IR() *ir.Module {
	return m.irModule
}

func (m *Module) SetOptimizationLevel(level string) {
	m.optimizationLevel = level
}

func (m *Module) GetOptimizationLevel() string {
	return m.optimizationLevel
}

func (m *Module) ApplyOptimizations() {
	if strings.Contains(m.optimizationLevel, "O0") {
		return
	}

	// Note: llir/llvm doesn't have built-in optimization passes like the C API
	// You would need to use a separate tool like 'opt' to apply optimizations
	// For now, we'll just return without doing anything
}

func (m *Module) String() string {
	return m.irModule.String()
}

func (m *Module) Dispose() {
	// llir/llvm handles memory management automatically
}
