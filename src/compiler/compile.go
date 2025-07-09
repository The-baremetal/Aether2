package compiler

import (
	"aether/src/parser"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func Compile(prog *parser.Program) string {
	return CompileWithOptions(prog, "main")
}

func CompileWithOptions(prog *parser.Program, moduleName string) string {
	return CompileWithOptionsAndModules(prog, moduleName, nil)
}

func CompileWithOptionsAndModules(prog *parser.Program, moduleName string, moduleSymbols map[string]map[string]interface{}) string {
	ctx := NewCompilerContext(moduleName)
	defer ctx.Dispose()

	ast := parser.ProgramToAST(prog)
	analysisResult := AnalyzeAST(ast)

	for _, include := range analysisResult.CIncludes {
		ctx.AddLibrary(include.Header)
	}

	// Set up module symbols
	if moduleSymbols != nil {
		for moduleName, symbols := range moduleSymbols {
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

	if moduleName == "main" {
		mainFn := createMainFunction(ctx)
		entry := addEntryBlock(ctx, mainFn)
		setInsertPoint(ctx, entry)

		for _, stmt := range prog.Statements {
			compileStmt(stmt, ctx)
		}

		createMainReturn(ctx)
	} else {
		dummyFn := ctx.module.NewFunc("__module_"+moduleName, types.I32)
		entry := dummyFn.NewBlock("entry")
		ctx.builder = entry

		for _, stmt := range prog.Statements {
			compileStmt(stmt, ctx)
		}

		entry.NewRet(constant.NewInt(types.I32, 0))
	}

	return ctx.GetModule().String()
}
