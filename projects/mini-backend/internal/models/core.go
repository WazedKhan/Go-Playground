package models

import "time"

type User map[string]string
type RouteMetrics struct {
	TotalDuration time.Duration `json:"total_duration"`
	Count         int64         `json:"count"`
}

type Metrics struct {
	TotalRequests int64                   `json:"total_requests"`
	GetRequests   int64                   `json:"get_requests"`
	PostRequests  int64                   `json:"post_requests"`
	Latency       map[string]RouteMetrics `json:"latency"`
}
