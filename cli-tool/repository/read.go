package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"cli-tool/models"
)

func GetTODOs() []models.Todos {
	content, err := os.ReadFile("db/todos.json")
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

func GetMaxTitleLength() int {
	content, err := os.ReadFile("db/setting.json")
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
