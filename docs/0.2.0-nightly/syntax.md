# 🍕 Aether Language Syntax Guide

Welcome to Aether! The language that stays fresh, simple, and delicious. No crusty old syntax, just pure code vibes. Here’s everything you need to know, all in one place!

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

- `import math` brings all functions from math into the global namespace. You can use `plus()` or `sqrt()` directly.
- `import math as math` brings all functions from math into the `math` namespace. You must use `math.plus()` or `math.sqrt()`.
- Imports are not linked immediately. The compiler generates separate object files for each import, and the linker (`ld`) only links them when you actually use the functions.
- No global pollution when using `as`.

```aether
import math
x = plus(2, 3)
y = sqrt(16)

import math as math
z = math.plus(2, 3)
w = math.sqrt(16)
```

---

## 3. Arrays Everywhere

- All data is an array under the hood.
- Lists, strings, structs, everything can be treated as an array.

```aether
numbers = [1, 2, 3, 4, 5]
print(numbers[2])

matrix = [
  [1, 2],
  [3, 4]
]
print(matrix[1][0])

users = [
  { name: "bob", age: 20 },
  { name: "alice", age: 22 }
]
print(users[0].name)
```

---

## 4. Concatenation and Varargs

- `..` is used to concatenate arrays or strings.
- `...` is used for varargs (variable number of arguments in functions).
- Varargs can be in any position in the parameter list!

```aether
// Concatenation
hello = ["h", "e", "l", "l", "o"]
world = ["w", "o", "r", "l", "d"]
greeting = hello .. world // ["h", "e", "l", "l", "o", "w", "o", "r", "l", "d"]

str1 = "foo"
str2 = "bar"
result = str1 .. str2 // "foobar"

// Varargs in any position
func printAll(...args) {
  repeat(args.length) {
    print(args[_])
  }
}

func bind(funcName, returnType, ...paramTypes) {
  return func(...args) {
    return c_call(funcName, returnType, paramTypes, args)
  }
}

printAll(1, 2, 3, "pizza")
```

---

## 5. Functions

```aether
func greet {
  print("hello world")
}

func greet(name) {
  print("hello", name)
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

Aether supports **tuple/multi-variable assignments**:

```aether
result, err = do_something()
```

You can assign multiple variables at once, making error handling and other patterns easy and expressive!

---

## 7. Conditionals

```aether
if x > y {
  print("x bigger")
} else if x == y {
  print("they are equal!")
} else {
  print("y bigger")
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
    print("zero")
  }
  case 1 {
    print("one")
  }
  case [a, b] {
    print("array with two elements")
  }
  case {name} {
    print("struct with name field")
  }
  case _ {
    print("something else")
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
  print("topping:", item)
}

// Loop with index and value
for i, topping in ["cheese", "pepperoni", "mushroom"] {
  print("slice", i, "has", topping)
}
```

- Use the `in` keyword to loop over arrays, strings, or anything iterable.
- You can get just the value, or both index and value (like a true pizza chef).
- No 'range' keyword needed—just pure, simple, delicious 'in'!

### Repeat Loop

```aether
repeat(5) {
  print("looping!")
}
```

### While Loop

```aether
while x < 10 {
  print("count", x)
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
  print(i)
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

Aether uses explicit, Go-style error handling! No more try/catch/finally—just assignments and if-statements. Tuple assignments make this even easier:

```aether
result, err = risky_stuff()
if err != nil {
  print("oops!", err)
}
```

If your function returns a value and an error, check the error before using the value:

```aether
value, err = get_pizza_slice()
if err != nil {
  print("no pizza for you!", err)
} else {
  print("yum!", value)
}
```

To return an error from a function:

```aether
func get_pizza_slice() {
  if pizza_box_empty() {
    return nil, error("no pizza left!")
  }
  return "🍕", nil
}
```

Create custom errors with `error("message")`. No exceptions, no magic, just pure pizza logic!

---

## 13. Lambdas / Anonymous Functions

```aether
myLambda = {
  print("I am a lambda!")
}

doSomething {
  print("inline lambda!")
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
  print("x is ten")
}
if y != 5 {
  print("y is not five")
}
if a <= b {
  print("a is less than or equal to b")
}
z = (x + y) * 2
w = x ^ 3
```

---

## 22. Whitespace

Indentation is for humans. The compiler only cares about blocks and dots for chaining.

---

## 23. Standard Library

Built-ins like `print` are always available. More to come!

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
  print(x + 5)
}
myLambda()

nums = [1, 2, 3]
nums.map({
  print("pizza!")
})
```

---

## 27. No GOTO (Ever!)

- GOTO is not allowed in Aether. No jumping around, no spaghetti code, no pineapple on pizza!
- Use structured control flow: `if`, `else`, `while`, `repeat`, `break`, `continue`, `return`.
- LLVM will squeeze all the performance out of your code, so you never need GOTO anyway. 🍋
- Your code stays readable, maintainable, and delicious!

---

## 🍕 Varargs and Spread Arguments

### Function Parameter Lists (Declarations)

- You can use a single vararg (e.g., `...args`) in a function parameter list.
- The vararg must be the last parameter. You cannot have parameters after a vararg.
- Only one vararg is allowed per parameter list.
- Example:

```aether
func printAll(a, b, ...args) {
  // 'a' and 'b' are required, 'args' is a list of all remaining arguments
}
```

If you try to put a vararg anywhere but the end, or use more than one, the compiler will throw a pizza at you (syntax error)!

### Function Call Argument Lists

- You can use as many spreads as you want, anywhere in the argument list.
- Each spread can be a named variable (e.g., `...args`, `...flags`) or just `...` (spread all remaining arguments, if supported).
- Example:

```aether
printAll(...args, "cheese", ...flags, "pepperoni")
```

- This will expand all the values in `args`, then add "cheese", then expand all the values in `flags`, then add "pepperoni".
- Spreads can be mixed with normal arguments in any order.

---

## 🍕 That’s it

Aether is always simple, always fresh, and always ready for new toppings. If you want to add or remove features, just clean the kitchen and start again!

---
