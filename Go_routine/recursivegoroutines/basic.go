package recursivegoroutines

import (
	"fmt"
	"sync"
)

// basic nested function...
func Recursion(n int) {
	if n == 1 {
		return
	}
	fmt.Println(n)
	Recursion(n - 1)
}

// nested go routines...
func SampleGoRoutine(n int, wg *sync.WaitGroup) {
	if n == 1 {
		wg.Done() // Mark the last task as done
		return
	}
	fmt.Println("go routine..", n)

	// Add first, then start goroutine
	wg.Add(1)
	go SampleGoRoutine(n-1, wg)

	wg.Done()
}

func Goroutine_nested() {
	var wg sync.WaitGroup
	wg.Add(1) // Initial goroutine
	go SampleGoRoutine(25, &wg)
	wg.Wait()
}
