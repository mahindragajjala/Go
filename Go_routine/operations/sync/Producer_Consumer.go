package sync

import "fmt"

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i // Send data
	}
	close(ch)
}

func consumer(ch chan int) {
	for val := range ch {
		fmt.Println("Received:", val)
	}
}
