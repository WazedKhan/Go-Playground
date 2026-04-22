// Define the interface — this is the CONTRACT
package main

import "fmt"

type Greeter interface {
    Greet() string
}

// Two DIFFERENT types, both satisfy the Greeter interface
type English struct{}
func (e English) Greet() string { return "Hello!" }

type Bangla struct{}
func (b Bangla) Greet() string { return "Assalamu Alaikum!" }

// This function accepts ANYTHING that satisfies Greeter
func SayHello(g Greeter) {
    fmt.Println(g.Greet())
}
