package closechannel

import "fmt"

//
func Using_break_keyword() {
	var ch = make(chan int)
	for {
		val, ok := <-ch
		if !ok {
			break // channel closed
		}
		fmt.Println("Received:", val)
	}
}

/*
✅ Behavior:
ok == false means the channel has been closed and drained.
More verbose, but useful if you need to check status per message.

✅ When to use:
If you want extra control over what happens on channel closure.
If using select (next example). */
