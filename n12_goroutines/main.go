package main

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(message, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printMessage("Goroutine")

	printMessage("Main")

	time.Sleep(2 * time.Second)
}
