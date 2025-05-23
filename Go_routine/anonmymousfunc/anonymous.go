package anonmymousfunc

import (
	"fmt"
	"sort"
	"time"
)

/*
In Go, you can use "anonymous functions as goroutines" to quickly launch concurrent tasks without the need to define a named function.

This is especially useful for short tasks or closures that capture local variables.
*/
func Anonymous_func() {
	go func() {
		fmt.Println("Printing Anonymous function")

	}()
	time.Sleep(2 * time.Millisecond)
}

//when we will use the anonymous functions in the go.

// When you need a function temporarily.
func Function_Temporarily() {
	result := func(a, b int) int {
		return a + b
	}(10, 5) // immediately invoked

	fmt.Println(result) // Output: 15
}

//When you want to pass a function as an argument (e.g., to sort.Slice, go func(), etc.)
/*

func sort.Slice(x any, less func(i int, j int) bool)
Slice sorts the slice x given the provided less function. It panics if x is not a slice.

The sort is not guaranteed to be stable: equal elements may be reversed from their original order. For a stable sort, use [SliceStable].

The less function must satisfy the same requirements as the Interface type's Less method.

Note: in many situations, the newer slices.SortFunc function is more ergonomic and runs faster.
*/
func Function_as_an_argument() {
	names := []string{"Zara", "Bob", "Alice", "Mahi"}

	// Sort using anonymous function as a comparator
	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})

	fmt.Println("Sorted names:", names)
}

// For inline callback or short logic in loops or conditions
// This avoids variable capture issues inside loops.
func Inline_callback_or_short_logic() {
	items := []int{1, 2, 3}
	for _, item := range items {
		func(n int) {
			fmt.Println(n * 2)
		}(item)
	}

}
