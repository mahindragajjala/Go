package semaphore

import (
	"fmt"
	"sync"
	"time"
)

/*
Go’s standard library does not provide a built-in semaphore type, but you can easily implement it using channels.

*/

// Semaphore struct
type Semaphore struct {
	tokens chan struct{}
}

// NewSemaphore creates a new semaphore with a given capacity
func NewSemaphore(capacity int) *Semaphore {
	return &Semaphore{
		tokens: make(chan struct{}, capacity),
	}
}

// Acquire reserves a token or waits if none available
func (s *Semaphore) Acquire() {
	s.tokens <- struct{}{} // Will block if capacity reached
}

// Release releases a token back to the semaphore
func (s *Semaphore) Release() {
	<-s.tokens
}

func worker(id int, sem *Semaphore) {
	sem.Acquire()
	fmt.Printf("Worker %d: started\n", id)
	time.Sleep(2 * time.Second) // Simulate work - we can see what work is happening...
	fmt.Printf("Worker %d: finished\n", id)
	sem.Release()
}

func Semaphore_function() {
	sem := NewSemaphore(3) // Allow max 3 concurrent workers

	for i := 1; i <= 5; i++ {
		go worker(i, sem)
	}

	// Wait for enough time to let all workers finish
	time.Sleep(10 * time.Second)
}

/*
How is this similar to a Mutex?

A mutex is like a semaphore with capacity = 1.

Mutex: only one goroutine can hold the lock at once.
Semaphore: N goroutines can access simultaneously.
*/
/*
DEFINITION OF THE SEMAPHORE :
A semaphore is a synchronization primitive used to control access to a shared resource in a concurrent system such as a multithreaded program.
It uses an internal counter to track how many units of a resource are available.

REAL TIME EXAMPLES :
A semaphore is like a "traffic controller" that helps manage how many people (or programs) can use something at the same time.
It keeps a count of how many "slots" or "spaces" are available.
When the slots are full, others have to wait until a slot becomes free.

Real-Life Example:
Parking Lot 🅿️
Imagine a parking lot with 3 parking spots.
If 3 cars are parked, any new car that comes must wait outside.
When a car leaves, a new car can enter.
Here, the parking lot is using a semaphore with a value of 3 to control how many cars can park at the same time.
*/
/*
DIFFERENT TYPES OF THE SEMAPHORES :
Counting Semaphore - Allows any number of threads (counter can be > 1). Used for resource pools.
Binary Semaphore	Only allows one thread at a time (counter is 0 or 1). Used as a mutex (lock).

*/
/*
Imagine a parking lot that has only 3 parking spaces.

Cars are coming (threads/goroutines).
Parking slots are limited (shared resource).
We want only 3 cars to park at a time.

A counting semaphore with a value of 3 can control access to these 3 parking slots.
*/

/*
1. Semaphore Counter
Initialized to a value (N) = number of resources.

Each time a thread enters:
wait() or acquire() → decrements the counter.
If counter < 0 → thread waits.
Each time a thread exits:

signal() or release() → increments the counter.
If any thread is waiting, one is unblocked.

2. Atomic Operations
The semaphore ensures that updates to the counter are atomic, preventing race conditions.

3. Blocking and Waking
Threads trying to enter when counter is 0 will block.

When another thread leaves and calls release(), one blocked thread is woken up.
*/

//EXAMPLE OF THE SEMAPHORE
/* type Semaphore struct {
	ch chan struct{}
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, n),
	}
} */

/*
	 func (s *Semaphore) Acquire() {
		s.ch <- struct{}{}
	}

	func (s *Semaphore) Release() {
		<-s.ch
	}
*/
func parkCar(id int, sem *Semaphore, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("🚗 Car %d is trying to park\n", id)

	sem.Acquire() // occupy one slot
	fmt.Printf("✅ Car %d has parked\n", id)
	time.Sleep(2 * time.Second) // simulate parking duration

	fmt.Printf("🅿️ Car %d is leaving\n", id)
	sem.Release() // free the slot
}

func main() {
	var wg sync.WaitGroup
	parkingSlots := 3
	sem := NewSemaphore(parkingSlots)

	totalCars := 6
	for i := 1; i <= totalCars; i++ {
		wg.Add(1)
		go parkCar(i, sem, &wg)
	}

	wg.Wait()
}
