package directionsofchannels

import (
	"fmt"
	"sync"
)

var Achannel = make(chan int)
var wgA sync.WaitGroup

func Receive_only(data <-chan int) {
	fmt.Println(<-data)
	defer wgA.Done()
}
func Sender_() {
	Achannel <- 5555
	defer wgA.Done()
}

func Receive_only_main() {
	wgA.Add(2)
	go Sender_()
	go Receive_only(Achannel)
	wgA.Wait()
}
