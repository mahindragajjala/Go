package codingquestions

import (
	"fmt"
	"sync"
	"time"
)

/*
Implement a safe counter that multiple goroutines can increment concurrently using Mutex.
*/
var count int
var mu sync.Mutex
var wg sync.WaitGroup

func Count_data() {
	now := time.Now()
	var n int = 100000
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {

		go Counter_func(&wg)
	}

	wg.Wait()
	fmt.Println(count)
	fmt.Println(time.Since(now))
}
func Counter_func(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	count++
	mu.Unlock()
}
