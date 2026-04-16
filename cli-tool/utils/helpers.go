package utils

import (
	"fmt"
	"strings"
)

func ExtractQuotedTitle(input string) (string, error) {
	start := strings.Index(input, "\"")
	end := strings.LastIndex(input, "\"")
	if start == -1 || end == -1 || start == end {
		return "", fmt.Errorf("no quoted title found")
	}

	return input[start+1 : end], nil
}
