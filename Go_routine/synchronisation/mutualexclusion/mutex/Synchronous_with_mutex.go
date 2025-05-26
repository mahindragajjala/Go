package mutex

import (
	"fmt"
	"sync"
)

/*
In synchronous code, the execution is sequential and you typically don't need a mutex unless multiple goroutines are introduced.

When you do use it with goroutines (to avoid race conditions), it becomes synchronous access to the critical section.
*/

func Synchronous() {
	fmt.Println("SYNCHRONOUS--->")
	var mu sync.Mutex
	var counter int

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock() // 🟡 SYNC: only one goroutine can enter here at a time
			counter++
			fmt.Printf("Goroutine %d increased counter to %d\n", id, counter)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

/*
| Feature            | Synchronous                        | Asynchronous                         |
| ------------------ | ---------------------------------- | ------------------------------------ |
| Real-world analogy | ATM queue                          | Food delivery app                    |
| Go keyword used    | No special keyword                 | `go` keyword to launch goroutine     |
| Behavior           | Wait for one to finish             | Start and move on immediately        |
| Use case           | Critical operations, ordered steps | Background tasks, parallel execution |

*/
