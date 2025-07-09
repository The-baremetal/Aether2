# 🍕 Learning Aether: From Beginner to Pro

Welcome to the ultimate learning guide for Aether! Whether you're a complete beginner or an experienced developer, this guide will help you master Aether programming with clear examples, best practices, and lots of pizza personality! 🍕

---

## 🍕 Getting Started

### What is Aether?
Aether is a **simple, modern programming language** designed to be:
- **Easy to learn** - Clean, readable syntax
- **Fast to write** - Less boilerplate than other languages
- **Powerful** - Built on LLVM for performance
- **Fun** - Designed with developer happiness in mind

### Your First Aether Program
Let's start with the classic "Hello, World!":

```aether
import stdlib

print("Hello, World!")
```

**What's happening:**
1. `import stdlib` - Import the standard library
2. `print("Hello, World!")` - Print a message to the console

**To run it:**
```bash
aether build
./aether
```

---

## 🍕 Core Concepts

### 1. Variables and Assignment
Aether uses simple assignment - no `let`, `var`, or `const` needed!

```aether
name = "Alice"
age = 25
is_student = true
```

**Key points:**
- No semicolons needed
- Variables are dynamically typed
- Names use snake_case by convention

### 2. Functions
Functions are the building blocks of Aether programs:

```aether
func greet(name) {
  return "Hello, " + name + "!"
}

func add(a, b) {
  return a + b
}

func calculate_area(width, height) {
  return width * height
}
```

**Function rules:**
- Use `func` keyword
- Parameters in parentheses
- Return statement for output
- No semicolons

### 3. Control Flow
Aether has simple, readable control structures:

```aether
age = 18

if age >= 18 {
  print("You are an adult")
} else {
  print("You are a minor")
}

x = 5
if x > 0 {
  print("Positive")
} else if x < 0 {
  print("Negative")
} else {
  print("Zero")
}
```

### 4. Loops
Aether uses `for` loops with the `in` keyword:

```aether
numbers = [1, 2, 3, 4, 5]

for number in numbers {
  print(number)
}

for i in range(0, 10) {
  print(i)
}
```

**Loop types:**
- `for x in array` - Iterate over arrays
- `for x in range(start, end)` - Count from start to end-1

---

## 🍕 Data Types

### Numbers
Aether supports integers and floating-point numbers:

```aether
integer = 42
float = 3.14
negative = -10
```

**Operations:**
```aether
a = 10
b = 5

sum = a + b
difference = a - b
product = a * b
quotient = a / b
remainder = a % b
power = a ^ b
```

### Strings
Strings are sequences of characters:

```aether
message = "Hello, World!"
name = 'Alice'
multi_line = "This is a
multi-line string"
```

**String operations:**
```aether
first = "Hello"
second = "World"
combined = first + " " + second

length = len(combined)
```

### Booleans
Boolean values for true/false logic:

```aether
is_active = true
is_complete = false
```

**Boolean operations:**
```aether
a = true
b = false

and_result = a and b
or_result = a or b
not_result = not a
```

### Arrays
Arrays store multiple values:

```aether
numbers = [1, 2, 3, 4, 5]
names = ["Alice", "Bob", "Charlie"]
mixed = [1, "hello", true]
```

**Array operations:**
```aether
numbers = [1, 2, 3]
first = numbers[0]
length = len(numbers)

for number in numbers {
  print(number)
}
```

---

## 🍕 Functions Deep Dive

### Function Parameters
Functions can have multiple parameters:

```aether
func greet(name, age) {
  return "Hello " + name + ", you are " + str(age) + " years old"
}

func calculate_rectangle_area(width, height) {
  return width * height
}

func format_name(first, last) {
  return first + " " + last
}
```

### Return Values
Functions can return values or nothing:

```aether
func add(a, b) {
  return a + b
}

func print_greeting(name) {
  print("Hello, " + name)
  return
}
```

### Variable Arguments
Functions can accept variable numbers of arguments:

```aether
func sum(...numbers) {
  total = 0
  for number in numbers {
    total = total + number
  }
  return total
}

result = sum(1, 2, 3, 4, 5)
```

---

## 🍕 Modules and Imports

### Importing Modules
Aether uses a simple import system:

```aether
import stdlib
import math
import my_module
```

### Using Imported Functions
```aether
import stdlib
import math

print("Hello, World!")
result = math.sqrt(16)
```

### Creating Your Own Modules
Create a file called `my_module.ae`:

```aether
func greet(name) {
  return "Hello, " + name + "!"
}

func add(a, b) {
  return a + b
}
```

Then import it in another file:

```aether
import my_module

message = my_module.greet("Alice")
sum = my_module.add(5, 3)
```

---

## 🍕 Error Handling

### Understanding Error Messages
Aether provides clear, helpful error messages:

```
🍕 SyntaxError: Unexpected token at line 2, column 1: expected IN, got ILLEGAL
    for x (array) { ... }
    ^
    Fix: Use 'in' for for-loops. See: specs.md §9
```

**What to do:**
1. Read the error message carefully
2. Look at the line and column numbers
3. Check the suggested fix
4. Reference the spec if needed

### Common Mistakes and Fixes

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

**❌ Wrong:**
```aether
let x = 10;
```

**✅ Right:**
```aether
x = 10
```

**❌ Wrong:**
```aether
import "stdlib"
```

**✅ Right:**
```aether
import stdlib
```

---

## 🍕 Best Practices

### 1. Naming Conventions
- **Variables and functions:** Use snake_case
- **Constants:** Use UPPER_SNAKE_CASE
- **Be descriptive:** `user_age` is better than `ua`

