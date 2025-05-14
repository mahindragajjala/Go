package sync

import (
	"fmt"
	"sync"
	"time"
)

/*
Purpose:
Wait for multiple goroutines to complete before proceeding.

🔧 Methods:
wg.Add(n) – Adds n tasks
wg.Done() – Called by each goroutine when done
wg.Wait() – Main goroutine blocks until counter is zero
*/

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func WAITGROUP() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers completed")
}
