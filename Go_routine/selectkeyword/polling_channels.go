package selectkeyword

import "fmt"

/*
Checking the status of multiple channels at the same time to see if any of them are ready (i.e., have data to read or can be written to) — without blocking.
This is done using a select statement with a default case.

*/
func Polling_channels() {
	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	default:
		fmt.Println("No data available on either channel.")
	}

}

/*
Go checks if ch1 or ch2 has data.
If one of them is ready, it executes that case.
If none are ready, it executes default immediately.
No blocking occurs — this is polling.
*/
