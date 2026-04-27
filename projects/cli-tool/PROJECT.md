
# CLI TODO Tool — Project Overview

A compact, interactive terminal TODO application built in **Go** to strengthen practical backend engineering skills through real project structure, storage abstraction, testing, and maintainable architecture.

---

# Overview

This project is designed as both a usable productivity tool and a structured Go learning environment.

Users run a compiled binary and interact through natural commands such as:

```bash
todo add "Buy milk"
todo list
todo done 3
```

Tasks are persisted locally using either:

* **JSON files** (default)
* **SQLite** (optional)

The application emphasizes:

* Clean package separation
* Interface-driven design
* Multiple persistence layers
* Testing core workflows
* Practical CLI engineering

---

# Tech Stack

| Area                | Choice                                 |
| ------------------- | -------------------------------------- |
| Language            | Go 1.26+                               |
| Interactive Prompts | `github.com/manifoldco/promptui`       |
| Default Storage     | JSON (`todos.json`,`setting.json`)     |
| Optional Database   | SQLite (`github.com/mattn/go-sqlite3`) |

---

# Project Structure

```txt
projects/cli-tool/
│
├── cmd/
│   └── main.go                # Entry point, flag parsing, app bootstrap
│
├── internal/
│   ├── app.go                 # Main REPL loop
│   ├── commands.go            # Command routing
│   └── todo_actions.go        # Core TODO operations
│
├── repository/
│   ├── store.go               # TodoStore interface
│   ├── json_store.go          # JSON backend
│   ├── sqlite_store.go        # SQLite backend
│   └── paths.go               # Data directory configuration
│
├── models/
│   ├── todo.go                # Todo entity
│   └── settings.go            # Validation settings
│
├── utils/
│   ├── parser.go              # Input parsing
│   ├── display.go             # Formatting output
│   ├── validation.go          # Title validation
│   └── prompt.go              # Prompt helpers
│
├── db/                        # Runtime storage
│
├── Makefile
└── go.mod
```

---

# Core Features

## Interactive REPL

* Continuous command loop
* Exit with:
  * `q`
  * `quit`

---

## Supported Commands

| Command                      | Description          |
| ---------------------------- | -------------------- |
| `todo add "task"`            | Add a new TODO       |
| `todo list`                  | List all TODOs       |
| `todo list --filter=pending` | Show pending tasks   |
| `todo list --filter=done`    | Show completed tasks |
| `todo done [id]`             | Mark task complete   |
| `todo delete [id]`           | Delete task          |
| `todo edit [id]`             | Edit task            |
| `todo help`                  | Help menu            |

---

## Validation System

* Max title length configurable
* Rules stored in `setting.json`
* User-friendly validation messages

---

## UX Enhancements

* Human-readable relative dates
* Interactive prompts when IDs are omitted
* Cleaner task displays

---

## Storage Backends

### JSON Mode (Default)

```txt
db/
├── todos.json
└── setting.json
```

---

### SQLite Mode (`-db` flag)

```txt
db/
└── todo.db
```

---

# Build & Run

## Build

```bash
cd projects/cli-tool
go build -o cli-tool ./cmd
```

---

## Run (JSON)

```bash
./cli-tool
```

---

## Run (SQLite)

```bash
./cli-tool -db
```

---

## Test

```bash
go test ./...
```

---

## Makefile Targets

```bash
make build
make run
make test
make fmt
```

---

# Architecture Strengths

## Strong Learning Value

This project introduces:

* Interfaces
* Dependency abstraction
* Storage backends
* REPL systems
* File persistence
* SQL integration
* Testable design

---

## Practical Software Engineering Concepts

* Separation of concerns
* Command routing
* Configurable validation
* Extensible feature growth
* Backend portability

---

# Key Improvement Areas

## 1. Configuration Management

### Current:

Global `DataDir`

### Better:

Inject config struct:

```go
type Config struct {
    DataDir string
    UseSQLite bool
}
```

### Benefit:

* Removes hidden global state
* Improves testability
* Supports future env variables

---

## 2. Error Handling

### Current:

Some silent failures

### Better:

Centralize error handling

### Benefit:

* Easier debugging
* Cleaner user feedback
* Better production practices

---

## 3. Repository Purity

### Current:

Repository prints directly

### Better:

Return results/errors only

### Benefit:

* Cleaner architecture
* Easier testing
* UI decoupling

---

## 4. Command Parsing

### Current:

`strings.Fields`

### Problem:

Poor quote handling

### Better Options:

* Custom lexer
* `regexp`
* `cobra`
* `kong`

---

## 5. Backend Consistency

### Problem:

JSON and SQLite assign IDs differently

### Better:

Unify through interface contracts

---

## 6. CI/CD Hygiene

Recommended:

```bash
go test ./...
go vet ./...
staticcheck ./...
```

---

# High-Value Feature Ideas

| Feature           | Learning Outcome        |
| ----------------- | ----------------------- |
| Cobra/Kong CLI    | Advanced CLI design     |
| `--data-dir`      | Config systems          |
| Env vars          | 12-factor principles    |
| Due dates         | Time APIs + migrations  |
| Priority/tags     | Data modeling           |
| Undo              | Event sourcing basics   |
| Export/import     | Data workflows          |
| Integration tests | Binary validation       |
| Fuzzing           | Parser reliability      |
| HTTP API          | Service architecture    |
| SQLite migrations | Production DB practices |

---

# Recommended Learning Roadmap

## Phase 1 — Polish Core

* Config struct
* Better parsing
* Error cleanup
* Logging
* CI pipeline

---

## Phase 2 — Feature Depth

* Due dates
* Tags
* Sorting
* Non-interactive mode

---

## Phase 3 — Advanced Engineering

* Cobra migration
* HTTP API
* Migration system
* Fuzz tests
* Cross-platform support

---

# Final Assessment

## Current Status:

A strong intermediate-level Go CLI learning project.

---

## Best Parts:

* Clear architecture
* Dual persistence
* Good testing opportunities
* Real-world engineering patterns

---

## Biggest Next Wins:

* Remove globals
* Improve parser
* Add configuration
* Strengthen CI
* Expand one major feature fully

---

# Bottom Line

This project is already more than a beginner TODO app.

It serves as a practical foundation for learning:

* Go project architecture
* Interfaces
* Persistence layers
* Testing
* CLI systems
* System extensibility

With targeted improvements, it can evolve into a portfolio-worthy backend engineering project.
