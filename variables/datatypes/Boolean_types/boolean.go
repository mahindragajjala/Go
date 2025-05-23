package booleantypes

import "fmt"

func Boolean() {
	//Conditional Statements
	var isLoggedIn bool = true

	if isLoggedIn {
		fmt.Println("Welcome back!")
	} else {
		fmt.Println("Please log in.")
	}

	//Loops
	count := 0
	for count < 5 {
		fmt.Println(count)
		count++
	}

	//Function Return Types
	fmt.Println(isEven(4)) // true
	fmt.Println(isEven(5)) // false

	//Boolean Expressions
	a := 5
	b := 10
	fmt.Println(a > b)          // false
	fmt.Println(a != b)         // true
	fmt.Println(a < b && b > 0) // true

	//Flags and Toggles
	var isDebugMode bool = false

	if isDebugMode {
		fmt.Println("Debugging enabled")
	}

	// Struct Fields
	// type User struct {
	// 	Name     string
	// 	IsActive bool
	// }

}
func isEven(n int) bool {
	return n%2 == 0
}
