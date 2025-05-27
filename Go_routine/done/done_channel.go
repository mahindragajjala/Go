package done

import (
	"fmt"
	"time"
)

/*
In Go, a done channel is a common pattern used to "signal cancellation or completion of goroutines".

It is typically a channel of type chan struct{} because we don’t need to send any data — just a signal that something is “done”.
*/
/*
When and Why to Use a done Channel

To signal child goroutines to stop gracefully.
To notify a main/parent goroutine that a task has completed.
To avoid blocking forever and to clean up resources.
*/
/*
Basic Pattern
done := make(chan struct{})

You send the signal like:
close(done) // signal completion

You listen for it like:
<-done // block until done is closed
*/
func worker(done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Worker received stop signal")
			return
		default:
			fmt.Println("Worker is doing work...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func Multiple_worker(done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Worker received stop signal")
			return
		default:
			fmt.Println("Worker is doing work...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func Done_channel() {
	done := make(chan struct{})

	go worker(done)
	go Multiple_worker(done)

	time.Sleep(2 * time.Second)
	fmt.Println("Main sending stop signal")
	close(done) // close done channel to signal stop

	time.Sleep(1 * time.Second) // Give worker time to exit
	fmt.Println("Main exited")
}

/*
🔁 Call Flow Explanation
1. Main creates the done channel
done := make(chan struct{})

2. Main starts a child goroutine
go worker(done)

3. Child goroutine (worker) loops
It checks the done channel every loop cycle.
If done is not closed, it continues working.

select {
case <-done:
    return
default:
    // work
}
4. After 2 seconds, the main goroutine closes the done channel
close(done)

5. The worker goroutine receives from done and exits gracefully
case <-done:
    fmt.Println("Worker received stop signal")
    return

🧠 Why Use close(done) Instead of done <- struct{}{}?

close(done) broadcasts to all receivers that the channel is closed.
No need to worry about multiple receivers.
Sending done <- struct{}{} works, but only one receiver gets the message.

*/
