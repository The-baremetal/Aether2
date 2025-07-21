# Corountines in Aether

---

## 1. What are coroutines?

**Corountines** let you create lightweight, multithreaded apps

- Perfect for multithreaded jobs and seperate layers
- UI and main loops
- Game engines that require multiple layers

---

## 2. How coroutines work in aether

- Coroutines share a worker pool of cpu threads.
- Each coroutine shares a scope
- Coroutines are handled on compile-time

---

## 3. Use Cases

- **Game development:**
  - Create games with different layers like UI and main loops.
- **Game development:**
  - Applications that are heavy on CPU resources.

---

## 4. How to Enable Coroutines

Simply use the coroutines included in the standard library

```aether
import "coroutine"
import "fmt"

func uiLayer{ // don't worry about the missing parens, aether doesn't need parens when there are no arguments
    Println("Do UI stuff")
    coroutine.yield()
}

func mainLoop{
    while true {
        coroutine.resume(ui) // resume coroutine
        coroutine.yield() // yield control to ui
    }
}

ui = coroutine.create(uiLayer) // coroutines are paused on creation
main = coroutine.create(mainLoop)

coroutine.resume(main) // start main
```

---

## 5. Best Practices

- **When it should be used:**
  - Every new object with any runtime logic should be a new coroutine (to clarify, runtime logic means code)
- **Debugging:**
  - If a coroutine crashes from any stack or memory issue, it will panic with a coroutine ID.

---

## 6. FAQ

**Q: What happens if a coroutine crashes?**
A: The entire app crashes, but it is super rare and usually caused by unsafe C bindings. The panic will include a coroutine ID.

**Q: Are coroutines lightweight?**
A: Coroutines are resolved on compile time, they don't come with runtime overhead.

**Q: Do coroutines add bloat?**
A: No. They do not add runtime bloat to your app, only a tiny binary increase.

**Q: What are Coroutine IDs?**
A: Coroutine IDs are incrementing numbers that identify a coroutine for debugging purposes.

---
