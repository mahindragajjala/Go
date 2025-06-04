package slicewithchannels

import (
	"fmt"
	"sync"
)

/*
Create a slice of 3 unbuffered channels and send a unique integer to each. Receive and print all values.
*/
//sender
func Groutine_Multiplechannels(channel chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	channel <- i
}
func Channels() {
	//declaration and initialization of the channel
	Multiplechannels := make([]chan int, 3)

	//Indivisible initialization
	//each channel is initialization
	for i := 0; i < 3; i++ {
		Multiplechannels[i] = make(chan int)
	}

	//waitgroup
	var wg sync.WaitGroup

	//goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Groutine_Multiplechannels(Multiplechannels[i], &wg, i)
	}

	//receive
	for i := 0; i < 3; i++ {
		fmt.Println(<-Multiplechannels[i])
	}

	for i := 0; i < 3; i++ {
		close(Multiplechannels[i])
	}

	wg.Wait()
}
