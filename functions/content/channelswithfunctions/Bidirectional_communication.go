package content

import "fmt"

func process(ch chan int) {
	ch <- 10    // Send
	val := <-ch // Receive
	fmt.Println("Processed:", val)
}

func Bidirectional_communication() {
	ch := make(chan int)

	go process(ch)

	value := <-ch
	fmt.Println("Main received:", value)
	ch <- value + 5
}
