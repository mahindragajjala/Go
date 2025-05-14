package blocking

import (
	"fmt"
	"time"
)

/*
receiver blocking
*/
func ReceiverBlocks() {
	/* 	channel := make(chan int)
	   	msg := <-channel // Receiver blocks here if nothing has been sent into `channel` yet
	   	fmt.Println(msg) */
	ch := make(chan string)

	go func() {
		msg := <-ch // RECEIVER blocks until something is sent
		fmt.Println("Received:", msg)
	}()

	time.Sleep(time.Second) // give time for goroutine to start

	ch <- "Hi!" // SENDER, unblocks receiver

}

/*
Here, the receiver waits.

It blocks (pauses execution) until a value is sent into the channel by a sender goroutine.
*/
