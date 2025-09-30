---
alwaysApply: true
---

# Go Exercises for Frontend Developers

This is a progressive set of Go exercises designed for developers with frontend experience (JavaScript/TypeScript). The exercises focus on Go-specific concepts and idioms that differ from frontend development.

## Structure

Each topic has at least two files:

- `N_topic.go` - Contains function/type skeletons to implement
- `tests/N_topic_test.go` - Contains tests that verify your implementation
- `N_topic_subtopic.go` - Contains subtopic function/type skeletons to implement
- `tests/N_topic_subtopic_test.go` - Contains subtopic tests that verify your implementation

## Running Tests

```bash
make check
```

This will run all tests. Initially, they will fail - your job is to implement the functions to make them pass!

## Topics

### 1. Basics (`1_basics.go`)

**New concepts for frontend devs:**

- Explicit error handling (no try/catch)
- Multiple return values
- "Comma ok" idiom for optional values
- Goroutines and channels (async without Promises)
- Reference vs value semantics

**Key differences from JS/TS:**

- No `undefined` or `null` - use zero values and comma-ok idiom
- No exceptions - errors are values
- Goroutines ≠ Promises (different concurrency model)

### 2. Structs & Methods (`2_structs_methods.go`)

**New concepts for frontend devs:**

- Struct methods (not class methods!)
- Pointer vs value receivers
- No inheritance - composition instead
- Generics with constraints
- Working with `any` type (like `unknown` in TS)

**Key differences from JS/TS:**

- No classes - structs + methods
- Explicit pointer semantics
- No `this` keyword - receiver is explicit

### 3. Interfaces

- Duck typing (implicit implementation)
- Empty interface and type assertions
- Interface composition
- Common patterns: io.Reader, io.Writer

### 4. Concurrency Patterns (Coming Soon)

- Worker pools
- Fan-in/fan-out
- Context and cancellation
- Select statement
- WaitGroup patterns (including `.Go()` in Go 1.25+)

### 5. Iterators (Coming Soon)

- `iter.Seq` and `iter.Seq2`
- Pull vs push iteration
- Custom iterators
- Range over functions (Go 1.23+)

### 6. HTTP & JSON (Coming Soon)

- HTTP handlers and middleware
- JSON encoding/decoding
- Context for request handling
- Common web patterns

## Tips

1. **Read the tests first** - they show exactly what's expected
2. **Error handling** - Go doesn't use exceptions. Return errors explicitly
3. **No magic** - Go is explicit. No implicit type coercion, no automatic async
4. **The zero value** - Every type has a zero value (0, "", nil, false, etc.)
5. **Pointers matter** - Unlike JS objects, Go structs are copied by value

## Go Version

This project uses **Go 1.25.1** with modern features:

- Generics (1.18+)
- Iterators - `iter.Seq` (1.23+)
- `sync.WaitGroup.Go()` method (1.25+)
- Range over integers (1.22+)
