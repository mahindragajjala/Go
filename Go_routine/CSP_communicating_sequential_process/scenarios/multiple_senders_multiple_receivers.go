package scenarios

import (
	"fmt"
	"sync"
	"time"
)

func Senders(buffered chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		buffered <- i
		time.Sleep(1 * time.Second)
	}
}

func Receivers(buffered chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range buffered {
		fmt.Println(data)
	}
}

func Multiple_Sender_Multiple_Receivers() {
	buffered := make(chan int, 3)
	var wg sync.WaitGroup

	senderCount := 3
	receiverCount := 3

	// Start receivers
	wg.Add(receiverCount)
	for i := 0; i < receiverCount; i++ {
		go Receivers(buffered, &wg)
	}

	// Start senders
	var senderWg sync.WaitGroup
	senderWg.Add(senderCount)
	for i := 0; i < senderCount; i++ {
		go Senders(buffered, &senderWg)
	}

	// Close channel after all senders finish
	go func() {
		senderWg.Wait()
		close(buffered)
	}()

	// Wait for all receivers
	wg.Wait()
}
