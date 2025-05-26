package directionsofchannels

import (
	"fmt"
	"sync"
)

var BIDIRECTIONALCHANNEL = make(chan int)
var wg sync.WaitGroup

func BirectionalChannel_send(data chan<- int) {

	data <- 45
	fmt.Printf("%T\n", data)
	defer wg.Done()
}
func BirectionalChannel_receive(data <-chan int) {
	fmt.Println(<-data)
	defer wg.Done()
}
func BirectionalChannel() {
	wg.Add(2)
	go BirectionalChannel_send(BIDIRECTIONALCHANNEL)
	go BirectionalChannel_receive(BIDIRECTIONALCHANNEL)
	wg.Wait()
}
