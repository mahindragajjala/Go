package slicewithchannels

import (
	"fmt"
	"sync"
)

/*
Create a function that accepts a slice of channels and sends data to each.
*/
func SliceOfFunction(slicedata []chan int) { //initilize that no of channels are there in the code
	for i := 0; i < len(slicedata); i++ {
		slicedata[i] = make(chan int)
	}
	var wg sync.WaitGroup
	for i := 0; i < len(slicedata); i++ {
		wg.Add(1)
		go Go_routine_SliceOfFunction(slicedata[i], i, &wg)
	}

	for i := 0; i < len(slicedata); i++ {
		fmt.Println(<-slicedata[i])
	}
	for i := 0; i < len(slicedata); i++ {
		close(slicedata[i])
	}
	wg.Wait()
}
func Go_routine_SliceOfFunction(channel chan int, in int, wg *sync.WaitGroup) {
	defer wg.Done()
	channel <- in
}
