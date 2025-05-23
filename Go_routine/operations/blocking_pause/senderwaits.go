package blocking

/*
Buffered channel: Blocks when it’s full.
Unbuffered channel: Blocks if there is no receiver.
*/

//unbuffered
func Unbuffered_blocking() {
	ch := make(chan string)

	go func() {
		ch <- "Hello" // SENDER blocks if no receiver is ready
	}()

}

//buffered
func Buffered_blocking() {
	ch := make(chan string, 2)

	ch <- "One"
	ch <- "Two"   // ok
	ch <- "Three" // blocks, buffer full!
}
