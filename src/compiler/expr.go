package compiler

import (
	"aether/src/parser"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func compileExpr(expr parser.Expression, ctx *CompilerContext) value.Value {
	switch e := expr.(type) {
	case *parser.Identifier:
		val, ok := ctx.GetSymbol(e.Value)
		if ok {
			return ctx.builder.NewLoad(val.Type(), val)
		}
		return nil
	case *parser.Literal:
		switch v := e.Value.(type) {
		case int:
			return constant.NewInt(types.I32, int64(v))
		case float64:
			return constant.NewFloat(types.Double, v)
		case string:
			return constant.NewCharArrayFromString(v)
		case bool:
			if v {
				return constant.NewInt(types.I1, 1)
			}
			return constant.NewInt(types.I1, 0)
		}
		return nil
	case *parser.Array:
		elems := make([]constant.Constant, len(e.Elements))
		for i, el := range e.Elements {
			val := compileExpr(el, ctx)
			if constVal, ok := val.(constant.Constant); ok {
				elems[i] = constVal
			} else {
				// If not a constant, we need to create a global variable
				// For now, just use a default value
				elems[i] = constant.NewInt(types.I32, 0)
			}
		}
		arrType := types.NewArray(uint64(len(elems)), types.I32)
		return constant.NewArray(arrType, elems...)
	case *parser.Call:
		if ident, ok := e.Function.(*parser.Identifier); ok {
			if isStdlibFunction(ident.Value) {
				return compileStdlibCall(ident.Value, e.Args, ctx)
			}
		}
		fn := compileExpr(e.Function, ctx)
		args := make([]value.Value, len(e.Args))
		for i, arg := range e.Args {
			args[i] = compileExpr(arg, ctx)
		}
		return ctx.builder.NewCall(fn, args...)
	case *parser.PropertyAccess:
		obj := compileExpr(e.Object, ctx)
		if obj == nil {
			// Try to handle module.property access
			if moduleIdent, ok := e.Object.(*parser.Identifier); ok {
				// Use proper module resolution
				if symbol, exists := ctx.GetModuleSymbol(moduleIdent.Value, e.Property.Value); exists {
					return symbol
				}
			}
		}
		return ctx.builder.NewExtractValue(obj, 0)
	case *parser.PartialApplication:
		fn := compileExpr(e.Function, ctx)
		args := make([]value.Value, len(e.Args))
		for i, arg := range e.Args {
			args[i] = compileExpr(arg, ctx)
		}
		return ctx.builder.NewCall(fn, args...)
	}
	return nil
}

func compileStdlibCall(funcName string, args []parser.Expression, ctx *CompilerContext) value.Value {
	// Try to resolve stdlib functions from modules
	if symbol, exists := ctx.GetModuleSymbol("print", funcName); exists {
		// This is a stdlib function call
		compiledArgs := make([]value.Value, len(args))
		for i, arg := range args {
			compiledArgs[i] = compileExpr(arg, ctx)
		}
		return ctx.builder.NewCall(symbol, compiledArgs...)
	}

	// Fallback to printf for print function
	if funcName == "print" {
		if len(args) == 0 {
			return constant.NewInt(types.I32, 0)
		}
		arg := compileExpr(args[0], ctx)
		if str, ok := arg.(*constant.CharArray); ok {
			printfFunc := getOrCreatePrintfFunction(ctx)
			return ctx.builder.NewCall(printfFunc, str)
		}
		return constant.NewInt(types.I32, 0)
	}

	return constant.NewInt(types.I32, 0)
}

func getOrCreatePrintfFunction(ctx *CompilerContext) *ir.Func {
	for _, fn := range ctx.module.Funcs {
		if fn.Name() == "printf" {
			return fn
		}
	}
	printfFunc := ctx.module.NewFunc("printf", types.I32, ir.NewParam("format", types.NewPointer(types.I8)))
	printfFunc.Sig.Variadic = true
	return printfFunc
}
