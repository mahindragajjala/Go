package code

import "fmt"

/*
These are standard named functions declared using the func keyword.
*/

// Function declaration
func greet(name string) string {
	return "Hello, " + name
}

func Function_Declaration() {
	fmt.Println(greet("Mahindra"))
}
