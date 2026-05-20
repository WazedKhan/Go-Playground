package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/WazedKhan/Go-Playground/tree/main/projects/mini-backend/internal/models"
	"github.com/WazedKhan/Go-Playground/tree/main/projects/mini-backend/internal/repository"
)

func Set(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	data := models.User{key: value}
	log.Println("Going sleep for two minute ... .")
	time.Sleep(2 * time.Minute)
	ok, err := repository.WriteJsonFile(data)
	if !ok {
		fmt.Fprint(w, err)
	}
	fmt.Fprintf(w, "value %s stored with key %s \n", value, key)
}
