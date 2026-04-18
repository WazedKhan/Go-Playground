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
	if len(parts) < 2 {
		fmt.Println("  Invalid command. Type 'todo --help' for help.")
		return
	} else if parts[0] != "todo" {
		fmt.Println("  Unknown command. Type 'todo --help' for help.")
		return
	}

	switch {
	case parts[1] == "-h" || parts[1] == "--help" || parts[1] == "help":
		utils.AvailableCommands()

	case len(parts) >= 3 && parts[0] == "todo" && parts[1] == "add" && !utils.ContainHelpFlag(input):
		title, err := utils.ExtractQuotedTitle(input)
		if err != nil {
			fmt.Println("  Title must be in quotes. Usage: todo add \"New todo\"")
			return
		}
		createErr := CreateTodos(title)
		if createErr != nil {
			fmt.Println("  Error creating todo:", createErr)
		}

	case len(parts) == 2 && parts[1] == "list" && !utils.ContainHelpFlag(input):
		todos, _ := GetTodos()
		utils.DisplayTodos(todos)

	case len(parts) == 3 && parts[1] == "done" && !utils.ContainHelpFlag(input):
		id, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("  Invalid ID. Usage: todo done <id>")
			return
		}
		MarkTodoDone(int64(id))

	case len(parts) == 3 && parts[1] == "delete" && !utils.ContainHelpFlag(input):
		id, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("  Invalid ID. Usage: todo delete <id>")
			return
		}
		err = DeleteTodo(int64(id))
		if err != nil {
			fmt.Println("  Error deleting todo:", err)
		}

	case len(parts) == 3 && strings.Contains(parts[2], "--filter=") && !utils.ContainHelpFlag(input):
		filter := strings.TrimPrefix(parts[2], "--filter=")
		validFilter, err := utils.ValidateFilter(filter)
		if err != nil {
			fmt.Println("  Invalid filter. Use 'pending' or 'done'.")
			return
		}
		todos, _ := GetFilteredTodos(validFilter)
		utils.DisplayTodos(todos)

	case len(parts) >= 4 && parts[0] == "todo" && parts[1] == "edit" && !utils.ContainHelpFlag(input):
		id, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("  Invalid ID. Usage: todo edit <id> \"New title\"")
			return
		}
		title, err := utils.ExtractQuotedTitle(input)
		if err != nil {
			fmt.Println("  Title must be in quotes. Usage: todo edit <id> \"New title\"")
			return
		}
		EditTodo(int64(id), title)

	default:
		if utils.ContainHelpFlag(input) {
			key := parts[1]
			utils.CommandHelp(key)
			return
		}
		fmt.Println("  Unknown command. Type '--help' for help.")
	}
}
