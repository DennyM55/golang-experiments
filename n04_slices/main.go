package main

import "fmt"

func main() {

	// Fixed-size array
	var numbers [3]int = [3]int{10, 20, 30}
	fmt.Println("Array:", numbers)

	// Slice: dynamic-size collection
	var scores []int = []int{10, 20, 30}

	scores = append(scores, 40)

	fmt.Println("Slice:", scores)
}
