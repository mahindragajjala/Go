package closechannel

import "fmt"

func USing_select_keyword() {
	var ch = make(chan int)
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("Received:", val)
		}
	}
}

/*
✅ Behavior:
Useful when you’re listening to multiple channels.
Allows non-blocking listening on multiple input sources.
Works the same — check ok to detect closure.

✅ When to use:
In concurrent systems where you monitor many signals (e.g., dataChan, errorChan, quitChan, etc.)
*/
