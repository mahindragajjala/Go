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

/*
A buffered channel has a capacity (e.g., 2, 5, etc.). You can send values into the channel without a receiver, until the buffer is full.
Asynchronous communication: send doesn’t block until the buffer is full.


Used when you want to allow some decoupling between sender and receiver.

ch := make(chan int, 2) // buffered channel with capacity 2

ch <- 1 // doesn't block
ch <- 2 // doesn't block
// ch <- 3 // would block because buffer is full

fmt.Println(<-ch) // receives 1
fmt.Println(<-ch) // receives 2


Buffered channels, often referred to as asynchronous channels, offer more flexibility in communication between senders and receivers.
These channels possess a buffer size, which is beneficial when the sender needs to send data in bursts, while the receiver processes it at a different pace.
The buffer stores data temporarily, allowing the sender to continue sending without immediate blocking.
However, it’s crucial to select an appropriate buffer size. Too large a buffer might consume more memory, whereas too small a buffer might not optimize performance.
To create a buffered channel of type string is: ch := make(chan string, 3) // this will create a buffered channel of size 3.
If the buffer is full or if there is nothing to receive, a buffered channel will behave very much like an unbuffered channel.

*/
func Buffered_channel() {
	buffered_channel := make(chan int, 3)
	buffered_channel <- 2
	value := <-buffered_channel
	fmt.Println(value)
}

func Unbuffered_channel() {}
