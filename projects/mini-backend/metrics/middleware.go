package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/WazedKhan/Go-Playground/tree/main/projects/mini-backend/internal/models"
)

var (
	totalGetRequest  atomic.Int64
	totalPostRequest atomic.Int64
	totalRequest     atomic.Int64
	mu sync.Mutex
)

var latency = make(map[string]models.RouteMetrics)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		fmt.Printf(
			"[%s] %s %s \n",
			start.Format("2006-01-02 03:04:05 PM"),
			r.Method,
			r.URL.String(),
		)

		next.ServeHTTP(w, r)
	})
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		end := time.Now()

		duration := end.Sub(start)
		mu.Lock()
		existing := latency[r.URL.Path]
		latency[r.URL.Path] = models.RouteMetrics{
			Count:         existing.Count + 1,
			TotalDuration: existing.TotalDuration + duration,
		}
		mu.Unlock()

		switch r.Method {
		case "GET":
			totalGetRequest.Add(1)
		case "POST":
			totalPostRequest.Add(1)
		}

		// track every request
		totalRequest.Add(1)
	})
}

func GetRouteMetrics() models.Metrics {
	mu.Lock()
	defer mu.Unlock()
	return models.Metrics{
		TotalRequests: totalRequest.Load(),
		GetRequests:   totalGetRequest.Load(),
		PostRequests:  totalPostRequest.Load(),
		Latency:       latency,
	}
}
