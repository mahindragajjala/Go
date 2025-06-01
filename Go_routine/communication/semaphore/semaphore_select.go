package semaphore

import (
	"fmt"
	"time"
)

// Semaphore with capacity of 2 tokens
var semaphore = make(chan struct{}, 2)

// acquire tries to get a token from the semaphore within the timeout duration.
// Returns true if token acquired, false if timed out.
func acquire(timeout time.Duration) bool {
	select {
	case semaphore <- struct{}{}:
		// Got token
		return true
	case <-time.After(timeout):
		// Timeout happened before acquiring token
		return false
	}
}

// release returns a token to the semaphore.
func release() {
	<-semaphore
}

func worker(id int, timeout time.Duration) {
	fmt.Printf("Worker %d: Trying to acquire semaphore...\n", id)
	if acquire(timeout) {
		fmt.Printf("Worker %d: Acquired semaphore!\n", id)

		// Simulate work
		time.Sleep(2 * time.Second)

		release()
		fmt.Printf("Worker %d: Released semaphore.\n", id)
	} else {
		fmt.Printf("Worker %d: Timeout! Could not acquire semaphore.\n", id)
	}
}

func Semaphore_Select() {
	// Start 4 workers trying to acquire a semaphore with a 1-second timeout
	for i := 1; i <= 4; i++ {
		go worker(i, 1*time.Second)
	}

	// Wait enough time for workers to finish
	time.Sleep(5 * time.Second)
}

// simple
// Routine tries to acquire a token from semaphore with timeout,
// does work if successful, then releases the token.
func Routine(id int, sem chan struct{}, timeout time.Duration) {
	fmt.Printf("Routine %d: trying to acquire semaphore\n", id)

	select {
	case sem <- struct{}{}: // Acquire token
		fmt.Printf("Routine %d: acquired semaphore\n", id)
		time.Sleep(3 * time.Second) // simulate work
		<-sem                       // Release token
		fmt.Printf("Routine %d: released semaphore\n", id)

	case <-time.After(timeout): // Timeout waiting for token
		fmt.Printf("Routine %d: timeout while waiting for semaphore\n", id)
	}
}

func Time_semaphore() {
	// Create a semaphore with capacity 4
	sem := make(chan struct{}, 4)

	// Start 10 routines competing for 4 tokens
	for i := 1; i <= 10; i++ {
		go Routine(i, sem, 2*time.Second)
	}

	// Wait long enough for routines to finish
	time.Sleep(15 * time.Second)
}
