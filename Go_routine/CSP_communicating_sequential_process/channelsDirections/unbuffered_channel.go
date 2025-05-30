package channelsdirections

import (
	"fmt"
	"sync"
	"time"
)

func Unbuffered_channel() {
	// Unbuffered channel
	unbufferedChan := make(chan int)

	// Semaphore (buffered channel of size 2)
	signal := make(chan struct{}, 2)

	// WaitGroup
	var wg sync.WaitGroup

	// Start a goroutine to receive and print data
	go func() {
		for val := range unbufferedChan {
			fmt.Println("channel data:", val)
		}
	}()

	// Launch goroutines
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go goroutineUnbuffered(signal, &wg, i, unbufferedChan)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the channel after all goroutines are done
	close(unbufferedChan)
}

// Goroutine
func goroutineUnbuffered(signal chan struct{}, wg *sync.WaitGroup, i int, ch chan int) {
	defer wg.Done()

	// Acquire lock
	signal <- struct{}{}

	// Critical section
	ch <- i
	time.Sleep(time.Second)

	// Release lock
	<-signal
}

/*
🔁 Execution Flow:
The main function waits until all goroutines are done writing.

Meanwhile, the separate goroutine is reading.

After the writers are done, the main function closes the channel.

The reader detects closure and exits the loop.
*/
