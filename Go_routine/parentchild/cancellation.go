package parentchild

import (
	"context"
	"fmt"
	"time"
)

func child(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Child received cancel signal")
			return
		default:
			fmt.Println("Child working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func CANCELLATION() {
	ctx, cancel := context.WithCancel(context.Background())
	go child(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("Parent exiting")
}
