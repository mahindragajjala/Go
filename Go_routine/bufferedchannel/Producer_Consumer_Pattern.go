package bufferedchannel

import "fmt"

func PRODUCER_CONSUMER() {
	var ch = make(chan int, 5)
	Producer(ch)
	Consumer(ch)
}

func Producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func Consumer(ch chan int) {
	for val := range ch {
		fmt.Println("Consumed:", val)
	}
}
