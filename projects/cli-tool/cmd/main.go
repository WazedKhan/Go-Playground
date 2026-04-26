package main

import (
	"fmt"
	"os"

	"cli-tool/internal"
	"cli-tool/repository"
)

func main() {
	fmt.Println("Welcome to Go Cli Tool")
	if err := repository.Init(); err != nil {
		fmt.Fprintln(os.Stderr, "storage:", err)
		os.Exit(1)
	}
	internal.AppLoop()
}
