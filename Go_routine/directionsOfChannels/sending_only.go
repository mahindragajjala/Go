package directionsofchannels

import (
	"fmt"
	"sync"
)

var channel chan int
var wgsend sync.WaitGroup

func Sending_data_only(data chan<- int) {
	defer wgsend.Done()
	data <- 2
	fmt.Println("Data sent")
}

func Sending_direction() {
	channel = make(chan int) // ✅ initialize the channel
	wgsend.Add(1)
	go Sending_data_only(channel)

	// Read from channel to prevent goroutine from blocking
	value := <-channel
	fmt.Println("Received:", value)

	wgsend.Wait()
}

func Type_chan(ch chan int) {
	fmt.Printf("channel type : %T\n", ch)
}
