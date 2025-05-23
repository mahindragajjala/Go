package pointertostructure

import "fmt"

//Structure Without Pointer
type Person struct {
	name string
	age  int
}

func changeName(p Person) {
	p.name = "John"
}
func Structure_without_Pointer() {
	person1 := Person{name: "Alice", age: 25}
	changeName(person1)
	fmt.Println(person1.name) // Output: Alice (unchanged!)
}

//Pointer to Structure

func changeNamePointer(p *Person) {
	p.name = "John"
}
func PointerTostructure() {
	person1 := Person{name: "Alice", age: 25}
	changeNamePointer(&person1)
	fmt.Println(person1.name) // Output: John (changed!)
}

/*
Why Use Pointer to Struct?

Memory Efficient: Large structs aren’t copied.
Allow Modifications: Functions can change the original.
Standard Practice: Often used in Go for efficiency.
*/
