package internal

import (
	"bufio"
	"os"
	"time"

	"cli-tool/models"
	"cli-tool/repository"
)

// GetTodoList is a function that retrieves a list of TODO items. Currently, it returns an empty list and no error.
func GetTodoList() ([]models.Todos, error) {
	return []models.Todos{}, nil
}

func CreateTodos() error {
	scanner := bufio.NewScanner(os.Stdin)
	now := time.Now().Format("2006-01-02 15:04:05")
	todo := models.Todos{
		CreatedAt: now,
		Status:    models.PENDING,
	}
	var title string
	if scanner.Scan() {
		title = scanner.Text()
	}
	todo.Title = title

	repository.AddTODO(todo)
	return nil
}
