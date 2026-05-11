package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WazedKhan/Go-Playground/tree/main/projects/mini-backend/internal/handler"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {

	// API routes
	http.HandleFunc("/", Greeting)

	// get route
	http.HandleFunc("/get", handler.Get)
	http.HandleFunc("/set", handler.Set)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	port := ":5000"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Printf("Starting server on :%s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
