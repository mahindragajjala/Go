package memoryallocation

import "fmt"

/*
Stack Allocation:

Happens automatically when functions are called.
Memory for local variables and function calls is allocated on the stack.
It's fast and efficient because it's managed in a "Last In, First Out" (LIFO) manner.

*/
func Stack_allocation() {
	x := 10 // Allocated on stack
	fmt.Println(x)
}
