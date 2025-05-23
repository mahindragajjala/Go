package mutex

import (
	"fmt"
	"sync"
)

/*
What is a Deadlock?
A deadlock is a situation where two or more goroutines are waiting for each other to release resources (like locks), but none of them can proceed, causing the program to hang indefinitely.

🧠 Real-World Analogy:
Imagine two people trying to pass each other in a narrow hallway:

Person A moves to the right.

Person B moves to the left.

Now both block each other's way, waiting for the other to move, resulting in a deadlock.
*/
/*
How Can Deadlocks Occur with sync.Mutex?
🔐 sync.Mutex:
A Mutex (mutual exclusion) is used to protect shared data.

Only one goroutine can lock it at a time.

Others trying to lock it will block (wait) until it's unlocked.
*/

//DEAD LOCKS HAPPEN BECAUSE OF...
/* 1
Double Locking the Same Mutex

var mu sync.Mutex
func main() {
    mu.Lock()
    mu.Lock() // This causes deadlock! Second lock waits forever.
    mu.Unlock()
    mu.Unlock()
}
*/
/* 2
💣 Deadlock Example 2: Goroutines Waiting on Each Other’s Locks

var mu1 sync.Mutex
var mu2 sync.Mutex

func main() {
    go func() {
        mu1.Lock()
        defer mu1.Unlock()

        time.Sleep(1 * time.Second)
        mu2.Lock()
        defer mu2.Unlock()
    }()

    go func() {
        mu2.Lock()
        defer mu2.Unlock()

        time.Sleep(1 * time.Second)
        mu1.Lock()
        defer mu1.Unlock()
    }()

    time.Sleep(3 * time.Second)
}

Goroutine 1 locks mu1 and waits for mu2.
Goroutine 2 locks mu2 and waits for mu1.
Neither can proceed: deadlock.
*/
var mu1 sync.Mutex
var mu2 sync.Mutex
var mu sync.Mutex
var Data int = 0

//TO AVOID DEADLOCKS

// Always Lock in a Consistent Order
func Consist_order() {

	// Always lock mu1 before mu2 everywhere
	mu1.Lock()
	mu2.Lock()
	// Do work
	Data++
	mu2.Unlock()
	mu1.Unlock()

}

// Use defer mu.Unlock() Right After Locking
func Defer_unlock() {
	mu.Lock()
	defer mu.Unlock()
	// Safe critical section

}

//Avoid Nested Locks Whenever Possible
/* Nested locks are complex and more error-prone. Simplify logic to use one lock if possible. */

//Use TryLock Pattern with Channels or Flags
/*
Go’s sync.Mutex does not support TryLock, but you can implement similar logic using select with channels or use sync.Mutex alternatives like sync.RWMutex, sync.Once, or atomic.
*/
func TryLock() {
	// Using a channel as a binary semaphore
	lock := make(chan struct{}, 1)
	lock <- struct{}{} // Acquire

	select {
	case <-lock:
		// Critical section
		lock <- struct{}{} // Release
	default:
		fmt.Println("Lock is busy")
	}

}

/*
Keep Critical Sections Small
Do only the absolutely necessary operations while holding the lock. This reduces contention and chance of deadlocks.
*/

/*
Avoid Blocking or Sleeping Inside Locked Sections
Never call time.Sleep, select, or other blocking code inside a locked section unless necessary.
*/
