package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"cli-tool/models"
)

func (j *jsonStore) GetTodos() ([]models.Todos, error) {
	content, err := os.ReadFile(j.filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading json file: %w", err)
	}

	var todos []models.Todos
	if err := json.Unmarshal(content, &todos); err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %w", err)
	}
	return todos, nil
}

func GetTODOs() []models.Todos {
	todos, err := activeStore.GetTodos()
	if err != nil {
		fmt.Println(err)
		return []models.Todos{}
	}

	return todos
}

func GetMaxTitleLength() int {
	content, err := os.ReadFile(settingPath())
	if err != nil {
		fmt.Println("error reading json file,", err)
		return 0
	}
	var setting models.Setting
	if err := json.Unmarshal(content, &setting); err != nil {
		fmt.Println("error un-marshalling json file,", err)
		return 0
	}

	return setting.MaxTitleLength
}
