package communication

import (
	"fmt"
	"sync"
	"time"
)

/*
Select with unbuffered and buffered channels:
Create an unbuffered and a buffered channel.
Use a select statement to receive from either channel whichever has data ready.
Send data to both channels in separate goroutines to observe behavior.

This means the channel can hold up to 5 values without needing a receiver to read them immediately. That is buffering.
*/
func Buffering() {
	var wg sync.WaitGroup

	buffered := make(chan int, 5)
	unbuffered := make(chan int)

	wg.Add(2)

	go Buffered(buffered, &wg)
	go Unbuffered(unbuffered, &wg)

	// Start a goroutine to wait and close channels after sending is done
	go func() {
		wg.Wait()
		close(buffered)
		close(unbuffered)
	}()

	// Receive from both channels until both are closed
	for buffered != nil || unbuffered != nil {
		select {
		case _, ok := <-buffered:
			if ok {
				fmt.Println("from buffered...")
			} else {
				buffered = nil
			}
		case _, ok := <-unbuffered:
			if ok {
				fmt.Println("from unbuffered...")
			} else {
				unbuffered = nil
			}
		}
	}
}

func Buffered(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
	}
}

func Unbuffered(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i
	}
}
