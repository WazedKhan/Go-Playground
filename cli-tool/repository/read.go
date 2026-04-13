package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"cli-tool/models"
)

func GetTODOList(id int64) []models.Todos {
	content, err := os.ReadFile("db/todos.json")
	if err != nil {
		fmt.Println("error reading json file,", err)
		return []models.Todos{}
	}
	var todos []models.Todos
	if err := json.Unmarshal(content, &todos); err != nil {
		fmt.Println("error unmarshalling json file,", err)
		return []models.Todos{}
	}

	return todos
}
