# aether.toml Reference

The `aether.toml` file configures your Aether project. It is read automatically when you build or run your project.

---

## Example `aether.toml`

```toml
[project]
name = "myproject"
version = "0.1.0"
author = "Your Name"
description = "A sample Aether project."

[build]
source_directories = ["src"]
output_directory = "bin"
target = "native"           # e.g., "linux-amd64", "windows-amd64"
optimization = "2"          # 0, 1, 2, 3, s, z
linker = "mold"             # mold, ld, lld
create_library = false
library_type = "shared"     # shared, static, both
compiler_flags = {}          # Advanced: see below

[dependencies]
fmt = "packages/fmt"
math = "packages/math"

[dev-dependencies]
# Add development dependencies here
```

---

## `[project]` Section
- `name` (string): Project name.
- `version` (string): Project version.
- `author` (string, optional): Author name.
- `description` (string, optional): Project description.

## `[build]` Section
- `source_directories` (array of strings): Where to find source files. Default: `["src", "."]`
- `output_directory` (string): Where to place build outputs. Default: `bin`
- `target` (string): Build target (e.g., `native`, `linux-amd64`).
- `optimization` (string): Optimization level (`0`, `1`, `2`, `3`, `s`, `z`).
- `linker` (string): Linker to use (`mold`, `ld`, `lld`).
- `create_library` (bool): Build as a library. Default: `false`
- `library_type` (string): Type of library (`shared`, `static`, `both`).
- `compiler_flags` (table): Advanced compiler flags (see below).
- `targets` (table): Per-target build overrides (optional).

## `[dependencies]` Section
- Key-value pairs: `name = "path"` for each dependency.

## `[dev-dependencies]` Section
- Key-value pairs: Development-only dependencies.

---

## Advanced: `compiler_flags` Table
You can set advanced compiler flags here. See [compiler-flags.md](./compiler-flags.md) for all options.

---

## Precedence
- **Command-line flags always override `aether.toml` settings.**
- If a flag is not set on the command line, the value from `aether.toml` (if present) is used.
- If neither is set, the default is used.

---

For more details on flags, see [compiler-flags.md](./compiler-flags.md). 