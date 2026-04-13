package models

// Todos structure for the TODO items in the application.
type Todos struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
}
