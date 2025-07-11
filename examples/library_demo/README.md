# 🍕 Aether Library Demo

This example demonstrates how to create and use libraries in Aether!

## Creating a Library

### 1. Create the Library Source

The `src/mathlib.ae` file contains math functions that will be compiled into a shared library.

### 2. Build the Library

```bash
# Create a shared library
aether build --create-library --library-type shared --library-name mathlib src/mathlib.ae

# Create both shared and static libraries
aether build --create-library --library-type both --library-name mathlib src/mathlib.ae

# Create with pkg-config file
aether build --create-library --library-type shared --library-name mathlib --generate-pkg-config src/mathlib.ae
```

### 3. Using the Library Command

```bash
# Create library using the library command
aether library --create mathlib --lib-type shared --generate-pc

# Analyze an existing library
aether library --analyze libc --verbose

# Generate binding for external library
aether library --bind openssl --output openssl_binding.ae
```

## Library Types

### Shared Libraries (.so/.dll/.dylib)
- **Linux**: `libmathlib.so`
- **Windows**: `mathlib.dll`
- **macOS**: `libmathlib.dylib`

### Static Libraries (.a/.lib)
- **Linux**: `libmathlib.a`
- **Windows**: `mathlib.lib`

### pkg-config Files (.pc)
- **File**: `mathlib.pc`
- **Usage**: `pkg-config --libs --cflags mathlib`

## Using the Library

The `src/main.ae` file shows how to import and use the library:

```aether
import "mathlib" as math

result = math.Add(5, 3)
print("5 + 3 =", result)
```

## Library Analysis

Aether can analyze external libraries and generate bindings:

```bash
# Analyze system library
aether library --analyze libc --verbose

# Generate binding for OpenSSL
aether library --bind openssl --output openssl_binding.ae
```

## Cross-Platform Support

Libraries work across all platforms:
- **Linux**: `.so` files
- **Windows**: `.dll` files  
- **macOS**: `.dylib` files

## Performance Benefits

- **Fast linking** with mold
- **LLVM optimization** for maximum performance
- **No runtime overhead** - direct library calls
- **Single binary** deployment

## The Future

Aether's library system enables:
- **Easy C/C++ integration**
- **System library access**
- **Cross-platform compatibility**
- **High performance**
- **Simple deployment**

**Aether = Go's simplicity + Rust's performance + C's library ecosystem!** 🚀🍕 