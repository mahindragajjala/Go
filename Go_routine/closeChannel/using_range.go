package closechannel

import (
	"fmt"
	"time"
)

/*
Using range on a channel (Most Idiomatic)
*/
func Using_range() {
	var ch = make(chan int)
	for msg := range ch {
		fmt.Println(msg)
	}
}

/*
✅ Behavior:
This loop automatically ends when the channel is closed and all values have been received.
You don’t need to check anything manually.
Very clean syntax.

⚠️ Notes:
Only works when the sender closes the channel.
This is the preferred pattern in Go for reading until the channel is done.
*/

func Using_range_close_after_the_sender() {
	bufferedChannel := make(chan string, 5)

	// Start sender in a separate goroutine
	go sender(bufferedChannel)

	// Receiver reads until channel is closed
	for msg := range bufferedChannel {
		fmt.Println("Received:", msg)
	}
}

func sender(ch chan string) {
	for i := 1; i <= 10; i++ {
		message := fmt.Sprintf("Message %d", i)
		ch <- message
		time.Sleep(300 * time.Millisecond) // Simulate work
	}

	// ✅ Sender closes the channel once done
	close(ch)
}
