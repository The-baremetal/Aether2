package compiler

import (
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func createMainReturn(ctx *CompilerContext) {
	if ctx.builder != nil {
		ctx.builder.NewRet(constant.NewInt(types.I32, 0))
	}
}
