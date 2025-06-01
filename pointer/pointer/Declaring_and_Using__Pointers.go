package pointer

import "fmt"

func Declaring_and_Using_Pointers() {
	var pointer *int      // Declare a pointer to an int (initially nil)
	fmt.Println(&pointer) // Print the address of the pointer variable itself

	var data int = 5888 // Declare and initialize an integer variable
	pointer = &data     // Assign the address of data to the pointer (reference)

	fmt.Println(*pointer) // Dereference the pointer to access the value stored at that address
}
