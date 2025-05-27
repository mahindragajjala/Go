package closing

import (
	"fmt"
	"time"
)

var ch = make(chan string)

func Closing() {
	/* 	In Go, closing a channel is a way to indicate that no more values will be sent on that channel.

		It’s important for memory management, graceful goroutine termination, and avoiding resource leaks.

	close(ch) marks the channel as closed.
	It means no more sends (ch <- value) can happen.
	Receivers (value := <-ch) can still receive values until the channel is empty.

	After that, receives return the zero value and a false status.
	*/
	go Goroutine1()
	go Goroutine2()
	Receiver()
}
func Receiver() {
	data := <-ch
	fmt.Println(data)
}
func Goroutine1() {
	time.Sleep(2 * time.Millisecond)
	ch <- "message from the go routine 1"

}
func Goroutine2() {
	time.Sleep(2 * time.Millisecond)
	ch <- "message from the go routine 2"
	close(ch)
}

/*
Goroutine1() sends a message, then closes the channel.
Goroutine2() may try to send after the channel is already closed,

Why Is This a Problem?
You should only close a channel from one place, and only when no more sends will occur.

Closing from Goroutine1() is unsafe because you don’t know if Goroutine2() has already sent or is about to send.
*/
/*
package closing

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan string, 2) // buffered to hold 2 messages

func Closing() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Millisecond)
		ch <- "message from goroutine 1"
	}()

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Millisecond)
		ch <- "message from goroutine 2"
	}()

	go func() {
		wg.Wait()
		close(ch) // safe to close here: after both sends are done
	}()

	Receiver()
}

func Receiver() {
	for msg := range ch {
		fmt.Println(msg)
	}
}

Before close(ch):
Channel has:

Space for messages (buffer if any),
Internal queues for blocked senders/receivers.
When close(ch) is called:
Status flag is updated (closed = true),
GC does not reclaim memory yet (still referenced).
After all references to ch are gone (e.g., out of scope, goroutines done):
Go's garbage collector reclaims the memory,
All internal structures (buffer, queues) are freed.
*/
