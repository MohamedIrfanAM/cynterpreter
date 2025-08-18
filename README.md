# Cynterpreter

A C interpreter written in Go that provides both interactive REPL (Read-Eval-Print Loop) mode and batch file execution capabilities.

Cynterpreter is an educational C interpreter implementation that demonstrates the core components of a programming language interpreter: lexical analysis, parsing, and evaluation. Built from scratch in Go, it follows a tree-walking interpreter architecture and supports a subset of the C programming language. The parser uses top down **recursive descent parsing** for statements and control structures, combined with **Pratt parsing** (operator precedence parsing) for expressions, providing clean and maintainable code while handling complex expression precedence correctly.


## Features

### Core Language Support

- **Data Types**: `int`, `float`, `char`, `bool`, `string`
- **Variable Declaration and Assignment**: Type-safe variable declarations with optional initialization
- **Arrays**: Static array declarations with literal initialization and index-based access
- **Functions**: Function declarations, parameters, return values, and function calls
- **Control Flow**: `if-else` statements, `while` loops, `for` loops
- **Operators**: 
  - Arithmetic: `+`, `-`, `*`, `/`, `%`
  - Comparison: `==`, `!=`, `<`, `<=`, `>`, `>=`
  - Logical: `&&`, `||`, `!`
  - Assignment: `=`, `+=`, `-=`, `*=`, `/=`, `%=`
  - Unary: `+`, `-` (prefix)

### Built-in Functions

- `print()` - Print values to stdout
- `printf()` - Formatted printing with format strings
- `input()` - Read input from stdin with optional prompt

### Execution Modes

- **REPL Mode**: Interactive command-line interface for executing C statements
- **Batch Mode**: Execute C programs from files with automatic `main()` function invocation

### Language Features

- **Type Safety**: Runtime type checking for variables and function parameters
- **Scope Management**: Proper variable scoping with nested environments
- **Error Handling**: Comprehensive error reporting for parsing and runtime errors
- **Expression Evaluation**: Support for complex nested expressions with proper operator precedence
- **Automatic Garbage Collection**: Memory management handled automatically by Go's runtime GC
- **Memory Safety**: No buffer overflows or dangling pointer issues due to Go's memory model
- **Unicode Support**: Full UTF-8 string handling
- **Stack Overflow Protection**: Automatic stack management and overflow detection
- **Bounds Checking**: Array and string access bounds checking to prevent memory corruption

## Missing Features

The following C language features are not yet implemented:
- **Pointers**: No pointer support (`*`, `&` operators)
- **Structs/Unions**: No composite data types
- **Preprocessor**: No `#include`, `#define`, or other preprocessor directives
- **Switch Statements**: No `switch-case` support
- **Multiple File Support**: Single file compilation only
- **Dynamic Memory**: No `malloc`/`free` support
- **Standard Library**: Limited built-in functions
- **Type Modifiers**: No `const`, `volatile`, `static`, etc.
- **Advanced Numeric Types**: No `long`, `double`, `unsigned` variants
- **Bit Fields**: No bit manipulation structures

## Installation and Usage

### Prerequisites
- Go 1.19 or later

### Building
```bash
go build -o cynterpreter main.go
```

### REPL Mode
```bash
./cynterpreter
```

### File Mode
```bash
./cynterpreter program.c
```

## Example Programs

### Interactive Example (REPL)
```c
>> int x = 42;
>> float y = 3.14;
>> bool flag = true;
>> printf("x = %d, y = %.2f, flag = %t\n", x, y, flag);
x = 42, y = 3.14, flag = true
>> if (x > 40){ 
>>>     print("x is large");
>>> }
x is large
>> for (int i = 0; i < 3; i = i + 1) {
>>>       printf("i = %d\n", i);
>>> }
i = 0
i = 1
i = 2
>> x + y * 2;
48.28
>> x > 30 && flag;
true
>> x / 10;
4
>> "Hello " + "World";
Hello World
>> x;
42
>> string text = "hello";
>> text[0];
h
>> text[1];
e
>> text[4];
o
>> int i = 2;
>> text[i*2];
o
```

### Hello World
```c
// hello.c
int main() {
    print("Hello, World!");
    return 0;
}
```

### Variables and Arithmetic
```c
// variables.c
int main() {
    int x = 10;
    int y = 20;
    int sum = x + y;
    printf("Sum: %d\n", sum);
    
    float pi = 3.14;
    float area = pi * 5 * 5;
    printf("Area: %.2f\n", area);
    
    return 0;
}
```

### Control Flow
```c
// control.c
int main() {
    int num = 5;
    
    // If-else statement
    if (num > 0) {
        print("Positive number");
    } else {
        print("Non-positive number");
    }
    
    // While loop
    int i = 0;
    while (i < 3) {
        printf("Count: %d\n", i);
        i = i + 1;
    }
    
    // For loop
    for (int j = 0; j < 3; j = j + 1) {
        printf("Loop: %d\n", j);
    }
    
    return 0;
}
```

### Functions
```c
// functions.c
int add(int a, int b) {
    return a + b;
}

int factorial(int n) {
    if (n <= 1) {
        return 1;
    }
    return n * factorial(n - 1);
}

int main() {
    int result = add(5, 3);
    printf("5 + 3 = %d\n", result);
    
    int fact = factorial(5);
    printf("5! = %d\n", fact);
    
    return 0;
}
```

### Arrays
```c
// arrays.c
int main() {
    // Array declaration and initialization
    int numbers[5] = {1, 2, 3, 4, 5};
    
    // Array access
    printf("First element: %d\n", numbers[0]);
    printf("Last element: %d\n", numbers[4]);
    
    // Array assignment
    numbers[2] = 10;
    printf("Modified element: %d\n", numbers[2]);
    
    return 0;
}
```



### String Operations
```c
// strings.c
int main() {
    string name = "Alice";
    string greeting = "Hello, ";
    
    printf("%s%s!\n", greeting, name);
    
    char initial = 'A';
    printf("Initial: %c\n", initial);
    
    return 0;
}
```

## Project Structure

```
cynterpreter/
├── main.go              # Entry point
├── lexer/               # Lexical analysis
│   ├── lexer.go
│   ├── next_token_test.go
│   └── token/
│       ├── token.go
│       └── map.go
├── parser/              # Syntax analysis
│   ├── parser.go
│   ├── expressions.go
│   ├── statements.go
│   ├── *_test.go
│   └── ast/
│       ├── ast.go
│       ├── expressions.go
│       └── statements.go
├── eval/                # Evaluation engine
│   ├── eval.go
│   ├── expressions.go
│   ├── statements.go
│   ├── builtin.go
│   ├── *_test.go
│   └── obj/
│       ├── obj.go
│       └── env.go
├── repl/                # Interactive mode
│   └── repl.go
└── batch/               # File execution mode
    └── batch.go
```
