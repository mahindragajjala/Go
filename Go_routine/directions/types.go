package directions

import "fmt"

/*
chan T: Bidirectional (default)
chan<- T: Send-only
<-chan T: Receive-only
*/
func Bidirectional() {
	ch := make(chan int) // bidirectional channel
	go func() {
		ch <- 10 // send
	}()
	val := <-ch // receive
	fmt.Println("Received:", val)
}
func sendData(ch chan<- int) {
	ch <- 20 // can only send
}

func SendOnly() {
	ch := make(chan int)

	go sendData(ch) // pass channel as send-only

	val := <-ch // receive in main
	fmt.Println("Received:", val)
}
func receiveData(ch <-chan int) {
	val := <-ch // can only receive
	fmt.Println("Received:", val)
}
func Receive_only_Channel() {
	ch := make(chan int)

	go receiveData(ch) // pass channel as receive-only

	ch <- 30 // send in main
}
