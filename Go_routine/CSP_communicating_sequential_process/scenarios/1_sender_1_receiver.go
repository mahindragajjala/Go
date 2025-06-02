package scenarios

import (
	"fmt"
	"time"
)

//1 SENDER - 1 RECEIVER

//BUFFERED

// RECEIVER
func One_Sender_One_Receiver_Receiver() {
	//token
	token := make(chan struct{}, 3)

	//buffered channel
	buffered := make(chan int, 3)

	go One_Sender_One_Receiver_Sender(buffered, token)

	fmt.Println(<-buffered)
	close(buffered)

}

// SENDER
func One_Sender_One_Receiver_Sender(buffered chan int, token chan struct{}) {
	//lock
	token <- struct{}{}
	//shared data
	buffered <- 2
	//unlock
	<-token
}

// UNBUFFERED
// RECEIVER
func One_Sender_One_Receiver_Receiver_receiver() {
	//token
	token := make(chan struct{}, 3)

	//unbuffered channel
	unbuffered := make(chan int)

	go One_Sender_One_Receiver_Receiver_sender(token, unbuffered)
	fmt.Println(<-unbuffered)
	close(unbuffered)
}

// SENDER
func One_Sender_One_Receiver_Receiver_sender(token chan struct{}, unbuffered chan int) {
	//lock
	token <- struct{}{}
	//shared data
	unbuffered <- 2
	time.Sleep(2 * time.Second)
	//unlock
	<-token
}
