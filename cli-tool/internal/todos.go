package internal

import "cli-tool/models"

// GetTodoList is a function that retrieves a list of TODO items. Currently, it returns an empty list and no error.
func GetTodoList() ([]models.Todos, error) {
	return []models.Todos{}, nil
}
