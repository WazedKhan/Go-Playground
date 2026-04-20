package utils

import "fmt"

func IsValidTitle(title string) (bool, error) {
	if len(title) == 0 {
		return false, fmt.Errorf("Title cannot be empty.")
	}
	return true, nil
}
