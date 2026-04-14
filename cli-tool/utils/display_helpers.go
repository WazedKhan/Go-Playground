package utils

import (
	"fmt"
	"strings"

	"cli-tool/models"
	"cli-tool/repository"
)

// DisplayTodos prints the list in a user-friendly format. If the list is empty, it informs the user that no TODOs are found.
func DisplayTodos(data []models.Todos) {
	if len(data) == 0 {
		fmt.Println("\n  No TODOs found.")
		return
	}

	maxTitleLength := repository.GetMaxTitleLength()
	if maxTitleLength < 5 {
		maxTitleLength = 5 // minimum column width
	}

	divider := strings.Repeat("─", maxTitleLength+46)

	fmt.Printf("\n  %d TODO(s)\n", len(data))
	fmt.Println("  " + divider)
	fmt.Printf("  %-4s  %-*s  %-10s  %s\n", "ID", maxTitleLength, "Title", "Status", "Created At")
	fmt.Println("  " + divider)

	for _, todo := range data {
		fmt.Printf("  %-4d  %-*s  %-10s  %s\n",
			todo.Id,
			maxTitleLength, todo.Title,
			todo.Status,
			todo.CreatedAt,
		)
	}

	fmt.Println("  " + divider)
}
