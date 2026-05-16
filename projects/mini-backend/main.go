package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WazedKhan/Go-Playground/tree/main/projects/mini-backend/internal/handler"
	middleware "github.com/WazedKhan/Go-Playground/tree/main/projects/mini-backend/metrics"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/get", handler.Get)
	mux.HandleFunc("/set", handler.Set)
	mux.HandleFunc("/health", middleware.GetHealth)

	port := ":8000"
	fmt.Println("Server is running on port" + port)
	middleware := middleware.LoggerMiddleware(mux)

	// Start server on port specified above
	log.Printf("Starting server on :%s...\n", port)
	log.Fatal(http.ListenAndServe(port, middleware))
}
