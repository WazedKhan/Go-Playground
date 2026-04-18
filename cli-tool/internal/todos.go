package internal

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"cli-tool/models"
	"cli-tool/repository"
)

// GetTodoList is a function that retrieves a list of TODO items. Currently, it returns an empty list and no error.
func GetTodos() ([]models.Todos, error) {
	return repository.GetTODOs(), nil
}

func CreateTodos() error {
	scanner := bufio.NewScanner(os.Stdin)
	now := time.Now().Format("2006-01-02 15:04:05")
	todo := models.Todos{
		CreatedAt: now,
		Status:    models.PENDING,
	}
	var title string
	fmt.Printf("Enter your todo: ")
	if scanner.Scan() {
		title = scanner.Text()
	}
	todo.Title = title

	repository.AddTODO(todo)
	return nil
}

func MarkTodoDone(id int64) error {
	todos := repository.GetTODOs()
	for index, todo := range todos {
		if todo.Id == id {
			todos[index].Status = models.DONE
		}
	}
	repository.UpdateTODO(todos)
	return nil
}

func DeleteTodo(id int64) error {
	todos := repository.GetTODOs()
	for index, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	repository.UpdateTODO(todos)
	return nil
}

func GetFilteredTodos(filter string) ([]models.Todos, error) {
	todos := repository.GetTODOs()
	var filtered []models.Todos
	for _, todo := range todos {
		if todo.Status == filter {
			filtered = append(filtered, todo)
		}
	}
	return filtered, nil
}

func EditTodo(id int64, title string) error {
	todos, _ := GetTodos()
	for index, value := range todos {
		if value.Id == id {
			todos[index].Title = title
		}
	}
	repository.UpdateTODO(todos)
	return nil
}
