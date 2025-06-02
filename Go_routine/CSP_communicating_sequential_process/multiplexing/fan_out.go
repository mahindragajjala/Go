package multiplexing

import (
	"fmt"
	"sync"
	"time"
)

/*
One-to-many (fan-out)
*/
func Fan_Out() {
	channel := make(chan int)
	var wg sync.WaitGroup

	// Start 25 receivers
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go receiverO(channel, &wg)
	}

	// Start sender
	go senderO(channel)

	// Wait for all receivers
	wg.Wait()
}

func senderO(channel chan int) {
	for i := 0; i < 25; i++ {
		channel <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(channel)
}

func receiverO(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range channel {
		fmt.Println(val)
	}
}
