package main

import (
	"errors"
	"fmt"
)

// N3 — Functions and multiple return values
// Go functions look like Java functions, but they can return more than one value.
// This is useful for the common Go pattern: value, error.

func add(a, b int) int {
	return a + b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("add(3, 4) =", add(3, 4))

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("divide failed:", err)
		return
	}
	fmt.Println("divide(10, 2) =", result)

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("divide(10, 0) ->", err)
		return
	}
	fmt.Println("divide(10, 0) =", result)

	// Go usually returns (value, error) instead of throwing an exception.
	// The caller checks the error immediately, which makes failures explicit.
	fmt.Println("\nError handling pattern:")
	value, err := divide(21, 3)
	if err != nil {
		fmt.Println("unexpected error:", err)
	} else {
		fmt.Println("value:", value)
	}
}
