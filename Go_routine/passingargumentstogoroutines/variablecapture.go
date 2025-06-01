package passingargumentstogoroutines

import (
	"fmt"
	"time"
)

/*
Common Pitfall: Variable Capture

When using goroutines inside loops, variables may not behave as you expect.

Go closures capture variables by reference, not by value.
*/
func Variable_Capture_wrong() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Print(i)
		}()
	}

	time.Sleep(time.Second)
}

/*
output := 44444
expected := 01234

or some other mix,

Because all goroutines share the same i from the loop, and by the time the goroutines run, the loop has already moved on.
*/
func Variable_Capture_right() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Print(i)
		}(i)
	}
	time.Sleep(time.Second)
}
