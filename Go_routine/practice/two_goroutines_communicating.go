package practice

import (
	"fmt"
	"time"
)

var channel = make(chan string)

func Communication() {
	go Goroutine_sender()
	go Goroutine_receiver()
	time.Sleep(2 * time.Millisecond)
}
func Goroutine_sender() {
	channel <- "message from sender"
	time.Sleep(2 * time.Millisecond)
}
func Goroutine_receiver() {
	msg := <-channel
	fmt.Println(msg)
	fmt.Println("received -->")
}
