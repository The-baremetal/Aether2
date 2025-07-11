# 🍕 Aether Error Handling: The Go-Style Way (No Pineapple Allowed!)

## 🎯 Rationale

Aether is moving away from try/catch blocks (which made our parser cry) and embracing a new, Go-inspired error handling system! This means:
- No more exceptions flying around like rogue pizza slices.
- Errors are handled explicitly, right where they happen.
- Your code is easier to read, reason about, and debug (and less likely to get pineapple on it).

## 🚫 What’s Gone?
- `try { ... } catch (err) { ... } finally { ... }` is **deprecated**.
- No more magical error-catching blocks.

## 🍕 The New System: Explicit Error Returns

### Function Signature
Functions that can fail now return a tuple: `(result, error)`.

```aether
result, err = do_something(arg1, arg2)
if err != nil {
  // handle the error (log, return, etc.)
}
```

### Example: Old vs. New

#### Old (Deprecated)
```aether
try {
  risky_stuff()
} catch (err) {
  print("oops!", err)
}
```

#### New (Go-Style)
```aether
err = risky_stuff()
if err != nil {
  print("oops!", err)
}
```

Or, if the function returns a value:
```aether
value, err = get_pizza_slice()
if err != nil {
  print("no pizza for you!", err)
} else {
  print("yum!", value)
}
```

### Returning Errors
Use a special `error` value to indicate failure:
```aether
func get_pizza_slice() {
  if pizza_box_empty() {
    return nil, error("no pizza left!")
  }
  return "🍕", nil
}
```

### Creating Custom Errors
```aether
err = error("the oven is on fire!")
```

## 🧀 Migration Notes
- Replace all try/catch/finally blocks with explicit error checks.
- Update function signatures to return error values where appropriate.
- Use `nil` to indicate no error.

## 🦄 Why This Is Awesome
- No more hidden control flow—errors are handled right where they happen.
- No more parser tantrums.
- Your code is now Go-approved (and pineapple-free)!

## 📝 FAQ
**Q: Can I still use try/catch?**
A: Nope! The parser will throw a fit (and maybe a pizza) if you try.

**Q: What about panics?**
A: For truly catastrophic errors, you can use `panic("message")`, but don’t use it for normal error handling. That’s like burning the whole pizza because of one bad topping.

**Q: Is this more work?**
A: Maybe a little, but your code will be tastier and easier to debug!

---

**Welcome to the new era of Aether error handling! 🍕✨** 