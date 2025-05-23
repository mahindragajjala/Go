package parentchild

import (
	"context"
	"fmt"
	"time"
)

func TIMEOUT_CANCELLATION() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Child timed out")
				return
			default:
				fmt.Println("Child still running...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("Parent exiting")
}
