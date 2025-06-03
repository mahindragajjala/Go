package pipeline

import (
	"fmt"
	"time"
)

// fanIn merges multiple input channels into a single output channel
/*
	Input: A variadic slice of receive-only channels of integers (<-chan int).
	Output: A single receive-only channel of integers (<-chan int) which merges all the input channels' values.
*/
func fanIn(channels ...<-chan int) <-chan int {
	/*
	   This is an unbuffered channel that will carry all the values received from the input channels.

	   This channel will be returned to the caller for receiving merged values.
	*/
	out := make(chan int)
	/*
	   For each input channel passed to fanIn, the code will:
	   Start a separate goroutine to read from that input channel.
	   Forward all values from that input channel into the out channel.
	*/
	for _, ch := range channels {
		/* This anonymous function takes the current input channel (ch) as parameter c.
		For each value v received from c, it sends the value into the output channel out.
		The for v := range c loop keeps reading values from the input channel until the input channel is closed.
		*/
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
		}(ch)
	}

	return out
}

/*
FAN-IN - multiple channels sends data into one channel
*/
func Fan_in_with_channel() {
	// create and initialize channels
	channel1 := make(chan int)
	channel2 := make(chan int)
	channel3 := make(chan int)
	channel4 := make(chan int)
	channel5 := make(chan int)

	// start goroutines to send values into each channel
	go func() {
		channel1 <- 1
		close(channel1)
	}()
	go func() {
		channel2 <- 2
		close(channel2)
	}()
	go func() {
		channel3 <- 3
		close(channel3)
	}()
	go func() {
		channel4 <- 4
		close(channel4)
	}()
	go func() {
		channel5 <- 5
		close(channel5)
	}()

	// fan-in all channels into one
	merged := fanIn(channel1, channel2, channel3, channel4, channel5)

	// receive from merged channel
	// Note: Since fanIn never closes the output channel,
	// we read 5 values here explicitly
	for i := 0; i < 5; i++ {
		val := <-merged
		fmt.Println(val)
	}

	// small delay to allow goroutines to finish printing
	time.Sleep(100 * time.Millisecond)
}
