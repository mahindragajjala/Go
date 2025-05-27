package deadlock

import "fmt"

/*
ch := make(chan int)
fmt.Println(<-ch) // blocks forever - no sender
*/
func Unbuffered_channel_with_no_sender() {
	ch := make(chan int)
	go func() {
		ch <- 10
	}()
	fmt.Println(<-ch)
}
