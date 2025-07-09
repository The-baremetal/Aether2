# 🍕 Aether Debugging & Error Guide

Welcome to the ultimate debugging guide for Aether! This guide will help you understand, fix, and learn from every error you encounter. No more cryptic messages—just clear, actionable advice with a dash of pizza personality! 🍕

---

## 🍕 Quick Start: Understanding Error Messages

Every Aether error follows this format:

```bash
🍕 SyntaxError: Unexpected token at line 2, column 1: expected IN, got ILLEGAL
    for x (something) { ... }
    ^
    Fix: Use 'in' for for-loops. See: specs.md §9
```

**What each part means:**

- **🍕 Emoji:** Error type (🍕 for syntax, 🚨 for warnings, etc.)
- **Error Type:** SyntaxError, TypeError, etc.
- **Location:** Line and column where the error occurred
- **Message:** What went wrong in plain English
- **Code Snippet:** The actual line with the error
- **Caret (^):** Points to the exact problematic character
- **Fix:** How to fix it
- **Spec Reference:** Link to the relevant language specification

---

## 🍕 Error Types & Categories

### Syntax Errors (Most Common)
These prevent your code from compiling. Fix these first!

| Error Code | What It Means | Common Causes | How to Fix |
|------------|---------------|---------------|------------|
| `UnexpectedToken` | Parser found a token it wasn't expecting | Wrong syntax, missing keywords | Check the syntax in specs.md |
| `UnexpectedEOF` | File ended unexpectedly | Missing closing brace/bracket | Add the missing closing character |
| `InvalidSyntax` | General syntax problem | Malformed expression | Rewrite the expression |
| `UnterminatedString` | String literal not closed | Missing quote | Add the missing quote |
| `InvalidNumber` | Number format is wrong | Invalid numeric literal | Fix the number format |
| `UnexpectedSemicolon` | Semicolon where it shouldn't be | C/Java habits | Remove the semicolon |

### Type Errors (Runtime)
These happen when your code runs but does something wrong.

| Error Code | What It Means | Common Causes | How to Fix |
|------------|---------------|---------------|------------|
| `TypeMismatch` | Wrong type for operation | Adding string to number | Convert types explicitly |
| `UndefinedVariable` | Variable doesn't exist | Misspelled name | Check spelling, declare variable |
| `UndefinedFunction` | Function doesn't exist | Misspelled function name | Check spelling, import function |

### Warning Errors (Non-blocking)
These don't stop compilation but suggest improvements.

| Error Code | What It Means | Common Causes | How to Fix |
|------------|---------------|---------------|------------|
| `UnusedVariable` | Variable declared but never used | Dead code | Remove or use the variable |
| `UnusedImport` | Imported module not used | Unnecessary import | Remove the import |
| `DeprecatedFeature` | Using old syntax | Outdated code | Update to new syntax |

---

## 🍕 Common Error Patterns & Solutions

### 1. Function Declaration Errors

**❌ Wrong:**
```aether
func add(a, b) {
  return a + b
}
```

**✅ Right:**
```aether
func add(a, b) {
  return a + b
}
```

**Common Issues:**
- Missing `func` keyword
- Wrong parameter syntax
- Missing return statement

### 2. For Loop Errors

**❌ Wrong:**
```aether
for x (array) {
  print(x)
}
```

**✅ Right:**
```aether
for x in array {
  print(x)
}
```

**Common Issues:**
- Missing `in` keyword
- Wrong parentheses syntax
- Using `range` instead of `in`

### 3. Import Errors

**❌ Wrong:**
```aether
import "math"
```

**✅ Right:**
```aether
import math
```

