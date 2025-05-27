package rangeoverchannel

import "fmt"

func RangeOverchannels() {
	ch := make(chan string)

	// Sender goroutine
	go func() {
		ch <- "Parcel 1"
		ch <- "Parcel 2"
		ch <- "Parcel 3"
		close(ch) // Important! Needed for 'range' to stop
		/*
			Because range on a channel waits forever if the channel is open but no data is coming.
			So, always close the channel from the sender side when done.
		*/
	}()

	// Receiver
	for msg := range ch {
		fmt.Println("Received:", msg)
	}
}

/*
| Concept            | Explanation                                            |
| ------------------ | ------------------------------------------------------ |
| Only receiving     | You can only use `range` on the receiving side         |
| Only sender closes | The sender should `close()` the channel                |
| Infinite wait      | If channel is not closed, `range` will block forever   |
| Not for sending    | `range` cannot be used to send values to a channel     |

*/
