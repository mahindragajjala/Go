package buffered

import "fmt"

/*
Buffered channel with multiple senders:
Create a buffered channel of size 5.
Launch 3 goroutines that send 2 values each into the channel.
Main goroutine receives and prints all values.
*/
func Multiple_sender() {
	Buffered := make(chan int, 5)
	go Goroutine1(Buffered)
	go Goroutine2(Buffered)
	go Goroutine3(Buffered)
	for i := 0; i < 6; i++ {
		fmt.Println(<-Buffered)
	}
}
func Goroutine1(ch chan<- int) {
	ch <- 1
	ch <- 1
}
func Goroutine2(ch chan<- int) {
	ch <- 2
	ch <- 2
}
func Goroutine3(ch chan<- int) {
	ch <- 3
	ch <- 3
}
