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
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Please enter your name (or type 'q' or 'quit' to exit): ")
		if scanner.Scan() {
			if input := scanner.Text(); input == "q" || input == "quit" {
				fmt.Println("Goodbye!")
				break
			}
			fmt.Printf("Hello, %s!\n", scanner.Text())
		}
	}
}
