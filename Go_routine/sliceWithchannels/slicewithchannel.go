package slicewithchannels

import "fmt"

func SliceWithChannel() {
	// Create a slice of 3 channels
	chs := make([]chan int, 3)

	// Initialize each channel
	for i := 0; i < 3; i++ {
		chs[i] = make(chan int)
	}

	// Launch goroutines using these channels
	for i := 0; i < 3; i++ {
		go func(index int, ch chan int) {
			ch <- index * 10
		}(i, chs[i])
	}

	// Receive from each channel
	for i := 0; i < 3; i++ {
		fmt.Println(<-chs[i])
	}
}
