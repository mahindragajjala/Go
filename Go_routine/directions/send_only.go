package directions

import (
	"fmt"
	"sync"
	"time"
)

/*
Write a function that only receives from a channel (<-chan int) and prints the values.
Write another that only sends (chan<- int).
*/
func Sending_data() {
	var wg sync.WaitGroup
	channel := make(chan int)
	wg.Add(2)
	go Receive(&wg, channel)
	go sender(&wg, channel)
	wg.Wait()

}
func sender(wg *sync.WaitGroup, receivedata chan<- int) {
	defer wg.Done()
	receivedata <- 5
	time.Sleep(time.Second)
	close(receivedata)
}
func Receive(wg *sync.WaitGroup, receivedata <-chan int) {
	defer wg.Done()
	for data := range receivedata {
		fmt.Println("received...", data)
	}

}
