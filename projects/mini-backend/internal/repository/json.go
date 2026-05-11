package repository

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"path/filepath"

	"github.com/WazedKhan/Go-Playground/tree/main/projects/mini-backend/internal/models"
)

var filePath = filepath.Join("./internal/db", "storage.json")

func WriteJsonFile(data models.User) (bool, error) {
	existingData, _ := ReadJsonFile()
	if existingData != nil {
		maps.Copy(data, existingData)
	}
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshall data: %q", err)
		return false, err
	}

	writeErr := os.WriteFile(filePath, fileData, 0644)
	if writeErr != nil {
		fmt.Printf("failed to write into json file: %q", writeErr)
		return false, writeErr
	}
	return true, nil
}

func ReadJsonFileByKey(key string) (*string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("failed to read from json file(%s): %q",filePath, err)
	}

	var res models.User
	json.Unmarshal(content, &res)
	value, ok := res[key]
	if !ok {
		return nil, fmt.Errorf("no data found")
	}
	return &value, nil
}

func ReadJsonFile() (map[string]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("failed to read from json file(%s): %q",filePath, err)
	}

	var res models.User
	json.Unmarshal(content, &res)
	return res, nil
}
