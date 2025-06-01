package polling

import (
	"fmt"
	"time"
)

/*
In Go (Golang), polling typically refers to periodically checking for a condition or data — often used to monitor channels or other shared resources in concurrent applications.

Polling is a programming technique where you repeatedly check if something is available (like data on a channel) at regular intervals, rather than getting notified immediately (as in event-driven models).
Go has excellent concurrency support with goroutines and channels, making it easy to implement polling patterns.

A goroutine performs work and sends data to a channel.
Another goroutine or the main function polls the channel to check if data is available.

This polling might use:
select with default (non-blocking)
A for loop with time delays (time.Sleep or time.Ticker)

When to Use Polling?
When you don’t control when data will arrive and can’t block forever.
For status checks, timeouts, or heartbeat monitoring.

When integrating with non-blocking I/O or external services.
*/
func worker(ch chan int) {
	time.Sleep(2 * time.Second)
	ch <- 42
}

func POLLING() {
	ch := make(chan int)
	go worker(ch)

	for {
		select {
		case val := <-ch:
			fmt.Println("Received:", val)
			return
		default:
			fmt.Println("No data yet...")
			time.Sleep(500 * time.Millisecond) // avoid busy waiting
		}
	}
}

//USING TRICKER...

func producer(ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- 100
}

func Tricker() {
	ch := make(chan int)
	go producer(ch)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case val := <-ch:
			fmt.Println("Received:", val)
			return
		case <-ticker.C:
			fmt.Println("Polling...")
		}
	}
}
