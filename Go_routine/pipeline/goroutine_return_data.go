package pipeline

import (
	"fmt"
	"time"
)

/*
In Go, goroutines are used to run functions concurrently. However, you cannot directly return a value from a goroutine like this:

data := go functionname() // ❌ This is invalid

The above line will throw a syntax error because go starts a goroutine but "does not wait for it to return a value", and "it cannot assign a return value directly".

Correct Way to Get Output from a Goroutine
To receive data from a goroutine, "you need to use channels".
*/
func functionname() string {
	time.Sleep(1 * time.Second) // Simulate some work
	return "Hello from goroutine"
}

func GoRoutine_return() {
	ch := make(chan string) // Create a channel

	go func() {
		result := functionname()
		ch <- result // Send the result to the channel
	}()

	data := <-ch // Receive the result from the channel
	fmt.Println(data)
}
