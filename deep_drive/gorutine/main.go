package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from the go-routine!")
}

func main() {
	go func(s string) {
		for i := 0; i < 3; i++ {
			fmt.Println(s)
			time.Sleep(500 * time.Millisecond)
		}
	}("Hello from Anonymous Goroutine!")

	time.Sleep(2 * time.Second)
	fmt.Println("Main function complete!")
}
