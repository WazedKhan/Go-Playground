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

	err = os.WriteFile("db/setting.json", fileData, 0644)
	if err != nil {
		fmt.Println("failed to write to json file with error:", err)
		return err
	}
	fmt.Println("Setting updated successfully!")
	return nil
}

func AddTODO(todo models.Todos) error {
	todos := GetTODOs()
	todo.Id = int64((len(todos))) + 1
	todos = append(todos, todo)

	// check and update max title length
	UpdateMaxTitleLength(int64(len(todo.Title)))

	// convert struct into json
	fileData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Println("failed to marshal data into json format with err: ", err)
		return err
	}

	err = os.WriteFile("db/todos.json", fileData, 0644)
	if err != nil {
		fmt.Println("failed to write to json file with error:", err)
		return err
	}
	fmt.Println("TODO added successfully! ID:", todo.Id)
	return nil
}
