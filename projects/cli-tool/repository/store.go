package repository

import (
	"cli-tool/models"
)

type TodoStore interface {
	GetTodos() ([]models.Todos, error)
	AddTodo(todo models.Todos) error
	ReplaceAll(todos []models.Todos) error
}

type jsonStore struct {
	filePath string
}

func NewJsonStore(filePath string) TodoStore {
	return &jsonStore{filePath: filePath}
}