```aether
user_name = "Alice"
MAX_RETRY_COUNT = 3

func calculate_user_score(user_id, game_level) {
  return user_id * game_level
}
```

### 2. Code Organization
- **Group related functions together**
- **Use meaningful function names**
- **Keep functions small and focused**

```aether
func validate_user_input(input) {
  if len(input) == 0 {
    return false
  }
  return true
}

func process_user_data(user_data) {
  if not validate_user_input(user_data) {
    return "Invalid input"
  }
  return "Processing: " + user_data
}
```

### 3. Comments and Documentation
Aether doesn't use traditional comments, but you can document your code:

```aether
func calculate_discount(price, discount_percent) {
  discount_amount = price * (discount_percent / 100)
  final_price = price - discount_amount
  return final_price
}
```

### 4. Error Prevention
- **Check for edge cases**
- **Validate inputs**
- **Use meaningful variable names**

```aether
func divide_numbers(a, b) {
  if b == 0 {
    return "Error: Division by zero"
  }
  return a / b
}
```

---

## 🍕 Advanced Topics

### 1. Recursion
Functions can call themselves:

```aether
func factorial(n) {
  if n <= 1 {
    return 1
  }
  return n * factorial(n - 1)
}

func fibonacci(n) {
  if n <= 1 {
    return n
  }
  return fibonacci(n - 1) + fibonacci(n - 2)
}
```

### 2. Higher-Order Functions
Functions that work with other functions:

```aether
func apply_operation(numbers, operation) {
  result = []
  for number in numbers {
    result.append(operation(number))
  }
  return result
}

func square(x) {
  return x * x
}

numbers = [1, 2, 3, 4, 5]
squared = apply_operation(numbers, square)
```

### 3. Data Structures
Build complex data structures:

```aether
func create_person(name, age, city) {
  return {
    "name": name,
    "age": age,
    "city": city
  }
}

alice = create_person("Alice", 25, "New York")
print(alice["name"])
```

---

## 🍕 Debugging Techniques

### 1. Print Debugging
Use print statements to understand your code:

```aether
func complex_calculation(a, b, c) {
  print("Starting calculation with: " + str(a) + ", " + str(b) + ", " + str(c))
  
  step1 = a * b
  print("Step 1 result: " + str(step1))
  
  step2 = step1 + c
  print("Step 2 result: " + str(step2))
  
  return step2
}
```

### 2. Step-by-Step Testing
Test your functions with simple inputs:

```aether
func test_add_function() {
  result1 = add(2, 3)
  print("add(2, 3) = " + str(result1))
  
  result2 = add(-1, 5)
  print("add(-1, 5) = " + str(result2))
  
  result3 = add(0, 0)
  print("add(0, 0) = " + str(result3))
}
```

### 3. Understanding Error Messages
When you get an error:

1. **Read the error type** (SyntaxError, TypeError, etc.)
2. **Check the line and column** numbers
3. **Look at the code snippet** and caret position
4. **Follow the suggested fix**

---

## 🍕 Project Structure

### Simple Project
```
my_project/
├── aether.toml
├── src/
│   ├── main.ae
│   └── utils.ae
└── README.md
```

### Complex Project
```
my_project/
├── aether.toml
├── aether.lock
├── src/
│   ├── main.ae
│   ├── models/
│   │   ├── user.ae
│   │   └── product.ae
│   ├── utils/
│   │   ├── math.ae
│   │   └── string.ae
│   └── tests/
│       ├── user_test.ae
│       └── math_test.ae
├── docs/
│   ├── api.md
│   └── examples.md
└── README.md
```

---

## 🍕 Learning Path

### Beginner (Week 1-2)
1. **Learn basic syntax** - variables, functions, control flow
2. **Write simple programs** - calculator, greeting functions
3. **Understand error messages** - read the debugging guide
4. **Practice with loops** - iterate over arrays and ranges

### Intermediate (Week 3-4)
1. **Master functions** - parameters, return values, recursion
2. **Work with modules** - imports, creating your own modules
3. **Build data structures** - arrays, objects, complex types
4. **Debug effectively** - use print statements and error analysis

### Advanced (Week 5+)
1. **Optimize code** - performance, readability, maintainability
2. **Build larger projects** - multiple files, complex logic
3. **Contribute to the language** - report bugs, suggest features
4. **Teach others** - share your knowledge with the community

---

## 🍕 Resources

### Official Documentation
- **[Language Specification](specs.md):** Complete language rules
- **[Syntax Guide](syntax.md):** Quick syntax reference
- **[Debugging Guide](debugging.md):** Error handling and debugging
- **[Error Codes](error-codes.md):** All error codes explained

### Examples and Tutorials
- **[Examples](../examples/):** Working code examples
- **[Patterns](patterns/):** Common programming patterns
- **[Best Practices](best-practices.md):** Code quality guidelines

### Community
- **GitHub Issues:** Report bugs and request features
- **Discussions:** Ask questions and share knowledge
- **Contributing:** Help improve the language

---

## 🍕 Next Steps

### What to Build
1. **Simple calculator** - Practice functions and operators
2. **Text adventure game** - Learn control flow and user input
3. **Data processing tool** - Work with arrays and loops
4. **Web scraper** - Use external libraries and APIs
5. **Full-stack application** - Build a complete system

### Learning Resources
- **Practice coding daily** - Even 30 minutes helps
- **Read other people's code** - Learn from examples
- **Build projects** - Apply what you learn
- **Join the community** - Share and learn from others

---

**Remember:** Learning to program is a journey, not a destination! Take your time, practice regularly, and don't be afraid to make mistakes. Every error is a learning opportunity! 🍕✨

Happy coding, pizza lovers! 🍕 