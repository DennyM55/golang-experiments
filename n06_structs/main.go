package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{
		Name: "Denny",
		Age:  36,
	}

	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
}
