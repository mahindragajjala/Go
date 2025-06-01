package code

import "fmt"

/* type House struct {
	Address string
	Next    House // ❌ Invalid
} */
/*
This won’t compile, because:
House contains another House.
The second House would need another House inside it... infinite memory.
The compiler can’t calculate the size, so Go throws an error.
*/

type House struct {
	Address string
	Next    *House // ✅ Valid
}

//Now the Next field is a pointer — which has a fixed size (usually 8 bytes on 64-bit systems).

func Recursive_Struct() {

	house3 := &House{Address: "C-303"}
	house2 := &House{Address: "B-202", Next: house3}
	house1 := &House{Address: "A-101", Next: house2}

	current := house1
	for current != nil {
		fmt.Println("House Address:", current.Address)
		current = current.Next
	}
}

/* Whenever you're modeling something like:

Linked lists

Trees (family trees, directories)

Graphs (networks, routes)

You must use pointers or slices for self-reference to avoid compile-time errors. */
