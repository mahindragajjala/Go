package content

import (
	"fmt"
)

func sendData(ch chan string) {
	ch <- "Hello from function"
}

func receiveData(ch chan string) {
	msg := <-ch
	fmt.Println("Received:", msg)
}

func Sender_and_receiver() {
	ch := make(chan string)

	go sendData(ch)
	receiveData(ch)
}
