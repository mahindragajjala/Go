package deadlock

import "fmt"

/*
ch := make(chan int)
ch <- 10 // blocks forever - no receiver
*/
func Unbuffered_Channel_with_No_Sender() {
	ch := make(chan int)
	go func() {
		fmt.Println(<-ch) // receives in background
	}()
	ch <- 10
}
