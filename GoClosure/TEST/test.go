package test

import "fmt"

// outer function returns a closure
func createCounter() func() int {
	count := 0 // variable in outer function

	// closure function - it "remembers" count
	return func() int {
		count++
		return count
	}
}

func TEst() {
	counter := createCounter() // returns the closure

	fmt.Println(counter()) // Output: 1
	fmt.Println(counter()) // Output: 2
	fmt.Println(counter()) // Output: 3
}
