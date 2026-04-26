package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"cli-tool/models"
)

// jsonStore keeps todos in todos.json under DataDir.
type jsonStore struct{}

func (jsonStore) List() []models.Todos {
	content, err := os.ReadFile(todosPath())
	if err != nil {
		fmt.Println("error reading json file,", err)
		return []models.Todos{}
	}
	var todos []models.Todos
	if err := json.Unmarshal(content, &todos); err != nil {
		fmt.Println("error un-marshalling json file,", err)
		return []models.Todos{}
	}
	return todos
}

func (jsonStore) Add(todo models.Todos) error {
	todos := GetTODOs()
	todo.Id = GetNextID(todos)
	todos = append(todos, todo)

	if err := UpdateMaxTitleLength(int64(len(todo.Title))); err != nil {
		return err
	}

	fileData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Println("failed to marshal data into json format with err: ", err)
		return err
	}
	if err := os.WriteFile(todosPath(), fileData, 0644); err != nil {
		fmt.Println("failed to write to json file with error:", err)
		return err
	}
	fmt.Println("TODO added successfully! ID:", todo.Id)
	return nil
}

func (jsonStore) ReplaceAll(todos []models.Todos) error {
	fileData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Println("failed to marshal data into json format with err: ", err)
		return err
	}
	if err := os.WriteFile(todosPath(), fileData, 0644); err != nil {
		fmt.Println("failed to write to json file with error:", err)
		return err
	}
	fmt.Println("TODO updated successfully!")
	return nil
}
