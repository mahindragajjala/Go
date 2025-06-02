package scenarios

import (
	"fmt"
	"sync"
	"time"
)

// Conntinous_sending
func Sender(wg *sync.WaitGroup, channel chan int) {
	fmt.Println("sender function")
	defer wg.Done()
	for {
		channel <- 1
		time.Sleep(1 * time.Second)
	}
}
func Receiver(wg *sync.WaitGroup, channel chan int) {
	defer wg.Done()
	for data := range channel {
		fmt.Println(data)
	}
}
func Continous_sending() {
	//buffered channel
	channel := make(chan int, 3)

	//wait group
	var wg sync.WaitGroup

	wg.Add(2)
	go Sender(&wg, channel)
	go Receiver(&wg, channel)
	wg.Wait()
	close(channel)
}
