package compiler

import (
	"aether/src/parser"
	"fmt"
	"os"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"aether/src/lexer"
	"aether/src/analysis"
)

func compileStmt(stmt parser.Statement, ctx *CompilerContext, analysisResult interface{}) {
	switch s := stmt.(type) {
	case *parser.Import:
		moduleName := s.Name.Value
		if s.As != nil && s.As.Value != "" {
			moduleName = s.As.Value
		}
		// Use analysisResult to get the resolved import path
		if analysisResultTyped, ok := analysisResult.(*analysis.AnalysisResult); ok {
			fmt.Println("[IMPORT-DEBUG] analysisResultTyped.Imports keys:", getMapKeys(analysisResultTyped.Imports))
			fmt.Println("[IMPORT-DEBUG] s.Name.Value:", s.Name.Value)
			importInfo, exists := analysisResultTyped.Imports[s.Name.Value]
			if exists && importInfo.Resolved != "" {
				if data, err := os.ReadFile(importInfo.Resolved); err == nil {
					l := lexer.NewLexer(string(data))
					p := parser.NewParser(l)
					prog := p.Parse()
					for _, stmt := range prog.Statements {
						fmt.Printf("[IMPORT-DEBUG] Statement type: %T, value: %#v\n", stmt, stmt)
						if fn, ok := stmt.(*parser.Function); ok && fn.Name != nil {
							if fn.Body == nil {
								fmt.Println("[IMPORT-DEBUG] Skipping foreign function:", fn.Name.Value)
								continue
							}
							f := ctx.module.NewFunc(fn.Name.Value, types.I32)
							ctx.SetSymbol(fn.Name.Value, f)
							fmt.Println("[IMPORT-DEBUG] Registered function from import:", fn.Name.Value)
						}
					}
				}
			}
		}
		ctx.SetSymbol(moduleName, constant.NewInt(types.I32, 0))
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
		ctx.DebugPrint("  [IR] Emitting function: " + s.Name.Value)
		fn := ctx.module.NewFunc(s.Name.Value, types.I32)
		ctx.SetSymbol(s.Name.Value, fn)
		block := fn.NewBlock("entry")
		ctx.builder = block
		for i, stmt := range s.Body.Statements {
			ctx.DebugPrint("    [IR] Compiling function body statement #" + fmt.Sprint(i) + ": " + fmt.Sprintf("%T", stmt))
			compileStmt(stmt, ctx, analysisResult)
		}
		ctx.builder.NewRet(constant.NewInt(types.I32, 0))
		ctx.DebugPrint("  [IR] Finished function: " + s.Name.Value)
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
			compileStmt(stmt, ctx, analysisResult)
		}
		ctx.builder.NewBr(mergeBlock)
		ctx.builder = elseBlock
		if s.Alternative != nil {
			for _, stmt := range s.Alternative.Statements {
				compileStmt(stmt, ctx, analysisResult)
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
			compileStmt(stmt, ctx, analysisResult)
		}
		ctx.builder.NewBr(condBlock)
		ctx.builder = endBlock
	case *parser.Repeat:
		// Not implemented: repeat
	case *parser.For:
		// Not implemented: for
	case *parser.Block:
		for _, stmt := range s.Statements {
			compileStmt(stmt, ctx, analysisResult)
		}
	case *parser.Match:
		// TODO: Proper match/case codegen. For now, just compile all case bodies.
		for _, c := range s.Cases {
			compileStmt(c.Body, ctx, analysisResult)
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
	case *parser.ExpressionStatement:
		compileExpr(s.Expr, ctx)
	case *parser.Call:
		compileExpr(s, ctx)
	}
}

// Helper function for debug printing map keys
func getMapKeys(m map[string]analysis.ImportInfo) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
