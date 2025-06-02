package emptystructgo

/*
What This Line Does - token := make(chan struct{})

This creates an unbuffered channel of type chan struct{}.
struct{} is an empty struct, meaning it takes 0 bytes.
The channel is used for synchronization or signaling, not data transmission.

How Signaling Works

Signaling in Go via chan struct{} relies on channel operations:
token <- struct{}{} → sends an empty signal (blocks if no receiver).
<-token → receives the signal (blocks if no sender).
This pair is like a handshake between goroutines.

go func() {
    // signal (send)
    token <- struct{}{}
}()

// wait for signal (receive)
<-token
Both goroutines will block until the send and receive meet — this is synchronization.


What Happens Internally
Channel Creation:
Go runtime allocates a small internal structure (not user-visible) that contains:

A queue of waiting senders.
A queue of waiting receivers.
Metadata: type size, buffer capacity (0 in this case), closed flag, etc.

Note: Although struct{} is zero-sized, the channel itself still has internal memory to manage synchronization logic.


Synchronization:
When a send happens:
If no receiver is waiting → sender blocks (parked).
If receiver exists → values are exchanged (zero-cost for struct{}), both goroutines continue.

When a receive happens:
If no sender is waiting → receiver blocks.
If sender exists → immediate exchange.
The Go scheduler handles waking up the blocked goroutine using M:N scheduling.
*/
