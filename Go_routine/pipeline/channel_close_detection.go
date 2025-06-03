package pipeline

import (
	"fmt"
	"time"
)

func Channel_detection() {
	ch := make(chan int)

	// Receiver goroutine
	go func() {
		for {
			value, ok := <-ch
			if !ok {
				// Channel is closed
				fmt.Println("Channel is closed. Stopping receiver.")
				break
			}
			fmt.Println("Received:", value)
		}
	}()

	// Send some values
	ch <- 10
	ch <- 20
	ch <- 30

	// Give goroutine some time to receive
	time.Sleep(1 * time.Second)

	// Close the channel
	close(ch)

	// Give time for receiver to finish
	time.Sleep(1 * time.Second)
}
