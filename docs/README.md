# 🍕 Aether Documentation

Welcome to the complete Aether documentation! This is your one-stop resource for everything Aether - from getting started to advanced topics. 🍕

---

## 🍕 Quick Start

### New to Aether?

1. **[Learning Guide](learning.md)** - Start here! Complete tutorial from beginner to pro
2. **[Language Specification](specs.md)** - Complete language rules and syntax
3. **[Syntax Guide](syntax.md)** - Quick reference for syntax

### Having Issues?

1. **[Debugging Guide](debugging.md)** - Understanding and fixing errors
2. **[Error Codes](error-codes.md)** - Complete error code reference
3. **[Troubleshooting Guide](troubleshooting.md)** - Common problems and solutions

---

## 🍕 Project Structure (2025+)

```sh
Aether2/
  cmd/                # All CLI source files (aether2.go, build_cmd.go, etc)
  build/bin/          # Temporary folder for built binaries and tars
  src/                # Main source code
  packages/           # Modular packages
  docs/               # Documentation
  Makefile            # Build and publish automation
  aether.toml         # Project config
```

---

## 🍕 Development & Contributing

### 🛠️ CLI & Build Changes

- **All CLI code is now in `/cmd`** (not `/bin`).
- **`/build/bin`** is used for temporary build outputs and tars.
- **`aether update`** checks for new versions, supports mirrors, and manages multiple versions/forks.
- **Publishing:** Use `make publish` to build and tar binaries for all supported platforms/architectures. All builds are `.tar` (no compression).
- **Linter:** Use `aether lint` to check for casing and modularity. Use `aether fix` to auto-correct casing and structure (safe, parses code first).
- **Modular code is enforced:** The CLI, linter, and scaffolding all encourage tree-structured, modular code. Bad code is hard to write!

### 📝 How to Contribute

1. **Report documentation bugs** - Missing, incorrect, or unclear content
2. **Suggest improvements** - Better examples, clearer explanations
3. **Add examples** - Working code that demonstrates features
4. **Fix typos** - Grammar, spelling, and formatting issues
5. **Follow the new project structure** - All new CLI code in `/cmd`, binaries in `/build/bin`.
6. **Use the linter and fix tools** - Ensure casing and modularity before submitting PRs.

### 🎯 Guidelines

- **Be clear and concise**
- **Include examples**
- **Use consistent formatting**
- **Test everything**
- **Keep it fun**
- **Always modular, always tree-structured!**

---

## 🍕 Publishing & Version Management

- **To publish a new version:**
  1. Run `make publish` to build and tar all binaries for all platforms/architectures.
  2. Upload the generated `.tar` files from `/build/bin` to your release/mirror.
- **To update:**
  - Run `aether update` (downloads latest stable by default, use `--nightly` for nightlies).
  - Use `--mirror` to change the update source (e.g., for forks or alternate mirrors).
  - The updater manages multiple versions and can switch between them.

---

## 🍕 Example Tree

```sh
Aether2/
  cmd/
    aether2.go
    build_cmd.go
    linter_cmd.go
    update_cmd.go
    ...
  build/bin/
    aether-linux_amd64.tar
    aether-windows_386.tar
    ...
  src/
    main.ae
    ...
  packages/
    ...
  docs/
    README.md
    ...
  Makefile
  aether.toml
```

---

## 🍕 For More Info

- **[CLI Reference](cli.md)**
- **[Build System](build.md)**
- **[Dependencies](dependencies.md)**
- **[Learning Guide](learning.md)**
- **[Troubleshooting Guide](troubleshooting.md)**
