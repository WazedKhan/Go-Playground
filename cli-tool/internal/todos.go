package internal

import (
	"fmt"
	"time"

	"cli-tool/models"
	"cli-tool/repository"
	"cli-tool/utils"
)

func GetTodos() ([]models.Todos, error) {
	todos := repository.GetTODOs()
	if len(todos) == 0 {
		return nil, fmt.Errorf("No todos found")
	}
	for index, todo := range todos {
		todos[index].CreatedAt = utils.ConvertDateToRelativeTime(todo.CreatedAt)
		switch todo.Status {
		case models.PENDING:
			symbol := "○"
			todos[index].Symbol = &symbol
		case models.DONE:
			symbol := "✓"
			todos[index].Symbol = &symbol
		}
	}
	return todos, nil
}

func CreateTodos(title string) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	valid, errMsg := utils.IsValidTitle(title)
	if !valid {
		return fmt.Errorf("%s", errMsg.Error())
	}
	todo := models.Todos{
		Title:     title,
		CreatedAt: now,
		Status:    models.PENDING,
	}

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
	found := false
	for index, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:index], todos[index+1:]...)
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("Todo with ID %d not found", id)
	}
	repository.UpdateTODO(todos)
	return nil
}

func GetFilteredTodos(filter string) ([]models.Todos, error) {
	todos, _ := GetTodos()
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

func GetPendingTodos() ([]models.Todos, error) {
	todos, _ := GetTodos()
	var pending []models.Todos
	for _, todo := range todos {
		if todo.Status == models.PENDING {
			pending = append(pending, todo)
		}
	}
	return pending, nil
}

func SelectAndDeleteTodo() {
	todos, _ := GetTodos()
	if len(todos) == 0 {
		fmt.Println("  No todos to delete.")
		return
	}

	labels := make([]string, len(todos))
	for i, todo := range todos {
		labels[i] = fmt.Sprintf("%d: %s", todo.Id, todo.Title)
	}

	index, _, err := utils.SelectPrompt("Select a TODO to delete", labels)
	if err != nil {
		fmt.Println("  No selection made. Returning to main menu.")
		return
	}

	selected := todos[index]
	confirmIndex, _, err := utils.SelectPrompt(
		fmt.Sprintf("Delete \"%s\"?", selected.Title),
		[]string{"No, cancel", "Yes, delete"},
	)
	if err != nil {
		fmt.Println("  Prompt failed:", err)
		return
	}

	if confirmIndex == 1 {
		err := DeleteTodo(selected.Id)
		if err != nil {
			fmt.Println("  Failed to delete:", err)
			return
		}
		fmt.Printf("  ✓ \"%s\" deleted.\n", selected.Title)
	} else {
		fmt.Println("  Cancelled.")
	}
}

func SelectAndMarkTodoDone() {
	todos, _ := GetPendingTodos()
	if len(todos) == 0 {
		fmt.Println("  No todos to mark as done.")
		return
	}

	labels := make([]string, len(todos))
	for i, todo := range todos {
		labels[i] = fmt.Sprintf("%d: %s", todo.Id, todo.Title)
	}
	index, _, err := utils.SelectPrompt("Select a TODO to mark as done", labels)
	if err != nil {
		fmt.Println("  No selection made. Returning to main menu.")
		return
	}

	selected := todos[index]
	err = MarkTodoDone(selected.Id)
	if err != nil {
		fmt.Println("  Failed to mark as done:", err)
		return
	}
	fmt.Printf("  ✓ \"%s\" marked as done.\n", selected.Title)
}
