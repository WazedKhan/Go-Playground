package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"cli-tool/utils"
)

// AppLoop is the main loop of the application that handles user input and executes commands.
func AppLoop() {
	fmt.Println("===================")
	fmt.Println("Select and Option: ")
	fmt.Println("===================")
	utils.AvailableCommands()
	for {
		input := userInput()
		HandleCommands(input)
		if input == "q" || input == "quit" {
			fmt.Println("Exiting the application...")
			break
		}
	}
}

// userInput reads user input from the command line and executes the corresponding command.
func userInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}
	return ""
}

func HandleCommands(input string) {
	parts := strings.Fields(input)

	switch {
	case input == "h" || input == "help":
		utils.AvailableCommands()
	case input == "a" || parts[1] == "add":
		CreateTodos()
	case len(parts) == 2 && parts[1] == "list":
		todos, _ := GetTodos()
		utils.DisplayTodos(todos)
	case len(parts) == 3 && parts[1] == "done":
		id, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("  Invalid ID. Usage: todo done <id>")
		}
		MarkTodoDone(int64(id))
	default:
		fmt.Println("  Unknown command. Type 'h' for help.")
	}
}
