package semaphore

import (
	"fmt"
	"time"
)

func CountingSemaphoreExample() {
	fmt.Println("\n=== Counting Semaphore Example ===")

	semaphore := make(chan struct{}, 2) // Counting semaphore (2 tokens)

	for i := 1; i <= 5; i++ {
		go func(id int) {
			semaphore <- struct{}{} // Acquire token (wait if full)
			fmt.Printf("Goroutine %d entered critical section\n", id)

			time.Sleep(2 * time.Second) // Simulate work

			fmt.Printf("Goroutine %d exiting critical section\n", id)
			<-semaphore // Release token
		}(i)
	}

	time.Sleep(10 * time.Second) // Wait for all goroutines to finish
}
