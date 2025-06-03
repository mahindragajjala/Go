package slicewithchannels

import "fmt"

func Slice_with_channels() {
	NoOfChannels := 5
	channelsda := make([]chan int, NoOfChannels)

	// ✅ Properly initialize each channel
	for i := 0; i < NoOfChannels; i++ {
		channelsda[i] = make(chan int)
	}

	// Launch goroutines
	for i := 0; i < NoOfChannels; i++ {
		go Routine(i, channelsda[i])
	}

	// Receive from each channel
	for i := 0; i < NoOfChannels; i++ {
		fmt.Println(<-channelsda[i])
	}
}

func Routine(i int, ch chan int) {
	ch <- i
}
