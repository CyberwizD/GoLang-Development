package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Operation completed successfully!") // If the operation completes within 5 seconds.
	case <-ctx.Done():
		fmt.Println("Operation timed out or cancelled!") // If the context is cancelled (e.g., due to a timeout).
	}
}

func main() {
	// A context with a timeout of 3 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	doSomething(ctx)

	// Sleep for 4 seconds which will cause the `doSomething` function to timeout.
	time.Sleep(4 * time.Second)
}
