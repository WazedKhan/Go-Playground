package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	Greet(os.Stdout, "Elodie \n")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
