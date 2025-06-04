package onerouterwithmultiplerouters

import (
	"fmt"
	"sync"
	"time"
)

func MultipleChannels() []chan int {
	//receivers
	channels := make([]chan int, 5)
	for i := 0; i < 5; i++ {
		channels[i] = make(chan int)
	}
	return channels
}
func MultipleGoRoutines(channels []chan int, wg *sync.WaitGroup) {
	for i := 0; i < len(channels); i++ {
		wg.Add(1)
		go func(ch chan int, id int) {
			defer wg.Done()
			fmt.Printf("Receiver Goroutine %d started\n", id)
			for val := range ch {
				fmt.Printf("Receiver %d received: %d\n", id, val)
			}
			fmt.Printf("Receiver %d exiting\n", id)
		}(channels[i], i)
	}
}

func Sender_function() {
	senderchannel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		senderchannel <- 1
	}()
	fmt.Println(<-senderchannel)
	wg.Wait()
}
func BroadcastSender(channels []chan int, data int) {
	for i, ch := range channels {
		go func(c chan int, id int) {
			c <- data
			close(c)
		}(ch, i)
	}
}
func Connections() {
	channels := MultipleChannels()
	var wg sync.WaitGroup

	MultipleGoRoutines(channels, &wg)

	// Give goroutines time to start
	time.Sleep(1 * time.Second)

	BroadcastSender(channels, 100)

	wg.Wait()
	fmt.Println("All receivers done.")
}
