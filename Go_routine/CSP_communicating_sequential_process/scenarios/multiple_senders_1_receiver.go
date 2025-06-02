package scenarios

import (
	"fmt"
	"sync"
	"time"
)

//Multiple senders 1 receiver

// buffered
// receiver - single
func Multiple_senders_single_receiver_receiver() {
	//buffered
	buffered := make(chan int, 3)

	//waitgroup
	var wg sync.WaitGroup
	//token
	token := make(chan struct{}, 3)
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go Multiple_senders_single_receiver_sender(i, buffered, token, &wg)
	}

	// Close the channel after all senders are done
	go func() {
		wg.Wait()
		close(buffered)
	}()

	// Single receiver
	for data := range buffered {
		fmt.Println("Received data - ", data)
	}

}

// sender - multiple
func Multiple_senders_single_receiver_sender(i int, buffered chan int, token chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("Go Routine - %v\n", i)
	defer wg.Done()
	//LOCKED
	token <- struct{}{}
	//shared data
	buffered <- i
	//UNLOCKED
	time.Sleep(5 * time.Second)
	<-token
}

func Multiple_senders_single_receiver_receiver_unbuffered() {
	// unbuffered channel
	unbuffered := make(chan int)

	// token channel for mutual exclusion (acts like a semaphore)
	token := make(chan struct{}, 5)

	// wait group
	var wg sync.WaitGroup

	// start 25 sender goroutines
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go Multiple_senders_single_receiver_sender_unbuffered(i, &wg, unbuffered, token)
	}

	// receiver goroutine
	go func() {
		wg.Wait()
		close(unbuffered) // close channel after all senders finish
	}()

	// receive data from unbuffered channel
	for data := range unbuffered {
		fmt.Println("received data", data)
	}
}

// sender function
func Multiple_senders_single_receiver_sender_unbuffered(i int, wg *sync.WaitGroup, unbuffered chan int, token chan struct{}) {
	defer wg.Done()

	// simulate some logic and locking using token
	token <- struct{}{} // LOCK
	unbuffered <- i     // send data
	time.Sleep(100 * time.Millisecond)
	<-token // UNLOCK
	fmt.Printf("Go Routine - %v completed\n", i)
}
