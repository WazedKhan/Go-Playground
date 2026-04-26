package repository

import (
	"fmt"
	"os"
	"strings"

	"cli-tool/models"
)

// TodoStore is the persistence contract for todos. Two implementations exist:
// jsonStore (todos.json) and sqliteStore (todos.db). Application code calls
// GetTODOs / AddTODO / UpdateTODO; those functions delegate here so the rest
// of the app stays the same regardless of backend.
type TodoStore interface {
	List() []models.Todos
	Add(todo models.Todos) error
	ReplaceAll(todos []models.Todos) error
}

var todoStore TodoStore = &jsonStore{}

// SetTodoStore swaps the active backend (e.g. tests). nil restores the default JSON store.
func SetTodoStore(s TodoStore) {
	if s == nil {
		todoStore = &jsonStore{}
		return
	}
	todoStore = s
}

// Init selects the todo backend from CLI_TOOL_STORAGE: "json" (default) or "sqlite".
// Settings (setting.json) always use JSON regardless of this value.
func Init() error {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("CLI_TOOL_STORAGE"))) {
	case "sqlite":
		if err := os.MkdirAll(DataDir, 0755); err != nil {
			return fmt.Errorf("create data dir: %w", err)
		}
		s, err := NewSQLiteTodoStore(sqlitePath())
		if err != nil {
			return err
		}
		todoStore = s
	default:
		todoStore = &jsonStore{}
	}
	return nil
}

// GetTODOs loads all todos from the active TodoStore.
func GetTODOs() []models.Todos {
	return todoStore.List()
}

// AddTODO persists a new todo via the active TodoStore.
func AddTODO(todo models.Todos) error {
	return todoStore.Add(todo)
}

// UpdateTODO replaces the full todo list (same semantics as rewriting todos.json).
func UpdateTODO(todos []models.Todos) error {
	return todoStore.ReplaceAll(todos)
}
