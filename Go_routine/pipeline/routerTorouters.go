package pipeline

import (
	"fmt"
	"sync"
)

// ROUTER TO ROUTERS
func Routers(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range channel {
		fmt.Println(data)
	}
}
func RouteToRouters() {
	channel := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			Routers(channel, &wg)
		}()
	}
	for i := 0; i < 4; i++ {
		channel <- i
	}
	close(channel)
	wg.Wait()
}
