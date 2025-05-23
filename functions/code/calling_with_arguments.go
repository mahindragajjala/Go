package code

import "fmt"

func Greet(name string) {
	fmt.Println("Hello,", name)
}
func Calling_with_arguments() {
	Greet("Mahi")
}

/*
Stack Allocation:

When Greet("Mahi") is called, Go allocates memory on the stack for the function’s parameters (name).

	The string "Mahi" is passed by value, but since strings in Go are pointers to underlying data, only a small amount of data is passed.

Pass-by-Value:

Go always passes arguments by value.

For large structs or slices, the pointer to the data is passed (not a deep copy), but technically it's still "by value" because you're passing a copy of the pointer.

Return Path:

When the function ends, local variables like name go out of scope.

Their memory is released automatically because they are on the stack.

Garbage Collection (GC):

If a function allocates memory on the heap (e.g., creating a slice or struct), that memory is cleaned up by Go’s garbage collector if there are no more references to it.





How Strings Work in Go (Internally)
In Go, a string is a struct with two fields:

type string struct {
    DataPointer *byte  // Pointer to the actual data in memory (on the heap)
    Length      int    // Length of the string
}
So, when you write:

Greet("Mahi")
the literal "Mahi" is a string value, but under the hood it looks like:

DataPointer --> points to memory like [77 97 104 105]  // "Mahi"
Length      --> 4
*/
