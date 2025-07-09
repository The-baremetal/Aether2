package compiler

import (
	"aether/src/parser"

	"tinygo.org/x/go-llvm"
)

func compileStmt(stmt parser.Statement, ctx *CompilerContext) {
	switch s := stmt.(type) {
	case *parser.Assignment:
		value := compileExpr(s.Value, ctx)
		alloca := ctx.builder.CreateAlloca(value.Type(), s.Name.Value)
		ctx.builder.CreateStore(value, alloca)
		ctx.SetSymbol(s.Name.Value, alloca)
	case *parser.Function:
		fnType := llvm.FunctionType(ctx.llvm_context.Int32Type(), nil, false)
		fn := llvm.AddFunction(ctx.module, s.Name.Value, fnType)
		ctx.SetSymbol(s.Name.Value, fn)
		block := ctx.llvm_context.AddBasicBlock(fn, "entry")
		ctx.builder.SetInsertPointAtEnd(block)
		for _, stmt := range s.Body.Statements {
			compileStmt(stmt, ctx)
		}
		ctx.builder.CreateRet(llvm.ConstInt(ctx.llvm_context.Int32Type(), 0, false))
	case *parser.StructDef:
		// Structs are not codegen'd directly in LLVM IR here
	case *parser.If:
		cond := compileExpr(s.Condition, ctx)
		parent := ctx.builder.GetInsertBlock().Parent()
		thenBlock := ctx.llvm_context.AddBasicBlock(parent, "then")
		elseBlock := ctx.llvm_context.AddBasicBlock(parent, "else")
		mergeBlock := ctx.llvm_context.AddBasicBlock(parent, "merge")
		ctx.builder.CreateCondBr(cond, thenBlock, elseBlock)
		ctx.builder.SetInsertPointAtEnd(thenBlock)
		for _, stmt := range s.Consequence.Statements {
			compileStmt(stmt, ctx)
		}
		ctx.builder.CreateBr(mergeBlock)
		ctx.builder.SetInsertPointAtEnd(elseBlock)
		if s.Alternative != nil {
			for _, stmt := range s.Alternative.Statements {
				compileStmt(stmt, ctx)
			}
		}
		ctx.builder.CreateBr(mergeBlock)
		ctx.builder.SetInsertPointAtEnd(mergeBlock)
	case *parser.While:
		parent := ctx.builder.GetInsertBlock().Parent()
		condBlock := ctx.llvm_context.AddBasicBlock(parent, "while.cond")
		bodyBlock := ctx.llvm_context.AddBasicBlock(parent, "while.body")
		endBlock := ctx.llvm_context.AddBasicBlock(parent, "while.end")
		ctx.builder.CreateBr(condBlock)
		ctx.builder.SetInsertPointAtEnd(condBlock)
		cond := compileExpr(s.Condition, ctx)
		ctx.builder.CreateCondBr(cond, bodyBlock, endBlock)
		ctx.builder.SetInsertPointAtEnd(bodyBlock)
		for _, stmt := range s.Body.Statements {
			compileStmt(stmt, ctx)
		}
		ctx.builder.CreateBr(condBlock)
		ctx.builder.SetInsertPointAtEnd(endBlock)
	case *parser.Repeat:
		// Not implemented: repeat
	case *parser.For:
		// Not implemented: for
	case *parser.Block:
		for _, stmt := range s.Statements {
			compileStmt(stmt, ctx)
		}
	case *parser.Return:
		if s.Value != nil {
			value := compileExpr(s.Value, ctx)
			ctx.builder.CreateRet(value)
		} else {
			ctx.builder.CreateRet(llvm.ConstInt(ctx.llvm_context.Int32Type(), 0, false))
		}
	case *parser.Import:
		// Imports handled in analysis
	}
}
