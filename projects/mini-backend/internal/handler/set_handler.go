package handler

import (
	"fmt"
	"net/http"

	"github.com/WazedKhan/Go-Playground/tree/main/projects/mini-backend/internal/models"
	"github.com/WazedKhan/Go-Playground/tree/main/projects/mini-backend/internal/repository"
)

func Set(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	data := models.User{key: value}
	ok, err := repository.WriteJsonFile(data)
	if !ok {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "value %s stored with key %s", value, key)
}
