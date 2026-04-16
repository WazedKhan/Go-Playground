package utils

import "fmt"

const (
	HelpCommand   = "todo help"
	QuitCommand   = "todo quit"
	AddCommand    = "todo add"
	ListCommand   = "todo list"
	DeleteCommand = "todo delete <id>"
	DoneCommand   = "todo done <id>"
)

func AvailableCommands() {
	fmt.Println(" -", HelpCommand)
	fmt.Println(" -", QuitCommand)
	fmt.Println(" -", AddCommand)
	fmt.Println(" -", ListCommand)
	fmt.Println(" -", DeleteCommand)
	fmt.Println(" -", DoneCommand)
}
