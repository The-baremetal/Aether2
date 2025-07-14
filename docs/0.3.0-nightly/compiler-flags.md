# Aether Compiler Flags Reference

This document lists all command-line flags supported by the Aether compiler (`aether build`).

---

## Output Flags

| Flag                | Description                        | Default           | Example                      |
|---------------------|------------------------------------|-------------------|------------------------------|
| `--output`, `-o`    | Output executable name              | `bin/aether.out`  | `-o myprog.exe`              |
| `--emit-ir`         | Emit LLVM IR (.ll)                  | false             | `--emit-ir`                  |
| `--emit-asm`        | Emit assembly (.s)                  | false             | `--emit-asm`                 |
| `--emit-bitcode`    | Emit bitcode (.bc)                  | false             | `--emit-bitcode`             |
| `--emit-obj`        | Emit object files (.o)              | false             | `--emit-obj`                 |
| `--emit-exe`        | Emit executable                     | true              | `--emit-exe`                 |
| `--emit-tokens`     | Emit lexer tokens for debugging     | false             | `--emit-tokens`              |

## Optimization Flags

| Flag                | Description                        | Default           | Example                      |
|---------------------|------------------------------------|-------------------|------------------------------|
| `-O`                | Optimization level (0,1,2,3,s,z)   | `2`               | `-O3`                        |
| `--no-optimize`     | Disable all optimizations           | false             | `--no-optimize`              |
| `--no-inline`       | Disable function inlining           | false             | `--no-inline`                |
| `--no-vectorize`    | Disable vectorization               | false             | `--no-vectorize`             |
| `--no-unroll`       | Disable loop unrolling              | false             | `--no-unroll`                |

## Debug Flags

| Flag                | Description                        | Default           | Example                      |
|---------------------|------------------------------------|-------------------|------------------------------|
| `--debug-info`      | Generate debug information          | false             | `--debug-info`               |
| `--debug-symbols`   | Include debug symbols               | false             | `--debug-symbols`            |
| `--strip`           | Strip debug symbols from output     | false             | `--strip`                    |

## Target/Platform Flags

| Flag                | Description                        | Default           | Example                      |
|---------------------|------------------------------------|-------------------|------------------------------|
| `--target-os`       | Target operating system             | host OS           | `--target-os=linux`          |
| `--target-arch`     | Target architecture                 | host arch         | `--target-arch=arm64`        |
| `--linker`          | Linker to use (mold, ld, lld)       | `mold`            | `--linker=ld`                |
| `--fuse-ld`         | Linker to use (like clang -fuse-ld) | (empty)           | `--fuse-ld=lld`              |

## Library/Linking Flags

| Flag                | Description                        | Default           | Example                      |
|---------------------|------------------------------------|-------------------|------------------------------|
| `--static`          | Create static executable            | false             | `--static`                   |
| `--shared`          | Create shared library               | false             | `--shared`                   |
| `--pie`             | Position independent executable     | false             | `--pie`                      |
| `--rdynamic`        | Add all symbols to dynamic table    | false             | `--rdynamic`                 |
| `--export-dynamic`  | Export all symbols                  | false             | `--export-dynamic`           |
| `--no-stdlib`       | Disable stdlib builtins             | false             | `--no-stdlib`                |
| `--no-default-libs` | Don't link default libraries        | false             | `--no-default-libs`          |
| `--nostdlib`        | Don't link standard library         | false             | `--nostdlib`                 |
| `--nostartfiles`    | Don't link startup files            | false             | `--nostartfiles`             |

## Analysis/Debugging Flags

| Flag                | Description                        | Default           | Example                      |
|---------------------|------------------------------------|-------------------|------------------------------|
| `--check-imports`   | Check import validity               | true              | `--check-imports=false`      |
| `--analyze-only`    | Only analyze, don't compile         | false             | `--analyze-only`             |
| `--time-compile`    | Time compilation phases             | false             | `--time-compile`             |
| `--stats`           | Show compilation statistics         | false             | `--stats`                    |
| `--profile`         | Enable profiling                    | false             | `--profile`                  |

## Miscellaneous Flags

| Flag                | Description                        | Default           | Example                      |
|---------------------|------------------------------------|-------------------|------------------------------|
| `--verbose`, `-v`   | Verbose output                     | false             | `-v`                         |
| `--quiet`, `-q`     | Suppress output                    | false             | `-q`                         |
| `--help`            | Show help                          | false             | `--help`                     |
| `--version`         | Show version                       | false             | `--version`                  |

---

## Flag Precedence
- **Command-line flags always override `aether.toml` settings.**
- If a flag is not set on the command line, the value from `aether.toml` (if present) is used.
- If neither is set, the default is used.

---

For more details, see the [aether.toml documentation](./aether-toml.md). 