package onerouterwithmultiplerouters

import (
	"fmt"
	"sync"
	"time"
)

func Broadcasting() {
	channel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(6)
	go func() {
		defer wg.Done()
		//Receiver
		for data := range channel {
			fmt.Println(data)
		}
	}()
	go func() {
		defer wg.Done()
		//Receiver
		for data := range channel {
			fmt.Println(data)
		}
	}()
	go func() {
		defer wg.Done()
		//Receiver
		for data := range channel {
			fmt.Println(data)
		}
	}()
	go func() {
		defer wg.Done()
		//Receiver
		for data := range channel {
			fmt.Println(data)
		}
	}()
	go func() {
		defer wg.Done()
		//Receiver
		for data := range channel {
			fmt.Println(data)
		}
	}()
	go func() {
		defer wg.Done()
		//sender
		for i := 1; i <= 5; i++ {
			fmt.Printf("sending route : %v\n", i)
			channel <- i
			time.Sleep(time.Second)
		}
		close(channel)
	}()

	wg.Wait()
}
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
