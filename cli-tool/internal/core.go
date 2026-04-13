package internal

import (
	"bufio"
	"fmt"
	"os"

	"cli-tool/repository"
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
		if input == "q" || input == "quit" {
			fmt.Println("Exiting the application...")
			break
		}
	}
}

// userInput reads user input from the command line and executes the corresponding command.
func userInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	data := repository.GetTODOList(1)
	var input string

	if scanner.Scan() {
		input = scanner.Text()

		switch input {
		case "h", "help":
			utils.AvailableCommands()
		case "a", "add":
			fmt.Println("Add Command")
		case "l", "list":
			utils.DisplayTodos(data)
		case "d", "delete":
			fmt.Println("Delete Command")
		default:
			fmt.Println("Unknown Command")
		}
	}
	return input
}
