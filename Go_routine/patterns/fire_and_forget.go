package patterns

import (
	"fmt"
	"time"
)

/*
"Fire and forget" is a "programming pattern" where a function or operation is started "asynchronously" (i.e., in the background), and the caller does not wait for the result or completion of that function.

In Go, this is typically implemented using goroutines.
*/
func logMessage(msg string) {
	time.Sleep(2 * time.Second) // simulate delay
	fmt.Println("Logged:", msg)
}

func Fire_and_forget() {
	fmt.Println("Start")

	// Fire and forget
	go logMessage("User signed in")

	fmt.Println("Continue working...")

	time.Sleep(3 * time.Second) // Wait to let goroutine finish (in real apps, use sync or channels)
}
