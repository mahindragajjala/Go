package differencebetweenmutexsemaphore

import (
	"fmt"
	"sync"
	"time"
)

func Semaphore_mutex() {
	var mu sync.Mutex
	counter := 0

	worker := func(id int, wg *sync.WaitGroup) {
		defer wg.Done()
		mu.Lock()         // Acquire lock - only one goroutine at a time
		defer mu.Unlock() // Release lock

		counter++
		fmt.Printf("Worker %d incremented counter to %d\n", id, counter)
		time.Sleep(500 * time.Millisecond)
	}

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}
