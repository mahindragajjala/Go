package cspcommunicatingsequentialprocess

import (
	"fmt"
	"sync"
)

/*
Explaination:

CSP stands for "Communicating Sequential Processes", a formal language for describing "patterns of interaction in concurrent systems".

It was introduced by Tony Hoare in 1978.

🔍 Core Ideas of CSP:

Processes are independent: Each process runs sequentially and independently.
Communication happens only via "message passing", not by "shared memory".
Channels are used for "synchronous communication between processes".
Processes wait (block) until the message is received or sent — "promoting safe interaction".
*/
/*
✅ Why Go Uses the CSP Model
Go was designed by "Google engineers" to solve real-world problems in systems programming and networked services, "where concurrency is critical".

📌 Go uses CSP because:
Simplicity:

Threads and mutexes can be error-prone and hard to reason about.
CSP promotes simple, structured communication patterns between goroutines.
Safety:

By avoiding shared memory, Go avoids race conditions and makes programs more predictable.

Scalability:

Go’s lightweight goroutines scale better than OS threads.
CSP allows efficient coordination between thousands (or more) of goroutines.

Readability and Maintainability:
Go code using channels looks cleaner and is easier to understand.
It encourages design where logic and synchronization are clearly separated.
*/
/*
✅ Go’s Concurrency Motto
“Don’t communicate by sharing memory; share memory by communicating.”

🔁 Traditional Model (e.g., C/Java):

Multiple threads share memory and use locks (mutex, semaphore) to coordinate.
This often leads to:
Race conditions
Deadlocks

Complex debugging

🧠 Go’s CSP Approach:
"Instead of sharing variables between goroutines, you send data over "channels".
This ensures only one goroutine has access to a "piece of data at any time"."
*/

// TRADITIONAL
var count int
var mu sync.Mutex

func Increment() {
	mu.Lock()
	count++
	mu.Unlock()
}

// Go's CSP (with channels)
func counter(ch chan int) {
	count := 0
	for {
		ch <- count
		count++
	}
}

func CSP() {
	ch := make(chan int)
	go counter(ch)
	fmt.Println(<-ch) // 0
	fmt.Println(<-ch) // 1
}
