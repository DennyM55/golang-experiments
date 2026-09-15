package main

import "fmt"

type Speaker interface {
	Speak()
}

type Person struct {
	Name string
}

func (p Person) Speak() {
	fmt.Println("Hello, I am", p.Name)
}

func introduce(s Speaker) {
	s.Speak()
}

func main() {
	person := Person{
		Name: "Denny",
	}

	introduce(person)
}
