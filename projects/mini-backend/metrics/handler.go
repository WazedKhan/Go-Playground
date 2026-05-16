package middleware

import (
	"encoding/json"
	"net/http"
	"time"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type health struct {
		Status  string `json:"status"`
		Service string `json:"service"`
		Time    string `json:"time"`
	}

	res := health{
		Status:  "ok",
		Service: "mini-backend",
		Time:    time.Now().Format("2006-01-02 03:04:05 PM"),
	}

	json.NewEncoder(w).Encode(res)
}
