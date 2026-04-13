package models

type Todos struct {
	Id        int64  `json:"id"`
	Age       int64  `json:"age"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
