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
		maxTitleLength = 5
	}

	divider := strings.Repeat("─", maxTitleLength+52)

	fmt.Printf("\n  %d TODO(s)\n", len(data))
	fmt.Println("  " + divider)
	fmt.Printf("  %-4s  %-2s  %-*s  %-10s  %s\n", "ID", "  ", maxTitleLength, "Title", "Status", "Created At")
	fmt.Println("  " + divider)

	for _, todo := range data {
		symbol := "  "
		if todo.Symbol != nil {
			symbol = *todo.Symbol
		}

		fmt.Printf("  %-4d  %-2s  %-*s  %-10s  %s\n",
			todo.Id,
			symbol,
			maxTitleLength, TruncateTitle(todo.Title, 50),
			todo.Status,
			todo.CreatedAt,
		)
	}

	fmt.Println("  " + divider)
}
