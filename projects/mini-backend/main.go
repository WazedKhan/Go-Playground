package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/WazedKhan/Go-Playground/tree/main/projects/mini-backend/internal/handler"
	middleware "github.com/WazedKhan/Go-Playground/tree/main/projects/mini-backend/metrics"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is missing")
	}
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL is missing")
	}
	// if err := repository.InitPostgres(connStr); err != nil {
	// 	log.Fatalf("failed to init postgres: %v", err)
	// }
	// repository.SetStorage(repository.NewPostgresStorage())

	mux := http.NewServeMux()
	mux.HandleFunc("/get", handler.Get)
	mux.HandleFunc("/set", handler.Set)
	mux.HandleFunc("/health", middleware.GetHealth)
	mux.HandleFunc("/metrics", middleware.RouteMetrics)

	var wg sync.WaitGroup
	const shutDownTimeOut = 25 * time.Second // found that kubernetes waits 30 before kill -9
	port := ":8000"
	middleware := middleware.LoggerMiddleware(
		middleware.MetricsMiddleware(
			middleware.TrackRequests(&wg,
				middleware.APIKeyMiddleware(apiKey, mux),
			),
		),
	)
	srv := &http.Server{
		Addr:    port,
		Handler: middleware,
	}

	go func() {
		log.Printf("Starting server on http://localhost%s/\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// handling gracefully shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	log.Println("Shutting down... (Ctrl+C again to force)")
	stop()

	shutCtx, cancel := context.WithTimeout(context.Background(), shutDownTimeOut)
	defer cancel()

	go srv.Shutdown(shutCtx)
	waitDone := make(chan struct{})
	go func() {
		wg.Wait()
		close(waitDone)
	}()

	select {
	case <-waitDone:
		log.Println("all request done, clean exit")
	case <-shutCtx.Done():
		log.Printf("waited %q, forcing shutdown\n", shutDownTimeOut)
	}
}
