package done

import (
	"context"
	"fmt"
	"time"
)

/*
Using done with context.Context (Best Practice)
context.WithCancel, context.WithTimeout, context.WithDeadline

Using ctx.Done() in place of custom done channels.
*/
func contextWorker(ctx context.Context, ch <-chan int) {
	for {
		select {
		case val := <-ch:
			fmt.Println("Received:", val)
		case <-ctx.Done():
			fmt.Println("Context cancelled:", ctx.Err())
			return
		}
	}
}

func Done_function_with_Context_package() {
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go contextWorker(ctx, ch)

	ch <- 10
	time.Sleep(1 * time.Second)

	// Cancel the context to stop the goroutine
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("Main exits")
}
