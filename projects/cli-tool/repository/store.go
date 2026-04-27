package repository

import (
	"cli-tool/models"
)

var activeStore TodoStore = NewJsonStore()

type TodoStore interface {
	GetTodos() ([]models.Todos, error)
	AddTodo(todo models.Todos) error
	ReplaceAll(todos []models.Todos) error
}

// jsonStore reads and writes via todosPath() so repository.DataDir changes
// (e.g. in tests) take effect without reinitializing the store.
type jsonStore struct{}

func NewJsonStore() TodoStore {
	return &jsonStore{}
}

func SetTodoStore(s TodoStore) {
	activeStore = s
}
