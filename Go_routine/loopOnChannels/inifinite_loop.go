package looponchannels

import "fmt"

func Infinite_loop() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	for {
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(val)
	}
	/*
	   ✅ Use case: When you want more control over handling ok and decide what to do when the channel is closed.

	   ok is a boolean:

	   true if the value was received successfully.
	   false if the channel is closed and no more values are available.


	*/
}
