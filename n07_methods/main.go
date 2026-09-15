package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) greet() {
	fmt.Println("Hello, my name is", u.Name)
}

func main() {
	user := User{
		Name: "Denny",
		Age:  36,
	}

	user.greet()
}
