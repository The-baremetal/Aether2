package compiler

import (
	analysis "aether/src/analysis"
	"aether/src/parser"

	"fmt"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
	"github.com/llir/llvm/ir"
)

func hasFunc(mod *ir.Module, name string) bool {
	for _, fn := range mod.Funcs {
		if fn.Name() == name {
			return true
		}
	}
	return false
}

func Compile(prog *parser.Program) string {
	return CompileWithOptions(prog, "main")
}

func CompileWithOptions(prog *parser.Program, moduleName string) string {
	return CompileWithOptionsAndModules(prog, moduleName, nil, true, "main")
}

func CompileWithOptionsAndModules(prog *parser.Program, moduleName string, moduleSymbols map[string]map[string]interface{}, verbose bool, funcName string) string {
	ctx := NewCompilerContext(moduleName)
	ctx.verbose = verbose
	ctx.DebugPrint("Starting compilation for module: " + moduleName)
	defer ctx.Dispose()

	// After parsing and before analysis, print all top-level statement types
	ctx.DebugPrint("Top-level statements in parsed program:")
	for _, stmt := range prog.Statements {
		switch fn := stmt.(type) {
		case *parser.Function:
			ctx.DebugPrint("  Function: " + fn.Name.Value)
			if fn.Name.Value == "main" {
				ctx.DebugPrint("  >>> Found main function!")
			}
		case *parser.StructDef:
			ctx.DebugPrint("  StructDef: " + fn.Name.Value)
		case *parser.Import:
			ctx.DebugPrint("  Import: " + fn.Name.Value)
		case *parser.Package:
			ctx.DebugPrint("  Package: " + fn.Name.Value)
		default:
			ctx.DebugPrint("  Other statement: " + fmt.Sprintf("%T", stmt))
		}
	}

	ast := parser.ProgramToAST(prog)
	ctx.DebugPrint("AST generated")
	analysisResult := analysis.AnalyzeAST(ast)
	ctx.DebugPrint("AST analysis complete")

	for _, include := range analysisResult.CIncludes {
		ctx.DebugPrint("Adding C include: " + include.Header)
		ctx.AddLibrary(include.Header)
	}
	if moduleSymbols != nil {
		for moduleName, symbols := range moduleSymbols {
			ctx.DebugPrint("Setting module symbols for: " + moduleName)
			moduleInfo := &ModuleInfo{
				Name:    moduleName,
				Symbols: make(map[string]value.Value),
			}
			for symbolName, symbolValue := range symbols {
				if str, ok := symbolValue.(string); ok {
					moduleInfo.Symbols[symbolName] = constant.NewCharArrayFromString(str)
				}
			}
			ctx.SetModule(moduleName, moduleInfo)
		}
	}
	ctx.DebugPrint("Compilation finished for module: " + moduleName)

	// After analysis, emit all top-level statements, passing analysisResult to compileStmt
	for _, stmt := range prog.Statements {
		compileStmt(stmt, ctx, analysisResult)
	}
	return ctx.GetModule().String()
}
