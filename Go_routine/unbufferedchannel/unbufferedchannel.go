package unbufferedchannel

/*
An unbuffered channel has no capacity to hold values. Communication using unbuffered channels requires both a sender and a receiver to be ready at the same time (synchronous communication).

ch := make(chan int) // unbuffered channel

A send operation blocks until another goroutine receives from the channel.
A receive operation blocks until another goroutine sends into the channel.
*/
/*
In an unbuffered channel, when a goroutine tries to:

Send (ch <- 10) — it will block (i.e., pause execution) until some other goroutine is ready to receive from that same channel (<-ch).

Receive (<-ch) — it will block until some other goroutine is ready to send into the channel.

The Go runtime uses an internal struct called hchan for each channel.

Important fields include:

qcount: number of elements in the queue.
dataqsiz: buffer size (0 for unbuffered).
recvq: queue of goroutines waiting to receive.
sendq: queue of goroutines waiting to send.
lock: a mutex to synchronize channel operations.
*/
