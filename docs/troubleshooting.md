# 🍕 Aether Troubleshooting Guide

Having trouble with Aether? This guide covers the most common issues and their solutions. From build errors to runtime problems, we've got you covered! 🍕

---

## 🍕 Build Issues

### "Command not found: aether"
**Problem:** The `aether` command isn't recognized.

**Solutions:**
1. **Check if Aether is installed:**
   ```bash
   which aether
   ```

2. **Add to PATH if needed:**
   ```bash
   export PATH=$PATH:/path/to/aether/bin
   ```

3. **Install Aether properly:**
   ```bash
   go install github.com/your-repo/aether@latest
   ```

### "Cannot find module"
**Problem:** Module resolution fails.

**Solutions:**
1. **Check your `aether.toml`:**
   ```toml
   [dependencies]
   stdlib = "0.1.0"
   ```

2. **Update dependencies:**
   ```bash
   aether dep update
   ```

3. **Clear cache and rebuild:**
   ```bash
   rm -rf .aether
   aether build
   ```

### "Import error: module not found"
**Problem:** Can't import a module.

**Solutions:**
1. **Check module name spelling:**
   ```aether
   import stdlib  # ✅ Correct
   import "stdlib"  # ❌ Wrong
   ```

2. **Verify module exists:**
   ```bash
   ls src/stdlib/
   ```

3. **Check import path:**
   ```aether
   import my_module  # For src/my_module.ae
   ```

---

## 🍕 Compilation Errors

### "Unexpected token" Errors
**Problem:** Parser can't understand your syntax.

**Common fixes:**
```aether
// ❌ Wrong
for x (array) { ... }

// ✅ Right
for x in array { ... }
```

```aether
// ❌ Wrong
let x = 10;

// ✅ Right
x = 10
```

```aether
// ❌ Wrong
import "stdlib"

// ✅ Right
import stdlib
```

### "Unexpected EOF" Errors
**Problem:** Missing closing characters.

**Solutions:**
1. **Check for missing braces:**
   ```aether
   func add(a, b) {
     return a + b
   }  // Make sure this brace is here
   ```

2. **Check for missing brackets:**
   ```aether
   numbers = [1, 2, 3]  // Make sure this bracket is here
   ```

3. **Check for missing quotes:**
   ```aether
   message = "Hello, World!"  // Make sure this quote is here
   ```

### "Invalid syntax" Errors
**Problem:** Malformed expressions or statements.

**Solutions:**
1. **Complete expressions:**
   ```aether
   // ❌ Wrong
   x = 5 +
   
   // ✅ Right
   x = 5 + 3
   ```

2. **Fix operator usage:**
   ```aether
   // ❌ Wrong
   x = + 5
   
   // ✅ Right
   x = 5
   ```

---

## 🍕 Runtime Errors

### "Undefined function" Errors
**Problem:** Function not found.

**Solutions:**
1. **Import the function:**
   ```aether
   import stdlib
   print("Hello")
   ```

2. **Check spelling:**
   ```aether
   // ❌ Wrong
   prin("Hello")
   
   // ✅ Right
   print("Hello")
   ```

3. **Declare the function:**
   ```aether
   func my_function() {
     return "Hello"
   }
   ```

### "Undefined variable" Errors
**Problem:** Variable used before declaration.

**Solutions:**
1. **Declare variables first:**
   ```aether
   x = 10
   print(x)
   ```

2. **Check scope:**
   ```aether
   func my_func() {
     local_var = 10
     return local_var
   }
   ```

3. **Check spelling:**
   ```aether
   // ❌ Wrong
   my_variable = 10
   print(my_variabel)
   
   // ✅ Right
   my_variable = 10
   print(my_variable)
   ```

### "Type mismatch" Errors
**Problem:** Incompatible types in operations.

**Solutions:**
1. **Convert types explicitly:**
   ```aether
   // ❌ Wrong
   result = "Hello" + 5
   
   // ✅ Right
   result = "Hello" + str(5)
   ```

2. **Use appropriate operators:**
   ```aether
   // ❌ Wrong
   if "hello" == 5 { ... }
   
   // ✅ Right
   if "hello" == "hello" { ... }
   ```

---

## 🍕 Performance Issues

### Slow Compilation
**Problem:** Build takes too long.

**Solutions:**
1. **Use incremental builds:**
   ```bash
   aether build --incremental
   ```

2. **Limit error output:**
   ```bash
   aether build --max-errors=5
   ```

3. **Clean and rebuild:**
   ```bash
   aether clean
   aether build
   ```

