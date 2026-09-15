package main

import "fmt"

func main() {
	ages := map[string]int{
		"Denny": 36,
		"John":  30,
	}

	fmt.Println("Denny age:", ages["Denny"])

	ages["Alice"] = 28
	fmt.Println("All ages:", ages)

	delete(ages, "John")
	fmt.Println("After delete:", ages)
}
