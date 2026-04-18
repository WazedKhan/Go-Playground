# 🐹 go-playground

A personal Go workspace for learning, experimenting, and building — organized around three main tracks.

## Structure

```
go-playground/
├── projects/           # Self-directed projects and experiments
│   └── ...
├── book-practice/      # Code written while following a Go book
│   └── ...
└── learn-go-with-tests/  # Exercises from Learn Go with Tests
    └── ...
```

## Tracks

### 🔨 Projects
Mini-projects and experiments built from scratch to apply Go concepts in a practical context. Each project lives in its own directory with its own `main.go` and `go.mod`.

### 📖 Book Practice
Code written while following along with a Go book. Organized by chapter or topic. Good for reference when revisiting concepts.

### ✅ Learn Go with Tests
Exercises from [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) by Chris James — a TDD-driven approach to learning Go. Each topic has its own package with `_test.go` files driving the implementation.

## Running Code

```bash
# Run a specific program
go run projects/my-project/make run

# Run tests in a package
go test ./learn-go-with-tests/arrays-and-slices/...

# Run all tests across the repo
go test ./...
```

## Goals

- Build comfort with Go's type system, interfaces, and concurrency model
- Practice test-driven development
- Explore the standard library
- Write idiomatic Go

## Resources

- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)
- [Go by Example](https://gobyexample.com/)
- [The Go Tour](https://tour.golang.org/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Standard Library Docs](https://pkg.go.dev/std)
