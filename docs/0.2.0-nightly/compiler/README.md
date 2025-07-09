# Compiler Directory

This directory contains the core files for the LLVM IR code generation backend for the Aether language. All design and implementation is based on the documentation in /docs/llvm.

- module.go: Entry point for LLVM IR module generation.
- context.go: Compilation context, variable scope, and block management.
- expr.go: Expression compilation.
- stmt.go: Statement compilation (define, return, if, etc).
- types.go: Common LLVM type helpers.

For more details, see /docs/llvm/src/user-guide and /docs/llvm/researchllvm. 