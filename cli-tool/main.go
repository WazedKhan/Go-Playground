package main

import (
	"bufio"
	"fmt"
	"os"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Welcome to Go Cli Tool")
	fmt.Println(quote.Go())
	userInput()
}

func userInput() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	if scanner.Scan() {
		name := scanner.Text()
		fmt.Printf("Hello, %s!\n", name)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
