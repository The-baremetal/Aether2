# Aether Self-Hosting Plan

## Overview

The goal is to make Aether capable of writing its own compiler. This involves creating Aether implementations of the core compiler components while maintaining compatibility with the existing Go-based compiler.

## Current Architecture

The Aether compiler currently consists of:

1. **Lexer** - Tokenizes source code
2. **Parser** - Creates AST from tokens  
3. **Analyzer** - Type checking, dependency analysis
4. **Code Generator** - LLVM IR generation
5. **Linker** - Final binary creation

## Self-Hosting Strategy

### Phase 1: Core Data Structures

Create Aether implementations of the fundamental data structures:

```aether
// Token representation
struct Token {
  type: string
  literal: string
  line: int
  column: int
}

// AST nodes
struct Program {
  statements: []
}

struct FunctionLiteral {
  name: string
  parameters: []
  body: BlockStatement
}

struct BlockStatement {
  statements: []
}

struct ExpressionStatement {
  expression: Expression
}
```

### Phase 2: Lexer Implementation

Implement the lexer in Aether:

```aether
struct Lexer {
  input: string
  position: int
  read_position: int
  ch: string
}

func new_lexer(input) {
  l = Lexer {
    input: input
    position: 0
    read_position: 0
    ch: ""
  }
  l.read_char()
  return l
}

func (l) read_char() {
  if l.read_position >= l.input.length {
    l.ch = ""
  } else {
    l.ch = l.input[l.read_position]
  }
  l.position = l.read_position
  l.read_position = l.read_position + 1
}

func (l) next_token() {
  l.skip_whitespace()
  
  tok = Token {
    type: ""
    literal: l.ch
    line: 1
    column: l.position
  }
  
  match l.ch {
    case "=" {
      if l.peek_char() == "=" {
        l.read_char()
        tok.type = "EQ"
        tok.literal = "=="
      } else {
        tok.type = "ASSIGN"
      }
    }
    case "+" {
      tok.type = "PLUS"
    }
    case "(" {
      tok.type = "LPAREN"
    }
    case ")" {
      tok.type = "RPAREN"
    }
    case "{" {
      tok.type = "LBRACE"
    }
    case "}" {
      tok.type = "RBRACE"
    }
    case "" {
      tok.type = "EOF"
      tok.literal = ""
    }
    case _ {
      if is_letter(l.ch) {
        tok.literal = l.read_identifier()
        tok.type = lookup_ident(tok.literal)
        return tok
      } else if is_digit(l.ch) {
        tok.type = "INT"
        tok.literal = l.read_number()
        return tok
      } else {
        tok.type = "ILLEGAL"
      }
    }
  }
  
  l.read_char()
  return tok
}
```

### Phase 3: Parser Implementation

Implement the parser in Aether:

```aether
struct Parser {
  l: Lexer
  cur_token: Token
  peek_token: Token
  errors: []
}

func new_parser(l) {
  p = Parser {
    l: l
    errors: []
  }
  
  p.next_token()
  p.next_token()
  
  return p
}

func (p) next_token() {
  p.cur_token = p.peek_token
  p.peek_token = p.l.next_token()
}

func (p) parse_program() {
  program = Program {
    statements: []
  }
  
  while p.cur_token.type != "EOF" {
    stmt = p.parse_statement()
    if stmt != null {
      program.statements.push(stmt)
    }
    p.next_token()
  }
  
  return program
}

func (p) parse_statement() {
  match p.cur_token.type {
    case "LET" {
      return p.parse_let_statement()
    }
    case "RETURN" {
      return p.parse_return_statement()
    }
    case _ {
      return p.parse_expression_statement()
    }
  }
}
```

### Phase 4: AST Walking and Code Generation

Implement AST traversal and LLVM IR generation:

```aether
struct CodeGenerator {
  module: string
  functions: []
  current_function: string
}

func new_code_generator() {
  return CodeGenerator {
    module: ""
    functions: []
    current_function: ""
  }
}

func (cg) generate_ir(program) {
  cg.module = "define i32 @main() {\n"
  
  for stmt in program.statements {
    cg.generate_statement(stmt)
  }
  
  cg.module = cg.module .. "  ret i32 0\n}\n"
  return cg.module
}

func (cg) generate_statement(stmt) {
  match stmt.type {
    case "ExpressionStatement" {
      cg.generate_expression(stmt.expression)
    }
    case "ReturnStatement" {
      cg.generate_expression(stmt.return_value)
      cg.module = cg.module .. "  ret i32 %result\n"
    }
  }
}
```

### Phase 5: Integration with Go Compiler

Create a hybrid approach where Aether components can be used by the Go compiler:

```aether
// Bridge between Aether and Go
func compile_aether_source(source) {
  lexer = new_lexer(source)
  parser = new_parser(lexer)
  program = parser.parse_program()
  
  if len(parser.errors) > 0 {
    return "Compilation failed"
  }
  
  generator = new_code_generator()
  ir = generator.generate_ir(program)
  
  return ir
}
```

## Implementation Steps

1. **Create Aether data structures** for tokens, AST nodes, and compiler state
2. **Implement lexer** in pure Aether using string operations and arrays
3. **Implement parser** using recursive descent with Aether functions
4. **Create code generator** that outputs LLVM IR as strings
5. **Build integration layer** to connect Aether components with Go compiler
6. **Test self-hosting** by compiling Aether code with Aether components

## Benefits

- **Self-hosting capability** - Aether can compile itself
- **Language evolution** - Easier to add new features
- **Educational value** - Great learning resource
- **Performance potential** - Aether-optimized compiler
- **Pride** - Complete language independence

## Timeline

- **Phase 1-2**: 2-3 months (Data structures and lexer)
- **Phase 3**: 2-3 months (Parser implementation)
- **Phase 4**: 3-4 months (Code generation)
- **Phase 5**: 1-2 months (Integration and testing)

**Total**: 8-12 months to achieve self-hosting capability. 