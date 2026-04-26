package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"cli-tool/models"
)

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
