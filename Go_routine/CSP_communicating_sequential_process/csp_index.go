package cspcommunicatingsequentialprocess

/*

🔰 1. Fundamentals of CSP in Go
						What is CSP (Communicating Sequential Processes)
						Why Go uses CSP model
						Go’s concurrency motto: “Don’t communicate by sharing memory...”

🧵 2. Goroutines (Processes in CSP)
						Creating goroutines: go func()
						Goroutine lifecycle
						Goroutine scheduling by Go runtime (M:N scheduler)
						Memory model and stack management

🔗 3. Channels (Communication in CSP)
						Creating channels: make(chan T)
						Sending and receiving: ch <- val, val := <-ch
						Blocking behavior of unbuffered channels
						Closing channels: close(ch)
						Checking closed channels: v, ok := <-ch

🧰 4. Channel Types and Directions
						Unbuffered vs Buffered channels
						make(chan int) vs make(chan int, 10)
						Receive-only and Send-only channels
						chan<- int (send-only)
						<-chan int (receive-only)
						Nil channels and blocking behavior

🔄 5. Select Statement (Non-deterministic Communication)
						Using select with multiple channels
						default case to avoid blocking
						select with timeout using time.After
						Avoiding deadlocks and starvation

🔁 6. CSP Communication Patterns
						One-to-one communication
						One-to-many (fan-out)
						Many-to-one (fan-in)
						Pipelines (chaining goroutines)
						Multiplexing input from multiple channels
						Worker pools
						Tee pattern (splitting one channel into many)

🧱 7. Synchronization with Channels
						Using channels instead of sync.Mutex
						Using empty struct channels: chan struct{}
						Signaling between goroutines
						Waiting for completion with channels
🧪 8. Advanced CSP Topics
						Channel direction in function parameters
						Buffered channel capacity and deadlock issues
						Backpressure in pipelines
						Leak prevention (avoiding goroutines stuck on send/receive)
						Using context.Context for goroutine cancellation

🚨 9. Common Pitfalls in Go CSP
Goroutine leaks

Reading from closed channels

Writing to closed channels (panic)

Blocking on nil channels

Race conditions when misusing shared memory

🧩 10. Real-World Use Cases and Patterns
Rate limiting using channels

Bounded concurrency with semaphore pattern

Publish-subscribe model with CSP

Task queues with channel buffers

CSP-style state machines

🧠 Bonus: CSP Theoretical Background
Tony Hoare’s CSP theory (for deep dives)

Comparison with:

Shared memory (mutexes)

Actor model (Erlang, Akka)

Async/await (JavaScript, Python)

Message queues (Kafka, RabbitMQ)
*/
/*
Link :- https://chatgpt.com/share/6836bbc6-2598-800b-a734-670f5422e1b1
*/
