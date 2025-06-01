package unbuffered

import (
	"fmt"
	"time"
)

/*
Synchronous communication:
Demonstrate that unbuffered channels block the sender until the receiver is ready.
Show this by sending a value in a goroutine and printing timestamps before send and after receive in main.
*/
func Synchronous() {
	now := time.Now()
	Unbuffered := make(chan int)
	go func() {
		Unbuffered <- 42
	}()
	time.Sleep(2 * time.Minute)
	fmt.Println(time.Since(now))
	fmt.Println(<-Unbuffered)
	close(Unbuffered)
}
