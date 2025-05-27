package done

import (
	"fmt"
	"time"
)

/*
Using done channels with time.After in a select statement is a common Go pattern to handle timeouts gracefully while waiting for a goroutine to complete or a signal to arrive.

To stop waiting for an operation after a certain time (timeout), using:
done channel: indicates an operation has completed.
time.After: triggers after a specified duration to implement a timeout.
select: waits for multiple channel operations simultaneously.
*/
func Time_out() {
	done := make(chan struct{})

	// Start a goroutine to simulate some work
	go func() {
		time.Sleep(5 * time.Second) // simulate a long task
		close(done)                 // signal that the task is done
	}()

	select {
	case <-done:
		fmt.Println("Operation completed.")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout! Operation took too long.")
	}
}
