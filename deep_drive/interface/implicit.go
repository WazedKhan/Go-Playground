package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct {name string}
func (d Dog) Sound() string {
	return "Woof"
}

type Cat struct {name string}
func (c Cat) Sound() string {return "Mewo"}

func Speak(a Animal) {
	fmt.Println(a.Sound())
}

func main() {
	Speak(Cat{})
	Speak(Dog{})
}