package main

import "fmt"

func main() {
	fmt.Println("Hello CLI-Tool")
	userInput()
}

func userInput(){
	input, _ := fmt.Scanf("Give your input: ")
	fmt.Printf("Hello %d", input)
}