package pipeline

import (
	"fmt"
	"sync"
	"time"
)

/*
FAN-OUT - Multiple goroutines read from the same channel
fan-out
Create a stage that sends data to multiple worker goroutines to process in parallel.
*/
func Fan_out(channel chan int) { //sender
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("go routine : %v\n", i)
			channel <- i
		}(i)
		time.Sleep(time.Second)
	}
	// Wait for all goroutines to finish sending
	wg.Wait()
	close(channel) // Close the channel once all sends are done
}
func Fan_in(channel chan int) {
	for {
		value, ok := <-channel
		if !ok {
			// Channel is closed
			fmt.Println("Channel is closed. Stopping receiver.")
			break
		}
		fmt.Println(value)
	}
}
func Merge_Fan_in_Fan_out() {
	channel := make(chan int)
	go func() {
		Fan_out(channel)
	}()

	go func() {
		Fan_in(channel)
	}()

}
