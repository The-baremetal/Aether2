package compiler

func createMainReturn(ctx *CompilerContext) {
	ctx.builder.CreateRet(ctx.llvm_context.ConstInt(ctx.llvm_context.Int32Type(), 0, false))
}
