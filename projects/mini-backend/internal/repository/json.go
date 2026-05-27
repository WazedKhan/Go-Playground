package repository

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"sync"

	"github.com/WazedKhan/Go-Playground/tree/main/projects/mini-backend/internal/models"
)

var (
	mu       sync.Mutex
	jsonFilePath = filepath.Join("./internal/db", "storage.json")
	sqlFilePath = filepath.Join("./internal/db", "data.db")
)

func (j *jsonStorage)Save(key, value string) (bool, error) {
	mu.Lock()
	defer mu.Unlock()
	data := models.User{
		key:value,
	}
	existingData, _ := readJsonFile()
	if existingData != nil {
		maps.Copy(existingData, data)
	}
	fileData, err := json.MarshalIndent(existingData, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshall data: %q", err)
		return false, err
	}

	writeErr := os.WriteFile(jsonFilePath, fileData, 0644)
	if writeErr != nil {
		fmt.Printf("failed to write into json file: %q", writeErr)
		return false, writeErr
	}
	return true, nil
}

func (j jsonStorage)Get(key string) (*string, error) {
	mu.Lock()
	defer mu.Unlock()

	content, err := os.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Printf("failed to read from json file(%s): %q", jsonFilePath, err)
	}

	var res models.User
	json.Unmarshal(content, &res)
	value, ok := res[key]
	if !ok {
		return nil, fmt.Errorf("no data found")
	}
	return &value, nil
}

func readJsonFile() (map[string]string, error) {
	content, err := os.ReadFile(jsonFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]string{}, nil
		}
		fmt.Printf("failed to read from json file(%s): %q", jsonFilePath, err)
	}

	var res models.User
	json.Unmarshal(content, &res)
	return res, nil
}

func ReadJsonFileSafe() (map[string]string, error) {
	mu.Lock()
	defer mu.Unlock()
	return readJsonFile()
}
