package mutex

import "fmt"

/*
✅ Purpose:
Synchronize and communicate between goroutines (no explicit locks required)

Unbuffered channel — synchronous (sender blocks until receiver)
Buffered channel — asynchronous (up to buffer size)
*/

func worker_CHANNEL(done chan bool) {
	fmt.Println("Working...")
	// Simulate work
	done <- true // Send signal
}

func CHANNEL() {
	done := make(chan bool)
	go worker_CHANNEL(done)
	<-done // Wait for signal
	fmt.Println("Done")
}
