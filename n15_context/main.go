package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Work completed")

	case <-ctx.Done():
		fmt.Println("Work cancelled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)

	defer cancel()

	doWork(ctx)
}
