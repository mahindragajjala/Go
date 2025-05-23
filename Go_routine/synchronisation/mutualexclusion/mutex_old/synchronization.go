package mutex

import (
	"fmt"
	"sync"
	"time"
)

/*
Synchronization in programming is the process of coordinating the execution of multiple threads or goroutines so that they can safely share resources like memory, files, or databases without interfering with each other.

When multiple threads or goroutines run concurrently, they may try to:

Read/write the same variable or file at the same time
Access shared data without coordination
This can lead to race conditions, data corruption, crashes, or inconsistent results.
*/

var counter int

func increment() {
	for i := 0; i < 1000; i++ {
		counter++ // Not thread-safe
	}
}

func RACE_CONDITION() {

	go increment()
	go increment()
	time.Sleep(1 * time.Second)
	fmt.Println("Counter:", counter) // Output will be unpredictable!
}

/*
Synchronization (Using sync.Mutex)
*/
var mu sync.Mutex
var counter_SYNC int

func increment_SYNC() {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func SYNCHRONIZATION() {
	go increment_SYNC()
	go increment_SYNC()
	time.Sleep(1 * time.Second)
	fmt.Println("Counter:", counter_SYNC)
}

/*
package mutex provides basic synchronization primitives such as mutual exclusion locks.

Other than the Once and WaitGroup types, most are intended for use by low-level library routines.

Higher-level synchronization is better done via channels and communication.
*/

/*
| Use Case                            | Why Synchronization is Needed                   |
| ----------------------------------- | ----------------------------------------------- |
| Shared counters / state updates     | Prevent race conditions                         |
| Logging from multiple threads       | Avoid interleaved or corrupted logs             |
| Concurrent database access          | Ensure data consistency and prevent dirty reads |
| Goroutine coordination              | Ensure goroutines finish in the correct order   |
| Producer-consumer problems          | Wait for resources to be available safely       |

*/
/*
| Tool             | Purpose                                        |
| ---------------- | ---------------------------------------------- |
| `sync.Mutex`     | Locking for exclusive access                   |
| `sync.RWMutex`   | Read-write locking for multiple readers        |
| `sync.WaitGroup` | Waiting for multiple goroutines to complete    |
| `sync.Cond`      | Waiting for and signaling conditions           |
| `channels`       | Built-in tool for safe goroutine communication |

*/
