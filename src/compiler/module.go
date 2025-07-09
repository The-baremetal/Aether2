package compiler

import (
	"strings"

	"tinygo.org/x/go-llvm"
)

type Module struct {
	irModule          llvm.Module
	context           llvm.Context
	optimizationLevel string
}

func NewModule() *Module {
	context := llvm.NewContext()
	module := context.NewModule("aether_module")
	return &Module{
		irModule:          module,
		context:           context,
		optimizationLevel: "default<O2>",
	}
}

func NewModuleWithOptimization(optimizationLevel string) *Module {
	context := llvm.NewContext()
	module := context.NewModule("aether_module")
	return &Module{
		irModule:          module,
		context:           context,
		optimizationLevel: optimizationLevel,
	}
}

func (m *Module) IR() llvm.Module {
	return m.irModule
}

func (m *Module) Context() llvm.Context {
	return m.context
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

	passManager := llvm.NewPassManager()
	defer passManager.Dispose()

	passManager.AddInstructionCombiningPass()
	passManager.AddReassociatePass()
	passManager.AddGVNPass()
	passManager.AddCFGSimplificationPass()

	passManager.Run(m.irModule)
}

func (m *Module) String() string {
	return m.irModule.String()
}

func (m *Module) Dispose() {
	m.irModule.Dispose()
	m.context.Dispose()
}
