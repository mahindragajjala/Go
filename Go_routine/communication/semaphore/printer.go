package semaphore

import (
	"fmt"
	"time"
)

func Printer() {
	printer := make(chan struct{}, 2)
	for i := 0; i < 5; i++ {
		go func(i int) {
			printer <- struct{}{}
			fmt.Printf("Person:%v is printing\n", i)
			<-printer
		}(i)
	}
	time.Sleep(2 * time.Second)
}

/*
test:


func Printer() {
	printer := make(chan struct{}, 2)

	var concurrent int32
	var maxConcurrent int32
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			printer <- struct{}{}

			current := atomic.AddInt32(&concurrent, 1)
			// Track the max concurrency seen
			atomic.CompareAndSwapInt32(&maxConcurrent, maxConcurrent, max(current, maxConcurrent))

			fmt.Printf("Person:%v is printing | Concurrent: %v\n", i, current)

			time.Sleep(300 * time.Millisecond) // simulate some printing time

			atomic.AddInt32(&concurrent, -1)
			<-printer
		}(i)
	}

	wg.Wait()
	fmt.Printf("Max Concurrent Goroutines: %v\n", maxConcurrent)
}

*/
/*
using CSP

func Printer() {
	printer := make(chan struct{}, 2) // Limits concurrency to 2
	concurrencyTracker := make(chan int) // Used to track how many are running concurrently
	done := make(chan struct{}) // Signals when all work is done

	total := 5
	go func() {
		concurrent := 0
		maxConcurrent := 0
		finished := 0

		for {
			select {
			case delta := <-concurrencyTracker:
				concurrent += delta
				if concurrent > maxConcurrent {
					maxConcurrent = concurrent
				}
			case <-done:
				finished++
				if finished == total {
					fmt.Println("Max Concurrent Goroutines:", maxConcurrent)
					return
				}
			}
		}
	}()

	for i := 0; i < total; i++ {
		go func(i int) {
			printer <- struct{}{}              // Block if 2 already running
			concurrencyTracker <- 1           // Announce start
			fmt.Printf("Person:%v is printing\n", i)
			time.Sleep(300 * time.Millisecond) // Simulate work
			concurrencyTracker <- -1          // Announce finish
			<-printer                          // Free up slot
			done <- struct{}{}
		}(i)
	}

	time.Sleep(3 * time.Second) // Let everything finish
}

*/
