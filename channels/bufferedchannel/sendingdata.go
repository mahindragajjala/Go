package bufferedchannel

import "fmt"

/*
buffered channels allow you to send data into the channel without an immediate corresponding receiver, up to the buffer capacity. Sending to a buffered channel will block only when the buffer is full.
*/
func Basic_Sending() {
	ch := make(chan int, 3) // buffered channel with capacity 3

	ch <- 10 // send 10 to channel
	ch <- 20
	ch <- 30

}

// Send in a Goroutine
func Sending_by_a_goroutine() {
	ch := make(chan int, 2)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3 // this will block until a receiver consumes
	}()

	// Receiving part (to unblock the sender eventually)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

//Send with Select and Default Case (Non-blocking Send)
func Using_Select_keyword() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	select {
	case ch <- 3:
		fmt.Println("Sent 3 to channel")
	default:
		fmt.Println("Channel buffer full, not sending")
	}
}

//Sending Multiple Values in a Loop
func Using_loop() {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}

}
