package semaphore

import "time"

/*
Buffered channel as a semaphore:
Use a buffered channel of size 2 as a semaphore to limit concurrency.
Launch 5 goroutines that try to enter a critical section.
Each goroutine should acquire the semaphore (send to channel), print a message, sleep 1 second, then release (receive from channel).
*/

func Semaphore() {
	Buffered := make(chan struct{}, 2)
	go func() {
		//acqiure the lock
		Buffered <- struct{}{}

		//critical section
		time.Sleep(time.Second)

		//release semaphore
		<-Buffered
	}()
	go func() {
		//acqiure the lock
		Buffered <- struct{}{}

		//critical section
		time.Sleep(time.Second)

		//release semaphore
		<-Buffered
	}()
	go func() {
		//acqiure the lock
		Buffered <- struct{}{}

		//critical section
		time.Sleep(time.Second)

		//release semaphore
		<-Buffered
	}()
	go func() {
		//acqiure the lock
		Buffered <- struct{}{}

		//critical section
		time.Sleep(time.Second)

		//release semaphore
		<-Buffered
	}()
	go func() {
		//acqiure the lock
		Buffered <- struct{}{}

		//critical section
		time.Sleep(time.Second)

		//release semaphore
		<-Buffered
	}()
}

/*
Counting Semaphore
A counting semaphore maintains a set number of "tokens" (or "permits").
If a goroutine wants to access a resource, it must acquire a token.
If no tokens are left, it waits (blocks).
After finishing the job, the goroutine releases the token.

Tokens
A token is like a permission slip.
Each goroutine needs one to proceed.
Number of tokens = number of goroutines allowed to enter simultaneously.

Critical Section
A critical section is a part of the code that accesses shared data or resources that must not be accessed concurrently by more than a fixed number of goroutines.
*/

//semaphore := make(chan struct{}, 2)
/*
make(chan struct{}, 2) creates a buffered channel of size 2.
The type is struct{} — an empty struct in Go that takes 0 bytes of memory, perfect for signaling (no data passed).
Think of this channel as holding 2 tokens.
*/

/*
How it works as a Semaphore
semaphore <- struct{}{}   // acquire a token (wait if full)
defer func() {
    <-semaphore          // release the token
}()

Working Flow:
A goroutine sends an empty struct into the channel to acquire a token.
If the channel has space (buffered size not full), it proceeds.
If the channel is full, it waits (blocks) until a token is released.
When done, the goroutine receives from the channel (<-semaphore) to release the token.
*/

//Different types of the semaphore
/*
Binary Semaphore
A binary semaphore is a synchronization primitive that has only two states — 0 (locked) or 1 (unlocked). It is typically used to ensure mutual exclusion to a critical section, meaning only one goroutine or thread can access a resource at a time.

Key Characteristics:
Acts like a mutex or lock.
Only one thread/goroutine can acquire the semaphore at a time.
Represented by a buffered channel of size 1 in Go.

Controlling access to a shared resource like a counter, database, or file — where only one goroutine should access it at any moment.

Counting Semaphore
A counting semaphore is a synchronization primitive that maintains a count (an integer) "which represents the number of available permits" (tokens). It allows up to N goroutines/threads to access a shared resource concurrently.

Key Characteristics:
Generalization of the binary semaphore.
Allows multiple threads/goroutines (up to the limit) to enter the critical section simultaneously.
Represented in Go by a buffered channel with capacity = N.

Limiting the number of concurrent:
Database connections (e.g., 5 connections max)
Worker goroutines (e.g., process 3 jobs at a time),
API calls (e.g., throttle 10 requests per second).

| Feature             | Binary Semaphore               | Counting Semaphore               |
| - |  | -- |
| Capacity            | 1 (either 0 or 1)              | Any integer ≥ 1 (e.g., 3, 5, 10) |
| Concurrency allowed | Only 1 goroutine               | Up to N goroutines               |
| Use case            | Mutual exclusion (like a lock) | Rate limiting / resource control |
| Go implementation   | `make(chan struct{}, 1)`       | `make(chan struct{}, N)`         |

*/

/*
RATE LIMITING
You're running an API Gateway that handles requests from thousands of users. Your backend server can safely handle only 10 requests per second. If more requests come in, the server gets overwhelmed and crashes.

solution :
You implement rate limiting at the gateway using a counting semaphore with 10 tokens (one for each allowed request per second).


*/
/*
Here are 10 coding questions based on Binary and Counting Semaphores using Buffered Channels in Go — from beginner to advanced level — with a brief description and goal for each:



 🔰 Basic Level

 1. Basic Binary Semaphore Lock

Task:
Write a Go program where multiple goroutines increment a shared counter using a binary semaphore to ensure mutual exclusion.

Goal:
Prevent race conditions using a buffered channel of size 1.



 2. Simple Counting Semaphore

Task:
Create 5 goroutines accessing a shared printer, but only 2 goroutines should print at the same time.

Goal:
Use a counting semaphore (channel buffer size 2) to simulate limited access.



 3. Semaphore as Connection Pool

Task:
Simulate a connection pool with 3 available connections. Launch 10 goroutines; each should acquire a connection, simulate a DB call, and then release it.

Goal:
Use a counting semaphore to simulate limited resources (connections).

DO IT LATER:)

 ⚙️ Intermediate Level

 4. Bounded Worker Pool

Task:
Create a worker pool using a counting semaphore to limit the number of concurrent jobs running to 4, even if 20 jobs are submitted.

Goal:
Limit goroutine concurrency using semaphores instead of wait groups or worker channels.

 5. Timeout With Semaphore

Task:
Modify the semaphore pattern to support a timeout when acquiring a token using `select` and `time.After`.

Goal:
Demonstrate non-blocking or timed access to limited resources.

[LOCK / UNLOCK - INSIDE THE GOROUTINE ]

 6. Read-Write Simulation

Task:
Create 5 reader and 3 writer goroutines. Only 1 writer can write at a time (binary semaphore), but 2 readers can read concurrently (counting semaphore with size 2).

Goal:
Use both binary and counting semaphores to model a simplified read-write lock.



 7. Rate Limiter Using Semaphore

Task:
Allow only 5 requests per second using a semaphore. Each request is handled by a goroutine. If 5 tokens are used, new requests wait until tokens are freed.

Goal:
Throttle request handling using a counting semaphore.



 🚀 Advanced Level

 8. Semaphore-Guarded Pipeline

Task:
Build a 3-stage pipeline (e.g., Read → Process → Write) where each stage is guarded by a semaphore to limit parallelism at each stage.

Goal:
Combine pipeline patterns and semaphores to control stage-wise concurrency.



 9. Semaphore With Panic Recovery

Task:
Simulate a critical section with a counting semaphore. Some goroutines may panic. Ensure the token is released even if panic occurs.

Goal:
Use `defer` and `recover()` to handle panics and safely release the semaphore.



 10. Custom Semaphore Type in Go

Task:
Create a reusable `Semaphore` struct with `Acquire()` and `Release()` methods internally backed by a `chan struct{}`.

Goal:
Encapsulate semaphore behavior in a clean API, improving code reuse and readability.



*/
