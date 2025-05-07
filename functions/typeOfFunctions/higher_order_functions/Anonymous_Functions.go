package higherorderfunctions

import "fmt"

func Anonymous_functions() {
	square := func(x int) int {
		return x * x
	}

	fmt.Println(square(6)) // Output: 36
}
