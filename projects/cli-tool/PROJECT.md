# CLI TODO tool — project overview

## What it is

A small **interactive terminal TODO application** written in Go. You run a binary, type natural commands like `todo add "Buy milk"`, and data is persisted under a configurable data directory (default `./db`).

It is structured as a learning playground: clear package boundaries, two persistence backends, and tests around core behavior.

## Tech stack

| Area | Choice |
|------|--------|
| Language | Go 1.26+ (`go.mod`) |
| Interactive prompts | `github.com/manifoldco/promptui` |
| Optional database | SQLite via `github.com/mattn/go-sqlite3` |
| Default storage | JSON files (`todos.json`, `setting.json`) |

## Package layout

| Package | Role |
|---------|------|
| `cmd` | `main`: parses `-db` flag, optionally wires SQLite store, starts `internal.AppLoop`. |
| `internal` | Application loop, command routing (`HandleCommands`), todo use-cases (list, add, edit, delete, mark done, filters). |
| `repository` | `TodoStore` interface, JSON implementation, SQLite implementation, paths, read/write helpers. |
| `models` | `Todos`, `Setting`, status constants. |
| `utils` | Help text, quoted-title parsing, display formatting, validation, prompt UI. |

## Current features

- **REPL-style loop**: prompts until `q` / `quit`.
- **Commands** (prefix `todo`): `add`, `list` (with optional `--filter=pending|done`), `done`, `delete`, `edit`, help variants.
- **Title rules**: validation and max length driven by `setting.json`.
- **Relative dates**: `createdAt` shown in a friendlier form in the list.
- **Dual backend**: default JSON store; `-db` uses SQLite (`db/todo.db` under the same data dir concept).
- **Interactive flows**: `promptui` for delete/done when no ID is passed.

## Data layout (default)

- `repository.DataDir` defaults to `./db` (see `repository/paths.go`).
- JSON mode: `todos.json`, `setting.json` in that directory.
- SQLite mode: `todo.db` in that directory (when using `-db`).

## Build and test

```bash
cd projects/cli-tool
go build -o cli-tool ./cmd
./cli-tool              # JSON backend
./cli-tool -db          # SQLite backend
go test ./...
```

The `Makefile` also defines `build`, `run`, `test`, and formatting targets.

---

## What to improve (robustness and code quality)

These are ordered roughly by impact vs. effort for a learning project.

1. **Configuration instead of only a global `DataDir`**  
   Today tests swap `repository.DataDir`. A small config struct (or `XDG_CONFIG_HOME`-style paths) passed into the app would remove hidden global state and make behavior obvious at the call site.

2. **Consistent error handling in `HandleCommands`**  
   Some branches ignore errors from helpers (for example list/filter paths). Surfacing errors to the user (or a single `log`/`slog` path) avoids silent failures.

3. **Replace ad-hoc `fmt.Println` in repository with structured logging**  
   Side effects in the store make testing noisy and couple persistence to UI. Prefer returning errors / result types and letting `internal` or `cmd` print user-facing messages.

4. **JSON vs SQLite parity**  
   SQLite `AddTodo` uses auto-increment IDs; JSON assigns IDs in application code. Document the differences or align behavior (e.g. always return assigned ID from the store interface).

5. **Command parsing**  
   `strings.Fields` splits quoted phrases awkwardly. For learning, try a tiny lexer, `regexp`, or a library (`cobra` / `kong`) so `todo add "multi word title"` is parsed reliably without edge-case bugs.

6. **Concurrency and globals**  
   Package-level `activeStore` and `DataDir` are fine for a CLI process, but document that they are not goroutine-safe if you ever add background work or HTTP mode.

7. **Dependencies**  
   `go-sqlite3` is CGO; builds need a C toolchain. For portability experiments, consider `modernc.org/sqlite` (pure Go) as an alternative.

8. **CI and hygiene**  
   Add a minimal GitHub Action (or similar): `go test ./...`, `go vet ./...`, and optionally `staticcheck` or `golangci-lint`.

9. **Documentation in-repo**  
   This file plus a short “Contributing / layout” section keeps onboarding cheap as the project grows.

---

## Feature ideas (learning value)

Pick a few; each teaches different Go or systems skills.

| Idea | What you learn |
|------|----------------|
| **Subcommands with Cobra or Kong** | CLI design, flags, shell completion, testing CLIs. |
| **`--data-dir` / `TODO_DATA_DIR` env** | Configuration, `flag` package, 12-factor style defaults. |
| **Non-interactive mode** | Same binary usable in scripts: `todo add "x" && todo list` then exit (no REPL). |
| **Due dates and sorting** | Time APIs, schema migration (JSON + SQL), list ordering. |
| **Priority or tags** | Modeling, filtering, SQLite indexes. |
| **Undo last change** | Small command stack or event log; careful JSON rewrite. |
| **Export / import** | `encoding/json` streaming, backup workflows. |
| **Integration tests** | Build binary with `testing`, run with `os/exec`, assert on files/DB. |
| **Fuzzing** | `go test -fuzz` on parsers (`ExtractQuotedTitle`, filters). |
| **HTTP + JSON API (optional)** | Same `TodoStore` behind `net/http`, context cancellation, timeouts. |
| **Migrations for SQLite** | Versioned `CREATE`/`ALTER`, forward-only migrations table. |
| **Windows paths and CI** | `filepath`, line endings, cross-OS tests if you care about portability. |

---

## Summary

This project is a **compact Go CLI** with a **repository abstraction**, **JSON and SQLite** storage, and **tests** around commands and persistence. The highest-leverage improvements are **less global state**, **clearer errors**, and **cleaner separation between storage and printing**. Feature work that pairs well with learning includes **real CLI parsing**, **env-based config**, and **one vertical slice** (due dates or undo) end to end including migrations and tests.
