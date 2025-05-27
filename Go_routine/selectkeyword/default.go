package selectkeyword

import "fmt"

//default

/*
In Go, the select statement is used to wait on multiple channel operations.

The default case within a select makes it non-blocking, meaning the select will not wait for any channel operation to proceed.

Instead, it executes the default case immediately if no channels are ready.
*/
func Default_select() {
	ch := make(chan int, 5)
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received.")
	}

}

/*
Explanation:

If there's data available on ch, it is received.
If not, default is executed immediately—no waiting.
Use case: Prevent a goroutine from getting stuck waiting on a channel.
*/
/*
You want "non-blocking behavior".
You need to "check channel status" without getting stuck.
You’re implementing time-sensitive operations (e.g., responsiveness, user input).
You’re running background tasks and want to check if they have finished.
*/
/*
When to Avoid default in select
❌ Avoid When:

You want to wait for communication.
You expect one of the channels to be ready soon.
You risk busy-waiting (e.g., polling in a loop
*/
