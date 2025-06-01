package code

import "fmt"

type Person_data struct {
	name string
	age  int
}

func changeName(p *Person_data) {
	fmt.Println(p)
	p.name = "John"
}

func Pointer_with_struct() {
	person1 := Person_data{name: "Alice", age: 25}
	changeName(&person1)
	fmt.Println(person1.name) // Output: John (changed!)
}
func Printing() {
	data := Person_data{name: "output", age: 52}
	fmt.Println(data)
}

/*
- A pointer to a structure is a variable that holds the memory address of a structure
rather than the structure itself.

Instead of copying the whole struct, the pointer refers to the memory address
where the struct is stored.

It allows efficient data manipulation, especially when the struct is large.
Changes made through the pointer reflect directly on the original structure.

Why Use Pointers to Structures?

Efficiency:
Passing large structs by value creates a copy, which is memory and CPU expensive.
Pointers avoid copying, only the memory address is passed around.

Mutability:
When you pass a struct by pointer, changes inside the function affect the original struct.
When passed by value, a new copy is created — changes won’t affect the original.

Shared Access:
Multiple functions or goroutines can work on the same structure using its pointer.
Useful in concurrent programming and when working with shared resources.

*/
