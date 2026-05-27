package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WazedKhan/Go-Playground/tree/main/projects/mini-backend/internal/repository"
)

func Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, err := repository.GetStorage().Get(key)
	if err != nil {
		fmt.Fprint(w, "no value found with given key:", key)
		log.Println(fmt.Errorf("no value found with given, %s\n", key))
		return
	}
	fmt.Fprint(w, *value)
}
