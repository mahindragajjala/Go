package newkeyword

import "fmt"

/*
In Go, the new keyword is used for memory allocation,
typically to create pointers to a new instance of a
variable,
struct, or other data types.

It is a built-in function that allocates memory for a given type and
returns a pointer to that memory. Here's a deeper look at how new works
and its different usages:

*/
type Person struct {
	Name string
	Age  int
}

func New_keyword() {
	// Creating a Pointer to a Value of a Type

	// Create a new integer pointer
	ptr := new(int)

	// Dereferencing the pointer to set a value
	*ptr = 10

	fmt.Println(*ptr) // Output: 10

	//Creating a Pointer to a Struct
	// Create a new Person struct pointer
	p := new(Person)

	// Set fields using the pointer
	p.Name = "John"
	p.Age = 30

	fmt.Println(*p) // Output: {John 30}

	//Comparison with & (Address-of Operator)
	/*
		Although new and & can both be used to create pointers,they differ in how they work.
		While new always allocates memory and returns a pointer to a zero-value instance of a type,
	 	& only gives the address of an existing value.
	*/
	// Using new
	ptr1 := new(int)
	fmt.Println(*ptr1) // Output: 0 (zero value of int)

	// Using & operator
	i := 42
	ptr2 := &i
	fmt.Println(*ptr2) // Output: 42
}
