package looponchannels

import (
	"fmt"
	"math/rand"
	"time"
)

func Done() {
	dataCh := make(chan int)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case dataCh <- rand.Intn(100):
				time.Sleep(500 * time.Millisecond)
			case <-done:
				fmt.Println("Stopped producing")
				close(dataCh)
				return
			}
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		close(done) // trigger cancellation
	}()

	for val := range dataCh {
		fmt.Println("Received:", val)
	}

}

/*
Use case: Graceful shutdown with control over cancellation.
*/
