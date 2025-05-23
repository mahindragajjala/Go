package selectkeyword

import (
	"fmt"
	"time"
)

func Memory_leaks() {
	go func() {
		var ch chan int
		select {
		case val := <-ch:
			fmt.Println("Received:", val)
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout in goroutine")
		}
	}()

}
