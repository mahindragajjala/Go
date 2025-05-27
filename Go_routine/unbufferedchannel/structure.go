package unbufferedchannel

/*
hchan structure (runtime representation of a channel)

The Go runtime uses an internal struct called hchan for each channel.
Important fields include:

qcount: number of elements in the queue.
dataqsiz: buffer size (0 for unbuffered).
recvq: queue of goroutines waiting to receive.
sendq: queue of goroutines waiting to send.
lock: a mutex to synchronize channel operations.

Go allocates a zero-capacity hchan in memory.
dataqsiz = 0, so it can't hold even one value.
Both send and receive must be simultaneous.


When ch <- 10 Happens (No Receiver Yet)
Sender Goroutine Behavior
The goroutine encounters ch <- 10.

The runtime checks recvq.
If no receiver is waiting, it:
Creates a sudog struct (represents the blocked goroutine).
Adds it to sendq.
Marks the goroutine as waiting.
Parks (suspends) the goroutine.

// pseudocode
if recvq.empty {
    block current goroutine
    append it to sendq
}

*/
/*
Memory Behavior
The sender doesn’t store the value anywhere in the channel (because buffer = 0).
Instead, the value is temporarily held in the blocked sudog struct (in heap).
The runtime keeps a pointer to this sudog inside sendq.

When Receiver (val := <-ch) is Called

The receiver goroutine checks sendq.
It finds a sender waiting in the queue.
The value is directly transferred from the sender's stack to the receiver's stack.
Both goroutines are now unblocked and continue.
*/
