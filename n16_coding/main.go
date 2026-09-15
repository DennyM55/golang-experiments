package main

import "fmt"

func secondLargest(numbers []int) int {
	largest := numbers[0]
	second := numbers[0]

	for _, number := range numbers {

		if number > largest {
			second = largest
			largest = number
		} else if number > second && number != largest {
			second = number
		}
	}

	return second
}

func main() {
	numbers := []int{10, 40, 20, 50, 30}

	result := secondLargest(numbers)

	fmt.Println("Second largest:", result)
}
