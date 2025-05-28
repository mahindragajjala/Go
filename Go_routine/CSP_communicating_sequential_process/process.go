package cspcommunicatingsequentialprocess

/*
Goroutines (Processes in CSP)
						Creating goroutines: go func()
						Goroutine lifecycle
						Goroutine scheduling by Go runtime (M:N scheduler)
						Memory model and stack management
*/
/*
goroutine life cycle :

| Stage       | Description                                                      |
|-------------| -------------------------------------------------------------------- |
| Creation    | Created using `go func()` or `go someFunction()`                     |
| Runnable    | Added to the run queue by the Go scheduler                           |
| Running     | Gets picked up by a worker thread (`M`) and executed                 |
| Blocked     | Waiting for I/O, channel operation, sleep, etc.                      |
| Ready Again | After blocking ends (e.g., channel receives), goes back to run queue |
| Terminated  | Function completes execution or panics                               |

*/
/*
Goroutine Scheduling by Go Runtime (M:N Scheduler)

Go uses an M:N scheduler to map M "OS threads" to N "goroutines".
| Symbol | Meaning                                       |
| ------ | --------------------------------------------- |
| G      | Goroutine                                     |
| M      | OS Thread                                     |
| P      | Processor (Logical context to run goroutines) |

 Scheduler Model:

Goroutines (G) -> Queued to Processor (P)
Processor (P) -> Assigned to OS Thread (M)
M: Represents an actual OS thread.
P: Provides a context to execute goroutines. There are usually GOMAXPROCS number of Ps.
G: Represents the actual goroutine.

⚙ How It Works:
Each P has its local run queue.
When a goroutine becomes ready, it's added to a P's run queue.
If P's run queue is empty, it steals from other Ps (work stealing).
If M is idle and there's work, it gets a P and starts executing goroutines.

🔁 Preemption:
Preemptive scheduling added in Go 1.14.
"Long-running goroutines can be paused to give others a chance to run.
"
*/
/*
Memory Model and Stack Management

Stack Allocation:

Each goroutine starts with a tiny stack (~2 KB).
Unlike threads (which pre-allocate large stacks), goroutine stacks grow and shrink dynamically.

// Goroutine stack starts small and grows as needed.
Garbage Collection Friendly:

Since stacks are small and grow on demand, memory pressure is low.
Helps Go scale to hundreds of thousands of goroutines.

Escape Analysis:
Determines whether variables should be allocated on the stack or heap.
If a variable “escapes” a function’s scope (e.g., returned or captured by closure), it's heap-allocated.
*/
