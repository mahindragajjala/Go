package readwritesimulation

import (
	"fmt"
	"time"
)

/*
Task:
Create 5 reader and 3 writer goroutines. Only 1 writer can write at a time (binary semaphore), but 2 readers can read concurrently (counting semaphore with size 2).

Goal:
Use both binary and counting semaphores to model a simplified read-write lock.
*/
func Read_write() {
	read := make(chan struct{}, 2)
	write := make(chan struct{}, 1)
	for i := 0; i <= 5; i++ {
		go Reader(read)
	}
	time.Sleep(5 * time.Second)
	for i := 0; i <= 3; i++ {
		go Writer(write)
	}
	time.Sleep(5 * time.Second)

}
func Reader(ch chan struct{}) {
	ch <- struct{}{}
	fmt.Println("Reading....")
	time.Sleep(2 * time.Second)
	<-ch
}
func Writer(ch chan struct{}) {
	ch <- struct{}{}
	fmt.Println("Wrinting....")
	time.Sleep(2 * time.Second)
	<-ch
}
