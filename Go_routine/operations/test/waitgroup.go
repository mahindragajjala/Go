package test

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

/*
A WaitGroup is used when you want to wait for a group of goroutines to finish before continuing with the rest of the program.
| Function    | Description                                        |
| ----------- | -------------------------------------------------- |
| `wg.Add(n)` | Informs that there are `n` goroutines to wait for. |
| `wg.Done()` | Signals that one goroutine is finished.            |
| `wg.Wait()` | Blocks until the counter goes to 0.                |
*/
var wg sync.WaitGroup

func WaitGroup_testing() {
	wg.Add(2)
	go Goroutine1()
	go Goroutine2()
	wg.Wait()

}
func Goroutine1() {
	fmt.Println("go routine 1 is executing...")
	wg.Done()
}
func Goroutine2() {
	fmt.Println("go routine 2 is executing...")
	wg.Done()
}
