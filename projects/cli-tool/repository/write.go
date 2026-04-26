package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"cli-tool/models"
)

func UpdateMaxTitleLength(newLength int64) error {
	currentML := GetMaxTitleLength()
	var setting models.Setting
	if newLength > int64(currentML) {
		setting.MaxTitleLength = int(newLength)
		WriteSetting(setting)
		fmt.Printf("Max title length updated to %d characters.\n", newLength)
	}

	return nil
}

func WriteSetting(setting models.Setting) error {
	// convert struct into json
	fileData, err := json.MarshalIndent(setting, "", "  ")
	if err != nil {
		fmt.Println("failed to marshal data into json format with err: ", err)
		return err
	}

	err = os.WriteFile(settingPath(), fileData, 0644)
	if err != nil {
		fmt.Println("failed to write to json file with error:", err)
		return err
	}
	fmt.Println("Setting updated successfully!")
	return nil
}

func (j *jsonStore) AddTodo(todo models.Todos) error {
	todos, err := j.GetTodos()
	if err != nil {
		return fmt.Errorf("error reading todos: %w", err)
	}

	todo.Id = GetNextID(todos)
	todos = append(todos, todo)

	UpdateMaxTitleLength(int64(len(todo.Title)))
	fileData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Println("failed to marshal data into json format with err: ", err)
		return err
	}

	err = os.WriteFile(j.filePath, fileData, 0644)
	if err != nil {
		fmt.Println("failed to write to json file with error:", err)
		return err
	}
	fmt.Println("TODO added successfully! ID:", todo.Id)
	return nil
}

func AddTODO(todo models.Todos) error {
	todos := GetTODOs()
	todo.Id = GetNextID(todos)
	todos = append(todos, todo)

	// check and update max title length
	UpdateMaxTitleLength(int64(len(todo.Title)))

	// convert struct into json
	fileData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Println("failed to marshal data into json format with err: ", err)
		return err
	}

	err = os.WriteFile(todosPath(), fileData, 0644)
	if err != nil {
		fmt.Println("failed to write to json file with error:", err)
		return err
	}
	fmt.Println("TODO added successfully! ID:", todo.Id)
	return nil
}

func UpdateTODO(todos []models.Todos) error {
	fileData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Println("failed to marshal data into json format with err: ", err)
		return err
	}

	err = os.WriteFile(todosPath(), fileData, 0644)
	if err != nil {
		fmt.Println("failed to write to json file with error:", err)
		return err
	}
	fmt.Println("TODO updated successfully!")
	return nil
}

func (j *jsonStore) ReplaceAll(todos []models.Todos) error {
	content, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Println("failed to marshal data into json format with err: ", err)
		return err
	}

	err = os.WriteFile(j.filePath, content, 0644)
	if err != nil {
		fmt.Println("failed to write to json file with error:", err)
		return err
	}
	fmt.Println("TODO updated successfully!")
	return nil
}
