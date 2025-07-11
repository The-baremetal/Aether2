# Reimagined Aether Imports - Go-Style Package System

## Current Aether Import System

Currently Aether uses:
```aether
import math
x = plus(2, 3)

import math as math
z = math.plus(2, 3)
```

## Proposed Go-Style Package System

### 1. Package Declaration

Every `.aeth` file starts with a package declaration:

```aether
package main

func hello() {
    print("Hello, Aether!")
}
```

```aether
package math

func add(a, b) {
    return a + b
}

func multiply(a, b) {
    return a * b
}
```

```aether
package utils

func format_string(template, ...args) {
    // String formatting logic
}
```

### 2. Import System

Import packages using Go-style syntax:

```aether
package main

import "math"
import "utils"

func main() {
    result = math.add(5, 3)
    message = utils.format_string("Result: {}", result)
    print(message)
}
```

### 3. Package Organization

#### Directory Structure
```
myproject/
├── main.aeth          # package main
├── math/
│   ├── arithmetic.aeth  # package math
│   └── geometry.aeth    # package math
├── utils/
│   ├── string.aeth      # package utils
│   └── file.aeth        # package utils
└── aether.toml
```

#### Package Rules
- **One package per directory**
- **Package name matches directory name**
- **All files in directory share same package name**
- **Main package is special** - entry point

### 4. Import Options

#### Standard Import
```aether
import "math"
result = math.add(5, 3)
```

#### Aliased Import
```aether
import "math" as m
result = m.add(5, 3)
```

#### Dot Import (Global namespace)
```aether
import "math" as .
result = add(5, 3)  // No prefix needed
```

#### Selective Import
```aether
import { add, multiply } from "math"
result = add(5, 3)
product = multiply(4, 2)
```

### 5. Package Visibility

#### Exported vs Unexported
```aether
package math

// Exported (capitalized) - can be imported
func Add(a, b) {
    return a + b
}

// Unexported (lowercase) - internal only
func internal_helper() {
    // Internal logic
}
```

#### Import Rules
- **Capitalized functions** are exported and importable
- **Lowercase functions** are internal to package
- **Package name** is always exported

### 6. Package Initialization

#### Init Functions
```aether
package math

var cache = []

func init() {
    // Package initialization
    cache = [1, 2, 3, 4, 5]
}

func Add(a, b) {
    return a + b
}
```

### 7. Subpackages

#### Nested Package Structure
```
myproject/
├── main.aeth
├── math/
│   ├── arithmetic.aeth
│   ├── geometry/
│   │   ├── circle.aeth
│   │   └── rectangle.aeth
│   └── statistics/
│       ├── mean.aeth
│       └── variance.aeth
```

#### Importing Subpackages
```aether
import "math/geometry"
import "math/statistics"

func main() {
    area = geometry.circle_area(5)
    avg = statistics.mean([1, 2, 3, 4, 5])
}
```

### 8. Standard Library Organization

#### Built-in Packages
```aether
import "fmt"      // Formatting and printing
import "os"       // Operating system interface
import "io"       // I/O primitives
import "strconv"  // String conversions
import "time"     // Time operations
import "math"     // Mathematical functions
```

### 9. Package Management

#### aether.toml Dependencies
```toml
[project]
name = "myapp"
version = "1.0.0"

[dependencies]
math = "packages/math"
utils = "packages/utils"
http = "packages/http"
```

#### Import Resolution
1. **Local packages** - `import "math"` → `./math/`
2. **Standard library** - `import "fmt"` → built-in
3. **External packages** - `import "github.com/user/pkg"` → downloaded

### 10. Benefits of Go-Style Packages

#### **Simplicity**
- One package per directory
- Clear naming conventions
- Simple import paths

#### **Organization**
- Logical grouping of related code
- Clear separation of concerns
- Easy to navigate

#### **Visibility Control**
- Capitalized = exported
- Lowercase = internal
- No complex visibility modifiers

#### **Standard Library**
- Consistent package structure
- Easy to learn and use
- Familiar to Go developers

#### **Tooling Support**
- Easy to implement package discovery
- Simple dependency resolution
- Clear build system integration

### 11. Migration Strategy

#### **Phase 1: Add Package Declarations**
```aether
// Old
func main() {
    print("Hello")
}

// New
package main

func main() {
    print("Hello")
}
```

#### **Phase 2: Update Import Syntax**
```aether
// Old
import math
x = plus(2, 3)

// New
import "math"
x = math.add(2, 3)
```

#### **Phase 3: Organize Standard Library**
```
stdlib/
├── fmt/
├── os/
├── io/
├── math/
└── time/
```

### 12. Implementation Plan

1. **Add package declaration parsing** to lexer/parser
2. **Update import resolution** to use package paths
3. **Implement visibility rules** (capitalized = exported)
4. **Create standard library packages**
5. **Update build system** to handle packages
6. **Add package management tools**

## Conclusion

Go's package system is **PERFECT** for Aether because it's:
- **Simple** - Easy to understand and use
- **Organized** - Clear structure and conventions
- **Scalable** - Works for small and large projects
- **Familiar** - Many developers already know it

**This would make Aether much more professional and easier to use!** 🚀 