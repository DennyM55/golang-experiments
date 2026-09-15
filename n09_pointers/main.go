package main

import "fmt"

type User struct {
	Name string
}

func changeName(user *User) {
	user.Name = "John"
}

func main() {
	user := User{
		Name: "Denny",
	}

	fmt.Println("Before:", user.Name)

	changeName(&user)

	fmt.Println("After:", user.Name)
}
