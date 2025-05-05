package code

import "fmt"

/*
A channel in Go is a way to communicate between goroutines.
Channels can be used to send and receive values between goroutines.
Syntax to create a channel:
channelName := make(chan dataType)


*/
func Creating_channel() {
	Buffered_channel()
}
func Buffered_channel() {
	buffered_channel := make(chan int, 3)
	buffered_channel <- 2
	fmt.Printf("%t", buffered_channel)
}
