package goroutinewithloop

import "fmt"

/*
In each iteration of the loop:
You write i := i, which creates a "new local variable i "within the scope of that iteration.
The goroutine captures this new local i, not the loop-scoped i.
Since each goroutine has its own copy, they print the correct values (0, 1, 2, 3, 4 — in any order).
*/
func Capture_by_Value() {
	for i := 0; i < 5; i++ {
		i := i // create a new local copy of i
		go func() {
			fmt.Println(i)
		}()
	}
}
