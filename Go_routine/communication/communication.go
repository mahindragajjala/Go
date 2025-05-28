package communication

import (
	"fmt"
	"time"
)

func goroutine1(ch chan int) {
	ch <- 111
}

func goroutine2(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func Synchronization() {
	channel := make(chan int)
	go goroutine1(channel)
	go goroutine2(channel)

	time.Sleep(2 * time.Second) // Let goroutines run briefly
	close(channel)
}
