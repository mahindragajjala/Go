package semaphore

import (
	"fmt"
	"time"
)

func BinarySemaphoreExample() {
	fmt.Println("=== Binary Semaphore Example ===")

	semaphore := make(chan struct{}, 1) // Binary semaphore (only 1 token)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			semaphore <- struct{}{} // Acquire token (blocks if already held)
			fmt.Printf("Goroutine %d entered critical section\n", id)

			time.Sleep(2 * time.Second) // Simulate work

			fmt.Printf("Goroutine %d exiting critical section\n", id)
			<-semaphore // Release token
		}(i)
	}

	time.Sleep(7 * time.Second) // Wait for all goroutines to finish
}
