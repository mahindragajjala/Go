package channelsdirections

import (
	"fmt"
	"sync"
	"time"
)

/*
Scenario:
You have a "sender goroutine".
The main function is the "receiver".
You're using a "buffered channel".
You want to know how and when to close the channel safely.
*/
/*
 1. Only the "sender should close the channel" - The one who sends data (writes into the channel) must also be the one to close it. |
| 2. Never send on a closed channel - Doing so will cause a runtime panic.                                                   |
| 3. Receiver must never close the channel - It doesn't make sense for a receiver to say "I'm done getting data from you."              |
| 4. Closing a buffered channel is same as unbuffered - But you must ensure all values are written before closing.                             |
| 5. After closing, no more values can be sent - But values already in the buffer can still be read.                                    |

*/
/*
Sender (Goroutine)
     |
     |  ---> ch <- value
     |  ---> close(ch) ✅
     ↓
Receiver (Main)
     |  ---> for val := range ch { ... } ✅
*/

func sender(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // sending to buffered channel
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // ✅ close only after done sending
}

func Buffered_close() {
	ch := make(chan int, 2) // Buffered channel with capacity 2

	go sender(ch) // Start sender goroutine

	// Receiver: Read till the channel is closed
	for val := range ch {
		fmt.Println("Received:", val)
	}

	fmt.Println("Channel closed. All data received.")
}

/*
✅ How Closing Works with Buffered Channels:
Even after calling close(ch), the channel still holds buffered values.
range ch will continue reading until all buffered values are consumed.
Once empty and closed, the loop ends.

✅ When Should You Close the Channel?
After the last send is done.
Typically at the end of the sender goroutine.
Do not close in main() if the sender is a goroutine.
*/

//MULTIPLE GO ROUTINES IN THE GO

/*
When working with multiple goroutines (senders) sending to a single channel, channel closing becomes tricky—because" only one sender can close the channel", and closing must happen after all senders are done.
*/
/*
Goal:
You want to close the channel safely when all goroutines are done sending.
*/
/*
| Rule                                                  | Description                                                                                                         |
| ----------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| 1. ✅ Only one sender should close the channel     | Never have multiple goroutines try to close the same channel. Only one goroutine should be responsible for closing. |
| 2. ✅ Use `sync.WaitGroup` to wait for all senders | Once all goroutines finish sending, one designated goroutine should close the channel.                              |
| 3. ❌ Do not close the channel from receiver       | The receiver just reads; it should not try to close the channel.                                                    |
| 4. ❌ Never close a channel that is still in use   | Closing while another goroutine is trying to send will cause a panic.                                               |
| 5. ✅ Use a separate "closer" goroutine if needed  | One goroutine can monitor others and close the channel when all are done.                                           |

*/

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		ch <- id*10 + i
	}
}

func Multiple_goroutines() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	numWorkers := 3
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, ch, &wg)
	}

	// Closer goroutine: waits for all senders to finish
	go func() {
		wg.Wait()
		close(ch) // ✅ close only after all senders are done
	}()

	// Receiver
	for val := range ch {
		fmt.Println("Received:", val)
	}

	fmt.Println("All data received. Channel closed.")
}
