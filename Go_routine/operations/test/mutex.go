package test

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var counter int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func SYNCMUTEX() {
	var wg sync.WaitGroup

	wg.Add(2)
	go increment(&wg)
	go increment(&wg)

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
