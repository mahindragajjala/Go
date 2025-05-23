package selectwait

import (
	"fmt"
	"time"
)

/*
In Go, the select statement is a powerful control structure used with channels to wait on multiple communication operations.

It allows a goroutine to wait until one of several possible channel operations is ready to proceed, making it ideal for handling concurrent communication.
*/
/*
select {
case val := <-ch1:
    // Code to handle data received from channel ch1
case ch2 <- val:
    // Code to send data to channel ch2
case <-time.After(time.Second):
    // Code to handle timeout after 1 second
default:
    // Optional: Code to execute if no channels are ready
}
*/
/*
Explanation of Key Components:
1. Each case must be a channel operation:
Receiving: val := <-ch1
Sending: ch2 <- val

2. Goroutine blocks until one case is ready:
Only one case executes – the one that is ready first.
If multiple cases are ready, one is randomly chosen.

3. default case (optional):
Runs immediately if no channels are ready.

Prevents the select from blocking
*/

func SELECT_KEYWORD() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Simulate async message sending
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello from ch1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout!")
	}
}

/*
select waits for either ch1 or ch2 to receive a value.
If neither receives within 3 seconds, the timeout triggers.

time.After(duration) returns a channel that sends the current time after the specified duration.

case <-time.After(time.Second):
    fmt.Println("Operation timed out")
This is extremely useful when:

You're waiting for a response from a slow goroutine.

You don’t want to wait forever.
*/

// Retry with Timeout
func FetchDataWithTimeout(ch chan string) {
	select {
	case data := <-ch:
		fmt.Println("Got data:", data)
	case <-time.After(2 * time.Second):
		fmt.Println("Request timed out!")
	}
}

// Use in Loop: Listen to Multiple Channels Continuously
func Continuously_LISTENGING() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	for {
		select {
		case msg := <-ch1:
			fmt.Println("From ch1:", msg)
		case msg := <-ch2:
			fmt.Println("From ch2:", msg)
		}
	}

}
