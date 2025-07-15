# C Bindings in Aether 0.4.0-nightly 🍕🔗

Welcome to the world of C bindings in Aether! Here you’ll learn how to use C libraries the Aether way: direct, lowercase, and as easy as eating pizza with your hands.

---

## 1. What Are Bindings?

**Bindings** let you call C functions from your Aether code. Want to use `printf`, `malloc`, or `sin`? You can! Bindings make C’s power available in your Aether programs, with zero fuss and zero pineapple.

---

## 2. How to Write Bindings (The Aether Way)

- **Direct mapping:** Use the C function name, all lowercase, no prefixes.
- **No wrappers:** Don’t wrap or rename unless you really need to.
- **Mark as foreign:** Add a `// foreign` comment above each function so the compiler knows it’s a C function.
- **Parameter names:** Match the C standard as closely as possible.

### Example: stdio.ae
```aether
// C stdio.h direct bindings for Aether
// #include <stdio.h>

// foreign
func printf(format)
// foreign
func scanf(format)
// foreign
func fopen(filename, mode)
// foreign
func fclose(stream)
```

### Example: string.ae
```aether
// C string.h direct bindings for Aether
// #include <string.h>

// foreign
func strcpy(dest, src)
// foreign
func strlen(s)
```

### Example: math.ae
```aether
// C math.h direct bindings for Aether
// #include <math.h>

// foreign
func sin(x)
// foreign
func pow(x, y)
```

### Example: stdlib.ae
```aether
// C stdlib.h direct bindings for Aether
// #include <stdlib.h>

// foreign
func malloc(size)
// foreign
func free(ptr)
```

---

## 3. How to Use Bindings

1. **Import the binding:**
   ```aether
   import "c/stdio"
   import "c/math"
   ```
2. **Call the function directly:**
   ```aether
   printf("hello, world!\n")
   x = sin(3.14)
   ```

---

## 4. Best Practices

- **Keep it lowercase:** All function names are lowercase, just like in C.
- **No prefixes:** Don’t use `c_` or any other prefix.
- **No wrappers unless needed:** Only add wrappers if you need to adapt types or add error handling.
- **Use the `// foreign` marker:** This tells the compiler to link to the C function.

---

## 5. FAQ

**Q: Can I use any C library?**  
A: If you have a header and the library is linked, yes! Just write the binding file.

**Q: What about types?**  
A: Map C types to the closest Aether types. For pointers, use `buffer` or `array` as appropriate.

**Q: Do I need to write wrappers?**  
A: Nope! Only if you want to change the interface or add extra checks.

---

Now go forth and bind C like a pro—no pineapple required! 🍕 