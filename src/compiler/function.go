package compiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

func createMainFunction(ctx *CompilerContext) *ir.Func {
	mainFn := ctx.module.NewFunc("main", types.I32)
	ctx.SetCurrentFunction(mainFn)
	return mainFn
}

func addEntryBlock(ctx *CompilerContext, fn *ir.Func) *ir.Block {
	return fn.NewBlock("entry")
}

func setInsertPoint(ctx *CompilerContext, block *ir.Block) {
	ctx.builder = block
}
