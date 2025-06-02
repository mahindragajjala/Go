package scenarios

import (
	"fmt"
	"sync"
	"time"
)

func sender(ch chan int) {
	// Send 5 numbers to the channel
	for i := 1; i <= 5; i++ {
		fmt.Println("Sending:", i)
		ch <- i
		time.Sleep(500 * time.Millisecond) // simulate some work
	}
	close(ch) // close the channel to signal no more data
}

func receiver(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range ch {
		fmt.Printf("Receiver %d received: %d\n", id, val)
	}
	fmt.Printf("Receiver %d exiting\n", id)
}

func One_sender_multiple_receiver() {
	ch := make(chan int)
	var wg sync.WaitGroup

	// Start sender
	go sender(ch)

	// Start multiple receivers
	numReceivers := 3
	wg.Add(numReceivers)
	for i := 1; i <= numReceivers; i++ {
		go receiver(i, ch, &wg)
	}

	wg.Wait() // wait for all receivers to finish
	fmt.Println("All receivers done.")
}
