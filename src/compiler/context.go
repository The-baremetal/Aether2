package compiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

type CompilerContext struct {
	module       *ir.Module
	builder      *ir.Block
	symbol_table map[string]value.Value
	scopes       []map[string]value.Value
	current_func *ir.Func
	modules      map[string]*ModuleInfo
	libraries    []string
}

type ModuleInfo struct {
	Name    string
	Symbols map[string]value.Value
}

func NewCompilerContext(module_name string) *CompilerContext {
	module := ir.NewModule()

	return &CompilerContext{
		module:       module,
		builder:      nil,
		symbol_table: make(map[string]value.Value),
		scopes:       []map[string]value.Value{make(map[string]value.Value)},
		current_func: nil,
		modules:      make(map[string]*ModuleInfo),
		libraries:    []string{},
	}
}

func (c *CompilerContext) EnterScope() {
	c.scopes = append(c.scopes, make(map[string]value.Value))
}

func (c *CompilerContext) ExitScope() {
	if len(c.scopes) > 1 {
		c.scopes = c.scopes[:len(c.scopes)-1]
	}
}

func (c *CompilerContext) SetSymbol(name string, val value.Value) {
	current_scope := c.scopes[len(c.scopes)-1]
	current_scope[name] = val
}

func (c *CompilerContext) GetSymbol(name string) (value.Value, bool) {
	for i := len(c.scopes) - 1; i >= 0; i-- {
		if val, exists := c.scopes[i][name]; exists {
			return val, true
		}
	}
	return nil, false
}

func (c *CompilerContext) SetModule(moduleName string, moduleInfo *ModuleInfo) {
	c.modules[moduleName] = moduleInfo
}

func (c *CompilerContext) GetModuleSymbol(moduleName, symbolName string) (value.Value, bool) {
	if module, exists := c.modules[moduleName]; exists {
		if symbol, exists := module.Symbols[symbolName]; exists {
			return symbol, true
		}
	}
	return nil, false
}

func (c *CompilerContext) SetCurrentFunction(func_val *ir.Func) {
	c.current_func = func_val
}

func (c *CompilerContext) GetCurrentFunction() *ir.Func {
	return c.current_func
}

func (c *CompilerContext) GetModule() *ir.Module {
	return c.module
}

func (c *CompilerContext) GetBuilder() *ir.Block {
	return c.builder
}

func (c *CompilerContext) AddLibrary(library string) {
	for _, lib := range c.libraries {
		if lib == library {
			return
		}
	}
	c.libraries = append(c.libraries, library)
}

func (c *CompilerContext) GetLibraries() []string {
	return c.libraries
}

func (c *CompilerContext) Dispose() {
	// llir/llvm handles memory management automatically
}
