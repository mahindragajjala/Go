package memoryallocation

/*
Happens manually or by the compiler/runtime when dynamic memory is needed (e.g., via pointers, structs, slices).
Used for longer-lived or larger data.
Memory is requested using Go’s new() or make() functions.
*/
func Heap_allocation() {
	x := new(int) // Allocated on heap
	*x = 10
}
