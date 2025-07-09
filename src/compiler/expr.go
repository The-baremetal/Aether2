package compiler

import (
	"aether/src/parser"

	"tinygo.org/x/go-llvm"
)

func compileExpr(expr parser.Expression, ctx *CompilerContext) llvm.Value {
	switch e := expr.(type) {
	case *parser.Identifier:
		val, ok := ctx.GetSymbol(e.Value)
		if ok {
			return ctx.builder.CreateLoad(val.Type().ElementType(), val, e.Value)
		}
		return llvm.Value{}
	case *parser.Literal:
		switch v := e.Value.(type) {
		case int:
			return llvm.ConstInt(ctx.llvm_context.Int32Type(), uint64(v), false)
		case float64:
			return llvm.ConstFloat(ctx.llvm_context.DoubleType(), v)
		case string:
			return llvm.ConstString(v, false)
		case bool:
			if v {
				return llvm.ConstInt(ctx.llvm_context.Int1Type(), 1, false)
			}
			return llvm.ConstInt(ctx.llvm_context.Int1Type(), 0, false)
		}
		return llvm.Value{}
	case *parser.Array:
		elems := make([]llvm.Value, len(e.Elements))
		for i, el := range e.Elements {
			elems[i] = compileExpr(el, ctx)
		}
		arrType := llvm.ArrayType(ctx.llvm_context.Int32Type(), len(elems))
		arr := llvm.Undef(arrType)
		for i, v := range elems {
			arr = ctx.builder.CreateInsertValue(arr, v, i, "")
		}
		return arr
	case *parser.Call:
		fn := compileExpr(e.Function, ctx)
		args := make([]llvm.Value, len(e.Args))
		for i, arg := range e.Args {
			args[i] = compileExpr(arg, ctx)
		}
		return ctx.builder.CreateCall(fn.Type(), fn, args, "call")
	case *parser.PropertyAccess:
		obj := compileExpr(e.Object, ctx)
		return ctx.builder.CreateExtractValue(obj, 0, "prop")
	case *parser.PartialApplication:
		fn := compileExpr(e.Function, ctx)
		args := make([]llvm.Value, len(e.Args))
		for i, arg := range e.Args {
			args[i] = compileExpr(arg, ctx)
		}
		return ctx.builder.CreateCall(fn.Type(), fn, args, "partial")
	}
	return llvm.Value{}
}
