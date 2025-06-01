package parentchild

/*
In Go, goroutines are lightweight threads managed by the Go runtime.

"""When one goroutine spawns another, the one spawning is often considered the parent and the one spawned is the child."""

While Go doesn’t enforce parent-child relationships like some systems (e.g., Unix processes), you can build structured relationships using Go's concurrency primitives.
*/
/*
✅ Parent Goroutine
The main goroutine or any goroutine that creates/spawns another goroutine.
Responsible for controlling or managing the child goroutine (e.g., via channels or context.Context).

✅ Child Goroutine
A goroutine started by another goroutine.
Typically performs a specific task and communicates back to the parent.
*/
/*
| Topic                    | Description                                                       |
| ------------------------ | ----------------------------------------------------------------- |
| 1. Basic spawning        | Using `go` keyword to start a child goroutine                     |
| 2. Communication         | Using channels to communicate between parent and child            |
| 3. Synchronization       | Waiting for child goroutines using `sync.WaitGroup`               |
| 4. Cancellation          | Using `context.Context` to cancel child goroutines                |
| 5. Timeout handling      | Child goroutine stops after a timeout using `context.WithTimeout` |
| 6. Error propagation     | Returning errors from child to parent                             |
| 7. Resource sharing      | Managing shared resources safely                                  |
| 8. Nested goroutines     | Child goroutines spawning their own children                      |
| 9. Panic handling        | Recovering from child goroutine panics                            |
| 10. Lifecycle management | Ensuring proper startup and shutdown of child goroutines          |

*/