**Common Issues:**
- Using quotes around module names
- Missing semicolons (Aether doesn't use them)
- Wrong import path

### 4. Variable Assignment Errors

**❌ Wrong:**
```aether
let x = 10;
```

**✅ Right:**
```aether
x = 10
```

**Common Issues:**
- Using `let`, `var`, or `const` (Aether doesn't use them)
- Adding semicolons
- Using `=` instead of `==` for comparison

---

## 🍕 Debugging Workflow

### Step 1: Read the Error Message
1. **Look at the error type** (SyntaxError, TypeError, etc.)
2. **Check the line and column** numbers
3. **Read the message** carefully
4. **Look at the code snippet** and caret position

### Step 2: Understand the Problem
1. **What was expected?** (shown in the error message)
2. **What was found instead?** (the actual token)
3. **Why did this happen?** (common causes listed above)

### Step 3: Apply the Fix
1. **Follow the suggested fix** if provided
2. **Check the spec reference** for detailed rules
3. **Test your fix** by running the code again

### Step 4: Learn from It
1. **Understand why the error happened**
2. **Remember the correct syntax** for next time
3. **Check similar code** for the same issue

---

## 🍕 Error Recovery Strategies

### When You Get Many Errors
If you see a flood of errors after one mistake:

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
2. **Check the examples** in this guide
3. **Look at working code** for reference
4. **Ask for help** in the community

---

## 🍕 Error Message Reference

### SyntaxError Messages

| Message | Meaning | Example | Fix |
|---------|---------|---------|-----|
| `expected IN, got ILLEGAL` | Missing `in` keyword in for loop | `for x (array)` | Use `for x in array` |
| `expected IDENT, got LPAREN` | Function parameter needs a name | `func add(, b)` | Add parameter name |
| `expected IDENT, got VARARG` | Varargs must be last parameter | `func f(...a, b)` | Move `...a` to end |
| `unexpected token in parsePrimary` | Invalid expression | `func f(1 +)` | Complete the expression |
| `expected expression for return value` | Return needs a value | `return` | Add return value or `return 0` |

### TypeError Messages

| Message | Meaning | Example | Fix |
|---------|---------|---------|-----|
| `cannot add string and number` | Type mismatch in addition | `"hello" + 5` | Convert to same type |
| `undefined variable 'x'` | Variable not declared | `x = 10` without declaration | Declare variable first |
| `undefined function 'print'` | Function not imported | `print("hello")` | Import stdlib or declare function |

---

## 🍕 Learning Resources

### Official Documentation
- **[Language Specification](specs.md):** Complete language rules
- **[Syntax Guide](syntax.md):** Quick syntax reference
- **[Examples](../examples/):** Working code examples

### Common Patterns
- **[Function Patterns](patterns/functions.md):** How to write functions
- **[Loop Patterns](patterns/loops.md):** Different loop types
- **[Import Patterns](patterns/imports.md):** Module importing

### Debugging Tools
- **`aether build --verbose`:** Show detailed compilation info
- **`aether build --analyze-only`:** Check for errors without compiling
- **`aether build --max-errors=3`:** Limit error output for debugging

---

## 🍕 Getting Help

### When to Ask for Help
- You've tried the fixes in this guide
- The error message doesn't make sense
- You've checked the spec and examples
- You're still stuck after 15 minutes

### How to Ask for Help
1. **Include the full error message**
2. **Show the code that caused it**
3. **Explain what you were trying to do**
4. **Mention what you've already tried**

### Example Good Help Request
```
I'm getting this error:
🍕 SyntaxError: Unexpected token at line 3, column 15: expected IDENT, got LPAREN

My code:
func add(a, b) {
  return a + b
}

I was trying to create a function that adds two numbers. I've checked the syntax guide and it looks correct to me. What am I missing?
```

---

## 🍕 Pro Tips

### 1. Start Small
- Write and test small pieces of code
- Add complexity gradually
- Don't try to write the whole program at once

### 2. Use the Spec
- The language specification is your friend
- Reference it when you're unsure about syntax
- It's more authoritative than examples

### 3. Learn from Errors
- Each error is a learning opportunity
- Understand why it happened
- Remember the correct syntax for next time

### 4. Keep It Simple
- Aether is designed to be simple
- If something feels complicated, you're probably doing it wrong
- Look for the simpler solution

---

## 🍕 Error Code Quick Reference

| Code | Type | Severity | Action |
|------|------|----------|--------|
| `UnexpectedToken` | Syntax | Critical | Fix syntax |
| `UnexpectedEOF` | Syntax | Critical | Add missing closing |
| `InvalidSyntax` | Syntax | Critical | Rewrite expression |
| `UnterminatedString` | Syntax | Critical | Add missing quote |
| `InvalidNumber` | Syntax | Critical | Fix number format |
| `UnexpectedSemicolon` | Syntax | Critical | Remove semicolon |
| `TypeMismatch` | Type | Error | Convert types |
| `UndefinedVariable` | Type | Error | Declare variable |
| `UndefinedFunction` | Type | Error | Import or declare |
| `UnusedVariable` | Warning | Info | Remove or use |
| `UnusedImport` | Warning | Info | Remove import |
| `DeprecatedFeature` | Warning | Info | Update syntax |

---

**Remember:** Every error is a step toward better code! Don't get frustrated—get curious! 🍕✨

Happy debugging, pizza lovers! 🍕 