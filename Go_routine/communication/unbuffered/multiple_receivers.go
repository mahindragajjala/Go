package unbuffered

import "fmt"

/*
Multiple receivers:
Create an unbuffered channel.
Start 3 goroutines that receive data from the channel.
Send 3 different values and ensure each goroutine receives one.


Ping-pong with unbuffered channels:
Implement a ping-pong pattern between two goroutines using an unbuffered channel.
Goroutine A sends "ping" to goroutine B, which sends back "pong" via the same channel.
Repeat 5 times.
*/
func Multiple_receivers() {
	Unbuffered := make(chan int)
	go goroutine1(Unbuffered)
	go goroutine2(Unbuffered)
	go goroutine3(Unbuffered)
	fmt.Println(<-Unbuffered)
	fmt.Println(<-Unbuffered)
	fmt.Println(<-Unbuffered)

}
func goroutine1(ch chan int) {
	ch <- 1
}
func goroutine2(ch chan int) {
	ch <- 2
}
func goroutine3(ch chan int) {
	ch <- 3
}
