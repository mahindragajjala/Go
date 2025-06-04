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

// ABOVE CODE IS WRONG:)
//ONE SENDER - ONLY ONE RECEIVER CAN RECEIVE
//SO WE HAVE TO USE SLICE OF CHANNELS AND SEND THE DATA TO INDIVIDUAL GO ROUTINES

func Sender(senderchannel chan int, wg *sync.WaitGroup, channels []chan int) {
	defer wg.Done()

	// Send 5 integers
	for i := 0; i < 5; i++ {
		senderchannel <- i
		time.Sleep(time.Second)
	}

	close(senderchannel) // Done sending
}

func Receiver(ch chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Receiver %d got data: %d\n", id, data)
	}
}

func Challelization() {
	var wg sync.WaitGroup

	// Create 5 receiver channels
	receivechannels := make([]chan int, 5)
	for i := 0; i < 5; i++ {
		receivechannels[i] = make(chan int)
	}

	// Shared sender channel
	senderchannel := make(chan int)

	wg.Add(1)
	go Sender(senderchannel, &wg, receivechannels)

	// Launch receivers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Receiver(receivechannels[i], i, &wg)
	}

	// Read from senderchannel and forward to all receiver channels
	go func() {
		for data := range senderchannel {
			for i := 0; i < len(receivechannels); i++ {
				receivechannels[i] <- data
			}
		}

		// Close all receive channels
		for i := 0; i < len(receivechannels); i++ {
			close(receivechannels[i])
		}
	}()

	wg.Wait()
}
