package main

import (
	"fmt"

	"cli-tool/internal"
	// main is the entry point of the application.
	// It prints a welcome message and calls internal.AppLoop() to start the application loop.
)

func main() {
	// prompt := promptui.Select{
	// 	Label: "Select Action",
	// 	Items: []string{"View TODOs: List TODOs", "Add TODO", "Mark TODO as Done", "Delete TODO"},
	// }

	// _, result, err := prompt.Run()

	// if err != nil {
	// 	fmt.Printf("Prompt failed %v\n", err)
	// 	return
	// }

	// fmt.Printf("You choose %q\n", result)
	fmt.Println("Welcome to Go Cli Tool")
	internal.AppLoop()
}
