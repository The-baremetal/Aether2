# 🍕 Aether Language Specification

---

## 1. Implicit Borrowing (Deep Dish Edition)

Aether’s implicit borrowing system is like a pizza chef who slices and serves your data perfectly every time—no burnt crust, no dropped toppings, and no pineapple unless you ask for it (and even then, the chef will make sure it’s safe).

### 🍕 How the Compiler Tracks References
- Every variable in Aether is tracked by the compiler with a unique ID (like a pizza order number).
- When you assign a variable or pass it to a function, the compiler records whether it’s a reference (borrow) or a copy (fresh slice).
- The compiler maintains a reference table for every scope, tracking who’s holding a slice of your data.

### 🍕 Mutation and Shared/Exclusive Borrows
- If a function only reads from a value, it gets a **shared borrow** (many can read, none can write).
- If a function wants to mutate a value, it must have an **exclusive borrow** (only one can write, nobody else can read or write until it’s done).
- The compiler enforces this at compile time—if you try to mutate while someone else is reading, you get a friendly error.

#### Example
```aether
x = [1, 2, 3]
func read(arr) {
  print(arr[0])
}
func write(arr) {
  arr[0] = 99
}
read(x)   // ok, shared borrow
write(x)  // ok, exclusive borrow
read(x)   // ok again
```
If you try to do this:
```aether
read(x)
write(x) // error: cannot borrow x as mutable while it is also borrowed as shared
```
You get a clear error: “Cannot mutate x while it’s being read!”

### 🍕 Aliasing Prevention
- The compiler prevents two variables from holding exclusive borrows to the same data at the same time.
- If you try to create an alias that would break safety, the compiler stops you.

#### Example
```aether
x = [1, 2, 3]
y = x // y is a reference to x
write(x) // ok
write(y) // error: y and x are the same, only one can mutate at a time
```

### 🍕 Lifetime Management
- The compiler tracks how long each reference lives (like a pizza timer).
- When a reference goes out of scope, the compiler knows it’s safe to let others borrow or mutate again.
- No manual lifetimes, no annotations, just automatic, delicious safety.

### 🍕 Closures and Borrowing
- Closures (blocks that capture outer variables) borrow those variables implicitly.
- The compiler tracks which variables are captured and ensures they’re not mutated unsafely.

#### Example
```aether
x = 10
myLambda = {
  print(x) // borrows x from outer scope
}
myLambda() // ok
```
If you try to mutate x in two closures at once, the compiler will stop you.

### 🍕 Arrays, Structs, and Nested Data
- Borrowing works recursively: if you borrow an array, you borrow all its elements.
- If you borrow a struct, you borrow all its fields.
- The compiler prevents partial mutation if it would break safety.

#### Example
```aether
struct Point { x: int, y: int }
p = Point { x: 1, y: 2 }
func move(point) {
  point.x = point.x + 1
}
move(p) // ok, exclusive borrow of p
```

### 🍕 Function Calls and Borrowing
- When you pass a variable to a function, the compiler decides if it’s a shared or exclusive borrow based on the function’s actions.
- If the function only reads, it’s shared. If it writes, it’s exclusive.
- If you want to force a copy, you can use a built-in `copy()` function (not required for safety, just for explicit duplication).

### 🍕 Error Reporting
- If you break the rules, the compiler gives you a clear, friendly error:
  - “Cannot mutate x while it’s being read!”
  - “Cannot borrow y as mutable more than once at a time!”
  - “Reference to z outlives its scope!”
- No cryptic messages, just pizza chef advice.

### 🍕 Why This Rocks
- You get all the safety of Rust, but with zero borrow syntax and zero headaches.
- You can write fast, safe, concurrent code without ever thinking about lifetimes or references.
- The compiler is your pizza sous-chef, always making sure your slices are safe and tasty.

---

## 2. Concurrency Roadmap: Coroutines First, Actors Next

Aether’s concurrency is like a pizza oven that gets hotter and fancier over time!

### 🍕 Phase 1: Coroutines (Powered by LLVM)
- Aether uses LLVM’s built-in coroutine support for async and lightweight concurrency.
- Coroutines are perfect for async I/O, generators, and pipelining tasks.
- Syntax:
  ```aether
  func myCoroutine {
    print("start")
    yield
    print("resume")
  }
  ```
- You can create, pause, and resume coroutines for cooperative multitasking.
- No threads, no message passing (yet)—just pure, simple async flavor.

### 🍕 Phase 2: Actor Runtime (Coming Soon!)
- Aether will add a full actor runtime for safe, scalable, message-passing concurrency.
- Each actor is a lightweight process with its own state and mailbox.
- Syntax:
  ```aether
  pid = spawn {
    while true {
      msg = receive()
      print("got message:", msg)
    }
  }
  send(pid, "hello pizza!")
  ```
- Actors can run their own coroutines for internal async tasks.
- Message passing between actors means no data races, no shared state, and no pineapple on pizza.

### 🍕 Transition Plan
- Coroutines are the default for now—easy, async, and fast.
- When the actor runtime lands, coroutines will still be useful inside actors.
- The transition will be smooth, and your code will stay delicious!

---

## 3. Error Handling: Clang-Style, Pizza Energy

Aether’s error messages are inspired by clang: clear, precise, and a little bit silly! You’ll always know what went wrong, where, and why—with a dash of pizza personality.

### 🍕 Example: Semicolon Error

```aether
x = 10;
```

Aether output:

```
🍕 SyntaxError: Unexpected `;` at line 1
    print("Hello world");
                        ^
    Fix: Remove the semicolon
    Do you want to apply the codemod to remove the semicolon? (y/n)
```

### 🍕 Error Message Features
- **Emoji and silly energy:** Every error comes with a friendly, human explanation.
- **Precise location:** Shows the line and character where the error happened.
- **Clear type:** SyntaxError, TypeError, etc., just like clang.
- **Helpful hints:** If possible, Aether will suggest how to fix your code (and maybe roast you a little).
- **Automatic resolutions:** If the error is easy to solve, the language can suggest a code mod or allow you to fix the error on the spot to avoid recompiling from a simple mistake.

### 🍕 More Examples

```aether
func add(a, b) {
  return a +
}
```

Aether output:
```output
🍕 SyntaxError: Unexpected end of input at line 2
    let result = (8*31)+
                        ^
    Looks like you forgot something after the operator!
    Do you want to fix this right now?
    (y/n)
    (your preferred editor opens to the file)
    Once you save the file, the language will check again and validate it
```

---

## Incremental Compilation and Smart Build Features

Aether is designed for fast, modern development workflows. The compiler supports:

### Incremental Compilation
- Only files and modules that have changed are recompiled.
- Unchanged code is reused from previous builds, dramatically reducing build times for large projects.

### Smart Dependency Tracking
- The compiler automatically tracks dependencies between files, modules, and packages.
- When a file changes, only its dependents are rebuilt.
- This ensures minimal recompilation and fast feedback.

### Hot Module Loading (coming soon)
- The runtime will support hot module loading for rapid development and live updates.
- Modules can be reloaded without restarting the whole program.
- Useful for REPLs, servers, and interactive development.

### Persistent Compiler State
- The compiler maintains a persistent state between builds.
- Caches parsed ASTs, type information, and intermediate representations.
- Further reduces build times and enables advanced tooling (like IDE integration and refactoring support).

---