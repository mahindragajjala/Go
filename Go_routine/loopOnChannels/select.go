package looponchannels

import "fmt"

/*
Loop with Select Statement
*/
func Using_selectkeyword() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 10
		ch2 <- 20
	}()

	for i := 0; i < 2; i++ {
		select {
		case val := <-ch1:
			fmt.Println("From ch1:", val)
		case val := <-ch2:
			fmt.Println("From ch2:", val)
		}
	}
}

/*
Use case: When listening to multiple channels concurrently.
*/
