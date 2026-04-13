package utils

import "fmt"

const (
	HelpCommand   = "help"
	QuitCommand   = "quit"
	AddCommand    = "add"
	ListCommand   = "list"
	DeleteCommand = "delete"
)

func AvailableCommands() {
	fmt.Println(" -", HelpCommand)
	fmt.Println(" -", QuitCommand)
	fmt.Println(" -", AddCommand)
	fmt.Println(" -", ListCommand)
	fmt.Println(" -", DeleteCommand)
}
