# SteinsLang Ψ

A small, interpreted programming language written in Go, inspired by
*Writing an Interpreter in Go* by Thorsten Ball.

SteinsLang includes a lexer, Pratt parser, abstract syntax tree (AST),
tree-walking evaluator, environment-based variable storage, functions,
conditionals, strings, arrays, and an interactive REPL.

The `.Steins` extension is used for SteinsLang source files.

## Features

- Lexical analysis and parsing
- Integer, boolean, and string values
- Variable declarations with `let`
- Arithmetic and comparison operators
- Prefix operators (`!`, `-`)
- Conditional expressions (`if` / `else`)
- Named and anonymous functions
- Function calls and return statements
- Array literals and integer indexing
- Built-in `len()` function for strings
- Interactive REPL
- Dockerized execution
- Automated Go tests

## Table of Contents

- [Getting Started](#getting-started)
- [Running with Docker](#running-with-docker)
- [Running with Go](#running-with-go)
- [REPL Usage](#repl-usage)
- [Language Syntax](#language-syntax)
- [Operators](#operators)
- [Built-in Functions](#built-in-functions)
- [Example Program](#example-program)
- [Project Structure](#project-structure)
- [Testing](#testing)
- [Current Limitations](#current-limitations)

## Getting Started

### Prerequisites

Choose one of the following:

- Docker Desktop or another working Docker Engine
- Go installed locally (Go 1.25 is the current Docker build target)

Clone the repository:

```bash
git clone https://github.com/Pluto2804/Interpreter.git
cd Interpreter
```

## Running with Docker

### 1. Build the image

```bash
docker build -t steinslang .
```

### 2. Start the interactive REPL

```bash
docker run --rm -it steinslang
```

Enter expressions at the `>>` prompt:

```text
>> 5 + 10
15
>> let x = 20;
>> x * 2
40
>> len("SteinsLang")
10
```

Exit using `Ctrl+D`.

### 3. Run a `.Steins` file

Mount the repository into the container:

```bash
docker run --rm -it \
  -v "$PWD:/workspace" \
  -w /workspace \
  steinslang
```

At the REPL prompt, enter the filename:

```text
>> first.Steins
```

The REPL recognizes a line ending in `.Steins` as a request to execute
that file. The file must be accessible inside the container, which is why
the volume mount is required.

### 4. Rebuild after changing the source

```bash
docker build -t steinslang .
```

Docker will rebuild the image using the source in your working directory.

## Running with Go

### Start the REPL

```bash
go run .
```

### Build a standalone executable

```bash
go build -o steinslang .
```

Run it:

```bash
./steinslang
```

### Run a source file

Start the REPL and enter the path to the source file:

```text
>> first.Steins
```

When running natively, the file path is resolved relative to the
interpreter process's working directory.

## REPL Usage

The REPL evaluates expressions interactively and keeps the environment
alive between input lines.

```text
>> let language = "SteinsLang";
>> let version = 1;
>> language
SteinsLang
>> version + 1
2
```

A variable declaration does not itself print a value. Evaluate an
expression to inspect its result.

The current REPL reads one line at a time. For multi-line constructs,
enter the complete expression on one line.

## Language Syntax

### 1. Variables

Declare variables using `let`.

```javascript
let name = "SteinsLang";
let age = 1;
let active = true;
let scores = [10, 20, 30];
```

Variable names can contain letters and underscores as supported by the
lexer. Use a declaration before referring to a variable.

### 2. Integers

SteinsLang supports integer literals and arithmetic.

```javascript
let a = 10;
let b = 3;

a + b;
a - b;
a * b;
a / b;
```

Integer division truncates the fractional part.

### 3. Booleans

Boolean literals are `true` and `false`.

```javascript
let enabled = true;
let disabled = false;

enabled == true;
disabled != true;
```

### 4. Strings

Write strings using double quotes.

```javascript
let greeting = "Hello";
let language = "SteinsLang";

greeting + ", " + language;
```

The `+` operator concatenates two strings.

The `len()` built-in returns the length of a string.

```javascript
len("hello");
```

### 5. Prefix operators

Use `!` for logical negation and `-` for integer negation.

```javascript
!true;
!false;
-10;
--5;
```

### 6. Comparison operators

```javascript
10 < 20;
20 > 10;

10 == 10;
10 != 5;

true == true;
false != true;
```

Supported comparisons are:

| Operator | Meaning |
|---|---|
| `==` | Equality |
| `!=` | Inequality |
| `<` | Less than |
| `>` | Greater than |

The implementation does not currently provide `<=` or `>=`.

### 7. Conditional expressions

Conditions must be enclosed in parentheses. Branches use braces.

```javascript
if (10 > 5) {
    "greater";
} else {
    "not greater";
};
```

A conditional can also be used as part of a larger expression.

```javascript
let age = 20;

if (age > 17) {
    "adult";
} else {
    "minor";
};
```

SteinsLang treats `false` and `null` as false-like values. Other
evaluated values are truthy in the current implementation.

### 8. Functions

Define a named function using `fn`, followed by its name, parameters,
and body.

```javascript
fn add(a, b) {
    a + b;
};

add(10, 20);
```

Functions can accept multiple parameters and be called with arguments.

```javascript
fn multiply(a, b) {
    return a * b;
};

multiply(6, 7);
```

The `return` statement exits the current function with a value.

```javascript
fn square(x) {
    return x * x;
};

square(5);
```

A function body can also return its final expression without an explicit
`return` statement.

```javascript
fn double(x) {
    x * 2;
};

double(8);
```

Anonymous functions are also represented by the parser:

```javascript
fn(x) {
    x * 2;
};
```

Anonymous function values can be evaluated, but the language does not
currently have a separate variable-assignment syntax for updating an
existing binding.

### 9. Arrays

Create arrays using square brackets. Separate elements with commas.

```javascript
let numbers = [10, 20, 30];
let mixed = [1, true, "hello"];
```

Access an element using its zero-based index.

```javascript
numbers[0];
numbers[1];
numbers[2];
```

An index outside the array's valid range evaluates to `null`.

Array indexing currently requires an integer index. The `len()` built-in
does not currently support arrays.

### 10. Return statements

Use `return` inside a function to return a result.

```javascript
fn absoluteLike(x) {
    if (x < 0) {
        return -x;
    };
    return x;
};

absoluteLike(-10);
```

A `return` statement must end with a semicolon.

### 11. Semicolons and formatting

Use semicolons after variable declarations and return statements.

```javascript
let x = 10;
return x;
```

Semicolons are optional after ordinary expression statements in the
parser, but using them consistently improves readability.

Whitespace and indentation can be used to make programs easier to read.

## Operators

| Category | Operators | Example |
|---|---|---|
| Arithmetic | `+`, `-`, `*`, `/` | `a + b` |
| Comparison | `<`, `>` | `a < b` |
| Equality | `==`, `!=` | `a == b` |
| Prefix | `!`, `-` | `!true`, `-5` |
| String concatenation | `+` | `"a" + "b"` |
| Function call | `()` | `add(1, 2)` |
| Array indexing | `[]` | `items[0]` |

The parser implements precedence so multiplication and division bind
more tightly than addition and subtraction. Parentheses can group
expressions explicitly.

```javascript
2 + 3 * 4;
(2 + 3) * 4;
```

## Built-in Functions

### `len(value)`

Returns the length of a string.

```javascript
len("");
len("hello");
len("SteinsLang");
```

The current implementation accepts exactly one argument and supports
strings only.

## Example Program

The repository includes `first.Steins`:

```javascript
let x = 10;
let y = 20;

x + y;

fn multiply(x, y) {
    x * y;
};

multiply(8, 9);
```

This demonstrates variable declarations, arithmetic, a named function,
and a function call.

You can run it using either Go or Docker:

```bash
go run .
```

Then enter:

```text
>> first.Steins
```

Or use the mounted Docker workflow described above.

## Project Structure

```text
Interpreter/
├── ast/          # AST node definitions and string representations
├── evaluator/    # Evaluation logic and built-in functions
├── lexer/        # Source-code tokenization
├── object/       # Runtime objects and environments
├── parser/       # Pratt parser and syntax analysis
├── repl/         # Interactive REPL and .Steins file loading
├── token/        # Token types and keyword definitions
├── first.Steins  # Example SteinsLang program
├── main.go       # Application entry point
├── go.mod        # Go module definition
├── Dockerfile    # Multi-stage Docker build
└── README.md     # Documentation
```

## Testing

Run the complete test suite:

```bash
go test ./...
```

Build the interpreter:

```bash
go build -o steinslang .
```

Format Go source files:

```bash
gofmt -w .
```

Run tests before submitting changes so that lexer, parser, AST, and
evaluator behavior remains consistent.

## Current Limitations

SteinsLang is an educational interpreter and is still evolving.

- No loop constructs such as `for` or `while`.
- No logical `&&` or `||` operators.
- No `<=` or `>=` comparison operators.
- No floating-point literals.
- No user-defined structs, maps, or modules.
- No dedicated `print()` built-in.
- `len()` currently supports strings, not arrays.
- Array indexing currently supports integer indices only.
- The REPL reads one line at a time; multi-line constructs must be
  entered on a single line.
- Source files are loaded by entering their `.Steins` path in the REPL,
  rather than through a dedicated command-line file argument.

These limitations describe the current implementation, not necessarily
the future direction of the language.

## Contributing

Contributions, bug reports, and suggestions are welcome.

When contributing:

1. Keep the code formatted with `gofmt`.
2. Add tests for new syntax or evaluation behavior.
3. Run `go test ./...`.
4. Update this README when language features change.

## Inspiration

SteinsLang is inspired by Thorsten Ball's
[*Writing an Interpreter in Go*](https://interpreterbook.com/).

The project uses a lexer, Pratt parser, AST, and tree-walking evaluator
to demonstrate the core components of an interpreted programming language.

---

Built with Go. El Psy Congroo. Ψ
