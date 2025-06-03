package pipeline

import (
	"fmt"
	"sync"
)

func Routerss(channel chan int, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	channel <- i
}
func Route(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range channel {
		fmt.Println(data)
	}
}
func RoutersToRouter() {
	channel := make(chan int)
	NoOfRouters := 25
	var wg sync.WaitGroup
	for i := 0; i < NoOfRouters; i++ {
		wg.Add(1)
		Routerss(channel, i, &wg)
	}
	go func() {
		close(channel)
		wg.Wait()
	}()
	Route(channel, &wg)
}
