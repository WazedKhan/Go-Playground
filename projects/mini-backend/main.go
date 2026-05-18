package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	mux.HandleFunc("/metrics", middleware.RouteMetrics)

	// to shutdown gracefully we need to allow main function to reach after server listen
	// and to allow go signal layer we need to use goroutine as go reads codes from up to down
	port := ":8000"
	middleware := middleware.LoggerMiddleware(
		middleware.MetricsMiddleware(mux),
	)
	srv := &http.Server{
		Addr: port,
		Handler: middleware,
	}

	go func() {
        log.Printf("Starting server on http://localhost%s/\n", port)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
    }()
	// now that we are running with go routine if we start the server it will close as main function completes before goroutine
	// we need to wait for signal for quite
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit
	log.Println("Shutting down the server ....")

	// right now we are just waiting for the signal to quite then quitting the server but still following forceful shutdown
	// set lets wait for some time to shutdown to give our server to handle its ongoing request/task
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
    defer cancel()

	if err := srv.Shutdown(ctx); err != nil{
		log.Fatal("forced shutdown:", err)
	}
	log.Println("server exited cleanly")
}
