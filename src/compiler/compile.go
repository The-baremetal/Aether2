package compiler

import (
	"aether/src/parser"
)

func Compile(prog *parser.Program) string {
	return CompileWithOptions(prog, "main")
}

func CompileWithOptions(prog *parser.Program, moduleName string) string {
	ctx := NewCompilerContext(moduleName)
	defer ctx.Dispose()

	mainFn := createMainFunction(ctx)
	entry := addEntryBlock(ctx, mainFn)
	setInsertPoint(ctx, entry)

	for _, stmt := range prog.Statements {
		compileStmt(stmt, ctx)
	}

	createMainReturn(ctx)

	return ctx.GetModule().String()
}
