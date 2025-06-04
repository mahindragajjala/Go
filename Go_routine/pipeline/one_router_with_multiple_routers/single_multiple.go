package onerouterwithmultiplerouters

import (
	"fmt"
	"sync"
	"time"
)

func Single_Sender_Multiple_Receiver() {
	// Number of receivers
	numReceivers := 3

	// Create a slice of channels
	channels := make([]chan string, numReceivers)

	// Initialize each channel
	for i := range channels {
		channels[i] = make(chan string)
	}

	var wg sync.WaitGroup

	// Start multiple receivers
	for i := 0; i < numReceivers; i++ {
		wg.Add(1)
		go func(id int, ch <-chan string) {
			defer wg.Done()
			for msg := range ch {
				fmt.Printf("Receiver %d got: %s\n", id, msg)
			}
			fmt.Printf("Receiver %d done.\n", id)
		}(i+1, channels[i])
	}

	// One sender sends messages to all receivers
	go func() {
		messages := []string{"Hello", "How are you?", "Goodbye"}
		for _, msg := range messages {
			for _, ch := range channels {
				ch <- msg // send message to all receivers
			}
			time.Sleep(time.Second) // simulate some delay
		}

		// Close all channels
		for _, ch := range channels {
			close(ch)
		}
	}()

	wg.Wait()
	fmt.Println("All receivers finished.")
}
