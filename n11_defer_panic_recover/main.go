package main

import "fmt"

func riskyOperation() {
	defer fmt.Println("Cleanup runs before function exits")

	panic("something went wrong")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from:", err)
		}
	}()

	riskyOperation()

	fmt.Println("This line will not run")
}
