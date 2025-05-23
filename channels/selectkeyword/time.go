package selectkeyword

import (
	"fmt"
	"time"
)

/*
	Using `select` with `time.After` (Timeouts)

* Implementing timeouts on channel operations.
* Preventing goroutines from waiting forever.(leaks)

Using select with time.After is a powerful way in Go to implement timeouts on channel operations and ensure goroutines don’t wait forever, especially when working with blocking operations.

To limit the wait time for receiving or sending data on a channel.
To avoid goroutine leaks when the operation might block indefinitely
*/
func Timeouts() {

	ch := make(chan int)

	select {
	case val := <-ch:
		fmt.Println("Received:", val)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout! No value received.")
	}
}
