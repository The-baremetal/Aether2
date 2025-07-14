Welcome to Aether! The language that stays fresh, simple, and delicious. No crusty old syntax, just pure code vibes. Here's everything you need to know, all in one place!

---

## 0. Program Structure - Minimal Typing! 🍕

Aether is designed for **minimal typing**. Write your code directly at the top level - no boilerplate needed!

```aether
// Just write your code directly!
import "fmt"
fmt.Print("🍕 Hello from Aether!")
x = 10
y = 20
fmt.Print("Sum:", x + y)
```

**Key principles:**
- ✅ **Top-level statements** - write code directly, no `func main()` required
- ✅ **Minimal boilerplate** - no unnecessary function wrappers
- ✅ **Direct execution** - your code runs immediately
- ✅ **Import what you need** - `import "fmt"` gives you `fmt.Print()`, `fmt.Printf()`, etc.

**Wrong (don't do this):**
```aether
func main() {
  fmt.Print("Hello!")
}
main()
```

**Right (do this):**
```aether
fmt.Print("Hello!")
```

---

## 0.5. Standard Library Usage 🍕

Aether provides built-in functions through imports:

```aether
// Import stdio for basic I/O
import "fmt"
fmt.Print("Hello!")           // Basic print
fmt.Printf("Formatted") // Formatted print

// Import math for calculations
import "math"
result = math.Add(5, 3)   // Addition
sqrt = math.Sqrt(16)      // Square root

// Import with alias
import "math" as m
result = m.Add(5, 3)
```

**Common stdlib functions:**
- `fmt.Print()` - Print to console
- `fmt.Printf()` - Formatted printing
- `math.Add()`, `math.Sub()`, `math.Mul()`, `math.Div()` - Basic math
- `math.Sqrt()`, `math.Pow()` - Advanced math
- `string.Len()`, `string.Concat()` - String operations

**No imports needed for basic operations:**
```aether
x = 10 + 5        // Basic arithmetic works without imports
name = "Alice"     // String literals work
arr = [1, 2, 3]   // Arrays work
```

---

## 1. Comments

```aether
// Single-line comment
/* Multi-line
   comment */
/// Doc comment
```

---

## 2. Imports & Linking

- `import "math"` brings all functions from math into the global namespace. You can use `plus()` or `sqrt()` directly.
- `import "math" as math` brings all functions from math into the `math` namespace. You must use `math.plus()` or `math.sqrt()`.
- Imports are not linked immediately. The compiler generates separate object files for each import, and the linker (`ld`) only links them when you actually use the functions.
- No global pollution when using `as`.

```aether
import "math"
x = plus(2, 3)
y = sqrt(16)

import "math" as math
z = math.plus(2, 3)
w = math.sqrt(16)
```

---

## 3. Arrays Everywhere

- All data is an array under the hood.
- Lists, strings, structs, everything can be treated as an array.

```aether
numbers = [1, 2, 3, 4, 5]
fmt.Print(numbers[2])

matrix = [
  [1, 2],
  [3, 4]
]
fmt.Print(matrix[1][0])

users = [
  { name: "bob", age: 20 },
  { name: "alice", age: 22 }
]
fmt.Print(users[0].name)
```

---

## 4. Concatenation and Varargs

- `..` is used to concatenate arrays or strings.
- `...` is used for varargs (variable number of arguments in functions).

```aether
// Concatenation
hello = ["h", "e", "l", "l", "o"]
world = ["w", "o", "r", "l", "d"]
greeting = hello .. world // ["h", "e", "l", "l", "o", "w", "o", "r", "l", "d"]

str1 = "foo"
str2 = "bar"
result = str1 .. str2 // "foobar"

// Varargs
func printAll(...args) {
  repeat(args.length) {
    fmt.Print(args[_])
  }
}

printAll(1, 2, 3, "pizza")
```

---

## 5. Functions

```aether
func greet {
  fmt.Print("hello world")
}

func greet(name) {
  fmt.Print("hello", name)
}

func add(a, b) {
  return a + b
}
```

---

## 6. Variables

```aether
x = 10
y = 20
```

---

## 7. Conditionals

```aether
if x > y {
  fmt.Print("x bigger")
} else if x == y {
  fmt.Print("they are equal!")
} else {
  fmt.Print("y bigger")
}
```

- Use `if`, `else if`, and `else` for branching logic.
- You can chain as many `else if` as you want, like stacking pizza toppings!

---

## 8. Pattern Matching

Pattern matching uses the match/case syntax:

```aether
match x {
  case 0 {
    fmt.Print("zero")
  }
  case 1 {
    fmt.Print("one")
  }
  case [a, b] {
    fmt.Print("array with two elements")
  }
  case {name} {
    fmt.Print("struct with name field")
  }
  case _ {
    fmt.Print("something else")
  }
}
```

- `match` keyword, followed by the value to match.
- Each `case` is followed by a pattern and a block.
- Patterns can be literals, arrays, structs, or `_` for wildcard.
- The first matching case is executed.

---

## 9. Loops

### For Loop (with 'in' keyword)

```aether
// Loop over values
for item in ["cheese", "pepperoni", "mushroom"] {
  fmt.Print("topping:", item)
}

// Loop with index and value
for i, topping in ["cheese", "pepperoni", "mushroom"] {
  fmt.Print("slice", i, "has", topping)
}
```

- Use the `in` keyword to loop over arrays, strings, or anything iterable.
- You can get just the value, or both index and value (like a true pizza chef).
- No 'range' keyword needed—just pure, simple, delicious 'in'!

### Repeat Loop

```aether
repeat(5) {
  fmt.Print("looping!")
}
```

### While Loop

```aether
while x < 10 {
  fmt.Print("count", x)
  x = x + 1
}
```

### Break and Continue

```aether
for i in [1, 2, 3, 4, 5, 6, 7, 8] {
  if i == 5 {
    break // Stop the loop when i is 5
  }
  if i % 2 == 0 {
    continue // Skip even numbers
  }
  fmt.Print(i)
}
```

- `break` exits the nearest loop early (like running out of pizza).
- `continue` skips to the next loop iteration (like skipping a pineapple slice).

---

## 10. Structs

```aether
struct Point {
  x: int
  y: int
}

p = Point { x: 5, y: 10 }
```

---

## 11. Chaining

```aether
pizza()
  .addCheese()
  .addPepperoni()
  .bake()
  .eat()
```

---

## 12. Error Handling

```aether
try {
  riskyStuff()
} catch (err) {
  fmt.Print("oops!", err)
} finally {
  fmt.Print("done!")
}
```

---

## 13. Lambdas / Anonymous Functions

```aether
myLambda = {
  fmt.Print("I am a lambda!")
}

doSomething {
  fmt.Print("inline lambda!")
}
```

---

## 14. Return

```aether
func double(x) {
  return x * 2
}
```

---

## 15. Types (Optional)

```aether
struct User {
  name: string
  age: int
}

func agePlusOne(user) {
  return user.age + 1
}
```

---

## 16. No Semicolons

No semicolons. Ever. Not even a little bit.

---

## 17. No Promises, No Consts

No async, no const, no depends, no bloat. Only pizza.

---

## 18. Packages

Packages are installed via a DNF-style repository with aria2. Fast, silly, and always fresh.

---

## 19. Everything is a Block

All code blocks use `{}`. No exceptions.

---

## 20. Operators

Standard math and comparison operators: `+`, `-`, `*`, `/`, `>`, `<`, `==`, `!=`, `=`, etc.

---

## 21. Comparison and Arithmetic Operators

Aether supports the following operators:

- `==`  equality
- `!=`  not equal
- `<`   less than
- `<=`  less than or equal
- `>`   greater than
- `>=`  greater than or equal
- `+`   addition
- `-`   subtraction
- `*`   multiplication
- `/`   division
- `%`   modulo
- `^`   exponentiation

### Examples

```aether
if x == 10 {
  fmt.Print("x is ten")
}
if y != 5 {
  fmt.Print("y is not five")
}
if a <= b {
  fmt.Print("a is less than or equal to b")
}
z = (x + y) * 2
w = x ^ 3
```

---

## 22. Whitespace

Indentation is for humans. The compiler only cares about blocks and dots for chaining.

---

## 23. Standard Library

Built-ins like `fmt.Print` are always available. More to come!

---

## 24. Implicit Borrowing

- Borrowing of data (references instead of copies) is handled implicitly by the compiler.
- You don’t need to write special syntax for borrowing—just use variables and pass them around!
- The compiler does the heavy lifting, so you can just vibe and not worry about memory safety.

---

## 25. Inline Assembly

- Inline assembly is **highly discouraged** in Aether code.
- If you really need it, it’s only available in the standard library (stdlib), not in user code.
- The language wants you to write safe, readable code—leave the assembly to the pizza chefs in the stdlib kitchen!

---

## 26. Lambdas Are Just Blocks (No Parameters)

- Lambdas (anonymous functions) are always written as blocks.
- No parameters, no arrows, no implicit variables, no parentheses.
- If you want to use data, capture it from the outer scope (closure style).

```aether
x = 10
myLambda = {
  fmt.Print(x + 5)
}
myLambda()

nums = [1, 2, 3]
nums.map({
  fmt.Print("pizza!")
})
```

---

## 27. No GOTO (Ever!)

- GOTO is not allowed in Aether. No jumping around, no spaghetti code, no pineapple on pizza!
- Use structured control flow: `if`, `