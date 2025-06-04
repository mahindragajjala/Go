package slicewithchannels

import (
	"fmt"
	"sync"
)

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

func Slicewihtchannels() {
	noofchannels := 5
	multliplechannels := make([]chan int, noofchannels)
	fmt.Println(multliplechannels)

	//initialize each and where channel
	for i := 0; i < noofchannels; i++ {
		multliplechannels[i] = make(chan int)
	}
	var wg sync.WaitGroup
	for i := 0; i < noofchannels; i++ {
		wg.Add(1)
		go goroutine(multliplechannels[i], &wg)
	}

	for i := 0; i < noofchannels; i++ {
		fmt.Println(<-multliplechannels[i])
	}
	for i := 0; i < noofchannels; i++ {
		close(multliplechannels[i])
	}
	wg.Wait()
}
func goroutine(channel chan int, ch *sync.WaitGroup) {
	defer ch.Done()
	channel <- 1
}
