package selectkeyword

import "fmt"

func Combining_Send_and_Receive_in_One_select() {
	sendCh := make(chan string, 1)
	recvCh := make(chan string)

	sendCh <- "Prepared message"

	go func() {
		recvCh <- "Incoming message"
	}()

	select {
	case msg := <-recvCh:
		fmt.Println("Received:", msg)
	case sendCh <- "Another message":
		fmt.Println("Sent another message")
	}
}
