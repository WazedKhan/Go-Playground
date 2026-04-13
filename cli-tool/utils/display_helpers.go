package utils

import (
	"fmt"

	"cli-tool/models"
)

func DisplayTodos(data []models.Todos) {
	for _, value := range data {
		status := "🔴 pending"
		if value.Status == "completed" {
			status = "✅ completed"
		}

		fmt.Println("┌────────────────────────────────────────┐")
		fmt.Printf("│  %-38s│\n", "TODO")
		fmt.Println("├────────────────────────────────────────┤")
		fmt.Printf("│  ID         : %-25v│\n", value.Id)
		fmt.Printf("│  Title      : %-25v│\n", value.Title)
		fmt.Printf("│  Status     : %-24v│\n", status)
		fmt.Printf("│  Created At : %-25v│\n", value.CreatedAt)
		fmt.Println("└────────────────────────────────────────┘")
		fmt.Println()
	}
}
