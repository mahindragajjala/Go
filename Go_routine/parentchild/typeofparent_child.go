package parentchild

/*
In Go, goroutines are lightweight threads managed by the Go runtime. When one goroutine spawns another, the one spawning is often considered the parent and the one spawned is the child. While Go doesn’t enforce parent-child relationships like some systems (e.g., Unix processes), you can build structured relationships using Go's concurrency primitives.

---

 🔸 Core Concept: Parent and Child Goroutines

 ✅ Parent Goroutine

* The main goroutine or any goroutine that creates/spawns another goroutine.
* Responsible for controlling or managing the child goroutine (e.g., via channels or `context.Context`).

 ✅ Child Goroutine

* A goroutine started by another goroutine.
* Typically performs a specific task and communicates back to the parent.

---

 📌 Table of Topics on Parent-Child Goroutines

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

---

 🔹 1. Basic Spawning

go
package main

import "fmt"

func main() {
    fmt.Println("Parent started")

    go func() {
        fmt.Println("Child goroutine running")
    }()

    fmt.Println("Parent finished")
}


Output is nondeterministic due to concurrency.

---

 🔹 2. Communication (using Channels)

go
package main

import "fmt"

func main() {
    ch := make(chan string)

    go func() {
        ch <- "Hello from child"
    }()

    msg := <-ch
    fmt.Println("Parent received:", msg)
}


---

 🔹 3. Synchronization (using `sync.WaitGroup`)

go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(1)

    go func() {
        defer wg.Done()
        fmt.Println("Child goroutine finished")
    }()

    wg.Wait()
    fmt.Println("Parent exiting")
}


---

 🔹 4. Cancellation (using `context.WithCancel`)

go
package main

import (
    "context"
    "fmt"
    "time"
)

func child(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Child received cancel signal")
            return
        default:
            fmt.Println("Child working...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    go child(ctx)

    time.Sleep(2 * time.Second)
    cancel()
    time.Sleep(1 * time.Second)
    fmt.Println("Parent exiting")
}


---

 🔹 5. Timeout Handling (using `context.WithTimeout`)

go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    go func(ctx context.Context) {
        for {
            select {
            case <-ctx.Done():
                fmt.Println("Child timed out")
                return
            default:
                fmt.Println("Child still running...")
                time.Sleep(500 * time.Millisecond)
            }
        }
    }(ctx)

    time.Sleep(3 * time.Second)
    fmt.Println("Parent exiting")
}


---

 🔹 6. Error Propagation

Using channels to send errors from child to parent:

go
package main

import (
    "errors"
    "fmt"
)

func child(errCh chan error) {
    // simulate error
    errCh <- errors.New("something went wrong in child")
}

func main() {
    errCh := make(chan error)

    go child(errCh)

    err := <-errCh
    fmt.Println("Parent received error:", err)
}


---

 🔹 7. Resource Sharing (Avoiding Race Conditions)

go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex
    counter := 0

    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        mu.Lock()
        counter++
        mu.Unlock()
    }()

    go func() {
        defer wg.Done()
        mu.Lock()
        counter++
        mu.Unlock()
    }()

    wg.Wait()
    fmt.Println("Counter:", counter)
}


---

 🔹 8. Nested Goroutines

go
package main

import "fmt"

func main() {
    go func() {
        fmt.Println("Child started")

        go func() {
            fmt.Println("Grandchild running")
        }()
    }()

    fmt.Scanln() // wait for input to see output
}


---

 🔹 9. Panic Handling in Child

go
package main

import (
    "fmt"
    "time"
)

func main() {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println("Recovered in child:", r)
            }
        }()
        panic("Child panic!")
    }()

    time.Sleep(1 * time.Second)
    fmt.Println("Parent exiting")
}


---

 🔹 10. Lifecycle Management Pattern

go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Child gracefully stopped")
            return
        default:
            fmt.Println("Child working...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    var wg sync.WaitGroup

    wg.Add(1)
    go worker(ctx, &wg)

    time.Sleep(2 * time.Second)
    cancel()

    wg.Wait()
    fmt.Println("Parent exited after child cleanup")
}


---

 ✅ Summary

| Feature              | Tool Used           |
| -------------------- | ------------------- |
| Communication        | Channels            |
| Sync                 | `sync.WaitGroup`    |
| Timeout/Cancel       | `context`           |
| Safe access          | `sync.Mutex`        |
| Panic recovery       | `recover()`         |
| Lifecycle management | Context + WaitGroup |

---

Would you like a real-time mini project demonstrating parent-child goroutines, like a web server handling requests and child goroutines processing them?

*/
