package mutex

import (
	"fmt"
	"sync"
)

/*
In Go, channels are not meant for mutual exclusion, but you can simulate mutual exclusion using channels by allowing only one goroutine at a time to access a shared resource via channel communication.

Increment a shared counter safely using channels instead of sync.Mutex.
*/
func Mutual_exclusion_using_channel() {
	counter := 0

	// Create a channel with capacity 1 to act as a lock
	lock := make(chan struct{}, 1)

	// Acquire the lock initially so it's available
	lock <- struct{}{}

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			// Acquire the lock by receiving from the channel
			<-lock

			// Critical section
			counter++

			// Release the lock by sending back into the channel
			lock <- struct{}{}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
