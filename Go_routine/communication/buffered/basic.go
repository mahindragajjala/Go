package buffered

import "fmt"

/*
Basic buffered channel:
Create a buffered channel with capacity 3.
Send 3 values without a goroutine.
Receive and print all 3 values.
*/
func Buffered() {
	Buffered := make(chan int, 3)
	Buffered <- 1
	Buffered <- 2
	Buffered <- 3
	fmt.Println(<-Buffered)
	fmt.Println(<-Buffered)
	fmt.Println(<-Buffered)
}
