package compiler

import (
	"tinygo.org/x/go-llvm"
)

type CompilerContext struct {
	llvm_context llvm.Context
	module       llvm.Module
	builder      llvm.Builder
	symbol_table map[string]llvm.Value
	scopes       []map[string]llvm.Value
	current_func llvm.Value
	optimizer    llvm.PassManager
}

func NewCompilerContext(module_name string) *CompilerContext {
	ctx := llvm.NewContext()
	module := ctx.NewModule(module_name)
	builder := ctx.NewBuilder()

	return &CompilerContext{
		llvm_context: ctx,
		module:       module,
		builder:      builder,
		symbol_table: make(map[string]llvm.Value),
		scopes:       []map[string]llvm.Value{make(map[string]llvm.Value)},
		current_func: llvm.Value{},
		optimizer:    llvm.NewPassManager(),
	}
}

func (c *CompilerContext) EnterScope() {
	c.scopes = append(c.scopes, make(map[string]llvm.Value))
}

func (c *CompilerContext) ExitScope() {
	if len(c.scopes) > 1 {
		c.scopes = c.scopes[:len(c.scopes)-1]
	}
}

func (c *CompilerContext) SetSymbol(name string, value llvm.Value) {
	current_scope := c.scopes[len(c.scopes)-1]
	current_scope[name] = value
}

func (c *CompilerContext) GetSymbol(name string) (llvm.Value, bool) {
	for i := len(c.scopes) - 1; i >= 0; i-- {
		if value, exists := c.scopes[i][name]; exists {
			return value, true
		}
	}
	return llvm.Value{}, false
}

func (c *CompilerContext) SetCurrentFunction(func_val llvm.Value) {
	c.current_func = func_val
}

func (c *CompilerContext) GetCurrentFunction() llvm.Value {
	return c.current_func
}

func (c *CompilerContext) GetModule() llvm.Module {
	return c.module
}

func (c *CompilerContext) GetBuilder() llvm.Builder {
	return c.builder
}

func (c *CompilerContext) GetContext() llvm.Context {
	return c.llvm_context
}

func (c *CompilerContext) Dispose() {
	c.llvm_context.Dispose()
	c.builder.Dispose()
	c.optimizer.Dispose()
}
