package bufferedchannel

import "fmt"

/*
In Go, a channel is used to communicate between goroutines. A buffered channel has a capacity to store values before they are received.

ch := make(chan int, 3) // buffered channel of size 3

You can send up to 3 values into this channel without blocking.
If the buffer is full, sending will block until a value is received.
If the buffer is empty, receiving will block until a value is sent.
*/

func CHANNEL() {
	ch := make(chan int, 2)

	ch <- 10
	ch <- 20
	// ch <- 30 // This would block if uncommented

	fmt.Println(<-ch) // 10
	fmt.Println(<-ch) // 20
}

/*
| Feature             | Buffered Channel              | Unbuffered Channel                 |
| ------------------- | ----------------------------- | ---------------------------------- |
| Blocking on send    | Only when buffer is full      | Always blocks                      |
| Blocking on receive | Only when buffer is empty     | Always blocks                      |
| Use case            | Temporary storage, pipelining | Synchronization between goroutines |
*/
/*
 When and Where to Use Buffered Channels

Rate Limiting – control how many goroutines can proceed.
Worker Pools – queue tasks for workers.
Message Queuing – store messages before consuming.
Decoupling – sender and receiver don't need to operate at the same pace.
*/
/*
Capacity and Length

ch := make(chan string, 5)
fmt.Println(cap(ch)) // 5
fmt.Println(len(ch)) // 0 (changes as you send/receive)
*/
/*
Range Over Buffered Channel
func main() {
    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch) // must close to avoid deadlock

    for val := range ch {
        fmt.Println(val)
    }
}

*/
