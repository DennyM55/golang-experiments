package main

import "fmt"

func sendMessage(ch chan string) {
	ch <- "Hello from goroutine"
}

func main() {
	ch := make(chan string)

	go sendMessage(ch)

	message := <-ch

	fmt.Println(message)
}
