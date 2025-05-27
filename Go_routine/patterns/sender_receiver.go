package patterns

import (
	"fmt"
	"sync"
	"time"
)

func SENDER_RECEIVER() {
	senderChannel := make(chan int)
	receiverChannel := make(chan int)

	messages := 5
	var wg sync.WaitGroup
	wg.Add(2)

	go Sender(senderChannel, receiverChannel, messages, &wg)
	go Receiver(senderChannel, receiverChannel, messages, &wg)

	wg.Wait() // Wait for both goroutines to finish
	fmt.Println("Done")
}

func Sender(sender chan int, receiver chan int, message int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sender started...")
	for i := 0; i < message; i++ {
		val := <-receiver
		fmt.Println("Sender received:", val)
		sender <- i
		time.Sleep(1 * time.Second)
	}
}

func Receiver(sender chan int, receiver chan int, message int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Receiver started...")
	for i := 0; i < message; i++ {
		receiver <- i // Initiates first communication
		val := <-sender
		fmt.Println("Receiver got:", val)
		time.Sleep(1 * time.Second)
	}
}
