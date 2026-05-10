package main

import (
	"os"
)

func main() {
	Greet(os.Stdout, "Elodie \n")
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper )
}
