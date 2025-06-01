package mutex

import (
	"fmt"
	"sync"
)

/*
Using defer m.Unlock() in Go is a common and often recommended pattern when working with sync.Mutex.

It ensures that the mutex is always unlocked—even if the function returns early or panics.

However, there are trade-offs in terms of performance and readability.
*/
func Defer_with_mutex() {
	criticalSection()
}

var m sync.Mutex

func criticalSection() {
	m.Lock()
	defer m.Unlock()

	// Do some critical work
	fmt.Println("Inside the critical section")
}