### Memory Issues
**Problem:** Program uses too much memory.

**Solutions:**
1. **Check for infinite loops:**
   ```aether
   // ❌ Wrong
   while true {
     print("Infinite loop")
   }
   
   // ✅ Right
   counter = 0
   while counter < 10 {
     print(counter)
     counter = counter + 1
   }
   ```

2. **Optimize data structures:**
   ```aether
   // Use arrays instead of large objects
   numbers = [1, 2, 3, 4, 5]
   ```

---

## 🍕 Development Environment Issues

### Editor Integration
**Problem:** Your editor doesn't recognize Aether syntax.

**Solutions:**
1. **Install Aether extension** (if available)
2. **Configure file associations:**
   ```json
   {
     "files.associations": {
       "*.ae": "aether"
     }
   }
   ```

3. **Use syntax highlighting:**
   - Most editors support basic syntax highlighting
   - Look for Aether language support

### Debugging Issues
**Problem:** Can't debug your code effectively.

**Solutions:**
1. **Use print statements:**
   ```aether
   func debug_function(x) {
     print("Input: " + str(x))
     result = x * 2
     print("Result: " + str(result))
     return result
   }
   ```

2. **Add error checking:**
   ```aether
   func safe_divide(a, b) {
     if b == 0 {
       print("Error: Division by zero")
       return 0
     }
     return a / b
   }
   ```

3. **Use verbose mode:**
   ```bash
   aether build --verbose
   ```

---

## 🍕 Project Structure Issues

### "Module not found" in Project
**Problem:** Can't find modules in your project.

**Solutions:**
1. **Check file structure:**
   ```
   my_project/
   ├── aether.toml
   ├── src/
   │   ├── main.ae
   │   └── utils.ae
   ```

2. **Verify import paths:**
   ```aether
   // In src/main.ae
   import utils  // For src/utils.ae
   ```

3. **Check `aether.toml` configuration:**
   ```toml
   [project]
   name = "my_project"
   version = "0.1.0"
   ```

### Dependency Issues
**Problem:** Can't resolve dependencies.

**Solutions:**
1. **Update lock file:**
   ```bash
   aether dep update
   ```

2. **Check dependency versions:**
   ```toml
   [dependencies]
   stdlib = "0.1.0"
   math = "0.1.0"
   ```

3. **Clear and rebuild:**
   ```bash
   rm aether.lock
   aether dep resolve
   ```

---

## 🍕 Common Patterns and Solutions

### Error Recovery Patterns

**When you get many errors:**
1. **Fix the first error only** - others are likely cascading
2. **Re-run the build** after each fix
3. **Don't try to fix everything at once**

**When errors don't make sense:**
1. **Check for hidden characters** (copy-paste issues)
2. **Verify file encoding** (should be UTF-8)
3. **Look for missing braces/brackets**
4. **Check indentation** (Aether is sensitive to this)

### Debugging Workflow

1. **Read the error message carefully**
2. **Look at the line and column numbers**
3. **Check the code snippet and caret position**
4. **Follow the suggested fix**
5. **Test your fix by running again**

### Prevention Strategies

1. **Start with small, working code**
2. **Test frequently** as you build
3. **Use consistent formatting**
4. **Follow naming conventions**
5. **Check the spec** when unsure

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

## 🍕 Quick Reference

### Common Error Fixes

| Error | Quick Fix |
|-------|-----------|
| `expected IN, got ILLEGAL` | Add `in` keyword: `for x in array` |
| `expected IDENT, got LPAREN` | Add parameter name: `func add(a, b)` |
| `expected IDENT, got VARARG` | Move varargs to end: `func f(a, ...b)` |
| `unexpected token in parsePrimary` | Complete the expression |
| `expected expression for return value` | Add return value or `return 0` |
| `cannot add string and number` | Convert to same type: `"hello" + str(5)` |
| `undefined variable 'x'` | Declare variable: `x = 10` |
| `undefined function 'print'` | Import stdlib: `import stdlib` |

### Build Commands

| Command | Purpose |
|---------|---------|
| `aether build` | Build the project |
| `aether build --verbose` | Show detailed output |
| `aether build --analyze-only` | Check for errors without compiling |
| `aether build --max-errors=3` | Limit error output |
| `aether dep update` | Update dependencies |
| `aether clean` | Clean build artifacts |

---

**Remember:** Every problem has a solution! Don't get frustrated—get curious! 🍕✨

Happy troubleshooting, pizza lovers! 🍕 