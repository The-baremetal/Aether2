package compiler

import (
	"tinygo.org/x/go-llvm"
)

func createMainFunction(ctx *CompilerContext) llvm.Value {
	mainFnType := llvm.FunctionType(ctx.llvm_context.Int32Type(), []llvm.Type{}, false)
	mainFn := llvm.AddFunction(ctx.module, "main", mainFnType)
	ctx.SetCurrentFunction(mainFn)
	return mainFn
}

func addEntryBlock(ctx *CompilerContext, fn llvm.Value) llvm.BasicBlock {
	return ctx.llvm_context.AddBasicBlock(fn, "entry")
}

func setInsertPoint(ctx *CompilerContext, block llvm.BasicBlock) {
	ctx.builder.SetInsertPointAtEnd(block)
}
