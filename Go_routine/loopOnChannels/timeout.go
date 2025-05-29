package looponchannels

import (
	"fmt"
	"time"
)

func Timeout() {
	ch := make(chan int)

	go func() {
		ch <- 100
		time.Sleep(2 * time.Second)
	}()

	for {
		select {
		case val := <-ch:
			fmt.Println("Received:", val)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout")
			return
		}
	}
}

/*
Use case: When you want to avoid blocking forever.
*/
