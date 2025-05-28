package unbuffered

import (
	"fmt"
	"time"
)

/*
Create an unbuffered channel and a
sender goroutine that sends 5 integers.
Close the channel afterward.
The receiver should range over the channel and print the values until closed.
*/
func Closing() {
	Unbuffered := make(chan int)
	go CSender(Unbuffered)
	go CReceiver(Unbuffered)
	time.Sleep(2 * time.Second)
}
func CSender(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}
func CReceiver(ch chan int) {
	for v := range ch {
		fmt.Println(v)
		fmt.Println(<-ch)
	}
}
