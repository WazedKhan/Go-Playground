package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/WazedKhan/Go-Playground/tree/main/projects/mini-backend/internal/models"
)

var (
	totalGetRequest  int64
	totalPostRequest int64
	totalRequest     int64
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
		existing := latency[r.URL.Path]

		latency[r.URL.Path] = models.RouteMetrics{
			Count:         existing.Count + 1,
			TotalDuration: existing.TotalDuration + duration,
		}

		switch r.Method {
		case "GET":
			totalGetRequest++
		case "POST":
			totalPostRequest++
		}

		// track every request
		totalRequest++
	})
}

func GetRouteMetrics() models.Metrics {
	return models.Metrics{
		TotalRequests: totalRequest,
		GetRequests:   totalGetRequest,
		PostRequests:  totalPostRequest,
		Latency:       latency,
	}
}
