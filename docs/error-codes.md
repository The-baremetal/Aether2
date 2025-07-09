# 🍕 Aether Error Codes Reference

Quick reference for all Aether error codes, their meanings, and how to fix them. Use this when you need to understand an error quickly! 🍕

---

## 🍕 Syntax Error Codes

### `UnexpectedToken`
**What it means:** The parser found a token it wasn't expecting in this context.

**Common causes:**
- Wrong syntax for the language construct
- Missing keywords
- Extra characters

**Examples:**
```aether
// ❌ Wrong
for x (array) { ... }
// Expected: for x in array { ... }

func add(, b) { ... }
// Expected: func add(a, b) { ... }
```

**How to fix:**
1. Check the syntax in the language specification
2. Look for missing keywords
3. Remove extra characters

---

### `UnexpectedEOF`
**What it means:** The file ended unexpectedly, usually missing a closing character.

**Common causes:**
- Missing closing brace `}`
- Missing closing bracket `]`
- Missing closing parenthesis `)`
- Unclosed string literal

**Examples:**
```aether
// ❌ Wrong
func add(a, b) {
  return a + b
// Missing closing brace

x = [1, 2, 3
// Missing closing bracket
```

**How to fix:**
1. Add the missing closing character
2. Check for balanced braces/brackets/parentheses
3. Use an editor with bracket matching

---

### `InvalidSyntax`
**What it means:** General syntax problem with an expression or statement.

**Common causes:**
- Malformed expression
- Invalid operator usage
- Wrong statement structure

**Examples:**
```aether
// ❌ Wrong
func f(1 +) { ... }
// Incomplete expression

x = + 5
// Invalid operator usage
```

**How to fix:**
1. Complete the expression
2. Check operator precedence
3. Rewrite the statement

---

### `UnterminatedString`
**What it means:** String literal is not properly closed with a quote.

**Common causes:**
- Missing closing quote
- Wrong quote type
- Escaped quote issues

**Examples:**
```aether
// ❌ Wrong
message = "Hello world
// Missing closing quote

message = "He said "Hello""
// Nested quotes without escaping
```

**How to fix:**
1. Add the missing closing quote
2. Use different quote types for nested strings
3. Escape quotes with backslash

---

### `InvalidNumber`
**What it means:** Number format is invalid or unsupported.

**Common causes:**
- Invalid numeric literal
- Unsupported number format
- Wrong base notation

**Examples:**
```aether
// ❌ Wrong
x = 123.456.789
// Invalid decimal format

x = 0xGG
// Invalid hexadecimal
```

**How to fix:**
1. Use valid number format
2. Check for typos in numbers
3. Use supported number bases

---

### `UnexpectedSemicolon`
**What it means:** Semicolon found where it shouldn't be (Aether doesn't use semicolons).

**Common causes:**
- C/Java/JavaScript habits
- Copy-paste from other languages
- Automatic semicolon insertion

**Examples:**
```aether
// ❌ Wrong
x = 10;
func add(a, b) {
  return a + b;
}
```

**How to fix:**
1. Remove all semicolons
2. Aether uses newlines for statement separation
3. No semicolons needed anywhere

---

## 🍕 Type Error Codes

### `TypeMismatch`
**What it means:** Operation between incompatible types.

**Common causes:**
- Adding string to number
- Comparing different types
- Wrong function argument types

**Examples:**
```aether
// ❌ Wrong
result = "hello" + 5
// Can't add string and number

if "hello" == 5 {
  // Can't compare string and number
}
```

**How to fix:**
1. Convert types explicitly
2. Use appropriate operators
3. Check function parameter types

---

### `UndefinedVariable`
**What it means:** Variable is used but not declared.

**Common causes:**
- Misspelled variable name
- Variable not declared
- Scope issues

**Examples:**
```aether
// ❌ Wrong
print(x)
// x not declared

func add(a, b) {
  return a + b + c
  // c not declared
}
```

**How to fix:**
1. Declare the variable first
2. Check spelling
3. Verify scope

---

### `UndefinedFunction`
**What it means:** Function is called but not defined or imported.

**Common causes:**
- Misspelled function name
- Function not imported
- Function not declared

**Examples:**
```aether
// ❌ Wrong
print("hello")
// print not imported

result = add(1, 2)
// add not declared
```

**How to fix:**
1. Import the function from stdlib
2. Declare the function
3. Check spelling

---

## 🍕 Warning Error Codes

### `UnusedVariable`
**What it means:** Variable is declared but never used.

**Common causes:**
- Dead code
- Unused parameters
- Debugging leftovers

**Examples:**
```aether
// ❌ Warning
x = 10
// x never used

func add(a, b) {
  return a + b
  // b parameter not used
}
```

**How to fix:**
1. Remove unused variables
2. Use the variable
3. Prefix with underscore for intentionally unused

---

### `UnusedImport`
**What it means:** Module is imported but not used.

**Common causes:**
- Unnecessary imports
- Dead code
- Copy-paste leftovers

**Examples:**
```aether
// ❌ Warning
import math
// math not used

import stdlib
// stdlib not used
```

**How to fix:**
1. Remove unused imports
2. Use the imported functions
3. Keep only necessary imports

---

### `DeprecatedFeature`
**What it means:** Using old syntax that will be removed.

**Common causes:**
- Outdated code
- Old tutorials
- Migration needed

**Examples:**
```aether
// ❌ Deprecated
old_syntax()
// Use new_syntax() instead
```

**How to fix:**
1. Update to new syntax
2. Check migration guide
3. Use current language features

---

## 🍕 Error Severity Levels

### Critical (Build Fails)
- `UnexpectedToken`
- `UnexpectedEOF`
- `InvalidSyntax`
- `UnterminatedString`
- `InvalidNumber`
- `UnexpectedSemicolon`

### Error (Runtime Issues)
- `TypeMismatch`
- `UndefinedVariable`
- `UndefinedFunction`

### Warning (Suggestions)
- `UnusedVariable`
- `UnusedImport`
- `DeprecatedFeature`

---

## 🍕 Quick Fix Patterns

### Common Syntax Fixes

| Error | Quick Fix |
|-------|-----------|
| `expected IN, got ILLEGAL` | Add `in` keyword: `for x in array` |
| `expected IDENT, got LPAREN` | Add parameter name: `func add(a, b)` |
| `expected IDENT, got VARARG` | Move varargs to end: `func f(a, ...b)` |
| `unexpected token in parsePrimary` | Complete the expression |
| `expected expression for return value` | Add return value or `return 0` |

### Common Type Fixes

| Error | Quick Fix |
|-------|-----------|
| `cannot add string and number` | Convert to same type: `"hello" + str(5)` |
| `undefined variable 'x'` | Declare variable: `x = 10` |
| `undefined function 'print'` | Import stdlib: `import stdlib` |

---

## 🍕 Error Recovery Tips

### When You Get Many Errors
1. **Fix the first error only** - others are likely cascading
2. **Re-run the build** after each fix
3. **Don't try to fix everything at once**

### When Errors Don't Make Sense
1. **Check for hidden characters** (copy-paste issues)
2. **Verify file encoding** (should be UTF-8)
3. **Look for missing braces/brackets**
4. **Check indentation** (Aether is sensitive to this)

### When You're Stuck
1. **Read the relevant spec section** (linked in error messages)
2. **Check the examples** in the debugging guide
3. **Look at working code** for reference
4. **Ask for help** in the community

---

**Remember:** Every error code is a learning opportunity! Use this reference to understand what went wrong and how to fix it quickly! 🍕✨ 