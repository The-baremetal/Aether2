# Simple Aether Package System 🍕

## The Problem
Aether's current package system is **too hard to create**. We need something **simple and fast**.

## The Solution: One-File Packages 📦

### 1. Simple Package Creation
```bash
# Create a package in ONE command:
aether new math
# That's it! No questions, no complex setup.
```

### 2. One-File Package Structure
```
myproject/
├── main.aeth
├── packages/
│   ├── math.aeth      # Just one file!
│   ├── utils.aeth     # Just one file!
│   └── http.aeth      # Just one file!
└── aether.toml
```

### 3. Simple Package Content
```aether
# packages/math.aeth
package math

func add(a, b) {
    return a + b
}

func multiply(a, b) {
    return a * b
}

func divide(a, b) {
    return a / b
}
```

### 4. Simple Import System
```aether
# main.aeth
import "math"

func main() {
    result = math.add(5, 3)
    print(result)
}
```

### 5. Automatic Package Discovery
```bash
# Aether automatically finds packages in packages/
# No need to declare them in aether.toml
```

## Implementation Plan 🚀

### Phase 1: Simple Package Creation
```bash
aether new <package_name>
```

This creates:
- `packages/<package_name>.aeth`
- Basic package template
- No complex configuration

### Phase 2: Simple Import Syntax
```aether
# Old (confusing):
import depmod
print(depmod.exported_value)

# New (simple):
import "math"
print(math.add(5, 3))
```

### Phase 3: Automatic Package Loading
- Aether scans `packages/` directory
- Automatically loads all `.aeth` files
- No manual configuration needed

### Phase 4: Package Visibility
```aether
# packages/math.aeth
package math

# Exported (can be imported)
func Add(a, b) {
    return a + b
}

# Internal (cannot be imported)
func internal_helper() {
    return "private"
}
```

## Benefits 🌟

### 1. **Super Simple Creation**
```bash
aether new math    # Creates packages/math.aeth
aether new utils   # Creates packages/utils.aeth
aether new http    # Creates packages/http.aeth
```

### 2. **No Complex Configuration**
- No `aether.toml` in packages
- No directory structure
- No build configuration
- Just one file!

### 3. **Familiar Import Syntax**
```aether
import "math"           # Standard import
import "utils" as u     # Aliased import
import "fmt" as .       # Dot import (global)
```

### 4. **Automatic Discovery**
- Aether finds packages automatically
- No manual dependency declaration
- Works out of the box

### 5. **Clear Visibility Rules**
- Capitalized functions = exported
- Lowercase functions = internal
- Simple and clear

## Migration Strategy 📋

### Step 1: Add Package Declaration
```aether
# Old
func add(a, b) {
    return a + b
}

# New
package math

func add(a, b) {
    return a + b
}
```

### Step 2: Update Import Syntax
```aether
# Old
import math
x = plus(2, 3)

# New
import "math"
x = math.add(2, 3)
```

### Step 3: Organize Standard Library
```
packages/
├── fmt.aeth      # Formatting and printing
├── os.aeth       # Operating system interface
├── io.aeth       # I/O primitives
├── math.aeth     # Mathematical functions
└── time.aeth     # Time operations
```

## Implementation Details 🔧

### 1. Package Declaration Parser
```go
// Add to lexer
PACKAGE = "package"

// Add to parser
func (p *Parser) parsePackage() *Package {
    if !p.expect(lexer.PACKAGE) {
        return nil
    }
    name := &Identifier{Value: p.curToken.Literal}
    if !p.expect(lexer.IDENT) {
        return nil
    }
    return &Package{Name: name}
}
```

### 2. Import Resolution
```go
func resolveImport(importPath string) string {
    // Check packages/ directory first
    packagePath := filepath.Join("packages", importPath+".aeth")
    if _, err := os.Stat(packagePath); err == nil {
        return packagePath
    }
    
    // Check standard library
    stdlibPath := filepath.Join("lib", importPath+".aeth")
    if _, err := os.Stat(stdlibPath); err == nil {
        return stdlibPath
    }
    
    return ""
}
```

### 3. Package Loading
```go
func loadPackages() map[string]*Package {
    packages := make(map[string]*Package)
    
    // Scan packages/ directory
    entries, err := os.ReadDir("packages")
    if err != nil {
        return packages
    }
    
    for _, entry := range entries {
        if strings.HasSuffix(entry.Name(), ".aeth") {
            packageName := strings.TrimSuffix(entry.Name(), ".aeth")
            packagePath := filepath.Join("packages", entry.Name())
            
            // Parse and load package
            if pkg := parsePackage(packagePath); pkg != nil {
                packages[packageName] = pkg
            }
        }
    }
    
    return packages
}
```

## Conclusion 🎉

This simplified package system makes Aether **much easier to use**:

1. **One command** to create packages
2. **One file** per package
3. **Simple imports** with quotes
4. **Automatic discovery** of packages
5. **Clear visibility** rules

**No more complex setup! No more confusing imports! Just simple, fast package creation!** 🚀 