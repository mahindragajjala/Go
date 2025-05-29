package semaphore

import (
	"fmt"
	"sync"
)

/*
Semaphore-Guarded Pipeline

Task:
Build a 3-stage pipeline (e.g., Read → Process → Write) where each stage is guarded by a semaphore to limit parallelism at each stage.

Goal:
Combine pipeline patterns and semaphores to control stage-wise concurrency.
*/
const (
	numItems         = 25
	genConcurrency   = 2
	procConcurrency  = 4
	writeConcurrency = 3
)

func SemaphoreGuardedPipelineCSP() {
	// Channels between stages
	genOut := make(chan int)
	procOut := make(chan int)

	var wg sync.WaitGroup

	// CSP-style workers using concurrency limits via buffered semaphores
	genLimiter := make(chan struct{}, genConcurrency)
	procLimiter := make(chan struct{}, procConcurrency)
	writeLimiter := make(chan struct{}, writeConcurrency)

	// Generator: Feed input values
	go func() {
		for i := 0; i < numItems; i++ {
			genLimiter <- struct{}{} // limit concurrency
			wg.Add(1)
			go func(val int) {
				defer wg.Done()
				defer func() { <-genLimiter }()
				genOut <- val
			}(i)
		}
		wg.Wait()
		close(genOut) // Close after all generation is done
	}()

	// Processor: Convert and forward values
	go func() {
		var procWg sync.WaitGroup
		for val := range genOut {
			procLimiter <- struct{}{}
			procWg.Add(1)
			go func(v int) {
				defer procWg.Done()
				defer func() { <-procLimiter }()
				if v%2 == 0 {
					fmt.Printf("Processed: %d is divisible by 2\n", v)
					procOut <- v
				} else {
					fmt.Printf("Processed: %d is not divisible by 2\n", v)
					procOut <- 0
				}
			}(val)
		}
		procWg.Wait()
		close(procOut)
	}()

	// Writer: Consume and print
	var writeWg sync.WaitGroup
	for val := range procOut {
		writeLimiter <- struct{}{}
		writeWg.Add(1)
		go func(v int) {
			defer writeWg.Done()
			defer func() { <-writeLimiter }()
			fmt.Printf("Written: %d\n", v)
		}(val)
	}
	writeWg.Wait()
}
