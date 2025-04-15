package referencedereference

import "fmt"

/*

Concept	Meaning
Reference (&)	Getting the memory address of a variable.
Dereference (*)	Accessing or modifying the value at the memory address (pointer).

*/

func Reference() {
	var age int = 30
	var pointerToAge *int = &age // Reference

	fmt.Println("Before:", age) // Output: 30
	*pointerToAge = 35          // Dereference: changing value at address
	fmt.Println("After:", age)  // Output: 35
}

func updateBalance(balance *float64) {
	*balance += 1000 // Dereferencing to add money
}

func Real_world() {
	var myBalance float64 = 5000
	fmt.Println("Before:", myBalance) // Output: 5000

	updateBalance(&myBalance) // Passing reference

	fmt.Println("After:", myBalance) // Output: 6000
}
