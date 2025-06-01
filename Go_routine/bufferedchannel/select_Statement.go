package bufferedchannel

import (
	"fmt"
	"time"
)

var channel1 = make(chan string)
var channel2 = make(chan string)

func Select_channel() {
	go Goroutine1()
	go Goroutine2()
	time.Sleep(2 * time.Millisecond)
	select {
	case ch1 := <-channel1:
		fmt.Println(ch1)

	case ch2 := <-channel2:
		fmt.Println(ch2)

	}

}
func Goroutine1() {
	channel1 <- "message sent from go routine 1"
}
func Goroutine2() {
	channel2 <- "message sent from go routine 2"
}
