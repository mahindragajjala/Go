package directions

/*
Restricting a channel’s direction when passing to a function
Helps in API design to prevent misuse
*/
func Producer(ch chan<- int) { //sender

}
func Consumer(ch <-chan int) { //receiver

}

/*

Inbound channel: the channel a goroutine reads from (i.e. receives data).

Outbound channel: the channel a goroutine writes to (i.e. sends data).

These aren't special channel types — they’re just a way to describe direction of data flow in a pipeline stage.
*/
func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in { // <-- in is the inbound channel
			out <- n * n // <-- out is the outbound channel
		}
		close(out)
	}()
	return out
}
