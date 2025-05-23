package mutex

import (
	"fmt"
	"sync"
	"time"
)

/*
Wait and signal goroutines when a specific condition is met.

Wait() – blocks until signaled
Signal() – wakes one waiter
Broadcast() – wakes all waiters
*/
//🧪 Example: Producer-Consumer (Wait for Queue)
var (
	lock  sync.Mutex
	cond  = sync.NewCond(&lock)
	queue []int
)

func remove() {
	lock.Lock()
	for len(queue) == 0 {
		cond.Wait() // Block until there’s something in the queue
	}
	item := queue[0]
	queue = queue[1:]
	fmt.Println("Consumed:", item)
	lock.Unlock()
}

func add(item int) {
	lock.Lock()
	queue = append(queue, item)
	fmt.Println("Produced:", item)
	cond.Signal() // Notify a waiting consumer
	lock.Unlock()
}

func Condition() {
	go remove()
	time.Sleep(1 * time.Second)
	add(42)
	time.Sleep(1 * time.Second)
}
