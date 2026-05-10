package main

import (
	"fmt"
	"log"
	"net/http"
)

var hashMap = make(map[string]string)

func Greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, ok := hashMap[key]
	if !ok {
		fmt.Fprint(w, "no value found with given key:", key)
		log.Println(fmt.Errorf("no value found with given, %s\n", key))
	}
	fmt.Fprint(w, value)
}

func Set(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	hashMap[key] = value
	fmt.Fprintf(w, "value %s stored with key %s", value, key)
}

func main() {

	// API routes
	http.HandleFunc("/", Greeting)

	// get route
	http.HandleFunc("/get", Get)
	http.HandleFunc("/set", Set)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	port := ":5000"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Printf("Starting server on :%s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
