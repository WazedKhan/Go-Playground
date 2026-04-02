package main

import (
	"bufio"
	"fmt"
	"os"

	"cli-tool/utils"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Welcome to Go Cli Tool")
	fmt.Println(quote.Go())
	userInput()
}

func userInput() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Select and Option: ")
		utils.AvailableCommands()
		if scanner.Scan() {
			if input := scanner.Text(); input == "q" || input == "quit" {
				fmt.Println("Goodbye!")
				break
			}

			input := scanner.Text()
			switch input {
			case "h", "help":
				fmt.Println("Help Command")
			case "a", "add":
				fmt.Println("Add Command")
			case "l", "list":
				fmt.Println("List Command")
			case "d", "delete":
				fmt.Println("Delete Command")
			default:
				fmt.Println("Unknown Command")
			}
		}
	}
}
