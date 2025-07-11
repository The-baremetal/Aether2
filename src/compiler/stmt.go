package compiler

import (
	"aether/src/parser"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func compileStmt(stmt parser.Statement, ctx *CompilerContext) {
	switch s := stmt.(type) {
	case *parser.Assignment:
		val := compileExpr(s.Value, ctx)
		// For now, only support single assignment for codegen
		if len(s.Names) > 0 {
			alloca := ctx.builder.NewAlloca(val.Type())
			alloca.SetName(s.Names[0].Value)
			ctx.builder.NewStore(val, alloca)
			ctx.SetSymbol(s.Names[0].Value, alloca)
		}
	case *parser.Function:
		fn := ctx.module.NewFunc(s.Name.Value, types.I32)
		ctx.SetSymbol(s.Name.Value, fn)
		block := fn.NewBlock("entry")
		ctx.builder = block
		for _, stmt := range s.Body.Statements {
			compileStmt(stmt, ctx)
		}
		ctx.builder.NewRet(constant.NewInt(types.I32, 0))
	case *parser.StructDef:
		// Structs are not codegen'd directly in LLVM IR here
	case *parser.If:
		cond := compileExpr(s.Condition, ctx)
		parent := ctx.current_func
		thenBlock := parent.NewBlock("then")
		elseBlock := parent.NewBlock("else")
		mergeBlock := parent.NewBlock("merge")
		ctx.builder.NewCondBr(cond, thenBlock, elseBlock)
		ctx.builder = thenBlock
		for _, stmt := range s.Consequence.Statements {
			compileStmt(stmt, ctx)
		}
		ctx.builder.NewBr(mergeBlock)
		ctx.builder = elseBlock
		if s.Alternative != nil {
			for _, stmt := range s.Alternative.Statements {
				compileStmt(stmt, ctx)
			}
		}
		ctx.builder.NewBr(mergeBlock)
		ctx.builder = mergeBlock
	case *parser.While:
		parent := ctx.current_func
		condBlock := parent.NewBlock("while.cond")
		bodyBlock := parent.NewBlock("while.body")
		endBlock := parent.NewBlock("while.end")
		ctx.builder.NewBr(condBlock)
		ctx.builder = condBlock
		cond := compileExpr(s.Condition, ctx)
		ctx.builder.NewCondBr(cond, bodyBlock, endBlock)
		ctx.builder = bodyBlock
		for _, stmt := range s.Body.Statements {
			compileStmt(stmt, ctx)
		}
		ctx.builder.NewBr(condBlock)
		ctx.builder = endBlock
	case *parser.Repeat:
		// Not implemented: repeat
	case *parser.For:
		// Not implemented: for
	case *parser.Block:
		for _, stmt := range s.Statements {
			compileStmt(stmt, ctx)
		}
	case *parser.Match:
		// TODO: Proper match/case codegen. For now, just compile all case bodies.
		for _, c := range s.Cases {
			compileStmt(c.Body, ctx)
		}
	case *parser.Break:
		// TODO: Proper break codegen. For now, do nothing.
	case *parser.Continue:
		// TODO: Proper continue codegen. For now, do nothing.
	case *parser.Return:
		if s.Value != nil {
			val := compileExpr(s.Value, ctx)
			ctx.builder.NewRet(val)
		} else {
			ctx.builder.NewRet(constant.NewInt(types.I32, 0))
		}
	case *parser.Import:
		// Handle imports by making symbols available
		moduleName := s.Name.Value
		if s.As != nil && s.As.Value != "" {
			moduleName = s.As.Value
		}
		// Create a dummy value for the module symbol
		dummyValue := constant.NewInt(types.I32, 0)
		ctx.SetSymbol(moduleName, dummyValue)
	}
}
