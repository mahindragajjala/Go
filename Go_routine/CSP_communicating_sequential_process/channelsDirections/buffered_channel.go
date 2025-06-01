package channelsdirections

import (
	"fmt"
	"sync"
	"time"
)

func Buffered_channel() {
	var wg sync.WaitGroup

	csp_lock := make(chan struct{}, 2)
	Buffered_channel := make(chan string, 5)

	// Add to WaitGroup BEFORE launching goroutines
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go Goroutine(csp_lock, Buffered_channel, i, &wg)
	}

	// Close the channel once all goroutines finish writing
	go func() {
		wg.Wait()
		close(Buffered_channel)
	}()

	// Read from the channel
	for data := range Buffered_channel {
		fmt.Println("data:", data)
	}
}

func Goroutine(csp_lock chan struct{}, Buffered_channel chan string, i int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Acquire lock (max 2 goroutines)
	csp_lock <- struct{}{}

	// Simulate work and send to channel
	data := fmt.Sprintf("Go Routine %v", i)
	Buffered_channel <- data
	time.Sleep(2 * time.Second)

	// Release lock
	<-csp_lock
}

/*
🔁 Execution Flow:
You launch goroutines to write.

You immediately start a reader loop in the main function.

Meanwhile, you wait in a separate goroutine (go func() { ... }) to detect when all writes are done and then close the channel.
*/
