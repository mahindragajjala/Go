package selectKeyword

import (
	"fmt"
	"time"
)

/*
Select - In Go, select is used to wait on multiple channel operations and proceed
		 with the one that’s ready first.

Handle multiple channels concurrently

Avoid blocking a goroutine waiting on a single channel

Implement timeouts and cancellation

Build non-blocking channel operations

*/

func Select_keyword() {
	var channel1 = make(chan (int)) // unbuffered channel
	var channel2 = make(chan (int)) // unbuffered channel
	go Goroutine1(channel1)
	go Goroutine2(channel2)
	i := 0
	for i < 2 {
		select {
		case val := <-channel1:
			fmt.Println("Received from channel1:", val) // unblocks the sender
			i++
		case val := <-channel2:
			fmt.Println("Received from channel2:", val)
			i++
		}
	}
}
func Goroutine1(s chan int) {
	time.Sleep(2 * time.Second)
	s <- 25252525 // blocked until someone receives
	fmt.Println("Working with the go routine 1 : ", s)
}

func Goroutine2(s chan int) {
	time.Sleep(1 * time.Second)
	s <- 656565 // blocked until someone receives
	fmt.Println("Working with the go routine 2 : ", s)
}
