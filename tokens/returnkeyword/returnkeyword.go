package returnkeyword

import "fmt"

func RetrunKeyword() {
	var a int = 10
	var b int = 2
	add(a, b)
	divide(a, b)
	getName()
	sayHi("mahindra")

}

// Basic Usage of return
func add(a int, b int) int {
	return a + b
}

// Multiple Return Values
func divide(a, b int) (int, int) {
	return a / b, a % b
}

// Named Return Values
func getName() (name string) {
	name = "Gopher"
	return // returns the named variable `name`
}

// Use with Early Exit
func sayHi(name string) {
	if name == "" {
		return // early exit if no name
	}
	fmt.Println("Hi", name)
}
