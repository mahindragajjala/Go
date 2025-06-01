package blocking

import "fmt"

func Close_blocking() {
	ch := make(chan int)
	close(ch) // safe to close; doesn't block
}

/*
Closing a channel is non-blocking.
But if receivers try to read after channel is closed and drained, they get zero value and do not block.
*/
func CLOSE_ROUTINE() {
	ch := make(chan int)

	go func() {
		for v := range ch { // receiver
			fmt.Println(v)
		}
		fmt.Println("Channel closed and drained")
	}()

	ch <- 1
	ch <- 2
	close(ch) // does not block

}

/*
If the channel is already closed and drained, reading like this:
val := <-ch
returns the zero value of the type and does not block.
*/
