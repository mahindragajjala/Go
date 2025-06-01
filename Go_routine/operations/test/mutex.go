package test

import (
	"fmt"
	"sync"
	"time"
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

/*
A Mutex is used to prevent multiple goroutines from accessing shared data at the same time, which could cause data races or wrong results.
*/
var sum int = 0
var mutex sync.Mutex

func MUTEX_TESTING() {
	go Goroutine3()
	go Goroutine4()

	time.Sleep(2 * time.Millisecond)
	fmt.Println(sum)
}
func Goroutine3() {
	mutex.Lock()
	sum = sum + 1
	mutex.Unlock()
	mutex.Lock()
	sum = sum + 1
	mutex.Unlock()
	time.Sleep(2 * time.Millisecond)

}
func Goroutine4() {
	mutex.Lock()
	sum = sum + 1
	mutex.Unlock()
	time.Sleep(2 * time.Millisecond)

}
