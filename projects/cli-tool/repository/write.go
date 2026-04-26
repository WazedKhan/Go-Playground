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

