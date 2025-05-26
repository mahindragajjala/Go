package differencebetweenmutexsemaphore

import (
	"fmt"
	"time"
)

type Semaphore struct {
	tokens chan struct{}
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{tokens: make(chan struct{}, n)}
}

func (s *Semaphore) Acquire() {
	s.tokens <- struct{}{} // Block if full (max concurrency)
}

func (s *Semaphore) Release() {
	<-s.tokens
}

func worker(id int, sem *Semaphore) {
	sem.Acquire()
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(2 * time.Second) // Simulate work
	fmt.Printf("Worker %d finished\n", id)
	sem.Release()
}

func Semaphore_with_channel() {
	sem := NewSemaphore(2) // Allow max 2 concurrent workers

	for i := 1; i <= 5; i++ {
		go worker(i, sem)
	}

	time.Sleep(11 * time.Second)
}
