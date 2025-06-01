package communication

import "fmt"

func TimeOur_on_Receive() {
	/*
			Timeout on receive:
		Create a buffered channel.
		Attempt to receive from the channel with a timeout using select.
		If no data received within 1 second, print "timeout".
	*/
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // required to exit range
	}()

	for val := range ch {
		fmt.Println(val)
	}
}

/*
Use case: When you want to receive all values from a producer until the channel is closed.

*/
