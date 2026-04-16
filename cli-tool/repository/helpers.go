package repository

import "cli-tool/models"

func GetNextID(todos []models.Todos) int64 {
	if len(todos) == 0 {
		return 1
	}
	maxId := todos[0].Id
	for _, todo := range todos {
		if todo.Id > maxId {
			maxId = todo.Id
		}
	}
	return maxId + 1
}
