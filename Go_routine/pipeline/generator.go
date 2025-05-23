package pipeline

import (
	"fmt"
	"time"
)

// repeatFunc repeatedly sends values returned by fn into a channel,
// until it receives a signal on the done channel.
func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

// Example usage
func GENERATOR() {
	done := make(chan struct{})
	fn := func() int {
		return 42
	}

	// Start the generator
	stream := repeatFunc(done, fn)

	// Read a few values and then stop
	for i := 0; i < 5; i++ {
		fmt.Println(<-stream)
	}

	// Signal the generator to stop
	close(done)

	// Optional: wait a bit to ensure the goroutine finishes
	time.Sleep(100 * time.Millisecond)
}

/*
The repeatFunc is a generic generator function that keeps pushing values into a channel until cancellation is requested.
It uses a goroutine, select, and generics for type flexibility and concurrent value streaming.
Very useful in pipelines, streaming, and concurrent data generation.
*/
