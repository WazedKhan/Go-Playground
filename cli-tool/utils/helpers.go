package utils

import (
	"fmt"
	"strings"
	"time"
)

func ExtractQuotedTitle(input string) (string, error) {
	start := strings.Index(input, "\"")
	end := strings.LastIndex(input, "\"")
	if start == -1 || end == -1 || start == end {
		return "", fmt.Errorf("no quoted title found")
	}

	return input[start+1 : end], nil
}

// ConvertDateToRelativeTime converts a time string to a relative time format (e.g., "2 hours ago").
func ConvertDateToRelativeTime(dateStr string) string {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return dateStr // Return original string if parsing fails
	}

	duration := time.Since(t)

	switch {
	case duration < time.Minute:
		return "Just now"
	case duration < time.Hour:
		return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
	case duration < 24*time.Hour:
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	case duration < 30*24*time.Hour:
		return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
	default:
		return t.Format("Jan 2, 2006")
	}
}

func ValidateFilter(filter string) (string, error) {
	switch strings.ToLower(filter) {
	case "pending":
		return "PENDING", nil
	case "done":
		return "DONE", nil
	}
	return "", fmt.Errorf("Invalid filter. Use 'pending' or 'done'.")
}

func TruncateTitle(title string, maxLength int) string {
	if len(title) <= maxLength {
		return title
	}
	return title[:maxLength-3] + "..."
}
