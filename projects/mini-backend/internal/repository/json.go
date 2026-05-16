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
	filePath = filepath.Join("./internal/db", "storage.json")
)

func WriteJsonFile(data models.User) (bool, error) {
	mu.Lock()
	defer mu.Unlock()

	existingData, _ := readJsonFile()
	if existingData != nil {
		maps.Copy(existingData, data)
	}
	fileData, err := json.MarshalIndent(existingData, "", "  ")
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
	mu.Lock()
	defer mu.Unlock()

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("failed to read from json file(%s): %q", filePath, err)
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
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("failed to read from json file(%s): %q", filePath, err)
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
