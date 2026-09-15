package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func increment() {
	mutex.Lock()

	counter++

	mutex.Unlock()
}

func main() {
	var waitGroup sync.WaitGroup

	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()
			increment()
		}()
	}

	waitGroup.Wait()

	fmt.Println("Counter:", counter)
}
