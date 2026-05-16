package middleware

import (
	"fmt"
	"net/http"
	"time"
)

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
