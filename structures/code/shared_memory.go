package code

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	Value int
	mu    sync.Mutex
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	c.Value++
	c.mu.Unlock()
}

func (c *SafeCounter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Value
}

func Shared_memory() {
	c := SafeCounter{}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final Value:", c.Get())
}
