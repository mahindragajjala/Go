package multiplexing

import (
	"fmt"
	"sync"
	"time"
)

/*
Many-to-one (fan-in)
*/
func Multiple_sender() {
	channel := make(chan int)
	var wg sync.WaitGroup
	token := make(chan struct{}, 2)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go sender(i, &wg, channel, token)
	}
	go func() {
		wg.Wait()
		close(channel)
	}()
	for channeldata := range channel {
		fmt.Println(channeldata)
	}

}
func sender(i int, wg *sync.WaitGroup, channel chan int, token chan struct{}) {
	defer wg.Done()
	token <- struct{}{}
	channel <- i
	time.Sleep(time.Second)
	<-token
}
